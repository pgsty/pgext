## Usage

Sources:

- [argm 1.1.1 README](https://github.com/bashtanov/argm/blob/1.1.1/README.md)
- [Extension control file](https://github.com/bashtanov/argm/blob/1.1.1/argm.control)
- [SQL definitions](https://github.com/bashtanov/argm/blob/1.1.1/argm--1.1.1.sql)

`argm` provides the polymorphic aggregates `argmax`, `argmin`, and `anyold`. They return a value from a selected row while grouping, avoiding a join or window-function pass when the row can be chosen by one or more sortable keys.

### Core Workflow

```sql
CREATE EXTENSION argm;

SELECT customer_id,
       argmax(order_id, total, ordered_at) AS largest_order
FROM orders
GROUP BY customer_id;
```

`argmax(value, key...)` returns the `value` belonging to the lexicographically greatest key tuple. `argmin` selects the least tuple. Additional keys break ties without building a composite value:

```sql
SELECT device_id,
       argmax(reading, measured_at, sequence_no) AS latest_reading
FROM measurements
GROUP BY device_id;
```

Use `anyold(value)` when any member of a group is acceptable:

```sql
SELECT account_id, anyold(display_name)
FROM account_aliases
GROUP BY account_id;
```

### Important Objects

- `argmax(value, key [, key ...])` selects the value associated with the greatest key tuple.
- `argmin(value, key [, key ...])` selects the value associated with the least key tuple.
- `anyold(value)` returns an arbitrary non-null value from the aggregate state.

The aggregates accept any value type; key types must have ordering support. The SQL definitions are parallel-safe and include combine and serialization functions for partial aggregation.

### Semantics and Caveats

Key tuples use one ordering direction and one collation for the whole tuple, with null keys sorted last. If complete key tuples tie, the chosen value is unspecified; add a stable final key when deterministic results matter. As with other PostgreSQL aggregates, empty input produces `NULL`.

`argm` 1.1.x requires PostgreSQL 9.6 or newer. The extension is relocatable. Upgrading from 1.0.3 to 1.1.x requires dropping and recreating the extension because the aggregate state changed; the 1.1.0-to-1.1.1 upgrade does not change the public SQL surface.
