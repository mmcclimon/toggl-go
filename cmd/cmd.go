package cmd

import (
	"time"

	"github.com/mmmcclimon/toggl-go/internal/toggl"
	"github.com/spf13/cobra"
)

type TogglCommand interface {
	Run(toggl *toggl.Client, args []string) error
	Cobra() *cobra.Command
}

func Execute() {
	cmd := setup()
	_ = cmd.Execute()
}

var allCommands = []TogglCommand{
	&AbortCommand{},
	&ConfigCommand{},
	&LoggedCommand{},
	&ProjectsCommand{},
	&ResumeCommand{},
	&ShortcutsCommand{},
	&StartCommand{},
	&StopCommand{},
	&TimelineCommand{},
	&TimerCommand{},
	&TodayCommand{},
	&WeekCommand{},
}

func setup() *cobra.Command {
	var toggl *toggl.Client

	rootCmd := &cobra.Command{
		Use:               "toggl",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) (err error) {
			toggl, err = maybeLoadConfig(cmd)
			return err
		},
	}

	// hide all the root help, it's just in the way
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.InitDefaultHelpFlag()
	_ = rootCmd.Flags().MarkHidden("help")

	for _, cmd := range allCommands {
		cmd := cmd
		cc := cmd.Cobra()
		cc.RunE = func(_ *cobra.Command, args []string) error { return cmd.Run(toggl, args) }
		rootCmd.AddCommand(cc)
	}

	return rootCmd
}

// If cmd is a child command (i.e., not the root), load up the config.
func maybeLoadConfig(cmd *cobra.Command) (*toggl.Client, error) {
	if !cmd.HasParent() {
		return nil, nil
	}

	// read config, etc.
	client := toggl.NewClient()
	err := client.ReadConfig()
	return client, err
}

// This is so goofy: time.Truncate() acts on absolute (roughly, Unix) time,
// and not on the local time, so if it's Monday at 4pm in Philadelphia,
// truncating to 24*hour will give you a time that's Sunday 7pm, rather than
// Monday at midnight, which is what I actually need.
//
// To get around this, we do a stupid hack of reparsing the date-only format
// in a local time zone.
func startOfToday() time.Time {
	now := time.Now() // always in Local zone
	format := time.DateOnly
	midnight, _ := time.ParseInLocation(format, now.Format(format), time.Local)
	return midnight
}
