## Usage

Sources:

- [Official upstream README](https://github.com/veranemoloko/pg_jewelry_tools/blob/702b3e118c2a0b0a4b941037a8f6d3093f59788d/README.md)
- [Official extension control file (pg_jewelry_tools.control)](https://github.com/veranemoloko/pg_jewelry_tools/blob/702b3e118c2a0b0a4b941037a8f6d3093f59788d/pg_jewelry_tools.control)
- [Official extension SQL (pg_jewelry_tools--0.0.1.sql)](https://github.com/veranemoloko/pg_jewelry_tools/blob/702b3e118c2a0b0a4b941037a8f6d3093f59788d/sql/pg_jewelry_tools--0.0.1.sql)

`pg_jewelry_tools` — An extension for PSQL with useful functionality for jewelry companies. So far, this is a project created for educational purposes. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_jewelry_tools;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `calculate_gemstone_carat_to_gr(carat double precision)` is an extension function and returns `double`.
- `calculate_gemstone_gr_to_carat(carat double precision)` is an extension function and returns `double`.
- `calculate_metal_weight_gr(metal TEXT, purity INTEGER, volume_mm3 double precision)` is an extension function and returns `double`.
- `public.test_add_one(i integer)` is an extension function and returns `integer`.
- `gemstone` is an extension-defined type.
- `jw_item` is an extension-defined type.
- `precious_metal` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
