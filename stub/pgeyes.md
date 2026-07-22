## Usage

Sources:

- [Official README](https://github.com/ycku/pgeyes/blob/e09e243faaf4837a05366c3fc9d573909dbd053d/README.md)
- [Version 0.0.2 extension SQL](https://github.com/ycku/pgeyes/blob/e09e243faaf4837a05366c3fc9d573909dbd053d/pgeyes--0.0.2.sql)
- [Extension control file](https://github.com/ycku/pgeyes/blob/e09e243faaf4837a05366c3fc9d573909dbd053d/pgeyes.control)

`pgeyes` version `0.0.2` installs one PL/pgSQL environment-check function. Although the project README describes a broader collection of remote administration tools, `CREATE EXTENSION pgeyes` does not install a foreign-data wrapper, remote inspection API, or metadata table in this reviewed release.

### Run the Check

The upstream instructions recommend a dedicated schema with access removed from `PUBLIC`:

```sql
CREATE SCHEMA pgeyes;
REVOKE ALL ON SCHEMA pgeyes FROM PUBLIC;
CREATE EXTENSION pgeyes WITH SCHEMA pgeyes;

SELECT * FROM pgeyes.pgeyes();
```

`pgeyes()` returns one row with four text/boolean fields:

- `item`: `PostgreSQL version number`
- `result`: the value of `server_version_num`
- `valid`: true when the number is at least `90600`
- `description`: the fixed threshold `>=90600`

This check only compares a version number. It does not validate configuration, security, connectivity, replication, backups, or extension compatibility, and the PostgreSQL 9.6 threshold is no longer a useful health gate for current systems.

The project provides no extension update script; its README recommends dropping and recreating the extension. Check dependent objects before doing so. Because the project is abandoned and its installed API is much smaller than its description suggests, use direct PostgreSQL catalog or monitoring queries for new work.
