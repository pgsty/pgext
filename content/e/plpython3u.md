---
title: "plpython3u"
linkTitle: "plpython3u"
description: "PL/Python3U untrusted procedural language"
weight: 3290
categories: ["LANG"]
width: full
---

[**plpython3u**](https://www.postgresql.org/docs/current/plpython.html) : PL/Python3U untrusted procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3290** | {{< badge content="plpython3u" link="https://www.postgresql.org/docs/current/plpython.html" >}} | {{< ext "plpython3u" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|    **Need By**    | {{< ext "hstore_plpython3u" >}} {{< ext "jsonb_plpython3u" >}} {{< ext "ltree_plpython3u" >}} {{< ext "omni_python" >}} {{< ext "pg4ml" >}} |
|   **See Also**    | {{< ext "faker" >}} {{< ext "plv8" >}} {{< ext "pllua" >}} {{< ext "plluau" >}} {{< ext "pltcl" >}} {{< ext "pltclu" >}} {{< ext "plperl" >}} {{< ext "plperlu" >}} |
|    **Siblings**   | {{< ext "jsonb_plpython3u" >}} {{< ext "ltree_plpython3u" >}} {{< ext "hstore_plpython3u" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plpython3u;
```



## Usage

> [plpython3u: PL/Python3 untrusted procedural language](https://www.postgresql.org/docs/current/plpython.html)

PL/Python3U allows writing PostgreSQL functions in Python 3. It is an untrusted language with full access to the Python ecosystem. Only superusers can create functions.

```sql
CREATE EXTENSION plpython3u;

-- Simple function
CREATE FUNCTION py_hello(name text) RETURNS text
LANGUAGE plpython3u AS $$
  return f"Hello, {name}!"
$$;

SELECT py_hello('world');

-- Using Python standard library
CREATE FUNCTION py_sha256(data text) RETURNS text
LANGUAGE plpython3u AS $$
  import hashlib
  return hashlib.sha256(data.encode()).hexdigest()
$$;

-- Returning a composite type
CREATE TYPE address AS (street text, city text, zip text);

CREATE FUNCTION parse_address(raw text) RETURNS address
LANGUAGE plpython3u AS $$
  import re
  m = re.match(r'(.+),\s*(.+)\s+(\d{5})', raw)
  if m:
    return (m.group(1), m.group(2), m.group(3))
  return None
$$;

-- Set-returning function
CREATE FUNCTION py_generate_dates(start text, days int) RETURNS SETOF date
LANGUAGE plpython3u AS $$
  from datetime import datetime, timedelta
  d = datetime.strptime(start, '%Y-%m-%d')
  for i in range(days):
    yield (d + timedelta(days=i)).strftime('%Y-%m-%d')
$$;

-- Database access via plpy
CREATE FUNCTION py_row_count(table_name text) RETURNS bigint
LANGUAGE plpython3u AS $$
  result = plpy.execute(f"SELECT count(*) AS cnt FROM {table_name}")
  return result[0]['cnt']
$$;

-- Using external packages (must be installed on the server)
CREATE FUNCTION py_parse_json(url text) RETURNS jsonb
LANGUAGE plpython3u AS $$
  import json, urllib.request
  response = urllib.request.urlopen(url)
  data = json.loads(response.read())
  return json.dumps(data)
$$;
```
