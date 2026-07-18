## Usage

Sources:

- [Upstream README and proof-of-concept warning](https://github.com/darold/datalink/blob/fe8853949aedc5ec7df0b0e6a1d8f49d9b57b4bd/README.md)
- [Extension control file](https://github.com/darold/datalink/blob/fe8853949aedc5ec7df0b0e6a1d8f49d9b57b4bd/datalink.control)
- [SQL installation script](https://github.com/darold/datalink/blob/fe8853949aedc5ec7df0b0e6a1d8f49d9b57b4bd/sql/datalink--0.5.0.sql)
- [Background-worker implementation](https://github.com/darold/datalink/blob/fe8853949aedc5ec7df0b0e6a1d8f49d9b57b4bd/datalink_bgw.c)

`datalink` version `0.5.0` is a proof of concept for the SQL/MED `DATALINK` type. A value references a URL or a file below a registered base directory and can carry a comment; database-controlled modes attempt token-based access, copy-on-write behavior, and cleanup/recovery of external files.

### Minimal construction

The module requires `uuid-ossp` and `uri`, must be preloaded for its background worker, and must be installed by a superuser:

```sql
CREATE EXTENSION "uuid-ossp";
CREATE EXTENSION uri;
CREATE EXTENSION datalink;
CREATE TABLE persons (id integer, picture datalink);
INSERT INTO persons VALUES (1, DLVALUE('img1.jpg', 'MyBaseDir', 'portrait'));
```

Register base directories in `pg_catalog.pg_datalink_bases` before using local paths. The PostgreSQL OS account needs the relevant filesystem permissions.

Upstream explicitly says this is not for production. It creates, copies, renames, symlinks, and deletes server-side files, while a background worker removes token files and some external copies. Defaults and semantics are early-stage and the source has no current compatibility matrix. Test only inside an isolated directory tree with backups and an unprivileged OS account.
