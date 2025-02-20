package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/zzxwill/aigit/llm"
)

var Version = "dev"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "aigit",
		Short: "Generate git commit message including title and body",
		Long:  `AI Git Commi streamlines the git commit process by automatically generating meaningful and standardized commit messages.`,
	}

	var authCmd = &cobra.Command{
		Use:                   "auth",
		Short:                 "Manage LLM providers and API keys",
		Long:                  `Manage Language Model providers and their API keys. Use subcommands to list, add, or select providers.`,
		DisableFlagsInUseLine: true,
	}

	var authListCmd = &cobra.Command{
		Use:                   "list",
		Aliases:               []string{"ls"},
		Short:                 "List configured LLM providers",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			config := llm.NewConfig()
			if err := config.Load(); err != nil {
				fmt.Printf("Error reading config: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Configured providers:")
			for _, provider := range config.ListProviders() {
				if provider == config.CurrentProvider {
					fmt.Printf("* %s (current)\n", provider)
				} else {
					fmt.Printf("  %s\n", provider)
				}
			}
		},
	}

	var authAddCmd = &cobra.Command{
		Use:                   "add [provider] [api_key]",
		Short:                 "Add or update API key for a provider",
		Long:                  "Add or update API key for a provider. Supported providers: openai, gemini, doubao, deepseek",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				color.Red("Not enough arguments")
				color.Red(cmd.Long)
				color.Red("\nUsage: aigit auth add <provider> <api_key> [endpoint_id]")
				os.Exit(1)
			}

			provider := strings.ToLower(args[0])
			apiKey := args[1]

			config := llm.NewConfig()
			if err := config.Load(); err != nil {
				fmt.Printf("Error reading config: %v\n", err)
				os.Exit(1)
			}

			// Validate provider
			switch provider {
			case llm.ProviderOpenAI, llm.ProviderGemini, llm.ProviderDeepseek:
				if err := config.AddProvider(provider, apiKey); err != nil {
					fmt.Printf("Error saving config: %v\n", err)
					os.Exit(1)
				}
			case llm.ProviderDoubao:
				if len(args) < 3 {
					color.Red("Endpoint ID is required for Doubao provider")
					color.Red("Please run `aigit auth add doubao <api_key> <endpoint_id>`")
					os.Exit(1)
				}
				if err := config.AddProvider(provider, apiKey, args[2]); err != nil {
					fmt.Printf("Error saving config: %v\n", err)
					os.Exit(1)
				}
			default:
				fmt.Printf("Unsupported provider: %s\nSupported providers are: openai, gemini, doubao, deepseek\n", provider)
				os.Exit(1)
			}

			color.Green("Successfully added API key for %s", provider)
		},
	}

	var authUseCmd = &cobra.Command{
		Use:                   "use [provider]",
		Short:                 "Set the current LLM provider",
		Args:                  cobra.ExactArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			provider := strings.ToLower(args[0])

			config := llm.NewConfig()
			if err := config.Load(); err != nil {
				fmt.Printf("Error reading config: %v\n", err)
				os.Exit(1)
			}

			if err := config.UseProvider(provider); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			color.Green("Now using %s as the current provider", provider)
		},
	}

	authCmd.AddCommand(authListCmd)
	authCmd.AddCommand(authAddCmd)
	authCmd.AddCommand(authUseCmd)
	rootCmd.AddCommand(authCmd)

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
				color.Yellow("No staged changes found.")
				fmt.Print("Would you like to run 'git add .' to stage all changes? (y/N): ")
				var choice string
				fmt.Scanln(&choice)

				if strings.ToLower(choice) == "y" || strings.ToLower(choice) == "yes" {
					cmd := exec.Command("git", "add", ".")
					if err := cmd.Run(); err != nil {
						fmt.Printf("Error staging changes: %v\n", err)
						os.Exit(1)
					}
					color.Green("All changes staged successfully!")

					// Re-run git diff to get the newly staged changes
					diffOutput, err = exec.Command("git", "diff", "--cached").Output()
					if err != nil {
						fmt.Printf("Error getting git diff: %v\n", err)
						os.Exit(1)
					}
				} else {
					color.Red("No changes staged. Please use 'git add' to stage your changes.")
					os.Exit(1)
				}
			}

			config := llm.NewConfig()
			if err := config.Load(); err != nil {
				fmt.Printf("Error reading config: %v\n", err)
				os.Exit(1)
			}

			// First message generation
			fmt.Println("\n🤖 Generating commit message by", config.CurrentProvider)
			var commitMessage string
			commitMessage, err = generateMessage(config, diffOutput)
			if err != nil {
				fmt.Printf("Error generating commit message: %v\n", err)
				os.Exit(1)
			}

			for {
				// Clear some space and show the message in a box
				fmt.Println("\n📝 Generated commit message:")
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
				fmt.Println(commitMessage)
				fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

				fmt.Println("\n🤔 What would you like to do?")
				color.Blue("1. Commit this message (default)")
				color.Blue("2. Regenerate message")
				fmt.Print("\nEnter your choice (press Enter for default): ")

				var choice string
				fmt.Scanln(&choice)

				if choice == "" {
					choice = "1"
				}

				switch choice {
				case "1":
					cmd := exec.Command("git", "commit", "-m", commitMessage)
					if err := cmd.Run(); err != nil {
						fmt.Printf("Error committing changes: %v\n", err)
						os.Exit(1)
					}
					color.Green("\n✅ Successfully committed changes!")

					fmt.Print("Would you like to push these changes to the remote repository? (y/N): ")
					var pushChoice string
					fmt.Scanln(&pushChoice)

					if strings.ToLower(pushChoice) == "y" || strings.ToLower(pushChoice) == "yes" {
						cmd := exec.Command("git", "push")
						if err := cmd.Run(); err != nil {
							color.Red("Error pushing changes: %v", err)
							os.Exit(1)
						}
						color.Green("✅ Successfully pushed changes to remote repository!")
					} else {
						color.Yellow("Changes committed locally. Remember to push when ready.")
					}
					return
				case "2":
					fmt.Println("\n🤖 Regenerating commit message...")
					commitMessage, err = generateMessage(config, diffOutput)
					if err != nil {
						fmt.Printf("Error generating commit message: %v\n", err)
						os.Exit(1)
					}
					continue
				default:
					color.Red("Invalid choice. Please enter 1 or 2.")
				}
			}
		},
	}

	rootCmd.AddCommand(commitCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Add this helper function before the main function
func generateMessage(config *llm.Config, diffOutput []byte) (string, error) {
	var commitMessage string
	var err error

	apiKey := config.Providers[config.CurrentProvider].APIKey

	switch config.CurrentProvider {
	case llm.ProviderGemini:
		commitMessage, err = llm.GenerateGeminiCommitMessage(string(diffOutput), apiKey)
	case llm.ProviderDoubao:
		endpoint := config.Providers[config.CurrentProvider].Endpoint
		commitMessage, err = llm.GenerateDoubaoCommitMessage(string(diffOutput), apiKey, endpoint)
	case llm.ProviderOpenAI:
		commitMessage, err = llm.GenerateOpenAICommitMessage(string(diffOutput), apiKey)
	default:
		apiKey, err1 := base64.StdEncoding.DecodeString(llm.DefaultApiKey)
		if err1 != nil {
			return "", fmt.Errorf("error decoding API key: %v", err1)
		}
		endpoint, err1 := base64.StdEncoding.DecodeString(llm.DefaultEndpoint)
		if err1 != nil {
			return "", fmt.Errorf("error decoding endpoint: %v", err1)
		}
		commitMessage, err = llm.GenerateDoubaoCommitMessage(string(diffOutput), string(apiKey), string(endpoint))
	}
	return commitMessage, err
}
