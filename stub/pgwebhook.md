## Usage

Sources:

- [Official upstream README](https://github.com/mansueli/tle/blob/9e1c8ce5a7d7a2f334eb21894d3e8994fb7f28ad/pgwebhook/README.md)
- [Official extension control file (pgwebhook.control)](https://github.com/mansueli/tle/blob/9e1c8ce5a7d7a2f334eb21894d3e8994fb7f28ad/pgwebhook/pgwebhook.control)
- [Official extension SQL (pgwebhook--0.1.1--0.1.2.sql)](https://github.com/mansueli/tle/blob/9e1c8ce5a7d7a2f334eb21894d3e8994fb7f28ad/pgwebhook/pgwebhook--0.1.1--0.1.2.sql)

`pgwebhook` — pgwebhook is a PostgreSQL extension designed to facilitate the creation and management of webhooks directly from your database. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pgwebhook;

SELECT dbdev.install('mansueli@pgwebhook');
CREATE EXTENSION "mansueli@pgwebhook" VERSION '0.1.1';
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hook.edge_wrapper` is an extension function.
- `hook.edgehook_trigger()` is an extension function and returns `trigger`.
- `hook.http_request` is an extension function.
- `hook.webhook_trigger()` is an extension function and returns `trigger`.
- `hook.migrations` is a table installed or managed by the extension.
- `hook` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.2`.
- Install the confirmed extension dependencies first: `http`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
