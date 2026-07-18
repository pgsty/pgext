## Usage

Sources:

- [Official extension control file](https://github.com/paroga/pg-cbor/blob/23afc682e6168e59c4d27d64fb10e8c5e624722a/cbor.control)
- [Official upstream documentation](https://github.com/paroga/pg-cbor/blob/23afc682e6168e59c4d27d64fb10e8c5e624722a/README.md)

`cbor` — Concise Binary Object Representation data type for PostgreSQL.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "cbor";
SELECT extversion
FROM pg_extension
WHERE extname = 'cbor';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
