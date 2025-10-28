---
title: "url_encode"
linkTitle: "url_encode"
description: "url_encode, url_decode functions"
weight: 4190
categories: ["UTIL"]
width: full
---

url_encode, url_decode functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4190** | {{< badge content="url_encode" link="https://github.com/okbob/url_encode" >}} | {{< ext "url_encode" >}} | `1.2.5` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_html5_email_address" >}} {{< ext "base36" >}} {{< ext "base62" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/url_encode" >}} | `1.2.5` | {{< bg "18" "url_encode_18*" "green" >}} {{< bg "17" "url_encode_17*" "green" >}} {{< bg "16" "url_encode_16*" "green" >}} {{< bg "15" "url_encode_15*" "green" >}} {{< bg "14" "url_encode_14*" "green" >}} {{< bg "13" "url_encode_13*" "green" >}} | `url_encode_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/url_encode" >}} | `1.2.5` | {{< bg "18" "postgresql-18-url-encode" "green" >}} {{< bg "17" "postgresql-17-url-encode" "green" >}} {{< bg "16" "postgresql-16-url-encode" "green" >}} {{< bg "15" "postgresql-15-url-encode" "green" >}} {{< bg "14" "postgresql-14-url-encode" "green" >}} {{< bg "13" "postgresql-13-url-encode" "green" >}} | `postgresql-$v-url-encode` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "url_encode_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.5" "url_encode_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "url_encode_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.5" "url_encode_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "url_encode_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.5" "url_encode_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "url_encode_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.5" "url_encode_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "url_encode_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "url_encode_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "url_encode_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "url_encode_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "url_encode_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "url_encode_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "url_encode_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "url_encode_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "url_encode_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "url_encode_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "url_encode_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "url_encode_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "url_encode_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-url-encode : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-url-encode : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-url-encode : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-url-encode : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-url-encode : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-url-encode : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-url-encode : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-url-encode : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-url-encode : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-url-encode : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-url-encode : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-url-encode : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-url-encode : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-url-encode : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-url-encode : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-url-encode : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-url-encode : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-url-encode : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.2.5" "postgresql-17-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-16-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-15-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-14-url-encode : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.5" "postgresql-13-url-encode : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `url_encode_17` | 1.2.5 | `el8.x86_64` | pigsty | 13.4 KiB | [url_encode_17-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/url_encode_17-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `url_encode_17` | 1.2.5 | `el8.aarch64` | pigsty | 13.3 KiB | [url_encode_17-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/url_encode_17-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `url_encode_17` | 1.2.5 | `el9.x86_64` | pigsty | 13.4 KiB | [url_encode_17-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/url_encode_17-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `url_encode_17` | 1.2.5 | `el9.aarch64` | pigsty | 13.4 KiB | [url_encode_17-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/url_encode_17-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-url-encode` | 1.2.5 | `d12.x86_64` | pigsty | 12.9 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-url-encode` | 1.2.5 | `d12.aarch64` | pigsty | 12.7 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-url-encode` | 1.2.5 | `u22.x86_64` | pigsty | 13.1 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-url-encode` | 1.2.5 | `u22.aarch64` | pigsty | 12.9 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-url-encode` | 1.2.5 | `u24.x86_64` | pigsty | 13.1 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-url-encode` | 1.2.5 | `u24.aarch64` | pigsty | 12.9 KiB | [postgresql-17-url-encode_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-17-url-encode_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `url_encode_16` | 1.2.5 | `el8.x86_64` | pigsty | 13.3 KiB | [url_encode_16-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/url_encode_16-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `url_encode_16` | 1.2.5 | `el8.aarch64` | pigsty | 13.2 KiB | [url_encode_16-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/url_encode_16-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `url_encode_16` | 1.2.5 | `el9.x86_64` | pigsty | 13.3 KiB | [url_encode_16-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/url_encode_16-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `url_encode_16` | 1.2.5 | `el9.aarch64` | pigsty | 13.2 KiB | [url_encode_16-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/url_encode_16-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-url-encode` | 1.2.5 | `d12.x86_64` | pigsty | 12.7 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-url-encode` | 1.2.5 | `d12.aarch64` | pigsty | 12.5 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-url-encode` | 1.2.5 | `u22.x86_64` | pigsty | 12.9 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-url-encode` | 1.2.5 | `u22.aarch64` | pigsty | 12.7 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-url-encode` | 1.2.5 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-url-encode` | 1.2.5 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-16-url-encode_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-16-url-encode_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `url_encode_15` | 1.2.5 | `el8.x86_64` | pigsty | 13.3 KiB | [url_encode_15-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/url_encode_15-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `url_encode_15` | 1.2.5 | `el8.aarch64` | pigsty | 13.3 KiB | [url_encode_15-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/url_encode_15-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `url_encode_15` | 1.2.5 | `el9.x86_64` | pigsty | 13.3 KiB | [url_encode_15-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/url_encode_15-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `url_encode_15` | 1.2.5 | `el9.aarch64` | pigsty | 13.2 KiB | [url_encode_15-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/url_encode_15-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-url-encode` | 1.2.5 | `d12.x86_64` | pigsty | 12.7 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-url-encode` | 1.2.5 | `d12.aarch64` | pigsty | 12.5 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-url-encode` | 1.2.5 | `u22.x86_64` | pigsty | 12.9 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-url-encode` | 1.2.5 | `u22.aarch64` | pigsty | 12.7 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-url-encode` | 1.2.5 | `u24.x86_64` | pigsty | 12.8 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-url-encode` | 1.2.5 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-15-url-encode_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-15-url-encode_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `url_encode_14` | 1.2.5 | `el8.x86_64` | pigsty | 13.3 KiB | [url_encode_14-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/url_encode_14-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `url_encode_14` | 1.2.5 | `el8.aarch64` | pigsty | 13.3 KiB | [url_encode_14-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/url_encode_14-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `url_encode_14` | 1.2.5 | `el9.x86_64` | pigsty | 13.3 KiB | [url_encode_14-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/url_encode_14-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `url_encode_14` | 1.2.5 | `el9.aarch64` | pigsty | 13.2 KiB | [url_encode_14-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/url_encode_14-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-url-encode` | 1.2.5 | `d12.x86_64` | pigsty | 12.6 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-url-encode` | 1.2.5 | `d12.aarch64` | pigsty | 12.4 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-url-encode` | 1.2.5 | `u22.x86_64` | pigsty | 12.9 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-url-encode` | 1.2.5 | `u22.aarch64` | pigsty | 12.7 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-url-encode` | 1.2.5 | `u24.x86_64` | pigsty | 12.7 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-url-encode` | 1.2.5 | `u24.aarch64` | pigsty | 12.6 KiB | [postgresql-14-url-encode_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-14-url-encode_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `url_encode_13` | 1.2.5 | `el8.x86_64` | pigsty | 13.1 KiB | [url_encode_13-1.2.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/url_encode_13-1.2.5-1PIGSTY.el8.x86_64.rpm) |
| `url_encode_13` | 1.2.5 | `el8.aarch64` | pigsty | 13.2 KiB | [url_encode_13-1.2.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/url_encode_13-1.2.5-1PIGSTY.el8.aarch64.rpm) |
| `url_encode_13` | 1.2.5 | `el9.x86_64` | pigsty | 13.4 KiB | [url_encode_13-1.2.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/url_encode_13-1.2.5-1PIGSTY.el9.x86_64.rpm) |
| `url_encode_13` | 1.2.5 | `el9.aarch64` | pigsty | 13.2 KiB | [url_encode_13-1.2.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/url_encode_13-1.2.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-url-encode` | 1.2.5 | `d12.x86_64` | pigsty | 12.3 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-url-encode` | 1.2.5 | `d12.aarch64` | pigsty | 12.6 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-url-encode` | 1.2.5 | `u22.x86_64` | pigsty | 12.9 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-url-encode` | 1.2.5 | `u22.aarch64` | pigsty | 12.8 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-url-encode` | 1.2.5 | `u24.x86_64` | pigsty | 12.6 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-url-encode` | 1.2.5 | `u24.aarch64` | pigsty | 12.7 KiB | [postgresql-13-url-encode_1.2.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/u/url-encode/postgresql-13-url-encode_1.2.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/okbob/url_encode" title="Repository" icon="github" subtitle="github.com/okbob/url_encode" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="url_encode-1.2.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build get url_encode; # get url_encode source code
pig build dep url_encode; # install build dependencies
pig build pkg url_encode; # build extension rpm or deb
pig build ext url_encode; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install url_encode; # install by extension name, for the current active PG version
pig ext install url_encode; # install via package alias, for the active PG version
pig ext install url_encode -v 18;   # install for PG 18
pig ext install url_encode -v 17;   # install for PG 17
pig ext install url_encode -v 16;   # install for PG 16
pig ext install url_encode -v 15;   # install for PG 15
pig ext install url_encode -v 14;   # install for PG 14
pig ext install url_encode -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION url_encode;
```

