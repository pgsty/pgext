## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/salesforce_id/)

`salesforce_id` — Compact 12-byte PostgreSQL type with casts, comparison operators, and B-tree/hash operator classes for Salesforce IDs.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "salesforce_id";
SELECT extversion
FROM pg_extension
WHERE extname = 'salesforce_id';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
