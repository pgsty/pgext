---
title: "faker"
linkTitle: "faker"
description: "Wrapper for the Faker Python library"
weight: 3210
categories: ["LANG"]
width: full
---

[**faker**](https://github.com/anpandu/postgresql_faker) : Wrapper for the Faker Python library


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3210** | {{< badge content="faker" link="https://github.com/anpandu/postgresql_faker" >}} | {{< ext "faker" >}} | `0.5.3` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "Python" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpython3u" >}} {{< ext "pgtap" >}} {{< ext "dbt2" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "ltree_plpython3u" >}} {{< ext "hstore_plpython3u" >}} {{< ext "random" >}} {{< ext "pg_tle" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.5.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `faker` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.5.3` | {{< bg "18" "postgresql_faker_18*" "green" >}} {{< bg "17" "postgresql_faker_17*" "green" >}} {{< bg "16" "postgresql_faker_16*" "green" >}} {{< bg "15" "postgresql_faker_15*" "green" >}} {{< bg "14" "postgresql_faker_14*" "green" >}} {{< bg "13" "postgresql_faker_13*" "green" >}} | `postgresql_faker_$v*` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_13 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.3" "postgresql_faker_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |      {{< bg "MISS" "faker : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql_faker_18` | `0.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.0 KiB | [postgresql_faker_18-0.5.3-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/postgresql_faker_18-0.5.3-7PGDG.rhel8.x86_64.rpm) |
| `postgresql_faker_18` | `0.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.1 KiB | [postgresql_faker_18-0.5.3-7PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/postgresql_faker_18-0.5.3-7PGDG.rhel8.aarch64.rpm) |
| `postgresql_faker_18` | `0.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.0 KiB | [postgresql_faker_18-0.5.3-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgresql_faker_18-0.5.3-7PGDG.rhel9.x86_64.rpm) |
| `postgresql_faker_18` | `0.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 43.8 KiB | [postgresql_faker_18-0.5.3-7PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgresql_faker_18-0.5.3-7PGDG.rhel9.aarch64.rpm) |
| `postgresql_faker_18` | `0.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [postgresql_faker_18-0.5.3-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgresql_faker_18-0.5.3-7PGDG.rhel10.x86_64.rpm) |
| `postgresql_faker_18` | `0.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.5 KiB | [postgresql_faker_18-0.5.3-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgresql_faker_18-0.5.3-7PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql_faker_17` | `0.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.9 KiB | [postgresql_faker_17-0.5.3-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgresql_faker_17-0.5.3-6PGDG.rhel8.x86_64.rpm) |
| `postgresql_faker_17` | `0.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.0 KiB | [postgresql_faker_17-0.5.3-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgresql_faker_17-0.5.3-6PGDG.rhel8.aarch64.rpm) |
| `postgresql_faker_17` | `0.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.2 KiB | [postgresql_faker_17-0.5.3-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgresql_faker_17-0.5.3-6PGDG.rhel9.x86_64.rpm) |
| `postgresql_faker_17` | `0.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.1 KiB | [postgresql_faker_17-0.5.3-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgresql_faker_17-0.5.3-6PGDG.rhel9.aarch64.rpm) |
| `postgresql_faker_17` | `0.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [postgresql_faker_17-0.5.3-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgresql_faker_17-0.5.3-7PGDG.rhel10.x86_64.rpm) |
| `postgresql_faker_17` | `0.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.5 KiB | [postgresql_faker_17-0.5.3-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgresql_faker_17-0.5.3-7PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql_faker_16` | `0.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.4 KiB | [postgresql_faker_16-0.5.3-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgresql_faker_16-0.5.3-3PGDG.rhel8.x86_64.rpm) |
| `postgresql_faker_16` | `0.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.6 KiB | [postgresql_faker_16-0.5.3-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgresql_faker_16-0.5.3-3PGDG.rhel8.aarch64.rpm) |
| `postgresql_faker_16` | `0.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 44.1 KiB | [postgresql_faker_16-0.5.3-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgresql_faker_16-0.5.3-3PGDG.rhel9.x86_64.rpm) |
| `postgresql_faker_16` | `0.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.1 KiB | [postgresql_faker_16-0.5.3-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgresql_faker_16-0.5.3-3PGDG.rhel9.aarch64.rpm) |
| `postgresql_faker_16` | `0.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [postgresql_faker_16-0.5.3-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgresql_faker_16-0.5.3-7PGDG.rhel10.x86_64.rpm) |
| `postgresql_faker_16` | `0.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.5 KiB | [postgresql_faker_16-0.5.3-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgresql_faker_16-0.5.3-7PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql_faker_15` | `0.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.6 KiB | [postgresql_faker_15-0.5.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgresql_faker_15-0.5.3-1.rhel8.x86_64.rpm) |
| `postgresql_faker_15` | `0.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.8 KiB | [postgresql_faker_15-0.5.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgresql_faker_15-0.5.3-1.rhel8.aarch64.rpm) |
| `postgresql_faker_15` | `0.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.6 KiB | [postgresql_faker_15-0.5.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgresql_faker_15-0.5.3-1.rhel9.x86_64.rpm) |
| `postgresql_faker_15` | `0.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [postgresql_faker_15-0.5.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgresql_faker_15-0.5.3-1.rhel9.aarch64.rpm) |
| `postgresql_faker_15` | `0.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [postgresql_faker_15-0.5.3-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgresql_faker_15-0.5.3-7PGDG.rhel10.x86_64.rpm) |
| `postgresql_faker_15` | `0.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.5 KiB | [postgresql_faker_15-0.5.3-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgresql_faker_15-0.5.3-7PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql_faker_14` | `0.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.9 KiB | [postgresql_faker_14-0.5.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgresql_faker_14-0.5.3-1.rhel8.x86_64.rpm) |
| `postgresql_faker_14` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.7 KiB | [postgresql_faker_14-0.4.0-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgresql_faker_14-0.4.0-1.rhel8.noarch.rpm) |
| `postgresql_faker_14` | `0.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.8 KiB | [postgresql_faker_14-0.5.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgresql_faker_14-0.5.3-1.rhel8.aarch64.rpm) |
| `postgresql_faker_14` | `0.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.6 KiB | [postgresql_faker_14-0.5.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresql_faker_14-0.5.3-1.rhel9.x86_64.rpm) |
| `postgresql_faker_14` | `0.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 47.9 KiB | [postgresql_faker_14-0.5.3-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresql_faker_14-0.5.3-1.rhel9.noarch.rpm) |
| `postgresql_faker_14` | `0.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [postgresql_faker_14-0.5.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgresql_faker_14-0.5.3-1.rhel9.aarch64.rpm) |
| `postgresql_faker_14` | `0.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [postgresql_faker_14-0.5.3-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgresql_faker_14-0.5.3-7PGDG.rhel10.x86_64.rpm) |
| `postgresql_faker_14` | `0.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.5 KiB | [postgresql_faker_14-0.5.3-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgresql_faker_14-0.5.3-7PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql_faker_13` | `0.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.9 KiB | [postgresql_faker_13-0.5.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/postgresql_faker_13-0.5.3-1.rhel8.x86_64.rpm) |
| `postgresql_faker_13` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 37.6 KiB | [postgresql_faker_13-0.4.0-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/postgresql_faker_13-0.4.0-1.rhel8.noarch.rpm) |
| `postgresql_faker_13` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.0 KiB | [postgresql_faker_13-0.3.0-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/postgresql_faker_13-0.3.0-1.rhel8.noarch.rpm) |
| `postgresql_faker_13` | `0.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.8 KiB | [postgresql_faker_13-0.5.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/postgresql_faker_13-0.5.3-1.rhel8.aarch64.rpm) |
| `postgresql_faker_13` | `0.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.6 KiB | [postgresql_faker_13-0.5.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/postgresql_faker_13-0.5.3-1.rhel9.x86_64.rpm) |
| `postgresql_faker_13` | `0.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 47.9 KiB | [postgresql_faker_13-0.5.3-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/postgresql_faker_13-0.5.3-1.rhel9.noarch.rpm) |
| `postgresql_faker_13` | `0.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [postgresql_faker_13-0.5.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/postgresql_faker_13-0.5.3-1.rhel9.aarch64.rpm) |
| `postgresql_faker_13` | `0.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 44.3 KiB | [postgresql_faker_13-0.5.3-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/postgresql_faker_13-0.5.3-7PGDG.rhel10.x86_64.rpm) |
| `postgresql_faker_13` | `0.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 44.5 KiB | [postgresql_faker_13-0.5.3-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/postgresql_faker_13-0.5.3-7PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/anpandu/postgresql_faker" title="Repository" icon="github" subtitle="github.com/anpandu/postgresql_faker" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install faker;		# install via package name, for the active PG version

pig install faker -v 18;   # install for PG 18
pig install faker -v 17;   # install for PG 17
pig install faker -v 16;   # install for PG 16
pig install faker -v 15;   # install for PG 15
pig install faker -v 14;   # install for PG 14
pig install faker -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION faker;
```
