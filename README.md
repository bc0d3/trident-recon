# 🔱 Trident Recon

[![Release](https://img.shields.io/github/v/release/bc0d3/trident-recon)](https://github.com/bc0d3/trident-recon/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/bc0d3/trident-recon)](https://goreportcard.com/report/github.com/bc0d3/trident-recon)
[![License](https://img.shields.io/github/license/bc0d3/trident-recon)](LICENSE)
[![CI](https://github.com/bc0d3/trident-recon/workflows/CI/badge.svg)](https://github.com/bc0d3/trident-recon/actions)

A powerful bug bounty tool orchestrator that automates reconnaissance workflows by executing multiple tools in background tmux sessions.

## Features

- 🎯 **Multi-tool support**: ffuf, gobuster, dirsearch, nuclei, httpx, and more
- 🔧 **Highly configurable**: YAML-based configuration for easy customization
- 📋 **Command generation**: Generate copy-pasteable commands in markdown format
- ⚡ **Batch processing**: Process multiple targets from a file
- 🎮 **Session management**: Easy tmux session management and monitoring
- 📊 **Organized output**: Clean directory structure with detailed logs

## Installation

### Prerequisites
- **Linux or macOS** (tmux is not available on Windows)
- Go 1.21 or higher
- tmux installed
- Your favorite recon tools (ffuf, gobuster, etc.)

### Via go install (Recommended)
```bash
# Install latest version
go install github.com/bc0d3/trident-recon@latest

# Install specific version
go install github.com/bc0d3/trident-recon@v1.0.0
```

### Download Pre-built Binary
Download from [releases page](https://github.com/bc0d3/trident-recon/releases/latest)

**Linux (amd64)**
```bash
wget https://github.com/bc0d3/trident-recon/releases/latest/download/trident-recon_linux_x86_64.tar.gz
tar -xzf trident-recon_linux_x86_64.tar.gz
sudo mv trident-recon /usr/local/bin/
```

**macOS (Intel)**
```bash
wget https://github.com/bc0d3/trident-recon/releases/latest/download/trident-recon_darwin_x86_64.tar.gz
tar -xzf trident-recon_darwin_x86_64.tar.gz
sudo mv trident-recon /usr/local/bin/
```

**macOS (Apple Silicon M1/M2)**
```bash
wget https://github.com/bc0d3/trident-recon/releases/latest/download/trident-recon_darwin_arm64.tar.gz
tar -xzf trident-recon_darwin_arm64.tar.gz
sudo mv trident-recon /usr/local/bin/
```

**Linux (ARM64 - Raspberry Pi, etc.)**
```bash
wget https://github.com/bc0d3/trident-recon/releases/latest/download/trident-recon_linux_arm64.tar.gz
tar -xzf trident-recon_linux_arm64.tar.gz
sudo mv trident-recon /usr/local/bin/
```

### Build from Source
```bash
git clone https://github.com/bc0d3/trident-recon.git
cd trident-recon
make build
sudo make install
```

## Quick Start

1. **Initialize configuration**
```bash
trident-recon init
```

2. **Edit config** (optional)
```bash
vim ~/.config/trident-recon/config.yaml
```

3. **Generate commands**
```bash
trident-recon generate -u http://testphp.vulnweb.com
```

4. **Run reconnaissance**
```bash
trident-recon run -u http://testphp.vulnweb.com
```

## Usage

### Basic Commands
```bash
# Generate commands only (creates commands.md)
trident-recon generate -u http://example.com

# Execute commands in tmux
trident-recon run -u http://example.com

# Specify output directory
trident-recon run -u http://example.com -o ~/scans/target1

# Process multiple targets
trident-recon run -l targets.txt
```

### Tool Filtering
```bash
# Run only specific tools
trident-recon run -u http://example.com --tools ffuf,gobuster

# Skip specific tools
trident-recon run -u http://example.com --skip nuclei,arjun
```

### Session Management
```bash
# List all active sessions
trident-recon list

# List sessions for specific tool
trident-recon list --tool ffuf

# Kill specific session
trident-recon kill <session-id>

# Kill all sessions
trident-recon kill-all

# Kill all sessions for specific tool
trident-recon kill-all --tool ffuf
```

## Configuration

Config location: `~/.config/trident-recon/config.yaml`

See [examples/config.yaml](examples/config.yaml) for a complete configuration example.

### Adding Custom Tools
```yaml
tools:
  your-tool:
    enabled: true
    tmux_prefix: "yourtool_"
    commands:
      - name: "scan"
        description: "Your tool scan"
        command: "yourtool -u {URL} -o {OUTPUT_DIR}/output.txt"
        wordlist: ""
```

### Template Variables

Available variables for command templates:
- `{URL}` - Full target URL
- `{DOMAIN}` - Domain without protocol
- `{PROTOCOL}` - http or https
- `{WORDLIST}` - Path to wordlist
- `{OUTPUT_DIR}` - Output directory
- `{ID}` - Unique session ID

## Output Structure

```
~/trident-output/
└── testphp.vulnweb.com_20251021_153045/
    ├── ffuf-testphp.vulnweb.com-content.json
    ├── ffuf-testphp.vulnweb.com-php.json
    ├── gobuster-testphp.vulnweb.com-dirs.txt
    └── commands.md
```

## Examples

### Single Target Recon
```bash
trident-recon run -u https://target.com -o ~/bug-bounty/target1
```

### Batch Processing
```bash
# Create targets file
cat > targets.txt <<EOF
https://target1.com
https://target2.com
https://target3.com
EOF

# Run batch scan
trident-recon run -l targets.txt
```

### Custom Workflow
```bash
# Only run ffuf commands
trident-recon run -u https://target.com --tools ffuf

# Skip vulnerability scanning
trident-recon run -u https://target.com --skip nuclei,sqlmap
```

## Development

### Build
```bash
make build
```

### Install locally
```bash
make install
```

### Clean
```bash
make clean
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see [LICENSE](LICENSE) file for details

## Author

[@bc0d3](https://github.com/bc0d3)

## Acknowledgments

Built for the bug bounty community. Special thanks to all the amazing tool developers:
- ffuf, gobuster, dirsearch, feroxbuster
- httpx, nuclei, subfinder
- And many more!

---

**⚠️ Disclaimer**: This tool is for authorized security testing only. Always obtain proper permission before testing.
