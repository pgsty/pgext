## Usage

Sources:

- [Upstream README](https://github.com/jaiminpan/cassandra2_fdw/blob/ae5e8145e341f50c150a05d464a9fa4a5eb91179/README.md)
- [Extension control file](https://github.com/jaiminpan/cassandra2_fdw/blob/ae5e8145e341f50c150a05d464a9fa4a5eb91179/cassandra2_fdw.control)
- [SQL installation script](https://github.com/jaiminpan/cassandra2_fdw/blob/ae5e8145e341f50c150a05d464a9fa4a5eb91179/cassandra2_fdw--2.0.sql)
- [C FDW implementation](https://github.com/jaiminpan/cassandra2_fdw/blob/ae5e8145e341f50c150a05d464a9fa4a5eb91179/cassandra2_fdw.c)

`cassandra2_fdw` version `2.0` is a read-only foreign data wrapper for Cassandra 2.x–4.x using the DataStax C/C++ driver. A server specifies Cassandra hosts, port, and protocol; a user mapping supplies credentials; a foreign table maps a PostgreSQL row shape to a Cassandra keyspace and table.

### Example

```sql
CREATE EXTENSION cassandra2_fdw;
CREATE SERVER cass_serv FOREIGN DATA WRAPPER cassandra2_fdw
  OPTIONS (host '127.0.0.1,127.0.0.2', port '9042', protocol '4');
CREATE USER MAPPING FOR CURRENT_USER SERVER cass_serv
  OPTIONS (username 'reader', password 'secret');
CREATE FOREIGN TABLE cass_orders (id integer, description text)
  SERVER cass_serv OPTIONS (schema_name 'example', table_name 'orders');
SELECT * FROM cass_orders LIMIT 5;
```

The wrapper does not write to Cassandra and supports only a limited scalar type mapping; collections, maps, and user-defined types are not implemented. Upstream reports PostgreSQL 11–15 and Cassandra 3/4 testing, not current PostgreSQL releases. Credentials are stored in PostgreSQL catalogs. Use a least-privilege Cassandra account, protect catalog visibility and network traffic, and validate protocol/driver compatibility and query behavior before deployment.
