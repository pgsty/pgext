---
title: "sepgsql"
linkTitle: "sepgsql"
description: "label-based mandatory access control (MAC) based on SELinux security policy."
weight: 7960
categories: ["SEC"]
width: full
---

[**sepgsql**](https://www.postgresql.org/docs/current/sepgsql.html) : label-based mandatory access control (MAC) based on SELinux security policy.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7960** | {{< badge content="sepgsql" link="https://www.postgresql.org/docs/current/sepgsql.html" >}} | {{< ext "sepgsql" >}} | `-` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_readonly" >}} {{< ext "pg_permissions" >}} {{< ext "set_user" >}} {{< ext "noset" >}} {{< ext "pgaudit" >}} {{< ext "credcheck" >}} {{< ext "login_hook" >}} {{< ext "passwordcheck_cracklib" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "-" "PostgreSQL 18: version -" "green" >}} | {{< bg "-" "PostgreSQL 17: version -" "green" >}} | {{< bg "-" "PostgreSQL 16: version -" "green" >}} | {{< bg "-" "PostgreSQL 15: version -" "green" >}} | {{< bg "-" "PostgreSQL 14: version -" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'sepgsql';
```


This extension does not need `CREATE EXTENSION` DDL command






## Usage

> [sepgsql: SELinux label-based mandatory access control for PostgreSQL](https://www.postgresql.org/docs/current/sepgsql.html)

`sepgsql` provides label-based mandatory access control (MAC) based on SELinux security policy. It adds an extra layer of security checking above PostgreSQL's standard SQL permissions.

### Configuration Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `sepgsql.permissive` | `off` | Enable permissive mode regardless of system SELinux settings |
| `sepgsql.debug_audit` | `off` | Force all possible logging regardless of policy |

### Functions

| Function | Returns | Description |
|----------|---------|-------------|
| `sepgsql_getcon()` | `text` | Get current client security label |
| `sepgsql_setcon(text)` | `boolean` | Switch client domain to new label (NULL to revert) |
| `sepgsql_mcstrans_in(text)` | `text` | Translate qualified MLS/MCS range to raw format |
| `sepgsql_mcstrans_out(text)` | `text` | Translate raw MLS/MCS range to qualified format |
| `sepgsql_restorecon(text)` | `boolean` | Set initial security labels for all objects in database |

### Security Labels

Security labels can be assigned to schemas, tables, columns, sequences, views, and functions:

```sql
SECURITY LABEL ON COLUMN customer.credit
    IS 'system_u:object_r:sepgsql_secret_table_t:s0';
```

### Dynamic Domain Transitions

```sql
SELECT sepgsql_getcon();
-- unconfined_u:unconfined_r:unconfined_t:s0-s0:c0.c1023

SELECT sepgsql_setcon('unconfined_u:unconfined_r:unconfined_t:s0-s0:c1.c4');
-- t
```

### Trusted Procedures

```sql
-- Create function to access sensitive data with masking
CREATE FUNCTION show_credit(int) RETURNS text
    AS 'SELECT regexp_replace(credit, ''-[0-9]+$'', ''-xxxx'', ''g'')
         FROM customer WHERE cid = $1'
    LANGUAGE sql;

-- Mark as trusted procedure
SECURITY LABEL ON FUNCTION show_credit(int)
    IS 'system_u:object_r:sepgsql_trusted_proc_exec_t:s0';
```

### Permission Classes

DML operations check: `db_table:{select|insert|update|delete}` and `db_column:{select|update|insert}`.
DDL operations check: `create`, `drop`, `setattr`, `add_name`, `remove_name`.
Schema access requires: `db_schema:search`.
