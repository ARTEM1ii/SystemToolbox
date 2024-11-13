/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"
	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// diskUsageCmd represents the diskUsage command
var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "Print diks usage of the current directory",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		defaultDirectory := "."
		if dir := viper.GetString("cmd.info.diskUsage.defaultDirectory"); dir != ""{
			defaultDirectory = dir
		}
		usage := du.NewDiskUsage(defaultDirectory)
		fmt.Printf("Free disk space:%d in directory %s\n", usage.Free(),defaultDirectory)
	},
}

func init() {
	InfoCmd.AddCommand(diskUsageCmd)
}
