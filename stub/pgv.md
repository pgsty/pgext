## Usage

Sources:

- [Official upstream README](https://github.com/jac30b/pgv/blob/75fc2437b607af80726ac0b2489d8ea30ab39c52/README.md)
- [Official extension control file (pgv.control)](https://github.com/jac30b/pgv/blob/75fc2437b607af80726ac0b2489d8ea30ab39c52/pgv.control)
- [Official extension SQL (pgv--0.0.1.sql)](https://github.com/jac30b/pgv/blob/75fc2437b607af80726ac0b2489d8ea30ab39c52/sql/pgv--0.0.1.sql)

`pgv` — Learning how postgres extensions work based on pgvector and postgresql extension. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgv;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `vec_cmp(vec, vec)` is an extension function and returns `integer`.
- `vec_cosine_distance(vec, vec)` is an extension function and returns `float4`.
- `vec_eq(vec, vec)` is an extension function and returns `boolean`.
- `vec_ge(vec, vec)` is an extension function and returns `boolean`.
- `vec_gt(vec, vec)` is an extension function and returns `boolean`.
- `vec_input(cstring, oid, integer)` is an extension function and returns `vec`.
- `vec_le(vec, vec)` is an extension function and returns `boolean`.
- `vec_lt(vec, vec)` is an extension function and returns `boolean`.
- `vec_ne(vec, vec)` is an extension function and returns `boolean`.
- `vec_output(vec)` is an extension function and returns `cstring`.
- `vec_typemodifier_in(cstring[])` is an extension function and returns `integer`.
- `vec_typemodifier_out(integer)` is an extension function and returns `cstring`.
- `vec` is an extension-defined type.
- `vec_ops` is an extension-defined operator class.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
