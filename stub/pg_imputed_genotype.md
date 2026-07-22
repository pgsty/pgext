## Usage

Sources:

- [Official README at the reviewed commit](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/README.md)
- [Version 1.0 type definition](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/pg_imputed_genotype--1.0.sql)
- [Type input and output implementation](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/pg_imputed_genotype.c)

`pg_imputed_genotype` adds the fixed-size `imputed_genotype` base type for storing three imputed genotype probabilities. It packs three 10-bit integers into one four-byte, pass-by-value datum. Use it only when compact storage and 0.001-scale quantization are acceptable.

### Core Workflow

```sql
CREATE EXTENSION pg_imputed_genotype;

CREATE TABLE sample_call (
  sample_id bigint PRIMARY KEY,
  probabilities imputed_genotype NOT NULL
);

INSERT INTO sample_call VALUES (1, '{0.2,0.5,0.1}');
SELECT probabilities FROM sample_call WHERE sample_id = 1;
-- {0.200,0.500,0.100}
```

Input is written as three comma-separated probabilities inside braces. Values are multiplied by 1000 and truncated rather than rounded, then packed. A total below one is allowed, so any residual probability remains implicit. The extension defines only the base type input/output path; it provides no comparison operators, casts, or index operator classes.

### Validation and Compatibility Boundaries

The implementation allows an integerized total up to 1001 as a rounding tolerance, so some inputs slightly above one can be accepted. Its parser also does not reliably reject trailing text, and it does not explicitly guard non-finite floating-point values. Validate exactly three finite, nonnegative values and their sum in application or SQL code before casting.

Quantization loses decimal precision, and the packed representation is an extension-specific binary format. The old repository has no release tags or PostgreSQL compatibility matrix. Test input, output, dump/restore, and binary behavior on the exact PostgreSQL major and CPU architecture before storing durable data.
