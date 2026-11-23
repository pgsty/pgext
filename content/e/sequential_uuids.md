---
title: "sequential_uuids"
linkTitle: "sequential_uuids"
description: "generator of sequential UUIDs"
weight: 4570
categories: ["FUNC"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**sequential_uuids**](https://github.com/tvondra/sequential-uuids) : generator of sequential UUIDs


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4570** | {{< badge content="sequential_uuids" link="https://github.com/tvondra/sequential-uuids" >}} | {{< ext "sequential_uuids" >}} | `1.0.3` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_uuidv7" >}} {{< ext "snowflake" >}} {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuid_v8" >}} {{< ext "uuid-ossp" >}} {{< ext "typeid" >}} {{< ext "permuteseq" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.3` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `sequential_uuids` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.3` | {{< bg "19" "sequential_uuids_19" "green" >}} {{< bg "18" "sequential_uuids_18" "green" >}} {{< bg "17" "sequential_uuids_17" "green" >}} {{< bg "16" "sequential_uuids_16" "green" >}} {{< bg "15" "sequential_uuids_15" "green" >}} | `sequential_uuids_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.3` | {{< bg "19" "postgresql-19-sequential-uuids" "green" >}} {{< bg "18" "postgresql-18-sequential-uuids" "green" >}} {{< bg "17" "postgresql-17-sequential-uuids" "green" >}} {{< bg "16" "postgresql-16-sequential-uuids" "green" >}} {{< bg "15" "postgresql-15-sequential-uuids" "green" >}} | `postgresql-$v-sequential-uuids` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "sequential_uuids_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_15 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "sequential_uuids_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_15 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_19 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_15 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_19 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_19 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_19 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.0.3" "sequential_uuids_15 : AVAIL 5" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-19-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-19-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-19-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-19-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-19-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-19-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-19-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-19-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-19-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-19-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-18-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sequential_uuids_19` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.8 KiB | [sequential_uuids_19-1.0.3-8PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/sequential_uuids_19-1.0.3-8PGDG.rhel9.9.x86_64.rpm) |
| `sequential_uuids_19` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.2 KiB | [sequential_uuids_19-1.0.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sequential_uuids_19-1.0.3-1PGSTY.el9.x86_64.rpm) |
| `sequential_uuids_19` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.6 KiB | [sequential_uuids_19-1.0.3-8PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/sequential_uuids_19-1.0.3-8PGDG.rhel9.9.aarch64.rpm) |
| `sequential_uuids_19` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [sequential_uuids_19-1.0.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sequential_uuids_19-1.0.3-1PGSTY.el9.aarch64.rpm) |
| `sequential_uuids_19` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.8 KiB | [sequential_uuids_19-1.0.3-8PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/sequential_uuids_19-1.0.3-8PGDG.rhel10.3.x86_64.rpm) |
| `sequential_uuids_19` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.0 KiB | [sequential_uuids_19-1.0.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sequential_uuids_19-1.0.3-1PGSTY.el10.x86_64.rpm) |
| `sequential_uuids_19` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.8 KiB | [sequential_uuids_19-1.0.3-8PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/sequential_uuids_19-1.0.3-8PGDG.rhel10.3.aarch64.rpm) |
| `sequential_uuids_19` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.0 KiB | [sequential_uuids_19-1.0.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sequential_uuids_19-1.0.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-19-sequential-uuids` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.9 KiB | [postgresql-19-sequential-uuids_1.0.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-19-sequential-uuids_1.0.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-19-sequential-uuids` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.8 KiB | [postgresql-19-sequential-uuids_1.0.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-19-sequential-uuids_1.0.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-19-sequential-uuids` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.8 KiB | [postgresql-19-sequential-uuids_1.0.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sequential-uuids/postgresql-19-sequential-uuids_1.0.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-19-sequential-uuids` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.9 KiB | [postgresql-19-sequential-uuids_1.0.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sequential-uuids/postgresql-19-sequential-uuids_1.0.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-19-sequential-uuids` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.1 KiB | [postgresql-19-sequential-uuids_1.0.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-19-sequential-uuids_1.0.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-19-sequential-uuids` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.1 KiB | [postgresql-19-sequential-uuids_1.0.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-19-sequential-uuids_1.0.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-19-sequential-uuids` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.2 KiB | [postgresql-19-sequential-uuids_1.0.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-19-sequential-uuids_1.0.3-1PGSTY~noble_amd64.deb) |
| `postgresql-19-sequential-uuids` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.2 KiB | [postgresql-19-sequential-uuids_1.0.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-19-sequential-uuids_1.0.3-1PGSTY~noble_arm64.deb) |
| `postgresql-19-sequential-uuids` | `1.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.4 KiB | [postgresql-19-sequential-uuids_1.0.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sequential-uuids/postgresql-19-sequential-uuids_1.0.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-19-sequential-uuids` | `1.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.5 KiB | [postgresql-19-sequential-uuids_1.0.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sequential-uuids/postgresql-19-sequential-uuids_1.0.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sequential_uuids_18` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.2 KiB | [sequential_uuids_18-1.0.3-8PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/sequential_uuids_18-1.0.3-8PGDG.rhel8.10.x86_64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.2 KiB | [sequential_uuids_18-1.0.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/sequential_uuids_18-1.0.3-2PGDG.rhel8.x86_64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.8 KiB | [sequential_uuids_18-1.0.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sequential_uuids_18-1.0.3-1PGSTY.el8.x86_64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.2 KiB | [sequential_uuids_18-1.0.3-8PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/sequential_uuids_18-1.0.3-8PGDG.rhel8.10.aarch64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 16.2 KiB | [sequential_uuids_18-1.0.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/sequential_uuids_18-1.0.3-2PGDG.rhel8.aarch64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.0 KiB | [sequential_uuids_18-1.0.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sequential_uuids_18-1.0.3-1PGSTY.el8.aarch64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.8 KiB | [sequential_uuids_18-1.0.3-8PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/sequential_uuids_18-1.0.3-8PGDG.rhel9.9.x86_64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.2 KiB | [sequential_uuids_18-1.0.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sequential_uuids_18-1.0.3-1PGSTY.el9.x86_64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.6 KiB | [sequential_uuids_18-1.0.3-8PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/sequential_uuids_18-1.0.3-8PGDG.rhel9.9.aarch64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [sequential_uuids_18-1.0.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sequential_uuids_18-1.0.3-1PGSTY.el9.aarch64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.8 KiB | [sequential_uuids_18-1.0.3-8PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/sequential_uuids_18-1.0.3-8PGDG.rhel10.3.x86_64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.0 KiB | [sequential_uuids_18-1.0.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sequential_uuids_18-1.0.3-1PGSTY.el10.x86_64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.8 KiB | [sequential_uuids_18-1.0.3-8PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/sequential_uuids_18-1.0.3-8PGDG.rhel10.3.aarch64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.8 KiB | [sequential_uuids_18-1.0.3-8PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/sequential_uuids_18-1.0.3-8PGDG.rhel10.2.aarch64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.0 KiB | [sequential_uuids_18-1.0.3-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/sequential_uuids_18-1.0.3-4PGDG.rhel10.2.aarch64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.3 KiB | [sequential_uuids_18-1.0.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/sequential_uuids_18-1.0.3-2PGDG.rhel10.aarch64.rpm) |
| `sequential_uuids_18` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.0 KiB | [sequential_uuids_18-1.0.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sequential_uuids_18-1.0.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-sequential-uuids` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-18-sequential-uuids_1.0.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-18-sequential-uuids_1.0.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-sequential-uuids` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.8 KiB | [postgresql-18-sequential-uuids_1.0.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-18-sequential-uuids_1.0.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-sequential-uuids` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.8 KiB | [postgresql-18-sequential-uuids_1.0.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sequential-uuids/postgresql-18-sequential-uuids_1.0.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-sequential-uuids` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.9 KiB | [postgresql-18-sequential-uuids_1.0.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sequential-uuids/postgresql-18-sequential-uuids_1.0.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-sequential-uuids` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.1 KiB | [postgresql-18-sequential-uuids_1.0.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-18-sequential-uuids_1.0.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-sequential-uuids` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.1 KiB | [postgresql-18-sequential-uuids_1.0.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-18-sequential-uuids_1.0.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-sequential-uuids` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.2 KiB | [postgresql-18-sequential-uuids_1.0.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-18-sequential-uuids_1.0.3-1PGSTY~noble_amd64.deb) |
| `postgresql-18-sequential-uuids` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.2 KiB | [postgresql-18-sequential-uuids_1.0.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-18-sequential-uuids_1.0.3-1PGSTY~noble_arm64.deb) |
| `postgresql-18-sequential-uuids` | `1.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.4 KiB | [postgresql-18-sequential-uuids_1.0.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sequential-uuids/postgresql-18-sequential-uuids_1.0.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-sequential-uuids` | `1.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.5 KiB | [postgresql-18-sequential-uuids_1.0.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sequential-uuids/postgresql-18-sequential-uuids_1.0.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sequential_uuids_17` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.2 KiB | [sequential_uuids_17-1.0.3-8PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/sequential_uuids_17-1.0.3-8PGDG.rhel8.10.x86_64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.8 KiB | [sequential_uuids_17-1.0.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sequential_uuids_17-1.0.3-1PGSTY.el8.x86_64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.2 KiB | [sequential_uuids_17-1.0.3-8PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/sequential_uuids_17-1.0.3-8PGDG.rhel8.10.aarch64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.0 KiB | [sequential_uuids_17-1.0.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sequential_uuids_17-1.0.3-1PGSTY.el8.aarch64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.8 KiB | [sequential_uuids_17-1.0.3-8PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/sequential_uuids_17-1.0.3-8PGDG.rhel9.9.x86_64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.2 KiB | [sequential_uuids_17-1.0.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sequential_uuids_17-1.0.3-1PGSTY.el9.x86_64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.6 KiB | [sequential_uuids_17-1.0.3-8PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/sequential_uuids_17-1.0.3-8PGDG.rhel9.9.aarch64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [sequential_uuids_17-1.0.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sequential_uuids_17-1.0.3-1PGSTY.el9.aarch64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.8 KiB | [sequential_uuids_17-1.0.3-8PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/sequential_uuids_17-1.0.3-8PGDG.rhel10.3.x86_64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [sequential_uuids_17-1.0.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sequential_uuids_17-1.0.3-1PGSTY.el10.x86_64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.8 KiB | [sequential_uuids_17-1.0.3-8PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/sequential_uuids_17-1.0.3-8PGDG.rhel10.3.aarch64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.8 KiB | [sequential_uuids_17-1.0.3-8PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/sequential_uuids_17-1.0.3-8PGDG.rhel10.2.aarch64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.0 KiB | [sequential_uuids_17-1.0.3-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/sequential_uuids_17-1.0.3-4PGDG.rhel10.2.aarch64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.3 KiB | [sequential_uuids_17-1.0.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/sequential_uuids_17-1.0.3-2PGDG.rhel10.aarch64.rpm) |
| `sequential_uuids_17` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.0 KiB | [sequential_uuids_17-1.0.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sequential_uuids_17-1.0.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-sequential-uuids` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-17-sequential-uuids_1.0.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-sequential-uuids` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.8 KiB | [postgresql-17-sequential-uuids_1.0.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-sequential-uuids` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.8 KiB | [postgresql-17-sequential-uuids_1.0.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-sequential-uuids` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.9 KiB | [postgresql-17-sequential-uuids_1.0.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-sequential-uuids` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.3 KiB | [postgresql-17-sequential-uuids_1.0.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-sequential-uuids` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.3 KiB | [postgresql-17-sequential-uuids_1.0.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-sequential-uuids` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.2 KiB | [postgresql-17-sequential-uuids_1.0.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PGSTY~noble_amd64.deb) |
| `postgresql-17-sequential-uuids` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.2 KiB | [postgresql-17-sequential-uuids_1.0.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PGSTY~noble_arm64.deb) |
| `postgresql-17-sequential-uuids` | `1.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.4 KiB | [postgresql-17-sequential-uuids_1.0.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-sequential-uuids` | `1.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.4 KiB | [postgresql-17-sequential-uuids_1.0.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sequential_uuids_16` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.2 KiB | [sequential_uuids_16-1.0.3-8PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/sequential_uuids_16-1.0.3-8PGDG.rhel8.10.x86_64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.8 KiB | [sequential_uuids_16-1.0.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sequential_uuids_16-1.0.3-1PGSTY.el8.x86_64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.2 KiB | [sequential_uuids_16-1.0.3-8PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/sequential_uuids_16-1.0.3-8PGDG.rhel8.10.aarch64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.0 KiB | [sequential_uuids_16-1.0.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sequential_uuids_16-1.0.3-1PGSTY.el8.aarch64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.8 KiB | [sequential_uuids_16-1.0.3-8PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/sequential_uuids_16-1.0.3-8PGDG.rhel9.9.x86_64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.2 KiB | [sequential_uuids_16-1.0.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sequential_uuids_16-1.0.3-1PGSTY.el9.x86_64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.6 KiB | [sequential_uuids_16-1.0.3-8PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/sequential_uuids_16-1.0.3-8PGDG.rhel9.9.aarch64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 18.9 KiB | [sequential_uuids_16-1.0.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sequential_uuids_16-1.0.3-1PGSTY.el9.aarch64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.8 KiB | [sequential_uuids_16-1.0.3-8PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/sequential_uuids_16-1.0.3-8PGDG.rhel10.3.x86_64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [sequential_uuids_16-1.0.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sequential_uuids_16-1.0.3-1PGSTY.el10.x86_64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.8 KiB | [sequential_uuids_16-1.0.3-8PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/sequential_uuids_16-1.0.3-8PGDG.rhel10.3.aarch64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.8 KiB | [sequential_uuids_16-1.0.3-8PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/sequential_uuids_16-1.0.3-8PGDG.rhel10.2.aarch64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.0 KiB | [sequential_uuids_16-1.0.3-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/sequential_uuids_16-1.0.3-4PGDG.rhel10.2.aarch64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.3 KiB | [sequential_uuids_16-1.0.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/sequential_uuids_16-1.0.3-2PGDG.rhel10.aarch64.rpm) |
| `sequential_uuids_16` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.0 KiB | [sequential_uuids_16-1.0.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sequential_uuids_16-1.0.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-sequential-uuids` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-16-sequential-uuids_1.0.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-sequential-uuids` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.8 KiB | [postgresql-16-sequential-uuids_1.0.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-sequential-uuids` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.8 KiB | [postgresql-16-sequential-uuids_1.0.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-sequential-uuids` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.9 KiB | [postgresql-16-sequential-uuids_1.0.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-sequential-uuids` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.3 KiB | [postgresql-16-sequential-uuids_1.0.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-sequential-uuids` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.3 KiB | [postgresql-16-sequential-uuids_1.0.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-sequential-uuids` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.2 KiB | [postgresql-16-sequential-uuids_1.0.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PGSTY~noble_amd64.deb) |
| `postgresql-16-sequential-uuids` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.2 KiB | [postgresql-16-sequential-uuids_1.0.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PGSTY~noble_arm64.deb) |
| `postgresql-16-sequential-uuids` | `1.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.4 KiB | [postgresql-16-sequential-uuids_1.0.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-sequential-uuids` | `1.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.4 KiB | [postgresql-16-sequential-uuids_1.0.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sequential_uuids_15` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.2 KiB | [sequential_uuids_15-1.0.3-8PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/sequential_uuids_15-1.0.3-8PGDG.rhel8.10.x86_64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 18.8 KiB | [sequential_uuids_15-1.0.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sequential_uuids_15-1.0.3-1PGSTY.el8.x86_64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.2 KiB | [sequential_uuids_15-1.0.3-8PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/sequential_uuids_15-1.0.3-8PGDG.rhel8.10.aarch64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.0 KiB | [sequential_uuids_15-1.0.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sequential_uuids_15-1.0.3-1PGSTY.el8.aarch64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.8 KiB | [sequential_uuids_15-1.0.3-8PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/sequential_uuids_15-1.0.3-8PGDG.rhel9.9.x86_64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.2 KiB | [sequential_uuids_15-1.0.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sequential_uuids_15-1.0.3-1PGSTY.el9.x86_64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 16.6 KiB | [sequential_uuids_15-1.0.3-8PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/sequential_uuids_15-1.0.3-8PGDG.rhel9.9.aarch64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [sequential_uuids_15-1.0.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sequential_uuids_15-1.0.3-1PGSTY.el9.aarch64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 16.8 KiB | [sequential_uuids_15-1.0.3-8PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/sequential_uuids_15-1.0.3-8PGDG.rhel10.3.x86_64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 18.9 KiB | [sequential_uuids_15-1.0.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sequential_uuids_15-1.0.3-1PGSTY.el10.x86_64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.8 KiB | [sequential_uuids_15-1.0.3-8PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/sequential_uuids_15-1.0.3-8PGDG.rhel10.3.aarch64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.8 KiB | [sequential_uuids_15-1.0.3-8PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/sequential_uuids_15-1.0.3-8PGDG.rhel10.2.aarch64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.0 KiB | [sequential_uuids_15-1.0.3-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/sequential_uuids_15-1.0.3-4PGDG.rhel10.2.aarch64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 16.3 KiB | [sequential_uuids_15-1.0.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/sequential_uuids_15-1.0.3-2PGDG.rhel10.aarch64.rpm) |
| `sequential_uuids_15` | `1.0.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.0 KiB | [sequential_uuids_15-1.0.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sequential_uuids_15-1.0.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-sequential-uuids` | `1.0.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-15-sequential-uuids_1.0.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-sequential-uuids` | `1.0.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.8 KiB | [postgresql-15-sequential-uuids_1.0.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-sequential-uuids` | `1.0.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.8 KiB | [postgresql-15-sequential-uuids_1.0.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-sequential-uuids` | `1.0.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.9 KiB | [postgresql-15-sequential-uuids_1.0.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-sequential-uuids` | `1.0.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.3 KiB | [postgresql-15-sequential-uuids_1.0.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-sequential-uuids` | `1.0.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.3 KiB | [postgresql-15-sequential-uuids_1.0.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-sequential-uuids` | `1.0.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.2 KiB | [postgresql-15-sequential-uuids_1.0.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PGSTY~noble_amd64.deb) |
| `postgresql-15-sequential-uuids` | `1.0.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.2 KiB | [postgresql-15-sequential-uuids_1.0.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PGSTY~noble_arm64.deb) |
| `postgresql-15-sequential-uuids` | `1.0.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 13.4 KiB | [postgresql-15-sequential-uuids_1.0.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-sequential-uuids` | `1.0.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 13.4 KiB | [postgresql-15-sequential-uuids_1.0.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/sequential-uuids" title="Repository" icon="github" subtitle="github.com/tvondra/sequential-uuids" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="sequential-uuids-1.0.3.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg sequential_uuids;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install sequential_uuids;		# install via package name, for the active PG version

pig install sequential_uuids -v 19;   # install for PG 19
pig install sequential_uuids -v 18;   # install for PG 18
pig install sequential_uuids -v 17;   # install for PG 17
pig install sequential_uuids -v 16;   # install for PG 16
pig install sequential_uuids -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION sequential_uuids;
```




## Usage

> [sequential_uuids: sequential UUID generators for better index locality](https://github.com/tvondra/sequential-uuids)

Generates UUIDs with sequential patterns to reduce random I/O in indexes while maintaining sufficient randomness to avoid collisions.

```sql
CREATE EXTENSION sequential_uuids;
```

### Functions

| Function | Description |
|---|---|
| `uuid_sequence_nextval(sequence regclass, block_size int DEFAULT 65536, block_count int DEFAULT 65536)` | Generate a sequential UUID based on a sequence |
| `uuid_time_nextval(interval_length int DEFAULT 60, interval_count int DEFAULT 65536)` | Generate a sequential UUID based on current timestamp |

### Examples

```sql
CREATE SEQUENCE my_seq;

-- Sequence-based UUID generation
SELECT uuid_sequence_nextval('my_seq'::regclass);

-- Time-based UUID generation (wraps around every ~45 days with defaults)
SELECT uuid_time_nextval();

-- Use as default for a column
CREATE TABLE orders (
  id uuid DEFAULT uuid_time_nextval() PRIMARY KEY,
  data text
);

-- Custom block size and count
SELECT uuid_sequence_nextval('my_seq', 256, 65536);
SELECT uuid_time_nextval(120, 32768);
```
