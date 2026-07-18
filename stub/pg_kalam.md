## Usage

Sources:

- [PostgreSQL extension README at the reviewed commit](https://github.com/kalamdb/KalamDB/blob/e59521f548e970902d3a71b8253109ebd1fd5731/pg/README.md)
- [Extension control file](https://github.com/kalamdb/KalamDB/blob/e59521f548e970902d3a71b8253109ebd1fd5731/pg/pg_kalam.control)
- [Cargo feature matrix](https://github.com/kalamdb/KalamDB/blob/e59521f548e970902d3a71b8253109ebd1fd5731/pg/Cargo.toml)
- [KalamDB project site](https://kalamdb.org)

`pg_kalam` bridges PostgreSQL to a separately running KalamDB server over gRPC. It registers a `pg_kalam` foreign data wrapper and also supports KalamDB-backed tables created with the `kalamdb` table access method.

```sql
CREATE EXTENSION pg_kalam;
CREATE SERVER kalam_server
  FOREIGN DATA WRAPPER pg_kalam
  OPTIONS (
    host '127.0.0.1', port '2910',
    auth_mode 'account_login',
    login_user 'pg_bridge', login_password 'replace-me'
  );
SELECT kalam_version(), kalam_compiled_mode();
```

The remote KalamDB service must already be running. Current documentation also provides a legacy `static_header` authentication mode; prefer account login and protect foreign-server options because they contain credentials. PostgreSQL `JSON` and `JSONB` map to KalamDB `JSON`, while the remote `FILE` type is mirrored as `JSONB`.

The cataloged version is 0.5.5-rc.1 and the project labels the extension preview. It supports PostgreSQL 13–18 through pgrx feature builds, but remote DDL, type mappings, JSON operators, authentication, network failures, and transactional semantics still require end-to-end validation before production use.
