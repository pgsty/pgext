---
title: "logerrors"
linkTitle: "logerrors"
description: "Function for collecting statistics about messages in logfile"
weight: 7140
categories: ["SEC"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["MIXED"]
page_width: full
---

[**logerrors**](https://github.com/munakoiso/logerrors) : Function for collecting statistics about messages in logfile


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7140** | {{< badge content="logerrors" link="https://github.com/munakoiso/logerrors" >}} | {{< ext "logerrors" >}} | `2.1.6` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_stat_log" >}} {{< ext "pgelog" >}} {{< ext "pg_sqlog" >}} {{< ext "log_fdw" >}} {{< ext "pgaudit" >}} {{< ext "pgdisablelogerror" >}} {{< ext "pg_auth_mon" >}} {{< ext "pgfr_record" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `2.1.6` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `logerrors` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.1.6` | {{< bg "19" "logerrors_19" "green" >}} {{< bg "18" "logerrors_18" "green" >}} {{< bg "17" "logerrors_17" "green" >}} {{< bg "16" "logerrors_16" "green" >}} {{< bg "15" "logerrors_15" "green" >}} | `logerrors_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.6` | {{< bg "19" "postgresql-19-logerrors" "green" >}} {{< bg "18" "postgresql-18-logerrors" "green" >}} {{< bg "17" "postgresql-17-logerrors" "green" >}} {{< bg "16" "postgresql-16-logerrors" "green" >}} {{< bg "15" "postgresql-15-logerrors" "green" >}} | `postgresql-$v-logerrors` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "logerrors_19 : N/A 0" "gray" >}} | {{< bg "PGDG 2.1.6" "logerrors_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_15 : AVAIL 9" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "logerrors_19 : N/A 0" "gray" >}} | {{< bg "PGDG 2.1.6" "logerrors_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_15 : AVAIL 9" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.1.6" "logerrors_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_15 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.1.6" "logerrors_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.1.6" "logerrors_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.1.6" "logerrors_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.1.6" "logerrors_15 : AVAIL 7" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "MISS" "postgresql-19-logerrors : MISS 0" "red" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "MISS" "postgresql-19-logerrors : MISS 0" "red" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "MISS" "postgresql-19-logerrors : MISS 0" "red" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "MISS" "postgresql-19-logerrors : MISS 0" "red" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "MISS" "postgresql-19-logerrors : MISS 0" "red" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "MISS" "postgresql-19-logerrors : MISS 0" "red" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "MISS" "postgresql-19-logerrors : MISS 0" "red" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "MISS" "postgresql-19-logerrors : MISS 0" "red" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "MISS" "postgresql-19-logerrors : MISS 0" "red" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "MISS" "postgresql-19-logerrors : MISS 0" "red" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-18-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-17-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-16-logerrors : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.5" "postgresql-15-logerrors : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logerrors_19` | `2.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.6 KiB | [logerrors_19-2.1.6-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/logerrors_19-2.1.6-1PGDG.rhel9.9.x86_64.rpm) |
| `logerrors_19` | `2.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [logerrors_19-2.1.6-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/logerrors_19-2.1.6-1PGDG.rhel9.9.aarch64.rpm) |
| `logerrors_19` | `2.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.8 KiB | [logerrors_19-2.1.6-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/logerrors_19-2.1.6-1PGDG.rhel10.3.x86_64.rpm) |
| `logerrors_19` | `2.1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [logerrors_19-2.1.6-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/logerrors_19-2.1.6-1PGDG.rhel10.3.aarch64.rpm) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logerrors_18` | `2.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.4 KiB | [logerrors_18-2.1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/logerrors_18-2.1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `logerrors_18` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [logerrors_18-2.1.5-7PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/logerrors_18-2.1.5-7PGDG.rhel8.10.x86_64.rpm) |
| `logerrors_18` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.3 KiB | [logerrors_18-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/logerrors_18-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_18` | `2.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.6 KiB | [logerrors_18-2.1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/logerrors_18-2.1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `logerrors_18` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.4 KiB | [logerrors_18-2.1.5-7PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/logerrors_18-2.1.5-7PGDG.rhel8.10.aarch64.rpm) |
| `logerrors_18` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.4 KiB | [logerrors_18-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/logerrors_18-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_18` | `2.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.6 KiB | [logerrors_18-2.1.6-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/logerrors_18-2.1.6-1PGDG.rhel9.9.x86_64.rpm) |
| `logerrors_18` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.5 KiB | [logerrors_18-2.1.5-7PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/logerrors_18-2.1.5-7PGDG.rhel9.9.x86_64.rpm) |
| `logerrors_18` | `2.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [logerrors_18-2.1.6-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/logerrors_18-2.1.6-1PGDG.rhel9.9.aarch64.rpm) |
| `logerrors_18` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.6 KiB | [logerrors_18-2.1.5-7PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/logerrors_18-2.1.5-7PGDG.rhel9.9.aarch64.rpm) |
| `logerrors_18` | `2.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.8 KiB | [logerrors_18-2.1.6-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/logerrors_18-2.1.6-1PGDG.rhel10.3.x86_64.rpm) |
| `logerrors_18` | `2.1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.7 KiB | [logerrors_18-2.1.5-7PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/logerrors_18-2.1.5-7PGDG.rhel10.3.x86_64.rpm) |
| `logerrors_18` | `2.1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.1 KiB | [logerrors_18-2.1.6-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/logerrors_18-2.1.6-1PGDG.rhel10.3.aarch64.rpm) |
| `logerrors_18` | `2.1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.2 KiB | [logerrors_18-2.1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/logerrors_18-2.1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `logerrors_18` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [logerrors_18-2.1.5-7PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/logerrors_18-2.1.5-7PGDG.rhel10.3.aarch64.rpm) |
| `logerrors_18` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [logerrors_18-2.1.5-7PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/logerrors_18-2.1.5-7PGDG.rhel10.2.aarch64.rpm) |
| `logerrors_18` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.5 KiB | [logerrors_18-2.1.5-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/logerrors_18-2.1.5-3PGDG.rhel10.2.aarch64.rpm) |
| `logerrors_18` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.4 KiB | [logerrors_18-2.1.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/logerrors_18-2.1.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-logerrors` | `2.1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.8 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.3 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.8 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.5 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 33.5 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 33.5 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 32.4 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 32.2 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 32.0 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-logerrors` | `2.1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 32.7 KiB | [postgresql-18-logerrors_2.1.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/logerrors/postgresql-18-logerrors_2.1.5-1PIGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logerrors_17` | `2.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.5 KiB | [logerrors_17-2.1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/logerrors_17-2.1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `logerrors_17` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [logerrors_17-2.1.5-7PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/logerrors_17-2.1.5-7PGDG.rhel8.10.x86_64.rpm) |
| `logerrors_17` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.3 KiB | [logerrors_17-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/logerrors_17-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_17` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.9 KiB | [logerrors_17-2.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/logerrors_17-2.1.3-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_17` | `2.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.7 KiB | [logerrors_17-2.1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/logerrors_17-2.1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `logerrors_17` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.5 KiB | [logerrors_17-2.1.5-7PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/logerrors_17-2.1.5-7PGDG.rhel8.10.aarch64.rpm) |
| `logerrors_17` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.5 KiB | [logerrors_17-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/logerrors_17-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_17` | `2.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.1 KiB | [logerrors_17-2.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/logerrors_17-2.1.3-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_17` | `2.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.7 KiB | [logerrors_17-2.1.6-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/logerrors_17-2.1.6-1PGDG.rhel9.9.x86_64.rpm) |
| `logerrors_17` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.5 KiB | [logerrors_17-2.1.5-7PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/logerrors_17-2.1.5-7PGDG.rhel9.9.x86_64.rpm) |
| `logerrors_17` | `2.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 24.0 KiB | [logerrors_17-2.1.6-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/logerrors_17-2.1.6-1PGDG.rhel9.9.aarch64.rpm) |
| `logerrors_17` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.8 KiB | [logerrors_17-2.1.5-7PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/logerrors_17-2.1.5-7PGDG.rhel9.9.aarch64.rpm) |
| `logerrors_17` | `2.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.9 KiB | [logerrors_17-2.1.6-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/logerrors_17-2.1.6-1PGDG.rhel10.3.x86_64.rpm) |
| `logerrors_17` | `2.1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.7 KiB | [logerrors_17-2.1.5-7PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/logerrors_17-2.1.5-7PGDG.rhel10.3.x86_64.rpm) |
| `logerrors_17` | `2.1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.2 KiB | [logerrors_17-2.1.6-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/logerrors_17-2.1.6-1PGDG.rhel10.3.aarch64.rpm) |
| `logerrors_17` | `2.1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.2 KiB | [logerrors_17-2.1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/logerrors_17-2.1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `logerrors_17` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [logerrors_17-2.1.5-7PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/logerrors_17-2.1.5-7PGDG.rhel10.3.aarch64.rpm) |
| `logerrors_17` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [logerrors_17-2.1.5-7PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/logerrors_17-2.1.5-7PGDG.rhel10.2.aarch64.rpm) |
| `logerrors_17` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.4 KiB | [logerrors_17-2.1.5-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/logerrors_17-2.1.5-3PGDG.rhel10.2.aarch64.rpm) |
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
| `postgresql-17-logerrors` | `2.1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.9 KiB | [postgresql-17-logerrors_2.1.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/logerrors/postgresql-17-logerrors_2.1.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-logerrors` | `2.1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 32.7 KiB | [postgresql-17-logerrors_2.1.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/logerrors/postgresql-17-logerrors_2.1.5-1PIGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logerrors_16` | `2.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.5 KiB | [logerrors_16-2.1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/logerrors_16-2.1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `logerrors_16` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [logerrors_16-2.1.5-7PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/logerrors_16-2.1.5-7PGDG.rhel8.10.x86_64.rpm) |
| `logerrors_16` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.3 KiB | [logerrors_16-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/logerrors_16-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_16` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.9 KiB | [logerrors_16-2.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/logerrors_16-2.1.3-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_16` | `2.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.1 KiB | [logerrors_16-2.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/logerrors_16-2.1.2-1.rhel8.x86_64.rpm) |
| `logerrors_16` | `2.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.7 KiB | [logerrors_16-2.1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/logerrors_16-2.1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `logerrors_16` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.5 KiB | [logerrors_16-2.1.5-7PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/logerrors_16-2.1.5-7PGDG.rhel8.10.aarch64.rpm) |
| `logerrors_16` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.5 KiB | [logerrors_16-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/logerrors_16-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_16` | `2.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.1 KiB | [logerrors_16-2.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/logerrors_16-2.1.3-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_16` | `2.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.2 KiB | [logerrors_16-2.1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/logerrors_16-2.1.2-1.rhel8.aarch64.rpm) |
| `logerrors_16` | `2.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.7 KiB | [logerrors_16-2.1.6-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/logerrors_16-2.1.6-1PGDG.rhel9.9.x86_64.rpm) |
| `logerrors_16` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.5 KiB | [logerrors_16-2.1.5-7PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/logerrors_16-2.1.5-7PGDG.rhel9.9.x86_64.rpm) |
| `logerrors_16` | `2.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [logerrors_16-2.1.6-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/logerrors_16-2.1.6-1PGDG.rhel9.9.aarch64.rpm) |
| `logerrors_16` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.8 KiB | [logerrors_16-2.1.5-7PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/logerrors_16-2.1.5-7PGDG.rhel9.9.aarch64.rpm) |
| `logerrors_16` | `2.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.9 KiB | [logerrors_16-2.1.6-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/logerrors_16-2.1.6-1PGDG.rhel10.3.x86_64.rpm) |
| `logerrors_16` | `2.1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.7 KiB | [logerrors_16-2.1.5-7PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/logerrors_16-2.1.5-7PGDG.rhel10.3.x86_64.rpm) |
| `logerrors_16` | `2.1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.2 KiB | [logerrors_16-2.1.6-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/logerrors_16-2.1.6-1PGDG.rhel10.3.aarch64.rpm) |
| `logerrors_16` | `2.1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.3 KiB | [logerrors_16-2.1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/logerrors_16-2.1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `logerrors_16` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [logerrors_16-2.1.5-7PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/logerrors_16-2.1.5-7PGDG.rhel10.3.aarch64.rpm) |
| `logerrors_16` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.0 KiB | [logerrors_16-2.1.5-7PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/logerrors_16-2.1.5-7PGDG.rhel10.2.aarch64.rpm) |
| `logerrors_16` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.5 KiB | [logerrors_16-2.1.5-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/logerrors_16-2.1.5-3PGDG.rhel10.2.aarch64.rpm) |
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
| `postgresql-16-logerrors` | `2.1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.9 KiB | [postgresql-16-logerrors_2.1.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/logerrors/postgresql-16-logerrors_2.1.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-logerrors` | `2.1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 32.7 KiB | [postgresql-16-logerrors_2.1.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/logerrors/postgresql-16-logerrors_2.1.5-1PIGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `logerrors_15` | `2.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.5 KiB | [logerrors_15-2.1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `logerrors_15` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.5 KiB | [logerrors_15-2.1.5-7PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.1.5-7PGDG.rhel8.10.x86_64.rpm) |
| `logerrors_15` | `2.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.4 KiB | [logerrors_15-2.1.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.1.5-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 23.1 KiB | [logerrors_15-2.1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.1.3-1PGDG.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.3 KiB | [logerrors_15-2.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.1.2-1.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.1 KiB | [logerrors_15-2.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.1-2.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.7 KiB | [logerrors_15-2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.1-1.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.9 KiB | [logerrors_15-2.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.0-2.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [logerrors_15-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/logerrors_15-2.0-1.rhel8.x86_64.rpm) |
| `logerrors_15` | `2.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.6 KiB | [logerrors_15-2.1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `logerrors_15` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.6 KiB | [logerrors_15-2.1.5-7PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.1.5-7PGDG.rhel8.10.aarch64.rpm) |
| `logerrors_15` | `2.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.6 KiB | [logerrors_15-2.1.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.1.5-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.2 KiB | [logerrors_15-2.1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.1.3-1PGDG.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.4 KiB | [logerrors_15-2.1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.1.2-1.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [logerrors_15-2.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.1-2.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.8 KiB | [logerrors_15-2.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.1-1.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.9 KiB | [logerrors_15-2.0-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.0-2.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.7 KiB | [logerrors_15-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/logerrors_15-2.0-1.rhel8.aarch64.rpm) |
| `logerrors_15` | `2.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.7 KiB | [logerrors_15-2.1.6-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/logerrors_15-2.1.6-1PGDG.rhel9.9.x86_64.rpm) |
| `logerrors_15` | `2.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.7 KiB | [logerrors_15-2.1.5-7PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/logerrors_15-2.1.5-7PGDG.rhel9.9.x86_64.rpm) |
| `logerrors_15` | `2.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [logerrors_15-2.1.6-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/logerrors_15-2.1.6-1PGDG.rhel9.9.aarch64.rpm) |
| `logerrors_15` | `2.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.9 KiB | [logerrors_15-2.1.5-7PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/logerrors_15-2.1.5-7PGDG.rhel9.9.aarch64.rpm) |
| `logerrors_15` | `2.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.9 KiB | [logerrors_15-2.1.6-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/logerrors_15-2.1.6-1PGDG.rhel10.3.x86_64.rpm) |
| `logerrors_15` | `2.1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.9 KiB | [logerrors_15-2.1.5-7PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/logerrors_15-2.1.5-7PGDG.rhel10.3.x86_64.rpm) |
| `logerrors_15` | `2.1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.2 KiB | [logerrors_15-2.1.6-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/logerrors_15-2.1.6-1PGDG.rhel10.3.aarch64.rpm) |
| `logerrors_15` | `2.1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.2 KiB | [logerrors_15-2.1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/logerrors_15-2.1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `logerrors_15` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.2 KiB | [logerrors_15-2.1.5-7PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/logerrors_15-2.1.5-7PGDG.rhel10.3.aarch64.rpm) |
| `logerrors_15` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 24.2 KiB | [logerrors_15-2.1.5-7PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/logerrors_15-2.1.5-7PGDG.rhel10.2.aarch64.rpm) |
| `logerrors_15` | `2.1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.6 KiB | [logerrors_15-2.1.5-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/logerrors_15-2.1.5-3PGDG.rhel10.2.aarch64.rpm) |
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
| `postgresql-15-logerrors` | `2.1.5` | [u26.x86_64](/os/u26.x86_64) | pigsty | 32.1 KiB | [postgresql-15-logerrors_2.1.5-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/logerrors/postgresql-15-logerrors_2.1.5-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-logerrors` | `2.1.5` | [u26.aarch64](/os/u26.aarch64) | pigsty | 32.8 KiB | [postgresql-15-logerrors_2.1.5-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/l/logerrors/postgresql-15-logerrors_2.1.5-1PIGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/munakoiso/logerrors" title="Repository" icon="github" subtitle="github.com/munakoiso/logerrors" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="logerrors-2.1.6.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg logerrors;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install logerrors;		# install via package name, for the active PG version

pig install logerrors -v 19;   # install for PG 19
pig install logerrors -v 18;   # install for PG 18
pig install logerrors -v 17;   # install for PG 17
pig install logerrors -v 16;   # install for PG 16
pig install logerrors -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'logerrors';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION logerrors;
```

## Usage

Sources:

- [Official source](https://github.com/munakoiso/logerrors/blob/v2.1.6/README.md)
- [Official source](https://github.com/munakoiso/logerrors/blob/v2.1.6/logerrors.control)
- [Official source](https://github.com/munakoiso/logerrors/releases/tag/v2.1.6)

`logerrors` collects counts of WARNING, ERROR and FATAL messages through server hooks. Release 2.1.6 fixes PostgreSQL 19 ABI/build compatibility while retaining SQL extension version 2.1.

### Enable and inspect

Add the library to the existing preload list, restart PostgreSQL, then create the SQL extension:

```ini
shared_preload_libraries = 'logerrors'
```

```sql
CREATE EXTENSION logerrors;
SELECT * FROM pg_log_errors_stats();
SELECT * FROM pg_slow_log_stats();
SELECT pg_log_errors_reset();
```

`pg_log_errors_stats()` returns collection intervals, message types, counts, role/database names and SQLSTATE. `pg_slow_log_stats()` reports slow-log counts and reset time; `pg_log_errors_reset()` resets statistics.

### Configuration

`logerrors.interval` controls the collection interval in milliseconds: default 5000, maximum 60000. `logerrors.intervals_count` retains 120 intervals by default, at most 360. `logerrors.excluded_errcodes` accepts comma-separated error codes.

These are bounded in-memory statistics, not a durable audit archive. Restrict access to operational information according to database policy. Replacing the library requires a restart because it is preloaded; SQL extversion 2.1 alone does not identify the 2.1.6 runtime fix.
