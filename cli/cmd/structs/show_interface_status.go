package structs

import (
	"errors"
	"fmt"
)

type ShowInterfaceStatus struct {
}

/**
TEST : 2022-08-30 15:40

> efa inventory device execute-cli --ip {ip} --command "show interface status"
*/
func (s *ShowInterfaceStatus) Print(ip string) (result string, err error) {
	if ip == "60.30.181.101" {
		result = `Execute CLI[success]
+---------------+-------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------+
|  IP Address   |  Host Name  | Fabric |        Command        | Status  | Reason |                                               Output                                               |
+---------------+-------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------+
| 60.30.181.101 | TB2-SLX-S1A | TB2    | show interface status | Success |        | TB2-SLX-S1A# show interface status                                                                 |
|               |             |        |                       |         |        | Flags:  L - Internal Loopback                                                                      |
|               |             |        |                       |         |        | --------------------------------------------------------------------------------                   |
|               |             |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                       |
|               |             |        |                       |         |        | --------------------------------------------------------------------------------                   |
|               |             |        |                       |         |        | Eth 0/1      connected (up)  --         100G     100G            Link to 60.30.181.103 Leaf        |
|               |             |        |                       |         |        | Eth 0/2      connected (up)  --         100G     100G            Link to 60.30.181.104 Leaf        |
|               |             |        |                       |         |        | Eth 0/3      connected (up)  --         100G     100G            Link to 60.30.181.105 Leaf        |
|               |             |        |                       |         |        | Eth 0/4      connected (up)  --         100G     100G            Link to 60.30.181.106 Leaf        |
|               |             |        |                       |         |        | Eth 0/5      connected (up)  --         100G     100G            Link to 60.30.181.107 Leaf        |
|               |             |        |                       |         |        | Eth 0/6      connected (up)  --         100G     100G            Link to 60.30.181.108 Leaf        |
|               |             |        |                       |         |        | Eth 0/7      connected (up)  --         100G     100G            Link to 60.30.181.109 Leaf        |
|               |             |        |                       |         |        | Eth 0/8      connected (up)  --         100G     100G            Link to 60.30.181.110 Leaf        |
|               |             |        |                       |         |        | Eth 0/9      adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/10     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/11     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/12     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/13     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/14     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/15     connected (up)  --         100G     100G            Link to 60.30.181.109 Leaf        |
|               |             |        |                       |         |        | Eth 0/16     connected (up)  --         100G     100G            Link to 60.30.181.110 Leaf        |
|               |             |        |                       |         |        | Eth 0/17     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/18     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/19     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/20     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/21     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/22     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/23     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/24     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/25     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/26     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/27     connected (up)  --         100G     100G            Link to 60.30.181.111 border-leaf |
|               |             |        |                       |         |        | Eth 0/28     connected (up)  --         100G     100G            Link to 60.30.181.112 border-leaf |
|               |             |        |                       |         |        | Eth 0/29     connected (up)  --         100G     100G            Link to 60.30.181.113 border-leaf |
|               |             |        |                       |         |        | Eth 0/30     connected (up)  --         100G     100G            Link to 60.30.181.114 border-leaf |
|               |             |        |                       |         |        | Eth 0/31     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/32:1   connected (up)  --         1G       1G-SFP-T                                          |
|               |             |        |                       |         |        | Eth 0/32:2   adminDown       --         --       1G-SFP-T                                          |
|               |             |        |                       |         |        | Eth 0/32:3   adminDown       --         --       1G-SFP-T                                          |
|               |             |        |                       |         |        | Eth 0/32:4   adminDown       --         --       1G-SFP-T                                          |
|               |             |        |                       |         |        |                                                                                                    |
+---------------+-------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 5.641453671s ---
`
	} else if ip == "60.30.181.102" {
		result = `Execute CLI[success]
+---------------+-------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------+
|  IP Address   |  Host Name  | Fabric |        Command        | Status  | Reason |                                               Output                                               |
+---------------+-------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------+
| 60.30.181.102 | TB2-SLX-S1B | TB2    | show interface status | Success |        | TB2-SLX-S1B# show interface status                                                                 |
|               |             |        |                       |         |        | Flags:  L - Internal Loopback                                                                      |
|               |             |        |                       |         |        | --------------------------------------------------------------------------------                   |
|               |             |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                       |
|               |             |        |                       |         |        | --------------------------------------------------------------------------------                   |
|               |             |        |                       |         |        | Eth 0/1      connected (up)  --         100G     100G            Link to 60.30.181.103 Leaf        |
|               |             |        |                       |         |        | Eth 0/2      connected (up)  --         100G     100G            Link to 60.30.181.104 Leaf        |
|               |             |        |                       |         |        | Eth 0/3      connected (up)  --         100G     100G            Link to 60.30.181.105 Leaf        |
|               |             |        |                       |         |        | Eth 0/4      connected (up)  --         100G     100G            Link to 60.30.181.106 Leaf        |
|               |             |        |                       |         |        | Eth 0/5      connected (up)  --         100G     100G            Link to 60.30.181.107 Leaf        |
|               |             |        |                       |         |        | Eth 0/6      connected (up)  --         100G     100G            Link to 60.30.181.108 Leaf        |
|               |             |        |                       |         |        | Eth 0/7      connected (up)  --         100G     100G            Link to 60.30.181.109 Leaf        |
|               |             |        |                       |         |        | Eth 0/8      connected (up)  --         100G     100G            Link to 60.30.181.110 Leaf        |
|               |             |        |                       |         |        | Eth 0/9      connected (up)  --         100G     100G            Link to 60.30.181.109 Leaf        |
|               |             |        |                       |         |        | Eth 0/10     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/11     connected (up)  --         100G     100G                                              |
|               |             |        |                       |         |        | Eth 0/12     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/13     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/14     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/15     connected (up)  --         100G     100G            Link to 60.30.181.110 Leaf        |
|               |             |        |                       |         |        | Eth 0/16     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/17     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/18     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/19     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/20     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/21     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/22     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/23     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/24     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/25     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/26     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/27     connected (up)  --         100G     100G            Link to 60.30.181.111 border-leaf |
|               |             |        |                       |         |        | Eth 0/28     connected (up)  --         100G     100G            Link to 60.30.181.112 border-leaf |
|               |             |        |                       |         |        | Eth 0/29     connected (up)  --         100G     100G            Link to 60.30.181.113 border-leaf |
|               |             |        |                       |         |        | Eth 0/30     connected (up)  --         100G     100G            Link to 60.30.181.114 border-leaf |
|               |             |        |                       |         |        | Eth 0/31     adminDown       --         --       --                                                |
|               |             |        |                       |         |        | Eth 0/32     adminDown       --         --       1G-SFP-T                                          |
|               |             |        |                       |         |        |                                                                                                    |
+---------------+-------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 5.550494214s ---
`
	} else if ip == "60.30.181.103" {
		result = `Execute CLI[success]
+---------------+---------------+--------+-----------------------+---------+--------+-------------------------------------------------------------------------------------------------------------+
|  IP Address   |   Host Name   | Fabric |        Command        | Status  | Reason |                                                   Output                                                    |
+---------------+---------------+--------+-----------------------+---------+--------+-------------------------------------------------------------------------------------------------------------+
| 60.30.181.103 | TB2-SLX-Com-A | TB2    | show interface status | Success |        | TB2-SLX-Com-A# show interface status                                                                        |
|               |               |        |                       |         |        | Flags:  L - Internal Loopback                                                                               |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                            |
|               |               |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                                |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                            |
|               |               |        |                       |         |        | Eth 0/1:1    connected (up)  --         25G      100G            Port-channel Master1_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/1:2    connected (up)  --         25G      100G            Port-channel Master2_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/1:3    connected (up)  --         25G      100G            Port-channel Master1_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/1:4    connected (up)  --         25G      100G            Port-channel Master2_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/2:1    connected (up)  --         25G      100G            Port-channel Master3_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/2:2    connected (up)  --         25G      100G            Port-channel Common_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/2:3    connected (up)  --         25G      100G            Port-channel Master3_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/2:4    connected (up)  --         25G      100G            Port-channel Common_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:1    connected (up)  --         25G      100G            Port-channel Common2_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/3:2    adminDown       --         --       100G                                                       |
|               |               |        |                       |         |        | Eth 0/3:3    connected (up)  --         25G      100G                                                       |
|               |               |        |                       |         |        | Eth 0/3:4    notconnected    --         --       100G                                                       |
|               |               |        |                       |         |        | Eth 0/4:1    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/4:2    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/4:3    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/4:4    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/5:1    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/5:2    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/5:3    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/5:4    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/6:1    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/6:2    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/6:3    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/6:4    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/7:1    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/7:2    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/7:3    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/7:4    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/8:1    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/8:2    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/8:3    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/8:4    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/9:1    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/9:2    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/9:3    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/9:4    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/10:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/10:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/10:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/10:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/11:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/11:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/11:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/11:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/12:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/12:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/12:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/12:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/13:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/13:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/13:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/13:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/14:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/14:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/14:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/14:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/15:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/15:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/15:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/15:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/16:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/16:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/16:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/16:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/17:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/17:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/17:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/17:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/18:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/18:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/18:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/18:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/19:1   connected (up)  --         25G      100G            Port-channel SDS-ctrl1_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/19:2   connected (up)  --         25G      100G            Port-channel SDS-ctrl2_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/19:3   connected (up)  --         25G      100G            Port-channel SDS-ctrl1_S7 Member interface |
|               |               |        |                       |         |        | Eth 0/19:4   connected (up)  --         25G      100G            Port-channel SDS-ctrl2_S7 Member interface |
|               |               |        |                       |         |        | Eth 0/20:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/20:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/20:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/20:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/21     connected (up)  --         100G     100G            Link to 60.30.181.101 Spine                |
|               |               |        |                       |         |        | Eth 0/22     connected (up)  --         100G     100G            Link to 60.30.181.102 Spine                |
|               |               |        |                       |         |        | Eth 0/23     adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/24     adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/25     adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/26     adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/27     adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/28     adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/29     connected (up)  --         100G     100G            clusterPeerIntfMember                      |
|               |               |        |                       |         |        | Eth 0/30     connected (up)  --         100G     100G            clusterPeerIntfMember                      |
|               |               |        |                       |         |        | Eth 0/31:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/31:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/31:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/31:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/32:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/32:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/32:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/32:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Po 1         connected (up)  Trunk      25G      --              Master1_S1                                 |
|               |               |        |                       |         |        | Po 2         connected (up)  Trunk      25G      --              Master2_S1                                 |
|               |               |        |                       |         |        | Po 3         notconnected    Trunk      --       --              Master1_S3                                 |
|               |               |        |                       |         |        | Po 4         notconnected    Trunk      --       --              Master2_S3                                 |
|               |               |        |                       |         |        | Po 5         connected (up)  Trunk      25G      --              Master3_S1                                 |
|               |               |        |                       |         |        | Po 6         connected (up)  Trunk      25G      --              Common_S1                                  |
|               |               |        |                       |         |        | Po 7         notconnected    Trunk      --       --              Master3_S3                                 |
|               |               |        |                       |         |        | Po 8         notconnected    Trunk      --       --              Common_S3                                  |
|               |               |        |                       |         |        | Po 9         connected (up)  Trunk      25G      --              Common2_S3                                 |
|               |               |        |                       |         |        | Po 50        connected (up)  --         10G      --                                                         |
|               |               |        |                       |         |        | Po 64        connected (up)  --         200G     --              MCTPeerInterface                           |
|               |               |        |                       |         |        | Po 101       connected (up)  Trunk      25G      --              SDS-ctrl1_S3                               |
|               |               |        |                       |         |        | Po 102       connected (up)  Trunk      25G      --              SDS-ctrl2_S3                               |
|               |               |        |                       |         |        | Po 103       connected (up)  Trunk      25G      --              SDS-ctrl1_S7                               |
|               |               |        |                       |         |        | Po 104       connected (up)  Trunk      25G      --              SDS-ctrl2_S7                               |
|               |               |        |                       |         |        |                                                                                                             |
+---------------+---------------+--------+-----------------------+---------+--------+-------------------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 5.834908968s ---
`
	} else if ip == "60.30.181.104" {
		result = `Execute CLI[success]
+---------------+---------------+--------+-----------------------+---------+--------+-------------------------------------------------------------------------------------------------------------+
|  IP Address   |   Host Name   | Fabric |        Command        | Status  | Reason |                                                   Output                                                    |
+---------------+---------------+--------+-----------------------+---------+--------+-------------------------------------------------------------------------------------------------------------+
| 60.30.181.104 | TB2-SLX-Com-B | TB2    | show interface status | Success |        | TB2-SLX-Com-B# show interface status                                                                        |
|               |               |        |                       |         |        | Flags:  L - Internal Loopback                                                                               |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                            |
|               |               |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                                |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                            |
|               |               |        |                       |         |        | Eth 0/1:1    connected (up)  --         25G      100G            Port-channel Master1_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/1:2    connected (up)  --         25G      100G            Port-channel Master2_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/1:3    connected (up)  --         25G      100G            Port-channel Master1_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/1:4    connected (up)  --         25G      100G            Port-channel Master2_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/2:1    connected (up)  --         25G      100G            Port-channel Master3_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/2:2    connected (up)  --         25G      100G            Port-channel Common_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/2:3    connected (up)  --         25G      100G            Port-channel Master3_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/2:4    notconnected    --         --       100G            Port-channel Common_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:1    connected (up)  --         25G      100G            Port-channel Common2_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/3:2    adminDown       --         --       100G                                                       |
|               |               |        |                       |         |        | Eth 0/3:3    adminDown       --         --       100G                                                       |
|               |               |        |                       |         |        | Eth 0/3:4    notconnected    --         --       100G                                                       |
|               |               |        |                       |         |        | Eth 0/4:1    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/4:2    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/4:3    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/4:4    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/5:1    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/5:2    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/5:3    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/5:4    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/6:1    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/6:2    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/6:3    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/6:4    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/7:1    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/7:2    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/7:3    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/7:4    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/8:1    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/8:2    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/8:3    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/8:4    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/9:1    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/9:2    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/9:3    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/9:4    adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/10:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/10:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/10:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/10:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/11:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/11:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/11:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/11:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/12:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/12:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/12:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/12:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/13:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/13:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/13:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/13:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/14:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/14:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/14:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/14:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/15:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/15:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/15:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/15:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/16:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/16:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/16:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/16:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/17:1   sfpAbsent       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/17:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/17:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/17:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/18:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/18:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/18:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/18:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/19:1   connected (up)  --         25G      100G            Port-channel SDS-ctrl1_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/19:2   connected (up)  --         25G      100G            Port-channel SDS-ctrl2_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/19:3   connected (up)  --         25G      100G            Port-channel SDS-ctrl1_S7 Member interface |
|               |               |        |                       |         |        | Eth 0/19:4   connected (up)  --         25G      100G            Port-channel SDS-ctrl2_S7 Member interface |
|               |               |        |                       |         |        | Eth 0/20:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/20:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/20:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/20:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/21     connected (up)  --         100G     100G            Link to 60.30.181.101 Spine                |
|               |               |        |                       |         |        | Eth 0/22     connected (up)  --         100G     100G            Link to 60.30.181.102 Spine                |
|               |               |        |                       |         |        | Eth 0/23     adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/24     adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/25     adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/26     adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/27     adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/28     adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/29     connected (up)  --         100G     100G            clusterPeerIntfMember                      |
|               |               |        |                       |         |        | Eth 0/30     connected (up)  --         100G     100G            clusterPeerIntfMember                      |
|               |               |        |                       |         |        | Eth 0/31:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/31:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/31:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/31:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/32:1   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/32:2   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/32:3   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Eth 0/32:4   adminDown       --         --       --                                                         |
|               |               |        |                       |         |        | Po 1         connected (up)  Trunk      25G      --              Master1_S1                                 |
|               |               |        |                       |         |        | Po 2         connected (up)  Trunk      25G      --              Master2_S1                                 |
|               |               |        |                       |         |        | Po 3         notconnected    Trunk      --       --              Master1_S3                                 |
|               |               |        |                       |         |        | Po 4         notconnected    Trunk      --       --              Master2_S3                                 |
|               |               |        |                       |         |        | Po 5         connected (up)  Trunk      25G      --              Master3_S1                                 |
|               |               |        |                       |         |        | Po 6         connected (up)  Trunk      25G      --              Common_S1                                  |
|               |               |        |                       |         |        | Po 7         notconnected    Trunk      --       --              Master3_S3                                 |
|               |               |        |                       |         |        | Po 8         notconnected    Trunk      --       --              Common_S3                                  |
|               |               |        |                       |         |        | Po 9         connected (up)  Trunk      25G      --              Common2_S3                                 |
|               |               |        |                       |         |        | Po 64        connected (up)  --         200G     --              MCTPeerInterface                           |
|               |               |        |                       |         |        | Po 101       connected (up)  Trunk      25G      --              SDS-ctrl1_S3                               |
|               |               |        |                       |         |        | Po 102       connected (up)  Trunk      25G      --              SDS-ctrl2_S3                               |
|               |               |        |                       |         |        | Po 103       connected (up)  Trunk      25G      --              SDS-ctrl1_S7                               |
|               |               |        |                       |         |        | Po 104       connected (up)  Trunk      25G      --              SDS-ctrl2_S7                               |
|               |               |        |                       |         |        |                                                                                                             |
+---------------+---------------+--------+-----------------------+---------+--------+-------------------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 5.517077407s ---
`
	} else if ip == "60.30.181.105" {
		result = `Execute CLI[success]
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------------+
|  IP Address   |   Host Name   | Fabric |        Command        | Status  | Reason |                                                  Output                                                  |
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------------+
| 60.30.181.105 | TB2-SLX-AMF-A | TB2    | show interface status | Success |        | TB2-SLX-AMF-A# show interface status                                                                     |
|               |               |        |                       |         |        | Flags:  L - Internal Loopback                                                                            |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                         |
|               |               |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                             |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                         |
|               |               |        |                       |         |        | Eth 0/1:1    connected (up)  --         25G      100G            Port-channel A-1_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/1:2    notconnected    --         --       100G            Port-channel A-2_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/1:3    connected (up)  --         25G      100G            Port-channel A-1_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/1:4    connected (up)  --         25G      100G            Port-channel A-2_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/2:1    connected (up)  --         25G      100G            Port-channel A-3_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/2:2    connected (up)  --         25G      100G            Port-channel A-4_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/2:3    connected (up)  --         25G      100G            Port-channel A-3_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/2:4    connected (up)  --         25G      100G            Port-channel A-4_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:1    connected (up)  --         25G      100G            Port-channel A-5_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:2    connected (up)  --         25G      100G            Port-channel A-6_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:3    connected (up)  --         25G      100G            Port-channel A-5_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:4    connected (up)  --         25G      100G            Port-channel A-6_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/4:1    connected (up)  --         25G      100G            Port-channel A-7_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/4:2    connected (up)  --         25G      100G            Port-channel A-8_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/4:3    connected (up)  --         25G      100G            Port-channel A-7_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/4:4    connected (up)  --         25G      100G            Port-channel A-8_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/5:1    connected (up)  --         25G      100G            Port-channel A-9_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/5:2    connected (up)  --         25G      100G            Port-channel A-10_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/5:3    connected (up)  --         25G      100G            Port-channel A-9_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/5:4    connected (up)  --         25G      100G            Port-channel A-10_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/6:1    connected (up)  --         25G      100G            Port-channel A-11_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/6:2    connected (up)  --         25G      100G            Port-channel A-12_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/6:3    connected (up)  --         25G      100G            Port-channel A-11_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/6:4    notconnected    --         --       100G            Port-channel A-12_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/7:1    connected (up)  --         25G      100G            Port-channel A-13_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/7:2    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/7:3    connected (up)  --         25G      100G            Port-channel A-13_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/7:4    notconnected    --         --       100G            Port-channel A-EMF1_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/8:1    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/8:2    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/8:3    notconnected    --         --       100G            Port-channel A-EMF2_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/8:4    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/9:1    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/9:2    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/9:3    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/9:4    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/10:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/10:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/10:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/10:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/11:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/11:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/11:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/11:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/12:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/12:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/12:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/12:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/13:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/13:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/13:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/13:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/14:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/14:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/14:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/14:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/15:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/15:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/15:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/15:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/16:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/16:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/16:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/16:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/17:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/17:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/17:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/17:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/18:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/18:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/18:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/18:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/19:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/19:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/19:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/19:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/20:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/20:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/20:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/20:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/21     connected (up)  --         100G     100G            Link to 60.30.181.101 Spine             |
|               |               |        |                       |         |        | Eth 0/22     connected (up)  --         100G     100G            Link to 60.30.181.102 Spine             |
|               |               |        |                       |         |        | Eth 0/23:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/23:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/23:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/23:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/24     adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/25     connected (up)  --         100G     100G                                                    |
|               |               |        |                       |         |        | Eth 0/26     adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/27     adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/28     adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/29     adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/30     connected (up)  --         100G     100G            clusterPeerIntfMember                   |
|               |               |        |                       |         |        | Eth 0/31:1   adminDown       --         --       10G-SFP-SR                                              |
|               |               |        |                       |         |        | Eth 0/31:2   adminDown       --         --       10G-SFP-SR                                              |
|               |               |        |                       |         |        | Eth 0/31:3   adminDown       --         --       10G-SFP-SR                                              |
|               |               |        |                       |         |        | Eth 0/31:4   adminDown       --         --       10G-SFP-SR                                              |
|               |               |        |                       |         |        | Eth 0/32:1   connected (up)  --         1G       1G-SFP-T                                                |
|               |               |        |                       |         |        | Eth 0/32:2   adminDown       --         --       1G-SFP-T                                                |
|               |               |        |                       |         |        | Eth 0/32:3   adminDown       --         --       1G-SFP-T                                                |
|               |               |        |                       |         |        | Eth 0/32:4   adminDown       --         --       1G-SFP-T                                                |
|               |               |        |                       |         |        | Po 1         connected (up)  Trunk      25G      --              A-1_S1                                  |
|               |               |        |                       |         |        | Po 2         notconnected    Trunk      --       --              A-2_S1                                  |
|               |               |        |                       |         |        | Po 3         connected (up)  Trunk      25G      --              A-3_S1                                  |
|               |               |        |                       |         |        | Po 4         connected (up)  Trunk      25G      --              A-4_S1                                  |
|               |               |        |                       |         |        | Po 5         connected (up)  Trunk      25G      --              A-5_S1                                  |
|               |               |        |                       |         |        | Po 6         connected (up)  Trunk      25G      --              A-6_S1                                  |
|               |               |        |                       |         |        | Po 7         connected (up)  Trunk      25G      --              A-7_S1                                  |
|               |               |        |                       |         |        | Po 8         connected (up)  Trunk      25G      --              A-8_S1                                  |
|               |               |        |                       |         |        | Po 9         connected (up)  Trunk      25G      --              A-9_S1                                  |
|               |               |        |                       |         |        | Po 10        connected (up)  Trunk      25G      --              A-10_S1                                 |
|               |               |        |                       |         |        | Po 11        connected (up)  Trunk      25G      --              A-11_S1                                 |
|               |               |        |                       |         |        | Po 12        notconnected    Trunk      --       --              A-12_S1                                 |
|               |               |        |                       |         |        | Po 13        connected (up)  Trunk      25G      --              A-13_S1                                 |
|               |               |        |                       |         |        | Po 31        notconnected    Trunk      --       --              A-1_S3                                  |
|               |               |        |                       |         |        | Po 32        notconnected    Trunk      --       --              A-2_S3                                  |
|               |               |        |                       |         |        | Po 33        notconnected    Trunk      --       --              A-3_S3                                  |
|               |               |        |                       |         |        | Po 34        notconnected    Trunk      --       --              A-4_S3                                  |
|               |               |        |                       |         |        | Po 35        notconnected    Trunk      --       --              A-5_S3                                  |
|               |               |        |                       |         |        | Po 36        notconnected    Trunk      --       --              A-6_S3                                  |
|               |               |        |                       |         |        | Po 37        notconnected    Trunk      --       --              A-7_S3                                  |
|               |               |        |                       |         |        | Po 38        notconnected    Trunk      --       --              A-8_S3                                  |
|               |               |        |                       |         |        | Po 39        notconnected    Trunk      --       --              A-9_S3                                  |
|               |               |        |                       |         |        | Po 40        notconnected    Trunk      --       --              A-10_S3                                 |
|               |               |        |                       |         |        | Po 41        notconnected    Trunk      --       --              A-11_S3                                 |
|               |               |        |                       |         |        | Po 42        notconnected    Trunk      --       --              A-12_S3                                 |
|               |               |        |                       |         |        | Po 43        notconnected    Trunk      --       --              A-13_S3                                 |
|               |               |        |                       |         |        | Po 64        connected (up)  --         100G     --              MCTPeerInterface                        |
|               |               |        |                       |         |        | Po 121       notconnected    Trunk      --       --              A-EMF1_S3                               |
|               |               |        |                       |         |        | Po 122       notconnected    Trunk      --       --              A-EMF2_S3                               |
|               |               |        |                       |         |        |                                                                                                          |
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 5.947725877s ---
`
	} else if ip == "60.30.181.106" {
		result = `Execute CLI[success]
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------------+
|  IP Address   |   Host Name   | Fabric |        Command        | Status  | Reason |                                                  Output                                                  |
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------------+
| 60.30.181.106 | TB2-SLX-AMF-B | TB2    | show interface status | Success |        | TB2-SLX-AMF-B# show interface status                                                                     |
|               |               |        |                       |         |        | Flags:  L - Internal Loopback                                                                            |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                         |
|               |               |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                             |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                         |
|               |               |        |                       |         |        | Eth 0/1:1    connected (up)  --         25G      100G            Port-channel A-1_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/1:2    connected (up)  --         25G      100G            Port-channel A-2_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/1:3    connected (up)  --         25G      100G            Port-channel A-1_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/1:4    connected (up)  --         25G      100G            Port-channel A-2_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/2:1    connected (up)  --         25G      100G            Port-channel A-3_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/2:2    connected (up)  --         25G      100G            Port-channel A-4_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/2:3    connected (up)  --         25G      100G            Port-channel A-3_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/2:4    connected (up)  --         25G      100G            Port-channel A-4_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:1    connected (up)  --         25G      100G            Port-channel A-5_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:2    connected (up)  --         25G      100G            Port-channel A-6_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:3    connected (up)  --         25G      100G            Port-channel A-5_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:4    connected (up)  --         25G      100G            Port-channel A-6_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/4:1    connected (up)  --         25G      100G            Port-channel A-7_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/4:2    connected (up)  --         25G      100G            Port-channel A-8_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/4:3    connected (up)  --         25G      100G            Port-channel A-7_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/4:4    connected (up)  --         25G      100G            Port-channel A-8_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/5:1    connected (up)  --         25G      100G            Port-channel A-9_S1 Member interface    |
|               |               |        |                       |         |        | Eth 0/5:2    connected (up)  --         25G      100G            Port-channel A-10_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/5:3    connected (up)  --         25G      100G            Port-channel A-9_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/5:4    connected (up)  --         25G      100G            Port-channel A-10_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/6:1    connected (up)  --         25G      100G            Port-channel A-11_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/6:2    connected (up)  --         25G      100G            Port-channel A-12_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/6:3    connected (up)  --         25G      100G            Port-channel A-11_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/6:4    connected (up)  --         25G      100G            Port-channel A-12_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/7:1    connected (up)  --         25G      100G            Port-channel A-13_S1 Member interface   |
|               |               |        |                       |         |        | Eth 0/7:2    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/7:3    connected (up)  --         25G      100G            Port-channel A-13_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/7:4    notconnected    --         --       100G            Port-channel A-EMF1_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/8:1    notconnected    --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/8:2    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/8:3    notconnected    --         --       100G            Port-channel A-EMF2_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/8:4    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/9:1    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/9:2    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/9:3    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/9:4    adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/10:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/10:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/10:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/10:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/11:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/11:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/11:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/11:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/12:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/12:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/12:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/12:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/13:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/13:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/13:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/13:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/14:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/14:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/14:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/14:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/15:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/15:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/15:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/15:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/16:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/16:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/16:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/16:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/17:1   sfpAbsent       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/17:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/17:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/17:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/18:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/18:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/18:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/18:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/19:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/19:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/19:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/19:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/20:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/20:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/20:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/20:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/21     connected (up)  --         100G     100G            Link to 60.30.181.101 Spine             |
|               |               |        |                       |         |        | Eth 0/22     connected (up)  --         100G     100G            Link to 60.30.181.102 Spine             |
|               |               |        |                       |         |        | Eth 0/23     adminDown       --         --       100G                                                    |
|               |               |        |                       |         |        | Eth 0/24     adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/25     connected (up)  --         100G     100G                                                    |
|               |               |        |                       |         |        | Eth 0/26     adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/27     adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/28     adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/29     adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/30     connected (up)  --         100G     100G            clusterPeerIntfMember                   |
|               |               |        |                       |         |        | Eth 0/31:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/31:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/31:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/31:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/32:1   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/32:2   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/32:3   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Eth 0/32:4   adminDown       --         --       --                                                      |
|               |               |        |                       |         |        | Po 1         connected (up)  Trunk      25G      --              A-1_S1                                  |
|               |               |        |                       |         |        | Po 2         connected (up)  Trunk      25G      --              A-2_S1                                  |
|               |               |        |                       |         |        | Po 3         connected (up)  Trunk      25G      --              A-3_S1                                  |
|               |               |        |                       |         |        | Po 4         connected (up)  Trunk      25G      --              A-4_S1                                  |
|               |               |        |                       |         |        | Po 5         connected (up)  Trunk      25G      --              A-5_S1                                  |
|               |               |        |                       |         |        | Po 6         connected (up)  Trunk      25G      --              A-6_S1                                  |
|               |               |        |                       |         |        | Po 7         connected (up)  Trunk      25G      --              A-7_S1                                  |
|               |               |        |                       |         |        | Po 8         connected (up)  Trunk      25G      --              A-8_S1                                  |
|               |               |        |                       |         |        | Po 9         connected (up)  Trunk      25G      --              A-9_S1                                  |
|               |               |        |                       |         |        | Po 10        connected (up)  Trunk      25G      --              A-10_S1                                 |
|               |               |        |                       |         |        | Po 11        connected (up)  Trunk      25G      --              A-11_S1                                 |
|               |               |        |                       |         |        | Po 12        notconnected    Trunk      --       --              A-12_S1                                 |
|               |               |        |                       |         |        | Po 13        connected (up)  Trunk      25G      --              A-13_S1                                 |
|               |               |        |                       |         |        | Po 31        notconnected    Trunk      --       --              A-1_S3                                  |
|               |               |        |                       |         |        | Po 32        notconnected    Trunk      --       --              A-2_S3                                  |
|               |               |        |                       |         |        | Po 33        notconnected    Trunk      --       --              A-3_S3                                  |
|               |               |        |                       |         |        | Po 34        notconnected    Trunk      --       --              A-4_S3                                  |
|               |               |        |                       |         |        | Po 35        notconnected    Trunk      --       --              A-5_S3                                  |
|               |               |        |                       |         |        | Po 36        notconnected    Trunk      --       --              A-6_S3                                  |
|               |               |        |                       |         |        | Po 37        notconnected    Trunk      --       --              A-7_S3                                  |
|               |               |        |                       |         |        | Po 38        notconnected    Trunk      --       --              A-8_S3                                  |
|               |               |        |                       |         |        | Po 39        notconnected    Trunk      --       --              A-9_S3                                  |
|               |               |        |                       |         |        | Po 40        notconnected    Trunk      --       --              A-10_S3                                 |
|               |               |        |                       |         |        | Po 41        notconnected    Trunk      --       --              A-11_S3                                 |
|               |               |        |                       |         |        | Po 42        notconnected    Trunk      --       --              A-12_S3                                 |
|               |               |        |                       |         |        | Po 43        notconnected    Trunk      --       --              A-13_S3                                 |
|               |               |        |                       |         |        | Po 50        connected (up)  --         10G      --                                                      |
|               |               |        |                       |         |        | Po 64        connected (up)  --         100G     --              MCTPeerInterface                        |
|               |               |        |                       |         |        | Po 121       notconnected    Trunk      --       --              A-EMF1_S3                               |
|               |               |        |                       |         |        | Po 122       notconnected    Trunk      --       --              A-EMF2_S3                               |
|               |               |        |                       |         |        |                                                                                                          |
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 5.670709247s ---
`
	} else if ip == "60.30.181.107" {
		result = `Execute CLI[success]
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------------------+
|  IP Address   |   Host Name   | Fabric |        Command        | Status  | Reason |                                                     Output                                                     |
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------------------+
| 60.30.181.107 | TB2-SLX-SMF-A | TB2    | show interface status | Success |        | TB2-SLX-SMF-A# show interface status                                                                           |
|               |               |        |                       |         |        | Flags:  L - Internal Loopback                                                                                  |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                               |
|               |               |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                                   |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                               |
|               |               |        |                       |         |        | Eth 0/1:1    connected (up)  --         25G      100G            Port-channel S-1_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/1:2    connected (up)  --         25G      100G            Port-channel S-2_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/1:3    connected (up)  --         25G      100G            Port-channel S-1_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/1:4    connected (up)  --         25G      100G            Port-channel S-2_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/2:1    connected (up)  --         25G      100G            Port-channel S-3_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/2:2    connected (up)  --         25G      100G            Port-channel S-4_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/2:3    connected (up)  --         25G      100G            Port-channel S-3_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/2:4    connected (up)  --         25G      100G            Port-channel S-4_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/3:1    connected (up)  --         25G      100G            Port-channel S-5_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/3:2    connected (up)  --         25G      100G            Port-channel S-6_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/3:3    connected (up)  --         25G      100G            Port-channel S-5_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/3:4    connected (up)  --         25G      100G            Port-channel S-6_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/4:1    connected (up)  --         25G      100G            Port-channel S-7_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/4:2    connected (up)  --         25G      100G            Port-channel S-8_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/4:3    connected (up)  --         25G      100G            Port-channel S-7_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/4:4    connected (up)  --         25G      100G            Port-channel S-8_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/5:1    connected (up)  --         25G      100G            Port-channel S-9_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/5:2    connected (up)  --         25G      100G            Port-channel S-CMF1_S1 Member interface       |
|               |               |        |                       |         |        | Eth 0/5:3    connected (up)  --         25G      100G            Port-channel S-9_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/5:4    connected (up)  --         25G      100G            Port-channel S-CMF1_S3 Member interface       |
|               |               |        |                       |         |        | Eth 0/6:1    connected (up)  --         25G      100G            Port-channel S-CMF2_S1 Member interface       |
|               |               |        |                       |         |        | Eth 0/6:2    connected (up)  --         25G      100G            Port-channel S-EMF1_S1 Member interface       |
|               |               |        |                       |         |        | Eth 0/6:3    connected (up)  --         25G      100G            Port-channel S-CMF2_S3 Member interface       |
|               |               |        |                       |         |        | Eth 0/6:4    notconnected    --         --       100G            Port-channel S-EMF1_S3 Member interface       |
|               |               |        |                       |         |        | Eth 0/7:1    connected (up)  --         25G      100G            Port-channel S-EMF2_S1 Member interface       |
|               |               |        |                       |         |        | Eth 0/7:2    connected (up)  --         25G      100G            Port-channel UPF-Data1-14_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/7:3    connected (up)  --         25G      100G            Port-channel S-EMF2_S3 Member interface       |
|               |               |        |                       |         |        | Eth 0/7:4    adminDown       --         --       100G                                                          |
|               |               |        |                       |         |        | Eth 0/8:1    connected (up)  Trunk      25G      100G                                                          |
|               |               |        |                       |         |        | Eth 0/8:2    connected (up)  Trunk      25G      100G                                                          |
|               |               |        |                       |         |        | Eth 0/8:3    adminDown       --         --       100G                                                          |
|               |               |        |                       |         |        | Eth 0/8:4    adminDown       --         --       100G                                                          |
|               |               |        |                       |         |        | Eth 0/9:1    adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/9:2    adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/9:3    adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/9:4    adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/10:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/10:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/10:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/10:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/11:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/11:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/11:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/11:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/12:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/12:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/12:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/12:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/13:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/13:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/13:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/13:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/14:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/14:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/14:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/14:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/15:1   connected (up)  --         25G      100G            Port-channel A-EMF1_S1 Member interface       |
|               |               |        |                       |         |        | Eth 0/15:2   connected (up)  --         25G      100G            Port-channel A-EMF1_S2 Member interface       |
|               |               |        |                       |         |        | Eth 0/15:3   connected (up)  --         25G      100G            Port-channel A-EMF2_S1 Member interface       |
|               |               |        |                       |         |        | Eth 0/15:4   connected (up)  --         25G      100G            Port-channel A-EMF2_S2 Member interface       |
|               |               |        |                       |         |        | Eth 0/16:1   connected (up)  --         25G      100G            Port-channel NMF1_S1 Member interface         |
|               |               |        |                       |         |        | Eth 0/16:2   connected (up)  --         25G      100G            Port-channel NMF2_S1 Member interface         |
|               |               |        |                       |         |        | Eth 0/16:3   adminDown       --         --       100G                                                          |
|               |               |        |                       |         |        | Eth 0/16:4   adminDown       --         --       100G                                                          |
|               |               |        |                       |         |        | Eth 0/17:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/17:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/17:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/17:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/18:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/18:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/18:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/18:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/19:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/19:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/19:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/19:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/20:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/20:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/20:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/20:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/21     connected (up)  --         100G     100G            Link to 60.30.181.101 Spine                   |
|               |               |        |                       |         |        | Eth 0/22     connected (up)  --         100G     100G            Link to 60.30.181.102 Spine                   |
|               |               |        |                       |         |        | Eth 0/23     adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/24     adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/25     adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/26     adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/27     adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/28     adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/29     connected (up)  --         100G     100G            clusterPeerIntfMember                         |
|               |               |        |                       |         |        | Eth 0/30     connected (up)  --         100G     100G            clusterPeerIntfMember                         |
|               |               |        |                       |         |        | Eth 0/31:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/31:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/31:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/31:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/32:1   adminDown       --         --       10G-SFP-SR                                                    |
|               |               |        |                       |         |        | Eth 0/32:2   adminDown       --         --       10G-SFP-SR                                                    |
|               |               |        |                       |         |        | Eth 0/32:3   adminDown       --         --       10G-SFP-SR                                                    |
|               |               |        |                       |         |        | Eth 0/32:4   adminDown       --         --       10G-SFP-SR                                                    |
|               |               |        |                       |         |        | Po 1         connected (up)  Trunk      25G      --              S-1_S1                                        |
|               |               |        |                       |         |        | Po 2         connected (up)  Trunk      25G      --              S-2_S1                                        |
|               |               |        |                       |         |        | Po 3         connected (up)  Trunk      25G      --              S-3_S1                                        |
|               |               |        |                       |         |        | Po 4         connected (up)  Trunk      25G      --              S-4_S1                                        |
|               |               |        |                       |         |        | Po 5         connected (up)  Trunk      25G      --              S-5_S1                                        |
|               |               |        |                       |         |        | Po 6         connected (up)  Trunk      25G      --              S-6_S1                                        |
|               |               |        |                       |         |        | Po 7         connected (up)  Trunk      25G      --              S-7_S1                                        |
|               |               |        |                       |         |        | Po 8         connected (up)  Trunk      25G      --              S-8_S1                                        |
|               |               |        |                       |         |        | Po 9         connected (up)  Trunk      25G      --              S-9_S1                                        |
|               |               |        |                       |         |        | Po 14        connected (up)  Trunk      25G      --              UPF-Data1-14_S3_Control                       |
|               |               |        |                       |         |        | Po 31        notconnected    Trunk      --       --              S-1_S3                                        |
|               |               |        |                       |         |        | Po 32        notconnected    Trunk      --       --              S-2_S3                                        |
|               |               |        |                       |         |        | Po 33        notconnected    Trunk      --       --              S-3_S3                                        |
|               |               |        |                       |         |        | Po 34        notconnected    Trunk      --       --              S-4_S3                                        |
|               |               |        |                       |         |        | Po 35        notconnected    Trunk      --       --              S-5_S3                                        |
|               |               |        |                       |         |        | Po 36        notconnected    Trunk      --       --              S-6_S3                                        |
|               |               |        |                       |         |        | Po 37        notconnected    Trunk      --       --              S-7_S3                                        |
|               |               |        |                       |         |        | Po 38        notconnected    Trunk      --       --              S-8_S3                                        |
|               |               |        |                       |         |        | Po 39        notconnected    Trunk      --       --              S-9_S3                                        |
|               |               |        |                       |         |        | Po 50        connected (up)  --         10G      --                                                            |
|               |               |        |                       |         |        | Po 64        connected (up)  --         200G     --              MCTPeerInterface                              |
|               |               |        |                       |         |        | Po 111       connected (up)  Trunk      25G      --              S-CMF1_S1                                     |
|               |               |        |                       |         |        | Po 112       connected (up)  Trunk      25G      --              S-CMF2_S1                                     |
|               |               |        |                       |         |        | Po 113       connected (up)  Trunk      25G      --              S-EMF1_S1                                     |
|               |               |        |                       |         |        | Po 114       connected (up)  Trunk      25G      --              S-EMF2_S1                                     |
|               |               |        |                       |         |        | Po 121       connected (up)  Trunk      25G      --              S-CMF1_S3                                     |
|               |               |        |                       |         |        | Po 122       connected (up)  Trunk      25G      --              S-CMF2_S3                                     |
|               |               |        |                       |         |        | Po 123       notconnected    Trunk      --       --              S-EMF1_S3                                     |
|               |               |        |                       |         |        | Po 124       connected (up)  Trunk      25G      --              S-EMF2_S3                                     |
|               |               |        |                       |         |        | Po 151       connected (up)  Trunk      25G      --              A-EMF1_S1                                     |
|               |               |        |                       |         |        | Po 152       connected (up)  Trunk      25G      --              A-EMF1_S2                                     |
|               |               |        |                       |         |        | Po 153       connected (up)  Trunk      25G      --              A-EMF2_S1                                     |
|               |               |        |                       |         |        | Po 154       connected (up)  Trunk      25G      --              A-EMF2_S2                                     |
|               |               |        |                       |         |        | Po 155       connected (up)  Trunk      25G      --              NMF1_S1                                       |
|               |               |        |                       |         |        | Po 156       connected (up)  Trunk      25G      --              NMF2_S1                                       |
|               |               |        |                       |         |        |                                                                                                                |
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 5.845217382s ---
`
	} else if ip == "60.30.181.108" {
		result = `Execute CLI[success]
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------------------+
|  IP Address   |   Host Name   | Fabric |        Command        | Status  | Reason |                                                     Output                                                     |
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------------------+
| 60.30.181.108 | TB2-SLX-SMF-B | TB2    | show interface status | Success |        | TB2-SLX-SMF-B# show interface status                                                                           |
|               |               |        |                       |         |        | Flags:  L - Internal Loopback                                                                                  |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                               |
|               |               |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                                   |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                               |
|               |               |        |                       |         |        | Eth 0/1:1    connected (up)  --         25G      100G            Port-channel S-1_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/1:2    connected (up)  --         25G      100G            Port-channel S-2_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/1:3    connected (up)  --         25G      100G            Port-channel S-1_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/1:4    connected (up)  --         25G      100G            Port-channel S-2_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/2:1    connected (up)  --         25G      100G            Port-channel S-3_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/2:2    connected (up)  --         25G      100G            Port-channel S-4_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/2:3    connected (up)  --         25G      100G            Port-channel S-3_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/2:4    connected (up)  --         25G      100G            Port-channel S-4_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/3:1    connected (up)  --         25G      100G            Port-channel S-5_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/3:2    connected (up)  --         25G      100G            Port-channel S-6_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/3:3    connected (up)  --         25G      100G            Port-channel S-5_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/3:4    connected (up)  --         25G      100G            Port-channel S-6_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/4:1    connected (up)  --         25G      100G            Port-channel S-7_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/4:2    connected (up)  --         25G      100G            Port-channel S-8_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/4:3    connected (up)  --         25G      100G            Port-channel S-7_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/4:4    connected (up)  --         25G      100G            Port-channel S-8_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/5:1    connected (up)  --         25G      100G            Port-channel S-9_S1 Member interface          |
|               |               |        |                       |         |        | Eth 0/5:2    connected (up)  --         25G      100G            Port-channel S-CMF1_S1 Member interface       |
|               |               |        |                       |         |        | Eth 0/5:3    connected (up)  --         25G      100G            Port-channel S-9_S3 Member interface          |
|               |               |        |                       |         |        | Eth 0/5:4    connected (up)  --         25G      100G            Port-channel S-CMF1_S3 Member interface       |
|               |               |        |                       |         |        | Eth 0/6:1    connected (up)  --         25G      100G            Port-channel S-CMF2_S1 Member interface       |
|               |               |        |                       |         |        | Eth 0/6:2    connected (up)  --         25G      100G            Port-channel S-EMF1_S1 Member interface       |
|               |               |        |                       |         |        | Eth 0/6:3    connected (up)  --         25G      100G            Port-channel S-CMF2_S3 Member interface       |
|               |               |        |                       |         |        | Eth 0/6:4    connected (up)  --         25G      100G            Port-channel S-EMF1_S3 Member interface       |
|               |               |        |                       |         |        | Eth 0/7:1    connected (up)  --         25G      100G            Port-channel S-EMF2_S1 Member interface       |
|               |               |        |                       |         |        | Eth 0/7:2    connected (up)  --         25G      100G            Port-channel UPF-Data1-14_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/7:3    connected (up)  --         25G      100G            Port-channel S-EMF2_S3 Member interface       |
|               |               |        |                       |         |        | Eth 0/7:4    adminDown       --         --       100G                                                          |
|               |               |        |                       |         |        | Eth 0/8:1    connected (up)  Trunk      25G      100G                                                          |
|               |               |        |                       |         |        | Eth 0/8:2    connected (up)  Trunk      25G      100G                                                          |
|               |               |        |                       |         |        | Eth 0/8:3    adminDown       --         --       100G                                                          |
|               |               |        |                       |         |        | Eth 0/8:4    adminDown       --         --       100G                                                          |
|               |               |        |                       |         |        | Eth 0/9:1    adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/9:2    adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/9:3    adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/9:4    adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/10:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/10:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/10:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/10:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/11:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/11:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/11:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/11:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/12:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/12:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/12:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/12:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/13:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/13:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/13:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/13:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/14:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/14:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/14:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/14:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/15:1   connected (up)  --         25G      100G            Port-channel A-EMF1_S1 Member interface       |
|               |               |        |                       |         |        | Eth 0/15:2   connected (up)  --         25G      100G            Port-channel A-EMF1_S2 Member interface       |
|               |               |        |                       |         |        | Eth 0/15:3   connected (up)  --         25G      100G            Port-channel A-EMF2_S1 Member interface       |
|               |               |        |                       |         |        | Eth 0/15:4   connected (up)  --         25G      100G            Port-channel A-EMF2_S2 Member interface       |
|               |               |        |                       |         |        | Eth 0/16:1   connected (up)  --         25G      100G            Port-channel NMF1_S1 Member interface         |
|               |               |        |                       |         |        | Eth 0/16:2   connected (up)  --         25G      100G            Port-channel NMF2_S1 Member interface         |
|               |               |        |                       |         |        | Eth 0/16:3   adminDown       --         --       100G                                                          |
|               |               |        |                       |         |        | Eth 0/16:4   adminDown       --         --       100G                                                          |
|               |               |        |                       |         |        | Eth 0/17:1   sfpAbsent       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/17:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/17:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/17:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/18:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/18:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/18:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/18:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/19:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/19:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/19:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/19:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/20:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/20:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/20:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/20:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/21     connected (up)  --         100G     100G            Link to 60.30.181.101 Spine                   |
|               |               |        |                       |         |        | Eth 0/22     connected (up)  --         100G     100G            Link to 60.30.181.102 Spine                   |
|               |               |        |                       |         |        | Eth 0/23     adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/24     adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/25     adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/26     adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/27     adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/28     adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/29     connected (up)  --         100G     100G            clusterPeerIntfMember                         |
|               |               |        |                       |         |        | Eth 0/30     connected (up)  --         100G     100G            clusterPeerIntfMember                         |
|               |               |        |                       |         |        | Eth 0/31:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/31:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/31:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/31:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/32:1   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/32:2   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/32:3   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Eth 0/32:4   adminDown       --         --       --                                                            |
|               |               |        |                       |         |        | Po 1         connected (up)  Trunk      25G      --              S-1_S1                                        |
|               |               |        |                       |         |        | Po 2         connected (up)  Trunk      25G      --              S-2_S1                                        |
|               |               |        |                       |         |        | Po 3         connected (up)  Trunk      25G      --              S-3_S1                                        |
|               |               |        |                       |         |        | Po 4         connected (up)  Trunk      25G      --              S-4_S1                                        |
|               |               |        |                       |         |        | Po 5         connected (up)  Trunk      25G      --              S-5_S1                                        |
|               |               |        |                       |         |        | Po 6         connected (up)  Trunk      25G      --              S-6_S1                                        |
|               |               |        |                       |         |        | Po 7         connected (up)  Trunk      25G      --              S-7_S1                                        |
|               |               |        |                       |         |        | Po 8         connected (up)  Trunk      25G      --              S-8_S1                                        |
|               |               |        |                       |         |        | Po 9         connected (up)  Trunk      25G      --              S-9_S1                                        |
|               |               |        |                       |         |        | Po 14        connected (up)  Trunk      25G      --              UPF-Data1-14_S3_Control                       |
|               |               |        |                       |         |        | Po 31        notconnected    Trunk      --       --              S-1_S3                                        |
|               |               |        |                       |         |        | Po 32        notconnected    Trunk      --       --              S-2_S3                                        |
|               |               |        |                       |         |        | Po 33        notconnected    Trunk      --       --              S-3_S3                                        |
|               |               |        |                       |         |        | Po 34        notconnected    Trunk      --       --              S-4_S3                                        |
|               |               |        |                       |         |        | Po 35        notconnected    Trunk      --       --              S-5_S3                                        |
|               |               |        |                       |         |        | Po 36        notconnected    Trunk      --       --              S-6_S3                                        |
|               |               |        |                       |         |        | Po 37        notconnected    Trunk      --       --              S-7_S3                                        |
|               |               |        |                       |         |        | Po 38        notconnected    Trunk      --       --              S-8_S3                                        |
|               |               |        |                       |         |        | Po 39        notconnected    Trunk      --       --              S-9_S3                                        |
|               |               |        |                       |         |        | Po 64        connected (up)  --         200G     --              MCTPeerInterface                              |
|               |               |        |                       |         |        | Po 111       connected (up)  Trunk      25G      --              S-CMF1_S1                                     |
|               |               |        |                       |         |        | Po 112       connected (up)  Trunk      25G      --              S-CMF2_S1                                     |
|               |               |        |                       |         |        | Po 113       connected (up)  Trunk      25G      --              S-EMF1_S1                                     |
|               |               |        |                       |         |        | Po 114       connected (up)  Trunk      25G      --              S-EMF2_S1                                     |
|               |               |        |                       |         |        | Po 121       connected (up)  Trunk      25G      --              S-CMF1_S3                                     |
|               |               |        |                       |         |        | Po 122       connected (up)  Trunk      25G      --              S-CMF2_S3                                     |
|               |               |        |                       |         |        | Po 123       connected (up)  Trunk      25G      --              S-EMF1_S3                                     |
|               |               |        |                       |         |        | Po 124       connected (up)  Trunk      25G      --              S-EMF2_S3                                     |
|               |               |        |                       |         |        | Po 151       connected (up)  Trunk      25G      --              A-EMF1_S1                                     |
|               |               |        |                       |         |        | Po 152       connected (up)  Trunk      25G      --              A-EMF1_S2                                     |
|               |               |        |                       |         |        | Po 153       connected (up)  Trunk      25G      --              A-EMF2_S1                                     |
|               |               |        |                       |         |        | Po 154       connected (up)  Trunk      25G      --              A-EMF2_S2                                     |
|               |               |        |                       |         |        | Po 155       connected (up)  Trunk      25G      --              NMF1_S1                                       |
|               |               |        |                       |         |        | Po 156       connected (up)  Trunk      25G      --              NMF2_S1                                       |
|               |               |        |                       |         |        |                                                                                                                |
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 5.838427108s ---
`
	} else if ip == "60.30.181.109" {
		result = `Execute CLI[success]
+---------------+---------------+--------+-----------------------+---------+--------+------------------------------------------------------------------------------------------------------------------+
|  IP Address   |   Host Name   | Fabric |        Command        | Status  | Reason |                                                      Output                                                      |
+---------------+---------------+--------+-----------------------+---------+--------+------------------------------------------------------------------------------------------------------------------+
| 60.30.181.109 | TB2-SLX-UPF-A | TB2    | show interface status | Success |        | TB2-SLX-UPF-A# show interface status                                                                             |
|               |               |        |                       |         |        | Flags:  L - Internal Loopback                                                                                    |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                                 |
|               |               |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                                     |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                                 |
|               |               |        |                       |         |        | Eth 0/1:1    connected (up)  --         25G      100G            Port-channel UPF-Data1-1_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/1:2    connected (up)  --         25G      100G            Port-channel UPF-Data1-2_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/1:3    notconnected    --         --       100G            Port-channel UPF-Data1-1_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/1:4    notconnected    --         --       100G            Port-channel UPF-Data1-2_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/2:1    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/2:2    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/2:3    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/2:4    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/3:1    connected (up)  --         25G      100G            Port-channel UPF-Data1-3_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:2    connected (up)  --         25G      100G            Port-channel UPF-Data1-4_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:3    notconnected    --         --       100G            Port-channel UPF-Data1-3_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:4    notconnected    --         --       100G            Port-channel UPF-Data1-4_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/4:1    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/4:2    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/4:3    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/4:4    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/5:1    connected (up)  --         25G      100G            Port-channel UPF-Data1-5_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/5:2    connected (up)  --         25G      100G            Port-channel UPF-Data1-6_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/5:3    notconnected    --         --       100G            Port-channel UPF-Data1-5_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/5:4    notconnected    --         --       100G            Port-channel UPF-Data1-6_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/6:1    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/6:2    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/6:3    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/6:4    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/7:1    connected (up)  --         25G      100G            Port-channel UPF-Data1-7_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/7:2    connected (up)  --         25G      100G            Port-channel UPF-Data1-8_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/7:3    notconnected    --         --       100G            Port-channel UPF-Data1-7_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/7:4    notconnected    --         --       100G            Port-channel UPF-Data1-8_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/8:1    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/8:2    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/8:3    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/8:4    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/9:1    connected (up)  --         25G      100G            Port-channel UPF-Data1-9_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/9:2    connected (up)  --         25G      100G            Port-channel UPF-Data1-10_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/9:3    notconnected    --         --       100G            Port-channel UPF-Data1-9_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/9:4    notconnected    --         --       100G            Port-channel UPF-Data1-10_S7 Member interface   |
|               |               |        |                       |         |        | Eth 0/10:1   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/10:2   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/10:3   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/10:4   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/11:1   connected (up)  --         25G      100G            Port-channel UPF-Data1-11_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/11:2   connected (up)  --         25G      100G            Port-channel UPF-Data1-12_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/11:3   connected (up)  --         25G      100G            Port-channel UPF-Data1-11_S7 Member interface   |
|               |               |        |                       |         |        | Eth 0/11:4   connected (up)  --         25G      100G            Port-channel UPF-Data1-12_S7 Member interface   |
|               |               |        |                       |         |        | Eth 0/12:1   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/12:2   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/12:3   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/12:4   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/13:1   connected (up)  --         25G      100G            Port-channel UPF-Data1-13_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/13:2   adminDown       --         --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/13:3   connected (up)  --         25G      100G            Port-channel UPF-Data1-13_S7 Member interface   |
|               |               |        |                       |         |        | Eth 0/13:4   adminDown       --         --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/14:1   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/14:2   notconnected    Trunk      --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/14:3   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/14:4   notconnected    Trunk      --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/15:1   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/15:2   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/15:3   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/15:4   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/16:1   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/16:2   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/16:3   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/16:4   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/17:1   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/17:2   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/17:3   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/17:4   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/18:1   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/18:2   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/18:3   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/18:4   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/19:1   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/19:2   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/19:3   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/19:4   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/20:1   connected (up)  --         25G      100G            Port-channel UPF-Data1-EMF1_S1 Member interface |
|               |               |        |                       |         |        | Eth 0/20:2   connected (up)  --         25G      100G            Port-channel UPF-Data1-EMF1_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/20:3   connected (up)  --         25G      100G            Port-channel UPF-Data1-EMF2_S1 Member interface |
|               |               |        |                       |         |        | Eth 0/20:4   connected (up)  --         25G      100G            Port-channel UPF-Data1-EMF2_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/21     connected (up)  --         100G     100G            Link to 60.30.181.101 Spine                     |
|               |               |        |                       |         |        | Eth 0/22     connected (up)  --         100G     100G            Link to 60.30.181.102 Spine                     |
|               |               |        |                       |         |        | Eth 0/23     connected (up)  --         100G     100G            Link to 60.30.181.101 Spine                     |
|               |               |        |                       |         |        | Eth 0/24     connected (up)  --         100G     100G            Link to 60.30.181.102 Spine                     |
|               |               |        |                       |         |        | Eth 0/25     adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/26     adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/27:1   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/27:2   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/27:3   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/27:4   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/28     adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/29     adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/30     connected (up)  --         100G     100G            clusterPeerIntfMember                           |
|               |               |        |                       |         |        | Eth 0/31     connected (up)  --         100G     100G            clusterPeerIntfMember                           |
|               |               |        |                       |         |        | Eth 0/32:1   connected (up)  --         1G       1G-SFP-T                                                        |
|               |               |        |                       |         |        | Eth 0/32:2   adminDown       --         --       1G-SFP-T                                                        |
|               |               |        |                       |         |        | Eth 0/32:3   adminDown       --         --       1G-SFP-T                                                        |
|               |               |        |                       |         |        | Eth 0/32:4   adminDown       --         --       1G-SFP-T                                                        |
|               |               |        |                       |         |        | Po 1         connected (up)  Trunk      25G      --              UPF-Data1-1_S3_Control                          |
|               |               |        |                       |         |        | Po 2         connected (up)  Trunk      25G      --              UPF-Data1-2_S3_Control                          |
|               |               |        |                       |         |        | Po 3         connected (up)  Trunk      25G      --              UPF-Data1-3_S3_Control                          |
|               |               |        |                       |         |        | Po 4         connected (up)  Trunk      25G      --              UPF-Data1-4_S3_Control                          |
|               |               |        |                       |         |        | Po 5         connected (up)  Trunk      25G      --              UPF-Data1-5_S3_Control                          |
|               |               |        |                       |         |        | Po 6         connected (up)  Trunk      25G      --              UPF-Data1-6_S3_Control                          |
|               |               |        |                       |         |        | Po 7         connected (up)  Trunk      25G      --              UPF-Data1-7_S3_Control                          |
|               |               |        |                       |         |        | Po 8         connected (up)  Trunk      25G      --              UPF-Data1-8_S3_Control                          |
|               |               |        |                       |         |        | Po 9         connected (up)  Trunk      25G      --              UPF-Data1-9_S3_Control                          |
|               |               |        |                       |         |        | Po 10        connected (up)  Trunk      25G      --              UPF-Data1-10_S3_Control                         |
|               |               |        |                       |         |        | Po 11        connected (up)  Trunk      25G      --              UPF-Data1-11_S3_Control                         |
|               |               |        |                       |         |        | Po 12        connected (up)  Trunk      25G      --              UPF-Data1-12_S3_Control                         |
|               |               |        |                       |         |        | Po 13        connected (up)  Trunk      25G      --              UPF-Data1-13_S3_Control                         |
|               |               |        |                       |         |        | Po 31        notconnected    Trunk      --       --              UPF-Data1-1_S7_Control                          |
|               |               |        |                       |         |        | Po 32        notconnected    Trunk      --       --              UPF-Data1-2_S7_Control                          |
|               |               |        |                       |         |        | Po 33        notconnected    Trunk      --       --              UPF-Data1-3_S7_Control                          |
|               |               |        |                       |         |        | Po 34        notconnected    Trunk      --       --              UPF-Data1-4_S7_Control                          |
|               |               |        |                       |         |        | Po 35        notconnected    Trunk      --       --              UPF-Data1-5_S7_Control                          |
|               |               |        |                       |         |        | Po 36        notconnected    Trunk      --       --              UPF-Data1-6_S7_Control                          |
|               |               |        |                       |         |        | Po 37        notconnected    Trunk      --       --              UPF-Data1-7_S7_Control                          |
|               |               |        |                       |         |        | Po 38        notconnected    Trunk      --       --              UPF-Data1-8_S7_Control                          |
|               |               |        |                       |         |        | Po 39        notconnected    Trunk      --       --              UPF-Data1-9_S7_Control                          |
|               |               |        |                       |         |        | Po 40        notconnected    Trunk      --       --              UPF-Data1-10_S7_Control                         |
|               |               |        |                       |         |        | Po 41        notconnected    Trunk      --       --              UPF-Data1-11_S7_Control                         |
|               |               |        |                       |         |        | Po 42        notconnected    Trunk      --       --              UPF-Data1-12_S7_Control                         |
|               |               |        |                       |         |        | Po 43        notconnected    Trunk      --       --              UPF-Data1-13_S7_Control                         |
|               |               |        |                       |         |        | Po 64        connected (up)  --         200G     --              MCTPeerInterface                                |
|               |               |        |                       |         |        | Po 111       connected (up)  Trunk      25G      --              UPF-Data1-EMF1_S1                               |
|               |               |        |                       |         |        | Po 112       connected (up)  Trunk      25G      --              UPF-Data1-EMF2_S1                               |
|               |               |        |                       |         |        | Po 121       connected (up)  Trunk      25G      --              UPF-Data1-EMF1_S3                               |
|               |               |        |                       |         |        | Po 124       connected (up)  Trunk      25G      --              UPF-Data1-EMF2_S3                               |
|               |               |        |                       |         |        |                                                                                                                  |
+---------------+---------------+--------+-----------------------+---------+--------+------------------------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 5.414954481s ---
`
	} else if ip == "60.30.181.110" {
		result = `Execute CLI[success]
+---------------+---------------+--------+-----------------------+---------+--------+------------------------------------------------------------------------------------------------------------------+
|  IP Address   |   Host Name   | Fabric |        Command        | Status  | Reason |                                                      Output                                                      |
+---------------+---------------+--------+-----------------------+---------+--------+------------------------------------------------------------------------------------------------------------------+
| 60.30.181.110 | TB2-SLX-UPF-B | TB2    | show interface status | Success |        | TB2-SLX-UPF-B# show interface status                                                                             |
|               |               |        |                       |         |        | Flags:  L - Internal Loopback                                                                                    |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                                 |
|               |               |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                                     |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                                 |
|               |               |        |                       |         |        | Eth 0/1:1    connected (up)  --         25G      100G            Port-channel UPF-Data1-1_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/1:2    connected (up)  --         25G      100G            Port-channel UPF-Data1-2_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/1:3    notconnected    --         --       100G            Port-channel UPF-Data1-1_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/1:4    notconnected    --         --       100G            Port-channel UPF-Data1-2_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/2:1    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/2:2    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/2:3    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/2:4    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/3:1    connected (up)  --         25G      100G            Port-channel UPF-Data1-3_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:2    connected (up)  --         25G      100G            Port-channel UPF-Data1-4_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:3    notconnected    --         --       100G            Port-channel UPF-Data1-3_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/3:4    notconnected    --         --       100G            Port-channel UPF-Data1-4_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/4:1    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/4:2    notconnected    Trunk      --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/4:3    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/4:4    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/5:1    connected (up)  --         25G      100G            Port-channel UPF-Data1-5_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/5:2    connected (up)  --         25G      100G            Port-channel UPF-Data1-6_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/5:3    notconnected    --         --       100G            Port-channel UPF-Data1-5_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/5:4    notconnected    --         --       100G            Port-channel UPF-Data1-6_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/6:1    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/6:2    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/6:3    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/6:4    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/7:1    connected (up)  --         25G      100G            Port-channel UPF-Data1-7_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/7:2    connected (up)  --         25G      100G            Port-channel UPF-Data1-8_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/7:3    notconnected    --         --       100G            Port-channel UPF-Data1-7_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/7:4    notconnected    --         --       100G            Port-channel UPF-Data1-8_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/8:1    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/8:2    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/8:3    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/8:4    connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/9:1    connected (up)  --         25G      100G            Port-channel UPF-Data1-9_S3 Member interface    |
|               |               |        |                       |         |        | Eth 0/9:2    connected (up)  --         25G      100G            Port-channel UPF-Data1-10_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/9:3    notconnected    --         --       100G            Port-channel UPF-Data1-9_S7 Member interface    |
|               |               |        |                       |         |        | Eth 0/9:4    notconnected    --         --       100G            Port-channel UPF-Data1-10_S7 Member interface   |
|               |               |        |                       |         |        | Eth 0/10:1   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/10:2   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/10:3   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/10:4   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/11:1   connected (up)  --         25G      100G            Port-channel UPF-Data1-11_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/11:2   connected (up)  --         25G      100G            Port-channel UPF-Data1-12_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/11:3   connected (up)  --         25G      100G            Port-channel UPF-Data1-11_S7 Member interface   |
|               |               |        |                       |         |        | Eth 0/11:4   connected (up)  --         25G      100G            Port-channel UPF-Data1-12_S7 Member interface   |
|               |               |        |                       |         |        | Eth 0/12:1   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/12:2   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/12:3   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/12:4   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/13:1   connected (up)  --         25G      100G            Port-channel UPF-Data1-13_S3 Member interface   |
|               |               |        |                       |         |        | Eth 0/13:2   adminDown       --         --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/13:3   connected (up)  --         25G      100G            Port-channel UPF-Data1-13_S7 Member interface   |
|               |               |        |                       |         |        | Eth 0/13:4   adminDown       --         --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/14:1   connected (up)  Trunk      25G      100G                                                            |
|               |               |        |                       |         |        | Eth 0/14:2   notconnected    Trunk      --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/14:3   notconnected    Trunk      --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/14:4   notconnected    Trunk      --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/15:1   adminDown       --         --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/15:2   adminDown       --         --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/15:3   adminDown       --         --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/15:4   adminDown       --         --       100G                                                            |
|               |               |        |                       |         |        | Eth 0/16:1   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/16:2   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/16:3   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/16:4   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/17:1   sfpAbsent       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/17:2   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/17:3   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/17:4   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/18:1   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/18:2   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/18:3   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/18:4   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/19:1   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/19:2   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/19:3   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/19:4   adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/20:1   connected (up)  --         25G      100G            Port-channel UPF-Data1-EMF1_S1 Member interface |
|               |               |        |                       |         |        | Eth 0/20:2   notconnected    --         --       100G            Port-channel UPF-Data1-EMF1_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/20:3   connected (up)  --         25G      100G            Port-channel UPF-Data1-EMF2_S1 Member interface |
|               |               |        |                       |         |        | Eth 0/20:4   connected (up)  --         25G      100G            Port-channel UPF-Data1-EMF2_S3 Member interface |
|               |               |        |                       |         |        | Eth 0/21     connected (up)  --         100G     100G            Link to 60.30.181.101 Spine                     |
|               |               |        |                       |         |        | Eth 0/22     connected (up)  --         100G     100G            Link to 60.30.181.102 Spine                     |
|               |               |        |                       |         |        | Eth 0/23     connected (up)  --         100G     100G            Link to 60.30.181.101 Spine                     |
|               |               |        |                       |         |        | Eth 0/24     connected (up)  --         100G     100G            Link to 60.30.181.102 Spine                     |
|               |               |        |                       |         |        | Eth 0/25     adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/26     adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/27     adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/28     adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/29     adminDown       --         --       --                                                              |
|               |               |        |                       |         |        | Eth 0/30     connected (up)  --         100G     100G            clusterPeerIntfMember                           |
|               |               |        |                       |         |        | Eth 0/31     connected (up)  --         100G     100G            clusterPeerIntfMember                           |
|               |               |        |                       |         |        | Eth 0/32:1   connected (up)  --         1G       1G-SFP-T                                                        |
|               |               |        |                       |         |        | Eth 0/32:2   adminDown       --         --       1G-SFP-T                                                        |
|               |               |        |                       |         |        | Eth 0/32:3   adminDown       --         --       1G-SFP-T                                                        |
|               |               |        |                       |         |        | Eth 0/32:4   adminDown       --         --       1G-SFP-T                                                        |
|               |               |        |                       |         |        | Po 1         connected (up)  Trunk      25G      --              UPF-Data1-1_S3_Control                          |
|               |               |        |                       |         |        | Po 2         connected (up)  Trunk      25G      --              UPF-Data1-2_S3_Control                          |
|               |               |        |                       |         |        | Po 3         connected (up)  Trunk      25G      --              UPF-Data1-3_S3_Control                          |
|               |               |        |                       |         |        | Po 4         connected (up)  Trunk      25G      --              UPF-Data1-4_S3_Control                          |
|               |               |        |                       |         |        | Po 5         connected (up)  Trunk      25G      --              UPF-Data1-5_S3_Control                          |
|               |               |        |                       |         |        | Po 6         connected (up)  Trunk      25G      --              UPF-Data1-6_S3_Control                          |
|               |               |        |                       |         |        | Po 7         connected (up)  Trunk      25G      --              UPF-Data1-7_S3_Control                          |
|               |               |        |                       |         |        | Po 8         connected (up)  Trunk      25G      --              UPF-Data1-8_S3_Control                          |
|               |               |        |                       |         |        | Po 9         connected (up)  Trunk      25G      --              UPF-Data1-9_S3_Control                          |
|               |               |        |                       |         |        | Po 10        connected (up)  Trunk      25G      --              UPF-Data1-10_S3_Control                         |
|               |               |        |                       |         |        | Po 11        connected (up)  Trunk      25G      --              UPF-Data1-11_S3_Control                         |
|               |               |        |                       |         |        | Po 12        connected (up)  Trunk      25G      --              UPF-Data1-12_S3_Control                         |
|               |               |        |                       |         |        | Po 13        connected (up)  Trunk      25G      --              UPF-Data1-13_S3_Control                         |
|               |               |        |                       |         |        | Po 31        notconnected    Trunk      --       --              UPF-Data1-1_S7_Control                          |
|               |               |        |                       |         |        | Po 32        notconnected    Trunk      --       --              UPF-Data1-2_S7_Control                          |
|               |               |        |                       |         |        | Po 33        notconnected    Trunk      --       --              UPF-Data1-3_S7_Control                          |
|               |               |        |                       |         |        | Po 34        notconnected    Trunk      --       --              UPF-Data1-4_S7_Control                          |
|               |               |        |                       |         |        | Po 35        notconnected    Trunk      --       --              UPF-Data1-5_S7_Control                          |
|               |               |        |                       |         |        | Po 36        notconnected    Trunk      --       --              UPF-Data1-6_S7_Control                          |
|               |               |        |                       |         |        | Po 37        notconnected    Trunk      --       --              UPF-Data1-7_S7_Control                          |
|               |               |        |                       |         |        | Po 38        notconnected    Trunk      --       --              UPF-Data1-8_S7_Control                          |
|               |               |        |                       |         |        | Po 39        notconnected    Trunk      --       --              UPF-Data1-9_S7_Control                          |
|               |               |        |                       |         |        | Po 40        notconnected    Trunk      --       --              UPF-Data1-10_S7_Control                         |
|               |               |        |                       |         |        | Po 41        notconnected    Trunk      --       --              UPF-Data1-11_S7_Control                         |
|               |               |        |                       |         |        | Po 42        notconnected    Trunk      --       --              UPF-Data1-12_S7_Control                         |
|               |               |        |                       |         |        | Po 43        notconnected    Trunk      --       --              UPF-Data1-13_S7_Control                         |
|               |               |        |                       |         |        | Po 64        connected (up)  --         200G     --              MCTPeerInterface                                |
|               |               |        |                       |         |        | Po 111       connected (up)  Trunk      25G      --              UPF-Data1-EMF1_S1                               |
|               |               |        |                       |         |        | Po 112       connected (up)  Trunk      25G      --              UPF-Data1-EMF2_S1                               |
|               |               |        |                       |         |        | Po 121       notconnected    Trunk      --       --              UPF-Data1-EMF1_S3                               |
|               |               |        |                       |         |        | Po 124       connected (up)  Trunk      25G      --              UPF-Data1-EMF2_S3                               |
|               |               |        |                       |         |        |                                                                                                                  |
+---------------+---------------+--------+-----------------------+---------+--------+------------------------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 7.047448854s ---
`
	} else if ip == "60.30.181.111" {
		result = `Execute CLI[success]
+---------------+---------------+--------+-----------------------+---------+--------+-----------------------------------------------------------------------------------------------------------+
|  IP Address   |   Host Name   | Fabric |        Command        | Status  | Reason |                                                  Output                                                   |
+---------------+---------------+--------+-----------------------+---------+--------+-----------------------------------------------------------------------------------------------------------+
| 60.30.181.111 | TB2-SLX-DBL-A | TB2    | show interface status | Success |        | TB2-SLX-DBL-A# show interface status                                                                      |
|               |               |        |                       |         |        | Flags:  L - Internal Loopback                                                                             |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                          |
|               |               |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                              |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                          |
|               |               |        |                       |         |        | Eth 0/1      connected (up)  --         100G     100G            Link to 60.30.181.101 Spine              |
|               |               |        |                       |         |        | Eth 0/2      connected (up)  --         100G     100G            Link to 60.30.181.102 Spine              |
|               |               |        |                       |         |        | Eth 0/3      adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/4      adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/5      adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/6      adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/7      adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/8      adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/9      connected (up)  --         100G     100G            Port-channel DBL1_NAT01 Member interface |
|               |               |        |                       |         |        | Eth 0/10     connected (up)  --         100G     100G            Port-channel DBL1_NAT01 Member interface |
|               |               |        |                       |         |        | Eth 0/11     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/12     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/13     connected (up)  --         100G     100G            Port-channel DBL1_NAT02 Member interface |
|               |               |        |                       |         |        | Eth 0/14     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/15     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/16     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/17     connected (up)  Trunk      100G     100G            PUB-DCGW-01-100GE4/0/6                   |
|               |               |        |                       |         |        | Eth 0/18     connected (up)  Trunk      100G     100G            PUB-DCGW-02-100GE4/0/6                   |
|               |               |        |                       |         |        | Eth 0/19     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/20     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/21     connected (up)  Trunk      100G     100G            Sim_X870_A                               |
|               |               |        |                       |         |        | Eth 0/22     connected (up)  Trunk      100G     100G            Sim_X870_A                               |
|               |               |        |                       |         |        | Eth 0/23:1   adminDown       --         --       1G-SFP-T                                                 |
|               |               |        |                       |         |        | Eth 0/23:2   adminDown       --         --       1G-SFP-T                                                 |
|               |               |        |                       |         |        | Eth 0/23:3   adminDown       --         --       1G-SFP-T                                                 |
|               |               |        |                       |         |        | Eth 0/23:4   adminDown       --         --       1G-SFP-T                                                 |
|               |               |        |                       |         |        | Eth 0/25     connected (up)  --         100G     100G            Port-channel DBL1_NATBK Member interface |
|               |               |        |                       |         |        | Eth 0/26     sfpAbsent       --         --       --              Port-channel DBL1_NATBK Member interface |
|               |               |        |                       |         |        | Eth 0/27:1   connected (up)  Trunk      25G      100G                                                     |
|               |               |        |                       |         |        | Eth 0/27:2   connected (up)  --         25G      100G                                                     |
|               |               |        |                       |         |        | Eth 0/27:3   connected (up)  Trunk      25G      100G                                                     |
|               |               |        |                       |         |        | Eth 0/27:4   adminDown       --         --       100G                                                     |
|               |               |        |                       |         |        | Eth 0/29     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/30     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/31     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/32     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/33     sfpAbsent       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/34     sfpAbsent       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/35     sfpAbsent       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/36     connected (up)  --         100G     100G            clusterPeerIntfMember                    |
|               |               |        |                       |         |        | Eth 0/37     connected (up)  --         100G     100G            clusterPeerIntfMember                    |
|               |               |        |                       |         |        | Eth 0/38     connected (up)  --         100G     100G            clusterPeerIntfMember                    |
|               |               |        |                       |         |        | Eth 0/39:1   notconnected    --         --       1G-SFP-T                                                 |
|               |               |        |                       |         |        | Eth 0/39:2   adminDown       --         --       1G-SFP-T                                                 |
|               |               |        |                       |         |        | Eth 0/39:3   adminDown       --         --       1G-SFP-T                                                 |
|               |               |        |                       |         |        | Eth 0/39:4   adminDown       --         --       1G-SFP-T                                                 |
|               |               |        |                       |         |        | Po 10        connected (up)  Trunk      200G     --              DBL1_NAT01                               |
|               |               |        |                       |         |        | Po 11        connected (up)  Trunk      100G     --              DBL1_NATBK                               |
|               |               |        |                       |         |        | Po 12        connected (up)  --         100G     --              DBL1_NAT02                               |
|               |               |        |                       |         |        | Po 64        connected (up)  --         300G     --              MCTPeerInterface                         |
|               |               |        |                       |         |        |                                                                                                           |
+---------------+---------------+--------+-----------------------+---------+--------+-----------------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 5.933116312s ---
`
	} else if ip == "60.30.181.112" {
		result = `Execute CLI[success]
+---------------+---------------+--------+-----------------------+---------+--------+-----------------------------------------------------------------------------------------------------------+
|  IP Address   |   Host Name   | Fabric |        Command        | Status  | Reason |                                                  Output                                                   |
+---------------+---------------+--------+-----------------------+---------+--------+-----------------------------------------------------------------------------------------------------------+
| 60.30.181.112 | TB2-SLX-DBL-B | TB2    | show interface status | Success |        | TB2-SLX-DBL-B# show interface status                                                                      |
|               |               |        |                       |         |        | Flags:  L - Internal Loopback                                                                             |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                          |
|               |               |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                              |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                          |
|               |               |        |                       |         |        | Eth 0/1      connected (up)  --         100G     100G            Link to 60.30.181.101 Spine              |
|               |               |        |                       |         |        | Eth 0/2      connected (up)  --         100G     100G            Link to 60.30.181.102 Spine              |
|               |               |        |                       |         |        | Eth 0/3      sfpAbsent       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/4      adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/5      adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/6      adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/7      adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/8      adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/9      connected (up)  --         100G     100G            Port-channel DBL2_NAT01 Member interface |
|               |               |        |                       |         |        | Eth 0/10     connected (up)  --         100G     100G            Port-channel DBL2_NAT01 Member interface |
|               |               |        |                       |         |        | Eth 0/11     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/12     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/13     connected (up)  --         100G     100G            Port-channel DBL2_NAT02 Member interface |
|               |               |        |                       |         |        | Eth 0/14     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/15     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/16     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/17     connected (up)  Trunk      100G     100G            PUB-DCGW-01-100GE4/0/7                   |
|               |               |        |                       |         |        | Eth 0/18     connected (up)  Trunk      100G     100G            PUB-DCGW-02-100GE4/0/7                   |
|               |               |        |                       |         |        | Eth 0/19     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/20     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/21     connected (up)  Trunk      100G     100G                                                     |
|               |               |        |                       |         |        | Eth 0/22     connected (up)  Trunk      100G     100G                                                     |
|               |               |        |                       |         |        | Eth 0/23     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/24     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/25     connected (up)  --         100G     100G            Port-channel DBL2_NATBK Member interface |
|               |               |        |                       |         |        | Eth 0/26     connected (up)  --         100G     100G            Port-channel DBL2_NATBK Member interface |
|               |               |        |                       |         |        | Eth 0/27:1   notconnected    Trunk      --       100G                                                     |
|               |               |        |                       |         |        | Eth 0/27:2   connected (up)  --         25G      100G                                                     |
|               |               |        |                       |         |        | Eth 0/27:3   adminDown       --         --       100G                                                     |
|               |               |        |                       |         |        | Eth 0/27:4   adminDown       --         --       100G                                                     |
|               |               |        |                       |         |        | Eth 0/29     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/30     adminDown       --         --       100G                                                     |
|               |               |        |                       |         |        | Eth 0/31     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/32     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/33     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/34     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/35     adminDown       --         --       --                                                       |
|               |               |        |                       |         |        | Eth 0/36     connected (up)  --         100G     100G            clusterPeerIntfMember                    |
|               |               |        |                       |         |        | Eth 0/37     connected (up)  --         100G     100G            clusterPeerIntfMember                    |
|               |               |        |                       |         |        | Eth 0/38     connected (up)  --         100G     100G            clusterPeerIntfMember                    |
|               |               |        |                       |         |        | Eth 0/39:1   adminDown       --         --       1G-SFP-T                                                 |
|               |               |        |                       |         |        | Eth 0/39:2   adminDown       --         --       1G-SFP-T                                                 |
|               |               |        |                       |         |        | Eth 0/39:3   adminDown       --         --       1G-SFP-T                                                 |
|               |               |        |                       |         |        | Eth 0/39:4   adminDown       --         --       1G-SFP-T                                                 |
|               |               |        |                       |         |        | Po 10        connected (up)  Trunk      200G     --              DBL2_NAT01                               |
|               |               |        |                       |         |        | Po 11        connected (up)  Trunk      200G     --              DBL2_NATBK                               |
|               |               |        |                       |         |        | Po 12        connected (up)  --         100G     --              DBL2_NAT02                               |
|               |               |        |                       |         |        | Po 64        connected (up)  --         300G     --              MCTPeerInterface                         |
|               |               |        |                       |         |        |                                                                                                           |
+---------------+---------------+--------+-----------------------+---------+--------+-----------------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 5.766672535s ---
`
	} else if ip == "60.30.181.113" {
		result = `Execute CLI[success]
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------+
|  IP Address   |   Host Name   | Fabric |        Command        | Status  | Reason |                                               Output                                               |
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------+
| 60.30.181.113 | TB2-SLX-MBL-A | TB2    | show interface status | Success |        | TB2-SLX-MBL-A# show interface status                                                               |
|               |               |        |                       |         |        | Flags:  L - Internal Loopback                                                                      |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                   |
|               |               |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                       |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                   |
|               |               |        |                       |         |        | Eth 0/1      connected (up)  --         100G     100G            Link to 60.30.181.101 Spine       |
|               |               |        |                       |         |        | Eth 0/2      connected (up)  --         100G     100G            Link to 60.30.181.102 Spine       |
|               |               |        |                       |         |        | Eth 0/3      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/4      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/5      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/6      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/7      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/8      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/9      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/10     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/11     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/12     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/13     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/14     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/15     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/16     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/17:1   connected (up)  Access     25G      100G                                              |
|               |               |        |                       |         |        | Eth 0/17:2   connected (up)  Access     25G      100G                                              |
|               |               |        |                       |         |        | Eth 0/17:3   adminDown       --         --       100G                                              |
|               |               |        |                       |         |        | Eth 0/17:4   adminDown       --         --       100G                                              |
|               |               |        |                       |         |        | Eth 0/18:1   connected (up)  --         10G      25G-SFP         Port-channel CAS Member interface |
|               |               |        |                       |         |        | Eth 0/18:2   adminDown       --         --       25G-SFP                                           |
|               |               |        |                       |         |        | Eth 0/18:3   adminDown       --         --       25G-SFP                                           |
|               |               |        |                       |         |        | Eth 0/18:4   adminDown       --         --       25G-SFP                                           |
|               |               |        |                       |         |        | Eth 0/19     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/20     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/21     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/22     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/23     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/24     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/25     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/26     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/27     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/28     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/29     connected (up)  --         100G     100G            clusterPeerIntfMember             |
|               |               |        |                       |         |        | Eth 0/30     connected (up)  --         100G     100G            clusterPeerIntfMember             |
|               |               |        |                       |         |        | Eth 0/31     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/32:1   adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/32:2   adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/32:3   adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/32:4   adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Po 1         connected (up)  Trunk      10G      --              CAS                               |
|               |               |        |                       |         |        | Po 64        connected (up)  --         200G     --              MCTPeerInterface                  |
|               |               |        |                       |         |        |                                                                                                    |
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 5.549038541s ---
`
	} else if ip == "60.30.181.114" {
		result = `Execute CLI[success]
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------+
|  IP Address   |   Host Name   | Fabric |        Command        | Status  | Reason |                                               Output                                               |
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------+
| 60.30.181.114 | TB2-SLX-MBL-B | TB2    | show interface status | Success |        | TB2-SLX-MBL-B# show interface status                                                               |
|               |               |        |                       |         |        | Flags:  L - Internal Loopback                                                                      |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                   |
|               |               |        |                       |         |        | Port         Status          Mode       Speed    Type            Description                       |
|               |               |        |                       |         |        | --------------------------------------------------------------------------------                   |
|               |               |        |                       |         |        | Eth 0/1      connected (up)  --         100G     100G            Link to 60.30.181.101 Spine       |
|               |               |        |                       |         |        | Eth 0/2      connected (up)  --         100G     100G            Link to 60.30.181.102 Spine       |
|               |               |        |                       |         |        | Eth 0/3      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/4      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/5      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/6      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/7      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/8      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/9      adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/10     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/11     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/12     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/13     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/14     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/15     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/16     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/17:1   connected (up)  Access     25G      100G                                              |
|               |               |        |                       |         |        | Eth 0/17:2   connected (up)  Access     25G      100G                                              |
|               |               |        |                       |         |        | Eth 0/17:3   adminDown       --         --       100G                                              |
|               |               |        |                       |         |        | Eth 0/17:4   adminDown       --         --       100G                                              |
|               |               |        |                       |         |        | Eth 0/18:1   connected (up)  --         10G      25G-SFP         Port-channel CAS Member interface |
|               |               |        |                       |         |        | Eth 0/18:2   adminDown       --         --       25G-SFP                                           |
|               |               |        |                       |         |        | Eth 0/18:3   adminDown       --         --       25G-SFP                                           |
|               |               |        |                       |         |        | Eth 0/18:4   adminDown       --         --       25G-SFP                                           |
|               |               |        |                       |         |        | Eth 0/19     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/20     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/21     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/22     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/23     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/24     sfpAbsent       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/25     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/26     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/27     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/28     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/29     connected (up)  --         100G     100G            clusterPeerIntfMember             |
|               |               |        |                       |         |        | Eth 0/30     connected (up)  --         100G     100G            clusterPeerIntfMember             |
|               |               |        |                       |         |        | Eth 0/31     adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/32:1   adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/32:2   adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/32:3   adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Eth 0/32:4   adminDown       --         --       --                                                |
|               |               |        |                       |         |        | Po 1         connected (up)  Trunk      10G      --              CAS                               |
|               |               |        |                       |         |        | Po 64        connected (up)  --         200G     --              MCTPeerInterface                  |
|               |               |        |                       |         |        |                                                                                                    |
+---------------+---------------+--------+-----------------------+---------+--------+----------------------------------------------------------------------------------------------------+
Execute CLI Details
--- Time Elapsed: 5.585280961s ---
`
	} else {
		err = errors.New(fmt.Sprintf("Device with IPs [[%s]] are not registered", ip))
	}
	return result, err
}
