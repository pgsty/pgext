---
title: "count_distinct"
linkTitle: "count_distinct"
description: "An alternative to COUNT(DISTINCT …) aggregate, usable with HashAggregate"
weight: 4630
categories: ["FUNC"]
width: full
---

An alternative to COUNT(DISTINCT …) aggregate, usable with HashAggregate


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4630** | {{< badge content="count_distinct" link="https://github.com/tvondra/count_distinct" >}} | {{< ext "count_distinct" >}} | `3.0.2` | {{< category "FUNC" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "topn" >}} {{< ext "hll" >}} {{< ext "omnisketch" >}} {{< ext "ddsketch" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "first_last_agg" >}} {{< ext "aggs_for_arrays" >}} |

> [!Note] no pg14,13,12 on el8/9 pgdg


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/count_distinct" >}} | `3.0.2` | {{< bg "18" "count_distinct_18*" "green" >}} {{< bg "17" "count_distinct_17*" "green" >}} {{< bg "16" "count_distinct_16*" "green" >}} {{< bg "15" "count_distinct_15*" "green" >}} {{< bg "14" "count_distinct_14*" "green" >}} {{< bg "13" "count_distinct_13*" "green" >}} | `count_distinct_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/count_distinct" >}} | `3.0.2` | {{< bg "18" "postgresql-18-count-distinct" "green" >}} {{< bg "17" "postgresql-17-count-distinct" "green" >}} {{< bg "16" "postgresql-16-count-distinct" "green" >}} {{< bg "15" "postgresql-15-count-distinct" "green" >}} {{< bg "14" "postgresql-14-count-distinct" "green" >}} {{< bg "13" "postgresql-13-count-distinct" "green" >}} | `postgresql-$v-count-distinct` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_13 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_13 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_13 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.0.2" "count_distinct_13 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-count-distinct : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-count-distinct : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-count-distinct : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-count-distinct : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-count-distinct : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-count-distinct : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-count-distinct : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-count-distinct : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-count-distinct : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-count-distinct : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-count-distinct : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-count-distinct : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-count-distinct : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-count-distinct : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-count-distinct : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-count-distinct : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-count-distinct : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-count-distinct : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-count-distinct : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `count_distinct_18` | 3.0.2 | `el8.x86_64` | pigsty | 16.8 KiB | [count_distinct_18-3.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/count_distinct_18-3.0.2-1PIGSTY.el8.x86_64.rpm) |
| `count_distinct_18` | 3.0.2 | `el8.x86_64` | pgdg | 23.1 KiB | [count_distinct_18-3.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/count_distinct_18-3.0.2-1PGDG.rhel8.x86_64.rpm) |
| `count_distinct_18` | 3.0.2 | `el8.aarch64` | pigsty | 16.6 KiB | [count_distinct_18-3.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/count_distinct_18-3.0.2-1PIGSTY.el8.aarch64.rpm) |
| `count_distinct_18` | 3.0.2 | `el8.aarch64` | pgdg | 22.8 KiB | [count_distinct_18-3.0.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/count_distinct_18-3.0.2-1PGDG.rhel8.aarch64.rpm) |
| `count_distinct_18` | 3.0.2 | `el9.x86_64` | pigsty | 16.6 KiB | [count_distinct_18-3.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/count_distinct_18-3.0.2-1PIGSTY.el9.x86_64.rpm) |
| `count_distinct_18` | 3.0.2 | `el9.x86_64` | pgdg | 22.9 KiB | [count_distinct_18-3.0.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/count_distinct_18-3.0.2-1PGDG.rhel9.x86_64.rpm) |
| `count_distinct_18` | 3.0.2 | `el9.aarch64` | pigsty | 16.5 KiB | [count_distinct_18-3.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/count_distinct_18-3.0.2-1PIGSTY.el9.aarch64.rpm) |
| `count_distinct_18` | 3.0.2 | `el9.aarch64` | pgdg | 22.3 KiB | [count_distinct_18-3.0.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/count_distinct_18-3.0.2-1PGDG.rhel9.aarch64.rpm) |
| `count_distinct_18` | 3.0.2 | `el10.x86_64` | pigsty | 16.7 KiB | [count_distinct_18-3.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/count_distinct_18-3.0.2-1PIGSTY.el10.x86_64.rpm) |
| `count_distinct_18` | 3.0.2 | `el10.x86_64` | pgdg | 23.3 KiB | [count_distinct_18-3.0.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/count_distinct_18-3.0.2-1PGDG.rhel10.x86_64.rpm) |
| `count_distinct_18` | 3.0.2 | `el10.aarch64` | pigsty | 16.6 KiB | [count_distinct_18-3.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/count_distinct_18-3.0.2-1PIGSTY.el10.aarch64.rpm) |
| `count_distinct_18` | 3.0.2 | `el10.aarch64` | pgdg | 22.9 KiB | [count_distinct_18-3.0.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/count_distinct_18-3.0.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-count-distinct` | 3.0.2 | `d12.x86_64` | pigsty | 34.8 KiB | [postgresql-18-count-distinct_3.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/count-distinct/postgresql-18-count-distinct_3.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-count-distinct` | 3.0.2 | `d12.aarch64` | pigsty | 34.6 KiB | [postgresql-18-count-distinct_3.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/count-distinct/postgresql-18-count-distinct_3.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-count-distinct` | 3.0.2 | `u22.x86_64` | pigsty | 36.4 KiB | [postgresql-18-count-distinct_3.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/count-distinct/postgresql-18-count-distinct_3.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-count-distinct` | 3.0.2 | `u22.aarch64` | pigsty | 36.4 KiB | [postgresql-18-count-distinct_3.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/count-distinct/postgresql-18-count-distinct_3.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-count-distinct` | 3.0.2 | `u24.x86_64` | pigsty | 35.8 KiB | [postgresql-18-count-distinct_3.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/count-distinct/postgresql-18-count-distinct_3.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-count-distinct` | 3.0.2 | `u24.aarch64` | pigsty | 35.7 KiB | [postgresql-18-count-distinct_3.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/count-distinct/postgresql-18-count-distinct_3.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `count_distinct_17` | 3.0.2 | `el8.x86_64` | pigsty | 16.8 KiB | [count_distinct_17-3.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/count_distinct_17-3.0.2-1PIGSTY.el8.x86_64.rpm) |
| `count_distinct_17` | 3.0.1 | `el8.x86_64` | pgdg | 20.7 KiB | [count_distinct_17-3.0.1-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/count_distinct_17-3.0.1-6PGDG.rhel8.x86_64.rpm) |
| `count_distinct_17` | 3.0.2 | `el8.aarch64` | pigsty | 16.6 KiB | [count_distinct_17-3.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/count_distinct_17-3.0.2-1PIGSTY.el8.aarch64.rpm) |
| `count_distinct_17` | 3.0.1 | `el8.aarch64` | pgdg | 20.4 KiB | [count_distinct_17-3.0.1-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/count_distinct_17-3.0.1-6PGDG.rhel8.aarch64.rpm) |
| `count_distinct_17` | 3.0.2 | `el9.x86_64` | pigsty | 16.6 KiB | [count_distinct_17-3.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/count_distinct_17-3.0.2-1PIGSTY.el9.x86_64.rpm) |
| `count_distinct_17` | 3.0.1 | `el9.x86_64` | pgdg | 20.6 KiB | [count_distinct_17-3.0.1-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/count_distinct_17-3.0.1-6PGDG.rhel9.x86_64.rpm) |
| `count_distinct_17` | 3.0.2 | `el9.aarch64` | pigsty | 16.5 KiB | [count_distinct_17-3.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/count_distinct_17-3.0.2-1PIGSTY.el9.aarch64.rpm) |
| `count_distinct_17` | 3.0.1 | `el9.aarch64` | pgdg | 20.2 KiB | [count_distinct_17-3.0.1-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/count_distinct_17-3.0.1-6PGDG.rhel9.aarch64.rpm) |
| `count_distinct_17` | 3.0.2 | `el10.x86_64` | pigsty | 16.7 KiB | [count_distinct_17-3.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/count_distinct_17-3.0.2-1PIGSTY.el10.x86_64.rpm) |
| `count_distinct_17` | 3.0.2 | `el10.x86_64` | pgdg | 23.3 KiB | [count_distinct_17-3.0.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/count_distinct_17-3.0.2-1PGDG.rhel10.x86_64.rpm) |
| `count_distinct_17` | 3.0.2 | `el10.aarch64` | pigsty | 16.6 KiB | [count_distinct_17-3.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/count_distinct_17-3.0.2-1PIGSTY.el10.aarch64.rpm) |
| `count_distinct_17` | 3.0.2 | `el10.aarch64` | pgdg | 22.9 KiB | [count_distinct_17-3.0.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/count_distinct_17-3.0.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-count-distinct` | 3.0.2 | `d12.x86_64` | pigsty | 34.7 KiB | [postgresql-17-count-distinct_3.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/count-distinct/postgresql-17-count-distinct_3.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-count-distinct` | 3.0.2 | `d12.aarch64` | pigsty | 34.5 KiB | [postgresql-17-count-distinct_3.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/count-distinct/postgresql-17-count-distinct_3.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-count-distinct` | 3.0.2 | `u22.x86_64` | pigsty | 37.4 KiB | [postgresql-17-count-distinct_3.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/count-distinct/postgresql-17-count-distinct_3.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-count-distinct` | 3.0.2 | `u22.aarch64` | pigsty | 37.3 KiB | [postgresql-17-count-distinct_3.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/count-distinct/postgresql-17-count-distinct_3.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-count-distinct` | 3.0.2 | `u24.x86_64` | pigsty | 35.7 KiB | [postgresql-17-count-distinct_3.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/count-distinct/postgresql-17-count-distinct_3.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-count-distinct` | 3.0.2 | `u24.aarch64` | pigsty | 35.6 KiB | [postgresql-17-count-distinct_3.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/count-distinct/postgresql-17-count-distinct_3.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `count_distinct_16` | 3.0.2 | `el8.x86_64` | pigsty | 16.8 KiB | [count_distinct_16-3.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/count_distinct_16-3.0.2-1PIGSTY.el8.x86_64.rpm) |
| `count_distinct_16` | 3.0.1 | `el8.x86_64` | pgdg | 20.6 KiB | [count_distinct_16-3.0.1-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/count_distinct_16-3.0.1-5PGDG.rhel8.x86_64.rpm) |
| `count_distinct_16` | 3.0.2 | `el8.aarch64` | pigsty | 16.6 KiB | [count_distinct_16-3.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/count_distinct_16-3.0.2-1PIGSTY.el8.aarch64.rpm) |
| `count_distinct_16` | 3.0.1 | `el8.aarch64` | pgdg | 20.3 KiB | [count_distinct_16-3.0.1-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/count_distinct_16-3.0.1-5PGDG.rhel8.aarch64.rpm) |
| `count_distinct_16` | 3.0.2 | `el9.x86_64` | pigsty | 16.6 KiB | [count_distinct_16-3.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/count_distinct_16-3.0.2-1PIGSTY.el9.x86_64.rpm) |
| `count_distinct_16` | 3.0.1 | `el9.x86_64` | pgdg | 20.4 KiB | [count_distinct_16-3.0.1-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/count_distinct_16-3.0.1-5PGDG.rhel9.x86_64.rpm) |
| `count_distinct_16` | 3.0.2 | `el9.aarch64` | pigsty | 16.5 KiB | [count_distinct_16-3.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/count_distinct_16-3.0.2-1PIGSTY.el9.aarch64.rpm) |
| `count_distinct_16` | 3.0.1 | `el9.aarch64` | pgdg | 19.8 KiB | [count_distinct_16-3.0.1-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/count_distinct_16-3.0.1-5PGDG.rhel9.aarch64.rpm) |
| `count_distinct_16` | 3.0.2 | `el10.x86_64` | pigsty | 16.7 KiB | [count_distinct_16-3.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/count_distinct_16-3.0.2-1PIGSTY.el10.x86_64.rpm) |
| `count_distinct_16` | 3.0.2 | `el10.x86_64` | pgdg | 23.3 KiB | [count_distinct_16-3.0.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/count_distinct_16-3.0.2-1PGDG.rhel10.x86_64.rpm) |
| `count_distinct_16` | 3.0.2 | `el10.aarch64` | pigsty | 16.6 KiB | [count_distinct_16-3.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/count_distinct_16-3.0.2-1PIGSTY.el10.aarch64.rpm) |
| `count_distinct_16` | 3.0.2 | `el10.aarch64` | pgdg | 22.9 KiB | [count_distinct_16-3.0.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/count_distinct_16-3.0.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-count-distinct` | 3.0.2 | `d12.x86_64` | pigsty | 34.7 KiB | [postgresql-16-count-distinct_3.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/count-distinct/postgresql-16-count-distinct_3.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-count-distinct` | 3.0.2 | `d12.aarch64` | pigsty | 34.5 KiB | [postgresql-16-count-distinct_3.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/count-distinct/postgresql-16-count-distinct_3.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-count-distinct` | 3.0.2 | `u22.x86_64` | pigsty | 37.4 KiB | [postgresql-16-count-distinct_3.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/count-distinct/postgresql-16-count-distinct_3.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-count-distinct` | 3.0.2 | `u22.aarch64` | pigsty | 37.3 KiB | [postgresql-16-count-distinct_3.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/count-distinct/postgresql-16-count-distinct_3.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-count-distinct` | 3.0.2 | `u24.x86_64` | pigsty | 35.7 KiB | [postgresql-16-count-distinct_3.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/count-distinct/postgresql-16-count-distinct_3.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-count-distinct` | 3.0.2 | `u24.aarch64` | pigsty | 35.6 KiB | [postgresql-16-count-distinct_3.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/count-distinct/postgresql-16-count-distinct_3.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `count_distinct_15` | 3.0.2 | `el8.x86_64` | pigsty | 16.7 KiB | [count_distinct_15-3.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/count_distinct_15-3.0.2-1PIGSTY.el8.x86_64.rpm) |
| `count_distinct_15` | 3.0.1 | `el8.x86_64` | pgdg | 31.7 KiB | [count_distinct_15-3.0.1-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/count_distinct_15-3.0.1-3.rhel8.x86_64.rpm) |
| `count_distinct_15` | 3.0.2 | `el8.aarch64` | pigsty | 16.6 KiB | [count_distinct_15-3.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/count_distinct_15-3.0.2-1PIGSTY.el8.aarch64.rpm) |
| `count_distinct_15` | 3.0.1 | `el8.aarch64` | pgdg | 31.2 KiB | [count_distinct_15-3.0.1-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/count_distinct_15-3.0.1-3.rhel8.aarch64.rpm) |
| `count_distinct_15` | 3.0.2 | `el9.x86_64` | pigsty | 16.7 KiB | [count_distinct_15-3.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/count_distinct_15-3.0.2-1PIGSTY.el9.x86_64.rpm) |
| `count_distinct_15` | 3.0.1 | `el9.x86_64` | pgdg | 32.4 KiB | [count_distinct_15-3.0.1-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/count_distinct_15-3.0.1-3.rhel9.x86_64.rpm) |
| `count_distinct_15` | 3.0.2 | `el9.aarch64` | pigsty | 16.4 KiB | [count_distinct_15-3.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/count_distinct_15-3.0.2-1PIGSTY.el9.aarch64.rpm) |
| `count_distinct_15` | 3.0.1 | `el9.aarch64` | pgdg | 31.6 KiB | [count_distinct_15-3.0.1-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/count_distinct_15-3.0.1-3.rhel9.aarch64.rpm) |
| `count_distinct_15` | 3.0.2 | `el10.x86_64` | pigsty | 16.6 KiB | [count_distinct_15-3.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/count_distinct_15-3.0.2-1PIGSTY.el10.x86_64.rpm) |
| `count_distinct_15` | 3.0.2 | `el10.x86_64` | pgdg | 23.2 KiB | [count_distinct_15-3.0.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/count_distinct_15-3.0.2-1PGDG.rhel10.x86_64.rpm) |
| `count_distinct_15` | 3.0.2 | `el10.aarch64` | pigsty | 16.6 KiB | [count_distinct_15-3.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/count_distinct_15-3.0.2-1PIGSTY.el10.aarch64.rpm) |
| `count_distinct_15` | 3.0.2 | `el10.aarch64` | pgdg | 22.9 KiB | [count_distinct_15-3.0.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/count_distinct_15-3.0.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-count-distinct` | 3.0.2 | `d12.x86_64` | pigsty | 34.5 KiB | [postgresql-15-count-distinct_3.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/count-distinct/postgresql-15-count-distinct_3.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-count-distinct` | 3.0.2 | `d12.aarch64` | pigsty | 34.2 KiB | [postgresql-15-count-distinct_3.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/count-distinct/postgresql-15-count-distinct_3.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-count-distinct` | 3.0.2 | `u22.x86_64` | pigsty | 37.1 KiB | [postgresql-15-count-distinct_3.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/count-distinct/postgresql-15-count-distinct_3.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-count-distinct` | 3.0.2 | `u22.aarch64` | pigsty | 37.1 KiB | [postgresql-15-count-distinct_3.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/count-distinct/postgresql-15-count-distinct_3.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-count-distinct` | 3.0.2 | `u24.x86_64` | pigsty | 35.5 KiB | [postgresql-15-count-distinct_3.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/count-distinct/postgresql-15-count-distinct_3.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-count-distinct` | 3.0.2 | `u24.aarch64` | pigsty | 35.4 KiB | [postgresql-15-count-distinct_3.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/count-distinct/postgresql-15-count-distinct_3.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `count_distinct_14` | 3.0.2 | `el8.x86_64` | pigsty | 16.7 KiB | [count_distinct_14-3.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/count_distinct_14-3.0.2-1PIGSTY.el8.x86_64.rpm) |
| `count_distinct_14` | 3.0.1 | `el8.x86_64` | pgdg | 32.1 KiB | [count_distinct_14-3.0.1-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/count_distinct_14-3.0.1-3.rhel8.x86_64.rpm) |
| `count_distinct_14` | 3.0.2 | `el8.aarch64` | pigsty | 16.6 KiB | [count_distinct_14-3.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/count_distinct_14-3.0.2-1PIGSTY.el8.aarch64.rpm) |
| `count_distinct_14` | 3.0.1 | `el8.aarch64` | pgdg | 31.1 KiB | [count_distinct_14-3.0.1-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/count_distinct_14-3.0.1-3.rhel8.aarch64.rpm) |
| `count_distinct_14` | 3.0.2 | `el9.x86_64` | pigsty | 16.6 KiB | [count_distinct_14-3.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/count_distinct_14-3.0.2-1PIGSTY.el9.x86_64.rpm) |
| `count_distinct_14` | 3.0.2 | `el9.aarch64` | pigsty | 16.4 KiB | [count_distinct_14-3.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/count_distinct_14-3.0.2-1PIGSTY.el9.aarch64.rpm) |
| `count_distinct_14` | 3.0.1 | `el9.aarch64` | pgdg | 31.7 KiB | [count_distinct_14-3.0.1-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/count_distinct_14-3.0.1-3.rhel9.aarch64.rpm) |
| `count_distinct_14` | 3.0.2 | `el10.x86_64` | pigsty | 16.6 KiB | [count_distinct_14-3.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/count_distinct_14-3.0.2-1PIGSTY.el10.x86_64.rpm) |
| `count_distinct_14` | 3.0.2 | `el10.x86_64` | pgdg | 23.2 KiB | [count_distinct_14-3.0.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/count_distinct_14-3.0.2-1PGDG.rhel10.x86_64.rpm) |
| `count_distinct_14` | 3.0.2 | `el10.aarch64` | pigsty | 16.5 KiB | [count_distinct_14-3.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/count_distinct_14-3.0.2-1PIGSTY.el10.aarch64.rpm) |
| `count_distinct_14` | 3.0.2 | `el10.aarch64` | pgdg | 22.9 KiB | [count_distinct_14-3.0.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/count_distinct_14-3.0.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-count-distinct` | 3.0.2 | `d12.x86_64` | pigsty | 34.5 KiB | [postgresql-14-count-distinct_3.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/count-distinct/postgresql-14-count-distinct_3.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-count-distinct` | 3.0.2 | `d12.aarch64` | pigsty | 34.2 KiB | [postgresql-14-count-distinct_3.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/count-distinct/postgresql-14-count-distinct_3.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-count-distinct` | 3.0.2 | `u22.x86_64` | pigsty | 37.0 KiB | [postgresql-14-count-distinct_3.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/count-distinct/postgresql-14-count-distinct_3.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-count-distinct` | 3.0.2 | `u22.aarch64` | pigsty | 37.0 KiB | [postgresql-14-count-distinct_3.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/count-distinct/postgresql-14-count-distinct_3.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-count-distinct` | 3.0.2 | `u24.x86_64` | pigsty | 35.4 KiB | [postgresql-14-count-distinct_3.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/count-distinct/postgresql-14-count-distinct_3.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-count-distinct` | 3.0.2 | `u24.aarch64` | pigsty | 35.4 KiB | [postgresql-14-count-distinct_3.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/count-distinct/postgresql-14-count-distinct_3.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `count_distinct_13` | 3.0.2 | `el8.x86_64` | pigsty | 16.6 KiB | [count_distinct_13-3.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/count_distinct_13-3.0.2-1PIGSTY.el8.x86_64.rpm) |
| `count_distinct_13` | 3.0.2 | `el8.aarch64` | pigsty | 16.6 KiB | [count_distinct_13-3.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/count_distinct_13-3.0.2-1PIGSTY.el8.aarch64.rpm) |
| `count_distinct_13` | 3.0.1 | `el8.aarch64` | pgdg | 31.1 KiB | [count_distinct_13-3.0.1-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/count_distinct_13-3.0.1-3.rhel8.aarch64.rpm) |
| `count_distinct_13` | 3.0.2 | `el9.x86_64` | pigsty | 16.6 KiB | [count_distinct_13-3.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/count_distinct_13-3.0.2-1PIGSTY.el9.x86_64.rpm) |
| `count_distinct_13` | 3.0.2 | `el9.aarch64` | pigsty | 16.5 KiB | [count_distinct_13-3.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/count_distinct_13-3.0.2-1PIGSTY.el9.aarch64.rpm) |
| `count_distinct_13` | 3.0.1 | `el9.aarch64` | pgdg | 31.7 KiB | [count_distinct_13-3.0.1-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/count_distinct_13-3.0.1-3.rhel9.aarch64.rpm) |
| `count_distinct_13` | 3.0.2 | `el10.x86_64` | pigsty | 16.6 KiB | [count_distinct_13-3.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/count_distinct_13-3.0.2-1PIGSTY.el10.x86_64.rpm) |
| `count_distinct_13` | 3.0.2 | `el10.x86_64` | pgdg | 23.2 KiB | [count_distinct_13-3.0.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/count_distinct_13-3.0.2-1PGDG.rhel10.x86_64.rpm) |
| `count_distinct_13` | 3.0.2 | `el10.aarch64` | pigsty | 16.5 KiB | [count_distinct_13-3.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/count_distinct_13-3.0.2-1PIGSTY.el10.aarch64.rpm) |
| `count_distinct_13` | 3.0.2 | `el10.aarch64` | pgdg | 22.9 KiB | [count_distinct_13-3.0.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/count_distinct_13-3.0.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-count-distinct` | 3.0.2 | `d12.x86_64` | pigsty | 34.2 KiB | [postgresql-13-count-distinct_3.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/count-distinct/postgresql-13-count-distinct_3.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-count-distinct` | 3.0.2 | `d12.aarch64` | pigsty | 33.9 KiB | [postgresql-13-count-distinct_3.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/count-distinct/postgresql-13-count-distinct_3.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-count-distinct` | 3.0.2 | `u22.x86_64` | pigsty | 37.1 KiB | [postgresql-13-count-distinct_3.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/count-distinct/postgresql-13-count-distinct_3.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-count-distinct` | 3.0.2 | `u22.aarch64` | pigsty | 36.8 KiB | [postgresql-13-count-distinct_3.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/count-distinct/postgresql-13-count-distinct_3.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-count-distinct` | 3.0.2 | `u24.x86_64` | pigsty | 35.3 KiB | [postgresql-13-count-distinct_3.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/count-distinct/postgresql-13-count-distinct_3.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-count-distinct` | 3.0.2 | `u24.aarch64` | pigsty | 35.0 KiB | [postgresql-13-count-distinct_3.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/count-distinct/postgresql-13-count-distinct_3.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/count_distinct" title="Repository" icon="github" subtitle="github.com/tvondra/count_distinct" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="count_distinct-3.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get count_distinct; # get count_distinct source code
pig build dep count_distinct; # install build dependencies
pig build pkg count_distinct; # build extension rpm or deb
pig build ext count_distinct; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install count_distinct; # install by extension name, for the current active PG version
pig ext install count_distinct; # install via package alias, for the active PG version
pig ext install count_distinct -v 18;   # install for PG 18
pig ext install count_distinct -v 17;   # install for PG 17
pig ext install count_distinct -v 16;   # install for PG 16
pig ext install count_distinct -v 15;   # install for PG 15
pig ext install count_distinct -v 14;   # install for PG 14
pig ext install count_distinct -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION count_distinct;
```

