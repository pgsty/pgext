## Usage

Sources:

- [Official extension control file](https://github.com/fabriziomello/hba_fdw/blob/fcc1cd8eee8f997538b917f095caef5373a83e7d/hba_fdw.control)
- [Official upstream documentation](https://github.com/fabriziomello/hba_fdw/blob/fcc1cd8eee8f997538b917f095caef5373a83e7d/README.md)

`hba_fdw` — Work-in-progress foreign data wrapper scaffold intended to expose pg_hba.conf through SQL.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "hba_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'hba_fdw';
```
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
