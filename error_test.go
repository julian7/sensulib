package sensulib

import (
	"errors"
	"testing"
)

func TestError_Error(t *testing.T) {
	type fields struct {
		criticality int
		err         error
	}

	tests := []struct {
		name   string
		fields fields
		error  string
		exits  int
	}{
		{"OK test", fields{0, errors.New("ok")}, "OK: ok", 0},
		{"WARN test", fields{1, errors.New("warn")}, "WARNING: warn", 1},
		{"CRIT test", fields{2, errors.New("crit")}, "CRITICAL: crit", 2},
		{"UNKNOWN test", fields{3, errors.New("unknown")}, "UNKNOWN: unknown", 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			serr := &Error{
				criticality: tt.fields.criticality,
				err:         tt.fields.err,
			}
			t.Run("Error", func(t *testing.T) {
				if got := serr.Error(); got != tt.error {
					t.Errorf("Error.Error() = %v, want %v", got, tt.error)
				}
			})
			t.Run("Exit", func(t *testing.T) {
				defer CatchExit(t, "Error.Error()", tt.exits)
				serr.Exit()
			})
		})
	}
}

func TestNewError(t *testing.T) {
	tests := []struct {
		name    string
		crit    int
		fields  []interface{}
		message string
	}{
		{"empty", 1, []interface{}{}, "WARNING: (no message)"},
		{"string", 1, []interface{}{"message"}, "WARNING: message"},
		{"error", 1, []interface{}{errors.New("errormsg")}, "WARNING: errormsg"},
		{"value", 2, []interface{}{[]interface{}{"one", 2, "three", 4}}, "CRITICAL: [one 2 three 4]"},
		{"errorf", 1, []interface{}{"%s %d %s %d", "one", 2, "three", 4}, "WARNING: one 2 three 4"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ret := NewError(tt.crit, tt.fields...)
			t.Run("Error", func(t *testing.T) {
				if got := ret.Error(); got != tt.message {
					t.Errorf("NewError.Error() = %v, want %v", got, tt.message)
				}
			})
			t.Run("Exit", func(t *testing.T) {
				if tt.crit != ret.criticality {
					t.Errorf("NewError.criticality = %d, want %d", ret.criticality, tt.crit)
				}
			})
		})
	}
}
