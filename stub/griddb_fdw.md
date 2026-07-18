## Usage

Sources:

- [Upstream README](https://github.com/pgspider/griddb_fdw/blob/a05fbf78a57d1941538aa11a0038037b08ff0e79/README.md)
- [Extension control file](https://github.com/pgspider/griddb_fdw/blob/a05fbf78a57d1941538aa11a0038037b08ff0e79/griddb_fdw.control)
- [Extension SQL](https://github.com/pgspider/griddb_fdw/blob/a05fbf78a57d1941538aa11a0038037b08ff0e79/griddb_fdw--1.2.sql)
- [Project release v2.4.0](https://github.com/pgspider/griddb_fdw/releases/tag/v2.4.0)

`griddb_fdw` connects PostgreSQL to GridDB containers. Build and runtime installation require the external GridDB C client, including `gridstore.h` and `libgridstore.so`.

Create a server and user mapping, then import the remote schema:

```sql
CREATE EXTENSION griddb_fdw;

CREATE SERVER griddb_remote
  FOREIGN DATA WRAPPER griddb_fdw
  OPTIONS (host '239.0.0.1', port '31999', clustername 'exampleCluster', database 'public');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER griddb_remote
  OPTIONS (username 'appuser', password 'change-me');

IMPORT FOREIGN SCHEMA public
  FROM SERVER griddb_remote
  INTO public;
```

### Versions and limitations

Project release `v2.4.0` contains extension SQL version `1.2`; these are different version namespaces. That release documents PostgreSQL 13 through 17 and testing with GridDB 5.5.0. It supports reads and writes, but remote containers need a row key for `UPDATE` or `DELETE`. It does not support `ON CONFLICT`, `RETURNING`, `TRUNCATE`, or `SAVEPOINT`, and PostgreSQL column definitions must match GridDB, so upstream recommends `IMPORT FOREIGN SCHEMA` over hand-written foreign tables.
