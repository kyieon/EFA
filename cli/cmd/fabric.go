/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fabricCmd represents the fabric command
var fabricCmd = &cobra.Command{
	Use:   "fabric",
	Short: "",
}

var fabricCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create fabric",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var fabricShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show fabric details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
Fabric Name: default, Fabric Description: Default Fabric, Fabric Stage: 3, Fabric Type: clos, Fabric Status: created
+------------+-----+-----------+-----+------+--------------+-----------+-------------------+-----------------+---------+-------+
| IP ADDRESS | POD | HOST NAME | ASN | ROLE | DEVICE STATE | APP STATE | CONFIG GEN REASON | PENDING CONFIGS | VTLB ID | LB ID |
+------------+-----+-----------+-----+------+--------------+-----------+-------------------+-----------------+---------+-------+
+------------+-----+-----------+-----+------+--------------+-----------+-------------------+-----------------+---------+-------+

Fabric Name: TB2, Fabric Description: , Fabric Stage: 3, Fabric Type: clos, Fabric Status: configure-success
+---------------+-----+---------------+-------+-------------+--------------+---------------+-------------------+----------------------------+---------+-------+
|  IP ADDRESS   | POD |   HOST NAME   |  ASN  |    ROLE     | DEVICE STATE |   APP STATE   | CONFIG GEN REASON |      PENDING CONFIGS       | VTLB ID | LB ID |
+---------------+-----+---------------+-------+-------------+--------------+---------------+-------------------+----------------------------+---------+-------+
| 60.30.181.101 |     | TB2-SLX-S1A   | 66000 | spine       | provisioned  | cfg refreshed | LA,LD             | SYSP-U,BGP-C,INTIP-C       | NA      | 1     |
| 60.30.181.102 |     | TB2-SLX-S1B   | 66000 | spine       | provisioned  | cfg refreshed | LA,LD             | SYSP-U,BGP-C,BGP-U,INTIP-C | NA      | 1     |
| 60.30.181.103 |     | TB2-SLX-Com-A | 66004 | leaf        | provisioned  | cfg in-sync   | NA                | NA                         | 2       | 1     |
| 60.30.181.104 |     | TB2-SLX-Com-B | 66004 | leaf        | provisioned  | cfg in-sync   | NA                | NA                         | 2       | 1     |
| 60.30.181.105 |     | TB2-SLX-AMF-A | 66001 | leaf        | provisioned  | cfg in-sync   | NA                | NA                         | 2       | 1     |
| 60.30.181.106 |     | TB2-SLX-AMF-B | 66001 | leaf        | provisioned  | cfg in-sync   | NA                | NA                         | 2       | 1     |
| 60.30.181.107 |     | TB2-SLX-SMF-A | 66002 | leaf        | provisioned  | cfg in-sync   | NA                | NA                         | 2       | 1     |
| 60.30.181.108 |     | TB2-SLX-SMF-B | 66002 | leaf        | provisioned  | cfg in-sync   | NA                | NA                         | 2       | 1     |
| 60.30.181.109 |     | TB2-SLX-UPF-A | 66003 | leaf        | provisioned  | cfg refreshed | LA                | SYSP-U,BGP-C,BGP-U,INTIP-C | 2       | 1     |
| 60.30.181.110 |     | TB2-SLX-UPF-B | 66003 | leaf        | provisioned  | cfg refreshed | LA,LD,POU         | SYSP-U,BGP-C,BGP-U,INTIP-C | 2       | 1     |
| 60.30.181.111 |     | TB2-SLX-DBL-A | 66011 | border-leaf | provisioned  | cfg in-sync   | NA                | NA                         | 2       | 1     |
| 60.30.181.112 |     | TB2-SLX-DBL-B | 66011 | border-leaf | provisioned  | cfg in-sync   | NA                | NA                         | 2       | 1     |
| 60.30.181.113 |     | TB2-SLX-MBL-A | 66012 | border-leaf | provisioned  | cfg in-sync   | NA                | NA                         | 2       | 1     |
| 60.30.181.114 |     | TB2-SLX-MBL-B | 66012 | border-leaf | provisioned  | cfg in-sync   | NA                | NA                         | 2       | 1     |
+---------------+-----+---------------+-------+-------------+--------------+---------------+-------------------+----------------------------+---------+-------+

CONFIG GEN REASON:
LD - Link Delete, LA - Link Add, IU - Interface Update, PLC - IPPrefixList Create, PLD - IPPrefixList Delete, PLU - IPPrefixList Update
MD/MU - MCT Delete/Update, OD - Overlay Gateway Delete, OU - Overlay Gateway Update, ED - Evpn Delete, PC - RouterPim Create, PD - RouterPim Delete
DD - Dependent Device Update, DA - Device Add, DR - Device ReAdd, ASN - Asn Update, PU - RouterPim Update, SYS - System Properties Update
MD5 - BGP MD5 Password, BGPU - Router BGP Update, BGPLL - BGP Listen Limit, POU - Port Channel Update, NA - Not Applicable

PENDING CONFIGS:
MCT - MCT Cluster, O - Overlay Gateway, SYSP - System Properties, INTIP - Interface IP, BGP - Router BGP
C/D/U - Create/Delete/Update, PA/PD - Port Add/Port Delete

For App or Device Error/Failure reason, run "efa fabric error show" for details
For config refresh reason, run "efa fabric debug config-gen-reason" for details
`)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("--- Time Elapsed: 20ms ---\n")
	},
}

func init() {
	rootCmd.AddCommand(fabricCmd)
	fabricCmd.AddCommand(fabricCreateCmd)
	fabricCmd.AddCommand(fabricShowCmd)
}
