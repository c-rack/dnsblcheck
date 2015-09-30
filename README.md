# dnsblcheck, reads one or more DNSBL hosts from stdin and checks if the IP address is listed

dnsblcheck is a command-line utility written in [Go](http://golang.org/) to check if an
IP address is listed at a [DNS-based Blackhole List (DNSBL)](https://en.wikipedia.org/wiki/DNSBL).

An example list of DNSBL hosts is included ([original source](http://blog.penumbra.be/2010/02/zabbix-monitor-dns-blacklists/)).

If the IP address is listed, the DNS TXT record is printed to stdout and program exits with status code 1.


## Installation

You need a working [Go installation](http://golang.org/doc/install) to compile dnsblcheck.

```bash
git clone https://github.com/c-rack/dnsblcheck
cd dnsblcheck
go build
```

## Usage

```bash
cat dnsbl.txt | ./dnsblcheck 127.0.0.1
```

## License

Copyright (C) 2013 [Constantin Rack](http://twitter.com/ConstantinRack)

Licensed under [the MIT License](http://opensource.org/licenses/MIT).
