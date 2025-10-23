---
title: "pgqr"
linkTitle: "pgqr"
description: "QR Code generator from PostgreSQL"
weight: 4250
categories: ["UTIL"]
width: full
---

QR Code generator from PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4250** | {{< badge content="pgqr" link="https://github.com/AbdulYadi/pgqr" >}} | {{< ext "pgqr" >}} | `1.0` | {{< category "UTIL" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_protobuf" >}} {{< ext "base36" >}} {{< ext "base62" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgqr" >}} | `1.0` | {{< bg "18" "pgqr_18*" "red" >}} {{< bg "17" "pgqr_17*" "green" >}} {{< bg "16" "pgqr_16*" "green" >}} {{< bg "15" "pgqr_15*" "green" >}} {{< bg "14" "pgqr_14*" "green" >}} | `pgqr_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgqr" >}} | `1.0` | {{< bg "18" "postgresql-18-pgqr" "red" >}} {{< bg "17" "postgresql-17-pgqr" "green" >}} {{< bg "16" "postgresql-16-pgqr" "green" >}} {{< bg "15" "postgresql-15-pgqr" "green" >}} {{< bg "14" "postgresql-14-pgqr" "green" >}} | `postgresql-$v-pgqr` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pgqr_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "pgqr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_14 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "pgqr_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "pgqr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "pgqr_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "pgqr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "pgqr_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "pgqr_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pgqr_14 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pgqr : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pgqr : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pgqr : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pgqr : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pgqr : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pgqr : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0" "postgresql-17-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pgqr : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pgqr : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgqr_17` | 1.0 | `el8.x86_64` | pigsty | 23.9 KiB | [pgqr_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgqr_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgqr_17` | 1.0 | `el8.aarch64` | pigsty | 23.0 KiB | [pgqr_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgqr_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgqr_17` | 1.0 | `el9.x86_64` | pigsty | 23.7 KiB | [pgqr_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgqr_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgqr_17` | 1.0 | `el9.aarch64` | pigsty | 23.1 KiB | [pgqr_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgqr_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pgqr` | 1.0 | `d12.x86_64` | pigsty | 48.9 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgqr` | 1.0 | `d12.aarch64` | pigsty | 47.3 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgqr` | 1.0 | `u22.x86_64` | pigsty | 50.9 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgqr` | 1.0 | `u22.aarch64` | pigsty | 49.6 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgqr` | 1.0 | `u24.x86_64` | pigsty | 50.4 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgqr` | 1.0 | `u24.aarch64` | pigsty | 49.1 KiB | [postgresql-17-pgqr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-17-pgqr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgqr_16` | 1.0 | `el8.x86_64` | pigsty | 23.8 KiB | [pgqr_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgqr_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgqr_16` | 1.0 | `el8.aarch64` | pigsty | 23.0 KiB | [pgqr_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgqr_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgqr_16` | 1.0 | `el9.x86_64` | pigsty | 23.8 KiB | [pgqr_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgqr_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgqr_16` | 1.0 | `el9.aarch64` | pigsty | 23.3 KiB | [pgqr_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgqr_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pgqr` | 1.0 | `d12.x86_64` | pigsty | 48.9 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgqr` | 1.0 | `d12.aarch64` | pigsty | 47.3 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgqr` | 1.0 | `u22.x86_64` | pigsty | 50.9 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgqr` | 1.0 | `u22.aarch64` | pigsty | 49.6 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgqr` | 1.0 | `u24.x86_64` | pigsty | 50.4 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgqr` | 1.0 | `u24.aarch64` | pigsty | 49.1 KiB | [postgresql-16-pgqr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-16-pgqr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgqr_15` | 1.0 | `el8.x86_64` | pigsty | 24.4 KiB | [pgqr_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgqr_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgqr_15` | 1.0 | `el8.aarch64` | pigsty | 23.6 KiB | [pgqr_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgqr_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgqr_15` | 1.0 | `el9.x86_64` | pigsty | 24.8 KiB | [pgqr_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgqr_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgqr_15` | 1.0 | `el9.aarch64` | pigsty | 24.2 KiB | [pgqr_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgqr_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pgqr` | 1.0 | `d12.x86_64` | pigsty | 48.7 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgqr` | 1.0 | `d12.aarch64` | pigsty | 47.3 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgqr` | 1.0 | `u22.x86_64` | pigsty | 51.3 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgqr` | 1.0 | `u22.aarch64` | pigsty | 50.1 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgqr` | 1.0 | `u24.x86_64` | pigsty | 50.9 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgqr` | 1.0 | `u24.aarch64` | pigsty | 49.7 KiB | [postgresql-15-pgqr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-15-pgqr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgqr_14` | 1.0 | `el8.x86_64` | pigsty | 24.3 KiB | [pgqr_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgqr_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgqr_14` | 1.0 | `el8.aarch64` | pigsty | 23.6 KiB | [pgqr_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgqr_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgqr_14` | 1.0 | `el9.x86_64` | pigsty | 24.8 KiB | [pgqr_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgqr_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgqr_14` | 1.0 | `el9.aarch64` | pigsty | 24.2 KiB | [pgqr_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgqr_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pgqr` | 1.0 | `d12.x86_64` | pigsty | 48.7 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgqr` | 1.0 | `d12.aarch64` | pigsty | 47.2 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgqr` | 1.0 | `u22.x86_64` | pigsty | 51.3 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgqr` | 1.0 | `u22.aarch64` | pigsty | 50.1 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgqr` | 1.0 | `u24.x86_64` | pigsty | 50.8 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgqr` | 1.0 | `u24.aarch64` | pigsty | 49.6 KiB | [postgresql-14-pgqr_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgqr/postgresql-14-pgqr_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/AbdulYadi/pgqr" title="Repository" icon="github" subtitle="github.com/AbdulYadi/pgqr" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgqr-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgqr; # get pgqr source code
pig build dep pgqr; # install build dependencies
pig build pkg pgqr; # build extension rpm or deb
pig build ext pgqr; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgqr; # install by extension name, for the current active PG version
pig ext install pgqr; # install via package alias, for the active PG version
pig ext install pgqr -v 18;   # install for PG 18
pig ext install pgqr -v 17;   # install for PG 17
pig ext install pgqr -v 16;   # install for PG 16
pig ext install pgqr -v 15;   # install for PG 15
pig ext install pgqr -v 14;   # install for PG 14
pig ext install pgqr -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgqr;
```

