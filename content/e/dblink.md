---
title: "dblink"
linkTitle: "dblink"
description: "connect to other PostgreSQL databases from within a database"
weight: 8970
categories: ["FDW"]
width: full
---

[**dblink**](https://www.postgresql.org/docs/current/dblink.html) : connect to other PostgreSQL databases from within a database


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8970** | {{< badge content="dblink" link="https://www.postgresql.org/docs/current/dblink.html" >}} | {{< ext "dblink" >}} | `1.2` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "emaj" >}} {{< ext "mimeo" >}} {{< ext "omni_schema" >}} {{< ext "omni_test" >}} {{< ext "omni_vfs" >}} {{< ext "pg_jobmon" >}} {{< ext "pg_profile" >}} |
|   **See Also**    | {{< ext "plproxy" >}} {{< ext "pgbouncer_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "citus" >}} {{< ext "wrappers" >}} {{< ext "pgspider_ext" >}} {{< ext "pglogical" >}} {{< ext "repmgr" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.2" "PostgreSQL 18: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 17: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 16: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 15: version 1.2" "green" >}} | {{< bg "1.2" "PostgreSQL 14: version 1.2" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION dblink;
```




## Usage

> [dblink: Connect to other PostgreSQL databases from within a database](https://www.postgresql.org/docs/current/dblink.html)

### Connect to a Remote Database

```sql
CREATE EXTENSION dblink;

-- Unnamed connection (only one allowed)
SELECT dblink_connect('dbname=remotedb host=remotehost options=-csearch_path=');

-- Named connection (multiple allowed)
SELECT dblink_connect('myconn', 'dbname=remotedb host=remotehost');
```

Or connect via a foreign server definition:

```sql
CREATE SERVER remote_srv FOREIGN DATA WRAPPER dblink_fdw
  OPTIONS (hostaddr '192.168.1.10', dbname 'remotedb');
CREATE USER MAPPING FOR local_user SERVER remote_srv
  OPTIONS (user 'remote_user', password 'secret');

SELECT dblink_connect('myconn', 'remote_srv');
```

### Query a Remote Database

```sql
-- Ad-hoc connection
SELECT * FROM dblink(
  'dbname=remotedb host=remotehost',
  'SELECT id, name, value FROM remote_table'
) AS t(id int, name text, value numeric);

-- Using a named connection
SELECT * FROM dblink(
  'myconn',
  'SELECT id, name FROM remote_table WHERE status = 1'
) AS t(id int, name text);
```

You must always specify the column definition list in the `AS` clause.

### Execute Commands (No Result Set)

```sql
-- INSERT, UPDATE, DELETE, DDL on the remote database
SELECT dblink_exec('myconn', 'INSERT INTO remote_table VALUES (1, ''test'', 42)');
SELECT dblink_exec('myconn', 'UPDATE remote_table SET value = 100 WHERE id = 1');
SELECT dblink_exec('myconn', 'DELETE FROM remote_table WHERE id = 1');
```

Returns the command status string (e.g., `INSERT 0 1`).

### Cursor-Based Access

```sql
SELECT dblink_open('myconn', 'mycursor', 'SELECT * FROM large_table');
SELECT * FROM dblink_fetch('myconn', 'mycursor', 100) AS t(id int, data text);
SELECT dblink_close('myconn', 'mycursor');
```

### Connection Management

```sql
SELECT dblink_get_connections();    -- List open named connections
SELECT dblink_disconnect('myconn'); -- Close a named connection
```

### Create a View for Convenience

```sql
CREATE VIEW remote_data AS
  SELECT * FROM dblink(
    'dbname=remotedb host=remotehost',
    'SELECT id, name, value FROM data_table'
  ) AS t(id int, name text, value numeric);

SELECT * FROM remote_data WHERE value > 100;
```

### Key Functions

| Function | Description |
|----------|-------------|
| `dblink_connect(connstr)` | Open an unnamed persistent connection |
| `dblink_connect(name, connstr)` | Open a named persistent connection |
| `dblink_disconnect(name)` | Close a named connection |
| `dblink(connstr, sql)` | Execute a query, return rows |
| `dblink_exec(connstr, sql)` | Execute a command, return status |
| `dblink_open(name, cursor, sql)` | Open a cursor on a remote database |
| `dblink_fetch(name, cursor, count)` | Fetch rows from a remote cursor |
| `dblink_close(name, cursor)` | Close a remote cursor |
| `dblink_get_connections()` | List all open named connections |
