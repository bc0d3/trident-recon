package config

// DefaultConfig returns a default configuration template
const DefaultConfig = `# Trident Recon Configuration File - OPTIMIZED FOR BUG BOUNTY 2025
# Save this at ~/.config/trident-recon/config.yaml
#
# ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
# TEMPLATE VARIABLES - Available in all command templates:
# ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
#
# {URL}          - Full target URL with protocol
#                  Example: http://example.com or https://api.example.com
#
# {DOMAIN}       - Extracted domain from URL (without port)
#                  Example: example.com (even if URL is http://example.com:8080)
#
# {PROTOCOL}     - Protocol extracted from URL
#                  Example: http or https
#
# {WORDLIST}     - Full path to the wordlist file specified in command
#                  Example: /usr/share/seclists/Discovery/Web-Content/common.txt
#                  This is resolved from the wordlist name defined below
#
# {OUTPUT_DIR}   - Output directory for this scan
#                  Single target: ~/trident-output/example.com/
#                  With -o flag: /custom/path/example.com/
#                  Multiple targets: ~/trident-output/example.com/, ~/trident-output/google.com/, etc.
#
# {ID}           - Unique session identifier (12 characters by default)
#                  Example: a1b2c3d4e5f6
#                  Generated using MD5 hash of tool+command+domain+timestamp
#
# {DOMAIN_LIST}  - Path to domains.txt file (only for multi-target scans)
#                  Example: ~/trident-output/domains.txt
#                  Contains list of all domains being scanned (one per line)
#                  Only available when using -l flag with multiple targets
#
# ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
# HEADER VARIABLES - Dynamic based on your config.yaml:
# ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
#
# Individual Headers (any header from your config):
# {HEADER-User-Agent}      - Replaced with: -H "User-Agent: Mozilla/5.0..."
# {HEADER-Accept}          - Replaced with: -H "Accept: */*"
# {HEADER-Accept-Language} - Replaced with: -H "Accept-Language: en-US,en;q=0.9"
# {HEADER-Accept-Encoding} - Replaced with: -H "Accept-Encoding: gzip, deflate"
# {HEADER-X-Bug-Bounty}    - Replaced with: -H "X-Bug-Bounty: hackeroneUser"
# {HEADER-<any-name>}      - Any header you add to config.yaml automatically works!
#
# Grouped Headers:
# {HEADERS-ALL}     - ALL headers (default + custom) combined
#                     Example: -H "User-Agent: ..." -H "Accept: ..." -H "X-Bug-Bounty: ..."
#
# {HEADERS-DEFAULT} - Only default headers from headers.default section
#                     Example: -H "User-Agent: ..." -H "Accept: ..." -H "Accept-Language: ..."
#
# {HEADERS-CUSTOM}  - Only custom headers from headers.custom section
#                     Example: -H "X-Bug-Bounty: hackeroneUser"
#
# Usage Examples:
#   ffuf -u {URL}/FUZZ {HEADERS-ALL} -w wordlist.txt
#   gobuster dir -u {URL} {HEADER-User-Agent} {HEADER-X-Bug-Bounty} -w wordlist.txt
#   curl {URL} {HEADERS-DEFAULT} -o output.txt
#   nuclei -u {URL} {HEADER-Accept} -H "Custom: value"
#
# ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
# USAGE EXAMPLES:
# ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
#
# Single target:
#   trident-recon generate -u http://example.com
#   → Output: ~/trident-output/example.com/comandos.md
#   → Output: ~/trident-output/example.com/comandos.txt
#
# Single target with custom output:
#   trident-recon generate -u http://example.com -o /tmp/my-scan
#   → Output: /tmp/my-scan/example.com/comandos.md
#   → Output: /tmp/my-scan/example.com/comandos.txt
#
# Multiple targets:
#   trident-recon generate -l targets.txt
#   → Output: ~/trident-output/example.com/comandos.md
#   → Output: ~/trident-output/google.com/comandos.md
#   → Output: ~/trident-output/domains.txt (shared list)
#
# Multiple targets with custom output:
#   trident-recon generate -l targets.txt -o /tmp/my-scan
#   → Output: /tmp/my-scan/example.com/comandos.md
#   → Output: /tmp/my-scan/google.com/comandos.md
#   → Output: /tmp/my-scan/domains.txt
#
# Filter specific tools:
#   trident-recon generate -u http://example.com --tools ffuf,gobuster
#
# Skip specific tools:
#   trident-recon generate -u http://example.com --skip dirsearch,feroxbuster
#
# Run commands (execute in tmux):
#   trident-recon run -u http://example.com
#   trident-recon run -l targets.txt -o ~/scans/project-name
#
# ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

global:
  output_dir: ~/trident-output
  id_length: 12

headers:
  default:
    User-Agent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
    Accept: "*/*"
    Accept-Language: "en-US,en;q=0.9"
    Accept-Encoding: "gzip, deflate"
  custom:
    - "X-Bug-Bounty: hackeroneUser"
    # Add any custom headers here. Each will be available as {HEADER-name}
    # Example: - "X-Custom-Header: value"

wordlists:
  common: /usr/share/wordlists/dirb/common.txt
  big: /usr/share/wordlists/dirb/big.txt
  raft-small-words: /usr/share/seclists/Discovery/Web-Content/raft-small-words.txt
  raft-small-dirs: /usr/share/seclists/Discovery/Web-Content/raft-small-directories.txt
  raft-medium-words: /usr/share/seclists/Discovery/Web-Content/raft-medium-words.txt
  raft-medium-dirs: /usr/share/seclists/Discovery/Web-Content/raft-medium-directories.txt
  raft-medium-files: /usr/share/seclists/Discovery/Web-Content/raft-medium-files.txt
  raft-large-dirs: /usr/share/seclists/Discovery/Web-Content/raft-large-directories.txt
  raft-large-files: /usr/share/seclists/Discovery/Web-Content/raft-large-files.txt
  quickhits: /usr/share/seclists/Discovery/Web-Content/quickhits.txt
  php: /usr/share/seclists/Discovery/Web-Content/Common-PHP-Filenames.txt
  api: /usr/share/seclists/Discovery/Web-Content/api/api-endpoints.txt
  api-v2: /usr/share/seclists/Discovery/Web-Content/api/api-endpoints-res.txt
  params: /usr/share/seclists/Discovery/Web-Content/burp-parameter-names.txt
  swagger: /usr/share/seclists/Discovery/Web-Content/swagger.txt
  graphql: /usr/share/seclists/Discovery/Web-Content/graphql.txt
  subdomain-top5000: /usr/share/seclists/Discovery/DNS/subdomains-top1million-5000.txt
  subdomain-top20000: /usr/share/seclists/Discovery/DNS/subdomains-top1million-20000.txt
  vhosts: /usr/share/seclists/Discovery/DNS/namelist.txt
  backups: /usr/share/seclists/Discovery/Web-Content/backup-files.txt

tools:
  ffuf:
    enabled: true
    tmux_prefix: "ffuf_"
    commands:
      # ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
      # DISCOVERY RÁPIDO - Fast initial scans with small wordlists
      # ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
      - name: "quickhits"
        description: "Fast scan with quickhits wordlist (immediate findings)"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404 -t 100 -rate 100 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-quickhits.json -of json -ac"
        wordlist: quickhits

      - name: "common"
        description: "Common paths and files (dirb common)"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404 -t 100 -rate 100 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-common.json -of json -ac"
        wordlist: common

      - name: "raft-small-dirs"
        description: "Raft small directories wordlist"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404 -t 100 -rate 150 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-raft-small-dirs.json -of json -ac"
        wordlist: raft-small-dirs

      - name: "raft-small-words"
        description: "Raft small words wordlist"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404 -t 100 -rate 150 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-raft-small-words.json -of json -ac"
        wordlist: raft-small-words

      # ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
      # DISCOVERY MEDIO - Medium depth scans with balanced wordlists
      # ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
      - name: "raft-medium-dirs"
        description: "Raft medium directories with recursion"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404,403 -fs 0 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-raft-medium-dirs.json -of json -t 100 -rate 200 -ac -recursion -recursion-depth 2"
        wordlist: raft-medium-dirs

      - name: "raft-medium-words"
        description: "Raft medium words with extensions"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404 -e .php,.html,.js -t 100 -rate 200 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-raft-medium-words.json -of json -ac"
        wordlist: raft-medium-words

      - name: "raft-medium-files"
        description: "Raft medium files with multiple extensions"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404 -e .php,.asp,.aspx,.jsp,.html,.js,.txt,.json,.xml,.bak,.old,.zip,.tar.gz,.sql,.db,.config,.env,.log -t 100 -rate 200 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-raft-medium-files.json -of json"
        wordlist: raft-medium-files

      # ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
      # DISCOVERY PROFUNDO - Deep scans with large wordlists
      # ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
      - name: "big"
        description: "Big wordlist comprehensive scan"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404 -fs 0 -t 80 -rate 150 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-big.json -of json -ac"
        wordlist: big

      - name: "raft-large-dirs"
        description: "Raft large directories extensive scan"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404 -fs 0 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-raft-large-dirs.json -of json -t 80 -rate 150 -ac"
        wordlist: raft-large-dirs

      - name: "raft-large-files"
        description: "Raft large files with extensions"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404 -e .php,.asp,.aspx,.jsp,.html,.js,.txt,.json,.xml,.bak,.old -t 80 -rate 150 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-raft-large-files.json -of json"
        wordlist: raft-large-files

      # ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
      # API DISCOVERY - API endpoints and documentation
      # ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
      - name: "api-endpoints"
        description: "API endpoints discovery (main list)"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -H 'Content-Type: application/json' -mc all -fc 404 -t 100 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-api.json -of json"
        wordlist: api

      - name: "api-endpoints-v2"
        description: "API endpoints discovery (extended list)"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -H 'Content-Type: application/json' -mc all -fc 404 -t 100 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-api-v2.json -of json"
        wordlist: api-v2

      - name: "swagger-docs"
        description: "Swagger/OpenAPI documentation discovery"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404 -t 100 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-swagger.json -of json"
        wordlist: swagger

      - name: "graphql-endpoints"
        description: "GraphQL endpoints discovery"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -H 'Content-Type: application/json' -mc all -fc 404 -t 100 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-graphql.json -of json"
        wordlist: graphql

      # ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
      # FUZZING ESPECIALIZADO - Specialized fuzzing targets
      # ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
      - name: "parameter-fuzzing"
        description: "GET parameter fuzzing"
        command: "ffuf -u {URL}?FUZZ=test -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404 -fs 0 -t 100 -rate 200 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-params.json -of json -ac"
        wordlist: params

      - name: "php-files"
        description: "Common PHP filenames discovery"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404 -t 100 -rate 150 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-php.json -of json -ac"
        wordlist: php

      - name: "backup-files"
        description: "Backup and sensitive files discovery"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} {HEADERS-ALL} -mc all -fc 404 -e .bak,.backup,.old,.swp,~,.git,.env,.sql,.db,.config,.log -t 100 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-backups.json -of json"
        wordlist: backups

      # ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
      # VHOST ENUMERATION - Virtual host and subdomain discovery
      # ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
      - name: "vhost-top5k"
        description: "Virtual host enumeration (top 5000)"
        command: "ffuf -u {URL} -w {WORDLIST} {HEADERS-ALL} -H 'Host: FUZZ.{DOMAIN}' -mc all -fc 404 -fs 0 -t 100 -rate 150 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-vhost-top5k.json -of json"
        wordlist: subdomain-top5000

      - name: "vhost-top20k"
        description: "Virtual host enumeration (top 20000)"
        command: "ffuf -u {URL} -w {WORDLIST} {HEADERS-ALL} -H 'Host: FUZZ.{DOMAIN}' -mc all -fc 404 -fs 0 -t 100 -rate 150 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-vhost-top20k.json -of json"
        wordlist: subdomain-top20000

      - name: "vhost-namelist"
        description: "Virtual host enumeration (namelist)"
        command: "ffuf -u {URL} -w {WORDLIST} {HEADERS-ALL} -H 'Host: FUZZ.{DOMAIN}' -mc all -fc 404 -fs 0 -t 100 -rate 150 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-vhost-namelist.json -of json"
        wordlist: vhosts

  gobuster:
    enabled: true
    tmux_prefix: "gobuster_"
    commands:
      - name: "dir-enum-fast"
        description: "Fast directory enumeration"
        command: "gobuster dir -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/gobuster-{DOMAIN}-dirs.txt -t 100 -k -e -q --no-error -s '200,204,301,302,307,401,403'"
        wordlist: raft-medium-dirs

      - name: "dir-enum-extensions"
        description: "Directory enumeration with multiple extensions"
        command: "gobuster dir -u {URL} -w {WORDLIST} -x php,asp,aspx,jsp,html,js,txt,json,xml,bak,zip,tar.gz,sql -o {OUTPUT_DIR}/gobuster-{DOMAIN}-ext.txt -t 100 -k -e -q --no-error"
        wordlist: raft-medium-files

      - name: "dir-enum-big"
        description: "Extensive directory enumeration"
        command: "gobuster dir -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/gobuster-{DOMAIN}-big.txt -t 80 -k -e -q --no-error -x php,html,txt"
        wordlist: raft-large-dirs

      - name: "api-endpoints"
        description: "API endpoint enumeration"
        command: "gobuster dir -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/gobuster-{DOMAIN}-api.txt -t 100 -k -e -q --no-error -x json,xml"
        wordlist: api

      - name: "vhost-enum"
        description: "Virtual host enumeration"
        command: "gobuster vhost -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/gobuster-{DOMAIN}-vhosts.txt -t 100 -k --append-domain -r"
        wordlist: subdomain-top5000

      - name: "dns-enum"
        description: "DNS subdomain enumeration"
        command: "gobuster dns -d {DOMAIN} -w {WORDLIST} -o {OUTPUT_DIR}/gobuster-{DOMAIN}-dns.txt -t 100"
        wordlist: subdomain-top20000

      - name: "sensitive-files"
        description: "Search for sensitive files"
        command: "gobuster dir -u {URL} -w {WORDLIST} -x bak,backup,old,swp,env,git,sql,db,config,log -o {OUTPUT_DIR}/gobuster-{DOMAIN}-sensitive.txt -t 100 -k -e -q"
        wordlist: backups

      - name: "bypass-filtering"
        description: "Attempt bypass with custom patterns"
        command: "gobuster dir -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/gobuster-{DOMAIN}-bypass.txt -t 100 -k -e -q --no-error -H 'X-Original-URL: /' -H 'X-Rewrite-URL: /'"
        wordlist: common

  dirsearch:
    enabled: true
    tmux_prefix: "dirsearch_"
    commands:
      - name: "default-scan"
        description: "Default fast scan with common wordlist"
        command: "dirsearch -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/dirsearch-{DOMAIN}-default.txt -t 100 --random-agent -i 200,204,301,302,307,401,403 -q"
        wordlist: raft-medium-dirs

      - name: "recursive-scan"
        description: "Recursive directory scanning"
        command: "dirsearch -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/dirsearch-{DOMAIN}-recursive.txt -t 100 -R 2 --random-agent -i 200,204,301,302,307,401,403 -q"
        wordlist: raft-medium-dirs

      - name: "deep-recursive"
        description: "Deep recursive scan (finds more endpoints)"
        command: "dirsearch -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/dirsearch-{DOMAIN}-deep.txt -t 80 --deep-recursive --random-agent -i 200,204,301,302,307,401,403 -q"
        wordlist: raft-medium-words

      - name: "multi-extension"
        description: "Scan with multiple important extensions"
        command: "dirsearch -u {URL} -w {WORDLIST} -e php,asp,aspx,jsp,html,js,txt,json,xml,yml,yaml,bak,old,zip,tar.gz,sql,db,config,env,log -o {OUTPUT_DIR}/dirsearch-{DOMAIN}-multi-ext.txt -t 100 --random-agent -q"
        wordlist: raft-medium-files

      - name: "backup-files"
        description: "Search for backup and sensitive files"
        command: "dirsearch -u {URL} -w {WORDLIST} -e bak,backup,old,swp,save,copy,orig,tmp,~ -o {OUTPUT_DIR}/dirsearch-{DOMAIN}-backups.txt -t 100 --random-agent --suffixes=~ --prefixes=. -q"
        wordlist: backups

      - name: "api-scan"
        description: "API endpoint discovery"
        command: "dirsearch -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/dirsearch-{DOMAIN}-api.txt -t 100 --random-agent -i 200,201,204,301,302,401,403 -q"
        wordlist: api

      - name: "config-files"
        description: "Search for configuration files"
        command: "dirsearch -u {URL} -w {WORDLIST} -e config,conf,cfg,ini,env,xml,yml,yaml,json,properties -o {OUTPUT_DIR}/dirsearch-{DOMAIN}-configs.txt -t 100 --random-agent -q"
        wordlist: common

      - name: "large-scan"
        description: "Comprehensive scan with large wordlist"
        command: "dirsearch -u {URL} -w {WORDLIST} -e php,html,txt,js,json -o {OUTPUT_DIR}/dirsearch-{DOMAIN}-large.txt -t 80 --random-agent -R 1 -q"
        wordlist: raft-large-dirs

      - name: "exclude-sizes"
        description: "Scan excluding common false positive sizes"
        command: "dirsearch -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/dirsearch-{DOMAIN}-filtered.txt -t 100 --random-agent --exclude-sizes=0B -q"
        wordlist: raft-medium-dirs

  feroxbuster:
    enabled: true
    tmux_prefix: "feroxbuster_"
    commands:
      - name: "fast-scan"
        description: "Fast scan with auto-tune"
        command: "feroxbuster -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/feroxbuster-{DOMAIN}-fast.txt -t 100 -k --auto-tune -s 200,204,301,302,307,401,403 -q"
        wordlist: raft-medium-dirs

      - name: "recursive-scan"
        description: "Recursive scan with intelligent depth"
        command: "feroxbuster -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/feroxbuster-{DOMAIN}-recursive.txt -t 100 -k -d 2 --auto-tune -s 200,204,301,302,307,401,403 -q"
        wordlist: raft-medium-dirs

      - name: "deep-recursive"
        description: "Deep recursive scan with word collection"
        command: "feroxbuster -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/feroxbuster-{DOMAIN}-deep.txt -t 80 -k -d 3 --auto-tune --collect-words --extract-links -s 200,204,301,302,307,401,403 -q"
        wordlist: raft-medium-words

      - name: "extensions-scan"
        description: "Scan with multiple extensions"
        command: "feroxbuster -u {URL} -w {WORDLIST} -x php,asp,aspx,jsp,html,js,txt,json,xml,yml,bak,old -o {OUTPUT_DIR}/feroxbuster-{DOMAIN}-ext.txt -t 100 -k --auto-tune -q"
        wordlist: raft-medium-files

      - name: "backup-discovery"
        description: "Discover backup files automatically"
        command: "feroxbuster -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/feroxbuster-{DOMAIN}-backups.txt -t 100 -k --collect-backups --auto-tune -s 200,204,301,302,401,403 -q"
        wordlist: common

      - name: "smart-scan"
        description: "Smart scan with link extraction and word collection"
        command: "feroxbuster -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/feroxbuster-{DOMAIN}-smart.txt -t 100 -k -d 2 --auto-tune --collect-words --extract-links --collect-backups -q"
        wordlist: raft-medium-dirs

      - name: "large-scan"
        description: "Comprehensive scan with large wordlist"
        command: "feroxbuster -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/feroxbuster-{DOMAIN}-large.txt -t 80 -k -d 2 --auto-tune --rate-limit 200 -q"
        wordlist: raft-large-dirs

      - name: "filtered-scan"
        description: "Scan with intelligent size filtering"
        command: "feroxbuster -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/feroxbuster-{DOMAIN}-filtered.txt -t 100 -k --auto-tune --filter-size 0 -C 404 -q"
        wordlist: raft-medium-dirs

      - name: "api-scan"
        description: "API endpoint discovery"
        command: "feroxbuster -u {URL} -w {WORDLIST} -x json,xml -o {OUTPUT_DIR}/feroxbuster-{DOMAIN}-api.txt -t 100 -k --auto-tune -s 200,201,204,401,403 -q"
        wordlist: api

      - name: "thorough-scan"
        description: "Thorough scan for maximum coverage"
        command: "feroxbuster -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/feroxbuster-{DOMAIN}-thorough.txt -t 80 -k -d 3 --auto-tune --collect-words --extract-links --collect-backups --rate-limit 150 -x php,html,js,txt,json,xml,bak,old -q"
        wordlist: raft-large-files

# Notes:
# - Todos los comandos incluyen rate limiting o thread control para evitar bloqueos
# - Se usa auto-calibration (-ac en ffuf) y auto-tune (en feroxbuster) cuando es posible
# - Los comandos de ffuf usan -mc all -fc 404 para capturar todos los códigos excepto 404
# - Se incluyen búsquedas específicas para APIs, backups, y archivos sensibles
# - Los comandos recursivos tienen profundidad limitada (2-3) para eficiencia
# - Se usan random-agent en dirsearch para evitar detección
# - Gobuster usa -k para skip SSL verification, útil en pruebas internas
# - Feroxbuster usa --collect-words y --extract-links para descubrimiento inteligente
# - Se incluyen extensiones críticas: .env, .git, .bak, .sql, .db, .config
# - Los outputs están en formato JSON cuando es posible para parsing automatizado
`
