## Usage

Sources:

- [Official extension control file](https://github.com/goldenhelix/tsf-fdw/blob/bf8d839bf447d3300cd371b9e059b612b9330923/tsf_fdw.control)

`tsf_fdw` — Golden Helix C foreign data wrapper for reading one or more TSF genomic data files as PostgreSQL foreign tables.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "tsf_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'tsf_fdw';
```

The upstream project is associated with `Golden Helix`; verify its current support, license, packaging, and deployment boundary from the linked source.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
