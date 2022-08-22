/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// tenantCmd represents the tenant command
var tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "",
}

var tenantCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create tenant",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var tenantShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show tenant details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`+------+---------+------------+-------------+-------------+-----------+-----------+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| Name |  Type   | VLAN Range | L2VNI Range | L3VNI Range | VRF Count | Enable BD |                                                                                               Ports                                                                                               |
+------+---------+------------+-------------+-------------+-----------+-----------+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
|  SA  | private |   2-4090   |             |             |    10     |   false   |          60.30.181.112[0/9-26,0/32,0/35,0/27:1-4] 60.30.181.111[0/4-5,0/9-22,0/25-26,0/32,0/35,0/27:1-4] 60.30.181.113[0/17:1-4,0/18:1-4,0/32:2] 60.30.181.114[0/17:1-4,0/18:1-4,0/32:2]          |
|      |         |            |             |             |           |           |     60.30.181.108[0/1:1-4,0/10:1-4,0/11:1-4,0/12:1-4,0/13:1-4,0/14:1-4,0/15:1-4,0/16:1-4,0/17:1-4,0/18:1-4,0/19:1-4,0/2:1-4,0/20:1-4,0/3:1-4,0/4:1-4,0/5:1-4,0/6:1-4,0/7:1-4,0/8:1-4,0/9:1-4]     |
|      |         |            |             |             |           |           |     60.30.181.103[0/1:1-4,0/10:1-4,0/11:1-4,0/12:1-4,0/13:1-4,0/14:1-4,0/15:1-4,0/16:1-4,0/17:1-4,0/18:1-4,0/19:1-4,0/2:1-4,0/20:1-4,0/3:1-4,0/4:1-4,0/5:1-4,0/6:1-4,0/7:1-4,0/8:1-4,0/9:1-4]     |
|      |         |            |             |             |           |           | 60.30.181.106[0/23-26,0/1:1-4,0/10:1-4,0/11:1-4,0/12:1-4,0/13:1-4,0/14:1-4,0/15:1-4,0/16:1-4,0/17:1-4,0/18:1-4,0/19:1-4,0/2:1-4,0/20:1-4,0/3:1-4,0/4:1-4,0/5:1-4,0/6:1-4,0/7:1-4,0/8:1-4,0/9:1-4] |
|      |         |            |             |             |           |           |  60.30.181.107[0/25,0/1:1-4,0/10:1-4,0/11:1-4,0/12:1-4,0/13:1-4,0/14:1-4,0/15:1-4,0/16:1-4,0/17:1-4,0/18:1-4,0/19:1-4,0/2:1-4,0/20:1-4,0/3:1-4,0/4:1-4,0/5:1-4,0/6:1-4,0/7:1-4,0/8:1-4,0/9:1-4]   |
|      |         |            |             |             |           |           |     60.30.181.104[0/1:1-4,0/10:1-4,0/11:1-4,0/12:1-4,0/13:1-4,0/14:1-4,0/15:1-4,0/16:1-4,0/17:1-4,0/18:1-4,0/19:1-4,0/2:1-4,0/20:1-4,0/3:1-4,0/4:1-4,0/5:1-4,0/6:1-4,0/7:1-4,0/8:1-4,0/9:1-4]     |
|      |         |            |             |             |           |           |     60.30.181.109[0/1:1-4,0/10:1-4,0/11:1-4,0/12:1-4,0/13:1-4,0/14:1-4,0/15:1-4,0/16:1-4,0/17:1-4,0/18:1-4,0/19:1-4,0/2:1-4,0/20:1-4,0/3:1-4,0/4:1-4,0/5:1-4,0/6:1-4,0/7:1-4,0/8:1-4,0/9:1-4]     |
|      |         |            |             |             |           |           |     60.30.181.110[0/1:1-4,0/10:1-4,0/11:1-4,0/12:1-4,0/13:1-4,0/14:1-4,0/15:1-4,0/16:1-4,0/17:1-4,0/18:1-4,0/19:1-4,0/2:1-4,0/20:1-4,0/3:1-4,0/4:1-4,0/5:1-4,0/6:1-4,0/7:1-4,0/8:1-4,0/9:1-4]     |
|      |         |            |             |             |           |           | 60.30.181.105[0/25-26,0/1:1-4,0/10:1-4,0/11:1-4,0/12:1-4,0/13:1-4,0/14:1-4,0/15:1-4,0/16:1-4,0/17:1-4,0/18:1-4,0/19:1-4,0/2:1-4,0/20:1-4,0/3:1-4,0/4:1-4,0/5:1-4,0/6:1-4,0/7:1-4,0/8:1-4,0/9:1-4] |
+------+---------+------------+-------------+-------------+-----------+-----------+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
Tenant Details
`)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("--- Time Elapsed: 20ms ---\n")
	},
}

func init() {
	rootCmd.AddCommand(tenantCmd)
	tenantCmd.AddCommand(tenantCreateCmd)
	tenantCmd.AddCommand(tenantShowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tenantCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tenantCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
