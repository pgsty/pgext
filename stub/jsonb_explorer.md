## Usage

Sources:

- [Official extension control file](https://github.com/erthalion/jsonb_explorer/blob/515b434872ded15b64c88bd4101c061afb15698e/jsonb_explorer.control)
- [Official upstream README](https://github.com/erthalion/jsonb_explorer/blob/515b434872ded15b64c88bd4101c061afb15698e/README.md)

`jsonb_explorer` — Inspect jsonb trees and paths

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "jsonb_explorer";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonb_explorer';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
