---
title: "basic_archive"
linkTitle: "basic_archive"
description: "an example of an archive module"
weight: 5940
categories: ["ADMIN"]
width: full
---

[**basic_archive**](https://www.postgresql.org/docs/current/basic-archive.html) : an example of an archive module


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5940** | {{< badge content="basic_archive" link="https://www.postgresql.org/docs/current/basic-archive.html" >}} | {{< ext "basic_archive" >}} | `-` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s----" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "basebackup_to_shell" >}} {{< ext "pg_walinspect" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} {{< ext "pg_dirtyread" >}} {{< ext "pgfincore" >}} {{< ext "pg_cooldown" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "-" "PostgreSQL 18: version -" "green" >}} | {{< bg "-" "PostgreSQL 17: version -" "green" >}} | {{< bg "-" "PostgreSQL 16: version -" "green" >}} | {{< bg "-" "PostgreSQL 15: version -" "green" >}} | {{< bg "N/A" "PostgreSQL 14: not available" "red" >}} | {{< bg "N/A" "PostgreSQL 13: not available" "red" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


This extension does not need `CREATE EXTENSION` DDL command


