## Usage

Sources:

- [README 0.3.0](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/README.md)
- [Installation and recovery](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/docs/INSTALLATION.md)
- [Control](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/pg_govlite_single.control)
- [Versioned SQL migration](https://github.com/AddressForAll/pg_govLite/blob/0.3.0/extension/pg_govlite_single--0.2.0--0.3.0.sql)

`pg_govlite_single` 0.3.0 is pg_govLite's single-schema governance layout. Framework objects use prefixes inside an existing chosen schema. It is an alternative to `pg_govlite`, and the two layouts should not be installed together in one database.

### Enable in a Chosen Schema

Upstream requires PostgreSQL 15 or later, `plpython3u`, and PyYAML plus Chevron in the server's PL/Python environment. A superuser installs this non-relocatable extension. The example creates a dedicated schema:

```sql
CREATE SCHEMA cliente;
CREATE EXTENSION pg_govlite_single
  WITH SCHEMA cliente VERSION '0.3.0' CASCADE;
SELECT * FROM cliente.gvlt_governance_check();
SELECT * FROM cliente.gvlt_policy_check();
SELECT cliente.lib_yaml_to_jsonb('name: sample');
SELECT cliente.lib_jsonb_mustache_render('Hello {{name}}', '{"name":"world"}'::jsonb);
```

Choose a lowercase schema identifier of at most 63 bytes and keep administrative ownership. Dependency packages must already exist on the server. No preload or server restart is specified.

### Objects and Policy

`cliente.gvlt_tagobj_include` associates governed tags with objects. `cliente.gvlt_object_name` and `cliente.gvlt_object_name_parse` resolve the layout's physical names; `cliente.gvlt_medallion_objects` inspects registered Medallion relations. The single layout records layers with prefixes instead of creating a schema for each layer.

Restrictive governance is disabled by default. When enabled by the DBA, it uses database-wide event triggers and requires an existing restricted master login. Policy changes can remove non-master writes; they do not supply application read authorization. Do not rename installed objects by editing the naming dictionary.

### Access and Upgrades

Functions start with owner-only access; grant schema `USAGE` and function `EXECUTE` only as needed. `cliente.lib_yamlfile_to_jsonb` uses `pg_read_file` permissions, while optional Mustache filesystem partials use the server process's access.

Upgrades reuse the registered schema. The 0.2.0-to-0.3.0 script preserves governance data and rejects function-name collisions. Standalone adoption and layout conversion require separate migrations. Ordinary `pg_dump` omits native extension-table contents in this baseline; recovery must preserve governance state and object OID references.
