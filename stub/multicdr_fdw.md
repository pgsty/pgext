## Usage

Sources:

- [Official upstream documentation](https://pgxn.org/dist/multicdr_fdw/1.2.2/doc/multicdr_fdw.html)
- [Official PGXN distribution page](https://pgxn.org/dist/multicdr_fdw/)
- [Official upstream README](https://github.com/theirix/multicdr_fdw/blob/ee1b813e8515a8ef8b21893796ba46b593b7bcc4/README.md)

`multicdr_fdw` — Foreign data wrapper for reading call-detail-record file streams as external SQL tables.

The reviewed catalog snapshot records version `1.2.2`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "multicdr_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'multicdr_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
