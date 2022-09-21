/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"efa/cmd/structs"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

// inventoryCmd represents the inventory command
var inventoryCmd = &cobra.Command{
	Use:   "inventory",
	Short: "Inventory service commands",
}

var inventoryConfigBackupCmd = &cobra.Command{
	Use:   "config-backup",
	Short: "Config Backup commands",
}

var inventoryConfigBackupDetailCmd = &cobra.Command{
	Use:   "detail",
	Short: "List the Config Backup Detail for given UUID",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := new(structs.ConfigBackupDetail).Print(inventoryConfigBackupDetail_uuid)
		if result != "" {
			fmt.Print(result)
		}
		return err
	},
}

var inventoryConfigBackupExecuteCmd = &cobra.Command{
	Use:   "execute",
	Short: "execute SLX configuration backup",
	Run: func(cmd *cobra.Command, args []string) {
		uuid := "19551554-f75f-4b7b-8902-ea8d5fc23bde"

		fmt.Printf(`Config Backup Execute [success]
Config Backup execution UUID: %s

execute the CLI to get details: efa inventory config-backup detail --uuid %s
--- Time Elapsed: 609.389693ms ---
`, uuid, uuid)
	},
}

var inventoryConfigBackupHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "History commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := new(structs.ConfigBackupHistory).Print(inventoryConfigBackupHistory_ip)
		if result != "" {
			fmt.Print(result)
		}
		return err
	},
}

var inventoryDeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Device commands",
}

/**
TEST : 2022-08-30 16:05

> efa inventory device list
*/
var inventoryDeviceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Devices",
	Long: `List all Devices for a fabric.If fabric name is not provided,
        list of all devices not associated to a fabric will be returned.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
|  IP Address   |   Host Name   | Model | Chassis Name | Firmware |  ASN  |    Role    | Fabric |    LastDiscoveryTime    | Cert/Key Saved | Admin State | Maintenance Mode | Maintenance Mode on Reboot | Syslog Registered | SNMP Registered |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.101 | TB2-SLX-S1A   | 3012  | SLX9250-32C  | 20.3.2fb | 66000 | Spine      | TB2    | 2022-08-30 16:03:23 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.111 | TB2-SLX-DBL-A | 4004  | SLX9740-40C  | 20.4.1ca | 66011 | BorderLeaf | TB2    | 2022-08-30 16:03:40 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.112 | TB2-SLX-DBL-B | 4004  | SLX9740-40C  | 20.4.1ca | 66011 | BorderLeaf | TB2    | 2022-08-30 15:13:23 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.113 | TB2-SLX-MBL-A | 3012  | SLX9250-32C  | 20.3.2fb | 66012 | BorderLeaf | TB2    | 2022-08-30 16:03:51 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.114 | TB2-SLX-MBL-B | 3012  | SLX9250-32C  | 20.3.2fb | 66012 | BorderLeaf | TB2    | 2022-08-30 16:04:10 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.103 | TB2-SLX-Com-A | 3012  | SLX9250-32C  | 20.3.2fb | 66004 | Leaf       | TB2    | 2022-08-30 16:04:32 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.104 | TB2-SLX-Com-B | 3012  | SLX9250-32C  | 20.3.2fb | 66004 | Leaf       | TB2    | 2022-08-30 15:34:03 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.105 | TB2-SLX-AMF-A | 3012  | SLX9250-32C  | 20.3.2fb | 66001 | Leaf       | TB2    | 2022-08-30 16:04:55 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.106 | TB2-SLX-AMF-B | 3012  | SLX9250-32C  | 20.3.2fb | 66001 | Leaf       | TB2    | 2022-08-30 15:53:23 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.107 | TB2-SLX-SMF-A | 3012  | SLX9250-32C  | 20.3.2fb | 66002 | Leaf       | TB2    | 2022-08-30 16:05:18 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.108 | TB2-SLX-SMF-B | 3012  | SLX9250-32C  | 20.3.2fb | 66002 | Leaf       | TB2    | 2022-08-30 16:05:43 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.109 | TB2-SLX-UPF-A | 3012  | SLX9250-32C  | 20.4.1ca | 66003 | Leaf       | TB2    | 2022-08-30 16:06:09 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.110 | TB2-SLX-UPF-B | 3012  | SLX9250-32C  | 20.4.1ca | 66003 | Leaf       | TB2    | 2022-08-30 16:06:19 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
| 60.30.181.102 | TB2-SLX-S1B   | 3012  | SLX9250-32C  | 20.3.2fb | 66000 | Spine      | TB2    | 2022-08-30 14:56:37 KST | Y              | up          | Disable          | Enable                     | Y                 | Y               |
+---------------+---------------+-------+--------------+----------+-------+------------+--------+-------------------------+----------------+-------------+------------------+----------------------------+-------------------+-----------------+
Device Details
--- Time Elapsed: 237.426756ms ---
`)
	},
}

