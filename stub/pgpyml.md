## Usage

Sources:

- [Official upstream documentation](https://minoro.is-a.dev/pgpyml/get-started/)
- [Official PGXN distribution page](https://pgxn.org/dist/pgpyml/)
- [Official project or provider page](https://minoro.is-a.dev/pgpyml/)

`pgpyml` — Run scikit-learn machine-learning models inside PostgreSQL through PL/Python functions and triggers.

The reviewed catalog snapshot records version `0.3.1`, kind `puresql`, and implementation language `Python`.
Install and validate the declared extension dependencies first: `plpython3u`.

```sql
CREATE EXTENSION "pgpyml";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgpyml';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
