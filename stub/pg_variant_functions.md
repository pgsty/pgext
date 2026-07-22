## Usage

Sources:

- [Official README](https://github.com/rlichtenwalter/pg_variant_functions/blob/8ad88883b840ee334aeb4d549aacfa102dceae11/README.md)
- [Version 1.0 extension SQL](https://github.com/rlichtenwalter/pg_variant_functions/blob/8ad88883b840ee334aeb4d549aacfa102dceae11/pg_variant_functions--1.0.sql)
- [C implementation](https://github.com/rlichtenwalter/pg_variant_functions/blob/8ad88883b840ee334aeb4d549aacfa102dceae11/pg_variant_functions.c)

`pg_variant_functions` computes call rates and alternate-allele frequencies from genotype dosage arrays stored as `TINYINT[]`. It is a small, abandoned work-in-progress rather than a general genomics toolkit, so use it only when its exact array convention matches the data model.

### Core Workflow

Install the required `tinyint` extension first. Genotype positions are addressed by one-based `INTEGER[]` subscripts, and non-NULL genotype values are summed as alternate-allele dosages and divided by twice the number of called genotypes.

```sql
CREATE EXTENSION tinyint;
CREATE EXTENSION pg_variant_functions;

SELECT alternate_allele_frequency(
    ARRAY[0, 1, 2, NULL]::tinyint[],
    ARRAY[1, 2, 3, 4]::integer[]
);

SELECT *
FROM summarize_variant(
    ARRAY[0, 1, 2, NULL]::tinyint[],
    ARRAY[1, 2]::integer[],
    ARRAY[3, 4]::integer[]
);
```

`alternate_allele_frequency(genotypes, indices)` returns a `REAL` value for the selected positions and returns NULL when none of those positions has a called genotype.

`summarize_variant(genotypes, subset1, subset2)` returns the `variant_summary` composite type with these fields:

- `call_rate` and `maf` for the complete genotype array.
- `subset1_call_rate` and `subset1_maf` for the first optional index array.
- `subset2_call_rate` and `subset2_maf` for the second optional index array.

Passing NULL for a subset leaves that subset's fields NULL. Despite the field name `maf`, the implementation reports the alternate-allele frequency directly and does not transform values above `0.5` into minor-allele frequency.

### Caveats

The code expects a one-dimensional dosage array and valid positive one-based indices. It does not validate biological ploidy or genotype range, and it should not receive zero, negative, or otherwise invalid subscripts. Empty arrays and empty subsets need application-level handling. Both functions are immutable and parallel safe, but carry a declared cost of `10000` because they scan arrays. Version `1.0` is marked abandoned; validate numeric results against a trusted implementation before analytical or clinical use.
