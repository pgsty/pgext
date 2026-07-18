## Usage

Sources:

- [Current upstream README at HEAD](http://sigaev.ru/git/gitweb.cgi?p=gevel.git;a=blob;f=README.gevel;hb=cb9416a0fae4b90eaaefeaa85948c50b481dc2a0)
- [Current extension control file](http://sigaev.ru/git/gitweb.cgi?p=gevel.git;a=blob;f=gevel.control;hb=cb9416a0fae4b90eaaefeaa85948c50b481dc2a0)
- [Current installation SQL](http://sigaev.ru/git/gitweb.cgi?p=gevel.git;a=blob;f=gevel.sql;hb=cb9416a0fae4b90eaaefeaa85948c50b481dc2a0)
- [Current index-page implementation](http://sigaev.ru/git/gitweb.cgi?p=gevel.git;a=blob;f=gevel.c;hb=cb9416a0fae4b90eaaefeaa85948c50b481dc2a0)
- [Official project documentation](http://www.sai.msu.su/~megera/oddmuse/index.cgi/Gevel)

gevel version 1.1 is a low-level index diagnostics module for advanced developers. It walks GiST, GIN, and SP-GiST pages to report tree structure, page statistics, stored keys, and estimated GIN matches. The current source also contains B-tree and BRIN diagnostics, but their build and installation are conditional and inconsistent across PostgreSQL versions.

### Diagnostic example

If the installed package supplies a working versioned extension script, inspect a test index as follows:

```sql
CREATE EXTENSION gevel;

SELECT gist_stat('public.demo_gist_idx');
SELECT gist_tree('public.demo_gist_idx', 1);

SELECT gin_count_estimate(
    'public.demo_gin_idx',
    'postgres'::tsquery
);
```

Printing functions return setof record, so callers must provide the exact output column types required by the index operator class. The upstream README gives concrete examples for GiST, GIN, SP-GiST, B-tree, and BRIN.

### Locking, privilege, and packaging risks

This code reads raw index pages through private PostgreSQL access-method structures. Several GiST, B-tree, and BRIN paths take AccessExclusiveLock and recursively scan an entire index; large or corrupt indexes can block workloads for a long time or crash a backend. Printing stored index keys may reveal indexed data. The SQL grants normal public execution by default, while the C paths do not implement a table-style SELECT privilege boundary. Revoke all functions from PUBLIC and expose only narrowly reviewed wrappers to a diagnostics role.

The current HEAD is from 2020. Its Makefile installs legacy gevel.sql and does not declare the normal EXTENSION/versioned-SQL PGXS variables even though gevel.control declares version 1.1. It conditionally installs B-tree and BRIN SQL only for PostgreSQL 12. Distribution packages may patch this, so verify the exact installed files and function inventory instead of assuming upstream CREATE EXTENSION works unchanged.

Use only on a replica or disposable copy, compile against the exact server source, set statement and lock timeouts, test one small index first, and never use page output as a supported stable API.
