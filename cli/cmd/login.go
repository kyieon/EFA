/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

/**
TEST : 2022-08-30 15:54

> efa login --password {pwd}
*/
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the EFA application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`
Login successful.
--- Time Elapsed: 163.484639ms ---
`)
	},
}

var login_username string
var login_password string

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.PersistentFlags().StringVar(&login_username, "username", "", "Name of the user to login")
	loginCmd.PersistentFlags().StringVar(&login_password, "password", "", "Password of the user")
	loginCmd.MarkPersistentFlagRequired("password")
}
