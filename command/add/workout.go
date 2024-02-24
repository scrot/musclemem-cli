package add

import (
	"encoding/json"
	"os"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/scrot/go-musclemem"
	"github.com/scrot/musclemem-cli/cli"
	"github.com/spf13/cobra"
)

type AddWorkoutOptions struct {
	FilePath string
}

// NewAddWorkoutCmd is the cli command for adding exercises to a user workout
// it should only be used in combination with the NewAddCmd
func NewAddWorkoutCmd(c *cli.CLIConfig) *cobra.Command {
	opts := AddWorkoutOptions{}

	cmd := &cobra.Command{
		Use:     "workout",
		Aliases: []string{"wo"},
		Short:   "Add one or more workouts",
		Long:    `Add a new workout to the list of user workouts `,
		Example: heredoc.Doc(`
      # add single workout from json file
      $ mm add workout -f path/to/workout.json

      # add multiple workouts from json file
      $ mm add workout -f path/to/workouts.json
    `),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			file, err := os.Open(opts.FilePath)
			if err != nil {
				return cli.NewCLIError(err)
			}

			var w musclemem.Workout
			if err := json.NewEncoder(file).Encode(w); err != nil {
				return cli.NewCLIError(err)
			}

			if _, _, err := c.Client.Workouts.Add(cmd.Context(), c.User, w); err != nil {
				return cli.NewAPIError(err)
			}

			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&opts.FilePath, "file", "f", "", "path to json file (required)")
	cmd.MarkPersistentFlagRequired("file")

	return cmd
}