var inventoryDeviceExecuteCLICmd = &cobra.Command{
	Use:   "execute-cli",
	Short: "Execute CLI on the device",
	RunE: func(cmd *cobra.Command, args []string) error {
		command := strings.Replace(inventoryDeviceExecuteCLI_command, " | nomore", "", 1)

		var result string
		var err error

		switch command {
		case "show interface status":
			result, err = new(structs.ShowInterfaceStatus).Print(inventoryDeviceExecuteCLI_ip)
		case "show lldp neighbors":
			result, err = new(structs.ShowLLDPNeighbors).Print(inventoryDeviceExecuteCLI_ip)
		case "show ip interface brief":
			result, err = new(structs.ShowIpInterfaceBrief).Print(inventoryDeviceExecuteCLI_ip)
		default:
			return errors.New(fmt.Sprintf("Not found Execute CLI Command :: [%s]", inventoryDeviceExecuteCLI_command))
		}
		if result != "" {
			fmt.Print(result)
		}
		return err
	},
}

var inventoryDeviceDiscoveryTimeCmd = &cobra.Command{
	Use:   "discovery-time",
	Short: "System commands",
}

/**
TEST : 2022-08-30 16:00

> efa inventory device discovery-time list --fabric {name}
*/
var inventoryDeviceDiscoveryTimeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List device discovery interval time",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch inventoryDeviceDiscoveryTimeList_fabric {
		case "TB2":
			fmt.Print(`+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
|  IP Address   |   Host Name   | Model | Chassis Name | Firmware | Discovery Interval |    LastDiscoveryTime    |  DiscoveryReason  |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.101 | TB2-SLX-S1A   | 3012  | SLX9250-32C  | 20.3.2fb | 1 Hour 0 Min       | 2022-08-30 14:53:23 KST | TimerCollection   |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.111 | TB2-SLX-DBL-A | 4004  | SLX9740-40C  | 20.4.1ca | 1 Hour 0 Min       | 2022-08-30 16:03:40 KST | TimerCollection   |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.112 | TB2-SLX-DBL-B | 4004  | SLX9740-40C  | 20.4.1ca | 1 Hour 0 Min       | 2022-08-30 15:13:23 KST | TimerCollection   |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.113 | TB2-SLX-MBL-A | 3012  | SLX9250-32C  | 20.3.2fb | 1 Hour 0 Min       | 2022-08-30 14:53:51 KST | TimerCollection   |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.114 | TB2-SLX-MBL-B | 3012  | SLX9250-32C  | 20.3.2fb | 1 Hour 0 Min       | 2022-08-30 14:54:12 KST | TimerCollection   |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.103 | TB2-SLX-Com-A | 3012  | SLX9250-32C  | 20.3.2fb | 1 Hour 0 Min       | 2022-08-30 14:54:32 KST | TimerCollection   |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.104 | TB2-SLX-Com-B | 3012  | SLX9250-32C  | 20.3.2fb | 1 Hour 0 Min       | 2022-08-30 15:34:03 KST | RaslogEventReason |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.105 | TB2-SLX-AMF-A | 3012  | SLX9250-32C  | 20.3.2fb | 1 Hour 0 Min       | 2022-08-30 14:54:58 KST | TimerCollection   |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.106 | TB2-SLX-AMF-B | 3012  | SLX9250-32C  | 20.3.2fb | 1 Hour 0 Min       | 2022-08-30 15:53:23 KST | TimerCollection   |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.107 | TB2-SLX-SMF-A | 3012  | SLX9250-32C  | 20.3.2fb | 1 Hour 0 Min       | 2022-08-30 14:55:20 KST | TimerCollection   |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.108 | TB2-SLX-SMF-B | 3012  | SLX9250-32C  | 20.3.2fb | 1 Hour 0 Min       | 2022-08-30 14:55:48 KST | TimerCollection   |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.109 | TB2-SLX-UPF-A | 3012  | SLX9250-32C  | 20.4.1ca | 1 Hour 0 Min       | 2022-08-30 14:56:19 KST | TimerCollection   |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.110 | TB2-SLX-UPF-B | 3012  | SLX9250-32C  | 20.4.1ca | 1 Hour 0 Min       | 2022-08-30 14:56:28 KST | TimerCollection   |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
| 60.30.181.102 | TB2-SLX-S1B   | 3012  | SLX9250-32C  | 20.3.2fb | 1 Hour 0 Min       | 2022-08-30 14:56:37 KST | TimerCollection   |
+---------------+---------------+-------+--------------+----------+--------------------+-------------------------+-------------------+
Device Discovery interval time Details
--- Time Elapsed: 150.579117ms ---
`)
		default:
			return errors.New(fmt.Sprintf("Error Message: No devices are part of the Fabric: %s", inventoryDeviceDiscoveryTimeList_fabric))
		}
		return nil
	},
}

