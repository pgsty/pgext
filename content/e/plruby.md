---
title: "plruby"
linkTitle: "plruby"
description: "Embed MRI Ruby as an untrusted PostgreSQL procedural language"
weight: 3160
categories: ["LANG"]
width: full
---

[**plruby**](https://github.com/commandprompt/plruby) : Embed MRI Ruby as an untrusted PostgreSQL procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3160** | {{< badge content="plruby" link="https://github.com/commandprompt/plruby" >}} | {{< ext "plruby" >}} | `2.5` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|    **Need By**    | {{< ext "hstore_plruby" >}} {{< ext "jsonb_plruby" >}} {{< ext "ltree_plruby" >}} |
|   **See Also**    | {{< ext "jsonb_plruby" >}} {{< ext "hstore_plruby" >}} {{< ext "ltree_plruby" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} {{< ext "pllua" >}} {{< ext "plv8" >}} {{< ext "plrust" >}} |
|    **Siblings**   | {{< ext "jsonb_plruby" >}} {{< ext "hstore_plruby" >}} {{< ext "ltree_plruby" >}} |

> [!Note] Extension control default_version is 2.5 while the project and package version is 2.5.0; PL/Ruby embeds MRI Ruby 3.x, is untrusted and superuser-only, and requires no preload. RPM builds also provide an llvmjit subpackage.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plruby` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.0` | {{< bg "18" "plruby_18" "green" >}} {{< bg "17" "plruby_17" "green" >}} {{< bg "16" "plruby_16" "green" >}} {{< bg "15" "plruby_15" "green" >}} {{< bg "14" "plruby_14" "green" >}} | `plruby_$v` | `ruby-libs` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.0` | {{< bg "18" "postgresql-18-plruby" "green" >}} {{< bg "17" "postgresql-17-plruby" "green" >}} {{< bg "16" "postgresql-16-plruby" "green" >}} {{< bg "15" "postgresql-15-plruby" "green" >}} {{< bg "14" "postgresql-14-plruby" "green" >}} | `postgresql-$v-plruby` | `libruby3.0 | libruby3.1 | libruby3.2 | libruby3.3` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "plruby_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "plruby_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-plruby : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-plruby : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plruby_18` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.1 KiB | [plruby_18-2.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plruby_18-2.5.0-1PIGSTY.el8.x86_64.rpm) |
| `plruby_18` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 61.4 KiB | [plruby_18-2.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plruby_18-2.5.0-1PIGSTY.el8.aarch64.rpm) |
| `plruby_18` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.3 KiB | [plruby_18-2.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plruby_18-2.5.0-1PIGSTY.el9.x86_64.rpm) |
| `plruby_18` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.1 KiB | [plruby_18-2.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plruby_18-2.5.0-1PIGSTY.el9.aarch64.rpm) |
| `plruby_18` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.1 KiB | [plruby_18-2.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plruby_18-2.5.0-1PIGSTY.el10.x86_64.rpm) |
| `plruby_18` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.1 KiB | [plruby_18-2.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plruby_18-2.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-plruby` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 138.5 KiB | [postgresql-18-plruby_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 135.9 KiB | [postgresql-18-plruby_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 135.6 KiB | [postgresql-18-plruby_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 133.3 KiB | [postgresql-18-plruby_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 151.1 KiB | [postgresql-18-plruby_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 148.6 KiB | [postgresql-18-plruby_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 143.4 KiB | [postgresql-18-plruby_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 141.8 KiB | [postgresql-18-plruby_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 140.3 KiB | [postgresql-18-plruby_2.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-plruby` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 139.1 KiB | [postgresql-18-plruby_2.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-18-plruby_2.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plruby_17` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 62.9 KiB | [plruby_17-2.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plruby_17-2.5.0-1PIGSTY.el8.x86_64.rpm) |
| `plruby_17` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 61.4 KiB | [plruby_17-2.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plruby_17-2.5.0-1PIGSTY.el8.aarch64.rpm) |
| `plruby_17` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.3 KiB | [plruby_17-2.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plruby_17-2.5.0-1PIGSTY.el9.x86_64.rpm) |
| `plruby_17` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.1 KiB | [plruby_17-2.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plruby_17-2.5.0-1PIGSTY.el9.aarch64.rpm) |
| `plruby_17` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.0 KiB | [plruby_17-2.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plruby_17-2.5.0-1PIGSTY.el10.x86_64.rpm) |
| `plruby_17` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.2 KiB | [plruby_17-2.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plruby_17-2.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-plruby` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 138.1 KiB | [postgresql-17-plruby_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 135.4 KiB | [postgresql-17-plruby_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 135.2 KiB | [postgresql-17-plruby_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 132.8 KiB | [postgresql-17-plruby_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 168.3 KiB | [postgresql-17-plruby_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 165.9 KiB | [postgresql-17-plruby_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 142.4 KiB | [postgresql-17-plruby_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 141.4 KiB | [postgresql-17-plruby_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 139.9 KiB | [postgresql-17-plruby_2.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-plruby` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 138.7 KiB | [postgresql-17-plruby_2.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-17-plruby_2.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plruby_16` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 62.9 KiB | [plruby_16-2.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plruby_16-2.5.0-1PIGSTY.el8.x86_64.rpm) |
| `plruby_16` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 61.4 KiB | [plruby_16-2.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plruby_16-2.5.0-1PIGSTY.el8.aarch64.rpm) |
| `plruby_16` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.3 KiB | [plruby_16-2.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plruby_16-2.5.0-1PIGSTY.el9.x86_64.rpm) |
| `plruby_16` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.2 KiB | [plruby_16-2.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plruby_16-2.5.0-1PIGSTY.el9.aarch64.rpm) |
| `plruby_16` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.0 KiB | [plruby_16-2.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plruby_16-2.5.0-1PIGSTY.el10.x86_64.rpm) |
| `plruby_16` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 61.0 KiB | [plruby_16-2.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plruby_16-2.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-plruby` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 138.0 KiB | [postgresql-16-plruby_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 135.3 KiB | [postgresql-16-plruby_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 134.9 KiB | [postgresql-16-plruby_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 132.6 KiB | [postgresql-16-plruby_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 167.0 KiB | [postgresql-16-plruby_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 164.5 KiB | [postgresql-16-plruby_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 142.4 KiB | [postgresql-16-plruby_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 141.2 KiB | [postgresql-16-plruby_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 139.7 KiB | [postgresql-16-plruby_2.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-plruby` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 138.4 KiB | [postgresql-16-plruby_2.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-16-plruby_2.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plruby_15` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.3 KiB | [plruby_15-2.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plruby_15-2.5.0-1PIGSTY.el8.x86_64.rpm) |
| `plruby_15` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 61.9 KiB | [plruby_15-2.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plruby_15-2.5.0-1PIGSTY.el8.aarch64.rpm) |
| `plruby_15` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 63.0 KiB | [plruby_15-2.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plruby_15-2.5.0-1PIGSTY.el9.x86_64.rpm) |
| `plruby_15` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 62.4 KiB | [plruby_15-2.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plruby_15-2.5.0-1PIGSTY.el9.aarch64.rpm) |
| `plruby_15` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.3 KiB | [plruby_15-2.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plruby_15-2.5.0-1PIGSTY.el10.x86_64.rpm) |
| `plruby_15` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 62.4 KiB | [plruby_15-2.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plruby_15-2.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-plruby` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 137.5 KiB | [postgresql-15-plruby_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 135.1 KiB | [postgresql-15-plruby_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 134.2 KiB | [postgresql-15-plruby_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 132.1 KiB | [postgresql-15-plruby_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 167.1 KiB | [postgresql-15-plruby_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 165.0 KiB | [postgresql-15-plruby_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 142.0 KiB | [postgresql-15-plruby_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 141.4 KiB | [postgresql-15-plruby_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 139.5 KiB | [postgresql-15-plruby_2.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-plruby` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 138.5 KiB | [postgresql-15-plruby_2.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-15-plruby_2.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plruby_14` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 63.2 KiB | [plruby_14-2.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plruby_14-2.5.0-1PIGSTY.el8.x86_64.rpm) |
| `plruby_14` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 61.9 KiB | [plruby_14-2.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plruby_14-2.5.0-1PIGSTY.el8.aarch64.rpm) |
| `plruby_14` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 63.0 KiB | [plruby_14-2.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plruby_14-2.5.0-1PIGSTY.el9.x86_64.rpm) |
| `plruby_14` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 62.5 KiB | [plruby_14-2.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plruby_14-2.5.0-1PIGSTY.el9.aarch64.rpm) |
| `plruby_14` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.5 KiB | [plruby_14-2.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plruby_14-2.5.0-1PIGSTY.el10.x86_64.rpm) |
| `plruby_14` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 62.4 KiB | [plruby_14-2.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plruby_14-2.5.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-plruby` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 138.3 KiB | [postgresql-14-plruby_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-14-plruby_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-plruby` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 135.8 KiB | [postgresql-14-plruby_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plruby/postgresql-14-plruby_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-plruby` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 134.2 KiB | [postgresql-14-plruby_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-14-plruby_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-plruby` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 131.9 KiB | [postgresql-14-plruby_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plruby/postgresql-14-plruby_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-plruby` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 164.4 KiB | [postgresql-14-plruby_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-14-plruby_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-plruby` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 162.3 KiB | [postgresql-14-plruby_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plruby/postgresql-14-plruby_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-plruby` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 142.1 KiB | [postgresql-14-plruby_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-14-plruby_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-plruby` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 141.6 KiB | [postgresql-14-plruby_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plruby/postgresql-14-plruby_2.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-plruby` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 139.3 KiB | [postgresql-14-plruby_2.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-14-plruby_2.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-plruby` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 138.5 KiB | [postgresql-14-plruby_2.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plruby/postgresql-14-plruby_2.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/commandprompt/plruby" title="Repository" icon="github" subtitle="github.com/commandprompt/plruby" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plruby-2.5.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg plruby;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install plruby;		# install via package name, for the active PG version

pig install plruby -v 18;   # install for PG 18
pig install plruby -v 17;   # install for PG 17
pig install plruby -v 16;   # install for PG 16
pig install plruby -v 15;   # install for PG 15
pig install plruby -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plruby;
```

## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/godfat/plruby/blob/548db739064a02dad4418376f46fc25e8e842f29/README.md)
- [Version 0.0.1 SQL objects](https://github.com/godfat/plruby/blob/548db739064a02dad4418376f46fc25e8e842f29/plruby--0.0.1.sql)
- [Call-handler implementation](https://github.com/godfat/plruby/blob/548db739064a02dad4418376f46fc25e8e842f29/plruby.c)

`plruby` is an extremely early procedural-language prototype embedding Ruby in a PostgreSQL backend. Version 0.0.1 installs the untrusted language `plruby` with call, inline, and validator handlers; upstream describes it only as “sort of working” on one computer.

```sql
CREATE EXTENSION plruby;
CREATE FUNCTION ruby_add(integer, integer)
RETURNS integer
LANGUAGE plruby
AS $$ args[0] + args[1] $$;
SELECT ruby_add(2, 3);
```

The example is suitable only for an isolated build test after reviewing the handler's actual argument and return conventions. An untrusted procedural language can execute with the database server process's operating-system authority and must remain superuser-only.

Do not deploy this prototype in production or expose it to tenant-supplied code. It copies substantial early PL/V8 patterns, has no current Ruby/PostgreSQL compatibility matrix, and makes no sandbox, memory, exception, signal, threading, garbage-collection, or transaction-safety guarantees. Prefer a maintained procedural language or execute Ruby outside the database. Any archival test belongs in a disposable host with no secrets or network trust.
