<p align="center">
  <img src="https://github.com/thisisiliya/httpr/assets/66384228/5495f1de-eebd-4fb3-a540-3c2af81f248b" width="400px">
  <br>
  HTTPR is an OSINT tool to Scrape the Undisclosed Data via Search Engines
  <br>
  Available Engines: Google, Bing, Yahoo 
  <br>
  <a href="#modes">Modes</a> |
  <a href="#installation">Installation</a> |
  <a href="#next-features">Next Features</a> |
  <a href="#license">License</a>
</p>

# Modes

- [sub](#sub-mode) - algorithmic subdomain enumeration for domains
- [key](#key-mode) - keywords enumeration for domains
- [path](#path-mode) - path enumeration for domains
- [custom](custom-mode) - custom dork command to scrape

![ray-so](https://github.com/thisisiliya/httpr/assets/66384228/750e3662-38b1-4211-9096-e46a08ec4bce)

## `sub` Mode
### Options
```
>>> httpr help sub
algorithmic subdomain enumeration for domain(s)

Usage:
  httpr sub [flags]

Examples:
httpr sub --domain hackerone.com --all

Flags:
  -a, --all              redo the process for the result
      --depth int        number of pages to scrape per result (default 5)
  -d, --domain string    target domain to search subdomains
      --domains string   target domains file path
  -h, --help             help for sub
      --show-sub         show subdomains as result
      --show-url         show URLs as result
```
### Example
```
>>> httpr sub -d hackerone.com -a --silent
www.hackerone.com
docs.hackerone.com
api.hackerone.com
...
```

## `key` Mode
### Options
```
>>> httpr help key
keyword(s) enumeration for domain(s)

Usage:
  httpr key [flags]

Examples:
httpr key --domain hackerone.com --keyword report --depth 3

Flags:
      --depth int         number of pages to scrape per key (default 3)
  -d, --domain string     target domain to search keyword(s)
      --domains string    target domains file path
  -h, --help              help for key
  -k, --keyword string    target keyword to search
      --keywords string   target keywords path
      --show-host         show hosts as result
      --show-path         show paths as result
      --show-sub          show subdomains as result
```
### Example
```
>>> httpr key -d hackerone.com -k report --silent
https://hackerone.com/reports/647130
https://hackerone.com/directory/programs
https://www.hackerone.com/hacker-powered-security-report-0
...
```

## `path` Mode
### Options
```
>>> httpr help path
path enumeration for domain(s)

Usage:
  httpr path [flags]

Examples:
usage: --domain hackerone.com --depth 10

Flags:
      --depth int        number of pages to scrape per domain (default 10)
  -d, --domain string    target domain to search
      --domains string   target domains file path
  -h, --help             help for path
      --show-path        show paths as result
```
### Example 
```
>>> httpr path -d hackerone.com --silent
https://hackerone.com/enter
https://hackerone.com/telegram
https://hackerone.com/rockstargames
...
```

## `custom` Mode
### Options
```
>>> httpr help custom
engine page(s) scrape by custom dork commands

Usage:
  httpr custom [flags]

Examples:
httpr custom --command "site:*.hackerone.com inurl:report" --target-host hackerone.com --engine Google --depth 1

Flags:
  -c, --command string       dork command to scrape
      --depth int            number of pages to scrape (default 1)
  -e, --engine string        target engine to scrape. available engines:Google, Bing, Yahoo (default "Google")
  -h, --help                 help for custom
      --show-host            show hosts as result
      --show-path            show paths as result
      --show-sub             show subdomains as result
      --split-by string      dork commands split character (default " ")
  -t, --target-host string   filter result by host
```
### Example
```
>>> httpr custom -c "site:*.hackerone.com inurl:report" --target-host hackerone.com --silent
https://docs.hackerone.com/en/articles/8475030-report-states
https://docs.hackerone.com/en/articles/8474574-report-actions
https://www.hackerone.com/reports/7th-annual-hacker-powered-security-report
...
```

# Installation
1. Go easy installation:
```bash
go install github.com/thisisiliya/httpr@lastest
```
Note that you need at least *go 1.20v* to compile httpr

2. Download the windows version for [release](https://github.com/thisisiliya/httpr/releases)

3. Or compile it from the [Source code](https://github.com/thisisiliya/httpr/releases) file

# Next features
- More search engines support
- New modes

# Support
Support me by a ⭐

# License
Please visit [License](https://github.com/thisisiliya/httpr/blob/main/LICENSE) file
