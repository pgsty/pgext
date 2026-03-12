

## Usage

> [pg_timeseries: Convenience API for time series stack](https://github.com/ChuckHend/pg_timeseries)

This extension provides a cohesive user experience around the creation, maintenance, and use of time-series tables.

### Getting Started

Assuming you already have a partitioned table created, simply call the `enable_ts_table` function with your table name.

```sql
CREATE EXTENSION timeseries CASCADE;

SELECT enable_ts_table('sensor_readings');
```

With this one call, several things will happen:

  * The table will be restructured as a series of partitions using PostgreSQL's [native PARTITION features](https://www.postgresql.org/docs/current/ddl-partitioning.html)
  * Each partition covers a particular range of time (one week by default)
  * New partitions will be created for some time in the future (one month by default)
  * Once an hour, a maintenance job will create any missing partitions as well as needed future ones


## Using Your Tables

### Indexes

The time-series tables you create start out life as little more than typical [partitioned PostgreSQL tables](https://www.postgresql.org/docs/current/ddl-partitioning.html). All of PostgreSQL's existing functionality will "just work" with them.

[Traditional B-Tree indexes](https://www.postgresql.org/docs/current/btree-intro.html) work well for time-series data, but you may wish to benchmark [BRIN indexes](https://www.postgresql.org/docs/current/brin-intro.html) as well, as they may perform better in specific query scenarios (often queries with _many_ results). Start with B-Tree if you don't anticipate more than a million records in each partition (by default, partitions are one week long).

### Partition Sizing

Because calculating the total size of partitioned tables can be tedious, this extension provides several easy-to-use views surfacing this information.

To examine the table (data), index, and total size for each of your partitions, query the time-series partition information view, `ts_part_info`. A general rule of thumb is that each partition should be able to fit within roughly one quarter of your available memory.

### Retention

Call `set_ts_retention_policy` with your time-series table and an interval (say, `'90 days'`) to establish a retention policy. Once an hour, any partitions falling entirely outside the retention window will be dropped. Use `clear_ts_retention_policy` to revert to the default behavior (infinite retention). Each of these functions will return the previous retention policy when called.

### Compression

By calling `set_ts_compression_policy` on a time-series table with an appropriate interval (perhaps `'1 month'`), this extension will compress partitions (using a columnar storage method) older than the specified interval, once an hour. A function is also provided for clearing any existing policy (existing partitions will not be decompressed, however).

The compression features depend on the [citus](https://github.com/citusdata/citus/tree/main) and [citus_columnar](https://github.com/citusdata/citus/tree/main/src/backend/columnar) extensions:

```sql
CREATE EXTENSION citus;
CREATE EXTENSION citus_columnar;
```


## Analytics Helpers

### `first` and `last`

These two functions help clean up the syntax of a fairly common pattern: a query is grouped by one dimension, but a user wants to know what the first or last row in a group is when ordered by a _different_ dimension.

```sql
SELECT machine_id,
       last(cpu_util, recorded_at)
FROM events
GROUP BY machine_id;
```

### `date_bin_table`

This function automates the tedium of aligning time-series values to a given width, or "stride", and makes sure to include NULL rows for any time periods where the source table has no data points.

```sql
SELECT * FROM date_bin_table(NULL::target_table, '1 hour', '[2024-02-01 00:00, 2024-02-02 15:00]');
```

The output of this query will differ from simply hitting the target table directly in three ways:

  * Rows will be sorted by time, ascending
  * The time column's values will be binned to the provided width
  * Extra rows will be added for periods with no data. They will include the time stamp for that bin and NULL in all other columns

### `make_view_incremental`

This function accepts a view and converts it into a materialized view which is kept up-to-date after every modification. This removes the need for users to pick between always up-to-date `VIEW`s and having to call `REFRESH` on `MATERIALIZED VIEW`s.

The underlying functionality is provided by [a fork](https://github.com/ChuckHend/pg_ivm) of [`pg_ivm`](https://github.com/sraoss/pg_ivm). Enable the `pg_ivm` extension if you want to use this feature:

```sql
CREATE EXTENSION pg_ivm;
```


## Requirements

* [pg_cron](https://github.com/citusdata/pg_cron)
* [pg_partman](https://github.com/pgpartman/pg_partman)

### Optional Dependencies

* [pg_ivm](https://github.com/sraoss/pg_ivm) — for incremental materialized views
* [Citus & Citus Columnar](https://github.com/citusdata/citus) — for compression features
