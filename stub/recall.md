## Usage

Sources:

- [Official extension control file](https://github.com/mreithub/pg_recall/blob/87126883d6272a324b259b83f96436a53980dbd1/recall.control)

`recall` — Tracks table changes in per-table _log tables for timestamp-based history queries.

The reviewed catalog snapshot records version `0.9.6`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `btree_gist`, `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "recall";
SELECT extversion
FROM pg_extension
WHERE extname = 'recall';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
