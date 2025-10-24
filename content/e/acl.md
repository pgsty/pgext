---
title: "acl"
linkTitle: "acl"
description: "ACL Data type"
weight: 3860
categories: ["TYPE"]
width: full
---

ACL Data type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3860** | {{< badge content="acl" link="https://github.com/arkhipov/acl" >}} | {{< ext "acl" "pg_acl" >}} | `1.0.4` | {{< category "TYPE" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |

> [!Note] +cast pg_uuid_t


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/acl" >}} | `1.0.4` | {{< bg "18" "acl_18*" "red" >}} {{< bg "17" "acl_17*" "green" >}} {{< bg "16" "acl_16*" "green" >}} {{< bg "15" "acl_15*" "green" >}} {{< bg "14" "acl_14*" "green" >}} | `acl_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/acl" >}} | `1.0.4` | {{< bg "18" "postgresql-18-acl" "red" >}} {{< bg "17" "postgresql-17-acl" "green" >}} {{< bg "16" "postgresql-16-acl" "green" >}} {{< bg "15" "postgresql-15-acl" "green" >}} {{< bg "14" "postgresql-14-acl" "green" >}} | `postgresql-$v-acl` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "acl_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.4" "acl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_14 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "acl_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.4" "acl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "acl_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.4" "acl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "acl_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.4" "acl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "acl_14 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-acl : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-acl : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-acl : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-acl : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-acl : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-acl : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.4" "postgresql-17-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-16-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-15-acl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.4" "postgresql-14-acl : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acl_17` | 1.0.4 | `el8.x86_64` | pigsty | 26.7 KiB | [acl_17-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acl_17-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `acl_17` | 1.0.4 | `el8.aarch64` | pigsty | 25.9 KiB | [acl_17-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acl_17-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `acl_17` | 1.0.4 | `el9.x86_64` | pigsty | 27.1 KiB | [acl_17-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acl_17-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `acl_17` | 1.0.4 | `el9.aarch64` | pigsty | 26.7 KiB | [acl_17-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acl_17-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-acl` | 1.0.4 | `d12.x86_64` | pigsty | 47.8 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-acl` | 1.0.4 | `d12.aarch64` | pigsty | 47.3 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-acl` | 1.0.4 | `u22.x86_64` | pigsty | 50.2 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-acl` | 1.0.4 | `u22.aarch64` | pigsty | 49.9 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-acl` | 1.0.4 | `u24.x86_64` | pigsty | 47.6 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-acl` | 1.0.4 | `u24.aarch64` | pigsty | 46.9 KiB | [postgresql-17-acl_1.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-17-acl_1.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acl_16` | 1.0.4 | `el8.x86_64` | pigsty | 26.7 KiB | [acl_16-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acl_16-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `acl_16` | 1.0.4 | `el8.aarch64` | pigsty | 25.9 KiB | [acl_16-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acl_16-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `acl_16` | 1.0.4 | `el9.x86_64` | pigsty | 27.1 KiB | [acl_16-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acl_16-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `acl_16` | 1.0.4 | `el9.aarch64` | pigsty | 26.7 KiB | [acl_16-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acl_16-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-acl` | 1.0.4 | `d12.x86_64` | pigsty | 47.8 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-acl` | 1.0.4 | `d12.aarch64` | pigsty | 47.3 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-acl` | 1.0.4 | `u22.x86_64` | pigsty | 50.2 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-acl` | 1.0.4 | `u22.aarch64` | pigsty | 49.9 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-acl` | 1.0.4 | `u24.x86_64` | pigsty | 47.7 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-acl` | 1.0.4 | `u24.aarch64` | pigsty | 46.8 KiB | [postgresql-16-acl_1.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-16-acl_1.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acl_15` | 1.0.4 | `el8.x86_64` | pigsty | 26.8 KiB | [acl_15-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acl_15-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `acl_15` | 1.0.4 | `el8.aarch64` | pigsty | 26.0 KiB | [acl_15-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acl_15-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `acl_15` | 1.0.4 | `el9.x86_64` | pigsty | 27.2 KiB | [acl_15-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acl_15-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `acl_15` | 1.0.4 | `el9.aarch64` | pigsty | 26.5 KiB | [acl_15-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acl_15-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-acl` | 1.0.4 | `d12.x86_64` | pigsty | 48.0 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-acl` | 1.0.4 | `d12.aarch64` | pigsty | 47.6 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-acl` | 1.0.4 | `u22.x86_64` | pigsty | 50.3 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-acl` | 1.0.4 | `u22.aarch64` | pigsty | 49.8 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-acl` | 1.0.4 | `u24.x86_64` | pigsty | 47.4 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-acl` | 1.0.4 | `u24.aarch64` | pigsty | 46.7 KiB | [postgresql-15-acl_1.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-15-acl_1.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `acl_14` | 1.0.4 | `el8.x86_64` | pigsty | 26.8 KiB | [acl_14-1.0.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/acl_14-1.0.4-1PIGSTY.el8.x86_64.rpm) |
| `acl_14` | 1.0.4 | `el8.aarch64` | pigsty | 26.0 KiB | [acl_14-1.0.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/acl_14-1.0.4-1PIGSTY.el8.aarch64.rpm) |
| `acl_14` | 1.0.4 | `el9.x86_64` | pigsty | 27.2 KiB | [acl_14-1.0.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/acl_14-1.0.4-1PIGSTY.el9.x86_64.rpm) |
| `acl_14` | 1.0.4 | `el9.aarch64` | pigsty | 26.5 KiB | [acl_14-1.0.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/acl_14-1.0.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-acl` | 1.0.4 | `d12.x86_64` | pigsty | 48.0 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-acl` | 1.0.4 | `d12.aarch64` | pigsty | 47.5 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-acl` | 1.0.4 | `u22.x86_64` | pigsty | 50.2 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-acl` | 1.0.4 | `u22.aarch64` | pigsty | 49.7 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-acl` | 1.0.4 | `u24.x86_64` | pigsty | 47.4 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-acl` | 1.0.4 | `u24.aarch64` | pigsty | 46.7 KiB | [postgresql-14-acl_1.0.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/acl/postgresql-14-acl_1.0.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/arkhipov/acl" title="Repository" icon="github" subtitle="github.com/arkhipov/acl" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="acl-1.0.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build get acl; # get acl source code
pig build dep acl; # install build dependencies
pig build pkg acl; # build extension rpm or deb
pig build ext acl; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install acl; # install by extension name, for the current active PG version
pig ext install pg_acl; # install via package alias, for the active PG version
pig ext install acl -v 18;   # install for PG 18
pig ext install acl -v 17;   # install for PG 17
pig ext install acl -v 16;   # install for PG 16
pig ext install acl -v 15;   # install for PG 15
pig ext install acl -v 14;   # install for PG 14
pig ext install acl -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION acl;
```

