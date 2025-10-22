package cmd

import (
	"fmt"
	"text/tabwriter"

	"github.com/bc0d3/trident-recon/pkg/config"
	"github.com/bc0d3/trident-recon/pkg/executor"
	"github.com/bc0d3/trident-recon/pkg/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List active trident-recon sessions",
	Long: `List all active reconnaissance sessions.

Shows session ID, tool, command, status, and target for each session.

Examples:
  trident-recon list
  trident-recon list --tool ffuf`,
	Aliases: []string{"ls"},
	RunE:    runList,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVar(&toolFilter, "tool", "", "Filter by tool name")
}

func runList(cmd *cobra.Command, args []string) error {
	stateDir := config.GetStateDir()
	sm := executor.NewSessionManager(stateDir)

	utils.PrintInfo("Fetching active sessions...")
	sessions, err := sm.ListSessions(toolFilter)
	if err != nil {
		return fmt.Errorf("failed to list sessions: %w", err)
	}

	if len(sessions) == 0 {
		utils.PrintInfo("No active sessions found")
		return nil
	}

	fmt.Printf("\n🔱 Trident Recon - Active Sessions (%d)\n\n", len(sessions))

	// Create table writer
	w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTOOL\tCOMMAND\tSTATUS\tTARGET")
	fmt.Fprintln(w, "──\t────\t───────\t──────\t──────")

	for _, s := range sessions {
		status := s.Status
		if status == "" {
			status = "unknown"
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
			s.ID,
			s.Tool,
			truncate(s.CommandName, 30),
			status,
			truncate(s.Target, 40))
	}

	w.Flush()
	fmt.Println()

	utils.PrintInfo("Use 'tmux attach -t <session-id>' to attach to a session")
	utils.PrintInfo("Use 'trident-recon kill <id>' to kill a session")

	return nil
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
