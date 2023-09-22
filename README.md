<p align="center">
  <img src="https://github.com/thisisiliya/httpr/assets/66384228/fdcc9f64-dd38-4f65-b597-1074c17720d1" width="400px">
  <br>
  HTTPR is an OSINT tool to Scrape the Undisclosed Data via Search Engines
  <br>
  <a href="#installation">Installation</a> |
  <a href="#modes">Modes</a> |
  <a href="#next-features">Next Features</a> |
  <a href="#license">License</a>
</p>

# Installation
Using go easy installation:
```bash
go install github.com/thisisiliya/httpr@v0.1
```
Note that you need at least *go 1.20v* to compile httpr

# Modes
- [sub](#sub-mode) - algorithmic subdomain enumeration for domains
- [key](#key-mode) - keywords enumeration for domains
- [path](#path-mode) - path enumeration for domains
- [custom](custom-mode) - custom dork command to scrape

![ray-so-export](https://github.com/thisisiliya/httpr/assets/66384228/33aff54d-8275-4522-b8be-d5329485d821)

## `sub` Mode
```
>>> httpr help sub

algorithmic subdomain enumeration for domain(s)
usage: -d google.com

Usage:
  httpr sub [flags]

Flags:
  -a, --all              redo the process for the result
      --depth int        number of pages to scrape per result (default 5)
  -d, --domain string    target domain to search subdomains
      --domains string   target domains file path
  -h, --help             help for sub
      --show-sub         show subdomains as result
      --show-url         show URLs as result
```

## `key` Mode
```
>>> httpr help key

keyword(s) enumeration for domain(s)
usage: -d www.google.com -k exploit --depth 3

Usage:
  httpr key [flags]

Flags:
      --depth int         number of pages to scrape per key (default 3)
  -d, --domain string     target domain to search keyword(s)
      --domains string    target domains file path
  -h, --help              help for key
  -k, --keyword string    target keyword to search
      --show-host         show hosts as result
      --show-path         show paths as result
      --show-sub          show subdomains as result
```

## `path` Mode
```
>>> httpr help path

path enumeration for domain(s)
usage: -d www.google.com --depth 20

Usage:
  httpr path [flags]

Flags:
      --depth int        number of pages to scrape per domain (default 20)
  -d, --domain string    target domain to search
      --domains string   target domains file path
  -h, --help             help for path
      --show-path        show paths as result
```

## `custom` Mode
```
>>> httpr help custom

google page(s) scrape by custom dork commands
usage: -c site:www.google.com,inurl:map -t google.com --depth 1

Usage:
  httpr custom [flags]

Flags:
  -c, --command string       dork command to scrape
      --depth int            number of pages to scrape (default 1)
  -h, --help                 help for custom
      --show-host            show hosts as result
      --show-path            show paths as result
      --show-sub             show subdomains as result
      --split-by string      dork commands split character (default ",")
  -t, --target-host string   filter result by host
```

# Next features
- More search engines support
- Colored output
- Verbose output
- New modes

# Support
Support me by a ⭐

# License
Please visit [License](https://github.com/thisisiliya/httpr/blob/main/LICENSE) file
