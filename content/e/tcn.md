---
title: "tcn"
linkTitle: "tcn"
description: "Triggered change notifications"
weight: 4920
categories: ["FUNC"]
width: full
---

[**tcn**](https://www.postgresql.org/docs/current/tcn.html) : Triggered change notifications


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4920** | {{< badge content="tcn" link="https://www.postgresql.org/docs/current/tcn.html" >}} | {{< ext "tcn" >}} | `1.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "pg_uuidv7" >}} {{< ext "permuteseq" >}} {{< ext "pg_hashids" >}} {{< ext "sequential_uuids" >}} {{< ext "topn" >}} {{< ext "quantile" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION tcn;
```



## Usage

> [tcn: triggered change notifications via LISTEN/NOTIFY](https://www.postgresql.org/docs/current/tcn.html)

Provides a trigger function that sends NOTIFY events with information about changed rows, enabling asynchronous change tracking.

```sql
CREATE EXTENSION tcn;
```

### Trigger Function

| Function | Description |
|---|---|
| `triggered_change_notification()` | Send NOTIFY on row changes with primary key info |

Optional parameter: custom channel name (defaults to `tcn`).

### Notification Payload Format

```
"table_name",operation,"column"='value',"column"='value'
```

Operations: `I` (INSERT), `U` (UPDATE), `D` (DELETE).

### Examples

```sql
CREATE TABLE tcndata (
  a int NOT NULL,
  b date NOT NULL,
  c text,
  PRIMARY KEY (a, b)
);

-- Attach the trigger
CREATE TRIGGER tcndata_tcn
  AFTER INSERT OR UPDATE OR DELETE ON tcndata
  FOR EACH ROW
  EXECUTE FUNCTION triggered_change_notification();

-- Listen for notifications
LISTEN tcn;

-- Changes trigger notifications:
INSERT INTO tcndata VALUES (1, '2024-01-01', 'test');
-- Notification: "tcndata",I,"a"='1',"b"='2024-01-01'

UPDATE tcndata SET c = 'updated' WHERE a = 1;
-- Notification: "tcndata",U,"a"='1',"b"='2024-01-01'

DELETE FROM tcndata WHERE a = 1;
-- Notification: "tcndata",D,"a"='1',"b"='2024-01-01'

-- Use a custom channel name
CREATE TRIGGER my_trigger
  AFTER INSERT OR UPDATE OR DELETE ON my_table
  FOR EACH ROW
  EXECUTE FUNCTION triggered_change_notification('my_channel');
```
