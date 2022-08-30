/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

// fabricCmd represents the fabric command
var fabricCmd = &cobra.Command{
	Use:   "fabric",
	Short: "Fabric service commands",
}

/**
TEST : 2022-08-30 15:50

> efa fabric show
*/
var fabricShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show fabric details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`
Fabric Name: default, Fabric Description: Default Fabric, Fabric Stage: 3, Fabric Type: clos, Fabric Status: created
+------------+-----+-----------+-----+------+--------------+-----------+-------------------+-----------------+---------+-------+
| IP ADDRESS | POD | HOST NAME | ASN | ROLE | DEVICE STATE | APP STATE | CONFIG GEN REASON | PENDING CONFIGS | VTLB ID | LB ID |
+------------+-----+-----------+-----+------+--------------+-----------+-------------------+-----------------+---------+-------+
+------------+-----+-----------+-----+------+--------------+-----------+-------------------+-----------------+---------+-------+

Fabric Name: TB2, Fabric Description: , Fabric Stage: 3, Fabric Type: clos, Fabric Status: configure-success
+---------------+-----+---------------+-------+------------+--------------+-------------+-------------------+-----------------+---------+-------+
|  IP ADDRESS   | POD |   HOST NAME   |  ASN  |    ROLE    | DEVICE STATE |  APP STATE  | CONFIG GEN REASON | PENDING CONFIGS | VTLB ID | LB ID |
+---------------+-----+---------------+-------+------------+--------------+-------------+-------------------+-----------------+---------+-------+
| 60.30.181.101 |     | TB2-SLX-S1A   | 66000 | Spine      | provisioned  | cfg in-sync | NA                | NA              | NA      | 1     |
| 60.30.181.102 |     | TB2-SLX-S1B   | 66000 | Spine      | provisioned  | cfg in-sync | NA                | NA              | NA      | 1     |
| 60.30.181.103 |     | TB2-SLX-Com-A | 66004 | Leaf       | provisioned  | cfg in-sync | NA                | NA              | 2       | 1     |
| 60.30.181.104 |     | TB2-SLX-Com-B | 66004 | Leaf       | provisioned  | cfg in-sync | NA                | NA              | 2       | 1     |
| 60.30.181.105 |     | TB2-SLX-AMF-A | 66001 | Leaf       | provisioned  | cfg in-sync | NA                | NA              | 2       | 1     |
| 60.30.181.106 |     | TB2-SLX-AMF-B | 66001 | Leaf       | provisioned  | cfg in-sync | NA                | NA              | 2       | 1     |
| 60.30.181.107 |     | TB2-SLX-SMF-A | 66002 | Leaf       | provisioned  | cfg in-sync | NA                | NA              | 2       | 1     |
| 60.30.181.108 |     | TB2-SLX-SMF-B | 66002 | Leaf       | provisioned  | cfg in-sync | NA                | NA              | 2       | 1     |
| 60.30.181.109 |     | TB2-SLX-UPF-A | 66003 | Leaf       | provisioned  | cfg in-sync | NA                | NA              | 2       | 1     |
| 60.30.181.110 |     | TB2-SLX-UPF-B | 66003 | Leaf       | provisioned  | cfg in-sync | NA                | NA              | 2       | 1     |
| 60.30.181.111 |     | TB2-SLX-DBL-A | 66011 | BorderLeaf | provisioned  | cfg in-sync | NA                | NA              | 2       | 1     |
| 60.30.181.112 |     | TB2-SLX-DBL-B | 66011 | BorderLeaf | provisioned  | cfg in-sync | NA                | NA              | 2       | 1     |
| 60.30.181.113 |     | TB2-SLX-MBL-A | 66012 | BorderLeaf | provisioned  | cfg in-sync | NA                | NA              | 2       | 1     |
| 60.30.181.114 |     | TB2-SLX-MBL-B | 66012 | BorderLeaf | provisioned  | cfg in-sync | NA                | NA              | 2       | 1     |
+---------------+-----+---------------+-------+------------+--------------+-------------+-------------------+-----------------+---------+-------+

CONFIG GEN REASON:
LA/LD - Link Add/Delete, IA/ID/IU - Interface Add/Delete/Update, PLC/PLD/PLU - IPPrefixList Create/Delete/Update
MD/MU - MCT Delete/Update, OD/OU - Overlay Gateway Delete/Update, EU/ED - Evpn Delete/Update, PC/PD/PU - RouterPim Create/Delete/Update
DD - Dependent Device Update, DA/DR - Device Add/ReAdd, ASN - Asn Update, SYS - System Properties Update
MD5 - BGP MD5 Password, BGPU - Router BGP Update, BGPLL - BGP Listen Limit, POU - Port Channel Update, NA - Not Applicable

PENDING CONFIGS:
MCT - MCT Cluster, O - Overlay Gateway, SYSP - System Properties, INTIP - Interface IP, BGP - Router BGP
C/D/U - Create/Delete/Update, PA/PD - Port Add/Port Delete

For App or Device Error/Failure reason, run "efa fabric error show" for details
For config refresh reason, run "efa fabric debug config-gen-reason" for details

--- Time Elapsed: 330.212931ms ---
`)
	},
}

