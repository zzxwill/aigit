package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "aigit",
		Short: "Generate git commit message including title and body",
		Long:  `AI Git Commi streamlines the git commit process by automatically generating meaningful and standardized commit messages.`,
	}

	var commitCmd = &cobra.Command{
		Use:   "commit",
		Short: "Generate git commit message including title and body",
		Run: func(cmd *cobra.Command, args []string) {
			// Execute git diff --cached command
			diffOutput, err := exec.Command("git", "diff", "--cached").Output()
			if err != nil {
				fmt.Printf("Error getting git diff: %v\n", err)
				os.Exit(1)
			}

			// If there are no staged changes
			if len(diffOutput) == 0 {
				fmt.Println("No staged changes found. Please stage your changes using 'git add'")
				os.Exit(1)
			}

			// TODO: Call LLM API to generate commit message
			fmt.Printf("Changes to be committed:\n%s\n", string(diffOutput))
		},
	}

	rootCmd.AddCommand(commitCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
