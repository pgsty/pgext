---
title: "url_encode"
linkTitle: "url_encode"
description: "url_encode, url_decode functions"
weight: 4190
categories: ["UTIL"]
width: full
---

[**url_encode**](https://github.com/okbob/url_encode) : url_encode, url_decode functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4190** | {{< badge content="url_encode" link="https://github.com/okbob/url_encode" >}} | {{< ext "url_encode" >}} | `1.2.5` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_html5_email_address" >}} {{< ext "base36" >}} {{< ext "base62" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `url_encode` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.5` | {{< bg "18" "url_encode_18*" "green" >}} {{< bg "17" "url_encode_17*" "green" >}} {{< bg "16" "url_encode_16*" "green" >}} {{< bg "15" "url_encode_15*" "green" >}} {{< bg "14" "url_encode_14*" "green" >}} {{< bg "13" "url_encode_13*" "green" >}} | `url_encode_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.5` | {{< bg "18" "postgresql-18-url-encode" "green" >}} {{< bg "17" "postgresql-17-url-encode" "green" >}} {{< bg "16" "postgresql-16-url-encode" "green" >}} {{< bg "15" "postgresql-15-url-encode" "green" >}} {{< bg "14" "postgresql-14-url-encode" "green" >}} {{< bg "13" "postgresql-13-url-encode" "green" >}} | `postgresql-$v-url-encode` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-18-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `url_encode_18` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.3 KiB | [url_encode_18-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/url_encode_18-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `url_encode_18` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.4 KiB | [url_encode_18-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/url_encode_18-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `url_encode_18` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.2 KiB | [url_encode_18-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/url_encode_18-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `url_encode_18` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.2 KiB | [url_encode_18-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/url_encode_18-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `url_encode_18` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.2 KiB | [url_encode_18-1.2.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/url_encode_18-1.2.5-1PIGSTY.el10.x86_64.rpm) |
| `url_encode_18` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.5 KiB | [url_encode_18-1.2.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/url_encode_18-1.2.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-url-encode` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-18-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-18-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-url-encode` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.8 KiB | [postgresql-18-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-18-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-url-encode` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.8 KiB | [postgresql-18-url-encode_1.2.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/url-encode/postgresql-18-url-encode_1.2.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-url-encode` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.0 KiB | [postgresql-18-url-encode_1.2.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/url-encode/postgresql-18-url-encode_1.2.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-url-encode` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.0 KiB | [postgresql-18-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-18-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-url-encode` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.8 KiB | [postgresql-18-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-18-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-url-encode` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.2 KiB | [postgresql-18-url-encode_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-18-url-encode_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-url-encode` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.2 KiB | [postgresql-18-url-encode_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-18-url-encode_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `url_encode_17` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.3 KiB | [url_encode_17-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/url_encode_17-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `url_encode_17` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.4 KiB | [url_encode_17-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/url_encode_17-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `url_encode_17` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.2 KiB | [url_encode_17-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/url_encode_17-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `url_encode_17` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.2 KiB | [url_encode_17-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/url_encode_17-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `url_encode_17` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.2 KiB | [url_encode_17-1.2.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/url_encode_17-1.2.5-1PIGSTY.el10.x86_64.rpm) |
| `url_encode_17` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.5 KiB | [url_encode_17-1.2.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/url_encode_17-1.2.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-url-encode` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.7 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-url-encode` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.8 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-url-encode` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.7 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-url-encode` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.0 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-url-encode` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.1 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-url-encode` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.9 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-url-encode` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.2 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-url-encode` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.2 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `url_encode_16` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.2 KiB | [url_encode_16-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/url_encode_16-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `url_encode_16` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.4 KiB | [url_encode_16-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/url_encode_16-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `url_encode_16` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.0 KiB | [url_encode_16-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/url_encode_16-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `url_encode_16` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.1 KiB | [url_encode_16-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/url_encode_16-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `url_encode_16` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.1 KiB | [url_encode_16-1.2.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/url_encode_16-1.2.5-1PIGSTY.el10.x86_64.rpm) |
| `url_encode_16` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.4 KiB | [url_encode_16-1.2.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/url_encode_16-1.2.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-url-encode` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.4 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-url-encode` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.5 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-url-encode` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.4 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-url-encode` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.7 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-url-encode` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.9 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-url-encode` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.7 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-url-encode` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.8 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-url-encode` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.9 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `url_encode_15` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.2 KiB | [url_encode_15-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/url_encode_15-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `url_encode_15` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.4 KiB | [url_encode_15-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/url_encode_15-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `url_encode_15` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.0 KiB | [url_encode_15-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/url_encode_15-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `url_encode_15` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.1 KiB | [url_encode_15-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/url_encode_15-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `url_encode_15` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.1 KiB | [url_encode_15-1.2.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/url_encode_15-1.2.5-1PIGSTY.el10.x86_64.rpm) |
| `url_encode_15` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.4 KiB | [url_encode_15-1.2.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/url_encode_15-1.2.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-url-encode` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.5 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-url-encode` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.5 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-url-encode` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.4 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-url-encode` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.7 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-url-encode` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.9 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-url-encode` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.7 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-url-encode` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.9 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-url-encode` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.9 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `url_encode_14` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.2 KiB | [url_encode_14-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/url_encode_14-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `url_encode_14` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.4 KiB | [url_encode_14-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/url_encode_14-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `url_encode_14` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.0 KiB | [url_encode_14-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/url_encode_14-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `url_encode_14` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.1 KiB | [url_encode_14-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/url_encode_14-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `url_encode_14` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.1 KiB | [url_encode_14-1.2.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/url_encode_14-1.2.5-1PIGSTY.el10.x86_64.rpm) |
| `url_encode_14` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.4 KiB | [url_encode_14-1.2.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/url_encode_14-1.2.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-url-encode` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.4 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-url-encode` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.5 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-url-encode` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.4 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-url-encode` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.7 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-url-encode` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.9 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-url-encode` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.7 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-url-encode` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.9 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-url-encode` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.8 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `url_encode_13` | `1.2.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.0 KiB | [url_encode_13-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/url_encode_13-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `url_encode_13` | `1.2.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.3 KiB | [url_encode_13-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/url_encode_13-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `url_encode_13` | `1.2.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.2 KiB | [url_encode_13-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/url_encode_13-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `url_encode_13` | `1.2.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.1 KiB | [url_encode_13-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/url_encode_13-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `url_encode_13` | `1.2.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.1 KiB | [url_encode_13-1.2.5-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/url_encode_13-1.2.5-1PIGSTY.el10.x86_64.rpm) |
| `url_encode_13` | `1.2.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.3 KiB | [url_encode_13-1.2.5-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/url_encode_13-1.2.5-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-url-encode` | `1.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.1 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-url-encode` | `1.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.4 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-url-encode` | `1.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.5 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-url-encode` | `1.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.3 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-url-encode` | `1.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.9 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-url-encode` | `1.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.8 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-url-encode` | `1.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.8 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-url-encode` | `1.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.8 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/okbob/url_encode" title="Repository" icon="github" subtitle="github.com/okbob/url_encode" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="url_encode-1.2.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg url_encode;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install url_encode;		# install via package name, for the active PG version

pig install url_encode -v 18;   # install for PG 18
pig install url_encode -v 17;   # install for PG 17
pig install url_encode -v 16;   # install for PG 16
pig install url_encode -v 15;   # install for PG 15
pig install url_encode -v 14;   # install for PG 14
pig install url_encode -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION url_encode;
```
