# Web Attack Detector

## 実装済み機能

- なし

## 実装する機能

- 設定したポートを眺めて異常なアクセスを検知
  - そもそも外部からアクセスされたくないポートを設定できる
- 攻撃を検知したらメールとかで通知
- できれば自己学習させたいけどな〜
- そもそものポート開けすぎFirewallの設定カスすぎとかは最初に警告出して欲しいが…
  - これはOS依存なのでいい感じにやらないといけない

### 検知したい攻撃

- インフラ回り
  - 踏み台攻撃
  - DoS

- アプリケーション
  - WebShell
  - SQLi
  - XSS
  - ディレクトリトラバーサル
  - OSコマンドインジェクション

- その他
  - 不正なクライアント(チートなど)

## 実装方針

- ~~Webサーバのアクセスログから検出~~
  - これだとWebサーバ以外(ssh, DBとか)への攻撃が取れない
- リアルタイムにパケットキャプチャ

## 動機

- Web開発をする上でセキュリティの重要性を感じた
- 同様なツールにIPAの公開しているiLogScannerがある
  - リアルタイム性に欠ける
  - サーバーログのみを見るのでDBなど他のプロセスへの攻撃を判定できない

## 進捗

**※ソースとバイナリ更新したら`vagrant reload`しないと反映されないぞ**

eth1のパケットを読みながらホストでping送り続けた↓
``` shell
PACKET: 98 bytes, wire length 98 cap length 98 @ 2019-12-16 07:37:27.263528 +0000 UTC
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..84..] SrcMAC=0a:00:27:00:00:01 DstMAC=08:00:27:4d:67:ff EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4	{Contents=[..20..] Payload=[..64..] Version=4 IHL=5 TOS=0 Length=84 Id=43410 Flags=DF FragOffset=0 TTL=64 Protocol=ICMPv4 Checksum=52666 SrcIP=192.168.33.1 DstIP=192.168.33.10 Options=[] Padding=[]}
- Layer 3 (08 bytes) = ICMPv4	{Contents=[..8..] Payload=[..56..] TypeCode=EchoRequest Checksum=38904 Id=5878 Seq=222}
- Layer 4 (56 bytes) = Payload	56 byte(s)

PACKET: 98 bytes, wire length 98 cap length 98 @ 2019-12-16 07:37:27.263578 +0000 UTC
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..84..] SrcMAC=08:00:27:4d:67:ff DstMAC=0a:00:27:00:00:01 EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4	{Contents=[..20..] Payload=[..64..] Version=4 IHL=5 TOS=0 Length=84 Id=49483 Flags= FragOffset=0 TTL=64 Protocol=ICMPv4 Checksum=62977 SrcIP=192.168.33.10 DstIP=192.168.33.1 Options=[] Padding=[]}
- Layer 3 (08 bytes) = ICMPv4	{Contents=[..8..] Payload=[..56..] TypeCode=EchoReply Checksum=40952 Id=5878 Seq=222}
- Layer 4 (56 bytes) = Payload	56 byte(s)

PACKET: 42 bytes, wire length 42 cap length 42 @ 2019-12-16 07:37:30.25102 +0000 UTC
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..28..] SrcMAC=08:00:27:4d:67:ff DstMAC=0a:00:27:00:00:01 EthernetType=ARP Length=0}
- Layer 2 (28 bytes) = ARP	{Contents=[..28..] Payload=[] AddrType=Ethernet Protocol=IPv4 HwAddressSize=6 ProtAddressSize=4 Operation=1 SourceHwAddress=[..6..] SourceProtAddress=[192, 168, 33, 10] DstHwAddress=[..6..] DstProtAddress=[192, 168, 33, 1]}

PACKET: 60 bytes, wire length 60 cap length 60 @ 2019-12-16 07:37:30.251231 +0000 UTC
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..46..] SrcMAC=0a:00:27:00:00:01 DstMAC=08:00:27:4d:67:ff EthernetType=ARP Length=0}
- Layer 2 (28 bytes) = ARP	{Contents=[..28..] Payload=[..18..] AddrType=Ethernet Protocol=IPv4 HwAddressSize=6 ProtAddressSize=4 Operation=2 SourceHwAddress=[..6..] SourceProtAddress=[192, 168, 33, 1] DstHwAddress=[..6..] DstProtAddress=[192, 168, 33, 10]}
- Layer 3 (18 bytes) = Payload	18 byte(s)

PACKET: 66 bytes, wire length 66 cap length 66 @ 2019-12-16 07:37:34.063823 +0000 UTC
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..52..] SrcMAC=0a:00:27:00:00:01 DstMAC=08:00:27:4d:67:ff EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4	{Contents=[..20..] Payload=[..32..] Version=4 IHL=5 TOS=0 Length=52 Id=38067 Flags=DF FragOffset=0 TTL=64 Protocol=TCP Checksum=58036 SrcIP=192.168.33.1 DstIP=192.168.33.10 Options=[] Padding=[]}
- Layer 3 (32 bytes) = TCP	{Contents=[..32..] Payload=[] SrcPort=46128 DstPort=80(http) Seq=3254196337 Ack=3163937474 DataOffset=8 FIN=false SYN=false RST=false PSH=false ACK=true URG=false ECE=false CWR=false NS=false Window=502 Checksum=20372 Urgent=0 Options=[TCPOption(NOP:), TCPOption(NOP:), TCPOption(Timestamps:3555010628/212840 0xd3e5284400033f68)] Padding=[]}

PACKET: 54 bytes, wire length 54 cap length 54 @ 2019-12-16 07:37:34.063871 +0000 UTC
- Layer 1 (14 bytes) = Ethernet	{Contents=[..14..] Payload=[..40..] SrcMAC=08:00:27:4d:67:ff DstMAC=0a:00:27:00:00:01 EthernetType=IPv4 Length=0}
- Layer 2 (20 bytes) = IPv4	{Contents=[..20..] Payload=[..20..] Version=4 IHL=5 TOS=0 Length=40 Id=8324 Flags=DF FragOffset=0 TTL=64 Protocol=TCP Checksum=22256 SrcIP=192.168.33.10 DstIP=192.168.33.1 Options=[] Padding=[]}
- Layer 3 (20 bytes) = TCP	{Contents=[..20..] Payload=[] SrcPort=80(http) DstPort=46128 Seq=3163937474 Ack=0 DataOffset=5 FIN=false SYN=false RST=true PSH=false ACK=false URG=false ECE=false CWR=false NS=false Window=0 Checksum=41131 Urgent=0 Options=[] Padding=[]}
```

ICMPv4って出てるのがpingで送られてきたパケット，なんでかしらんけどARPも付いてくるしpingやめたあとも定期的に謎のTCPパケットが送られてる　これはもしかして別タブで開いてたやつだろうか
