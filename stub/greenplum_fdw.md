## Usage

Sources:

- [Upstream README](https://github.com/luxms/greenplum-fdw/blob/fc8f03883391b72940ff8ceb9dc0aa56644a3c67/README.md)
- [Extension SQL](https://github.com/luxms/greenplum-fdw/blob/f643f9193fa43c672b2f8fec1fb9e1c90900d1be/greenplum_fdw--1.0.sql)
- [PostgreSQL postgres_fdw documentation](https://www.postgresql.org/docs/current/postgres-fdw.html)

`greenplum_fdw` is a fork of PostgreSQL's `postgres_fdw` for connections to Greenplum. It changes the remote transaction isolation level from `REPEATABLE READ` to `SERIALIZABLE`, addressing the documented Greenplum 5 rejection of the former level.

After installing the server-side library, configure it with the normal `postgres_fdw` object model:

```sql
CREATE EXTENSION greenplum_fdw;

CREATE SERVER gp_remote
  FOREIGN DATA WRAPPER greenplum_fdw
  OPTIONS (host 'greenplum.example.net', port '5432', dbname 'analytics');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER gp_remote
  OPTIONS (user 'gp_reader', password 'change-me');

CREATE FOREIGN TABLE gp_sales (
  sale_id bigint,
  amount numeric
) SERVER gp_remote OPTIONS (schema_name 'public', table_name 'sales');
```

### Compatibility

Upstream version `1.0` says the published branches target PostgreSQL 12 and 13, and that the remaining `postgres_fdw` documentation applies. The repository does not publish a current Greenplum/PostgreSQL compatibility matrix, so build and transaction tests are required before using it with newer releases. Protect mapping credentials and grant foreign-server access only to intended roles.
