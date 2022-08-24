package structs

import (
	"fmt"
)

type ShowInterfaceStatus struct {
}

func (s *ShowInterfaceStatus) Print(ip string) {
	if ip == "60.30.181.101" {
		fmt.Print(`Execute CLI[success]
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
--- Time Elapsed: 500ms ---
`)
	}
}
