---
title: "index_advisor"
linkTitle: "index_advisor"
description: "Query index advisor"
weight: 2840
categories: ["FEAT"]
width: full
---

[**index_advisor**](https://github.com/supabase/index_advisor) : Query index advisor


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2840** | {{< badge content="index_advisor" link="https://github.com/supabase/index_advisor" >}} | {{< ext "index_advisor" >}} | `0.2.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hypopg" >}} {{< ext "pg_qualstats" >}} {{< ext "powa" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_hint_plan" >}} {{< ext "auto_explain" >}} {{< ext "pg_profile" >}} {{< ext "pg_show_plans" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `index_advisor` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "index_advisor_18" "green" >}} {{< bg "17" "index_advisor_17" "green" >}} {{< bg "16" "index_advisor_16" "green" >}} {{< bg "15" "index_advisor_15" "green" >}} {{< bg "14" "index_advisor_14" "green" >}} {{< bg "13" "index_advisor_13" "green" >}} | `index_advisor_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "postgresql-18-index-advisor" "green" >}} {{< bg "17" "postgresql-17-index-advisor" "green" >}} {{< bg "16" "postgresql-16-index-advisor" "green" >}} {{< bg "15" "postgresql-15-index-advisor" "green" >}} {{< bg "14" "postgresql-14-index-advisor" "green" >}} {{< bg "13" "postgresql-13-index-advisor" "green" >}} | `postgresql-$v-index-advisor` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "index_advisor_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-13-index-advisor : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-13-index-advisor : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-13-index-advisor : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-13-index-advisor : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-13-index-advisor : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-13-index-advisor : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-13-index-advisor : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-index-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-13-index-advisor : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `index_advisor_13` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.2 KiB | [index_advisor_13-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/index_advisor_13-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `index_advisor_13` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.2 KiB | [index_advisor_13-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/index_advisor_13-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `index_advisor_13` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.1 KiB | [index_advisor_13-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/index_advisor_13-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `index_advisor_13` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.1 KiB | [index_advisor_13-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/index_advisor_13-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `index_advisor_13` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.2 KiB | [index_advisor_13-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/index_advisor_13-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `index_advisor_13` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.2 KiB | [index_advisor_13-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/index_advisor_13-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-index-advisor` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-index-advisor` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-index-advisor` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-index-advisor` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-index-advisor` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-index-advisor` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-index-advisor` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-index-advisor` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.6 KiB | [postgresql-13-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/index-advisor/postgresql-13-index-advisor_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/index_advisor" title="Repository" icon="github" subtitle="github.com/supabase/index_advisor" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="index_advisor-0.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg index_advisor;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install index_advisor;		# install via package name, for the active PG version

pig install index_advisor -v 18;   # install for PG 18
pig install index_advisor -v 17;   # install for PG 17
pig install index_advisor -v 16;   # install for PG 16
pig install index_advisor -v 15;   # install for PG 15
pig install index_advisor -v 14;   # install for PG 14
pig install index_advisor -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION index_advisor;
```
