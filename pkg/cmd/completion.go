package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/stripe/stripe-cli/pkg/validators"
)

type completionCmd struct {
	cmd *cobra.Command

	shell string
}

func newCompletionCmd() *completionCmd {
	cc := &completionCmd{}

	cc.cmd = &cobra.Command{
		Use:   "completion",
		Short: "Generate bash and zsh completion scripts",
		Args:  validators.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cc.shell == "zsh" {
				fmt.Println("Generated zsh completion file: stripe-completion.zsh")
				return rootCmd.GenZshCompletionFile("stripe-completion.zsh")
			}
			fmt.Println("Generated bash completion file: stripe-completion.bash")
			return rootCmd.GenBashCompletionFile("stripe-completion.bash")
		},
	}

	cc.cmd.Flags().StringVar(&cc.shell, "shell", "bash", "The shell to generate completion commands for. Supports \"bash\" or \"zsh\"")

	return cc
}
