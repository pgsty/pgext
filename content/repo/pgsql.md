---
title: PGSQL Repo
description: The repo for PostgreSQL Extensions & Kernel Forks
icon: fa-solid fa-database
weight: 100
---

The `pigsty-pgsql` repo contains packages that are ad hoc to specific PostgreSQL Major Versions.
(Often ad hoc to a specific Linux distro major version, too) Including extensions, and some kernel forks.

You can check the [Release - RPM Changelog](/release/rpm) / [Release - DEB Changelog](/release/deb)  for the latest updates.

{{< pgext-platforms >}}



The catalog covers PostgreSQL 15–19, with PostgreSQL 18 remaining the configuration default. PG14 was retired from active build targets on 2026-11-21; EL8 remains supported for PostgreSQL 15–18 (`PG19` is `N/A` on EL8), while EL9/EL10, Debian, and Ubuntu cover PostgreSQL 15–19. See the [package matrix](/os/matrix/) for per-target availability.

## Quick Start

### PIG

You can install [pig](https://pig.pgsty.com) - the cli tool, and add pgdg / pigsty repo with it (recommended):

```bash tab="pig"
pig repo add pigsty                         # add pigsty-pgsql repo
pig repo add pigsty -u                      # add pigsty-pgsql repo, and update cache
pig repo add pigsty -u --region=default     # add pigsty-pgsql repo and enforce default region (pigsty.io)
pig repo add pigsty -u --region=china       # add pigsty-pgsql repo with china region   (pigsty.cc)
pig repo add pgsql -u                       # pgsql = pgdg + pigsty-pgsql (add pigsty + official PGDG)
pig repo add -u                             # all = node + pgsql (pgdg + pigsty) + infra
```

### APT

You can also enable this repo with `apt` directly on Debian / Ubuntu:

```bash tab="apt"
# Add Pigsty's GPG public key to your system keychain to verify package signatures
curl -fsSL https://repo.pigsty.io/key | sudo gpg --dearmor -o /etc/apt/keyrings/pigsty.gpg

# Get Debian distribution codename (distro_codename=jammy, focal, bullseye, bookworm), and write the corresponding upstream repository address to the APT List file
distro_codename=$(lsb_release -cs)
sudo tee /etc/apt/sources.list.d/pigsty-io.list > /dev/null <<EOF
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.io/apt/pgsql/${distro_codename} ${distro_codename} main
EOF

# Refresh APT repository cache
sudo apt update
```

### DNF

You can also enable this repo with `dnf`/`yum` directly on EL-compatible systems:

```bash tab="dnf"
# Add Pigsty's GPG public key to your system keychain to verify package signatures
curl -fsSL https://repo.pigsty.io/key | sudo tee /etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty >/dev/null

# Add Pigsty Repo definition files to /etc/yum.repos.d/ directory, including two repositories
sudo tee /etc/yum.repos.d/pigsty-pgsql.repo > /dev/null <<-'EOF'
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
sudo dnf makecache;
```



--------

## Source

Building specs of this repo is open-sourced on GitHub:

- https://github.com/pgsty/rpm
- https://github.com/pgsty/deb

If the platform is not supported, you can also build the packages from source code by yourself.

The [`pig build`](https://pig.pgsty.com/build/) provides an easy way to build extension RPM/DEB packages on your own systems.
