---
title: "lo"
linkTitle: "lo"
description: "Large Object maintenance"
weight: 5930
categories: ["ADMIN"]
width: full
---

[**lo**](https://www.postgresql.org/docs/current/lo.html) : Large Object maintenance


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5930** | {{< badge content="lo" link="https://www.postgresql.org/docs/current/lo.html" >}} | {{< ext "lo" >}} | `1.1` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgcrypto" >}} {{< ext "adminpack" >}} {{< ext "file_fdw" >}} {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.1" "PostgreSQL 18: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 17: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 16: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 15: version 1.1" "green" >}} | {{< bg "1.1" "PostgreSQL 14: version 1.1" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION lo;
```




## Usage

> [lo: Large Object maintenance](https://www.postgresql.org/docs/current/lo.html)

The `lo` extension provides a data type and trigger function for managing PostgreSQL Large Objects, preventing orphaned objects when references are updated or deleted.

### Data Type

The `lo` type is a domain over `oid`, used to identify columns that hold Large Object references. This is especially useful for ODBC driver compatibility.

```sql
CREATE TABLE image (
    title  text,
    raster lo      -- large object reference column
);
```

### Trigger Function

The `lo_manage()` trigger automatically calls `lo_unlink()` to delete the associated Large Object when a row is updated or deleted:

```sql
CREATE TRIGGER t_raster
    BEFORE UPDATE OR DELETE ON image
    FOR EACH ROW EXECUTE FUNCTION lo_manage(raster);
```

For multiple `lo` columns, create a separate trigger for each:

```sql
CREATE TABLE gallery (
    title     text,
    thumbnail lo,
    fullsize  lo
);

CREATE TRIGGER t_thumbnail
    BEFORE UPDATE OR DELETE ON gallery
    FOR EACH ROW EXECUTE FUNCTION lo_manage(thumbnail);

CREATE TRIGGER t_fullsize
    BEFORE UPDATE OR DELETE ON gallery
    FOR EACH ROW EXECUTE FUNCTION lo_manage(fullsize);
```

To restrict the trigger to column updates only:

```sql
CREATE TRIGGER t_raster
    BEFORE UPDATE OF raster OR DELETE ON image
    FOR EACH ROW EXECUTE FUNCTION lo_manage(raster);
```

### Limitations

- `DROP TABLE` and `TRUNCATE` do not fire row-level triggers, so Large Objects will be orphaned. Run `DELETE FROM table` before dropping.
- The trigger assumes each Large Object is referenced by only one column/row.
- Use the `vacuumlo` utility to clean up any orphaned Large Objects.

The extension is trusted and can be installed by non-superusers with `CREATE` privilege.
