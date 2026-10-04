---
title: "tdigest"
linkTitle: "tdigest"
description: "Provides tdigest aggregate function."
weight: 4710
categories: ["FUNC"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**tdigest**](https://github.com/tvondra/tdigest) : Provides tdigest aggregate function.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4710** | {{< badge content="tdigest" link="https://github.com/tvondra/tdigest" >}} | {{< ext "tdigest" >}} | `1.4.7` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "ddsketch" >}} {{< ext "count_distinct" >}} {{< ext "topn" >}} {{< ext "omnisketch" >}} {{< ext "datasketches" >}} {{< ext "hll" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "weighted_statistics" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.4.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `tdigest` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.4.7` | {{< bg "18" "tdigest_18" "green" >}} {{< bg "17" "tdigest_17" "green" >}} {{< bg "16" "tdigest_16" "green" >}} {{< bg "15" "tdigest_15" "green" >}} {{< bg "14" "tdigest_14" "green" >}} | `tdigest_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.4.7` | {{< bg "18" "postgresql-18-tdigest" "green" >}} {{< bg "17" "postgresql-17-tdigest" "green" >}} {{< bg "16" "postgresql-16-tdigest" "green" >}} {{< bg "15" "postgresql-15-tdigest" "green" >}} {{< bg "14" "postgresql-14-tdigest" "green" >}} | `postgresql-$v-tdigest` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.4.7" "tdigest_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_14 : AVAIL 6" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.4.7" "tdigest_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_14 : AVAIL 6" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.4.7" "tdigest_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_14 : AVAIL 7" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.4.7" "tdigest_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_14 : AVAIL 7" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.4.7" "tdigest_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_14 : AVAIL 6" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.4.7" "tdigest_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.4.7" "tdigest_14 : AVAIL 6" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.4.7" "postgresql-18-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-17-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-16-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-15-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-14-tdigest : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.4.7" "postgresql-18-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-17-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-16-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-15-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-14-tdigest : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.4.7" "postgresql-18-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-17-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-16-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-15-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-14-tdigest : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.4.7" "postgresql-18-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-17-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-16-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-15-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-14-tdigest : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.4.7" "postgresql-18-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-17-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-16-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-15-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-14-tdigest : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.4.7" "postgresql-18-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-17-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-16-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-15-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-14-tdigest : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.4.7" "postgresql-18-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-17-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-16-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-15-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-14-tdigest : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.4.7" "postgresql-18-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-17-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-16-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-15-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-14-tdigest : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.4.7" "postgresql-18-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-17-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-16-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-15-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-14-tdigest : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.4.7" "postgresql-18-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-17-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-16-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-15-tdigest : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.7" "postgresql-14-tdigest : AVAIL 3" "blue" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tdigest_18` | `1.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.3 KiB | [tdigest_18-1.4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/tdigest_18-1.4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_18` | `1.4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.4 KiB | [tdigest_18-1.4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/tdigest_18-1.4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_18` | `1.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.1 KiB | [tdigest_18-1.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/tdigest_18-1.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_18` | `1.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.7 KiB | [tdigest_18-1.4.4-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/tdigest_18-1.4.4-4PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_18` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.5 KiB | [tdigest_18-1.4.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/tdigest_18-1.4.2-2PGDG.rhel8.x86_64.rpm) |
| `tdigest_18` | `1.4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.6 KiB | [tdigest_18-1.4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/tdigest_18-1.4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_18` | `1.4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.8 KiB | [tdigest_18-1.4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/tdigest_18-1.4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_18` | `1.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [tdigest_18-1.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/tdigest_18-1.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_18` | `1.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.4 KiB | [tdigest_18-1.4.4-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/tdigest_18-1.4.4-4PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_18` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.1 KiB | [tdigest_18-1.4.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/tdigest_18-1.4.2-2PGDG.rhel8.aarch64.rpm) |
| `tdigest_18` | `1.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.7 KiB | [tdigest_18-1.4.7-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/tdigest_18-1.4.7-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_18` | `1.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.8 KiB | [tdigest_18-1.4.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/tdigest_18-1.4.6-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_18` | `1.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.3 KiB | [tdigest_18-1.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/tdigest_18-1.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_18` | `1.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.5 KiB | [tdigest_18-1.4.4-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/tdigest_18-1.4.4-4PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_18` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.6 KiB | [tdigest_18-1.4.2-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/tdigest_18-1.4.2-4PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_18` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.6 KiB | [tdigest_18-1.4.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/tdigest_18-1.4.2-2PGDG.rhel9.x86_64.rpm) |
| `tdigest_18` | `1.4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.7 KiB | [tdigest_18-1.4.7-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/tdigest_18-1.4.7-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_18` | `1.4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [tdigest_18-1.4.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/tdigest_18-1.4.6-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_18` | `1.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.1 KiB | [tdigest_18-1.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/tdigest_18-1.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_18` | `1.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 33.6 KiB | [tdigest_18-1.4.4-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/tdigest_18-1.4.4-4PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_18` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.4 KiB | [tdigest_18-1.4.2-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/tdigest_18-1.4.2-4PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_18` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.3 KiB | [tdigest_18-1.4.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/tdigest_18-1.4.2-2PGDG.rhel9.aarch64.rpm) |
| `tdigest_18` | `1.4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.1 KiB | [tdigest_18-1.4.7-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/tdigest_18-1.4.7-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_18` | `1.4.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.2 KiB | [tdigest_18-1.4.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/tdigest_18-1.4.6-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_18` | `1.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.7 KiB | [tdigest_18-1.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/tdigest_18-1.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_18` | `1.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 34.6 KiB | [tdigest_18-1.4.4-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/tdigest_18-1.4.4-4PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_18` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.5 KiB | [tdigest_18-1.4.2-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/tdigest_18-1.4.2-4PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_18` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.8 KiB | [tdigest_18-1.4.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/tdigest_18-1.4.2-2PGDG.rhel10.x86_64.rpm) |
| `tdigest_18` | `1.4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.3 KiB | [tdigest_18-1.4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/tdigest_18-1.4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_18` | `1.4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [tdigest_18-1.4.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/tdigest_18-1.4.6-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_18` | `1.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 38.7 KiB | [tdigest_18-1.4.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/tdigest_18-1.4.5-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_18` | `1.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 34.3 KiB | [tdigest_18-1.4.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/tdigest_18-1.4.4-4PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_18` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.9 KiB | [tdigest_18-1.4.2-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/tdigest_18-1.4.2-4PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_18` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.2 KiB | [tdigest_18-1.4.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/tdigest_18-1.4.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-tdigest` | `1.4.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 77.7 KiB | [postgresql-18-tdigest_1.4.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.7-1.pgdg12+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 72.2 KiB | [postgresql-18-tdigest_1.4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 59.7 KiB | [postgresql-18-tdigest_1.4.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.4-1.pgdg12+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 76.4 KiB | [postgresql-18-tdigest_1.4.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.7-1.pgdg12+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 71.1 KiB | [postgresql-18-tdigest_1.4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 59.2 KiB | [postgresql-18-tdigest_1.4.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.4-1.pgdg12+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 77.8 KiB | [postgresql-18-tdigest_1.4.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.7-1.pgdg13+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 72.2 KiB | [postgresql-18-tdigest_1.4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 59.8 KiB | [postgresql-18-tdigest_1.4.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.4-1.pgdg13+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 77.0 KiB | [postgresql-18-tdigest_1.4.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.7-1.pgdg13+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 71.6 KiB | [postgresql-18-tdigest_1.4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 59.6 KiB | [postgresql-18-tdigest_1.4.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.4-1.pgdg13+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 76.7 KiB | [postgresql-18-tdigest_1.4.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 71.3 KiB | [postgresql-18-tdigest_1.4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 59.7 KiB | [postgresql-18-tdigest_1.4.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 75.3 KiB | [postgresql-18-tdigest_1.4.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 70.5 KiB | [postgresql-18-tdigest_1.4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 58.9 KiB | [postgresql-18-tdigest_1.4.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 76.6 KiB | [postgresql-18-tdigest_1.4.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 71.5 KiB | [postgresql-18-tdigest_1.4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 59.7 KiB | [postgresql-18-tdigest_1.4.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 75.5 KiB | [postgresql-18-tdigest_1.4.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 70.5 KiB | [postgresql-18-tdigest_1.4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 59.1 KiB | [postgresql-18-tdigest_1.4.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 75.9 KiB | [postgresql-18-tdigest_1.4.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 71.5 KiB | [postgresql-18-tdigest_1.4.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 59.2 KiB | [postgresql-18-tdigest_1.4.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-tdigest` | `1.4.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 75.2 KiB | [postgresql-18-tdigest_1.4.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 70.7 KiB | [postgresql-18-tdigest_1.4.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-tdigest` | `1.4.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 59.0 KiB | [postgresql-18-tdigest_1.4.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-18-tdigest_1.4.4-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tdigest_17` | `1.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.3 KiB | [tdigest_17-1.4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/tdigest_17-1.4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_17` | `1.4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.4 KiB | [tdigest_17-1.4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/tdigest_17-1.4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_17` | `1.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.1 KiB | [tdigest_17-1.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/tdigest_17-1.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_17` | `1.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.7 KiB | [tdigest_17-1.4.4-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/tdigest_17-1.4.4-4PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_17` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.5 KiB | [tdigest_17-1.4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/tdigest_17-1.4.2-1PGDG.rhel8.x86_64.rpm) |
| `tdigest_17` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.2 KiB | [tdigest_17-1.4.1-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/tdigest_17-1.4.1-3PGDG.rhel8.x86_64.rpm) |
| `tdigest_17` | `1.4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.7 KiB | [tdigest_17-1.4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/tdigest_17-1.4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_17` | `1.4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.8 KiB | [tdigest_17-1.4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/tdigest_17-1.4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_17` | `1.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [tdigest_17-1.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/tdigest_17-1.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_17` | `1.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.4 KiB | [tdigest_17-1.4.4-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/tdigest_17-1.4.4-4PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_17` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.1 KiB | [tdigest_17-1.4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/tdigest_17-1.4.2-1PGDG.rhel8.aarch64.rpm) |
| `tdigest_17` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.8 KiB | [tdigest_17-1.4.1-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/tdigest_17-1.4.1-3PGDG.rhel8.aarch64.rpm) |
| `tdigest_17` | `1.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.7 KiB | [tdigest_17-1.4.7-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/tdigest_17-1.4.7-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_17` | `1.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.8 KiB | [tdigest_17-1.4.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/tdigest_17-1.4.6-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_17` | `1.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.3 KiB | [tdigest_17-1.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/tdigest_17-1.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_17` | `1.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.5 KiB | [tdigest_17-1.4.4-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/tdigest_17-1.4.4-4PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_17` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.6 KiB | [tdigest_17-1.4.2-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/tdigest_17-1.4.2-4PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_17` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.6 KiB | [tdigest_17-1.4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/tdigest_17-1.4.2-1PGDG.rhel9.x86_64.rpm) |
| `tdigest_17` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.3 KiB | [tdigest_17-1.4.1-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/tdigest_17-1.4.1-3PGDG.rhel9.x86_64.rpm) |
| `tdigest_17` | `1.4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.7 KiB | [tdigest_17-1.4.7-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/tdigest_17-1.4.7-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_17` | `1.4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [tdigest_17-1.4.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/tdigest_17-1.4.6-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_17` | `1.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.1 KiB | [tdigest_17-1.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/tdigest_17-1.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_17` | `1.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 33.6 KiB | [tdigest_17-1.4.4-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/tdigest_17-1.4.4-4PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_17` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.4 KiB | [tdigest_17-1.4.2-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/tdigest_17-1.4.2-4PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_17` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.4 KiB | [tdigest_17-1.4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/tdigest_17-1.4.2-1PGDG.rhel9.aarch64.rpm) |
| `tdigest_17` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.1 KiB | [tdigest_17-1.4.1-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/tdigest_17-1.4.1-3PGDG.rhel9.aarch64.rpm) |
| `tdigest_17` | `1.4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.1 KiB | [tdigest_17-1.4.7-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/tdigest_17-1.4.7-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_17` | `1.4.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.2 KiB | [tdigest_17-1.4.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/tdigest_17-1.4.6-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_17` | `1.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.7 KiB | [tdigest_17-1.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/tdigest_17-1.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_17` | `1.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 34.6 KiB | [tdigest_17-1.4.4-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/tdigest_17-1.4.4-4PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_17` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.5 KiB | [tdigest_17-1.4.2-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/tdigest_17-1.4.2-4PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_17` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.8 KiB | [tdigest_17-1.4.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/tdigest_17-1.4.2-2PGDG.rhel10.x86_64.rpm) |
| `tdigest_17` | `1.4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.3 KiB | [tdigest_17-1.4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/tdigest_17-1.4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_17` | `1.4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [tdigest_17-1.4.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/tdigest_17-1.4.6-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_17` | `1.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 38.7 KiB | [tdigest_17-1.4.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/tdigest_17-1.4.5-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_17` | `1.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 34.3 KiB | [tdigest_17-1.4.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/tdigest_17-1.4.4-4PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_17` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.9 KiB | [tdigest_17-1.4.2-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/tdigest_17-1.4.2-4PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_17` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.2 KiB | [tdigest_17-1.4.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/tdigest_17-1.4.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-tdigest` | `1.4.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 77.3 KiB | [postgresql-17-tdigest_1.4.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.7-1.pgdg12+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 72.2 KiB | [postgresql-17-tdigest_1.4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 59.8 KiB | [postgresql-17-tdigest_1.4.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.4-1.pgdg12+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 76.1 KiB | [postgresql-17-tdigest_1.4.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.7-1.pgdg12+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 71.1 KiB | [postgresql-17-tdigest_1.4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 59.3 KiB | [postgresql-17-tdigest_1.4.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.4-1.pgdg12+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 77.4 KiB | [postgresql-17-tdigest_1.4.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.7-1.pgdg13+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 72.2 KiB | [postgresql-17-tdigest_1.4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 59.8 KiB | [postgresql-17-tdigest_1.4.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.4-1.pgdg13+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 76.7 KiB | [postgresql-17-tdigest_1.4.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.7-1.pgdg13+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 71.6 KiB | [postgresql-17-tdigest_1.4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 59.7 KiB | [postgresql-17-tdigest_1.4.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.4-1.pgdg13+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 80.9 KiB | [postgresql-17-tdigest_1.4.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 74.7 KiB | [postgresql-17-tdigest_1.4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 62.5 KiB | [postgresql-17-tdigest_1.4.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 79.2 KiB | [postgresql-17-tdigest_1.4.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 73.7 KiB | [postgresql-17-tdigest_1.4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 61.5 KiB | [postgresql-17-tdigest_1.4.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 76.2 KiB | [postgresql-17-tdigest_1.4.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 71.5 KiB | [postgresql-17-tdigest_1.4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 59.8 KiB | [postgresql-17-tdigest_1.4.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 75.2 KiB | [postgresql-17-tdigest_1.4.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 70.5 KiB | [postgresql-17-tdigest_1.4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 59.2 KiB | [postgresql-17-tdigest_1.4.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 75.6 KiB | [postgresql-17-tdigest_1.4.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 71.5 KiB | [postgresql-17-tdigest_1.4.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 59.1 KiB | [postgresql-17-tdigest_1.4.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-tdigest` | `1.4.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 74.8 KiB | [postgresql-17-tdigest_1.4.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 70.6 KiB | [postgresql-17-tdigest_1.4.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-tdigest` | `1.4.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 59.0 KiB | [postgresql-17-tdigest_1.4.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-17-tdigest_1.4.4-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tdigest_16` | `1.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.3 KiB | [tdigest_16-1.4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/tdigest_16-1.4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_16` | `1.4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.4 KiB | [tdigest_16-1.4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/tdigest_16-1.4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_16` | `1.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.1 KiB | [tdigest_16-1.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/tdigest_16-1.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_16` | `1.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.7 KiB | [tdigest_16-1.4.4-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/tdigest_16-1.4.4-4PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_16` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.0 KiB | [tdigest_16-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/tdigest_16-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `tdigest_16` | `1.4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.7 KiB | [tdigest_16-1.4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/tdigest_16-1.4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_16` | `1.4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.8 KiB | [tdigest_16-1.4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/tdigest_16-1.4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_16` | `1.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [tdigest_16-1.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/tdigest_16-1.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_16` | `1.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.4 KiB | [tdigest_16-1.4.4-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/tdigest_16-1.4.4-4PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_16` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.6 KiB | [tdigest_16-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/tdigest_16-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `tdigest_16` | `1.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.7 KiB | [tdigest_16-1.4.7-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/tdigest_16-1.4.7-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_16` | `1.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.8 KiB | [tdigest_16-1.4.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/tdigest_16-1.4.6-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_16` | `1.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.3 KiB | [tdigest_16-1.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/tdigest_16-1.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_16` | `1.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.5 KiB | [tdigest_16-1.4.4-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/tdigest_16-1.4.4-4PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_16` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.6 KiB | [tdigest_16-1.4.2-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/tdigest_16-1.4.2-4PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_16` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.1 KiB | [tdigest_16-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/tdigest_16-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `tdigest_16` | `1.4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.7 KiB | [tdigest_16-1.4.7-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/tdigest_16-1.4.7-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_16` | `1.4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [tdigest_16-1.4.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/tdigest_16-1.4.6-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_16` | `1.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.1 KiB | [tdigest_16-1.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/tdigest_16-1.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_16` | `1.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 33.6 KiB | [tdigest_16-1.4.4-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/tdigest_16-1.4.4-4PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_16` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.4 KiB | [tdigest_16-1.4.2-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/tdigest_16-1.4.2-4PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_16` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [tdigest_16-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/tdigest_16-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `tdigest_16` | `1.4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.1 KiB | [tdigest_16-1.4.7-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/tdigest_16-1.4.7-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_16` | `1.4.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.2 KiB | [tdigest_16-1.4.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/tdigest_16-1.4.6-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_16` | `1.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.7 KiB | [tdigest_16-1.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/tdigest_16-1.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_16` | `1.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 34.6 KiB | [tdigest_16-1.4.4-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/tdigest_16-1.4.4-4PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_16` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.5 KiB | [tdigest_16-1.4.2-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/tdigest_16-1.4.2-4PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_16` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.8 KiB | [tdigest_16-1.4.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/tdigest_16-1.4.2-2PGDG.rhel10.x86_64.rpm) |
| `tdigest_16` | `1.4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.3 KiB | [tdigest_16-1.4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/tdigest_16-1.4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_16` | `1.4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [tdigest_16-1.4.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/tdigest_16-1.4.6-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_16` | `1.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 38.7 KiB | [tdigest_16-1.4.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/tdigest_16-1.4.5-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_16` | `1.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 34.3 KiB | [tdigest_16-1.4.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/tdigest_16-1.4.4-4PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_16` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.9 KiB | [tdigest_16-1.4.2-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/tdigest_16-1.4.2-4PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_16` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.2 KiB | [tdigest_16-1.4.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/tdigest_16-1.4.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-tdigest` | `1.4.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 77.3 KiB | [postgresql-16-tdigest_1.4.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.7-1.pgdg12+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 72.2 KiB | [postgresql-16-tdigest_1.4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 59.8 KiB | [postgresql-16-tdigest_1.4.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.4-1.pgdg12+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 76.1 KiB | [postgresql-16-tdigest_1.4.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.7-1.pgdg12+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 71.1 KiB | [postgresql-16-tdigest_1.4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 59.3 KiB | [postgresql-16-tdigest_1.4.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.4-1.pgdg12+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 77.4 KiB | [postgresql-16-tdigest_1.4.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.7-1.pgdg13+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 72.3 KiB | [postgresql-16-tdigest_1.4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 59.8 KiB | [postgresql-16-tdigest_1.4.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.4-1.pgdg13+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 76.7 KiB | [postgresql-16-tdigest_1.4.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.7-1.pgdg13+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 71.6 KiB | [postgresql-16-tdigest_1.4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 59.7 KiB | [postgresql-16-tdigest_1.4.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.4-1.pgdg13+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 80.9 KiB | [postgresql-16-tdigest_1.4.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 74.7 KiB | [postgresql-16-tdigest_1.4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 62.5 KiB | [postgresql-16-tdigest_1.4.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 79.2 KiB | [postgresql-16-tdigest_1.4.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 73.6 KiB | [postgresql-16-tdigest_1.4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 61.5 KiB | [postgresql-16-tdigest_1.4.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 76.2 KiB | [postgresql-16-tdigest_1.4.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 71.5 KiB | [postgresql-16-tdigest_1.4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 59.8 KiB | [postgresql-16-tdigest_1.4.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 75.2 KiB | [postgresql-16-tdigest_1.4.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 70.5 KiB | [postgresql-16-tdigest_1.4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 59.2 KiB | [postgresql-16-tdigest_1.4.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 75.6 KiB | [postgresql-16-tdigest_1.4.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 71.4 KiB | [postgresql-16-tdigest_1.4.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 59.1 KiB | [postgresql-16-tdigest_1.4.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-tdigest` | `1.4.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 74.8 KiB | [postgresql-16-tdigest_1.4.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 70.6 KiB | [postgresql-16-tdigest_1.4.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-tdigest` | `1.4.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 59.0 KiB | [postgresql-16-tdigest_1.4.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-16-tdigest_1.4.4-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tdigest_15` | `1.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.3 KiB | [tdigest_15-1.4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/tdigest_15-1.4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_15` | `1.4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.4 KiB | [tdigest_15-1.4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/tdigest_15-1.4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_15` | `1.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.1 KiB | [tdigest_15-1.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/tdigest_15-1.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_15` | `1.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.7 KiB | [tdigest_15-1.4.4-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/tdigest_15-1.4.4-4PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_15` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.0 KiB | [tdigest_15-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/tdigest_15-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `tdigest_15` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 70.1 KiB | [tdigest_15-1.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/tdigest_15-1.4.0-1.rhel8.x86_64.rpm) |
| `tdigest_15` | `1.4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.6 KiB | [tdigest_15-1.4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/tdigest_15-1.4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_15` | `1.4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.8 KiB | [tdigest_15-1.4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/tdigest_15-1.4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_15` | `1.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.6 KiB | [tdigest_15-1.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/tdigest_15-1.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_15` | `1.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.4 KiB | [tdigest_15-1.4.4-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/tdigest_15-1.4.4-4PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_15` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.6 KiB | [tdigest_15-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/tdigest_15-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `tdigest_15` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 68.6 KiB | [tdigest_15-1.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/tdigest_15-1.4.0-1.rhel8.aarch64.rpm) |
| `tdigest_15` | `1.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.7 KiB | [tdigest_15-1.4.7-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/tdigest_15-1.4.7-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_15` | `1.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.8 KiB | [tdigest_15-1.4.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/tdigest_15-1.4.6-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_15` | `1.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.3 KiB | [tdigest_15-1.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/tdigest_15-1.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_15` | `1.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.5 KiB | [tdigest_15-1.4.4-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/tdigest_15-1.4.4-4PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_15` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.6 KiB | [tdigest_15-1.4.2-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/tdigest_15-1.4.2-4PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_15` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.1 KiB | [tdigest_15-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/tdigest_15-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `tdigest_15` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.1 KiB | [tdigest_15-1.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/tdigest_15-1.4.0-1.rhel9.x86_64.rpm) |
| `tdigest_15` | `1.4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.7 KiB | [tdigest_15-1.4.7-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/tdigest_15-1.4.7-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_15` | `1.4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [tdigest_15-1.4.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/tdigest_15-1.4.6-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_15` | `1.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.1 KiB | [tdigest_15-1.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/tdigest_15-1.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_15` | `1.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 33.6 KiB | [tdigest_15-1.4.4-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/tdigest_15-1.4.4-4PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_15` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.4 KiB | [tdigest_15-1.4.2-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/tdigest_15-1.4.2-4PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_15` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [tdigest_15-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/tdigest_15-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `tdigest_15` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 70.6 KiB | [tdigest_15-1.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/tdigest_15-1.4.0-1.rhel9.aarch64.rpm) |
| `tdigest_15` | `1.4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.1 KiB | [tdigest_15-1.4.7-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/tdigest_15-1.4.7-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_15` | `1.4.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.2 KiB | [tdigest_15-1.4.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/tdigest_15-1.4.6-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_15` | `1.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.7 KiB | [tdigest_15-1.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/tdigest_15-1.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_15` | `1.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 34.6 KiB | [tdigest_15-1.4.4-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/tdigest_15-1.4.4-4PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_15` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.5 KiB | [tdigest_15-1.4.2-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/tdigest_15-1.4.2-4PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_15` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.8 KiB | [tdigest_15-1.4.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/tdigest_15-1.4.2-2PGDG.rhel10.x86_64.rpm) |
| `tdigest_15` | `1.4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.3 KiB | [tdigest_15-1.4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/tdigest_15-1.4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_15` | `1.4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [tdigest_15-1.4.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/tdigest_15-1.4.6-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_15` | `1.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 38.7 KiB | [tdigest_15-1.4.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/tdigest_15-1.4.5-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_15` | `1.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 34.3 KiB | [tdigest_15-1.4.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/tdigest_15-1.4.4-4PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_15` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.9 KiB | [tdigest_15-1.4.2-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/tdigest_15-1.4.2-4PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_15` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.2 KiB | [tdigest_15-1.4.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/tdigest_15-1.4.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-tdigest` | `1.4.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 77.3 KiB | [postgresql-15-tdigest_1.4.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.7-1.pgdg12+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 72.2 KiB | [postgresql-15-tdigest_1.4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 59.8 KiB | [postgresql-15-tdigest_1.4.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.4-1.pgdg12+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 76.2 KiB | [postgresql-15-tdigest_1.4.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.7-1.pgdg12+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 71.1 KiB | [postgresql-15-tdigest_1.4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 59.3 KiB | [postgresql-15-tdigest_1.4.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.4-1.pgdg12+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 77.5 KiB | [postgresql-15-tdigest_1.4.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.7-1.pgdg13+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 72.3 KiB | [postgresql-15-tdigest_1.4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 59.9 KiB | [postgresql-15-tdigest_1.4.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.4-1.pgdg13+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 76.7 KiB | [postgresql-15-tdigest_1.4.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.7-1.pgdg13+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 71.6 KiB | [postgresql-15-tdigest_1.4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 59.7 KiB | [postgresql-15-tdigest_1.4.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.4-1.pgdg13+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.0 KiB | [postgresql-15-tdigest_1.4.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 74.8 KiB | [postgresql-15-tdigest_1.4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 62.6 KiB | [postgresql-15-tdigest_1.4.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 79.5 KiB | [postgresql-15-tdigest_1.4.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 73.7 KiB | [postgresql-15-tdigest_1.4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 61.7 KiB | [postgresql-15-tdigest_1.4.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 76.2 KiB | [postgresql-15-tdigest_1.4.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 71.5 KiB | [postgresql-15-tdigest_1.4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 59.8 KiB | [postgresql-15-tdigest_1.4.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 75.3 KiB | [postgresql-15-tdigest_1.4.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 70.5 KiB | [postgresql-15-tdigest_1.4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 59.3 KiB | [postgresql-15-tdigest_1.4.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 75.7 KiB | [postgresql-15-tdigest_1.4.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 71.6 KiB | [postgresql-15-tdigest_1.4.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 59.1 KiB | [postgresql-15-tdigest_1.4.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-tdigest` | `1.4.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 74.7 KiB | [postgresql-15-tdigest_1.4.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 70.6 KiB | [postgresql-15-tdigest_1.4.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-tdigest` | `1.4.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 59.0 KiB | [postgresql-15-tdigest_1.4.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-15-tdigest_1.4.4-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `tdigest_14` | `1.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.3 KiB | [tdigest_14-1.4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/tdigest_14-1.4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_14` | `1.4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.4 KiB | [tdigest_14-1.4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/tdigest_14-1.4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_14` | `1.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.1 KiB | [tdigest_14-1.4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/tdigest_14-1.4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_14` | `1.4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.6 KiB | [tdigest_14-1.4.4-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/tdigest_14-1.4.4-4PGDG.rhel8.10.x86_64.rpm) |
| `tdigest_14` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.0 KiB | [tdigest_14-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/tdigest_14-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `tdigest_14` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 60.3 KiB | [tdigest_14-1.2.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/tdigest_14-1.2.0-2.rhel8.x86_64.rpm) |
| `tdigest_14` | `1.4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.6 KiB | [tdigest_14-1.4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/tdigest_14-1.4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_14` | `1.4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.8 KiB | [tdigest_14-1.4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/tdigest_14-1.4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_14` | `1.4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.5 KiB | [tdigest_14-1.4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/tdigest_14-1.4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_14` | `1.4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.4 KiB | [tdigest_14-1.4.4-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/tdigest_14-1.4.4-4PGDG.rhel8.10.aarch64.rpm) |
| `tdigest_14` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.6 KiB | [tdigest_14-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/tdigest_14-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `tdigest_14` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 68.5 KiB | [tdigest_14-1.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/tdigest_14-1.4.0-1.rhel8.aarch64.rpm) |
| `tdigest_14` | `1.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.7 KiB | [tdigest_14-1.4.7-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/tdigest_14-1.4.7-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_14` | `1.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.8 KiB | [tdigest_14-1.4.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/tdigest_14-1.4.6-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_14` | `1.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.3 KiB | [tdigest_14-1.4.5-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/tdigest_14-1.4.5-1PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_14` | `1.4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.5 KiB | [tdigest_14-1.4.4-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/tdigest_14-1.4.4-4PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_14` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.6 KiB | [tdigest_14-1.4.2-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/tdigest_14-1.4.2-4PGDG.rhel9.8.x86_64.rpm) |
| `tdigest_14` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.1 KiB | [tdigest_14-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/tdigest_14-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `tdigest_14` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.2 KiB | [tdigest_14-1.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/tdigest_14-1.4.0-1.rhel9.x86_64.rpm) |
| `tdigest_14` | `1.4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.7 KiB | [tdigest_14-1.4.7-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/tdigest_14-1.4.7-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_14` | `1.4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [tdigest_14-1.4.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/tdigest_14-1.4.6-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_14` | `1.4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.1 KiB | [tdigest_14-1.4.5-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/tdigest_14-1.4.5-1PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_14` | `1.4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 33.6 KiB | [tdigest_14-1.4.4-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/tdigest_14-1.4.4-4PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_14` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.4 KiB | [tdigest_14-1.4.2-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/tdigest_14-1.4.2-4PGDG.rhel9.8.aarch64.rpm) |
| `tdigest_14` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.6 KiB | [tdigest_14-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/tdigest_14-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `tdigest_14` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 70.6 KiB | [tdigest_14-1.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/tdigest_14-1.4.0-1.rhel9.aarch64.rpm) |
| `tdigest_14` | `1.4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 45.1 KiB | [tdigest_14-1.4.7-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/tdigest_14-1.4.7-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_14` | `1.4.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.2 KiB | [tdigest_14-1.4.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/tdigest_14-1.4.6-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_14` | `1.4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.7 KiB | [tdigest_14-1.4.5-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/tdigest_14-1.4.5-1PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_14` | `1.4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 34.6 KiB | [tdigest_14-1.4.4-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/tdigest_14-1.4.4-4PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_14` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.5 KiB | [tdigest_14-1.4.2-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/tdigest_14-1.4.2-4PGDG.rhel10.2.x86_64.rpm) |
| `tdigest_14` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.8 KiB | [tdigest_14-1.4.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/tdigest_14-1.4.2-2PGDG.rhel10.x86_64.rpm) |
| `tdigest_14` | `1.4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.3 KiB | [tdigest_14-1.4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/tdigest_14-1.4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_14` | `1.4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [tdigest_14-1.4.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/tdigest_14-1.4.6-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_14` | `1.4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 38.7 KiB | [tdigest_14-1.4.5-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/tdigest_14-1.4.5-1PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_14` | `1.4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 34.2 KiB | [tdigest_14-1.4.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/tdigest_14-1.4.4-4PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_14` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.9 KiB | [tdigest_14-1.4.2-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/tdigest_14-1.4.2-4PGDG.rhel10.2.aarch64.rpm) |
| `tdigest_14` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 33.2 KiB | [tdigest_14-1.4.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/tdigest_14-1.4.2-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-tdigest` | `1.4.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 77.3 KiB | [postgresql-14-tdigest_1.4.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.7-1.pgdg12+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 72.2 KiB | [postgresql-14-tdigest_1.4.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.6-1.pgdg12+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 59.7 KiB | [postgresql-14-tdigest_1.4.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.4-1.pgdg12+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 76.1 KiB | [postgresql-14-tdigest_1.4.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.7-1.pgdg12+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 71.1 KiB | [postgresql-14-tdigest_1.4.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.6-1.pgdg12+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 59.2 KiB | [postgresql-14-tdigest_1.4.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.4-1.pgdg12+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 77.4 KiB | [postgresql-14-tdigest_1.4.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.7-1.pgdg13+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 72.3 KiB | [postgresql-14-tdigest_1.4.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.6-1.pgdg13+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 59.8 KiB | [postgresql-14-tdigest_1.4.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.4-1.pgdg13+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 76.6 KiB | [postgresql-14-tdigest_1.4.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.7-1.pgdg13+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 71.6 KiB | [postgresql-14-tdigest_1.4.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.6-1.pgdg13+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 59.6 KiB | [postgresql-14-tdigest_1.4.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.4-1.pgdg13+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.1 KiB | [postgresql-14-tdigest_1.4.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 74.8 KiB | [postgresql-14-tdigest_1.4.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 62.6 KiB | [postgresql-14-tdigest_1.4.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 79.4 KiB | [postgresql-14-tdigest_1.4.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 73.7 KiB | [postgresql-14-tdigest_1.4.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 61.6 KiB | [postgresql-14-tdigest_1.4.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 76.1 KiB | [postgresql-14-tdigest_1.4.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 71.5 KiB | [postgresql-14-tdigest_1.4.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 59.7 KiB | [postgresql-14-tdigest_1.4.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 75.2 KiB | [postgresql-14-tdigest_1.4.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 70.5 KiB | [postgresql-14-tdigest_1.4.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 59.1 KiB | [postgresql-14-tdigest_1.4.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 75.6 KiB | [postgresql-14-tdigest_1.4.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 71.5 KiB | [postgresql-14-tdigest_1.4.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 59.1 KiB | [postgresql-14-tdigest_1.4.4-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.4-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-tdigest` | `1.4.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 74.8 KiB | [postgresql-14-tdigest_1.4.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 70.6 KiB | [postgresql-14-tdigest_1.4.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.6-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-tdigest` | `1.4.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 58.9 KiB | [postgresql-14-tdigest_1.4.4-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/t/tdigest/postgresql-14-tdigest_1.4.4-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/tdigest" title="Repository" icon="github" subtitle="github.com/tvondra/tdigest" />}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install tdigest;		# install via package name, for the active PG version

pig install tdigest -v 18;   # install for PG 18
pig install tdigest -v 17;   # install for PG 17
pig install tdigest -v 16;   # install for PG 16
pig install tdigest -v 15;   # install for PG 15
pig install tdigest -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION tdigest;
```

## Usage

Sources:

- [README.md](https://github.com/tvondra/tdigest/blob/c0af713163e78fe7d7717559fd4862adcb8162d7/README.md)
- [tdigest.control](https://github.com/tvondra/tdigest/blob/c0af713163e78fe7d7717559fd4862adcb8162d7/tdigest.control)
- [Changes](https://github.com/tvondra/tdigest/blob/c0af713163e78fe7d7717559fd4862adcb8162d7/Changes)
- [tdigest--1.4.6--1.4.7.sql](https://github.com/tvondra/tdigest/blob/c0af713163e78fe7d7717559fd4862adcb8162d7/tdigest--1.4.6--1.4.7.sql)

`tdigest` 1.4.7 supplies approximate, mergeable rank statistics. Store partial digests and combine them to calculate quantiles without sorting the complete original data.

### Core Workflow

```sql
CREATE EXTENSION tdigest;
SELECT tdigest_percentile(v, 100, ARRAY[0.5, 0.95, 0.99])
FROM generate_series(1, 1000) AS g(v);
CREATE TABLE digest_daily AS
SELECT current_date AS day, tdigest(v, 100) AS digest
FROM generate_series(1, 1000) AS g(v);
SELECT tdigest_percentile(digest, 0.95) FROM digest_daily;
```

### Functions and Accuracy

`tdigest(value, compression)` builds a digest; `tdigest(digest)` merges stored digests. `tdigest_percentile` estimates quantiles and `tdigest_percentile_of` estimates ranks, with scalar and array forms. `tdigest_add` and `tdigest_union` update or combine digests; `tdigest_count`, `tdigest_sum` and `tdigest_avg` report counts and trimmed aggregates. The low and high arguments to trimmed aggregates are quantile thresholds, not raw value bounds. `tdigest_is_valid` checks serialized values.

`compression` must be between 10 and 10000. Higher values trade memory and CPU for generally better accuracy, but do not provide a fixed error bound. Validate estimates against exact results on representative data and use consistent compression when merging states.

### Upgrade and Boundaries

Version 1.4.7 fixes closely spaced percentile interpolation, large-count inverse ranks, reusable-state finalization and NULL-compression handling, and marks more scalar functions parallel safe. Upgrade the installed library and SQL together with `ALTER EXTENSION tdigest UPDATE TO '1.4.7'`. Review stored digests with `tdigest_is_valid`; older malformed values can be rejected by the stricter validation. Installation requires superuser rights; the extension is relocatable and requires no preload. The release metadata sets PostgreSQL 13 as the minimum, and the changelog records PostgreSQL 19 build fixes without making a broad performance guarantee.
