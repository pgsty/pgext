---
title: "sepgsql"
linkTitle: "sepgsql"
description: "label-based mandatory access control (MAC) based on SELinux security policy."
weight: 7960
categories: ["Sec"]
width: full
---

label-based mandatory access control (MAC) based on SELinux security policy.

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7960** | {{< badge content="sepgsql" link="https://www.postgresql.org/docs/current/sepgsql.html" >}} | {{< ext "sepgsql" "sepgsql" >}} | `-` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sL---" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_readonly" >}} {{< ext "pg_permissions" >}} {{< ext "set_user" >}} {{< ext "noset" >}} {{< ext "pgaudit" >}} {{< ext "credcheck" >}} {{< ext "login_hook" >}} {{< ext "passwordcheck_cracklib" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< badge content="-" color="green" >}} | {{< badge content="-" color="green" >}} | {{< badge content="-" color="green" >}} | {{< badge content="-" color="green" >}} | {{< badge content="-" color="green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install

Add this extension to [`shared_preload_libraries`](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'sepgsql';  -- comma-separated list
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION sepgsql;
```
