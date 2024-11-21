# rdt

RDT "Redirect Data Traffic" runs on a local host redirecting network traffic from local LAN address to localhost/127.0.0.1

## Requirements
- Go 1.21 or higher
- Linux/MacOS/Windows

## Building
```bash
# Clone the repository
git clone https://github.com/gcclinux/rdt.git
cd rdt

# Build the binary (linux)
go build -o rdt-$(arch)-$(uname -s)

# Build the binary (windows)
go build -o rdt.exe
```

## Run directly with Go
```
go run main.go
```

## Or run the built binary

```bash
# Linux
./rdt-$(arch)-$(uname -s)

# Windows

./rdt.exe
```

## Configuration

The application uses config.json to define inbound and outbound connections. Create or modify the config.json file in the same directory as the binary:

```
{
    "inbound_address": "192.168.0.45",  // Address to listen on
    "inbound_port": "4891",             // Port to listen on
    "outbound_address": "127.0.0.1",    // Target address to forward to
    "outbound_port": "4891",            // Target port to forward to
    "verbose": "true"                   // Verbose set will log inbound and outbound traffic
}
```