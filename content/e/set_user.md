---
title: "set_user"
linkTitle: "set_user"
description: "similar to SET ROLE but with added logging"
weight: 7370
categories: ["SEC"]
width: full
---

[**set_user**](https://github.com/pgaudit/set_user) : similar to SET ROLE but with added logging


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7370** | {{< badge content="set_user" link="https://github.com/pgaudit/set_user" >}} | {{< ext "set_user" >}} | `4.2.0` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_readonly" >}} {{< ext "pg_permissions" >}} {{< ext "pgaudit" >}} {{< ext "login_hook" >}} {{< ext "pgauditlogtofile" >}} {{< ext "pg_auth_mon" >}} {{< ext "credcheck" >}} {{< ext "pgextwlist" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `set_user` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.0` | {{< bg "18" "set_user_18*" "green" >}} {{< bg "17" "set_user_17*" "green" >}} {{< bg "16" "set_user_16*" "green" >}} {{< bg "15" "set_user_15*" "green" >}} {{< bg "14" "set_user_14*" "green" >}} {{< bg "13" "set_user_13*" "green" >}} | `set_user_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.0` | {{< bg "18" "postgresql-18-set-user" "green" >}} {{< bg "17" "postgresql-17-set-user" "green" >}} {{< bg "16" "postgresql-16-set-user" "green" >}} {{< bg "15" "postgresql-15-set-user" "green" >}} {{< bg "14" "postgresql-14-set-user" "green" >}} {{< bg "13" "postgresql-13-set-user" "green" >}} | `postgresql-$v-set-user` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.2.0" "set_user_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_13 : AVAIL 5" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.2.0" "set_user_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_13 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.2.0" "set_user_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.2.0" "set_user_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_13 : AVAIL 4" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.2.0" "set_user_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.2.0" "set_user_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.1.0" "set_user_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 4.2.0" "postgresql-18-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-17-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-16-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-15-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-14-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-13-set-user : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 4.2.0" "postgresql-18-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-17-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-16-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-15-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-14-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-13-set-user : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 4.2.0" "postgresql-18-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-17-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-16-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-15-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-14-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-13-set-user : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 4.2.0" "postgresql-18-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-17-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-16-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-15-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-14-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-13-set-user : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 4.2.0" "postgresql-18-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-17-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-16-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-15-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-14-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-13-set-user : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 4.2.0" "postgresql-18-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-17-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-16-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-15-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-14-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-13-set-user : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 4.2.0" "postgresql-18-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-17-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-16-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-15-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-14-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-13-set-user : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 4.2.0" "postgresql-18-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-17-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-16-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-15-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-14-set-user : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.0" "postgresql-13-set-user : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `set_user_18` | `4.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.8 KiB | [set_user_18-4.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/set_user_18-4.2.0-1PGDG.rhel8.x86_64.rpm) |
| `set_user_18` | `4.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.5 KiB | [set_user_18-4.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/set_user_18-4.2.0-1PGDG.rhel8.aarch64.rpm) |
| `set_user_18` | `4.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.5 KiB | [set_user_18-4.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/set_user_18-4.2.0-1PGDG.rhel9.x86_64.rpm) |
| `set_user_18` | `4.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.8 KiB | [set_user_18-4.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/set_user_18-4.2.0-1PGDG.rhel9.aarch64.rpm) |
| `set_user_18` | `4.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 27.0 KiB | [set_user_18-4.2.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/set_user_18-4.2.0-1PGDG.rhel10.x86_64.rpm) |
| `set_user_18` | `4.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [set_user_18-4.2.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10.0-aarch64/set_user_18-4.2.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-set-user` | `4.2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 35.1 KiB | [postgresql-18-set-user_4.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-18-set-user_4.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-set-user` | `4.2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 34.7 KiB | [postgresql-18-set-user_4.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-18-set-user_4.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-set-user` | `4.2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 35.1 KiB | [postgresql-18-set-user_4.2.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-18-set-user_4.2.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-set-user` | `4.2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.8 KiB | [postgresql-18-set-user_4.2.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-18-set-user_4.2.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-set-user` | `4.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 35.1 KiB | [postgresql-18-set-user_4.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-18-set-user_4.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-set-user` | `4.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 34.7 KiB | [postgresql-18-set-user_4.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-18-set-user_4.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-set-user` | `4.2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.6 KiB | [postgresql-18-set-user_4.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-18-set-user_4.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-set-user` | `4.2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.1 KiB | [postgresql-18-set-user_4.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-18-set-user_4.2.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `set_user_17` | `4.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.4 KiB | [set_user_17-4.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/set_user_17-4.1.0-1PGDG.rhel8.x86_64.rpm) |
| `set_user_17` | `4.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.1 KiB | [set_user_17-4.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/set_user_17-4.1.0-1PGDG.rhel8.aarch64.rpm) |
| `set_user_17` | `4.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.5 KiB | [set_user_17-4.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/set_user_17-4.1.0-1PGDG.rhel9.x86_64.rpm) |
| `set_user_17` | `4.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.8 KiB | [set_user_17-4.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/set_user_17-4.1.0-1PGDG.rhel9.aarch64.rpm) |
| `set_user_17` | `4.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.8 KiB | [set_user_17-4.1.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/set_user_17-4.1.0-2PGDG.rhel10.x86_64.rpm) |
| `set_user_17` | `4.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.4 KiB | [set_user_17-4.1.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10.0-aarch64/set_user_17-4.1.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-set-user` | `4.2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 35.0 KiB | [postgresql-17-set-user_4.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-17-set-user_4.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-set-user` | `4.2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 34.6 KiB | [postgresql-17-set-user_4.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-17-set-user_4.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-set-user` | `4.2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 35.0 KiB | [postgresql-17-set-user_4.2.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-17-set-user_4.2.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-set-user` | `4.2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.6 KiB | [postgresql-17-set-user_4.2.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-17-set-user_4.2.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-set-user` | `4.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 39.0 KiB | [postgresql-17-set-user_4.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-17-set-user_4.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-set-user` | `4.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.6 KiB | [postgresql-17-set-user_4.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-17-set-user_4.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-set-user` | `4.2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.6 KiB | [postgresql-17-set-user_4.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-17-set-user_4.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-set-user` | `4.2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.1 KiB | [postgresql-17-set-user_4.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-17-set-user_4.2.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `set_user_16` | `4.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.4 KiB | [set_user_16-4.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/set_user_16-4.1.0-1PGDG.rhel8.x86_64.rpm) |
| `set_user_16` | `4.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.2 KiB | [set_user_16-4.0.1-2.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/set_user_16-4.0.1-2.rhel8.1.x86_64.rpm) |
| `set_user_16` | `4.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.2 KiB | [set_user_16-4.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/set_user_16-4.1.0-1PGDG.rhel8.aarch64.rpm) |
| `set_user_16` | `4.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.9 KiB | [set_user_16-4.0.1-2.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/set_user_16-4.0.1-2.rhel8.1.aarch64.rpm) |
| `set_user_16` | `4.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.5 KiB | [set_user_16-4.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/set_user_16-4.1.0-1PGDG.rhel9.x86_64.rpm) |
| `set_user_16` | `4.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.3 KiB | [set_user_16-4.0.1-2.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/set_user_16-4.0.1-2.rhel9.1.x86_64.rpm) |
| `set_user_16` | `4.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.9 KiB | [set_user_16-4.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/set_user_16-4.1.0-1PGDG.rhel9.aarch64.rpm) |
| `set_user_16` | `4.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.6 KiB | [set_user_16-4.0.1-2.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/set_user_16-4.0.1-2.rhel9.1.aarch64.rpm) |
| `set_user_16` | `4.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.8 KiB | [set_user_16-4.1.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/set_user_16-4.1.0-2PGDG.rhel10.x86_64.rpm) |
| `set_user_16` | `4.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.4 KiB | [set_user_16-4.1.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10.0-aarch64/set_user_16-4.1.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-set-user` | `4.2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 35.0 KiB | [postgresql-16-set-user_4.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-16-set-user_4.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-set-user` | `4.2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 34.6 KiB | [postgresql-16-set-user_4.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-16-set-user_4.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-set-user` | `4.2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 35.0 KiB | [postgresql-16-set-user_4.2.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-16-set-user_4.2.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-set-user` | `4.2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.6 KiB | [postgresql-16-set-user_4.2.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-16-set-user_4.2.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-set-user` | `4.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.5 KiB | [postgresql-16-set-user_4.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-16-set-user_4.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-set-user` | `4.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 38.1 KiB | [postgresql-16-set-user_4.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-16-set-user_4.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-set-user` | `4.2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.5 KiB | [postgresql-16-set-user_4.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-16-set-user_4.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-set-user` | `4.2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.0 KiB | [postgresql-16-set-user_4.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-16-set-user_4.2.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `set_user_15` | `4.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [set_user_15-4.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/set_user_15-4.1.0-1PGDG.rhel8.x86_64.rpm) |
| `set_user_15` | `4.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.0 KiB | [set_user_15-4.0.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/set_user_15-4.0.1-2.rhel8.x86_64.rpm) |
| `set_user_15` | `4.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 25.5 KiB | [set_user_15-4.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/set_user_15-4.0.0-1.rhel8.x86_64.rpm) |
| `set_user_15` | `4.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.1 KiB | [set_user_15-4.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/set_user_15-4.1.0-1PGDG.rhel8.aarch64.rpm) |
| `set_user_15` | `4.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.8 KiB | [set_user_15-4.0.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/set_user_15-4.0.1-2.rhel8.aarch64.rpm) |
| `set_user_15` | `4.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.2 KiB | [set_user_15-4.0.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/set_user_15-4.0.0-1.rhel8.aarch64.rpm) |
| `set_user_15` | `4.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.4 KiB | [set_user_15-4.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/set_user_15-4.1.0-1PGDG.rhel9.x86_64.rpm) |
| `set_user_15` | `4.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.1 KiB | [set_user_15-4.0.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/set_user_15-4.0.1-2.rhel9.x86_64.rpm) |
| `set_user_15` | `4.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.8 KiB | [set_user_15-4.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/set_user_15-4.0.0-1.rhel9.x86_64.rpm) |
| `set_user_15` | `4.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.8 KiB | [set_user_15-4.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/set_user_15-4.1.0-1PGDG.rhel9.aarch64.rpm) |
| `set_user_15` | `4.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.4 KiB | [set_user_15-4.0.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/set_user_15-4.0.1-2.rhel9.aarch64.rpm) |
| `set_user_15` | `4.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.1 KiB | [set_user_15-4.0.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/set_user_15-4.0.0-1.rhel9.aarch64.rpm) |
| `set_user_15` | `4.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.7 KiB | [set_user_15-4.1.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/set_user_15-4.1.0-2PGDG.rhel10.x86_64.rpm) |
| `set_user_15` | `4.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.3 KiB | [set_user_15-4.1.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10.0-aarch64/set_user_15-4.1.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-set-user` | `4.2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.6 KiB | [postgresql-15-set-user_4.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-15-set-user_4.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-set-user` | `4.2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 34.2 KiB | [postgresql-15-set-user_4.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-15-set-user_4.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-set-user` | `4.2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 34.7 KiB | [postgresql-15-set-user_4.2.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-15-set-user_4.2.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-set-user` | `4.2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.3 KiB | [postgresql-15-set-user_4.2.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-15-set-user_4.2.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-set-user` | `4.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.2 KiB | [postgresql-15-set-user_4.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-15-set-user_4.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-set-user` | `4.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 37.8 KiB | [postgresql-15-set-user_4.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-15-set-user_4.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-set-user` | `4.2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.2 KiB | [postgresql-15-set-user_4.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-15-set-user_4.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-set-user` | `4.2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 33.7 KiB | [postgresql-15-set-user_4.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-15-set-user_4.2.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `set_user_14` | `4.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [set_user_14-4.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/set_user_14-4.1.0-1PGDG.rhel8.x86_64.rpm) |
| `set_user_14` | `4.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.0 KiB | [set_user_14-4.0.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/set_user_14-4.0.1-2.rhel8.x86_64.rpm) |
| `set_user_14` | `4.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 25.5 KiB | [set_user_14-4.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/set_user_14-4.0.0-1.rhel8.x86_64.rpm) |
| `set_user_14` | `3.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 25.3 KiB | [set_user_14-3.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/set_user_14-3.0.0-1.rhel8.x86_64.rpm) |
| `set_user_14` | `4.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.1 KiB | [set_user_14-4.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/set_user_14-4.1.0-1PGDG.rhel8.aarch64.rpm) |
| `set_user_14` | `4.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.7 KiB | [set_user_14-4.0.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/set_user_14-4.0.1-2.rhel8.aarch64.rpm) |
| `set_user_14` | `4.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.2 KiB | [set_user_14-4.0.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/set_user_14-4.0.0-1.rhel8.aarch64.rpm) |
| `set_user_14` | `3.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.1 KiB | [set_user_14-3.0.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/set_user_14-3.0.0-1.rhel8.aarch64.rpm) |
| `set_user_14` | `4.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.4 KiB | [set_user_14-4.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/set_user_14-4.1.0-1PGDG.rhel9.x86_64.rpm) |
| `set_user_14` | `4.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.1 KiB | [set_user_14-4.0.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/set_user_14-4.0.1-2.rhel9.x86_64.rpm) |
| `set_user_14` | `4.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.8 KiB | [set_user_14-4.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/set_user_14-4.0.0-1.rhel9.x86_64.rpm) |
| `set_user_14` | `4.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.8 KiB | [set_user_14-4.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/set_user_14-4.1.0-1PGDG.rhel9.aarch64.rpm) |
| `set_user_14` | `4.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.4 KiB | [set_user_14-4.0.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/set_user_14-4.0.1-2.rhel9.aarch64.rpm) |
| `set_user_14` | `4.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.1 KiB | [set_user_14-4.0.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/set_user_14-4.0.0-1.rhel9.aarch64.rpm) |
| `set_user_14` | `3.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.0 KiB | [set_user_14-3.0.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/set_user_14-3.0.0-1.rhel9.aarch64.rpm) |
| `set_user_14` | `4.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.7 KiB | [set_user_14-4.1.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/set_user_14-4.1.0-2PGDG.rhel10.x86_64.rpm) |
| `set_user_14` | `4.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.3 KiB | [set_user_14-4.1.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10.0-aarch64/set_user_14-4.1.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-set-user` | `4.2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.7 KiB | [postgresql-14-set-user_4.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-14-set-user_4.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-set-user` | `4.2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 34.2 KiB | [postgresql-14-set-user_4.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-14-set-user_4.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-set-user` | `4.2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 34.7 KiB | [postgresql-14-set-user_4.2.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-14-set-user_4.2.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-set-user` | `4.2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.2 KiB | [postgresql-14-set-user_4.2.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-14-set-user_4.2.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-set-user` | `4.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.1 KiB | [postgresql-14-set-user_4.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-14-set-user_4.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-set-user` | `4.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 37.8 KiB | [postgresql-14-set-user_4.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-14-set-user_4.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-set-user` | `4.2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.2 KiB | [postgresql-14-set-user_4.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-14-set-user_4.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-set-user` | `4.2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 33.7 KiB | [postgresql-14-set-user_4.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-14-set-user_4.2.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `set_user_13` | `4.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.2 KiB | [set_user_13-4.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/set_user_13-4.1.0-1PGDG.rhel8.x86_64.rpm) |
| `set_user_13` | `4.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 25.9 KiB | [set_user_13-4.0.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/set_user_13-4.0.1-2.rhel8.x86_64.rpm) |
| `set_user_13` | `4.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 25.3 KiB | [set_user_13-4.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/set_user_13-4.0.0-1.rhel8.x86_64.rpm) |
| `set_user_13` | `3.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 25.1 KiB | [set_user_13-3.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/set_user_13-3.0.0-1.rhel8.x86_64.rpm) |
| `set_user_13` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.5 KiB | [set_user_13-2.0.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/set_user_13-2.0.1-1.rhel8.x86_64.rpm) |
| `set_user_13` | `4.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.0 KiB | [set_user_13-4.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/set_user_13-4.1.0-1PGDG.rhel8.aarch64.rpm) |
| `set_user_13` | `4.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.7 KiB | [set_user_13-4.0.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/set_user_13-4.0.1-2.rhel8.aarch64.rpm) |
| `set_user_13` | `4.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.2 KiB | [set_user_13-4.0.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/set_user_13-4.0.0-1.rhel8.aarch64.rpm) |
| `set_user_13` | `3.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 25.1 KiB | [set_user_13-3.0.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/set_user_13-3.0.0-1.rhel8.aarch64.rpm) |
| `set_user_13` | `4.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.3 KiB | [set_user_13-4.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/set_user_13-4.1.0-1PGDG.rhel9.x86_64.rpm) |
| `set_user_13` | `4.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.0 KiB | [set_user_13-4.0.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/set_user_13-4.0.1-2.rhel9.x86_64.rpm) |
| `set_user_13` | `4.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.8 KiB | [set_user_13-4.0.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/set_user_13-4.0.0-1.rhel9.x86_64.rpm) |
| `set_user_13` | `4.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.8 KiB | [set_user_13-4.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/set_user_13-4.1.0-1PGDG.rhel9.aarch64.rpm) |
| `set_user_13` | `4.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.4 KiB | [set_user_13-4.0.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/set_user_13-4.0.1-2.rhel9.aarch64.rpm) |
| `set_user_13` | `4.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.1 KiB | [set_user_13-4.0.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/set_user_13-4.0.0-1.rhel9.aarch64.rpm) |
| `set_user_13` | `3.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.0 KiB | [set_user_13-3.0.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/set_user_13-3.0.0-1.rhel9.aarch64.rpm) |
| `set_user_13` | `4.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.7 KiB | [set_user_13-4.1.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/set_user_13-4.1.0-2PGDG.rhel10.x86_64.rpm) |
| `set_user_13` | `4.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.3 KiB | [set_user_13-4.1.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10.0-aarch64/set_user_13-4.1.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-set-user` | `4.2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.4 KiB | [postgresql-13-set-user_4.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-13-set-user_4.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-13-set-user` | `4.2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 34.3 KiB | [postgresql-13-set-user_4.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-13-set-user_4.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-13-set-user` | `4.2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 34.5 KiB | [postgresql-13-set-user_4.2.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-13-set-user_4.2.0-1.pgdg13+1_amd64.deb) |
| `postgresql-13-set-user` | `4.2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.4 KiB | [postgresql-13-set-user_4.2.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-13-set-user_4.2.0-1.pgdg13+1_arm64.deb) |
| `postgresql-13-set-user` | `4.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 37.2 KiB | [postgresql-13-set-user_4.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-13-set-user_4.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-set-user` | `4.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 36.5 KiB | [postgresql-13-set-user_4.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-13-set-user_4.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-set-user` | `4.2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.1 KiB | [postgresql-13-set-user_4.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-13-set-user_4.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-set-user` | `4.2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 33.7 KiB | [postgresql-13-set-user_4.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-set-user/postgresql-13-set-user_4.2.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgaudit/set_user" title="Repository" icon="github" subtitle="github.com/pgaudit/set_user" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install set_user;		# install via package name, for the active PG version

pig install set_user -v 18;   # install for PG 18
pig install set_user -v 17;   # install for PG 17
pig install set_user -v 16;   # install for PG 16
pig install set_user -v 15;   # install for PG 15
pig install set_user -v 14;   # install for PG 14
pig install set_user -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION set_user;
```
