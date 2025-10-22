package cmd

import (
	"fmt"

	"github.com/bc0d3/trident-recon/pkg/config"
	"github.com/bc0d3/trident-recon/pkg/executor"
	"github.com/bc0d3/trident-recon/pkg/utils"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill [session-id]",
	Short: "Kill a specific session",
	Long: `Kill a specific reconnaissance session by its ID.

Examples:
  trident-recon kill abc123def456`,
	Args: cobra.ExactArgs(1),
	RunE: runKill,
}

var killAllCmd = &cobra.Command{
	Use:   "kill-all",
	Short: "Kill all active sessions",
	Long: `Kill all active reconnaissance sessions.

You can optionally filter by tool using the --tool flag.

Examples:
  trident-recon kill-all
  trident-recon kill-all --tool ffuf`,
	RunE: runKillAll,
}

func init() {
	rootCmd.AddCommand(killCmd)
	rootCmd.AddCommand(killAllCmd)
	killAllCmd.Flags().StringVar(&toolFilter, "tool", "", "Filter by tool name")
}

func runKill(cmd *cobra.Command, args []string) error {
	sessionID := args[0]

	stateDir := config.GetStateDir()
	sm := executor.NewSessionManager(stateDir)

	utils.PrintInfo(fmt.Sprintf("Killing session %s...", sessionID))

	if err := sm.KillSession(sessionID); err != nil {
		return fmt.Errorf("failed to kill session: %w", err)
	}

	utils.PrintSuccess(fmt.Sprintf("Session %s killed successfully", sessionID))

	return nil
}

func runKillAll(cmd *cobra.Command, args []string) error {
	stateDir := config.GetStateDir()
	sm := executor.NewSessionManager(stateDir)

	// Get sessions to be killed
	sessions, err := sm.ListSessions(toolFilter)
	if err != nil {
		return fmt.Errorf("failed to list sessions: %w", err)
	}

	if len(sessions) == 0 {
		utils.PrintInfo("No sessions to kill")
		return nil
	}

	// Ask for confirmation
	msg := fmt.Sprintf("Kill %d session(s)?", len(sessions))
	if toolFilter != "" {
		msg = fmt.Sprintf("Kill %d %s session(s)?", len(sessions), toolFilter)
	}

	confirm, err := utils.PromptConfirm(msg)
	if err != nil {
		return err
	}

	if !confirm {
		utils.PrintInfo("Operation cancelled")
		return nil
	}

	utils.PrintInfo("Killing sessions...")

	killed, err := sm.KillAllSessions(toolFilter)
	if err != nil {
		return fmt.Errorf("failed to kill sessions: %w", err)
	}

	utils.PrintSuccess(fmt.Sprintf("Killed %d session(s)", killed))

	return nil
}
