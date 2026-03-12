---
title: "uuid-ossp"
linkTitle: "uuid-ossp"
description: "generate universally unique identifiers (UUIDs)"
weight: 4930
categories: ["FUNC"]
width: full
---

[**uuid-ossp**](https://www.postgresql.org/docs/current/uuid-ossp.html) : generate universally unique identifiers (UUIDs)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4930** | {{< badge content="uuid-ossp" link="https://www.postgresql.org/docs/current/uuid-ossp.html" >}} | {{< ext "uuid-ossp" >}} | `1.1` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "babelfishpg_tsql" >}} |
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "permuteseq" >}} {{< ext "ddsketch" >}} {{< ext "vasco" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.1" "PostgreSQL 18: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 17: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 16: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 15: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 14: version 1.1" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION uuid-ossp;
```



## Usage

> [uuid-ossp: UUID generation functions](https://www.postgresql.org/docs/current/uuid-ossp.html)

Provides functions to generate UUIDs using several standard algorithms. Note: for simple random UUIDs, consider using the built-in `gen_random_uuid()` instead.

```sql
CREATE EXTENSION "uuid-ossp";
```

### UUID Generation Functions

| Function | Description |
|---|---|
| `uuid_generate_v1()` | Version 1: MAC address + timestamp |
| `uuid_generate_v1mc()` | Version 1 with random multicast MAC |
| `uuid_generate_v3(namespace uuid, name text)` | Version 3: MD5 hash of namespace + name |
| `uuid_generate_v4()` | Version 4: fully random |
| `uuid_generate_v5(namespace uuid, name text)` | Version 5: SHA-1 hash of namespace + name (preferred over v3) |

### Namespace Constants

| Function | Description |
|---|---|
| `uuid_nil()` | Nil UUID (all zeros) |
| `uuid_ns_dns()` | DNS namespace |
| `uuid_ns_url()` | URL namespace |
| `uuid_ns_oid()` | ISO OID namespace |
| `uuid_ns_x500()` | X.500 DN namespace |

### Examples

```sql
-- Random UUID (v4)
SELECT uuid_generate_v4();

-- Timestamp-based UUID (v1)
SELECT uuid_generate_v1();

-- Deterministic UUID from name (v5, preferred over v3)
SELECT uuid_generate_v5(uuid_ns_url(), 'http://www.postgresql.org');

-- Use as default primary key
CREATE TABLE items (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  name text
);
```
