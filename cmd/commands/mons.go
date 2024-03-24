package command

import (
	"github.com/spf13/cobra"
)

// MonCmd represents the mons command
var MonCmd = &cobra.Command{
	Use:                "mons",
	Short:              "Output mon endpoints",
	DisableFlagParsing: true,
	Args:               cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		println("Moncmd")
	},
}

// RestoreQuorum represents the mons command
var RestoreQuorum = &cobra.Command{
	Use:                "restore-quorum",
	Short:              "When quorum is lost, restore quorum to the remaining healthy mon",
	DisableFlagParsing: true,
	Args:               cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		println(args[0])
	},
}

func init() {
	MonCmd.AddCommand(RookCmd)
}
