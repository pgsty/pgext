---
title: "prioritize"
linkTitle: "prioritize"
description: "get and set the priority of PostgreSQL backends"
weight: 5100
categories: ["ADMIN"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**pg_prioritize**](https://github.com/schmiddy/pg_prioritize) : get and set the priority of PostgreSQL backends


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5100** | {{< badge content="prioritize" link="https://github.com/schmiddy/pg_prioritize" >}} | {{< ext "prioritize" "pg_prioritize" >}} | `1.0.4` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plan_filter" >}} {{< ext "pg_kpart" >}} {{< ext "pg_readonly" >}} {{< ext "qos" >}} {{< ext "block_copy_command" >}} {{< ext "safeupdate" >}} {{< ext "pg_command_fw" >}} {{< ext "pg_strict" >}} {{< ext "pg_hint_plan" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_prioritize` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.4` | {{< bg "18" "pg_prioritize_18" "green" >}} {{< bg "17" "pg_prioritize_17" "green" >}} {{< bg "16" "pg_prioritize_16" "green" >}} {{< bg "15" "pg_prioritize_15" "green" >}} {{< bg "14" "pg_prioritize_14" "green" >}} | `pg_prioritize_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.4` | {{< bg "18" "postgresql-18-prioritize" "green" >}} {{< bg "17" "postgresql-17-prioritize" "green" >}} {{< bg "16" "postgresql-16-prioritize" "green" >}} {{< bg "15" "postgresql-15-prioritize" "green" >}} {{< bg "14" "postgresql-14-prioritize" "green" >}} | `postgresql-$v-prioritize` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.0.4" "pg_prioritize_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-prioritize : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-prioritize : AVAIL 1" "blue" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_prioritize_18` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.4 KiB | [pg_prioritize_18-1.0.4-7PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_prioritize_18-1.0.4-7PGDG.rhel8.x86_64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.3 KiB | [pg_prioritize_18-1.0.4-7PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_prioritize_18-1.0.4-7PGDG.rhel8.aarch64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.2 KiB | [pg_prioritize_18-1.0.4-9PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_prioritize_18-1.0.4-9PGDG.rhel9.8.x86_64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.1 KiB | [pg_prioritize_18-1.0.4-7PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_prioritize_18-1.0.4-7PGDG.rhel9.x86_64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.8 KiB | [pg_prioritize_18-1.0.4-9PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_prioritize_18-1.0.4-9PGDG.rhel9.8.aarch64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.6 KiB | [pg_prioritize_18-1.0.4-7PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_prioritize_18-1.0.4-7PGDG.rhel9.aarch64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.2 KiB | [pg_prioritize_18-1.0.4-9PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_prioritize_18-1.0.4-9PGDG.rhel10.2.x86_64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.4 KiB | [pg_prioritize_18-1.0.4-7PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_prioritize_18-1.0.4-7PGDG.rhel10.x86_64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.1 KiB | [pg_prioritize_18-1.0.4-9PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_prioritize_18-1.0.4-9PGDG.rhel10.2.aarch64.rpm) |
| `pg_prioritize_18` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.2 KiB | [pg_prioritize_18-1.0.4-7PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_prioritize_18-1.0.4-7PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-prioritize` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.7 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg12+1_amd64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.6 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg12+1_arm64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.7 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg13+1_amd64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.7 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg13+1_arm64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.3 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.1 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 11.8 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 11.7 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.0 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg26.04+1_amd64.deb) |
| `postgresql-18-prioritize` | `1.0.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.0 KiB | [postgresql-18-prioritize_1.0.4-13.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-18-prioritize_1.0.4-13.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_prioritize_17` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.1 KiB | [pg_prioritize_17-1.0.4-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_prioritize_17-1.0.4-5PGDG.rhel8.x86_64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 14.0 KiB | [pg_prioritize_17-1.0.4-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_prioritize_17-1.0.4-5PGDG.rhel8.aarch64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.2 KiB | [pg_prioritize_17-1.0.4-9PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_prioritize_17-1.0.4-9PGDG.rhel9.8.x86_64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.0 KiB | [pg_prioritize_17-1.0.4-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_prioritize_17-1.0.4-5PGDG.rhel9.x86_64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.8 KiB | [pg_prioritize_17-1.0.4-9PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_prioritize_17-1.0.4-9PGDG.rhel9.8.aarch64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.7 KiB | [pg_prioritize_17-1.0.4-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_prioritize_17-1.0.4-5PGDG.rhel9.aarch64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.2 KiB | [pg_prioritize_17-1.0.4-9PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_prioritize_17-1.0.4-9PGDG.rhel10.2.x86_64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.3 KiB | [pg_prioritize_17-1.0.4-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_prioritize_17-1.0.4-6PGDG.rhel10.x86_64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.1 KiB | [pg_prioritize_17-1.0.4-9PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_prioritize_17-1.0.4-9PGDG.rhel10.2.aarch64.rpm) |
| `pg_prioritize_17` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.2 KiB | [pg_prioritize_17-1.0.4-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_prioritize_17-1.0.4-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-prioritize` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.7 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg12+1_amd64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.6 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg12+1_arm64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.7 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg13+1_amd64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.7 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg13+1_arm64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.6 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.4 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 11.8 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 11.7 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.0 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg26.04+1_amd64.deb) |
| `postgresql-17-prioritize` | `1.0.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 11.9 KiB | [postgresql-17-prioritize_1.0.4-13.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-17-prioritize_1.0.4-13.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_prioritize_16` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 14.0 KiB | [pg_prioritize_16-1.0.4-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_prioritize_16-1.0.4-4PGDG.rhel8.x86_64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 13.9 KiB | [pg_prioritize_16-1.0.4-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_prioritize_16-1.0.4-4PGDG.rhel8.aarch64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.2 KiB | [pg_prioritize_16-1.0.4-9PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_prioritize_16-1.0.4-9PGDG.rhel9.8.x86_64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 13.8 KiB | [pg_prioritize_16-1.0.4-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_prioritize_16-1.0.4-4PGDG.rhel9.x86_64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.8 KiB | [pg_prioritize_16-1.0.4-9PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_prioritize_16-1.0.4-9PGDG.rhel9.8.aarch64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.3 KiB | [pg_prioritize_16-1.0.4-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_prioritize_16-1.0.4-4PGDG.rhel9.aarch64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.2 KiB | [pg_prioritize_16-1.0.4-9PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_prioritize_16-1.0.4-9PGDG.rhel10.2.x86_64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.3 KiB | [pg_prioritize_16-1.0.4-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_prioritize_16-1.0.4-6PGDG.rhel10.x86_64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.1 KiB | [pg_prioritize_16-1.0.4-9PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_prioritize_16-1.0.4-9PGDG.rhel10.2.aarch64.rpm) |
| `pg_prioritize_16` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.2 KiB | [pg_prioritize_16-1.0.4-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_prioritize_16-1.0.4-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-prioritize` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.7 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg12+1_amd64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.6 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg12+1_arm64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.7 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg13+1_amd64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.7 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg13+1_arm64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.6 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.3 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 11.8 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 11.7 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.0 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg26.04+1_amd64.deb) |
| `postgresql-16-prioritize` | `1.0.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 11.9 KiB | [postgresql-16-prioritize_1.0.4-13.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-16-prioritize_1.0.4-13.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_prioritize_15` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.2 KiB | [pg_prioritize_15-1.0.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_prioritize_15-1.0.4-2.rhel8.x86_64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.3 KiB | [pg_prioritize_15-1.0.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_prioritize_15-1.0.4-2.rhel8.aarch64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.2 KiB | [pg_prioritize_15-1.0.4-9PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_prioritize_15-1.0.4-9PGDG.rhel9.8.x86_64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.6 KiB | [pg_prioritize_15-1.0.4-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_prioritize_15-1.0.4-2.rhel9.x86_64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.8 KiB | [pg_prioritize_15-1.0.4-9PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_prioritize_15-1.0.4-9PGDG.rhel9.8.aarch64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.3 KiB | [pg_prioritize_15-1.0.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_prioritize_15-1.0.4-2.rhel9.aarch64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.2 KiB | [pg_prioritize_15-1.0.4-9PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_prioritize_15-1.0.4-9PGDG.rhel10.2.x86_64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.3 KiB | [pg_prioritize_15-1.0.4-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_prioritize_15-1.0.4-6PGDG.rhel10.x86_64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.1 KiB | [pg_prioritize_15-1.0.4-9PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_prioritize_15-1.0.4-9PGDG.rhel10.2.aarch64.rpm) |
| `pg_prioritize_15` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.2 KiB | [pg_prioritize_15-1.0.4-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_prioritize_15-1.0.4-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-prioritize` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.7 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg12+1_amd64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.6 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg12+1_arm64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.7 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg13+1_amd64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.7 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg13+1_arm64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.6 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.4 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 11.8 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 11.7 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 12.0 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg26.04+1_amd64.deb) |
| `postgresql-15-prioritize` | `1.0.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 12.0 KiB | [postgresql-15-prioritize_1.0.4-13.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-15-prioritize_1.0.4-13.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_prioritize_14` | `1.0.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.0 KiB | [pg_prioritize_14-1.0.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_prioritize_14-1.0.4-2.rhel8.x86_64.rpm) |
| `pg_prioritize_14` | `1.0.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.3 KiB | [pg_prioritize_14-1.0.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_prioritize_14-1.0.4-2.rhel8.aarch64.rpm) |
| `pg_prioritize_14` | `1.0.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 14.1 KiB | [pg_prioritize_14-1.0.4-9PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_prioritize_14-1.0.4-9PGDG.rhel9.8.x86_64.rpm) |
| `pg_prioritize_14` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 13.8 KiB | [pg_prioritize_14-1.0.4-9PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_prioritize_14-1.0.4-9PGDG.rhel9.8.aarch64.rpm) |
| `pg_prioritize_14` | `1.0.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.3 KiB | [pg_prioritize_14-1.0.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_prioritize_14-1.0.4-2.rhel9.aarch64.rpm) |
| `pg_prioritize_14` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.2 KiB | [pg_prioritize_14-1.0.4-9PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_prioritize_14-1.0.4-9PGDG.rhel10.2.x86_64.rpm) |
| `pg_prioritize_14` | `1.0.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 14.3 KiB | [pg_prioritize_14-1.0.4-6PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_prioritize_14-1.0.4-6PGDG.rhel10.x86_64.rpm) |
| `pg_prioritize_14` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.1 KiB | [pg_prioritize_14-1.0.4-9PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_prioritize_14-1.0.4-9PGDG.rhel10.2.aarch64.rpm) |
| `pg_prioritize_14` | `1.0.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 14.2 KiB | [pg_prioritize_14-1.0.4-6PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_prioritize_14-1.0.4-6PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-prioritize` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 11.6 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg12+1_amd64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 11.6 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg12+1_arm64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 11.6 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg13+1_amd64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 11.7 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg13+1_arm64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 12.6 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg22.04+1_amd64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 12.3 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg22.04+1_arm64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 11.8 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg24.04+1_amd64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 11.7 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg24.04+1_arm64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 11.9 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg26.04+1_amd64.deb) |
| `postgresql-14-prioritize` | `1.0.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 11.9 KiB | [postgresql-14-prioritize_1.0.4-13.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-prioritize/postgresql-14-prioritize_1.0.4-13.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/schmiddy/pg_prioritize" title="Repository" icon="github" subtitle="github.com/schmiddy/pg_prioritize" />}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_prioritize;		# install via package name, for the active PG version
pig install prioritize;		# install by extension name, for the current active PG version

pig install prioritize -v 18;   # install for PG 18
pig install prioritize -v 17;   # install for PG 17
pig install prioritize -v 16;   # install for PG 16
pig install prioritize -v 15;   # install for PG 15
pig install prioritize -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION prioritize;
```

## Usage

Sources:

- [Official README](https://github.com/schmiddy/pg_prioritize/blob/c160271202ca8a713dea10cf1cc30448db786533/README.md)
- [SQL functions](https://github.com/schmiddy/pg_prioritize/blob/c160271202ca8a713dea10cf1cc30448db786533/prioritize.sql.in)
- [Control file](https://github.com/schmiddy/pg_prioritize/blob/c160271202ca8a713dea10cf1cc30448db786533/prioritize.control)

`prioritize` exposes operating-system priority controls for PostgreSQL backend processes. Use it to lower the scheduling priority of selected sessions; it is not a PostgreSQL query scheduler.

### Inspect and Adjust a Backend

```sql
CREATE EXTENSION prioritize;
SELECT get_backend_priority(pg_backend_pid());
SELECT set_backend_priority(pg_backend_pid(), 10);
```

Any user may query a backend's priority. A PostgreSQL superuser may request adjustments for any backend; other users may adjust only backends running under the same database role.

### Adjust Related Sessions

```sql
SELECT set_backend_priority(pid, get_backend_priority(pid) + 5)
FROM pg_stat_activity
WHERE usename = CURRENT_USER;
```

Increasing the numeric nice value lowers operating-system scheduling priority. The example therefore reduces those backends' priority, rather than speeding them up.

### Privileges and Limits

Installation requires a superuser, with no preload or restart. Operating-system permissions still apply: ordinary PostgreSQL processes normally cannot lower the numeric nice value to increase priority. Database superuser privileges do not confer root privileges. Review platform scheduling policy and process identity before using bulk adjustments. The control file uses SQL version 1.0; distribution package version 1.0.4 is separate.
