---
title: "pg_strict"
linkTitle: "pg_strict"
description: "Prevent dangerous UPDATE and DELETE without WHERE clause"
weight: 5830
categories: ["ADMIN"]
width: full
---

[**pg_strict**](https://github.com/spa5k/pg_strict) : Prevent dangerous UPDATE and DELETE without WHERE clause


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5830** | {{< badge content="pg_strict" link="https://github.com/spa5k/pg_strict" >}} | {{< ext "pg_strict" >}} | `1.0.2` | {{< category "ADMIN" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "safeupdate" >}} {{< ext "pg_savior" >}} {{< ext "pg_upless" >}} {{< ext "pg_drop_events" >}} {{< ext "pg_readonly" >}} {{< ext "table_log" >}} {{< ext "pgaudit" >}} {{< ext "pg_permissions" >}} |

> [!Note] manually patched


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pg_strict` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.2` | {{< bg "18" "pg_strict_18" "green" >}} {{< bg "17" "pg_strict_17" "green" >}} {{< bg "16" "pg_strict_16" "green" >}} {{< bg "15" "pg_strict_15" "green" >}} {{< bg "14" "pg_strict_14" "green" >}} {{< bg "13" "pg_strict_13" "red" >}} | `pg_strict_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.2` | {{< bg "18" "pg_strict_18" "green" >}} {{< bg "17" "pg_strict_17" "green" >}} {{< bg "16" "pg_strict_16" "green" >}} {{< bg "15" "pg_strict_15" "green" >}} {{< bg "14" "pg_strict_14" "green" >}} {{< bg "13" "pg_strict_13" "red" >}} | `pg_strict_$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_strict_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_strict_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_strict_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_strict_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "pg_strict_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_13 : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "pg_strict_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_13 : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "pg_strict_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_13 : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "pg_strict_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_13 : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "pg_strict_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_13 : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "pg_strict_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_13 : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "pg_strict_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_13 : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "pg_strict_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_strict_13 : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strict_18` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 324.1 KiB | [pg_strict_18-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_strict_18-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_strict_18` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 217.9 KiB | [pg_strict_18-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_strict_18-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_strict_18` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 338.9 KiB | [pg_strict_18-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_strict_18-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_strict_18` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 232.5 KiB | [pg_strict_18-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_strict_18-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_strict_18` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 339.1 KiB | [pg_strict_18-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_strict_18-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_strict_18` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 232.6 KiB | [pg_strict_18-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_strict_18-1.0.2-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strict_17` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 324.4 KiB | [pg_strict_17-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_strict_17-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_strict_17` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 217.8 KiB | [pg_strict_17-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_strict_17-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_strict_17` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 339.0 KiB | [pg_strict_17-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_strict_17-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_strict_17` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 232.8 KiB | [pg_strict_17-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_strict_17-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_strict_17` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 339.3 KiB | [pg_strict_17-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_strict_17-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_strict_17` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 232.3 KiB | [pg_strict_17-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_strict_17-1.0.2-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strict_16` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 324.5 KiB | [pg_strict_16-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_strict_16-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_strict_16` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 217.8 KiB | [pg_strict_16-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_strict_16-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_strict_16` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 339.1 KiB | [pg_strict_16-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_strict_16-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_strict_16` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 232.8 KiB | [pg_strict_16-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_strict_16-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_strict_16` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 339.2 KiB | [pg_strict_16-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_strict_16-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_strict_16` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 232.7 KiB | [pg_strict_16-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_strict_16-1.0.2-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strict_15` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 323.7 KiB | [pg_strict_15-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_strict_15-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_strict_15` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 217.8 KiB | [pg_strict_15-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_strict_15-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_strict_15` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 338.9 KiB | [pg_strict_15-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_strict_15-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_strict_15` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 232.8 KiB | [pg_strict_15-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_strict_15-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_strict_15` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 339.2 KiB | [pg_strict_15-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_strict_15-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_strict_15` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 232.4 KiB | [pg_strict_15-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_strict_15-1.0.2-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strict_14` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 323.8 KiB | [pg_strict_14-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_strict_14-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_strict_14` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 217.8 KiB | [pg_strict_14-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_strict_14-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_strict_14` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 338.6 KiB | [pg_strict_14-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_strict_14-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_strict_14` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 232.7 KiB | [pg_strict_14-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_strict_14-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_strict_14` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 338.8 KiB | [pg_strict_14-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_strict_14-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_strict_14` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 232.7 KiB | [pg_strict_14-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_strict_14-1.0.2-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_strict_13` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 324.1 KiB | [pg_strict_13-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_strict_13-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_strict_13` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 339.2 KiB | [pg_strict_13-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_strict_13-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_strict_13` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 339.4 KiB | [pg_strict_13-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_strict_13-1.0.2-1PIGSTY.el10.x86_64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/spa5k/pg_strict" title="Repository" icon="github" subtitle="github.com/spa5k/pg_strict" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_strict-1.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_strict;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_strict;		# install via package name, for the active PG version

pig install pg_strict -v 18;   # install for PG 18
pig install pg_strict -v 17;   # install for PG 17
pig install pg_strict -v 16;   # install for PG 16
pig install pg_strict -v 15;   # install for PG 15
pig install pg_strict -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_strict';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_strict;
```
