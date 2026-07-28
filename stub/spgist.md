## Usage

Sources:

- [Official upstream README](https://github.com/duckbill/sp-gist/blob/fd8516ec426ae0c7f0ae508e326fae7dcf674d1d/README.md)
- [Official extension control file (spgist.control)](https://github.com/duckbill/sp-gist/blob/fd8516ec426ae0c7f0ae508e326fae7dcf674d1d/spgist.control)
- [Official extension SQL (spgist--1.0.sql)](https://github.com/duckbill/sp-gist/blob/fd8516ec426ae0c7f0ae508e326fae7dcf674d1d/spgist--1.0.sql)

`spgist` — you can user it to test the sp-gist index (postresql). Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION spgist;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `spg_quad_choose(internal, internal)` is an extension function and returns `internal`.
- `spg_quad_config(internal)` is an extension function and returns `internal`.
- `spg_quad_inner_consistent(internal, internal)` is an extension function and returns `internal`.
- `spg_quad_leaf_consistent(internal, internal, internal)` is an extension function and returns `internal`.
- `spg_quad_picksplit(internal, internal)` is an extension function and returns `internal`.
- `spg_text_choose(internal, internal)` is an extension function and returns `internal`.
- `spg_text_config(internal)` is an extension function and returns `internal`.
- `spg_text_inner_consistent(internal, internal)` is an extension function and returns `internal`.
- `spg_text_leaf_consistent(internal, internal, internal)` is an extension function and returns `internal`.
- `spg_text_picksplit(internal, internal)` is an extension function and returns `internal`.
- `spgbeginscan(internal)` is an extension function and returns `internal`.
- `spgbuild(internal)` is an extension function and returns `internal`.
- `spgbuildempty(internal)` is an extension function and returns `internal`.
- `spgbulkdelete(internal)` is an extension function and returns `internal`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
