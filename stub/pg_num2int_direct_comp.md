## Usage

Sources:

- [Official documentation](https://github.com/datastone-inc/pg_num2int_direct_comp/blob/7dba5a74b9a7e4515664cbd601847e9abbd687fa/README.md)
- [Extension control file](https://github.com/datastone-inc/pg_num2int_direct_comp/blob/7dba5a74b9a7e4515664cbd601847e9abbd687fa/pg_num2int_direct_comp.control)
- [Official repository](https://github.com/datastone-inc/pg_num2int_direct_comp)

`pg_num2int_direct_comp` Exact numeric/float-to-integer comparison operators that preserve index use.

### Enablement

Install the files for the intended server, then create `pg_num2int_direct_comp` in the target database:

```sql
CREATE EXTENSION pg_num2int_direct_comp;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- Example: API passes product ID as numeric parameter
CREATE TABLE products(id int4 PRIMARY KEY, parent int4 REFERENCES products, name text);
PREPARE find_product(numeric) AS SELECT * FROM products WHERE id = $1;

-- Without this extension: sequential scan (casts indexed column)
EXPLAIN (COSTS OFF) EXECUTE find_product(42);
--  Seq Scan on products
--    Filter: ((id)::numeric = '42'::numeric)   ← full table scan!

-- With this extension: index scan (transforms to integer comparison)
EXPLAIN (COSTS OFF) EXECUTE find_product(42);
--  Index Scan using products_pkey on products
--    Index Cond: (id = 42)                     ← uses primary key index
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `float4_cmp_int2` | FUNCTION | Callable function from the reviewed install surface. |
| `float4_cmp_int4` | FUNCTION | Callable function from the reviewed install surface. |
| `float4_cmp_int8` | FUNCTION | Callable function from the reviewed install surface. |
| `float4_eq_int2` | FUNCTION | Callable function from the reviewed install surface. |
| `float4_eq_int4` | FUNCTION | Callable function from the reviewed install surface. |
| `float4_eq_int8` | FUNCTION | Callable function from the reviewed install surface. |
| `float4_ge_int2` | FUNCTION | Callable function from the reviewed install surface. |
| `float4_ge_int4` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 12, 13, 14, 15, 16; do not infer unlisted majors.
