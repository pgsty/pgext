## Usage

Sources:

- [Official upstream README](https://github.com/jnidzwetzki/pg-dev-container/blob/d7cd5db9481e8fbc7f2d88139e40564463657b10/README.md)
- [Official extension control file (scan.control)](https://github.com/jnidzwetzki/pg-dev-container/blob/d7cd5db9481e8fbc7f2d88139e40564463657b10/src/extensions/06_scan/scan.control)
- [Official extension SQL (scan--1.0.sql)](https://github.com/jnidzwetzki/pg-dev-container/blob/d7cd5db9481e8fbc7f2d88139e40564463657b10/src/extensions/06_scan/scan--1.0.sql)

`scan` — Visual Studio Code - Development Container - PostgreSQL. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION scan;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `full_table_scan(REGCLASS)` is an extension function and returns `VOID`.
- `get_attribute_type(tablename REGCLASS, attrname TEXT)` is an extension function and returns `OID`.
- `table_scan_and_sort_attribute(tablename REGCLASS, attrname TEXT)` is an extension function and returns `VOID`.
- `table_scan_with_index(tablename REGCLASS, indexname REGCLASS)` is an extension function and returns `VOID`.
- `table_scan_with_scankeys(REGCLASS)` is an extension function and returns `VOID`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
