## Usage

Sources:

- [Upstream README](https://github.com/rlichtenwalter/pg_genotype/blob/e7fa2e88e33c3f4fe4d912c8f86884d904af7a1b/README.md)
- [Extension control file](https://github.com/rlichtenwalter/pg_genotype/blob/e7fa2e88e33c3f4fe4d912c8f86884d904af7a1b/pg_genotype.control)
- [Version 1.0 SQL definition](https://github.com/rlichtenwalter/pg_genotype/blob/e7fa2e88e33c3f4fe4d912c8f86884d904af7a1b/pg_genotype--1.0.sql)
- [Type implementation](https://github.com/rlichtenwalter/pg_genotype/blob/e7fa2e88e33c3f4fe4d912c8f86884d904af7a1b/pg_genotype.c)

`pg_genotype` version `1.0` adds a compact `genotype` base type for two-character allele pairs. Its input routine accepts the symbols A, C, G, T, I, and D in either position and preserves their order, so `AT` and `TA` are distinct values.

### Store allele pairs

```sql
CREATE EXTENSION pg_genotype;

CREATE TABLE sample_genotypes (
    sample_id bigint PRIMARY KEY,
    call genotype NOT NULL
);

INSERT INTO sample_genotypes VALUES
    (1, 'AT'),
    (2, 'TA'),
    (3, 'ID');

SELECT sample_id, call::text
FROM sample_genotypes
ORDER BY sample_id;
```

The extension defines type input/output only: it does not provide genotype operators, comparison semantics, casts, or an index operator class. Although the README describes a one-byte packed representation, the SQL definition declares a four-byte pass-by-value datum. The C parser reads the first two characters without proving that the input ends there, so applications should enforce an exact `^[ACGTID]{2}$` check before casting. The repository predates current PostgreSQL releases; compile and regression-test it against the exact server version before depending on its storage representation.
