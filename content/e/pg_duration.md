---
title: "pg_duration"
linkTitle: "pg_duration"
description: "data type for representing durations"
weight: 3830
categories: ["TYPE"]
width: full
---

[**pg_duration**](https://github.com/jkosh44/pg_duration) : data type for representing durations


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3830** | {{< badge content="pg_duration" link="https://github.com/jkosh44/pg_duration" >}} | {{< ext "pg_duration" >}} | `1.0.2` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `pg_duration` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.2` | {{< bg "18" "pg_duration_18*" "green" >}} {{< bg "17" "pg_duration_17*" "green" >}} {{< bg "16" "pg_duration_16*" "red" >}} {{< bg "15" "pg_duration_15*" "red" >}} {{< bg "14" "pg_duration_14*" "red" >}} {{< bg "13" "pg_duration_13*" "red" >}} | `pg_duration_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.2` | {{< bg "18" "postgresql-18-pg-duration" "green" >}} {{< bg "17" "postgresql-17-pg-duration" "green" >}} {{< bg "16" "postgresql-16-pg-duration" "red" >}} {{< bg "15" "postgresql-15-pg-duration" "red" >}} {{< bg "14" "postgresql-14-pg-duration" "red" >}} {{< bg "13" "postgresql-13-pg-duration" "red" >}} | `postgresql-$v-pg-duration` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duration_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duration_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duration_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_17 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pg_duration_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duration_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duration_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-duration : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-duration : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-duration : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-duration : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-duration : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-duration : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-duration : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-duration : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-duration : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-duration : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-duration : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-duration : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-duration : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-duration : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pg-duration : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-duration : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duration_18` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.0 KiB | [pg_duration_18-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duration_18-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_duration_18` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.3 KiB | [pg_duration_18-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duration_18-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_duration_18` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.3 KiB | [pg_duration_18-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duration_18-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_duration_18` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.8 KiB | [pg_duration_18-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duration_18-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_duration_18` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 23.3 KiB | [pg_duration_18-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duration_18-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_duration_18` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.9 KiB | [pg_duration_18-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duration_18-1.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-duration` | `1.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.3 KiB | [postgresql-18-pg-duration_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duration/postgresql-18-pg-duration_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-duration` | `1.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.8 KiB | [postgresql-18-pg-duration_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duration/postgresql-18-pg-duration_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-duration` | `1.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 29.3 KiB | [postgresql-18-pg-duration_1.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duration/postgresql-18-pg-duration_1.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-duration` | `1.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.9 KiB | [postgresql-18-pg-duration_1.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duration/postgresql-18-pg-duration_1.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-duration` | `1.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.5 KiB | [postgresql-18-pg-duration_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duration/postgresql-18-pg-duration_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-duration` | `1.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 30.7 KiB | [postgresql-18-pg-duration_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duration/postgresql-18-pg-duration_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-duration` | `1.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.5 KiB | [postgresql-18-pg-duration_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duration/postgresql-18-pg-duration_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-duration` | `1.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.3 KiB | [postgresql-18-pg-duration_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duration/postgresql-18-pg-duration_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duration_17` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.0 KiB | [pg_duration_17-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duration_17-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_duration_17` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.3 KiB | [pg_duration_17-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duration_17-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_duration_17` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.5 KiB | [pg_duration_17-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duration_17-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_duration_17` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.8 KiB | [pg_duration_17-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duration_17-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_duration_17` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.3 KiB | [pg_duration_17-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duration_17-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_duration_17` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 23.4 KiB | [pg_duration_17-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duration_17-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_duration_17` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.9 KiB | [pg_duration_17-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duration_17-1.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-duration` | `1.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.2 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-duration` | `1.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.8 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-duration` | `1.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 29.2 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-duration` | `1.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 28.9 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-duration` | `1.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.3 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-duration` | `1.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 31.5 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-duration` | `1.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.5 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-duration` | `1.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.3 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jkosh44/pg_duration" title="Repository" icon="github" subtitle="github.com/jkosh44/pg_duration" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_duration-1.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_duration;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_duration;		# install via package name, for the active PG version

pig install pg_duration -v 18;   # install for PG 18
pig install pg_duration -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_duration;
```
