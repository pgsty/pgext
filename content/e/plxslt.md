---
title: "plxslt"
linkTitle: "plxslt"
description: "XSLT procedural language for PostgreSQL"
weight: 3110
categories: ["LANG"]
width: full
---

[**plxslt**](https://github.com/petere/plxslt) : XSLT procedural language for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3110** | {{< badge content="plxslt" link="https://github.com/petere/plxslt" >}} | {{< ext "plxslt" >}} | `0.20140221` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "pgml" >}} {{< ext "plpython3u" >}} {{< ext "pg_tle" >}} {{< ext "plv8" >}} {{< ext "pljava" >}} {{< ext "plperl" >}} {{< ext "pllua" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.20140221` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `plxslt` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.20140221` | {{< bg "18" "plxslt_18" "green" >}} {{< bg "17" "plxslt_17" "green" >}} {{< bg "16" "plxslt_16" "green" >}} {{< bg "15" "plxslt_15" "green" >}} {{< bg "14" "plxslt_14" "green" >}} {{< bg "13" "plxslt_13" "green" >}} | `plxslt_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.20140221` | {{< bg "18" "postgresql-18-plxslt" "green" >}} {{< bg "17" "postgresql-17-plxslt" "green" >}} {{< bg "16" "postgresql-16-plxslt" "green" >}} {{< bg "15" "postgresql-15-plxslt" "green" >}} {{< bg "14" "postgresql-14-plxslt" "green" >}} {{< bg "13" "postgresql-13-plxslt" "green" >}} | `postgresql-$v-plxslt` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.20140221" "plxslt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.20140221" "plxslt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.20140221" "plxslt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.20140221" "plxslt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.20140221" "plxslt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.20140221" "plxslt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-18-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-17-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-16-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-15-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-14-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-13-plxslt : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-18-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-17-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-16-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-15-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-14-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-13-plxslt : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-18-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-17-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-16-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-15-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-14-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-13-plxslt : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-18-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-17-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-16-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-15-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-14-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-13-plxslt : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-18-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-17-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-16-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-15-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-14-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-13-plxslt : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-18-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-17-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-16-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-15-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-14-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-13-plxslt : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-18-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-17-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-16-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-15-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-14-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-13-plxslt : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-18-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-17-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-16-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-15-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-14-plxslt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.20140221" "postgresql-13-plxslt : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plxslt_18` | `0.20140221` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.7 KiB | [plxslt_18-0.20140221-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plxslt_18-0.20140221-1PGDG.rhel8.x86_64.rpm) |
| `plxslt_18` | `0.20140221` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.7 KiB | [plxslt_18-0.20140221-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plxslt_18-0.20140221-1PGDG.rhel8.aarch64.rpm) |
| `plxslt_18` | `0.20140221` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.9 KiB | [plxslt_18-0.20140221-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plxslt_18-0.20140221-1PGDG.rhel9.x86_64.rpm) |
| `plxslt_18` | `0.20140221` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.6 KiB | [plxslt_18-0.20140221-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plxslt_18-0.20140221-1PGDG.rhel9.aarch64.rpm) |
| `plxslt_18` | `0.20140221` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.3 KiB | [plxslt_18-0.20140221-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plxslt_18-0.20140221-1PGDG.rhel10.x86_64.rpm) |
| `plxslt_18` | `0.20140221` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.2 KiB | [plxslt_18-0.20140221-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plxslt_18-0.20140221-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-plxslt` | `0.20140221` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-18-plxslt_0.20140221-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plxslt/postgresql-18-plxslt_0.20140221-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-plxslt` | `0.20140221` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-18-plxslt_0.20140221-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plxslt/postgresql-18-plxslt_0.20140221-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-plxslt` | `0.20140221` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.9 KiB | [postgresql-18-plxslt_0.20140221-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plxslt/postgresql-18-plxslt_0.20140221-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-plxslt` | `0.20140221` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.9 KiB | [postgresql-18-plxslt_0.20140221-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plxslt/postgresql-18-plxslt_0.20140221-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-plxslt` | `0.20140221` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.3 KiB | [postgresql-18-plxslt_0.20140221-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plxslt/postgresql-18-plxslt_0.20140221-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-plxslt` | `0.20140221` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.6 KiB | [postgresql-18-plxslt_0.20140221-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plxslt/postgresql-18-plxslt_0.20140221-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-plxslt` | `0.20140221` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.4 KiB | [postgresql-18-plxslt_0.20140221-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plxslt/postgresql-18-plxslt_0.20140221-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-plxslt` | `0.20140221` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.4 KiB | [postgresql-18-plxslt_0.20140221-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plxslt/postgresql-18-plxslt_0.20140221-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plxslt_17` | `0.20140221` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.7 KiB | [plxslt_17-0.20140221-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plxslt_17-0.20140221-1PGDG.rhel8.x86_64.rpm) |
| `plxslt_17` | `0.20140221` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.7 KiB | [plxslt_17-0.20140221-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plxslt_17-0.20140221-1PGDG.rhel8.aarch64.rpm) |
| `plxslt_17` | `0.20140221` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.9 KiB | [plxslt_17-0.20140221-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plxslt_17-0.20140221-1PGDG.rhel9.x86_64.rpm) |
| `plxslt_17` | `0.20140221` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.7 KiB | [plxslt_17-0.20140221-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plxslt_17-0.20140221-1PGDG.rhel9.aarch64.rpm) |
| `plxslt_17` | `0.20140221` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.3 KiB | [plxslt_17-0.20140221-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plxslt_17-0.20140221-1PGDG.rhel10.x86_64.rpm) |
| `plxslt_17` | `0.20140221` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.2 KiB | [plxslt_17-0.20140221-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plxslt_17-0.20140221-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-plxslt` | `0.20140221` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-17-plxslt_0.20140221-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plxslt/postgresql-17-plxslt_0.20140221-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-plxslt` | `0.20140221` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-17-plxslt_0.20140221-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plxslt/postgresql-17-plxslt_0.20140221-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-plxslt` | `0.20140221` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.8 KiB | [postgresql-17-plxslt_0.20140221-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plxslt/postgresql-17-plxslt_0.20140221-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-plxslt` | `0.20140221` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.9 KiB | [postgresql-17-plxslt_0.20140221-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plxslt/postgresql-17-plxslt_0.20140221-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-plxslt` | `0.20140221` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.5 KiB | [postgresql-17-plxslt_0.20140221-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plxslt/postgresql-17-plxslt_0.20140221-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-plxslt` | `0.20140221` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.8 KiB | [postgresql-17-plxslt_0.20140221-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plxslt/postgresql-17-plxslt_0.20140221-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-plxslt` | `0.20140221` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.4 KiB | [postgresql-17-plxslt_0.20140221-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plxslt/postgresql-17-plxslt_0.20140221-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-plxslt` | `0.20140221` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.3 KiB | [postgresql-17-plxslt_0.20140221-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plxslt/postgresql-17-plxslt_0.20140221-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plxslt_16` | `0.20140221` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.7 KiB | [plxslt_16-0.20140221-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plxslt_16-0.20140221-1PGDG.rhel8.x86_64.rpm) |
| `plxslt_16` | `0.20140221` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.7 KiB | [plxslt_16-0.20140221-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plxslt_16-0.20140221-1PGDG.rhel8.aarch64.rpm) |
| `plxslt_16` | `0.20140221` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.9 KiB | [plxslt_16-0.20140221-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plxslt_16-0.20140221-1PGDG.rhel9.x86_64.rpm) |
| `plxslt_16` | `0.20140221` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.7 KiB | [plxslt_16-0.20140221-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plxslt_16-0.20140221-1PGDG.rhel9.aarch64.rpm) |
| `plxslt_16` | `0.20140221` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.3 KiB | [plxslt_16-0.20140221-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plxslt_16-0.20140221-1PGDG.rhel10.x86_64.rpm) |
| `plxslt_16` | `0.20140221` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.2 KiB | [plxslt_16-0.20140221-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plxslt_16-0.20140221-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-plxslt` | `0.20140221` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-16-plxslt_0.20140221-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plxslt/postgresql-16-plxslt_0.20140221-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-plxslt` | `0.20140221` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-16-plxslt_0.20140221-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plxslt/postgresql-16-plxslt_0.20140221-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-plxslt` | `0.20140221` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.8 KiB | [postgresql-16-plxslt_0.20140221-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plxslt/postgresql-16-plxslt_0.20140221-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-plxslt` | `0.20140221` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.9 KiB | [postgresql-16-plxslt_0.20140221-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plxslt/postgresql-16-plxslt_0.20140221-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-plxslt` | `0.20140221` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.5 KiB | [postgresql-16-plxslt_0.20140221-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plxslt/postgresql-16-plxslt_0.20140221-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-plxslt` | `0.20140221` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.8 KiB | [postgresql-16-plxslt_0.20140221-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plxslt/postgresql-16-plxslt_0.20140221-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-plxslt` | `0.20140221` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.4 KiB | [postgresql-16-plxslt_0.20140221-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plxslt/postgresql-16-plxslt_0.20140221-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-plxslt` | `0.20140221` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.3 KiB | [postgresql-16-plxslt_0.20140221-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plxslt/postgresql-16-plxslt_0.20140221-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plxslt_15` | `0.20140221` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.7 KiB | [plxslt_15-0.20140221-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plxslt_15-0.20140221-1PGDG.rhel8.x86_64.rpm) |
| `plxslt_15` | `0.20140221` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.7 KiB | [plxslt_15-0.20140221-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plxslt_15-0.20140221-1PGDG.rhel8.aarch64.rpm) |
| `plxslt_15` | `0.20140221` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.9 KiB | [plxslt_15-0.20140221-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plxslt_15-0.20140221-1PGDG.rhel9.x86_64.rpm) |
| `plxslt_15` | `0.20140221` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.7 KiB | [plxslt_15-0.20140221-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plxslt_15-0.20140221-1PGDG.rhel9.aarch64.rpm) |
| `plxslt_15` | `0.20140221` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.3 KiB | [plxslt_15-0.20140221-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plxslt_15-0.20140221-1PGDG.rhel10.x86_64.rpm) |
| `plxslt_15` | `0.20140221` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.2 KiB | [plxslt_15-0.20140221-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plxslt_15-0.20140221-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-plxslt` | `0.20140221` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-15-plxslt_0.20140221-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plxslt/postgresql-15-plxslt_0.20140221-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-plxslt` | `0.20140221` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-15-plxslt_0.20140221-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plxslt/postgresql-15-plxslt_0.20140221-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-plxslt` | `0.20140221` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.8 KiB | [postgresql-15-plxslt_0.20140221-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plxslt/postgresql-15-plxslt_0.20140221-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-plxslt` | `0.20140221` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.9 KiB | [postgresql-15-plxslt_0.20140221-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plxslt/postgresql-15-plxslt_0.20140221-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-plxslt` | `0.20140221` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.5 KiB | [postgresql-15-plxslt_0.20140221-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plxslt/postgresql-15-plxslt_0.20140221-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-plxslt` | `0.20140221` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.8 KiB | [postgresql-15-plxslt_0.20140221-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plxslt/postgresql-15-plxslt_0.20140221-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-plxslt` | `0.20140221` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.4 KiB | [postgresql-15-plxslt_0.20140221-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plxslt/postgresql-15-plxslt_0.20140221-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-plxslt` | `0.20140221` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.3 KiB | [postgresql-15-plxslt_0.20140221-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plxslt/postgresql-15-plxslt_0.20140221-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plxslt_14` | `0.20140221` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.7 KiB | [plxslt_14-0.20140221-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plxslt_14-0.20140221-1PGDG.rhel8.x86_64.rpm) |
| `plxslt_14` | `0.20140221` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.7 KiB | [plxslt_14-0.20140221-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plxslt_14-0.20140221-1PGDG.rhel8.aarch64.rpm) |
| `plxslt_14` | `0.20140221` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.9 KiB | [plxslt_14-0.20140221-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plxslt_14-0.20140221-1PGDG.rhel9.x86_64.rpm) |
| `plxslt_14` | `0.20140221` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.7 KiB | [plxslt_14-0.20140221-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plxslt_14-0.20140221-1PGDG.rhel9.aarch64.rpm) |
| `plxslt_14` | `0.20140221` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.3 KiB | [plxslt_14-0.20140221-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plxslt_14-0.20140221-1PGDG.rhel10.x86_64.rpm) |
| `plxslt_14` | `0.20140221` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.2 KiB | [plxslt_14-0.20140221-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plxslt_14-0.20140221-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-plxslt` | `0.20140221` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-14-plxslt_0.20140221-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plxslt/postgresql-14-plxslt_0.20140221-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-plxslt` | `0.20140221` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.8 KiB | [postgresql-14-plxslt_0.20140221-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plxslt/postgresql-14-plxslt_0.20140221-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-plxslt` | `0.20140221` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.8 KiB | [postgresql-14-plxslt_0.20140221-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plxslt/postgresql-14-plxslt_0.20140221-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-plxslt` | `0.20140221` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.8 KiB | [postgresql-14-plxslt_0.20140221-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plxslt/postgresql-14-plxslt_0.20140221-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-plxslt` | `0.20140221` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.5 KiB | [postgresql-14-plxslt_0.20140221-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plxslt/postgresql-14-plxslt_0.20140221-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-plxslt` | `0.20140221` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.8 KiB | [postgresql-14-plxslt_0.20140221-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plxslt/postgresql-14-plxslt_0.20140221-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-plxslt` | `0.20140221` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.4 KiB | [postgresql-14-plxslt_0.20140221-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plxslt/postgresql-14-plxslt_0.20140221-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-plxslt` | `0.20140221` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.3 KiB | [postgresql-14-plxslt_0.20140221-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plxslt/postgresql-14-plxslt_0.20140221-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plxslt_13` | `0.20140221` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.6 KiB | [plxslt_13-0.20140221-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/plxslt_13-0.20140221-1PGDG.rhel8.x86_64.rpm) |
| `plxslt_13` | `0.20140221` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.7 KiB | [plxslt_13-0.20140221-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/plxslt_13-0.20140221-1PGDG.rhel8.aarch64.rpm) |
| `plxslt_13` | `0.20140221` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.9 KiB | [plxslt_13-0.20140221-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/plxslt_13-0.20140221-1PGDG.rhel9.x86_64.rpm) |
| `plxslt_13` | `0.20140221` | [el9.aarch64](/os/el9.aarch64) | pgdg | 14.7 KiB | [plxslt_13-0.20140221-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/plxslt_13-0.20140221-1PGDG.rhel9.aarch64.rpm) |
| `plxslt_13` | `0.20140221` | [el10.x86_64](/os/el10.x86_64) | pgdg | 15.3 KiB | [plxslt_13-0.20140221-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/plxslt_13-0.20140221-1PGDG.rhel10.x86_64.rpm) |
| `plxslt_13` | `0.20140221` | [el10.aarch64](/os/el10.aarch64) | pgdg | 15.2 KiB | [plxslt_13-0.20140221-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/plxslt_13-0.20140221-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-plxslt` | `0.20140221` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.6 KiB | [postgresql-13-plxslt_0.20140221-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plxslt/postgresql-13-plxslt_0.20140221-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-plxslt` | `0.20140221` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.8 KiB | [postgresql-13-plxslt_0.20140221-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plxslt/postgresql-13-plxslt_0.20140221-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-plxslt` | `0.20140221` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.6 KiB | [postgresql-13-plxslt_0.20140221-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plxslt/postgresql-13-plxslt_0.20140221-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-plxslt` | `0.20140221` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.8 KiB | [postgresql-13-plxslt_0.20140221-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plxslt/postgresql-13-plxslt_0.20140221-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-plxslt` | `0.20140221` | [u22.x86_64](/os/u22.x86_64) | pigsty | 14.6 KiB | [postgresql-13-plxslt_0.20140221-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plxslt/postgresql-13-plxslt_0.20140221-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-plxslt` | `0.20140221` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.6 KiB | [postgresql-13-plxslt_0.20140221-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plxslt/postgresql-13-plxslt_0.20140221-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-plxslt` | `0.20140221` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.2 KiB | [postgresql-13-plxslt_0.20140221-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plxslt/postgresql-13-plxslt_0.20140221-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-plxslt` | `0.20140221` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.3 KiB | [postgresql-13-plxslt_0.20140221-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plxslt/postgresql-13-plxslt_0.20140221-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petere/plxslt" title="Repository" icon="github" subtitle="github.com/petere/plxslt" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plxslt-0.20140221.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg plxslt;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install plxslt;		# install via package name, for the active PG version

pig install plxslt -v 18;   # install for PG 18
pig install plxslt -v 17;   # install for PG 17
pig install plxslt -v 16;   # install for PG 16
pig install plxslt -v 15;   # install for PG 15
pig install plxslt -v 14;   # install for PG 14
pig install plxslt -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plxslt;
```
