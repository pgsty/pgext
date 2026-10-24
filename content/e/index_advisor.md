---
title: "index_advisor"
linkTitle: "index_advisor"
description: "Query index advisor"
weight: 2800
categories: ["FEAT"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**index_advisor**](https://github.com/supabase/index_advisor) : Query index advisor


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2800** | {{< badge content="index_advisor" link="https://github.com/supabase/index_advisor" >}} | {{< ext "index_advisor" >}} | `0.2.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "hypopg" >}} |
|   **See Also**    | {{< ext "hypopg" >}} {{< ext "online_advisor" >}} {{< ext "pg_qualstats" >}} {{< ext "pg_hint_plan" >}} {{< ext "pg_stat_plans" >}} {{< ext "plan_filter" >}} {{< ext "pg_track_optimizer" >}} {{< ext "pg_strom" >}} {{< ext "pg_orca" >}} {{< ext "pg_relusage" >}} |

> [!Note] Upstream 0.2.1; pkg 0.2.0; planner estimates only.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `index_advisor` | `hypopg` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "index_advisor_18" "green" >}} {{< bg "17" "index_advisor_17" "green" >}} {{< bg "16" "index_advisor_16" "green" >}} {{< bg "15" "index_advisor_15" "green" >}} {{< bg "14" "index_advisor_14" "green" >}} | `index_advisor_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "postgresql-18-index-advisor" "green" >}} {{< bg "17" "postgresql-17-index-advisor" "green" >}} {{< bg "16" "postgresql-16-index-advisor" "green" >}} {{< bg "15" "postgresql-15-index-advisor" "green" >}} {{< bg "14" "postgresql-14-index-advisor" "green" >}} | `postgresql-$v-index-advisor` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `index_advisor_18` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.2 KiB | [index_advisor_18-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_18-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `index_advisor_18` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.2 KiB | [index_advisor_18-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_18-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `index_advisor_18` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.1 KiB | [index_advisor_18-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_18-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `index_advisor_18` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.1 KiB | [index_advisor_18-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_18-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `index_advisor_18` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.2 KiB | [index_advisor_18-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/index_advisor_18-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `index_advisor_18` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.2 KiB | [index_advisor_18-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/index_advisor_18-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-index-advisor` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.6 KiB | [postgresql-18-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-18-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-index-advisor` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.6 KiB | [postgresql-18-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-18-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-index-advisor` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.6 KiB | [postgresql-18-index-advisor_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/index-advisor/postgresql-18-index-advisor_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-index-advisor` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.6 KiB | [postgresql-18-index-advisor_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/index-advisor/postgresql-18-index-advisor_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-index-advisor` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.6 KiB | [postgresql-18-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-18-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-index-advisor` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.6 KiB | [postgresql-18-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-18-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-index-advisor` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.6 KiB | [postgresql-18-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-18-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-index-advisor` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.6 KiB | [postgresql-18-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-18-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-index-advisor` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.6 KiB | [postgresql-18-index-advisor_0.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/index-advisor/postgresql-18-index-advisor_0.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-index-advisor` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.6 KiB | [postgresql-18-index-advisor_0.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/index-advisor/postgresql-18-index-advisor_0.2.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `index_advisor_17` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.2 KiB | [index_advisor_17-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_17-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `index_advisor_17` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.2 KiB | [index_advisor_17-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_17-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `index_advisor_17` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.1 KiB | [index_advisor_17-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_17-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `index_advisor_17` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.1 KiB | [index_advisor_17-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_17-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `index_advisor_17` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.2 KiB | [index_advisor_17-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/index_advisor_17-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `index_advisor_17` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.2 KiB | [index_advisor_17-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/index_advisor_17-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-index-advisor` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-index-advisor` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-index-advisor` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-index-advisor` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-index-advisor` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-index-advisor` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-index-advisor` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-index-advisor` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-index-advisor` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-index-advisor` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.6 KiB | [postgresql-17-index-advisor_0.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/index-advisor/postgresql-17-index-advisor_0.2.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `index_advisor_16` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.2 KiB | [index_advisor_16-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_16-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `index_advisor_16` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.2 KiB | [index_advisor_16-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_16-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `index_advisor_16` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.1 KiB | [index_advisor_16-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_16-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `index_advisor_16` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.1 KiB | [index_advisor_16-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_16-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `index_advisor_16` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.2 KiB | [index_advisor_16-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/index_advisor_16-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `index_advisor_16` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.2 KiB | [index_advisor_16-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/index_advisor_16-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-index-advisor` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-index-advisor` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-index-advisor` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-index-advisor` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-index-advisor` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-index-advisor` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-index-advisor` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-index-advisor` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-index-advisor` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-index-advisor` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.6 KiB | [postgresql-16-index-advisor_0.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/index-advisor/postgresql-16-index-advisor_0.2.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `index_advisor_15` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.2 KiB | [index_advisor_15-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_15-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `index_advisor_15` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.2 KiB | [index_advisor_15-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_15-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `index_advisor_15` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.1 KiB | [index_advisor_15-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_15-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `index_advisor_15` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.1 KiB | [index_advisor_15-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_15-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `index_advisor_15` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.2 KiB | [index_advisor_15-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/index_advisor_15-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `index_advisor_15` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.2 KiB | [index_advisor_15-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/index_advisor_15-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-index-advisor` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-index-advisor` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-index-advisor` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-index-advisor` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-index-advisor` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-index-advisor` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-index-advisor` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-index-advisor` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-index-advisor` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-index-advisor` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.6 KiB | [postgresql-15-index-advisor_0.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/index-advisor/postgresql-15-index-advisor_0.2.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `index_advisor_14` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.2 KiB | [index_advisor_14-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_14-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `index_advisor_14` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.2 KiB | [index_advisor_14-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_14-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `index_advisor_14` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.1 KiB | [index_advisor_14-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_14-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `index_advisor_14` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.1 KiB | [index_advisor_14-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_14-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `index_advisor_14` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.2 KiB | [index_advisor_14-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/index_advisor_14-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `index_advisor_14` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.2 KiB | [index_advisor_14-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/index_advisor_14-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-index-advisor` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-index-advisor` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-index-advisor` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-index-advisor` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-index-advisor` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-index-advisor` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-index-advisor` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-index-advisor` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-index-advisor` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-index-advisor` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.6 KiB | [postgresql-14-index-advisor_0.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/index-advisor/postgresql-14-index-advisor_0.2.0-1PIGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/index_advisor" title="Repository" icon="github" subtitle="github.com/supabase/index_advisor" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="index_advisor-0.2.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg index_advisor;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install index_advisor;		# install via package name, for the active PG version

pig install index_advisor -v 18;   # install for PG 18
pig install index_advisor -v 17;   # install for PG 17
pig install index_advisor -v 16;   # install for PG 16
pig install index_advisor -v 15;   # install for PG 15
pig install index_advisor -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION index_advisor CASCADE; -- requires hypopg
```

## Usage

Sources:

- [Official source](https://github.com/supabase/index_advisor/blob/b41d1f5f42f8909091edd6752d1f5006d6535db3/README.md)
- [Official source](https://github.com/supabase/index_advisor/blob/b41d1f5f42f8909091edd6752d1f5006d6535db3/index_advisor.control)
- [Official source](https://github.com/supabase/index_advisor/blob/b41d1f5f42f8909091edd6752d1f5006d6535db3/index_advisor--0.2.0--0.2.1.sql)

`index_advisor` 0.2.1 uses HypoPG to compare planner costs with hypothetical indexes. It returns recommendations, not measured execution-time guarantees, and does not create the recommended physical indexes.

### Core Workflow

Install the required `hypopg` extension first, with the privileges its installation requires. This SQL extension needs no preload or restart.

```sql
CREATE EXTENSION hypopg;
CREATE EXTENSION index_advisor;
CREATE TABLE book (id bigint PRIMARY KEY, title text NOT NULL);
SELECT * FROM index_advisor('SELECT id FROM book WHERE title = $1');
```

`index_advisor(query text)` accepts a parameterized query and returns before/after startup and total cost values, `index_statements` and `errors`. Inspect the suggested DDL and real workload before creating an index. View and materialized-view support depends on what the advisor can resolve to underlying tables.

### 0.2.1 Corrections and Limits

Version 0.2.1 fixes qualification of HypoPG calls when HypoPG uses another schema, non-greedy removal of comments, and cleanup of the advisor's own prepared statement. Update installed databases with `ALTER EXTENSION index_advisor UPDATE TO '0.2.1'` after installing the new SQL files.

Run analysis with deliberate table permissions and representative statistics. Planning cost does not account for every operational cost: additional indexes consume disk and slow writes, and should be validated against production query shapes before deployment.