var fabricTopologyCmd = &cobra.Command{
	Use:   "topology",
	Short: "Show Fabric topology commands",
}

var fabricTopologyShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show Fabric topology commands",
}

/**
TEST : 2022-08-30 15:50

> efa fabric topology show physical --name default
*/
var fabricTopologyShowPhysicalCmd = &cobra.Command{
	Use:   "physical",
	Short: "fabric topology show physical",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch topologyShowPhysical_name {
		case "default":
			fmt.Print(`+----------------+------------------+---------------------+-----------------------+-----------------------+----------------------------+---------------------------+--------------------------------+
| SOURCE NODE IP | SOURCE NODE ROLE | DESTINATION NODE IP | DESTINATION NODE ROLE | SOURCE NODE INTERFACE | DESTINATION NODE INTERFACE | SOURCE DEVICE MULTI HOMED | DESTINATION DEVICE MULTI HOMED |
+----------------+------------------+---------------------+-----------------------+-----------------------+----------------------------+---------------------------+--------------------------------+
+----------------+------------------+---------------------+-----------------------+-----------------------+----------------------------+---------------------------+--------------------------------+
--- Time Elapsed: 117.938307ms ---
`)
		case "TB2":
			fmt.Print(`+----------------+------------------+---------------------+-----------------------+-----------------------+----------------------------+---------------------------+--------------------------------+
| SOURCE NODE IP | SOURCE NODE ROLE | DESTINATION NODE IP | DESTINATION NODE ROLE | SOURCE NODE INTERFACE | DESTINATION NODE INTERFACE | SOURCE DEVICE MULTI HOMED | DESTINATION DEVICE MULTI HOMED |
+----------------+------------------+---------------------+-----------------------+-----------------------+----------------------------+---------------------------+--------------------------------+
| 60.30.181.104  | Leaf             | 60.30.181.101       | Spine                 | 0/21                  | 0/2                        | true                      | false                          |
| 60.30.181.104  | Leaf             | 60.30.181.102       | Spine                 | 0/22                  | 0/2                        | true                      | false                          |
| 60.30.181.104  | Leaf             | 60.30.181.103       | Leaf                  | 0/29                  | 0/29                       | true                      | true                           |
| 60.30.181.104  | Leaf             | 60.30.181.103       | Leaf                  | 0/30                  | 0/30                       | true                      | true                           |
| 60.30.181.111  | BorderLeaf       | 60.30.181.101       | Spine                 | 0/1                   | 0/27                       | true                      | false                          |
| 60.30.181.111  | BorderLeaf       | 60.30.181.102       | Spine                 | 0/2                   | 0/27                       | true                      | false                          |
| 60.30.181.111  | BorderLeaf       | 60.30.181.112       | BorderLeaf            | 0/36                  | 0/36                       | true                      | true                           |
| 60.30.181.111  | BorderLeaf       | 60.30.181.112       | BorderLeaf            | 0/37                  | 0/37                       | true                      | true                           |
| 60.30.181.111  | BorderLeaf       | 60.30.181.112       | BorderLeaf            | 0/38                  | 0/38                       | true                      | true                           |
| 60.30.181.105  | Leaf             | 60.30.181.101       | Spine                 | 0/21                  | 0/3                        | true                      | false                          |
| 60.30.181.105  | Leaf             | 60.30.181.102       | Spine                 | 0/22                  | 0/3                        | true                      | false                          |
| 60.30.181.105  | Leaf             | 60.30.181.106       | Leaf                  | 0/25                  | 0/25                       | true                      | true                           |
| 60.30.181.105  | Leaf             | 60.30.181.106       | Leaf                  | 0/30                  | 0/30                       | true                      | true                           |
| 60.30.181.106  | Leaf             | 60.30.181.101       | Spine                 | 0/21                  | 0/4                        | true                      | false                          |
| 60.30.181.106  | Leaf             | 60.30.181.102       | Spine                 | 0/22                  | 0/4                        | true                      | false                          |
| 60.30.181.106  | Leaf             | 60.30.181.105       | Leaf                  | 0/25                  | 0/25                       | true                      | true                           |
| 60.30.181.106  | Leaf             | 60.30.181.105       | Leaf                  | 0/30                  | 0/30                       | true                      | true                           |
| 60.30.181.110  | Leaf             | 60.30.181.101       | Spine                 | 0/21                  | 0/8                        | true                      | false                          |
| 60.30.181.110  | Leaf             | 60.30.181.102       | Spine                 | 0/22                  | 0/8                        | true                      | false                          |
| 60.30.181.110  | Leaf             | 60.30.181.101       | Spine                 | 0/23                  | 0/16                       | true                      | false                          |
| 60.30.181.110  | Leaf             | 60.30.181.102       | Spine                 | 0/24                  | 0/15                       | true                      | false                          |
| 60.30.181.110  | Leaf             | 60.30.181.109       | Leaf                  | 0/30                  | 0/30                       | true                      | true                           |
| 60.30.181.110  | Leaf             | 60.30.181.109       | Leaf                  | 0/31                  | 0/31                       | true                      | true                           |
| 60.30.181.107  | Leaf             | 60.30.181.101       | Spine                 | 0/21                  | 0/5                        | true                      | false                          |
| 60.30.181.107  | Leaf             | 60.30.181.102       | Spine                 | 0/22                  | 0/5                        | true                      | false                          |
| 60.30.181.107  | Leaf             | 60.30.181.108       | Leaf                  | 0/29                  | 0/29                       | true                      | true                           |
| 60.30.181.107  | Leaf             | 60.30.181.108       | Leaf                  | 0/30                  | 0/30                       | true                      | true                           |
| 60.30.181.112  | BorderLeaf       | 60.30.181.101       | Spine                 | 0/1                   | 0/28                       | true                      | false                          |
| 60.30.181.112  | BorderLeaf       | 60.30.181.102       | Spine                 | 0/2                   | 0/28                       | true                      | false                          |
| 60.30.181.112  | BorderLeaf       | 60.30.181.111       | BorderLeaf            | 0/36                  | 0/36                       | true                      | true                           |
| 60.30.181.112  | BorderLeaf       | 60.30.181.111       | BorderLeaf            | 0/37                  | 0/37                       | true                      | true                           |
| 60.30.181.112  | BorderLeaf       | 60.30.181.111       | BorderLeaf            | 0/38                  | 0/38                       | true                      | true                           |
| 60.30.181.101  | Spine            | 60.30.181.103       | Leaf                  | 0/1                   | 0/21                       | false                     | true                           |
| 60.30.181.101  | Spine            | 60.30.181.104       | Leaf                  | 0/2                   | 0/21                       | false                     | true                           |
| 60.30.181.101  | Spine            | 60.30.181.105       | Leaf                  | 0/3                   | 0/21                       | false                     | true                           |
| 60.30.181.101  | Spine            | 60.30.181.106       | Leaf                  | 0/4                   | 0/21                       | false                     | true                           |
| 60.30.181.101  | Spine            | 60.30.181.107       | Leaf                  | 0/5                   | 0/21                       | false                     | true                           |
| 60.30.181.101  | Spine            | 60.30.181.108       | Leaf                  | 0/6                   | 0/21                       | false                     | true                           |
| 60.30.181.101  | Spine            | 60.30.181.109       | Leaf                  | 0/7                   | 0/21                       | false                     | true                           |
| 60.30.181.101  | Spine            | 60.30.181.110       | Leaf                  | 0/8                   | 0/21                       | false                     | true                           |
| 60.30.181.101  | Spine            | 60.30.181.109       | Leaf                  | 0/15                  | 0/23                       | false                     | true                           |
| 60.30.181.101  | Spine            | 60.30.181.110       | Leaf                  | 0/16                  | 0/23                       | false                     | true                           |
| 60.30.181.101  | Spine            | 60.30.181.111       | BorderLeaf            | 0/27                  | 0/1                        | false                     | true                           |
| 60.30.181.101  | Spine            | 60.30.181.112       | BorderLeaf            | 0/28                  | 0/1                        | false                     | true                           |
| 60.30.181.101  | Spine            | 60.30.181.113       | BorderLeaf            | 0/29                  | 0/1                        | false                     | true                           |
| 60.30.181.101  | Spine            | 60.30.181.114       | BorderLeaf            | 0/30                  | 0/1                        | false                     | true                           |
| 60.30.181.103  | Leaf             | 60.30.181.101       | Spine                 | 0/21                  | 0/1                        | true                      | false                          |
| 60.30.181.103  | Leaf             | 60.30.181.102       | Spine                 | 0/22                  | 0/1                        | true                      | false                          |
| 60.30.181.103  | Leaf             | 60.30.181.104       | Leaf                  | 0/29                  | 0/29                       | true                      | true                           |
| 60.30.181.103  | Leaf             | 60.30.181.104       | Leaf                  | 0/30                  | 0/30                       | true                      | true                           |
| 60.30.181.114  | BorderLeaf       | 60.30.181.101       | Spine                 | 0/1                   | 0/30                       | true                      | false                          |
| 60.30.181.114  | BorderLeaf       | 60.30.181.102       | Spine                 | 0/2                   | 0/30                       | true                      | false                          |
| 60.30.181.114  | BorderLeaf       | 60.30.181.113       | BorderLeaf            | 0/29                  | 0/29                       | true                      | true                           |
| 60.30.181.114  | BorderLeaf       | 60.30.181.113       | BorderLeaf            | 0/30                  | 0/30                       | true                      | true                           |
| 60.30.181.113  | BorderLeaf       | 60.30.181.101       | Spine                 | 0/1                   | 0/29                       | true                      | false                          |
| 60.30.181.113  | BorderLeaf       | 60.30.181.102       | Spine                 | 0/2                   | 0/29                       | true                      | false                          |
| 60.30.181.113  | BorderLeaf       | 60.30.181.114       | BorderLeaf            | 0/29                  | 0/29                       | true                      | true                           |
| 60.30.181.113  | BorderLeaf       | 60.30.181.114       | BorderLeaf            | 0/30                  | 0/30                       | true                      | true                           |
| 60.30.181.109  | Leaf             | 60.30.181.101       | Spine                 | 0/21                  | 0/7                        | true                      | false                          |
| 60.30.181.109  | Leaf             | 60.30.181.102       | Spine                 | 0/22                  | 0/7                        | true                      | false                          |
| 60.30.181.109  | Leaf             | 60.30.181.101       | Spine                 | 0/23                  | 0/15                       | true                      | false                          |
| 60.30.181.109  | Leaf             | 60.30.181.102       | Spine                 | 0/24                  | 0/9                        | true                      | false                          |
| 60.30.181.109  | Leaf             | 60.30.181.110       | Leaf                  | 0/30                  | 0/30                       | true                      | true                           |
| 60.30.181.109  | Leaf             | 60.30.181.110       | Leaf                  | 0/31                  | 0/31                       | true                      | true                           |
| 60.30.181.108  | Leaf             | 60.30.181.101       | Spine                 | 0/21                  | 0/6                        | true                      | false                          |
| 60.30.181.108  | Leaf             | 60.30.181.102       | Spine                 | 0/22                  | 0/6                        | true                      | false                          |
| 60.30.181.108  | Leaf             | 60.30.181.107       | Leaf                  | 0/29                  | 0/29                       | true                      | true                           |
| 60.30.181.108  | Leaf             | 60.30.181.107       | Leaf                  | 0/30                  | 0/30                       | true                      | true                           |
| 60.30.181.102  | Spine            | 60.30.181.103       | Leaf                  | 0/1                   | 0/22                       | false                     | true                           |
| 60.30.181.102  | Spine            | 60.30.181.104       | Leaf                  | 0/2                   | 0/22                       | false                     | true                           |
| 60.30.181.102  | Spine            | 60.30.181.105       | Leaf                  | 0/3                   | 0/22                       | false                     | true                           |
| 60.30.181.102  | Spine            | 60.30.181.106       | Leaf                  | 0/4                   | 0/22                       | false                     | true                           |
| 60.30.181.102  | Spine            | 60.30.181.107       | Leaf                  | 0/5                   | 0/22                       | false                     | true                           |
| 60.30.181.102  | Spine            | 60.30.181.108       | Leaf                  | 0/6                   | 0/22                       | false                     | true                           |
| 60.30.181.102  | Spine            | 60.30.181.109       | Leaf                  | 0/7                   | 0/22                       | false                     | true                           |
| 60.30.181.102  | Spine            | 60.30.181.110       | Leaf                  | 0/8                   | 0/22                       | false                     | true                           |
| 60.30.181.102  | Spine            | 60.30.181.109       | Leaf                  | 0/9                   | 0/24                       | false                     | true                           |
| 60.30.181.102  | Spine            | 60.30.181.110       | Leaf                  | 0/15                  | 0/24                       | false                     | true                           |
| 60.30.181.102  | Spine            | 60.30.181.111       | BorderLeaf            | 0/27                  | 0/2                        | false                     | true                           |
| 60.30.181.102  | Spine            | 60.30.181.112       | BorderLeaf            | 0/28                  | 0/2                        | false                     | true                           |
| 60.30.181.102  | Spine            | 60.30.181.113       | BorderLeaf            | 0/29                  | 0/2                        | false                     | true                           |
| 60.30.181.102  | Spine            | 60.30.181.114       | BorderLeaf            | 0/30                  | 0/2                        | false                     | true                           |
+----------------+------------------+---------------------+-----------------------+-----------------------+----------------------------+---------------------------+--------------------------------+
--- Time Elapsed: 9.08233086s ---
`)
		default:
			return errors.New(fmt.Sprintf("Fabric %s does not exist", topologyShowPhysical_name))
		}
		return nil
	},
}

var topologyShowPhysical_name string

func init() {
	rootCmd.AddCommand(fabricCmd)
	fabricCmd.AddCommand(fabricShowCmd)

	fabricCmd.AddCommand(fabricTopologyCmd)
	fabricTopologyCmd.AddCommand(fabricTopologyShowCmd)
	fabricTopologyShowCmd.AddCommand(fabricTopologyShowPhysicalCmd)

	fabricTopologyShowPhysicalCmd.PersistentFlags().StringVar(&topologyShowPhysical_name, "name", "", "Name")
	fabricTopologyShowPhysicalCmd.MarkPersistentFlagRequired("name")
}
