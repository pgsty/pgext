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

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "-" "PostgreSQL 18: version -" "green" >}} | {{< bg "-" "PostgreSQL 17: version -" "green" >}} | {{< bg "-" "PostgreSQL 16: version -" "green" >}} | {{< bg "-" "PostgreSQL 15: version -" "green" >}} | {{< bg "-" "PostgreSQL 14: version -" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = '$libdir/passwordcheck';
```


This extension does not need `CREATE EXTENSION` DDL command






## Usage

> [passwordcheck: Check password strength on CREATE/ALTER ROLE](https://www.postgresql.org/docs/current/passwordcheck.html)

`passwordcheck` validates password strength whenever passwords are set using `CREATE ROLE` or `ALTER ROLE`. Weak passwords are rejected with an error.

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'passwordcheck'
```

### Configuration Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `passwordcheck.min_password_length` | `8` | Minimum password length in bytes (superuser only) |

### How It Works

The module checks passwords set via `CREATE ROLE` or `ALTER ROLE`:

```sql
-- Rejected if password is too short or too weak
CREATE ROLE myuser WITH LOGIN PASSWORD 'abc';
-- ERROR: password is too short

-- Accepted with a strong enough password
CREATE ROLE myuser WITH LOGIN PASSWORD 'Str0ng_P@ssword!';
```

### Default Checks

Without CrackLib, the module enforces:
- Minimum password length (configurable via `passwordcheck.min_password_length`)
- Password must not be the username
- Basic complexity requirements

### Limitations

- Pre-encrypted passwords sent by client programs cannot be fully validated
- The module can only guess the actual password from encrypted submissions
- For stronger security, consider external authentication methods (e.g., GSSAPI)
- No `CREATE EXTENSION` is required -- this is a shared library module only
