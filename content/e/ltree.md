---
title: "ltree"
linkTitle: "ltree"
description: "data type for hierarchical tree-like structures"
weight: 3960
categories: ["TYPE"]
width: full
---

[**ltree**](https://www.postgresql.org/docs/current/ltree.html) : data type for hierarchical tree-like structures


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3960** | {{< badge content="ltree" link="https://www.postgresql.org/docs/current/ltree.html" >}} | {{< ext "ltree" >}} | `1.3` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "ltree_plpython3u" >}} |
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "citext" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.3" "PostgreSQL 18: version 1.3" "green" >}} | {{< bg "1.3" "PostgreSQL 17: version 1.3" "green" >}} | {{< bg "1.3" "PostgreSQL 16: version 1.3" "green" >}} | {{< bg "1.3" "PostgreSQL 15: version 1.3" "green" >}} | {{< bg "1.3" "PostgreSQL 14: version 1.3" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION ltree;
```



## Usage

> [ltree: hierarchical tree-like label data type](https://www.postgresql.org/docs/current/ltree.html)

The `ltree` extension provides data types for hierarchical tree-structured data with extensive search facilities.

```sql
CREATE EXTENSION ltree;
```

### Data Types

- **`ltree`**: Label path (e.g., `Top.Science.Astronomy`)
- **`lquery`**: Regex-like pattern for matching ltree values
- **`ltxtquery`**: Full-text-search-like pattern

### Basic Usage

```sql
CREATE TABLE categories (path ltree);
INSERT INTO categories VALUES
    ('Top'), ('Top.Science'), ('Top.Science.Astronomy'),
    ('Top.Hobbies'), ('Top.Collections.Pictures');

-- Find descendants
SELECT path FROM categories WHERE path <@ 'Top.Science';

-- Pattern matching
SELECT path FROM categories WHERE path ~ '*.Astronomy.*';

-- Full-text search
SELECT path FROM categories WHERE path @ 'Science & !Pictures';
```

### Operators

| Operator | Description |
|----------|-------------|
| `@>` | Is ancestor of (or equal) |
| `<@` | Is descendant of (or equal) |
| `~` | Matches lquery pattern |
| `?` | Matches any lquery in array |
| `@` | Matches ltxtquery |
| `\|\|` | Concatenate paths |

### lquery Patterns

```sql
'*.Science.*'           -- any path containing Science
'Top.*{1,2}.Astronomy'  -- 1-2 labels between Top and Astronomy
'*.astro*'              -- prefix matching
'*.Astro*@'             -- case-insensitive prefix
```

### Functions

```sql
SELECT nlevel('Top.Science.Astronomy');                     -- 3
SELECT subltree('Top.Science.Astronomy', 1, 2);            -- Science
SELECT subpath('Top.Science.Astronomy', 1);                 -- Science.Astronomy
SELECT index('a.b.c.d', 'b.c');                             -- 1
SELECT lca('1.2.3', '1.2.3.4.5');                          -- 1.2
SELECT lca(ARRAY['1.2.3'::ltree, '1.2.4'::ltree]);        -- 1.2
```

### Index Support

```sql
-- GiST index (supports @>, <@, ~, ?, @)
CREATE INDEX path_gist_idx ON categories USING gist (path);

-- B-tree index (supports <, <=, =, >=, >)
CREATE INDEX path_idx ON categories USING btree (path);
```
