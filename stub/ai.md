## Usage

Sources:

- [Official extension control file](https://github.com/timescale/pgai/blob/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/projects/extension/sql/output/ai.control)
- [Official upstream documentation](https://github.com/timescale/pgai/blob/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/projects/extension/README.md)
- [Official upstream README](https://github.com/timescale/pgai/blob/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/README.md)

`ai` — Helper functions for AI workflows in PostgreSQL

The reviewed catalog snapshot records version `0.11.3-dev`, kind `puresql`, and implementation language `Python`.
Install and validate the declared extension dependencies first: `plpython3u`, `vector`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "ai";
SELECT extversion
FROM pg_extension
WHERE extname = 'ai';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
