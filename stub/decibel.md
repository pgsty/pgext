## Usage

Sources:

- [Project README](https://github.com/macenv/pgdecibel/blob/decbd087e08075dee86fa4bbdddadfc736261cc4/README.md)
- [Version 0.0.1 installation SQL](https://github.com/macenv/pgdecibel/blob/decbd087e08075dee86fa4bbdddadfc736261cc4/decibel--0.0.1.sql)
- [C conversion implementation](https://github.com/macenv/pgdecibel/blob/decbd087e08075dee86fa4bbdddadfc736261cc4/decibel.c)

`decibel` 0.0.1 defines a PostgreSQL base type that displays and accepts decibel values while storing a double-precision linear value. Its input path applies 10^(dB/10), and its output path converts the stored value back with 10*log10(value).

### Create and convert values

```sql
CREATE EXTENSION decibel;

SELECT 65.23::decibel AS level_db,
       pascals(65.23::decibel) AS linear_value;
```

The extension exposes `decibelpascal(float8)`, `pascaldecibel(float8)`, and `pascals(decibel)`, plus casts between `decibel` and common numeric types. It also defines comparison and arithmetic operators and `sum`, `avg`, `min`, and `max` aggregates for the type.

### Semantics and compatibility

Addition, subtraction, and averaging between `decibel` values operate on the stored linear representation; adding or subtracting a `float8` adjusts the displayed decibel value. This distinction matters when interpreting results. The API calls its linear quantity pascals, but the implementation supplies no reference-pressure constant and performs the power-ratio conversion shown above. The project is a C extension with a single old 0.0.1 source snapshot and no current compatibility matrix, so test builds and numerical behavior on the target PostgreSQL release.
