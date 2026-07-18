## Usage

Sources:

- [Official upstream documentation](https://pgxn.org/dist/test_factory/doc/test_factory.html)
- [Official PGXN distribution page](https://pgxn.org/dist/test_factory/)
- [Official upstream README](https://github.com/BlueTreble/test_factory/blob/0eb02e0fd8fe0fbd59dc22d4d3b9f86c27678054/README.md)

`test_factory` — Framework for registering, creating, caching, and reusing named PostgreSQL test data sets.

The reviewed catalog snapshot records version `0.5.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "test_factory";
SELECT extversion
FROM pg_extension
WHERE extname = 'test_factory';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
