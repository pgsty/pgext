## Usage

Sources:

- [Official README](https://github.com/adjust/ajbool/blob/b5a47572713f475006fa81e45eda87cc4e965315/README.md)
- [Version 0.0.3 extension SQL](https://github.com/adjust/ajbool/blob/b5a47572713f475006fa81e45eda87cc4e965315/ajbool--0.0.3.sql)
- [C implementation](https://github.com/adjust/ajbool/blob/b5a47572713f475006fa81e45eda87cc4e965315/src/ajbool.c)

`ajbool` adds a one-byte three-state Boolean type whose stored values are `t`, `f`, and `u`. Unlike SQL `NULL`, the explicit `u` value can participate in equality, ordering, hash indexes, B-tree indexes, unique constraints, and primary keys. An `ajbool` column can still separately contain SQL `NULL`.

### Core Workflow

```sql
CREATE EXTENSION ajbool;

CREATE TABLE feature_state (
    feature text PRIMARY KEY,
    state ajbool NOT NULL
);

INSERT INTO feature_state VALUES
    ('search', 't'),
    ('billing', 'f'),
    ('recommendations', 'u');

SELECT feature, state, state::boolean
FROM feature_state
ORDER BY state, feature;
```

Casts preserve ordinary Boolean values and translate between explicit unknown and SQL nullness:

```sql
SELECT true::ajbool, false::ajbool, NULL::boolean::ajbool;
SELECT 't'::ajbool::boolean, 'f'::ajbool::boolean, 'u'::ajbool::boolean;
```

The first query returns `t`, `f`, and `u`; the second returns `true`, `false`, and SQL `NULL`.

### Object Index

- `ajbool` is the scalar type; its text output is `t`, `f`, or `u`.
- Cast `boolean AS ajbool` is assignment-only and maps SQL `NULL` to `u`.
- Cast `ajbool AS boolean` is implicit and maps `u` to SQL `NULL`.
- Operators `=`, `<>`, `<`, `<=`, `>`, and `>=` support comparisons.
- `btree_ajbool_ops` and `hash_ajbool_ops` provide default B-tree and hash operator classes.

### Operational Notes

Version `0.0.3` is relocatable and declares no dependency or preload requirement. Input parsing examines only the first character: `t` becomes true, `f` becomes false, and every other input becomes `u`. Validate application input before casting if silent conversion of misspellings to unknown is unacceptable.

Do not confuse the stored `u` value with SQL `NULL`: `u = u` is a normal comparison, while `NULL = NULL` remains unknown. Casting `u` to `boolean` re-enters SQL three-valued logic, so predicates may need `IS NULL` handling after that cast.
