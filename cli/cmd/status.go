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

> efa status
*/
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of EFA application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`+-----------+---------+--------+---------------+
| Node Name | Role    | Status | IP            |
+-----------+---------+--------+---------------+
| tpvm1     | active  | up     | 60.30.181.99  |
+-----------+---------+--------+---------------+
| tpvm2     | standby | up     | 60.30.181.100 |
+-----------+---------+--------+---------------+
--- Time Elapsed: 7.133318546s ---
`)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
