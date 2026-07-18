## Usage

Sources:

- [Official upstream documentation](https://github.com/hooopo/pg-fuzzywuzzy/blob/b02f84ac42f7cc1c5bc51f4daf0c58a68752bf56/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/fuzzywuzzy/)

`fuzzywuzzy` — fuzzywuzzy ratio for postgres

The reviewed catalog snapshot records version `0.0.2`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `fuzzystrmatch`.

```sql
CREATE EXTENSION "fuzzywuzzy";
SELECT extversion
FROM pg_extension
WHERE extname = 'fuzzywuzzy';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
