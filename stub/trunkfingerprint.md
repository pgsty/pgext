## Usage

Sources:

- [Official upstream documentation](https://github.com/BrandwatchLtd/trunkfingerprint/blob/a4ed83cd533f233e70b23aac3ebb7d91248c7087/README.md)

`trunkfingerprint` — Computes comparable fingerprints of PostgreSQL database structure and optionally data.

The reviewed catalog snapshot records version `1.3.0`, kind `puresql`, and implementation language `SQL`.
The curated compatibility set is `10,11,12,13,14,15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "trunkfingerprint";
SELECT extversion
FROM pg_extension
WHERE extname = 'trunkfingerprint';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
