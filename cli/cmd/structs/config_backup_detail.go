package structs

import (
	"errors"
	"fmt"
)

type ConfigBackupDetail struct {
}

/**
TEST : 2022-09-21 11:20

> efa inventory config-backup detail --uuid {uuid} --show-config
*/
func (s *ConfigBackupDetail) Print(uuid string) (result string, err error) {
	if uuid == "19551554-f75f-4b7b-8902-ea8d5fc23bde" {
		result = `+---------------------+--------------------------------------+
|        NAME         |                VALUE                 |
+---------------------+--------------------------------------+
| UUID                | 19551554-f75f-4b7b-8902-ea8d5fc23bde |
+---------------------+--------------------------------------+
| Device IP           | 60.30.181.101                        |
+---------------------+--------------------------------------+
| Status              | success                              |
+---------------------+--------------------------------------+
| Execution Reason    | manual                               |
+---------------------+--------------------------------------+
| operation           | manual trigger                       |
+---------------------+--------------------------------------+
| Snapshot ID         | 13                                   |
+---------------------+--------------------------------------+
| Start Time          | 2022-08-29 13:06:57 +0900 KST        |
+---------------------+--------------------------------------+
| Last Modified Time  | 2022-08-29 13:07:20 +0900 KST        |
+---------------------+--------------------------------------+
| Last Sync Time      | 2022-08-29 13:06:57 +0900 KST        |
+---------------------+--------------------------------------+
| Duration            | 23.702528673s                        |
+---------------------+--------------------------------------+
Config Text     :
acl-policy
 acl-log-raslog log-interval 1
!
clock timezone Asia/Seoul
ha
 no process-restart bgp
 no process-restart isis
 no process-restart ospfv2
 no process-restart ospfv3
!
hardware
 profile route default
 profile route maximum-paths 64
 profile qos lossy
 connector 0/32
  breakout mode 4x1g
 !
!
switch-attributes chassis-name SLX9250-32C
switch-attributes host-name TB2-SLX-S1A
no support autoupload enable
support ffdc
resource-monitor cpu enable threshold 90 action raslog
resource-monitor memory enable threshold 100 action raslog
resource-monitor process memory enable alarm 1500 critical 1800
system-monitor fan threshold marginal-threshold 1 down-threshold 2
system-monitor fan alert state removed action raslog
system-monitor power threshold marginal-threshold 1 down-threshold 2
system-monitor power alert state removed action raslog
system-monitor temp threshold marginal-threshold 1 down-threshold 2
system-monitor cid-card threshold marginal-threshold 1 down-threshold 2
system-monitor cid-card alert state none action none
system-monitor compact-flash threshold marginal-threshold 1 down-threshold 0
system-monitor MM threshold marginal-threshold 1 down-threshold 0
system-monitor LineCard threshold marginal-threshold 1 down-threshold 2
system-monitor LineCard alert state none action none
system-monitor SFM threshold marginal-threshold 1 down-threshold 2
system-monitor port
 crc enable
!
system-monitor tm
!
system interface utilization-watermark
system maintenance
 enable-on-reboot
!
line vty
 exec-timeout 10
!
threshold-monitor Buffer limit 70
vrf mgmt-vrf
 address-family ipv4 unicast
  ip route 0.0.0.0/0 60.30.181.126
 !
 address-family ipv6 unicast
 !
!
ssh server max-idle-timeout 300
ssh server key rsa 2048
ssh server key ecdsa 256
ssh server key dsa
ssh server use-vrf mgmt-vrf
telnet server use-vrf default-vrf
telnet server use-vrf mgmt-vrf
http server use-vrf default-vrf
http server use-vrf mgmt-vrf
banner login "  ----------------------------------------------------------------------\n    !!!! WARNING !!!!WARNING !!!!WARNING !!!! WARNING !!!!\n  ----------------------------------------------------------------------\n     Unauthorized use of this system is prohibited.\n     All access to this equipment is logged.\n     Disconnect immediately if you are not an authorized user.\n     Violators will be prosecuted to the fullest extent of the law.\n  ----------------------------------------------------------------------\n    !!!! WARNING !!!!WARNING !!!!WARNING !!!! WARNING !!!!\n  ----------------------------------------------------------------------\n"
role name admin desc Administrator
role name user desc User
aaa authentication login oauth2 local-auth-fallback
aaa accounting exec default start-stop none
aaa accounting commands default start-stop none
aaa authorization command none
service password-encryption
username admin password $6$Uph9M2pYdE2Lft28$o4LI8Kgi2dh44LuLidC9vHmFNUjixlE3nw8JtgBogNW2iOEzVQTln2p7e.abZZK.z5qVqJnfF7YGn1Sic5e1i1 encryption-level 10 role admin desc Administrator
username user password $6$mAog0c./JxVGu1zy$6wFogQmek0KOEgTav.0DVKXz1vRodc1UCAbipYft/DWnT5R6/Y3qpq7V3JHlhRNVtwguLgXnzdtBDKPKaXbBg/ encryption-level 10 role user desc User
snmp-server sys-descr "Extreme SLX9250-32C Switch"
snmp-server engineID local 80:0:6:34:b2:4:0:0:10:1d:ef:d6:96
snmp-server community $9$4qOrvwtKcShb/JCoXGb7nw==
snmp-server community $9$pdTj0ueN2KcUO36QTF1PIQ== groupname read-only
snmp-server community $9$uZrcr+xAbndcYWJPaOUb1g==
snmp-server host 60.11.26.5 $9$uZrcr+xAbndcYWJPaOUb1g==
 source-interface management chassis-ip
!
snmp-server host 60.11.8.28 $9$4qOrvwtKcShb/JCoXGb7nw==
 source-interface management chassis-ip
!
snmp-server user efav3User groupname efav3Group auth sha auth-password LlgMJxy1dZfAfPydPDdFTNNtkpw= priv AES128 priv-password 53IJSJKiK/A6UFLVd9+ZkkclIJs= encrypted
snmp-server user skt groupname SKT auth sha auth-password xLMpyXWUPBeK7XI9pyaETkwBPgE= encrypted
snmp-server user sktip groupname SKT auth sha auth-password VhsrToxSUvuPh6nWXwFbuDALXrU= encrypted
snmp-server v3host 60.30.181.98 efav3User
 source-interface management chassis-ip
!
snmp-server view efav3View 1.3.6.1 included
snmp-server view read-only 1.3.6.1 included
snmp-server view view2 1.3.6.1 included
snmp-server group efav3Group v3 notify efav3View
snmp-server group read-only v2c read view2
snmp-server group SKT v3 auth read read-only notify read-only
ntp server 60.30.181.97
bpdu-drop-enable
ip mtu 9100
ip access-list extended ANTI-SPOOFING
 seq 10 permit udp host 0.0.0.0 eq bootpc host 255.255.255.255 eq bootps count
 seq 20 deny ip host 0.0.0.0 any count log
 seq 30 deny ip host 255.255.255.255 any count log
 seq 100 permit ip any any count
!
ip access-list extended MGMT-SNMP-PERMIT
 seq 10 permit udp host 0.0.0.0 eq bootpc host 255.255.255.255 eq bootps count
 seq 20 permit ip host 60.30.181.115 any count
 seq 30 permit ip host 60.30.181.97 any count
 seq 40 permit ip host 60.30.181.98 any count
 seq 50 permit ip host 60.30.181.99 any count
 seq 60 permit ip host 60.30.181.100 any count
 seq 70 permit ip host 60.30.181.116 any count
 seq 80 permit ip host 60.30.181.117 any count
 seq 90 permit ip host 60.30.181.118 any count
 seq 100 permit ip host 60.30.181.120 any count
 seq 110 permit ip host 60.30.181.121 any count
 seq 120 permit ip host 60.30.131.209 any count
 seq 130 permit ip host 60.30.181.101 any count
 seq 140 permit ip host 60.30.181.102 any count
 seq 150 permit ip host 60.30.181.103 any count
 seq 160 permit ip host 60.30.181.104 any count
 seq 170 permit ip host 60.30.181.105 any count
 seq 180 permit ip host 60.30.181.106 any count
 seq 190 permit ip host 60.30.181.107 any count
 seq 200 permit ip host 60.30.181.108 any count
 seq 210 permit ip host 60.30.181.109 any count
 seq 220 permit ip host 60.30.181.110 any count
 seq 230 permit ip host 60.30.181.111 any count
 seq 240 permit ip host 60.30.181.112 any count
 seq 250 permit ip host 60.30.181.113 any count
 seq 260 permit ip host 60.30.181.114 any count
 seq 1000 deny ip any any count
!
vlan 1
!
protocol lldp
 system-description Extreme SLX9250-32C Switch
!
vlan dot1q tag native
ip router-id 198.19.96.13
class-map default
!
no protocol vrrp
no protocol vrrp-extended
router bgp
 local-as 66000
 capability as4-enable
 fast-external-fallover
 listen-limit 255
 neighbor leaf-group peer-group
 neighbor leaf-group description To border-leaf
 neighbor leaf-group bfd
 neighbor 198.19.97.5 remote-as 66011
 neighbor 198.19.97.5 peer-group leaf-group
 neighbor 198.19.97.6 remote-as 66011
 neighbor 198.19.97.6 peer-group leaf-group
 neighbor 198.19.97.12 remote-as 66012
 neighbor 198.19.97.12 peer-group leaf-group
 neighbor 198.19.97.16 remote-as 66012
 neighbor 198.19.97.16 peer-group leaf-group
 neighbor 198.19.97.20 remote-as 66004
 neighbor 198.19.97.20 peer-group leaf-group
 neighbor 198.19.97.24 remote-as 66004
 neighbor 198.19.97.24 peer-group leaf-group
 neighbor 198.19.97.29 remote-as 66001
 neighbor 198.19.97.29 peer-group leaf-group
 neighbor 198.19.97.33 remote-as 66001
 neighbor 198.19.97.33 peer-group leaf-group
 neighbor 198.19.97.36 remote-as 66002
 neighbor 198.19.97.36 peer-group leaf-group
 neighbor 198.19.97.42 remote-as 66002
 neighbor 198.19.97.42 peer-group leaf-group
 neighbor 198.19.97.50 remote-as 66003
 neighbor 198.19.97.50 peer-group leaf-group
 neighbor 198.19.97.56 remote-as 66003
 neighbor 198.19.97.56 peer-group leaf-group
 neighbor 198.19.97.69 remote-as 66003
 neighbor 198.19.97.69 peer-group leaf-group
 neighbor 198.19.97.71 remote-as 66003
 neighbor 198.19.97.71 peer-group leaf-group
 address-family ipv4 unicast
  maximum-paths 8
  graceful-restart
 !
 address-family ipv6 unicast
 !
 address-family l2vpn evpn
  graceful-restart
  retain route-target all
  neighbor leaf-group encapsulation vxlan
  neighbor leaf-group next-hop-unchanged
  neighbor leaf-group enable-peer-as-check
  neighbor leaf-group activate
 !
!
interface Loopback 1
 ip address 198.19.96.13/32
 no shutdown
!
interface Management 0
 no tcp burstrate
 no shutdown
 vrf forwarding mgmt-vrf
 ip icmp unreachable
 ip icmp echo-reply
 no ip address dhcp
 ip address 60.30.181.101/27
 ip access-group MGMT-SNMP-PERMIT in
 no ipv6 address autoconfig
 no ipv6 address dhcp
!
interface Ethernet 0/1
 description Link to 60.30.181.103 Leaf
 ip address 198.19.97.21/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/2
 description Link to 60.30.181.104 Leaf
 ip address 198.19.97.25/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/3
 description Link to 60.30.181.105 Leaf
 ip address 198.19.97.28/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/4
 description Link to 60.30.181.106 Leaf
 ip address 198.19.97.32/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/5
 description Link to 60.30.181.107 Leaf
 ip address 198.19.97.37/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/6
 description Link to 60.30.181.108 Leaf
 ip address 198.19.97.43/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/7
 description Link to 60.30.181.109 Leaf
 ip address 198.19.97.51/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/8
 description Link to 60.30.181.110 Leaf
 ip address 198.19.97.57/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/9
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/10
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/11
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/12
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/13
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/14
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/15
 description Link to 60.30.181.109 Leaf
 ip address 198.19.97.70/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/16
 description Link to 60.30.181.110 Leaf
 ip address 198.19.97.68/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/17
 shutdown
!
interface Ethernet 0/18
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/19
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/20
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/21
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/22
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/23
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/24
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/25
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/26
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/27
 description Link to 60.30.181.111 border-leaf
 ip address 198.19.97.4/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/28
 description Link to 60.30.181.112 border-leaf
 ip address 198.19.97.7/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/29
 description Link to 60.30.181.113 border-leaf
 ip address 198.19.97.13/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/30
 description Link to 60.30.181.114 border-leaf
 ip address 198.19.97.17/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/31
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/32:1
 no shutdown
!
interface Ethernet 0/32:2
 shutdown
!
interface Ethernet 0/32:3
 shutdown
!
interface Ethernet 0/32:4
 shutdown
!
logging raslog message RAS-2006 suppress
logging raslog message RAS-2007 suppress
logging raslog message SEC-1203 suppress
logging raslog message SEC-1206 suppress
logging raslog message SEC-3022 suppress
logging raslog console INFO
logging syslog-server 60.30.181.98 use-vrf mgmt-vrf
 secure
!
logging auditlog class CONFIGURATION
logging auditlog class FIRMWARE
logging syslog-facility local LOG_LOCAL6
telemetry profile system-utilization default_system_utilization_statistics
 interval 60
 add total-system-memory
 add total-used-memory
 add total-free-memory
 add cached-memory
 add buffers
 add total-swap-memory
 add total-free-swap-memory
 add total-used-swap-memory
 add user-process
 add system-process
 add niced-process
 add iowait
 add hw-interrupt
 add sw-interrupt
 add idle-state
 add steal-time
 add uptime
!
telemetry profile interface default_interface_statistics
 interval 30
 add out-pkts
 add in-pkts
 add out-unicast-pkts
 add in-unicast-pkts
 add out-broadcast-pkts
 add in-broadcast-pkts
 add out-multicast-pkts
 add in-multicast-pkts
 add out-pkts-per-second
 add in-pkts-per-second
 add out-bandwidth
 add in-bandwidth
 add out-octets
 add in-octets
 add out-errors
 add in-errors
 add out-crc-errors
 add in-crc-errors
 add out-discards
 add in-discards
!
tpvm TPVM
 password $6$lSexdbj5$9pCijwxf251d7B1RPKs8zvmEUhqLPda/oLOlTlcsVYQOq5U8QFb70ZEObTmJgR6eXBBkiXR8viWMovqVVf44q1
 auto-boot
 ntp 60.30.181.101
 hostname tpvm1
 timezone Asia/Seoul
 trusted-peer ip 60.30.181.100 password $9$BwrsDbB+tABWGWpINOVKoQ==
 interface management ip 60.30.181.99/27 gw 60.30.181.101
 deploy
!

--- Time Elapsed: 670.776483ms ---
`
	} else if uuid == "b8765eeb-1499-4186-9521-2ad06563d87b" {
		result = `+---------------------+--------------------------------------+
|        NAME         |                VALUE                 |
+---------------------+--------------------------------------+
| UUID                | b8765eeb-1499-4186-9521-2ad06563d87b |
+---------------------+--------------------------------------+
| Device IP           | 60.30.181.111                        |
+---------------------+--------------------------------------+
| Status              | success                              |
+---------------------+--------------------------------------+
| Execution Reason    | manual                               |
+---------------------+--------------------------------------+
| operation           | manual trigger                       |
+---------------------+--------------------------------------+
| Snapshot ID         | 31                                   |
+---------------------+--------------------------------------+
| Start Time          | 2022-08-29 13:07:34 +0900 KST        |
+---------------------+--------------------------------------+
| Last Modified Time  | 2022-08-29 13:08:02 +0900 KST        |
+---------------------+--------------------------------------+
| Last Sync Time      | 2022-08-29 13:07:34 +0900 KST        |
+---------------------+--------------------------------------+
| Duration            | 27.872901586s                        |
+---------------------+--------------------------------------+
Config Text     :
acl-policy
 acl-log-raslog log-interval 1
!
clock timezone Asia/Seoul
ha
 process-restart mpls
 no process-restart bgp
 no process-restart isis
 no process-restart ospfv2
 no process-restart ospfv3
!
hardware
 profile tcam default
 profile route default
 profile route maximum-paths 64
 connector 0/23
  breakout mode 4x10g
 !
 connector 0/27
  breakout mode 4x25g
 !
 connector 0/39
  breakout mode 4x10g
 !
!
pw-profile Tenant-profile
!
switch-attributes chassis-name SLX9740-40C
switch-attributes host-name TB2-SLX-DBL-A
no support autoupload enable
support ffdc
resource-monitor cpu enable threshold 90 action raslog
resource-monitor memory enable threshold 100 action raslog
resource-monitor process memory enable alarm 1500 critical 1800
system-monitor fan threshold marginal-threshold 1 down-threshold 2
system-monitor fan alert state removed action raslog
system-monitor power threshold marginal-threshold 1 down-threshold 2
system-monitor power alert state removed action raslog
system-monitor temp threshold marginal-threshold 1 down-threshold 2
system-monitor cid-card threshold marginal-threshold 1 down-threshold 2
system-monitor cid-card alert state none action none
system-monitor compact-flash threshold marginal-threshold 1 down-threshold 0
system-monitor MM threshold marginal-threshold 1 down-threshold 0
system-monitor LineCard threshold marginal-threshold 1 down-threshold 2
system-monitor LineCard alert state none action none
system-monitor SFM threshold marginal-threshold 1 down-threshold 2
system-monitor port
 crc enable
!
system-monitor tm
!
system interface utilization-watermark
system maintenance
 enable-on-reboot
!
line vty
 exec-timeout 10
!
threshold-monitor Buffer limit 70
vrf 5G_VRF
 rd 198.19.96.93:3
 address-family ipv4 unicast
  route-target export 12:12 evpn
  route-target import 12:12 evpn
  ip route static bfd 40.172.2.11 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.12 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.13 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.14 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.15 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.16 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.17 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.18 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.19 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.20 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.21 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.22 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.23 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.24 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.25 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.26 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.27 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.28 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.29 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.30 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.31 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.32 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.33 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.34 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.35 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.36 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.37 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.38 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.39 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.40 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.41 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.42 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.43 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.44 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.45 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.46 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.47 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.48 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.49 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.50 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.51 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.52 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.53 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.2.54 40.172.2.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.142.5 172.1.142.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.142.6 172.1.142.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.142.7 172.1.142.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.142.8 172.1.142.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.142.9 172.1.142.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.142.10 172.1.142.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.142.11 172.1.142.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.142.12 172.1.142.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.142.125 172.1.142.122 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.142.225 172.1.142.226 interval 50 min-rx 50 multiplier 4
  ip route static bfd 172.1.142.229 172.1.142.230 interval 50 min-rx 50 multiplier 4
  ip route 0.0.0.0/0 172.1.142.225
  ip route 0.0.0.0/0 172.1.142.229
  ip route 10.101.0.0/16 40.255.255.1
  ip route 10.101.0.0/16 40.255.255.101
  ip route 10.102.0.0/16 40.255.255.1
  ip route 10.102.0.0/16 40.255.255.101
  ip route 10.103.0.0/16 40.255.255.1
  ip route 10.103.0.0/16 40.255.255.101
  ip route 10.104.0.0/16 40.255.255.1
  ip route 10.104.0.0/16 40.255.255.101
  ip route 10.105.0.0/16 40.255.255.1
  ip route 10.105.0.0/16 40.255.255.101
  ip route 10.106.0.0/16 40.255.255.1
  ip route 10.106.0.0/16 40.255.255.101
  ip route 10.107.0.0/16 40.255.255.1
  ip route 10.107.0.0/16 40.255.255.101
  ip route 10.108.0.0/16 40.255.255.101
  ip route 10.111.0.0/16 40.255.255.1
  ip route 10.111.0.0/16 40.255.255.101
  ip route 10.112.0.0/16 40.255.255.1
  ip route 10.112.0.0/16 40.255.255.101
  ip route 10.113.0.0/16 40.255.255.1
  ip route 10.113.0.0/16 40.255.255.101
  ip route 10.114.0.0/16 40.255.255.1
  ip route 10.114.0.0/16 40.255.255.101
  ip route 10.115.0.0/16 40.255.255.1
  ip route 10.115.0.0/16 40.255.255.101
  ip route 10.116.0.0/16 40.255.255.1
  ip route 10.116.0.0/16 40.255.255.101
  ip route 10.117.0.0/16 40.255.255.101
  ip route 10.118.0.0/16 40.255.255.101
  ip route 20.1.13.1/32 40.255.255.1
  ip route 20.1.13.41/32 40.255.255.1
  ip route 20.1.13.55/32 40.255.255.1
  ip route 20.1.13.57/32 40.255.255.1
  ip route 20.1.13.76/32 40.255.255.1
  ip route 40.173.2.11/32 40.172.2.11
  ip route 40.173.2.12/32 40.172.2.12
  ip route 40.173.2.13/32 40.172.2.13
  ip route 40.173.2.14/32 40.172.2.14
  ip route 40.173.2.15/32 40.172.2.15
  ip route 40.173.2.16/32 40.172.2.16
  ip route 40.173.2.17/32 40.172.2.17
  ip route 40.173.2.18/32 40.172.2.18
  ip route 40.173.2.19/32 40.172.2.19
  ip route 40.173.2.20/32 40.172.2.20
  ip route 40.173.2.21/32 40.172.2.21
  ip route 40.173.2.22/32 40.172.2.22
  ip route 40.173.2.23/32 40.172.2.23
  ip route 40.173.2.24/32 40.172.2.24
  ip route 40.173.2.25/32 40.172.2.25
  ip route 40.173.2.26/32 40.172.2.26
  ip route 40.173.2.27/32 40.172.2.27
  ip route 40.173.2.28/32 40.172.2.28
  ip route 40.173.2.29/32 40.172.2.29
  ip route 40.173.2.30/32 40.172.2.30
  ip route 40.173.2.31/32 40.172.2.31
  ip route 40.173.2.32/32 40.172.2.32
  ip route 40.173.2.33/32 40.172.2.33
  ip route 40.173.2.34/32 40.172.2.34
  ip route 40.173.2.35/32 40.172.2.35
  ip route 40.173.2.36/32 40.172.2.36
  ip route 40.173.2.37/32 40.172.2.37
  ip route 40.173.2.38/32 40.172.2.38
  ip route 40.173.2.39/32 40.172.2.39
  ip route 40.173.2.40/32 40.172.2.40
  ip route 40.173.2.41/32 40.172.2.41
  ip route 40.173.2.42/32 40.172.2.42
  ip route 40.173.2.43/32 40.172.2.43
  ip route 40.173.2.44/32 40.172.2.44
  ip route 40.173.2.45/32 40.172.2.45
  ip route 40.173.2.46/32 40.172.2.46
  ip route 40.173.2.47/32 40.172.2.47
  ip route 40.173.2.48/32 40.172.2.48
  ip route 40.173.2.49/32 40.172.2.49
  ip route 40.173.2.50/32 40.172.2.50
  ip route 40.173.2.51/32 40.172.2.51
  ip route 40.173.2.52/32 40.172.2.52
  ip route 40.173.2.53/32 40.172.2.53
  ip route 40.173.2.54/32 40.172.2.54
  ip route 172.1.142.32/30 172.1.142.5
  ip route 172.1.142.32/30 172.1.142.6
  ip route 172.1.142.32/30 172.1.142.7
  ip route 172.1.142.32/30 172.1.142.8
  ip route 172.1.142.32/30 172.1.142.9
  ip route 172.1.142.32/30 172.1.142.10
  ip route 172.1.142.32/30 172.1.142.11
  ip route 172.1.142.32/30 172.1.142.12
  ip route 172.1.142.40/32 40.172.2.11
  ip route 172.1.142.40/32 40.172.2.12
  ip route 172.1.142.40/32 40.172.2.13
  ip route 172.1.142.40/32 40.172.2.14
  ip route 172.1.142.40/32 40.172.2.15
  ip route 172.1.142.40/32 40.172.2.16
  ip route 172.1.142.40/32 40.172.2.17
  ip route 172.1.142.40/32 40.172.2.18
  ip route 172.1.142.40/32 40.172.2.19
  ip route 172.1.142.40/32 40.172.2.20
  ip route 172.1.142.40/32 40.172.2.21
  ip route 172.1.142.40/32 40.172.2.22
  ip route 172.1.142.40/32 40.172.2.23
  ip route 172.1.142.40/32 40.172.2.24
  ip route 172.1.142.40/32 40.172.2.25
  ip route 172.1.142.40/32 40.172.2.26
  ip route 172.1.142.40/32 40.172.2.27
  ip route 172.1.142.40/32 40.172.2.28
  ip route 172.1.142.40/32 40.172.2.29
  ip route 172.1.142.40/32 40.172.2.30
  ip route 172.1.142.40/32 40.172.2.31
  ip route 172.1.142.40/32 40.172.2.32
  ip route 172.1.142.40/32 40.172.2.33
  ip route 172.1.142.40/32 40.172.2.34
  ip route 172.1.142.40/32 40.172.2.35
  ip route 172.1.142.40/32 40.172.2.36
  ip route 172.1.142.40/32 40.172.2.37
  ip route 172.1.142.40/32 40.172.2.38
  ip route 172.1.142.40/32 40.172.2.39
  ip route 172.1.142.40/32 40.172.2.40
  ip route 172.1.142.40/32 40.172.2.41
  ip route 172.1.142.40/32 40.172.2.42
  ip route 172.1.142.40/32 40.172.2.43
  ip route 172.1.142.40/32 40.172.2.44
  ip route 172.1.142.40/32 40.172.2.45
  ip route 172.1.142.40/32 40.172.2.46
  ip route 172.1.142.40/32 40.172.2.47
  ip route 172.1.142.40/32 40.172.2.48
  ip route 172.1.142.40/32 40.172.2.49
  ip route 172.1.142.40/32 40.172.2.50
  ip route 172.1.142.40/32 40.172.2.51
  ip route 172.1.142.40/32 40.172.2.52
  ip route 172.1.142.40/32 40.172.2.53
  ip route 172.1.142.40/32 40.172.2.54
  ip route 172.1.142.216/29 172.1.142.125
  ip route 50.15.0.0/24 next-hop-vrf NAT01_VRF 40.172.1.139
  ip route 50.15.3.0/24 next-hop-vrf NAT01_VRF 40.172.1.139
 !
 address-family ipv6 unicast
  route-target export 12:12 evpn
  route-target import 12:12 evpn
 !
!
vrf GLOBAL_VRF
 rd 198.19.96.93:4
 address-family ipv4 unicast
  route-target export 10:10 evpn
  route-target import 10:10 evpn
  ip route static bfd 40.172.1.11 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.12 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.13 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.14 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.15 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.16 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.17 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.18 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.19 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.20 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.21 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.22 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.23 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.24 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.25 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.26 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.27 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.28 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.29 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.30 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.31 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.32 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.33 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.34 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.35 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.36 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.37 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.38 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.39 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.40 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.41 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.42 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.43 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.44 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.45 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.46 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.47 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.48 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.49 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.50 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.51 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.52 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.53 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.54 40.172.1.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.75 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.76 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.77 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.78 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.79 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.80 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.81 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.82 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.83 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.84 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.85 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.86 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.87 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.88 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.89 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.90 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.91 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.92 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.93 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.94 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.95 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.96 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.97 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.98 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.99 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.100 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.101 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.102 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.103 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.104 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.105 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.106 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.107 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.108 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.109 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.110 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.111 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.112 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.113 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.114 40.172.1.66 interval 300 min-rx 300 multiplier 3
  ip route static bfd 40.172.1.115 40.172.1.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.116 40.172.1.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.117 40.172.1.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.118 40.172.1.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.203 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.204 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.205 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.206 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.207 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.208 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.209 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.210 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.211 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.212 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.213 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.214 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.215 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.216 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.217 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.218 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.219 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.220 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.221 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.222 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.223 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.224 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.225 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.226 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.227 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.228 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.229 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.230 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.231 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.232 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.233 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.234 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.235 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.236 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.237 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.238 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.239 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.240 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.241 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.242 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.243 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.244 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.245 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.246 40.172.1.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.141.80 40.172.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.141.81 40.172.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.141.82 40.172.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.141.83 40.172.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.141.84 40.172.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.141.85 40.172.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.141.86 40.172.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.141.87 40.172.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.141.88 40.172.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.141.89 40.172.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.5 172.1.141.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.6 172.1.141.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.7 172.1.141.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.8 172.1.141.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.9 172.1.141.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.10 172.1.141.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.11 172.1.141.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.12 172.1.141.2 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.23 172.1.141.18 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.24 172.1.141.18 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.25 172.1.141.18 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.26 172.1.141.18 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.27 172.1.141.18 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.28 172.1.141.18 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.29 172.1.141.18 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.30 172.1.141.18 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.45 172.1.141.42 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.80 172.1.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.81 172.1.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.82 172.1.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.83 172.1.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.84 172.1.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.85 172.1.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.86 172.1.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.87 172.1.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.88 172.1.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.89 172.1.141.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.137 172.1.141.138 interval 50 min-rx 50 multiplier 4
  ip route static bfd 172.1.141.141 172.1.141.142 interval 50 min-rx 50 multiplier 4
  ip route static bfd 172.1.141.145 172.1.141.146 interval 50 min-rx 50 multiplier 4
  ip route static bfd 172.1.141.199 172.1.141.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.200 172.1.141.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.201 172.1.141.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.202 172.1.141.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.203 172.1.141.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.204 172.1.141.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.205 172.1.141.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.206 172.1.141.194 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.225 172.1.141.226 interval 50 min-rx 50 multiplier 4
  ip route static bfd 172.1.141.229 172.1.141.230 interval 50 min-rx 50 multiplier 4
  ip route 0.0.0.0/0 172.1.141.225
  ip route 0.0.0.0/0 172.1.141.229
  ip route 10.0.0.0/12 40.255.255.109
  ip route 10.0.0.0/13 40.255.255.9
  ip route 10.8.0.0/13 40.255.255.9
  ip route 10.101.0.0/16 40.255.255.9
  ip route 10.101.0.0/16 40.255.255.109
  ip route 10.102.0.0/16 40.255.255.9
  ip route 10.102.0.0/16 40.255.255.109
  ip route 10.103.0.0/16 40.255.255.9
  ip route 10.103.0.0/16 40.255.255.109
  ip route 10.104.0.0/16 40.255.255.9
  ip route 10.104.0.0/16 40.255.255.109
  ip route 10.105.0.0/16 40.255.255.9
  ip route 10.105.0.0/16 40.255.255.109
  ip route 10.106.0.0/16 40.255.255.9
  ip route 10.106.0.0/16 40.255.255.109
  ip route 10.107.0.0/16 40.255.255.9
  ip route 10.107.0.0/16 40.255.255.109
  ip route 10.108.0.0/16 40.255.255.109
  ip route 10.111.0.0/16 40.255.255.9
  ip route 10.111.0.0/16 40.255.255.109
  ip route 10.112.0.0/16 40.255.255.9
  ip route 10.112.0.0/16 40.255.255.109
  ip route 10.113.0.0/16 40.255.255.9
  ip route 10.113.0.0/16 40.255.255.109
  ip route 10.114.0.0/16 40.255.255.9
  ip route 10.114.0.0/16 40.255.255.109
  ip route 10.115.0.0/16 40.255.255.9
  ip route 10.115.0.0/16 40.255.255.109
  ip route 10.116.0.0/16 40.255.255.9
  ip route 10.116.0.0/16 40.255.255.109
  ip route 10.117.0.0/16 40.255.255.9
  ip route 10.117.0.0/16 40.255.255.109
  ip route 10.118.0.0/16 40.255.255.109
  ip route 20.0.0.0/27 40.172.141.19
  ip route 22.0.0.0/8 172.1.141.145
  ip route 22.0.0.0/9 172.1.141.137
  ip route 22.128.0.0/9 172.1.141.137
  ip route 40.172.141.59/32 40.172.141.99
  ip route 40.172.141.59/32 40.172.141.100
  ip route 40.172.141.59/32 40.172.141.101
  ip route 40.172.141.59/32 40.172.141.102
  ip route 40.172.141.59/32 40.172.141.103
  ip route 40.172.141.59/32 40.172.141.104
  ip route 40.172.141.59/32 40.172.141.105
  ip route 40.172.141.59/32 40.172.141.106
  ip route 40.172.141.59/32 40.172.141.107
  ip route 40.172.141.59/32 40.172.141.108
  ip route 40.172.141.99/32 40.172.141.80
  ip route 40.172.141.100/32 40.172.141.81
  ip route 40.172.141.101/32 40.172.141.82
  ip route 40.172.141.102/32 40.172.141.83
  ip route 40.172.141.103/32 40.172.141.84
  ip route 40.172.141.104/32 40.172.141.85
  ip route 40.172.141.105/32 40.172.141.86
  ip route 40.172.141.106/32 40.172.141.87
  ip route 40.172.141.107/32 40.172.141.88
  ip route 40.172.141.108/32 40.172.141.89
  ip route 40.172.141.129/32 172.1.141.21
  ip route 40.172.141.130/32 172.1.141.22
  ip route 40.172.141.133/32 172.1.141.197
  ip route 40.172.141.134/32 172.1.141.198
  ip route 40.172.141.136/32 172.1.141.23
  ip route 40.172.141.136/32 172.1.141.24
  ip route 40.172.141.136/32 172.1.141.25
  ip route 40.172.141.136/32 172.1.141.26
  ip route 40.172.141.136/32 172.1.141.27
  ip route 40.172.141.136/32 172.1.141.28
  ip route 40.172.141.136/32 172.1.141.29
  ip route 40.172.141.136/32 172.1.141.30
  ip route 40.172.141.140/32 172.1.141.199
  ip route 40.172.141.140/32 172.1.141.200
  ip route 40.172.141.140/32 172.1.141.201
  ip route 40.172.141.140/32 172.1.141.202
  ip route 40.172.141.140/32 172.1.141.203
  ip route 40.172.141.140/32 172.1.141.204
  ip route 40.172.141.140/32 172.1.141.205
  ip route 40.172.141.140/32 172.1.141.206
  ip route 40.173.1.11/32 40.172.1.11
  ip route 40.173.1.12/32 40.172.1.12
  ip route 40.173.1.13/32 40.172.1.13
  ip route 40.173.1.14/32 40.172.1.14
  ip route 40.173.1.15/32 40.172.1.15
  ip route 40.173.1.16/32 40.172.1.16
  ip route 40.173.1.17/32 40.172.1.17
  ip route 40.173.1.18/32 40.172.1.18
  ip route 40.173.1.19/32 40.172.1.19
  ip route 40.173.1.20/32 40.172.1.20
  ip route 40.173.1.21/32 40.172.1.21
  ip route 40.173.1.22/32 40.172.1.22
  ip route 40.173.1.23/32 40.172.1.23
  ip route 40.173.1.24/32 40.172.1.24
  ip route 40.173.1.25/32 40.172.1.25
  ip route 40.173.1.26/32 40.172.1.26
  ip route 40.173.1.27/32 40.172.1.27
  ip route 40.173.1.28/32 40.172.1.28
  ip route 40.173.1.29/32 40.172.1.29
  ip route 40.173.1.30/32 40.172.1.30
  ip route 40.173.1.31/32 40.172.1.31
  ip route 40.173.1.32/32 40.172.1.32
  ip route 40.173.1.33/32 40.172.1.33
  ip route 40.173.1.34/32 40.172.1.34
  ip route 40.173.1.35/32 40.172.1.35
  ip route 40.173.1.36/32 40.172.1.36
  ip route 40.173.1.37/32 40.172.1.37
  ip route 40.173.1.38/32 40.172.1.38
  ip route 40.173.1.39/32 40.172.1.39
  ip route 40.173.1.40/32 40.172.1.40
  ip route 40.173.1.41/32 40.172.1.41
  ip route 40.173.1.42/32 40.172.1.42
  ip route 40.173.1.43/32 40.172.1.43
  ip route 40.173.1.44/32 40.172.1.44
  ip route 40.173.1.45/32 40.172.1.45
  ip route 40.173.1.46/32 40.172.1.46
  ip route 40.173.1.47/32 40.172.1.47
  ip route 40.173.1.48/32 40.172.1.48
  ip route 40.173.1.49/32 40.172.1.49
  ip route 40.173.1.50/32 40.172.1.50
  ip route 40.173.1.51/32 40.172.1.51
  ip route 40.173.1.52/32 40.172.1.52
  ip route 40.173.1.53/32 40.172.1.53
  ip route 40.173.1.54/32 40.172.1.54
  ip route 40.173.1.75/32 40.172.1.75
  ip route 40.173.1.76/32 40.172.1.76
  ip route 40.173.1.77/32 40.172.1.77
  ip route 40.173.1.78/32 40.172.1.78
  ip route 40.173.1.79/32 40.172.1.79
  ip route 40.173.1.80/32 40.172.1.80
  ip route 40.173.1.81/32 40.172.1.81
  ip route 40.173.1.82/32 40.172.1.82
  ip route 40.173.1.83/32 40.172.1.83
  ip route 40.173.1.84/32 40.172.1.84
  ip route 40.173.1.85/32 40.172.1.85
  ip route 40.173.1.86/32 40.172.1.86
  ip route 40.173.1.87/32 40.172.1.87
  ip route 40.173.1.88/32 40.172.1.88
  ip route 40.173.1.89/32 40.172.1.89
  ip route 40.173.1.90/32 40.172.1.90
  ip route 40.173.1.91/32 40.172.1.91
  ip route 40.173.1.92/32 40.172.1.92
  ip route 40.173.1.93/32 40.172.1.93
  ip route 40.173.1.94/32 40.172.1.94
  ip route 40.173.1.95/32 40.172.1.95
  ip route 40.173.1.96/32 40.172.1.96
  ip route 40.173.1.97/32 40.172.1.97
  ip route 40.173.1.98/32 40.172.1.98
  ip route 40.173.1.99/32 40.172.1.99
  ip route 40.173.1.100/32 40.172.1.100
  ip route 40.173.1.101/32 40.172.1.101
  ip route 40.173.1.102/32 40.172.1.102
  ip route 40.173.1.103/32 40.172.1.103
  ip route 40.173.1.104/32 40.172.1.104
  ip route 40.173.1.105/32 40.172.1.105
  ip route 40.173.1.106/32 40.172.1.106
  ip route 40.173.1.107/32 40.172.1.107
  ip route 40.173.1.108/32 40.172.1.108
  ip route 40.173.1.109/32 40.172.1.109
  ip route 40.173.1.110/32 40.172.1.110
  ip route 40.173.1.111/32 40.172.1.111
  ip route 40.173.1.112/32 40.172.1.112
  ip route 40.173.1.113/32 40.172.1.113
  ip route 40.173.1.114/32 40.172.1.114
  ip route 40.173.1.115/32 40.172.1.115
  ip route 40.173.1.116/32 40.172.1.116
  ip route 40.173.1.117/32 40.172.1.117
  ip route 40.173.1.118/32 40.172.1.118
  ip route 40.173.1.203/32 40.172.1.203
  ip route 40.173.1.204/32 40.172.1.204
  ip route 40.173.1.205/32 40.172.1.205
  ip route 40.173.1.206/32 40.172.1.206
  ip route 40.173.1.207/32 40.172.1.207
  ip route 40.173.1.208/32 40.172.1.208
  ip route 40.173.1.209/32 40.172.1.209
  ip route 40.173.1.210/32 40.172.1.210
  ip route 40.173.1.211/32 40.172.1.211
  ip route 40.173.1.212/32 40.172.1.212
  ip route 40.173.1.213/32 40.172.1.213
  ip route 40.173.1.214/32 40.172.1.214
  ip route 40.173.1.215/32 40.172.1.215
  ip route 40.173.1.216/32 40.172.1.216
  ip route 40.173.1.217/32 40.172.1.217
  ip route 40.173.1.218/32 40.172.1.218
  ip route 40.173.1.219/32 40.172.1.219
  ip route 40.173.1.220/32 40.172.1.220
  ip route 40.173.1.221/32 40.172.1.221
  ip route 40.173.1.222/32 40.172.1.222
  ip route 40.173.1.223/32 40.172.1.223
  ip route 40.173.1.224/32 40.172.1.224
  ip route 40.173.1.225/32 40.172.1.225
  ip route 40.173.1.226/32 40.172.1.226
  ip route 40.173.1.227/32 40.172.1.227
  ip route 40.173.1.228/32 40.172.1.228
  ip route 40.173.1.229/32 40.172.1.229
  ip route 40.173.1.230/32 40.172.1.230
  ip route 40.173.1.231/32 40.172.1.231
  ip route 40.173.1.232/32 40.172.1.232
  ip route 40.173.1.233/32 40.172.1.233
  ip route 40.173.1.234/32 40.172.1.234
  ip route 40.173.1.235/32 40.172.1.235
  ip route 40.173.1.236/32 40.172.1.236
  ip route 40.173.1.237/32 40.172.1.237
  ip route 40.173.1.238/32 40.172.1.238
  ip route 40.173.1.239/32 40.172.1.239
  ip route 40.173.1.240/32 40.172.1.240
  ip route 40.173.1.241/32 40.172.1.241
  ip route 40.173.1.242/32 40.172.1.242
  ip route 40.173.1.243/32 40.172.1.243
  ip route 40.173.1.244/32 40.172.1.244
  ip route 40.173.1.245/32 40.172.1.245
  ip route 40.173.1.246/32 40.172.1.246
  ip route 172.1.141.32/30 172.1.141.5
  ip route 172.1.141.32/30 172.1.141.6
  ip route 172.1.141.32/30 172.1.141.7
  ip route 172.1.141.32/30 172.1.141.8
  ip route 172.1.141.32/30 172.1.141.9
  ip route 172.1.141.32/30 172.1.141.10
  ip route 172.1.141.32/30 172.1.141.11
  ip route 172.1.141.32/30 172.1.141.12
  ip route 172.1.141.36/30 172.1.141.45
  ip route 172.1.141.56/32 40.172.1.11
  ip route 172.1.141.56/32 40.172.1.12
  ip route 172.1.141.56/32 40.172.1.13
  ip route 172.1.141.56/32 40.172.1.14
  ip route 172.1.141.56/32 40.172.1.15
  ip route 172.1.141.56/32 40.172.1.16
  ip route 172.1.141.56/32 40.172.1.17
  ip route 172.1.141.56/32 40.172.1.18
  ip route 172.1.141.56/32 40.172.1.19
  ip route 172.1.141.56/32 40.172.1.20
  ip route 172.1.141.56/32 40.172.1.21
  ip route 172.1.141.56/32 40.172.1.22
  ip route 172.1.141.56/32 40.172.1.23
  ip route 172.1.141.56/32 40.172.1.24
  ip route 172.1.141.56/32 40.172.1.25
  ip route 172.1.141.56/32 40.172.1.26
  ip route 172.1.141.56/32 40.172.1.27
  ip route 172.1.141.56/32 40.172.1.28
  ip route 172.1.141.56/32 40.172.1.29
  ip route 172.1.141.56/32 40.172.1.30
  ip route 172.1.141.56/32 40.172.1.31
  ip route 172.1.141.56/32 40.172.1.32
  ip route 172.1.141.56/32 40.172.1.33
  ip route 172.1.141.56/32 40.172.1.34
  ip route 172.1.141.56/32 40.172.1.35
  ip route 172.1.141.56/32 40.172.1.36
  ip route 172.1.141.56/32 40.172.1.37
  ip route 172.1.141.56/32 40.172.1.38
  ip route 172.1.141.56/32 40.172.1.39
  ip route 172.1.141.56/32 40.172.1.40
  ip route 172.1.141.56/32 40.172.1.41
  ip route 172.1.141.56/32 40.172.1.42
  ip route 172.1.141.56/32 40.172.1.43
  ip route 172.1.141.56/32 40.172.1.44
  ip route 172.1.141.56/32 40.172.1.45
  ip route 172.1.141.56/32 40.172.1.46
  ip route 172.1.141.56/32 40.172.1.47
  ip route 172.1.141.56/32 40.172.1.48
  ip route 172.1.141.56/32 40.172.1.49
  ip route 172.1.141.56/32 40.172.1.50
  ip route 172.1.141.56/32 40.172.1.51
  ip route 172.1.141.56/32 40.172.1.52
  ip route 172.1.141.56/32 40.172.1.53
  ip route 172.1.141.56/32 40.172.1.54
  ip route 172.1.141.57/32 40.172.1.75
  ip route 172.1.141.57/32 40.172.1.76
  ip route 172.1.141.57/32 40.172.1.77
  ip route 172.1.141.57/32 40.172.1.78
  ip route 172.1.141.57/32 40.172.1.79
  ip route 172.1.141.57/32 40.172.1.80
  ip route 172.1.141.57/32 40.172.1.81
  ip route 172.1.141.57/32 40.172.1.82
  ip route 172.1.141.57/32 40.172.1.83
  ip route 172.1.141.57/32 40.172.1.84
  ip route 172.1.141.57/32 40.172.1.85
  ip route 172.1.141.57/32 40.172.1.86
  ip route 172.1.141.57/32 40.172.1.87
  ip route 172.1.141.57/32 40.172.1.88
  ip route 172.1.141.57/32 40.172.1.89
  ip route 172.1.141.57/32 40.172.1.90
  ip route 172.1.141.57/32 40.172.1.91
  ip route 172.1.141.57/32 40.172.1.92
  ip route 172.1.141.57/32 40.172.1.93
  ip route 172.1.141.57/32 40.172.1.94
  ip route 172.1.141.57/32 40.172.1.95
  ip route 172.1.141.57/32 40.172.1.96
  ip route 172.1.141.57/32 40.172.1.97
  ip route 172.1.141.57/32 40.172.1.98
  ip route 172.1.141.57/32 40.172.1.99
  ip route 172.1.141.57/32 40.172.1.100
  ip route 172.1.141.57/32 40.172.1.101
  ip route 172.1.141.57/32 40.172.1.102
  ip route 172.1.141.57/32 40.172.1.103
  ip route 172.1.141.57/32 40.172.1.104
  ip route 172.1.141.57/32 40.172.1.105
  ip route 172.1.141.57/32 40.172.1.106
  ip route 172.1.141.57/32 40.172.1.107
  ip route 172.1.141.57/32 40.172.1.108
  ip route 172.1.141.57/32 40.172.1.109
  ip route 172.1.141.57/32 40.172.1.110
  ip route 172.1.141.57/32 40.172.1.111
  ip route 172.1.141.57/32 40.172.1.112
  ip route 172.1.141.57/32 40.172.1.113
  ip route 172.1.141.57/32 40.172.1.114
  ip route 172.1.141.57/32 40.172.1.115
  ip route 172.1.141.57/32 40.172.1.116
  ip route 172.1.141.57/32 40.172.1.117
  ip route 172.1.141.57/32 40.172.1.118
  ip route 172.1.141.59/32 40.172.1.203
  ip route 172.1.141.59/32 40.172.1.204
  ip route 172.1.141.59/32 40.172.1.205
  ip route 172.1.141.59/32 40.172.1.206
  ip route 172.1.141.59/32 40.172.1.207
  ip route 172.1.141.59/32 40.172.1.208
  ip route 172.1.141.59/32 40.172.1.209
  ip route 172.1.141.59/32 40.172.1.210
  ip route 172.1.141.59/32 40.172.1.211
  ip route 172.1.141.59/32 40.172.1.212
  ip route 172.1.141.59/32 40.172.1.213
  ip route 172.1.141.59/32 40.172.1.214
  ip route 172.1.141.59/32 40.172.1.215
  ip route 172.1.141.59/32 40.172.1.216
  ip route 172.1.141.59/32 40.172.1.217
  ip route 172.1.141.59/32 40.172.1.218
  ip route 172.1.141.59/32 40.172.1.219
  ip route 172.1.141.59/32 40.172.1.220
  ip route 172.1.141.59/32 40.172.1.221
  ip route 172.1.141.59/32 40.172.1.222
  ip route 172.1.141.59/32 40.172.1.223
  ip route 172.1.141.59/32 40.172.1.224
  ip route 172.1.141.59/32 40.172.1.225
  ip route 172.1.141.59/32 40.172.1.226
  ip route 172.1.141.59/32 40.172.1.227
  ip route 172.1.141.59/32 40.172.1.228
  ip route 172.1.141.59/32 40.172.1.229
  ip route 172.1.141.59/32 40.172.1.230
  ip route 172.1.141.59/32 40.172.1.231
  ip route 172.1.141.59/32 40.172.1.232
  ip route 172.1.141.59/32 40.172.1.233
  ip route 172.1.141.59/32 40.172.1.234
  ip route 172.1.141.59/32 40.172.1.235
  ip route 172.1.141.59/32 40.172.1.236
  ip route 172.1.141.59/32 40.172.1.237
  ip route 172.1.141.59/32 40.172.1.238
  ip route 172.1.141.59/32 40.172.1.239
  ip route 172.1.141.59/32 40.172.1.240
  ip route 172.1.141.59/32 40.172.1.241
  ip route 172.1.141.59/32 40.172.1.242
  ip route 172.1.141.59/32 40.172.1.243
  ip route 172.1.141.59/32 40.172.1.244
  ip route 172.1.141.59/32 40.172.1.245
  ip route 172.1.141.59/32 40.172.1.246
  ip route 172.1.141.176/32 172.1.141.166
  ip route 172.1.141.176/32 172.1.141.167
  ip route 172.6.1.0/24 172.1.141.225
  ip route 172.6.1.0/24 172.1.141.229
  ip route 172.27.65.0/24 172.1.141.225
  ip route 172.27.65.0/24 172.1.141.229
  ip route 223.39.6.142/32 172.1.141.137
  ip route 223.39.6.143/32 172.1.141.137
  ip route 223.39.6.144/32 172.1.141.141
  ip route 223.39.6.145/32 172.1.141.141
  ip route 223.39.6.148/32 172.1.141.145
  ip route 223.39.6.149/32 172.1.141.145
 !
 address-family ipv6 unicast
  route-target export 10:10 evpn
  route-target import 10:10 evpn
  ipv6 route static bfd fd00:0:0:47::410 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::411 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::412 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::413 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::414 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::415 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::416 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::417 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::418 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::419 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::41a fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::41b fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::41c fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::41d fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::41e fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::41f fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::420 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::421 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::422 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::423 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::424 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::425 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::426 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::427 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::428 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::429 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::42a fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::42b fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::42c fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::42d fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::42e fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::42f fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::430 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::431 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::432 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::433 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::434 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::435 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::436 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::437 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::438 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::439 fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::43a fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::43b fd00:0:0:47::402 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47:c000:: fd00::47:c000:0:0:1 interval 50 min-rx 50 multiplier 4
  ipv6 route static bfd fd00::47:c000:0:0:2 fd00::47:c000:0:0:3 interval 50 min-rx 50 multiplier 4
  ipv6 route static bfd fd00::47:c001:0:0:21 fd00::47:c001:0:0:22 interval 50 min-rx 50 multiplier 4
  ipv6 route static bfd fd00::47:c001:0:0:25 fd00::47:c001:0:0:26 interval 50 min-rx 50 multiplier 4
  ipv6 route static bfd fd00::47:c001:0:0:29 fd00::47:c001:0:0:2a interval 50 min-rx 50 multiplier 4
  ipv6 route static bfd fd00::47:c001:0:0:31 fd00::47:c001:0:0:32 interval 50 min-rx 50 multiplier 4
  ipv6 route ::/0 fd00:0:0:47:c000::
  ipv6 route ::/0 fd00::47:c000:0:0:2
  ipv6 route 2001:2d8:2fa2:c000::/53 fd00::47:c001:0:0:21
  ipv6 route 2001:2d8:2fa2:c800::/53 fd00::47:c001:0:0:21
  ipv6 route 2001:2d8:2fa2:d000::/52 fd00::47:c001:0:0:29
  ipv6 route 2001:2d8:2fa2:d000::/53 fd00::47:c001:0:0:25
  ipv6 route 2001:2d8:2fa2:d800::/53 fd00::47:c001:0:0:25
  ipv6 route 2001:2d8:ff00::/40 fd00::47:0:0:f000:101
  ipv6 route 2001:2d8:ff00::/40 fd00::47:0:0:f000:111
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::410
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::411
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::412
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::413
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::414
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::415
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::416
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::417
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::418
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::419
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::41a
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::41b
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::41c
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::41d
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::41e
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::41f
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::420
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::421
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::422
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::423
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::424
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::425
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::426
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::427
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::428
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::429
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::42a
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::42b
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::42c
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::42d
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::42e
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::42f
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::430
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::431
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::432
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::433
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::434
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::435
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::436
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::437
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::438
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::439
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::43a
  ipv6 route fd00:0:0:47::201/128 fd00:0:0:47::43b
  ipv6 route fd00:0:0:47::510/128 fd00:0:0:47::410
  ipv6 route fd00:0:0:47::511/128 fd00:0:0:47::411
  ipv6 route fd00:0:0:47::512/128 fd00:0:0:47::412
  ipv6 route fd00:0:0:47::513/128 fd00:0:0:47::413
  ipv6 route fd00:0:0:47::514/128 fd00:0:0:47::414
  ipv6 route fd00:0:0:47::515/128 fd00:0:0:47::415
  ipv6 route fd00:0:0:47::516/128 fd00:0:0:47::416
  ipv6 route fd00:0:0:47::517/128 fd00:0:0:47::417
  ipv6 route fd00:0:0:47::518/128 fd00:0:0:47::418
  ipv6 route fd00:0:0:47::519/128 fd00:0:0:47::419
  ipv6 route fd00:0:0:47::51a/128 fd00:0:0:47::41a
  ipv6 route fd00:0:0:47::51b/128 fd00:0:0:47::41b
  ipv6 route fd00:0:0:47::51c/128 fd00:0:0:47::41c
  ipv6 route fd00:0:0:47::51d/128 fd00:0:0:47::41d
  ipv6 route fd00:0:0:47::51e/128 fd00:0:0:47::41e
  ipv6 route fd00:0:0:47::51f/128 fd00:0:0:47::41f
  ipv6 route fd00:0:0:47::520/128 fd00:0:0:47::420
  ipv6 route fd00:0:0:47::521/128 fd00:0:0:47::421
  ipv6 route fd00:0:0:47::522/128 fd00:0:0:47::422
  ipv6 route fd00:0:0:47::523/128 fd00:0:0:47::423
  ipv6 route fd00:0:0:47::524/128 fd00:0:0:47::424
  ipv6 route fd00:0:0:47::525/128 fd00:0:0:47::425
  ipv6 route fd00:0:0:47::526/128 fd00:0:0:47::426
  ipv6 route fd00:0:0:47::527/128 fd00:0:0:47::427
  ipv6 route fd00:0:0:47::528/128 fd00:0:0:47::428
  ipv6 route fd00:0:0:47::529/128 fd00:0:0:47::429
  ipv6 route fd00:0:0:47::52a/128 fd00:0:0:47::42a
  ipv6 route fd00:0:0:47::52b/128 fd00:0:0:47::42b
  ipv6 route fd00:0:0:47::52c/128 fd00:0:0:47::42c
  ipv6 route fd00:0:0:47::52d/128 fd00:0:0:47::42d
  ipv6 route fd00:0:0:47::52e/128 fd00:0:0:47::42e
  ipv6 route fd00:0:0:47::52f/128 fd00:0:0:47::42f
  ipv6 route fd00:0:0:47::530/128 fd00:0:0:47::430
  ipv6 route fd00:0:0:47::531/128 fd00:0:0:47::431
  ipv6 route fd00:0:0:47::532/128 fd00:0:0:47::432
  ipv6 route fd00:0:0:47::533/128 fd00:0:0:47::433
  ipv6 route fd00:0:0:47::534/128 fd00:0:0:47::434
  ipv6 route fd00:0:0:47::535/128 fd00:0:0:47::435
  ipv6 route fd00:0:0:47::536/128 fd00:0:0:47::436
  ipv6 route fd00:0:0:47::537/128 fd00:0:0:47::437
  ipv6 route fd00:0:0:47::538/128 fd00:0:0:47::438
  ipv6 route fd00:0:0:47::539/128 fd00:0:0:47::439
  ipv6 route fd00:0:0:47::53a/128 fd00:0:0:47::43a
  ipv6 route fd00:0:0:47::53b/128 fd00:0:0:47::43b
  ipv6 route fd00:22::/45 fd00::47:c001:0:0:21
  ipv6 route fd00:22:8::/45 fd00::47:c001:0:0:21
  ipv6 route fd00:150::/120 fd00:a01:a01:1::2
  ipv6 route 2001:2d8:2fa2:c000::/52 next-hop-vrf NAT01_VRF fd00:0:0:47::200
  ipv6 route 2001:2d8:2fa2:c000::fffe/128 next-hop-vrf NAT01_VRF fd00:606:2c05:1::2
  ipv6 route fd00:22::/44 next-hop-vrf NAT01_VRF fd00:0:0:47::200
 !
!
vrf mgmt-vrf
 address-family ipv4 unicast
  ip route 0.0.0.0/0 60.30.181.126
 !
 address-family ipv6 unicast
 !
!
vrf NAT01_VRF
 rd 198.19.96.93:5
 address-family ipv4 unicast
  route-target export 13:13 evpn
  route-target import 13:13 evpn
  ip route static bfd 40.172.1.139 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.140 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.141 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.142 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.143 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.144 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.145 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.146 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.147 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.148 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.149 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.150 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.151 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.152 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.153 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.154 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.155 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.156 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.157 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.158 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.159 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.160 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.161 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.162 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.163 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.164 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.165 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.166 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.167 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.168 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.169 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.170 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.171 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.172 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.173 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.174 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.175 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.176 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.177 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.178 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.179 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.180 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.181 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 40.172.1.182 40.172.1.130 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.141.105 172.1.141.106 interval 50 min-rx 50 multiplier 4
  ip route static bfd 172.1.141.113 172.1.141.114 interval 50 min-rx 50 multiplier 4
  ip route 0.0.0.0/0 172.1.141.113
  ip route 0.0.0.0/1 172.1.141.105 10
  ip route 10.101.0.0/16 40.255.255.13
  ip route 10.101.0.0/16 40.255.255.113
  ip route 10.102.0.0/16 40.255.255.13
  ip route 10.102.0.0/16 40.255.255.113
  ip route 10.103.0.0/16 40.255.255.13
  ip route 10.103.0.0/16 40.255.255.113
  ip route 10.104.0.0/16 40.255.255.13
  ip route 10.104.0.0/16 40.255.255.113
  ip route 10.105.0.0/16 40.255.255.13
  ip route 10.105.0.0/16 40.255.255.113
  ip route 10.106.0.0/16 40.255.255.13
  ip route 10.106.0.0/16 40.255.255.113
  ip route 10.107.0.0/16 40.255.255.13
  ip route 10.107.0.0/16 40.255.255.113
  ip route 10.108.0.0/16 40.255.255.13
  ip route 10.111.0.0/16 40.255.255.113
  ip route 10.112.0.0/16 40.255.255.113
  ip route 10.113.0.0/16 40.255.255.113
  ip route 10.114.0.0/16 40.255.255.113
  ip route 10.115.0.0/16 40.255.255.113
  ip route 10.116.0.0/16 40.255.255.113
  ip route 10.117.0.0/16 40.255.255.113
  ip route 40.173.1.139/32 40.172.1.139
  ip route 40.173.1.140/32 40.172.1.140
  ip route 40.173.1.141/32 40.172.1.141
  ip route 40.173.1.142/32 40.172.1.142
  ip route 40.173.1.143/32 40.172.1.143
  ip route 40.173.1.144/32 40.172.1.144
  ip route 40.173.1.145/32 40.172.1.145
  ip route 40.173.1.146/32 40.172.1.146
  ip route 40.173.1.147/32 40.172.1.147
  ip route 40.173.1.148/32 40.172.1.148
  ip route 40.173.1.149/32 40.172.1.149
  ip route 40.173.1.150/32 40.172.1.150
  ip route 40.173.1.151/32 40.172.1.151
  ip route 40.173.1.152/32 40.172.1.152
  ip route 40.173.1.153/32 40.172.1.153
  ip route 40.173.1.154/32 40.172.1.154
  ip route 40.173.1.155/32 40.172.1.155
  ip route 40.173.1.156/32 40.172.1.156
  ip route 40.173.1.157/32 40.172.1.157
  ip route 40.173.1.158/32 40.172.1.158
  ip route 40.173.1.159/32 40.172.1.159
  ip route 40.173.1.160/32 40.172.1.160
  ip route 40.173.1.161/32 40.172.1.161
  ip route 40.173.1.162/32 40.172.1.162
  ip route 40.173.1.163/32 40.172.1.163
  ip route 40.173.1.164/32 40.172.1.164
  ip route 40.173.1.165/32 40.172.1.165
  ip route 40.173.1.166/32 40.172.1.166
  ip route 40.173.1.167/32 40.172.1.167
  ip route 40.173.1.168/32 40.172.1.168
  ip route 40.173.1.169/32 40.172.1.169
  ip route 40.173.1.170/32 40.172.1.170
  ip route 40.173.1.171/32 40.172.1.171
  ip route 40.173.1.172/32 40.172.1.172
  ip route 40.173.1.173/32 40.172.1.173
  ip route 40.173.1.174/32 40.172.1.174
  ip route 40.173.1.175/32 40.172.1.175
  ip route 40.173.1.176/32 40.172.1.176
  ip route 40.173.1.177/32 40.172.1.177
  ip route 40.173.1.178/32 40.172.1.178
  ip route 40.173.1.179/32 40.172.1.179
  ip route 40.173.1.180/32 40.172.1.180
  ip route 40.173.1.181/32 40.172.1.181
  ip route 40.173.1.182/32 40.172.1.182
  ip route 128.0.0.0/1 172.1.141.105 10
  ip route 172.1.141.58/32 40.172.1.139
  ip route 172.1.141.58/32 40.172.1.140
  ip route 172.1.141.58/32 40.172.1.141
  ip route 172.1.141.58/32 40.172.1.142
  ip route 172.1.141.58/32 40.172.1.143
  ip route 172.1.141.58/32 40.172.1.144
  ip route 172.1.141.58/32 40.172.1.145
  ip route 172.1.141.58/32 40.172.1.146
  ip route 172.1.141.58/32 40.172.1.147
  ip route 172.1.141.58/32 40.172.1.148
  ip route 172.1.141.58/32 40.172.1.149
  ip route 172.1.141.58/32 40.172.1.150
  ip route 172.1.141.58/32 40.172.1.151
  ip route 172.1.141.58/32 40.172.1.152
  ip route 172.1.141.58/32 40.172.1.153
  ip route 172.1.141.58/32 40.172.1.154
  ip route 172.1.141.58/32 40.172.1.155
  ip route 172.1.141.58/32 40.172.1.156
  ip route 172.1.141.58/32 40.172.1.157
  ip route 172.1.141.58/32 40.172.1.158
  ip route 172.1.141.58/32 40.172.1.159
  ip route 172.1.141.58/32 40.172.1.160
  ip route 172.1.141.58/32 40.172.1.161
  ip route 172.1.141.58/32 40.172.1.162
  ip route 172.1.141.58/32 40.172.1.163
  ip route 172.1.141.58/32 40.172.1.164
  ip route 172.1.141.58/32 40.172.1.165
  ip route 172.1.141.58/32 40.172.1.166
  ip route 172.1.141.58/32 40.172.1.167
  ip route 172.1.141.58/32 40.172.1.168
  ip route 172.1.141.58/32 40.172.1.169
  ip route 172.1.141.58/32 40.172.1.170
  ip route 172.1.141.58/32 40.172.1.171
  ip route 172.1.141.58/32 40.172.1.172
  ip route 172.1.141.58/32 40.172.1.173
  ip route 172.1.141.58/32 40.172.1.174
  ip route 172.1.141.58/32 40.172.1.175
  ip route 172.1.141.58/32 40.172.1.176
  ip route 172.1.141.58/32 40.172.1.177
  ip route 172.1.141.58/32 40.172.1.178
  ip route 172.1.141.58/32 40.172.1.179
  ip route 172.1.141.58/32 40.172.1.180
  ip route 172.1.141.58/32 40.172.1.181
  ip route 172.1.141.58/32 40.172.1.182
 !
 address-family ipv6 unicast
  route-target export 13:13 evpn
  route-target import 13:13 evpn
  ipv6 route static bfd fd00:0:0:47::10 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::11 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::12 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::13 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::14 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::15 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::16 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::17 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::18 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::19 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::1a fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::1b fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::1c fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::1d fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::1e fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::1f fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::20 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::21 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::22 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::23 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::24 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::25 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::26 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::27 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::28 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::29 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::2a fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::2b fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::2c fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::2d fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::2e fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::2f fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::30 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::31 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::32 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::33 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::34 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::35 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::36 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::37 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::38 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::39 fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::3a fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00:0:0:47::3b fd00:0:0:47::2 interval 200 min-rx 200 multiplier 3
  ipv6 route static bfd fd00::47:c001:0:0:1 fd00::47:c001:0:0:2 interval 50 min-rx 50 multiplier 4
  ipv6 route static bfd fd00::47:c001:0:0:9 fd00::47:c001:0:0:a interval 50 min-rx 50 multiplier 4
  ipv6 route ::/1 fd00::47:c001:0:0:1 10
  ipv6 route 64:ff9b::/96 fd00::47:c001:0:0:9
  ipv6 route 64:ff9b::/97 fd00::47:c001:0:0:1
  ipv6 route 64:ff9b::8000:0/97 fd00::47:c001:0:0:1
  ipv6 route 2001:2d8:ff00::/40 fd00::47:0:0:f000:1
  ipv6 route 2001:2d8:ff00::/40 fd00::47:0:0:f000:11
  ipv6 route 8000::/1 fd00::47:c001:0:0:1 10
  ipv6 route fd00:0:0:47::110/128 fd00:0:0:47::10
  ipv6 route fd00:0:0:47::111/128 fd00:0:0:47::11
  ipv6 route fd00:0:0:47::112/128 fd00:0:0:47::12
  ipv6 route fd00:0:0:47::113/128 fd00:0:0:47::13
  ipv6 route fd00:0:0:47::114/128 fd00:0:0:47::14
  ipv6 route fd00:0:0:47::115/128 fd00:0:0:47::15
  ipv6 route fd00:0:0:47::116/128 fd00:0:0:47::16
  ipv6 route fd00:0:0:47::117/128 fd00:0:0:47::17
  ipv6 route fd00:0:0:47::118/128 fd00:0:0:47::18
  ipv6 route fd00:0:0:47::119/128 fd00:0:0:47::19
  ipv6 route fd00:0:0:47::11a/128 fd00:0:0:47::1a
  ipv6 route fd00:0:0:47::11b/128 fd00:0:0:47::1b
  ipv6 route fd00:0:0:47::11c/128 fd00:0:0:47::1c
  ipv6 route fd00:0:0:47::11d/128 fd00:0:0:47::1d
  ipv6 route fd00:0:0:47::11e/128 fd00:0:0:47::1e
  ipv6 route fd00:0:0:47::11f/128 fd00:0:0:47::1f
  ipv6 route fd00:0:0:47::120/128 fd00:0:0:47::20
  ipv6 route fd00:0:0:47::121/128 fd00:0:0:47::21
  ipv6 route fd00:0:0:47::122/128 fd00:0:0:47::22
  ipv6 route fd00:0:0:47::123/128 fd00:0:0:47::23
  ipv6 route fd00:0:0:47::124/128 fd00:0:0:47::24
  ipv6 route fd00:0:0:47::125/128 fd00:0:0:47::25
  ipv6 route fd00:0:0:47::126/128 fd00:0:0:47::26
  ipv6 route fd00:0:0:47::127/128 fd00:0:0:47::27
  ipv6 route fd00:0:0:47::128/128 fd00:0:0:47::28
  ipv6 route fd00:0:0:47::129/128 fd00:0:0:47::29
  ipv6 route fd00:0:0:47::12a/128 fd00:0:0:47::2a
  ipv6 route fd00:0:0:47::12b/128 fd00:0:0:47::2b
  ipv6 route fd00:0:0:47::12c/128 fd00:0:0:47::2c
  ipv6 route fd00:0:0:47::12d/128 fd00:0:0:47::2d
  ipv6 route fd00:0:0:47::12e/128 fd00:0:0:47::2e
  ipv6 route fd00:0:0:47::12f/128 fd00:0:0:47::2f
  ipv6 route fd00:0:0:47::130/128 fd00:0:0:47::30
  ipv6 route fd00:0:0:47::131/128 fd00:0:0:47::31
  ipv6 route fd00:0:0:47::132/128 fd00:0:0:47::32
  ipv6 route fd00:0:0:47::133/128 fd00:0:0:47::33
  ipv6 route fd00:0:0:47::134/128 fd00:0:0:47::34
  ipv6 route fd00:0:0:47::135/128 fd00:0:0:47::35
  ipv6 route fd00:0:0:47::136/128 fd00:0:0:47::36
  ipv6 route fd00:0:0:47::137/128 fd00:0:0:47::37
  ipv6 route fd00:0:0:47::138/128 fd00:0:0:47::38
  ipv6 route fd00:0:0:47::139/128 fd00:0:0:47::39
  ipv6 route fd00:0:0:47::13a/128 fd00:0:0:47::3a
  ipv6 route fd00:0:0:47::13b/128 fd00:0:0:47::3b
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::10
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::11
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::12
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::13
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::14
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::15
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::16
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::17
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::18
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::19
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::1a
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::1b
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::1c
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::1d
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::1e
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::1f
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::20
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::21
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::22
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::23
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::24
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::25
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::26
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::27
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::28
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::29
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::2a
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::2b
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::2c
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::2d
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::2e
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::2f
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::30
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::31
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::32
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::33
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::34
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::35
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::36
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::37
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::38
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::39
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::3a
  ipv6 route fd00:0:0:47::200/128 fd00:0:0:47::3b
  ipv6 route fd00:22:8000::/45 fd00::47:c001:0:0:1
  ipv6 route fd00:22:8000::/45 fd00::47:c001:0:0:2d
  ipv6 route ::/0 next-hop-vrf GLOBAL_VRF fd00::47:0:0:f000:101
  ipv6 route ::/0 next-hop-vrf GLOBAL_VRF fd00::47:0:0:f000:111
 !
!
vrf SBI_VRF
 rd 198.19.96.93:2
 address-family ipv4 unicast
  route-target export 11:11 evpn
  route-target import 11:11 evpn
  ip route static bfd 40.255.255.5 40.255.255.6
  ip route static bfd 172.1.140.69 172.1.140.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.70 172.1.140.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.71 172.1.140.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.72 172.1.140.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.73 172.1.140.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.74 172.1.140.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.75 172.1.140.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.76 172.1.140.66 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.148 172.1.140.146 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.167 172.1.140.162 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.168 172.1.140.162 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.169 172.1.140.162 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.170 172.1.140.162 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.171 172.1.140.162 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.172 172.1.140.162 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.173 172.1.140.162 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.174 172.1.140.162 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.175 172.1.140.162 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.176 172.1.140.162 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.177 172.1.140.162 interval 200 min-rx 200 multiplier 3
  ip route static bfd 172.1.140.225 172.1.140.226 interval 50 min-rx 50 multiplier 4
  ip route static bfd 172.1.140.229 172.1.140.230 interval 50 min-rx 50 multiplier 4
  ip route 0.0.0.0/0 172.1.140.225
  ip route 0.0.0.0/0 172.1.140.229
  ip route 10.101.0.0/16 40.255.255.5
  ip route 10.101.0.0/16 40.255.255.105
  ip route 10.102.0.0/16 40.255.255.5
  ip route 10.102.0.0/16 40.255.255.105
  ip route 10.103.0.0/16 40.255.255.5
  ip route 10.103.0.0/16 40.255.255.105
  ip route 10.104.0.0/16 40.255.255.5
  ip route 10.104.0.0/16 40.255.255.105
  ip route 10.105.0.0/16 40.255.255.5
  ip route 10.105.0.0/16 40.255.255.105
  ip route 10.106.0.0/16 40.255.255.5
  ip route 10.106.0.0/16 40.255.255.105
  ip route 10.107.0.0/16 40.255.255.105
  ip route 10.108.0.0/16 40.255.255.105
  ip route 10.111.0.0/16 40.255.255.5
  ip route 10.111.0.0/16 40.255.255.105
  ip route 10.112.0.0/16 40.255.255.5
  ip route 10.112.0.0/16 40.255.255.105
  ip route 10.113.0.0/16 40.255.255.5
  ip route 10.113.0.0/16 40.255.255.105
  ip route 10.114.0.0/16 40.255.255.5
  ip route 10.114.0.0/16 40.255.255.105
  ip route 10.115.0.0/16 40.255.255.5
  ip route 10.115.0.0/16 40.255.255.105
  ip route 10.116.0.0/16 40.255.255.5
  ip route 10.116.0.0/16 40.255.255.105
  ip route 10.117.0.0/16 40.255.255.105
  ip route 10.118.0.0/16 40.255.255.105
  ip route 40.172.140.1/32 172.1.140.165
  ip route 40.172.140.2/32 172.1.140.166
  ip route 40.172.140.8/32 172.1.140.167
  ip route 40.172.140.8/32 172.1.140.168
  ip route 40.172.140.8/32 172.1.140.169
  ip route 40.172.140.8/32 172.1.140.170
  ip route 40.172.140.8/32 172.1.140.171
  ip route 40.172.140.8/32 172.1.140.172
  ip route 40.172.140.8/32 172.1.140.173
  ip route 40.172.140.8/32 172.1.140.174
  ip route 172.1.75.0/24 172.1.140.241
  ip route 172.1.75.0/24 172.1.140.245
  ip route 172.1.140.80/28 172.1.140.69
  ip route 172.1.140.80/28 172.1.140.70
  ip route 172.1.140.80/28 172.1.140.71
  ip route 172.1.140.80/28 172.1.140.72
  ip route 172.1.140.80/28 172.1.140.73
  ip route 172.1.140.80/28 172.1.140.74
  ip route 172.1.140.80/28 172.1.140.75
  ip route 172.1.140.80/28 172.1.140.76
  ip route 172.1.140.84/30 172.1.140.148
  ip route 172.1.218.252/30 172.1.140.241
  ip route 172.1.218.252/30 172.1.140.249
  ip route 172.1.237.0/24 172.1.140.241
  ip route 172.1.237.0/24 172.1.140.249
  ip route 172.1.238.0/24 172.1.140.241
  ip route 172.1.238.0/24 172.1.140.249
  ip route 172.27.65.242/32 172.1.140.241
  ip route 172.27.65.242/32 172.1.140.245
  ip route 172.28.81.0/24 172.1.140.241
  ip route 172.28.81.0/24 172.1.140.245
 !
 address-family ipv6 unicast
  route-target export 11:11 evpn
  route-target import 11:11 evpn
 !
!
ssh server max-idle-timeout 300
ssh server key rsa 2048
ssh server key ecdsa 256
ssh server key dsa
ssh server use-vrf mgmt-vrf
telnet server use-vrf default-vrf shutdown
telnet server use-vrf mgmt-vrf shutdown
http server use-vrf default-vrf
http server use-vrf mgmt-vrf
banner login "  ----------------------------------------------------------------------\n    !!!! WARNING !!!!WARNING !!!!WARNING !!!! WARNING !!!!\n  ----------------------------------------------------------------------\n     Unauthorized use of this system is prohibited.\n     All access to this equipment is logged.\n     Disconnect immediately if you are not an authorized user.\n     Violators will be prosecuted to the fullest extent of the law.\n  ----------------------------------------------------------------------\n    !!!! WARNING !!!!WARNING !!!!WARNING !!!! WARNING !!!!\n  ----------------------------------------------------------------------\n"
role name admin desc Administrator
role name user desc User
aaa authentication login oauth2 local-auth-fallback
aaa accounting exec default start-stop none
aaa accounting commands default start-stop none
aaa authorization command none
service password-encryption
username admin password $6$iBVWrBRXxJlXWoXy$9uv3K4yZ5rLGfiW.r931npNAcfAayIpkA.GtiJaPmA0mzQi/.nrNZvnNjBXinr2WPz9eA1O.0NZR5BRywnD7q1 encryption-level 10 role admin desc Administrator
username sdn_suser password $6$ginQxSfnRhKsb04d$V3.qpTrHSFVfqjcCUVnkSTNwam1QQIDs5Jn0ljpR1ZR2SezbaNFQbSDzVLvmLZ0kCFZM63DesEeVDg8NiF8p.. encryption-level 10 role user
username user password $6$1QCZK/4T3p7ATJhI$LGqDn50NQPJbcpPB9hVyKMkatCC.fkKlCqTPeyA22hDYruvlLmHfWODYuBkXJtVA.8zFHuj9RWVpLXiyurv600 encryption-level 10 role user desc User
snmp-server sys-descr "Extreme SLX9740-40C Switch/Router"
snmp-server engineID local 80:0:6:34:b2:7e:88:0:10:a:e9:d3:25
snmp-server community $9$4qOrvwtKcShb/JCoXGb7nw==
snmp-server community $9$pdTj0ueN2KcUO36QTF1PIQ== groupname read-only
snmp-server community $9$uZrcr+xAbndcYWJPaOUb1g==
snmp-server host 60.11.26.5 $9$uZrcr+xAbndcYWJPaOUb1g==
 source-interface management chassis-ip
!
snmp-server host 60.11.8.28 $9$4qOrvwtKcShb/JCoXGb7nw==
 source-interface management chassis-ip
!
snmp-server user efav3User groupname efav3Group auth sha auth-password LlgMJxy1dZfAfPydPDdFTNNtkpw= priv AES128 priv-password 53IJSJKiK/A6UFLVd9+ZkkclIJs= encrypted
snmp-server user skt groupname SKT auth sha auth-password xLMpyXWUPBeK7XI9pyaETkwBPgE= encrypted
snmp-server user sktip groupname SKT auth sha auth-password VhsrToxSUvuPh6nWXwFbuDALXrU= encrypted
snmp-server v3host 60.30.181.98 efav3User
 source-interface management chassis-ip
!
snmp-server view efav3View 1.3.6.1 included
snmp-server view read-only 1.3.6.1 included
snmp-server view view2 1.3.6.1 included
snmp-server group efav3Group v3 notify efav3View
snmp-server group read-only v2c read view2
snmp-server group SKT v3 auth read read-only notify read-only
ntp server 60.30.181.97
ntp server 60.30.181.101
ntp server 127.127.1.0
tunneled-arp-trap enable
ip mtu 9100
ip access-list extended ANTI-SPOOFING
 seq 10 permit udp host 0.0.0.0 eq bootpc host 255.255.255.255 eq bootps count
 seq 20 deny ip host 0.0.0.0 any count log
 seq 30 deny ip host 255.255.255.255 any count log
 seq 100 permit ip any any count
!
ip access-list extended MGMT-SNMP-PERMIT
 seq 10 permit udp host 0.0.0.0 eq bootpc host 255.255.255.255 eq bootps count
 seq 20 permit ip host 60.30.181.115 any count
 seq 30 permit ip host 60.30.181.97 any count
 seq 40 permit ip host 60.30.181.98 any count
 seq 50 permit ip host 60.30.181.99 any count
 seq 60 permit ip host 60.30.181.100 any count
 seq 70 permit ip host 60.30.181.116 any count
 seq 80 permit ip host 60.30.181.117 any count
 seq 90 permit ip host 60.30.181.118 any count
 seq 100 permit ip host 60.30.181.120 any count
 seq 110 permit ip host 60.30.181.121 any count
 seq 120 permit ip host 60.30.131.209 any count
 seq 130 permit ip host 60.30.181.101 any count
 seq 140 permit ip host 60.30.181.102 any count
 seq 150 permit ip host 60.30.181.103 any count
 seq 160 permit ip host 60.30.181.104 any count
 seq 170 permit ip host 60.30.181.105 any count
 seq 180 permit ip host 60.30.181.106 any count
 seq 190 permit ip host 60.30.181.107 any count
 seq 200 permit ip host 60.30.181.108 any count
 seq 210 permit ip host 60.30.181.109 any count
 seq 220 permit ip host 60.30.181.110 any count
 seq 230 permit ip host 60.30.181.111 any count
 seq 240 permit ip host 60.30.181.112 any count
 seq 250 permit ip host 60.30.181.113 any count
 seq 260 permit ip host 60.30.181.114 any count
 seq 1000 deny ip any any count
!
ip access-list extended SYSTEM_to_UE_Deny
 seq 10 deny icmp 172.16.0.0 0.15.255.255 211.235.64.0 0.0.63.255 count log
 seq 20 deny icmp 192.168.0.0 0.0.255.255 211.235.64.0 0.0.63.255 count log
 seq 1000 permit ip any any count
!
ip access-list extended UE_to_SYSTEM_Deny
 seq 10 deny ip host 0.0.0.0 any count log
 seq 20 deny ip host 255.255.255.255 any count log
 seq 30 deny icmp 211.235.64.0 0.0.63.255 172.16.0.0 0.15.255.255 count log
 seq 40 deny icmp 211.235.64.0 0.0.63.255 192.168.0.0 0.0.255.255 count log
 seq 50 deny tcp 211.235.64.0 0.0.63.255 172.16.0.0 0.15.255.255 eq telnet count log
 seq 60 deny tcp 211.235.64.0 0.0.63.255 192.168.0.0 0.0.255.255 eq telnet count log
 seq 70 deny tcp 211.235.64.0 0.0.63.255 172.16.0.0 0.15.255.255 eq 22 count log
 seq 80 deny tcp 211.235.64.0 0.0.63.255 192.168.0.0 0.0.255.255 eq 22 count log
 seq 90 deny udp 211.235.64.0 0.0.63.255 172.16.0.0 0.15.255.255 eq snmp count log
 seq 100 deny udp 211.235.64.0 0.0.63.255 192.168.0.0 0.0.255.255 eq snmp count log
 seq 110 deny udp 211.235.64.0 0.0.63.255 172.16.0.0 0.15.255.255 eq 162 count log
 seq 120 deny udp 211.235.64.0 0.0.63.255 192.168.0.0 0.0.255.255 eq 162 count log
 seq 1000 permit ip any any count
!
vlan 1
!
vlan 801
 description Tenant L2 Extended VLAN
!
vlan 810
 router-interface Ve 810
 description Tenant L2 Extended VLAN
!
vlan 811
 router-interface Ve 811
 description Tenant L2 Extended VLAN
!
vlan 820
 router-interface Ve 820
 description Tenant L2 Extended VLAN
!
vlan 821
 router-interface Ve 821
 description Tenant L2 Extended VLAN
!
vlan 822
 router-interface Ve 822
 description Tenant L2 Extended VLAN
!
vlan 825
 router-interface Ve 825
 description Tenant L2 Extended VLAN
!
vlan 826
 router-interface Ve 826
 description Tenant L2 Extended VLAN
!
vlan 827
 router-interface Ve 827
 description Tenant L2 Extended VLAN
!
vlan 831
 router-interface Ve 831
 description Tenant L2 Extended VLAN
!
vlan 832
 router-interface Ve 832
 description Tenant L2 Extended VLAN
!
vlan 833
 router-interface Ve 833
 description Tenant L2 Extended VLAN
!
vlan 834
 router-interface Ve 834
 description Tenant L2 Extended VLAN
!
vlan 835
 router-interface Ve 835
 description Tenant L2 Extended VLAN
!
vlan 836
 router-interface Ve 836
 description Tenant L2 Extended VLAN
!
vlan 840
 router-interface Ve 840
 description Tenant L2 Extended VLAN
!
vlan 841
 router-interface Ve 841
 description Tenant L2 Extended VLAN
!
vlan 842
 router-interface Ve 842
 description Tenant L2 Extended VLAN
!
vlan 931
 router-interface Ve 931
 description Tenant L3 Hand-off VLAN
!
vlan 932
 router-interface Ve 932
 description Tenant L3 Hand-off VLAN
!
vlan 933
 router-interface Ve 933
 description Tenant L3 Hand-off VLAN
!
vlan 934
 router-interface Ve 934
 description Tenant L3 Hand-off VLAN
!
vlan 935
 router-interface Ve 935
 description Tenant L3 Hand-off VLAN
!
vlan 936
 router-interface Ve 936
 description Tenant L3 Hand-off VLAN
!
vlan 937
 router-interface Ve 937
 description Tenant L3 Hand-off VLAN
!
vlan 938
 router-interface Ve 938
 description Tenant L3 Hand-off VLAN
!
vlan 939
 router-interface Ve 939
 description Tenant L3 Hand-off VLAN
!
vlan 940
 router-interface Ve 940
 description Tenant L3 Hand-off VLAN
!
vlan 941
 router-interface Ve 941
 description Tenant L3 Hand-off VLAN
!
vlan 942
 router-interface Ve 942
 description Tenant L3 Hand-off VLAN
!
vlan 943
 router-interface Ve 943
 description Tenant L3 Hand-off VLAN
!
vlan 944
 router-interface Ve 944
 description Tenant L3 Hand-off VLAN
!
vlan 945
 router-interface Ve 945
 description Tenant L3 Hand-off VLAN
!
vlan 946
 router-interface Ve 946
 description Tenant L3 Hand-off VLAN
!
vlan 947
 router-interface Ve 947
 description Tenant L3 Hand-off VLAN
!
vlan 948
 router-interface Ve 948
 description Tenant L3 Hand-off VLAN
!
vlan 949
 router-interface Ve 949
 description Tenant L3 Hand-off VLAN
!
vlan 950
 router-interface Ve 950
 description Tenant L3 Hand-off VLAN
!
vlan 951
 router-interface Ve 951
 description Tenant L3 Hand-off VLAN
!
vlan 952
 router-interface Ve 952
 description Tenant L3 Hand-off VLAN
!
vlan 953
 router-interface Ve 953
 description Tenant L3 Hand-off VLAN
!
vlan 954
 router-interface Ve 954
 description Tenant L3 Hand-off VLAN
!
vlan 955
 router-interface Ve 955
 description Tenant L3 Hand-off VLAN
!
vlan 956
 router-interface Ve 956
 description Tenant L3 Hand-off VLAN
!
vlan 965
 router-interface Ve 965
 description Tenant L3 Hand-off VLAN
!
vlan 966
 router-interface Ve 966
 description Tenant L3 Hand-off VLAN
!
vlan 970
 router-interface Ve 970
 description Tenant L3 Hand-off VLAN
!
vlan 972
 router-interface Ve 972
 description Tenant L3 Hand-off VLAN
!
vlan 974
 router-interface Ve 974
 description Tenant L3 Hand-off VLAN
!
vlan 976
 router-interface Ve 976
 description Tenant L3 Hand-off VLAN
!
vlan 978
 router-interface Ve 978
 description Tenant L3 Hand-off VLAN
!
vlan 980
 router-interface Ve 980
 description Tenant L3 Hand-off VLAN
!
vlan 981
 router-interface Ve 981
 description Tenant L3 Hand-off VLAN
!
vlan 983
 router-interface Ve 983
 description Tenant L3 Hand-off VLAN
!
vlan 990
 router-interface Ve 990
 description Tenant L3 Hand-off VLAN
!
vlan 991
 router-interface Ve 991
 description Tenant L3 Hand-off VLAN
!
vlan 992
 router-interface Ve 992
 description Tenant L3 Hand-off VLAN
!
vlan 993
 router-interface Ve 993
 description Tenant L3 Hand-off VLAN
!
vlan 995
 router-interface Ve 995
 description Tenant L3 Hand-off VLAN
!
vlan 996
 router-interface Ve 996
 description Tenant L3 Hand-off VLAN
!
vlan 997
 router-interface Ve 997
 description Tenant L3 Hand-off VLAN
!
vlan 998
 router-interface Ve 998
 description Tenant L3 Hand-off VLAN
!
vlan 1003
 router-interface Ve 1003
 description Tenant L2 Extended VLAN
!
vlan 1004
 router-interface Ve 1004
 description Tenant L2 Extended VLAN
!
vlan 1005
 router-interface Ve 1005
 description Tenant L2 Extended VLAN
!
vlan 1007
 router-interface Ve 1007
 description Tenant L2 Extended VLAN
!
vlan 1008
 router-interface Ve 1008
 description Tenant L2 Extended VLAN
!
protocol lldp
 system-description Extreme SLX9740-40C Switch/Router
!
vlan dot1q tag native
ip router-id 198.19.96.93
class-map default
!
no protocol vrrp
no protocol vrrp-extended
ip anycast-gateway-mac 0201.0101.0101
evpn TB2
 route-target both auto ignore-as
 rd auto
 duplicate-mac-timer 5 max-count 3
 vlan add 801,810-811,820-822,825-827,831-836,840-842,1003-1005,1007-1008
!
router bgp
 local-as 66011
 capability as4-enable
 fast-external-fallover
 listen-limit 255
 neighbor k8s_ext_svc peer-group
 neighbor k8s_ext_svc remote-as 66111
 neighbor k8s_ext_svc bfd
 neighbor k8s_ext_svc bfd interval 200 min-rx 200 multiplier 3
 neighbor smf_global_link peer-group
 neighbor smf_global_link remote-as 64534
 neighbor spine-group peer-group
 neighbor spine-group remote-as 66000
 neighbor spine-group description To Spine
 neighbor spine-group bfd
 neighbor upf_ims peer-group
 neighbor upf_ims remote-as 66164
 neighbor upf_ims bfd
 neighbor upf_ims bfd interval 200 min-rx 200 multiplier 3
 neighbor upftb_nat01 peer-group
 neighbor upftb_nat01 remote-as 66144
 neighbor upftb_nat01 bfd
 neighbor upftb_nat01 bfd interval 200 min-rx 200 multiplier 3
 neighbor upftb_s1u_n3 peer-group
 neighbor upftb_s1u_n3 remote-as 66141
 neighbor upftb_s1u_n3 bfd
 neighbor upftb_s1u_n3 bfd interval 200 min-rx 200 multiplier 3
 neighbor upftb_s5u_n9 peer-group
 neighbor upftb_s5u_n9 remote-as 66143
 neighbor upftb_s5u_n9 bfd
 neighbor upftb_s5u_n9 bfd interval 200 min-rx 200 multiplier 3
 neighbor upftb_sx_n4 peer-group
 neighbor upftb_sx_n4 remote-as 66142
 neighbor upftb_sx_n4 bfd
 neighbor upftb_sx_n4 bfd interval 200 min-rx 200 multiplier 3
 neighbor vpngw_media peer-group
 neighbor vpngw_media remote-as 66131
 neighbor vpngw_media ebgp-multihop
 neighbor vpngw_signal peer-group
 neighbor vpngw_signal remote-as 66132
 neighbor vpngw_signal ebgp-multihop
 neighbor vpngw_tunnel peer-group
 neighbor vpngw_tunnel remote-as 66130
 neighbor vpngw_tunnel ebgp-multihop
 neighbor 198.19.97.4 peer-group spine-group
 neighbor 198.19.97.112 peer-group spine-group
 neighbor 198.19.98.2 remote-as 66011
 neighbor 198.19.98.2 next-hop-self
 neighbor 198.19.98.2 bfd
 address-family ipv4 unicast
  network 198.19.96.167/32
  maximum-paths 8
  graceful-restart
 !
 address-family ipv4 unicast vrf 5G_VRF
  local-as 65051
  redistribute connected
  redistribute static
  listen-range 40.172.2.0/26 peer-group upftb_s1u_n3 limit 64
  neighbor 198.19.98.243 remote-as 65051
  neighbor 198.19.98.243 next-hop-self
  maximum-paths 64
  default-information-originate
 !
 address-family ipv4 unicast vrf GLOBAL_VRF
  local-as 65051
  redistribute connected
  redistribute static
  listen-range 40.172.1.0/26 peer-group upftb_sx_n4 limit 64
  listen-range 40.172.1.64/26 peer-group upftb_s5u_n9 limit 64
  listen-range 40.172.1.192/26 peer-group upf_ims limit 64
  listen-range 40.172.141.128/30 peer-group vpngw_media limit 64
  listen-range 40.172.141.132/30 peer-group vpngw_signal limit 64
  listen-range 172.1.141.16/28 peer-group smf_global_link limit 64
  listen-range 172.1.141.192/28 peer-group smf_global_link limit 64
  neighbor 198.19.98.243 remote-as 65051
  neighbor 198.19.98.243 next-hop-self
  maximum-paths 64
  default-information-originate
 !
 address-family ipv4 unicast vrf NAT01_VRF
  local-as 66201
  redistribute connected
  redistribute static
  listen-range 40.172.1.128/26 peer-group upftb_nat01 limit 64
  neighbor 40.172.1.135 soft-reconfiguration inbound
  neighbor 40.172.1.136 soft-reconfiguration inbound
  neighbor 198.19.98.243 remote-as 66201
  neighbor 198.19.98.243 next-hop-self
  maximum-paths 64
  default-information-originate
 !
 address-family ipv4 unicast vrf SBI_VRF
  local-as 65051
  redistribute connected
  redistribute static
  listen-range 40.172.140.0/30 peer-group vpngw_tunnel limit 64
  listen-range 172.1.140.0/26 peer-group k8s_ext_svc limit 128
  neighbor 198.19.98.243 remote-as 65051
  neighbor 198.19.98.243 next-hop-self
  maximum-paths 64
  default-information-originate
 !
 address-family ipv6 unicast
 !
 address-family ipv6 unicast vrf 5G_VRF
  redistribute connected
  redistribute static
  neighbor fd00:60c:0:12::75 remote-as 65051
  neighbor fd00:60c:0:12::75 next-hop-self
  neighbor fd00:60c:0:12::75 activate
  maximum-paths 64
 !
 address-family ipv6 unicast vrf GLOBAL_VRF
  redistribute connected
  redistribute static
  listen-range fd00:0:0:47::400/120 peer-group upf_ims limit 64
  neighbor fd00:60c:0:12::75 remote-as 65051
  neighbor fd00:60c:0:12::75 next-hop-self
  neighbor fd00:60c:0:12::75 activate
  maximum-paths 64
  default-information-originate
 !
 address-family ipv6 unicast vrf NAT01_VRF
  redistribute connected
  redistribute static
  listen-range fd00:0:0:47::/120 peer-group upftb_nat01 limit 64
  neighbor fd00:0:0:47::7 soft-reconfiguration inbound
  neighbor fd00:0:0:47::8 soft-reconfiguration inbound
  neighbor fd00:60c:0:12::75 remote-as 66201
  neighbor fd00:60c:0:12::75 next-hop-self
  neighbor fd00:60c:0:12::75 activate
  maximum-paths 64
  default-information-originate
 !
 address-family ipv6 unicast vrf SBI_VRF
  redistribute connected
  redistribute static
  neighbor fd00:60c:0:12::75 remote-as 65051
  neighbor fd00:60c:0:12::75 next-hop-self
  neighbor fd00:60c:0:12::75 activate
  maximum-paths 64
 !
 address-family l2vpn evpn
  graceful-restart
  neighbor spine-group encapsulation vxlan
  neighbor spine-group next-hop-unchanged
  neighbor spine-group enable-peer-as-check
  neighbor spine-group activate
 !
!
ipv6 anycast-gateway-mac 0201.0101.0102
interface Loopback 1
 ip address 198.19.96.93/32
 no shutdown
!
interface Loopback 2
 ip address 198.19.96.167/32
 no shutdown
!
interface Ve 810
 vrf forwarding SBI_VRF
 ip anycast-address 172.1.140.1/26
 ip address 172.1.140.2/26
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ve 811
 vrf forwarding SBI_VRF
 ip anycast-address 172.1.140.177/28
 ip address 172.1.140.178/28
 no shutdown
!
interface Ve 820
 vrf forwarding 5G_VRF
 ip anycast-address 172.1.142.1/27
 ip address 172.1.142.2/27
 no shutdown
!
interface Ve 821
 vrf forwarding SBI_VRF
 ip anycast-address 172.1.140.65/28
 ip address 172.1.140.66/28
 no shutdown
!
interface Ve 822
 vrf forwarding GLOBAL_VRF
 ip anycast-address 172.1.141.1/28
 ip address 172.1.141.2/28
 no shutdown
!
interface Ve 825
 vrf forwarding SBI_VRF
 ip anycast-address 172.1.140.161/28
 ip address 172.1.140.162/28
 no shutdown
!
interface Ve 826
 vrf forwarding GLOBAL_VRF
 ip anycast-address 172.1.141.17/28
 ip address 172.1.141.18/28
 no shutdown
!
interface Ve 827
 vrf forwarding GLOBAL_VRF
 ip anycast-address 172.1.141.193/28
 ip address 172.1.141.194/28
 no shutdown
!
interface Ve 831
 vrf forwarding 5G_VRF
 ip anycast-address 40.172.2.1/26
 ip address 40.172.2.2/26
 bfd interval 1000 min-rx 1000 multiplier 3
 no shutdown
!
interface Ve 832
 description SMF_vpngw_GTP/Diameter
 vrf forwarding GLOBAL_VRF
 ip anycast-address 40.172.1.1/26
 ip address 40.172.1.2/26
 bfd interval 1000 min-rx 1000 multiplier 3
 no shutdown
!
interface Ve 833
 vrf forwarding GLOBAL_VRF
 ip anycast-address 40.172.1.65/26
 ip address 40.172.1.66/26
 bfd interval 1000 min-rx 1000 multiplier 3
 no shutdown
!
interface Ve 834
 vrf forwarding NAT01_VRF
 ip anycast-address 40.172.1.129/26
 ip address 40.172.1.130/26
 ipv6 anycast-address fd00:0:0:47::1/120
 ipv6 address fd00:0:0:47::2/120
 bfd interval 1000 min-rx 1000 multiplier 3
 no shutdown
!
interface Ve 835
 vrf forwarding GLOBAL_VRF
 ip anycast-address 40.172.1.193/26
 ip address 40.172.1.194/26
 ipv6 anycast-address fd00:0:0:47::401/120
 ipv6 address fd00:0:0:47::402/120
 bfd interval 1000 min-rx 1000 multiplier 3
 no shutdown
!
interface Ve 836
 vrf forwarding GLOBAL_VRF
 ip anycast-address 172.1.141.161/28
 ip address 172.1.141.162/28
 no shutdown
!
interface Ve 840
 vrf forwarding 5G_VRF
 ip anycast-address 172.1.142.121/29
 ip address 172.1.142.122/29
 no shutdown
!
interface Ve 841
 vrf forwarding SBI_VRF
 ip anycast-address 172.1.140.145/29
 ip address 172.1.140.146/29
 no shutdown
!
interface Ve 842
 vrf forwarding GLOBAL_VRF
 ip anycast-address 172.1.141.41/29
 ip address 172.1.141.42/29
 no shutdown
!
interface Ve 931
 vrf forwarding 5G_VRF
 ip address 172.1.142.226/30
 no shutdown
!
interface Ve 932
 vrf forwarding 5G_VRF
 ip address 172.1.142.230/30
 no shutdown
!
interface Ve 933
 vrf forwarding 5G_VRF
 no shutdown
!
interface Ve 934
 vrf forwarding 5G_VRF
 no shutdown
!
interface Ve 935
 vrf forwarding 5G_VRF
 ip address 40.255.255.2/30
 no shutdown
!
interface Ve 936
 vrf forwarding 5G_VRF
 no shutdown
!
interface Ve 937
 vrf forwarding 5G_VRF
 ip address 172.1.142.242/30
 no shutdown
!
interface Ve 938
 vrf forwarding 5G_VRF
 ip address 172.1.140.246/30
 no shutdown
!
interface Ve 939
 vrf forwarding 5G_VRF
 no shutdown
!
interface Ve 940
 vrf forwarding 5G_VRF
 no shutdown
!
interface Ve 941
 vrf forwarding SBI_VRF
 ip address 172.1.140.226/30
 no shutdown
!
interface Ve 942
 vrf forwarding SBI_VRF
 ip address 172.1.140.230/30
 no shutdown
!
interface Ve 943
 vrf forwarding SBI_VRF
 no shutdown
!
interface Ve 944
 vrf forwarding SBI_VRF
 no shutdown
!
interface Ve 945
 vrf forwarding SBI_VRF
 ip address 40.255.255.6/30
 no shutdown
!
interface Ve 946
 vrf forwarding SBI_VRF
 no shutdown
!
interface Ve 947
 vrf forwarding SBI_VRF
 ip address 172.1.140.242/30
 no shutdown
!
interface Ve 948
 vrf forwarding SBI_VRF
 no shutdown
!
interface Ve 949
 vrf forwarding SBI_VRF
 ip address 172.1.140.250/30
 no shutdown
!
interface Ve 950
 vrf forwarding SBI_VRF
 no shutdown
!
interface Ve 951
 vrf forwarding GLOBAL_VRF
 ip address 172.1.141.226/30
 ipv6 address fd00::47:c000:0:0:1/127
 no shutdown
!
interface Ve 952
 vrf forwarding GLOBAL_VRF
 ip address 172.1.141.230/30
 ipv6 address fd00::47:c000:0:0:3/127
 no shutdown
!
interface Ve 953
 vrf forwarding GLOBAL_VRF
 no shutdown
!
interface Ve 954
 vrf forwarding GLOBAL_VRF
 no shutdown
!
interface Ve 955
 vrf forwarding GLOBAL_VRF
 ip address 40.255.255.10/30
 ipv6 address fd00::47:0:0:f000:100/127
 no shutdown
!
interface Ve 956
 vrf forwarding GLOBAL_VRF
 no shutdown
!
interface Ve 965
 vrf forwarding NAT01_VRF
 ip address 40.255.255.14/30
 ipv6 address fd00::47:0:0:f000:0/127
 no shutdown
!
interface Ve 966
 vrf forwarding NAT01_VRF
 no shutdown
!
interface Ve 970
 vrf forwarding NAT01_VRF
 ip address 172.1.141.106/30
 ipv6 address fd00::47:c001:0:0:2/126
 no shutdown
!
interface Ve 972
 vrf forwarding NAT01_VRF
 ip address 172.1.141.114/30
 ipv6 address fd00::47:c001:0:0:a/126
 no shutdown
!
interface Ve 974
 vrf forwarding NAT01_VRF
 no shutdown
!
interface Ve 976
 vrf forwarding NAT01_VRF
 no shutdown
!
interface Ve 978
 vrf forwarding GLOBAL_VRF
 ip access-group UE_to_SYSTEM_Deny in
 ip access-group SYSTEM_to_UE_Deny out
 ip address 172.1.141.138/30
 ipv6 address fd00::47:c001:0:0:22/126
 no shutdown
!
interface Ve 980
 vrf forwarding GLOBAL_VRF
 ip access-group UE_to_SYSTEM_Deny in
 ip access-group SYSTEM_to_UE_Deny out
 ip address 172.1.141.146/30
 ipv6 address fd00::47:c001:0:0:2a/126
 no shutdown
!
interface Ve 981
 vrf forwarding GLOBAL_VRF
 no shutdown
!
interface Ve 983
 vrf forwarding GLOBAL_VRF
 no shutdown
!
interface Ve 990
 vrf forwarding 5G_VRF
 ip address 40.255.255.102/30
 no shutdown
!
interface Ve 991
 vrf forwarding SBI_VRF
 ip address 40.255.255.106/30
 no shutdown
!
interface Ve 992
 vrf forwarding GLOBAL_VRF
 ip address 40.255.255.110/30
 ipv6 address fd00::47:0:0:f000:110/127
 no shutdown
!
interface Ve 993
 vrf forwarding NAT01_VRF
 ip address 40.255.255.114/30
 ipv6 address fd00::47:0:0:f000:10/127
 no shutdown
!
interface Ve 995
 vrf forwarding 5G_VRF
 no shutdown
!
interface Ve 996
 vrf forwarding SBI_VRF
 no shutdown
!
interface Ve 997
 vrf forwarding GLOBAL_VRF
 no shutdown
!
interface Ve 998
 vrf forwarding NAT01_VRF
 no shutdown
!
interface Ve 1003
 vrf forwarding SBI_VRF
 ip anycast-address 172.1.140.113/28
 ip address 172.1.140.114/28
 no shutdown
!
interface Ve 1004
 vrf forwarding SBI_VRF
 ip anycast-address 172.1.140.129/28
 ip address 172.1.140.130/28
 no shutdown
!
interface Ve 1005
 vrf forwarding GLOBAL_VRF
 ip anycast-address 172.1.141.209/28
 ip address 172.1.141.210/28
 no shutdown
!
interface Ve 1007
 vrf forwarding NAT01_VRF
 ip anycast-address 172.1.141.185/29
 ipv6 anycast-address fd00:172:1:141:185::f001/120
 ipv6 address fd00:172:1:141:185::f002/120
 no shutdown
!
interface Ve 1008
 vrf forwarding 5G_VRF
 ip anycast-address 172.1.142.113/29
 no shutdown
!
interface Ve 8188
 vrf forwarding NAT01_VRF
 ip address 198.19.98.242/31
 ipv6 address fd00:60c:0:12::74/127
 no shutdown
!
interface Ve 8189
 vrf forwarding GLOBAL_VRF
 ip address 198.19.98.242/31
 ipv6 address fd00:60c:0:12::74/127
 no shutdown
!
interface Ve 8190
 vrf forwarding 5G_VRF
 ip address 198.19.98.242/31
 ipv6 address fd00:60c:0:12::74/127
 no shutdown
!
interface Ve 8191
 vrf forwarding SBI_VRF
 ip address 198.19.98.242/31
 ipv6 address fd00:60c:0:12::74/127
 no shutdown
!
cluster TB2-cluster-1
 peer 198.19.98.2
 peer-interface Port-channel 64
 bringup-delay 200
 peer-keepalive
  auto
 !
 member vlan all
 member vlan remove 931-956,965-966,970,972,974,976,978,980-981,983,990-993,995-998
 member bridge-domain all
!
interface Management 0
 no tcp burstrate
 no shutdown
 vrf forwarding mgmt-vrf
 ip icmp unreachable
 ip icmp echo-reply
 no ip address dhcp
 ip address 60.30.181.111/27
 ip access-group MGMT-SNMP-PERMIT in
 no ipv6 address autoconfig
 no ipv6 address dhcp
!
interface Ethernet 0/1
 description Link to 60.30.181.101 Spine
 ip address 198.19.97.5/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/2
 description Link to 60.30.181.102 Spine
 ip address 198.19.97.113/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
interface Ethernet 0/3
 shutdown
!
interface Ethernet 0/4
 shutdown
!
interface Ethernet 0/5
 shutdown
!
interface Ethernet 0/6
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/7
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/8
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/9
 link-error-disable 1 1
 description Port-channel DBL1_NAT01 Member interface
 channel-group 10 mode active type standard
 lacp timeout short
 no shutdown
!
interface Ethernet 0/10
 link-error-disable 1 1
 description Port-channel DBL1_NAT01 Member interface
 channel-group 10 mode active type standard
 lacp timeout short
 no shutdown
!
interface Ethernet 0/11
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/12
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/13
 description Port-channel DBL1_NAT02 Member interface
 channel-group 12 mode active type standard
 lacp timeout short
 no shutdown
!
interface Ethernet 0/14
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/15
 shutdown
!
interface Ethernet 0/16
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/17
 description PUB-DCGW-01-100GE4/0/6
 cluster-track
 switchport
 switchport mode trunk-no-default-native
 switchport trunk allowed vlan add 931,937,941,947,951
 ip access-group ANTI-SPOOFING in
 no shutdown
!
interface Ethernet 0/18
 description PUB-DCGW-02-100GE4/0/6
 cluster-track
 switchport
 switchport mode trunk-no-default-native
 switchport trunk allowed vlan add 932,938-939,942-943,949,952
 ip access-group ANTI-SPOOFING in
 no shutdown
!
interface Ethernet 0/19
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/20
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/21
 description Sim_X870_A
 cluster-track
 switchport
 switchport mode trunk-no-default-native
 switchport trunk allowed vlan add 935,945,955,965
 no shutdown
!
interface Ethernet 0/22
 fec mode RS-FEC
 description Sim_X870_A
 cluster-track
 switchport
 switchport mode trunk-no-default-native
 switchport trunk allowed vlan add 990-993
 no shutdown
!
interface Ethernet 0/23:1
 shutdown
!
interface Ethernet 0/23:2
 shutdown
!
interface Ethernet 0/23:3
 shutdown
!
interface Ethernet 0/23:4
 shutdown
!
interface Ethernet 0/25
 description Port-channel DBL1_NATBK Member interface
 channel-group 11 mode active type standard
 lacp timeout short
 no shutdown
!
interface Ethernet 0/26
 link-error-disable 1 1
 description Port-channel DBL1_NATBK Member interface
 channel-group 11 mode active type standard
 lacp timeout short
 no shutdown
!
interface Ethernet 0/27:1
 cluster-track
 switchport
 switchport mode trunk-no-default-native
 switchport trunk allowed vlan add 822,1004,1007-1008
 no shutdown
!
interface Ethernet 0/27:2
 lldp disable
 no shutdown
!
interface Ethernet 0/27:3
 cluster-track
 switchport
 switchport mode trunk
 switchport trunk allowed vlan add 801
 no switchport trunk tag native-vlan
 switchport trunk native-vlan 801
 no shutdown
!
interface Ethernet 0/27:4
 shutdown
!
interface Ethernet 0/29
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/30
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/31
 ip access-group ANTI-SPOOFING in
 shutdown
!
interface Ethernet 0/32
 shutdown
!
interface Ethernet 0/33
 ip access-group ANTI-SPOOFING in
 no shutdown
!
interface Ethernet 0/34
 ip access-group ANTI-SPOOFING in
 no shutdown
!
interface Ethernet 0/35
 ip access-group ANTI-SPOOFING in
 no shutdown
!
interface Ethernet 0/36
 description clusterPeerIntfMember
 channel-group 64 mode active type standard
 lacp timeout short
 no shutdown
!
interface Ethernet 0/37
 description clusterPeerIntfMember
 channel-group 64 mode active type standard
 lacp timeout short
 no shutdown
!
interface Ethernet 0/38
 description clusterPeerIntfMember
 channel-group 64 mode active type standard
 lacp timeout short
 no shutdown
!
interface Ethernet 0/39:1
 redundant-management enable
 no shutdown
!
interface Ethernet 0/39:2
 shutdown
!
interface Ethernet 0/39:3
 shutdown
!
interface Ethernet 0/39:4
 shutdown
!
interface Port-channel 10
 speed 100000
 description DBL1_NAT01
 cluster-track
 switchport
 switchport mode trunk-no-default-native
 switchport trunk allowed vlan add 970,978
 no shutdown
!
interface Port-channel 11
 speed 100000
 description DBL1_NATBK
 cluster-track
 switchport
 switchport mode trunk-no-default-native
 switchport trunk allowed vlan add 972,980
 no shutdown
!
interface Port-channel 12
 speed 100000
 description DBL1_NAT02
 no shutdown
!
interface Port-channel 64
 speed 100000
 description MCTPeerInterface
 ip address 198.19.98.3/31
 bfd interval 200 min-rx 200 multiplier 3
 no shutdown
!
mac-address-table mac-move detect
monitor session 1
 source ethernet 0/1 destination ethernet 0/27:2 direction rx
!
monitor session 2
 source ethernet 0/2 destination ethernet 0/27:2 direction rx
!
monitor session 3
 source ethernet 0/17 destination ethernet 0/27:2 direction rx
!
monitor session 4
 source ethernet 0/18 destination ethernet 0/27:2 direction rx
!
monitor session 5
 source ethernet 0/21 destination ethernet 0/27:2 direction rx
!
monitor session 6
 source ethernet 0/27:1 destination ethernet 0/27:2 direction rx
!
bridge-domain 4092 p2mp
 description Tenant L3 Extended BR BD
 pw-profile Tenant-profile
 router-interface Ve 8188
 !
 bpdu-drop-enable
 local-switching
!
bridge-domain 4093 p2mp
 description Tenant L3 Extended BR BD
 pw-profile Tenant-profile
 router-interface Ve 8189
 !
 bpdu-drop-enable
 local-switching
!
bridge-domain 4094 p2mp
 description Tenant L3 Extended BR BD
 pw-profile Tenant-profile
 router-interface Ve 8190
 !
 bpdu-drop-enable
 local-switching
!
bridge-domain 4095 p2mp
 description Tenant L3 Extended BR BD
 pw-profile Tenant-profile
 router-interface Ve 8191
 !
 bpdu-drop-enable
 local-switching
!
logging raslog message HIL-1530 suppress
logging raslog message RAS-2006 suppress
logging raslog message RAS-2007 suppress
logging raslog message SEC-1203 suppress
logging raslog message SEC-1206 suppress
logging raslog message SEC-3022 suppress
logging raslog console INFO
logging syslog-server 60.30.181.98 use-vrf mgmt-vrf
 secure
!
logging auditlog class CONFIGURATION
logging auditlog class FIRMWARE
logging syslog-facility local LOG_LOCAL6
overlay-gateway TB2
 ip interface Loopback 2
 map vni auto
 activate
!
telemetry profile system-utilization default_system_utilization_statistics
 interval 60
 add total-system-memory
 add total-used-memory
 add total-free-memory
 add cached-memory
 add buffers
 add total-swap-memory
 add total-free-swap-memory
 add total-used-swap-memory
 add user-process
 add system-process
 add niced-process
 add iowait
 add hw-interrupt
 add sw-interrupt
 add idle-state
 add steal-time
 add uptime
!
telemetry profile interface default_interface_statistics
 interval 30
 add out-pkts
 add in-pkts
 add out-unicast-pkts
 add in-unicast-pkts
 add out-broadcast-pkts
 add in-broadcast-pkts
 add out-multicast-pkts
 add in-multicast-pkts
 add out-pkts-per-second
 add in-pkts-per-second
 add out-bandwidth
 add in-bandwidth
 add out-octets
 add in-octets
 add out-errors
 add in-errors
 add out-crc-errors
 add in-crc-errors
 add out-discards
 add in-discards
!
telemetry profile mpls-traffic-lsp default_mpls_traffic_lsp_statistics
 interval 240
 add out-packets
 add out-bytes
!
telemetry profile mpls-traffic-bypass default_mpls_traffic_bypass_statistics
 interval 240
 add out-packets
 add out-bytes
!
telemetry profile mpls-traffic-fec default_mpls_traffic_fec_statistics
 interval 240
 add out-packets
 add out-bytes
!

--- Time Elapsed: 370.336093ms ---
`
	} else {
		err = errors.New(fmt.Sprintf(" UUID %s Not Found", uuid))
	}
	return result, err
}
