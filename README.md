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


git clone <repository-url>
cd network-tool

# Build for Windows
go build -o network-tool.exe main.go

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o network-tool-linux main.go
