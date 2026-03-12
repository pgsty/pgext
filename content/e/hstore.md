---
title: "hstore"
linkTitle: "hstore"
description: "data type for storing sets of (key, value) pairs"
weight: 3970
categories: ["TYPE"]
width: full
---

[**hstore**](https://www.postgresql.org/docs/current/hstore.html) : data type for storing sets of (key, value) pairs


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3970** | {{< badge content="hstore" link="https://www.postgresql.org/docs/current/hstore.html" >}} | {{< ext "hstore" >}} | `1.8` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "hstore_pllua" >}} {{< ext "hstore_plluau" >}} {{< ext "hstore_plpython3u" >}} {{< ext "pg_readme" >}} {{< ext "pg_readme_test_extension" >}} |
|   **See Also**    | {{< ext "intarray" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.8" "PostgreSQL 18: version 1.8" "green" >}} | {{< bg "1.8" "PostgreSQL 17: version 1.8" "green" >}} | {{< bg "1.8" "PostgreSQL 16: version 1.8" "green" >}} | {{< bg "1.8" "PostgreSQL 15: version 1.8" "green" >}} | {{< bg "1.8" "PostgreSQL 14: version 1.8" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION hstore;
```



## Usage

> [hstore: key-value pair data type](https://www.postgresql.org/docs/current/hstore.html)

The `hstore` extension provides a key/value pair data type for storing semi-structured data in a single column.

```sql
CREATE EXTENSION hstore;
```

### Basic Usage

```sql
SELECT 'name => Alice, age => 30'::hstore;
SELECT 'name => Alice'::hstore -> 'name';           -- Alice
SELECT 'a => 1, b => 2'::hstore ? 'a';              -- true
SELECT 'a => 1'::hstore || 'b => 2'::hstore;        -- "a"=>"1", "b"=>"2"
```

### Operators

| Operator | Description | Example |
|----------|-------------|---------|
| `->` | Get value by key | `h -> 'key'` |
| `\|\|` | Concatenate | `h1 \|\| h2` |
| `?` | Contains key | `h ? 'key'` |
| `?&` | Contains all keys | `h ?& ARRAY['a','b']` |
| `?\|` | Contains any key | `h ?\| ARRAY['a','b']` |
| `@>` | Contains | `h @> 'a=>1'` |
| `<@` | Contained by | `h <@ 'a=>1, b=>2'` |
| `-` | Delete key(s) | `h - 'key'` or `h - ARRAY['a','b']` |

### Subscript Access

```sql
SELECT h['name'] FROM mytable;
UPDATE mytable SET h['age'] = '31';
```

### Functions

```sql
-- Construction
SELECT hstore('key', 'value');
SELECT hstore(ARRAY['a','b'], ARRAY['1','2']);
SELECT hstore(ROW(1, 'hello'));

-- Extraction
SELECT akeys(h);                    -- text[] of keys
SELECT avals(h);                    -- text[] of values
SELECT skeys(h);                    -- set of keys
SELECT svals(h);                    -- set of values
SELECT each(h);                     -- set of (key, value) records

-- Query
SELECT exist(h, 'key');             -- boolean
SELECT defined(h, 'key');           -- true if non-NULL value

-- Modification
SELECT delete(h, 'key');
SELECT slice(h, ARRAY['a','b']);    -- subset of keys

-- JSON conversion
SELECT hstore_to_json(h);
SELECT hstore_to_jsonb(h);
SELECT hstore_to_json_loose(h);    -- distinguishes numbers/booleans

-- Record conversion
SELECT populate_record(NULL::my_table, h);
```

### Index Support

```sql
CREATE INDEX idx ON t USING gin (h);    -- supports @>, ?, ?&, ?|
CREATE INDEX idx ON t USING gist (h);   -- supports @>, ?, ?&, ?|
CREATE INDEX idx ON t USING btree (h);  -- supports =
CREATE INDEX idx ON t USING hash (h);   -- supports =
```
