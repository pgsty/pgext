## Usage

Sources:

- [Official README at the reviewed commit](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/README.md)
- [Version 0.0.1 SQL implementation](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/norm--0.0.1.sql)
- [Extension control file](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/NORM.control)

`norm` is a pure SQL/PLpgSQL code generator for simple CRUD functions. Its generators inspect catalog metadata and create table-specific `SECURITY DEFINER` functions. Use it only in a trusted schema-migration workflow where every identifier and every generated definition is reviewed.

### Core Workflow

```sql
CREATE EXTENSION norm;

CREATE TABLE games (
  game_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  game_name text NOT NULL
);

SELECT norm_insert('games', ARRAY['game_name']);
SELECT norm_get('games', ARRAY['game_id', 'game_name']);

SELECT insert_games('chess');
SELECT * FROM get_games();
```

The public generators are `norm_insert(table, columns, function_name)`, `norm_update(table, columns, filters, function_name)`, `norm_delete(table, filters, function_name)`, and `norm_get(tables, columns, filters, function_name)`. The final name argument is optional; defaults are derived from the table names.

Generated update functions use `coalesce(parameter, column)`, so they cannot set a target column to SQL NULL. Multi-table reads infer inner joins from foreign keys, reject tables with multiple foreign keys in the supported path, and do not provide outer joins, aggregates, or arbitrary subqueries.

### Security Boundary

The reviewed SQL interpolates caller-supplied table, column, and custom function names without consistently using identifier quoting. Generated `SECURITY DEFINER` functions also do not establish a hardened `search_path`. Untrusted generator input can therefore become SQL injection, object substitution, or privilege escalation.

Do not grant generator execution to application roles. Run it only as a controlled deployment step, avoid ambiguous same-named tables across schemas, inspect the resulting function definitions and owners, set a safe execution path, and revoke default privileges before exposing any generated helper.
