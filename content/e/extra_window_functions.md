---
title: "extra_window_functions"
linkTitle: "extra_window_functions"
description: "Extra Window Functions for PostgreSQL"
weight: 4720
categories: ["FUNC"]
width: full
---

[**extra_window_functions**](https://github.com/xocolatl/extra_window_functions) : Extra Window Functions for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4720** | {{< badge content="extra_window_functions" link="https://github.com/xocolatl/extra_window_functions" >}} | {{< ext "extra_window_functions" >}} | `1.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "permuteseq" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "topn" >}} {{< ext "quantile" >}} |

> [!Note] no pg14,13,12 on el8/9


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `extra_window_functions` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0` | {{< bg "18" "extra_window_functions_18" "green" >}} {{< bg "17" "extra_window_functions_17" "green" >}} {{< bg "16" "extra_window_functions_16" "green" >}} {{< bg "15" "extra_window_functions_15" "green" >}} {{< bg "14" "extra_window_functions_14" "green" >}} {{< bg "13" "extra_window_functions_13" "green" >}} | `extra_window_functions_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0` | {{< bg "18" "postgresql-18-extra-window-functions" "green" >}} {{< bg "17" "postgresql-17-extra-window-functions" "green" >}} {{< bg "16" "postgresql-16-extra-window-functions" "green" >}} {{< bg "15" "postgresql-15-extra-window-functions" "green" >}} {{< bg "14" "postgresql-14-extra-window-functions" "green" >}} {{< bg "13" "postgresql-13-extra-window-functions" "green" >}} | `postgresql-$v-extra-window-functions` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0" "extra_window_functions_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "extra_window_functions_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0" "extra_window_functions_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0" "extra_window_functions_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_15 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "extra_window_functions_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "extra_window_functions_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0" "extra_window_functions_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0" "extra_window_functions_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0" "extra_window_functions_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "extra_window_functions_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-extra-window-functions : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-extra-window-functions : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-extra-window-functions : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-extra-window-functions : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-extra-window-functions : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-extra-window-functions : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.0" "postgresql-18-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-extra-window-functions : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.0" "postgresql-18-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-17-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-16-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-15-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-14-extra-window-functions : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "postgresql-13-extra-window-functions : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `extra_window_functions_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.5 KiB | [extra_window_functions_18-1.0-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/extra_window_functions_18-1.0-6PGDG.rhel8.x86_64.rpm) |
| `extra_window_functions_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.3 KiB | [extra_window_functions_18-1.0-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/extra_window_functions_18-1.0-6PGDG.rhel8.aarch64.rpm) |
| `extra_window_functions_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.8 KiB | [extra_window_functions_18-1.0-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/extra_window_functions_18-1.0-6PGDG.rhel9.x86_64.rpm) |
| `extra_window_functions_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.4 KiB | [extra_window_functions_18-1.0-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/extra_window_functions_18-1.0-6PGDG.rhel9.aarch64.rpm) |
| `extra_window_functions_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.2 KiB | [extra_window_functions_18-1.0-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/extra_window_functions_18-1.0-6PGDG.rhel10.x86_64.rpm) |
| `extra_window_functions_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 25.1 KiB | [extra_window_functions_18-1.0-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/extra_window_functions_18-1.0-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-extra-window-functions` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 15.8 KiB | [postgresql-18-extra-window-functions_1.0-7.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-18-extra-window-functions_1.0-7.pgdg12+1_amd64.deb) |
| `postgresql-18-extra-window-functions` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 15.9 KiB | [postgresql-18-extra-window-functions_1.0-7.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-18-extra-window-functions_1.0-7.pgdg12+1_arm64.deb) |
| `postgresql-18-extra-window-functions` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 15.8 KiB | [postgresql-18-extra-window-functions_1.0-7.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-18-extra-window-functions_1.0-7.pgdg13+1_amd64.deb) |
| `postgresql-18-extra-window-functions` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 15.9 KiB | [postgresql-18-extra-window-functions_1.0-7.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-18-extra-window-functions_1.0-7.pgdg13+1_arm64.deb) |
| `postgresql-18-extra-window-functions` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 15.5 KiB | [postgresql-18-extra-window-functions_1.0-7.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-18-extra-window-functions_1.0-7.pgdg22.04+1_amd64.deb) |
| `postgresql-18-extra-window-functions` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 15.6 KiB | [postgresql-18-extra-window-functions_1.0-7.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-18-extra-window-functions_1.0-7.pgdg22.04+1_arm64.deb) |
| `postgresql-18-extra-window-functions` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 15.8 KiB | [postgresql-18-extra-window-functions_1.0-7.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-18-extra-window-functions_1.0-7.pgdg24.04+1_amd64.deb) |
| `postgresql-18-extra-window-functions` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 15.9 KiB | [postgresql-18-extra-window-functions_1.0-7.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-18-extra-window-functions_1.0-7.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `extra_window_functions_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [extra_window_functions_17-1.0-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/extra_window_functions_17-1.0-5PGDG.rhel8.x86_64.rpm) |
| `extra_window_functions_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.2 KiB | [extra_window_functions_17-1.0-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/extra_window_functions_17-1.0-5PGDG.rhel8.aarch64.rpm) |
| `extra_window_functions_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.7 KiB | [extra_window_functions_17-1.0-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/extra_window_functions_17-1.0-5PGDG.rhel9.x86_64.rpm) |
| `extra_window_functions_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.5 KiB | [extra_window_functions_17-1.0-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/extra_window_functions_17-1.0-5PGDG.rhel9.aarch64.rpm) |
| `extra_window_functions_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.1 KiB | [extra_window_functions_17-1.0-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/extra_window_functions_17-1.0-6PGDG.rhel10.x86_64.rpm) |
| `extra_window_functions_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 25.1 KiB | [extra_window_functions_17-1.0-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/extra_window_functions_17-1.0-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-extra-window-functions` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 15.8 KiB | [postgresql-17-extra-window-functions_1.0-7.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-17-extra-window-functions_1.0-7.pgdg12+1_amd64.deb) |
| `postgresql-17-extra-window-functions` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 15.8 KiB | [postgresql-17-extra-window-functions_1.0-7.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-17-extra-window-functions_1.0-7.pgdg12+1_arm64.deb) |
| `postgresql-17-extra-window-functions` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 15.8 KiB | [postgresql-17-extra-window-functions_1.0-7.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-17-extra-window-functions_1.0-7.pgdg13+1_amd64.deb) |
| `postgresql-17-extra-window-functions` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 15.9 KiB | [postgresql-17-extra-window-functions_1.0-7.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-17-extra-window-functions_1.0-7.pgdg13+1_arm64.deb) |
| `postgresql-17-extra-window-functions` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 15.6 KiB | [postgresql-17-extra-window-functions_1.0-7.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-17-extra-window-functions_1.0-7.pgdg22.04+1_amd64.deb) |
| `postgresql-17-extra-window-functions` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 15.7 KiB | [postgresql-17-extra-window-functions_1.0-7.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-17-extra-window-functions_1.0-7.pgdg22.04+1_arm64.deb) |
| `postgresql-17-extra-window-functions` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 15.9 KiB | [postgresql-17-extra-window-functions_1.0-7.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-17-extra-window-functions_1.0-7.pgdg24.04+1_amd64.deb) |
| `postgresql-17-extra-window-functions` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 15.9 KiB | [postgresql-17-extra-window-functions_1.0-7.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-17-extra-window-functions_1.0-7.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `extra_window_functions_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [extra_window_functions_16-1.0-3.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/extra_window_functions_16-1.0-3.rhel8.1.x86_64.rpm) |
| `extra_window_functions_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.0 KiB | [extra_window_functions_16-1.0-3.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/extra_window_functions_16-1.0-3.rhel8.1.aarch64.rpm) |
| `extra_window_functions_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.4 KiB | [extra_window_functions_16-1.0-3.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/extra_window_functions_16-1.0-3.rhel9.1.x86_64.rpm) |
| `extra_window_functions_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.2 KiB | [extra_window_functions_16-1.0-3.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/extra_window_functions_16-1.0-3.rhel9.1.aarch64.rpm) |
| `extra_window_functions_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.1 KiB | [extra_window_functions_16-1.0-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/extra_window_functions_16-1.0-6PGDG.rhel10.x86_64.rpm) |
| `extra_window_functions_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 25.1 KiB | [extra_window_functions_16-1.0-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/extra_window_functions_16-1.0-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-extra-window-functions` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 15.8 KiB | [postgresql-16-extra-window-functions_1.0-7.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-16-extra-window-functions_1.0-7.pgdg12+1_amd64.deb) |
| `postgresql-16-extra-window-functions` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 15.8 KiB | [postgresql-16-extra-window-functions_1.0-7.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-16-extra-window-functions_1.0-7.pgdg12+1_arm64.deb) |
| `postgresql-16-extra-window-functions` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 15.8 KiB | [postgresql-16-extra-window-functions_1.0-7.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-16-extra-window-functions_1.0-7.pgdg13+1_amd64.deb) |
| `postgresql-16-extra-window-functions` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 15.9 KiB | [postgresql-16-extra-window-functions_1.0-7.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-16-extra-window-functions_1.0-7.pgdg13+1_arm64.deb) |
| `postgresql-16-extra-window-functions` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 15.6 KiB | [postgresql-16-extra-window-functions_1.0-7.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-16-extra-window-functions_1.0-7.pgdg22.04+1_amd64.deb) |
| `postgresql-16-extra-window-functions` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 15.7 KiB | [postgresql-16-extra-window-functions_1.0-7.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-16-extra-window-functions_1.0-7.pgdg22.04+1_arm64.deb) |
| `postgresql-16-extra-window-functions` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 15.9 KiB | [postgresql-16-extra-window-functions_1.0-7.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-16-extra-window-functions_1.0-7.pgdg24.04+1_amd64.deb) |
| `postgresql-16-extra-window-functions` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 15.9 KiB | [postgresql-16-extra-window-functions_1.0-7.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-16-extra-window-functions_1.0-7.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `extra_window_functions_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [extra_window_functions_15-1.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/extra_window_functions_15-1.0-2.rhel8.x86_64.rpm) |
| `extra_window_functions_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.0 KiB | [extra_window_functions_15-1.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/extra_window_functions_15-1.0-2.rhel8.aarch64.rpm) |
| `extra_window_functions_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.8 KiB | [extra_window_functions_15-1.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/extra_window_functions_15-1.0-2.rhel9.x86_64.rpm) |
| `extra_window_functions_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.6 KiB | [extra_window_functions_15-1.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/extra_window_functions_15-1.0-2.rhel9.aarch64.rpm) |
| `extra_window_functions_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.1 KiB | [extra_window_functions_15-1.0-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/extra_window_functions_15-1.0-6PGDG.rhel10.x86_64.rpm) |
| `extra_window_functions_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 25.1 KiB | [extra_window_functions_15-1.0-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/extra_window_functions_15-1.0-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-extra-window-functions` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 15.8 KiB | [postgresql-15-extra-window-functions_1.0-7.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-15-extra-window-functions_1.0-7.pgdg12+1_amd64.deb) |
| `postgresql-15-extra-window-functions` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 15.8 KiB | [postgresql-15-extra-window-functions_1.0-7.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-15-extra-window-functions_1.0-7.pgdg12+1_arm64.deb) |
| `postgresql-15-extra-window-functions` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 15.8 KiB | [postgresql-15-extra-window-functions_1.0-7.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-15-extra-window-functions_1.0-7.pgdg13+1_amd64.deb) |
| `postgresql-15-extra-window-functions` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 15.9 KiB | [postgresql-15-extra-window-functions_1.0-7.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-15-extra-window-functions_1.0-7.pgdg13+1_arm64.deb) |
| `postgresql-15-extra-window-functions` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 15.6 KiB | [postgresql-15-extra-window-functions_1.0-7.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-15-extra-window-functions_1.0-7.pgdg22.04+1_amd64.deb) |
| `postgresql-15-extra-window-functions` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 15.7 KiB | [postgresql-15-extra-window-functions_1.0-7.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-15-extra-window-functions_1.0-7.pgdg22.04+1_arm64.deb) |
| `postgresql-15-extra-window-functions` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 15.9 KiB | [postgresql-15-extra-window-functions_1.0-7.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-15-extra-window-functions_1.0-7.pgdg24.04+1_amd64.deb) |
| `postgresql-15-extra-window-functions` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 15.9 KiB | [postgresql-15-extra-window-functions_1.0-7.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-15-extra-window-functions_1.0-7.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `extra_window_functions_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [extra_window_functions_14-1.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/extra_window_functions_14-1.0-2.rhel8.x86_64.rpm) |
| `extra_window_functions_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.9 KiB | [extra_window_functions_14-1.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/extra_window_functions_14-1.0-2.rhel8.aarch64.rpm) |
| `extra_window_functions_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.5 KiB | [extra_window_functions_14-1.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/extra_window_functions_14-1.0-2.rhel9.aarch64.rpm) |
| `extra_window_functions_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 25.1 KiB | [extra_window_functions_14-1.0-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/extra_window_functions_14-1.0-6PGDG.rhel10.x86_64.rpm) |
| `extra_window_functions_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 25.1 KiB | [extra_window_functions_14-1.0-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/extra_window_functions_14-1.0-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-extra-window-functions` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 15.7 KiB | [postgresql-14-extra-window-functions_1.0-7.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-14-extra-window-functions_1.0-7.pgdg12+1_amd64.deb) |
| `postgresql-14-extra-window-functions` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 15.8 KiB | [postgresql-14-extra-window-functions_1.0-7.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-14-extra-window-functions_1.0-7.pgdg12+1_arm64.deb) |
| `postgresql-14-extra-window-functions` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 15.7 KiB | [postgresql-14-extra-window-functions_1.0-7.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-14-extra-window-functions_1.0-7.pgdg13+1_amd64.deb) |
| `postgresql-14-extra-window-functions` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 15.8 KiB | [postgresql-14-extra-window-functions_1.0-7.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-14-extra-window-functions_1.0-7.pgdg13+1_arm64.deb) |
| `postgresql-14-extra-window-functions` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 15.7 KiB | [postgresql-14-extra-window-functions_1.0-7.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-14-extra-window-functions_1.0-7.pgdg22.04+1_amd64.deb) |
| `postgresql-14-extra-window-functions` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 15.7 KiB | [postgresql-14-extra-window-functions_1.0-7.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-14-extra-window-functions_1.0-7.pgdg22.04+1_arm64.deb) |
| `postgresql-14-extra-window-functions` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 15.9 KiB | [postgresql-14-extra-window-functions_1.0-7.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-14-extra-window-functions_1.0-7.pgdg24.04+1_amd64.deb) |
| `postgresql-14-extra-window-functions` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 15.8 KiB | [postgresql-14-extra-window-functions_1.0-7.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-14-extra-window-functions_1.0-7.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `extra_window_functions_13` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.8 KiB | [extra_window_functions_13-1.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/extra_window_functions_13-1.0-2.rhel8.aarch64.rpm) |
| `extra_window_functions_13` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.3 KiB | [extra_window_functions_13-1.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/extra_window_functions_13-1.0-2.rhel9.aarch64.rpm) |
| `extra_window_functions_13` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.9 KiB | [extra_window_functions_13-1.0-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/extra_window_functions_13-1.0-6PGDG.rhel10.x86_64.rpm) |
| `extra_window_functions_13` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.9 KiB | [extra_window_functions_13-1.0-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/extra_window_functions_13-1.0-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-extra-window-functions` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 15.6 KiB | [postgresql-13-extra-window-functions_1.0-7.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-13-extra-window-functions_1.0-7.pgdg12+1_amd64.deb) |
| `postgresql-13-extra-window-functions` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 15.6 KiB | [postgresql-13-extra-window-functions_1.0-7.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-13-extra-window-functions_1.0-7.pgdg12+1_arm64.deb) |
| `postgresql-13-extra-window-functions` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 15.5 KiB | [postgresql-13-extra-window-functions_1.0-7.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-13-extra-window-functions_1.0-7.pgdg13+1_amd64.deb) |
| `postgresql-13-extra-window-functions` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 15.6 KiB | [postgresql-13-extra-window-functions_1.0-7.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-13-extra-window-functions_1.0-7.pgdg13+1_arm64.deb) |
| `postgresql-13-extra-window-functions` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 15.6 KiB | [postgresql-13-extra-window-functions_1.0-7.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-13-extra-window-functions_1.0-7.pgdg22.04+1_amd64.deb) |
| `postgresql-13-extra-window-functions` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 15.7 KiB | [postgresql-13-extra-window-functions_1.0-7.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-13-extra-window-functions_1.0-7.pgdg22.04+1_arm64.deb) |
| `postgresql-13-extra-window-functions` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 15.7 KiB | [postgresql-13-extra-window-functions_1.0-7.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-13-extra-window-functions_1.0-7.pgdg24.04+1_amd64.deb) |
| `postgresql-13-extra-window-functions` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 15.6 KiB | [postgresql-13-extra-window-functions_1.0-7.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/e/extra-window-functions/postgresql-13-extra-window-functions_1.0-7.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/xocolatl/extra_window_functions" title="Repository" icon="github" subtitle="github.com/xocolatl/extra_window_functions" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install extra_window_functions;		# install via package name, for the active PG version

pig install extra_window_functions -v 18;   # install for PG 18
pig install extra_window_functions -v 17;   # install for PG 17
pig install extra_window_functions -v 16;   # install for PG 16
pig install extra_window_functions -v 15;   # install for PG 15
pig install extra_window_functions -v 14;   # install for PG 14
pig install extra_window_functions -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION extra_window_functions;
```
