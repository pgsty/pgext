---
title: "credcheck"
linkTitle: "credcheck"
description: "credcheck - postgresql plain text credential checker"
weight: 7310
categories: ["SEC"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PGDG"]
page_width: full
---

[**credcheck**](https://github.com/HexaCluster/credcheck) : credcheck - postgresql plain text credential checker


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7310** | {{< badge content="credcheck" link="https://github.com/HexaCluster/credcheck" >}} | {{< ext "credcheck" >}} | `5.0` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_pwhash" >}} {{< ext "passwordcheck" >}} {{< ext "passwordcheck_cracklib" >}} {{< ext "passwordpolicy" >}} {{< ext "chkpass" >}} {{< ext "pg_enigma" >}} {{< ext "column_encrypt" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `5.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `credcheck` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.7` | {{< bg "18" "credcheck_18" "green" >}} {{< bg "17" "credcheck_17" "green" >}} {{< bg "16" "credcheck_16" "green" >}} {{< bg "15" "credcheck_15" "green" >}} {{< bg "14" "credcheck_14" "green" >}} | `credcheck_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `5.0` | {{< bg "18" "postgresql-18-credcheck" "green" >}} {{< bg "17" "postgresql-17-credcheck" "green" >}} {{< bg "16" "postgresql-16-credcheck" "green" >}} {{< bg "15" "postgresql-15-credcheck" "green" >}} {{< bg "14" "postgresql-14-credcheck" "green" >}} | `postgresql-$v-credcheck` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.7" "credcheck_18 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_16 : AVAIL 12" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_15 : AVAIL 17" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_14 : AVAIL 17" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.7" "credcheck_18 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_17 : AVAIL 9" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_16 : AVAIL 12" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_15 : AVAIL 17" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_14 : AVAIL 17" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.7" "credcheck_18 : AVAIL 14" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_17 : AVAIL 15" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_16 : AVAIL 18" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_15 : AVAIL 23" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_14 : AVAIL 22" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.7" "credcheck_18 : AVAIL 14" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_17 : AVAIL 15" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_16 : AVAIL 18" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_15 : AVAIL 23" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_14 : AVAIL 23" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.7" "credcheck_18 : AVAIL 13" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_17 : AVAIL 13" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_16 : AVAIL 13" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_15 : AVAIL 13" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_14 : AVAIL 13" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.7" "credcheck_18 : AVAIL 14" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_17 : AVAIL 14" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_16 : AVAIL 14" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_15 : AVAIL 14" "blue" >}} | {{< bg "PGDG 4.7" "credcheck_14 : AVAIL 14" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 5.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-14-credcheck : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 5.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-14-credcheck : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 5.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-14-credcheck : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 5.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-14-credcheck : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 5.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-14-credcheck : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 5.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-14-credcheck : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 5.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-14-credcheck : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 5.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-14-credcheck : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 5.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-14-credcheck : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 5.0" "postgresql-18-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-17-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-16-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-15-credcheck : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.0" "postgresql-14-credcheck : AVAIL 3" "blue" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_18` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.3 KiB | [credcheck_18-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.8 KiB | [credcheck_18-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.5 KiB | [credcheck_18-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.8 KiB | [credcheck_18-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.6 KiB | [credcheck_18-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_18` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [credcheck_18-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_18` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.4 KiB | [credcheck_18-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_18` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.6 KiB | [credcheck_18-3.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/credcheck_18-3.0-2PGDG.rhel8.x86_64.rpm) |
| `credcheck_18` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.5 KiB | [credcheck_18-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.1 KiB | [credcheck_18-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.8 KiB | [credcheck_18-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.2 KiB | [credcheck_18-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.9 KiB | [credcheck_18-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_18` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.2 KiB | [credcheck_18-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_18` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.8 KiB | [credcheck_18-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_18` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.1 KiB | [credcheck_18-3.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/credcheck_18-3.0-2PGDG.rhel8.aarch64.rpm) |
| `credcheck_18` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.3 KiB | [credcheck_18-4.7-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.7-1PGDG.rhel9.8.x86_64.rpm) |
| `credcheck_18` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.3 KiB | [credcheck_18-4.7-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.7-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_18` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.4 KiB | [credcheck_18-4.7-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.7-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_18` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_18-4.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.6-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_18` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.0 KiB | [credcheck_18-4.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.6-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_18` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.8 KiB | [credcheck_18-4.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.5-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_18` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_18-4.5-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.5-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_18` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.3 KiB | [credcheck_18-4.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.4-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_18` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.4 KiB | [credcheck_18-4.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.4-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_18` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.0 KiB | [credcheck_18-4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_18` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.1 KiB | [credcheck_18-4.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.3-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_18` | `4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.6 KiB | [credcheck_18-4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_18` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.2 KiB | [credcheck_18-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_18` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.9 KiB | [credcheck_18-3.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/credcheck_18-3.0-2PGDG.rhel9.x86_64.rpm) |
| `credcheck_18` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.7 KiB | [credcheck_18-4.7-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.7-1PGDG.rhel9.8.aarch64.rpm) |
| `credcheck_18` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.7 KiB | [credcheck_18-4.7-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.7-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_18` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.8 KiB | [credcheck_18-4.7-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.7-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_18` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_18-4.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.6-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_18` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [credcheck_18-4.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.6-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_18` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.5 KiB | [credcheck_18-4.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.5-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_18` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [credcheck_18-4.5-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.5-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_18` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.6 KiB | [credcheck_18-4.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.4-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_18` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [credcheck_18-4.4-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.4-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_18` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [credcheck_18-4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_18` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.8 KiB | [credcheck_18-4.3-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.3-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_18` | `4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.1 KiB | [credcheck_18-4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_18` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.7 KiB | [credcheck_18-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_18` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.6 KiB | [credcheck_18-3.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/credcheck_18-3.0-2PGDG.rhel9.aarch64.rpm) |
| `credcheck_18` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [credcheck_18-4.7-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.7-1PGDG.rhel10.2.x86_64.rpm) |
| `credcheck_18` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [credcheck_18-4.7-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.7-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_18` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.0 KiB | [credcheck_18-4.7-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.7-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_18` | `4.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.5 KiB | [credcheck_18-4.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.6-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_18` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.1 KiB | [credcheck_18-4.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.5-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_18` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.4 KiB | [credcheck_18-4.5-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.5-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_18` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.6 KiB | [credcheck_18-4.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.4-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_18` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.9 KiB | [credcheck_18-4.4-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.4-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_18` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.4 KiB | [credcheck_18-4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_18` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.7 KiB | [credcheck_18-4.3-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.3-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_18` | `4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.3 KiB | [credcheck_18-4.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.2-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_18` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.7 KiB | [credcheck_18-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_18` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 36.3 KiB | [credcheck_18-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/credcheck_18-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_18` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.1 KiB | [credcheck_18-4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_18` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.1 KiB | [credcheck_18-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.1 KiB | [credcheck_18-4.7-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.7-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_18` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_18-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_18-4.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.6-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_18` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_18-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_18-4.5-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.5-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_18` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.3 KiB | [credcheck_18-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.3 KiB | [credcheck_18-4.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.4-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_18` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_18-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_18` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_18-4.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.3-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_18` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.9 KiB | [credcheck_18-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_18` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.5 KiB | [credcheck_18-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_18` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.3 KiB | [credcheck_18-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/credcheck_18-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.8 KiB | [postgresql-18-credcheck_5.0-2.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg12+2_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.8 KiB | [postgresql-18-credcheck_5.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg12+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.8 KiB | [postgresql-18-credcheck_5.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.7 KiB | [postgresql-18-credcheck_5.0-2.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg12+2_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.7 KiB | [postgresql-18-credcheck_5.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg12+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.6 KiB | [postgresql-18-credcheck_5.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.8 KiB | [postgresql-18-credcheck_5.0-2.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg13+2_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.8 KiB | [postgresql-18-credcheck_5.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg13+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.7 KiB | [postgresql-18-credcheck_5.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.3 KiB | [postgresql-18-credcheck_5.0-2.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg13+2_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.3 KiB | [postgresql-18-credcheck_5.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg13+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.3 KiB | [postgresql-18-credcheck_5.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 74.6 KiB | [postgresql-18-credcheck_5.0-2.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg22.04+2_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 74.5 KiB | [postgresql-18-credcheck_5.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 74.5 KiB | [postgresql-18-credcheck_5.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 73.1 KiB | [postgresql-18-credcheck_5.0-2.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg22.04+2_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 73.1 KiB | [postgresql-18-credcheck_5.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 73.0 KiB | [postgresql-18-credcheck_5.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 74.1 KiB | [postgresql-18-credcheck_5.0-2.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg24.04+2_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 74.1 KiB | [postgresql-18-credcheck_5.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 74.1 KiB | [postgresql-18-credcheck_5.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.8 KiB | [postgresql-18-credcheck_5.0-2.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg24.04+2_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.8 KiB | [postgresql-18-credcheck_5.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.7 KiB | [postgresql-18-credcheck_5.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.5 KiB | [postgresql-18-credcheck_5.0-2.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg26.04+2_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.5 KiB | [postgresql-18-credcheck_5.0-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg26.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.6 KiB | [postgresql-18-credcheck_5.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 72.0 KiB | [postgresql-18-credcheck_5.0-2.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg26.04+2_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 72.0 KiB | [postgresql-18-credcheck_5.0-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-2.pgdg26.04+1_arm64.deb) |
| `postgresql-18-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 72.0 KiB | [postgresql-18-credcheck_5.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-18-credcheck_5.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_17` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.4 KiB | [credcheck_17-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [credcheck_17-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.5 KiB | [credcheck_17-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.9 KiB | [credcheck_17-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.6 KiB | [credcheck_17-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_17` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [credcheck_17-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.5 KiB | [credcheck_17-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.5 KiB | [credcheck_17-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | `2.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.1 KiB | [credcheck_17-2.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/credcheck_17-2.8-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_17` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.6 KiB | [credcheck_17-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.2 KiB | [credcheck_17-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.8 KiB | [credcheck_17-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.2 KiB | [credcheck_17-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.0 KiB | [credcheck_17-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_17` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.3 KiB | [credcheck_17-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.9 KiB | [credcheck_17-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.0 KiB | [credcheck_17-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | `2.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.7 KiB | [credcheck_17-2.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/credcheck_17-2.8-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_17` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.5 KiB | [credcheck_17-4.7-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.7-1PGDG.rhel9.8.x86_64.rpm) |
| `credcheck_17` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.4 KiB | [credcheck_17-4.7-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.7-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_17` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.5 KiB | [credcheck_17-4.7-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.7-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_17` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_17-4.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.6-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_17` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.0 KiB | [credcheck_17-4.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.6-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_17` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_17-4.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.5-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_17` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.0 KiB | [credcheck_17-4.5-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.5-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_17` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.3 KiB | [credcheck_17-4.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.4-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_17` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.4 KiB | [credcheck_17-4.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.4-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_17` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.1 KiB | [credcheck_17-4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_17` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.2 KiB | [credcheck_17-4.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.3-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_17` | `4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.6 KiB | [credcheck_17-4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_17` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.2 KiB | [credcheck_17-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_17` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.9 KiB | [credcheck_17-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-3.0-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_17` | `2.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.6 KiB | [credcheck_17-2.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/credcheck_17-2.8-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_17` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.8 KiB | [credcheck_17-4.7-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.7-1PGDG.rhel9.8.aarch64.rpm) |
| `credcheck_17` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.8 KiB | [credcheck_17-4.7-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.7-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_17` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.9 KiB | [credcheck_17-4.7-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.7-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_17` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [credcheck_17-4.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.6-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_17` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.4 KiB | [credcheck_17-4.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.6-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_17` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_17-4.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.5-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_17` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [credcheck_17-4.5-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.5-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_17` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.0 KiB | [credcheck_17-4.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.4-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_17` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.1 KiB | [credcheck_17-4.4-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.4-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_17` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.8 KiB | [credcheck_17-4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_17` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.9 KiB | [credcheck_17-4.3-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.3-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_17` | `4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.2 KiB | [credcheck_17-4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_17` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.8 KiB | [credcheck_17-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_17` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.7 KiB | [credcheck_17-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-3.0-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_17` | `2.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.4 KiB | [credcheck_17-2.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/credcheck_17-2.8-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_17` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.8 KiB | [credcheck_17-4.7-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.7-1PGDG.rhel10.2.x86_64.rpm) |
| `credcheck_17` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.7 KiB | [credcheck_17-4.7-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.7-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_17` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.1 KiB | [credcheck_17-4.7-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.7-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_17` | `4.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [credcheck_17-4.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.6-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_17` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.1 KiB | [credcheck_17-4.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.5-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_17` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.5 KiB | [credcheck_17-4.5-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.5-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_17` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.6 KiB | [credcheck_17-4.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.4-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_17` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.0 KiB | [credcheck_17-4.4-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.4-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_17` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.5 KiB | [credcheck_17-4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_17` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.8 KiB | [credcheck_17-4.3-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.3-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_17` | `4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.3 KiB | [credcheck_17-4.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.2-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_17` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.8 KiB | [credcheck_17-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_17` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 36.5 KiB | [credcheck_17-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/credcheck_17-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_17` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_17-4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_17` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_17-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_17-4.7-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.7-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_17` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_17-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_17-4.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.6-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_17` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_17-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_17-4.5-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.5-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_17` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [credcheck_17-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [credcheck_17-4.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.4-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_17` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.1 KiB | [credcheck_17-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_17` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.1 KiB | [credcheck_17-4.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.3-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_17` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_17-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_17` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.6 KiB | [credcheck_17-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_17` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.4 KiB | [credcheck_17-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/credcheck_17-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.6 KiB | [postgresql-17-credcheck_5.0-2.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg12+2_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.6 KiB | [postgresql-17-credcheck_5.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg12+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.5 KiB | [postgresql-17-credcheck_5.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.6 KiB | [postgresql-17-credcheck_5.0-2.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg12+2_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.6 KiB | [postgresql-17-credcheck_5.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg12+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.6 KiB | [postgresql-17-credcheck_5.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.8 KiB | [postgresql-17-credcheck_5.0-2.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg13+2_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 81.0 KiB | [postgresql-17-credcheck_5.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg13+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.7 KiB | [postgresql-17-credcheck_5.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.3 KiB | [postgresql-17-credcheck_5.0-2.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg13+2_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.3 KiB | [postgresql-17-credcheck_5.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg13+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.3 KiB | [postgresql-17-credcheck_5.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.9 KiB | [postgresql-17-credcheck_5.0-2.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg22.04+2_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.9 KiB | [postgresql-17-credcheck_5.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.9 KiB | [postgresql-17-credcheck_5.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 80.7 KiB | [postgresql-17-credcheck_5.0-2.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg22.04+2_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 80.6 KiB | [postgresql-17-credcheck_5.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 80.7 KiB | [postgresql-17-credcheck_5.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 74.1 KiB | [postgresql-17-credcheck_5.0-2.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg24.04+2_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 74.0 KiB | [postgresql-17-credcheck_5.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 74.0 KiB | [postgresql-17-credcheck_5.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.8 KiB | [postgresql-17-credcheck_5.0-2.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg24.04+2_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.8 KiB | [postgresql-17-credcheck_5.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.7 KiB | [postgresql-17-credcheck_5.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.6 KiB | [postgresql-17-credcheck_5.0-2.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg26.04+2_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.6 KiB | [postgresql-17-credcheck_5.0-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg26.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.5 KiB | [postgresql-17-credcheck_5.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 71.9 KiB | [postgresql-17-credcheck_5.0-2.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg26.04+2_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 71.9 KiB | [postgresql-17-credcheck_5.0-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-2.pgdg26.04+1_arm64.deb) |
| `postgresql-17-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 71.8 KiB | [postgresql-17-credcheck_5.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-17-credcheck_5.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_16` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.4 KiB | [credcheck_16-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [credcheck_16-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.5 KiB | [credcheck_16-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.9 KiB | [credcheck_16-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.6 KiB | [credcheck_16-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_16` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.0 KiB | [credcheck_16-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.5 KiB | [credcheck_16-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.5 KiB | [credcheck_16-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `2.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.7 KiB | [credcheck_16-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.7-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.3 KiB | [credcheck_16-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.6-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.8 KiB | [credcheck_16-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.8 KiB | [credcheck_16-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/credcheck_16-2.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_16` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.6 KiB | [credcheck_16-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.2 KiB | [credcheck_16-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.8 KiB | [credcheck_16-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.2 KiB | [credcheck_16-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.0 KiB | [credcheck_16-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_16` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.3 KiB | [credcheck_16-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.9 KiB | [credcheck_16-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.1 KiB | [credcheck_16-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `2.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.2 KiB | [credcheck_16-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.7-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.9 KiB | [credcheck_16-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.6-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.5 KiB | [credcheck_16-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [credcheck_16-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/credcheck_16-2.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_16` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.5 KiB | [credcheck_16-4.7-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.7-1PGDG.rhel9.8.x86_64.rpm) |
| `credcheck_16` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.5 KiB | [credcheck_16-4.7-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.7-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_16` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.5 KiB | [credcheck_16-4.7-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.7-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_16` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_16-4.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.6-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_16` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.0 KiB | [credcheck_16-4.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.6-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_16` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_16-4.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.5-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_16` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.0 KiB | [credcheck_16-4.5-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.5-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_16` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.3 KiB | [credcheck_16-4.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.4-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_16` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.4 KiB | [credcheck_16-4.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.4-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_16` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.1 KiB | [credcheck_16-4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_16` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.2 KiB | [credcheck_16-4.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.3-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_16` | `4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.7 KiB | [credcheck_16-4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.2 KiB | [credcheck_16-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.2 KiB | [credcheck_16-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-3.0-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `2.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.1 KiB | [credcheck_16-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-2.7-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.7 KiB | [credcheck_16-2.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-2.6-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.5 KiB | [credcheck_16-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-2.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.3 KiB | [credcheck_16-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/credcheck_16-2.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_16` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.8 KiB | [credcheck_16-4.7-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.7-1PGDG.rhel9.8.aarch64.rpm) |
| `credcheck_16` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.8 KiB | [credcheck_16-4.7-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.7-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_16` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.9 KiB | [credcheck_16-4.7-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.7-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_16` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [credcheck_16-4.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.6-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_16` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.4 KiB | [credcheck_16-4.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.6-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_16` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_16-4.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.5-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_16` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.4 KiB | [credcheck_16-4.5-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.5-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_16` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.0 KiB | [credcheck_16-4.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.4-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_16` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.8 KiB | [credcheck_16-4.4-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.4-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_16` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.5 KiB | [credcheck_16-4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_16` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.6 KiB | [credcheck_16-4.3-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.3-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_16` | `4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.2 KiB | [credcheck_16-4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.8 KiB | [credcheck_16-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.7 KiB | [credcheck_16-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-3.0-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `2.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.8 KiB | [credcheck_16-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-2.7-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.5 KiB | [credcheck_16-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-2.6-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.9 KiB | [credcheck_16-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-2.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.8 KiB | [credcheck_16-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/credcheck_16-2.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_16` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.7 KiB | [credcheck_16-4.7-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.7-1PGDG.rhel10.2.x86_64.rpm) |
| `credcheck_16` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.7 KiB | [credcheck_16-4.7-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.7-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_16` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.1 KiB | [credcheck_16-4.7-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.7-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_16` | `4.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.5 KiB | [credcheck_16-4.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.6-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_16` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.1 KiB | [credcheck_16-4.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.5-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_16` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.5 KiB | [credcheck_16-4.5-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.5-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_16` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.6 KiB | [credcheck_16-4.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.4-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_16` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.0 KiB | [credcheck_16-4.4-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.4-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_16` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.4 KiB | [credcheck_16-4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_16` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.9 KiB | [credcheck_16-4.3-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.3-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_16` | `4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.3 KiB | [credcheck_16-4.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.2-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_16` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.8 KiB | [credcheck_16-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_16` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 36.4 KiB | [credcheck_16-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/credcheck_16-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_16` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_16-4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_16` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_16-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_16-4.7-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.7-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_16` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_16-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_16-4.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.6-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_16` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_16-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [credcheck_16-4.5-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.5-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_16` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [credcheck_16-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.4 KiB | [credcheck_16-4.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.4-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_16` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.1 KiB | [credcheck_16-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_16` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.1 KiB | [credcheck_16-4.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.3-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_16` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_16-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_16` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.6 KiB | [credcheck_16-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_16` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.4 KiB | [credcheck_16-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/credcheck_16-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.8 KiB | [postgresql-16-credcheck_5.0-2.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg12+2_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.8 KiB | [postgresql-16-credcheck_5.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg12+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.7 KiB | [postgresql-16-credcheck_5.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.6 KiB | [postgresql-16-credcheck_5.0-2.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg12+2_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.6 KiB | [postgresql-16-credcheck_5.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg12+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.6 KiB | [postgresql-16-credcheck_5.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.7 KiB | [postgresql-16-credcheck_5.0-2.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg13+2_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.8 KiB | [postgresql-16-credcheck_5.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg13+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.7 KiB | [postgresql-16-credcheck_5.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.3 KiB | [postgresql-16-credcheck_5.0-2.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg13+2_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.3 KiB | [postgresql-16-credcheck_5.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg13+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.3 KiB | [postgresql-16-credcheck_5.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.6 KiB | [postgresql-16-credcheck_5.0-2.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg22.04+2_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.6 KiB | [postgresql-16-credcheck_5.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.4 KiB | [postgresql-16-credcheck_5.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 80.2 KiB | [postgresql-16-credcheck_5.0-2.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg22.04+2_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 80.2 KiB | [postgresql-16-credcheck_5.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 80.2 KiB | [postgresql-16-credcheck_5.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 74.1 KiB | [postgresql-16-credcheck_5.0-2.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg24.04+2_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 74.0 KiB | [postgresql-16-credcheck_5.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.9 KiB | [postgresql-16-credcheck_5.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.8 KiB | [postgresql-16-credcheck_5.0-2.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg24.04+2_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.8 KiB | [postgresql-16-credcheck_5.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.7 KiB | [postgresql-16-credcheck_5.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.6 KiB | [postgresql-16-credcheck_5.0-2.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg26.04+2_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.6 KiB | [postgresql-16-credcheck_5.0-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.4 KiB | [postgresql-16-credcheck_5.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 71.9 KiB | [postgresql-16-credcheck_5.0-2.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg26.04+2_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 71.9 KiB | [postgresql-16-credcheck_5.0-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-2.pgdg26.04+1_arm64.deb) |
| `postgresql-16-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 71.9 KiB | [postgresql-16-credcheck_5.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-16-credcheck_5.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_15` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.5 KiB | [credcheck_15-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [credcheck_15-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.6 KiB | [credcheck_15-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.0 KiB | [credcheck_15-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.8 KiB | [credcheck_15-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_15` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.2 KiB | [credcheck_15-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.6 KiB | [credcheck_15-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.6 KiB | [credcheck_15-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.7 KiB | [credcheck_15-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.7-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.4 KiB | [credcheck_15-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.6-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 33.0 KiB | [credcheck_15-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.9 KiB | [credcheck_15-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.1 KiB | [credcheck_15-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-2.0-1.rhel8.x86_64.rpm) |
| `credcheck_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.7 KiB | [credcheck_15-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-1.2-1.rhel8.x86_64.rpm) |
| `credcheck_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.1 KiB | [credcheck_15-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-1.0-1.rhel8.x86_64.rpm) |
| `credcheck_15` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.6 KiB | [credcheck_15-0.2.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-0.2.0-3.rhel8.x86_64.rpm) |
| `credcheck_15` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.0 KiB | [credcheck_15-0.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/credcheck_15-0.2.0-1.rhel8.x86_64.rpm) |
| `credcheck_15` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.5 KiB | [credcheck_15-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.1 KiB | [credcheck_15-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.8 KiB | [credcheck_15-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.1 KiB | [credcheck_15-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.9 KiB | [credcheck_15-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_15` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.2 KiB | [credcheck_15-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.8 KiB | [credcheck_15-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.0 KiB | [credcheck_15-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.2 KiB | [credcheck_15-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.7-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.9 KiB | [credcheck_15-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.6-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.5 KiB | [credcheck_15-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.3 KiB | [credcheck_15-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.5 KiB | [credcheck_15-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-2.0-1.rhel8.aarch64.rpm) |
| `credcheck_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.2 KiB | [credcheck_15-1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-1.2-1.rhel8.aarch64.rpm) |
| `credcheck_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.6 KiB | [credcheck_15-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-1.0-1.rhel8.aarch64.rpm) |
| `credcheck_15` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.3 KiB | [credcheck_15-0.2.0-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-0.2.0-3.rhel8.aarch64.rpm) |
| `credcheck_15` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.9 KiB | [credcheck_15-0.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/credcheck_15-0.2.0-1.rhel8.aarch64.rpm) |
| `credcheck_15` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.5 KiB | [credcheck_15-4.7-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.7-1PGDG.rhel9.8.x86_64.rpm) |
| `credcheck_15` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.5 KiB | [credcheck_15-4.7-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.7-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_15` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.6 KiB | [credcheck_15-4.7-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.7-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_15` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_15-4.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.6-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_15` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.0 KiB | [credcheck_15-4.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.6-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_15` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_15-4.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.5-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_15` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.0 KiB | [credcheck_15-4.5-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.5-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_15` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.3 KiB | [credcheck_15-4.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.4-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_15` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.4 KiB | [credcheck_15-4.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.4-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_15` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.0 KiB | [credcheck_15-4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_15` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.2 KiB | [credcheck_15-4.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.3-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_15` | `4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.7 KiB | [credcheck_15-4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.3 KiB | [credcheck_15-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.1 KiB | [credcheck_15-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-3.0-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `2.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.2 KiB | [credcheck_15-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.7-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.9 KiB | [credcheck_15-2.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.6-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.4 KiB | [credcheck_15-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.5 KiB | [credcheck_15-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_15` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.6 KiB | [credcheck_15-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-2.0-1.rhel9.x86_64.rpm) |
| `credcheck_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.1 KiB | [credcheck_15-1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-1.2-1.rhel9.x86_64.rpm) |
| `credcheck_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.5 KiB | [credcheck_15-1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-1.0-1.rhel9.x86_64.rpm) |
| `credcheck_15` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.8 KiB | [credcheck_15-0.2.0-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-0.2.0-3.rhel9.x86_64.rpm) |
| `credcheck_15` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.9 KiB | [credcheck_15-0.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/credcheck_15-0.2.0-1.rhel9.x86_64.rpm) |
| `credcheck_15` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.7 KiB | [credcheck_15-4.7-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.7-1PGDG.rhel9.8.aarch64.rpm) |
| `credcheck_15` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.7 KiB | [credcheck_15-4.7-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.7-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_15` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.8 KiB | [credcheck_15-4.7-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.7-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_15` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_15-4.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.6-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_15` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [credcheck_15-4.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.6-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_15` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_15-4.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.5-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_15` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [credcheck_15-4.5-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.5-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_15` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.6 KiB | [credcheck_15-4.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.4-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_15` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [credcheck_15-4.4-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.4-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_15` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.6 KiB | [credcheck_15-4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_15` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.7 KiB | [credcheck_15-4.3-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.3-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_15` | `4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.9 KiB | [credcheck_15-4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.7 KiB | [credcheck_15-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.8 KiB | [credcheck_15-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-3.0-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `2.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.8 KiB | [credcheck_15-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.7-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.5 KiB | [credcheck_15-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.6-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.9 KiB | [credcheck_15-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.8 KiB | [credcheck_15-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_15` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.9 KiB | [credcheck_15-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-2.0-1.rhel9.aarch64.rpm) |
| `credcheck_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.5 KiB | [credcheck_15-1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-1.2-1.rhel9.aarch64.rpm) |
| `credcheck_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.9 KiB | [credcheck_15-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-1.0-1.rhel9.aarch64.rpm) |
| `credcheck_15` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.1 KiB | [credcheck_15-0.2.0-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-0.2.0-3.rhel9.aarch64.rpm) |
| `credcheck_15` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.5 KiB | [credcheck_15-0.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/credcheck_15-0.2.0-1.rhel9.aarch64.rpm) |
| `credcheck_15` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [credcheck_15-4.7-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.7-1PGDG.rhel10.2.x86_64.rpm) |
| `credcheck_15` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [credcheck_15-4.7-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.7-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_15` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.1 KiB | [credcheck_15-4.7-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.7-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_15` | `4.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [credcheck_15-4.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.6-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_15` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.1 KiB | [credcheck_15-4.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.5-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_15` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [credcheck_15-4.5-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.5-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_15` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.7 KiB | [credcheck_15-4.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.4-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_15` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.1 KiB | [credcheck_15-4.4-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.4-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_15` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.4 KiB | [credcheck_15-4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_15` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.8 KiB | [credcheck_15-4.3-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.3-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_15` | `4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.3 KiB | [credcheck_15-4.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.2-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_15` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.8 KiB | [credcheck_15-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_15` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 36.5 KiB | [credcheck_15-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/credcheck_15-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_15` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_15-4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_15` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_15-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.2 KiB | [credcheck_15-4.7-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.7-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_15` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_15-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_15-4.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.6-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_15` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_15-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [credcheck_15-4.5-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.5-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_15` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_15-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_15-4.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.4-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_15` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_15-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_15` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.0 KiB | [credcheck_15-4.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.3-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_15` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.9 KiB | [credcheck_15-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_15` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.5 KiB | [credcheck_15-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_15` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.5 KiB | [credcheck_15-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/credcheck_15-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.4 KiB | [postgresql-15-credcheck_5.0-2.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg12+2_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.4 KiB | [postgresql-15-credcheck_5.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg12+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 80.3 KiB | [postgresql-15-credcheck_5.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.3 KiB | [postgresql-15-credcheck_5.0-2.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg12+2_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.3 KiB | [postgresql-15-credcheck_5.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg12+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 79.2 KiB | [postgresql-15-credcheck_5.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.4 KiB | [postgresql-15-credcheck_5.0-2.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg13+2_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.4 KiB | [postgresql-15-credcheck_5.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg13+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 80.5 KiB | [postgresql-15-credcheck_5.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.1 KiB | [postgresql-15-credcheck_5.0-2.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg13+2_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.1 KiB | [postgresql-15-credcheck_5.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg13+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 79.0 KiB | [postgresql-15-credcheck_5.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.2 KiB | [postgresql-15-credcheck_5.0-2.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg22.04+2_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.2 KiB | [postgresql-15-credcheck_5.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 81.3 KiB | [postgresql-15-credcheck_5.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 79.8 KiB | [postgresql-15-credcheck_5.0-2.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg22.04+2_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 79.8 KiB | [postgresql-15-credcheck_5.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 79.8 KiB | [postgresql-15-credcheck_5.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.8 KiB | [postgresql-15-credcheck_5.0-2.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg24.04+2_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.8 KiB | [postgresql-15-credcheck_5.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 73.6 KiB | [postgresql-15-credcheck_5.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.5 KiB | [postgresql-15-credcheck_5.0-2.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg24.04+2_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.5 KiB | [postgresql-15-credcheck_5.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 72.4 KiB | [postgresql-15-credcheck_5.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.3 KiB | [postgresql-15-credcheck_5.0-2.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg26.04+2_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.3 KiB | [postgresql-15-credcheck_5.0-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg26.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 73.2 KiB | [postgresql-15-credcheck_5.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 71.7 KiB | [postgresql-15-credcheck_5.0-2.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg26.04+2_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 71.7 KiB | [postgresql-15-credcheck_5.0-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-2.pgdg26.04+1_arm64.deb) |
| `postgresql-15-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 71.6 KiB | [postgresql-15-credcheck_5.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-15-credcheck_5.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `credcheck_14` | `4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 42.4 KiB | [credcheck_14-4.7-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.7-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_14` | `4.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.9 KiB | [credcheck_14-4.6-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.6-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_14` | `4.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.6 KiB | [credcheck_14-4.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.5-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_14` | `4.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.0 KiB | [credcheck_14-4.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.4-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_14` | `4.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.7 KiB | [credcheck_14-4.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.3-1PGDG.rhel8.10.x86_64.rpm) |
| `credcheck_14` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 40.1 KiB | [credcheck_14-4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 39.6 KiB | [credcheck_14-4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-4.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.6 KiB | [credcheck_14-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-3.0-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `2.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.6 KiB | [credcheck_14-2.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.7-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 34.3 KiB | [credcheck_14-2.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.6-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 32.9 KiB | [credcheck_14-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.2-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.8 KiB | [credcheck_14-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.1-1PGDG.rhel8.x86_64.rpm) |
| `credcheck_14` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.0 KiB | [credcheck_14-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-2.0-1.rhel8.x86_64.rpm) |
| `credcheck_14` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.7 KiB | [credcheck_14-1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-1.2-1.rhel8.x86_64.rpm) |
| `credcheck_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 27.1 KiB | [credcheck_14-1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-1.0-1.rhel8.x86_64.rpm) |
| `credcheck_14` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.6 KiB | [credcheck_14-0.2.0-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-0.2.0-3.rhel8.x86_64.rpm) |
| `credcheck_14` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 35.3 KiB | [credcheck_14-0.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/credcheck_14-0.2.0-1.rhel8.x86_64.rpm) |
| `credcheck_14` | `4.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.4 KiB | [credcheck_14-4.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.7-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_14` | `4.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 41.0 KiB | [credcheck_14-4.6-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.6-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_14` | `4.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.7 KiB | [credcheck_14-4.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.5-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_14` | `4.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 40.0 KiB | [credcheck_14-4.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.4-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_14` | `4.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.8 KiB | [credcheck_14-4.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.3-1PGDG.rhel8.10.aarch64.rpm) |
| `credcheck_14` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 39.1 KiB | [credcheck_14-4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 38.7 KiB | [credcheck_14-4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-4.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 35.0 KiB | [credcheck_14-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-3.0-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `2.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.1 KiB | [credcheck_14-2.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.7-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 33.8 KiB | [credcheck_14-2.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.6-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 32.4 KiB | [credcheck_14-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.2-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.2 KiB | [credcheck_14-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.1-1PGDG.rhel8.aarch64.rpm) |
| `credcheck_14` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.4 KiB | [credcheck_14-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-2.0-1.rhel8.aarch64.rpm) |
| `credcheck_14` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 27.2 KiB | [credcheck_14-1.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-1.2-1.rhel8.aarch64.rpm) |
| `credcheck_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.6 KiB | [credcheck_14-1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-1.0-1.rhel8.aarch64.rpm) |
| `credcheck_14` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.3 KiB | [credcheck_14-0.2.0-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-0.2.0-3.rhel8.aarch64.rpm) |
| `credcheck_14` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 34.8 KiB | [credcheck_14-0.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/credcheck_14-0.2.0-1.rhel8.aarch64.rpm) |
| `credcheck_14` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.4 KiB | [credcheck_14-4.7-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.7-1PGDG.rhel9.8.x86_64.rpm) |
| `credcheck_14` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.4 KiB | [credcheck_14-4.7-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.7-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_14` | `4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.6 KiB | [credcheck_14-4.7-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.7-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_14` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_14-4.6-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.6-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_14` | `4.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 41.0 KiB | [credcheck_14-4.6-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.6-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_14` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_14-4.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.5-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_14` | `4.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.9 KiB | [credcheck_14-4.5-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.5-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_14` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.2 KiB | [credcheck_14-4.4-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.4-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_14` | `4.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.3 KiB | [credcheck_14-4.4-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.4-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_14` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.0 KiB | [credcheck_14-4.3-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.3-1PGDG.rhel9.7.x86_64.rpm) |
| `credcheck_14` | `4.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 40.1 KiB | [credcheck_14-4.3-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.3-1PGDG.rhel9.6.x86_64.rpm) |
| `credcheck_14` | `4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.7 KiB | [credcheck_14-4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 39.3 KiB | [credcheck_14-4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-4.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 36.0 KiB | [credcheck_14-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-3.0-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `2.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 35.1 KiB | [credcheck_14-2.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.7-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `2.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 34.8 KiB | [credcheck_14-2.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.6-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.3 KiB | [credcheck_14-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.2-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.2 KiB | [credcheck_14-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.1-1PGDG.rhel9.x86_64.rpm) |
| `credcheck_14` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.3 KiB | [credcheck_14-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-2.0-1.rhel9.x86_64.rpm) |
| `credcheck_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.0 KiB | [credcheck_14-1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-1.2-1.rhel9.x86_64.rpm) |
| `credcheck_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 27.4 KiB | [credcheck_14-1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-1.0-1.rhel9.x86_64.rpm) |
| `credcheck_14` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.8 KiB | [credcheck_14-0.2.0-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/credcheck_14-0.2.0-3.rhel9.x86_64.rpm) |
| `credcheck_14` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.6 KiB | [credcheck_14-4.7-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.7-1PGDG.rhel9.8.aarch64.rpm) |
| `credcheck_14` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.6 KiB | [credcheck_14-4.7-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.7-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_14` | `4.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.7 KiB | [credcheck_14-4.7-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.7-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_14` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_14-4.6-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.6-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_14` | `4.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [credcheck_14-4.6-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.6-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_14` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.2 KiB | [credcheck_14-4.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.5-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_14` | `4.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [credcheck_14-4.5-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.5-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_14` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.5 KiB | [credcheck_14-4.4-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.4-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_14` | `4.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.9 KiB | [credcheck_14-4.4-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.4-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_14` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.5 KiB | [credcheck_14-4.3-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.3-1PGDG.rhel9.7.aarch64.rpm) |
| `credcheck_14` | `4.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.6 KiB | [credcheck_14-4.3-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.3-1PGDG.rhel9.6.aarch64.rpm) |
| `credcheck_14` | `4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 39.0 KiB | [credcheck_14-4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 38.6 KiB | [credcheck_14-4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-4.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.6 KiB | [credcheck_14-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-3.0-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `2.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.8 KiB | [credcheck_14-2.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.7-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `2.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 34.4 KiB | [credcheck_14-2.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.6-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 32.8 KiB | [credcheck_14-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.2-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.7 KiB | [credcheck_14-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.1-1PGDG.rhel9.aarch64.rpm) |
| `credcheck_14` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.8 KiB | [credcheck_14-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-2.0-1.rhel9.aarch64.rpm) |
| `credcheck_14` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 27.4 KiB | [credcheck_14-1.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-1.2-1.rhel9.aarch64.rpm) |
| `credcheck_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.8 KiB | [credcheck_14-1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-1.0-1.rhel9.aarch64.rpm) |
| `credcheck_14` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.0 KiB | [credcheck_14-0.2.0-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-0.2.0-3.rhel9.aarch64.rpm) |
| `credcheck_14` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 35.4 KiB | [credcheck_14-0.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/credcheck_14-0.2.0-1.rhel9.aarch64.rpm) |
| `credcheck_14` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [credcheck_14-4.7-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.7-1PGDG.rhel10.2.x86_64.rpm) |
| `credcheck_14` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.6 KiB | [credcheck_14-4.7-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.7-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_14` | `4.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 42.1 KiB | [credcheck_14-4.7-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.7-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_14` | `4.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.5 KiB | [credcheck_14-4.6-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.6-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_14` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.1 KiB | [credcheck_14-4.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.5-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_14` | `4.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.5 KiB | [credcheck_14-4.5-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.5-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_14` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.6 KiB | [credcheck_14-4.4-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.4-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_14` | `4.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 41.0 KiB | [credcheck_14-4.4-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.4-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_14` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.4 KiB | [credcheck_14-4.3-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.3-1PGDG.rhel10.1.x86_64.rpm) |
| `credcheck_14` | `4.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.7 KiB | [credcheck_14-4.3-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.3-1PGDG.rhel10.0.x86_64.rpm) |
| `credcheck_14` | `4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 40.3 KiB | [credcheck_14-4.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.2-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_14` | `4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 39.8 KiB | [credcheck_14-4.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-4.1-1PGDG.rhel10.x86_64.rpm) |
| `credcheck_14` | `3.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 36.5 KiB | [credcheck_14-3.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/credcheck_14-3.0-2PGDG.rhel10.x86_64.rpm) |
| `credcheck_14` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.0 KiB | [credcheck_14-4.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.7-1PGDG.rhel10.2.aarch64.rpm) |
| `credcheck_14` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.0 KiB | [credcheck_14-4.7-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.7-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_14` | `4.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.0 KiB | [credcheck_14-4.7-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.7-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_14` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.5 KiB | [credcheck_14-4.6-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.6-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_14` | `4.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.5 KiB | [credcheck_14-4.6-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.6-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_14` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.5 KiB | [credcheck_14-4.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.5-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_14` | `4.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.5 KiB | [credcheck_14-4.5-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.5-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_14` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.9 KiB | [credcheck_14-4.4-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.4-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_14` | `4.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.9 KiB | [credcheck_14-4.4-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.4-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_14` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.9 KiB | [credcheck_14-4.3-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.3-1PGDG.rhel10.1.aarch64.rpm) |
| `credcheck_14` | `4.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.9 KiB | [credcheck_14-4.3-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.3-1PGDG.rhel10.0.aarch64.rpm) |
| `credcheck_14` | `4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.6 KiB | [credcheck_14-4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.2-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_14` | `4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.2 KiB | [credcheck_14-4.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-4.1-1PGDG.rhel10.aarch64.rpm) |
| `credcheck_14` | `3.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 36.4 KiB | [credcheck_14-3.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/credcheck_14-3.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 75.8 KiB | [postgresql-14-credcheck_5.0-2.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg12+2_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 75.8 KiB | [postgresql-14-credcheck_5.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg12+1_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 75.5 KiB | [postgresql-14-credcheck_5.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 74.4 KiB | [postgresql-14-credcheck_5.0-2.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg12+2_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 74.4 KiB | [postgresql-14-credcheck_5.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg12+1_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 74.4 KiB | [postgresql-14-credcheck_5.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 75.6 KiB | [postgresql-14-credcheck_5.0-2.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg13+2_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 75.6 KiB | [postgresql-14-credcheck_5.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg13+1_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 75.3 KiB | [postgresql-14-credcheck_5.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 74.1 KiB | [postgresql-14-credcheck_5.0-2.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg13+2_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 74.1 KiB | [postgresql-14-credcheck_5.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg13+1_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 74.0 KiB | [postgresql-14-credcheck_5.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 75.8 KiB | [postgresql-14-credcheck_5.0-2.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg22.04+2_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 75.8 KiB | [postgresql-14-credcheck_5.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 75.7 KiB | [postgresql-14-credcheck_5.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.2 KiB | [postgresql-14-credcheck_5.0-2.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg22.04+2_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.1 KiB | [postgresql-14-credcheck_5.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 74.0 KiB | [postgresql-14-credcheck_5.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 69.0 KiB | [postgresql-14-credcheck_5.0-2.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg24.04+2_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 69.0 KiB | [postgresql-14-credcheck_5.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 68.9 KiB | [postgresql-14-credcheck_5.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 67.8 KiB | [postgresql-14-credcheck_5.0-2.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg24.04+2_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 67.7 KiB | [postgresql-14-credcheck_5.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 67.7 KiB | [postgresql-14-credcheck_5.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 68.5 KiB | [postgresql-14-credcheck_5.0-2.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg26.04+2_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 68.4 KiB | [postgresql-14-credcheck_5.0-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg26.04+1_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 68.3 KiB | [postgresql-14-credcheck_5.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 66.9 KiB | [postgresql-14-credcheck_5.0-2.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg26.04+2_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 66.9 KiB | [postgresql-14-credcheck_5.0-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-2.pgdg26.04+1_arm64.deb) |
| `postgresql-14-credcheck` | `5.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 66.9 KiB | [postgresql-14-credcheck_5.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/c/credcheck/postgresql-14-credcheck_5.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HexaCluster/credcheck" title="Repository" icon="github" subtitle="github.com/HexaCluster/credcheck" />}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install credcheck;		# install via package name, for the active PG version

pig install credcheck -v 18;   # install for PG 18
pig install credcheck -v 17;   # install for PG 17
pig install credcheck -v 16;   # install for PG 16
pig install credcheck -v 15;   # install for PG 15
pig install credcheck -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'credcheck';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION credcheck;
```

## Usage

Sources:

- [v5.0 README](https://github.com/HexaCluster/credcheck/blob/v5.0/README.md)
- [v5.0 changelog](https://github.com/HexaCluster/credcheck/blob/v5.0/ChangeLog)
- [SQL objects 5.0.0](https://github.com/HexaCluster/credcheck/blob/v5.0/sql/credcheck--5.0.0.sql)
- [Password history WAL implementation](https://github.com/HexaCluster/credcheck/blob/v5.0/credcheck.c)
- [Login-event setup](https://github.com/HexaCluster/credcheck/blob/v5.0/event_trigger.sql)

`credcheck` enforces username and plaintext-password rules during role creation, password changes and role renames. It also tracks password reuse, bans repeated authentication failures and can require a password change at first login. Configure policy as a superuser; the defaults do not enforce a comprehensive password-strength policy.

### Enable and Set Policy

Add the library to the existing preload list and restart PostgreSQL. Install SQL objects in each database where administrators need its views and reset functions:

```ini
shared_preload_libraries = 'credcheck'
credcheck.password_min_length = 12
credcheck.password_contain_username = on
credcheck.password_reuse_history = 2
credcheck.password_reuse_interval = 365
```

```sql
CREATE EXTENSION credcheck;
CREATE ROLE app_user LOGIN PASSWORD 'example-Strong-Pass#123';
SELECT rolename, password_date FROM pg_password_history;
```

The interval is in days. Installing the SQL extension is separate from loading the server-wide hooks. Upstream release 5.0 uses SQL extension version 5.0.0; installing upgraded files requires a restart to reload the library.

### Policy Index

| Settings | Purpose |
|---|---|
| `credcheck.username_min_length`, `credcheck.username_min_special`, `credcheck.username_min_digit`, `credcheck.username_min_upper`, `credcheck.username_min_lower` | Username length and character requirements |
| `credcheck.password_min_length`, `credcheck.password_min_special`, `credcheck.password_min_digit`, `credcheck.password_min_upper`, `credcheck.password_min_lower` | Password length and character requirements |
| `credcheck.username_min_repeat`, `credcheck.password_min_repeat` | Maximum adjacent repetitions, despite the parameter names |
| `credcheck.username_contain`, `credcheck.username_not_contain`, `credcheck.password_contain`, `credcheck.password_not_contain` | Required or forbidden content |
| `credcheck.username_contain_password`, `credcheck.password_contain_username` | Reject credentials containing one another |
| `credcheck.username_ignore_case`, `credcheck.password_ignore_case` | Case handling |
| `credcheck.password_min_length_su`, `credcheck.password_valid_until_su` | Separate superuser requirements |
| `credcheck.password_valid_until`, `credcheck.password_valid_max` | Minimum and maximum password lifetime; the minimum also supplies an omitted expiry when a password is changed |
| `credcheck.whitelist`, `credcheck.superuser_nocheck` | Explicit policy exemptions |
| `credcheck.no_password_logging` | Suppress passwords in policy-error logs; enabled by default |

CrackLib strength checking is available only when the library was built with that support and its dictionary is available.

### Password History and Replication

History contains SHA-256 password hashes, is shared across databases, and is persisted in `$PGDATA/pg_password_history`. Include that file in backup planning and protect access to the SQL history view, which is granted to PUBLIC by default. `credcheck.history_max_size` changes the shared-memory capacity and requires a restart.

Version 5.0 replicates history changes through custom WAL resource manager ID 150 on PostgreSQL 15 and later. Keep the matching library preloaded on replicas and recovery servers that replay its WAL. Earlier PostgreSQL versions retain the file-backed history path without this replication support. History-reset and timestamp-test functions reject execution during recovery.

```sql
SELECT pg_password_history_reset('app_user');
```

Resetting history removes reuse protection for those records; reserve it for administrators.

### Authentication and Password Changes

```ini
credcheck.max_auth_failure = 3
credcheck.auth_delay_ms = 1000
credcheck.whitelist_auth_failure = 'service_user'
credcheck.password_change_first_login = true
```

```sql
SELECT * FROM pg_banned_role;
SELECT pg_banned_role_reset('app_user');
ALTER ROLE app_user SET credcheck_internal.force_change_password = true;
```

Bans remain until reset and their cache is lost at restart. `credcheck.reset_superuser` provides the documented superuser recovery path; `credcheck.auth_failure_cache_size` requires a restart.

Use the actual parameter `credcheck.disallow_change_password` to prohibit password changes. Even superusers are affected unless they enable `credcheck.superuser_nocheck` in their session. This exemption bypasses all corresponding role checks and must be controlled.

`credcheck.password_valid_warning` needs PostgreSQL 17 or later and the official login event trigger installed separately in every relevant database; SQL extension creation does not install that trigger.

### Plaintext Boundary

Strength and reuse checks need plaintext at password-change time. Already hashed passwords are rejected by default, including passwords sent by psql's `\password`. Setting `credcheck.encrypted_password_allowed` accepts them without providing equivalent plaintext checks. Protect the password-change connection and do not assume existing credentials are scanned retroactively. Username checks are skipped when creating a role without a password or renaming a role with no password.
