


## Usage

Sources: [README](https://github.com/ringsaturn/pg-tzf/blob/v0.3.0/README.md), [v0.3.0 release](https://github.com/ringsaturn/pg-tzf/releases/tag/v0.3.0)

`tzf` is a PostgreSQL extension for fast timezone lookup from longitude and latitude coordinates. The pgext catalog maps package `pg_tzf` to extension `tzf` and tracks version `0.3.0` for PostgreSQL 14-18.

### Create the extension

```sql
CREATE EXTENSION tzf;
```

The upstream project packages one build artifact per PostgreSQL major version.

### Functions

Coordinate lookup:

```sql
SELECT tzf_tzname(116.3883, 39.9289) AS timezone;
```

Batch coordinate lookup:

```sql
SELECT unnest(
  tzf_tzname_batch(
    ARRAY[-74.0060, -118.2437, 139.6917],
    ARRAY[40.7128, 34.0522, 35.6895]
  )
) AS timezones;
```

Point lookup:

```sql
SELECT tzf_tzname_point(point(-74.0060, 40.7128)) AS timezone;
```

Batch point lookup:

```sql
SELECT unnest(
  tzf_tzname_batch_points(
    ARRAY[
      point(-74.0060, 40.7128),
      point(-118.2437, 34.0522),
      point(139.6917, 35.6895)
    ]
  )
) AS timezones;
```

### Notes

- Upstream README documents support for PostgreSQL 14 through 18 builds.
- Pre-built release tarballs contain `tzf.so`, `tzf.control`, and `tzf--<version>.sql`.
- The current README still points to a complete schema in `sql/tzf.sql` and includes benchmark figures for the four lookup functions above.
- Upstream v0.3.0 adds city-json powered tests for timezone lookup queries and dependency updates; the SQL API remains the four lookup functions above.
