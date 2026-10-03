## Usage

Sources:

- [Official README](https://bitbucket.org/adunstan/blackhole_fdw/src/b97c35672aa8618fd84dae6bb11714fdf23acef2/README.md)
- [Extension control file](https://bitbucket.org/adunstan/blackhole_fdw/src/b97c35672aa8618fd84dae6bb11714fdf23acef2/blackhole_fdw.control)
- [Official SQL example](https://bitbucket.org/adunstan/blackhole_fdw/src/b97c35672aa8618fd84dae6bb11714fdf23acef2/sql/blackhole_fdw.sql)

`blackhole_fdw` is Andrew Dunstan's standalone FDW teaching skeleton and an intentional data sink. It discards inserted rows, treats updates and deletes as no-ops, and returns an empty result set. Never use it for data that must be retained.

### Core Workflow

After installing the library and extension SQL, an administrator can create a sink table. A successful insert does not mean the row was stored.

```sql
CREATE EXTENSION blackhole_fdw;
CREATE SERVER sink FOREIGN DATA WRAPPER blackhole_fdw;
CREATE FOREIGN TABLE discarded (id integer, payload text) SERVER sink;
INSERT INTO discarded VALUES (1, 'intentionally discarded');
SELECT * FROM discarded;
```

### Development Boundary

The author presents the implementation as a starting point for FDW development, including scan and modification callbacks. The control version is `0.0.1`, with relocatable SQL objects and no documented preload requirement. The source does not provide a current tested PostgreSQL-major matrix. The canonical upstream is Bitbucket; the GitHub import is a mirror. Treat this as demonstration infrastructure with deliberate data loss, not an archival or replication mechanism.
