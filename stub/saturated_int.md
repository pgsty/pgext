## Usage

Sources:

- [Official README](https://github.com/adjust/pg-saturated_int/blob/a1228ed87d0b27a365eb6d0b241a27f654f49156/README.md)
- [Extension SQL and object definitions](https://github.com/adjust/pg-saturated_int/blob/a1228ed87d0b27a365eb6d0b241a27f654f49156/saturated_int--0.0.1.sql)
- [Saturating arithmetic implementation](https://github.com/adjust/pg-saturated_int/blob/a1228ed87d0b27a365eb6d0b241a27f654f49156/src/saturated_int.c)

`saturated_int` version `0.0.1` supplies a signed 32-bit integer type whose input and arithmetic clamp overflow to `-2147483648` or `2147483647` instead of raising `integer out of range`.

### Core Workflow

```sql
CREATE EXTENSION saturated_int;

SELECT 999999999999999::saturated_int;
SELECT 2147483640::saturated_int + 100::saturated_int;
SELECT (-2147483648)::saturated_int / (-1)::saturated_int;
```

The results of these expressions are respectively `2147483647`, `2147483647`, and `2147483647`.

### Objects

- `saturated_int`: int4-sized custom type with a range of `-2147483648` through `2147483647`.
- Comparison operators: `=`, `<>`, `<`, `<=`, `>`, and `>=`.
- Arithmetic operators: `+`, `-`, `*`, `/`, and `%`.
- `sum(saturated_int)`: aggregate with a saturated accumulator.
- B-tree and hash operator classes for indexes, ordering, grouping, and equality lookup.

### Cast and Arithmetic Notes

Casts from `bigint` clamp to the type's range. Assignment casts exist between `integer` and `saturated_int`, but the extension does not define mixed-type operators: explicitly cast both operands when comparing or calculating with built-in integers.

Saturation is a data-model choice, not merely error handling: distinct out-of-range inputs collapse to the same boundary value, and aggregate overflow also clamps. Division by zero still raises an error. Use the type only when this loss of magnitude is intended. No preload or restart is required.
