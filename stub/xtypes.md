## Usage

Sources:

- [Official extension control file](https://github.com/mkindahl/pg_xtypes/blob/eee55dfe3fb8d02baa36926692d1ccbcee930df6/xtypes.control)

`xtypes` — Unsigned 64-bit byte-size type with human-readable binary-unit input/output, arithmetic, comparison, and rounding.

The reviewed catalog snapshot records version `0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "xtypes";
SELECT extversion
FROM pg_extension
WHERE extname = 'xtypes';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
