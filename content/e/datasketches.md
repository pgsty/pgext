---
title: "datasketches"
linkTitle: "datasketches"
description: "Approximate analytics sketches and aggregates for PostgreSQL"
weight: 4700
categories: ["FUNC"]
languages: ["C++"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**datasketches**](https://github.com/apache/datasketches-postgresql) : Approximate analytics sketches and aggregates for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4700** | {{< badge content="datasketches" link="https://github.com/apache/datasketches-postgresql" >}} | {{< ext "datasketches" >}} | `1.7.0` | {{< category "FUNC" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "ddsketch" >}} {{< ext "count_distinct" >}} {{< ext "tdigest" >}} {{< ext "topn" >}} {{< ext "omnisketch" >}} {{< ext "hll" >}} |

> [!Note] DataSketches C++ core 5.2.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `datasketches` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "19" "datasketches_19" "green" >}} {{< bg "18" "datasketches_18" "green" >}} {{< bg "17" "datasketches_17" "green" >}} {{< bg "16" "datasketches_16" "green" >}} {{< bg "15" "datasketches_15" "green" >}} | `datasketches_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "19" "postgresql-19-datasketches" "green" >}} {{< bg "18" "postgresql-18-datasketches" "green" >}} {{< bg "17" "postgresql-17-datasketches" "green" >}} {{< bg "16" "postgresql-16-datasketches" "green" >}} {{< bg "15" "postgresql-15-datasketches" "green" >}} | `postgresql-$v-datasketches` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "datasketches_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.7.0" "datasketches_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_15 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "datasketches_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.7.0" "datasketches_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_15 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_19 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.7.0" "datasketches_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_15 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_19 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.7.0" "datasketches_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_19 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.7.0" "datasketches_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "datasketches_19 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.7.0" "datasketches_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.7.0" "datasketches_15 : AVAIL 4" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-19-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-19-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-19-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-19-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-19-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-19-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-19-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-19-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-19-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-19-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-datasketches : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-datasketches : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `datasketches_19` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 965.8 KiB | [datasketches_19-1.7.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/datasketches_19-1.7.0-2PGSTY.el9.x86_64.rpm) |
| `datasketches_19` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 979.3 KiB | [datasketches_19-1.7.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/datasketches_19-1.7.0-2PGSTY.el9.aarch64.rpm) |
| `datasketches_19` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 989.3 KiB | [datasketches_19-1.7.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/datasketches_19-1.7.0-2PGSTY.el10.x86_64.rpm) |
| `datasketches_19` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 994.4 KiB | [datasketches_19-1.7.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/datasketches_19-1.7.0-2PGSTY.el10.aarch64.rpm) |
| `postgresql-19-datasketches` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 918.8 KiB | [postgresql-19-datasketches_1.7.0-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-19-datasketches_1.7.0-2PGSTY~bookworm_amd64.deb) |
| `postgresql-19-datasketches` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 919.9 KiB | [postgresql-19-datasketches_1.7.0-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-19-datasketches_1.7.0-2PGSTY~bookworm_arm64.deb) |
| `postgresql-19-datasketches` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 944.1 KiB | [postgresql-19-datasketches_1.7.0-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-19-datasketches_1.7.0-2PGSTY~trixie_amd64.deb) |
| `postgresql-19-datasketches` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 945.5 KiB | [postgresql-19-datasketches_1.7.0-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-19-datasketches_1.7.0-2PGSTY~trixie_arm64.deb) |
| `postgresql-19-datasketches` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1017.0 KiB | [postgresql-19-datasketches_1.7.0-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-19-datasketches_1.7.0-2PGSTY~jammy_amd64.deb) |
| `postgresql-19-datasketches` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1021.7 KiB | [postgresql-19-datasketches_1.7.0-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-19-datasketches_1.7.0-2PGSTY~jammy_arm64.deb) |
| `postgresql-19-datasketches` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 979.6 KiB | [postgresql-19-datasketches_1.7.0-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-19-datasketches_1.7.0-2PGSTY~noble_amd64.deb) |
| `postgresql-19-datasketches` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 992.2 KiB | [postgresql-19-datasketches_1.7.0-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-19-datasketches_1.7.0-2PGSTY~noble_arm64.deb) |
| `postgresql-19-datasketches` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 947.7 KiB | [postgresql-19-datasketches_1.7.0-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/datasketches/postgresql-19-datasketches_1.7.0-2PGSTY~resolute_amd64.deb) |
| `postgresql-19-datasketches` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 953.7 KiB | [postgresql-19-datasketches_1.7.0-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/datasketches/postgresql-19-datasketches_1.7.0-2PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `datasketches_18` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 315.2 KiB | [datasketches_18-1.7.0-3PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/datasketches_18-1.7.0-3PGDG.rhel8.10.x86_64.rpm) |
| `datasketches_18` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 989.4 KiB | [datasketches_18-1.7.0-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/datasketches_18-1.7.0-2PGSTY.el8.x86_64.rpm) |
| `datasketches_18` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 315.1 KiB | [datasketches_18-1.7.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/datasketches_18-1.7.0-2PGDG.rhel8.10.x86_64.rpm) |
| `datasketches_18` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 314.9 KiB | [datasketches_18-1.7.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/datasketches_18-1.7.0-1PGDG.rhel8.10.x86_64.rpm) |
| `datasketches_18` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 301.9 KiB | [datasketches_18-1.7.0-3PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/datasketches_18-1.7.0-3PGDG.rhel8.10.aarch64.rpm) |
| `datasketches_18` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 970.9 KiB | [datasketches_18-1.7.0-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/datasketches_18-1.7.0-2PGSTY.el8.aarch64.rpm) |
| `datasketches_18` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 301.8 KiB | [datasketches_18-1.7.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/datasketches_18-1.7.0-2PGDG.rhel8.10.aarch64.rpm) |
| `datasketches_18` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 301.5 KiB | [datasketches_18-1.7.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/datasketches_18-1.7.0-1PGDG.rhel8.10.aarch64.rpm) |
| `datasketches_18` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 312.8 KiB | [datasketches_18-1.7.0-3PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/datasketches_18-1.7.0-3PGDG.rhel9.9.x86_64.rpm) |
| `datasketches_18` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 966.0 KiB | [datasketches_18-1.7.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/datasketches_18-1.7.0-2PGSTY.el9.x86_64.rpm) |
| `datasketches_18` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 315.3 KiB | [datasketches_18-1.7.0-3PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/datasketches_18-1.7.0-3PGDG.rhel9.9.aarch64.rpm) |
| `datasketches_18` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 979.4 KiB | [datasketches_18-1.7.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/datasketches_18-1.7.0-2PGSTY.el9.aarch64.rpm) |
| `datasketches_18` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 319.7 KiB | [datasketches_18-1.7.0-3PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/datasketches_18-1.7.0-3PGDG.rhel10.3.x86_64.rpm) |
| `datasketches_18` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 989.3 KiB | [datasketches_18-1.7.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/datasketches_18-1.7.0-2PGSTY.el10.x86_64.rpm) |
| `datasketches_18` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 317.5 KiB | [datasketches_18-1.7.0-3PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/datasketches_18-1.7.0-3PGDG.rhel10.3.aarch64.rpm) |
| `datasketches_18` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 315.3 KiB | [datasketches_18-1.7.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/datasketches_18-1.7.0-3PGDG.rhel10.2.aarch64.rpm) |
| `datasketches_18` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 994.4 KiB | [datasketches_18-1.7.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/datasketches_18-1.7.0-2PGSTY.el10.aarch64.rpm) |
| `datasketches_18` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 315.2 KiB | [datasketches_18-1.7.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/datasketches_18-1.7.0-2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-datasketches` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 918.7 KiB | [postgresql-18-datasketches_1.7.0-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-2PGSTY~bookworm_amd64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 919.6 KiB | [postgresql-18-datasketches_1.7.0-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-2PGSTY~bookworm_arm64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 944.5 KiB | [postgresql-18-datasketches_1.7.0-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-2PGSTY~trixie_amd64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 945.8 KiB | [postgresql-18-datasketches_1.7.0-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-2PGSTY~trixie_arm64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1017.1 KiB | [postgresql-18-datasketches_1.7.0-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-2PGSTY~jammy_amd64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1021.9 KiB | [postgresql-18-datasketches_1.7.0-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-2PGSTY~jammy_arm64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 979.8 KiB | [postgresql-18-datasketches_1.7.0-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-2PGSTY~noble_amd64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 992.4 KiB | [postgresql-18-datasketches_1.7.0-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-2PGSTY~noble_arm64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 947.8 KiB | [postgresql-18-datasketches_1.7.0-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-2PGSTY~resolute_amd64.deb) |
| `postgresql-18-datasketches` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 953.8 KiB | [postgresql-18-datasketches_1.7.0-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/datasketches/postgresql-18-datasketches_1.7.0-2PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `datasketches_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 315.2 KiB | [datasketches_17-1.7.0-3PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/datasketches_17-1.7.0-3PGDG.rhel8.10.x86_64.rpm) |
| `datasketches_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 989.3 KiB | [datasketches_17-1.7.0-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/datasketches_17-1.7.0-2PGSTY.el8.x86_64.rpm) |
| `datasketches_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 315.0 KiB | [datasketches_17-1.7.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/datasketches_17-1.7.0-2PGDG.rhel8.10.x86_64.rpm) |
| `datasketches_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 314.8 KiB | [datasketches_17-1.7.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/datasketches_17-1.7.0-1PGDG.rhel8.10.x86_64.rpm) |
| `datasketches_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 301.9 KiB | [datasketches_17-1.7.0-3PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/datasketches_17-1.7.0-3PGDG.rhel8.10.aarch64.rpm) |
| `datasketches_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 970.9 KiB | [datasketches_17-1.7.0-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/datasketches_17-1.7.0-2PGSTY.el8.aarch64.rpm) |
| `datasketches_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 301.8 KiB | [datasketches_17-1.7.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/datasketches_17-1.7.0-2PGDG.rhel8.10.aarch64.rpm) |
| `datasketches_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 301.5 KiB | [datasketches_17-1.7.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/datasketches_17-1.7.0-1PGDG.rhel8.10.aarch64.rpm) |
| `datasketches_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 312.9 KiB | [datasketches_17-1.7.0-3PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/datasketches_17-1.7.0-3PGDG.rhel9.9.x86_64.rpm) |
| `datasketches_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 966.3 KiB | [datasketches_17-1.7.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/datasketches_17-1.7.0-2PGSTY.el9.x86_64.rpm) |
| `datasketches_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 315.2 KiB | [datasketches_17-1.7.0-3PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/datasketches_17-1.7.0-3PGDG.rhel9.9.aarch64.rpm) |
| `datasketches_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 979.4 KiB | [datasketches_17-1.7.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/datasketches_17-1.7.0-2PGSTY.el9.aarch64.rpm) |
| `datasketches_17` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 319.6 KiB | [datasketches_17-1.7.0-3PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/datasketches_17-1.7.0-3PGDG.rhel10.3.x86_64.rpm) |
| `datasketches_17` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 989.2 KiB | [datasketches_17-1.7.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/datasketches_17-1.7.0-2PGSTY.el10.x86_64.rpm) |
| `datasketches_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 317.5 KiB | [datasketches_17-1.7.0-3PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/datasketches_17-1.7.0-3PGDG.rhel10.3.aarch64.rpm) |
| `datasketches_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 315.4 KiB | [datasketches_17-1.7.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/datasketches_17-1.7.0-3PGDG.rhel10.2.aarch64.rpm) |
| `datasketches_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 994.5 KiB | [datasketches_17-1.7.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/datasketches_17-1.7.0-2PGSTY.el10.aarch64.rpm) |
| `datasketches_17` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 315.3 KiB | [datasketches_17-1.7.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/datasketches_17-1.7.0-2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-datasketches` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 919.5 KiB | [postgresql-17-datasketches_1.7.0-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-2PGSTY~bookworm_amd64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 919.9 KiB | [postgresql-17-datasketches_1.7.0-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-2PGSTY~bookworm_arm64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 944.8 KiB | [postgresql-17-datasketches_1.7.0-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-2PGSTY~trixie_amd64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 946.3 KiB | [postgresql-17-datasketches_1.7.0-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-2PGSTY~trixie_arm64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.1 MiB | [postgresql-17-datasketches_1.7.0-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-2PGSTY~jammy_amd64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-17-datasketches_1.7.0-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-2PGSTY~jammy_arm64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 979.9 KiB | [postgresql-17-datasketches_1.7.0-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-2PGSTY~noble_amd64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 992.4 KiB | [postgresql-17-datasketches_1.7.0-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-2PGSTY~noble_arm64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 947.8 KiB | [postgresql-17-datasketches_1.7.0-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-2PGSTY~resolute_amd64.deb) |
| `postgresql-17-datasketches` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 953.8 KiB | [postgresql-17-datasketches_1.7.0-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/datasketches/postgresql-17-datasketches_1.7.0-2PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `datasketches_16` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 315.2 KiB | [datasketches_16-1.7.0-3PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/datasketches_16-1.7.0-3PGDG.rhel8.10.x86_64.rpm) |
| `datasketches_16` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 989.3 KiB | [datasketches_16-1.7.0-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/datasketches_16-1.7.0-2PGSTY.el8.x86_64.rpm) |
| `datasketches_16` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 315.0 KiB | [datasketches_16-1.7.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/datasketches_16-1.7.0-2PGDG.rhel8.10.x86_64.rpm) |
| `datasketches_16` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 314.8 KiB | [datasketches_16-1.7.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/datasketches_16-1.7.0-1PGDG.rhel8.10.x86_64.rpm) |
| `datasketches_16` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 301.9 KiB | [datasketches_16-1.7.0-3PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/datasketches_16-1.7.0-3PGDG.rhel8.10.aarch64.rpm) |
| `datasketches_16` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 970.9 KiB | [datasketches_16-1.7.0-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/datasketches_16-1.7.0-2PGSTY.el8.aarch64.rpm) |
| `datasketches_16` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 301.8 KiB | [datasketches_16-1.7.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/datasketches_16-1.7.0-2PGDG.rhel8.10.aarch64.rpm) |
| `datasketches_16` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 301.5 KiB | [datasketches_16-1.7.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/datasketches_16-1.7.0-1PGDG.rhel8.10.aarch64.rpm) |
| `datasketches_16` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 312.9 KiB | [datasketches_16-1.7.0-3PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/datasketches_16-1.7.0-3PGDG.rhel9.9.x86_64.rpm) |
| `datasketches_16` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 965.8 KiB | [datasketches_16-1.7.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/datasketches_16-1.7.0-2PGSTY.el9.x86_64.rpm) |
| `datasketches_16` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 315.2 KiB | [datasketches_16-1.7.0-3PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/datasketches_16-1.7.0-3PGDG.rhel9.9.aarch64.rpm) |
| `datasketches_16` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 979.4 KiB | [datasketches_16-1.7.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/datasketches_16-1.7.0-2PGSTY.el9.aarch64.rpm) |
| `datasketches_16` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 319.6 KiB | [datasketches_16-1.7.0-3PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/datasketches_16-1.7.0-3PGDG.rhel10.3.x86_64.rpm) |
| `datasketches_16` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 989.3 KiB | [datasketches_16-1.7.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/datasketches_16-1.7.0-2PGSTY.el10.x86_64.rpm) |
| `datasketches_16` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 317.4 KiB | [datasketches_16-1.7.0-3PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/datasketches_16-1.7.0-3PGDG.rhel10.3.aarch64.rpm) |
| `datasketches_16` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 315.3 KiB | [datasketches_16-1.7.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/datasketches_16-1.7.0-3PGDG.rhel10.2.aarch64.rpm) |
| `datasketches_16` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 994.5 KiB | [datasketches_16-1.7.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/datasketches_16-1.7.0-2PGSTY.el10.aarch64.rpm) |
| `datasketches_16` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 315.2 KiB | [datasketches_16-1.7.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/datasketches_16-1.7.0-2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-datasketches` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 919.1 KiB | [postgresql-16-datasketches_1.7.0-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-2PGSTY~bookworm_amd64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 919.1 KiB | [postgresql-16-datasketches_1.7.0-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-2PGSTY~bookworm_arm64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 944.5 KiB | [postgresql-16-datasketches_1.7.0-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-2PGSTY~trixie_amd64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 945.3 KiB | [postgresql-16-datasketches_1.7.0-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-2PGSTY~trixie_arm64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.1 MiB | [postgresql-16-datasketches_1.7.0-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-2PGSTY~jammy_amd64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-16-datasketches_1.7.0-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-2PGSTY~jammy_arm64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 979.8 KiB | [postgresql-16-datasketches_1.7.0-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-2PGSTY~noble_amd64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 992.4 KiB | [postgresql-16-datasketches_1.7.0-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-2PGSTY~noble_arm64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 947.7 KiB | [postgresql-16-datasketches_1.7.0-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-2PGSTY~resolute_amd64.deb) |
| `postgresql-16-datasketches` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 953.8 KiB | [postgresql-16-datasketches_1.7.0-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/datasketches/postgresql-16-datasketches_1.7.0-2PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `datasketches_15` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 326.0 KiB | [datasketches_15-1.7.0-3PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/datasketches_15-1.7.0-3PGDG.rhel8.10.x86_64.rpm) |
| `datasketches_15` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 997.3 KiB | [datasketches_15-1.7.0-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/datasketches_15-1.7.0-2PGSTY.el8.x86_64.rpm) |
| `datasketches_15` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 325.8 KiB | [datasketches_15-1.7.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/datasketches_15-1.7.0-2PGDG.rhel8.10.x86_64.rpm) |
| `datasketches_15` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 325.6 KiB | [datasketches_15-1.7.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/datasketches_15-1.7.0-1PGDG.rhel8.10.x86_64.rpm) |
| `datasketches_15` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 312.9 KiB | [datasketches_15-1.7.0-3PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/datasketches_15-1.7.0-3PGDG.rhel8.10.aarch64.rpm) |
| `datasketches_15` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 979.1 KiB | [datasketches_15-1.7.0-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/datasketches_15-1.7.0-2PGSTY.el8.aarch64.rpm) |
| `datasketches_15` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 312.7 KiB | [datasketches_15-1.7.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/datasketches_15-1.7.0-2PGDG.rhel8.10.aarch64.rpm) |
| `datasketches_15` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 312.4 KiB | [datasketches_15-1.7.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/datasketches_15-1.7.0-1PGDG.rhel8.10.aarch64.rpm) |
| `datasketches_15` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 323.6 KiB | [datasketches_15-1.7.0-3PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/datasketches_15-1.7.0-3PGDG.rhel9.9.x86_64.rpm) |
| `datasketches_15` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 973.9 KiB | [datasketches_15-1.7.0-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/datasketches_15-1.7.0-2PGSTY.el9.x86_64.rpm) |
| `datasketches_15` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 325.7 KiB | [datasketches_15-1.7.0-3PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/datasketches_15-1.7.0-3PGDG.rhel9.9.aarch64.rpm) |
| `datasketches_15` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 985.9 KiB | [datasketches_15-1.7.0-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/datasketches_15-1.7.0-2PGSTY.el9.aarch64.rpm) |
| `datasketches_15` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 326.4 KiB | [datasketches_15-1.7.0-3PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/datasketches_15-1.7.0-3PGDG.rhel10.3.x86_64.rpm) |
| `datasketches_15` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 995.2 KiB | [datasketches_15-1.7.0-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/datasketches_15-1.7.0-2PGSTY.el10.x86_64.rpm) |
| `datasketches_15` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 323.9 KiB | [datasketches_15-1.7.0-3PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/datasketches_15-1.7.0-3PGDG.rhel10.3.aarch64.rpm) |
| `datasketches_15` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 320.9 KiB | [datasketches_15-1.7.0-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/datasketches_15-1.7.0-3PGDG.rhel10.2.aarch64.rpm) |
| `datasketches_15` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 998.5 KiB | [datasketches_15-1.7.0-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/datasketches_15-1.7.0-2PGSTY.el10.aarch64.rpm) |
| `datasketches_15` | `1.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 320.7 KiB | [datasketches_15-1.7.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/datasketches_15-1.7.0-2PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-datasketches` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 934.2 KiB | [postgresql-15-datasketches_1.7.0-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-2PGSTY~bookworm_amd64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 934.1 KiB | [postgresql-15-datasketches_1.7.0-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-2PGSTY~bookworm_arm64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 959.0 KiB | [postgresql-15-datasketches_1.7.0-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-2PGSTY~trixie_amd64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 958.6 KiB | [postgresql-15-datasketches_1.7.0-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-2PGSTY~trixie_arm64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.1 MiB | [postgresql-15-datasketches_1.7.0-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-2PGSTY~jammy_amd64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.1 MiB | [postgresql-15-datasketches_1.7.0-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-2PGSTY~jammy_arm64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 986.9 KiB | [postgresql-15-datasketches_1.7.0-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-2PGSTY~noble_amd64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1000.6 KiB | [postgresql-15-datasketches_1.7.0-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-2PGSTY~noble_arm64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 954.1 KiB | [postgresql-15-datasketches_1.7.0-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-2PGSTY~resolute_amd64.deb) |
| `postgresql-15-datasketches` | `1.7.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 961.1 KiB | [postgresql-15-datasketches_1.7.0-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/datasketches/postgresql-15-datasketches_1.7.0-2PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/apache/datasketches-postgresql" title="Repository" icon="github" subtitle="github.com/apache/datasketches-postgresql" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="apache-datasketches-postgresql-1.7.0-src.tar.gz apache-datasketches-cpp-5.2.0-src.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg datasketches;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install datasketches;		# install via package name, for the active PG version

pig install datasketches -v 19;   # install for PG 19
pig install datasketches -v 18;   # install for PG 18
pig install datasketches -v 17;   # install for PG 17
pig install datasketches -v 16;   # install for PG 16
pig install datasketches -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION datasketches;
```




## Usage

Sources: [README](https://github.com/apache/datasketches-postgresql/blob/master/README.md), [latest release 1.7.0](https://github.com/apache/datasketches-postgresql/releases/tag/1.7.0), [Apache DataSketches](https://datasketches.apache.org/)

`datasketches` adds approximate analytics sketch types and aggregates to PostgreSQL. The upstream README lists CPC, HLL, Theta, Array Of Doubles, KLL, Quantiles, and Frequent Strings sketches; the 1.7.0 release is the latest published GitHub release, while the default branch has already moved on to `1.8.0-SNAPSHOT`.

```sql
CREATE EXTENSION datasketches;
```

### Core Sketch Families

- `cpc_sketch` and `hll_sketch` for approximate distinct counting.
- `theta_sketch` for distinct counting plus set operations such as union, intersection, and A-not-B.
- `aod_sketch` for tuple-style metrics keyed by identifiers with arrays of doubles.
- `kll_*_sketch` and `quantiles_*_sketch` for quantiles, ranks, PMF, and CDF.
- `frequent_strings_sketch` for heavy-hitter detection.

### Common Patterns

Build a sketch from raw values:

```sql
SELECT cpc_sketch_build(1);
SELECT kll_float_sketch_build(value) FROM normal;
```

Use one-shot approximate aggregates:

```sql
SELECT cpc_sketch_distinct(id) FROM random_ints_100m;
```

Merge sketches across groups or cube dimensions:

```sql
SELECT cpc_sketch_get_estimate(cpc_sketch_union(sketch)) FROM cpc_sketch_test;
SELECT hll_sketch_get_estimate(hll_sketch_union(sketch)) FROM hll_sketch_test;
SELECT kll_float_sketch_get_quantile(kll_float_sketch_merge(sketch), 0.5)
FROM kll_float_sketch_test;
```

Run set operations on Theta sketches:

```sql
SELECT theta_sketch_get_estimate(theta_sketch_intersection(sketch1, sketch2))
FROM theta_set_op_test;
```

Find frequent items above a threshold:

```sql
SELECT frequent_strings_sketch_result_no_false_negatives(
  frequent_strings_sketch_build(9, value),
  1000000
)
FROM zipf_1p1_8k_100m;
```

### Caveats

- Upstream documents PostgreSQL 9.6+ plus Boost 1.75.0 and DataSketches C++ core 5.0.0 or later as build dependencies.
- These are approximate structures meant to be mergeable across dimensions; they are not exact replacements for `COUNT(DISTINCT ...)` or exact histograms.
