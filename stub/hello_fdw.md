## Usage

Sources:

- [Official extension control file](https://github.com/wikrsh/hello_fdw/blob/925ade901f9badcc6ed7e01a1c99b6c0f9e3fab1/hello_fdw.control)
- [Official upstream documentation](https://github.com/wikrsh/hello_fdw/blob/925ade901f9badcc6ed7e01a1c99b6c0f9e3fab1/README.md)

`hello_fdw` — foreign-data wrapper for hello world

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "hello_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'hello_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
