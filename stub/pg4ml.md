

## Usage

> [pg4ml](https://gitee.com/guotiecheng/plpgsql_pg4ml): Machine learning framework for PostgreSQL.
> Source: [README.md](https://gitee.com/guotiecheng/plpgsql_pg4ml)

`pg4ml` is a PostgreSQL extension that implements a machine learning framework entirely within the database using PL/pgSQL and PL/Python. It provides matrix operations, neural network construction and training, clustering algorithms, and scientific computing -- all through SQL.


--------

### Prerequisites

- PostgreSQL >= 14 with Python3 support
- Required extensions: `plpgsql`, `tablefunc`, `cube`, `plpython3u`

### Getting Started

```sql
CREATE EXTENSION pg4ml CASCADE;
-- This will also create the required dependencies: plpgsql, tablefunc, cube, plpython3u
```


--------

### Features

#### Matrix Operations

The framework provides a comprehensive matrix operation library under the `sm_sc` schema:

- **Element-wise operations**: arithmetic, comparison, rounding, concatenation, boolean, bitwise, complex number, and broadcast operations
- **Matrix operations**: multiplication, transpose, flip, rotate, concatenation
- **Construction**: sampling, replacement, padding, character matching, random generation
- **Trigonometric functions**: broadcast operations on matrices
- **Aggregation**: slice-level aggregation, matrix-level aggregation, sorting by slice values, locating extremum positions

#### Slice Aggregation Examples

Average over vertical slices (groups of 2):

```sql
SELECT sm_sc.fv_aggr_slice_avg(
    array[[1.5, 11.5],
          [2.1, 12.1],
          [3.3, 13.3],
          [4.3, 14.3],
          [5.5, 15.5],
          [6.1, 16.1]],
    array[2, 1]
);
-- Returns: array[[1.8, 11.8],[3.8, 13.8],[5.8, 15.8]]
```

Max pooling over 2x3 blocks:

```sql
SELECT sm_sc.fv_aggr_slice_max(
    array[[2.3, 5.1, 8.2, 2.56, 3.33, -1.9],
          [3.25, 6.4, 6.6, 6.9, -2.65, -4.6],
          [-2.3, 5.1, -8.2, 2.56, -3.33, -1.9],
          [3.25, -6.4, -6.6, 6.9, -2.65, -4.6]],
    array[2, 3]
);
-- Returns: array[[8.2, 6.9],[5.1, 6.9]]
```

#### Neural Networks

The framework supports deep neural network construction and training:

- **Node and Path tables**: `sm_sc.tb_nn_node` / `sm_sc.tb_nn_path` for defining network structure
- **Training input buffer**: `sm_sc.tb_nn_train_input_buff` for receiving training data
- **Task management**: `sm_sc.tb_classify_task` for deploying and managing training tasks
- **Activation functions**, **convolution**, **pooling**, **lambda operations**
- **Loss functions**, **derivative computation**, **backpropagation**
- **Inference**: `sm_sc.ft_nn_in_out` for running test/validation data through a trained model

#### Clustering

- **K-means++**: via `sm_sc.prc_kmeans_pp` procedure
- **DBSCAN**: via `sm_sc.prc_dbscan_pp` procedure

Both use `sm_sc.tb_cluster_task` for task deployment and management.

#### Scientific Computing

- Waveform processing
- Computational graph JSON serialization/deserialization
- Complex number operations
- Linear algebra


--------

### Performance Tips

- Enable debug mode with: `SET session pg4ml._v_is_debug_check = '1';`
- Matrix multiplication uses `plpython3u` to call numpy for optimization
- Adjust PostgreSQL parallel parameters for multi-threaded training:
  - `max_parallel_workers_per_gather`
  - `force_parallel_mode`
  - `parallel_setup_cost`, `parallel_tuple_cost`
- Consider using `pg_strom` extension for GPU acceleration
