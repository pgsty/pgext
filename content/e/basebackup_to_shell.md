---
title: "basebackup_to_shell"
linkTitle: "basebackup_to_shell"
description: "adds a custom basebackup target called shell"
weight: 5950
categories: ["Admin"]
width: full
---

adds a custom basebackup target called shell

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5950** | {{< badge content="basebackup_to_shell" link="https://www.postgresql.org/docs/current/basebackup-to-shell.html" >}} | {{< ext "basebackup_to_shell" "basebackup_to_shell" >}} | `-` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s----" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "basic_archive" >}} {{< ext "pg_walinspect" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} {{< ext "pg_dirtyread" >}} {{< ext "pgfincore" >}} {{< ext "pg_cooldown" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="-" color="green" >}} | {{< badge content="-" color="green" >}} | {{< badge content="-" color="green" >}} | {{< badge content="-" color="green" >}} | {{< badge content="N/A" color="red" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION basebackup_to_shell;
```
