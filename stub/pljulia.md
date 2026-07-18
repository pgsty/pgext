## Usage

Sources:

- [Official upstream source](https://gitlab.com/pljulia/pljulia)

`pljulia` — Embeds Julia as an untrusted PostgreSQL procedural language with SPI access, DO blocks, triggers, arrays, and composite-type conversion.

The reviewed catalog snapshot records version `0.8`, kind `standard`, and implementation language `C`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pljulia";
SELECT extversion
FROM pg_extension
WHERE extname = 'pljulia';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
