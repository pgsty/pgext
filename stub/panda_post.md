## Usage

Sources:

- [Official extension control file (panda_post.control)](https://api.pgxn.org/src/panda_post/panda_post-0.2.1/panda_post.control)
- [Official extension SQL (panda_post.sql)](https://api.pgxn.org/src/panda_post/panda_post-0.2.1/sql/panda_post.sql)

`panda_post` — Python Pandas data in Postgres. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION panda_post;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `create_cast(data_type text , transform text DEFAULT '' , cast_type text DEFAULT NULL , create_array_cast boolean DEFAULT true)` is an extension function and returns `void`.
- `eval(i text)` is an extension function and returns `ndarray`.
- `ndall(i ndarray , axis int , keepdims boolean=False)` is an extension function and returns `ndarray`.
- `ndall(i ndarray , axis int[] , keepdims boolean=False)` is an extension function and returns `ndarray`.
- `ndall(i ndarray , keepdims boolean=False)` is an extension function and returns `ndarray`.
- `ndany(i ndarray , axis int , keepdims boolean=False)` is an extension function and returns `ndarray`.
- `ndany(i ndarray , axis int[] , keepdims boolean=False)` is an extension function and returns `ndarray`.
- `ndany(i ndarray , keepdims boolean=False)` is an extension function and returns `ndarray`.
- `ndarray_from_plpython(internal)` is an extension function and returns `ndarray`.
- `ndarray_in(cstring)` is an extension function and returns `ndarray`.
- `ndarray_out(ndarray)` is an extension function and returns `cstring`.
- `ndarray_to_plpython(internal)` is an extension function and returns `internal`.
- `ndunique(ar ndarray , return_index boolean = False , return_inverse boolean = False , return_counts boolean = False)` is an extension function and returns `ndarray[]`.
- `ndunique1(ar ndarray)` is an extension function and returns `ndarray`.

### Requirements and Caveats

- The reviewed control file declares default version `0.2.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
