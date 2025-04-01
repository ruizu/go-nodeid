# NodeID - IPv4-Based Node ID Generator

## Overview
This tool generates a unique node identifier by extracting and utilizing the last two bytes of the local IPv4 address assigned to the instance. This approach provides a simple way to create distinct identifiers for nodes in a network.

## How It Works
The generator performs the following steps:
1. Retrieves the local IPv4 address of the instance
2. Extracts the last two octets (bytes) of the IP address
3. Combines these bytes to create a unique node identifier

For example:
- If local IP is `192.168.9.120`
- Last two bytes are `9` and `120`
- Generated node ID: `2424`

## Installation
```bash
go install github.com/ruizu/go-nodeid
```