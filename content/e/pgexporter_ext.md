---
title: "pgexporter_ext"
linkTitle: "pgexporter_ext"
description: "Operating-system and host metrics for pgexporter"
weight: 6430
categories: ["STAT"]
languages: ["C"]
licenses: ["BSD-3-Clause"]
repos: ["PGDG"]
page_width: full
---

[**pgexporter_ext**](https://github.com/pgexporter/pgexporter_ext) : Operating-system and host metrics for pgexporter


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6430** | {{< badge content="pgexporter_ext" link="https://github.com/pgexporter/pgexporter_ext" >}} | {{< ext "pgexporter_ext" >}} | `0.2.5` | {{< category "STAT" >}} | {{< license "BSD-3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgnodemx" >}} {{< ext "system_stats" >}} {{< ext "pg_proctab" >}} {{< ext "pg_statvfs" >}} {{< ext "pgmonitor" >}} {{< ext "pgtelemetry" >}} |

> [!Note] Release/control and PIGSTY DEB version are 0.2.5; principal PGDG RPM is 0.2.4, except EL8 PG14-16 at 0.2.3. The guide configures preload, but 0.2.5 metrics work without it (verified). Metric access requires pg_monitor; local build capability does not authorize replacing PGDG packages.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.2.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgexporter_ext` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.2.4` | {{< bg "18" "pgexporter_ext_18" "green" >}} {{< bg "17" "pgexporter_ext_17" "green" >}} {{< bg "16" "pgexporter_ext_16" "green" >}} {{< bg "15" "pgexporter_ext_15" "green" >}} {{< bg "14" "pgexporter_ext_14" "green" >}} | `pgexporter_ext_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.5` | {{< bg "18" "postgresql-18-pgexporter-ext" "green" >}} {{< bg "17" "postgresql-17-pgexporter-ext" "green" >}} {{< bg "16" "postgresql-16-pgexporter-ext" "green" >}} {{< bg "15" "postgresql-15-pgexporter-ext" "green" >}} {{< bg "14" "postgresql-14-pgexporter-ext" "green" >}} | `postgresql-$v-pgexporter-ext` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.3" "pgexporter_ext_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.3" "pgexporter_ext_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.3" "pgexporter_ext_14 : AVAIL 5" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.3" "pgexporter_ext_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.3" "pgexporter_ext_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.3" "pgexporter_ext_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_14 : AVAIL 5" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.2.4" "pgexporter_ext_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-18-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-17-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-16-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-15-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-14-pgexporter-ext : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-18-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-17-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-16-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-15-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-14-pgexporter-ext : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-18-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-17-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-16-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-15-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-14-pgexporter-ext : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-18-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-17-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-16-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-15-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-14-pgexporter-ext : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-18-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-17-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-16-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-15-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-14-pgexporter-ext : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-18-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-17-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-16-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-15-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-14-pgexporter-ext : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-18-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-17-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-16-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-15-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-14-pgexporter-ext : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-18-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-17-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-16-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-15-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-14-pgexporter-ext : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-18-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-17-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-16-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-15-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-14-pgexporter-ext : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-18-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-17-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-16-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-15-pgexporter-ext : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.5" "postgresql-14-pgexporter-ext : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgexporter_ext_18` | `0.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.3 KiB | [pgexporter_ext_18-0.2.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgexporter_ext_18-0.2.4-1PGDG.rhel8.x86_64.rpm) |
| `pgexporter_ext_18` | `0.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.2 KiB | [pgexporter_ext_18-0.2.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgexporter_ext_18-0.2.4-1PGDG.rhel8.aarch64.rpm) |
| `pgexporter_ext_18` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.4 KiB | [pgexporter_ext_18-0.2.4-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgexporter_ext_18-0.2.4-2PGDG.rhel9.8.x86_64.rpm) |
| `pgexporter_ext_18` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.6 KiB | [pgexporter_ext_18-0.2.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgexporter_ext_18-0.2.4-1PGDG.rhel9.x86_64.rpm) |
| `pgexporter_ext_18` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [pgexporter_ext_18-0.2.4-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgexporter_ext_18-0.2.4-2PGDG.rhel9.8.aarch64.rpm) |
| `pgexporter_ext_18` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.5 KiB | [pgexporter_ext_18-0.2.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgexporter_ext_18-0.2.4-1PGDG.rhel9.aarch64.rpm) |
| `pgexporter_ext_18` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.3 KiB | [pgexporter_ext_18-0.2.4-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgexporter_ext_18-0.2.4-2PGDG.rhel10.2.x86_64.rpm) |
| `pgexporter_ext_18` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.8 KiB | [pgexporter_ext_18-0.2.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgexporter_ext_18-0.2.4-1PGDG.rhel10.x86_64.rpm) |
| `pgexporter_ext_18` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.8 KiB | [pgexporter_ext_18-0.2.4-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgexporter_ext_18-0.2.4-2PGDG.rhel10.2.aarch64.rpm) |
| `pgexporter_ext_18` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pgexporter_ext_18-0.2.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgexporter_ext_18-0.2.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgexporter-ext` | `0.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.7 KiB | [postgresql-18-pgexporter-ext_0.2.5-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgexporter-ext/postgresql-18-pgexporter-ext_0.2.5-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgexporter-ext` | `0.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.5 KiB | [postgresql-18-pgexporter-ext_0.2.5-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgexporter-ext/postgresql-18-pgexporter-ext_0.2.5-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgexporter-ext` | `0.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.8 KiB | [postgresql-18-pgexporter-ext_0.2.5-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgexporter-ext/postgresql-18-pgexporter-ext_0.2.5-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pgexporter-ext` | `0.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.5 KiB | [postgresql-18-pgexporter-ext_0.2.5-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgexporter-ext/postgresql-18-pgexporter-ext_0.2.5-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pgexporter-ext` | `0.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.4 KiB | [postgresql-18-pgexporter-ext_0.2.5-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgexporter-ext/postgresql-18-pgexporter-ext_0.2.5-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pgexporter-ext` | `0.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.3 KiB | [postgresql-18-pgexporter-ext_0.2.5-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgexporter-ext/postgresql-18-pgexporter-ext_0.2.5-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pgexporter-ext` | `0.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.2 KiB | [postgresql-18-pgexporter-ext_0.2.5-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgexporter-ext/postgresql-18-pgexporter-ext_0.2.5-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pgexporter-ext` | `0.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.3 KiB | [postgresql-18-pgexporter-ext_0.2.5-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgexporter-ext/postgresql-18-pgexporter-ext_0.2.5-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pgexporter-ext` | `0.2.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.3 KiB | [postgresql-18-pgexporter-ext_0.2.5-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgexporter-ext/postgresql-18-pgexporter-ext_0.2.5-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pgexporter-ext` | `0.2.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.2 KiB | [postgresql-18-pgexporter-ext_0.2.5-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgexporter-ext/postgresql-18-pgexporter-ext_0.2.5-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgexporter_ext_17` | `0.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.3 KiB | [pgexporter_ext_17-0.2.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgexporter_ext_17-0.2.4-1PGDG.rhel8.x86_64.rpm) |
| `pgexporter_ext_17` | `0.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.2 KiB | [pgexporter_ext_17-0.2.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgexporter_ext_17-0.2.4-1PGDG.rhel8.aarch64.rpm) |
| `pgexporter_ext_17` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.4 KiB | [pgexporter_ext_17-0.2.4-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgexporter_ext_17-0.2.4-2PGDG.rhel9.8.x86_64.rpm) |
| `pgexporter_ext_17` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.7 KiB | [pgexporter_ext_17-0.2.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgexporter_ext_17-0.2.4-1PGDG.rhel9.x86_64.rpm) |
| `pgexporter_ext_17` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [pgexporter_ext_17-0.2.4-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgexporter_ext_17-0.2.4-2PGDG.rhel9.8.aarch64.rpm) |
| `pgexporter_ext_17` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.5 KiB | [pgexporter_ext_17-0.2.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgexporter_ext_17-0.2.4-1PGDG.rhel9.aarch64.rpm) |
| `pgexporter_ext_17` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.3 KiB | [pgexporter_ext_17-0.2.4-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgexporter_ext_17-0.2.4-2PGDG.rhel10.2.x86_64.rpm) |
| `pgexporter_ext_17` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.8 KiB | [pgexporter_ext_17-0.2.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgexporter_ext_17-0.2.4-1PGDG.rhel10.x86_64.rpm) |
| `pgexporter_ext_17` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.8 KiB | [pgexporter_ext_17-0.2.4-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgexporter_ext_17-0.2.4-2PGDG.rhel10.2.aarch64.rpm) |
| `pgexporter_ext_17` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pgexporter_ext_17-0.2.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgexporter_ext_17-0.2.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgexporter-ext` | `0.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.7 KiB | [postgresql-17-pgexporter-ext_0.2.5-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgexporter-ext/postgresql-17-pgexporter-ext_0.2.5-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgexporter-ext` | `0.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.5 KiB | [postgresql-17-pgexporter-ext_0.2.5-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgexporter-ext/postgresql-17-pgexporter-ext_0.2.5-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgexporter-ext` | `0.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.8 KiB | [postgresql-17-pgexporter-ext_0.2.5-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgexporter-ext/postgresql-17-pgexporter-ext_0.2.5-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pgexporter-ext` | `0.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.5 KiB | [postgresql-17-pgexporter-ext_0.2.5-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgexporter-ext/postgresql-17-pgexporter-ext_0.2.5-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pgexporter-ext` | `0.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.5 KiB | [postgresql-17-pgexporter-ext_0.2.5-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgexporter-ext/postgresql-17-pgexporter-ext_0.2.5-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pgexporter-ext` | `0.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.3 KiB | [postgresql-17-pgexporter-ext_0.2.5-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgexporter-ext/postgresql-17-pgexporter-ext_0.2.5-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pgexporter-ext` | `0.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.3 KiB | [postgresql-17-pgexporter-ext_0.2.5-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgexporter-ext/postgresql-17-pgexporter-ext_0.2.5-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pgexporter-ext` | `0.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.3 KiB | [postgresql-17-pgexporter-ext_0.2.5-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgexporter-ext/postgresql-17-pgexporter-ext_0.2.5-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pgexporter-ext` | `0.2.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.3 KiB | [postgresql-17-pgexporter-ext_0.2.5-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgexporter-ext/postgresql-17-pgexporter-ext_0.2.5-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pgexporter-ext` | `0.2.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.2 KiB | [postgresql-17-pgexporter-ext_0.2.5-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgexporter-ext/postgresql-17-pgexporter-ext_0.2.5-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgexporter_ext_16` | `0.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.9 KiB | [pgexporter_ext_16-0.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgexporter_ext_16-0.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgexporter_ext_16` | `0.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.8 KiB | [pgexporter_ext_16-0.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgexporter_ext_16-0.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgexporter_ext_16` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.4 KiB | [pgexporter_ext_16-0.2.4-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgexporter_ext_16-0.2.4-2PGDG.rhel9.8.x86_64.rpm) |
| `pgexporter_ext_16` | `0.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.6 KiB | [pgexporter_ext_16-0.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgexporter_ext_16-0.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgexporter_ext_16` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [pgexporter_ext_16-0.2.4-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgexporter_ext_16-0.2.4-2PGDG.rhel9.8.aarch64.rpm) |
| `pgexporter_ext_16` | `0.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.2 KiB | [pgexporter_ext_16-0.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgexporter_ext_16-0.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgexporter_ext_16` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.3 KiB | [pgexporter_ext_16-0.2.4-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgexporter_ext_16-0.2.4-2PGDG.rhel10.2.x86_64.rpm) |
| `pgexporter_ext_16` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.8 KiB | [pgexporter_ext_16-0.2.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgexporter_ext_16-0.2.4-1PGDG.rhel10.x86_64.rpm) |
| `pgexporter_ext_16` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.8 KiB | [pgexporter_ext_16-0.2.4-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgexporter_ext_16-0.2.4-2PGDG.rhel10.2.aarch64.rpm) |
| `pgexporter_ext_16` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pgexporter_ext_16-0.2.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgexporter_ext_16-0.2.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgexporter-ext` | `0.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.7 KiB | [postgresql-16-pgexporter-ext_0.2.5-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgexporter-ext/postgresql-16-pgexporter-ext_0.2.5-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgexporter-ext` | `0.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.5 KiB | [postgresql-16-pgexporter-ext_0.2.5-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgexporter-ext/postgresql-16-pgexporter-ext_0.2.5-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgexporter-ext` | `0.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.8 KiB | [postgresql-16-pgexporter-ext_0.2.5-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgexporter-ext/postgresql-16-pgexporter-ext_0.2.5-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pgexporter-ext` | `0.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.5 KiB | [postgresql-16-pgexporter-ext_0.2.5-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgexporter-ext/postgresql-16-pgexporter-ext_0.2.5-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pgexporter-ext` | `0.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.5 KiB | [postgresql-16-pgexporter-ext_0.2.5-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgexporter-ext/postgresql-16-pgexporter-ext_0.2.5-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pgexporter-ext` | `0.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.3 KiB | [postgresql-16-pgexporter-ext_0.2.5-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgexporter-ext/postgresql-16-pgexporter-ext_0.2.5-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pgexporter-ext` | `0.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.3 KiB | [postgresql-16-pgexporter-ext_0.2.5-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgexporter-ext/postgresql-16-pgexporter-ext_0.2.5-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pgexporter-ext` | `0.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.3 KiB | [postgresql-16-pgexporter-ext_0.2.5-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgexporter-ext/postgresql-16-pgexporter-ext_0.2.5-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pgexporter-ext` | `0.2.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.3 KiB | [postgresql-16-pgexporter-ext_0.2.5-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgexporter-ext/postgresql-16-pgexporter-ext_0.2.5-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pgexporter-ext` | `0.2.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.2 KiB | [postgresql-16-pgexporter-ext_0.2.5-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgexporter-ext/postgresql-16-pgexporter-ext_0.2.5-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgexporter_ext_15` | `0.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.9 KiB | [pgexporter_ext_15-0.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgexporter_ext_15-0.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgexporter_ext_15` | `0.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.1 KiB | [pgexporter_ext_15-0.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgexporter_ext_15-0.1.2-1.rhel8.x86_64.rpm) |
| `pgexporter_ext_15` | `0.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.8 KiB | [pgexporter_ext_15-0.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgexporter_ext_15-0.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgexporter_ext_15` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.4 KiB | [pgexporter_ext_15-0.2.4-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgexporter_ext_15-0.2.4-2PGDG.rhel9.8.x86_64.rpm) |
| `pgexporter_ext_15` | `0.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.6 KiB | [pgexporter_ext_15-0.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgexporter_ext_15-0.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgexporter_ext_15` | `0.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.4 KiB | [pgexporter_ext_15-0.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgexporter_ext_15-0.1.2-1.rhel9.x86_64.rpm) |
| `pgexporter_ext_15` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [pgexporter_ext_15-0.2.4-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgexporter_ext_15-0.2.4-2PGDG.rhel9.8.aarch64.rpm) |
| `pgexporter_ext_15` | `0.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.2 KiB | [pgexporter_ext_15-0.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgexporter_ext_15-0.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgexporter_ext_15` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.3 KiB | [pgexporter_ext_15-0.2.4-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgexporter_ext_15-0.2.4-2PGDG.rhel10.2.x86_64.rpm) |
| `pgexporter_ext_15` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.8 KiB | [pgexporter_ext_15-0.2.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgexporter_ext_15-0.2.4-1PGDG.rhel10.x86_64.rpm) |
| `pgexporter_ext_15` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.8 KiB | [pgexporter_ext_15-0.2.4-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgexporter_ext_15-0.2.4-2PGDG.rhel10.2.aarch64.rpm) |
| `pgexporter_ext_15` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pgexporter_ext_15-0.2.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgexporter_ext_15-0.2.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgexporter-ext` | `0.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.7 KiB | [postgresql-15-pgexporter-ext_0.2.5-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgexporter-ext/postgresql-15-pgexporter-ext_0.2.5-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgexporter-ext` | `0.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.5 KiB | [postgresql-15-pgexporter-ext_0.2.5-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgexporter-ext/postgresql-15-pgexporter-ext_0.2.5-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgexporter-ext` | `0.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.8 KiB | [postgresql-15-pgexporter-ext_0.2.5-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgexporter-ext/postgresql-15-pgexporter-ext_0.2.5-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pgexporter-ext` | `0.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.5 KiB | [postgresql-15-pgexporter-ext_0.2.5-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgexporter-ext/postgresql-15-pgexporter-ext_0.2.5-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pgexporter-ext` | `0.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.4 KiB | [postgresql-15-pgexporter-ext_0.2.5-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgexporter-ext/postgresql-15-pgexporter-ext_0.2.5-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pgexporter-ext` | `0.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.3 KiB | [postgresql-15-pgexporter-ext_0.2.5-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgexporter-ext/postgresql-15-pgexporter-ext_0.2.5-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pgexporter-ext` | `0.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.3 KiB | [postgresql-15-pgexporter-ext_0.2.5-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgexporter-ext/postgresql-15-pgexporter-ext_0.2.5-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pgexporter-ext` | `0.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.3 KiB | [postgresql-15-pgexporter-ext_0.2.5-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgexporter-ext/postgresql-15-pgexporter-ext_0.2.5-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pgexporter-ext` | `0.2.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.3 KiB | [postgresql-15-pgexporter-ext_0.2.5-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgexporter-ext/postgresql-15-pgexporter-ext_0.2.5-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pgexporter-ext` | `0.2.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.2 KiB | [postgresql-15-pgexporter-ext_0.2.5-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgexporter-ext/postgresql-15-pgexporter-ext_0.2.5-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgexporter_ext_14` | `0.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.9 KiB | [pgexporter_ext_14-0.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgexporter_ext_14-0.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgexporter_ext_14` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.1 KiB | [pgexporter_ext_14-0.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgexporter_ext_14-0.2.0-1.rhel8.x86_64.rpm) |
| `pgexporter_ext_14` | `0.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 16.1 KiB | [pgexporter_ext_14-0.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgexporter_ext_14-0.1.2-1.rhel8.x86_64.rpm) |
| `pgexporter_ext_14` | `0.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.9 KiB | [pgexporter_ext_14-0.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgexporter_ext_14-0.1.1-1.rhel8.x86_64.rpm) |
| `pgexporter_ext_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 15.6 KiB | [pgexporter_ext_14-0.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgexporter_ext_14-0.1.0-1.rhel8.x86_64.rpm) |
| `pgexporter_ext_14` | `0.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.8 KiB | [pgexporter_ext_14-0.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgexporter_ext_14-0.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgexporter_ext_14` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.0 KiB | [pgexporter_ext_14-0.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgexporter_ext_14-0.2.0-1.rhel8.aarch64.rpm) |
| `pgexporter_ext_14` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.4 KiB | [pgexporter_ext_14-0.2.4-2PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgexporter_ext_14-0.2.4-2PGDG.rhel9.8.x86_64.rpm) |
| `pgexporter_ext_14` | `0.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.5 KiB | [pgexporter_ext_14-0.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgexporter_ext_14-0.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgexporter_ext_14` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.8 KiB | [pgexporter_ext_14-0.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgexporter_ext_14-0.2.0-1.rhel9.x86_64.rpm) |
| `pgexporter_ext_14` | `0.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.4 KiB | [pgexporter_ext_14-0.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgexporter_ext_14-0.1.2-1.rhel9.x86_64.rpm) |
| `pgexporter_ext_14` | `0.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 16.1 KiB | [pgexporter_ext_14-0.1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgexporter_ext_14-0.1.1-1.rhel9.x86_64.rpm) |
| `pgexporter_ext_14` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.3 KiB | [pgexporter_ext_14-0.2.4-2PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgexporter_ext_14-0.2.4-2PGDG.rhel9.8.aarch64.rpm) |
| `pgexporter_ext_14` | `0.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.2 KiB | [pgexporter_ext_14-0.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgexporter_ext_14-0.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgexporter_ext_14` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.5 KiB | [pgexporter_ext_14-0.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgexporter_ext_14-0.2.0-1.rhel9.aarch64.rpm) |
| `pgexporter_ext_14` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.3 KiB | [pgexporter_ext_14-0.2.4-2PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgexporter_ext_14-0.2.4-2PGDG.rhel10.2.x86_64.rpm) |
| `pgexporter_ext_14` | `0.2.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.8 KiB | [pgexporter_ext_14-0.2.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgexporter_ext_14-0.2.4-1PGDG.rhel10.x86_64.rpm) |
| `pgexporter_ext_14` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.8 KiB | [pgexporter_ext_14-0.2.4-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgexporter_ext_14-0.2.4-2PGDG.rhel10.2.aarch64.rpm) |
| `pgexporter_ext_14` | `0.2.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.4 KiB | [pgexporter_ext_14-0.2.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgexporter_ext_14-0.2.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgexporter-ext` | `0.2.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.6 KiB | [postgresql-14-pgexporter-ext_0.2.5-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgexporter-ext/postgresql-14-pgexporter-ext_0.2.5-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgexporter-ext` | `0.2.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.4 KiB | [postgresql-14-pgexporter-ext_0.2.5-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgexporter-ext/postgresql-14-pgexporter-ext_0.2.5-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgexporter-ext` | `0.2.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.7 KiB | [postgresql-14-pgexporter-ext_0.2.5-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgexporter-ext/postgresql-14-pgexporter-ext_0.2.5-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pgexporter-ext` | `0.2.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.4 KiB | [postgresql-14-pgexporter-ext_0.2.5-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgexporter-ext/postgresql-14-pgexporter-ext_0.2.5-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pgexporter-ext` | `0.2.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.5 KiB | [postgresql-14-pgexporter-ext_0.2.5-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgexporter-ext/postgresql-14-pgexporter-ext_0.2.5-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pgexporter-ext` | `0.2.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.3 KiB | [postgresql-14-pgexporter-ext_0.2.5-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgexporter-ext/postgresql-14-pgexporter-ext_0.2.5-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pgexporter-ext` | `0.2.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.3 KiB | [postgresql-14-pgexporter-ext_0.2.5-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgexporter-ext/postgresql-14-pgexporter-ext_0.2.5-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pgexporter-ext` | `0.2.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.2 KiB | [postgresql-14-pgexporter-ext_0.2.5-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgexporter-ext/postgresql-14-pgexporter-ext_0.2.5-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pgexporter-ext` | `0.2.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 12.3 KiB | [postgresql-14-pgexporter-ext_0.2.5-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgexporter-ext/postgresql-14-pgexporter-ext_0.2.5-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pgexporter-ext` | `0.2.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 12.2 KiB | [postgresql-14-pgexporter-ext_0.2.5-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgexporter-ext/postgresql-14-pgexporter-ext_0.2.5-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgexporter/pgexporter_ext" title="Repository" icon="github" subtitle="github.com/pgexporter/pgexporter_ext" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgexporter_ext-0.2.5.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgexporter_ext;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgexporter_ext;		# install via package name, for the active PG version

pig install pgexporter_ext -v 18;   # install for PG 18
pig install pgexporter_ext -v 17;   # install for PG 17
pig install pgexporter_ext -v 16;   # install for PG 16
pig install pgexporter_ext -v 15;   # install for PG 15
pig install pgexporter_ext -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgexporter_ext;
```

## Usage

Sources:

- [Version 0.2.5 README](https://github.com/pgexporter/pgexporter_ext/blob/0.2.5/README.md)
- [Version 0.2.5 setup guide](https://github.com/pgexporter/pgexporter_ext/blob/0.2.5/doc/GETTING_STARTED.md)
- [Version 0.2.5 SQL definitions](https://github.com/pgexporter/pgexporter_ext/tree/0.2.5/sql)
- [Version 0.2.3 SQL definitions](https://github.com/pgexporter/pgexporter_ext/tree/0.2.3/sql)
- [Version 0.2.4 SQL definitions](https://github.com/pgexporter/pgexporter_ext/tree/0.2.4/sql)
- [Version 0.2.5 filesystem implementation](https://github.com/pgexporter/pgexporter_ext/blob/0.2.5/src/pgexporter_ext/utils.c)

`pgexporter_ext` exposes Linux host and filesystem metrics through SQL for collection by pgexporter. The examples below apply to the packaged 0.2.3–0.2.5 APIs. Functions run with the PostgreSQL operating-system account's access, so restrict the monitoring role to trusted users.

### Setup and Core Workflow

The upstream setup guide configures the module in `shared_preload_libraries`, followed by a PostgreSQL restart. Preserve any other libraries already configured.

```ini
shared_preload_libraries = 'pgexporter_ext'
```

In the `postgres` database, a privileged administrator installs the extension and grants monitoring access to the existing exporter login:

```sql
CREATE EXTENSION pgexporter_ext;
GRANT pg_monitor TO pgexporter;

SET ROLE pgexporter;
SELECT pgexporter_version_ext();
SELECT * FROM pgexporter_get_functions();
SELECT * FROM pgexporter_os_info();
SELECT * FROM pgexporter_load_avg();
RESET ROLE;
```

The SQL scripts revoke public execution and grant it to `pg_monitor`. That predefined role also grants broad PostgreSQL monitoring access. The release includes a base installation script and an upgrade chain; PostgreSQL can use that chain when installing the default version with `CREATE EXTENSION`.

### Metric Functions

- `pgexporter_information_ext()` and `pgexporter_version_ext()` return extension information and version text.
- `pgexporter_get_functions()` lists metric names, whether they take input, descriptions, and types. `pgexporter_is_supported(text)` checks a metric name.
- `pgexporter_os_info()`, `pgexporter_cpu_info()`, `pgexporter_memory_info()`, `pgexporter_network_info()`, and `pgexporter_load_avg()` return host metrics.
- `pgexporter_used_space(text)`, `pgexporter_free_space(text)`, and `pgexporter_total_space(text)` return byte counts for a filesystem path.

```sql
SELECT pgexporter_is_supported('pgexporter_load_avg');
SELECT pgexporter_free_space('/var/lib/postgresql');
```

Choose a path that exists on the server and is accessible to its operating-system account. These releases do not provide the later FIPS or log-count APIs.

### Operational Boundaries

Filesystem calls use the server's filesystem view, which may be a container rather than the host. The used-space function walks directories, so large trees can make scrapes expensive. A `pg_monitor` member can inspect accessible paths; restrict database access and choose collection paths deliberately.

Upstream lists Linux and PostgreSQL 13+; this catalog supplies packages for PostgreSQL 14–18. Use the exact PostgreSQL-major package. The principal RPM supplier is PGDG at 0.2.4 (EL8 PostgreSQL 14–16 use 0.2.3), while Pigsty supplies DEB 0.2.5; enable the Pigsty repository as well when following the cross-platform installation instructions.
