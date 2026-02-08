---
title: "nominatim_fdw"
linkTitle: "nominatim_fdw"
description: "Nominatim Foreign Data Wrapper for PostgreSQL"
weight: 8680
categories: ["FDW"]
width: full
---

[**nominatim_fdw**](https://github.com/jimjonesbr/nominatim_fdw) : Nominatim Foreign Data Wrapper for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8680** | {{< badge content="nominatim_fdw" link="https://github.com/jimjonesbr/nominatim_fdw" >}} | {{< ext "nominatim_fdw" >}} | `1.1.0` | {{< category "FDW" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `nominatim_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.0` | {{< bg "18" "nominatim_fdw_18" "green" >}} {{< bg "17" "nominatim_fdw_17" "green" >}} {{< bg "16" "nominatim_fdw_16" "green" >}} {{< bg "15" "nominatim_fdw_15" "green" >}} {{< bg "14" "nominatim_fdw_14" "green" >}} {{< bg "13" "nominatim_fdw_13" "red" >}} | `nominatim_fdw_$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "nominatim_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "nominatim_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "nominatim_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "nominatim_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "nominatim_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "nominatim_fdw_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "nominatim_fdw : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.6 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jimjonesbr/nominatim_fdw" title="Repository" icon="github" subtitle="github.com/jimjonesbr/nominatim_fdw" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install nominatim_fdw;		# install via package name, for the active PG version

pig install nominatim_fdw -v 18;   # install for PG 18
pig install nominatim_fdw -v 17;   # install for PG 17
pig install nominatim_fdw -v 16;   # install for PG 16
pig install nominatim_fdw -v 15;   # install for PG 15
pig install nominatim_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION nominatim_fdw;
```
