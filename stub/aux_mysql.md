## Usage

Sources:

- [19.0.0 manual](https://github.com/pgsty/openhalo/blob/v19.0.0/README.md)
- [2.0 control file](https://github.com/pgsty/openhalo/blob/v19.0.0/contrib/aux_mysql/aux_mysql.control)

`aux_mysql` 2.0 ships with OpenHalo 19.0.0 and provides MySQL wire protocol, syntax, and functions. Run `pig install openhalo-19` to install the `openhalodb-19` kernel package. This command selects PG19 even when another PostgreSQL major is active.

### Enable and connect

New OpenHalo 19 clusters load `aux_mysql` through `shared_preload_libraries` by default. The extension creates ICU collations, so use a database with an ICU-compatible encoding such as `UTF8`. As an administrator, connect over PostgreSQL to the database named by `mysql.database` (default `postgres`) and create the extension.

```sql
CREATE EXTENSION aux_mysql;
SELECT extversion FROM pg_extension WHERE extname = 'aux_mysql';
```

On OpenHalo 19 the query returns `2.0`. With `mysql.listener_on` enabled, MySQL clients can connect to `mysql.port` (default 3306). MySQL databases map to schemas in the selected PostgreSQL database. The extension installs casts in that database. Use a separate database for workloads that require stock PostgreSQL casts.
