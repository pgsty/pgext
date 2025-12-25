---
title: "ip4r"
linkTitle: "ip4r"
description: "IPv4/v6 and IPv4/v6 range index type for PostgreSQL"
weight: 3820
categories: ["TYPE"]
width: full
---

[**ip4r**](https://github.com/RhodiumToad/ip4r) : IPv4/v6 and IPv4/v6 range index type for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3820** | {{< badge content="ip4r" link="https://github.com/RhodiumToad/ip4r" >}} | {{< ext "ip4r" >}} | `2.4.2` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


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
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `ip4r` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.2` | {{< bg "18" "ip4r_18" "green" >}} {{< bg "17" "ip4r_17" "green" >}} {{< bg "16" "ip4r_16" "green" >}} {{< bg "15" "ip4r_15" "green" >}} {{< bg "14" "ip4r_14" "green" >}} {{< bg "13" "ip4r_13" "green" >}} | `ip4r_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.4.2` | {{< bg "18" "postgresql-18-ip4r" "green" >}} {{< bg "17" "postgresql-17-ip4r" "green" >}} {{< bg "16" "postgresql-16-ip4r" "green" >}} {{< bg "15" "postgresql-15-ip4r" "green" >}} {{< bg "14" "postgresql-14-ip4r" "green" >}} {{< bg "13" "postgresql-13-ip4r" "green" >}} | `postgresql-$v-ip4r` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 2.4.2" "ip4r_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 2.4.2" "ip4r_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.4.2" "ip4r_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.4.2" "ip4r_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.4.2" "ip4r_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.4.2" "ip4r_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "ip4r_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.4.2" "postgresql-18-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-17-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-16-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-15-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-14-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-13-ip4r : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.4.2" "postgresql-18-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-17-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-16-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-15-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-14-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-13-ip4r : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.4.2" "postgresql-18-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-17-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-16-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-15-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-14-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-13-ip4r : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.4.2" "postgresql-18-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-17-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-16-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-15-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-14-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-13-ip4r : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.4.2" "postgresql-18-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-17-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-16-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-15-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-14-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-13-ip4r : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.4.2" "postgresql-18-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-17-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-16-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-15-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-14-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-13-ip4r : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.4.2" "postgresql-18-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-17-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-16-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-15-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-14-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-13-ip4r : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.4.2" "postgresql-18-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-17-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-16-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-15-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-14-ip4r : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.2" "postgresql-13-ip4r : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ip4r_18` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.0 KiB | [ip4r_18-2.4.2-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/ip4r_18-2.4.2-3PGDG.rhel8.x86_64.rpm) |
| `ip4r_18` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.2 KiB | [ip4r_18-2.4.2-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/ip4r_18-2.4.2-3PGDG.rhel8.aarch64.rpm) |
| `ip4r_18` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.9 KiB | [ip4r_18-2.4.2-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ip4r_18-2.4.2-3PGDG.rhel9.x86_64.rpm) |
| `ip4r_18` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.8 KiB | [ip4r_18-2.4.2-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/ip4r_18-2.4.2-3PGDG.rhel9.aarch64.rpm) |
| `ip4r_18` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 79.0 KiB | [ip4r_18-2.4.2-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/ip4r_18-2.4.2-3PGDG.rhel10.x86_64.rpm) |
| `ip4r_18` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.5 KiB | [ip4r_18-2.4.2-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ip4r_18-2.4.2-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 180.5 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg12+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 173.9 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg12+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 180.6 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg13+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 175.3 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg13+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 181.3 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 176.4 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 176.9 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-18-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 171.9 KiB | [postgresql-18-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-18-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ip4r_17` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.9 KiB | [ip4r_17-2.4.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ip4r_17-2.4.2-2PGDG.rhel8.x86_64.rpm) |
| `ip4r_17` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.0 KiB | [ip4r_17-2.4.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ip4r_17-2.4.2-2PGDG.rhel8.aarch64.rpm) |
| `ip4r_17` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.8 KiB | [ip4r_17-2.4.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ip4r_17-2.4.2-2PGDG.rhel9.x86_64.rpm) |
| `ip4r_17` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.9 KiB | [ip4r_17-2.4.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ip4r_17-2.4.2-2PGDG.rhel9.aarch64.rpm) |
| `ip4r_17` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 79.0 KiB | [ip4r_17-2.4.2-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ip4r_17-2.4.2-3PGDG.rhel10.x86_64.rpm) |
| `ip4r_17` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.5 KiB | [ip4r_17-2.4.2-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ip4r_17-2.4.2-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 180.4 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg12+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 174.0 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg12+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 180.6 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg13+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 175.2 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg13+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 194.9 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 189.6 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 176.8 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-17-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 171.7 KiB | [postgresql-17-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-17-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ip4r_16` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.7 KiB | [ip4r_16-2.4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ip4r_16-2.4.2-1PGDG.rhel8.x86_64.rpm) |
| `ip4r_16` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.0 KiB | [ip4r_16-2.4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ip4r_16-2.4.2-1PGDG.rhel8.aarch64.rpm) |
| `ip4r_16` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 76.7 KiB | [ip4r_16-2.4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ip4r_16-2.4.2-1PGDG.rhel9.x86_64.rpm) |
| `ip4r_16` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.4 KiB | [ip4r_16-2.4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ip4r_16-2.4.2-1PGDG.rhel9.aarch64.rpm) |
| `ip4r_16` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.9 KiB | [ip4r_16-2.4.2-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ip4r_16-2.4.2-3PGDG.rhel10.x86_64.rpm) |
| `ip4r_16` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 74.5 KiB | [ip4r_16-2.4.2-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ip4r_16-2.4.2-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 180.4 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg12+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 174.0 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg12+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 180.5 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg13+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 175.2 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg13+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 194.5 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 189.4 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 176.9 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-16-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 171.9 KiB | [postgresql-16-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-16-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ip4r_15` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.1 KiB | [ip4r_15-2.4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ip4r_15-2.4.2-1PGDG.rhel8.x86_64.rpm) |
| `ip4r_15` | `2.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 208.7 KiB | [ip4r_15-2.4.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ip4r_15-2.4.1-2.rhel8.x86_64.rpm) |
| `ip4r_15` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 72.0 KiB | [ip4r_15-2.4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ip4r_15-2.4.2-1PGDG.rhel8.aarch64.rpm) |
| `ip4r_15` | `2.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 203.2 KiB | [ip4r_15-2.4.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ip4r_15-2.4.1-2.rhel8.aarch64.rpm) |
| `ip4r_15` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.6 KiB | [ip4r_15-2.4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ip4r_15-2.4.2-1PGDG.rhel9.x86_64.rpm) |
| `ip4r_15` | `2.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 209.9 KiB | [ip4r_15-2.4.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ip4r_15-2.4.1-2.rhel9.x86_64.rpm) |
| `ip4r_15` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.3 KiB | [ip4r_15-2.4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ip4r_15-2.4.2-1PGDG.rhel9.aarch64.rpm) |
| `ip4r_15` | `2.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 204.6 KiB | [ip4r_15-2.4.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ip4r_15-2.4.1-2.rhel9.aarch64.rpm) |
| `ip4r_15` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.6 KiB | [ip4r_15-2.4.2-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ip4r_15-2.4.2-3PGDG.rhel10.x86_64.rpm) |
| `ip4r_15` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.8 KiB | [ip4r_15-2.4.2-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ip4r_15-2.4.2-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 179.2 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg12+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 172.7 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg12+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 179.2 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg13+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 173.9 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg13+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 192.9 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 187.5 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 175.0 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-15-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 170.2 KiB | [postgresql-15-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-15-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ip4r_14` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 77.1 KiB | [ip4r_14-2.4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ip4r_14-2.4.2-1PGDG.rhel8.x86_64.rpm) |
| `ip4r_14` | `2.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 210.6 KiB | [ip4r_14-2.4.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ip4r_14-2.4.1-2.rhel8.x86_64.rpm) |
| `ip4r_14` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.9 KiB | [ip4r_14-2.4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ip4r_14-2.4.2-1PGDG.rhel8.aarch64.rpm) |
| `ip4r_14` | `2.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 203.2 KiB | [ip4r_14-2.4.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ip4r_14-2.4.1-2.rhel8.aarch64.rpm) |
| `ip4r_14` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.6 KiB | [ip4r_14-2.4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ip4r_14-2.4.2-1PGDG.rhel9.x86_64.rpm) |
| `ip4r_14` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.3 KiB | [ip4r_14-2.4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ip4r_14-2.4.2-1PGDG.rhel9.aarch64.rpm) |
| `ip4r_14` | `2.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 204.7 KiB | [ip4r_14-2.4.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ip4r_14-2.4.1-2.rhel9.aarch64.rpm) |
| `ip4r_14` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.6 KiB | [ip4r_14-2.4.2-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ip4r_14-2.4.2-3PGDG.rhel10.x86_64.rpm) |
| `ip4r_14` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.8 KiB | [ip4r_14-2.4.2-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/ip4r_14-2.4.2-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 179.2 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg12+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 172.7 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg12+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 179.2 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg13+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 173.7 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg13+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 192.7 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 187.4 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 175.0 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-14-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 170.2 KiB | [postgresql-14-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-14-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ip4r_13` | `2.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.2 KiB | [ip4r_13-2.4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/ip4r_13-2.4.2-1PGDG.rhel8.x86_64.rpm) |
| `ip4r_13` | `2.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.9 KiB | [ip4r_13-2.4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ip4r_13-2.4.2-1PGDG.rhel8.aarch64.rpm) |
| `ip4r_13` | `2.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 202.7 KiB | [ip4r_13-2.4.1-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/ip4r_13-2.4.1-2.rhel8.aarch64.rpm) |
| `ip4r_13` | `2.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.2 KiB | [ip4r_13-2.4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/ip4r_13-2.4.2-1PGDG.rhel9.x86_64.rpm) |
| `ip4r_13` | `2.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 71.4 KiB | [ip4r_13-2.4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ip4r_13-2.4.2-1PGDG.rhel9.aarch64.rpm) |
| `ip4r_13` | `2.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 204.7 KiB | [ip4r_13-2.4.1-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/ip4r_13-2.4.1-2.rhel9.aarch64.rpm) |
| `ip4r_13` | `2.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 78.1 KiB | [ip4r_13-2.4.2-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/ip4r_13-2.4.2-3PGDG.rhel10.x86_64.rpm) |
| `ip4r_13` | `2.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.7 KiB | [ip4r_13-2.4.2-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/ip4r_13-2.4.2-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-ip4r` | `2.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 178.6 KiB | [postgresql-13-ip4r_2.4.2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-13-ip4r_2.4.2-4.pgdg12+1_amd64.deb) |
| `postgresql-13-ip4r` | `2.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 172.5 KiB | [postgresql-13-ip4r_2.4.2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-13-ip4r_2.4.2-4.pgdg12+1_arm64.deb) |
| `postgresql-13-ip4r` | `2.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 178.8 KiB | [postgresql-13-ip4r_2.4.2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-13-ip4r_2.4.2-4.pgdg13+1_amd64.deb) |
| `postgresql-13-ip4r` | `2.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 173.7 KiB | [postgresql-13-ip4r_2.4.2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-13-ip4r_2.4.2-4.pgdg13+1_arm64.deb) |
| `postgresql-13-ip4r` | `2.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 192.1 KiB | [postgresql-13-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-13-ip4r_2.4.2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-13-ip4r` | `2.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 187.1 KiB | [postgresql-13-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-13-ip4r_2.4.2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-13-ip4r` | `2.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 174.5 KiB | [postgresql-13-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-13-ip4r_2.4.2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-13-ip4r` | `2.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 169.9 KiB | [postgresql-13-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/i/ip4r/postgresql-13-ip4r_2.4.2-4.pgdg24.04+1_arm64.deb) |

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
pig install ip4r -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION ip4r;
```
