# cmus-time

Very simple application which communicates with `cmus` Linux music player
to fetch current song's details.

It uses default Linux socket location to communicate with cmus:
`/run/user/$UID/cmus-socket`.

It's meant to be run from `polybar` so it does what it needs to do quickly 
and then it exits.

## Installation

Application is built using Go 1.10.

```bash
go get github.com/milanaleksic/cmus-time
```

## Running

> Don't forget to run cmus otherwise you will (obviously) get an error

```
milan → time cmus-time
Nightwish - Kuolema Tekee Taiteilijan [03:43 / 04:04]

real    0m0.004s
user    0m0.004s
sys     0m0.000s
```