---
title: Installation
description: How to download and install the pig Package Manager
icon: fa-solid fa-download
weight: 300
breadcrumb: false
---

Install PIG with the Linux installer, a system package, or a release archive. The current version is **v{{< param pig_version >}}**. The [download page](https://pig.pgsty.com/download/) lists all release files.

## Script installation

On Linux, run:

```bash
curl -fsSL https://repo.pigsty.io/pig | bash
```

In mainland China, use:

```bash
curl -fsSL https://repo.pigsty.cc/pig | bash
```

The script detects the architecture, downloads the RPM or DEB package, and installs it with `rpm` or `dpkg`. Installation requires root or `sudo` access.

Check the installed version:

```bash
pig version
```

To select a published version, pass its version number:

```bash
curl -fsSL https://repo.pigsty.io/pig | bash -s {{< param pig_version >}}
```

## Download release artifacts

Choose a file for your operating system and architecture on the [download page](https://pig.pgsty.com/download/#packages) or [GitHub release](https://github.com/pgsty/pig/releases/tag/v{{< param pig_version >}}).

| System | Format | Architecture names |
|:---|:---|:---|
| EL | RPM | `x86_64`, `aarch64` |
| Debian, Ubuntu | DEB | `amd64`, `arm64` |
| Linux | tar.gz | `amd64`, `arm64` |
| macOS | tar.gz | `amd64` (Intel), `arm64` (Apple Silicon) |

For a release archive, extract it and install the `pig` executable in a directory on your `PATH`. For example, on Apple Silicon:

```bash
curl -fLO https://github.com/pgsty/pig/releases/download/v{{< param pig_version >}}/pig-v{{< param pig_version >}}.darwin-arm64.tar.gz
tar -xzf pig-v{{< param pig_version >}}.darwin-arm64.tar.gz
sudo mkdir -p /usr/local/bin
sudo install -m 0755 pig /usr/local/bin/pig
pig version
```

The release includes [SHA-256 checksums](https://github.com/pgsty/pig/releases/download/v{{< param pig_version >}}/checksums.txt). Use `sha256sum` on Linux or `shasum -a 256` on macOS to compare your downloaded file with its listed hash.

The macOS binary supports catalog queries, inventory editing, and other commands that do not need Linux package managers. Install PostgreSQL packages on a [supported Linux host](/pig/intro/#linux-compatibility).

## Repository installation

To manage PIG with the system package manager, add the [Pigsty Infra repository](https://pigsty.io/docs/repo/infra/). Choose the tab for your host:

```bash {tab="APT (Debian / Ubuntu)" group="install-pig-repository" value="apt" copy="all"}
sudo tee /etc/apt/sources.list.d/pigsty-infra.list > /dev/null <<EOF
deb [trusted=yes] https://repo.pigsty.io/apt/infra generic main
EOF

sudo apt update
sudo apt install -y pig
```

```bash {tab="DNF (EL)" value="dnf" copy="all"}
sudo tee /etc/yum.repos.d/pigsty-infra.repo > /dev/null <<-'EOF'
[pigsty-infra]
name=Pigsty Infra for $basearch
baseurl=https://repo.pigsty.io/yum/infra/$basearch
enabled=1
gpgcheck=0
module_hotfixes=1
EOF

sudo dnf makecache
sudo dnf install -y pig
```

These definitions disable signature checks for the Pigsty repository. For key handling and verification settings, see [repository keys](/pig/cmd/repo/#repo-key).

## Update

On Linux, update the executable with:

```bash
sudo pig update
```

Use `-m` for the China mirror or `-v {{< param pig_version >}}` to select that version. For an archive installation on macOS, download the new archive and replace the installed executable.

Refresh the extension catalog separately:

```bash
pig ext reload
```

This updates catalog metadata, leaving installed packages unchanged.

## Uninstall

Choose the command that matches your installation:

```bash
sudo apt remove pig              # Debian / Ubuntu package
sudo dnf remove pig              # EL package
sudo rm /usr/local/bin/pig        # Archive installed at this path
```

If you installed the archive elsewhere, remove that executable instead. PostgreSQL and its extensions continue to run without PIG.

## Build from source

Install Git, Make, and the Go toolchain specified in the project's [go.mod](https://github.com/pgsty/pig/blob/v{{< param pig_version >}}/go.mod). Build the current release:

```bash
git clone --branch v{{< param pig_version >}} https://github.com/pgsty/pig.git
cd pig
go mod download
make build
./pig version
```

`make build` produces `./pig` for the current platform. The release workflow uses GoReleaser to produce RPM, DEB, and tar.gz files.

Continue with [Getting started](/pig/start/) to install PostgreSQL and enable an extension.
