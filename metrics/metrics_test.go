package metrics

import (
	"testing"

	"github.com/go-test/deep"
)

func fakeTimesrc() int64 {
	return 1400000
}

func newMetrics() *Metrics {
	m := New("test")
	m.timesrc = fakeTimesrc

	return m
}

func ExampleMetrics_Log() {
	m := newMetrics()
	m.Log("value", 15)

	m2 := m.With("a", "b")
	m2.Log("othervalue", "full")

	m3 := m2.With("c", "d")
	m3.Log("different", false)
	// Output:
	// test.value 1400000 15
	// test.othervalue 1400000 full a=b
	// test.different 1400000 false a=b c=d
}

func TestMetrics_With(t *testing.T) {
	tests := []struct {
		name    string
		metrics *Metrics
		newTags []string
		want    *Metrics
	}{
		{
			"empty metrics gets tags",
			newMetrics(),
			[]string{"a", "b"},
			&Metrics{name: "test", tags: map[string]string{"a": "b"}, taglist: []string{"a=b"}},
		},
		{
			"keys migrate",
			newMetrics().With("c", "d"),
			[]string{"a", "b"},
			&Metrics{name: "test", tags: map[string]string{"a": "b", "c": "d"}, taglist: []string{"a=b", "c=d"}},
		},
		{
			"keys override",
			newMetrics().With("a", "d"),
			[]string{"a", "b"},
			&Metrics{name: "test", tags: map[string]string{"a": "b"}, taglist: []string{"a=b"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diff := deep.Equal(
				tt.want,
				tt.metrics.With(tt.newTags...),
			)

			if diff != nil {
				t.Error(diff)
			}
		})
	}
}
