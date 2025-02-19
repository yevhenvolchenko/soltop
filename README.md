# It's like TOP, but for solana

## Install
First download and install go 1.23.6 from [https://go.dev/dl/](https://go.dev/dl/)

To install soltop simply run
```bash
  go install github.com/yevhenvolchenko/soltop@latest
```

## Usage
`soltop validators` outputs top 100 validators by stake list. The output is structured in a way that it can be used as a part of SQL compatible predicate.

Example:
```bash
echo "SELECT * FROM cpu_load WHERE node IN $(go run main.go validators) - 1h"
```
