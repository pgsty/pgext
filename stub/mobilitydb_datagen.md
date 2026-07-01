


## Usage

Sources: [repository](https://github.com/MobilityDB/MobilityDB), [synthetic data generator docs](https://docs.mobilitydb.com/MobilityDB/develop/apb.html), [control file](https://github.com/MobilityDB/MobilityDB/blob/master/mobilitydb/datagen/mobilitydb_datagen.in.control), [temporal generators](https://github.com/MobilityDB/MobilityDB/blob/master/mobilitydb/datagen/temporal/random_temporal.sql), [temporal point generators](https://github.com/MobilityDB/MobilityDB/blob/master/mobilitydb/datagen/geo/random_tpoint.sql)

`mobilitydb_datagen` provides PL/pgSQL functions for generating synthetic PostgreSQL, PostGIS, and MobilityDB values. It is mainly useful for regression data, demos, and benchmark fixtures that need random temporal values or trajectories.

```sql
-- After the main MobilityDB extension is loaded:
CREATE EXTENSION mobilitydb_datagen;
```

### Generating Random Temporal Values

```sql
-- A random temporal float sequence.
SELECT random_tfloat_seq(
    -100.0, 100.0,                                  -- value bounds
    '2025-06-01 00:00+00', '2025-06-02 00:00+00',  -- time bounds
    10.0,                                           -- max value delta
    10,                                             -- max minutes between instants
    5, 10                                           -- min/max instants
);

-- Step interpolation instead of the default linear interpolation.
SELECT random_tfloat_seq(
    -100.0, 100.0,
    '2025-06-01 00:00+00', '2025-06-02 00:00+00',
    10.0, 10, 5, 10,
    false
);

-- A random temporal geometry point sequence.
SELECT asEWKT(
    random_tgeompoint_contseq(
        2.20, 2.50, 48.80, 48.95,                  -- x/y bounds
        '2025-06-01 08:00+00', '2025-06-01 18:00+00',
        0.02, 5, 20, 40,                           -- max delta, max minutes, min/max instants
        srid => 4326
    )
);
```

Other confirmed generator families include scalar helpers such as `random_bool`, `random_int`, `random_float`, `random_text`, and `random_timestamptz`; array, set, span, and range helpers; temporal helpers such as `random_tbool_inst`, `random_tint_discseq`, `random_tfloat_seq`, and `random_tfloat_seqset`; and spatial/temporal-point helpers such as `random_geom_point`, `random_geom_linestring`, `random_tgeompoint_contseq`, `random_tgeompoint_seqset`, `random_tgeogpoint_contseq`, and `random_tgeogpoint_seqset`.

### Generating Test Datasets

Create bulk test data for benchmarking trip queries:

```sql
CREATE TABLE trip_samples AS
SELECT
    vehicle_id,
    random_tgeompoint_contseq(
        2.20, 2.50, 48.80, 48.95,
        '2025-06-01 08:00+00', '2025-06-01 18:00+00',
        0.02, 5, 20, 40,
        srid => 4326
    ) AS trip
FROM generate_series(1, 1000) AS vehicle_id;
```

### Caveats

- The control file requires the main `mobilitydb` extension; `mobilitydb_datagen` is not standalone.
- The package row in `db/extension.csv` lists version `1.3.0`, package `mobilitydb`, and PostgreSQL support for 14 through 18.
- Upstream docs intentionally omit detailed parameter lists for many generator functions and point users to the SQL source files for exact signatures.
