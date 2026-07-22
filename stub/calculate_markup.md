## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/chobostar/calculate_markup/blob/c70a994c8a215e11d3fd444f0d1412289bcadf3b/README.md)
- [Version 1.0 SQL objects](https://github.com/chobostar/calculate_markup/blob/c70a994c8a215e11d3fd444f0d1412289bcadf3b/calculate_markup--1.0.sql)
- [C implementation](https://github.com/chobostar/calculate_markup/blob/c70a994c8a215e11d3fd444f0d1412289bcadf3b/calculate_markup.c)

`calculate_markup` installs the strict, immutable C function `price_modifiers._calculate_markup(numeric, numeric[])`. It selects or linearly interpolates a markup from a two-dimensional array of `(price, markup)` rows.

### Core Workflow

Pass rows in descending price order. Prices above the first breakpoint use its markup, prices below the last breakpoint use the last markup, and values between breakpoints are interpolated and rounded to two decimal places.

```sql
CREATE EXTENSION calculate_markup;

SELECT price_modifiers._calculate_markup(
  75::numeric,
  ARRAY[[100, 10], [50, 20], [0, 30]]::numeric[]
);
-- 15.00
```

The real SQL name begins with an underscore. Two README examples use `price_modifiers.c_calculate_markup`, but version 1.0 does not create that function.

### Constraints and Safety

The implementation checks only that the array has two dimensions. It assumes exactly two values per row, descending and distinct price breakpoints, at least one row, and no NULL elements. It does not correctly inspect PostgreSQL's NULL bitmap, and malformed, empty, odd-shaped, duplicate, or unsorted arrays can produce wrong results, division by zero, out-of-bounds access, or a backend crash. Validate the complete array shape and contents before calling it and test boundary prices explicitly.

This archived project has no current PostgreSQL compatibility statement. Treat it as unsafe native code, isolate evaluation in a disposable database, and prefer a small SQL implementation with explicit constraints for production pricing rules.
