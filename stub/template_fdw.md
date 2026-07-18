## Usage

Sources:

- [Pinned upstream README](https://github.com/okbob/template_fdw/blob/621f30e37f6ad530c1ed47d8764383368a7a3831/README.md)
- [Pinned extension control file](https://github.com/okbob/template_fdw/blob/621f30e37f6ad530c1ed47d8764383368a7a3831/template_fdw.control)
- [Official PGXN distribution page](https://pgxn.org/dist/template_fdw/)

`template_fdw` creates schema-only foreign tables that intentionally reject both reads and writes. They act as persistent templates for temporary tables, allowing tools such as `plpgsql_check` to resolve table shapes during static analysis without permitting accidental use of the persistent template relation.

```sql
CREATE EXTENSION template_fdw;
CREATE SERVER template_server FOREIGN DATA WRAPPER template_fdw;

CREATE FOREIGN TABLE public.template_rows (a integer, b integer)
  SERVER template_server;

-- Both statements against the foreign table are expected to fail.
SELECT * FROM public.template_rows;
INSERT INTO public.template_rows VALUES (10, 20);

-- A temporary table can safely copy and shadow the template definition.
CREATE TEMP TABLE template_rows
  (LIKE public.template_rows INCLUDING ALL);
INSERT INTO template_rows VALUES (10, 20);
SELECT * FROM template_rows;
```

The FDW is a development and validation guard, not remote storage. Any successful `SELECT` or DML against the foreign table contradicts its intended behavior and should fail a regression test. Control version `0.0` differs from upstream/PGXN release `0.0.2`; identify builds by source commit and test error behavior after every PostgreSQL upgrade.

Keep templates in a dedicated schema, grant only the privileges needed by validation roles, and qualify persistent names explicitly so search-path changes do not target a real table unexpectedly.
