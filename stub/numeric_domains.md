## Usage

Sources:

- [Official extension control file](https://github.com/cahutton/numeric_domains/blob/50496725313939a340f0c8129d9e9f2491009e51/numeric_domains.control)
- [Official upstream documentation](https://github.com/cahutton/numeric_domains/blob/50496725313939a340f0c8129d9e9f2491009e51/README.md)

`numeric_domains` — Custom numeric domains for PostgreSQL

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.
The curated compatibility set is `10`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "numeric_domains";
SELECT extversion
FROM pg_extension
WHERE extname = 'numeric_domains';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
