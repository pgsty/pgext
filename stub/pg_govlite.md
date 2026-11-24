## Usage

Sources:

- [README 0.3.0](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/README.md)
- [Installation and recovery](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/docs/INSTALLATION.md)
- [Control](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/pg_govlite.control)
- [Versioned SQL migration](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/extension/pg_govlite--0.2.0--0.3.0.sql)

`pg_govlite` 0.3.0 supplies a metadata and governed-tag framework using multiple schemas, including `gvlt` and `lib`. It is the multi-schema layout of pg_govLite; choose only one layout per database.

### Enable and Inspect

Upstream requires PostgreSQL 15 or later. Install the server's `plpython3u` package and make PyYAML and Chevron available to the Python interpreter used by PL/Python. Activate the extension as a superuser:

```sql
CREATE EXTENSION pg_govlite VERSION '0.3.0' CASCADE;
SELECT * FROM gvlt.governance_check();
SELECT * FROM gvlt.policy_check();
SELECT lib.yaml_to_jsonb('name: sample');
SELECT lib.jsonb_mustache_render('Hello {{name}}', '{"name":"world"}'::jsonb);
```

`CASCADE` enables available extension dependencies; it does not install operating-system or Python packages. No preload or server restart is specified.

### Governance API

`gvlt.tag_include` defines governed tags; `gvlt.tagobj_include` associates tags with database objects. `gvlt.obj_tags`, `gvlt.tag_search` and `gvlt.medallion_objects` inspect metadata. Disabling tags or associations preserves catalog history.

Restrictive governance is off by default. Enabling it is a DBA operation with an existing restricted master login. It changes write permissions and installs database-wide event-trigger enforcement; read authorization and row-security policy remain separate DBA decisions. Review ownership and policy before applying it to existing business tables.

### Permissions and Recovery

Framework functions initially have owner-only access. Grant schema `USAGE` and function `EXECUTE` deliberately. `lib.yamlfile_to_jsonb` reads through `pg_read_file`; Mustache filesystem partials use the PostgreSQL process's filesystem permissions. Keep the optional partials path empty unless server-file access is intended.

The versioned 0.2.0-to-0.3.0 migration preserves governance data. A standalone installation or a change of layout requires a separate migration. Upstream warns that ordinary `pg_dump` omits native extension-table contents in this baseline: use a recovery method that preserves governance state and its object OID references.
