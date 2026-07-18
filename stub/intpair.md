## Usage

Sources:

- [pg_intpair README at the reviewed commit](https://github.com/citusdata/pg_intpair/blob/fa274d8aa6a73670b9cee95e94373022c2777926/README.md)
- [intpair 0.2 install SQL at the reviewed commit](https://github.com/citusdata/pg_intpair/blob/fa274d8aa6a73670b9cee95e94373022c2777926/intpair--0.2.sql)
- [intpair regression SQL at the reviewed commit](https://github.com/citusdata/pg_intpair/blob/fa274d8aa6a73670b9cee95e94373022c2777926/sql/intpair.sql)

`intpair` version 0.2 provides the fixed-length `int64pair` type: two signed 64-bit integers stored in 16 bytes. Construct values with `int64pair(a, b)` and read their zero-based components with subscripts. Its comparison, B-tree, and hash support allow ordinary keys and indexes.

```sql
CREATE EXTENSION intpair;

CREATE TABLE pair_demo (
  id int64pair PRIMARY KEY
);

INSERT INTO pair_demo VALUES
  (int64pair(-1, 1)),
  (int64pair(10, 20));

SELECT id, id[0] AS first_value, id[1] AS second_value
FROM pair_demo
ORDER BY id;
```

### Caveats

- Subscripts are zero-based: `p[0]` is the first integer and `p[1]` is the second.
- This is a compact scalar type, not a PostgreSQL composite. Migrating a matching composite column requires an explicit text round trip or compatible casts, as shown by upstream.
- Index ordering follows the extension's comparison implementation. Validate ordering-dependent application behavior before replacing a composite or array representation.
