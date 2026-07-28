## Usage

Sources:

- [Official upstream README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_debug_numsegments/README.md)
- [Official extension control file (gp_debug_numsegments.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_debug_numsegments/gp_debug_numsegments.control)
- [Official extension SQL (gp_debug_numsegments--1.0.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_debug_numsegments/gp_debug_numsegments--1.0.sql)

`gp_debug_numsegments` — By default all tables are created on all the segments, with this extension we could check or change the default behavior. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION gp_debug_numsegments;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gp_debug_get_create_table_default_numsegments()` is an extension function and returns `text`.
- `gp_debug_reset_create_table_default_numsegments()` is an extension function and returns `void`.
- `gp_debug_reset_create_table_default_numsegments(integer)` is an extension function and returns `void`.
- `gp_debug_reset_create_table_default_numsegments(text)` is an extension function and returns `void`.
- `gp_debug_set_create_table_default_numsegments(integer)` is an extension function and returns `text`.
- `gp_debug_set_create_table_default_numsegments(text)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
