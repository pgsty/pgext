## Usage

Sources:

- [Official extension control file](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/31ee56f66b9df44e4fca083d83add588630af7a3/iban.control)
- [Official upstream documentation](https://github.com/yorickdewid/PostgreSQL-IBAN/blob/31ee56f66b9df44e4fca083d83add588630af7a3/README.md)

`iban` — IBAN data type and validation functions for PostgreSQL

The reviewed catalog snapshot records version `1.1`, kind `standard`, and implementation language `C++`.
The curated compatibility set is `14`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "iban";
SELECT extversion
FROM pg_extension
WHERE extname = 'iban';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
