## Usage

Sources:

- [Extension control file](https://github.com/idrawone/get_sum/blob/43f8ed114da79321c4125c190044ab4e06e1322e/get_sum.control)
- [Version 0.0.1 extension SQL](https://github.com/idrawone/get_sum/blob/43f8ed114da79321c4125c190044ab4e06e1322e/get_sum--0.0.1.sql)
- [C implementation](https://github.com/idrawone/get_sum/blob/43f8ed114da79321c4125c190044ab4e06e1322e/get_sum.c)
- [Upstream regression expectations](https://github.com/idrawone/get_sum/blob/43f8ed114da79321c4125c190044ab4e06e1322e/expected/get_sum_test.out)

`get_sum` 0.0.1 is a minimal C-extension example that adds two PostgreSQL `integer` values. It is educational sample code, not a safer or more capable replacement for the built-in `+` operator.

### Core Workflow

```sql
CREATE EXTENSION get_sum;

SELECT get_sum(101, 202);
SELECT get_sum(NULL, 11);
```

The first call returns `303`. The SQL declaration is `STRICT`, so a NULL argument returns NULL without entering the C function.

### Interface

`get_sum(integer, integer) RETURNS integer` is declared `IMMUTABLE STRICT`. It accepts exactly two `int4` values and returns their C `int32` sum. There are no `bigint`, `numeric`, aggregate, variadic, or checked-arithmetic overloads.

### Overflow Caveat

The implementation performs plain signed 32-bit C addition and does not call PostgreSQL's checked integer arithmetic. The upstream regression output records `get_sum(2147483647, 2147483647)` returning `-2` rather than raising `integer out of range`. C signed overflow is not a stable application contract and can vary with compiler assumptions.

Use PostgreSQL's built-in arithmetic for normal queries. If this sample is exposed at all, keep it to demonstrations with inputs proven to remain within the `integer` range.
