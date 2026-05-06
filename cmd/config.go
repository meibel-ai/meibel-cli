package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/meibel-ai/meibel-cli/internal/output"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage CLI configuration",
	Long:  `Commands for managing meibel CLI configuration.`,
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration interactively",
	Long:  `Set up meibel CLI with your API credentials via an interactive wizard.`,
	// Override PersistentPreRunE so config init works without existing config
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var apiKey string
		var environment string
		var customURL string
		var projectID string

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("API Key").
					Description("Your meibel API key").
					EchoMode(huh.EchoModePassword).
					Value(&apiKey).
					Validate(func(s string) error {
						if strings.TrimSpace(s) == "" {
							return fmt.Errorf("API key is required")
						}
						return nil
					}),
				huh.NewSelect[string]().
					Title("Environment").
					Description("Select the API environment").
					Options(
					huh.NewOption[string]("Meibel API (https://api.meibel.ai/v1)", "https://api.meibel.ai/v1"),
					huh.NewOption[string]("Meibel API (https://api.dev.meibel.ai/v1)", "https://api.dev.meibel.ai/v1"),
					huh.NewOption[string]("Local Development Server (http://localhost:8000)", "http://localhost:8000"),
					huh.NewOption[string]("Custom URL", "custom"),
					).
					Value(&environment),
				huh.NewInput().
					Title("Project ID").
					Description("Your project ID (e.g. prj_3bXMbz...)").
					Value(&projectID),
			),
		)

		if err := form.Run(); err != nil {
			return err
		}

		// Handle custom URL
		baseURL := environment
		if environment == "custom" {
			customForm := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Base URL").
						Description("Enter the full API base URL").
						Value(&customURL).
						Validate(func(s string) error {
							if strings.TrimSpace(s) == "" {
								return fmt.Errorf("base URL is required")
							}
							return nil
						}),
				),
			)
			if err := customForm.Run(); err != nil {
				return err
			}
			baseURL = customURL
		}

		// Create config directory
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("could not find home directory: %w", err)
		}
		configDir := filepath.Join(home, ".meibel")
		if err := os.MkdirAll(configDir, 0o755); err != nil {
			return fmt.Errorf("could not create config directory: %w", err)
		}

		// Write config
		configPath := filepath.Join(configDir, "config.yaml")
		viper.Set("api_key", apiKey)
		viper.Set("base_url", baseURL)
		if projectID != "" {
			viper.Set("project_id", projectID)
		}
		if err := viper.WriteConfigAs(configPath); err != nil {
			return fmt.Errorf("could not write config: %w", err)
		}

		output.PrintSuccess(fmt.Sprintf("Configuration saved to %s", configPath))
		return nil
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Display current configuration",
	Long:  `Show the current meibel CLI configuration with masked secrets.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := viper.GetString("api_key")
		baseURL := viper.GetString("base_url")
		configFile := viper.ConfigFileUsed()

		if configFile != "" {
			fmt.Printf("Config file: %s\n", configFile)
		} else {
			fmt.Println("Config file: (none)")
		}
		fmt.Println()

		if apiKey != "" {
			masked := apiKey[:4] + strings.Repeat("*", len(apiKey)-4)
			fmt.Printf("API Key:     %s\n", masked)
		} else {
			fmt.Println("API Key:     (not set)")
		}

		if baseURL != "" {
			fmt.Printf("Base URL:    %s\n", baseURL)
		} else {
			fmt.Println("Base URL:    (not set)")
		}

		projectID := viper.GetString("project_id")
		if projectID != "" {
			fmt.Printf("Project ID:  %s\n", projectID)
		} else {
			fmt.Println("Project ID:  (not set)")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configShowCmd)
}
