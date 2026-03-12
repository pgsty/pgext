

## Usage

> [mobilitydb_datagen: Synthetic mobility data generator for MobilityDB](https://github.com/MobilityDB/MobilityDB)

MobilityDB DataGen provides functions for generating synthetic mobility data for testing and benchmarking MobilityDB workloads. It creates random temporal values including trips, trajectories, and time-varying measurements.

### Generating Random Temporal Values

```sql
-- Generate a random temporal float over a time span
SELECT random_tfloat(
    '2025-06-01 00:00+00', '2025-06-02 00:00+00',  -- time span
    0.0, 100.0,                                      -- value range
    10                                               -- number of instants
);

-- Generate a random temporal geometric point (trajectory)
SELECT random_tgeompoint(
    '2025-06-01 08:00+00', '2025-06-01 18:00+00',   -- time span
    ST_MakeEnvelope(2.2, 48.8, 2.4, 48.9, 4326),    -- spatial bounds
    20                                               -- number of instants
);
```

### Generating Test Datasets

Create bulk test data for benchmarking trip queries:

```sql
INSERT INTO trips (vehicle_id, trip, trip_date)
SELECT
    i,
    random_tgeompoint(
        '2025-06-01 08:00+00', '2025-06-01 18:00+00',
        ST_MakeEnvelope(2.2, 48.8, 2.5, 48.9, 4326),
        50
    ),
    '2025-06-01'
FROM generate_series(1, 1000) AS i;
```
