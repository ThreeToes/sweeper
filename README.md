# sweeper
Tool to do SYN sweeps across given CIDR and port ranges. 

## Building
Run `make`. This will build a `sweeper` binary in the `build` folder

## Usage
```
sweeper -cidr 192.168.0.1/24 -ports 53-443 -workers 100 -timeout 100
```