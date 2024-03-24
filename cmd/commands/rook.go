package command

import (
	"github.com/spf13/cobra"
)

// MonCmd represents the mons command
var RookCmd = &cobra.Command{
	Use:                "rooks",
	Short:              "Output mon endpoints",
	DisableFlagParsing: true,
	Args:               cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		println("RookCmd")
	},
}
