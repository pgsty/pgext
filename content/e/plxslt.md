---
title: "plxslt"
linkTitle: "plxslt"
description: "XSLT procedural language for PostgreSQL"
weight: 3110
categories: ["LANG"]
width: full
---

XSLT procedural language for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3110** | {{< badge content="plxslt" link="https://github.com/petere/plxslt" >}} | {{< ext "plxslt" >}} | `0.20140221` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpgsql" >}} {{< ext "pgml" >}} {{< ext "plpython3u" >}} {{< ext "pg_tle" >}} {{< ext "plv8" >}} {{< ext "pljava" >}} {{< ext "plperl" >}} {{< ext "pllua" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/plxslt" >}} | `0.20140221` | {{< bg "18" "plxslt_18*" "red" >}} {{< bg "17" "plxslt_17*" "green" >}} {{< bg "16" "plxslt_16*" "green" >}} {{< bg "15" "plxslt_15*" "green" >}} {{< bg "14" "plxslt_14*" "green" >}} {{< bg "13" "plxslt_13*" "green" >}} | `plxslt_$v*` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.20140221" "plxslt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.20140221" "plxslt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.20140221" "plxslt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.20140221" "plxslt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.20140221" "plxslt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.20140221" "plxslt_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.20140221" "plxslt_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |      {{< bg "MISS" "plxslt : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plxslt_18` | 0.20140221 | `el8.x86_64` | pgdg | 14.7 KiB | [plxslt_18-0.20140221-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plxslt_18-0.20140221-1PGDG.rhel8.x86_64.rpm) |
| `plxslt_18` | 0.20140221 | `el8.aarch64` | pgdg | 14.7 KiB | [plxslt_18-0.20140221-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plxslt_18-0.20140221-1PGDG.rhel8.aarch64.rpm) |
| `plxslt_18` | 0.20140221 | `el9.x86_64` | pgdg | 14.9 KiB | [plxslt_18-0.20140221-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plxslt_18-0.20140221-1PGDG.rhel9.x86_64.rpm) |
| `plxslt_18` | 0.20140221 | `el9.aarch64` | pgdg | 14.6 KiB | [plxslt_18-0.20140221-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plxslt_18-0.20140221-1PGDG.rhel9.aarch64.rpm) |
| `plxslt_18` | 0.20140221 | `el10.x86_64` | pgdg | 15.3 KiB | [plxslt_18-0.20140221-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plxslt_18-0.20140221-1PGDG.rhel10.x86_64.rpm) |
| `plxslt_18` | 0.20140221 | `el10.aarch64` | pgdg | 15.2 KiB | [plxslt_18-0.20140221-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plxslt_18-0.20140221-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plxslt_17` | 0.20140221 | `el8.x86_64` | pgdg | 14.7 KiB | [plxslt_17-0.20140221-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plxslt_17-0.20140221-1PGDG.rhel8.x86_64.rpm) |
| `plxslt_17` | 0.20140221 | `el8.aarch64` | pgdg | 14.7 KiB | [plxslt_17-0.20140221-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plxslt_17-0.20140221-1PGDG.rhel8.aarch64.rpm) |
| `plxslt_17` | 0.20140221 | `el9.x86_64` | pgdg | 14.9 KiB | [plxslt_17-0.20140221-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plxslt_17-0.20140221-1PGDG.rhel9.x86_64.rpm) |
| `plxslt_17` | 0.20140221 | `el9.aarch64` | pgdg | 14.7 KiB | [plxslt_17-0.20140221-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plxslt_17-0.20140221-1PGDG.rhel9.aarch64.rpm) |
| `plxslt_17` | 0.20140221 | `el10.x86_64` | pgdg | 15.3 KiB | [plxslt_17-0.20140221-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plxslt_17-0.20140221-1PGDG.rhel10.x86_64.rpm) |
| `plxslt_17` | 0.20140221 | `el10.aarch64` | pgdg | 15.2 KiB | [plxslt_17-0.20140221-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plxslt_17-0.20140221-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plxslt_16` | 0.20140221 | `el8.x86_64` | pgdg | 14.7 KiB | [plxslt_16-0.20140221-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plxslt_16-0.20140221-1PGDG.rhel8.x86_64.rpm) |
| `plxslt_16` | 0.20140221 | `el8.aarch64` | pgdg | 14.7 KiB | [plxslt_16-0.20140221-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plxslt_16-0.20140221-1PGDG.rhel8.aarch64.rpm) |
| `plxslt_16` | 0.20140221 | `el9.x86_64` | pgdg | 14.9 KiB | [plxslt_16-0.20140221-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plxslt_16-0.20140221-1PGDG.rhel9.x86_64.rpm) |
| `plxslt_16` | 0.20140221 | `el9.aarch64` | pgdg | 14.7 KiB | [plxslt_16-0.20140221-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plxslt_16-0.20140221-1PGDG.rhel9.aarch64.rpm) |
| `plxslt_16` | 0.20140221 | `el10.x86_64` | pgdg | 15.3 KiB | [plxslt_16-0.20140221-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plxslt_16-0.20140221-1PGDG.rhel10.x86_64.rpm) |
| `plxslt_16` | 0.20140221 | `el10.aarch64` | pgdg | 15.2 KiB | [plxslt_16-0.20140221-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plxslt_16-0.20140221-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plxslt_15` | 0.20140221 | `el8.x86_64` | pgdg | 14.7 KiB | [plxslt_15-0.20140221-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plxslt_15-0.20140221-1PGDG.rhel8.x86_64.rpm) |
| `plxslt_15` | 0.20140221 | `el8.aarch64` | pgdg | 14.7 KiB | [plxslt_15-0.20140221-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plxslt_15-0.20140221-1PGDG.rhel8.aarch64.rpm) |
| `plxslt_15` | 0.20140221 | `el9.x86_64` | pgdg | 14.9 KiB | [plxslt_15-0.20140221-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plxslt_15-0.20140221-1PGDG.rhel9.x86_64.rpm) |
| `plxslt_15` | 0.20140221 | `el9.aarch64` | pgdg | 14.7 KiB | [plxslt_15-0.20140221-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plxslt_15-0.20140221-1PGDG.rhel9.aarch64.rpm) |
| `plxslt_15` | 0.20140221 | `el10.x86_64` | pgdg | 15.3 KiB | [plxslt_15-0.20140221-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plxslt_15-0.20140221-1PGDG.rhel10.x86_64.rpm) |
| `plxslt_15` | 0.20140221 | `el10.aarch64` | pgdg | 15.2 KiB | [plxslt_15-0.20140221-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plxslt_15-0.20140221-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plxslt_14` | 0.20140221 | `el8.x86_64` | pgdg | 14.7 KiB | [plxslt_14-0.20140221-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plxslt_14-0.20140221-1PGDG.rhel8.x86_64.rpm) |
| `plxslt_14` | 0.20140221 | `el8.aarch64` | pgdg | 14.7 KiB | [plxslt_14-0.20140221-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plxslt_14-0.20140221-1PGDG.rhel8.aarch64.rpm) |
| `plxslt_14` | 0.20140221 | `el9.x86_64` | pgdg | 14.9 KiB | [plxslt_14-0.20140221-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plxslt_14-0.20140221-1PGDG.rhel9.x86_64.rpm) |
| `plxslt_14` | 0.20140221 | `el9.aarch64` | pgdg | 14.7 KiB | [plxslt_14-0.20140221-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plxslt_14-0.20140221-1PGDG.rhel9.aarch64.rpm) |
| `plxslt_14` | 0.20140221 | `el10.x86_64` | pgdg | 15.3 KiB | [plxslt_14-0.20140221-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plxslt_14-0.20140221-1PGDG.rhel10.x86_64.rpm) |
| `plxslt_14` | 0.20140221 | `el10.aarch64` | pgdg | 15.2 KiB | [plxslt_14-0.20140221-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plxslt_14-0.20140221-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plxslt_13` | 0.20140221 | `el8.x86_64` | pgdg | 14.6 KiB | [plxslt_13-0.20140221-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/plxslt_13-0.20140221-1PGDG.rhel8.x86_64.rpm) |
| `plxslt_13` | 0.20140221 | `el8.aarch64` | pgdg | 14.7 KiB | [plxslt_13-0.20140221-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/plxslt_13-0.20140221-1PGDG.rhel8.aarch64.rpm) |
| `plxslt_13` | 0.20140221 | `el9.x86_64` | pgdg | 14.9 KiB | [plxslt_13-0.20140221-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/plxslt_13-0.20140221-1PGDG.rhel9.x86_64.rpm) |
| `plxslt_13` | 0.20140221 | `el9.aarch64` | pgdg | 14.7 KiB | [plxslt_13-0.20140221-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/plxslt_13-0.20140221-1PGDG.rhel9.aarch64.rpm) |
| `plxslt_13` | 0.20140221 | `el10.x86_64` | pgdg | 15.3 KiB | [plxslt_13-0.20140221-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/plxslt_13-0.20140221-1PGDG.rhel10.x86_64.rpm) |
| `plxslt_13` | 0.20140221 | `el10.aarch64` | pgdg | 15.2 KiB | [plxslt_13-0.20140221-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/plxslt_13-0.20140221-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petere/plxslt" title="Repository" icon="github" subtitle="github.com/petere/plxslt" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install plxslt; # install by extension name, for the current active PG version
pig ext install plxslt; # install via package alias, for the active PG version
pig ext install plxslt -v 17;   # install for PG 17
pig ext install plxslt -v 16;   # install for PG 16
pig ext install plxslt -v 15;   # install for PG 15
pig ext install plxslt -v 14;   # install for PG 14
pig ext install plxslt -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION plxslt;
```

