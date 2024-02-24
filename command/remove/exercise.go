package remove

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/scrot/musclemem-cli/cli"
	"github.com/spf13/cobra"
)

type RemoveExerciseOptions struct{}

func NewRemoveExerciseCmd(c *cli.CLIConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "exercise <workout-index>/<exercise-index>",
		Aliases: []string{"ex"},
		Short:   "Remove an exercise",
		Long:    `Remove an exercise from a workout of a user, the user must be logged-in`,
		Example: heredoc.Doc(`
      $ mm remove exercise 1/2
    `),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wi, ei, err := cli.ParseExerciseRef(args[0])
			if err != nil {
				return cli.NewCLIError(err)
			}

			if _, _, err := c.Client.Exercises.Delete(cmd.Context(), c.User, wi, ei); err != nil {
				return cli.NewAPIError(err)
			}

			return nil
		},
	}

	return cmd
}
