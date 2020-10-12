# sweeper
Tool to do SYN sweeps across given CIDR and port ranges. 

## Building
Run `make`. This will build a `sweeper` binary in the `build` folder

## Usage
```
sweeper -cidr 192.168.0.1/24 -ports 53-443 -workers 100 -timeout 100
```
Other flags:
* `local` - Polls the connected interfaces and uses the networks it finds
* `ipv4` - Use only specified IPv4 networks
* `ipv6` - Use only specified IPv6 networks