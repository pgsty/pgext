---
title: Repository
description: The infrastructure to deliver PostgreSQL Extensions
icon: Warehouse
weight: 300
breadcrumbs: false
---


Pigsty has a repository that provides 200+ extra PostgreSQL extensions on 10
mainstream [Linux Distros](/os).
It is designed to work together with the official PostgreSQL Global Development
Group ([PGDG](https://www.postgresql.org/download/linux/)) repo.
Together, they can provide up to [461 PostgreSQL Extensions](https://ext.pgsty.com/list) out-of-the-box.

{{< cards cols=2 >}}
{{< card link="/repo/pgsql"  title="PGSQL Repo" subtitle="Pigsty Extension Repository"  icon="play"     >}}
{{< card link="/repo/infra"  title="INFRA Repo" subtitle="Pigsty Infrastructure Repo"   icon="sparkles" >}}
{{< /cards >}}

|  OS / Arch   |  OS  |                                                                                                                        x86_64                                                                                                                         |                                                                                                                        aarch64                                                                                                                        |
|:------------:|:----:|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|
|     EL8      | el8  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |
|     EL9      | el9  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |
|     EL10     | el10 | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |
|  Debian 12   | d12  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |
|  Debian 13   | d13  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |
| Ubuntu 22.04 | u22  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |
| Ubuntu 24.04 | u24  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}} |

## Get Started

You can enable the pigsty [infra](/repo/infra) & [pgsql](/repo/pgsql) repo with the [pig](/pig/) CLI tool:

{{< tabs items="Default,Mirror" defaultIndex="0" >}}
{{< tab >}}

```bash tab="pig"
curl https://repo.pigsty.io/pig | bash      # download and install the pig CLI tool
pig repo add all -u                         # add linux, pgdg, pigsty repo and update cache
```

{{< /tab >}}

{{< tab >}}

```bash tab="pig"
curl https://repo.pigsty.cc/pig | bash      # download from mirror site
pig repo add -u                             # add linux, pgdg, pigsty repo and update cache
```

{{< /tab >}}
{{< /tabs >}}

## Manual Install

You can also add these repo to your system [manually](#manual-install) with default `apt`, `dnf`, `yum` approach.

{{< tabs items="APT,DNF/YUM" defaultIndex="0" >}}
{{< tab >}}

```bash tab="apt"
# Add Pigsty's GPG public key to your system keychain to verify package signatures
curl -fsSL https://repo.pigsty.io/key | sudo gpg --dearmor -o /etc/apt/keyrings/pigsty.gpg

# Get Debian distribution codename (distro_codename=jammy, focal, bullseye, bookworm), and write the corresponding upstream repository address to the APT List file
distro_codename=$(lsb_release -cs)
sudo tee /etc/apt/sources.list.d/pigsty-io.list > /dev/null <<EOF
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.io/apt/infra generic main
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.io/apt/pgsql/${distro_codename} ${distro_codename} main
EOF

# Refresh APT repository cache
sudo apt update
```

{{< /tab >}}

{{< tab >}}

```bash tab="yum"
# Add Pigsty's GPG public key to your system keychain to verify package signatures
curl -fsSL https://repo.pigsty.io/key | sudo tee /etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty >/dev/null

# Add Pigsty Repo definition files to /etc/yum.repos.d/ directory, including two repositories
sudo tee /etc/yum.repos.d/pigsty-io.repo > /dev/null <<-'EOF'
[pigsty-infra]
name=Pigsty Infra for $basearch
baseurl=https://repo.pigsty.io/yum/infra/$basearch
skip_if_unavailable = 1
enabled = 1
priority = 1
gpgcheck = 1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty
module_hotfixes=1

[pigsty-pgsql]
name=Pigsty PGSQL For el$releasever.$basearch
baseurl=https://repo.pigsty.io/yum/pgsql/el$releasever.$basearch
skip_if_unavailable = 1
enabled = 1
priority = 1
gpgcheck = 1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty
module_hotfixes=1
EOF

# Refresh YUM/DNF repository cache
sudo yum makecache;
```

{{< /tab >}}
{{< /tabs >}}

All the RPM / DEB packages are signed with [GPG Key](/repo/gpg) fingerprint (`B9BD8B20`) in Pigsty repository.

## Compatibility

Pigsty has two major repos: [INFRA](/repo/infra/) and [PGSQL](/repo/pgsql/),
provide DEB / RPM for `x86_64` and `aarch64` architecture.

The [`infra`](/repo/infra) repo contains packages that are generic to any PostgreSQL version and Linux major version,
including prometheus & grafana stack, admin tools for postgres, and many utils written in go.

| Linux  | Package |                 x86_64                  |                 aarch64                 |
|:------:|:-------:|:---------------------------------------:|:---------------------------------------:|
|   EL   |  `rpm`  | {{< badge content="✓" color="green" >}} | {{< badge content="✓" color="green" >}} |
| Debian |  `deb`  | {{< badge content="✓" color="green" >}} | {{< badge content="✓" color="green" >}} |

The [`pgsql`](/repo/pgsql) repo contains packages that are ad hoc to specific PostgreSQL Major Versions.
(Often ad hoc to a specific Linux distro major version, too) Including extensions, and some kernel forks.

| OS                                                                                      | Vendor | Major |  Minor  | Fullname          | PG Major Version                                                                                                                                                                                                                                            |                     Comment                     |
|:----------------------------------------------------------------------------------------|:-------|:-----:|:-------:|:------------------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-----------------------------------------------:|
| {{< badge content="el7.x86_64"  color="orange" >}}                                      | EL     |   7   |   7.9   | CentOS 7 x86      | {{< badge content="18" color="red" >}} {{< badge content="17" color="red" >}} {{< badge content="16" color="red" >}} {{< badge content="15" color="orange" >}} {{< badge content="14" color="orange" >}} {{< badge content="13" color="orange" >}}          |     {{< badge content="EOL" color="red" >}}     |
| {{< badge content="el8.x86_64"  link="/os/el8.x86_64"  color="green" >}}              | EL     |   8   |  8.10   | RockyLinux 8 x86  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       | {{< badge content="Near EOL" color="orange" >}} |
| {{< badge content="el8.aarch64" link="/os/el8.aarch64" color="green" border=false >}} | EL     |   8   |  8.10   | RockyLinux 8 ARM  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       | {{< badge content="Near EOL" color="orange" >}} |
| {{< badge content="el9.x86_64"  link="/os/el9.x86_64"  color="green" >}}              | EL     |   9   |   9.6   | RockyLinux 9 x86  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="el9.aarch64" link="/os/el9.aarch64" color="green" border=false >}} | EL     |   9   |   9.6   | RockyLinux 9 ARM  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="el10.x86_64" link="/os/el10.x86_64" color="green" >}}              | EL     |  10   |  10.0   | RockyLinux 10 x86 | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="el10.aarch64"link="/os/el10.aarch64"color="green" border=false >}} | EL     |  10   |  10.0   | RockyLinux 10 ARM | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="d11.x86_64"  color="orange" >}}                                      | Debian |  11   |  11.11  | Debian 11 x86     | {{< badge content="18" color="orange" >}} {{< badge content="17" color="orange" >}} {{< badge content="16" color="orange" >}} {{< badge content="15" color="orange" >}} {{< badge content="14" color="orange" >}} {{< badge content="13" color="orange" >}} |     {{< badge content="EOL" color="red" >}}     |
| {{< badge content="d11.aarch64" color="orange" border=false >}}                         | Debian |  11   |  11.11  | Debian 11 ARM     | {{< badge content="18" color="orange" >}} {{< badge content="17" color="orange" >}} {{< badge content="16" color="orange" >}} {{< badge content="15" color="orange" >}} {{< badge content="14" color="orange" >}} {{< badge content="13" color="orange" >}} |     {{< badge content="EOL" color="red" >}}     |
| {{< badge content="d12.x86_64"  link="/os/d12.x86_64"  color="green" >}}              | Debian |  12   |  12.12  | Debian 12 x86     | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="d12.aarch64" link="/os/d12.aarch64" color="green" border=false >}} | Debian |  12   |  12.12  | Debian 12 ARM     | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="d13.x86_64"  link="/os/d13.x86_64"  color="green" >}}              | Debian |  13   |  13.1   | Debian 13 x86     | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="d13.aarch64" link="/os/d13.aarch64" color="green" border=false >}} | Debian |  13   |  13.1   | Debian 13 ARM     | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="u20.x86_64" color="orange" >}}                                       | Ubuntu |  20   | 20.04.6 | Ubuntu 20.04 x86  | {{< badge content="18" color="orange" >}} {{< badge content="17" color="orange" >}} {{< badge content="16" color="orange" >}} {{< badge content="15" color="orange" >}} {{< badge content="14" color="orange" >}} {{< badge content="13" color="orange" >}} |     {{< badge content="EOL" color="red" >}}     |
| {{< badge content="u20.aarch64"color="orange" border=false >}}                          | Ubuntu |  20   | 20.04.6 | Ubuntu 20.04 ARM  | {{< badge content="18" color="orange" >}} {{< badge content="17" color="orange" >}} {{< badge content="16" color="orange" >}} {{< badge content="15" color="orange" >}} {{< badge content="14" color="orange" >}} {{< badge content="13" color="orange" >}} |     {{< badge content="EOL" color="red" >}}     |
| {{< badge content="u22.x86_64"  link="/os/u22.x86_64"  color="green" >}}              | Ubuntu |  22   | 22.04.5 | Ubuntu 22.04 x86  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="u22.aarch64" link="/os/u22.aarch64" color="green" border=false >}} | Ubuntu |  22   | 22.04.5 | Ubuntu 22.04 ARM  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="u24.x86_64"  link="/os/u24.x86_64"  color="green" >}}              | Ubuntu |  24   | 24.04.3 | Ubuntu 24.04 x86  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |
| {{< badge content="u24.aarch64" link="/os/u24.aarch64" color="green" border=false >}} | Ubuntu |  24   | 24.04.3 | Ubuntu 24.04 ARM  | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} {{< badge content="13" color="green" >}}       |    {{< badge content="OK" color="green" >}}     |


------

## Source

Building specs of these repos and packages are open-sourced on GitHub:

- https://github.com/pgsty/rpm
- https://github.com/pgsty/deb
- https://github.com/pgsty/infra-pkg
