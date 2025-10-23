
--------

## Usage


### Create Extension

After install the `duckdb_fdw` yum package, you can create the extension inside PostgreSQL database:

```sql
-- create extension
CREATE EXTENSION duckdb_fdw;

-- create duckdb_fdw server
CREATE SERVER duckdb_server FOREIGN DATA WRAPPER duckdb_fdw OPTIONS (database '/tmp/duck.db');

-- create user mapping [OPTIONAL]
-- GRANT USAGE ON FOREIGN SERVER duckdb_server TO PUBLIC;

SELECT duckdb_fdw_version();

-- You can execute duckdb command with `duckdb_execute`, for example, to create a table inside duckdb:
-- create a table in duckdb
SELECT duckdb_execute('duckdb_server', 'CREATE TABLE t1 (a integer,b varchar);');

-- create duckdb foreign table mapping that duckdb table
CREATE FOREIGN TABLE t1 (
    a integer,
    b text
) SERVER duckdb_server OPTIONS (
    table 't1'
);

-- write some data and read it back
INSERT INTO t1 SELECT i AS a,i::text AS b FROM generate_series(1,10) i;
SELECT * FROM t1;
```


You can also import foreign schema from duckdb server, for example, create a table with duckdb cli:

```bash
duckdb /tmp/duck.db

CREATE TABLE t1 (
  a integer,
  b text
);
  
INSERT INTO t1 VALUES (1, 'a'), (2 , 'b'), (3, 'c');
SELECT * FROM t1;
```

Then import the schema into PostgreSQL:

```sql
IMPORT FOREIGN SCHEMA public FROM SERVER duckdb_server INTO public;
```

### Other Resource

- [DuckDB Website](https://duckdb.org/)
- [GitHub: duckdb_fdw](https://github.com/alitrack/duckdb_fdw/)
- [Building libduckdb](https://github.com/digoal/blog/blob/master/202401/20240124_01.md)

