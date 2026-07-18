## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/pjungwir/inetrange/blob/9d62352afda4b84be01d2fb3375805a04a64995f/README.md)
- [Extension control file](https://github.com/pjungwir/inetrange/blob/9d62352afda4b84be01d2fb3375805a04a64995f/inetrange.control)
- [Version 1.0.2 SQL objects](https://github.com/pjungwir/inetrange/blob/9d62352afda4b84be01d2fb3375805a04a64995f/inetrange--1.0.2.sql)
- [Current PGXN distribution](https://pgxn.org/dist/inetrange/)

`inetrange` adds a PostgreSQL range type whose subtype is `inet`. It supports ordinary range containment and overlap operators, GiST indexing, and exclusion constraints for non-overlapping IP address intervals.

```sql
CREATE EXTENSION inetrange;
SELECT '1.0.1.1'::inet <@ inetrange('1.0.0.0', '1.1.0.0', '[]');
CREATE TABLE network_block (
  addresses inetrange NOT NULL,
  owner_name text NOT NULL,
  EXCLUDE USING gist (addresses WITH &&)
);
```

Bounds follow PostgreSQL range rules, while the subtype follows `inet` ordering across IPv4 and IPv6 values. Choose inclusive and exclusive endpoints deliberately and test canonical application representations before using equality as a business key.

Version 1.0.2 has current source and PGXN distribution but no explicit modern PostgreSQL compatibility matrix. Test GiST index builds, exclusion constraints, mixed address families, dump/restore, and upgrades on the target major before storing long-lived values.
