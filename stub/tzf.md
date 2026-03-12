

## Usage

> [tzf: Fast timezone lookup for PostgreSQL](https://github.com/ringsaturn/pg-tzf)

Find timezone names from coordinates. Built on [tzf-rs](https://github.com/ringsaturn/tzf-rs) with timezone boundary data from [timezone-boundary-builder](https://github.com/evansiroky/timezone-boundary-builder).

```sql
CREATE EXTENSION tzf;
```

### Functions

Look up timezone for a coordinate (longitude, latitude):

```sql
SELECT tzf_tzname(116.3883, 39.9289) AS timezone;
-- Asia/Shanghai
```

Look up timezone for a batch of coordinates:

```sql
SELECT unnest(
  tzf_tzname_batch(
    ARRAY[-74.0060, -118.2437, 139.6917],
    ARRAY[40.7128, 34.0522, 35.6895]
  )
) AS timezones;
-- America/New_York
-- America/Los_Angeles
-- Asia/Tokyo
```

Look up timezone for a point:

```sql
SELECT tzf_tzname_point(point(-74.0060, 40.7128)) AS timezone;
-- America/New_York
```

Look up timezone for a batch of points:

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
-- America/New_York
-- America/Los_Angeles
-- Asia/Tokyo
```

### Performance

| Function | TPS | Note |
|----------|-----|------|
| `tzf_tzname` | ~17,700 | Single coordinate lookup |
| `tzf_tzname_point` | ~17,600 | Single point lookup |
| `tzf_tzname_batch` | ~51 | Batch of 1000 = ~51,000 TPS |
| `tzf_tzname_batch_points` | ~32 | Batch of 1000 = ~32,000 TPS |
