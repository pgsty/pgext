

## Usage

> [xicor: XI correlation coefficient for PostgreSQL](https://github.com/Florents-Tselai/pgxicor)

Provides the XI (xi) correlation coefficient, which can detect functional relationships between X and Y -- a more powerful alternative to `corr(X, Y)` that works beyond linear relationships.

```sql
CREATE EXTENSION xicor;
```

### Functions

| Function | Description |
|---|---|
| `xicor(x, y)` | Compute the XI correlation coefficient between two variables |

### Configuration

For reproducible results with tied data:

```sql
SET xicor.ties = true;
SET xicor.seed = 42;
```

### Examples

```sql
CREATE TABLE xicor_test (x float8, y float8);
INSERT INTO xicor_test (x, y) VALUES
  (1.0, 2.0), (2.5, 3.5), (3.0, 4.0), (4.5, 5.5), (5.0, 6.0);

SELECT xicor(x, y) FROM xicor_test;
```
