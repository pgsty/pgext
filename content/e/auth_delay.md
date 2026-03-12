---
title: "auth_delay"
linkTitle: "auth_delay"
description: "pause briefly before reporting authentication failure"
weight: 7970
categories: ["SEC"]
width: full
---

[**auth_delay**](https://www.postgresql.org/docs/current/auth-delay.html) : pause briefly before reporting authentication failure


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7970** | {{< badge content="auth_delay" link="https://www.postgresql.org/docs/current/auth-delay.html" >}} | {{< ext "auth_delay" >}} | `-` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_auth_mon" >}} {{< ext "credcheck" >}} {{< ext "login_hook" >}} {{< ext "passwordcheck" >}} {{< ext "passwordcheck_cracklib" >}} {{< ext "pgaudit" >}} {{< ext "set_user" >}} {{< ext "pg_permissions" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "-" "PostgreSQL 18: version -" "green" >}} | {{< bg "-" "PostgreSQL 17: version -" "green" >}} | {{< bg "-" "PostgreSQL 16: version -" "green" >}} | {{< bg "-" "PostgreSQL 15: version -" "green" >}} | {{< bg "-" "PostgreSQL 14: version -" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'auth_delay';
```


This extension does not need `CREATE EXTENSION` DDL command






## Usage

> [auth_delay: Pause before reporting authentication failure](https://www.postgresql.org/docs/current/auth-delay.html)

`auth_delay` pauses the server briefly before reporting authentication failures, making brute-force password attacks more difficult.

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'auth_delay'
auth_delay.milliseconds = '500'
```

### Configuration Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `auth_delay.milliseconds` | `0` | Milliseconds to wait before reporting auth failure |

### Notes

- Must be loaded via `shared_preload_libraries`
- Does not prevent denial-of-service attacks (waiting processes still hold connection slots)
- No `CREATE EXTENSION` is required -- this is a shared library module only
