## Usage

Sources:

- [Official extension control file](https://github.com/AjinkyaTaranekar/ai_toolkit/blob/49c3e95059e2c050e61400829004b70dc08992b2/ai_toolkit.control)
- [Official upstream README](https://github.com/AjinkyaTaranekar/ai_toolkit/blob/49c3e95059e2c050e61400829004b70dc08992b2/README.md)

`ai_toolkit` — AI-powered query generation and explanation tools for PostgreSQL.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C++`.
The curated compatibility set is `18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "ai_toolkit";
SELECT extversion
FROM pg_extension
WHERE extname = 'ai_toolkit';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
