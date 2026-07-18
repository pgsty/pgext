## Usage

Sources:

- [Official extension control file](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/pgsfti.control)
- [Official upstream documentation](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/docs/source/intro.rst)
- [Official upstream README](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/README.md)

`pgsfti` — PostgreSQL data type for simple trapezoidal fuzzy time intervals and fuzzy temporal relations.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pgsfti";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgsfti';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
