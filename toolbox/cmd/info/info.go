package info

import (

	"github.com/spf13/cobra"
)

// InfoCmd represents the info command
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "All things information",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

}
