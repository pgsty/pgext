## Usage

Sources:

- [Official source](https://github.com/Mazuco/pg_synonym/blob/77fa1f92327d2c26f21000293b623941f946dd2f/README.md)
- [Official source](https://github.com/Mazuco/pg_synonym/blob/77fa1f92327d2c26f21000293b623941f946dd2f/pg_synonym.control)
- [Official source](https://github.com/Mazuco/pg_synonym/blob/77fa1f92327d2c26f21000293b623941f946dd2f/pg_synonym--1.0.sql)

`pg_synonym` maintains table and view aliases as invoker-security views. Release 1.0.0 installs extension version 1.0; PostgreSQL 15 or later and superuser installation are required. It needs no preload or server restart.

### Core Workflow

```sql
CREATE EXTENSION pg_synonym;
CREATE SCHEMA app;
CREATE SCHEMA erp;
CREATE TABLE erp.customers (id bigint PRIMARY KEY, name text);
SELECT synonym.create('app.customer', 'erp.customers');
INSERT INTO app.customer VALUES (1, 'Ada');
SELECT * FROM app.customer;
SELECT * FROM synonym.list;
```

All helpers live in the `synonym` schema.

### Operations and Permissions

`synonym.create` can register an alias before its target exists; event triggers resolve it later. Target renames follow the object identity, and supported column changes rebuild the alias. `synonym.refresh()` reconciles aliases manually. `synonym.drop` removes one, with optional `if_exists => true`.

Access through an alias still checks target permissions and row security. Creating an alias requires creation rights in its schema and usage rights on the target schema. Removing one requires ownership, role membership or superuser rights.

### Boundaries

Targets include ordinary, partitioned and foreign tables, views and materialized views. Aliases to aliases, cross-database objects, functions and sequences are unsupported. Use a query for copying data out; the alias is not a substitute for target DDL or truncation. Merging through a view requires PostgreSQL 17 or later.

Manual changes to an alias view can break synchronization. Dependent views may block destructive target changes, and rebuilding an alias can reset grants applied directly to it. Dropping the extension leaves ordinary alias views behind; plan their cleanup separately.
