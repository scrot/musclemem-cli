package list

import (
	"strconv"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/scrot/musclemem-cli/cli"
	"github.com/spf13/cobra"
)

type ListWorkoutOption struct{}

func ListWorkoutCmd(c *cli.CLIConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "workout",
		Aliases: []string{"wo"},
		Short:   "list workouts of user",
		Long:    `lists all workouts belonging the logged-in user`,
		Example: heredoc.Doc(`
      $ mm list workout
    `),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			ws, _, err := c.Client.Workouts.List(cmd.Context(), c.User)
			if err != nil {
				return cli.NewAPIError(err)
			}

			t := cli.NewSimpleTable(c)
			t.SetHeader([]string{"#", "NAME"})
			for _, w := range *ws {
				t.Append([]string{strconv.Itoa(w.Index), w.Name})
			}
			t.Render()

			return nil
		},
	}

	return cmd
}
