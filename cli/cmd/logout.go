/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

/**
TEST 2022-08-30 15:55

> efa logout
*/
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout of the EFA application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--- Time Elapsed: 20ms ---")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
