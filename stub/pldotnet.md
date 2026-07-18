## Usage

Sources:

- [Official project or provider page](https://pldotnet.brickabode.com)

`pldotnet` — add C# and F# procedural languages to PostgreSQL through .NET

The reviewed catalog snapshot records version `0.9`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pldotnet";
SELECT extversion
FROM pg_extension
WHERE extname = 'pldotnet';
```

The curated lifecycle is `preview`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
