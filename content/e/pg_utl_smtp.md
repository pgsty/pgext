---
title: "pg_utl_smtp"
linkTitle: "pg_utl_smtp"
description: "Oracle UTL_SMTP compatibility extension for PostgreSQL"
weight: 9290
categories: ["SIM"]
width: full
---

[**pg_utl_smtp**](https://github.com/hexacluster/pg_utl_smtp) : Oracle UTL_SMTP compatibility extension for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9290** | {{< badge content="pg_utl_smtp" link="https://github.com/hexacluster/pg_utl_smtp" >}} | {{< ext "pg_utl_smtp" >}} | `1.0.0` | {{< category "SIM" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `utl_smtp` |
|   **Requires**    | {{< ext "plperlu" >}} |

> [!Note] runtime requires plperlu and Perl Net::SMTP but not defined in control


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pg_utl_smtp` | `plperlu` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0` | {{< bg "18" "pg_utl_smtp_18" "green" >}} {{< bg "17" "pg_utl_smtp_17" "green" >}} {{< bg "16" "pg_utl_smtp_16" "green" >}} {{< bg "15" "pg_utl_smtp_15" "green" >}} {{< bg "14" "pg_utl_smtp_14" "green" >}} {{< bg "13" "pg_utl_smtp_13" "red" >}} | `pg_utl_smtp_$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_utl_smtp_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_utl_smtp_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_utl_smtp_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_utl_smtp_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_utl_smtp_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_utl_smtp_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_utl_smtp_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_utl_smtp : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_utl_smtp_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.5 KiB | [pg_utl_smtp_18-1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_utl_smtp_18-1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pg_utl_smtp_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.4 KiB | [pg_utl_smtp_18-1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_utl_smtp_18-1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pg_utl_smtp_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.4 KiB | [pg_utl_smtp_18-1.0-2PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_utl_smtp_18-1.0-2PGDG.rhel9.7.noarch.rpm) |
| `pg_utl_smtp_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.4 KiB | [pg_utl_smtp_18-1.0-2PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_utl_smtp_18-1.0-2PGDG.rhel9.7.noarch.rpm) |
| `pg_utl_smtp_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.6 KiB | [pg_utl_smtp_18-1.0-2PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_utl_smtp_18-1.0-2PGDG.rhel10.1.noarch.rpm) |
| `pg_utl_smtp_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.5 KiB | [pg_utl_smtp_18-1.0-2PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_utl_smtp_18-1.0-2PGDG.rhel10.1.noarch.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_utl_smtp_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.5 KiB | [pg_utl_smtp_17-1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_utl_smtp_17-1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pg_utl_smtp_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.4 KiB | [pg_utl_smtp_17-1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_utl_smtp_17-1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pg_utl_smtp_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.4 KiB | [pg_utl_smtp_17-1.0-2PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_utl_smtp_17-1.0-2PGDG.rhel9.7.noarch.rpm) |
| `pg_utl_smtp_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.4 KiB | [pg_utl_smtp_17-1.0-2PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_utl_smtp_17-1.0-2PGDG.rhel9.7.noarch.rpm) |
| `pg_utl_smtp_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.6 KiB | [pg_utl_smtp_17-1.0-2PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_utl_smtp_17-1.0-2PGDG.rhel10.1.noarch.rpm) |
| `pg_utl_smtp_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.5 KiB | [pg_utl_smtp_17-1.0-2PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_utl_smtp_17-1.0-2PGDG.rhel10.1.noarch.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_utl_smtp_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.5 KiB | [pg_utl_smtp_16-1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_utl_smtp_16-1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pg_utl_smtp_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.4 KiB | [pg_utl_smtp_16-1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_utl_smtp_16-1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pg_utl_smtp_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.4 KiB | [pg_utl_smtp_16-1.0-2PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_utl_smtp_16-1.0-2PGDG.rhel9.7.noarch.rpm) |
| `pg_utl_smtp_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.4 KiB | [pg_utl_smtp_16-1.0-2PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_utl_smtp_16-1.0-2PGDG.rhel9.7.noarch.rpm) |
| `pg_utl_smtp_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.6 KiB | [pg_utl_smtp_16-1.0-2PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_utl_smtp_16-1.0-2PGDG.rhel10.1.noarch.rpm) |
| `pg_utl_smtp_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.5 KiB | [pg_utl_smtp_16-1.0-2PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_utl_smtp_16-1.0-2PGDG.rhel10.1.noarch.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_utl_smtp_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.5 KiB | [pg_utl_smtp_15-1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_utl_smtp_15-1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pg_utl_smtp_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.4 KiB | [pg_utl_smtp_15-1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_utl_smtp_15-1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pg_utl_smtp_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.4 KiB | [pg_utl_smtp_15-1.0-2PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_utl_smtp_15-1.0-2PGDG.rhel9.7.noarch.rpm) |
| `pg_utl_smtp_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.4 KiB | [pg_utl_smtp_15-1.0-2PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_utl_smtp_15-1.0-2PGDG.rhel9.7.noarch.rpm) |
| `pg_utl_smtp_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.6 KiB | [pg_utl_smtp_15-1.0-2PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_utl_smtp_15-1.0-2PGDG.rhel10.1.noarch.rpm) |
| `pg_utl_smtp_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.5 KiB | [pg_utl_smtp_15-1.0-2PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_utl_smtp_15-1.0-2PGDG.rhel10.1.noarch.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_utl_smtp_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 12.4 KiB | [pg_utl_smtp_14-1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_utl_smtp_14-1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pg_utl_smtp_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 12.4 KiB | [pg_utl_smtp_14-1.0-2PGDG.rhel8.10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_utl_smtp_14-1.0-2PGDG.rhel8.10.noarch.rpm) |
| `pg_utl_smtp_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 12.4 KiB | [pg_utl_smtp_14-1.0-2PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_utl_smtp_14-1.0-2PGDG.rhel9.7.noarch.rpm) |
| `pg_utl_smtp_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 12.4 KiB | [pg_utl_smtp_14-1.0-2PGDG.rhel9.7.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_utl_smtp_14-1.0-2PGDG.rhel9.7.noarch.rpm) |
| `pg_utl_smtp_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 12.6 KiB | [pg_utl_smtp_14-1.0-2PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_utl_smtp_14-1.0-2PGDG.rhel10.1.noarch.rpm) |
| `pg_utl_smtp_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 12.5 KiB | [pg_utl_smtp_14-1.0-2PGDG.rhel10.1.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_utl_smtp_14-1.0-2PGDG.rhel10.1.noarch.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/hexacluster/pg_utl_smtp" title="Repository" icon="github" subtitle="github.com/hexacluster/pg_utl_smtp" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_utl_smtp;		# install via package name, for the active PG version

pig install pg_utl_smtp -v 18;   # install for PG 18
pig install pg_utl_smtp -v 17;   # install for PG 17
pig install pg_utl_smtp -v 16;   # install for PG 16
pig install pg_utl_smtp -v 15;   # install for PG 15
pig install pg_utl_smtp -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_utl_smtp CASCADE; -- requires plperlu
```
