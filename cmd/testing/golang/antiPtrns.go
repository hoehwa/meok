/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package golang

import (
	"github.com/hoehwa/meok/internal"
	"github.com/spf13/cobra"
)

// antiPtrnsCmd represents the antiPtrns command
var antiPtrnsCmd = &cobra.Command{
	Use:   "antiPtrns",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.PrettyPrint("/testing/golang/antiPtrns")
	},
}

func init() {
	golangCmd.AddCommand(antiPtrnsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// antiPtrnsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// antiPtrnsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
