# VPS Manager

[![Go Version](https://img.shields.io/github/go-mod/go-version/YasiruLaki/vps)](https://github.com/YasiruLaki/vps/blob/main/go.mod)
[![Release](https://img.shields.io/github/v/release/YasiruLaki/vps)](https://github.com/YasiruLaki/vps/releases)
[![License](https://img.shields.io/github/license/YasiruLaki/vps)](https://github.com/YasiruLaki/vps/blob/main/LICENSE)

A lightweight CLI tool to save, list, remove, and connect to VPS servers quickly using named entries.

## Features

- Add VPS entries with name, IP, username, and port
- List all saved VPS entries
- Remove VPS entries by name
- Connect to a VPS using SSH by name
- Store VPS data locally in JSON

## Pre-Install Requirements

Before installing VPS Manager, make sure you have:

- Go 1.20 or newer
- An SSH client available on your machine (`ssh` command)
- Access to your Go binary path in `PATH`

Check your environment:

```bash
go version
ssh -V
echo "$PATH"
```

## Installation

Install with Go:

```bash
go install github.com/YasiruLaki/vps@latest
```

Repository:

- https://github.com/YasiruLaki/vps

After installation, make sure your Go bin directory is in your `PATH`.

Typical location:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Usage

```bash
vps <command> [arguments]
```

## Commands

### Add a VPS

```bash
vps add <name> <ip> <username> [port]
```

- `name`: Friendly name for the server (example: `prod-app`)
- `ip`: Server IP address
- `username`: SSH username
- `port`: Optional, defaults to `22`

Example:

```bash
vps add prod-app 203.0.113.25 ubuntu 22
```

### List VPS entries

```bash
vps list
```

### Remove a VPS

```bash
vps remove <name>
```

Example:

```bash
vps remove prod-app
```

### Connect to a VPS

```bash
vps connect <name>
```

Example:

```bash
vps connect prod-app
```

## Data Storage

VPS records are stored in a local JSON file:

- `vps_data.json`

Keep this file private because it contains server connection metadata.

## Typical Workflow

```bash
vps add web-1 198.51.100.10 ubuntu 22
vps list
vps connect web-1
vps remove web-1
```

## Troubleshooting

- `command not found: vps`
  - Ensure Go bin path is added to your `PATH`.
- `Failed to connect to VPS`
  - Verify IP, username, SSH key access, and firewall rules.
- Duplicate name errors
  - Use unique VPS names for each entry.

## Uninstall

Remove the installed binary from your Go bin directory:

```bash
rm "$(go env GOPATH)/bin/vps"
```

## License

This project is licensed under the terms in the `LICENSE` file.
