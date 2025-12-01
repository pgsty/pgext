---
title: "pgagent"
linkTitle: "pgagent"
description: "A PostgreSQL job scheduler"
weight: 5880
categories: ["ADMIN"]
width: full
---

[**pgagent**](https://www.pgadmin.org/docs/pgadmin4/development/pgagent.html) : A PostgreSQL job scheduler


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5880** | {{< badge content="pgagent" link="https://www.pgadmin.org/docs/pgadmin4/development/pgagent.html" >}} | {{< ext "pgagent" >}} | `4.2.3` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_jobmon" >}} {{< ext "pg_partman" >}} {{< ext "pglogical" >}} {{< ext "pg_background" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pgagent` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.3` | {{< bg "18" "pgagent_18*" "green" >}} {{< bg "17" "pgagent_17*" "green" >}} {{< bg "16" "pgagent_16*" "green" >}} {{< bg "15" "pgagent_15*" "green" >}} {{< bg "14" "pgagent_14*" "green" >}} {{< bg "13" "pgagent_13*" "green" >}} | `pgagent_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.3` | {{< bg "18" "pgagent" "green" >}} {{< bg "17" "pgagent" "green" >}} {{< bg "16" "pgagent" "green" >}} {{< bg "15" "pgagent" "green" >}} {{< bg "14" "pgagent" "green" >}} {{< bg "13" "pgagent" "green" >}} | `pgagent` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.2.3" "pgagent_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_13 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.2.3" "pgagent_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.2.3" "pgagent_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.2.3" "pgagent_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.2.3" "pgagent_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.2.3" "pgagent_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.3" "pgagent_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |      {{< bg "MISS" "pgagent : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgagent_18` | `4.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 135.5 KiB | [pgagent_18-4.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgagent_18-4.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgagent_18` | `4.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 129.6 KiB | [pgagent_18-4.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgagent_18-4.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgagent_18` | `4.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 121.5 KiB | [pgagent_18-4.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgagent_18-4.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgagent_18` | `4.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 117.8 KiB | [pgagent_18-4.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgagent_18-4.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgagent_18` | `4.2.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 126.9 KiB | [pgagent_18-4.2.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgagent_18-4.2.3-1PGDG.rhel10.x86_64.rpm) |
| `pgagent_18` | `4.2.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 116.6 KiB | [pgagent_18-4.2.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgagent_18-4.2.3-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgagent_17` | `4.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 135.5 KiB | [pgagent_17-4.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgagent_17-4.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgagent_17` | `4.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 133.9 KiB | [pgagent_17-4.2.2-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgagent_17-4.2.2-5PGDG.rhel8.x86_64.rpm) |
| `pgagent_17` | `4.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 129.6 KiB | [pgagent_17-4.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgagent_17-4.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgagent_17` | `4.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 128.2 KiB | [pgagent_17-4.2.2-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgagent_17-4.2.2-5PGDG.rhel8.aarch64.rpm) |
| `pgagent_17` | `4.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 121.9 KiB | [pgagent_17-4.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgagent_17-4.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgagent_17` | `4.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 120.3 KiB | [pgagent_17-4.2.2-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgagent_17-4.2.2-5PGDG.rhel9.x86_64.rpm) |
| `pgagent_17` | `4.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 118.0 KiB | [pgagent_17-4.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgagent_17-4.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgagent_17` | `4.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 116.7 KiB | [pgagent_17-4.2.2-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgagent_17-4.2.2-5PGDG.rhel9.aarch64.rpm) |
| `pgagent_17` | `4.2.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 126.9 KiB | [pgagent_17-4.2.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgagent_17-4.2.3-1PGDG.rhel10.x86_64.rpm) |
| `pgagent_17` | `4.2.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 116.6 KiB | [pgagent_17-4.2.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgagent_17-4.2.3-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgagent_16` | `4.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 135.5 KiB | [pgagent_16-4.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgagent_16-4.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgagent_16` | `4.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 133.5 KiB | [pgagent_16-4.2.2-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgagent_16-4.2.2-3.rhel8.x86_64.rpm) |
| `pgagent_16` | `4.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 129.6 KiB | [pgagent_16-4.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgagent_16-4.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgagent_16` | `4.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 127.9 KiB | [pgagent_16-4.2.2-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgagent_16-4.2.2-3.rhel8.aarch64.rpm) |
| `pgagent_16` | `4.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 122.0 KiB | [pgagent_16-4.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgagent_16-4.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgagent_16` | `4.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 120.0 KiB | [pgagent_16-4.2.2-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgagent_16-4.2.2-3.rhel9.x86_64.rpm) |
| `pgagent_16` | `4.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 118.0 KiB | [pgagent_16-4.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgagent_16-4.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgagent_16` | `4.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 116.4 KiB | [pgagent_16-4.2.2-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgagent_16-4.2.2-3.rhel9.aarch64.rpm) |
| `pgagent_16` | `4.2.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 126.9 KiB | [pgagent_16-4.2.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgagent_16-4.2.3-1PGDG.rhel10.x86_64.rpm) |
| `pgagent_16` | `4.2.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 116.6 KiB | [pgagent_16-4.2.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgagent_16-4.2.3-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgagent_15` | `4.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 135.5 KiB | [pgagent_15-4.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgagent_15-4.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgagent_15` | `4.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 133.4 KiB | [pgagent_15-4.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgagent_15-4.2.2-1.rhel8.x86_64.rpm) |
| `pgagent_15` | `4.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 129.6 KiB | [pgagent_15-4.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgagent_15-4.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgagent_15` | `4.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 127.7 KiB | [pgagent_15-4.2.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgagent_15-4.2.2-1.rhel8.aarch64.rpm) |
| `pgagent_15` | `4.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 122.0 KiB | [pgagent_15-4.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgagent_15-4.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgagent_15` | `4.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 119.1 KiB | [pgagent_15-4.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgagent_15-4.2.2-1.rhel9.x86_64.rpm) |
| `pgagent_15` | `4.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 118.1 KiB | [pgagent_15-4.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgagent_15-4.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgagent_15` | `4.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 114.4 KiB | [pgagent_15-4.2.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgagent_15-4.2.2-1.rhel9.aarch64.rpm) |
| `pgagent_15` | `4.2.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 126.9 KiB | [pgagent_15-4.2.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgagent_15-4.2.3-1PGDG.rhel10.x86_64.rpm) |
| `pgagent_15` | `4.2.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 116.6 KiB | [pgagent_15-4.2.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgagent_15-4.2.3-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgagent_14` | `4.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 135.5 KiB | [pgagent_14-4.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgagent_14-4.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgagent_14` | `4.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 133.4 KiB | [pgagent_14-4.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgagent_14-4.2.2-1.rhel8.x86_64.rpm) |
| `pgagent_14` | `4.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 153.1 KiB | [pgagent_14-4.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgagent_14-4.2.1-1.rhel8.x86_64.rpm) |
| `pgagent_14` | `4.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 129.6 KiB | [pgagent_14-4.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgagent_14-4.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgagent_14` | `4.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 127.7 KiB | [pgagent_14-4.2.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgagent_14-4.2.2-1.rhel8.aarch64.rpm) |
| `pgagent_14` | `4.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 122.0 KiB | [pgagent_14-4.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgagent_14-4.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgagent_14` | `4.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 119.1 KiB | [pgagent_14-4.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgagent_14-4.2.2-1.rhel9.x86_64.rpm) |
| `pgagent_14` | `4.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 138.6 KiB | [pgagent_14-4.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgagent_14-4.2.1-1.rhel9.x86_64.rpm) |
| `pgagent_14` | `4.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 118.0 KiB | [pgagent_14-4.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgagent_14-4.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgagent_14` | `4.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 114.4 KiB | [pgagent_14-4.2.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgagent_14-4.2.2-1.rhel9.aarch64.rpm) |
| `pgagent_14` | `4.2.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 126.9 KiB | [pgagent_14-4.2.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgagent_14-4.2.3-1PGDG.rhel10.x86_64.rpm) |
| `pgagent_14` | `4.2.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 116.7 KiB | [pgagent_14-4.2.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgagent_14-4.2.3-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgagent_13` | `4.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 135.5 KiB | [pgagent_13-4.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgagent_13-4.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgagent_13` | `4.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 133.4 KiB | [pgagent_13-4.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgagent_13-4.2.2-1.rhel8.x86_64.rpm) |
| `pgagent_13` | `4.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 153.1 KiB | [pgagent_13-4.2.1-0.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgagent_13-4.2.1-0.rhel8.x86_64.rpm) |
| `pgagent_13` | `4.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.1 KiB | [pgagent_13-4.0.0-4.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgagent_13-4.0.0-4.rhel8.x86_64.rpm) |
| `pgagent_13` | `4.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 129.6 KiB | [pgagent_13-4.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgagent_13-4.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgagent_13` | `4.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 127.6 KiB | [pgagent_13-4.2.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgagent_13-4.2.2-1.rhel8.aarch64.rpm) |
| `pgagent_13` | `4.2.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 121.9 KiB | [pgagent_13-4.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgagent_13-4.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgagent_13` | `4.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 119.1 KiB | [pgagent_13-4.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgagent_13-4.2.2-1.rhel9.x86_64.rpm) |
| `pgagent_13` | `4.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 138.6 KiB | [pgagent_13-4.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgagent_13-4.2.1-1.rhel9.x86_64.rpm) |
| `pgagent_13` | `4.2.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 118.1 KiB | [pgagent_13-4.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgagent_13-4.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgagent_13` | `4.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 114.4 KiB | [pgagent_13-4.2.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgagent_13-4.2.2-1.rhel9.aarch64.rpm) |
| `pgagent_13` | `4.2.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 126.9 KiB | [pgagent_13-4.2.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgagent_13-4.2.3-1PGDG.rhel10.x86_64.rpm) |
| `pgagent_13` | `4.2.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 116.6 KiB | [pgagent_13-4.2.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgagent_13-4.2.3-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://www.pgadmin.org/docs/pgadmin4/development/pgagent.html" title="Repository" icon="link" subtitle="www.pgadmin.org/docs/pgadmin4/development/pgagent.html" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgagent;		# install via package name, for the active PG version

pig install pgagent -v 18;   # install for PG 18
pig install pgagent -v 17;   # install for PG 17
pig install pgagent -v 16;   # install for PG 16
pig install pgagent -v 15;   # install for PG 15
pig install pgagent -v 14;   # install for PG 14
pig install pgagent -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgagent;
```
