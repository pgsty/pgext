---
title: "logerrors"
linkTitle: "logerrors"
description: "Function for collecting statistics about messages in logfile"
weight: 7140
categories: ["SEC"]
width: full
---

[**logerrors**](https://github.com/munakoiso/logerrors) : Function for collecting statistics about messages in logfile


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7140** | {{< badge content="logerrors" link="https://github.com/munakoiso/logerrors" >}} | {{< ext "logerrors" >}} | `2.1.5` | {{< category "SEC" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgauditlogtofile" >}} {{< ext "pg_auth_mon" >}} {{< ext "pg_jobmon" >}} {{< ext "pg_stat_monitor" >}} {{< ext "auto_explain" >}} {{< ext "pg_track_settings" >}} {{< ext "pgaudit" >}} {{< ext "pgsentinel" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `2.1.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `logerrors` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.1.5` | {{< bg "18" "logerrors_18*" "green" >}} {{< bg "17" "logerrors_17*" "green" >}} {{< bg "16" "logerrors_16*" "green" >}} {{< bg "15" "logerrors_15*" "green" >}} {{< bg "14" "logerrors_14*" "green" >}} {{< bg "13" "logerrors_13*" "green" >}} | `logerrors_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.5` | {{< bg "18" "postgresql-18-logerrors" "green" >}} {{< bg "17" "postgresql-17-logerrors" "green" >}} {{< bg "16" "postgresql-16-logerrors" "green" >}} {{< bg "15" "postgresql-15-logerrors" "green" >}} {{< bg "14" "postgresql-14-logerrors" "green" >}} {{< bg "13" "postgresql-13-logerrors" "green" >}} | `postgresql-$v-logerrors` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.1.5" "logerrors_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_14 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_13 : AVAIL 7" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.1.5" "logerrors_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_14 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_13 : AVAIL 7" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.1.5" "logerrors_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_13 : AVAIL 6" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.1.5" "logerrors_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_14 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_13 : AVAIL 7" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.1.5" "logerrors_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.1.5" "logerrors_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.5" "logerrors_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-14-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-13-logerrors : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-14-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-13-logerrors : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-14-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-13-logerrors : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-14-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-13-logerrors : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-14-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-13-logerrors : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-14-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-13-logerrors : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-14-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-13-logerrors : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-14-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-13-logerrors : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logerrors_18` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.3 KiB | [logerrors_18-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/logerrors_18-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_18` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [logerrors_18-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/logerrors_18-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_18` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.7 KiB | [logerrors_18-2.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/logerrors_18-2.1.5-1PGDG.rhel9.x86_64.rpm) |
| `logerrors_18` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.9 KiB | [logerrors_18-2.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/logerrors_18-2.1.5-1PGDG.rhel9.aarch64.rpm) |
| `logerrors_18` | `2.1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.2 KiB | [logerrors_18-2.1.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/logerrors_18-2.1.5-1PGDG.rhel10.x86_64.rpm) |
| `logerrors_18` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.4 KiB | [logerrors_18-2.1.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/logerrors_18-2.1.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-logerrors` | `2.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.8 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.3 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.8 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.5 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 33.5 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 33.5 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.4 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 32.2 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logerrors_17` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.3 KiB | [logerrors_17-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/logerrors_17-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_17` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.9 KiB | [logerrors_17-2.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/logerrors_17-2.1.3-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_17` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.5 KiB | [logerrors_17-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/logerrors_17-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_17` | `2.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.1 KiB | [logerrors_17-2.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/logerrors_17-2.1.3-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_17` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.7 KiB | [logerrors_17-2.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/logerrors_17-2.1.5-1PGDG.rhel9.x86_64.rpm) |
| `logerrors_17` | `2.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.3 KiB | [logerrors_17-2.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/logerrors_17-2.1.3-1PGDG.rhel9.x86_64.rpm) |
| `logerrors_17` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.9 KiB | [logerrors_17-2.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/logerrors_17-2.1.5-1PGDG.rhel9.aarch64.rpm) |
| `logerrors_17` | `2.1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.4 KiB | [logerrors_17-2.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/logerrors_17-2.1.3-1PGDG.rhel9.aarch64.rpm) |
| `logerrors_17` | `2.1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.2 KiB | [logerrors_17-2.1.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/logerrors_17-2.1.5-1PGDG.rhel10.x86_64.rpm) |
| `logerrors_17` | `2.1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.6 KiB | [logerrors_17-2.1.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/logerrors_17-2.1.3-2PGDG.rhel10.x86_64.rpm) |
| `logerrors_17` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.5 KiB | [logerrors_17-2.1.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/logerrors_17-2.1.5-1PGDG.rhel10.aarch64.rpm) |
| `logerrors_17` | `2.1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.9 KiB | [logerrors_17-2.1.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/logerrors_17-2.1.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-logerrors` | `2.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.8 KiB | [postgresql-17-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-17-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-logerrors` | `2.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.3 KiB | [postgresql-17-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-17-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-logerrors` | `2.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.8 KiB | [postgresql-17-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-17-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-logerrors` | `2.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.5 KiB | [postgresql-17-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-17-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-logerrors` | `2.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 38.7 KiB | [postgresql-17-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-17-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-logerrors` | `2.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 38.3 KiB | [postgresql-17-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-17-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-logerrors` | `2.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.4 KiB | [postgresql-17-logerrors_2.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-17-logerrors_2.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-logerrors` | `2.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 32.2 KiB | [postgresql-17-logerrors_2.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-17-logerrors_2.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logerrors_16` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.3 KiB | [logerrors_16-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/logerrors_16-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_16` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.9 KiB | [logerrors_16-2.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/logerrors_16-2.1.3-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_16` | `2.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.1 KiB | [logerrors_16-2.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/logerrors_16-2.1.2-1.rhel8.x86_64.rpm) |
| `logerrors_16` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.5 KiB | [logerrors_16-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/logerrors_16-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_16` | `2.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.1 KiB | [logerrors_16-2.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/logerrors_16-2.1.3-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_16` | `2.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.2 KiB | [logerrors_16-2.1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/logerrors_16-2.1.2-1.rhel8.aarch64.rpm) |
| `logerrors_16` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.7 KiB | [logerrors_16-2.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/logerrors_16-2.1.5-1PGDG.rhel9.x86_64.rpm) |
| `logerrors_16` | `2.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.3 KiB | [logerrors_16-2.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/logerrors_16-2.1.3-1PGDG.rhel9.x86_64.rpm) |
| `logerrors_16` | `2.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.0 KiB | [logerrors_16-2.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/logerrors_16-2.1.2-1.rhel9.x86_64.rpm) |
| `logerrors_16` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.9 KiB | [logerrors_16-2.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/logerrors_16-2.1.5-1PGDG.rhel9.aarch64.rpm) |
| `logerrors_16` | `2.1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.4 KiB | [logerrors_16-2.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/logerrors_16-2.1.3-1PGDG.rhel9.aarch64.rpm) |
| `logerrors_16` | `2.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.8 KiB | [logerrors_16-2.1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/logerrors_16-2.1.2-1.rhel9.aarch64.rpm) |
| `logerrors_16` | `2.1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.2 KiB | [logerrors_16-2.1.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/logerrors_16-2.1.5-1PGDG.rhel10.x86_64.rpm) |
| `logerrors_16` | `2.1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.6 KiB | [logerrors_16-2.1.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/logerrors_16-2.1.3-2PGDG.rhel10.x86_64.rpm) |
| `logerrors_16` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.6 KiB | [logerrors_16-2.1.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/logerrors_16-2.1.5-1PGDG.rhel10.aarch64.rpm) |
| `logerrors_16` | `2.1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.9 KiB | [logerrors_16-2.1.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/logerrors_16-2.1.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-logerrors` | `2.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.8 KiB | [postgresql-16-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-16-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-logerrors` | `2.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.3 KiB | [postgresql-16-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-16-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-logerrors` | `2.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.8 KiB | [postgresql-16-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-16-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-logerrors` | `2.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.5 KiB | [postgresql-16-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-16-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-logerrors` | `2.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 38.3 KiB | [postgresql-16-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-16-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-logerrors` | `2.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.9 KiB | [postgresql-16-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-16-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-logerrors` | `2.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.4 KiB | [postgresql-16-logerrors_2.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-16-logerrors_2.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-logerrors` | `2.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 32.2 KiB | [postgresql-16-logerrors_2.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-16-logerrors_2.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logerrors_15` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.4 KiB | [logerrors_15-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.1 KiB | [logerrors_15-2.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.1.3-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.3 KiB | [logerrors_15-2.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.1.2-1.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.1 KiB | [logerrors_15-2.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.1-2.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.7 KiB | [logerrors_15-2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.1-1.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.9 KiB | [logerrors_15-2.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.0-2.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [logerrors_15-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.0-1.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.6 KiB | [logerrors_15-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.2 KiB | [logerrors_15-2.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.1.3-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.4 KiB | [logerrors_15-2.1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.1.2-1.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [logerrors_15-2.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.1-2.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.8 KiB | [logerrors_15-2.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.1-1.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.9 KiB | [logerrors_15-2.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.0-2.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.7 KiB | [logerrors_15-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.0-1.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.9 KiB | [logerrors_15-2.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/logerrors_15-2.1.5-1PGDG.rhel9.x86_64.rpm) |
| `logerrors_15` | `2.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.5 KiB | [logerrors_15-2.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/logerrors_15-2.1.3-1PGDG.rhel9.x86_64.rpm) |
| `logerrors_15` | `2.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.7 KiB | [logerrors_15-2.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/logerrors_15-2.1.2-1.rhel9.x86_64.rpm) |
| `logerrors_15` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.6 KiB | [logerrors_15-2.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/logerrors_15-2.1-2.rhel9.x86_64.rpm) |
| `logerrors_15` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 KiB | [logerrors_15-2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/logerrors_15-2.1-1.rhel9.x86_64.rpm) |
| `logerrors_15` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.1 KiB | [logerrors_15-2.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/logerrors_15-2.0-2.rhel9.x86_64.rpm) |
| `logerrors_15` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 42.8 KiB | [logerrors_15-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/logerrors_15-2.0-1.rhel9.x86_64.rpm) |
| `logerrors_15` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.0 KiB | [logerrors_15-2.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/logerrors_15-2.1.5-1PGDG.rhel9.aarch64.rpm) |
| `logerrors_15` | `2.1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.8 KiB | [logerrors_15-2.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/logerrors_15-2.1.3-1PGDG.rhel9.aarch64.rpm) |
| `logerrors_15` | `2.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.6 KiB | [logerrors_15-2.1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/logerrors_15-2.1.2-1.rhel9.aarch64.rpm) |
| `logerrors_15` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.5 KiB | [logerrors_15-2.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/logerrors_15-2.1-2.rhel9.aarch64.rpm) |
| `logerrors_15` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.9 KiB | [logerrors_15-2.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/logerrors_15-2.1-1.rhel9.aarch64.rpm) |
| `logerrors_15` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.0 KiB | [logerrors_15-2.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/logerrors_15-2.0-2.rhel9.aarch64.rpm) |
| `logerrors_15` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 42.4 KiB | [logerrors_15-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/logerrors_15-2.0-1.rhel9.aarch64.rpm) |
| `logerrors_15` | `2.1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.4 KiB | [logerrors_15-2.1.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/logerrors_15-2.1.5-1PGDG.rhel10.x86_64.rpm) |
| `logerrors_15` | `2.1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.8 KiB | [logerrors_15-2.1.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/logerrors_15-2.1.3-2PGDG.rhel10.x86_64.rpm) |
| `logerrors_15` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [logerrors_15-2.1.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/logerrors_15-2.1.5-1PGDG.rhel10.aarch64.rpm) |
| `logerrors_15` | `2.1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [logerrors_15-2.1.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/logerrors_15-2.1.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-logerrors` | `2.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.9 KiB | [postgresql-15-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-15-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-logerrors` | `2.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.5 KiB | [postgresql-15-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-15-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-logerrors` | `2.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 31.0 KiB | [postgresql-15-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-15-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-logerrors` | `2.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.7 KiB | [postgresql-15-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-15-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-logerrors` | `2.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 38.5 KiB | [postgresql-15-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-15-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-logerrors` | `2.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 38.1 KiB | [postgresql-15-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-15-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-logerrors` | `2.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.6 KiB | [postgresql-15-logerrors_2.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-15-logerrors_2.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-logerrors` | `2.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 32.4 KiB | [postgresql-15-logerrors_2.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-15-logerrors_2.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logerrors_14` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.3 KiB | [logerrors_14-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/logerrors_14-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_14` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.0 KiB | [logerrors_14-2.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/logerrors_14-2.1.3-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_14` | `2.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.2 KiB | [logerrors_14-2.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/logerrors_14-2.1.2-1.rhel8.x86_64.rpm) |
| `logerrors_14` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.1 KiB | [logerrors_14-2.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/logerrors_14-2.1-2.rhel8.x86_64.rpm) |
| `logerrors_14` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.7 KiB | [logerrors_14-2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/logerrors_14-2.1-1.rhel8.x86_64.rpm) |
| `logerrors_14` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.8 KiB | [logerrors_14-2.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/logerrors_14-2.0-2.rhel8.x86_64.rpm) |
| `logerrors_14` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.2 KiB | [logerrors_14-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/logerrors_14-2.0-1.rhel8.x86_64.rpm) |
| `logerrors_14` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.5 KiB | [logerrors_14-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/logerrors_14-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_14` | `2.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.1 KiB | [logerrors_14-2.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/logerrors_14-2.1.3-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_14` | `2.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.3 KiB | [logerrors_14-2.1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/logerrors_14-2.1.2-1.rhel8.aarch64.rpm) |
| `logerrors_14` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.1 KiB | [logerrors_14-2.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/logerrors_14-2.1-2.rhel8.aarch64.rpm) |
| `logerrors_14` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.7 KiB | [logerrors_14-2.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/logerrors_14-2.1-1.rhel8.aarch64.rpm) |
| `logerrors_14` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.9 KiB | [logerrors_14-2.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/logerrors_14-2.0-2.rhel8.aarch64.rpm) |
| `logerrors_14` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.9 KiB | [logerrors_14-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/logerrors_14-2.0-1.rhel8.aarch64.rpm) |
| `logerrors_14` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.8 KiB | [logerrors_14-2.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/logerrors_14-2.1.5-1PGDG.rhel9.x86_64.rpm) |
| `logerrors_14` | `2.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.5 KiB | [logerrors_14-2.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/logerrors_14-2.1.3-1PGDG.rhel9.x86_64.rpm) |
| `logerrors_14` | `2.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.6 KiB | [logerrors_14-2.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/logerrors_14-2.1.2-1.rhel9.x86_64.rpm) |
| `logerrors_14` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.5 KiB | [logerrors_14-2.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/logerrors_14-2.1-2.rhel9.x86_64.rpm) |
| `logerrors_14` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.1 KiB | [logerrors_14-2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/logerrors_14-2.1-1.rhel9.x86_64.rpm) |
| `logerrors_14` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.1 KiB | [logerrors_14-2.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/logerrors_14-2.0-2.rhel9.x86_64.rpm) |
| `logerrors_14` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.0 KiB | [logerrors_14-2.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/logerrors_14-2.1.5-1PGDG.rhel9.aarch64.rpm) |
| `logerrors_14` | `2.1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.8 KiB | [logerrors_14-2.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/logerrors_14-2.1.3-1PGDG.rhel9.aarch64.rpm) |
| `logerrors_14` | `2.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.6 KiB | [logerrors_14-2.1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/logerrors_14-2.1.2-1.rhel9.aarch64.rpm) |
| `logerrors_14` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.3 KiB | [logerrors_14-2.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/logerrors_14-2.1-2.rhel9.aarch64.rpm) |
| `logerrors_14` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.1 KiB | [logerrors_14-2.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/logerrors_14-2.1-1.rhel9.aarch64.rpm) |
| `logerrors_14` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.2 KiB | [logerrors_14-2.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/logerrors_14-2.0-2.rhel9.aarch64.rpm) |
| `logerrors_14` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.5 KiB | [logerrors_14-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/logerrors_14-2.0-1.rhel9.aarch64.rpm) |
| `logerrors_14` | `2.1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.4 KiB | [logerrors_14-2.1.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/logerrors_14-2.1.5-1PGDG.rhel10.x86_64.rpm) |
| `logerrors_14` | `2.1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.7 KiB | [logerrors_14-2.1.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/logerrors_14-2.1.3-2PGDG.rhel10.x86_64.rpm) |
| `logerrors_14` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [logerrors_14-2.1.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/logerrors_14-2.1.5-1PGDG.rhel10.aarch64.rpm) |
| `logerrors_14` | `2.1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [logerrors_14-2.1.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/logerrors_14-2.1.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-logerrors` | `2.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.7 KiB | [postgresql-14-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-14-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-logerrors` | `2.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.5 KiB | [postgresql-14-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-14-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-logerrors` | `2.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.7 KiB | [postgresql-14-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-14-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-logerrors` | `2.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.6 KiB | [postgresql-14-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-14-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-logerrors` | `2.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.6 KiB | [postgresql-14-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-14-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-logerrors` | `2.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.8 KiB | [postgresql-14-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-14-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-logerrors` | `2.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.3 KiB | [postgresql-14-logerrors_2.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-14-logerrors_2.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-logerrors` | `2.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 32.3 KiB | [postgresql-14-logerrors_2.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-14-logerrors_2.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logerrors_13` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.3 KiB | [logerrors_13-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/logerrors_13-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_13` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.9 KiB | [logerrors_13-2.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/logerrors_13-2.1.3-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_13` | `2.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.1 KiB | [logerrors_13-2.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/logerrors_13-2.1.2-1.rhel8.x86_64.rpm) |
| `logerrors_13` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.0 KiB | [logerrors_13-2.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/logerrors_13-2.1-2.rhel8.x86_64.rpm) |
| `logerrors_13` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.6 KiB | [logerrors_13-2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/logerrors_13-2.1-1.rhel8.x86_64.rpm) |
| `logerrors_13` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.7 KiB | [logerrors_13-2.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/logerrors_13-2.0-2.rhel8.x86_64.rpm) |
| `logerrors_13` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.8 KiB | [logerrors_13-1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/logerrors_13-1.1-1.rhel8.x86_64.rpm) |
| `logerrors_13` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.5 KiB | [logerrors_13-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/logerrors_13-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_13` | `2.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.2 KiB | [logerrors_13-2.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/logerrors_13-2.1.3-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_13` | `2.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.3 KiB | [logerrors_13-2.1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/logerrors_13-2.1.2-1.rhel8.aarch64.rpm) |
| `logerrors_13` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.1 KiB | [logerrors_13-2.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/logerrors_13-2.1-2.rhel8.aarch64.rpm) |
| `logerrors_13` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.8 KiB | [logerrors_13-2.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/logerrors_13-2.1-1.rhel8.aarch64.rpm) |
| `logerrors_13` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.9 KiB | [logerrors_13-2.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/logerrors_13-2.0-2.rhel8.aarch64.rpm) |
| `logerrors_13` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.0 KiB | [logerrors_13-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/logerrors_13-2.0-1.rhel8.aarch64.rpm) |
| `logerrors_13` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.8 KiB | [logerrors_13-2.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/logerrors_13-2.1.5-1PGDG.rhel9.x86_64.rpm) |
| `logerrors_13` | `2.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.4 KiB | [logerrors_13-2.1.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/logerrors_13-2.1.3-1PGDG.rhel9.x86_64.rpm) |
| `logerrors_13` | `2.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.6 KiB | [logerrors_13-2.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/logerrors_13-2.1.2-1.rhel9.x86_64.rpm) |
| `logerrors_13` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.5 KiB | [logerrors_13-2.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/logerrors_13-2.1-2.rhel9.x86_64.rpm) |
| `logerrors_13` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.1 KiB | [logerrors_13-2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/logerrors_13-2.1-1.rhel9.x86_64.rpm) |
| `logerrors_13` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.1 KiB | [logerrors_13-2.0-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/logerrors_13-2.0-2.rhel9.x86_64.rpm) |
| `logerrors_13` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.0 KiB | [logerrors_13-2.1.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/logerrors_13-2.1.5-1PGDG.rhel9.aarch64.rpm) |
| `logerrors_13` | `2.1.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.8 KiB | [logerrors_13-2.1.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/logerrors_13-2.1.3-1PGDG.rhel9.aarch64.rpm) |
| `logerrors_13` | `2.1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.6 KiB | [logerrors_13-2.1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/logerrors_13-2.1.2-1.rhel9.aarch64.rpm) |
| `logerrors_13` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.4 KiB | [logerrors_13-2.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/logerrors_13-2.1-2.rhel9.aarch64.rpm) |
| `logerrors_13` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.9 KiB | [logerrors_13-2.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/logerrors_13-2.1-1.rhel9.aarch64.rpm) |
| `logerrors_13` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.1 KiB | [logerrors_13-2.0-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/logerrors_13-2.0-2.rhel9.aarch64.rpm) |
| `logerrors_13` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.8 KiB | [logerrors_13-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/logerrors_13-2.0-1.rhel9.aarch64.rpm) |
| `logerrors_13` | `2.1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.3 KiB | [logerrors_13-2.1.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/logerrors_13-2.1.5-1PGDG.rhel10.x86_64.rpm) |
| `logerrors_13` | `2.1.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.7 KiB | [logerrors_13-2.1.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/logerrors_13-2.1.3-2PGDG.rhel10.x86_64.rpm) |
| `logerrors_13` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [logerrors_13-2.1.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/logerrors_13-2.1.5-1PGDG.rhel10.aarch64.rpm) |
| `logerrors_13` | `2.1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [logerrors_13-2.1.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/logerrors_13-2.1.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-logerrors` | `2.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.4 KiB | [postgresql-13-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-13-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-logerrors` | `2.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.5 KiB | [postgresql-13-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-13-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-logerrors` | `2.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.6 KiB | [postgresql-13-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-13-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-logerrors` | `2.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.6 KiB | [postgresql-13-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-13-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-logerrors` | `2.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 36.1 KiB | [postgresql-13-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-13-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-logerrors` | `2.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 36.5 KiB | [postgresql-13-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-13-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-logerrors` | `2.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.9 KiB | [postgresql-13-logerrors_2.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-13-logerrors_2.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-logerrors` | `2.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 32.6 KiB | [postgresql-13-logerrors_2.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-13-logerrors_2.1.5-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/munakoiso/logerrors" title="Repository" icon="github" subtitle="github.com/munakoiso/logerrors" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="logerrors-2.1.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg logerrors;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install logerrors;		# install via package name, for the active PG version

pig install logerrors -v 18;   # install for PG 18
pig install logerrors -v 17;   # install for PG 17
pig install logerrors -v 16;   # install for PG 16
pig install logerrors -v 15;   # install for PG 15
pig install logerrors -v 14;   # install for PG 14
pig install logerrors -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION logerrors;
```
