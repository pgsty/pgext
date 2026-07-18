## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/README.md)
- [Version 1.0 type definition](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/pg_imputed_genotype--1.0.sql)
- [C implementation](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/pg_imputed_genotype.c)
- [PostgreSQL license](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/LICENSE)

`pg_imputed_genotype` adds the `imputed_genotype` base type for compact storage of three genotype probabilities. Each component is quantized to a 10-bit integer from 0 through 1000, and the complete triple occupies four bytes.

```sql
CREATE EXTENSION pg_imputed_genotype;
SELECT '{0.2,0.5,0.1}'::imputed_genotype;
-- {0.200,0.500,0.100}
```

Input must contain exactly three probabilities, and their sum must not exceed 1.0. The type is intended for dense arrays of imputed genotype data where fixed precision and storage size matter more than preserving arbitrary decimal precision.

The repository has no release tags or compatibility matrix and its last reviewed source is old. Test binary compatibility against the exact PostgreSQL major version and CPU architecture in use, and verify that 0.001 quantization is acceptable for downstream statistical analysis before migrating data.
