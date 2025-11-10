---
title: "plproxy"
linkTitle: "plproxy"
description: "Database partitioning implemented as procedural language"
weight: 2520
categories: ["OLAP"]
width: full
---

[**plproxy**](https://github.com/plproxy/plproxy)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2520** | {{< badge content="plproxy" link="https://github.com/plproxy/plproxy" >}} | {{< ext "plproxy" >}} | `2.11.0` | {{< category "OLAP" >}} | {{< license "BSD 0-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "citus" >}} {{< ext "dblink" >}} {{< ext "postgres_fdw" >}} {{< ext "pg_partman" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} {{< ext "citus_columnar" >}} {{< ext "columnar" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/plproxy" >}} | `2.11.0` | {{< bg "18" "plproxy_18*" "green" >}} {{< bg "17" "plproxy_17*" "green" >}} {{< bg "16" "plproxy_16*" "green" >}} {{< bg "15" "plproxy_15*" "green" >}} {{< bg "14" "plproxy_14*" "green" >}} {{< bg "13" "plproxy_13*" "green" >}} | `plproxy_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/plproxy" >}} | `2.11.0` | {{< bg "18" "postgresql-18-plproxy" "green" >}} {{< bg "17" "postgresql-17-plproxy" "green" >}} {{< bg "16" "postgresql-16-plproxy" "green" >}} {{< bg "15" "postgresql-15-plproxy" "green" >}} {{< bg "14" "postgresql-14-plproxy" "green" >}} {{< bg "13" "postgresql-13-plproxy" "green" >}} | `postgresql-$v-plproxy` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.11.0" "plproxy_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_13 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.11.0" "plproxy_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_13 : AVAIL 3" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.11.0" "plproxy_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.11.0" "plproxy_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.11.0" "plproxy_13 : AVAIL 3" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.11.0" "plproxy_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.11.0" "plproxy_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "plproxy_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.11.0" "postgresql-18-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-17-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-16-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-15-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-14-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-13-plproxy : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.11.0" "postgresql-18-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-17-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-16-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-15-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-14-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-13-plproxy : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.11.0" "postgresql-18-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-17-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-16-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-15-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-14-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-13-plproxy : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.11.0" "postgresql-18-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-17-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-16-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-15-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-14-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-13-plproxy : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.11.0" "postgresql-18-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-17-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-16-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-15-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-14-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-13-plproxy : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.11.0" "postgresql-18-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-17-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-16-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-15-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-14-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-13-plproxy : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.11.0" "postgresql-18-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-17-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-16-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-15-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-14-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-13-plproxy : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.11.0" "postgresql-18-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-17-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-16-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-15-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-14-plproxy : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.11.0" "postgresql-13-plproxy : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plproxy_18` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 44.4 KiB | [plproxy_18-2.11.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plproxy_18-2.11.0-1PIGSTY.el8.x86_64.rpm) |
| `plproxy_18` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 48.2 KiB | [plproxy_18-2.11.0-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plproxy_18-2.11.0-4PGDG.rhel8.x86_64.rpm) |
| `plproxy_18` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.0 KiB | [plproxy_18-2.11.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plproxy_18-2.11.0-1PIGSTY.el8.aarch64.rpm) |
| `plproxy_18` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.8 KiB | [plproxy_18-2.11.0-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plproxy_18-2.11.0-4PGDG.rhel8.aarch64.rpm) |
| `plproxy_18` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.8 KiB | [plproxy_18-2.11.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plproxy_18-2.11.0-1PIGSTY.el9.x86_64.rpm) |
| `plproxy_18` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.8 KiB | [plproxy_18-2.11.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plproxy_18-2.11.0-4PGDG.rhel9.x86_64.rpm) |
| `plproxy_18` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.4 KiB | [plproxy_18-2.11.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plproxy_18-2.11.0-1PIGSTY.el9.aarch64.rpm) |
| `plproxy_18` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.4 KiB | [plproxy_18-2.11.0-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plproxy_18-2.11.0-4PGDG.rhel9.aarch64.rpm) |
| `plproxy_18` | `2.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.7 KiB | [plproxy_18-2.11.0-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plproxy_18-2.11.0-4PGDG.rhel10.x86_64.rpm) |
| `plproxy_18` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.9 KiB | [plproxy_18-2.11.0-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plproxy_18-2.11.0-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-plproxy` | `2.11.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 133.8 KiB | [postgresql-18-plproxy_2.11.0-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-18-plproxy_2.11.0-13.pgdg12+1_amd64.deb) |
| `postgresql-18-plproxy` | `2.11.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 130.0 KiB | [postgresql-18-plproxy_2.11.0-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-18-plproxy_2.11.0-13.pgdg12+1_arm64.deb) |
| `postgresql-18-plproxy` | `2.11.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 133.9 KiB | [postgresql-18-plproxy_2.11.0-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-18-plproxy_2.11.0-13.pgdg13+1_amd64.deb) |
| `postgresql-18-plproxy` | `2.11.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 130.1 KiB | [postgresql-18-plproxy_2.11.0-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-18-plproxy_2.11.0-13.pgdg13+1_arm64.deb) |
| `postgresql-18-plproxy` | `2.11.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 138.4 KiB | [postgresql-18-plproxy_2.11.0-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-18-plproxy_2.11.0-13.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plproxy` | `2.11.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 133.9 KiB | [postgresql-18-plproxy_2.11.0-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-18-plproxy_2.11.0-13.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plproxy` | `2.11.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 132.0 KiB | [postgresql-18-plproxy_2.11.0-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-18-plproxy_2.11.0-13.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plproxy` | `2.11.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 128.6 KiB | [postgresql-18-plproxy_2.11.0-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-18-plproxy_2.11.0-13.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plproxy_17` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 44.4 KiB | [plproxy_17-2.11.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plproxy_17-2.11.0-1PIGSTY.el8.x86_64.rpm) |
| `plproxy_17` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 48.1 KiB | [plproxy_17-2.11.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plproxy_17-2.11.0-2PGDG.rhel8.x86_64.rpm) |
| `plproxy_17` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.0 KiB | [plproxy_17-2.11.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plproxy_17-2.11.0-1PIGSTY.el8.aarch64.rpm) |
| `plproxy_17` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.6 KiB | [plproxy_17-2.11.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plproxy_17-2.11.0-2PGDG.rhel8.aarch64.rpm) |
| `plproxy_17` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.9 KiB | [plproxy_17-2.11.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plproxy_17-2.11.0-1PIGSTY.el9.x86_64.rpm) |
| `plproxy_17` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.7 KiB | [plproxy_17-2.11.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plproxy_17-2.11.0-2PGDG.rhel9.x86_64.rpm) |
| `plproxy_17` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.6 KiB | [plproxy_17-2.11.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plproxy_17-2.11.0-1PIGSTY.el9.aarch64.rpm) |
| `plproxy_17` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.5 KiB | [plproxy_17-2.11.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plproxy_17-2.11.0-2PGDG.rhel9.aarch64.rpm) |
| `plproxy_17` | `2.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.7 KiB | [plproxy_17-2.11.0-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plproxy_17-2.11.0-4PGDG.rhel10.x86_64.rpm) |
| `plproxy_17` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.9 KiB | [plproxy_17-2.11.0-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plproxy_17-2.11.0-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-plproxy` | `2.11.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 133.7 KiB | [postgresql-17-plproxy_2.11.0-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-17-plproxy_2.11.0-13.pgdg12+1_amd64.deb) |
| `postgresql-17-plproxy` | `2.11.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 130.0 KiB | [postgresql-17-plproxy_2.11.0-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-17-plproxy_2.11.0-13.pgdg12+1_arm64.deb) |
| `postgresql-17-plproxy` | `2.11.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 133.7 KiB | [postgresql-17-plproxy_2.11.0-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-17-plproxy_2.11.0-13.pgdg13+1_amd64.deb) |
| `postgresql-17-plproxy` | `2.11.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 130.1 KiB | [postgresql-17-plproxy_2.11.0-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-17-plproxy_2.11.0-13.pgdg13+1_arm64.deb) |
| `postgresql-17-plproxy` | `2.11.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.8 KiB | [postgresql-17-plproxy_2.11.0-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-17-plproxy_2.11.0-13.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plproxy` | `2.11.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 147.9 KiB | [postgresql-17-plproxy_2.11.0-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-17-plproxy_2.11.0-13.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plproxy` | `2.11.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 131.9 KiB | [postgresql-17-plproxy_2.11.0-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-17-plproxy_2.11.0-13.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plproxy` | `2.11.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 128.4 KiB | [postgresql-17-plproxy_2.11.0-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-17-plproxy_2.11.0-13.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plproxy_16` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 44.5 KiB | [plproxy_16-2.11.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plproxy_16-2.11.0-1PIGSTY.el8.x86_64.rpm) |
| `plproxy_16` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.0 KiB | [plproxy_16-2.11.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plproxy_16-2.11.0-1PIGSTY.el8.aarch64.rpm) |
| `plproxy_16` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.9 KiB | [plproxy_16-2.11.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plproxy_16-2.11.0-1PIGSTY.el9.x86_64.rpm) |
| `plproxy_16` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.6 KiB | [plproxy_16-2.11.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plproxy_16-2.11.0-1PIGSTY.el9.aarch64.rpm) |
| `plproxy_16` | `2.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 46.8 KiB | [plproxy_16-2.11.0-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plproxy_16-2.11.0-4PGDG.rhel10.x86_64.rpm) |
| `plproxy_16` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.9 KiB | [plproxy_16-2.11.0-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plproxy_16-2.11.0-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-plproxy` | `2.11.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 133.7 KiB | [postgresql-16-plproxy_2.11.0-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-16-plproxy_2.11.0-13.pgdg12+1_amd64.deb) |
| `postgresql-16-plproxy` | `2.11.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 129.9 KiB | [postgresql-16-plproxy_2.11.0-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-16-plproxy_2.11.0-13.pgdg12+1_arm64.deb) |
| `postgresql-16-plproxy` | `2.11.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 133.7 KiB | [postgresql-16-plproxy_2.11.0-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-16-plproxy_2.11.0-13.pgdg13+1_amd64.deb) |
| `postgresql-16-plproxy` | `2.11.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 130.1 KiB | [postgresql-16-plproxy_2.11.0-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-16-plproxy_2.11.0-13.pgdg13+1_arm64.deb) |
| `postgresql-16-plproxy` | `2.11.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.8 KiB | [postgresql-16-plproxy_2.11.0-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-16-plproxy_2.11.0-13.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plproxy` | `2.11.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 147.6 KiB | [postgresql-16-plproxy_2.11.0-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-16-plproxy_2.11.0-13.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plproxy` | `2.11.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 131.9 KiB | [postgresql-16-plproxy_2.11.0-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-16-plproxy_2.11.0-13.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plproxy` | `2.11.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 128.4 KiB | [postgresql-16-plproxy_2.11.0-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-16-plproxy_2.11.0-13.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plproxy_15` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 45.7 KiB | [plproxy_15-2.11.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plproxy_15-2.11.0-1PIGSTY.el8.x86_64.rpm) |
| `plproxy_15` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.2 KiB | [plproxy_15-2.11.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plproxy_15-2.11.0-1PGDG.rhel8.x86_64.rpm) |
| `plproxy_15` | `2.10.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 145.3 KiB | [plproxy_15-2.10.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plproxy_15-2.10.0-3.rhel8.x86_64.rpm) |
| `plproxy_15` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.0 KiB | [plproxy_15-2.11.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plproxy_15-2.11.0-1PIGSTY.el8.aarch64.rpm) |
| `plproxy_15` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.6 KiB | [plproxy_15-2.11.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plproxy_15-2.11.0-1PGDG.rhel8.aarch64.rpm) |
| `plproxy_15` | `2.10.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 142.2 KiB | [plproxy_15-2.10.0-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plproxy_15-2.10.0-3.rhel8.aarch64.rpm) |
| `plproxy_15` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 47.3 KiB | [plproxy_15-2.11.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plproxy_15-2.11.0-1PIGSTY.el9.x86_64.rpm) |
| `plproxy_15` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.0 KiB | [plproxy_15-2.11.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plproxy_15-2.11.0-1PGDG.rhel9.x86_64.rpm) |
| `plproxy_15` | `2.10.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 146.6 KiB | [plproxy_15-2.10.0-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plproxy_15-2.10.0-3.rhel9.x86_64.rpm) |
| `plproxy_15` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 45.4 KiB | [plproxy_15-2.11.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plproxy_15-2.11.0-1PIGSTY.el9.aarch64.rpm) |
| `plproxy_15` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.8 KiB | [plproxy_15-2.11.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plproxy_15-2.11.0-1PGDG.rhel9.aarch64.rpm) |
| `plproxy_15` | `2.10.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 144.3 KiB | [plproxy_15-2.10.0-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plproxy_15-2.10.0-3.rhel9.aarch64.rpm) |
| `plproxy_15` | `2.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 49.9 KiB | [plproxy_15-2.11.0-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plproxy_15-2.11.0-4PGDG.rhel10.x86_64.rpm) |
| `plproxy_15` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 48.1 KiB | [plproxy_15-2.11.0-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plproxy_15-2.11.0-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-plproxy` | `2.11.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 134.9 KiB | [postgresql-15-plproxy_2.11.0-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-15-plproxy_2.11.0-13.pgdg12+1_amd64.deb) |
| `postgresql-15-plproxy` | `2.11.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 131.3 KiB | [postgresql-15-plproxy_2.11.0-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-15-plproxy_2.11.0-13.pgdg12+1_arm64.deb) |
| `postgresql-15-plproxy` | `2.11.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 135.1 KiB | [postgresql-15-plproxy_2.11.0-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-15-plproxy_2.11.0-13.pgdg13+1_amd64.deb) |
| `postgresql-15-plproxy` | `2.11.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 131.3 KiB | [postgresql-15-plproxy_2.11.0-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-15-plproxy_2.11.0-13.pgdg13+1_arm64.deb) |
| `postgresql-15-plproxy` | `2.11.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 154.1 KiB | [postgresql-15-plproxy_2.11.0-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-15-plproxy_2.11.0-13.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plproxy` | `2.11.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 150.4 KiB | [postgresql-15-plproxy_2.11.0-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-15-plproxy_2.11.0-13.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plproxy` | `2.11.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 134.5 KiB | [postgresql-15-plproxy_2.11.0-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-15-plproxy_2.11.0-13.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plproxy` | `2.11.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 131.1 KiB | [postgresql-15-plproxy_2.11.0-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-15-plproxy_2.11.0-13.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plproxy_14` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 45.7 KiB | [plproxy_14-2.11.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plproxy_14-2.11.0-1PIGSTY.el8.x86_64.rpm) |
| `plproxy_14` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.2 KiB | [plproxy_14-2.11.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plproxy_14-2.11.0-1PGDG.rhel8.x86_64.rpm) |
| `plproxy_14` | `2.10.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 143.8 KiB | [plproxy_14-2.10.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plproxy_14-2.10.0-3.rhel8.x86_64.rpm) |
| `plproxy_14` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.0 KiB | [plproxy_14-2.11.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plproxy_14-2.11.0-1PIGSTY.el8.aarch64.rpm) |
| `plproxy_14` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.6 KiB | [plproxy_14-2.11.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plproxy_14-2.11.0-1PGDG.rhel8.aarch64.rpm) |
| `plproxy_14` | `2.10.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 140.5 KiB | [plproxy_14-2.10.0-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plproxy_14-2.10.0-3.rhel8.aarch64.rpm) |
| `plproxy_14` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 47.1 KiB | [plproxy_14-2.11.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plproxy_14-2.11.0-1PIGSTY.el9.x86_64.rpm) |
| `plproxy_14` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.7 KiB | [plproxy_14-2.11.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plproxy_14-2.11.0-1PGDG.rhel9.x86_64.rpm) |
| `plproxy_14` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 45.3 KiB | [plproxy_14-2.11.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plproxy_14-2.11.0-1PIGSTY.el9.aarch64.rpm) |
| `plproxy_14` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.8 KiB | [plproxy_14-2.11.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plproxy_14-2.11.0-1PGDG.rhel9.aarch64.rpm) |
| `plproxy_14` | `2.10.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.6 KiB | [plproxy_14-2.10.0-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plproxy_14-2.10.0-3.rhel9.aarch64.rpm) |
| `plproxy_14` | `2.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 49.8 KiB | [plproxy_14-2.11.0-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plproxy_14-2.11.0-4PGDG.rhel10.x86_64.rpm) |
| `plproxy_14` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 48.0 KiB | [plproxy_14-2.11.0-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plproxy_14-2.11.0-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-plproxy` | `2.11.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 134.6 KiB | [postgresql-14-plproxy_2.11.0-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-14-plproxy_2.11.0-13.pgdg12+1_amd64.deb) |
| `postgresql-14-plproxy` | `2.11.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 130.9 KiB | [postgresql-14-plproxy_2.11.0-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-14-plproxy_2.11.0-13.pgdg12+1_arm64.deb) |
| `postgresql-14-plproxy` | `2.11.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 134.8 KiB | [postgresql-14-plproxy_2.11.0-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-14-plproxy_2.11.0-13.pgdg13+1_amd64.deb) |
| `postgresql-14-plproxy` | `2.11.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 131.3 KiB | [postgresql-14-plproxy_2.11.0-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-14-plproxy_2.11.0-13.pgdg13+1_arm64.deb) |
| `postgresql-14-plproxy` | `2.11.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 152.5 KiB | [postgresql-14-plproxy_2.11.0-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-14-plproxy_2.11.0-13.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plproxy` | `2.11.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.8 KiB | [postgresql-14-plproxy_2.11.0-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-14-plproxy_2.11.0-13.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plproxy` | `2.11.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 134.4 KiB | [postgresql-14-plproxy_2.11.0-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-14-plproxy_2.11.0-13.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plproxy` | `2.11.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 131.1 KiB | [postgresql-14-plproxy_2.11.0-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-14-plproxy_2.11.0-13.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plproxy_13` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 45.3 KiB | [plproxy_13-2.11.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plproxy_13-2.11.0-1PIGSTY.el8.x86_64.rpm) |
| `plproxy_13` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 48.9 KiB | [plproxy_13-2.11.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/plproxy_13-2.11.0-1PGDG.rhel8.x86_64.rpm) |
| `plproxy_13` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 43.0 KiB | [plproxy_13-2.11.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plproxy_13-2.11.0-1PIGSTY.el8.aarch64.rpm) |
| `plproxy_13` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.5 KiB | [plproxy_13-2.11.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/plproxy_13-2.11.0-1PGDG.rhel8.aarch64.rpm) |
| `plproxy_13` | `2.10.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 140.1 KiB | [plproxy_13-2.10.0-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/plproxy_13-2.10.0-3.rhel8.aarch64.rpm) |
| `plproxy_13` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 47.2 KiB | [plproxy_13-2.11.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plproxy_13-2.11.0-1PIGSTY.el9.x86_64.rpm) |
| `plproxy_13` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.9 KiB | [plproxy_13-2.11.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/plproxy_13-2.11.0-1PGDG.rhel9.x86_64.rpm) |
| `plproxy_13` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 45.3 KiB | [plproxy_13-2.11.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plproxy_13-2.11.0-1PIGSTY.el9.aarch64.rpm) |
| `plproxy_13` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.8 KiB | [plproxy_13-2.11.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/plproxy_13-2.11.0-1PGDG.rhel9.aarch64.rpm) |
| `plproxy_13` | `2.10.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.3 KiB | [plproxy_13-2.10.0-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/plproxy_13-2.10.0-3.rhel9.aarch64.rpm) |
| `plproxy_13` | `2.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 49.9 KiB | [plproxy_13-2.11.0-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/plproxy_13-2.11.0-4PGDG.rhel10.x86_64.rpm) |
| `plproxy_13` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 48.1 KiB | [plproxy_13-2.11.0-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/plproxy_13-2.11.0-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-plproxy` | `2.11.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 134.5 KiB | [postgresql-13-plproxy_2.11.0-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-13-plproxy_2.11.0-13.pgdg12+1_amd64.deb) |
| `postgresql-13-plproxy` | `2.11.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 130.8 KiB | [postgresql-13-plproxy_2.11.0-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-13-plproxy_2.11.0-13.pgdg12+1_arm64.deb) |
| `postgresql-13-plproxy` | `2.11.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 134.4 KiB | [postgresql-13-plproxy_2.11.0-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-13-plproxy_2.11.0-13.pgdg13+1_amd64.deb) |
| `postgresql-13-plproxy` | `2.11.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 131.0 KiB | [postgresql-13-plproxy_2.11.0-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-13-plproxy_2.11.0-13.pgdg13+1_arm64.deb) |
| `postgresql-13-plproxy` | `2.11.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 152.0 KiB | [postgresql-13-plproxy_2.11.0-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-13-plproxy_2.11.0-13.pgdg22.04+1_amd64.deb) |
| `postgresql-13-plproxy` | `2.11.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.3 KiB | [postgresql-13-plproxy_2.11.0-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-13-plproxy_2.11.0-13.pgdg22.04+1_arm64.deb) |
| `postgresql-13-plproxy` | `2.11.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 134.3 KiB | [postgresql-13-plproxy_2.11.0-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-13-plproxy_2.11.0-13.pgdg24.04+1_amd64.deb) |
| `postgresql-13-plproxy` | `2.11.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 130.9 KiB | [postgresql-13-plproxy_2.11.0-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-plproxy/postgresql-13-plproxy_2.11.0-13.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/plproxy/plproxy" title="Repository" icon="github" subtitle="github.com/plproxy/plproxy" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plproxy-2.11.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get plproxy; # get plproxy source code
pig build dep plproxy; # install build dependencies
pig build pkg plproxy; # build extension rpm or deb
pig build ext plproxy; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install plproxy; # install by extension name, for the current active PG version
pig ext install plproxy; # install via package alias, for the active PG version
pig ext install plproxy -v 18;   # install for PG 18
pig ext install plproxy -v 17;   # install for PG 17
pig ext install plproxy -v 16;   # install for PG 16
pig ext install plproxy -v 15;   # install for PG 15
pig ext install plproxy -v 14;   # install for PG 14
pig ext install plproxy -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION plproxy;
```

