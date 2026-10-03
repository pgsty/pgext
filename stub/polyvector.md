## Usage

Sources:

- [polyvector.control](https://github.com/mynelis/polyvector/blob/1393fc58eee67894d2bd5850eef2dd12717422da/polyvector.control)
- [README.md](https://github.com/mynelis/polyvector/blob/1393fc58eee67894d2bd5850eef2dd12717422da/README.md)
- [Makefile](https://github.com/mynelis/polyvector/blob/1393fc58eee67894d2bd5850eef2dd12717422da/Makefile)
- [sql/polyvector--1.0.sql](https://github.com/mynelis/polyvector/blob/1393fc58eee67894d2bd5850eef2dd12717422da/sql/polyvector--1.0.sql)

`polyvector` is a SQL-only wrapper around pgvector that stores one active vector slot in a composite value. The installed 1.0 script deliberately uses dimensions 3, 4, 5 and 6, despite slot names suggesting larger dimensions.

### Core Workflow

```sql
CREATE EXTENSION vector;
CREATE EXTENSION polyvector;
SELECT polyvec_get_embedding(polyvec_convert('[1,2,3]'::vector));
CREATE TABLE polyvector_sample (
  id bigint PRIMARY KEY,
  embedding polyvec CHECK (polyvec_validate(embedding))
);
```

### Operational Boundaries

Enable `vector` first and install with superuser privileges; PL/pgSQL is used internally. No preload is involved. `polyvec_validate` checks the single-slot invariant, while conversion/getter/setter functions route by dimension. L2, cosine and inner-product search helpers operate on selected slots. Add the validation check to application tables. The separate production example is not the SQL selected by the Makefile; do not assume 384/768/1024/1536-dimensional support in this installed version.
