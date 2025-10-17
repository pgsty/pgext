---
title: "lo"
linkTitle: "lo"
description: "Large Object maintenance"
weight: 5930
categories: ["Admin"]
width: full
---

Large Object maintenance

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5930** | {{< badge content="lo" link="https://www.postgresql.org/docs/current/lo.html" >}} | {{< ext "lo" "lo" >}} | `1.1` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgcrypto" >}} {{< ext "adminpack" >}} {{< ext "file_fdw" >}} {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.1" color="green" >}} | {{< badge content="1.1" color="green" >}} | {{< badge content="1.1" color="green" >}} | {{< badge content="1.1" color="green" >}} | {{< badge content="1.1" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION lo;
```
