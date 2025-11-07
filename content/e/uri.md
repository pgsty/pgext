---
title: "uri"
linkTitle: "uri"
description: "URI Data type for PostgreSQL"
weight: 3840
categories: ["TYPE"]
width: full
---

URI Data type for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3840** | {{< badge content="uri" link="https://github.com/petere/pguri" >}} | {{< ext "uri" "pg_uri" >}} | `1.20151224` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |

> [!Note] +int flag


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/uri" >}} | `1.20151224` | {{< bg "18" "pg_uri_18*" "green" >}} {{< bg "17" "pg_uri_17*" "green" >}} {{< bg "16" "pg_uri_16*" "green" >}} {{< bg "15" "pg_uri_15*" "green" >}} {{< bg "14" "pg_uri_14*" "green" >}} {{< bg "13" "pg_uri_13*" "green" >}} | `pg_uri_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/uri" >}} | `1.20151224` | {{< bg "18" "postgresql-18-pg-uri" "green" >}} {{< bg "17" "postgresql-17-pg-uri" "green" >}} {{< bg "16" "postgresql-16-pg-uri" "green" >}} {{< bg "15" "postgresql-15-pg-uri" "green" >}} {{< bg "14" "postgresql-14-pg-uri" "green" >}} {{< bg "13" "postgresql-13-pg-uri" "green" >}} | `postgresql-$v-pg-uri` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "pg_uri_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-18-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-17-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-16-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-15-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-14-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-13-pg-uri : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-18-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-17-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-16-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-15-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-14-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-13-pg-uri : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-18-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-17-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-16-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-15-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-14-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-13-pg-uri : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-uri : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-uri : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-uri : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-uri : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-uri : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-uri : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-18-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-17-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-16-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-15-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-14-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-13-pg-uri : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-18-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-17-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-16-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-15-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-14-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-13-pg-uri : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-18-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-17-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-16-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-15-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-14-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-13-pg-uri : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-18-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-17-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-16-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-15-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-14-pg-uri : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20151224" "postgresql-13-pg-uri : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uri_18` | 1.20151224 | `el8.x86_64` | pigsty | 19.2 KiB | [pg_uri_18-1.20151224-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_18-1.20151224-1PIGSTY.el8.x86_64.rpm) |
| `pg_uri_18` | 1.20151224 | `el8.aarch64` | pigsty | 18.9 KiB | [pg_uri_18-1.20151224-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_18-1.20151224-1PIGSTY.el8.aarch64.rpm) |
| `pg_uri_18` | 1.20151224 | `el9.x86_64` | pigsty | 18.9 KiB | [pg_uri_18-1.20151224-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_18-1.20151224-1PIGSTY.el9.x86_64.rpm) |
| `pg_uri_18` | 1.20151224 | `el9.aarch64` | pigsty | 18.5 KiB | [pg_uri_18-1.20151224-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_18-1.20151224-1PIGSTY.el9.aarch64.rpm) |
| `pg_uri_18` | 1.20151224 | `el10.x86_64` | pigsty | 19.0 KiB | [pg_uri_18-1.20151224-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uri_18-1.20151224-1PIGSTY.el10.x86_64.rpm) |
| `pg_uri_18` | 1.20151224 | `el10.aarch64` | pigsty | 18.7 KiB | [pg_uri_18-1.20151224-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uri_18-1.20151224-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-uri` | 1.20151224 | `d12.x86_64` | pigsty | 21.6 KiB | [postgresql-18-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-18-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-uri` | 1.20151224 | `d12.aarch64` | pigsty | 21.3 KiB | [postgresql-18-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-18-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-uri` | 1.20151224 | `d13.x86_64` | pigsty | 21.8 KiB | [postgresql-18-pg-uri_1.20151224-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uri/postgresql-18-pg-uri_1.20151224-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-uri` | 1.20151224 | `u22.x86_64` | pigsty | 22.8 KiB | [postgresql-18-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-18-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-uri` | 1.20151224 | `u22.aarch64` | pigsty | 22.4 KiB | [postgresql-18-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-18-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-uri` | 1.20151224 | `u24.x86_64` | pigsty | 22.7 KiB | [postgresql-18-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-18-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-uri` | 1.20151224 | `u24.aarch64` | pigsty | 22.4 KiB | [postgresql-18-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-18-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uri_17` | 1.20151224 | `el8.x86_64` | pigsty | 19.2 KiB | [pg_uri_17-1.20151224-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_17-1.20151224-1PIGSTY.el8.x86_64.rpm) |
| `pg_uri_17` | 1.20151224 | `el8.aarch64` | pigsty | 18.9 KiB | [pg_uri_17-1.20151224-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_17-1.20151224-1PIGSTY.el8.aarch64.rpm) |
| `pg_uri_17` | 1.20151224 | `el9.x86_64` | pigsty | 18.8 KiB | [pg_uri_17-1.20151224-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_17-1.20151224-1PIGSTY.el9.x86_64.rpm) |
| `pg_uri_17` | 1.20151224 | `el9.aarch64` | pigsty | 18.5 KiB | [pg_uri_17-1.20151224-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_17-1.20151224-1PIGSTY.el9.aarch64.rpm) |
| `pg_uri_17` | 1.20151224 | `el10.x86_64` | pigsty | 18.8 KiB | [pg_uri_17-1.20151224-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uri_17-1.20151224-1PIGSTY.el10.x86_64.rpm) |
| `pg_uri_17` | 1.20151224 | `el10.aarch64` | pigsty | 18.7 KiB | [pg_uri_17-1.20151224-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uri_17-1.20151224-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-uri` | 1.20151224 | `d12.x86_64` | pigsty | 21.5 KiB | [postgresql-17-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-uri` | 1.20151224 | `d12.aarch64` | pigsty | 21.4 KiB | [postgresql-17-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-uri` | 1.20151224 | `d13.x86_64` | pigsty | 21.8 KiB | [postgresql-17-pg-uri_1.20151224-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-uri` | 1.20151224 | `u22.x86_64` | pigsty | 23.1 KiB | [postgresql-17-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-uri` | 1.20151224 | `u22.aarch64` | pigsty | 23.0 KiB | [postgresql-17-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-uri` | 1.20151224 | `u24.x86_64` | pigsty | 22.6 KiB | [postgresql-17-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-uri` | 1.20151224 | `u24.aarch64` | pigsty | 22.4 KiB | [postgresql-17-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uri_16` | 1.20151224 | `el8.x86_64` | pigsty | 19.2 KiB | [pg_uri_16-1.20151224-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_16-1.20151224-1PIGSTY.el8.x86_64.rpm) |
| `pg_uri_16` | 1.20151224 | `el8.aarch64` | pigsty | 18.9 KiB | [pg_uri_16-1.20151224-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_16-1.20151224-1PIGSTY.el8.aarch64.rpm) |
| `pg_uri_16` | 1.20151224 | `el9.x86_64` | pigsty | 18.8 KiB | [pg_uri_16-1.20151224-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_16-1.20151224-1PIGSTY.el9.x86_64.rpm) |
| `pg_uri_16` | 1.20151224 | `el9.aarch64` | pigsty | 18.5 KiB | [pg_uri_16-1.20151224-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_16-1.20151224-1PIGSTY.el9.aarch64.rpm) |
| `pg_uri_16` | 1.20151224 | `el10.x86_64` | pigsty | 18.8 KiB | [pg_uri_16-1.20151224-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uri_16-1.20151224-1PIGSTY.el10.x86_64.rpm) |
| `pg_uri_16` | 1.20151224 | `el10.aarch64` | pigsty | 18.7 KiB | [pg_uri_16-1.20151224-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uri_16-1.20151224-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-uri` | 1.20151224 | `d12.x86_64` | pigsty | 21.5 KiB | [postgresql-16-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-uri` | 1.20151224 | `d12.aarch64` | pigsty | 21.4 KiB | [postgresql-16-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-uri` | 1.20151224 | `d13.x86_64` | pigsty | 21.7 KiB | [postgresql-16-pg-uri_1.20151224-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-uri` | 1.20151224 | `u22.x86_64` | pigsty | 23.1 KiB | [postgresql-16-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-uri` | 1.20151224 | `u22.aarch64` | pigsty | 23.0 KiB | [postgresql-16-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-uri` | 1.20151224 | `u24.x86_64` | pigsty | 22.6 KiB | [postgresql-16-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-uri` | 1.20151224 | `u24.aarch64` | pigsty | 22.4 KiB | [postgresql-16-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uri_15` | 1.20151224 | `el8.x86_64` | pigsty | 19.2 KiB | [pg_uri_15-1.20151224-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_15-1.20151224-1PIGSTY.el8.x86_64.rpm) |
| `pg_uri_15` | 1.20151224 | `el8.aarch64` | pigsty | 18.9 KiB | [pg_uri_15-1.20151224-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_15-1.20151224-1PIGSTY.el8.aarch64.rpm) |
| `pg_uri_15` | 1.20151224 | `el9.x86_64` | pigsty | 18.8 KiB | [pg_uri_15-1.20151224-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_15-1.20151224-1PIGSTY.el9.x86_64.rpm) |
| `pg_uri_15` | 1.20151224 | `el9.aarch64` | pigsty | 18.5 KiB | [pg_uri_15-1.20151224-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_15-1.20151224-1PIGSTY.el9.aarch64.rpm) |
| `pg_uri_15` | 1.20151224 | `el10.x86_64` | pigsty | 18.8 KiB | [pg_uri_15-1.20151224-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uri_15-1.20151224-1PIGSTY.el10.x86_64.rpm) |
| `pg_uri_15` | 1.20151224 | `el10.aarch64` | pigsty | 18.7 KiB | [pg_uri_15-1.20151224-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uri_15-1.20151224-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-uri` | 1.20151224 | `d12.x86_64` | pigsty | 21.8 KiB | [postgresql-15-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-uri` | 1.20151224 | `d12.aarch64` | pigsty | 21.2 KiB | [postgresql-15-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-uri` | 1.20151224 | `d13.x86_64` | pigsty | 21.8 KiB | [postgresql-15-pg-uri_1.20151224-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-uri` | 1.20151224 | `u22.x86_64` | pigsty | 23.1 KiB | [postgresql-15-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-uri` | 1.20151224 | `u22.aarch64` | pigsty | 23.0 KiB | [postgresql-15-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-uri` | 1.20151224 | `u24.x86_64` | pigsty | 22.6 KiB | [postgresql-15-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-uri` | 1.20151224 | `u24.aarch64` | pigsty | 22.2 KiB | [postgresql-15-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uri_14` | 1.20151224 | `el8.x86_64` | pigsty | 19.1 KiB | [pg_uri_14-1.20151224-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_14-1.20151224-1PIGSTY.el8.x86_64.rpm) |
| `pg_uri_14` | 1.20151224 | `el8.aarch64` | pigsty | 18.9 KiB | [pg_uri_14-1.20151224-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_14-1.20151224-1PIGSTY.el8.aarch64.rpm) |
| `pg_uri_14` | 1.20151224 | `el9.x86_64` | pigsty | 18.8 KiB | [pg_uri_14-1.20151224-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_14-1.20151224-1PIGSTY.el9.x86_64.rpm) |
| `pg_uri_14` | 1.20151224 | `el9.aarch64` | pigsty | 18.5 KiB | [pg_uri_14-1.20151224-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_14-1.20151224-1PIGSTY.el9.aarch64.rpm) |
| `pg_uri_14` | 1.20151224 | `el10.x86_64` | pigsty | 18.8 KiB | [pg_uri_14-1.20151224-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uri_14-1.20151224-1PIGSTY.el10.x86_64.rpm) |
| `pg_uri_14` | 1.20151224 | `el10.aarch64` | pigsty | 18.7 KiB | [pg_uri_14-1.20151224-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uri_14-1.20151224-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-uri` | 1.20151224 | `d12.x86_64` | pigsty | 21.8 KiB | [postgresql-14-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-uri` | 1.20151224 | `d12.aarch64` | pigsty | 21.2 KiB | [postgresql-14-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-uri` | 1.20151224 | `d13.x86_64` | pigsty | 21.7 KiB | [postgresql-14-pg-uri_1.20151224-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-uri` | 1.20151224 | `u22.x86_64` | pigsty | 23.1 KiB | [postgresql-14-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-uri` | 1.20151224 | `u22.aarch64` | pigsty | 22.9 KiB | [postgresql-14-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-uri` | 1.20151224 | `u24.x86_64` | pigsty | 22.6 KiB | [postgresql-14-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-uri` | 1.20151224 | `u24.aarch64` | pigsty | 22.2 KiB | [postgresql-14-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uri_13` | 1.20151224 | `el8.x86_64` | pigsty | 19.1 KiB | [pg_uri_13-1.20151224-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_13-1.20151224-1PIGSTY.el8.x86_64.rpm) |
| `pg_uri_13` | 1.20151224 | `el8.aarch64` | pigsty | 18.9 KiB | [pg_uri_13-1.20151224-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_13-1.20151224-1PIGSTY.el8.aarch64.rpm) |
| `pg_uri_13` | 1.20151224 | `el9.x86_64` | pigsty | 18.7 KiB | [pg_uri_13-1.20151224-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_13-1.20151224-1PIGSTY.el9.x86_64.rpm) |
| `pg_uri_13` | 1.20151224 | `el9.aarch64` | pigsty | 18.5 KiB | [pg_uri_13-1.20151224-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_13-1.20151224-1PIGSTY.el9.aarch64.rpm) |
| `pg_uri_13` | 1.20151224 | `el10.x86_64` | pigsty | 18.9 KiB | [pg_uri_13-1.20151224-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_uri_13-1.20151224-1PIGSTY.el10.x86_64.rpm) |
| `pg_uri_13` | 1.20151224 | `el10.aarch64` | pigsty | 18.7 KiB | [pg_uri_13-1.20151224-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_uri_13-1.20151224-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-uri` | 1.20151224 | `d12.x86_64` | pigsty | 21.5 KiB | [postgresql-13-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-13-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-uri` | 1.20151224 | `d12.aarch64` | pigsty | 21.1 KiB | [postgresql-13-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-13-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-uri` | 1.20151224 | `d13.x86_64` | pigsty | 21.4 KiB | [postgresql-13-pg-uri_1.20151224-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uri/postgresql-13-pg-uri_1.20151224-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-uri` | 1.20151224 | `u22.x86_64` | pigsty | 22.9 KiB | [postgresql-13-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-13-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-uri` | 1.20151224 | `u22.aarch64` | pigsty | 22.7 KiB | [postgresql-13-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-13-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-uri` | 1.20151224 | `u24.x86_64` | pigsty | 22.2 KiB | [postgresql-13-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-13-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-uri` | 1.20151224 | `u24.aarch64` | pigsty | 22.1 KiB | [postgresql-13-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-13-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petere/pguri" title="Repository" icon="github" subtitle="github.com/petere/pguri" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pguri-1.20151224.tar.gz" >}}
{{< /cards >}}


```bash
pig build get uri; # get uri source code
pig build dep uri; # install build dependencies
pig build pkg uri; # build extension rpm or deb
pig build ext uri; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install uri; # install by extension name, for the current active PG version
pig ext install pg_uri; # install via package alias, for the active PG version
pig ext install uri -v 18;   # install for PG 18
pig ext install uri -v 17;   # install for PG 17
pig ext install uri -v 16;   # install for PG 16
pig ext install uri -v 15;   # install for PG 15
pig ext install uri -v 14;   # install for PG 14
pig ext install uri -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION uri;
```

