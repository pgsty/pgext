## Usage

Sources:

- [Official upstream README](https://github.com/open-gpdb/heapy/blob/3b16e36b362be564492a61d1f01e66086167e453/README.HOT)
- [Official extension control file (heapy.control)](https://github.com/open-gpdb/heapy/blob/3b16e36b362be564492a61d1f01e66086167e453/heapy.control)
- [Official extension SQL (heapy--1.0.sql)](https://github.com/open-gpdb/heapy/blob/3b16e36b362be564492a61d1f01e66086167e453/heapy--1.0.sql)

`heapy` — The Heap Only Tuple (HOT) feature eliminates redundant index entries and allows the re-use of space taken by DELETEd or obsoleted UPDATEd tuples without performing a table-wide vacuum. It does this by allowing single-page vacuuming, also called "defragmentation". Use it for the corresponding analytical or storage workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION heapy;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `heapy_define_relation_offload_policy_internal(reloid OID)` is an extension function and returns `void`.
- `heapy_define_relation_offload_policy_internal_seg(reloid OID)` is an extension function and returns `void`.
- `heapy_tableam_handler(internal)` is an extension function and returns `table_am_handler`.
- `heapy` is an extension-defined access method.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
