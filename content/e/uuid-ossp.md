---
title: "uuid-ossp"
linkTitle: "uuid-ossp"
description: "generate universally unique identifiers (UUIDs)"
weight: 4930
categories: ["FUNC"]
width: full
---

[**uuid-ossp**](https://www.postgresql.org/docs/current/uuid-ossp.html)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4930** | {{< badge content="uuid-ossp" link="https://www.postgresql.org/docs/current/uuid-ossp.html" >}} | {{< ext "uuid-ossp" >}} | `1.1` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "babelfishpg_tsql" >}} |
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "permuteseq" >}} {{< ext "ddsketch" >}} {{< ext "vasco" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.1" "PostgreSQL 18: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 17: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 16: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 15: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 14: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 13: version 1.1" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION uuid-ossp;
```
