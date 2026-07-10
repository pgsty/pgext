---
title: "pgextwlist"
linkTitle: "pgextwlist"
description: "PostgreSQL Extension Whitelisting"
weight: 7390
categories: ["SEC"]
width: full
---

[**pgextwlist**](https://github.com/dimitri/pgextwlist) : PostgreSQL Extension Whitelisting


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7390** | {{< badge content="pgextwlist" link="https://github.com/dimitri/pgextwlist" >}} | {{< ext "pgextwlist" >}} | `1.20` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "ddlx" >}} {{< ext "pgdd" >}} {{< ext "pg_permissions" >}} {{< ext "adminpack" >}} {{< ext "pgaudit" >}} {{< ext "set_user" >}} {{< ext "pg_catcheck" >}} {{< ext "noset" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.20` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgextwlist` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.20` | {{< bg "18" "pgextwlist_18" "green" >}} {{< bg "17" "pgextwlist_17" "green" >}} {{< bg "16" "pgextwlist_16" "green" >}} {{< bg "15" "pgextwlist_15" "green" >}} {{< bg "14" "pgextwlist_14" "green" >}} | `pgextwlist_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.20` | {{< bg "18" "postgresql-18-pgextwlist" "green" >}} {{< bg "17" "postgresql-17-pgextwlist" "green" >}} {{< bg "16" "postgresql-16-pgextwlist" "green" >}} {{< bg "15" "postgresql-15-pgextwlist" "green" >}} {{< bg "14" "postgresql-14-pgextwlist" "green" >}} | `postgresql-$v-pgextwlist` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.20" "pgextwlist_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_14 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.20" "pgextwlist_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.20" "pgextwlist_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_14 : AVAIL 5" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.20" "pgextwlist_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_14 : AVAIL 5" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.20" "pgextwlist_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_14 : AVAIL 5" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.20" "pgextwlist_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.20" "pgextwlist_14 : AVAIL 5" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.20" "postgresql-18-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-17-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-16-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-15-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-14-pgextwlist : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.20" "postgresql-18-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-17-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-16-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-15-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-14-pgextwlist : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.20" "postgresql-18-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-17-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-16-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-15-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-14-pgextwlist : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.20" "postgresql-18-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-17-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-16-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-15-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-14-pgextwlist : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.20" "postgresql-18-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-17-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-16-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-15-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-14-pgextwlist : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.20" "postgresql-18-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-17-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-16-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-15-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-14-pgextwlist : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.20" "postgresql-18-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-17-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-16-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-15-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-14-pgextwlist : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.20" "postgresql-18-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-17-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-16-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-15-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-14-pgextwlist : AVAIL 2" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.20" "postgresql-18-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-17-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-16-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-15-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-14-pgextwlist : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.20" "postgresql-18-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-17-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-16-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-15-pgextwlist : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.20" "postgresql-14-pgextwlist : AVAIL 2" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgextwlist_18` | `1.20` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.7 KiB | [pgextwlist_18-1.20-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgextwlist_18-1.20-1PGDG.rhel8.10.x86_64.rpm) |
| `pgextwlist_18` | `1.19` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.3 KiB | [pgextwlist_18-1.19-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_18-1.19-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_18` | `1.19` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.7 KiB | [pgextwlist_18-1.19-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgextwlist_18-1.19-1PGDG.rhel8.10.x86_64.rpm) |
| `pgextwlist_18` | `1.20` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.3 KiB | [pgextwlist_18-1.20-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgextwlist_18-1.20-1PGDG.rhel8.10.aarch64.rpm) |
| `pgextwlist_18` | `1.19` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.2 KiB | [pgextwlist_18-1.19-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_18-1.19-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_18` | `1.19` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.4 KiB | [pgextwlist_18-1.19-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgextwlist_18-1.19-1PGDG.rhel8.10.aarch64.rpm) |
| `pgextwlist_18` | `1.20` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.8 KiB | [pgextwlist_18-1.20-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgextwlist_18-1.20-1PGDG.rhel9.8.x86_64.rpm) |
| `pgextwlist_18` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.4 KiB | [pgextwlist_18-1.19-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_18-1.19-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_18` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.8 KiB | [pgextwlist_18-1.19-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgextwlist_18-1.19-1PGDG.rhel9.8.x86_64.rpm) |
| `pgextwlist_18` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.8 KiB | [pgextwlist_18-1.19-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgextwlist_18-1.19-1PGDG.rhel9.7.x86_64.rpm) |
| `pgextwlist_18` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.9 KiB | [pgextwlist_18-1.19-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgextwlist_18-1.19-1PGDG.rhel9.6.x86_64.rpm) |
| `pgextwlist_18` | `1.20` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.9 KiB | [pgextwlist_18-1.20-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgextwlist_18-1.20-1PGDG.rhel9.8.aarch64.rpm) |
| `pgextwlist_18` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.0 KiB | [pgextwlist_18-1.19-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_18-1.19-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_18` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.1 KiB | [pgextwlist_18-1.19-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgextwlist_18-1.19-1PGDG.rhel9.8.aarch64.rpm) |
| `pgextwlist_18` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.1 KiB | [pgextwlist_18-1.19-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgextwlist_18-1.19-1PGDG.rhel9.7.aarch64.rpm) |
| `pgextwlist_18` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.2 KiB | [pgextwlist_18-1.19-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgextwlist_18-1.19-1PGDG.rhel9.6.aarch64.rpm) |
| `pgextwlist_18` | `1.20` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.8 KiB | [pgextwlist_18-1.20-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgextwlist_18-1.20-1PGDG.rhel10.2.x86_64.rpm) |
| `pgextwlist_18` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.2 KiB | [pgextwlist_18-1.19-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgextwlist_18-1.19-1PIGSTY.el10.x86_64.rpm) |
| `pgextwlist_18` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 20.8 KiB | [pgextwlist_18-1.19-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgextwlist_18-1.19-1PGDG.rhel10.2.x86_64.rpm) |
| `pgextwlist_18` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 20.8 KiB | [pgextwlist_18-1.19-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgextwlist_18-1.19-1PGDG.rhel10.1.x86_64.rpm) |
| `pgextwlist_18` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.2 KiB | [pgextwlist_18-1.19-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgextwlist_18-1.19-1PGDG.rhel10.0.x86_64.rpm) |
| `pgextwlist_18` | `1.20` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.3 KiB | [pgextwlist_18-1.20-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgextwlist_18-1.20-1PGDG.rhel10.2.aarch64.rpm) |
| `pgextwlist_18` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.2 KiB | [pgextwlist_18-1.19-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgextwlist_18-1.19-1PIGSTY.el10.aarch64.rpm) |
| `pgextwlist_18` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.4 KiB | [pgextwlist_18-1.19-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgextwlist_18-1.19-1PGDG.rhel10.2.aarch64.rpm) |
| `pgextwlist_18` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.4 KiB | [pgextwlist_18-1.19-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgextwlist_18-1.19-1PGDG.rhel10.1.aarch64.rpm) |
| `pgextwlist_18` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.4 KiB | [pgextwlist_18-1.19-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgextwlist_18-1.19-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-pgextwlist` | `1.20` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.6 KiB | [postgresql-18-pgextwlist_1.20-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.20-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pgextwlist` | `1.19` | [d12.x86_64](/os/d12.x86_64) | pgdg | 29.1 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pgextwlist` | `1.20` | [d12.aarch64](/os/d12.aarch64) | pgdg | 34.0 KiB | [postgresql-18-pgextwlist_1.20-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.20-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pgextwlist` | `1.19` | [d12.aarch64](/os/d12.aarch64) | pgdg | 28.7 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pgextwlist` | `1.20` | [d13.x86_64](/os/d13.x86_64) | pgdg | 34.6 KiB | [postgresql-18-pgextwlist_1.20-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.20-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pgextwlist` | `1.19` | [d13.x86_64](/os/d13.x86_64) | pgdg | 29.1 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pgextwlist` | `1.20` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.1 KiB | [postgresql-18-pgextwlist_1.20-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.20-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pgextwlist` | `1.19` | [d13.aarch64](/os/d13.aarch64) | pgdg | 28.8 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pgextwlist` | `1.20` | [u22.x86_64](/os/u22.x86_64) | pgdg | 35.2 KiB | [postgresql-18-pgextwlist_1.20-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.20-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgextwlist` | `1.19` | [u22.x86_64](/os/u22.x86_64) | pgdg | 30.1 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgextwlist` | `1.20` | [u22.aarch64](/os/u22.aarch64) | pgdg | 34.8 KiB | [postgresql-18-pgextwlist_1.20-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.20-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgextwlist` | `1.19` | [u22.aarch64](/os/u22.aarch64) | pgdg | 29.3 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgextwlist` | `1.20` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.6 KiB | [postgresql-18-pgextwlist_1.20-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.20-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgextwlist` | `1.19` | [u24.x86_64](/os/u24.x86_64) | pgdg | 29.2 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgextwlist` | `1.20` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.0 KiB | [postgresql-18-pgextwlist_1.20-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.20-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgextwlist` | `1.19` | [u24.aarch64](/os/u24.aarch64) | pgdg | 28.6 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgextwlist` | `1.20` | [u26.x86_64](/os/u26.x86_64) | pgdg | 33.9 KiB | [postgresql-18-pgextwlist_1.20-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.20-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgextwlist` | `1.19` | [u26.x86_64](/os/u26.x86_64) | pgdg | 29.1 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgextwlist` | `1.20` | [u26.aarch64](/os/u26.aarch64) | pgdg | 33.4 KiB | [postgresql-18-pgextwlist_1.20-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.20-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pgextwlist` | `1.19` | [u26.aarch64](/os/u26.aarch64) | pgdg | 28.6 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgextwlist_17` | `1.20` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.7 KiB | [pgextwlist_17-1.20-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgextwlist_17-1.20-1PGDG.rhel8.10.x86_64.rpm) |
| `pgextwlist_17` | `1.19` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.3 KiB | [pgextwlist_17-1.19-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_17-1.19-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_17` | `1.19` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.7 KiB | [pgextwlist_17-1.19-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgextwlist_17-1.19-1PGDG.rhel8.10.x86_64.rpm) |
| `pgextwlist_17` | `1.20` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.3 KiB | [pgextwlist_17-1.20-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgextwlist_17-1.20-1PGDG.rhel8.10.aarch64.rpm) |
| `pgextwlist_17` | `1.19` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.2 KiB | [pgextwlist_17-1.19-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_17-1.19-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_17` | `1.19` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.4 KiB | [pgextwlist_17-1.19-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgextwlist_17-1.19-1PGDG.rhel8.10.aarch64.rpm) |
| `pgextwlist_17` | `1.20` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.8 KiB | [pgextwlist_17-1.20-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgextwlist_17-1.20-1PGDG.rhel9.8.x86_64.rpm) |
| `pgextwlist_17` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.4 KiB | [pgextwlist_17-1.19-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_17-1.19-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_17` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.8 KiB | [pgextwlist_17-1.19-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgextwlist_17-1.19-1PGDG.rhel9.8.x86_64.rpm) |
| `pgextwlist_17` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.8 KiB | [pgextwlist_17-1.19-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgextwlist_17-1.19-1PGDG.rhel9.7.x86_64.rpm) |
| `pgextwlist_17` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.9 KiB | [pgextwlist_17-1.19-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgextwlist_17-1.19-1PGDG.rhel9.6.x86_64.rpm) |
| `pgextwlist_17` | `1.20` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.9 KiB | [pgextwlist_17-1.20-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgextwlist_17-1.20-1PGDG.rhel9.8.aarch64.rpm) |
| `pgextwlist_17` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.0 KiB | [pgextwlist_17-1.19-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_17-1.19-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_17` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.1 KiB | [pgextwlist_17-1.19-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgextwlist_17-1.19-1PGDG.rhel9.8.aarch64.rpm) |
| `pgextwlist_17` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.1 KiB | [pgextwlist_17-1.19-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgextwlist_17-1.19-1PGDG.rhel9.7.aarch64.rpm) |
| `pgextwlist_17` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.2 KiB | [pgextwlist_17-1.19-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgextwlist_17-1.19-1PGDG.rhel9.6.aarch64.rpm) |
| `pgextwlist_17` | `1.20` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.9 KiB | [pgextwlist_17-1.20-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgextwlist_17-1.20-1PGDG.rhel10.2.x86_64.rpm) |
| `pgextwlist_17` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.3 KiB | [pgextwlist_17-1.19-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgextwlist_17-1.19-1PIGSTY.el10.x86_64.rpm) |
| `pgextwlist_17` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 20.8 KiB | [pgextwlist_17-1.19-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgextwlist_17-1.19-1PGDG.rhel10.2.x86_64.rpm) |
| `pgextwlist_17` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 20.8 KiB | [pgextwlist_17-1.19-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgextwlist_17-1.19-1PGDG.rhel10.1.x86_64.rpm) |
| `pgextwlist_17` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.2 KiB | [pgextwlist_17-1.19-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgextwlist_17-1.19-1PGDG.rhel10.0.x86_64.rpm) |
| `pgextwlist_17` | `1.20` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.3 KiB | [pgextwlist_17-1.20-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgextwlist_17-1.20-1PGDG.rhel10.2.aarch64.rpm) |
| `pgextwlist_17` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.2 KiB | [pgextwlist_17-1.19-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgextwlist_17-1.19-1PIGSTY.el10.aarch64.rpm) |
| `pgextwlist_17` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.4 KiB | [pgextwlist_17-1.19-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgextwlist_17-1.19-1PGDG.rhel10.2.aarch64.rpm) |
| `pgextwlist_17` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.4 KiB | [pgextwlist_17-1.19-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgextwlist_17-1.19-1PGDG.rhel10.1.aarch64.rpm) |
| `pgextwlist_17` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.4 KiB | [pgextwlist_17-1.19-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgextwlist_17-1.19-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-17-pgextwlist` | `1.20` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.5 KiB | [postgresql-17-pgextwlist_1.20-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.20-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pgextwlist` | `1.19` | [d12.x86_64](/os/d12.x86_64) | pgdg | 29.1 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pgextwlist` | `1.20` | [d12.aarch64](/os/d12.aarch64) | pgdg | 33.9 KiB | [postgresql-17-pgextwlist_1.20-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.20-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pgextwlist` | `1.19` | [d12.aarch64](/os/d12.aarch64) | pgdg | 28.7 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pgextwlist` | `1.20` | [d13.x86_64](/os/d13.x86_64) | pgdg | 34.6 KiB | [postgresql-17-pgextwlist_1.20-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.20-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pgextwlist` | `1.19` | [d13.x86_64](/os/d13.x86_64) | pgdg | 29.0 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pgextwlist` | `1.20` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.0 KiB | [postgresql-17-pgextwlist_1.20-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.20-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pgextwlist` | `1.19` | [d13.aarch64](/os/d13.aarch64) | pgdg | 28.8 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pgextwlist` | `1.20` | [u22.x86_64](/os/u22.x86_64) | pgdg | 44.1 KiB | [postgresql-17-pgextwlist_1.20-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.20-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgextwlist` | `1.19` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.3 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgextwlist` | `1.20` | [u22.aarch64](/os/u22.aarch64) | pgdg | 43.3 KiB | [postgresql-17-pgextwlist_1.20-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.20-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgextwlist` | `1.19` | [u22.aarch64](/os/u22.aarch64) | pgdg | 37.6 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgextwlist` | `1.20` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.6 KiB | [postgresql-17-pgextwlist_1.20-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.20-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgextwlist` | `1.19` | [u24.x86_64](/os/u24.x86_64) | pgdg | 29.2 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgextwlist` | `1.20` | [u24.aarch64](/os/u24.aarch64) | pgdg | 33.9 KiB | [postgresql-17-pgextwlist_1.20-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.20-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgextwlist` | `1.19` | [u24.aarch64](/os/u24.aarch64) | pgdg | 28.6 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgextwlist` | `1.20` | [u26.x86_64](/os/u26.x86_64) | pgdg | 33.9 KiB | [postgresql-17-pgextwlist_1.20-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.20-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgextwlist` | `1.19` | [u26.x86_64](/os/u26.x86_64) | pgdg | 29.1 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgextwlist` | `1.20` | [u26.aarch64](/os/u26.aarch64) | pgdg | 33.4 KiB | [postgresql-17-pgextwlist_1.20-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.20-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pgextwlist` | `1.19` | [u26.aarch64](/os/u26.aarch64) | pgdg | 28.7 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgextwlist_16` | `1.20` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.7 KiB | [pgextwlist_16-1.20-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgextwlist_16-1.20-1PGDG.rhel8.10.x86_64.rpm) |
| `pgextwlist_16` | `1.19` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.3 KiB | [pgextwlist_16-1.19-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_16-1.19-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_16` | `1.19` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.7 KiB | [pgextwlist_16-1.19-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgextwlist_16-1.19-1PGDG.rhel8.10.x86_64.rpm) |
| `pgextwlist_16` | `1.20` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.3 KiB | [pgextwlist_16-1.20-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgextwlist_16-1.20-1PGDG.rhel8.10.aarch64.rpm) |
| `pgextwlist_16` | `1.19` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.2 KiB | [pgextwlist_16-1.19-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_16-1.19-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_16` | `1.19` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.4 KiB | [pgextwlist_16-1.19-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgextwlist_16-1.19-1PGDG.rhel8.10.aarch64.rpm) |
| `pgextwlist_16` | `1.20` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.8 KiB | [pgextwlist_16-1.20-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgextwlist_16-1.20-1PGDG.rhel9.8.x86_64.rpm) |
| `pgextwlist_16` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.4 KiB | [pgextwlist_16-1.19-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_16-1.19-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_16` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.8 KiB | [pgextwlist_16-1.19-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgextwlist_16-1.19-1PGDG.rhel9.8.x86_64.rpm) |
| `pgextwlist_16` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.8 KiB | [pgextwlist_16-1.19-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgextwlist_16-1.19-1PGDG.rhel9.7.x86_64.rpm) |
| `pgextwlist_16` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.9 KiB | [pgextwlist_16-1.19-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgextwlist_16-1.19-1PGDG.rhel9.6.x86_64.rpm) |
| `pgextwlist_16` | `1.20` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.9 KiB | [pgextwlist_16-1.20-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgextwlist_16-1.20-1PGDG.rhel9.8.aarch64.rpm) |
| `pgextwlist_16` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.0 KiB | [pgextwlist_16-1.19-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_16-1.19-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_16` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.1 KiB | [pgextwlist_16-1.19-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgextwlist_16-1.19-1PGDG.rhel9.8.aarch64.rpm) |
| `pgextwlist_16` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.1 KiB | [pgextwlist_16-1.19-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgextwlist_16-1.19-1PGDG.rhel9.7.aarch64.rpm) |
| `pgextwlist_16` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.2 KiB | [pgextwlist_16-1.19-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgextwlist_16-1.19-1PGDG.rhel9.6.aarch64.rpm) |
| `pgextwlist_16` | `1.20` | [el10.x86_64](/os/el10.x86_64) | pgdg | 22.8 KiB | [pgextwlist_16-1.20-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgextwlist_16-1.20-1PGDG.rhel10.2.x86_64.rpm) |
| `pgextwlist_16` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.2 KiB | [pgextwlist_16-1.19-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgextwlist_16-1.19-1PIGSTY.el10.x86_64.rpm) |
| `pgextwlist_16` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 20.8 KiB | [pgextwlist_16-1.19-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgextwlist_16-1.19-1PGDG.rhel10.2.x86_64.rpm) |
| `pgextwlist_16` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 20.8 KiB | [pgextwlist_16-1.19-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgextwlist_16-1.19-1PGDG.rhel10.1.x86_64.rpm) |
| `pgextwlist_16` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.2 KiB | [pgextwlist_16-1.19-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgextwlist_16-1.19-1PGDG.rhel10.0.x86_64.rpm) |
| `pgextwlist_16` | `1.20` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.3 KiB | [pgextwlist_16-1.20-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgextwlist_16-1.20-1PGDG.rhel10.2.aarch64.rpm) |
| `pgextwlist_16` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.2 KiB | [pgextwlist_16-1.19-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgextwlist_16-1.19-1PIGSTY.el10.aarch64.rpm) |
| `pgextwlist_16` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.4 KiB | [pgextwlist_16-1.19-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgextwlist_16-1.19-1PGDG.rhel10.2.aarch64.rpm) |
| `pgextwlist_16` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.4 KiB | [pgextwlist_16-1.19-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgextwlist_16-1.19-1PGDG.rhel10.1.aarch64.rpm) |
| `pgextwlist_16` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.4 KiB | [pgextwlist_16-1.19-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgextwlist_16-1.19-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-16-pgextwlist` | `1.20` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.5 KiB | [postgresql-16-pgextwlist_1.20-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.20-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pgextwlist` | `1.19` | [d12.x86_64](/os/d12.x86_64) | pgdg | 29.1 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pgextwlist` | `1.20` | [d12.aarch64](/os/d12.aarch64) | pgdg | 33.9 KiB | [postgresql-16-pgextwlist_1.20-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.20-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pgextwlist` | `1.19` | [d12.aarch64](/os/d12.aarch64) | pgdg | 28.7 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pgextwlist` | `1.20` | [d13.x86_64](/os/d13.x86_64) | pgdg | 34.6 KiB | [postgresql-16-pgextwlist_1.20-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.20-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pgextwlist` | `1.19` | [d13.x86_64](/os/d13.x86_64) | pgdg | 29.0 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pgextwlist` | `1.20` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.1 KiB | [postgresql-16-pgextwlist_1.20-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.20-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pgextwlist` | `1.19` | [d13.aarch64](/os/d13.aarch64) | pgdg | 28.8 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pgextwlist` | `1.20` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.2 KiB | [postgresql-16-pgextwlist_1.20-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.20-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgextwlist` | `1.19` | [u22.x86_64](/os/u22.x86_64) | pgdg | 37.7 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgextwlist` | `1.20` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.7 KiB | [postgresql-16-pgextwlist_1.20-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.20-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgextwlist` | `1.19` | [u22.aarch64](/os/u22.aarch64) | pgdg | 37.0 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgextwlist` | `1.20` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.6 KiB | [postgresql-16-pgextwlist_1.20-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.20-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgextwlist` | `1.19` | [u24.x86_64](/os/u24.x86_64) | pgdg | 29.2 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgextwlist` | `1.20` | [u24.aarch64](/os/u24.aarch64) | pgdg | 33.9 KiB | [postgresql-16-pgextwlist_1.20-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.20-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgextwlist` | `1.19` | [u24.aarch64](/os/u24.aarch64) | pgdg | 28.6 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgextwlist` | `1.20` | [u26.x86_64](/os/u26.x86_64) | pgdg | 33.9 KiB | [postgresql-16-pgextwlist_1.20-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.20-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgextwlist` | `1.19` | [u26.x86_64](/os/u26.x86_64) | pgdg | 29.2 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgextwlist` | `1.20` | [u26.aarch64](/os/u26.aarch64) | pgdg | 33.4 KiB | [postgresql-16-pgextwlist_1.20-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.20-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pgextwlist` | `1.19` | [u26.aarch64](/os/u26.aarch64) | pgdg | 28.7 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgextwlist_15` | `1.20` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.8 KiB | [pgextwlist_15-1.20-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgextwlist_15-1.20-1PGDG.rhel8.10.x86_64.rpm) |
| `pgextwlist_15` | `1.19` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.3 KiB | [pgextwlist_15-1.19-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_15-1.19-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_15` | `1.19` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.7 KiB | [pgextwlist_15-1.19-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgextwlist_15-1.19-1PGDG.rhel8.10.x86_64.rpm) |
| `pgextwlist_15` | `1.20` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.3 KiB | [pgextwlist_15-1.20-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgextwlist_15-1.20-1PGDG.rhel8.10.aarch64.rpm) |
| `pgextwlist_15` | `1.19` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.2 KiB | [pgextwlist_15-1.19-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_15-1.19-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_15` | `1.19` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pgextwlist_15-1.19-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgextwlist_15-1.19-1PGDG.rhel8.10.aarch64.rpm) |
| `pgextwlist_15` | `1.20` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.1 KiB | [pgextwlist_15-1.20-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgextwlist_15-1.20-1PGDG.rhel9.8.x86_64.rpm) |
| `pgextwlist_15` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.4 KiB | [pgextwlist_15-1.19-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_15-1.19-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_15` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.9 KiB | [pgextwlist_15-1.19-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgextwlist_15-1.19-1PGDG.rhel9.8.x86_64.rpm) |
| `pgextwlist_15` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.9 KiB | [pgextwlist_15-1.19-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgextwlist_15-1.19-1PGDG.rhel9.7.x86_64.rpm) |
| `pgextwlist_15` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.0 KiB | [pgextwlist_15-1.19-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgextwlist_15-1.19-1PGDG.rhel9.6.x86_64.rpm) |
| `pgextwlist_15` | `1.20` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.2 KiB | [pgextwlist_15-1.20-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgextwlist_15-1.20-1PGDG.rhel9.8.aarch64.rpm) |
| `pgextwlist_15` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.1 KiB | [pgextwlist_15-1.19-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_15-1.19-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_15` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.2 KiB | [pgextwlist_15-1.19-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgextwlist_15-1.19-1PGDG.rhel9.8.aarch64.rpm) |
| `pgextwlist_15` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.2 KiB | [pgextwlist_15-1.19-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgextwlist_15-1.19-1PGDG.rhel9.7.aarch64.rpm) |
| `pgextwlist_15` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.4 KiB | [pgextwlist_15-1.19-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgextwlist_15-1.19-1PGDG.rhel9.6.aarch64.rpm) |
| `pgextwlist_15` | `1.20` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.1 KiB | [pgextwlist_15-1.20-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgextwlist_15-1.20-1PGDG.rhel10.2.x86_64.rpm) |
| `pgextwlist_15` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.4 KiB | [pgextwlist_15-1.19-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgextwlist_15-1.19-1PIGSTY.el10.x86_64.rpm) |
| `pgextwlist_15` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 20.9 KiB | [pgextwlist_15-1.19-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgextwlist_15-1.19-1PGDG.rhel10.2.x86_64.rpm) |
| `pgextwlist_15` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 20.9 KiB | [pgextwlist_15-1.19-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgextwlist_15-1.19-1PGDG.rhel10.1.x86_64.rpm) |
| `pgextwlist_15` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.3 KiB | [pgextwlist_15-1.19-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgextwlist_15-1.19-1PGDG.rhel10.0.x86_64.rpm) |
| `pgextwlist_15` | `1.20` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.5 KiB | [pgextwlist_15-1.20-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgextwlist_15-1.20-1PGDG.rhel10.2.aarch64.rpm) |
| `pgextwlist_15` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.3 KiB | [pgextwlist_15-1.19-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgextwlist_15-1.19-1PIGSTY.el10.aarch64.rpm) |
| `pgextwlist_15` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.6 KiB | [pgextwlist_15-1.19-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgextwlist_15-1.19-1PGDG.rhel10.2.aarch64.rpm) |
| `pgextwlist_15` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.6 KiB | [pgextwlist_15-1.19-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgextwlist_15-1.19-1PGDG.rhel10.1.aarch64.rpm) |
| `pgextwlist_15` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.6 KiB | [pgextwlist_15-1.19-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgextwlist_15-1.19-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-15-pgextwlist` | `1.20` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.5 KiB | [postgresql-15-pgextwlist_1.20-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.20-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pgextwlist` | `1.19` | [d12.x86_64](/os/d12.x86_64) | pgdg | 28.9 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pgextwlist` | `1.20` | [d12.aarch64](/os/d12.aarch64) | pgdg | 33.8 KiB | [postgresql-15-pgextwlist_1.20-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.20-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pgextwlist` | `1.19` | [d12.aarch64](/os/d12.aarch64) | pgdg | 28.4 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pgextwlist` | `1.20` | [d13.x86_64](/os/d13.x86_64) | pgdg | 34.4 KiB | [postgresql-15-pgextwlist_1.20-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.20-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pgextwlist` | `1.19` | [d13.x86_64](/os/d13.x86_64) | pgdg | 28.9 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pgextwlist` | `1.20` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.0 KiB | [postgresql-15-pgextwlist_1.20-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.20-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pgextwlist` | `1.19` | [d13.aarch64](/os/d13.aarch64) | pgdg | 28.7 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pgextwlist` | `1.20` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.0 KiB | [postgresql-15-pgextwlist_1.20-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.20-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgextwlist` | `1.19` | [u22.x86_64](/os/u22.x86_64) | pgdg | 37.5 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgextwlist` | `1.20` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.6 KiB | [postgresql-15-pgextwlist_1.20-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.20-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgextwlist` | `1.19` | [u22.aarch64](/os/u22.aarch64) | pgdg | 36.8 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgextwlist` | `1.20` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.5 KiB | [postgresql-15-pgextwlist_1.20-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.20-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgextwlist` | `1.19` | [u24.x86_64](/os/u24.x86_64) | pgdg | 29.0 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgextwlist` | `1.20` | [u24.aarch64](/os/u24.aarch64) | pgdg | 33.8 KiB | [postgresql-15-pgextwlist_1.20-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.20-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgextwlist` | `1.19` | [u24.aarch64](/os/u24.aarch64) | pgdg | 28.5 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgextwlist` | `1.20` | [u26.x86_64](/os/u26.x86_64) | pgdg | 33.9 KiB | [postgresql-15-pgextwlist_1.20-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.20-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgextwlist` | `1.19` | [u26.x86_64](/os/u26.x86_64) | pgdg | 29.0 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgextwlist` | `1.20` | [u26.aarch64](/os/u26.aarch64) | pgdg | 33.4 KiB | [postgresql-15-pgextwlist_1.20-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.20-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pgextwlist` | `1.19` | [u26.aarch64](/os/u26.aarch64) | pgdg | 28.6 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgextwlist_14` | `1.20` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.8 KiB | [pgextwlist_14-1.20-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgextwlist_14-1.20-1PGDG.rhel8.10.x86_64.rpm) |
| `pgextwlist_14` | `1.19` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.3 KiB | [pgextwlist_14-1.19-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_14-1.19-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_14` | `1.19` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.7 KiB | [pgextwlist_14-1.19-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgextwlist_14-1.19-1PGDG.rhel8.10.x86_64.rpm) |
| `pgextwlist_14` | `1.20` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.3 KiB | [pgextwlist_14-1.20-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgextwlist_14-1.20-1PGDG.rhel8.10.aarch64.rpm) |
| `pgextwlist_14` | `1.19` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.2 KiB | [pgextwlist_14-1.19-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_14-1.19-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_14` | `1.19` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.3 KiB | [pgextwlist_14-1.19-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgextwlist_14-1.19-1PGDG.rhel8.10.aarch64.rpm) |
| `pgextwlist_14` | `1.20` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.1 KiB | [pgextwlist_14-1.20-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgextwlist_14-1.20-1PGDG.rhel9.8.x86_64.rpm) |
| `pgextwlist_14` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.4 KiB | [pgextwlist_14-1.19-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_14-1.19-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_14` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.9 KiB | [pgextwlist_14-1.19-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgextwlist_14-1.19-1PGDG.rhel9.8.x86_64.rpm) |
| `pgextwlist_14` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.9 KiB | [pgextwlist_14-1.19-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgextwlist_14-1.19-1PGDG.rhel9.7.x86_64.rpm) |
| `pgextwlist_14` | `1.19` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.0 KiB | [pgextwlist_14-1.19-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgextwlist_14-1.19-1PGDG.rhel9.6.x86_64.rpm) |
| `pgextwlist_14` | `1.20` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.2 KiB | [pgextwlist_14-1.20-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgextwlist_14-1.20-1PGDG.rhel9.8.aarch64.rpm) |
| `pgextwlist_14` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.0 KiB | [pgextwlist_14-1.19-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_14-1.19-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_14` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.2 KiB | [pgextwlist_14-1.19-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgextwlist_14-1.19-1PGDG.rhel9.8.aarch64.rpm) |
| `pgextwlist_14` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.2 KiB | [pgextwlist_14-1.19-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgextwlist_14-1.19-1PGDG.rhel9.7.aarch64.rpm) |
| `pgextwlist_14` | `1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.3 KiB | [pgextwlist_14-1.19-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgextwlist_14-1.19-1PGDG.rhel9.6.aarch64.rpm) |
| `pgextwlist_14` | `1.20` | [el10.x86_64](/os/el10.x86_64) | pgdg | 23.1 KiB | [pgextwlist_14-1.20-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgextwlist_14-1.20-1PGDG.rhel10.2.x86_64.rpm) |
| `pgextwlist_14` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.3 KiB | [pgextwlist_14-1.19-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgextwlist_14-1.19-1PIGSTY.el10.x86_64.rpm) |
| `pgextwlist_14` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 20.9 KiB | [pgextwlist_14-1.19-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgextwlist_14-1.19-1PGDG.rhel10.2.x86_64.rpm) |
| `pgextwlist_14` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.0 KiB | [pgextwlist_14-1.19-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgextwlist_14-1.19-1PGDG.rhel10.1.x86_64.rpm) |
| `pgextwlist_14` | `1.19` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.3 KiB | [pgextwlist_14-1.19-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgextwlist_14-1.19-1PGDG.rhel10.0.x86_64.rpm) |
| `pgextwlist_14` | `1.20` | [el10.aarch64](/os/el10.aarch64) | pgdg | 22.4 KiB | [pgextwlist_14-1.20-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgextwlist_14-1.20-1PGDG.rhel10.2.aarch64.rpm) |
| `pgextwlist_14` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.2 KiB | [pgextwlist_14-1.19-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgextwlist_14-1.19-1PIGSTY.el10.aarch64.rpm) |
| `pgextwlist_14` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.5 KiB | [pgextwlist_14-1.19-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgextwlist_14-1.19-1PGDG.rhel10.2.aarch64.rpm) |
| `pgextwlist_14` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.5 KiB | [pgextwlist_14-1.19-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgextwlist_14-1.19-1PGDG.rhel10.1.aarch64.rpm) |
| `pgextwlist_14` | `1.19` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.5 KiB | [pgextwlist_14-1.19-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgextwlist_14-1.19-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-14-pgextwlist` | `1.20` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.3 KiB | [postgresql-14-pgextwlist_1.20-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.20-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pgextwlist` | `1.19` | [d12.x86_64](/os/d12.x86_64) | pgdg | 28.8 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pgextwlist` | `1.20` | [d12.aarch64](/os/d12.aarch64) | pgdg | 33.8 KiB | [postgresql-14-pgextwlist_1.20-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.20-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pgextwlist` | `1.19` | [d12.aarch64](/os/d12.aarch64) | pgdg | 28.4 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pgextwlist` | `1.20` | [d13.x86_64](/os/d13.x86_64) | pgdg | 34.4 KiB | [postgresql-14-pgextwlist_1.20-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.20-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pgextwlist` | `1.19` | [d13.x86_64](/os/d13.x86_64) | pgdg | 28.8 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pgextwlist` | `1.20` | [d13.aarch64](/os/d13.aarch64) | pgdg | 33.9 KiB | [postgresql-14-pgextwlist_1.20-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.20-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pgextwlist` | `1.19` | [d13.aarch64](/os/d13.aarch64) | pgdg | 28.5 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pgextwlist` | `1.20` | [u22.x86_64](/os/u22.x86_64) | pgdg | 43.1 KiB | [postgresql-14-pgextwlist_1.20-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.20-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgextwlist` | `1.19` | [u22.x86_64](/os/u22.x86_64) | pgdg | 37.5 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgextwlist` | `1.20` | [u22.aarch64](/os/u22.aarch64) | pgdg | 42.5 KiB | [postgresql-14-pgextwlist_1.20-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.20-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgextwlist` | `1.19` | [u22.aarch64](/os/u22.aarch64) | pgdg | 36.8 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgextwlist` | `1.20` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.5 KiB | [postgresql-14-pgextwlist_1.20-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.20-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgextwlist` | `1.19` | [u24.x86_64](/os/u24.x86_64) | pgdg | 29.0 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgextwlist` | `1.20` | [u24.aarch64](/os/u24.aarch64) | pgdg | 33.7 KiB | [postgresql-14-pgextwlist_1.20-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.20-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgextwlist` | `1.19` | [u24.aarch64](/os/u24.aarch64) | pgdg | 28.4 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgextwlist` | `1.20` | [u26.x86_64](/os/u26.x86_64) | pgdg | 33.9 KiB | [postgresql-14-pgextwlist_1.20-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.20-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgextwlist` | `1.19` | [u26.x86_64](/os/u26.x86_64) | pgdg | 28.9 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgextwlist` | `1.20` | [u26.aarch64](/os/u26.aarch64) | pgdg | 33.3 KiB | [postgresql-14-pgextwlist_1.20-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.20-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pgextwlist` | `1.19` | [u26.aarch64](/os/u26.aarch64) | pgdg | 28.5 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dimitri/pgextwlist" title="Repository" icon="github" subtitle="github.com/dimitri/pgextwlist" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgextwlist-1.19.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgextwlist;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgextwlist;		# install via package name, for the active PG version

pig install pgextwlist -v 18;   # install for PG 18
pig install pgextwlist -v 17;   # install for PG 17
pig install pgextwlist -v 16;   # install for PG 16
pig install pgextwlist -v 15;   # install for PG 15
pig install pgextwlist -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pgextwlist';
```


This extension does not need `CREATE EXTENSION` DDL command






## Usage

> [pgextwlist: PostgreSQL extension whitelisting](https://github.com/dimitri/pgextwlist)

`pgextwlist` implements extension whitelisting: only explicitly allowed extensions can be installed, and whitelisted extensions are installed with superuser privileges even when requested by non-superusers.

### Configuration

Add to `postgresql.conf`:

```ini
local_preload_libraries = 'pgextwlist'
extwlist.extensions = 'hstore,cube,pg_stat_statements'
```

Or per-role:

```sql
ALTER ROLE adminuser SET extwlist.extensions = 'pg_stat_statements, postgis';
```

| Parameter | Description |
|-----------|-------------|
| `extwlist.extensions` | Comma-separated list of whitelisted extensions |
| `extwlist.custom_path` | Filesystem path for custom pre/post scripts |

### Behavior

Non-superusers can install whitelisted extensions:

```sql
-- Allowed (hstore is whitelisted)
CREATE EXTENSION hstore;

-- Blocked (not whitelisted)
CREATE EXTENSION earthdistance;
-- ERROR: extension "earthdistance" is not whitelisted
```

Operations `CREATE EXTENSION`, `DROP EXTENSION`, `ALTER EXTENSION ... UPDATE`, and `COMMENT ON EXTENSION` are run as superuser for whitelisted extensions.

### Custom Scripts

Place scripts in `${extwlist.custom_path}/extname/`:

| Script | When |
|--------|------|
| `before--1.0.sql` | Before installing version 1.0 |
| `before-create.sql` | Before CREATE (fallback) |
| `after--1.0.sql` | After installing version 1.0 |
| `after-create.sql` | After CREATE (fallback) |
| `before-update.sql` / `after-update.sql` | Around ALTER EXTENSION UPDATE |
| `before-drop.sql` / `after-drop.sql` | Around DROP EXTENSION |

Custom scripts support template variables: `@extschema@`, `@current_user@`, `@database_owner@`.
