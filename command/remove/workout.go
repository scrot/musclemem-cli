package remove

import (
	"strconv"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/scrot/musclemem-cli/cli"
	"github.com/spf13/cobra"
)

type RemoveWorkoutOptions struct{}

func NewRemoveWorkoutCmd(c *cli.CLIConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "workout <workout-index>",
		Aliases: []string{"wo"},
		Short:   "Remove a workout",
		Long:    `Remove a workout, the user must be logged-in`,
		Example: heredoc.Doc(`
      $ mm remove workout 1
    `),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wi, err := strconv.Atoi(args[0])
			if err != nil {
				return cli.NewCLIError(err)
			}

			if _, _, err := c.Client.Workouts.Delete(cmd.Context(), c.User, wi); err != nil {
				return cli.NewAPIError(err)
			}

			return nil
		},
	}

	return cmd
}
