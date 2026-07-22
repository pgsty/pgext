## Usage

Sources:

- [Official usage documentation](https://github.com/umitanuki/kmeans-postgresql/blob/d6a57eda9e4908ef5768ad2a5c70bd92d0abb3ae/doc/kmeans.md)
- [Extension SQL](https://github.com/umitanuki/kmeans-postgresql/blob/d6a57eda9e4908ef5768ad2a5c70bd92d0abb3ae/kmeans.sql.in)
- [C implementation](https://github.com/umitanuki/kmeans-postgresql/blob/d6a57eda9e4908ef5768ad2a5c70bd92d0abb3ae/kmeans.c)

`kmeans` version `1.1.0` implements K-means clustering as a PostgreSQL window function. It returns a zero-based cluster number for every input row and is intended for relatively small, in-database vector classifications.

### Core Workflow

```sql
CREATE EXTENSION kmeans;

SELECT id,
       kmeans(ARRAY[x, y, z]::float8[], 3) OVER () AS cluster
FROM samples;
```

All vectors in a window partition must be one-dimensional and have the same length. The second argument is the requested number of clusters.

For reproducible initialization, supply centroids explicitly:

```sql
SELECT id,
       kmeans(
         ARRAY[x, y]::float8[],
         2,
         ARRAY[0.5, 0.5, 1.0, 1.0]::float8[]
       ) OVER (PARTITION BY group_key) AS cluster
FROM samples;
```

The third argument may be a two-dimensional centroid array or a flat array whose length equals cluster count times vector dimension.

### Behavior and Caveats

- `kmeans(float[], int) returns int`: choose initial centroids from input vectors after interpolating between per-dimension minima and maxima.
- `kmeans(float[], int, float[]) returns int`: use caller-supplied initial centroids.

K-means results depend on initialization and scale. Normalize dimensions when appropriate, supply fixed centroids when repeatability matters, and test NULLs and degenerate partitions. This upstream release dates from 2011 and was designed for PostgreSQL 8.4-era APIs; verify source compatibility and results on the target PostgreSQL version. No preload is required.