var inventoryDriftReconcileCmd = &cobra.Command{
	Use:   "drift-reconcile",
	Short: "Drift Reconcile commands",
}

var inventoryDriftReconcileHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "List the Drift Reconcile History by default. If device IP is given, list history for that device.",
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := new(structs.ConfigBackupHistory).Print(inventoryConfigBackupHistory_ip)
		if result != "" {
			fmt.Print(result)
		}
		return err
	},
}

var inventoryFirmwareHostCmd = &cobra.Command{
	Use:   "firmware-host",
	Short: "Firmware host commands",
}

/**
TEST : 2022-08-30 16:35

> efa inventory firmware-host list
*/
var inventoryFirmwareHostListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered Firmware Hosts.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`+----+---------------+----------+
| ID |  IP Address   | Protocol |
+----+---------------+----------+
| 1  | 60.30.181.116 | sftp     |
+----+---------------+----------+
| 7  | 60.30.181.115 | scp      |
+----+---------------+----------+
Firmware Host Details
--- Time Elapsed: 151.968036ms ---
`)
	},
}

/**


 */

var inventoryConfigBackupHistory_ip string

var inventoryConfigBackupDetail_uuid string

var inventoryConfigBackupExecute_ip string

var inventoryDeviceExecuteCLI_ip string
var inventoryDeviceExecuteCLI_command string

var inventoryDeviceDiscoveryTimeList_fabric string

var inventoryDriftReconcileHistory_ip string

func init() {
	rootCmd.AddCommand(inventoryCmd)

	// efa inventory config-backup
	inventoryCmd.AddCommand(inventoryConfigBackupCmd)
	inventoryConfigBackupCmd.AddCommand(inventoryConfigBackupHistoryCmd)
	inventoryConfigBackupHistoryCmd.PersistentFlags().StringVar(&inventoryConfigBackupHistory_ip, "ip", "", "IP")

	inventoryConfigBackupCmd.AddCommand(inventoryConfigBackupDetailCmd)

	inventoryConfigBackupDetailCmd.PersistentFlags().StringVar(&inventoryConfigBackupDetail_uuid, "uuid", "", "UUID")
	inventoryConfigBackupDetailCmd.MarkPersistentFlagRequired("uuid")

	inventoryConfigBackupDetailCmd.Flags().String("show-config", "", "Show Config")
	inventoryConfigBackupDetailCmd.Flags().Lookup("show-config").NoOptDefVal = "none"

	inventoryConfigBackupCmd.AddCommand(inventoryConfigBackupExecuteCmd)
	inventoryConfigBackupExecuteCmd.PersistentFlags().StringVar(&inventoryConfigBackupExecute_ip, "ip", "", "IP")

	// efa inventory device
	inventoryCmd.AddCommand(inventoryDeviceCmd)
	inventoryDeviceCmd.AddCommand(inventoryDeviceListCmd)
	inventoryDeviceCmd.AddCommand(inventoryDeviceExecuteCLICmd)

	inventoryDeviceExecuteCLICmd.PersistentFlags().StringVar(&inventoryDeviceExecuteCLI_command, "command", "", "Command")
	inventoryDeviceExecuteCLICmd.MarkPersistentFlagRequired("command")

	inventoryDeviceExecuteCLICmd.PersistentFlags().StringVar(&inventoryDeviceExecuteCLI_ip, "ip", "", "IP")
	inventoryDeviceExecuteCLICmd.MarkPersistentFlagRequired("ip")

	inventoryDeviceCmd.AddCommand(inventoryDeviceDiscoveryTimeCmd)
	inventoryDeviceDiscoveryTimeCmd.AddCommand(inventoryDeviceDiscoveryTimeListCmd)

	inventoryDeviceDiscoveryTimeListCmd.PersistentFlags().StringVar(&inventoryDeviceDiscoveryTimeList_fabric, "fabric", "", "FABRIC")
	inventoryDeviceDiscoveryTimeListCmd.MarkPersistentFlagRequired("fabric")

	// efa inventory drift-reconcile
	inventoryCmd.AddCommand(inventoryDriftReconcileCmd)
	inventoryDriftReconcileCmd.AddCommand(inventoryDriftReconcileHistoryCmd)

	inventoryDriftReconcileHistoryCmd.PersistentFlags().StringVar(&inventoryDriftReconcileHistory_ip, "ip", "", "IP")

	// efa inventory firmware-host
	inventoryCmd.AddCommand(inventoryFirmwareHostCmd)
	inventoryFirmwareHostCmd.AddCommand(inventoryFirmwareHostListCmd)
}
