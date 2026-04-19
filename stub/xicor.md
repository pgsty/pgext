## Usage

Sources: [README](https://github.com/Florents-Tselai/pgxicor/blob/main/README.md), [release 0.1.1](https://github.com/Florents-Tselai/pgxicor/releases/tag/v0.1.1)

`xicor` exposes the XI (Chatterjee's xi) correlation coefficient as a PostgreSQL aggregate. It is meant for detecting functional dependence, including non-linear relationships that Pearson's `corr()` can miss.

```sql
CREATE EXTENSION xicor;
```

### Main Aggregate

```sql
SELECT xicor(x, y) FROM xicor_test;
```

The upstream example contrasts it with `corr()` on a parabola-shaped dataset, where `corr()` is near zero while `xicor()` remains high.

### Example

```sql
CREATE TABLE xicor_test (x float8, y float8);
INSERT INTO xicor_test (x, y) VALUES
  (1.0, 2.0),
  (2.5, 3.5),
  (3.0, 4.0),
  (4.5, 5.5),
  (5.0, 6.0);

SELECT xicor(x, y) FROM xicor_test;
```

### Reproducibility Controls

For tied data, upstream recommends enabling deterministic tie handling:

```sql
SET xicor.ties = true;
SET xicor.seed = 42;
```

### Caveats

- `xicor()` is an aggregate over two numeric inputs, not a general-purpose statistical framework.
- Tie handling can change results unless you enable the documented GUCs for reproducible behavior.
