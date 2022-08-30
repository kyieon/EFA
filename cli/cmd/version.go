/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

/**
TEST 2022-08-30 15:58

> efa version
*/
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`Version : 2.5.4
Build: GA
Time Stamp: 21-11-04:00:52:09
Mode: Secure
Deployment Type: multi-node
Deployment Platform: TPVM
Virtual IP: 60.30.181.98
Node IPs: 60.30.181.99,60.30.181.100
--- Time Elapsed: 6.931585ms ---
`)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
