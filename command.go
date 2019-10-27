package sensulib

import "github.com/spf13/cobra"

// Runnable is an interface command.New can take to generate a new cobra command
type Runnable interface {
	Run(*cobra.Command, []string) error
}

// NewCommand returns a now *cobra.Command suitable for self-contained configuration
func NewCommand(runnable Runnable, use, short, long string) *cobra.Command {
	return &cobra.Command{
		Use:           use,
		Short:         short,
		Long:          long,
		RunE:          runnable.Run,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
}
