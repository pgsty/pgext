---
title: "citext"
linkTitle: "citext"
description: "data type for case-insensitive character strings"
weight: 3980
categories: ["TYPE"]
width: full
---

[**citext**](https://www.postgresql.org/docs/current/citext.html) : data type for case-insensitive character strings


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3980** | {{< badge content="citext" link="https://www.postgresql.org/docs/current/citext.html" >}} | {{< ext "citext" >}} | `1.6` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "ltree" >}} {{< ext "unaccent" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.6" "PostgreSQL 18: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 17: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 16: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 15: version 1.6" "green" >}} | {{< bg "1.6" "PostgreSQL 14: version 1.6" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION citext;
```



## Usage

> [citext: case-insensitive character string type](https://www.postgresql.org/docs/current/citext.html)

The `citext` extension provides a case-insensitive text type that eliminates the need for `lower()` calls in queries.

```sql
CREATE EXTENSION citext;
```

### Basic Usage

```sql
CREATE TABLE users (
    nick citext PRIMARY KEY,
    pass text NOT NULL
);

INSERT INTO users VALUES ('Larry', 'secret123');

-- Case-insensitive matching
SELECT * FROM users WHERE nick = 'larry';   -- matches 'Larry'
SELECT * FROM users WHERE nick = 'LARRY';   -- matches 'Larry'
```

### Behavior

`citext` performs comparisons by internally converting strings to lowercase. The following operations are case-insensitive with `citext`:

- Comparison operators: `=`, `<>`, `<`, `>`, `<=`, `>=`
- Pattern matching: `LIKE`, `ILIKE`, `~~`, `~~*`
- Regular expressions: `~`, `~*`, `!~`, `!~*`

### Case-Insensitive Functions

When arguments are `citext`, these functions perform case-insensitive matching:

`regexp_match()`, `regexp_matches()`, `regexp_replace()`, `regexp_split_to_array()`, `regexp_split_to_table()`, `replace()`, `split_part()`, `strpos()`, `translate()`

### Advantages Over lower()

- Eliminates verbose `lower()` calls in WHERE clauses
- Supports case-insensitive PRIMARY KEY and UNIQUE constraints
- No need for functional indexes
- Transparent case-folding in all operations

### Limitations

- Case-folding depends on `LC_CTYPE` at database creation
- Slightly less efficient than `text` (copying and conversion overhead)
- Does not support B-tree deduplication
- For better Unicode handling, consider nondeterministic collations instead
