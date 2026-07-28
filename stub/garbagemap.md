## Usage

Sources:

- [Official upstream README](https://github.com/masahikosawada/walker/blob/4246967426b8c4c86077fc4bef53b638ef3e6124/garbagemap/README.md)
- [Official extension control file (garbagemap.control)](https://github.com/masahikosawada/walker/blob/4246967426b8c4c86077fc4bef53b638ef3e6124/garbagemap/garbagemap.control)
- [Official extension SQL (garbagemap--1.0.sql)](https://github.com/masahikosawada/walker/blob/4246967426b8c4c86077fc4bef53b638ef3e6124/garbagemap/garbagemap--1.0.sql)

`garbagemap` — GarbageMap logs the summary of all garbagemaps whenever CHECKPOINT is executed to PostgreSQL sever log (LOG level). Example is,. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION garbagemap;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gs(rel regclass, OUT rangeno INT, OUT freespace INT, OUT n_tuples INT, OUT n_dead_tuples INT, OUT n_all_visible INT, OUT dead_tuple_ratio NUMERIC(15,4))` is an extension function and returns `SETOF`.
- `gs_rank(rel regclass, OUT rownum INT, OUT percent_blocks NUMERIC(15,4), OUT rangeno INT, OUT freespace INT, OUT n_tuples INT, OUT n_dead_tuples INT, OUT n_all_visible INT, OUT dead_tuple_ratio NUMERIC(15,4), OUT percent NUMERIC(15,4))` is an extension function and returns `SETOF`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
