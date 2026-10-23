---
title: "codesystem"
linkTitle: "codesystem"
description: "Hierarchical classification code types with validation and revision tracking"
weight: 3540
categories: ["TYPE"]
languages: ["C"]
licenses: ["BSD-2-Clause"]
repos: ["PIGSTY"]
page_width: full
---

[**codesystem**](https://github.com/John-Cameron-Burns/pg-codesystem) : Hierarchical classification code types with validation and revision tracking


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3540** | {{< badge content="codesystem" link="https://github.com/John-Cameron-Burns/pg-codesystem" >}} | {{< ext "codesystem" >}} | `0.2.0` | {{< category "TYPE" >}} | {{< license "BSD-2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] Distribution 0.2.1 keeps SQL/control 0.2.0; no extension upgrade from 0.2.0 is needed. Optional GSS entity updates and publisher-fetched ISIC/ISCO loaders need plpython3u, superuser installation and network access. Core extension dependencies unchanged. Upstream CI covers PG13-17; local EL9 aarch64 builds and regression checks additionally validate PG18.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `codesystem` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.1` | {{< bg "18" "codesystem_18" "green" >}} {{< bg "17" "codesystem_17" "green" >}} {{< bg "16" "codesystem_16" "green" >}} {{< bg "15" "codesystem_15" "green" >}} {{< bg "14" "codesystem_14" "green" >}} | `codesystem_$v` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.1" "codesystem_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} | {{< bg "N/A" "codesystem : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `codesystem_18` | `0.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 122.3 KiB | [codesystem_18-0.2.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/codesystem_18-0.2.1-1PGSTY.el8.x86_64.rpm) |
| `codesystem_18` | `0.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 120.2 KiB | [codesystem_18-0.2.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/codesystem_18-0.2.1-1PGSTY.el8.aarch64.rpm) |
| `codesystem_18` | `0.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 122.5 KiB | [codesystem_18-0.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/codesystem_18-0.2.1-1PGSTY.el9.x86_64.rpm) |
| `codesystem_18` | `0.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 121.4 KiB | [codesystem_18-0.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/codesystem_18-0.2.1-1PGSTY.el9.aarch64.rpm) |
| `codesystem_18` | `0.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 122.8 KiB | [codesystem_18-0.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/codesystem_18-0.2.1-1PGSTY.el10.x86_64.rpm) |
| `codesystem_18` | `0.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 122.2 KiB | [codesystem_18-0.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/codesystem_18-0.2.1-1PGSTY.el10.aarch64.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `codesystem_17` | `0.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 122.2 KiB | [codesystem_17-0.2.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/codesystem_17-0.2.1-1PGSTY.el8.x86_64.rpm) |
| `codesystem_17` | `0.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 120.2 KiB | [codesystem_17-0.2.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/codesystem_17-0.2.1-1PGSTY.el8.aarch64.rpm) |
| `codesystem_17` | `0.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 122.5 KiB | [codesystem_17-0.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/codesystem_17-0.2.1-1PGSTY.el9.x86_64.rpm) |
| `codesystem_17` | `0.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 121.4 KiB | [codesystem_17-0.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/codesystem_17-0.2.1-1PGSTY.el9.aarch64.rpm) |
| `codesystem_17` | `0.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 122.9 KiB | [codesystem_17-0.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/codesystem_17-0.2.1-1PGSTY.el10.x86_64.rpm) |
| `codesystem_17` | `0.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 122.2 KiB | [codesystem_17-0.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/codesystem_17-0.2.1-1PGSTY.el10.aarch64.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `codesystem_16` | `0.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 122.2 KiB | [codesystem_16-0.2.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/codesystem_16-0.2.1-1PGSTY.el8.x86_64.rpm) |
| `codesystem_16` | `0.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 120.2 KiB | [codesystem_16-0.2.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/codesystem_16-0.2.1-1PGSTY.el8.aarch64.rpm) |
| `codesystem_16` | `0.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 122.5 KiB | [codesystem_16-0.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/codesystem_16-0.2.1-1PGSTY.el9.x86_64.rpm) |
| `codesystem_16` | `0.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 121.4 KiB | [codesystem_16-0.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/codesystem_16-0.2.1-1PGSTY.el9.aarch64.rpm) |
| `codesystem_16` | `0.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 122.9 KiB | [codesystem_16-0.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/codesystem_16-0.2.1-1PGSTY.el10.x86_64.rpm) |
| `codesystem_16` | `0.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 122.2 KiB | [codesystem_16-0.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/codesystem_16-0.2.1-1PGSTY.el10.aarch64.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `codesystem_15` | `0.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 122.6 KiB | [codesystem_15-0.2.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/codesystem_15-0.2.1-1PGSTY.el8.x86_64.rpm) |
| `codesystem_15` | `0.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 120.4 KiB | [codesystem_15-0.2.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/codesystem_15-0.2.1-1PGSTY.el8.aarch64.rpm) |
| `codesystem_15` | `0.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 124.1 KiB | [codesystem_15-0.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/codesystem_15-0.2.1-1PGSTY.el9.x86_64.rpm) |
| `codesystem_15` | `0.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 123.2 KiB | [codesystem_15-0.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/codesystem_15-0.2.1-1PGSTY.el9.aarch64.rpm) |
| `codesystem_15` | `0.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 124.5 KiB | [codesystem_15-0.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/codesystem_15-0.2.1-1PGSTY.el10.x86_64.rpm) |
| `codesystem_15` | `0.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 123.8 KiB | [codesystem_15-0.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/codesystem_15-0.2.1-1PGSTY.el10.aarch64.rpm) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `codesystem_14` | `0.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 122.5 KiB | [codesystem_14-0.2.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/codesystem_14-0.2.1-1PGSTY.el8.x86_64.rpm) |
| `codesystem_14` | `0.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 120.4 KiB | [codesystem_14-0.2.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/codesystem_14-0.2.1-1PGSTY.el8.aarch64.rpm) |
| `codesystem_14` | `0.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 124.0 KiB | [codesystem_14-0.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/codesystem_14-0.2.1-1PGSTY.el9.x86_64.rpm) |
| `codesystem_14` | `0.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 123.1 KiB | [codesystem_14-0.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/codesystem_14-0.2.1-1PGSTY.el9.aarch64.rpm) |
| `codesystem_14` | `0.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 124.5 KiB | [codesystem_14-0.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/codesystem_14-0.2.1-1PGSTY.el10.x86_64.rpm) |
| `codesystem_14` | `0.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 123.8 KiB | [codesystem_14-0.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/codesystem_14-0.2.1-1PGSTY.el10.aarch64.rpm) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/John-Cameron-Burns/pg-codesystem" title="Repository" icon="github" subtitle="github.com/John-Cameron-Burns/pg-codesystem" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="codesystem-0.2.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg codesystem;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install codesystem;		# install via package name, for the active PG version

pig install codesystem -v 18;   # install for PG 18
pig install codesystem -v 17;   # install for PG 17
pig install codesystem -v 16;   # install for PG 16
pig install codesystem -v 15;   # install for PG 15
pig install codesystem -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION codesystem;
```

## Usage

Sources:

- [README.md](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/README.md)
- [docs/USER_GUIDE.md](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/docs/USER_GUIDE.md)
- [codesystem--0.2.0.sql](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/codesystem--0.2.0.sql)
- [codesystem--0.1.0--0.2.0.sql](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/codesystem--0.1.0--0.2.0.sql)
- [codesystem.control](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/codesystem.control)
- [.github/workflows/test.yml](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/.github/workflows/test.yml)
- [tools/gss_update.sql](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/tools/gss_update.sql)
- [tools/install_isic_isco.sql](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/tools/install_isic_isco.sql)
- [NOTICE](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/NOTICE)

`codesystem` stores named classification systems and typed `code` values. Version 0.2.0 supports list hierarchies, revisions, derived standards, crosswalks and pattern-defined identifiers. A pattern validates syntax, not whether an identifier has actually been issued.

### Core Workflow

```sql
CREATE EXTENSION codesystem;
CREATE TABLE classification_source (code text, level text, parent text, label text);
INSERT INTO classification_source VALUES ('A', 'section', '', 'Agriculture');
SELECT define_code_system(
  'toy',
  '[{"name":"section","width":1,"charset":"[A-Z]","link":"none","expected":1}]',
  'SELECT * FROM classification_source', 'Toy classification', 'Example');
SELECT name, current_revision, ncodes FROM code_systems;
SELECT code_label('toy:A'::code);
SELECT define_code_system_pattern('my_id', 'AA-NNN');
SELECT code_from('AB-123', 'my_id');
```

### Operational Boundaries

Upstream CI targets PostgreSQL 13–17. Install this native extension as an administrator. Run `ALTER EXTENSION codesystem UPDATE` only when the SQL extension version changes. Distribution 0.2.1 still uses SQL/control version 0.2.0, so an existing 0.2.0 database needs no extension upgrade.

`define_code_system` validates levels and the input query’s code, level, parent and label columns. SQL literals include a system prefix; `code_from` converts external text. `code_level`, `code_parent`, `code_ancestor`, `code_children` and `code_label` expose the hierarchy. Equal text in different revisions represents different values; `code_current` converts to the current revision when the code remains valid.

`define_code_system` also accepts `refines` and `shared_levels` for derived classifications, queried through `code_rollup`. `define_crosswalk` registers correspondence pairs; `code_map` can return multiple matches, while `code_map_unique` requires exactly one. `define_code_system_pattern` defines bounded templates or regular expressions; `code_part` and `code_parts` return named parts. `code_from_padded` handles lost leading zeros, and `code_first_with_prefix` plus `code_last_with_prefix` provide indexable prefix bounds. A system cannot switch between list and pattern forms.

Registry structural edits are protected; define a new revision when accepted values change. Registry writes need appropriate table permissions, and readers need SELECT on registry tables. Custom-format restores require the upstream `tools/restore.sh` helper to restore registry data before typed table values. Ordinary restore ordering can fail.

### Optional Classification Data Tools

Distribution 0.2.1 includes GSS code patterns and entity-type data. After separately loading `tools/gss_update.sql`, `gss_update_entity_types()` updates entity types from ONS. Loading `tools/install_isic_isco.sql` provides `codesystem_install_isic_rev4()` and `codesystem_install_isco_08()`, which download classifications from the UN and ILO. These helpers require an existing `plpython3u`, superuser installation and database-server access to the corresponding sources. They are not dependencies of the core extension.

Use `dry_run` to inspect fetched data first, and connect to the primary when writing the classification registry. ISIC/ISCO data is not redistributed in the package and remains subject to publisher terms. Upstream `NOTICE` records attribution and licensing for other bundled data.
