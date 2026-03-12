


## Usage

> [floatfile: Float array file I/O for PostgreSQL](https://github.com/pjungwir/floatfile)

Store float arrays in individual files for very fast querying without MVCC overhead. Ideal for time series data requiring low-latency reads and low-cost appends.

### Functions

#### Save a float array to a file

```sql
SELECT save_floatfile('my_data', ARRAY[1.0, 2.0, 3.0, 4.0]::float[]);
```

#### Load a float array from a file

```sql
SELECT load_floatfile('my_data');
-- {1,2,3,4}
```

#### Append values to an existing file

```sql
SELECT extend_floatfile('my_data', ARRAY[5.0, 6.0]::float[]);
```

#### Delete a floatfile

```sql
SELECT drop_floatfile('my_data');
```

### Tablespace Variants

All functions accept an optional tablespace name as first argument:

```sql
SELECT save_floatfile('my_ts', 'my_data', ARRAY[1.0, 2.0]::float[]);
SELECT load_floatfile('my_ts', 'my_data');
SELECT extend_floatfile('my_ts', 'my_data', ARRAY[3.0]::float[]);
SELECT drop_floatfile('my_ts', 'my_data');
```

### Histogram Functions

Compute histograms directly from floatfiles (useful when arrays exceed the 1GB Postgres limit):

```sql
SELECT floatfile_to_hist('my_data', 0.0, 1.0, 10);
-- Returns int[] with histogram counts

SELECT floatfile_to_hist2d('xs_file', 'ys_file', 0.0, 0.0, 1.0, 1.0, 10, 10);
-- Returns 2D int[] histogram
```

### Concurrency

Functions use PostgreSQL advisory locks: `load_floatfile` takes a shared lock; `save`, `extend`, and `drop` take exclusive locks.
