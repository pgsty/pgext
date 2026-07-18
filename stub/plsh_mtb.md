## Usage

Sources:

- [Official extension control file](https://github.com/credativ/plsh-mtb/blob/ec5de1202ba185acb96561d799797811c909a389/plsh_mtb.control)
- [Official upstream documentation](https://github.com/credativ/plsh-mtb/blob/ec5de1202ba185acb96561d799797811c909a389/README.md)

`plsh_mtb` — A PostgreSQL multi-tenant backup controller implemented with PL/SH and an installed Korn shell command.

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `Shell`.
Install and validate the declared extension dependencies first: `plsh`.

```sql
CREATE EXTENSION "plsh_mtb";
SELECT extversion
FROM pg_extension
WHERE extname = 'plsh_mtb';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
