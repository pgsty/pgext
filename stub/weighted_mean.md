## Usage

Sources:

- [Upstream usage documentation](https://github.com/Kozea/weighted_mean/blob/4bf67038b72acadf5e83e7d58f9fff26ce202eb9/doc/weighted_mean.md)
- [C aggregate implementation](https://github.com/Kozea/weighted_mean/blob/4bf67038b72acadf5e83e7d58f9fff26ce202eb9/src/weighted_mean.c)
- [SQL aggregate declaration](https://github.com/Kozea/weighted_mean/blob/4bf67038b72acadf5e83e7d58f9fff26ce202eb9/sql/weighted_mean.sql)

`weighted_mean` provides a two-argument aggregate over `numeric` values and weights. It computes the sum of value times weight divided by the sum of weights, returning zero when no rows are processed or the total weight is zero.

```sql
CREATE EXTENSION weighted_mean;
SELECT weighted_mean(unitprice, quantity)
FROM sales
WHERE unitprice IS NOT NULL AND quantity IS NOT NULL;
```

The transition function is not declared strict yet dereferences both inputs, so filter NULL values explicitly; malformed NULL calls may crash a backend. Negative weights and a zero total can also make the result misleading. The repository is archived and the aggregate lacks modern parallel-state support; regression-test exact numeric behavior and prefer an explicit maintained SQL expression when practical.
