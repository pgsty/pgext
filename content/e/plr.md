---
title: "plr"
linkTitle: "plr"
description: "load R interpreter and execute R script from within a database"
weight: 3100
categories: ["LANG"]
languages: ["C"]
licenses: ["GPL-2.0"]
repos: ["PGDG"]
page_width: full
---

[**plr**](https://github.com/postgres-plr/plr) : load R interpreter and execute R script from within a database


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3100** | {{< badge content="plr" link="https://github.com/postgres-plr/plr" >}} | {{< ext "plr" >}} | `8.4.8.7` | {{< category "LANG" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plxslt" >}} {{< ext "pltcl" >}} {{< ext "plperl" >}} {{< ext "pljava" >}} {{< ext "plsh" >}} {{< ext "plpython3u" >}} {{< ext "plpgsql" >}} {{< ext "plperlu" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `8.4.8.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plr` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `8.4.8.6` | {{< bg "18" "plr_18" "green" >}} {{< bg "17" "plr_17" "green" >}} {{< bg "16" "plr_16" "green" >}} {{< bg "15" "plr_15" "green" >}} {{< bg "14" "plr_14" "green" >}} | `plr_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `8.4.8.7` | {{< bg "18" "postgresql-18-plr" "green" >}} {{< bg "17" "postgresql-17-plr" "green" >}} {{< bg "16" "postgresql-16-plr" "green" >}} {{< bg "15" "postgresql-15-plr" "green" >}} {{< bg "14" "postgresql-14-plr" "green" >}} | `postgresql-$v-plr` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 8.4.8.6" "plr_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_14 : AVAIL 7" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 8.4.8.6" "plr_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_14 : AVAIL 5" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 8.4.8.6" "plr_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_16 : AVAIL 10" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_14 : AVAIL 9" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 8.4.8.6" "plr_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_15 : AVAIL 8" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_14 : AVAIL 8" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 8.4.8.6" "plr_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_14 : AVAIL 5" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 8.4.8.6" "plr_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 8.4.8.6" "plr_14 : AVAIL 6" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-14-plr : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-18-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-17-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-16-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-15-plr : AVAIL 3" "blue" >}} | {{< bg "PGDG 8.4.8.7" "postgresql-14-plr : AVAIL 3" "blue" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_18` | `8.4.8.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.8 KiB | [plr_18-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plr_18-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_18` | `8.4.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.4 KiB | [plr_18-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plr_18-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_18` | `8.4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.1 KiB | [plr_18-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plr_18-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_18` | `8.4.8.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.5 KiB | [plr_18-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plr_18-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_18` | `8.4.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.0 KiB | [plr_18-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plr_18-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_18` | `8.4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.7 KiB | [plr_18-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plr_18-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_18` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.4 KiB | [plr_18-8.4.8.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plr_18-8.4.8.6-1PGDG.rhel9.8.x86_64.rpm) |
| `plr_18` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.4 KiB | [plr_18-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plr_18-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_18` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.6 KiB | [plr_18-8.4.8.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plr_18-8.4.8.6-1PGDG.rhel9.6.x86_64.rpm) |
| `plr_18` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.0 KiB | [plr_18-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plr_18-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_18` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.1 KiB | [plr_18-8.4.8.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plr_18-8.4.8.4-1PGDG.rhel9.6.x86_64.rpm) |
| `plr_18` | `8.4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.9 KiB | [plr_18-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plr_18-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_18` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.7 KiB | [plr_18-8.4.8.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plr_18-8.4.8.6-1PGDG.rhel9.8.aarch64.rpm) |
| `plr_18` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.8 KiB | [plr_18-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plr_18-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_18` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.9 KiB | [plr_18-8.4.8.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plr_18-8.4.8.6-1PGDG.rhel9.6.aarch64.rpm) |
| `plr_18` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.3 KiB | [plr_18-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plr_18-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_18` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.4 KiB | [plr_18-8.4.8.4-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plr_18-8.4.8.4-1PGDG.rhel9.6.aarch64.rpm) |
| `plr_18` | `8.4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.1 KiB | [plr_18-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plr_18-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_18` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.8 KiB | [plr_18-8.4.8.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plr_18-8.4.8.6-1PGDG.rhel10.2.x86_64.rpm) |
| `plr_18` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.8 KiB | [plr_18-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plr_18-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_18` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 77.1 KiB | [plr_18-8.4.8.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plr_18-8.4.8.6-1PGDG.rhel10.0.x86_64.rpm) |
| `plr_18` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.4 KiB | [plr_18-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plr_18-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_18` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.8 KiB | [plr_18-8.4.8.4-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plr_18-8.4.8.4-1PGDG.rhel10.0.x86_64.rpm) |
| `plr_18` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.7 KiB | [plr_18-8.4.8.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plr_18-8.4.8.6-1PGDG.rhel10.2.aarch64.rpm) |
| `plr_18` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.7 KiB | [plr_18-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plr_18-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_18` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.8 KiB | [plr_18-8.4.8.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plr_18-8.4.8.6-1PGDG.rhel10.0.aarch64.rpm) |
| `plr_18` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.3 KiB | [plr_18-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plr_18-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_18` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.3 KiB | [plr_18-8.4.8.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plr_18-8.4.8.4-1PGDG.rhel10.0.aarch64.rpm) |
| `plr_18` | `8.4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.7 KiB | [plr_18-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plr_18-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-plr` | `8.4.8.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.9 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg12+2_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.9 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg12+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.9 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg12+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.5 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg12+2_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.5 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg12+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.5 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg12+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.2 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg13+2_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.2 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg13+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.1 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg13+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.7 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg13+2_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.7 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg13+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.7 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg13+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 131.7 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg22.04+2_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 131.6 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 131.6 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 128.5 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg22.04+2_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 128.5 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 128.5 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.3 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg24.04+2_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.3 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.2 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.9 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg24.04+2_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.9 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.8 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.6 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg26.04+2_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.6 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.6 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.4 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg26.04+2_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.4 KiB | [postgresql-18-plr_8.4.8.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-plr` | `8.4.8.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.3 KiB | [postgresql-18-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-18-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_17` | `8.4.8.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.7 KiB | [plr_17-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plr_17-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_17` | `8.4.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.3 KiB | [plr_17-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plr_17-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_17` | `8.4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.0 KiB | [plr_17-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plr_17-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_17` | `8.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.7 KiB | [plr_17-8.4.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plr_17-8.4.7-1PGDG.rhel8.x86_64.rpm) |
| `plr_17` | `8.4.8.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.4 KiB | [plr_17-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plr_17-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_17` | `8.4.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.0 KiB | [plr_17-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plr_17-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_17` | `8.4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.7 KiB | [plr_17-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plr_17-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_17` | `8.4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.4 KiB | [plr_17-8.4.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plr_17-8.4.7-1PGDG.rhel8.aarch64.rpm) |
| `plr_17` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.3 KiB | [plr_17-8.4.8.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plr_17-8.4.8.6-1PGDG.rhel9.8.x86_64.rpm) |
| `plr_17` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.3 KiB | [plr_17-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plr_17-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_17` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.4 KiB | [plr_17-8.4.8.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plr_17-8.4.8.6-1PGDG.rhel9.6.x86_64.rpm) |
| `plr_17` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.9 KiB | [plr_17-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plr_17-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_17` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.0 KiB | [plr_17-8.4.8.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plr_17-8.4.8.4-1PGDG.rhel9.6.x86_64.rpm) |
| `plr_17` | `8.4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.8 KiB | [plr_17-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plr_17-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_17` | `8.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.6 KiB | [plr_17-8.4.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plr_17-8.4.7-1PGDG.rhel9.x86_64.rpm) |
| `plr_17` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.6 KiB | [plr_17-8.4.8.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plr_17-8.4.8.6-1PGDG.rhel9.8.aarch64.rpm) |
| `plr_17` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.7 KiB | [plr_17-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plr_17-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_17` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.8 KiB | [plr_17-8.4.8.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plr_17-8.4.8.6-1PGDG.rhel9.6.aarch64.rpm) |
| `plr_17` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.2 KiB | [plr_17-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plr_17-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_17` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.3 KiB | [plr_17-8.4.8.4-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plr_17-8.4.8.4-1PGDG.rhel9.6.aarch64.rpm) |
| `plr_17` | `8.4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.2 KiB | [plr_17-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plr_17-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_17` | `8.4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.0 KiB | [plr_17-8.4.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plr_17-8.4.7-1PGDG.rhel9.aarch64.rpm) |
| `plr_17` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.7 KiB | [plr_17-8.4.8.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plr_17-8.4.8.6-1PGDG.rhel10.2.x86_64.rpm) |
| `plr_17` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.7 KiB | [plr_17-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plr_17-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_17` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 77.1 KiB | [plr_17-8.4.8.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plr_17-8.4.8.6-1PGDG.rhel10.0.x86_64.rpm) |
| `plr_17` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.2 KiB | [plr_17-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plr_17-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_17` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.6 KiB | [plr_17-8.4.8.4-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plr_17-8.4.8.4-1PGDG.rhel10.0.x86_64.rpm) |
| `plr_17` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.6 KiB | [plr_17-8.4.8.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plr_17-8.4.8.6-1PGDG.rhel10.2.aarch64.rpm) |
| `plr_17` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.6 KiB | [plr_17-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plr_17-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_17` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.6 KiB | [plr_17-8.4.8.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plr_17-8.4.8.6-1PGDG.rhel10.0.aarch64.rpm) |
| `plr_17` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.3 KiB | [plr_17-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plr_17-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_17` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.3 KiB | [plr_17-8.4.8.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plr_17-8.4.8.4-1PGDG.rhel10.0.aarch64.rpm) |
| `plr_17` | `8.4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.5 KiB | [plr_17-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plr_17-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-plr` | `8.4.8.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.0 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg12+2_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.9 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg12+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.8 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg12+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.3 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg12+2_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.3 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg12+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.2 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg12+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.0 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg13+2_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.0 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg13+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.0 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg13+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.5 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg13+2_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.5 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg13+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.6 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg13+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 155.6 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg22.04+2_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 155.6 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 155.5 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 152.3 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg22.04+2_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 152.3 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 152.3 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.3 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg24.04+2_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.4 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.2 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.7 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg24.04+2_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.6 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.5 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.3 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg26.04+2_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.2 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.2 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.1 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg26.04+2_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.1 KiB | [postgresql-17-plr_8.4.8.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-plr` | `8.4.8.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.0 KiB | [postgresql-17-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-17-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_16` | `8.4.8.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.7 KiB | [plr_16-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plr_16-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_16` | `8.4.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.3 KiB | [plr_16-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plr_16-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_16` | `8.4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.0 KiB | [plr_16-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plr_16-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_16` | `8.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.7 KiB | [plr_16-8.4.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plr_16-8.4.7-1PGDG.rhel8.x86_64.rpm) |
| `plr_16` | `8.4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.8 KiB | [plr_16-8.4.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plr_16-8.4.6-1PGDG.rhel8.x86_64.rpm) |
| `plr_16` | `8.4.8.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.4 KiB | [plr_16-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plr_16-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_16` | `8.4.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.0 KiB | [plr_16-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plr_16-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_16` | `8.4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.6 KiB | [plr_16-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plr_16-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_16` | `8.4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.4 KiB | [plr_16-8.4.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plr_16-8.4.7-1PGDG.rhel8.aarch64.rpm) |
| `plr_16` | `8.4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 72.6 KiB | [plr_16-8.4.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plr_16-8.4.6-1PGDG.rhel8.aarch64.rpm) |
| `plr_16` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.3 KiB | [plr_16-8.4.8.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.8.6-1PGDG.rhel9.8.x86_64.rpm) |
| `plr_16` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.3 KiB | [plr_16-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_16` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.4 KiB | [plr_16-8.4.8.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.8.6-1PGDG.rhel9.6.x86_64.rpm) |
| `plr_16` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.9 KiB | [plr_16-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_16` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.0 KiB | [plr_16-8.4.8.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.8.4-1PGDG.rhel9.6.x86_64.rpm) |
| `plr_16` | `8.4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.8 KiB | [plr_16-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | `8.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.6 KiB | [plr_16-8.4.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.7-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | `8.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.8 KiB | [plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | `8.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.6 KiB | [plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | `8.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.6 KiB | [plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_16-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_16` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.7 KiB | [plr_16-8.4.8.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.8.6-1PGDG.rhel9.8.aarch64.rpm) |
| `plr_16` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.7 KiB | [plr_16-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_16` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.8 KiB | [plr_16-8.4.8.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.8.6-1PGDG.rhel9.6.aarch64.rpm) |
| `plr_16` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.2 KiB | [plr_16-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_16` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.3 KiB | [plr_16-8.4.8.4-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.8.4-1PGDG.rhel9.6.aarch64.rpm) |
| `plr_16` | `8.4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.2 KiB | [plr_16-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_16` | `8.4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.0 KiB | [plr_16-8.4.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.7-1PGDG.rhel9.aarch64.rpm) |
| `plr_16` | `8.4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.1 KiB | [plr_16-8.4.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plr_16-8.4.6-1PGDG.rhel9.aarch64.rpm) |
| `plr_16` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.7 KiB | [plr_16-8.4.8.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plr_16-8.4.8.6-1PGDG.rhel10.2.x86_64.rpm) |
| `plr_16` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.7 KiB | [plr_16-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plr_16-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_16` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 77.1 KiB | [plr_16-8.4.8.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plr_16-8.4.8.6-1PGDG.rhel10.0.x86_64.rpm) |
| `plr_16` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.2 KiB | [plr_16-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plr_16-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_16` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.6 KiB | [plr_16-8.4.8.4-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plr_16-8.4.8.4-1PGDG.rhel10.0.x86_64.rpm) |
| `plr_16` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.6 KiB | [plr_16-8.4.8.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plr_16-8.4.8.6-1PGDG.rhel10.2.aarch64.rpm) |
| `plr_16` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.6 KiB | [plr_16-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plr_16-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_16` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.6 KiB | [plr_16-8.4.8.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plr_16-8.4.8.6-1PGDG.rhel10.0.aarch64.rpm) |
| `plr_16` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.3 KiB | [plr_16-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plr_16-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_16` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.3 KiB | [plr_16-8.4.8.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plr_16-8.4.8.4-1PGDG.rhel10.0.aarch64.rpm) |
| `plr_16` | `8.4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.5 KiB | [plr_16-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plr_16-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-plr` | `8.4.8.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.2 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg12+2_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.2 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg12+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.0 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg12+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.3 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg12+2_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.3 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg12+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.2 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg12+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.2 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg13+2_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.2 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg13+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.0 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg13+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.6 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg13+2_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.5 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg13+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.4 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg13+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.6 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg22.04+2_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.6 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.6 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.4 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg22.04+2_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.4 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.3 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.4 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg24.04+2_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.3 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.4 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.7 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg24.04+2_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.7 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.6 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.3 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg26.04+2_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.3 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.2 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.1 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg26.04+2_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.1 KiB | [postgresql-16-plr_8.4.8.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-plr` | `8.4.8.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.1 KiB | [postgresql-16-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-16-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_15` | `8.4.8.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.3 KiB | [plr_15-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_15` | `8.4.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.9 KiB | [plr_15-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_15` | `8.4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.5 KiB | [plr_15-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_15` | `8.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.2 KiB | [plr_15-8.4.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.7-1PGDG.rhel8.x86_64.rpm) |
| `plr_15` | `8.4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.4 KiB | [plr_15-8.4.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.6-1PGDG.rhel8.x86_64.rpm) |
| `plr_15` | `8.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 167.4 KiB | [plr_15-8.4.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plr_15-8.4.5-1.rhel8.x86_64.rpm) |
| `plr_15` | `8.4.8.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.9 KiB | [plr_15-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plr_15-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_15` | `8.4.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.4 KiB | [plr_15-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plr_15-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_15` | `8.4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.1 KiB | [plr_15-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plr_15-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_15` | `8.4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.9 KiB | [plr_15-8.4.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plr_15-8.4.7-1PGDG.rhel8.aarch64.rpm) |
| `plr_15` | `8.4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.1 KiB | [plr_15-8.4.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plr_15-8.4.6-1PGDG.rhel8.aarch64.rpm) |
| `plr_15` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.0 KiB | [plr_15-8.4.8.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.8.6-1PGDG.rhel9.8.x86_64.rpm) |
| `plr_15` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.0 KiB | [plr_15-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_15` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.1 KiB | [plr_15-8.4.8.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.8.6-1PGDG.rhel9.6.x86_64.rpm) |
| `plr_15` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.4 KiB | [plr_15-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_15` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.6 KiB | [plr_15-8.4.8.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.8.4-1PGDG.rhel9.6.x86_64.rpm) |
| `plr_15` | `8.4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.5 KiB | [plr_15-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_15` | `8.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.4 KiB | [plr_15-8.4.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.7-1PGDG.rhel9.x86_64.rpm) |
| `plr_15` | `8.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.6 KiB | [plr_15-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_15` | `8.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 167.8 KiB | [plr_15-8.4.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plr_15-8.4.5-1.rhel9.x86_64.rpm) |
| `plr_15` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 74.2 KiB | [plr_15-8.4.8.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.8.6-1PGDG.rhel9.8.aarch64.rpm) |
| `plr_15` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 74.2 KiB | [plr_15-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_15` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 74.4 KiB | [plr_15-8.4.8.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.8.6-1PGDG.rhel9.6.aarch64.rpm) |
| `plr_15` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.8 KiB | [plr_15-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_15` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.9 KiB | [plr_15-8.4.8.4-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.8.4-1PGDG.rhel9.6.aarch64.rpm) |
| `plr_15` | `8.4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.8 KiB | [plr_15-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_15` | `8.4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.6 KiB | [plr_15-8.4.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.7-1PGDG.rhel9.aarch64.rpm) |
| `plr_15` | `8.4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.7 KiB | [plr_15-8.4.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plr_15-8.4.6-1PGDG.rhel9.aarch64.rpm) |
| `plr_15` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 77.2 KiB | [plr_15-8.4.8.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plr_15-8.4.8.6-1PGDG.rhel10.2.x86_64.rpm) |
| `plr_15` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 77.2 KiB | [plr_15-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plr_15-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_15` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 77.6 KiB | [plr_15-8.4.8.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plr_15-8.4.8.6-1PGDG.rhel10.0.x86_64.rpm) |
| `plr_15` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.8 KiB | [plr_15-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plr_15-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_15` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 77.2 KiB | [plr_15-8.4.8.4-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plr_15-8.4.8.4-1PGDG.rhel10.0.x86_64.rpm) |
| `plr_15` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 75.1 KiB | [plr_15-8.4.8.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plr_15-8.4.8.6-1PGDG.rhel10.2.aarch64.rpm) |
| `plr_15` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 75.1 KiB | [plr_15-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plr_15-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_15` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 75.1 KiB | [plr_15-8.4.8.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plr_15-8.4.8.6-1PGDG.rhel10.0.aarch64.rpm) |
| `plr_15` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.8 KiB | [plr_15-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plr_15-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_15` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.8 KiB | [plr_15-8.4.8.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plr_15-8.4.8.4-1PGDG.rhel10.0.aarch64.rpm) |
| `plr_15` | `8.4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.1 KiB | [plr_15-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plr_15-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-plr` | `8.4.8.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.3 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg12+2_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.3 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg12+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.1 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg12+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.6 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg12+2_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.6 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg12+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.6 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg12+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.4 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg13+2_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.3 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg13+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.4 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg13+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.9 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg13+2_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.9 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg13+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.8 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg13+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.4 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg22.04+2_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.4 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.4 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.3 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg22.04+2_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.4 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.3 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.7 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg24.04+2_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.6 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.4 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.8 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg24.04+2_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.9 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.7 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.8 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg26.04+2_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.7 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.5 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.3 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg26.04+2_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.2 KiB | [postgresql-15-plr_8.4.8.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-plr` | `8.4.8.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.2 KiB | [postgresql-15-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-15-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plr_14` | `8.4.8.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.4 KiB | [plr_14-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.8.6-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_14` | `8.4.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.9 KiB | [plr_14-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plr_14` | `8.4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.5 KiB | [plr_14-8.4.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.8-1PGDG.rhel8.x86_64.rpm) |
| `plr_14` | `8.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.2 KiB | [plr_14-8.4.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.7-1PGDG.rhel8.x86_64.rpm) |
| `plr_14` | `8.4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.4 KiB | [plr_14-8.4.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.6-1PGDG.rhel8.x86_64.rpm) |
| `plr_14` | `8.4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 166.7 KiB | [plr_14-8.4.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.5-1.rhel8.x86_64.rpm) |
| `plr_14` | `8.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 166.5 KiB | [plr_14-8.4.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plr_14-8.4.3-1.rhel8.x86_64.rpm) |
| `plr_14` | `8.4.8.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.9 KiB | [plr_14-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plr_14-8.4.8.6-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_14` | `8.4.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.4 KiB | [plr_14-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plr_14-8.4.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plr_14` | `8.4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.1 KiB | [plr_14-8.4.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plr_14-8.4.8-1PGDG.rhel8.aarch64.rpm) |
| `plr_14` | `8.4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.9 KiB | [plr_14-8.4.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plr_14-8.4.7-1PGDG.rhel8.aarch64.rpm) |
| `plr_14` | `8.4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.1 KiB | [plr_14-8.4.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plr_14-8.4.6-1PGDG.rhel8.aarch64.rpm) |
| `plr_14` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.9 KiB | [plr_14-8.4.8.6-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.8.6-1PGDG.rhel9.8.x86_64.rpm) |
| `plr_14` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.9 KiB | [plr_14-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.8.6-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_14` | `8.4.8.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.1 KiB | [plr_14-8.4.8.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.8.6-1PGDG.rhel9.6.x86_64.rpm) |
| `plr_14` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.4 KiB | [plr_14-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.8.4-1PGDG.rhel9.7.x86_64.rpm) |
| `plr_14` | `8.4.8.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.6 KiB | [plr_14-8.4.8.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.8.4-1PGDG.rhel9.6.x86_64.rpm) |
| `plr_14` | `8.4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.5 KiB | [plr_14-8.4.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.8-1PGDG.rhel9.x86_64.rpm) |
| `plr_14` | `8.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.3 KiB | [plr_14-8.4.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.7-1PGDG.rhel9.x86_64.rpm) |
| `plr_14` | `8.4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 73.6 KiB | [plr_14-8.4.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.6-1PGDG.rhel9.x86_64.rpm) |
| `plr_14` | `8.4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 167.4 KiB | [plr_14-8.4.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plr_14-8.4.5-1.rhel9.x86_64.rpm) |
| `plr_14` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 74.2 KiB | [plr_14-8.4.8.6-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.8.6-1PGDG.rhel9.8.aarch64.rpm) |
| `plr_14` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 74.2 KiB | [plr_14-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.8.6-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_14` | `8.4.8.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 74.4 KiB | [plr_14-8.4.8.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.8.6-1PGDG.rhel9.6.aarch64.rpm) |
| `plr_14` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.8 KiB | [plr_14-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.8.4-1PGDG.rhel9.7.aarch64.rpm) |
| `plr_14` | `8.4.8.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.9 KiB | [plr_14-8.4.8.4-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.8.4-1PGDG.rhel9.6.aarch64.rpm) |
| `plr_14` | `8.4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.8 KiB | [plr_14-8.4.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.8-1PGDG.rhel9.aarch64.rpm) |
| `plr_14` | `8.4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.6 KiB | [plr_14-8.4.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.7-1PGDG.rhel9.aarch64.rpm) |
| `plr_14` | `8.4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.7 KiB | [plr_14-8.4.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plr_14-8.4.6-1PGDG.rhel9.aarch64.rpm) |
| `plr_14` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 77.2 KiB | [plr_14-8.4.8.6-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plr_14-8.4.8.6-1PGDG.rhel10.2.x86_64.rpm) |
| `plr_14` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 77.2 KiB | [plr_14-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plr_14-8.4.8.6-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_14` | `8.4.8.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 77.6 KiB | [plr_14-8.4.8.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plr_14-8.4.8.6-1PGDG.rhel10.0.x86_64.rpm) |
| `plr_14` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 76.8 KiB | [plr_14-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plr_14-8.4.8.4-1PGDG.rhel10.1.x86_64.rpm) |
| `plr_14` | `8.4.8.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 77.2 KiB | [plr_14-8.4.8.4-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plr_14-8.4.8.4-1PGDG.rhel10.0.x86_64.rpm) |
| `plr_14` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 75.1 KiB | [plr_14-8.4.8.6-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plr_14-8.4.8.6-1PGDG.rhel10.2.aarch64.rpm) |
| `plr_14` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 75.1 KiB | [plr_14-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plr_14-8.4.8.6-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_14` | `8.4.8.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 75.1 KiB | [plr_14-8.4.8.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plr_14-8.4.8.6-1PGDG.rhel10.0.aarch64.rpm) |
| `plr_14` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.8 KiB | [plr_14-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plr_14-8.4.8.4-1PGDG.rhel10.1.aarch64.rpm) |
| `plr_14` | `8.4.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.8 KiB | [plr_14-8.4.8.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plr_14-8.4.8.4-1PGDG.rhel10.0.aarch64.rpm) |
| `plr_14` | `8.4.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.1 KiB | [plr_14-8.4.8-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plr_14-8.4.8-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-plr` | `8.4.8.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.2 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg12+2_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.2 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg12+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 136.2 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg12+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.7 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg12+2_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.7 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg12+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 132.6 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg12+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.5 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg13+2_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.4 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg13+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.2 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg13+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.9 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg13+2_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.9 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg13+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 132.9 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg13+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.5 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg22.04+2_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.5 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 151.6 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.6 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg22.04+2_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.6 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 148.5 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.6 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg24.04+2_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.5 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.4 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.8 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg24.04+2_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.8 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 123.8 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.6 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg26.04+2_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.6 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [u26.x86_64](/os/u26.x86_64) | pgdg | 125.6 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.2 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg26.04+2_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.2 KiB | [postgresql-14-plr_8.4.8.7-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.7-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-plr` | `8.4.8.6` | [u26.aarch64](/os/u26.aarch64) | pgdg | 122.2 KiB | [postgresql-14-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plr/postgresql-14-plr_8.4.8.6-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgres-plr/plr" title="Repository" icon="github" subtitle="github.com/postgres-plr/plr" />}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install plr;		# install via package name, for the active PG version

pig install plr -v 18;   # install for PG 18
pig install plr -v 17;   # install for PG 17
pig install plr -v 16;   # install for PG 16
pig install plr -v 15;   # install for PG 15
pig install plr -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plr;
```

## Usage

Sources:

- [8.4.8.7 README](https://github.com/postgres-plr/plr/blob/REL8_4_8_7/README.md)
- [Versioned user guide](https://github.com/postgres-plr/plr/blob/REL8_4_8_7/userguide.md)
- [Control file](https://github.com/postgres-plr/plr/blob/REL8_4_8_7/plr.control)
- [Version 8.4.8.7 SQL](https://github.com/postgres-plr/plr/blob/REL8_4_8_7/plr--8.4.8.7.sql)

`plr` enables writing PostgreSQL functions in the R programming language, providing full access to R's statistical and data analysis capabilities.

```sql
CREATE EXTENSION plr;
```

### Create Functions

```sql
CREATE OR REPLACE FUNCTION r_max(integer, integer) RETURNS integer AS '
if (arg1 > arg2)
  return(arg1)
else
  return(arg2)
' LANGUAGE plr STRICT;

SELECT r_max(10, 20);  -- 20
```

With named arguments:

```sql
CREATE OR REPLACE FUNCTION sd(vals float8[]) RETURNS float AS '
sd(vals)
' LANGUAGE plr STRICT;

SELECT sd(ARRAY[1.0, 2.0, 3.0, 4.0, 5.0]);
```

### Argument Handling

- Unnamed arguments are available as `arg1`, `arg2`, ...; an explicitly named parameter replaces its corresponding `argN` variable.
- Scalar SQL NULL becomes R NULL; null array elements become R `NA`. A `STRICT` function skips calls with a whole NULL argument, not arrays containing null elements.
- Composite types (rows) are passed as R data.frames
- One-dimensional arrays become R vectors; two-dimensional arrays become matrices and three-dimensional arrays become R arrays. Higher dimensions are unsupported.

```sql
CREATE OR REPLACE FUNCTION r_max(integer, integer) RETURNS integer AS '
if (is.null(arg1) && is.null(arg2))
  return(NULL)
if (is.null(arg1))
  return(arg2)
if (is.null(arg2))
  return(arg1)
if (arg1 > arg2)
  return(arg1)
return(arg2)
' LANGUAGE plr;
```

### Database Access via SPI

```sql
CREATE OR REPLACE FUNCTION test_spi(text) RETURNS SETOF record AS '
pg.spi.exec(arg1)
' LANGUAGE plr;

SELECT * FROM test_spi('SELECT oid, typname FROM pg_type LIMIT 5')
  AS t(oid oid, typname name);
```

Initialize the type OID variables in the same connection before calling a function that prepares a parameterized query:

```sql
SELECT load_r_typenames();

CREATE OR REPLACE FUNCTION lookup_type(type_name text)
RETURNS SETOF record AS $$
sp <- pg.spi.prepare(
  'SELECT oid, typname FROM pg_type WHERE typname = $1',
  c(NAMEOID)
)
pg.spi.execp(sp, list(type_name))
$$ LANGUAGE plr;

SELECT * FROM lookup_type('text') AS t(oid oid, typname name);
```

### Set-Returning Functions

Return an R vector for a set of scalar values:

```sql
CREATE OR REPLACE FUNCTION get_numbers(n int) RETURNS SETOF integer AS '
1:n
' LANGUAGE plr;

SELECT * FROM get_numbers(5);
```

### Window Functions

```sql
CREATE OR REPLACE FUNCTION r_regr_slope(float8, float8, int)
RETURNS float8 AS '
slope <- NA
y <- farg1
x <- farg2
if (fnumrows == arg3 + 1L)
  try(slope <- lm(y ~ x)$coefficients[2])
return(slope)
' LANGUAGE plr WINDOW;
```

Window functions receive `farg1..fargN` (vectors of values in the window frame), `fnumrows` (frame size), and `prownum` (current row position in partition).

### Global Variables

Persist data across function calls using R's global environment:

```sql
CREATE OR REPLACE FUNCTION set_state(key text, val text) RETURNS void AS '
assign(key, val, env=.GlobalEnv)
' LANGUAGE plr;
```

### Useful Support Functions

```sql
SELECT load_r_typenames();  -- Load type OID variables
SELECT * FROM r_typenames(); -- List available type OIDs
SELECT plr_version();        -- PL/R version
```

### Trigger Functions

PL/R supports trigger functions with access to `pg.tg.name`, `pg.tg.relname`, `pg.tg.when`, `pg.tg.level`, `pg.tg.op`, `pg.tg.new`, and `pg.tg.old`.

### Runtime and Privileges

This page follows PL/R 8.4.8.7. PL/R is an untrusted language: creating its functions requires a superuser, and R code runs with the PostgreSQL operating-system user's access to files and processes. Review function bodies and EXECUTE grants accordingly. The R shared library must be available, and upstream requires `R_HOME` in the PostgreSQL server process environment before startup on Unix systems. Adding it only to an interactive client shell does not configure the server.

SQL scalar NULL converts to R NULL, while null elements within an array convert to R NA; declaring a function STRICT prevents calls with null arguments. R global state belongs to a backend process, not to all sessions or to a durable database table. Upstream revokes PUBLIC execution of environment-changing helpers including `plr_set_rhome(text)`; use an administrator-managed runtime rather than exposing these helpers to application roles. Normal usage does not require shared preload.
