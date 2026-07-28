## Usage

Sources:

- [Official upstream README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_sparse_vector/README)
- [Official extension control file (gp_sparse_vector.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_sparse_vector/gp_sparse_vector.control)
- [Official extension SQL (gp_sparse_vector--1.0.0--1.0.1.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_sparse_vector/gp_sparse_vector--1.0.0--1.0.1.sql)

`gp_sparse_vector` — Say for example that we have the following array of doubles in Postgres stored as a "float8[]": '{0, 33,...40,000 zeros..., 12, 22 }'::float8[]. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION gp_sparse_vector;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `dimension(svec)` is an extension function and returns `integer`.
- `dmax(float8,float8)` is an extension function and returns `float8`.
- `dmin(float8,float8)` is an extension function and returns `float8`.
- `dot(float8[],float8[])` is an extension function and returns `float8`.
- `dot(float8[],svec)` is an extension function and returns `float8`.
- `dot(svec,float8[])` is an extension function and returns `float8`.
- `dot(svec,svec)` is an extension function and returns `float8`.
- `float8arr_cast_float4(float4)` is an extension function and returns `float8[]`.
- `float8arr_cast_float8(float8)` is an extension function and returns `float8[]`.
- `float8arr_cast_int2(int2)` is an extension function and returns `float8[]`.
- `float8arr_cast_int4(int4)` is an extension function and returns `float8[]`.
- `float8arr_cast_int8(bigint)` is an extension function and returns `float8[]`.
- `float8arr_cast_numeric(numeric)` is an extension function and returns `float8[]`.
- `float8arr_div_float8arr(float8[],float8[])` is an extension function and returns `svec`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.1`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
