---
title: "ip4r"
linkTitle: "ip4r"
description: "IPv4/v6 and IPv4/v6 range index type for PostgreSQL"
weight: 3770
categories: ["TYPE"]
width: full
---

[**ip4r**](https://github.com/RhodiumToad/ip4r) : IPv4/v6 and IPv4/v6 range index type for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3770** | {{< badge content="ip4r" link="https://github.com/RhodiumToad/ip4r" >}} | {{< ext "ip4r" >}} | `2.4.3` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "geoip" >}} |
|   **See Also**    | {{< ext "pg_net" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `ip4r` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.3` | {{< bg "18" "ip4r_18" "green" >}} {{< bg "17" "ip4r_17" "green" >}} {{< bg "16" "ip4r_16" "green" >}} {{< bg "15" "ip4r_15" "green" >}} {{< bg "14" "ip4r_14" "green" >}} | `ip4r_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.3` | {{< bg "18" "postgresql-18-ip4r" "green" >}} {{< bg "17" "postgresql-17-ip4r" "green" >}} {{< bg "16" "postgresql-16-ip4r" "green" >}} {{< bg "15" "postgresql-15-ip4r" "green" >}} {{< bg "14" "postgresql-14-ip4r" "green" >}} | `postgresql-$v-ip4r` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.4.3" "ip4r_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_14 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.4.3" "ip4r_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.4.3" "ip4r_18 : AVAIL 8" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_14 : AVAIL 8" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.4.3" "ip4r_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_14 : AVAIL 9" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.4.3" "ip4r_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.4.3" "ip4r_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "ip4r_14 : AVAIL 3" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.4.3" "postgresql-18-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-17-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-16-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-15-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-14-ip4r : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.4.3" "postgresql-18-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-17-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-16-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-15-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-14-ip4r : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.4.3" "postgresql-18-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-17-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-16-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-15-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-14-ip4r : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.4.3" "postgresql-18-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-17-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-16-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-15-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-14-ip4r : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.4.3" "postgresql-18-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-17-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-16-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-15-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-14-ip4r : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.4.3" "postgresql-18-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-17-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-16-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-15-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-14-ip4r : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.4.3" "postgresql-18-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-17-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-16-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-15-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-14-ip4r : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.4.3" "postgresql-18-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-17-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-16-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-15-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-14-ip4r : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.4.3" "postgresql-18-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-17-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-16-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-15-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-14-ip4r : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.4.3" "postgresql-18-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-17-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-16-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-15-ip4r : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.3" "postgresql-14-ip4r : AVAIL 3" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ip4r_18` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.8 KiB | [ip4r_18-2.4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/ip4r_18-2.4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `ip4r_18` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.5 KiB | [ip4r_18-2.4.2-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/ip4r_18-2.4.2-6PGDG.rhel8.10.x86_64.rpm) |
| `ip4r_18` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.0 KiB | [ip4r_18-2.4.2-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/ip4r_18-2.4.2-3PGDG.rhel8.x86_64.rpm) |
| `ip4r_18` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.0 KiB | [ip4r_18-2.4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/ip4r_18-2.4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `ip4r_18` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.8 KiB | [ip4r_18-2.4.2-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/ip4r_18-2.4.2-6PGDG.rhel8.10.aarch64.rpm) |
| `ip4r_18` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.2 KiB | [ip4r_18-2.4.2-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/ip4r_18-2.4.2-3PGDG.rhel8.aarch64.rpm) |
| `ip4r_18` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.3 KiB | [ip4r_18-2.4.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ip4r_18-2.4.3-1PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_18` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.3 KiB | [ip4r_18-2.4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ip4r_18-2.4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `ip4r_18` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.5 KiB | [ip4r_18-2.4.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ip4r_18-2.4.3-1PGDG.rhel9.6.x86_64.rpm) |
| `ip4r_18` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.1 KiB | [ip4r_18-2.4.2-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ip4r_18-2.4.2-6PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_18` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.1 KiB | [ip4r_18-2.4.2-6PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ip4r_18-2.4.2-6PGDG.rhel9.7.x86_64.rpm) |
| `ip4r_18` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.3 KiB | [ip4r_18-2.4.2-6PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ip4r_18-2.4.2-6PGDG.rhel9.6.x86_64.rpm) |
| `ip4r_18` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.9 KiB | [ip4r_18-2.4.2-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ip4r_18-2.4.2-5PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_18` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.9 KiB | [ip4r_18-2.4.2-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ip4r_18-2.4.2-3PGDG.rhel9.x86_64.rpm) |
| `ip4r_18` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.3 KiB | [ip4r_18-2.4.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/ip4r_18-2.4.3-1PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_18` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.2 KiB | [ip4r_18-2.4.2-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/ip4r_18-2.4.2-6PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_18` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.9 KiB | [ip4r_18-2.4.2-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/ip4r_18-2.4.2-5PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_18` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 79.2 KiB | [ip4r_18-2.4.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/ip4r_18-2.4.3-1PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_18` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.9 KiB | [ip4r_18-2.4.2-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/ip4r_18-2.4.2-6PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_18` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.7 KiB | [ip4r_18-2.4.2-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/ip4r_18-2.4.2-5PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_18` | `2.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.8 KiB | [ip4r_18-2.4.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ip4r_18-2.4.3-1PGDG.rhel10.2.aarch64.rpm) |
| `ip4r_18` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.6 KiB | [ip4r_18-2.4.2-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ip4r_18-2.4.2-6PGDG.rhel10.2.aarch64.rpm) |
| `ip4r_18` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.3 KiB | [ip4r_18-2.4.2-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ip4r_18-2.4.2-5PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-ip4r` | `2.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 181.0 KiB | [postgresql-18-ip4r_2.4.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.3-1.pgdg12+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 180.9 KiB | [postgresql-18-ip4r_2.4.2-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-5.pgdg12+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 180.5 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg12+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 174.4 KiB | [postgresql-18-ip4r_2.4.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.3-1.pgdg12+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 174.4 KiB | [postgresql-18-ip4r_2.4.2-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-5.pgdg12+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 173.9 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg12+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 181.2 KiB | [postgresql-18-ip4r_2.4.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.3-1.pgdg13+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 181.2 KiB | [postgresql-18-ip4r_2.4.2-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-5.pgdg13+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 180.6 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg13+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 175.6 KiB | [postgresql-18-ip4r_2.4.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.3-1.pgdg13+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 175.7 KiB | [postgresql-18-ip4r_2.4.2-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-5.pgdg13+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 175.3 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg13+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 182.1 KiB | [postgresql-18-ip4r_2.4.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 181.9 KiB | [postgresql-18-ip4r_2.4.2-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-5.pgdg22.04+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 181.3 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 177.0 KiB | [postgresql-18-ip4r_2.4.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 177.0 KiB | [postgresql-18-ip4r_2.4.2-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-5.pgdg22.04+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 176.4 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 177.4 KiB | [postgresql-18-ip4r_2.4.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 177.4 KiB | [postgresql-18-ip4r_2.4.2-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-5.pgdg24.04+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 176.9 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 172.2 KiB | [postgresql-18-ip4r_2.4.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 172.3 KiB | [postgresql-18-ip4r_2.4.2-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-5.pgdg24.04+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 171.9 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 174.9 KiB | [postgresql-18-ip4r_2.4.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 174.4 KiB | [postgresql-18-ip4r_2.4.2-5.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-5.pgdg26.04+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 174.5 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg26.04+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 170.3 KiB | [postgresql-18-ip4r_2.4.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 170.6 KiB | [postgresql-18-ip4r_2.4.2-5.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-5.pgdg26.04+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 170.8 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ip4r_17` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.8 KiB | [ip4r_17-2.4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ip4r_17-2.4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `ip4r_17` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.6 KiB | [ip4r_17-2.4.2-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ip4r_17-2.4.2-6PGDG.rhel8.10.x86_64.rpm) |
| `ip4r_17` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.9 KiB | [ip4r_17-2.4.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ip4r_17-2.4.2-2PGDG.rhel8.x86_64.rpm) |
| `ip4r_17` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.0 KiB | [ip4r_17-2.4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ip4r_17-2.4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `ip4r_17` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.8 KiB | [ip4r_17-2.4.2-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ip4r_17-2.4.2-6PGDG.rhel8.10.aarch64.rpm) |
| `ip4r_17` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.0 KiB | [ip4r_17-2.4.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ip4r_17-2.4.2-2PGDG.rhel8.aarch64.rpm) |
| `ip4r_17` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.4 KiB | [ip4r_17-2.4.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ip4r_17-2.4.3-1PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_17` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.3 KiB | [ip4r_17-2.4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ip4r_17-2.4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `ip4r_17` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.4 KiB | [ip4r_17-2.4.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ip4r_17-2.4.3-1PGDG.rhel9.6.x86_64.rpm) |
| `ip4r_17` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.0 KiB | [ip4r_17-2.4.2-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ip4r_17-2.4.2-6PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_17` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.0 KiB | [ip4r_17-2.4.2-6PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ip4r_17-2.4.2-6PGDG.rhel9.7.x86_64.rpm) |
| `ip4r_17` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.3 KiB | [ip4r_17-2.4.2-6PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ip4r_17-2.4.2-6PGDG.rhel9.6.x86_64.rpm) |
| `ip4r_17` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.8 KiB | [ip4r_17-2.4.2-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ip4r_17-2.4.2-5PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_17` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.8 KiB | [ip4r_17-2.4.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ip4r_17-2.4.2-2PGDG.rhel9.x86_64.rpm) |
| `ip4r_17` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.4 KiB | [ip4r_17-2.4.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ip4r_17-2.4.3-1PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_17` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.1 KiB | [ip4r_17-2.4.2-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ip4r_17-2.4.2-6PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_17` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.9 KiB | [ip4r_17-2.4.2-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ip4r_17-2.4.2-5PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_17` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 79.1 KiB | [ip4r_17-2.4.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ip4r_17-2.4.3-1PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_17` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.9 KiB | [ip4r_17-2.4.2-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ip4r_17-2.4.2-6PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_17` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.7 KiB | [ip4r_17-2.4.2-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ip4r_17-2.4.2-5PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_17` | `2.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.7 KiB | [ip4r_17-2.4.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ip4r_17-2.4.3-1PGDG.rhel10.2.aarch64.rpm) |
| `ip4r_17` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.5 KiB | [ip4r_17-2.4.2-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ip4r_17-2.4.2-6PGDG.rhel10.2.aarch64.rpm) |
| `ip4r_17` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.3 KiB | [ip4r_17-2.4.2-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ip4r_17-2.4.2-5PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-ip4r` | `2.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 180.7 KiB | [postgresql-17-ip4r_2.4.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.3-1.pgdg12+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 180.7 KiB | [postgresql-17-ip4r_2.4.2-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-5.pgdg12+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 180.4 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg12+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 174.3 KiB | [postgresql-17-ip4r_2.4.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.3-1.pgdg12+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 174.4 KiB | [postgresql-17-ip4r_2.4.2-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-5.pgdg12+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 174.0 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg12+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 181.1 KiB | [postgresql-17-ip4r_2.4.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.3-1.pgdg13+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 180.9 KiB | [postgresql-17-ip4r_2.4.2-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-5.pgdg13+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 180.6 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg13+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 175.5 KiB | [postgresql-17-ip4r_2.4.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.3-1.pgdg13+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 175.6 KiB | [postgresql-17-ip4r_2.4.2-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-5.pgdg13+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 175.2 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg13+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 194.6 KiB | [postgresql-17-ip4r_2.4.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 194.4 KiB | [postgresql-17-ip4r_2.4.2-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-5.pgdg22.04+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 194.9 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 189.4 KiB | [postgresql-17-ip4r_2.4.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 189.3 KiB | [postgresql-17-ip4r_2.4.2-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-5.pgdg22.04+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 189.6 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 177.2 KiB | [postgresql-17-ip4r_2.4.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 177.2 KiB | [postgresql-17-ip4r_2.4.2-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-5.pgdg24.04+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 176.8 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 172.1 KiB | [postgresql-17-ip4r_2.4.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 172.0 KiB | [postgresql-17-ip4r_2.4.2-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-5.pgdg24.04+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 171.7 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 174.7 KiB | [postgresql-17-ip4r_2.4.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 174.6 KiB | [postgresql-17-ip4r_2.4.2-5.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-5.pgdg26.04+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 174.6 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg26.04+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 170.6 KiB | [postgresql-17-ip4r_2.4.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 170.7 KiB | [postgresql-17-ip4r_2.4.2-5.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-5.pgdg26.04+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 170.3 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ip4r_16` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.8 KiB | [ip4r_16-2.4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ip4r_16-2.4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `ip4r_16` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.6 KiB | [ip4r_16-2.4.2-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ip4r_16-2.4.2-6PGDG.rhel8.10.x86_64.rpm) |
| `ip4r_16` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.7 KiB | [ip4r_16-2.4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ip4r_16-2.4.2-1PGDG.rhel8.x86_64.rpm) |
| `ip4r_16` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.0 KiB | [ip4r_16-2.4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ip4r_16-2.4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `ip4r_16` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.8 KiB | [ip4r_16-2.4.2-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ip4r_16-2.4.2-6PGDG.rhel8.10.aarch64.rpm) |
| `ip4r_16` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.0 KiB | [ip4r_16-2.4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ip4r_16-2.4.2-1PGDG.rhel8.aarch64.rpm) |
| `ip4r_16` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.3 KiB | [ip4r_16-2.4.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ip4r_16-2.4.3-1PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_16` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.2 KiB | [ip4r_16-2.4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ip4r_16-2.4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `ip4r_16` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.3 KiB | [ip4r_16-2.4.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ip4r_16-2.4.3-1PGDG.rhel9.6.x86_64.rpm) |
| `ip4r_16` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.1 KiB | [ip4r_16-2.4.2-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ip4r_16-2.4.2-6PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_16` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.2 KiB | [ip4r_16-2.4.2-6PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ip4r_16-2.4.2-6PGDG.rhel9.7.x86_64.rpm) |
| `ip4r_16` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 77.1 KiB | [ip4r_16-2.4.2-6PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ip4r_16-2.4.2-6PGDG.rhel9.6.x86_64.rpm) |
| `ip4r_16` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.9 KiB | [ip4r_16-2.4.2-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ip4r_16-2.4.2-5PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_16` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.7 KiB | [ip4r_16-2.4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ip4r_16-2.4.2-1PGDG.rhel9.x86_64.rpm) |
| `ip4r_16` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.3 KiB | [ip4r_16-2.4.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ip4r_16-2.4.3-1PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_16` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 73.1 KiB | [ip4r_16-2.4.2-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ip4r_16-2.4.2-6PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_16` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.9 KiB | [ip4r_16-2.4.2-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ip4r_16-2.4.2-5PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_16` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 79.1 KiB | [ip4r_16-2.4.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ip4r_16-2.4.3-1PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_16` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.9 KiB | [ip4r_16-2.4.2-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ip4r_16-2.4.2-6PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_16` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.7 KiB | [ip4r_16-2.4.2-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ip4r_16-2.4.2-5PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_16` | `2.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.7 KiB | [ip4r_16-2.4.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ip4r_16-2.4.3-1PGDG.rhel10.2.aarch64.rpm) |
| `ip4r_16` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.5 KiB | [ip4r_16-2.4.2-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ip4r_16-2.4.2-6PGDG.rhel10.2.aarch64.rpm) |
| `ip4r_16` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.3 KiB | [ip4r_16-2.4.2-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ip4r_16-2.4.2-5PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-ip4r` | `2.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 180.8 KiB | [postgresql-16-ip4r_2.4.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.3-1.pgdg12+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 180.7 KiB | [postgresql-16-ip4r_2.4.2-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-5.pgdg12+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 180.4 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg12+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 174.3 KiB | [postgresql-16-ip4r_2.4.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.3-1.pgdg12+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 174.3 KiB | [postgresql-16-ip4r_2.4.2-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-5.pgdg12+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 174.0 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg12+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 180.9 KiB | [postgresql-16-ip4r_2.4.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.3-1.pgdg13+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 181.0 KiB | [postgresql-16-ip4r_2.4.2-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-5.pgdg13+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 180.5 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg13+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 175.6 KiB | [postgresql-16-ip4r_2.4.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.3-1.pgdg13+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 175.6 KiB | [postgresql-16-ip4r_2.4.2-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-5.pgdg13+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 175.2 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg13+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 194.2 KiB | [postgresql-16-ip4r_2.4.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 194.3 KiB | [postgresql-16-ip4r_2.4.2-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-5.pgdg22.04+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 194.5 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 189.3 KiB | [postgresql-16-ip4r_2.4.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 189.2 KiB | [postgresql-16-ip4r_2.4.2-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-5.pgdg22.04+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 189.4 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 177.4 KiB | [postgresql-16-ip4r_2.4.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 177.2 KiB | [postgresql-16-ip4r_2.4.2-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-5.pgdg24.04+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 176.9 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 172.1 KiB | [postgresql-16-ip4r_2.4.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 172.0 KiB | [postgresql-16-ip4r_2.4.2-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-5.pgdg24.04+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 171.9 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 174.1 KiB | [postgresql-16-ip4r_2.4.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 174.8 KiB | [postgresql-16-ip4r_2.4.2-5.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-5.pgdg26.04+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 175.0 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg26.04+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 170.7 KiB | [postgresql-16-ip4r_2.4.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 169.9 KiB | [postgresql-16-ip4r_2.4.2-5.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-5.pgdg26.04+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 170.3 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ip4r_15` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.1 KiB | [ip4r_15-2.4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ip4r_15-2.4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `ip4r_15` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.9 KiB | [ip4r_15-2.4.2-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ip4r_15-2.4.2-6PGDG.rhel8.10.x86_64.rpm) |
| `ip4r_15` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.1 KiB | [ip4r_15-2.4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ip4r_15-2.4.2-1PGDG.rhel8.x86_64.rpm) |
| `ip4r_15` | `2.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 208.7 KiB | [ip4r_15-2.4.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ip4r_15-2.4.1-2.rhel8.x86_64.rpm) |
| `ip4r_15` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.0 KiB | [ip4r_15-2.4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ip4r_15-2.4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `ip4r_15` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 72.8 KiB | [ip4r_15-2.4.2-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ip4r_15-2.4.2-6PGDG.rhel8.10.aarch64.rpm) |
| `ip4r_15` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 72.0 KiB | [ip4r_15-2.4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ip4r_15-2.4.2-1PGDG.rhel8.aarch64.rpm) |
| `ip4r_15` | `2.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 203.2 KiB | [ip4r_15-2.4.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ip4r_15-2.4.1-2.rhel8.aarch64.rpm) |
| `ip4r_15` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.3 KiB | [ip4r_15-2.4.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ip4r_15-2.4.3-1PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_15` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.3 KiB | [ip4r_15-2.4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ip4r_15-2.4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `ip4r_15` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.4 KiB | [ip4r_15-2.4.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ip4r_15-2.4.3-1PGDG.rhel9.6.x86_64.rpm) |
| `ip4r_15` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.1 KiB | [ip4r_15-2.4.2-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ip4r_15-2.4.2-6PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_15` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.2 KiB | [ip4r_15-2.4.2-6PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ip4r_15-2.4.2-6PGDG.rhel9.7.x86_64.rpm) |
| `ip4r_15` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.2 KiB | [ip4r_15-2.4.2-6PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ip4r_15-2.4.2-6PGDG.rhel9.6.x86_64.rpm) |
| `ip4r_15` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.8 KiB | [ip4r_15-2.4.2-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ip4r_15-2.4.2-5PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_15` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.6 KiB | [ip4r_15-2.4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ip4r_15-2.4.2-1PGDG.rhel9.x86_64.rpm) |
| `ip4r_15` | `2.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 209.9 KiB | [ip4r_15-2.4.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ip4r_15-2.4.1-2.rhel9.x86_64.rpm) |
| `ip4r_15` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.1 KiB | [ip4r_15-2.4.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ip4r_15-2.4.3-1PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_15` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.9 KiB | [ip4r_15-2.4.2-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ip4r_15-2.4.2-6PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_15` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.7 KiB | [ip4r_15-2.4.2-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ip4r_15-2.4.2-5PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_15` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.7 KiB | [ip4r_15-2.4.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ip4r_15-2.4.3-1PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_15` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.5 KiB | [ip4r_15-2.4.2-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ip4r_15-2.4.2-6PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_15` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.3 KiB | [ip4r_15-2.4.2-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ip4r_15-2.4.2-5PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_15` | `2.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.0 KiB | [ip4r_15-2.4.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ip4r_15-2.4.3-1PGDG.rhel10.2.aarch64.rpm) |
| `ip4r_15` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.9 KiB | [ip4r_15-2.4.2-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ip4r_15-2.4.2-6PGDG.rhel10.2.aarch64.rpm) |
| `ip4r_15` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.6 KiB | [ip4r_15-2.4.2-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ip4r_15-2.4.2-5PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-ip4r` | `2.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 179.7 KiB | [postgresql-15-ip4r_2.4.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.3-1.pgdg12+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 179.6 KiB | [postgresql-15-ip4r_2.4.2-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-5.pgdg12+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 179.2 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg12+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 173.0 KiB | [postgresql-15-ip4r_2.4.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.3-1.pgdg12+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 172.8 KiB | [postgresql-15-ip4r_2.4.2-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-5.pgdg12+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 172.7 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg12+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 179.9 KiB | [postgresql-15-ip4r_2.4.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.3-1.pgdg13+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 179.4 KiB | [postgresql-15-ip4r_2.4.2-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-5.pgdg13+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 179.2 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg13+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 174.1 KiB | [postgresql-15-ip4r_2.4.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.3-1.pgdg13+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 173.9 KiB | [postgresql-15-ip4r_2.4.2-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-5.pgdg13+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 173.9 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg13+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 192.6 KiB | [postgresql-15-ip4r_2.4.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 192.6 KiB | [postgresql-15-ip4r_2.4.2-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-5.pgdg22.04+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 192.9 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 187.4 KiB | [postgresql-15-ip4r_2.4.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 187.3 KiB | [postgresql-15-ip4r_2.4.2-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-5.pgdg22.04+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 187.5 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 175.6 KiB | [postgresql-15-ip4r_2.4.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 175.4 KiB | [postgresql-15-ip4r_2.4.2-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-5.pgdg24.04+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 175.0 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 170.6 KiB | [postgresql-15-ip4r_2.4.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 170.3 KiB | [postgresql-15-ip4r_2.4.2-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-5.pgdg24.04+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 170.2 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 172.7 KiB | [postgresql-15-ip4r_2.4.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 173.2 KiB | [postgresql-15-ip4r_2.4.2-5.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-5.pgdg26.04+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 172.9 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg26.04+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 169.1 KiB | [postgresql-15-ip4r_2.4.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 168.8 KiB | [postgresql-15-ip4r_2.4.2-5.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-5.pgdg26.04+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 169.3 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ip4r_14` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.1 KiB | [ip4r_14-2.4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ip4r_14-2.4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `ip4r_14` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.9 KiB | [ip4r_14-2.4.2-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ip4r_14-2.4.2-6PGDG.rhel8.10.x86_64.rpm) |
| `ip4r_14` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.1 KiB | [ip4r_14-2.4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ip4r_14-2.4.2-1PGDG.rhel8.x86_64.rpm) |
| `ip4r_14` | `2.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 210.6 KiB | [ip4r_14-2.4.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ip4r_14-2.4.1-2.rhel8.x86_64.rpm) |
| `ip4r_14` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 72.9 KiB | [ip4r_14-2.4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ip4r_14-2.4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `ip4r_14` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 72.8 KiB | [ip4r_14-2.4.2-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ip4r_14-2.4.2-6PGDG.rhel8.10.aarch64.rpm) |
| `ip4r_14` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.9 KiB | [ip4r_14-2.4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ip4r_14-2.4.2-1PGDG.rhel8.aarch64.rpm) |
| `ip4r_14` | `2.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 203.2 KiB | [ip4r_14-2.4.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ip4r_14-2.4.1-2.rhel8.aarch64.rpm) |
| `ip4r_14` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.3 KiB | [ip4r_14-2.4.3-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ip4r_14-2.4.3-1PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_14` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.3 KiB | [ip4r_14-2.4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ip4r_14-2.4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `ip4r_14` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.4 KiB | [ip4r_14-2.4.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ip4r_14-2.4.3-1PGDG.rhel9.6.x86_64.rpm) |
| `ip4r_14` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.1 KiB | [ip4r_14-2.4.2-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ip4r_14-2.4.2-6PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_14` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.1 KiB | [ip4r_14-2.4.2-6PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ip4r_14-2.4.2-6PGDG.rhel9.7.x86_64.rpm) |
| `ip4r_14` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.2 KiB | [ip4r_14-2.4.2-6PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ip4r_14-2.4.2-6PGDG.rhel9.6.x86_64.rpm) |
| `ip4r_14` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.9 KiB | [ip4r_14-2.4.2-5PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ip4r_14-2.4.2-5PGDG.rhel9.8.x86_64.rpm) |
| `ip4r_14` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.6 KiB | [ip4r_14-2.4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ip4r_14-2.4.2-1PGDG.rhel9.x86_64.rpm) |
| `ip4r_14` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.1 KiB | [ip4r_14-2.4.3-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ip4r_14-2.4.3-1PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_14` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.1 KiB | [ip4r_14-2.4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ip4r_14-2.4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `ip4r_14` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.2 KiB | [ip4r_14-2.4.3-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ip4r_14-2.4.3-1PGDG.rhel9.6.aarch64.rpm) |
| `ip4r_14` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.0 KiB | [ip4r_14-2.4.2-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ip4r_14-2.4.2-6PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_14` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.9 KiB | [ip4r_14-2.4.2-6PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ip4r_14-2.4.2-6PGDG.rhel9.7.aarch64.rpm) |
| `ip4r_14` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.1 KiB | [ip4r_14-2.4.2-6PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ip4r_14-2.4.2-6PGDG.rhel9.6.aarch64.rpm) |
| `ip4r_14` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.7 KiB | [ip4r_14-2.4.2-5PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ip4r_14-2.4.2-5PGDG.rhel9.8.aarch64.rpm) |
| `ip4r_14` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.3 KiB | [ip4r_14-2.4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ip4r_14-2.4.2-1PGDG.rhel9.aarch64.rpm) |
| `ip4r_14` | `2.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 204.7 KiB | [ip4r_14-2.4.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ip4r_14-2.4.1-2.rhel9.aarch64.rpm) |
| `ip4r_14` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.8 KiB | [ip4r_14-2.4.3-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ip4r_14-2.4.3-1PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_14` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.6 KiB | [ip4r_14-2.4.2-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ip4r_14-2.4.2-6PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_14` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.3 KiB | [ip4r_14-2.4.2-5PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ip4r_14-2.4.2-5PGDG.rhel10.2.x86_64.rpm) |
| `ip4r_14` | `2.4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.1 KiB | [ip4r_14-2.4.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/ip4r_14-2.4.3-1PGDG.rhel10.2.aarch64.rpm) |
| `ip4r_14` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.8 KiB | [ip4r_14-2.4.2-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/ip4r_14-2.4.2-6PGDG.rhel10.2.aarch64.rpm) |
| `ip4r_14` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.6 KiB | [ip4r_14-2.4.2-5PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/ip4r_14-2.4.2-5PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-14-ip4r` | `2.4.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 179.5 KiB | [postgresql-14-ip4r_2.4.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.3-1.pgdg12+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 179.5 KiB | [postgresql-14-ip4r_2.4.2-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-5.pgdg12+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 179.2 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg12+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 173.1 KiB | [postgresql-14-ip4r_2.4.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.3-1.pgdg12+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 173.0 KiB | [postgresql-14-ip4r_2.4.2-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-5.pgdg12+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 172.7 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg12+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 179.8 KiB | [postgresql-14-ip4r_2.4.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.3-1.pgdg13+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 179.8 KiB | [postgresql-14-ip4r_2.4.2-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-5.pgdg13+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 179.2 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg13+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 174.3 KiB | [postgresql-14-ip4r_2.4.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.3-1.pgdg13+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 174.2 KiB | [postgresql-14-ip4r_2.4.2-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-5.pgdg13+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 173.7 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg13+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 192.8 KiB | [postgresql-14-ip4r_2.4.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 192.5 KiB | [postgresql-14-ip4r_2.4.2-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-5.pgdg22.04+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 192.7 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 187.2 KiB | [postgresql-14-ip4r_2.4.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 187.1 KiB | [postgresql-14-ip4r_2.4.2-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-5.pgdg22.04+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 187.4 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 175.5 KiB | [postgresql-14-ip4r_2.4.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 175.3 KiB | [postgresql-14-ip4r_2.4.2-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-5.pgdg24.04+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 175.0 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 170.4 KiB | [postgresql-14-ip4r_2.4.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 170.4 KiB | [postgresql-14-ip4r_2.4.2-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-5.pgdg24.04+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 170.2 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 173.1 KiB | [postgresql-14-ip4r_2.4.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 172.7 KiB | [postgresql-14-ip4r_2.4.2-5.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-5.pgdg26.04+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 173.5 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg26.04+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 169.1 KiB | [postgresql-14-ip4r_2.4.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.3-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 168.5 KiB | [postgresql-14-ip4r_2.4.2-5.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-5.pgdg26.04+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 168.7 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/RhodiumToad/ip4r" title="Repository" icon="github" subtitle="github.com/RhodiumToad/ip4r" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install ip4r;		# install via package name, for the active PG version

pig install ip4r -v 18;   # install for PG 18
pig install ip4r -v 17;   # install for PG 17
pig install ip4r -v 16;   # install for PG 16
pig install ip4r -v 15;   # install for PG 15
pig install ip4r -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION ip4r;
```



## Usage

> [ip4r: IPv4/IPv6 address and range types with GiST indexing](https://github.com/RhodiumToad/ip4r)

The `ip4r` extension provides specialized data types for IPv4/IPv6 addresses and ranges with superior indexing for containment queries.

```sql
CREATE EXTENSION ip4r;
```

### Data Types

| Type | Description |
|------|-------------|
| `ip4` | Single IPv4 address (32-bit) |
| `ip6` | Single IPv6 address (dual 64-bit) |
| `ip4r` | IPv4 address range |
| `ip6r` | IPv6 address range |
| `ipaddress` | Mixed IPv4/IPv6 address |
| `iprange` | Mixed IPv4/IPv6 range |

### Address Input

```sql
SELECT '192.168.1.1'::ip4;
SELECT '2001:db8::1'::ip6;
SELECT '10.0.0.0/24'::ip4r;                   -- CIDR notation
SELECT '192.168.1.100-192.168.1.200'::ip4r;   -- explicit range
```

### Address Operators

- **Comparison**: `=`, `<>`, `<`, `>`, `<=`, `>=`
- **Arithmetic**: `+`, `-` with integers
- **Bitwise**: `&` (AND), `|` (OR), `#` (XOR), `~` (NOT)

### Address Functions

```sql
SELECT family('192.168.1.1'::ipaddress);       -- 4
SELECT ip4_netmask(24);                         -- 255.255.255.0
```

### Range Operators

| Operator | Description |
|----------|-------------|
| `>>=` | Contains or equal |
| `>>` | Strictly contains |
| `<<=` | Contained in or equal |
| `<<` | Strictly contained in |
| `&&` | Overlaps |

### Range Functions

```sql
SELECT lower('10.0.0.0/24'::ip4r);           -- 10.0.0.0
SELECT upper('10.0.0.0/24'::ip4r);           -- 10.0.0.255
SELECT is_cidr('10.0.0.0/24'::ip4r);         -- true
SELECT cidr_split('10.0.0.0-10.0.0.5'::ip4r); -- decompose to CIDRs
SELECT @ '10.0.0.0/24'::ip4r;                 -- approximate size
```

### Indexing

```sql
-- GiST index for containment queries
CREATE INDEX idx ON ipranges USING gist (range);

-- Find ranges containing a specific IP
SELECT * FROM ipranges WHERE range >>= '10.0.1.15'::ip4;

-- Find most specific match
SELECT * FROM ipranges
WHERE range >>= '10.0.1.15'::ip4
ORDER BY @ range LIMIT 1;
```
