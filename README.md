# rd-traffic

Redirect traffic runs local server to mimic localhost traffic when required

## Requirements
- Go 1.21 or higher
- Linux/MacOS/Windows

## Building
```bash
# Clone the repository
git clone https://github.com/gcclinux/rd-traffic
cd rd-traffic

# Build the binary
go build -o rd-traffic-$(arch)-$(uname -s)
```

## Run directly with Go
```
go run main.go
```

## Or run the built binary
```
./rd-traffic-$(arch)-$(uname -s)
```

## Configuration

The application uses config.json to define inbound and outbound connections. Create or modify the config.json file in the same directory as the binary:

```
{
    "inbound_address": "192.168.0.45",  // Address to listen on
    "inbound_port": "4891",             // Port to listen on
    "outbound_address": "127.0.0.1",    // Target address to forward to
    "outbound_port": "4891"             // Target port to forward to
}
```