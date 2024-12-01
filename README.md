# RDT

RDT "Redirect Data Traffic" runs on a local host redirecting network traffic from local LAN address to localhost/127.0.0.1 or another other `<LAN_IP_ADDRESS>` that the main PC has connection too.

## Requirements
- Go 1.21 or higher
- Linux/MacOS/Windows

## Usability example

This example demonstrates how to use RDT to access GPT4All queries and LLM on the local area network when hosting GPT4All. Follow these steps to download the model, start RDT, and access it using the LAN address rather than localhost.

## Download the GPT4All model
   Download the desired model from the GPT4All repository or [website](https://docs.gpt4all.io/index.html).

## Start the GPT4All server
   Start the GPT4All server on your machine, Download your favourite LLM model!

## Setup GTP4All API Server
1. Download your favourite LLM model!
2. Follow instructions in [GPT4All API Server](https://docs.gpt4all.io/gpt4all_api_server/home.html)


## Starting with RDT
   
Starting with RDT you will need to download, build it on your machine, ensuring config.json is configured to listen on LAN network `<LAN_IP_ADDRESS>`.

## Downloading and building RDT
```bash
# Clone the repository
git clone https://github.com/gcclinux/rdt.git
cd rdt

# Build the binary (linux)
go build -o rdt-$(arch)-$(uname -s)

# Build the binary (windows)
go build -o rdt.exe
```

*The application uses config.json to define inbound and outbound connections. Create or modify the config.json file in the same directory as the binary:*
```
{
    "inbound_address": "192.168.0.45",  // Address to listen on
    "inbound_port": "4891",             // Port to listen on
    "outbound_address": "127.0.0.1",    // Target address to forward to
    "outbound_port": "4891",            // Target port to forward to
    "verbose": "true"                   // Verbose set will log inbound and outbound traffic
}
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

## Access GPT4All using the LAN address

Replace `localhost` with the LAN IP address of the machine hosting GPT4All. You can find the LAN IP address using the `ipconfig` command on Windows or `ifconfig`/`ip a` command on Linux/Mac.

```sh
curl -X POST http://<LAN_IP_ADDRESS>:4891/v1/chat/completions -d '{
"model": "Phi-3 Mini Instruct",
"messages": [{"role":"user","content":"Who is Lionel Messi?"}],
"max_tokens": 50,
"temperature": 0.28
}'
```

Replace `<LAN_IP_ADDRESS>` with the actual IP address of the machine hosting GPT4All and configured as "inbound_address" in the config.json.

By following these steps, you can access GPT4All queries and LLM on the local area network using the LAN address.

## ChangeLog
- 1.0.0 - Initial version with single port redirect