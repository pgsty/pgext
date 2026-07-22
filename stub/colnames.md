## Usage

Sources:

- [Current upstream documentation at the reviewed commit](https://github.com/theory/colnames/blob/00bdd0f4becdde261c0395ebfbc7943b49e5d519/doc/colnames.md)
- [Version 1.7.0 SQL objects](https://github.com/theory/colnames/blob/00bdd0f4becdde261c0395ebfbc7943b49e5d519/sql/colnames.sql)
- [C implementation](https://github.com/theory/colnames/blob/00bdd0f4becdde261c0395ebfbc7943b49e5d519/src/colnames.c)

`colnames` provides one function, `colnames(record)`, which returns the input record's field names as a `text[]`. It is useful when generic trigger or metaprogramming code needs the shape of a composite value.

```sql
CREATE EXTENSION colnames;
SELECT colnames(ROW(1, 'foo', 458.0));
SELECT colnames(t)
FROM (SELECT 1 AS id, 'Ada'::text AS name) AS t;
```

Anonymous rows yield generated names such as `f1`, while named subqueries or table rows preserve their attribute names. Dropped columns, duplicate aliases, domains, and evolving composite types should be tested if output drives generated SQL.

PostgreSQL 9.3 and later can obtain names without a C extension through `row_to_json` and `json_object_keys`. Prefer the pure-SQL alternative when extension deployment is undesirable. If names are interpolated into dynamic SQL, quote them with PostgreSQL identifier functions and keep values parameterized; field discovery does not make SQL construction safe.
