## Usage

Sources:

- [Alibaba Cloud RDS varbitx documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/use-the-varbitx-plug-in)

`varbitx` version `1.0` is an Alibaba Cloud ApsaraDB RDS for PostgreSQL extension that adds slicing, bulk mutation, position lookup, generation, and counting functions for variable-length bit strings. It is a provider-managed feature, not a documented portable community package.

### Core Workflow

Alibaba Cloud documents the extension for supported RDS PostgreSQL 10 and 11 instances.

```sql
CREATE EXTENSION varbitx;

SELECT get_bit('111110000011'::varbit, 3, 5);
SELECT bit_count('1111000011110000'::varbit, 1, 5, 4);
SELECT set_bit_array_record(
  '111100001111'::varbit,
  1,
  0,
  ARRAY[1,4,5,6,7],
  2
);
```

Positions in the documented functions are zero-based. `get_bit` extracts a slice. `set_bit_array` returns a modified bit string; `set_bit_array_record` also returns changed positions. Optional limits stop after a specified number of changes.

### Function Groups and Caveats

Extraction uses `get_bit` and `get_bit_array`. Generation uses `bit_fill` and probabilistic `bit_rand`. Counting uses `bit_count` and `bit_count_array`. `bit_posite` returns matching positions in ascending or descending order.

Mutation functions return new values rather than updating a table in place. Their fill, result-length, ordering, and limit arguments materially change results, so test boundary positions and strings shorter than the requested indexes. Availability and upgrades follow the RDS instance family and provider release; verify the exact database version and function signatures on the target instance before storing procedures that depend on them.
