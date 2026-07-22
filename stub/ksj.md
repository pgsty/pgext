## Usage

Sources:

- [Version 1.0 SQL objects](https://github.com/nuko-yokohama/ksj/blob/6ae22bfa2d1fcfb59d55824f388d193cc0252a7e/ksj--1.0.sql)
- [Type implementation](https://github.com/nuko-yokohama/ksj/blob/6ae22bfa2d1fcfb59d55824f388d193cc0252a7e/ksj.c)
- [Extension control file](https://github.com/nuko-yokohama/ksj/blob/6ae22bfa2d1fcfb59d55824f388d193cc0252a7e/ksj.control)

`ksj` defines a fixed-length PostgreSQL type whose text form uses Japanese kanji numerals while its internal value is a 32-bit integer. It supplies arithmetic, comparisons, `sum`, `min`, `max`, integer casts, and a B-tree operator class.

```sql
CREATE EXTENSION ksj;
SELECT '百二十三'::ksj + '七'::ksj;
CREATE INDEX ON ledger (amount_kanji);
```

The reviewed repository has no user documentation or regression tests, so accepted grammar and output spelling must be established from the C implementation on the exact build. The 32-bit representation also means arithmetic and casts need explicit boundary and overflow tests.

Do not use display text as a stable interchange or collation key. Preserve the original source text when wording matters, and use a core numeric type for calculations, constraints, and external APIs. Treat the `ksj` B-tree ordering as numeric ordering only after validating negative values, zero, malformed input, dump/restore, and upgrade behavior.
