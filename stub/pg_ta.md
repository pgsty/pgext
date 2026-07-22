## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/magnusp/pg_ta/blob/899a2ce4ad674543d07bd007d2d85f5594baa30a/README.md)
- [Version 1.0 SQL function](https://github.com/magnusp/pg_ta/blob/899a2ce4ad674543d07bd007d2d85f5594baa30a/sql/pg_ta--1.0.sql)
- [C implementation](https://github.com/magnusp/pg_ta/blob/899a2ce4ad674543d07bd007d2d85f5594baa30a/src/pg_ta.c)
- [Upstream test query](https://github.com/magnusp/pg_ta/blob/899a2ce4ad674543d07bd007d2d85f5594baa30a/test.sql)

`pg_ta` is a 2014 proof of concept for calling TA-Lib from PostgreSQL. Version 1.0 exposes only `ta_f(double precision[])`, which returns a set representing a simple moving average with a hard-coded period of 5.

### Core Workflow

Builds require the external TA-Lib headers and library. Array order defines observation order; PostgreSQL does not infer ordering from a source query. The first four returned rows are NULL because the five-value moving-average window is not yet complete.

```sql
CREATE EXTENSION pg_ta;

SELECT *
FROM ta_f(ARRAY[1, 2, 3, 4, 5, 6]::double precision[]);
-- NULL, NULL, NULL, NULL, 3, 4
```

The function rejects arrays containing NULL. There is no parameter for the period or TA-Lib moving-average type, and no other indicators are implemented.

### Critical Safety Boundary

Do not execute this extension outside an isolated test instance. The reviewed C code allocates its PostgreSQL result array for TA-Lib's shorter `outNbElement` count, then writes one entry for every input element. For a period of 5 this writes four `Datum` values beyond the allocation and can corrupt backend memory. It also treats a zero `Datum` as the NULL marker, so a real moving-average result of `0` can be emitted as SQL NULL.

Upstream explicitly says the code is experimental, unstable, and not for anything remotely close to production. The source predates current PostgreSQL and TA-Lib releases; even apparently correct sample output is not evidence of memory safety or compatibility.
