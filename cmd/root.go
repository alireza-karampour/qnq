package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/alireza-karampour/qnq/components/login"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "qnq",
	Short: "Simple Command & Control center for all your DevOps needs",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Initialize viper configuration
		confDir, err := os.UserConfigDir()
		if err != nil {
			return err
		}
		viper.AddConfigPath(confDir)
		viper.SetConfigName(".qnqrc")
		viper.SetEnvPrefix("QNQ")
		replacer := strings.NewReplacer(".", "_")
		viper.SetEnvKeyReplacer(replacer)
		viper.AutomaticEnv()

		// Start Bubble Tea TUI with full screen mode
		p := tea.NewProgram(
			login.InitialModel(),
			tea.WithAltScreen(),       // Use alternate screen buffer
			tea.WithMouseCellMotion(), // Enable mouse support
		)
		f, err := tea.LogToFile("qnq.log", "qnq")
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := p.Run(); err != nil {
			return fmt.Errorf("error running TUI: %w", err)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
