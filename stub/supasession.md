## Usage

Sources:

- [Official upstream README](https://github.com/snehil-shah/supasession/blob/f09bed26cdac28b2394cd05c68c122b67337075a/README.md)
- [Official extension control file (supasession.control)](https://github.com/snehil-shah/supasession/blob/f09bed26cdac28b2394cd05c68c122b67337075a/supasession.control)
- [Official extension SQL (supasession.sql)](https://github.com/snehil-shah/supasession/blob/f09bed26cdac28b2394cd05c68c122b67337075a/supasession.sql)

`supasession` — > [!WARNING] > This extension is installed in the supasession schema and can potentially cause namespace collisions if you already had one before. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION supasession;

SELECT dbdev.install('Snehil_Shah@supasession');
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `supasession.disable()` is an extension function and returns `void`.
- `supasession.enable()` is an extension function and returns `void`.
- `supasession.get_config()` is an extension function and returns `supasession`.
- `supasession.set_config(enabled BOOLEAN DEFAULT NULL, max_sessions INTEGER DEFAULT NULL, strategy supasession.enforcement_strategy DEFAULT NULL)` is an extension function and returns `supasession`.
- `supasession.sid()` is an extension function and returns `uuid`.
- `supasession.enforcement_strategy` is an extension-defined type.
- `supasession.config` is a table installed or managed by the extension.
- `supasession` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.2`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
