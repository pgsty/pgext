

## Usage

> [vasco: Maximal Information Coefficient (MIC) extension for PostgreSQL](https://github.com/Florents-Tselai/vasco)

Discover hidden correlations in your data using the Maximal Information Coefficient (MIC) and the MINE family of statistics.

```sql
CREATE EXTENSION vasco;
```

### Aggregate Functions

| Function | Description |
|---|---|
| `mic(x, y)` | Maximal Information Coefficient -- detects any relationship |
| `mas(x, y)` | Maximum Asymmetry Score -- deviation from monotonicity |
| `mev(x, y)` | Maximum Edge Value -- degree of continuous function sampling |
| `mcn(x, y)` | Minimum Cell Number -- complexity of association |
| `mcn_general(x, y)` | MCN with `eps = 1 - MIC` |
| `tic(x, y)` | Total Information Coefficient |
| `gmic(x, y)` | Generalized Mean Information Coefficient |

### Utility Functions

| Function | Description |
|---|---|
| `vasco_corr_matrix(table_name, output_table)` | Compute MIC for all column pairs and store as a correlation matrix table |

### Configuration

```sql
SET vasco.mic_estimator = 'ApproxMIC';  -- or 'MIC_e'
SET vasco.mine_c = ...;
SET vasco.mine_alpha = ...;
```

### Examples

```sql
-- Compute MIC between column pairs
SELECT mic(x, cubic), mic(x, periodic), mic(x, rand_y)
FROM vasco_data;
-- 1, 1, 0.15

-- Create a full correlation matrix
SELECT vasco_corr_matrix('my_table', 'mic_my_table');
SELECT * FROM mic_my_table;
```
