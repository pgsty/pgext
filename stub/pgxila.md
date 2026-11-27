## Usage

Sources:

- [README.md](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/README.md)
- [extension/Cargo.toml](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/extension/Cargo.toml)
- [extension/pgxila.control](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/extension/pgxila.control)
- [extension/src/lib.rs](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/extension/src/lib.rs)
- [extension/src/tls_guard.rs](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/extension/src/tls_guard.rs)
- [extension/src/workers.rs](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/extension/src/workers.rs)
- [extension/sql/schema.sql](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/extension/sql/schema.sql)
- [docs/api.md](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/docs/api.md)
- [docker/initdb/10-pgxila.sql](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/docker/initdb/10-pgxila.sql)
- [docker/postgresql.conf](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/docker/postgresql.conf)
- [LICENSE](https://github.com/fredericalix/pgXila/blob/16f2d70299e893b610fdef503f904c28a8c75d84/LICENSE)

`pgxila` 0.1.0 is an early PostgreSQL 19 extension for storing and serving card-based applications. Screens, stacks, widgets, palette data and history live in the `xila` schema; separate terminal and graphical clients render the SQL API. The reviewed source has no published release or SQL upgrade path.

### Enable and Connect

Configure valid server certificates and add the library to the existing preload list before restarting PostgreSQL:

```ini
shared_preload_libraries = 'pgxila'
ssl = on
ssl_min_protocol_version = 'TLSv1.3'
```

This untrusted, non-relocatable extension requires a superuser installer. A normal build refuses to load outside server startup. Its authentication hook applies across the cluster: TCP connections must use TLS 1.3, including connections to other databases. Local Unix-socket connections remain available for administration. The guard has no runtime setting to disable it.

Connect over the administrative Unix socket, or a verified TLS connection, and create the extension in the application database:

```sql
CREATE EXTENSION pgxila;
SELECT xila.version();
SELECT xila.hello();
SELECT xila.open_stack('desktop');
SELECT * FROM xila.palette();
```

The installation creates `xila_user` and `xila_author` group roles when absent. Give existing application roles membership in `xila_user`; authors additionally need `xila_author`. The upstream image also installs `vector`, but the reviewed extension control and current SQL do not require it. The SQL layer uses `plpgsql`.

### Application Workflow

The client starts with `xila.hello()` and checks the API version, then uses `xila.open_stack()` for the screen definition. `xila.card()` reads a card; `xila.save_card()` writes fields; `xila.act()` dispatches an action. Cards remain rows in their backing tables. Ordinary table privileges and row-level security apply, and stack definitions are editable by their owner and members of that owner role.

Author mode creates stacks and fields through the SQL API. History records changes and supports rewinding, but concurrent edits use last-save-wins semantics. Clients can subscribe to notifications to refresh displayed data. The version currently provides no scripting engine, model agent or inter-instance link; the worker registration function is still empty.

### Maintenance

The extension registers application configuration/history tables for `pg_dump`; system stacks are recreated by installation. Preserve backing tables and role grants in backups. There is no supported in-place SQL upgrade script in this source revision: assess migration separately instead of assuming `ALTER EXTENSION` can upgrade it. Removing the library from preload and restarting affects its cluster-wide connection guard.
