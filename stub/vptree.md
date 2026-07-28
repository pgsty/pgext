## Usage

Sources:

- [Official upstream README](https://github.com/fake-name/pg-spgist_hamming/blob/9fa70b08e0f0108de6a6673ce095c86a987d261d/README.md)
- [Official extension control file (vptree.control)](https://github.com/fake-name/pg-spgist_hamming/blob/9fa70b08e0f0108de6a6673ce095c86a987d261d/vptree/vptree.control)
- [Official extension SQL (vptree--1.0.sql)](https://github.com/fake-name/pg-spgist_hamming/blob/9fa70b08e0f0108de6a6673ce095c86a987d261d/vptree/vptree--1.0.sql)

`vptree` — VP-tree implementation. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION vptree;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `vptree_area_match(int8, vptree_area)` is an extension function and returns `boolean`.
- `vptree_choose(internal, internal)` is an extension function and returns `void`.
- `vptree_config(internal, internal)` is an extension function and returns `void`.
- `vptree_eq_match(int8, int8)` is an extension function and returns `boolean`.
- `vptree_get_distance(int8, int8)` is an extension function and returns `float8`.
- `vptree_inner_consistent(internal, internal)` is an extension function and returns `void`.
- `vptree_leaf_consistent(internal, internal)` is an extension function and returns `boolean`.
- `vptree_picksplit(internal, internal)` is an extension function and returns `void`.
- `vptree_area` is an extension-defined type.
- `vptree_ops` is an extension-defined operator class.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
