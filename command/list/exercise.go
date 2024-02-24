package list

import (
	"fmt"
	"strconv"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/scrot/musclemem-cli/cli"
	"github.com/spf13/cobra"
)

type ListExerciseOptions struct{}

func ListExerciseCmd(c *cli.CLIConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "exercise <workout-index>",
		Aliases: []string{"ex"},
		Short:   "list exercises of a workout",
		Long: `lists all exercises belonging to a workout index
    only exercises can be listed that belongs to a workout
    of the logged-in user`,
		Example: heredoc.Doc(`
    $ mm list exercise 1
    `),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wi, err := strconv.Atoi(args[0])
			if err != nil {
				return cli.NewCLIError(err)
			}

			xs, _, err := c.Client.Exercises.List(cmd.Context(), c.User, wi)
			if err != nil {
				return cli.NewAPIError(err)
			}

			t := cli.NewSimpleTable(c)
			t.SetHeader([]string{"#", "NAME", "WEIGHT", "REPS"})
			for _, x := range *xs {
				t.Append([]string{
					strconv.Itoa(x.Index),
					x.Name,
					fmt.Sprintf("%.1f", x.Weight),
					fmt.Sprintf("%d", x.Repetitions),
				})
			}
			t.Render()

			return nil
		},
	}
	return cmd
}
