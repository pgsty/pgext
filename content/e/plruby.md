---
title: "plruby"
linkTitle: "plruby"
description: "Embed MRI Ruby as an untrusted PostgreSQL procedural language"
weight: 3160
categories: ["LANG"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**plruby**](https://github.com/commandprompt/plruby) : Embed MRI Ruby as an untrusted PostgreSQL procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3160** | {{< badge content="plruby" link="https://github.com/commandprompt/plruby" >}} | {{< ext "plruby" >}} | `2.5` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|    **Need By**    | {{< ext "hstore_plruby" >}} {{< ext "jsonb_plruby" >}} {{< ext "ltree_plruby" >}} |
|   **See Also**    | {{< ext "jsonb_plruby" >}} {{< ext "hstore_plruby" >}} {{< ext "ltree_plruby" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "pllua" >}} {{< ext "plv8" >}} {{< ext "plphp" >}} |
|    **Siblings**   | {{< ext "jsonb_plruby" >}} {{< ext "hstore_plruby" >}} {{< ext "ltree_plruby" >}} |

> [!Note] Extension control default_version is 2.5 while the project and package version is 2.5.0; PL/Ruby embeds MRI Ruby 3.x, is untrusted and superuser-only, and requires no preload. RPM builds also provide an llvmjit subpackage.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5` | {{< bg "19" "" "red" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `plruby` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.0` | {{< bg "19" "plruby_19" "red" >}} {{< bg "18" "plruby_18" "green" >}} {{< bg "17" "plruby_17" "green" >}} {{< bg "16" "plruby_16" "green" >}} {{< bg "15" "plruby_15" "green" >}} | `plruby_$v` | `ruby-libs` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.0` | {{< bg "19" "postgresql-19-plruby" "red" >}} {{< bg "18" "postgresql-18-plruby" "green" >}} {{< bg "17" "postgresql-17-plruby" "green" >}} {{< bg "16" "postgresql-16-plruby" "green" >}} {{< bg "15" "postgresql-15-plruby" "green" >}} | `postgresql-$v-plruby` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "plruby_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "plruby_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "N/A" "plruby_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "plruby_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "N/A" "plruby_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "plruby_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "postgresql-19-plruby : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "postgresql-19-plruby : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "postgresql-19-plruby : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "postgresql-19-plruby : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "postgresql-19-plruby : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "postgresql-19-plruby : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "postgresql-19-plruby : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "postgresql-19-plruby : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "postgresql-19-plruby : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "postgresql-19-plruby : N/A 0" "gray" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plruby_18` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 138.8 KiB | [plruby_18-2.5.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plruby_18-2.5.0-1PGSTY.el8.x86_64.rpm) |
| `plruby_18` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 136.2 KiB | [plruby_18-2.5.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plruby_18-2.5.0-1PGSTY.el8.aarch64.rpm) |
| `plruby_18` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 145.9 KiB | [plruby_18-2.5.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plruby_18-2.5.0-1PGSTY.el9.x86_64.rpm) |
| `plruby_18` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 143.5 KiB | [plruby_18-2.5.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plruby_18-2.5.0-1PGSTY.el9.aarch64.rpm) |
| `plruby_18` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 140.6 KiB | [plruby_18-2.5.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plruby_18-2.5.0-1PGSTY.el10.x86_64.rpm) |
| `plruby_18` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 139.3 KiB | [plruby_18-2.5.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plruby_18-2.5.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-plruby` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 124.7 KiB | [postgresql-18-plruby_2.5.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 122.2 KiB | [postgresql-18-plruby_2.5.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 121.6 KiB | [postgresql-18-plruby_2.5.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 119.3 KiB | [postgresql-18-plruby_2.5.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 136.5 KiB | [postgresql-18-plruby_2.5.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 134.2 KiB | [postgresql-18-plruby_2.5.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 129.5 KiB | [postgresql-18-plruby_2.5.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 128.3 KiB | [postgresql-18-plruby_2.5.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 126.1 KiB | [postgresql-18-plruby_2.5.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 124.9 KiB | [postgresql-18-plruby_2.5.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plruby_17` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 138.6 KiB | [plruby_17-2.5.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plruby_17-2.5.0-1PGSTY.el8.x86_64.rpm) |
| `plruby_17` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 136.1 KiB | [plruby_17-2.5.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plruby_17-2.5.0-1PGSTY.el8.aarch64.rpm) |
| `plruby_17` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 145.6 KiB | [plruby_17-2.5.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plruby_17-2.5.0-1PGSTY.el9.x86_64.rpm) |
| `plruby_17` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 143.4 KiB | [plruby_17-2.5.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plruby_17-2.5.0-1PGSTY.el9.aarch64.rpm) |
| `plruby_17` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 140.3 KiB | [plruby_17-2.5.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plruby_17-2.5.0-1PGSTY.el10.x86_64.rpm) |
| `plruby_17` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 138.9 KiB | [plruby_17-2.5.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plruby_17-2.5.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-plruby` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 124.5 KiB | [postgresql-17-plruby_2.5.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 121.9 KiB | [postgresql-17-plruby_2.5.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 121.4 KiB | [postgresql-17-plruby_2.5.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 119.0 KiB | [postgresql-17-plruby_2.5.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 154.3 KiB | [postgresql-17-plruby_2.5.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 151.4 KiB | [postgresql-17-plruby_2.5.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 129.0 KiB | [postgresql-17-plruby_2.5.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 127.8 KiB | [postgresql-17-plruby_2.5.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 125.9 KiB | [postgresql-17-plruby_2.5.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 124.6 KiB | [postgresql-17-plruby_2.5.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plruby_16` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 138.5 KiB | [plruby_16-2.5.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plruby_16-2.5.0-1PGSTY.el8.x86_64.rpm) |
| `plruby_16` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 136.1 KiB | [plruby_16-2.5.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plruby_16-2.5.0-1PGSTY.el8.aarch64.rpm) |
| `plruby_16` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 145.7 KiB | [plruby_16-2.5.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plruby_16-2.5.0-1PGSTY.el9.x86_64.rpm) |
| `plruby_16` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 143.4 KiB | [plruby_16-2.5.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plruby_16-2.5.0-1PGSTY.el9.aarch64.rpm) |
| `plruby_16` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 140.2 KiB | [plruby_16-2.5.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plruby_16-2.5.0-1PGSTY.el10.x86_64.rpm) |
| `plruby_16` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 139.0 KiB | [plruby_16-2.5.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plruby_16-2.5.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-plruby` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 124.5 KiB | [postgresql-16-plruby_2.5.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 122.0 KiB | [postgresql-16-plruby_2.5.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 121.4 KiB | [postgresql-16-plruby_2.5.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 119.0 KiB | [postgresql-16-plruby_2.5.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 153.2 KiB | [postgresql-16-plruby_2.5.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 150.3 KiB | [postgresql-16-plruby_2.5.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 129.1 KiB | [postgresql-16-plruby_2.5.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 127.9 KiB | [postgresql-16-plruby_2.5.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 125.9 KiB | [postgresql-16-plruby_2.5.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 124.6 KiB | [postgresql-16-plruby_2.5.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plruby_15` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 138.3 KiB | [plruby_15-2.5.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plruby_15-2.5.0-1PGSTY.el8.x86_64.rpm) |
| `plruby_15` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 135.9 KiB | [plruby_15-2.5.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plruby_15-2.5.0-1PGSTY.el8.aarch64.rpm) |
| `plruby_15` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 145.8 KiB | [plruby_15-2.5.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plruby_15-2.5.0-1PGSTY.el9.x86_64.rpm) |
| `plruby_15` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 144.0 KiB | [plruby_15-2.5.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plruby_15-2.5.0-1PGSTY.el9.aarch64.rpm) |
| `plruby_15` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 140.4 KiB | [plruby_15-2.5.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plruby_15-2.5.0-1PGSTY.el10.x86_64.rpm) |
| `plruby_15` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 139.7 KiB | [plruby_15-2.5.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plruby_15-2.5.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-plruby` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 124.7 KiB | [postgresql-15-plruby_2.5.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 122.1 KiB | [postgresql-15-plruby_2.5.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 121.0 KiB | [postgresql-15-plruby_2.5.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 118.7 KiB | [postgresql-15-plruby_2.5.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 153.4 KiB | [postgresql-15-plruby_2.5.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 151.0 KiB | [postgresql-15-plruby_2.5.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 129.1 KiB | [postgresql-15-plruby_2.5.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 128.7 KiB | [postgresql-15-plruby_2.5.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 126.1 KiB | [postgresql-15-plruby_2.5.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 125.3 KiB | [postgresql-15-plruby_2.5.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/commandprompt/plruby" title="Repository" icon="github" subtitle="github.com/commandprompt/plruby" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plruby-2.5.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg plruby;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install plruby;		# install via package name, for the active PG version

pig install plruby -v 18;   # install for PG 18
pig install plruby -v 17;   # install for PG 17
pig install plruby -v 16;   # install for PG 16
pig install plruby -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION plruby;
```

## Usage

Sources:

- [PL/Ruby v2.5.0 README](https://github.com/commandprompt/plruby/blob/v2.5.0/README.md)
- [PL/Ruby language reference](https://github.com/commandprompt/plruby/blob/v2.5.0/doc/plruby.md)
- [PL/Ruby cookbook](https://github.com/commandprompt/plruby/blob/v2.5.0/doc/cookbook.md)
- [PL/Ruby v2.5.0 control file](https://github.com/commandprompt/plruby/blob/v2.5.0/plruby.control)
- [PL/Ruby changelog](https://github.com/commandprompt/plruby/blob/v2.5.0/CHANGELOG.md)

`plruby` is the maintained Command Prompt procedural-language extension that embeds Ruby 3 in PostgreSQL. Package release 2.5.0 installs SQL extension version `2.5`. It supports scalar and set-returning functions, triggers, event triggers, procedures, anonymous `DO` blocks, SPI queries, cursors, and prepared plans.

### Create a Function

```sql
CREATE EXTENSION plruby;

CREATE FUNCTION ruby_add(integer, integer)
RETURNS integer
LANGUAGE plruby
AS $$
  args[0] + args[1]
$$;

SELECT ruby_add(2, 3);
```

Arguments are exposed through `args`; Ruby's final expression becomes the SQL return value. PostgreSQL scalar, array, composite, and record conversion rules are documented in the language reference.

### Set-Returning Functions

Use `return_next` to emit rows from a set-returning function:

```sql
CREATE FUNCTION ruby_series(integer)
RETURNS SETOF integer
LANGUAGE plruby
AS $$
  1.upto(args[0]) { |n| return_next(n) }
$$;

SELECT * FROM ruby_series(3);
```

### SPI and Database Work

PL/Ruby exposes PostgreSQL's Server Programming Interface for SQL execution, prepared plans, and cursors. Keep SQL values in parameters rather than interpolating them into command text, and release long-lived cursors or prepared state when the session no longer needs them.

Procedures can use the documented transaction-control surface where PostgreSQL permits `COMMIT` or `ROLLBACK`. Functions and triggers remain subject to PostgreSQL's normal transactional restrictions.

### Triggers and Session State

Trigger functions receive trigger metadata through `$_TD` and return the row action documented by PL/Ruby. Event triggers, anonymous `DO` blocks, backend-local session data, and shared data are also available. These features run inside the database backend, so an exception, blocking call, or memory leak directly affects that backend.

### Version 2.5.0

- `bytea` now maps to a raw, NUL-safe Ruby `String` with `ASCII-8BIT` encoding instead of PostgreSQL hex text. This is a breaking conversion change: audit functions that parse or construct `\x...` strings and build bytes explicitly, for example with `Array#pack`.
- `$_SD` adds per-function state that persists across calls in one session and resets when the function is recompiled. `$_SHARED` remains session-wide across PL/Ruby functions.
- `spi_colnames`, `spi_coltypes`, and `spi_coltypmods` expose result-column metadata, and `ltree_plruby` adds the opt-in `ltree` transform.
- After installing the 2.5.0 shared library and SQL files, run `ALTER EXTENSION plruby UPDATE` in each database that already has the extension.

### Security and Requirements

- `plruby` is an untrusted language. Ruby 3 provides no safe in-process sandbox, so creating PL/Ruby functions is restricted to superusers and code executes with the PostgreSQL server process's operating-system authority.
- Review all PL/Ruby source as privileged server code. Never allow tenants or ordinary application roles to submit arbitrary Ruby.
- Upstream v2.5.0 supports PostgreSQL 11-18 and Ruby 3.x. Current Pigsty packages target PostgreSQL 14-18.
- No `shared_preload_libraries` setting is required. Existing sessions must reconnect after server-side library replacement before assuming a new runtime is active.
- `jsonb_plruby`, `hstore_plruby`, and `ltree_plruby` are companion transforms. A function must explicitly declare `TRANSFORM FOR TYPE ...` to receive native Ruby structures instead of the normal datum wrapper/conversion path.
