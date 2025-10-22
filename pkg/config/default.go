package config

// DefaultConfig returns a default configuration template
const DefaultConfig = `# Trident Recon Configuration File
# Save this at ~/.config/trident-recon/config.yaml

global:
  output_dir: ~/trident-output
  id_length: 12

headers:
  default:
    User-Agent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36"
    Accept: "*/*"
  custom:
    - "X-Bug-Bounty: true"

wordlists:
  common: /usr/share/wordlists/dirb/common.txt
  big: /usr/share/wordlists/dirb/big.txt
  raft-small: /usr/share/seclists/Discovery/Web-Content/raft-small-words.txt
  raft-medium: /usr/share/seclists/Discovery/Web-Content/raft-medium-words.txt
  raft-large: /usr/share/seclists/Discovery/Web-Content/raft-large-words.txt
  php: /usr/share/seclists/Discovery/Web-Content/Common-PHP-Filenames.txt
  asp: /usr/share/seclists/Discovery/Web-Content/Common-ASP-Filenames.txt
  api: /usr/share/seclists/Discovery/Web-Content/api/api-endpoints.txt
  params: /usr/share/seclists/Discovery/Web-Content/burp-parameter-names.txt

tools:
  ffuf:
    enabled: true
    tmux_prefix: "ffuf_"
    commands:
      - name: "content-discovery"
        description: "Content discovery with common wordlist"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} -mc 200,204,301,302,307,401,403 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-content.json -of json -t 40"
        wordlist: common

      - name: "content-discovery-big"
        description: "Content discovery with big wordlist"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} -mc 200,204,301,302,307,401,403 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-big.json -of json -t 40"
        wordlist: big

      - name: "php-files"
        description: "PHP file fuzzing"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} -mc 200,204,301,302,307,401,403 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-php.json -of json -t 40"
        wordlist: php

      - name: "asp-files"
        description: "ASP file fuzzing"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} -mc 200,204,301,302,307,401,403 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-asp.json -of json -t 40"
        wordlist: asp

      - name: "api-endpoints"
        description: "API endpoint discovery"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} -mc 200,204,301,302,307,401,403 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-api.json -of json -t 40"
        wordlist: api

      - name: "recursive-scan"
        description: "Recursive directory discovery"
        command: "ffuf -u {URL}/FUZZ -w {WORDLIST} -mc 200,204,301,302,307,401,403 -recursion -recursion-depth 2 -o {OUTPUT_DIR}/ffuf-{DOMAIN}-recursive.json -of json -t 40"
        wordlist: common

  gobuster:
    enabled: true
    tmux_prefix: "gobuster_"
    commands:
      - name: "dir-enum"
        description: "Directory enumeration with gobuster"
        command: "gobuster dir -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/gobuster-{DOMAIN}-dirs.txt -t 40 -k"
        wordlist: common

      - name: "dir-enum-big"
        description: "Directory enumeration with big wordlist"
        command: "gobuster dir -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/gobuster-{DOMAIN}-big.txt -t 40 -k"
        wordlist: big

      - name: "vhost-enum"
        description: "Virtual host enumeration"
        command: "gobuster vhost -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/gobuster-{DOMAIN}-vhosts.txt -t 40 -k"
        wordlist: raft-small

  dirsearch:
    enabled: true
    tmux_prefix: "dirsearch_"
    commands:
      - name: "default-scan"
        description: "Default dirsearch scan"
        command: "dirsearch -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/dirsearch-{DOMAIN}-default.txt -t 40"
        wordlist: common

      - name: "recursive-scan"
        description: "Recursive dirsearch scan"
        command: "dirsearch -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/dirsearch-{DOMAIN}-recursive.txt -t 40 -R 2"
        wordlist: common

  feroxbuster:
    enabled: true
    tmux_prefix: "feroxbuster_"
    commands:
      - name: "default-scan"
        description: "Default feroxbuster scan"
        command: "feroxbuster -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/feroxbuster-{DOMAIN}-default.txt -t 40 -k"
        wordlist: common

      - name: "recursive-scan"
        description: "Recursive feroxbuster scan"
        command: "feroxbuster -u {URL} -w {WORDLIST} -o {OUTPUT_DIR}/feroxbuster-{DOMAIN}-recursive.txt -t 40 -k --depth 2"
        wordlist: common

  httpx:
    enabled: true
    tmux_prefix: "httpx_"
    commands:
      - name: "probe"
        description: "HTTP probe with httpx"
        command: "echo {URL} | httpx -o {OUTPUT_DIR}/httpx-{DOMAIN}-probe.txt -json -title -tech-detect -status-code"
        wordlist: ""

      - name: "full-scan"
        description: "Full httpx scan with all features"
        command: "echo {URL} | httpx -o {OUTPUT_DIR}/httpx-{DOMAIN}-full.txt -json -title -tech-detect -status-code -follow-redirects -screenshot -system-chrome"
        wordlist: ""

  nuclei:
    enabled: true
    tmux_prefix: "nuclei_"
    commands:
      - name: "default-scan"
        description: "Default nuclei scan"
        command: "nuclei -u {URL} -o {OUTPUT_DIR}/nuclei-{DOMAIN}-default.txt -severity critical,high,medium"
        wordlist: ""

      - name: "full-scan"
        description: "Full nuclei scan with all templates"
        command: "nuclei -u {URL} -o {OUTPUT_DIR}/nuclei-{DOMAIN}-full.txt -severity critical,high,medium,low,info"
        wordlist: ""

      - name: "cves"
        description: "CVE scanning with nuclei"
        command: "nuclei -u {URL} -o {OUTPUT_DIR}/nuclei-{DOMAIN}-cves.txt -tags cve"
        wordlist: ""

  arjun:
    enabled: true
    tmux_prefix: "arjun_"
    commands:
      - name: "param-discovery"
        description: "Parameter discovery with arjun"
        command: "arjun -u {URL} -o {OUTPUT_DIR}/arjun-{DOMAIN}-params.json -oJ"
        wordlist: ""

  katana:
    enabled: true
    tmux_prefix: "katana_"
    commands:
      - name: "crawl"
        description: "Web crawling with katana"
        command: "katana -u {URL} -o {OUTPUT_DIR}/katana-{DOMAIN}-crawl.txt -d 3 -jc -kf all"
        wordlist: ""

      - name: "js-crawl"
        description: "JavaScript crawling with katana"
        command: "katana -u {URL} -o {OUTPUT_DIR}/katana-{DOMAIN}-js.txt -jc -kf all -ef css,png,jpg,jpeg,gif,svg"
        wordlist: ""

  waybackurls:
    enabled: true
    tmux_prefix: "wayback_"
    commands:
      - name: "fetch-urls"
        description: "Fetch URLs from Wayback Machine"
        command: "echo {DOMAIN} | waybackurls > {OUTPUT_DIR}/waybackurls-{DOMAIN}.txt"
        wordlist: ""

  gau:
    enabled: true
    tmux_prefix: "gau_"
    commands:
      - name: "fetch-urls"
        description: "Fetch URLs with gau"
        command: "echo {DOMAIN} | gau --o {OUTPUT_DIR}/gau-{DOMAIN}.txt"
        wordlist: ""

  gospider:
    enabled: true
    tmux_prefix: "gospider_"
    commands:
      - name: "crawl"
        description: "Web crawling with gospider"
        command: "gospider -s {URL} -o {OUTPUT_DIR}/gospider-{DOMAIN} -c 10 -d 3"
        wordlist: ""

  hakrawler:
    enabled: true
    tmux_prefix: "hakrawler_"
    commands:
      - name: "crawl"
        description: "Web crawling with hakrawler"
        command: "echo {URL} | hakrawler -d 3 > {OUTPUT_DIR}/hakrawler-{DOMAIN}.txt"
        wordlist: ""

  dalfox:
    enabled: false
    tmux_prefix: "dalfox_"
    commands:
      - name: "xss-scan"
        description: "XSS scanning with dalfox"
        command: "dalfox url {URL} -o {OUTPUT_DIR}/dalfox-{DOMAIN}.txt"
        wordlist: ""

  sqlmap:
    enabled: false
    tmux_prefix: "sqlmap_"
    commands:
      - name: "sqli-scan"
        description: "SQL injection scan with sqlmap"
        command: "sqlmap -u {URL} --batch --output-dir={OUTPUT_DIR}/sqlmap-{DOMAIN}"
        wordlist: ""
`
