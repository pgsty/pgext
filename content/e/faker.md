---
title: "faker"
linkTitle: "faker"
description: "Wrapper for the Faker Python library"
weight: 3210
categories: ["LANG"]
width: full
---

Wrapper for the Faker Python library


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3210** | {{< badge content="faker" link="https://github.com/anpandu/postgresql_faker" >}} | {{< ext "faker" >}} | `0.5.3` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "Python" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plpython3u" >}} {{< ext "pgtap" >}} {{< ext "dbt2" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "ltree_plpython3u" >}} {{< ext "hstore_plpython3u" >}} {{< ext "random" >}} {{< ext "pg_tle" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/faker" >}} | `0.5.3` | {{< bg "18" "postgresql_faker_18*" "green" >}} {{< bg "17" "postgresql_faker_17*" "green" >}} {{< bg "16" "postgresql_faker_16*" "green" >}} {{< bg "15" "postgresql_faker_15*" "green" >}} {{< bg "14" "postgresql_faker_14*" "green" >}} {{< bg "13" "postgresql_faker_13*" "green" >}} | `postgresql_faker_$v*` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |  {{< bg "PGDG 0.5.3" "postgresql_faker_18 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_17 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_16 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_15 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_14 : HIDE 2" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_13 : HIDE 3" >}}   |
|    `el8.aarch64`    |  {{< bg "PGDG 0.5.3" "postgresql_faker_18 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_17 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_16 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_15 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_14 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_13 : HIDE 1" >}}   |
|    `el9.x86_64`    |  {{< bg "PGDG 0.5.3" "postgresql_faker_18 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_17 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_16 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_15 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_14 : HIDE 2" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_13 : HIDE 2" >}}   |
|    `el9.aarch64`    |  {{< bg "PGDG 0.5.3" "postgresql_faker_18 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_17 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_16 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_15 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_14 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_13 : HIDE 1" >}}   |
|    `el10.x86_64`    |  {{< bg "PGDG 0.5.3" "postgresql_faker_18 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_17 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_16 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_15 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_14 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_13 : HIDE 1" >}}   |
|    `el10.aarch64`    |  {{< bg "PGDG 0.5.3" "postgresql_faker_18 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_17 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_16 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_15 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_14 : HIDE 1" >}}   |  {{< bg "PGDG 0.5.3" "postgresql_faker_13 : HIDE 1" >}}   |
|    `d12.x86_64`    |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |
|    `d12.aarch64`    |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |
|    `d13.x86_64`    |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |
|    `d13.aarch64`    |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |
|    `u22.x86_64`    |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |
|    `u22.aarch64`    |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |
|    `u24.x86_64`    |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |
|    `u24.aarch64`    |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |  {{< bg "MISS" "faker : HIDE 0" >}}   |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql_faker_18` | 0.5.3 | `el8.x86_64` | pgdg | 46.0 KiB | [postgresql_faker_18-0.5.3-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/postgresql_faker_18-0.5.3-7PGDG.rhel8.x86_64.rpm) |
| `postgresql_faker_18` | 0.5.3 | `el8.aarch64` | pgdg | 46.1 KiB | [postgresql_faker_18-0.5.3-7PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/postgresql_faker_18-0.5.3-7PGDG.rhel8.aarch64.rpm) |
| `postgresql_faker_18` | 0.5.3 | `el9.x86_64` | pgdg | 44.0 KiB | [postgresql_faker_18-0.5.3-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgresql_faker_18-0.5.3-7PGDG.rhel9.x86_64.rpm) |
| `postgresql_faker_18` | 0.5.3 | `el9.aarch64` | pgdg | 43.8 KiB | [postgresql_faker_18-0.5.3-7PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgresql_faker_18-0.5.3-7PGDG.rhel9.aarch64.rpm) |
| `postgresql_faker_18` | 0.5.3 | `el10.x86_64` | pgdg | 44.3 KiB | [postgresql_faker_18-0.5.3-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgresql_faker_18-0.5.3-7PGDG.rhel10.x86_64.rpm) |
| `postgresql_faker_18` | 0.5.3 | `el10.aarch64` | pgdg | 44.5 KiB | [postgresql_faker_18-0.5.3-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgresql_faker_18-0.5.3-7PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql_faker_17` | 0.5.3 | `el8.x86_64` | pgdg | 45.9 KiB | [postgresql_faker_17-0.5.3-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/postgresql_faker_17-0.5.3-6PGDG.rhel8.x86_64.rpm) |
| `postgresql_faker_17` | 0.5.3 | `el8.aarch64` | pgdg | 46.0 KiB | [postgresql_faker_17-0.5.3-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/postgresql_faker_17-0.5.3-6PGDG.rhel8.aarch64.rpm) |
| `postgresql_faker_17` | 0.5.3 | `el9.x86_64` | pgdg | 44.2 KiB | [postgresql_faker_17-0.5.3-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgresql_faker_17-0.5.3-6PGDG.rhel9.x86_64.rpm) |
| `postgresql_faker_17` | 0.5.3 | `el9.aarch64` | pgdg | 44.1 KiB | [postgresql_faker_17-0.5.3-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgresql_faker_17-0.5.3-6PGDG.rhel9.aarch64.rpm) |
| `postgresql_faker_17` | 0.5.3 | `el10.x86_64` | pgdg | 44.3 KiB | [postgresql_faker_17-0.5.3-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgresql_faker_17-0.5.3-7PGDG.rhel10.x86_64.rpm) |
| `postgresql_faker_17` | 0.5.3 | `el10.aarch64` | pgdg | 44.5 KiB | [postgresql_faker_17-0.5.3-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgresql_faker_17-0.5.3-7PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql_faker_16` | 0.5.3 | `el8.x86_64` | pgdg | 45.4 KiB | [postgresql_faker_16-0.5.3-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/postgresql_faker_16-0.5.3-3PGDG.rhel8.x86_64.rpm) |
| `postgresql_faker_16` | 0.5.3 | `el8.aarch64` | pgdg | 45.6 KiB | [postgresql_faker_16-0.5.3-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/postgresql_faker_16-0.5.3-3PGDG.rhel8.aarch64.rpm) |
| `postgresql_faker_16` | 0.5.3 | `el9.x86_64` | pgdg | 44.1 KiB | [postgresql_faker_16-0.5.3-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgresql_faker_16-0.5.3-3PGDG.rhel9.x86_64.rpm) |
| `postgresql_faker_16` | 0.5.3 | `el9.aarch64` | pgdg | 44.1 KiB | [postgresql_faker_16-0.5.3-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgresql_faker_16-0.5.3-3PGDG.rhel9.aarch64.rpm) |
| `postgresql_faker_16` | 0.5.3 | `el10.x86_64` | pgdg | 44.3 KiB | [postgresql_faker_16-0.5.3-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgresql_faker_16-0.5.3-7PGDG.rhel10.x86_64.rpm) |
| `postgresql_faker_16` | 0.5.3 | `el10.aarch64` | pgdg | 44.5 KiB | [postgresql_faker_16-0.5.3-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgresql_faker_16-0.5.3-7PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql_faker_15` | 0.5.3 | `el8.x86_64` | pgdg | 49.6 KiB | [postgresql_faker_15-0.5.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/postgresql_faker_15-0.5.3-1.rhel8.x86_64.rpm) |
| `postgresql_faker_15` | 0.5.3 | `el8.aarch64` | pgdg | 49.8 KiB | [postgresql_faker_15-0.5.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/postgresql_faker_15-0.5.3-1.rhel8.aarch64.rpm) |
| `postgresql_faker_15` | 0.5.3 | `el9.x86_64` | pgdg | 48.6 KiB | [postgresql_faker_15-0.5.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/postgresql_faker_15-0.5.3-1.rhel9.x86_64.rpm) |
| `postgresql_faker_15` | 0.5.3 | `el9.aarch64` | pgdg | 48.5 KiB | [postgresql_faker_15-0.5.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/postgresql_faker_15-0.5.3-1.rhel9.aarch64.rpm) |
| `postgresql_faker_15` | 0.5.3 | `el10.x86_64` | pgdg | 44.3 KiB | [postgresql_faker_15-0.5.3-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/postgresql_faker_15-0.5.3-7PGDG.rhel10.x86_64.rpm) |
| `postgresql_faker_15` | 0.5.3 | `el10.aarch64` | pgdg | 44.5 KiB | [postgresql_faker_15-0.5.3-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/postgresql_faker_15-0.5.3-7PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql_faker_14` | 0.5.3 | `el8.x86_64` | pgdg | 49.9 KiB | [postgresql_faker_14-0.5.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgresql_faker_14-0.5.3-1.rhel8.x86_64.rpm) |
| `postgresql_faker_14` | 0.4.0 | `el8.x86_64` | pgdg | 37.7 KiB | [postgresql_faker_14-0.4.0-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/postgresql_faker_14-0.4.0-1.rhel8.noarch.rpm) |
| `postgresql_faker_14` | 0.5.3 | `el8.aarch64` | pgdg | 49.8 KiB | [postgresql_faker_14-0.5.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/postgresql_faker_14-0.5.3-1.rhel8.aarch64.rpm) |
| `postgresql_faker_14` | 0.5.3 | `el9.x86_64` | pgdg | 47.9 KiB | [postgresql_faker_14-0.5.3-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresql_faker_14-0.5.3-1.rhel9.noarch.rpm) |
| `postgresql_faker_14` | 0.5.3 | `el9.x86_64` | pgdg | 48.6 KiB | [postgresql_faker_14-0.5.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/postgresql_faker_14-0.5.3-1.rhel9.x86_64.rpm) |
| `postgresql_faker_14` | 0.5.3 | `el9.aarch64` | pgdg | 48.5 KiB | [postgresql_faker_14-0.5.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/postgresql_faker_14-0.5.3-1.rhel9.aarch64.rpm) |
| `postgresql_faker_14` | 0.5.3 | `el10.x86_64` | pgdg | 44.3 KiB | [postgresql_faker_14-0.5.3-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/postgresql_faker_14-0.5.3-7PGDG.rhel10.x86_64.rpm) |
| `postgresql_faker_14` | 0.5.3 | `el10.aarch64` | pgdg | 44.5 KiB | [postgresql_faker_14-0.5.3-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/postgresql_faker_14-0.5.3-7PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql_faker_13` | 0.5.3 | `el8.x86_64` | pgdg | 49.9 KiB | [postgresql_faker_13-0.5.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/postgresql_faker_13-0.5.3-1.rhel8.x86_64.rpm) |
| `postgresql_faker_13` | 0.4.0 | `el8.x86_64` | pgdg | 37.6 KiB | [postgresql_faker_13-0.4.0-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/postgresql_faker_13-0.4.0-1.rhel8.noarch.rpm) |
| `postgresql_faker_13` | 0.3.0 | `el8.x86_64` | pgdg | 14.0 KiB | [postgresql_faker_13-0.3.0-1.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/postgresql_faker_13-0.3.0-1.rhel8.noarch.rpm) |
| `postgresql_faker_13` | 0.5.3 | `el8.aarch64` | pgdg | 49.8 KiB | [postgresql_faker_13-0.5.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/postgresql_faker_13-0.5.3-1.rhel8.aarch64.rpm) |
| `postgresql_faker_13` | 0.5.3 | `el9.x86_64` | pgdg | 47.9 KiB | [postgresql_faker_13-0.5.3-1.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/postgresql_faker_13-0.5.3-1.rhel9.noarch.rpm) |
| `postgresql_faker_13` | 0.5.3 | `el9.x86_64` | pgdg | 48.6 KiB | [postgresql_faker_13-0.5.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/postgresql_faker_13-0.5.3-1.rhel9.x86_64.rpm) |
| `postgresql_faker_13` | 0.5.3 | `el9.aarch64` | pgdg | 48.5 KiB | [postgresql_faker_13-0.5.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/postgresql_faker_13-0.5.3-1.rhel9.aarch64.rpm) |
| `postgresql_faker_13` | 0.5.3 | `el10.x86_64` | pgdg | 44.3 KiB | [postgresql_faker_13-0.5.3-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/postgresql_faker_13-0.5.3-7PGDG.rhel10.x86_64.rpm) |
| `postgresql_faker_13` | 0.5.3 | `el10.aarch64` | pgdg | 44.5 KiB | [postgresql_faker_13-0.5.3-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/postgresql_faker_13-0.5.3-7PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/anpandu/postgresql_faker" title="Repository" icon="github" subtitle="github.com/anpandu/postgresql_faker" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install faker; # install by extension name, for the current active PG version
pig ext install faker; # install via package alias, for the active PG version
pig ext install faker -v 18;   # install for PG 18
pig ext install faker -v 17;   # install for PG 17
pig ext install faker -v 16;   # install for PG 16
pig ext install faker -v 15;   # install for PG 15
pig ext install faker -v 14;   # install for PG 14
pig ext install faker -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION faker;
```

