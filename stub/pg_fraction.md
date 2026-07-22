## Usage

Sources:

- [Official README](https://github.com/nuko-yokohama/pg_fraction/blob/812f0599f82a47024b96ed0c745c0464bf1e2f2b/README.md)
- [Official extension control file](https://github.com/nuko-yokohama/pg_fraction/blob/812f0599f82a47024b96ed0c745c0464bf1e2f2b/pg_fraction.control)

`pg_fraction` adds a `fraction` type for exact numerator/denominator storage, basic arithmetic, comparisons, aggregates, casts, and B-tree indexing. Upstream describes it as a sample extension rather than software intended for commercial use.

### Core Workflow

Create the extension, store fraction literals, and use the supplied operators:

```sql
CREATE EXTENSION pg_fraction;

CREATE TABLE ratios (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  value fraction NOT NULL
);

INSERT INTO ratios (value) VALUES ('1/2'), ('2/4'), ('3/5');

SELECT value,
       value + '1/3'::fraction AS increased,
       get_value(value) AS approximate_value
FROM ratios
ORDER BY value;
```

Fractions are reduced automatically, so equivalent inputs such as `2/4` and `1/2` normalize to the same value.

### Important Objects

- `fraction` stores a numerator and denominator and accepts slash-separated input.
- `+`, `-`, `*`, and `/` perform fraction arithmetic.
- Comparison operators plus the B-tree operator class support ordering and indexes.
- `max(fraction)` and `min(fraction)` aggregate fraction values.
- `get_value(fraction)` returns an approximate floating-point value.

### Limits and Caveats

- The implementation restricts each input component to at most five decimal digits and a value no greater than `99999`; arithmetic reports an error when its result exceeds those limits.
- The digit check applies to input before reduction, so a large reducible spelling can be rejected even when its normalized result would fit.
- Upstream comparisons convert fractions to floating point. Verify precision and ordering behavior for close or large values; do not use this type where exact rational comparison is a hard requirement.
- Division by zero, sign handling, casts, overflow, dump/restore, and index ordering should be tested against the exact PostgreSQL build before adopting this sample extension.
