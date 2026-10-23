---
title: "pg_lease"
linkTitle: "pg_lease"
description: "SQL leases with TTL expiry and monotonically increasing fencing epochs"
weight: 4145
categories: ["UTIL"]
languages: ["SQL"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_lease**](https://pgxn.org/dist/pg_lease/0.1.0/) : SQL leases with TTL expiry and monotonically increasing fencing epochs


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4145** | {{< badge content="pg_lease" link="https://pgxn.org/dist/pg_lease/0.1.0/" >}} | {{< ext "pg_lease" >}} | `0.1.0` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `lease` |
|   **See Also**    | {{< ext "pglock" >}} {{< ext "pg_retry" >}} {{< ext "pgmq" >}} |

> [!Note] PGXN testing release. Requires PG14+; upstream tests PG14/15, local EL9 aarch64 regression and isolation tests pass PG14-18. No failover fencing guarantee; owner strings are application identities. PGXN-listed GitHub repository azghr/pg-primitive returned 404 on 2026-10-23; PGXN hosts the verified source.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_lease` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_lease_18" "green" >}} {{< bg "17" "pg_lease_17" "green" >}} {{< bg "16" "pg_lease_16" "green" >}} {{< bg "15" "pg_lease_15" "green" >}} {{< bg "14" "pg_lease_14" "green" >}} | `pg_lease_$v` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_lease_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} | {{< bg "N/A" "pg_lease : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lease_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.0 KiB | [pg_lease_18-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_lease_18-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_lease_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.0 KiB | [pg_lease_18-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_lease_18-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_lease_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 22.8 KiB | [pg_lease_18-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lease_18-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_lease_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.8 KiB | [pg_lease_18-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lease_18-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_lease_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 22.9 KiB | [pg_lease_18-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lease_18-0.1.0-1PGSTY.el10.noarch.rpm) |
| `pg_lease_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.9 KiB | [pg_lease_18-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lease_18-0.1.0-1PGSTY.el10.noarch.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lease_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.0 KiB | [pg_lease_17-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_lease_17-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_lease_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.0 KiB | [pg_lease_17-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_lease_17-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_lease_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 22.8 KiB | [pg_lease_17-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lease_17-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_lease_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.8 KiB | [pg_lease_17-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lease_17-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_lease_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 22.9 KiB | [pg_lease_17-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lease_17-0.1.0-1PGSTY.el10.noarch.rpm) |
| `pg_lease_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.9 KiB | [pg_lease_17-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lease_17-0.1.0-1PGSTY.el10.noarch.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lease_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.0 KiB | [pg_lease_16-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_lease_16-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_lease_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.0 KiB | [pg_lease_16-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_lease_16-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_lease_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 22.8 KiB | [pg_lease_16-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lease_16-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_lease_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.8 KiB | [pg_lease_16-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lease_16-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_lease_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 22.9 KiB | [pg_lease_16-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lease_16-0.1.0-1PGSTY.el10.noarch.rpm) |
| `pg_lease_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.9 KiB | [pg_lease_16-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lease_16-0.1.0-1PGSTY.el10.noarch.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lease_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.0 KiB | [pg_lease_15-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_lease_15-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_lease_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.0 KiB | [pg_lease_15-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_lease_15-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_lease_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 22.8 KiB | [pg_lease_15-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lease_15-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_lease_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.8 KiB | [pg_lease_15-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lease_15-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_lease_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 22.9 KiB | [pg_lease_15-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lease_15-0.1.0-1PGSTY.el10.noarch.rpm) |
| `pg_lease_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.9 KiB | [pg_lease_15-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lease_15-0.1.0-1PGSTY.el10.noarch.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_lease_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.0 KiB | [pg_lease_14-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_lease_14-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_lease_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.0 KiB | [pg_lease_14-0.1.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_lease_14-0.1.0-1PGSTY.el8.noarch.rpm) |
| `pg_lease_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 22.8 KiB | [pg_lease_14-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_lease_14-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_lease_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.8 KiB | [pg_lease_14-0.1.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_lease_14-0.1.0-1PGSTY.el9.noarch.rpm) |
| `pg_lease_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 22.9 KiB | [pg_lease_14-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_lease_14-0.1.0-1PGSTY.el10.noarch.rpm) |
| `pg_lease_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.9 KiB | [pg_lease_14-0.1.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_lease_14-0.1.0-1PGSTY.el10.noarch.rpm) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://pgxn.org/dist/pg_lease/0.1.0/" title="Repository" icon="link" subtitle="pgxn.org/dist/pg_lease/0.1.0/" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_lease-0.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_lease;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_lease;		# install via package name, for the active PG version

pig install pg_lease -v 18;   # install for PG 18
pig install pg_lease -v 17;   # install for PG 17
pig install pg_lease -v 16;   # install for PG 16
pig install pg_lease -v 15;   # install for PG 15
pig install pg_lease -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_lease;
```

## Usage

Sources:

- [PGXN 0.1.0](https://pgxn.org/dist/pg_lease/0.1.0/)
- [Official README](https://api.pgxn.org/src/pg_lease/pg_lease-0.1.0/README.md)
- [Extension SQL](https://api.pgxn.org/src/pg_lease/pg_lease-0.1.0/sql/pg_lease--0.1.0.sql)
- [Control file](https://api.pgxn.org/src/pg_lease/pg_lease-0.1.0/pg_lease.control)

`pg_lease` provides expiring ownership of named resources with increasing fencing epochs. It is a pure SQL/PL/pgSQL extension with lazy expiry, no background worker, and no preload requirement. Version 0.1.0 is marked testing on PGXN; upstream requires PostgreSQL 14 or later and documents its compatibility tests on PostgreSQL 14 and 15.

### Core Workflow

The extension is trusted: a role with CREATE privilege on the database can install it. It creates the fixed `lease` schema and is not relocatable.

```sql
CREATE EXTENSION pg_lease;

SELECT acquired, epoch, expires_at
FROM lease.acquire('job-42', 'worker-a', interval '30 seconds');

SELECT status, expires_at
FROM lease.renew('job-42', 'worker-a', 1, interval '30 seconds');

SELECT status FROM lease.release('job-42', 'worker-a', 1);
SELECT held, owner, epoch, expires_at FROM lease.inspect('job-42');
```

The example uses the first epoch, 1. Applications must pass the epoch actually returned by acquisition to later renewal, release, and external resource guards. A successful release advances the epoch; re-acquisition must use the newly returned value.

### Operations

| Operation | Behavior |
|---|---|
| `lease.acquire(key, owner, ttl, wait)` | Acquires a free or expired key. The optional wait defaults to zero; a busy key returns a denial. Positive waits poll until acquisition or timeout. |
| `lease.renew(key, owner, epoch, ttl)` | Extends an active lease without changing its epoch. |
| `lease.release(key, owner, epoch)` | Releases an active lease and advances its epoch. |
| `lease.inspect(key)` | Observes ownership, expiry, and the current fencing counter; an unknown key returns no rows. |

Renewal and release return `ok`, `not_owner`, `lapsed`, or `epoch_mismatch`. Invalid inputs raise SQLSTATE `22023`; an exhausted blocking wait raises `57014`. Re-acquiring an unexpired lease with the same owner is idempotent and does not extend its expiry.

### Transactions and Boundaries

Use READ COMMITTED for blocking acquisition. State changes follow the caller's transaction and are undone by rollback. Expiry follows the database server's clock; session termination does not release a lease immediately. The functions use `SECURITY DEFINER` with a fixed search path; direct public access to `lease.leases` is revoked.

Owner strings are application-supplied identifiers, not authenticated database identities. Applications must control access to the lease API and enforce fencing epochs at the resource being protected. The extension provides no safety guarantee across primary promotion or independent databases. Dropping the extension destroys lease state and resets the epoch history on reinstallation; version 0.1.0 does not register its lease table for extension configuration dumps.
