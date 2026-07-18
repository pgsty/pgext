## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/pgzint/)

`pgzint` — Generate PNG barcodes and QR codes inside PostgreSQL using the Zint barcode library.

The reviewed catalog snapshot records version `0.1.4`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgzint";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgzint';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
