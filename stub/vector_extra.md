## Usage

Sources:

- [README.md](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/README.md)
- [pgv_extra/sql/vector_extra--0.1.0.sql](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/pgv_extra/sql/vector_extra--0.1.0.sql)
- [pgv_extra/meson.build](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/pgv_extra/meson.build)
- [pgv_extra/vector_extra.control](https://github.com/MarchLiu/tensor-dancer/blob/7630942c4d0a82af0e9ca62bc1aa88abdcdb8097/pgv_extra/vector_extra.control)

`vector_extra` is the matrix-vector helper in the Tensor Dancer source distribution. Its versioned SQL defines `pgv_mulmv(matrix bytea, vec vector)`, returning `vector`.

### Core Workflow

```sql
CREATE EXTENSION vector;
CREATE EXTENSION vector_extra;
SELECT oid::regprocedure FROM pg_proc WHERE proname = 'pgv_mulmv';
```

### Operational Boundaries

Install `vector` first: the SQL uses its type although the control omits a dependency declaration. The matrix is an upstream-specific binary representation, not an arbitrary byte array; construct it with the matching project tools. The build links BLAS, LAPACK and GGML. This is an early source prototype with no verified PostgreSQL-major support range or stable matrix interchange contract. Installation is not marked trusted, and no preload is specified.
