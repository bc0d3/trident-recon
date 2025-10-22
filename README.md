# 🔱 Trident Recon

[![Release](https://img.shields.io/github/v/release/bc0d3/trident-recon)](https://github.com/bc0d3/trident-recon/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/bc0d3/trident-recon)](https://goreportcard.com/report/github.com/bc0d3/trident-recon)
[![License](https://img.shields.io/github/license/bc0d3/trident-recon)](LICENSE)
[![CI](https://github.com/bc0d3/trident-recon/workflows/CI/badge.svg)](https://github.com/bc0d3/trident-recon/actions)

A powerful bug bounty tool orchestrator that automates reconnaissance workflows by executing multiple tools in background tmux sessions.

## Features

- 🎯 **Multi-tool support**: ffuf, gobuster, dirsearch, feroxbuster, gowitness, and more
- 🔧 **Highly configurable**: YAML-based configuration for easy customization
- 📋 **Dual output formats**: Generates both detailed markdown AND plain text commands for easy copy-paste
- ⚡ **Batch processing**: Process multiple targets from a file with automatic domain list generation
- 🎮 **Session management**: Easy tmux session management and monitoring
- 📊 **Organized output**: Clean directory structure with detailed logs
- 🖼️ **Screenshot support**: Integrated gowitness for visual reconnaissance
- 📝 **Copy-paste ready**: Plain text commands.txt file for instant execution

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
# Generate commands only (creates commands.md AND commands.txt)
trident-recon generate -u http://example.com

# Execute commands in tmux sessions
trident-recon run -u http://example.com

# Specify output directory
trident-recon run -u http://example.com -o ~/scans/target1

# Process multiple targets (auto-creates domains.txt for tools like gowitness)
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
- `{DOMAIN_LIST}` - Path to auto-generated domains list file (for multi-target scans)

### Tools that Support Domain Lists

Some tools can process multiple domains from a file. Use `use_domain_list: true` in your config:

```yaml
tools:
  gowitness:
    enabled: true
    tmux_prefix: "gowitness_"
    commands:
      - name: "multi-urls"
        description: "Screenshot multiple URLs"
        command: "gowitness file -f {DOMAIN_LIST} --threads 10 --write-db"
        wordlist: ""
        use_domain_list: true  # This command will use the domains.txt file
```

## Output Structure

### Single Target Scan
```
~/trident-output/
└── example.com_20251022_153045/
    ├── commands.md                              # Detailed markdown with session info
    ├── commands.txt                             # Plain text commands (copy-paste ready)
    ├── ffuf-example.com-content.json
    ├── ffuf-example.com-quickhits.json
    ├── gobuster-example.com-dirs.txt
    ├── feroxbuster-example.com-fast.txt
    └── gowitness-screenshots/
        ├── screenshot-example.com.png
        └── gowitness.sqlite3
```

### Multiple Targets Scan
```
~/trident-output/
└── multi-target_20251022_153045/
    ├── commands.md                              # All commands for all targets
    ├── commands.txt                             # Copy-paste ready commands
    ├── domains.txt                              # Auto-generated domain list
    ├── ffuf-target1.com-content.json
    ├── ffuf-target2.com-content.json
    ├── gobuster-target1.com-dirs.txt
    └── gowitness-screenshots/                   # Screenshots from all targets
        ├── screenshot-target1.com.png
        ├── screenshot-target2.com.png
        └── gowitness.sqlite3
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

# Only take screenshots with gowitness
trident-recon run -l urls.txt --tools gowitness
```

### Using Generated Commands

After running `generate` or `run`, you get two files:

**1. commands.md** - Detailed documentation with:
- Session IDs for each command
- Full command with context
- Instructions for attaching to sessions
- Output file locations

**2. commands.txt** - One command per line, ready to copy-paste:
```bash
# Just copy and paste the entire file!
tmux new-session -d -s "ffuf_abc123" bash -c "ffuf -u http://example.com/FUZZ ..."
tmux new-session -d -s "gobuster_def456" bash -c "gobuster dir -u http://example.com ..."
tmux new-session -d -s "gowitness_ghi789" bash -c "gowitness file -f domains.txt ..."
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

## Supported Tools

Trident Recon comes pre-configured with optimized commands for:

### Fuzzing & Directory Discovery
- **ffuf** - Fast web fuzzer with auto-calibration
- **gobuster** - Directory/file & DNS busting tool
- **dirsearch** - Web path scanner with recursive scanning
- **feroxbuster** - Fast content discovery with auto-tune

### Visual Reconnaissance
- **gowitness** - Web screenshot utility with full-page capture

All tools include intelligent rate limiting, multi-threading, and optimized wordlists for maximum efficiency.

## Acknowledgments

Built for the bug bounty community. Special thanks to all the amazing tool developers:
- [ffuf](https://github.com/ffuf/ffuf) by @joohoi
- [gobuster](https://github.com/OJ/gobuster) by @OJ
- [dirsearch](https://github.com/maurosoria/dirsearch) by @maurosoria
- [feroxbuster](https://github.com/epi052/feroxbuster) by @epi052
- [gowitness](https://github.com/sensepost/gowitness) by @sensepost

---

**⚠️ Disclaimer**: This tool is for authorized security testing only. Always obtain proper permission before testing.
