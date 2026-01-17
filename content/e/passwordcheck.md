---
title: "passwordcheck"
linkTitle: "passwordcheck"
description: "checks user passwords and reject weak password"
weight: 7990
categories: ["SEC"]
width: full
---

[**passwordcheck**](https://www.postgresql.org/docs/current/passwordcheck.html) : checks user passwords and reject weak password


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7990** | {{< badge content="passwordcheck" link="https://www.postgresql.org/docs/current/passwordcheck.html" >}} | {{< ext "passwordcheck" >}} | `-` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_auth_mon" >}} {{< ext "credcheck" >}} {{< ext "pgaudit" >}} {{< ext "login_hook" >}} {{< ext "auth_delay" >}} {{< ext "set_user" >}} {{< ext "sepgsql" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** | **PG13** |
|:--------:|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "-" "PostgreSQL 18: version -" "green" >}} | {{< bg "-" "PostgreSQL 17: version -" "green" >}} | {{< bg "-" "PostgreSQL 16: version -" "green" >}} | {{< bg "-" "PostgreSQL 15: version -" "green" >}} | {{< bg "-" "PostgreSQL 14: version -" "green" >}} | {{< bg "-" "PostgreSQL 13: version -" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = '$libdir/passwordcheck';
```


This extension does not need `CREATE EXTENSION` DDL command


