---
title: "amcheck"
linkTitle: "amcheck"
description: "functions for verifying relation integrity"
weight: 5980
categories: ["ADMIN"]
width: full
---

functions for verifying relation integrity


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5980** | {{< badge content="amcheck" link="https://www.postgresql.org/docs/current/amcheck.html" >}} | {{< ext "amcheck" >}} | `1.4` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_checksums" >}} {{< ext "pg_catcheck" >}} {{< ext "pg_visibility" >}} {{< ext "pg_surgery" >}} {{< ext "toastinfo" >}} {{< ext "pagevis" >}} {{< ext "pageinspect" >}} {{< ext "pg_freespacemap" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.4" "PostgreSQL 18: version 1.4" "green" >}} | {{< bg "1.4" "PostgreSQL 17: version 1.4" "green" >}} | {{< bg "1.4" "PostgreSQL 16: version 1.4" "green" >}} | {{< bg "1.4" "PostgreSQL 15: version 1.4" "green" >}} | {{< bg "1.4" "PostgreSQL 14: version 1.4" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION amcheck;
```
