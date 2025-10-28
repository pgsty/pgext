---
title: "shacrypt"
linkTitle: "shacrypt"
description: "Implements SHA256-CRYPT and SHA512-CRYPT password encryption schemes"
weight: 4440
categories: ["UTIL"]
width: full
---

Implements SHA256-CRYPT and SHA512-CRYPT password encryption schemes


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4440** | {{< badge content="shacrypt" link="https://github.com/dverite/postgres-shacrypt" >}} | {{< ext "shacrypt" >}} | `1.1` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hashlib" >}} {{< ext "xxhash" >}} {{< ext "cryptint" >}} {{< ext "pguecc" >}} {{< ext "pgcrypto" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/shacrypt" >}} | `1.1` | {{< bg "18" "shacrypt_18*" "green" >}} {{< bg "17" "shacrypt_17*" "green" >}} {{< bg "16" "shacrypt_16*" "green" >}} {{< bg "15" "shacrypt_15*" "green" >}} {{< bg "14" "shacrypt_14*" "green" >}} {{< bg "13" "shacrypt_13*" "green" >}} | `shacrypt_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/shacrypt" >}} | `1.1` | {{< bg "18" "postgresql-18-shacrypt" "green" >}} {{< bg "17" "postgresql-17-shacrypt" "green" >}} {{< bg "16" "postgresql-16-shacrypt" "green" >}} {{< bg "15" "postgresql-15-shacrypt" "green" >}} {{< bg "14" "postgresql-14-shacrypt" "green" >}} {{< bg "13" "postgresql-13-shacrypt" "green" >}} | `postgresql-$v-shacrypt` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "shacrypt_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_13 : MISS 0" "red" >}}      |
|    `el8.aarch64`    |      {{< bg "MISS" "shacrypt_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_13 : MISS 0" "red" >}}      |
|    `el9.x86_64`    |      {{< bg "MISS" "shacrypt_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_13 : MISS 0" "red" >}}      |
|    `el9.aarch64`    |      {{< bg "MISS" "shacrypt_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_13 : MISS 0" "red" >}}      |
|    `el10.x86_64`    |      {{< bg "MISS" "shacrypt_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "shacrypt_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "shacrypt_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-shacrypt : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1" "postgresql-17-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-13-shacrypt : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-shacrypt : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1" "postgresql-17-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-13-shacrypt : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-shacrypt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-shacrypt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-shacrypt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-shacrypt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-shacrypt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-shacrypt : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-shacrypt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-shacrypt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-shacrypt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-shacrypt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-shacrypt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-shacrypt : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-shacrypt : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1" "postgresql-17-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-13-shacrypt : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-shacrypt : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1" "postgresql-17-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-13-shacrypt : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-shacrypt : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1" "postgresql-17-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-13-shacrypt : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-shacrypt : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.1" "postgresql-17-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-shacrypt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-13-shacrypt : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-17-shacrypt` | 1.1 | `d12.x86_64` | pigsty | 2.5 KiB | [postgresql-17-shacrypt_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgres-shacrypt/postgresql-17-shacrypt_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-shacrypt` | 1.1 | `d12.aarch64` | pigsty | 2.5 KiB | [postgresql-17-shacrypt_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgres-shacrypt/postgresql-17-shacrypt_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-shacrypt` | 1.1 | `u22.x86_64` | pigsty | 2.4 KiB | [postgresql-17-shacrypt_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgres-shacrypt/postgresql-17-shacrypt_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-shacrypt` | 1.1 | `u22.aarch64` | pigsty | 2.4 KiB | [postgresql-17-shacrypt_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgres-shacrypt/postgresql-17-shacrypt_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-shacrypt` | 1.1 | `u24.x86_64` | pigsty | 2.4 KiB | [postgresql-17-shacrypt_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgres-shacrypt/postgresql-17-shacrypt_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-shacrypt` | 1.1 | `u24.aarch64` | pigsty | 2.4 KiB | [postgresql-17-shacrypt_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgres-shacrypt/postgresql-17-shacrypt_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-16-shacrypt` | 1.1 | `d12.x86_64` | pigsty | 2.5 KiB | [postgresql-16-shacrypt_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgres-shacrypt/postgresql-16-shacrypt_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-shacrypt` | 1.1 | `d12.aarch64` | pigsty | 2.5 KiB | [postgresql-16-shacrypt_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgres-shacrypt/postgresql-16-shacrypt_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-shacrypt` | 1.1 | `u22.x86_64` | pigsty | 2.4 KiB | [postgresql-16-shacrypt_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgres-shacrypt/postgresql-16-shacrypt_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-shacrypt` | 1.1 | `u22.aarch64` | pigsty | 2.4 KiB | [postgresql-16-shacrypt_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgres-shacrypt/postgresql-16-shacrypt_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-shacrypt` | 1.1 | `u24.x86_64` | pigsty | 2.4 KiB | [postgresql-16-shacrypt_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgres-shacrypt/postgresql-16-shacrypt_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-shacrypt` | 1.1 | `u24.aarch64` | pigsty | 2.4 KiB | [postgresql-16-shacrypt_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgres-shacrypt/postgresql-16-shacrypt_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-15-shacrypt` | 1.1 | `d12.x86_64` | pigsty | 2.5 KiB | [postgresql-15-shacrypt_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgres-shacrypt/postgresql-15-shacrypt_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-shacrypt` | 1.1 | `d12.aarch64` | pigsty | 2.5 KiB | [postgresql-15-shacrypt_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgres-shacrypt/postgresql-15-shacrypt_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-shacrypt` | 1.1 | `u22.x86_64` | pigsty | 2.4 KiB | [postgresql-15-shacrypt_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgres-shacrypt/postgresql-15-shacrypt_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-shacrypt` | 1.1 | `u22.aarch64` | pigsty | 2.4 KiB | [postgresql-15-shacrypt_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgres-shacrypt/postgresql-15-shacrypt_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-shacrypt` | 1.1 | `u24.x86_64` | pigsty | 2.4 KiB | [postgresql-15-shacrypt_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgres-shacrypt/postgresql-15-shacrypt_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-shacrypt` | 1.1 | `u24.aarch64` | pigsty | 2.4 KiB | [postgresql-15-shacrypt_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgres-shacrypt/postgresql-15-shacrypt_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-14-shacrypt` | 1.1 | `d12.x86_64` | pigsty | 2.5 KiB | [postgresql-14-shacrypt_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgres-shacrypt/postgresql-14-shacrypt_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-shacrypt` | 1.1 | `d12.aarch64` | pigsty | 2.5 KiB | [postgresql-14-shacrypt_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgres-shacrypt/postgresql-14-shacrypt_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-shacrypt` | 1.1 | `u22.x86_64` | pigsty | 2.4 KiB | [postgresql-14-shacrypt_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgres-shacrypt/postgresql-14-shacrypt_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-shacrypt` | 1.1 | `u22.aarch64` | pigsty | 2.4 KiB | [postgresql-14-shacrypt_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgres-shacrypt/postgresql-14-shacrypt_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-shacrypt` | 1.1 | `u24.x86_64` | pigsty | 2.4 KiB | [postgresql-14-shacrypt_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgres-shacrypt/postgresql-14-shacrypt_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-shacrypt` | 1.1 | `u24.aarch64` | pigsty | 2.4 KiB | [postgresql-14-shacrypt_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgres-shacrypt/postgresql-14-shacrypt_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-13-shacrypt` | 1.1 | `d12.x86_64` | pigsty | 2.5 KiB | [postgresql-13-shacrypt_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgres-shacrypt/postgresql-13-shacrypt_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-shacrypt` | 1.1 | `d12.aarch64` | pigsty | 2.5 KiB | [postgresql-13-shacrypt_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/postgres-shacrypt/postgresql-13-shacrypt_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-shacrypt` | 1.1 | `u22.x86_64` | pigsty | 2.4 KiB | [postgresql-13-shacrypt_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgres-shacrypt/postgresql-13-shacrypt_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-shacrypt` | 1.1 | `u22.aarch64` | pigsty | 2.4 KiB | [postgresql-13-shacrypt_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/postgres-shacrypt/postgresql-13-shacrypt_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-shacrypt` | 1.1 | `u24.x86_64` | pigsty | 2.4 KiB | [postgresql-13-shacrypt_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgres-shacrypt/postgresql-13-shacrypt_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-shacrypt` | 1.1 | `u24.aarch64` | pigsty | 2.4 KiB | [postgresql-13-shacrypt_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/postgres-shacrypt/postgresql-13-shacrypt_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dverite/postgres-shacrypt" title="Repository" icon="github" subtitle="github.com/dverite/postgres-shacrypt" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgres_shacrypt-1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get shacrypt; # get shacrypt source code
pig build dep shacrypt; # install build dependencies
pig build pkg shacrypt; # build extension rpm or deb
pig build ext shacrypt; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install shacrypt; # install by extension name, for the current active PG version
pig ext install shacrypt; # install via package alias, for the active PG version
pig ext install shacrypt -v 18;   # install for PG 18
pig ext install shacrypt -v 17;   # install for PG 17
pig ext install shacrypt -v 16;   # install for PG 16
pig ext install shacrypt -v 15;   # install for PG 15
pig ext install shacrypt -v 14;   # install for PG 14
pig ext install shacrypt -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION shacrypt;
```

