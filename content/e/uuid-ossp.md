---
title: "uuid-ossp"
linkTitle: "uuid-ossp"
description: "generate universally unique identifiers (UUIDs)"
weight: 4930
categories: ["Func"]
width: full
---

generate universally unique identifiers (UUIDs)

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4930** | {{< badge content="uuid-ossp" link="https://www.postgresql.org/docs/current/uuid-ossp.html" >}} | {{< ext "uuid-ossp" "uuid-ossp" >}} | `1.1` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "babelfishpg_tsql" >}} |
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "permuteseq" >}} {{< ext "ddsketch" >}} {{< ext "vasco" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="1.1" color="green" >}} | {{< badge content="1.1" color="green" >}} | {{< badge content="1.1" color="green" >}} | {{< badge content="1.1" color="green" >}} | {{< badge content="1.1" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION uuid-ossp;
```
