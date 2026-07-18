## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/zxdvd/pghelper/blob/04e91fb5722396cfe772cee9be9edc07a67f035f/README.md)
- [Extension SQL directory](https://github.com/zxdvd/pghelper/tree/04e91fb5722396cfe772cee9be9edc07a67f035f/sql)
- [Extension control file](https://github.com/zxdvd/pghelper/blob/04e91fb5722396cfe772cee9be9edc07a67f035f/myhelper.control)

`myhelper` is a personal pure-SQL toolkit collected into the `helper` schema. The reviewed tree includes array helpers, safe-looking conversion helpers, URL decoding, date utilities, privilege inspection, table description, lock and bloat reports, and table-switching scripts.

```sql
CREATE EXTENSION myhelper;
SELECT helper.array_uniq(ARRAY[3, 1, 3, 2]);
SELECT * FROM helper.unnest_with_index(ARRAY['a', 'b']);
SELECT helper.get_days_of_month(current_timestamp);
```

Treat the examples as discovery only and inspect the installed signatures with `\df helper.*`. This is a heterogeneous personal toolkit rather than a versioned public API; administrative and table-switching modules can inspect or modify database objects and may use dynamic SQL.

Install first in a disposable database, compare every created object with the pinned SQL, revoke execution from `PUBLIC`, and grant only individual read-only helpers. Test search paths, identifier quoting, null and exception behavior, privilege checks, concurrent DDL, and version compatibility. Prefer maintained dedicated tools for bloat, locks, privileges, and online table changes; do not import the combined SQL into a managed service without a full statement-by-statement review.
