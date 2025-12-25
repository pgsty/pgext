---
title: "jsquery"
linkTitle: "jsquery"
description: "data type for jsonb inspection"
weight: 2810
categories: ["FEAT"]
width: full
---

[**jsquery**](https://github.com/postgrespro/jsquery) : data type for jsonb inspection


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2810** | {{< badge content="jsquery" link="https://github.com/postgrespro/jsquery" >}} | {{< ext "jsquery" >}} | `1.2` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "plv8" >}} {{< ext "jsonb_plperl" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "pg_net" >}} {{< ext "pg_summarize" >}} {{< ext "age" >}} |

> [!Note] no pg13,12 on el9


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `jsquery` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2` | {{< bg "18" "jsquery_18" "green" >}} {{< bg "17" "jsquery_17" "green" >}} {{< bg "16" "jsquery_16" "green" >}} {{< bg "15" "jsquery_15" "green" >}} {{< bg "14" "jsquery_14" "green" >}} {{< bg "13" "jsquery_13" "green" >}} | `jsquery_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2` | {{< bg "18" "postgresql-18-jsquery" "green" >}} {{< bg "17" "postgresql-17-jsquery" "green" >}} {{< bg "16" "postgresql-16-jsquery" "green" >}} {{< bg "15" "postgresql-15-jsquery" "green" >}} {{< bg "14" "postgresql-14-jsquery" "green" >}} {{< bg "13" "postgresql-13-jsquery" "green" >}} | `postgresql-$v-jsquery` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.2" "jsquery_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.1" "jsquery_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.2" "jsquery_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.1" "jsquery_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.2" "jsquery_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "jsquery_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.2" "jsquery_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.1" "jsquery_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.2" "jsquery_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.2" "jsquery_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "jsquery_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.2" "postgresql-18-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-17-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-16-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-15-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-14-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-13-jsquery : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.2" "postgresql-18-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-17-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-16-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-15-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-14-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-13-jsquery : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.2" "postgresql-18-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-17-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-16-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-15-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-14-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-13-jsquery : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.2" "postgresql-18-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-17-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-16-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-15-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-14-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-13-jsquery : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.2" "postgresql-18-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-17-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-16-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-15-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-14-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-13-jsquery : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.2" "postgresql-18-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-17-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-16-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-15-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-14-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-13-jsquery : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.2" "postgresql-18-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-17-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-16-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-15-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-14-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-13-jsquery : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.2" "postgresql-18-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-17-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-16-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-15-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-14-jsquery : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2" "postgresql-13-jsquery : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jsquery_18` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.3 KiB | [jsquery_18-1.2-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/jsquery_18-1.2-4PGDG.rhel8.x86_64.rpm) |
| `jsquery_18` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.6 KiB | [jsquery_18-1.2-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/jsquery_18-1.2-4PGDG.rhel8.aarch64.rpm) |
| `jsquery_18` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.8 KiB | [jsquery_18-1.2-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/jsquery_18-1.2-4PGDG.rhel9.x86_64.rpm) |
| `jsquery_18` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 47.5 KiB | [jsquery_18-1.2-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/jsquery_18-1.2-4PGDG.rhel9.aarch64.rpm) |
| `jsquery_18` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.0 KiB | [jsquery_18-1.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/jsquery_18-1.2-4PGDG.rhel10.x86_64.rpm) |
| `jsquery_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 48.6 KiB | [jsquery_18-1.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/jsquery_18-1.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-jsquery` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 123.1 KiB | [postgresql-18-jsquery_1.2-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-18-jsquery_1.2-3.pgdg12+1_amd64.deb) |
| `postgresql-18-jsquery` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 120.1 KiB | [postgresql-18-jsquery_1.2-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-18-jsquery_1.2-3.pgdg12+1_arm64.deb) |
| `postgresql-18-jsquery` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 123.4 KiB | [postgresql-18-jsquery_1.2-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-18-jsquery_1.2-3.pgdg13+1_amd64.deb) |
| `postgresql-18-jsquery` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 120.5 KiB | [postgresql-18-jsquery_1.2-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-18-jsquery_1.2-3.pgdg13+1_arm64.deb) |
| `postgresql-18-jsquery` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 123.8 KiB | [postgresql-18-jsquery_1.2-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-18-jsquery_1.2-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-jsquery` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 120.5 KiB | [postgresql-18-jsquery_1.2-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-18-jsquery_1.2-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-jsquery` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 122.2 KiB | [postgresql-18-jsquery_1.2-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-18-jsquery_1.2-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-jsquery` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 119.9 KiB | [postgresql-18-jsquery_1.2-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-18-jsquery_1.2-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jsquery_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.1 KiB | [jsquery_17-1.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/jsquery_17-1.2-2PGDG.rhel8.x86_64.rpm) |
| `jsquery_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.6 KiB | [jsquery_17-1.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/jsquery_17-1.2-2PGDG.rhel8.aarch64.rpm) |
| `jsquery_17` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.9 KiB | [jsquery_17-1.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/jsquery_17-1.2-2PGDG.rhel9.x86_64.rpm) |
| `jsquery_17` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 47.6 KiB | [jsquery_17-1.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/jsquery_17-1.2-2PGDG.rhel9.aarch64.rpm) |
| `jsquery_17` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.0 KiB | [jsquery_17-1.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/jsquery_17-1.2-4PGDG.rhel10.x86_64.rpm) |
| `jsquery_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 48.6 KiB | [jsquery_17-1.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/jsquery_17-1.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-jsquery` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 123.1 KiB | [postgresql-17-jsquery_1.2-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-17-jsquery_1.2-3.pgdg12+1_amd64.deb) |
| `postgresql-17-jsquery` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 120.0 KiB | [postgresql-17-jsquery_1.2-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-17-jsquery_1.2-3.pgdg12+1_arm64.deb) |
| `postgresql-17-jsquery` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 123.4 KiB | [postgresql-17-jsquery_1.2-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-17-jsquery_1.2-3.pgdg13+1_amd64.deb) |
| `postgresql-17-jsquery` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 120.3 KiB | [postgresql-17-jsquery_1.2-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-17-jsquery_1.2-3.pgdg13+1_arm64.deb) |
| `postgresql-17-jsquery` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 130.8 KiB | [postgresql-17-jsquery_1.2-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-17-jsquery_1.2-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-jsquery` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 127.8 KiB | [postgresql-17-jsquery_1.2-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-17-jsquery_1.2-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-jsquery` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 122.3 KiB | [postgresql-17-jsquery_1.2-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-17-jsquery_1.2-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-jsquery` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 119.8 KiB | [postgresql-17-jsquery_1.2-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-17-jsquery_1.2-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jsquery_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.0 KiB | [jsquery_16-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/jsquery_16-1.2-1PGDG.rhel8.x86_64.rpm) |
| `jsquery_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.3 KiB | [jsquery_16-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/jsquery_16-1.2-1PGDG.rhel8.aarch64.rpm) |
| `jsquery_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.5 KiB | [jsquery_16-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/jsquery_16-1.2-1PGDG.rhel9.x86_64.rpm) |
| `jsquery_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 47.0 KiB | [jsquery_16-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/jsquery_16-1.2-1PGDG.rhel9.aarch64.rpm) |
| `jsquery_16` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 49.9 KiB | [jsquery_16-1.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/jsquery_16-1.2-4PGDG.rhel10.x86_64.rpm) |
| `jsquery_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 48.4 KiB | [jsquery_16-1.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/jsquery_16-1.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-jsquery` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 122.8 KiB | [postgresql-16-jsquery_1.2-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-16-jsquery_1.2-3.pgdg12+1_amd64.deb) |
| `postgresql-16-jsquery` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 119.7 KiB | [postgresql-16-jsquery_1.2-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-16-jsquery_1.2-3.pgdg12+1_arm64.deb) |
| `postgresql-16-jsquery` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 123.1 KiB | [postgresql-16-jsquery_1.2-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-16-jsquery_1.2-3.pgdg13+1_amd64.deb) |
| `postgresql-16-jsquery` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 120.3 KiB | [postgresql-16-jsquery_1.2-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-16-jsquery_1.2-3.pgdg13+1_arm64.deb) |
| `postgresql-16-jsquery` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 130.6 KiB | [postgresql-16-jsquery_1.2-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-16-jsquery_1.2-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-jsquery` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 127.6 KiB | [postgresql-16-jsquery_1.2-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-16-jsquery_1.2-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-jsquery` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 122.1 KiB | [postgresql-16-jsquery_1.2-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-16-jsquery_1.2-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-jsquery` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 119.5 KiB | [postgresql-16-jsquery_1.2-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-16-jsquery_1.2-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jsquery_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.3 KiB | [jsquery_15-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/jsquery_15-1.2-1PGDG.rhel8.x86_64.rpm) |
| `jsquery_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.6 KiB | [jsquery_15-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/jsquery_15-1.2-1PGDG.rhel8.aarch64.rpm) |
| `jsquery_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.3 KiB | [jsquery_15-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/jsquery_15-1.2-1PGDG.rhel9.x86_64.rpm) |
| `jsquery_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.6 KiB | [jsquery_15-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/jsquery_15-1.2-1PGDG.rhel9.aarch64.rpm) |
| `jsquery_15` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.5 KiB | [jsquery_15-1.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/jsquery_15-1.2-4PGDG.rhel10.x86_64.rpm) |
| `jsquery_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.7 KiB | [jsquery_15-1.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/jsquery_15-1.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-jsquery` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.3 KiB | [postgresql-15-jsquery_1.2-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-15-jsquery_1.2-3.pgdg12+1_amd64.deb) |
| `postgresql-15-jsquery` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 120.9 KiB | [postgresql-15-jsquery_1.2-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-15-jsquery_1.2-3.pgdg12+1_arm64.deb) |
| `postgresql-15-jsquery` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.5 KiB | [postgresql-15-jsquery_1.2-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-15-jsquery_1.2-3.pgdg13+1_amd64.deb) |
| `postgresql-15-jsquery` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 121.5 KiB | [postgresql-15-jsquery_1.2-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-15-jsquery_1.2-3.pgdg13+1_arm64.deb) |
| `postgresql-15-jsquery` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 132.6 KiB | [postgresql-15-jsquery_1.2-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-15-jsquery_1.2-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-jsquery` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 129.5 KiB | [postgresql-15-jsquery_1.2-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-15-jsquery_1.2-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-jsquery` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.2 KiB | [postgresql-15-jsquery_1.2-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-15-jsquery_1.2-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-jsquery` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 121.2 KiB | [postgresql-15-jsquery_1.2-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-15-jsquery_1.2-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jsquery_14` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.3 KiB | [jsquery_14-1.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/jsquery_14-1.2-1PGDG.rhel8.x86_64.rpm) |
| `jsquery_14` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.6 KiB | [jsquery_14-1.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/jsquery_14-1.2-1PGDG.rhel8.aarch64.rpm) |
| `jsquery_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.3 KiB | [jsquery_14-1.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/jsquery_14-1.2-1PGDG.rhel9.x86_64.rpm) |
| `jsquery_14` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.7 KiB | [jsquery_14-1.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/jsquery_14-1.2-1PGDG.rhel9.aarch64.rpm) |
| `jsquery_14` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.5 KiB | [jsquery_14-1.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/jsquery_14-1.2-4PGDG.rhel10.x86_64.rpm) |
| `jsquery_14` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.7 KiB | [jsquery_14-1.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/jsquery_14-1.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-jsquery` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.3 KiB | [postgresql-14-jsquery_1.2-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-14-jsquery_1.2-3.pgdg12+1_amd64.deb) |
| `postgresql-14-jsquery` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 121.0 KiB | [postgresql-14-jsquery_1.2-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-14-jsquery_1.2-3.pgdg12+1_arm64.deb) |
| `postgresql-14-jsquery` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.3 KiB | [postgresql-14-jsquery_1.2-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-14-jsquery_1.2-3.pgdg13+1_amd64.deb) |
| `postgresql-14-jsquery` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 121.3 KiB | [postgresql-14-jsquery_1.2-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-14-jsquery_1.2-3.pgdg13+1_arm64.deb) |
| `postgresql-14-jsquery` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 132.6 KiB | [postgresql-14-jsquery_1.2-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-14-jsquery_1.2-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-jsquery` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 129.3 KiB | [postgresql-14-jsquery_1.2-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-14-jsquery_1.2-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-jsquery` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.2 KiB | [postgresql-14-jsquery_1.2-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-14-jsquery_1.2-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-jsquery` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 121.3 KiB | [postgresql-14-jsquery_1.2-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-14-jsquery_1.2-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jsquery_13` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 138.3 KiB | [jsquery_13-1.1.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/jsquery_13-1.1.1-2.rhel8.x86_64.rpm) |
| `jsquery_13` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 136.7 KiB | [jsquery_13-1.1.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/jsquery_13-1.1.1-2.rhel8.aarch64.rpm) |
| `jsquery_13` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 140.7 KiB | [jsquery_13-1.1.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/jsquery_13-1.1.1-2.rhel9.aarch64.rpm) |
| `jsquery_13` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.5 KiB | [jsquery_13-1.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/jsquery_13-1.2-4PGDG.rhel10.x86_64.rpm) |
| `jsquery_13` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.7 KiB | [jsquery_13-1.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/jsquery_13-1.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-jsquery` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 124.1 KiB | [postgresql-13-jsquery_1.2-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-13-jsquery_1.2-3.pgdg12+1_amd64.deb) |
| `postgresql-13-jsquery` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 120.8 KiB | [postgresql-13-jsquery_1.2-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-13-jsquery_1.2-3.pgdg12+1_arm64.deb) |
| `postgresql-13-jsquery` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 124.3 KiB | [postgresql-13-jsquery_1.2-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-13-jsquery_1.2-3.pgdg13+1_amd64.deb) |
| `postgresql-13-jsquery` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 121.1 KiB | [postgresql-13-jsquery_1.2-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-13-jsquery_1.2-3.pgdg13+1_arm64.deb) |
| `postgresql-13-jsquery` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 132.4 KiB | [postgresql-13-jsquery_1.2-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-13-jsquery_1.2-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-jsquery` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 129.1 KiB | [postgresql-13-jsquery_1.2-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-13-jsquery_1.2-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-jsquery` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 124.0 KiB | [postgresql-13-jsquery_1.2-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-13-jsquery_1.2-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-jsquery` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 121.2 KiB | [postgresql-13-jsquery_1.2-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/j/jsquery/postgresql-13-jsquery_1.2-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgrespro/jsquery" title="Repository" icon="github" subtitle="github.com/postgrespro/jsquery" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="jsquery-1.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg jsquery;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install jsquery;		# install via package name, for the active PG version

pig install jsquery -v 18;   # install for PG 18
pig install jsquery -v 17;   # install for PG 17
pig install jsquery -v 16;   # install for PG 16
pig install jsquery -v 15;   # install for PG 15
pig install jsquery -v 14;   # install for PG 14
pig install jsquery -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION jsquery;
```
