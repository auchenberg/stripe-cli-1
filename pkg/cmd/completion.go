package cmd

import (
	"fmt"
	"os/exec"
	"strings"

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
			return selectShell(cc.shell)
		},
	}

	cc.cmd.Flags().StringVar(&cc.shell, "shell", "", "The shell to generate completion commands for. Supports \"bash\" or \"zsh\"")

	return cc
}

func selectShell(shell string) error {
	selected := shell
	if selected == "" {
		fmt.Println("Trying to automatically detect your shell.")

		selected = detectShell()
	}

	switch {
	case selected == "zsh":
		fmt.Println("Detected `zsh`, generating zsh completion file: stripe-completion.zsh")
		return rootCmd.GenZshCompletionFile("stripe-completion.zsh")
	case selected == "bash":
		fmt.Println("Detected `bash`, generating bash completion file: stripe-completion.bash")
		return rootCmd.GenBashCompletionFile("stripe-completion.bash")
	default:
		return fmt.Errorf("Could not automatically detect your shell. Please run the command with the `--shell` flag for either bash or zsh")
	}
}

func detectShell() string {
	cmd := exec.Command("echo", "$0")

	out, err := cmd.Output()
	if err != nil {
		return ""
	}

	switch {
	case strings.Contains(string(out), "zsh"):
		return "zsh"
	case strings.Contains(string(out), "bash"):
		return "bash"
	default:
		return ""
	}
}
