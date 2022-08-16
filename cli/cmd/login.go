/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the EFA application",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Login successful.")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("--- Time Elapsed: 20ms ---\n")
	},
}

var username string
var password string

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.PersistentFlags().StringVar(&username, "username", "", "Name of the user to login")
	loginCmd.PersistentFlags().StringVar(&password, "password", "", "Password of the user")
	loginCmd.MarkPersistentFlagRequired("password")
}
