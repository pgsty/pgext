## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/README.md)
- [Version 0.0.1 SQL implementation](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/norm--0.0.1.sql)
- [PGXS Makefile](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/Makefile)
- [Upstream control file](https://github.com/Chessbrain/NORM/blob/5d44da1bc220d6c09dd6ac4132891568e1b70d85/NORM.control)

`norm` is a pure SQL/PLpgSQL code generator for simple CRUD functions. `norm_insert`, `norm_update`, `norm_delete`, and `norm_get` inspect table metadata and create table-specific functions.

```sql
CREATE EXTENSION norm;

CREATE TABLE games (
  game_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  game_name text NOT NULL
);

SELECT norm_insert('games', ARRAY['game_name']);
SELECT insert_games('chess');
```

Generated functions use `SECURITY DEFINER`, and the implementation builds dynamic SQL from supplied table and column names. Accept identifiers only from trusted callers, review ownership and `search_path`, and inspect generated functions before granting access.

### Early-Version Limits

Version 0.0.1 does not support multi-table writes, left joins, aggregates, or joins involving tables with multiple foreign keys. The PGXS Makefile and versioned SQL use lowercase `norm`, but the repository stores `NORM.control`; a case-sensitive source build may need that file installed as `norm.control`.
