---
title: "cat_tools"
linkTitle: "cat_tools"
description: "Tools for interfacing with the PostgreSQL catalog"
weight: 5290
categories: ["ADMIN"]
languages: ["SQL"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**cat_tools**](https://github.com/Postgres-Extensions/cat_tools) : Tools for interfacing with the PostgreSQL catalog


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5290** | {{< badge content="cat_tools" link="https://github.com/Postgres-Extensions/cat_tools" >}} | {{< ext "cat_tools" >}} | `0.3.0` | {{< category "ADMIN" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `cat_tools` |
|   **Requires**    | {{< ext "plpgsql" >}} |
|    **Need By**    | {{< ext "extension_drop" >}} {{< ext "object_reference" >}} |
|   **See Also**    | {{< ext "pg_catalog_get_defs" >}} {{< ext "pg_global_catalog" >}} {{< ext "meta_triggers" >}} {{< ext "pg_catcheck" >}} {{< ext "pgdd" >}} {{< ext "ddlx" >}} {{< ext "meta" >}} {{< ext "object_reference" >}} |

> [!Note] Promoted from a source-only universe row to PIGSTY RPM and DEB packages at 0.3.0; control fixes schema cat_tools and META declares the plpgsql runtime dependency.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `cat_tools` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "cat_tools_18" "green" >}} {{< bg "17" "cat_tools_17" "green" >}} {{< bg "16" "cat_tools_16" "green" >}} {{< bg "15" "cat_tools_15" "green" >}} {{< bg "14" "cat_tools_14" "green" >}} | `cat_tools_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-cat-tools" "green" >}} {{< bg "17" "postgresql-17-cat-tools" "green" >}} {{< bg "16" "postgresql-16-cat-tools" "green" >}} {{< bg "15" "postgresql-15-cat-tools" "green" >}} {{< bg "14" "postgresql-14-cat-tools" "green" >}} | `postgresql-$v-cat-tools` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "cat_tools_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-cat-tools : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-cat-tools : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cat_tools_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.0 KiB | [cat_tools_18-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cat_tools_18-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.0 KiB | [cat_tools_18-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cat_tools_18-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.7 KiB | [cat_tools_18-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cat_tools_18-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.6 KiB | [cat_tools_18-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cat_tools_18-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.8 KiB | [cat_tools_18-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cat_tools_18-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `cat_tools_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.8 KiB | [cat_tools_18-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cat_tools_18-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-cat-tools` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.1 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.1 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.1 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.1 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.7 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.7 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.6 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.6 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.6 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
| `postgresql-18-cat-tools` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.6 KiB | [postgresql-18-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-18-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cat_tools_17` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.0 KiB | [cat_tools_17-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cat_tools_17-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_17` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.0 KiB | [cat_tools_17-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cat_tools_17-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_17` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.7 KiB | [cat_tools_17-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cat_tools_17-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_17` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.6 KiB | [cat_tools_17-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cat_tools_17-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_17` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.8 KiB | [cat_tools_17-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cat_tools_17-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `cat_tools_17` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.8 KiB | [cat_tools_17-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cat_tools_17-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-cat-tools` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.1 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.1 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.1 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.1 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.7 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.7 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.6 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.6 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.6 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
| `postgresql-17-cat-tools` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.6 KiB | [postgresql-17-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-17-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cat_tools_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.0 KiB | [cat_tools_16-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cat_tools_16-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.0 KiB | [cat_tools_16-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cat_tools_16-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.7 KiB | [cat_tools_16-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cat_tools_16-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.6 KiB | [cat_tools_16-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cat_tools_16-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.8 KiB | [cat_tools_16-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cat_tools_16-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `cat_tools_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.8 KiB | [cat_tools_16-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cat_tools_16-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-cat-tools` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.1 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.1 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.1 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.1 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.7 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.7 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.6 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.6 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.6 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
| `postgresql-16-cat-tools` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.6 KiB | [postgresql-16-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-16-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cat_tools_15` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.0 KiB | [cat_tools_15-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cat_tools_15-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_15` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.0 KiB | [cat_tools_15-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cat_tools_15-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_15` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.7 KiB | [cat_tools_15-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cat_tools_15-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_15` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.6 KiB | [cat_tools_15-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cat_tools_15-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_15` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.8 KiB | [cat_tools_15-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cat_tools_15-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `cat_tools_15` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.8 KiB | [cat_tools_15-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cat_tools_15-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-cat-tools` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.1 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.1 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.1 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.1 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.7 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.7 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.6 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.6 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.6 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
| `postgresql-15-cat-tools` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.6 KiB | [postgresql-15-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-15-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cat_tools_14` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 35.0 KiB | [cat_tools_14-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cat_tools_14-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_14` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 35.0 KiB | [cat_tools_14-0.3.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cat_tools_14-0.3.0-1PIGSTY.el8.noarch.rpm) |
| `cat_tools_14` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 33.7 KiB | [cat_tools_14-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cat_tools_14-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_14` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 33.6 KiB | [cat_tools_14-0.3.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cat_tools_14-0.3.0-1PIGSTY.el9.noarch.rpm) |
| `cat_tools_14` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.8 KiB | [cat_tools_14-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cat_tools_14-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `cat_tools_14` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 33.8 KiB | [cat_tools_14-0.3.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cat_tools_14-0.3.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-14-cat-tools` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.1 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.1 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.1 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.1 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.7 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.7 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.6 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.6 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 27.6 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
| `postgresql-14-cat-tools` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 27.6 KiB | [postgresql-14-cat-tools_0.3.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cat-tools/postgresql-14-cat-tools_0.3.0-1PIGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Postgres-Extensions/cat_tools" title="Repository" icon="github" subtitle="github.com/Postgres-Extensions/cat_tools" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="cat_tools-0.3.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg cat_tools;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install cat_tools;		# install via package name, for the active PG version

pig install cat_tools -v 18;   # install for PG 18
pig install cat_tools -v 17;   # install for PG 17
pig install cat_tools -v 16;   # install for PG 16
pig install cat_tools -v 15;   # install for PG 15
pig install cat_tools -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION cat_tools CASCADE; -- requires plpgsql
```

## Usage

Sources:

- [cat_tools 0.3.0 README](https://github.com/Postgres-Extensions/cat_tools/blob/0.3.0/README.asc)
- [cat_tools 0.3.0 history](https://github.com/Postgres-Extensions/cat_tools/blob/0.3.0/HISTORY.asc)
- [cat_tools 0.3.0 control file](https://github.com/Postgres-Extensions/cat_tools/blob/0.3.0/cat_tools.control)
- [cat_tools 0.3.0 install SQL](https://github.com/Postgres-Extensions/cat_tools/blob/0.3.0/sql/cat_tools--0.3.0.sql.in)

`cat_tools` provides typed views, enums, and helper functions for PostgreSQL catalog introspection. It is designed for database code that needs a more stable and readable interface than repeatedly decoding raw `pg_catalog` fields; the views still track PostgreSQL's catalogs and must be reviewed across major-version upgrades.

### Install and Grant Access

```sql
CREATE EXTENSION cat_tools;
GRANT cat_tools__usage TO app_introspection;
```

The extension installs in the fixed `cat_tools` schema, requires `plpgsql`, and is not relocatable. Grant the `cat_tools__usage` role rather than exposing internal `_cat_tools` helpers directly.

### Inspect Relations and Columns

```sql
SELECT cat_tools.relation__kind(c.relkind::text)
FROM pg_catalog.pg_class AS c
WHERE c.oid = 'public.orders'::regclass;

SELECT cat_tools.relation__column_names('public.orders'::regclass);
SELECT cat_tools.pg_attribute__get('public.orders'::regclass, 'id');
```

Useful relation helpers include `pg_class(regclass)`, `relation__is_catalog`, `relation__is_temp`, `relation__kind`, and `relation__relkind`. Typed mapping functions make the one-character catalog codes explicit.

### Inspect Routines

Version 0.3 adds functions and types that cover both functions and procedures:

```sql
SELECT cat_tools.routine__arg_types(
  'public.calculate_total(integer, numeric)'::regprocedure
);

SELECT cat_tools.routine__parse_arg_names(
  'IN account_id integer, INOUT total numeric'
);
```

The routine surface includes `routine__parse_arg_types`, `routine__parse_arg_names`, `routine__arg_types`, `routine__arg_names`, their text variants, and mappings for routine kind, argument mode, volatility, and parallel safety. `function__arg_types` and `function__arg_types_text` are deprecated; use the routine parsers.

### Version 0.3.0 and Caveats

- Version 0.3.0 supports PostgreSQL 12-18+ upstream; current Pigsty packages cover PostgreSQL 14-18.
- The release corrects the `c`, `f`, and `m` mappings for composite types, foreign tables, and materialized views. Re-test any code that worked around the old mapping.
- Internal `_cat_tools` helpers now revoke `EXECUTE` from `PUBLIC`; callers should inherit `cat_tools__usage` and use the supported surface.
- The 0.2.3-to-0.3.0 update adds enum values and therefore cannot run on PostgreSQL 11 or earlier. Upgrade the database major version and extension in the order documented upstream.
- PostgreSQL does not promise catalog compatibility across major releases. Pin tests to every supported PostgreSQL major even when using these wrappers.
