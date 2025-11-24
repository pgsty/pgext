---
title: "dbt2"
linkTitle: "dbt2"
description: "OSDL-DBT-2 test kit"
weight: 3220
categories: ["LANG"]
width: full
---

[**dbt2**](https://github.com/osdldbt/dbt2) : OSDL-DBT-2 test kit


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3220** | {{< badge content="dbt2" link="https://github.com/osdldbt/dbt2" >}} | {{< ext "dbt2" >}} | `0.61.7` | {{< category "LANG" >}} | {{< license "Artistic" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgtap" >}} {{< ext "faker" >}} {{< ext "plpgsql" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_tle" >}} {{< ext "plv8" >}} {{< ext "pllua" >}} {{< ext "hstore_pllua" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.61.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `dbt2` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.61.7` | {{< bg "18" "dbt2-pg18-extensions*" "green" >}} {{< bg "17" "dbt2-pg17-extensions*" "green" >}} {{< bg "16" "dbt2-pg16-extensions*" "green" >}} {{< bg "15" "dbt2-pg15-extensions*" "green" >}} {{< bg "14" "dbt2-pg14-extensions*" "green" >}} {{< bg "13" "dbt2-pg13-extensions*" "green" >}} | `dbt2-pg$v-extensions*` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "dbt2-pg18-extensions : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2-pg17-extensions : MISS 0" "red" >}}      | {{< bg "PGDG 0.53.7" "dbt2-pg16-extensions : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.53.7" "dbt2-pg15-extensions : AVAIL 5" "blue" >}} | {{< bg "PGDG 0.53.7" "dbt2-pg14-extensions : AVAIL 6" "blue" >}} | {{< bg "PGDG 0.53.7" "dbt2-pg13-extensions : AVAIL 6" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "dbt2-pg18-extensions : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2-pg17-extensions : MISS 0" "red" >}}      | {{< bg "PGDG 0.53.7" "dbt2-pg16-extensions : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.53.7" "dbt2-pg15-extensions : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.53.7" "dbt2-pg14-extensions : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.53.7" "dbt2-pg13-extensions : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg18-extensions : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg17-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg16-extensions : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg15-extensions : AVAIL 8" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg14-extensions : AVAIL 7" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg13-extensions : AVAIL 7" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg18-extensions : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg17-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg16-extensions : AVAIL 4" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg15-extensions : AVAIL 8" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg14-extensions : AVAIL 8" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg13-extensions : AVAIL 8" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg18-extensions : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg17-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg16-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg15-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg14-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg13-extensions : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg18-extensions : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg17-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg16-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg15-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg14-extensions : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.61.7" "dbt2-pg13-extensions : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |      {{< bg "MISS" "dbt2 : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg18-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [dbt2-pg18-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/dbt2-pg18-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg18-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg18-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/dbt2-pg18-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg18-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg18-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/dbt2-pg18-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg18-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg18-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10.0-aarch64/dbt2-pg18-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg17-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [dbt2-pg17-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/dbt2-pg17-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg17-extensions` | `0.61.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg17-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/dbt2-pg17-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg17-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg17-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/dbt2-pg17-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg17-extensions` | `0.61.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg17-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/dbt2-pg17-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg17-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg17-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/dbt2-pg17-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg17-extensions` | `0.61.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg17-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/dbt2-pg17-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg17-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg17-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10.0-aarch64/dbt2-pg17-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm) |
| `dbt2-pg17-extensions` | `0.61.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.3 KiB | [dbt2-pg17-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10.0-aarch64/dbt2-pg17-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg16-extensions` | `0.53.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.9 KiB | [dbt2-pg16-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/dbt2-pg16-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.53.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [dbt2-pg16-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/dbt2-pg16-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [dbt2-pg16-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/dbt2-pg16-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.61.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg16-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/dbt2-pg16-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.53.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [dbt2-pg16-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/dbt2-pg16-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.53.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.4 KiB | [dbt2-pg16-extensions-0.53.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/dbt2-pg16-extensions-0.53.6-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg16-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/dbt2-pg16-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.61.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg16-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/dbt2-pg16-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.53.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.0 KiB | [dbt2-pg16-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/dbt2-pg16-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.53.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [dbt2-pg16-extensions-0.53.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/dbt2-pg16-extensions-0.53.6-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg16-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/dbt2-pg16-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.61.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg16-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/dbt2-pg16-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg16-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg16-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10.0-aarch64/dbt2-pg16-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm) |
| `dbt2-pg16-extensions` | `0.61.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.3 KiB | [dbt2-pg16-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10.0-aarch64/dbt2-pg16-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg15-extensions` | `0.53.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.9 KiB | [dbt2-pg15-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.53.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.8 KiB | [dbt2-pg15-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.50.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.50.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.50.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.49.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.49.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.49.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.48.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.48.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.48.7-1.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.53.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [dbt2-pg15-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/dbt2-pg15-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.53.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/dbt2-pg15-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.50.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [dbt2-pg15-extensions-0.50.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/dbt2-pg15-extensions-0.50.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.49.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [dbt2-pg15-extensions-0.49.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/dbt2-pg15-extensions-0.49.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg15-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.61.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg15-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.53.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [dbt2-pg15-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.53.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [dbt2-pg15-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.50.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.1 KiB | [dbt2-pg15-extensions-0.50.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.50.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.49.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg15-extensions-0.49.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.49.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.48.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg15-extensions-0.48.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.48.7-1.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.48.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [dbt2-pg15-extensions-0.48.3-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.48.3-2.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.6 KiB | [dbt2-pg15-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.61.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg15-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.53.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [dbt2-pg15-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.53.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.8 KiB | [dbt2-pg15-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.50.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.50.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.50.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.49.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.49.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.49.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.48.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.4 KiB | [dbt2-pg15-extensions-0.48.7-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.48.7-1.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.48.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.3 KiB | [dbt2-pg15-extensions-0.48.3-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.48.3-2.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg15-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/dbt2-pg15-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.61.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg15-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/dbt2-pg15-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg15-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg15-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10.0-aarch64/dbt2-pg15-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm) |
| `dbt2-pg15-extensions` | `0.61.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg15-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10.0-aarch64/dbt2-pg15-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg14-extensions` | `0.53.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.9 KiB | [dbt2-pg14-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.53.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.7 KiB | [dbt2-pg14-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.50.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.5 KiB | [dbt2-pg14-extensions-0.50.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.50.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.49.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.5 KiB | [dbt2-pg14-extensions-0.49.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.49.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.48.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.4 KiB | [dbt2-pg14-extensions-0.48.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.48.7-1.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.48.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.3 KiB | [dbt2-pg14-extensions-0.48.3-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.48.3-2.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.53.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [dbt2-pg14-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/dbt2-pg14-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.53.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.4 KiB | [dbt2-pg14-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/dbt2-pg14-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.50.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [dbt2-pg14-extensions-0.50.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/dbt2-pg14-extensions-0.50.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.49.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [dbt2-pg14-extensions-0.49.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/dbt2-pg14-extensions-0.49.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [dbt2-pg14-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.61.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg14-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.53.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [dbt2-pg14-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.53.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [dbt2-pg14-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.50.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.1 KiB | [dbt2-pg14-extensions-0.50.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.50.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.49.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg14-extensions-0.49.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.49.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.48.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg14-extensions-0.48.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.48.7-1.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.6 KiB | [dbt2-pg14-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.61.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg14-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.53.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [dbt2-pg14-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.53.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.8 KiB | [dbt2-pg14-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.50.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.5 KiB | [dbt2-pg14-extensions-0.50.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.50.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.49.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.5 KiB | [dbt2-pg14-extensions-0.49.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.49.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.48.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.4 KiB | [dbt2-pg14-extensions-0.48.7-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.48.7-1.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.48.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.3 KiB | [dbt2-pg14-extensions-0.48.3-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.48.3-2.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.6 KiB | [dbt2-pg14-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/dbt2-pg14-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.61.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg14-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/dbt2-pg14-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg14-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg14-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10.0-aarch64/dbt2-pg14-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm) |
| `dbt2-pg14-extensions` | `0.61.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg14-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10.0-aarch64/dbt2-pg14-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg13-extensions` | `0.53.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.7 KiB | [dbt2-pg13-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/dbt2-pg13-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.53.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.5 KiB | [dbt2-pg13-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/dbt2-pg13-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.50.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.3 KiB | [dbt2-pg13-extensions-0.50.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/dbt2-pg13-extensions-0.50.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.49.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.2 KiB | [dbt2-pg13-extensions-0.49.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/dbt2-pg13-extensions-0.49.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.48.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.2 KiB | [dbt2-pg13-extensions-0.48.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/dbt2-pg13-extensions-0.48.7-1.rhel8.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.48.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.0 KiB | [dbt2-pg13-extensions-0.48.3-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/dbt2-pg13-extensions-0.48.3-2.rhel8.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.53.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.6 KiB | [dbt2-pg13-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/dbt2-pg13-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg13-extensions` | `0.53.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.4 KiB | [dbt2-pg13-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/dbt2-pg13-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg13-extensions` | `0.50.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [dbt2-pg13-extensions-0.50.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/dbt2-pg13-extensions-0.50.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg13-extensions` | `0.49.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [dbt2-pg13-extensions-0.49.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/dbt2-pg13-extensions-0.49.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg13-extensions` | `0.61.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.9 KiB | [dbt2-pg13-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/dbt2-pg13-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.61.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg13-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/dbt2-pg13-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.53.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [dbt2-pg13-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/dbt2-pg13-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.53.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.3 KiB | [dbt2-pg13-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/dbt2-pg13-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.50.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.1 KiB | [dbt2-pg13-extensions-0.50.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/dbt2-pg13-extensions-0.50.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.49.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg13-extensions-0.49.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/dbt2-pg13-extensions-0.49.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.48.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.0 KiB | [dbt2-pg13-extensions-0.48.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/dbt2-pg13-extensions-0.48.7-1.rhel9.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.61.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.6 KiB | [dbt2-pg13-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/dbt2-pg13-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg13-extensions` | `0.61.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.6 KiB | [dbt2-pg13-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/dbt2-pg13-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg13-extensions` | `0.53.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [dbt2-pg13-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/dbt2-pg13-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg13-extensions` | `0.53.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [dbt2-pg13-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/dbt2-pg13-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg13-extensions` | `0.50.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.5 KiB | [dbt2-pg13-extensions-0.50.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/dbt2-pg13-extensions-0.50.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg13-extensions` | `0.49.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.4 KiB | [dbt2-pg13-extensions-0.49.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/dbt2-pg13-extensions-0.49.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg13-extensions` | `0.48.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.4 KiB | [dbt2-pg13-extensions-0.48.7-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/dbt2-pg13-extensions-0.48.7-1.rhel9.aarch64.rpm) |
| `dbt2-pg13-extensions` | `0.48.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.2 KiB | [dbt2-pg13-extensions-0.48.3-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/dbt2-pg13-extensions-0.48.3-2.rhel9.aarch64.rpm) |
| `dbt2-pg13-extensions` | `0.61.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.6 KiB | [dbt2-pg13-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/dbt2-pg13-extensions-0.61.7-1PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.61.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.7 KiB | [dbt2-pg13-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/dbt2-pg13-extensions-0.61.6-2PGDG.rhel10.x86_64.rpm) |
| `dbt2-pg13-extensions` | `0.61.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg13-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10.0-aarch64/dbt2-pg13-extensions-0.61.7-1PGDG.rhel10.aarch64.rpm) |
| `dbt2-pg13-extensions` | `0.61.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.2 KiB | [dbt2-pg13-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10.0-aarch64/dbt2-pg13-extensions-0.61.6-2PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/osdldbt/dbt2" title="Repository" icon="github" subtitle="github.com/osdldbt/dbt2" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install dbt2;		# install via package name, for the active PG version

pig install dbt2 -v 18;   # install for PG 18
pig install dbt2 -v 17;   # install for PG 17
pig install dbt2 -v 16;   # install for PG 16
pig install dbt2 -v 15;   # install for PG 15
pig install dbt2 -v 14;   # install for PG 14
pig install dbt2 -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION dbt2;
```
