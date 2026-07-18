## Usage

Sources:

- [Official extension control file](https://github.com/MuhammadTahaNaveed/pg-gsheets/blob/f223b892ff2a0dcd98c5f2ac0ca8d748a3b5bb28/gsheets.control)

`gsheets` — Read and write Google Sheets from PostgreSQL.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "gsheets";
SELECT extversion
FROM pg_extension
WHERE extname = 'gsheets';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
