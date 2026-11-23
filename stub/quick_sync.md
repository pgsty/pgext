## Usage

Sources:

- [README.md](https://github.com/angelmorgaldev/pg-quick-sync/blob/45c160cec4f72673ec1a34fb1a93cb832d17d18d/README.md)
- [quick_sync.control](https://github.com/angelmorgaldev/pg-quick-sync/blob/45c160cec4f72673ec1a34fb1a93cb832d17d18d/quick_sync.control)
- [sql/quick_sync--0.1.0.sql](https://github.com/angelmorgaldev/pg-quick-sync/blob/45c160cec4f72673ec1a34fb1a93cb832d17d18d/sql/quick_sync--0.1.0.sql)
- [LICENSE](https://github.com/angelmorgaldev/pg-quick-sync/blob/45c160cec4f72673ec1a34fb1a93cb832d17d18d/LICENSE)

`quick_sync` 0.1.0 manages postgres_fdw connection objects from PostgreSQL URIs. The project is named pg-quick-sync, but the SQL extension name is quick_sync. The implemented module covers connection creation, listing, editing and deletion; table mapping and incremental synchronization remain under development. Upstream requires PostgreSQL 13 or later.

### Manage a Connection

Install its dependency and then the extension with sufficient administrative privileges:

```sql
CREATE EXTENSION postgres_fdw;
CREATE EXTENSION quick_sync;
SELECT quick_sync.create_connection(
    'source', 'postgresql://reader:example@db.example.net:5432/app'
);
SELECT * FROM quick_sync.list_connections();
SELECT quick_sync.edit_connection('source', p_fetch_size => 10000);
```

The URI above is illustrative; bind real connection details as a client parameter and protect SQL logs and history. `quick_sync.create_connection()` creates a foreign server named from the normalized alias, creates a mapping for the current database user and records non-password metadata in `quick_sync.connections`. The default fetch size is 50000 and remote estimates are enabled. Creation configures catalog objects; it does not prove that the remote server is reachable.

### API and Privileges

`quick_sync.list_connections()` reports registered metadata, foreign-table counts and whether the foreign server still exists. Its active flag is a catalog-existence check, not a network-health check. `quick_sync.edit_connection()` accepts either a new URI or named host, port, database, user, password and fetch-size arguments.

The functions use the caller's privileges. The caller needs appropriate foreign-server/user-mapping rights and access to the extension's metadata table. Passwords are stored in PostgreSQL user mappings, not in the custom inventory table. The simple URI parser is not a complete libpq URI implementation; do not assume IPv6 literals, percent-decoding or arbitrary connection options work.

### Delete Carefully

```sql
SELECT quick_sync.delete_connection('source', false);
```

The second argument defaults to true, which drops the server with CASCADE and can remove dependent foreign tables. Passing false uses RESTRICT so dependencies block the deletion. The extension uses the fixed `quick_sync` schema, is not relocatable and needs no shared preload or restart. Its SQL module installs no background synchronization worker.
