---
title: "toastinfo"
linkTitle: "toastinfo"
description: "show details on toasted datums"
weight: 6530
categories: ["STAT"]
width: full
---

[**toastinfo**](https://github.com/df7cb/toastinfo) : show details on toasted datums


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6530** | {{< badge content="toastinfo" link="https://github.com/df7cb/toastinfo" >}} | {{< ext "toastinfo" >}} | `1.7` | {{< category "STAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pgstattuple" >}} {{< ext "amcheck" >}} {{< ext "pg_relusage" >}} {{< ext "pg_buffercache" >}} {{< ext "pg_freespacemap" >}} {{< ext "pg_repack" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `toastinfo` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7` | {{< bg "18" "toastinfo_18" "green" >}} {{< bg "17" "toastinfo_17" "green" >}} {{< bg "16" "toastinfo_16" "green" >}} {{< bg "15" "toastinfo_15" "green" >}} {{< bg "14" "toastinfo_14" "green" >}} | `toastinfo_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.7` | {{< bg "18" "postgresql-18-toastinfo" "green" >}} {{< bg "17" "postgresql-17-toastinfo" "green" >}} {{< bg "16" "postgresql-16-toastinfo" "green" >}} {{< bg "15" "postgresql-15-toastinfo" "green" >}} {{< bg "14" "postgresql-14-toastinfo" "green" >}} | `postgresql-$v-toastinfo` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.7" "toastinfo_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_14 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.7" "toastinfo_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_14 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.7" "toastinfo_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_14 : AVAIL 4" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.7" "toastinfo_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_14 : AVAIL 4" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.7" "toastinfo_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_14 : AVAIL 4" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.7" "toastinfo_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_15 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.7" "toastinfo_14 : AVAIL 4" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.7" "postgresql-18-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-toastinfo : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.7" "postgresql-18-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-toastinfo : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.7" "postgresql-18-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-toastinfo : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.7" "postgresql-18-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-toastinfo : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.7" "postgresql-18-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-toastinfo : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.7" "postgresql-18-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-toastinfo : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.7" "postgresql-18-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-toastinfo : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.7" "postgresql-18-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-toastinfo : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.7" "postgresql-18-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-toastinfo : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.7" "postgresql-18-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-toastinfo : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-toastinfo : AVAIL 3" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `toastinfo_18` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [toastinfo_18-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_18-1.7-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_18` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.3 KiB | [toastinfo_18-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/toastinfo_18-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `toastinfo_18` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [toastinfo_18-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_18-1.7-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_18` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.2 KiB | [toastinfo_18-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/toastinfo_18-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `toastinfo_18` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [toastinfo_18-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_18-1.7-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_18` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [toastinfo_18-1.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/toastinfo_18-1.6-1PGDG.rhel9.8.x86_64.rpm) |
| `toastinfo_18` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [toastinfo_18-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/toastinfo_18-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `toastinfo_18` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.1 KiB | [toastinfo_18-1.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/toastinfo_18-1.6-1PGDG.rhel9.6.x86_64.rpm) |
| `toastinfo_18` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [toastinfo_18-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_18-1.7-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_18` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [toastinfo_18-1.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/toastinfo_18-1.6-1PGDG.rhel9.8.aarch64.rpm) |
| `toastinfo_18` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [toastinfo_18-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/toastinfo_18-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `toastinfo_18` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [toastinfo_18-1.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/toastinfo_18-1.6-1PGDG.rhel9.6.aarch64.rpm) |
| `toastinfo_18` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [toastinfo_18-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/toastinfo_18-1.7-1PIGSTY.el10.x86_64.rpm) |
| `toastinfo_18` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [toastinfo_18-1.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/toastinfo_18-1.6-1PGDG.rhel10.2.x86_64.rpm) |
| `toastinfo_18` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [toastinfo_18-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/toastinfo_18-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `toastinfo_18` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.5 KiB | [toastinfo_18-1.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/toastinfo_18-1.6-1PGDG.rhel10.0.x86_64.rpm) |
| `toastinfo_18` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.7 KiB | [toastinfo_18-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/toastinfo_18-1.7-1PIGSTY.el10.aarch64.rpm) |
| `toastinfo_18` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [toastinfo_18-1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/toastinfo_18-1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `toastinfo_18` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [toastinfo_18-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/toastinfo_18-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `toastinfo_18` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [toastinfo_18-1.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/toastinfo_18-1.6-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-toastinfo` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-18-toastinfo_1.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.7-1.pgdg12+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-18-toastinfo_1.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.6-1.pgdg12+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.6 KiB | [postgresql-18-toastinfo_1.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg12+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.7 KiB | [postgresql-18-toastinfo_1.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.7-1.pgdg12+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-18-toastinfo_1.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.6-1.pgdg12+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-18-toastinfo_1.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg12+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-18-toastinfo_1.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.7-1.pgdg13+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-18-toastinfo_1.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.6-1.pgdg13+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.6 KiB | [postgresql-18-toastinfo_1.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg13+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-18-toastinfo_1.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.7-1.pgdg13+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-18-toastinfo_1.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.6-1.pgdg13+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-18-toastinfo_1.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg13+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.8 KiB | [postgresql-18-toastinfo_1.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.8 KiB | [postgresql-18-toastinfo_1.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.7 KiB | [postgresql-18-toastinfo_1.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.5 KiB | [postgresql-18-toastinfo_1.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.5 KiB | [postgresql-18-toastinfo_1.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.5 KiB | [postgresql-18-toastinfo_1.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-18-toastinfo_1.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-18-toastinfo_1.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-18-toastinfo_1.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-18-toastinfo_1.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-18-toastinfo_1.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-18-toastinfo_1.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg24.04+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.5 KiB | [postgresql-18-toastinfo_1.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.5 KiB | [postgresql-18-toastinfo_1.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.9 KiB | [postgresql-18-toastinfo_1.5-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg26.04+1_amd64.deb) |
| `postgresql-18-toastinfo` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-18-toastinfo_1.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-18-toastinfo_1.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-toastinfo` | `1.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 13.1 KiB | [postgresql-18-toastinfo_1.5-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-18-toastinfo_1.5-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `toastinfo_17` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [toastinfo_17-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_17-1.7-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_17` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.3 KiB | [toastinfo_17-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/toastinfo_17-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `toastinfo_17` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [toastinfo_17-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_17-1.7-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_17` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.2 KiB | [toastinfo_17-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/toastinfo_17-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `toastinfo_17` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [toastinfo_17-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_17-1.7-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_17` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [toastinfo_17-1.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/toastinfo_17-1.6-1PGDG.rhel9.8.x86_64.rpm) |
| `toastinfo_17` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [toastinfo_17-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/toastinfo_17-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `toastinfo_17` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.1 KiB | [toastinfo_17-1.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/toastinfo_17-1.6-1PGDG.rhel9.6.x86_64.rpm) |
| `toastinfo_17` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [toastinfo_17-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_17-1.7-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_17` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [toastinfo_17-1.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/toastinfo_17-1.6-1PGDG.rhel9.8.aarch64.rpm) |
| `toastinfo_17` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [toastinfo_17-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/toastinfo_17-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `toastinfo_17` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [toastinfo_17-1.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/toastinfo_17-1.6-1PGDG.rhel9.6.aarch64.rpm) |
| `toastinfo_17` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [toastinfo_17-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/toastinfo_17-1.7-1PIGSTY.el10.x86_64.rpm) |
| `toastinfo_17` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [toastinfo_17-1.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/toastinfo_17-1.6-1PGDG.rhel10.2.x86_64.rpm) |
| `toastinfo_17` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [toastinfo_17-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/toastinfo_17-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `toastinfo_17` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.5 KiB | [toastinfo_17-1.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/toastinfo_17-1.6-1PGDG.rhel10.0.x86_64.rpm) |
| `toastinfo_17` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.7 KiB | [toastinfo_17-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/toastinfo_17-1.7-1PIGSTY.el10.aarch64.rpm) |
| `toastinfo_17` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.1 KiB | [toastinfo_17-1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/toastinfo_17-1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `toastinfo_17` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [toastinfo_17-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/toastinfo_17-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `toastinfo_17` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [toastinfo_17-1.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/toastinfo_17-1.6-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-17-toastinfo` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-17-toastinfo_1.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.7-1.pgdg12+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-17-toastinfo_1.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.6-1.pgdg12+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.6 KiB | [postgresql-17-toastinfo_1.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg12+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-17-toastinfo_1.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.7-1.pgdg12+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-17-toastinfo_1.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.6-1.pgdg12+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-17-toastinfo_1.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg12+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-17-toastinfo_1.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.7-1.pgdg13+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-17-toastinfo_1.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.6-1.pgdg13+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.6 KiB | [postgresql-17-toastinfo_1.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg13+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-17-toastinfo_1.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.7-1.pgdg13+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-17-toastinfo_1.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.6-1.pgdg13+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-17-toastinfo_1.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg13+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.2 KiB | [postgresql-17-toastinfo_1.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.2 KiB | [postgresql-17-toastinfo_1.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.1 KiB | [postgresql-17-toastinfo_1.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.9 KiB | [postgresql-17-toastinfo_1.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.9 KiB | [postgresql-17-toastinfo_1.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.8 KiB | [postgresql-17-toastinfo_1.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-17-toastinfo_1.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-17-toastinfo_1.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-17-toastinfo_1.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-17-toastinfo_1.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-17-toastinfo_1.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-17-toastinfo_1.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.5 KiB | [postgresql-17-toastinfo_1.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.5 KiB | [postgresql-17-toastinfo_1.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.9 KiB | [postgresql-17-toastinfo_1.5-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg26.04+1_amd64.deb) |
| `postgresql-17-toastinfo` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-17-toastinfo_1.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-17-toastinfo_1.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-toastinfo` | `1.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 13.1 KiB | [postgresql-17-toastinfo_1.5-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-17-toastinfo_1.5-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `toastinfo_16` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [toastinfo_16-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_16-1.7-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_16` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.3 KiB | [toastinfo_16-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/toastinfo_16-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `toastinfo_16` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [toastinfo_16-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_16-1.7-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_16` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.2 KiB | [toastinfo_16-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/toastinfo_16-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `toastinfo_16` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [toastinfo_16-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_16-1.7-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_16` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [toastinfo_16-1.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/toastinfo_16-1.6-1PGDG.rhel9.8.x86_64.rpm) |
| `toastinfo_16` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.0 KiB | [toastinfo_16-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/toastinfo_16-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `toastinfo_16` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.1 KiB | [toastinfo_16-1.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/toastinfo_16-1.6-1PGDG.rhel9.6.x86_64.rpm) |
| `toastinfo_16` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [toastinfo_16-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_16-1.7-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_16` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [toastinfo_16-1.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/toastinfo_16-1.6-1PGDG.rhel9.8.aarch64.rpm) |
| `toastinfo_16` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [toastinfo_16-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/toastinfo_16-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `toastinfo_16` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [toastinfo_16-1.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/toastinfo_16-1.6-1PGDG.rhel9.6.aarch64.rpm) |
| `toastinfo_16` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [toastinfo_16-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/toastinfo_16-1.7-1PIGSTY.el10.x86_64.rpm) |
| `toastinfo_16` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [toastinfo_16-1.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/toastinfo_16-1.6-1PGDG.rhel10.2.x86_64.rpm) |
| `toastinfo_16` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.1 KiB | [toastinfo_16-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/toastinfo_16-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `toastinfo_16` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.5 KiB | [toastinfo_16-1.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/toastinfo_16-1.6-1PGDG.rhel10.0.x86_64.rpm) |
| `toastinfo_16` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.7 KiB | [toastinfo_16-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/toastinfo_16-1.7-1PIGSTY.el10.aarch64.rpm) |
| `toastinfo_16` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [toastinfo_16-1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/toastinfo_16-1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `toastinfo_16` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [toastinfo_16-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/toastinfo_16-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `toastinfo_16` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.0 KiB | [toastinfo_16-1.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/toastinfo_16-1.6-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-16-toastinfo` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-16-toastinfo_1.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.7-1.pgdg12+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-16-toastinfo_1.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.6-1.pgdg12+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.6 KiB | [postgresql-16-toastinfo_1.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg12+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-16-toastinfo_1.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.7-1.pgdg12+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-16-toastinfo_1.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.6-1.pgdg12+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-16-toastinfo_1.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg12+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-16-toastinfo_1.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.7-1.pgdg13+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-16-toastinfo_1.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.6-1.pgdg13+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.6 KiB | [postgresql-16-toastinfo_1.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg13+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-16-toastinfo_1.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.7-1.pgdg13+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-16-toastinfo_1.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.6-1.pgdg13+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-16-toastinfo_1.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg13+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.2 KiB | [postgresql-16-toastinfo_1.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.2 KiB | [postgresql-16-toastinfo_1.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.1 KiB | [postgresql-16-toastinfo_1.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.9 KiB | [postgresql-16-toastinfo_1.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.9 KiB | [postgresql-16-toastinfo_1.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.8 KiB | [postgresql-16-toastinfo_1.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-16-toastinfo_1.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-16-toastinfo_1.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-16-toastinfo_1.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-16-toastinfo_1.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-16-toastinfo_1.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-16-toastinfo_1.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg24.04+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.5 KiB | [postgresql-16-toastinfo_1.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.5 KiB | [postgresql-16-toastinfo_1.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.9 KiB | [postgresql-16-toastinfo_1.5-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg26.04+1_amd64.deb) |
| `postgresql-16-toastinfo` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-16-toastinfo_1.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-16-toastinfo_1.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-toastinfo` | `1.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 13.1 KiB | [postgresql-16-toastinfo_1.5-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-16-toastinfo_1.5-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `toastinfo_15` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [toastinfo_15-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_15-1.7-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_15` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.3 KiB | [toastinfo_15-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/toastinfo_15-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `toastinfo_15` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [toastinfo_15-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_15-1.7-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_15` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.2 KiB | [toastinfo_15-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/toastinfo_15-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `toastinfo_15` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [toastinfo_15-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_15-1.7-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_15` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.1 KiB | [toastinfo_15-1.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/toastinfo_15-1.6-1PGDG.rhel9.8.x86_64.rpm) |
| `toastinfo_15` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.1 KiB | [toastinfo_15-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/toastinfo_15-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `toastinfo_15` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.2 KiB | [toastinfo_15-1.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/toastinfo_15-1.6-1PGDG.rhel9.6.x86_64.rpm) |
| `toastinfo_15` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.6 KiB | [toastinfo_15-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_15-1.7-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_15` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [toastinfo_15-1.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/toastinfo_15-1.6-1PGDG.rhel9.8.aarch64.rpm) |
| `toastinfo_15` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [toastinfo_15-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/toastinfo_15-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `toastinfo_15` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.0 KiB | [toastinfo_15-1.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/toastinfo_15-1.6-1PGDG.rhel9.6.aarch64.rpm) |
| `toastinfo_15` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.7 KiB | [toastinfo_15-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/toastinfo_15-1.7-1PIGSTY.el10.x86_64.rpm) |
| `toastinfo_15` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.2 KiB | [toastinfo_15-1.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/toastinfo_15-1.6-1PGDG.rhel10.2.x86_64.rpm) |
| `toastinfo_15` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.2 KiB | [toastinfo_15-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/toastinfo_15-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `toastinfo_15` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.6 KiB | [toastinfo_15-1.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/toastinfo_15-1.6-1PGDG.rhel10.0.x86_64.rpm) |
| `toastinfo_15` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.8 KiB | [toastinfo_15-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/toastinfo_15-1.7-1PIGSTY.el10.aarch64.rpm) |
| `toastinfo_15` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.1 KiB | [toastinfo_15-1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/toastinfo_15-1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `toastinfo_15` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.1 KiB | [toastinfo_15-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/toastinfo_15-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `toastinfo_15` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.1 KiB | [toastinfo_15-1.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/toastinfo_15-1.6-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-15-toastinfo` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.7-1.pgdg12+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.6-1.pgdg12+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg12+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.7-1.pgdg12+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-15-toastinfo_1.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.6-1.pgdg12+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-15-toastinfo_1.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg12+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.7-1.pgdg13+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.6-1.pgdg13+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg13+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.7-1.pgdg13+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.6-1.pgdg13+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-15-toastinfo_1.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg13+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.3 KiB | [postgresql-15-toastinfo_1.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.3 KiB | [postgresql-15-toastinfo_1.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.2 KiB | [postgresql-15-toastinfo_1.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.0 KiB | [postgresql-15-toastinfo_1.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.0 KiB | [postgresql-15-toastinfo_1.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.0 KiB | [postgresql-15-toastinfo_1.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-15-toastinfo_1.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-15-toastinfo_1.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.8 KiB | [postgresql-15-toastinfo_1.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-15-toastinfo_1.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg24.04+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.5 KiB | [postgresql-15-toastinfo_1.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.5 KiB | [postgresql-15-toastinfo_1.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.9 KiB | [postgresql-15-toastinfo_1.5-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg26.04+1_amd64.deb) |
| `postgresql-15-toastinfo` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-15-toastinfo_1.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-15-toastinfo_1.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-toastinfo` | `1.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 13.1 KiB | [postgresql-15-toastinfo_1.5-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-15-toastinfo_1.5-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `toastinfo_14` | `1.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.8 KiB | [toastinfo_14-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/toastinfo_14-1.7-1PIGSTY.el8.x86_64.rpm) |
| `toastinfo_14` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 13.3 KiB | [toastinfo_14-1.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/toastinfo_14-1.6-1PGDG.rhel8.10.x86_64.rpm) |
| `toastinfo_14` | `1.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.9 KiB | [toastinfo_14-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/toastinfo_14-1.7-1PIGSTY.el8.aarch64.rpm) |
| `toastinfo_14` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.2 KiB | [toastinfo_14-1.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/toastinfo_14-1.6-1PGDG.rhel8.10.aarch64.rpm) |
| `toastinfo_14` | `1.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.6 KiB | [toastinfo_14-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/toastinfo_14-1.7-1PIGSTY.el9.x86_64.rpm) |
| `toastinfo_14` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.1 KiB | [toastinfo_14-1.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/toastinfo_14-1.6-1PGDG.rhel9.8.x86_64.rpm) |
| `toastinfo_14` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.1 KiB | [toastinfo_14-1.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/toastinfo_14-1.6-1PGDG.rhel9.7.x86_64.rpm) |
| `toastinfo_14` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.2 KiB | [toastinfo_14-1.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/toastinfo_14-1.6-1PGDG.rhel9.6.x86_64.rpm) |
| `toastinfo_14` | `1.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [toastinfo_14-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/toastinfo_14-1.7-1PIGSTY.el9.aarch64.rpm) |
| `toastinfo_14` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [toastinfo_14-1.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/toastinfo_14-1.6-1PGDG.rhel9.8.aarch64.rpm) |
| `toastinfo_14` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.8 KiB | [toastinfo_14-1.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/toastinfo_14-1.6-1PGDG.rhel9.7.aarch64.rpm) |
| `toastinfo_14` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.9 KiB | [toastinfo_14-1.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/toastinfo_14-1.6-1PGDG.rhel9.6.aarch64.rpm) |
| `toastinfo_14` | `1.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.6 KiB | [toastinfo_14-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/toastinfo_14-1.7-1PIGSTY.el10.x86_64.rpm) |
| `toastinfo_14` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.2 KiB | [toastinfo_14-1.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/toastinfo_14-1.6-1PGDG.rhel10.2.x86_64.rpm) |
| `toastinfo_14` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.2 KiB | [toastinfo_14-1.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/toastinfo_14-1.6-1PGDG.rhel10.1.x86_64.rpm) |
| `toastinfo_14` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 13.5 KiB | [toastinfo_14-1.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/toastinfo_14-1.6-1PGDG.rhel10.0.x86_64.rpm) |
| `toastinfo_14` | `1.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.7 KiB | [toastinfo_14-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/toastinfo_14-1.7-1PIGSTY.el10.aarch64.rpm) |
| `toastinfo_14` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.1 KiB | [toastinfo_14-1.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/toastinfo_14-1.6-1PGDG.rhel10.2.aarch64.rpm) |
| `toastinfo_14` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.1 KiB | [toastinfo_14-1.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/toastinfo_14-1.6-1PGDG.rhel10.1.aarch64.rpm) |
| `toastinfo_14` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 13.1 KiB | [toastinfo_14-1.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/toastinfo_14-1.6-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-14-toastinfo` | `1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-14-toastinfo_1.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.7-1.pgdg12+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-14-toastinfo_1.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.6-1.pgdg12+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.6 KiB | [postgresql-14-toastinfo_1.5-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg12+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-14-toastinfo_1.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.7-1.pgdg12+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-14-toastinfo_1.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.6-1.pgdg12+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.6 KiB | [postgresql-14-toastinfo_1.5-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg12+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-14-toastinfo_1.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.7-1.pgdg13+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-14-toastinfo_1.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.6-1.pgdg13+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.6 KiB | [postgresql-14-toastinfo_1.5-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg13+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-14-toastinfo_1.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.7-1.pgdg13+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.7 KiB | [postgresql-14-toastinfo_1.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.6-1.pgdg13+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.6 KiB | [postgresql-14-toastinfo_1.5-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg13+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.3 KiB | [postgresql-14-toastinfo_1.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.3 KiB | [postgresql-14-toastinfo_1.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.2 KiB | [postgresql-14-toastinfo_1.5-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.0 KiB | [postgresql-14-toastinfo_1.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.0 KiB | [postgresql-14-toastinfo_1.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.9 KiB | [postgresql-14-toastinfo_1.5-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-14-toastinfo_1.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-14-toastinfo_1.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.7 KiB | [postgresql-14-toastinfo_1.5-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-14-toastinfo_1.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.7 KiB | [postgresql-14-toastinfo_1.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.6 KiB | [postgresql-14-toastinfo_1.5-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg24.04+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.5 KiB | [postgresql-14-toastinfo_1.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.5 KiB | [postgresql-14-toastinfo_1.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.8 KiB | [postgresql-14-toastinfo_1.5-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg26.04+1_amd64.deb) |
| `postgresql-14-toastinfo` | `1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-14-toastinfo_1.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.8 KiB | [postgresql-14-toastinfo_1.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-toastinfo` | `1.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 13.1 KiB | [postgresql-14-toastinfo_1.5-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/toastinfo/postgresql-14-toastinfo_1.5-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/df7cb/toastinfo" title="Repository" icon="github" subtitle="github.com/df7cb/toastinfo" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="toastinfo-1.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg toastinfo;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install toastinfo;		# install via package name, for the active PG version

pig install toastinfo -v 18;   # install for PG 18
pig install toastinfo -v 17;   # install for PG 17
pig install toastinfo -v 16;   # install for PG 16
pig install toastinfo -v 15;   # install for PG 15
pig install toastinfo -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION toastinfo;
```




## Usage

Sources: [upstream README](https://github.com/credativ/toastinfo), [upstream tags](https://github.com/credativ/toastinfo/tags), [PGDG package metadata via local `db/extension.csv`](../db/extension.csv).

`toastinfo` exposes how PostgreSQL stores variable-length (`varlena`) values, including inline, compressed, and out-of-line TOAST storage forms.

### Functions

`pg_toastinfo(anyelement)` describes the storage form of a datum:

```sql
CREATE EXTENSION toastinfo;

SELECT a, length(b), pg_column_size(b), pg_toastinfo(b), pg_toastpointer(b)
FROM t;
```

Possible storage forms include:

- `null`, for NULL values.
- `ordinary`, for non-varlena data types.
- `short inline varlena`, up to 126 bytes with a 1-byte header.
- `long inline varlena, uncompressed`, up to 1 GiB with a 4-byte header.
- `long inline varlena, compressed (pglz|lz4)`.
- `toasted varlena, uncompressed`.
- `toasted varlena, compressed (pglz|lz4)`.

Compressed varlenas show the compression method on PostgreSQL 14+.

`pg_toastpointer(anyelement)` returns the TOAST table `chunk_id` OID for out-of-line toasted values, or NULL for non-toasted input:

```sql
SELECT pg_toastpointer(large_column)
FROM my_table;
```

### Storage Example

```sql
CREATE TABLE t (a text, b text);

ALTER TABLE t ALTER COLUMN b SET STORAGE EXTERNAL;
INSERT INTO t VALUES ('external-10000', repeat('x', 10000));

ALTER TABLE t ALTER COLUMN b SET STORAGE EXTENDED;
INSERT INTO t VALUES ('extended-1000000', repeat('x', 1000000));

ALTER TABLE t ALTER COLUMN b SET COMPRESSION lz4;
INSERT INTO t VALUES ('extended-lz4', repeat('x', 1000000));

SELECT a, pg_column_size(b), pg_toastinfo(b), pg_toastpointer(b)
FROM t;
```

### Caveats

- Pigsty metadata carries `toastinfo` 1.6 for PostgreSQL 14-18, with Pigsty RPMs and PGDG DEBs.
- The upstream GitHub README documents the same SQL surface, but the public GitHub tags visible during review stop at `v1.5`; no upstream 1.6 changelog was found in that repository.
- `pg_toastpointer` is meaningful only for out-of-line toasted data; ordinary, inline, or NULL values return NULL.
