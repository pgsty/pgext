## Usage

Sources:

- [Upstream README and option reference](https://github.com/pgspider/dynamodb_fdw/blob/1d79f84b455705a1fb778a0c3ae35483ebd0d40c/README.md)
- [Extension control file](https://github.com/pgspider/dynamodb_fdw/blob/1d79f84b455705a1fb778a0c3ae35483ebd0d40c/dynamodb_fdw.control)
- [Authoritative option validator](https://github.com/pgspider/dynamodb_fdw/blob/1d79f84b455705a1fb778a0c3ae35483ebd0d40c/option.c)
- [Distribution metadata](https://github.com/pgspider/dynamodb_fdw/blob/1d79f84b455705a1fb778a0c3ae35483ebd0d40c/META.json)

`dynamodb_fdw` maps Amazon DynamoDB tables into PostgreSQL foreign tables. It supports `SELECT`, `INSERT`, `UPDATE`, and `DELETE`, pushes supported `WHERE` expressions and the DynamoDB `SIZE` operation, and can address nested map/list values with `->` and `->>`.

### Define a foreign table

```sql
CREATE EXTENSION dynamodb_fdw;

CREATE SERVER dynamodb_local
  FOREIGN DATA WRAPPER dynamodb_fdw
  OPTIONS (endpoint 'http://localhost:8000');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER dynamodb_local
  OPTIONS (user 'local-access-key', password 'local-secret-key');

CREATE FOREIGN TABLE public.orders (
  order_id text,
  created_at numeric,
  payload jsonb
)
SERVER dynamodb_local
OPTIONS (table_name 'orders', partition_key 'order_id', sort_key 'created_at');

SELECT order_id, payload->>'status'
FROM public.orders
WHERE order_id = 'A-100';
```

The server option is `endpoint`; user-mapping options are `user` and `password`; table options are `partition_key`, `sort_key`, and `table_name`; the column option is `column_name`. One README example uses `username`, but the validator accepts `user`, as shown above. Protect user mappings because they contain credentials.

The FDW does not support `IMPORT FOREIGN SCHEMA`, foreign-table `TRUNCATE`, or `COPY FROM`. DynamoDB sets are exposed as arrays but their order can change; null set elements are replaced with type defaults, and a missing map attribute does not satisfy the documented `IS NULL` pushdown. The README targets PostgreSQL 13 through 17 and requires the AWS C++ SDK. The catalog/control version is `1.1`, while the same commit's distribution metadata says `2.1.1`; treat these as separate version channels rather than assuming interchangeability.
