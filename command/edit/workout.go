package edit

import (
	"strconv"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/scrot/go-musclemem"
	"github.com/scrot/musclemem-cli/cli"
	"github.com/spf13/cobra"
)

type EditWorkoutOptions struct {
	musclemem.Workout
}

func NewEditWorkoutCmd(c *cli.CLIConfig) *cobra.Command {
	opts := EditWorkoutOptions{}

	cmd := &cobra.Command{
		Use:     "workout <index>",
		Aliases: []string{"wo"},
		Short:   "Edit a workout",
		Long: `Edit a existing workout belonging to a user,
    The workout must exist and the user must be logged-in.`,
		Example: heredoc.Doc(`
      $ mm edit workout 1 --name "Full-body workout"
    `),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wi, err := strconv.Atoi(args[0])
			if err != nil {
				return cli.NewCLIError(err)
			}

			if _, _, err := c.Client.Workouts.Update(cmd.Context(), c.User, wi, opts.Workout); err != nil {
				return cli.NewAPIError(err)
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&opts.Name, "name", "", "change workout name")

	return cmd
}
