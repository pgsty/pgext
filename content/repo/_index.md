---
title: Repository
description: The infrastructure to deliver PostgreSQL Extensions
icon: fa-solid fa-warehouse
weight: 300
breadcrumb: false
---


Pigsty provides a PostgreSQL extension repository across 16 [Linux platform targets](/os)
(8 distribution releases on both `x86_64` and `aarch64`).
It is designed to work together with the official PostgreSQL Global Development
Group ([PGDG](https://www.postgresql.org/download/linux/)) repo.
Together, they can provide up to [627 packaged PostgreSQL extensions](https://ext.pgsty.com/list) out-of-the-box.

{{< cards cols=2 >}}
{{< card link="/repo/pgsql"  title="PGSQL Repo" subtitle="Pigsty Extension Repository"  icon="play" />}}
{{< card link="/repo/infra"  title="INFRA Repo" subtitle="Pigsty Infrastructure Repo"   icon="sparkles" />}}
{{< /cards >}}

Current active targets are **PG15–19** across **16 Linux platforms** (EL8 covers PG15–18).

## Get Started

You can enable the pigsty [infra](/repo/infra) & [pgsql](/repo/pgsql) repo with the [pig](https://pig.pgsty.com) CLI tool:

{{< tabs group="mirror" >}}
{{< tab label="Default" value="global" >}}

```bash tab="pig"
curl https://repo.pigsty.io/pig | bash      # download and install the pig CLI tool
pig repo add all -u                         # add linux, pgdg, pigsty repo and update cache
```

{{< /tab >}}

{{< tab label="Mirror" value="mirror" >}}

```bash tab="pig"
curl https://repo.pigsty.cc/pig | bash      # download from mirror site
pig repo add -u                             # add linux, pgdg, pigsty repo and update cache
```

{{< /tab >}}
{{< /tabs >}}

## Manual Install

You can also add these repo to your system [manually](#manual-install) with default `apt`, `dnf`, `yum` approach.

{{< tabs group="pkgmgr" >}}
{{< tab label="APT" value="apt" >}}

```bash tab="apt"
# Add Pigsty's GPG public key to your system keychain to verify package signatures
curl -fsSL https://repo.pigsty.io/key | sudo gpg --dearmor -o /etc/apt/keyrings/pigsty.gpg

# Get the Debian/Ubuntu codename (bookworm, trixie, jammy, noble, resolute), and write the matching repository address to the APT list file
distro_codename=$(lsb_release -cs)
sudo tee /etc/apt/sources.list.d/pigsty-io.list > /dev/null <<EOF
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.io/apt/infra generic main
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.io/apt/pgsql/${distro_codename} ${distro_codename} main
EOF

# Refresh APT repository cache
sudo apt update
```

{{< /tab >}}

{{< tab label="DNF/YUM" value="dnf" >}}

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

{{< pgext-platforms >}}
{.wide .expand}


------

## Source

Building specs of these repos and packages are open-sourced on GitHub:

- https://github.com/pgsty/rpm
- https://github.com/pgsty/deb
- https://github.com/pgsty/infra-pkg
