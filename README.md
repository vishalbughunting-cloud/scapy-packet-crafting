# scapy-packet-crafting
🚀 Network Attack Tool - Scapy Alternative A powerful Golang-based network testing tool that replicates Scapy functionality for penetration testing and network security assessment.
This tool provides Scapy-like network packet manipulation capabilities in a single executable file. It's designed for ethical hacking, penetration testing, and network security research.

⚠️ Legal Disclaimer: Use only on systems you own or have explicit permission to test. Unauthorized use may violate laws.
🛠️ Features
✅ RST Flood Attack - TCP connection reset attacks

✅ SYN Flood Attack - DDoS simulation

✅ Port Scanning - Network reconnaissance

✅ General Flood Attack - Resource exhaustion testing

✅ Cross-Platform - Windows, Linux, macOS support

✅ Single Executable - No dependencies required

Go 1.19 or higher

Windows/Linux/macOS

# Build for Windows
git clone https://github.com/vishalbughunting-cloud/scapy-packet-crafting
cd network-tool
go build -o network-tool.exe main.go

<img width="1445" height="688" alt="image" src="https://github.com/user-attachments/assets/5f9c4fbd-8d73-4a1b-8f07-377d518ab4c0" />


<img width="1457" height="564" alt="image" src="https://github.com/user-attachments/assets/06428258-f841-4732-92ac-f67105c5e50d" />

<img width="1913" height="853" alt="image" src="https://github.com/user-attachments/assets/eacaa5b6-cbde-4716-9704-4de6ddcd38ce" />

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o network-tool-linux main.go
