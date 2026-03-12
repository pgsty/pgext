


## Usage

> [ddlx: DDL eXtractor functions](https://github.com/lacanoid/pgddl)

ddlx is a SQL-only extension that generates DDL scripts from PostgreSQL system catalogs. It provides three main functions that accept an OID and work with all `reg*` object identifier types.

### Core Functions

```sql
-- Generate CREATE statement for an object
SELECT ddlx_create('my_table'::regclass);
SELECT ddlx_create('my_type'::regtype);
SELECT ddlx_create('my_function'::regproc);
SELECT ddlx_create(current_role::regrole);

-- Generate DROP statement
SELECT ddlx_drop('my_table'::regclass);

-- Generate full DDL script with dependency tree
SELECT ddlx_script('my_table'::regclass);
SELECT ddlx_script('my_enum');
SELECT ddlx_script(current_role::regrole);
```

### Options

Options are passed as a text array (e.g., `'{ine,nodcl}'`):

```sql
SELECT ddlx_create('my_table'::regclass, '{ine}');        -- add IF NOT EXISTS
SELECT ddlx_create('my_type'::regtype, '{noowner}');       -- omit ALTER SET OWNER
SELECT ddlx_script('my_table'::regclass, '{drop}');        -- include DROP statements
```

Available options: `drop`, `nodrop`, `owner`, `noowner`, `nogrants`, `nodcl`, `noalter`, `ine` (IF NOT EXISTS), `ie` (IF EXISTS), `ext`, `lite`, `nowrap`, `nopartitions`, `comments`, `nocomments`, `nostorage`, `noconstraints`, `noindexes`, `nosettings`, `notriggers`, `grantor`, `data`.

### For Objects Without reg* Types

```sql
SELECT ddlx_create(oid) FROM pg_foreign_data_wrapper WHERE fdwname = 'postgres_fdw';
SELECT ddlx_create(oid) FROM pg_database WHERE datname = current_database();
```

### Additional Functions

```sql
-- Identify any object by OID
SELECT * FROM ddlx_identify(oid);

-- Describe columns of a class
SELECT * FROM ddlx_describe('my_table'::regclass);

-- Get individual definition parts
SELECT * FROM ddlx_definitions(oid);

-- Generate pre-data creation statements only
SELECT ddlx_createonly('my_table'::regclass);

-- Generate post-data alteration statements
SELECT ddlx_alter('my_table'::regclass);

-- Search function/view bodies by regex
SELECT ddlx_create(objid) FROM ddlx_apropos('users');

-- Get GRANT statements
SELECT ddlx_grants('my_table'::regclass);
```
