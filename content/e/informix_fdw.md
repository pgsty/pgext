---
title: "informix_fdw"
linkTitle: "informix_fdw"
description: "Foreign data wrapper for Informix access"
weight: 8670
categories: ["FDW"]
width: full
---

[**informix_fdw**](https://github.com/credativ/informix_fdw) : Foreign data wrapper for Informix access


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8670** | {{< badge content="informix_fdw" link="https://github.com/credativ/informix_fdw" >}} | {{< ext "informix_fdw" >}} | `0.6.3` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] PGDG non-free (pgnf) only; no SQL-level extension dependency; runtime requires IBM Informix Client SDK (libifsql15a/libifasf15a/libifgen15a/libifos15a/libifgls)


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.6.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `informix_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.6.3` | {{< bg "18" "informix_fdw_18" "green" >}} {{< bg "17" "informix_fdw_17" "green" >}} {{< bg "16" "informix_fdw_16" "green" >}} {{< bg "15" "informix_fdw_15" "green" >}} {{< bg "14" "informix_fdw_14" "green" >}} {{< bg "13" "informix_fdw_13" "green" >}} | `informix_fdw_$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "informix_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "informix_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.6.3" "informix_fdw_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "informix_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "informix_fdw : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `informix_fdw_18` | `0.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 61.3 KiB | [informix_fdw_18-0.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-8-x86_64/informix_fdw_18-0.6.3-1PGDG.rhel8.x86_64.rpm) |
| `informix_fdw_18` | `0.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 59.8 KiB | [informix_fdw_18-0.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-9-x86_64/informix_fdw_18-0.6.3-1PGDG.rhel9.x86_64.rpm) |
| `informix_fdw_18` | `0.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 60.6 KiB | [informix_fdw_18-0.6.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/18/redhat/rhel-10-x86_64/informix_fdw_18-0.6.3-1PGDG.rhel10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `informix_fdw_17` | `0.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 61.3 KiB | [informix_fdw_17-0.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-8-x86_64/informix_fdw_17-0.6.3-1PGDG.rhel8.x86_64.rpm) |
| `informix_fdw_17` | `0.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 59.9 KiB | [informix_fdw_17-0.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-9-x86_64/informix_fdw_17-0.6.3-1PGDG.rhel9.x86_64.rpm) |
| `informix_fdw_17` | `0.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 60.5 KiB | [informix_fdw_17-0.6.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-10-x86_64/informix_fdw_17-0.6.3-1PGDG.rhel10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `informix_fdw_16` | `0.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 61.3 KiB | [informix_fdw_16-0.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-8-x86_64/informix_fdw_16-0.6.3-1PGDG.rhel8.x86_64.rpm) |
| `informix_fdw_16` | `0.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 59.8 KiB | [informix_fdw_16-0.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/informix_fdw_16-0.6.3-1PGDG.rhel9.x86_64.rpm) |
| `informix_fdw_16` | `0.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 59.5 KiB | [informix_fdw_16-0.6.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-9-x86_64/informix_fdw_16-0.6.2-2PGDG.rhel9.x86_64.rpm) |
| `informix_fdw_16` | `0.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 60.6 KiB | [informix_fdw_16-0.6.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/informix_fdw_16-0.6.3-1PGDG.rhel10.x86_64.rpm) |
| `informix_fdw_16` | `0.6.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 60.2 KiB | [informix_fdw_16-0.6.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-10-x86_64/informix_fdw_16-0.6.2-2PGDG.rhel10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `informix_fdw_15` | `0.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.4 KiB | [informix_fdw_15-0.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-8-x86_64/informix_fdw_15-0.6.3-1PGDG.rhel8.x86_64.rpm) |
| `informix_fdw_15` | `0.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.2 KiB | [informix_fdw_15-0.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/informix_fdw_15-0.6.3-1PGDG.rhel9.x86_64.rpm) |
| `informix_fdw_15` | `0.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.9 KiB | [informix_fdw_15-0.6.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-9-x86_64/informix_fdw_15-0.6.2-2PGDG.rhel9.x86_64.rpm) |
| `informix_fdw_15` | `0.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.9 KiB | [informix_fdw_15-0.6.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/informix_fdw_15-0.6.3-1PGDG.rhel10.x86_64.rpm) |
| `informix_fdw_15` | `0.6.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.4 KiB | [informix_fdw_15-0.6.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-10-x86_64/informix_fdw_15-0.6.2-2PGDG.rhel10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `informix_fdw_14` | `0.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.4 KiB | [informix_fdw_14-0.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-8-x86_64/informix_fdw_14-0.6.3-1PGDG.rhel8.x86_64.rpm) |
| `informix_fdw_14` | `0.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.2 KiB | [informix_fdw_14-0.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/informix_fdw_14-0.6.3-1PGDG.rhel9.x86_64.rpm) |
| `informix_fdw_14` | `0.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.8 KiB | [informix_fdw_14-0.6.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-9-x86_64/informix_fdw_14-0.6.2-2PGDG.rhel9.x86_64.rpm) |
| `informix_fdw_14` | `0.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.9 KiB | [informix_fdw_14-0.6.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/informix_fdw_14-0.6.3-1PGDG.rhel10.x86_64.rpm) |
| `informix_fdw_14` | `0.6.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.5 KiB | [informix_fdw_14-0.6.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-10-x86_64/informix_fdw_14-0.6.2-2PGDG.rhel10.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `informix_fdw_13` | `0.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.9 KiB | [informix_fdw_13-0.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-8-x86_64/informix_fdw_13-0.6.3-1PGDG.rhel8.x86_64.rpm) |
| `informix_fdw_13` | `0.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.3 KiB | [informix_fdw_13-0.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/informix_fdw_13-0.6.3-1PGDG.rhel9.x86_64.rpm) |
| `informix_fdw_13` | `0.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.8 KiB | [informix_fdw_13-0.6.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-9-x86_64/informix_fdw_13-0.6.2-2PGDG.rhel9.x86_64.rpm) |
| `informix_fdw_13` | `0.6.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.8 KiB | [informix_fdw_13-0.6.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-10-x86_64/informix_fdw_13-0.6.3-1PGDG.rhel10.x86_64.rpm) |
| `informix_fdw_13` | `0.6.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.2 KiB | [informix_fdw_13-0.6.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-10-x86_64/informix_fdw_13-0.6.2-2PGDG.rhel10.x86_64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/credativ/informix_fdw" title="Repository" icon="github" subtitle="github.com/credativ/informix_fdw" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install informix_fdw;		# install via package name, for the active PG version

pig install informix_fdw -v 18;   # install for PG 18
pig install informix_fdw -v 17;   # install for PG 17
pig install informix_fdw -v 16;   # install for PG 16
pig install informix_fdw -v 15;   # install for PG 15
pig install informix_fdw -v 14;   # install for PG 14
pig install informix_fdw -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION informix_fdw;
```
