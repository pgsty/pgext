---
title: "plphp"
linkTitle: "plphp"
description: "Embed PHP 8 as an untrusted PostgreSQL procedural language"
weight: 3170
categories: ["LANG"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**plphp**](https://github.com/commandprompt/plPHP) : Embed PHP 8 as an untrusted PostgreSQL procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3170** | {{< badge content="plphp" link="https://github.com/commandprompt/plPHP" >}} | {{< ext "plphp" >}} | `2.6` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "bytea_plphp" >}} {{< ext "hstore_plphp" >}} {{< ext "jsonb_plphp" >}} |
|   **See Also**    | {{< ext "plruby" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "pllua" >}} {{< ext "plv8" >}} {{< ext "jsonb_plphp" >}} {{< ext "hstore_plphp" >}} {{< ext "bytea_plphp" >}} |

> [!Note] Uses non-ZTS PHP embed. EL8/EL9 RPMs require the PHP 8.2/8.3 module streams respectively; DEBs pin the matching PHP embed runtime for each OS. No preload or fixed schema. The three optional transforms are not included in this package.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plphp` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.6` | {{< bg "18" "plphp_18" "green" >}} {{< bg "17" "plphp_17" "green" >}} {{< bg "16" "plphp_16" "green" >}} {{< bg "15" "plphp_15" "green" >}} {{< bg "14" "plphp_14" "green" >}} | `plphp_$v` | `php-embedded` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.6` | {{< bg "18" "postgresql-18-plphp" "green" >}} {{< bg "17" "postgresql-17-plphp" "green" >}} {{< bg "16" "postgresql-16-plphp" "green" >}} {{< bg "15" "postgresql-15-plphp" "green" >}} {{< bg "14" "postgresql-14-plphp" "green" >}} | `postgresql-$v-plphp` | `libphp-embed` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.6" "plphp_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.6" "plphp_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.6" "plphp_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.6" "plphp_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.6" "plphp_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.6" "plphp_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "plphp_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.6" "postgresql-18-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-17-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-16-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-15-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-14-plphp : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.6" "postgresql-18-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-17-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-16-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-15-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-14-plphp : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.6" "postgresql-18-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-17-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-16-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-15-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-14-plphp : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.6" "postgresql-18-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-17-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-16-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-15-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-14-plphp : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.6" "postgresql-18-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-17-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-16-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-15-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-14-plphp : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.6" "postgresql-18-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-17-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-16-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-15-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-14-plphp : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.6" "postgresql-18-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-17-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-16-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-15-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-14-plphp : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.6" "postgresql-18-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-17-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-16-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-15-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-14-plphp : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.6" "postgresql-18-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-17-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-16-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-15-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-14-plphp : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.6" "postgresql-18-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-17-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-16-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-15-plphp : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6" "postgresql-14-plphp : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plphp_18` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 120.5 KiB | [plphp_18-2.6-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plphp_18-2.6-1PGSTY.el8.x86_64.rpm) |
| `plphp_18` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 117.2 KiB | [plphp_18-2.6-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plphp_18-2.6-1PGSTY.el8.aarch64.rpm) |
| `plphp_18` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 123.7 KiB | [plphp_18-2.6-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plphp_18-2.6-1PGSTY.el9.x86_64.rpm) |
| `plphp_18` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 120.8 KiB | [plphp_18-2.6-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plphp_18-2.6-1PGSTY.el9.aarch64.rpm) |
| `plphp_18` | `2.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 123.4 KiB | [plphp_18-2.6-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plphp_18-2.6-1PGSTY.el10.x86_64.rpm) |
| `plphp_18` | `2.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 121.2 KiB | [plphp_18-2.6-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plphp_18-2.6-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-plphp` | `2.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 114.1 KiB | [postgresql-18-plphp_2.6-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plphp/postgresql-18-plphp_2.6-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-plphp` | `2.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 110.8 KiB | [postgresql-18-plphp_2.6-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plphp/postgresql-18-plphp_2.6-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-plphp` | `2.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 114.5 KiB | [postgresql-18-plphp_2.6-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plphp/postgresql-18-plphp_2.6-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-plphp` | `2.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 111.8 KiB | [postgresql-18-plphp_2.6-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plphp/postgresql-18-plphp_2.6-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-plphp` | `2.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 117.7 KiB | [postgresql-18-plphp_2.6-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plphp/postgresql-18-plphp_2.6-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-plphp` | `2.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 115.1 KiB | [postgresql-18-plphp_2.6-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plphp/postgresql-18-plphp_2.6-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-plphp` | `2.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 114.0 KiB | [postgresql-18-plphp_2.6-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plphp/postgresql-18-plphp_2.6-1PGSTY~noble_amd64.deb) |
| `postgresql-18-plphp` | `2.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 111.8 KiB | [postgresql-18-plphp_2.6-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plphp/postgresql-18-plphp_2.6-1PGSTY~noble_arm64.deb) |
| `postgresql-18-plphp` | `2.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 112.8 KiB | [postgresql-18-plphp_2.6-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plphp/postgresql-18-plphp_2.6-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-plphp` | `2.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 109.7 KiB | [postgresql-18-plphp_2.6-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plphp/postgresql-18-plphp_2.6-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plphp_17` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 120.3 KiB | [plphp_17-2.6-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plphp_17-2.6-1PGSTY.el8.x86_64.rpm) |
| `plphp_17` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 117.1 KiB | [plphp_17-2.6-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plphp_17-2.6-1PGSTY.el8.aarch64.rpm) |
| `plphp_17` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 123.6 KiB | [plphp_17-2.6-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plphp_17-2.6-1PGSTY.el9.x86_64.rpm) |
| `plphp_17` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 120.6 KiB | [plphp_17-2.6-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plphp_17-2.6-1PGSTY.el9.aarch64.rpm) |
| `plphp_17` | `2.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 123.4 KiB | [plphp_17-2.6-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plphp_17-2.6-1PGSTY.el10.x86_64.rpm) |
| `plphp_17` | `2.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 121.0 KiB | [plphp_17-2.6-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plphp_17-2.6-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-plphp` | `2.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 114.2 KiB | [postgresql-17-plphp_2.6-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plphp/postgresql-17-plphp_2.6-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-plphp` | `2.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 111.2 KiB | [postgresql-17-plphp_2.6-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plphp/postgresql-17-plphp_2.6-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-plphp` | `2.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 114.4 KiB | [postgresql-17-plphp_2.6-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plphp/postgresql-17-plphp_2.6-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-plphp` | `2.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 111.7 KiB | [postgresql-17-plphp_2.6-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plphp/postgresql-17-plphp_2.6-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-plphp` | `2.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 138.6 KiB | [postgresql-17-plphp_2.6-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plphp/postgresql-17-plphp_2.6-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-plphp` | `2.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 136.0 KiB | [postgresql-17-plphp_2.6-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plphp/postgresql-17-plphp_2.6-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-plphp` | `2.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 113.5 KiB | [postgresql-17-plphp_2.6-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plphp/postgresql-17-plphp_2.6-1PGSTY~noble_amd64.deb) |
| `postgresql-17-plphp` | `2.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 111.6 KiB | [postgresql-17-plphp_2.6-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plphp/postgresql-17-plphp_2.6-1PGSTY~noble_arm64.deb) |
| `postgresql-17-plphp` | `2.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 112.4 KiB | [postgresql-17-plphp_2.6-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plphp/postgresql-17-plphp_2.6-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-plphp` | `2.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 109.3 KiB | [postgresql-17-plphp_2.6-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plphp/postgresql-17-plphp_2.6-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plphp_16` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 120.3 KiB | [plphp_16-2.6-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plphp_16-2.6-1PGSTY.el8.x86_64.rpm) |
| `plphp_16` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 117.0 KiB | [plphp_16-2.6-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plphp_16-2.6-1PGSTY.el8.aarch64.rpm) |
| `plphp_16` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 123.6 KiB | [plphp_16-2.6-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plphp_16-2.6-1PGSTY.el9.x86_64.rpm) |
| `plphp_16` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 120.6 KiB | [plphp_16-2.6-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plphp_16-2.6-1PGSTY.el9.aarch64.rpm) |
| `plphp_16` | `2.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 123.4 KiB | [plphp_16-2.6-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plphp_16-2.6-1PGSTY.el10.x86_64.rpm) |
| `plphp_16` | `2.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 121.1 KiB | [plphp_16-2.6-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plphp_16-2.6-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-plphp` | `2.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 114.3 KiB | [postgresql-16-plphp_2.6-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plphp/postgresql-16-plphp_2.6-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-plphp` | `2.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 111.2 KiB | [postgresql-16-plphp_2.6-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plphp/postgresql-16-plphp_2.6-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-plphp` | `2.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 114.5 KiB | [postgresql-16-plphp_2.6-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plphp/postgresql-16-plphp_2.6-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-plphp` | `2.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 111.8 KiB | [postgresql-16-plphp_2.6-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plphp/postgresql-16-plphp_2.6-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-plphp` | `2.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 137.3 KiB | [postgresql-16-plphp_2.6-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plphp/postgresql-16-plphp_2.6-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-plphp` | `2.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 134.7 KiB | [postgresql-16-plphp_2.6-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plphp/postgresql-16-plphp_2.6-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-plphp` | `2.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 113.5 KiB | [postgresql-16-plphp_2.6-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plphp/postgresql-16-plphp_2.6-1PGSTY~noble_amd64.deb) |
| `postgresql-16-plphp` | `2.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 111.6 KiB | [postgresql-16-plphp_2.6-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plphp/postgresql-16-plphp_2.6-1PGSTY~noble_arm64.deb) |
| `postgresql-16-plphp` | `2.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 112.4 KiB | [postgresql-16-plphp_2.6-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plphp/postgresql-16-plphp_2.6-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-plphp` | `2.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 109.3 KiB | [postgresql-16-plphp_2.6-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plphp/postgresql-16-plphp_2.6-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plphp_15` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 120.6 KiB | [plphp_15-2.6-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plphp_15-2.6-1PGSTY.el8.x86_64.rpm) |
| `plphp_15` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 117.3 KiB | [plphp_15-2.6-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plphp_15-2.6-1PGSTY.el8.aarch64.rpm) |
| `plphp_15` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 123.7 KiB | [plphp_15-2.6-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plphp_15-2.6-1PGSTY.el9.x86_64.rpm) |
| `plphp_15` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 121.3 KiB | [plphp_15-2.6-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plphp_15-2.6-1PGSTY.el9.aarch64.rpm) |
| `plphp_15` | `2.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 123.8 KiB | [plphp_15-2.6-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plphp_15-2.6-1PGSTY.el10.x86_64.rpm) |
| `plphp_15` | `2.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 121.4 KiB | [plphp_15-2.6-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plphp_15-2.6-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-plphp` | `2.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 114.6 KiB | [postgresql-15-plphp_2.6-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plphp/postgresql-15-plphp_2.6-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-plphp` | `2.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 111.3 KiB | [postgresql-15-plphp_2.6-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plphp/postgresql-15-plphp_2.6-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-plphp` | `2.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 115.0 KiB | [postgresql-15-plphp_2.6-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plphp/postgresql-15-plphp_2.6-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-plphp` | `2.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 111.7 KiB | [postgresql-15-plphp_2.6-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plphp/postgresql-15-plphp_2.6-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-plphp` | `2.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 137.8 KiB | [postgresql-15-plphp_2.6-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plphp/postgresql-15-plphp_2.6-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-plphp` | `2.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 135.5 KiB | [postgresql-15-plphp_2.6-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plphp/postgresql-15-plphp_2.6-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-plphp` | `2.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 114.1 KiB | [postgresql-15-plphp_2.6-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plphp/postgresql-15-plphp_2.6-1PGSTY~noble_amd64.deb) |
| `postgresql-15-plphp` | `2.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 112.2 KiB | [postgresql-15-plphp_2.6-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plphp/postgresql-15-plphp_2.6-1PGSTY~noble_arm64.deb) |
| `postgresql-15-plphp` | `2.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 113.0 KiB | [postgresql-15-plphp_2.6-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plphp/postgresql-15-plphp_2.6-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-plphp` | `2.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 109.8 KiB | [postgresql-15-plphp_2.6-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plphp/postgresql-15-plphp_2.6-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plphp_14` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 120.7 KiB | [plphp_14-2.6-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plphp_14-2.6-1PGSTY.el8.x86_64.rpm) |
| `plphp_14` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 117.3 KiB | [plphp_14-2.6-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plphp_14-2.6-1PGSTY.el8.aarch64.rpm) |
| `plphp_14` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 123.8 KiB | [plphp_14-2.6-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plphp_14-2.6-1PGSTY.el9.x86_64.rpm) |
| `plphp_14` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 121.2 KiB | [plphp_14-2.6-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plphp_14-2.6-1PGSTY.el9.aarch64.rpm) |
| `plphp_14` | `2.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 123.7 KiB | [plphp_14-2.6-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plphp_14-2.6-1PGSTY.el10.x86_64.rpm) |
| `plphp_14` | `2.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 121.3 KiB | [plphp_14-2.6-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plphp_14-2.6-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-plphp` | `2.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 117.3 KiB | [postgresql-14-plphp_2.6-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plphp/postgresql-14-plphp_2.6-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-plphp` | `2.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 114.1 KiB | [postgresql-14-plphp_2.6-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plphp/postgresql-14-plphp_2.6-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-plphp` | `2.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 117.8 KiB | [postgresql-14-plphp_2.6-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plphp/postgresql-14-plphp_2.6-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-plphp` | `2.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 114.5 KiB | [postgresql-14-plphp_2.6-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plphp/postgresql-14-plphp_2.6-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-plphp` | `2.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 140.7 KiB | [postgresql-14-plphp_2.6-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plphp/postgresql-14-plphp_2.6-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-plphp` | `2.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 138.3 KiB | [postgresql-14-plphp_2.6-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plphp/postgresql-14-plphp_2.6-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-plphp` | `2.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 116.9 KiB | [postgresql-14-plphp_2.6-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plphp/postgresql-14-plphp_2.6-1PGSTY~noble_amd64.deb) |
| `postgresql-14-plphp` | `2.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 115.0 KiB | [postgresql-14-plphp_2.6-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plphp/postgresql-14-plphp_2.6-1PGSTY~noble_arm64.deb) |
| `postgresql-14-plphp` | `2.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 115.8 KiB | [postgresql-14-plphp_2.6-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plphp/postgresql-14-plphp_2.6-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-plphp` | `2.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 112.8 KiB | [postgresql-14-plphp_2.6-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plphp/postgresql-14-plphp_2.6-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/commandprompt/plPHP" title="Repository" icon="github" subtitle="github.com/commandprompt/plPHP" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plphp-2.6.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg plphp;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install plphp;		# install via package name, for the active PG version

pig install plphp -v 18;   # install for PG 18
pig install plphp -v 17;   # install for PG 17
pig install plphp -v 16;   # install for PG 16
pig install plphp -v 15;   # install for PG 15
pig install plphp -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plphp;
```

## Usage

Sources:

- [PGXN 2.6.0 README](https://pgxn.org/dist/plphp/2.6.0/README.html)
- [PL/php language reference](https://pgxn.org/dist/plphp/2.6.0/doc/plphp.html)
- [plphp control file](https://api.pgxn.org/src/plphp/plphp-2.6.0/plphp.control)
- [PL/php 2.6.0 changelog](https://api.pgxn.org/src/plphp/plphp-2.6.0/CHANGELOG.md)

`plphp` embeds PHP as an untrusted PostgreSQL procedural language for functions, procedures, triggers, event triggers, and anonymous blocks. Use it only when PHP APIs are worth running with the operating-system privileges of the PostgreSQL server process; it is not a sandboxed language.

### Core Workflow

```sql
CREATE EXTENSION plphp;

CREATE FUNCTION php_square(integer)
RETURNS integer
LANGUAGE plphp
AS $$
    return $args[0] * $args[0];
$$;

SELECT php_square(12);
```

PL/php 2.6 supports PHP 8.1 through 8.4 using the embed SAPI without ZTS and PostgreSQL 11 through 18. Installation and creation require a superuser because the language is deliberately untrusted.

### Main Interfaces

- `spi_exec(...)`, `spi_prepare(...)`, `spi_exec_prepared(...)`, `spi_query(...)`, and cursor helpers access PostgreSQL through SPI.
- `return_next()` emits rows from set-returning functions.
- `spi_commit()` and `spi_rollback()` provide transaction control inside procedures.
- `subtransaction(...)` creates a catchable subtransaction boundary.
- `$_TD` exposes trigger and event-trigger context.
- `$_SHARED` is session-global storage, while `$_SD` is private per-function session storage.
- `plphp.on_init`, `plphp.start_proc`, and `plphp_modules` initialize a session interpreter.

The optional `jsonb_plphp`, `hstore_plphp`, and `bytea_plphp` extensions add native transforms. A function must declare `TRANSFORM FOR TYPE ...` to use a transform; installing a companion does not change every PL/php function automatically.

### Security and Operations

A PL/php function can read or write files, open network connections, invoke PHP facilities, and affect anything accessible to the PostgreSQL operating-system account. Only grant function ownership and creation rights to roles trusted like a server administrator. Apply normal `SECURITY DEFINER` and `search_path` hardening when a privileged wrapper is unavoidable.

SPI query text, dynamic identifiers, session initialization code, loaded modules, and persistent interpreter state all require review. Long-running PHP code blocks a PostgreSQL backend, and memory retained in `$_SHARED` or `$_SD` lasts for the session. Use prepared statements for untrusted values, bound resource use, test cancellation/error paths, and recycle sessions when application code retains excessive state.

Version 2.6 adds per-function `$_SD` and the binary-safe `bytea_plphp` transform. PostgreSQL tracks the control version as `2.6` while the source distribution is `2.6.0`.
