---
title: Installation
description: How to download and install the pig Package Manager
icon: Download
weight: 300
breadcrumbs: false
---

## Script Installation

The simplest way to install `pig` is to run the following installation script:

**Default Installation** (Cloudflare CDN):

```bash
curl -fsSL https://repo.pigsty.io/pig | bash
```

**China Mirror**:

```bash
curl -fsSL https://repo.pigsty.cc/pig | bash
```

This script downloads the latest `pig` RPM/DEB package from the Pigsty [software repository](/repo/) and installs it using `rpm` or `dpkg`.
Script installation targets Linux x86_64 / aarch64 RPM or DEB distributions. On macOS, use the binary from the release tarball.


## Specify Version

You can specify a particular version to install by passing the version number as an argument:

**Default Installation** (Cloudflare CDN):

```bash
curl -fsSL https://repo.pigsty.io/pig | bash -s 1.5.1
```

**China Mirror**:

```bash
curl -fsSL https://repo.pigsty.cc/pig | bash -s 1.5.1
```


## Download Release Artifacts

You can also download `pig` installation packages (`RPM`/`DEB`/tarball) directly from the Pigsty software repository. Current `v1.5.1` artifacts are available at:

- <https://repo.pigsty.io/pkg/pig/v1.5.1/>
- <https://repo.pigsty.cc/pkg/pig/v1.5.1/>

```
latest
└── v1.5.1
    ├── pig_1.5.1-1_amd64.deb
    ├── pig_1.5.1-1_arm64.deb
    ├── pig-1.5.1-1.aarch64.rpm
    ├── pig-1.5.1-1.x86_64.rpm
    ├── pig-v1.5.1.linux-amd64.tar.gz
    ├── pig-v1.5.1.linux-arm64.tar.gz
    ├── pig-v1.5.1.darwin-amd64.tar.gz
    └── pig-v1.5.1.darwin-arm64.tar.gz
```

After extracting, place the binary file in your system PATH.


## Repository Installation

The `pig` software is located in the [`pigsty-infra`](/repo/infra/) repository. You can add this repository to your operating system and then install using the OS package manager:

### YUM

For RHEL, RockyLinux, CentOS, Alma Linux, OracleLinux, and other EL distributions:

```bash
sudo tee /etc/yum.repos.d/pigsty-infra.repo > /dev/null <<-'EOF'
[pigsty-infra]
name=Pigsty Infra for $basearch
baseurl=https://repo.pigsty.io/yum/infra/$basearch
enabled = 1
gpgcheck = 0
module_hotfixes=1
EOF

sudo yum makecache;
sudo yum install -y pig
```

### APT

For Debian, Ubuntu, and other DEB distributions:

```bash
sudo tee /etc/apt/sources.list.d/pigsty-infra.list > /dev/null <<EOF
deb [trusted=yes] https://repo.pigsty.io/apt/infra generic main
EOF

sudo apt update;
sudo apt install -y pig
```


## Update

To upgrade an existing `pig` version to the latest available version, use the following command:

```bash
pig update            # Upgrade pig itself to the latest version
pig update -m         # Upgrade through the pigsty.cc mirror
```

To update the extension data of an existing `pig` to the latest available version, use the following command:

```bash
pig ext reload        # Update pig extension data to the latest version
```


## Uninstall

```bash
apt remove -y pig     # Debian / Ubuntu and other Debian-based systems
yum remove -y pig     # RHEL / CentOS / RockyLinux and other EL distributions
rm -rf /usr/bin/pig   # If installed directly from binary, just delete the binary file
```


## Build from Source

You can also build `pig` yourself. `pig` is developed in Go and is very easy to build. The source code is hosted at [github.com/pgsty/pig](https://github.com/pgsty/pig)

```bash
git clone https://github.com/pgsty/pig.git; cd pig
go mod download
make build
```

All RPM/DEB packages are automatically built through GitHub CI/CD workflow using goreleaser.
