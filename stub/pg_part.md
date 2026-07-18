## Usage

Sources:

- [Upstream function documentation](https://github.com/uptimejp/pg_part/blob/6a5c96db670f75fa486006536e659d5c5f4b876d/README.md)
- [Version 0.1.0 SQL implementation](https://github.com/uptimejp/pg_part/blob/6a5c96db670f75fa486006536e659d5c5f4b876d/pg_part--0.1.0.sql)
- [PGXN distribution page](https://pgxn.org/dist/pg_part/)

`pg_part` provides PL/pgSQL helpers in the fixed `pgpart` schema for legacy inheritance-based partitioning. It can add, merge, attach, detach, and list child tables.

```sql
CREATE EXTENSION pg_part;
SELECT pgpart.show_partition('public', 'orders');
SELECT pgpart.attach_partition('public', 'orders', 'orders_2025',
                               'order_date >= ''2025-01-01''');
```

The add and merge paths build dynamic DDL and use server-side `COPY` files to move rows. Identifiers, check expressions, and file paths are concatenated into SQL, so restrict execution to trusted administrators and never pass user-controlled text. This predates declarative partitioning; rehearse locking, constraints, indexes, privileges, failure rollback, disk capacity, and cleanup on a copy before touching important data, and prefer core declarative partitioning for new systems.
