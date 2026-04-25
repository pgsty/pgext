---
title: "credcheck"
linkTitle: "credcheck"
description: "credcheck - postgresql plain text credential checker"
weight: 7310
categories: ["SEC"]
width: full
---

[**credcheck**](https://github.com/MigOpsRepos/credcheck) : credcheck - postgresql plain text credential checker


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7310** | {{< badge content="credcheck" link="https://github.com/MigOpsRepos/credcheck" >}} | {{< ext "credcheck" >}} | `4.7` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "passwordcheck_cracklib" >}} {{< ext "login_hook" >}} {{< ext "passwordcheck" >}} {{< ext "pgaudit" >}} {{< ext "pg_auth_mon" >}} {{< ext "set_user" >}} {{< ext "auth_delay" >}} {{< ext "pg_permissions" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `credcheck` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.7` | {{< bg "18" "credcheck_18" "green" >}} {{< bg "17" "credcheck_17" "green" >}} {{< bg "16" "credcheck_16" "green" >}} {{< bg "15" "credcheck_15" "green" >}} {{< bg "14" "credcheck_14" "green" >}} | `credcheck_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.7` | {{< bg "18" "postgresql-18-credcheck" "green" >}} {{< bg "17" "postgresql-17-credcheck" "green" >}} {{< bg "16" "postgresql-16-credcheck" "green" >}} {{< bg "15" "postgresql-15-credcheck" "green" >}} {{< bg "14" "postgresql-14-credcheck" "green" >}} | `postgresql-$v-credcheck` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.7" "credcheck_18 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_16 : AVAIL 12" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_15 : AVAIL 17" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_14 : AVAIL 17" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.7" "credcheck_18 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_16 : AVAIL 12" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_15 : AVAIL 17" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_14 : AVAIL 17" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.7" "credcheck_18 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_16 : AVAIL 12" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_15 : AVAIL 17" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_14 : AVAIL 16" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.7" "credcheck_18 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_16 : AVAIL 12" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_15 : AVAIL 17" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_14 : AVAIL 17" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.7" "credcheck_18 : AVAIL 7" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_14 : AVAIL 7" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.7" "credcheck_18 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_15 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_14 : AVAIL 8" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 4.7" "postgresql-18-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-17-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-16-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-15-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-14-credcheck : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 4.7" "postgresql-18-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-17-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-16-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-15-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-14-credcheck : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 4.7" "postgresql-18-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-17-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-16-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-15-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-14-credcheck : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 4.7" "postgresql-18-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-17-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-16-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-15-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-14-credcheck : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 4.7" "postgresql-18-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-17-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-16-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-15-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-14-credcheck : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 4.7" "postgresql-18-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-17-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-16-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-15-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-14-credcheck : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 4.7" "postgresql-18-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-17-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-16-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-15-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-14-credcheck : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 4.7" "postgresql-18-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-17-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-16-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-15-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-14-credcheck : AVAIL 2" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 4.7" "postgresql-18-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-17-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-16-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-15-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-14-credcheck : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 4.7" "postgresql-18-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-17-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-16-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-15-credcheck : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.7" "postgresql-14-credcheck : AVAIL 2" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_18` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.3 KiB | [credcheck_18-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [credcheck_18-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.5 KiB | [credcheck_18-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.8 KiB | [credcheck_18-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.6 KiB | [credcheck_18-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [credcheck_18-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_18` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.4 KiB | [credcheck_18-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_18` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.6 KiB | [credcheck_18-3.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-3.0-2PGDG.rhel8.x86_64.rpm) |
| `credcheck_18` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.5 KiB | [credcheck_18-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.1 KiB | [credcheck_18-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.8 KiB | [credcheck_18-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.2 KiB | [credcheck_18-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.9 KiB | [credcheck_18-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.2 KiB | [credcheck_18-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_18` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.8 KiB | [credcheck_18-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_18` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.1 KiB | [credcheck_18-3.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-3.0-2PGDG.rhel8.aarch64.rpm) |
| `credcheck_18` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.3 KiB | [credcheck_18-4.7-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.7-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_18` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_18-4.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.6-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_18` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.8 KiB | [credcheck_18-4.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.5-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_18` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.3 KiB | [credcheck_18-4.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.4-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_18` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.0 KiB | [credcheck_18-4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_18` | `4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.6 KiB | [credcheck_18-4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_18` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.2 KiB | [credcheck_18-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_18` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.9 KiB | [credcheck_18-3.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-3.0-2PGDG.rhel9.x86_64.rpm) |
| `credcheck_18` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.7 KiB | [credcheck_18-4.7-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.7-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_18` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_18-4.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.6-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_18` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.5 KiB | [credcheck_18-4.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.5-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_18` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.6 KiB | [credcheck_18-4.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.4-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_18` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [credcheck_18-4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_18` | `4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.1 KiB | [credcheck_18-4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_18` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.7 KiB | [credcheck_18-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_18` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.6 KiB | [credcheck_18-3.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-3.0-2PGDG.rhel9.aarch64.rpm) |
| `credcheck_18` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [credcheck_18-4.7-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.7-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_18` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.1 KiB | [credcheck_18-4.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.5-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_18` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.6 KiB | [credcheck_18-4.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.4-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_18` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.4 KiB | [credcheck_18-4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_18` | `4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.3 KiB | [credcheck_18-4.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.2-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_18` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.7 KiB | [credcheck_18-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_18` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 36.3 KiB | [credcheck_18-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_18` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.1 KiB | [credcheck_18-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_18-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_18-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.3 KiB | [credcheck_18-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_18-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.9 KiB | [credcheck_18-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_18` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.5 KiB | [credcheck_18-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_18` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.3 KiB | [credcheck_18-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-credcheck` | `4.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 75.2 KiB | [postgresql-18-credcheck_4.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.7-1.pgdg12+1_amd64.deb) |
| `postgresql-18-credcheck` | `4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 74.4 KiB | [postgresql-18-credcheck_4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-18-credcheck` | `4.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 74.0 KiB | [postgresql-18-credcheck_4.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.7-1.pgdg12+1_arm64.deb) |
| `postgresql-18-credcheck` | `4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 73.1 KiB | [postgresql-18-credcheck_4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-18-credcheck` | `4.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 75.0 KiB | [postgresql-18-credcheck_4.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.7-1.pgdg13+1_amd64.deb) |
| `postgresql-18-credcheck` | `4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 74.1 KiB | [postgresql-18-credcheck_4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-18-credcheck` | `4.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.8 KiB | [postgresql-18-credcheck_4.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.7-1.pgdg13+1_arm64.deb) |
| `postgresql-18-credcheck` | `4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.0 KiB | [postgresql-18-credcheck_4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-18-credcheck` | `4.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 69.6 KiB | [postgresql-18-credcheck_4.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 68.5 KiB | [postgresql-18-credcheck_4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `4.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 68.1 KiB | [postgresql-18-credcheck_4.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-credcheck` | `4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 67.2 KiB | [postgresql-18-credcheck_4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-credcheck` | `4.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.8 KiB | [postgresql-18-credcheck_4.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.2 KiB | [postgresql-18-credcheck_4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `4.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 67.5 KiB | [postgresql-18-credcheck_4.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-credcheck` | `4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 66.9 KiB | [postgresql-18-credcheck_4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-credcheck` | `4.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 68.4 KiB | [postgresql-18-credcheck_4.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 67.6 KiB | [postgresql-18-credcheck_4.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `4.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.0 KiB | [postgresql-18-credcheck_4.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-credcheck` | `4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 66.2 KiB | [postgresql-18-credcheck_4.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_4.6-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_17` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.4 KiB | [credcheck_17-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [credcheck_17-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.5 KiB | [credcheck_17-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.9 KiB | [credcheck_17-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.6 KiB | [credcheck_17-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [credcheck_17-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.5 KiB | [credcheck_17-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.5 KiB | [credcheck_17-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | `2.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.1 KiB | [credcheck_17-2.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-2.8-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.6 KiB | [credcheck_17-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.2 KiB | [credcheck_17-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.8 KiB | [credcheck_17-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.2 KiB | [credcheck_17-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.0 KiB | [credcheck_17-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.3 KiB | [credcheck_17-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.9 KiB | [credcheck_17-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.0 KiB | [credcheck_17-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | `2.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.7 KiB | [credcheck_17-2.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-2.8-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.4 KiB | [credcheck_17-4.7-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.7-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_17` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_17-4.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.6-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_17` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_17-4.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.5-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_17` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.3 KiB | [credcheck_17-4.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.4-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_17` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.1 KiB | [credcheck_17-4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_17` | `4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.6 KiB | [credcheck_17-4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_17` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.2 KiB | [credcheck_17-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_17` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.9 KiB | [credcheck_17-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-3.0-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_17` | `2.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.6 KiB | [credcheck_17-2.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-2.8-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_17` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.8 KiB | [credcheck_17-4.7-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.7-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_17` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [credcheck_17-4.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.6-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_17` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_17-4.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.5-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_17` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.0 KiB | [credcheck_17-4.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.4-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_17` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.8 KiB | [credcheck_17-4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_17` | `4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.2 KiB | [credcheck_17-4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_17` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.8 KiB | [credcheck_17-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_17` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.7 KiB | [credcheck_17-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-3.0-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_17` | `2.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.4 KiB | [credcheck_17-2.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-2.8-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_17` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.7 KiB | [credcheck_17-4.7-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.7-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_17` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.1 KiB | [credcheck_17-4.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.5-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_17` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.6 KiB | [credcheck_17-4.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.4-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_17` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.5 KiB | [credcheck_17-4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_17` | `4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.3 KiB | [credcheck_17-4.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.2-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_17` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.8 KiB | [credcheck_17-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_17` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 36.5 KiB | [credcheck_17-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_17` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_17-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_17-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_17-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [credcheck_17-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.1 KiB | [credcheck_17-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_17-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_17` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.6 KiB | [credcheck_17-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_17` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.4 KiB | [credcheck_17-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-credcheck` | `4.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 75.2 KiB | [postgresql-17-credcheck_4.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.7-1.pgdg12+1_amd64.deb) |
| `postgresql-17-credcheck` | `4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 74.3 KiB | [postgresql-17-credcheck_4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-17-credcheck` | `4.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 74.0 KiB | [postgresql-17-credcheck_4.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.7-1.pgdg12+1_arm64.deb) |
| `postgresql-17-credcheck` | `4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 73.2 KiB | [postgresql-17-credcheck_4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-17-credcheck` | `4.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 74.9 KiB | [postgresql-17-credcheck_4.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.7-1.pgdg13+1_amd64.deb) |
| `postgresql-17-credcheck` | `4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 74.1 KiB | [postgresql-17-credcheck_4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-17-credcheck` | `4.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.9 KiB | [postgresql-17-credcheck_4.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.7-1.pgdg13+1_arm64.deb) |
| `postgresql-17-credcheck` | `4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.1 KiB | [postgresql-17-credcheck_4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-17-credcheck` | `4.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 76.5 KiB | [postgresql-17-credcheck_4.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 75.4 KiB | [postgresql-17-credcheck_4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `4.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.9 KiB | [postgresql-17-credcheck_4.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-credcheck` | `4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.0 KiB | [postgresql-17-credcheck_4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-credcheck` | `4.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.8 KiB | [postgresql-17-credcheck_4.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.3 KiB | [postgresql-17-credcheck_4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `4.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 67.5 KiB | [postgresql-17-credcheck_4.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-credcheck` | `4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 67.1 KiB | [postgresql-17-credcheck_4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-credcheck` | `4.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 68.3 KiB | [postgresql-17-credcheck_4.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 67.7 KiB | [postgresql-17-credcheck_4.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `4.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.0 KiB | [postgresql-17-credcheck_4.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-credcheck` | `4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 66.2 KiB | [postgresql-17-credcheck_4.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_4.6-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_16` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.4 KiB | [credcheck_16-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [credcheck_16-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.5 KiB | [credcheck_16-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.9 KiB | [credcheck_16-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.6 KiB | [credcheck_16-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [credcheck_16-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.5 KiB | [credcheck_16-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.5 KiB | [credcheck_16-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `2.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.7 KiB | [credcheck_16-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.7-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.3 KiB | [credcheck_16-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.6-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.8 KiB | [credcheck_16-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.8 KiB | [credcheck_16-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.6 KiB | [credcheck_16-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.2 KiB | [credcheck_16-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.8 KiB | [credcheck_16-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.2 KiB | [credcheck_16-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.0 KiB | [credcheck_16-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.3 KiB | [credcheck_16-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.9 KiB | [credcheck_16-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.1 KiB | [credcheck_16-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `2.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.2 KiB | [credcheck_16-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.7-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.9 KiB | [credcheck_16-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.6-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.5 KiB | [credcheck_16-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [credcheck_16-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.5 KiB | [credcheck_16-4.7-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.7-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_16` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_16-4.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.6-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_16` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_16-4.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.5-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_16` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.3 KiB | [credcheck_16-4.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.4-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_16` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.1 KiB | [credcheck_16-4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_16` | `4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.7 KiB | [credcheck_16-4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.2 KiB | [credcheck_16-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.2 KiB | [credcheck_16-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-3.0-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `2.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.1 KiB | [credcheck_16-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-2.7-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.7 KiB | [credcheck_16-2.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-2.6-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.5 KiB | [credcheck_16-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-2.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.3 KiB | [credcheck_16-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-2.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.8 KiB | [credcheck_16-4.7-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.7-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_16` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [credcheck_16-4.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.6-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_16` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_16-4.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.5-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_16` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.0 KiB | [credcheck_16-4.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.4-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_16` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.5 KiB | [credcheck_16-4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_16` | `4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.2 KiB | [credcheck_16-4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.8 KiB | [credcheck_16-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.7 KiB | [credcheck_16-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-3.0-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `2.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.8 KiB | [credcheck_16-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-2.7-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.5 KiB | [credcheck_16-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-2.6-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.9 KiB | [credcheck_16-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-2.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.8 KiB | [credcheck_16-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-2.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.7 KiB | [credcheck_16-4.7-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.7-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_16` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.1 KiB | [credcheck_16-4.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.5-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_16` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.6 KiB | [credcheck_16-4.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.4-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_16` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.4 KiB | [credcheck_16-4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_16` | `4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.3 KiB | [credcheck_16-4.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.2-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_16` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.8 KiB | [credcheck_16-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_16` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 36.4 KiB | [credcheck_16-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_16` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_16-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_16-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_16-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [credcheck_16-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.1 KiB | [credcheck_16-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_16-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_16` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.6 KiB | [credcheck_16-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_16` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.4 KiB | [credcheck_16-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-credcheck` | `4.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 75.1 KiB | [postgresql-16-credcheck_4.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.7-1.pgdg12+1_amd64.deb) |
| `postgresql-16-credcheck` | `4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 74.3 KiB | [postgresql-16-credcheck_4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-16-credcheck` | `4.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 74.0 KiB | [postgresql-16-credcheck_4.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.7-1.pgdg12+1_arm64.deb) |
| `postgresql-16-credcheck` | `4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 73.3 KiB | [postgresql-16-credcheck_4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-16-credcheck` | `4.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 74.9 KiB | [postgresql-16-credcheck_4.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.7-1.pgdg13+1_amd64.deb) |
| `postgresql-16-credcheck` | `4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 74.1 KiB | [postgresql-16-credcheck_4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-16-credcheck` | `4.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.9 KiB | [postgresql-16-credcheck_4.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.7-1.pgdg13+1_arm64.deb) |
| `postgresql-16-credcheck` | `4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.1 KiB | [postgresql-16-credcheck_4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-16-credcheck` | `4.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 76.1 KiB | [postgresql-16-credcheck_4.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 75.0 KiB | [postgresql-16-credcheck_4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `4.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.5 KiB | [postgresql-16-credcheck_4.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-credcheck` | `4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 73.5 KiB | [postgresql-16-credcheck_4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-credcheck` | `4.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.7 KiB | [postgresql-16-credcheck_4.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.2 KiB | [postgresql-16-credcheck_4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `4.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 67.5 KiB | [postgresql-16-credcheck_4.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-credcheck` | `4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 67.0 KiB | [postgresql-16-credcheck_4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-credcheck` | `4.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 68.3 KiB | [postgresql-16-credcheck_4.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 67.5 KiB | [postgresql-16-credcheck_4.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `4.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 67.0 KiB | [postgresql-16-credcheck_4.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-credcheck` | `4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 66.2 KiB | [postgresql-16-credcheck_4.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_4.6-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_15` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.5 KiB | [credcheck_15-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [credcheck_15-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.6 KiB | [credcheck_15-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.0 KiB | [credcheck_15-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.8 KiB | [credcheck_15-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.2 KiB | [credcheck_15-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.6 KiB | [credcheck_15-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.6 KiB | [credcheck_15-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.7 KiB | [credcheck_15-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.7-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.4 KiB | [credcheck_15-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.6-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.0 KiB | [credcheck_15-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.9 KiB | [credcheck_15-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.1 KiB | [credcheck_15-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.0-1.rhel8.x86_64.rpm) |
| `credcheck_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.7 KiB | [credcheck_15-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-1.2-1.rhel8.x86_64.rpm) |
| `credcheck_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.1 KiB | [credcheck_15-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-1.0-1.rhel8.x86_64.rpm) |
| `credcheck_15` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.6 KiB | [credcheck_15-0.2.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-0.2.0-3.rhel8.x86_64.rpm) |
| `credcheck_15` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.0 KiB | [credcheck_15-0.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-0.2.0-1.rhel8.x86_64.rpm) |
| `credcheck_15` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.5 KiB | [credcheck_15-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.1 KiB | [credcheck_15-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.8 KiB | [credcheck_15-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.1 KiB | [credcheck_15-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.9 KiB | [credcheck_15-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.2 KiB | [credcheck_15-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.8 KiB | [credcheck_15-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.0 KiB | [credcheck_15-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.2 KiB | [credcheck_15-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.7-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.9 KiB | [credcheck_15-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.6-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.5 KiB | [credcheck_15-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [credcheck_15-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.5 KiB | [credcheck_15-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.0-1.rhel8.aarch64.rpm) |
| `credcheck_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.2 KiB | [credcheck_15-1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-1.2-1.rhel8.aarch64.rpm) |
| `credcheck_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.6 KiB | [credcheck_15-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-1.0-1.rhel8.aarch64.rpm) |
| `credcheck_15` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.3 KiB | [credcheck_15-0.2.0-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-0.2.0-3.rhel8.aarch64.rpm) |
| `credcheck_15` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.9 KiB | [credcheck_15-0.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-0.2.0-1.rhel8.aarch64.rpm) |
| `credcheck_15` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.5 KiB | [credcheck_15-4.7-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.7-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_15` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_15-4.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.6-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_15` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_15-4.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.5-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_15` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.3 KiB | [credcheck_15-4.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.4-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_15` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.0 KiB | [credcheck_15-4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_15` | `4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.7 KiB | [credcheck_15-4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.3 KiB | [credcheck_15-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.1 KiB | [credcheck_15-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-3.0-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `2.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.2 KiB | [credcheck_15-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.7-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.9 KiB | [credcheck_15-2.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.6-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.4 KiB | [credcheck_15-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [credcheck_15-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.6 KiB | [credcheck_15-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.0-1.rhel9.x86_64.rpm) |
| `credcheck_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.1 KiB | [credcheck_15-1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-1.2-1.rhel9.x86_64.rpm) |
| `credcheck_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.5 KiB | [credcheck_15-1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-1.0-1.rhel9.x86_64.rpm) |
| `credcheck_15` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.8 KiB | [credcheck_15-0.2.0-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-0.2.0-3.rhel9.x86_64.rpm) |
| `credcheck_15` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.9 KiB | [credcheck_15-0.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-0.2.0-1.rhel9.x86_64.rpm) |
| `credcheck_15` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.7 KiB | [credcheck_15-4.7-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.7-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_15` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_15-4.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.6-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_15` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_15-4.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.5-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_15` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.6 KiB | [credcheck_15-4.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.4-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_15` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.6 KiB | [credcheck_15-4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_15` | `4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.9 KiB | [credcheck_15-4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.7 KiB | [credcheck_15-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.8 KiB | [credcheck_15-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-3.0-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `2.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.8 KiB | [credcheck_15-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.7-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.5 KiB | [credcheck_15-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.6-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.9 KiB | [credcheck_15-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.8 KiB | [credcheck_15-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.9 KiB | [credcheck_15-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.0-1.rhel9.aarch64.rpm) |
| `credcheck_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.5 KiB | [credcheck_15-1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-1.2-1.rhel9.aarch64.rpm) |
| `credcheck_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.9 KiB | [credcheck_15-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-1.0-1.rhel9.aarch64.rpm) |
| `credcheck_15` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.1 KiB | [credcheck_15-0.2.0-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-0.2.0-3.rhel9.aarch64.rpm) |
| `credcheck_15` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.5 KiB | [credcheck_15-0.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-0.2.0-1.rhel9.aarch64.rpm) |
| `credcheck_15` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [credcheck_15-4.7-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.7-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_15` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.1 KiB | [credcheck_15-4.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.5-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_15` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.7 KiB | [credcheck_15-4.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.4-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_15` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.4 KiB | [credcheck_15-4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_15` | `4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.3 KiB | [credcheck_15-4.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.2-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_15` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.8 KiB | [credcheck_15-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_15` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 36.5 KiB | [credcheck_15-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_15` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_15-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_15-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_15-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_15-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_15-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.9 KiB | [credcheck_15-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_15` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.5 KiB | [credcheck_15-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_15` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.5 KiB | [credcheck_15-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-credcheck` | `4.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 74.9 KiB | [postgresql-15-credcheck_4.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.7-1.pgdg12+1_amd64.deb) |
| `postgresql-15-credcheck` | `4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 74.1 KiB | [postgresql-15-credcheck_4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-15-credcheck` | `4.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 73.8 KiB | [postgresql-15-credcheck_4.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.7-1.pgdg12+1_arm64.deb) |
| `postgresql-15-credcheck` | `4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 72.9 KiB | [postgresql-15-credcheck_4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-15-credcheck` | `4.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 74.8 KiB | [postgresql-15-credcheck_4.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.7-1.pgdg13+1_amd64.deb) |
| `postgresql-15-credcheck` | `4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 73.8 KiB | [postgresql-15-credcheck_4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-15-credcheck` | `4.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.6 KiB | [postgresql-15-credcheck_4.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.7-1.pgdg13+1_arm64.deb) |
| `postgresql-15-credcheck` | `4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 72.7 KiB | [postgresql-15-credcheck_4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-15-credcheck` | `4.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 75.7 KiB | [postgresql-15-credcheck_4.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 74.7 KiB | [postgresql-15-credcheck_4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `4.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.2 KiB | [postgresql-15-credcheck_4.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-credcheck` | `4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 73.1 KiB | [postgresql-15-credcheck_4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-credcheck` | `4.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.5 KiB | [postgresql-15-credcheck_4.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 67.8 KiB | [postgresql-15-credcheck_4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `4.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 67.2 KiB | [postgresql-15-credcheck_4.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-credcheck` | `4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 66.6 KiB | [postgresql-15-credcheck_4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-credcheck` | `4.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 68.3 KiB | [postgresql-15-credcheck_4.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 67.3 KiB | [postgresql-15-credcheck_4.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `4.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 66.9 KiB | [postgresql-15-credcheck_4.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-credcheck` | `4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 65.8 KiB | [postgresql-15-credcheck_4.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_4.6-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_14` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.4 KiB | [credcheck_14-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_14` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [credcheck_14-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_14` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.6 KiB | [credcheck_14-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_14` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.0 KiB | [credcheck_14-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_14` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.7 KiB | [credcheck_14-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_14` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.1 KiB | [credcheck_14-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.6 KiB | [credcheck_14-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.6 KiB | [credcheck_14-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `2.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.6 KiB | [credcheck_14-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.7-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.3 KiB | [credcheck_14-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.6-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.9 KiB | [credcheck_14-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.8 KiB | [credcheck_14-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.0 KiB | [credcheck_14-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.0-1.rhel8.x86_64.rpm) |
| `credcheck_14` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.7 KiB | [credcheck_14-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-1.2-1.rhel8.x86_64.rpm) |
| `credcheck_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.1 KiB | [credcheck_14-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-1.0-1.rhel8.x86_64.rpm) |
| `credcheck_14` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.6 KiB | [credcheck_14-0.2.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-0.2.0-3.rhel8.x86_64.rpm) |
| `credcheck_14` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.3 KiB | [credcheck_14-0.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-0.2.0-1.rhel8.x86_64.rpm) |
| `credcheck_14` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.4 KiB | [credcheck_14-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_14` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.0 KiB | [credcheck_14-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_14` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [credcheck_14-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_14` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.0 KiB | [credcheck_14-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_14` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.8 KiB | [credcheck_14-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_14` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.1 KiB | [credcheck_14-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.7 KiB | [credcheck_14-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.0 KiB | [credcheck_14-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `2.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.1 KiB | [credcheck_14-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.7-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.8 KiB | [credcheck_14-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.6-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.4 KiB | [credcheck_14-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.2 KiB | [credcheck_14-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.4 KiB | [credcheck_14-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.0-1.rhel8.aarch64.rpm) |
| `credcheck_14` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.2 KiB | [credcheck_14-1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-1.2-1.rhel8.aarch64.rpm) |
| `credcheck_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.6 KiB | [credcheck_14-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-1.0-1.rhel8.aarch64.rpm) |
| `credcheck_14` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.3 KiB | [credcheck_14-0.2.0-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-0.2.0-3.rhel8.aarch64.rpm) |
| `credcheck_14` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.8 KiB | [credcheck_14-0.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-0.2.0-1.rhel8.aarch64.rpm) |
| `credcheck_14` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.4 KiB | [credcheck_14-4.7-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.7-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_14` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_14-4.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.6-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_14` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_14-4.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.5-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_14` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.2 KiB | [credcheck_14-4.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.4-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_14` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.0 KiB | [credcheck_14-4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_14` | `4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.7 KiB | [credcheck_14-4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.3 KiB | [credcheck_14-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.0 KiB | [credcheck_14-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-3.0-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `2.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.1 KiB | [credcheck_14-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.7-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.8 KiB | [credcheck_14-2.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.6-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.3 KiB | [credcheck_14-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.2 KiB | [credcheck_14-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.3 KiB | [credcheck_14-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.0-1.rhel9.x86_64.rpm) |
| `credcheck_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.0 KiB | [credcheck_14-1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-1.2-1.rhel9.x86_64.rpm) |
| `credcheck_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.4 KiB | [credcheck_14-1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-1.0-1.rhel9.x86_64.rpm) |
| `credcheck_14` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.8 KiB | [credcheck_14-0.2.0-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-0.2.0-3.rhel9.x86_64.rpm) |
| `credcheck_14` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.6 KiB | [credcheck_14-4.7-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.7-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_14` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_14-4.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.6-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_14` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_14-4.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.5-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_14` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.5 KiB | [credcheck_14-4.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.4-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_14` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.5 KiB | [credcheck_14-4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_14` | `4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.0 KiB | [credcheck_14-4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.6 KiB | [credcheck_14-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.6 KiB | [credcheck_14-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-3.0-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `2.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.8 KiB | [credcheck_14-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.7-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.4 KiB | [credcheck_14-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.6-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.8 KiB | [credcheck_14-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.7 KiB | [credcheck_14-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.8 KiB | [credcheck_14-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.0-1.rhel9.aarch64.rpm) |
| `credcheck_14` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.4 KiB | [credcheck_14-1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-1.2-1.rhel9.aarch64.rpm) |
| `credcheck_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.8 KiB | [credcheck_14-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-1.0-1.rhel9.aarch64.rpm) |
| `credcheck_14` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.0 KiB | [credcheck_14-0.2.0-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-0.2.0-3.rhel9.aarch64.rpm) |
| `credcheck_14` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.4 KiB | [credcheck_14-0.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-0.2.0-1.rhel9.aarch64.rpm) |
| `credcheck_14` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [credcheck_14-4.7-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.7-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_14` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.1 KiB | [credcheck_14-4.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.5-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_14` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.6 KiB | [credcheck_14-4.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.4-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_14` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.4 KiB | [credcheck_14-4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_14` | `4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.3 KiB | [credcheck_14-4.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.2-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_14` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.8 KiB | [credcheck_14-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_14` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 36.5 KiB | [credcheck_14-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_14` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.0 KiB | [credcheck_14-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_14` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.5 KiB | [credcheck_14-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_14` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.5 KiB | [credcheck_14-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_14` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.9 KiB | [credcheck_14-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_14` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.9 KiB | [credcheck_14-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_14` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.6 KiB | [credcheck_14-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_14` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.2 KiB | [credcheck_14-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_14` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.4 KiB | [credcheck_14-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-credcheck` | `4.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 74.7 KiB | [postgresql-14-credcheck_4.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.7-1.pgdg12+1_amd64.deb) |
| `postgresql-14-credcheck` | `4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 73.8 KiB | [postgresql-14-credcheck_4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-14-credcheck` | `4.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 73.5 KiB | [postgresql-14-credcheck_4.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.7-1.pgdg12+1_arm64.deb) |
| `postgresql-14-credcheck` | `4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 72.7 KiB | [postgresql-14-credcheck_4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-14-credcheck` | `4.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 74.5 KiB | [postgresql-14-credcheck_4.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.7-1.pgdg13+1_amd64.deb) |
| `postgresql-14-credcheck` | `4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 73.8 KiB | [postgresql-14-credcheck_4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-14-credcheck` | `4.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 73.4 KiB | [postgresql-14-credcheck_4.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.7-1.pgdg13+1_arm64.deb) |
| `postgresql-14-credcheck` | `4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 72.6 KiB | [postgresql-14-credcheck_4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-14-credcheck` | `4.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 75.6 KiB | [postgresql-14-credcheck_4.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-credcheck` | `4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 74.4 KiB | [postgresql-14-credcheck_4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-credcheck` | `4.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.0 KiB | [postgresql-14-credcheck_4.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-credcheck` | `4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 73.0 KiB | [postgresql-14-credcheck_4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-credcheck` | `4.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.3 KiB | [postgresql-14-credcheck_4.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-credcheck` | `4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 67.8 KiB | [postgresql-14-credcheck_4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-credcheck` | `4.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 67.1 KiB | [postgresql-14-credcheck_4.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-credcheck` | `4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 66.5 KiB | [postgresql-14-credcheck_4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-credcheck` | `4.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 68.0 KiB | [postgresql-14-credcheck_4.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-credcheck` | `4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 67.2 KiB | [postgresql-14-credcheck_4.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-credcheck` | `4.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 66.6 KiB | [postgresql-14-credcheck_4.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-credcheck` | `4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 65.7 KiB | [postgresql-14-credcheck_4.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_4.6-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/MigOpsRepos/credcheck" title="Repository" icon="github" subtitle="github.com/MigOpsRepos/credcheck" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install credcheck;		# install via package name, for the active PG version

pig install credcheck -v 18;   # install for PG 18
pig install credcheck -v 17;   # install for PG 17
pig install credcheck -v 16;   # install for PG 16
pig install credcheck -v 15;   # install for PG 15
pig install credcheck -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'credcheck';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION credcheck;
```




## Usage

> [credcheck: Credential checking for PostgreSQL usernames and passwords](https://github.com/MigOpsRepos/credcheck)

`credcheck` enforces configurable rules for username and password strength during `CREATE ROLE`, `ALTER ROLE`, and password changes. It also supports password reuse policies and authentication failure banning.

### Configuration Parameters

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'credcheck'
```

#### Username Checks

| Parameter | Description | Example |
|-----------|-------------|---------|
| `credcheck.username_min_length` | Minimum username length | `4` |
| `credcheck.username_min_special` | Minimum special characters | `1` |
| `credcheck.username_min_digit` | Minimum digit characters | `1` |
| `credcheck.username_min_upper` | Minimum uppercase characters | `2` |
| `credcheck.username_min_lower` | Minimum lowercase characters | `1` |
| `credcheck.username_min_repeat` | Max adjacent repeat characters | `2` |
| `credcheck.username_contain` | Must contain one of these chars | `a,b,c` |
| `credcheck.username_not_contain` | Must not contain these chars | `x,y,z` |
| `credcheck.username_contain_password` | Username must not contain password | `on` |

#### Password Checks

| Parameter | Description | Example |
|-----------|-------------|---------|
| `credcheck.password_min_length` | Minimum password length | `8` |
| `credcheck.password_min_special` | Minimum special characters | `1` |
| `credcheck.password_min_digit` | Minimum digit characters | `1` |
| `credcheck.password_min_upper` | Minimum uppercase characters | `1` |
| `credcheck.password_min_lower` | Minimum lowercase characters | `1` |
| `credcheck.password_min_repeat` | Max adjacent repeat characters | `3` |
| `credcheck.password_contain_username` | Password must not contain username | `on` |
| `credcheck.password_valid_until` | Minimum days for VALID UNTIL | `60` |
| `credcheck.password_valid_max` | Maximum days for VALID UNTIL | `365` |
| `credcheck.whitelist` | Usernames excluded from checks | `admin,super` |

### Examples

```sql
-- Rejected: username too short
CREATE USER abc WITH PASSWORD 'pass';
-- ERROR: username length should match the configured credcheck.username_min_length

-- Rejected: password contains username
CREATE USER abcd$ WITH PASSWORD 'abcd$xyz';
-- ERROR: password should not contain username
```

### Password Reuse Policy

```sql
SET credcheck.password_reuse_history = 2;
SET credcheck.password_reuse_interval = 365;  -- days
```

View password history:

```sql
SELECT rolename, password_hash FROM pg_password_history;
```

### Authentication Failure Ban

```sql
SET credcheck.max_auth_failure = 3;  -- ban after 3 failures
```

Reset banned users:

```sql
SELECT pg_banned_role_reset();              -- reset all
SELECT pg_banned_role_reset('username');     -- reset specific user
```
