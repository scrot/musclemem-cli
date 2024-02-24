package move

import (
	"errors"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/scrot/go-musclemem"
	"github.com/scrot/musclemem-cli/cli"
	"github.com/spf13/cobra"
)

func NewMoveExerciseCmd(c *cli.CLIConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "exercise <command>",
		Aliases: []string{"ex"},
		Short:   "Move exercise",
		Example: heredoc.Doc(`
      $ mm move exercise 1/2 up
      $ mm move exercise 1/1 down
      $ mm move exercise swap 1/1 1/2
    `),
		Args: cobra.NoArgs,
	}

	cmd.AddCommand(
		NewMoveExerciseDownCmd(c),
		NewMoveExerciseUpCmd(c),
		NewMoveExerciseSwapCmd(c),
	)

	return cmd
}

func NewMoveExerciseDownCmd(c *cli.CLIConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "down <workout-index>/<exercise-index>",
		Short: "Move exercise down",
		Long: `Move an exercise down in the list of workout exercises
    if the exercise is already the last exercise then nothing happens`,
		Example: heredoc.Doc(`
      $ mm move exercise down 1/2
    `),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wi, ei, err := cli.ParseExerciseRef(args[0])
			if err != nil {
				return cli.NewCLIError(err)
			}

			if _, err := c.Client.Exercises.Move(cmd.Context(), c.User, wi, ei, musclemem.MoveDown, nil); err != nil {
				return cli.NewAPIError(err)
			}

			return nil
		},
	}

	return cmd
}

func NewMoveExerciseUpCmd(c *cli.CLIConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up <workout-index>/<exercise-index>",
		Short: "Move exercise up",
		Long: `Move an exercise up in the list of workout exercises
    if the exercise is already the first exercise then nothing happens`,
		Example: heredoc.Doc(`
      $ mm move exercise up 1/2
    `),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			wi, ei, err := cli.ParseExerciseRef(args[0])
			if err != nil {
				return cli.NewCLIError(err)
			}

			if _, err := c.Client.Exercises.Move(cmd.Context(), c.User, wi, ei, musclemem.MoveUp, nil); err != nil {
				return cli.NewAPIError(err)
			}

			return nil
		},
	}

	return cmd
}

func NewMoveExerciseSwapCmd(c *cli.CLIConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap <workout-index>/<exercise-index> <workout-index>/<exercise-index>",
		Short: "swap two exercises",
		Long: `swap the exercise provided by the first argument 
    with the exercise from the second argument. Only exercises
    within the same workout can be swapped`,
		Example: heredoc.Doc(`
      $ mm move exercise swap 1/2 1/3
    `),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			wi, ei, err := cli.ParseExerciseRef(args[0])
			if err != nil {
				return cli.NewCLIError(err)
			}

			wi2, ei2, err := cli.ParseExerciseRef(args[1])
			if err != nil {
				return cli.NewCLIError(err)
			}

			if wi != wi2 {
				return cli.NewCLIError(errors.New("only exercises within same workout allowed"))
			}

			if _, err := c.Client.Exercises.Move(cmd.Context(), c.User, wi, ei, musclemem.MoveSwap, &ei2); err != nil {
				return cli.NewAPIError(err)
			}

			return nil
		},
	}

	return cmd
}
