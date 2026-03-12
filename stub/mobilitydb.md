

## Usage

> [mobilitydb: Temporal and spatio-temporal data management for PostgreSQL](https://github.com/MobilityDB/MobilityDB)

MobilityDB extends PostgreSQL and PostGIS with temporal and spatio-temporal data types, enabling efficient storage, indexing, and querying of moving object data such as vehicle trajectories, sensor readings, and time-varying attributes.

**Key Documentation:**

- [MobilityDB Manual](https://docs.mobilitydb.com/MobilityDB/master/)
- [Temporal Types](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html#temporal-types)
- [Temporal Operations](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html#temporal-operations)
- [Spatial-Temporal Types](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html#spatial-temporal-types)
- [Indexing](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html#indexing)
- [MobilityDB Workshop](https://mobilitydb.com/documentation/)
- [API Reference](https://docs.mobilitydb.com/MobilityDB/master/mobilitydb-manual.html)

### Getting Started

MobilityDB requires PostGIS. Enable both extensions:

```sql
CREATE EXTENSION PostGIS;
CREATE EXTENSION MobilityDB;
```

### Temporal Types

MobilityDB provides temporal variants of base types:

| Temporal Type | Base Type | Description |
|---------------|-----------|-------------|
| `tbool`       | `boolean` | Time-varying boolean |
| `tint`        | `integer` | Time-varying integer |
| `tfloat`      | `float`   | Time-varying float |
| `ttext`       | `text`    | Time-varying text |
| `tgeompoint`  | `geometry(Point)` | Time-varying geometric point |
| `tgeogpoint`  | `geography(Point)` | Time-varying geographic point |

### Temporal Subtypes

Each temporal type can be represented in different subtypes depending on how values change over time:

| Subtype | Description | Example |
|---------|-------------|---------|
| **Instant** | Single value at a single timestamp | `'25.5@2025-01-01 08:00'` |
| **Sequence** | Continuous values over a time interval | `'[25.5@08:00, 28.1@09:00, 30.0@10:00]'` |
| **SequenceSet** | Set of non-overlapping sequences | `'{[25.5@08:00, 28.1@09:00], [30.0@11:00, 31.2@12:00]}'` |

Sequences use brackets to indicate inclusive `[` or exclusive `(` bounds, just like PostgreSQL range types.

### Creating Temporal Values

**Instant values:**

```sql
SELECT tfloat '25.5@2025-06-01 08:00:00+00';
SELECT tgeompoint 'SRID=4326;Point(2.3522 48.8566)@2025-06-01 08:00:00+00';
```

**Sequence values (continuous interpolation):**

```sql
SELECT tfloat '[20.0@2025-06-01 08:00, 25.5@2025-06-01 09:00, 22.0@2025-06-01 10:00]';
```

**Discrete sequences (stepwise interpolation):**

```sql
SELECT tint 'Interp=Step;[10@2025-06-01 08:00, 20@2025-06-01 09:00, 15@2025-06-01 10:00]';
```

**SequenceSet values:**

```sql
SELECT tfloat '{[20.0@08:00, 25.5@09:00], [22.0@11:00, 28.0@12:00]}';
```

**Constructing from components:**

```sql
SELECT tgeompoint_inst(ST_Point(2.3522, 48.8566, 4326), '2025-06-01 08:00+00');
SELECT tgeompoint_seq(ARRAY[
    tgeompoint_inst(ST_Point(2.3522, 48.8566, 4326), '2025-06-01 08:00+00'),
    tgeompoint_inst(ST_Point(2.2945, 48.8584, 4326), '2025-06-01 08:30+00'),
    tgeompoint_inst(ST_Point(2.3364, 48.8606, 4326), '2025-06-01 09:00+00')
]);
```

### Temporal Operations

**Extracting values at a specific time:**

```sql
SELECT valueAtTimestamp(temp, '2025-06-01 08:30:00+00')
FROM (SELECT tfloat '[20.0@08:00, 30.0@09:00]' AS temp) t;
-- Returns 25.0 (linear interpolation)
```

**Restricting to a time period:**

```sql
SELECT atTime(trip, tstzspan '[2025-06-01 08:00, 2025-06-01 09:00]')
FROM trips;
```

**Getting the time span of a temporal value:**

```sql
SELECT duration(trip), startTimestamp(trip), endTimestamp(trip)
FROM trips;
```

**Temporal comparisons:**

```sql
-- Time periods when temperature exceeded 30 degrees
SELECT atValue(temperature, true)
FROM (SELECT tfloat '[20@08:00, 35@09:00, 25@10:00]' #> 30.0 AS temperature) t;
```

### Spatial-Temporal Operations

**Trajectory: extract the spatial path as a geometry:**

```sql
SELECT ST_AsText(trajectory(trip))
FROM trips
WHERE vehicle_id = 42;
```

**Speed calculation:**

```sql
-- Speed in units per second (m/s for geographic points)
SELECT speed(trip)
FROM trips
WHERE vehicle_id = 42;
```

**Length of trajectory:**

```sql
SELECT length(trip)
FROM trips
WHERE vehicle_id = 42;
```

**Space-time bounding box (stbox):**

```sql
-- Get the space-time bounding box
SELECT stbox(trip)
FROM trips;

-- Construct an stbox for querying
SELECT stbox(
    ST_MakeEnvelope(2.2, 48.8, 2.4, 48.9, 4326),
    tstzspan '[2025-06-01, 2025-06-02]'
);
```

**Spatial restriction: values within an area:**

```sql
-- Portions of a trip within a polygon
SELECT atGeometry(trip, ST_Buffer(ST_Point(2.35, 48.86, 4326), 0.01))
FROM trips;
```

**Distance between two temporal points:**

```sql
SELECT distance(t1.trip, t2.trip)
FROM trips t1, trips t2
WHERE t1.vehicle_id = 1 AND t2.vehicle_id = 2;
```

**Nearest approach distance and time:**

```sql
SELECT nearestApproachDistance(t1.trip, t2.trip),
       nearestApproachInstant(t1.trip, t2.trip)
FROM trips t1, trips t2
WHERE t1.vehicle_id = 1 AND t2.vehicle_id = 2;
```

### Indexing

MobilityDB supports GiST and SP-GiST indexes for efficient temporal and spatio-temporal queries.

**SP-GiST index for temporal types (time dimension):**

```sql
CREATE INDEX ON measurements USING spgist(temperature);
```

**GiST index for spatio-temporal types (space + time):**

```sql
CREATE INDEX ON trips USING gist(trip);
```

These indexes accelerate bounding box queries, temporal overlap checks, and spatial-temporal intersection:

```sql
-- Uses GiST index for space-time filtering
SELECT vehicle_id
FROM trips
WHERE trip && stbox(
    ST_MakeEnvelope(2.2, 48.8, 2.4, 48.9, 4326),
    tstzspan '[2025-06-01, 2025-06-02]'
);
```

### Example: Vehicle Tracking

A complete example storing and querying vehicle GPS trajectories:

```sql
CREATE TABLE vehicles (
    vehicle_id  INT PRIMARY KEY,
    plate       TEXT,
    type        TEXT
);

CREATE TABLE trips (
    trip_id     BIGSERIAL PRIMARY KEY,
    vehicle_id  INT REFERENCES vehicles(vehicle_id),
    trip        tgeompoint,
    trip_date   DATE
);

CREATE INDEX ON trips USING gist(trip);

-- Insert a trip as a sequence of GPS points
INSERT INTO trips (vehicle_id, trip, trip_date) VALUES (
    1,
    tgeompoint_seq(ARRAY[
        tgeompoint_inst(ST_Point(2.3522, 48.8566, 4326), '2025-06-01 08:00+00'),
        tgeompoint_inst(ST_Point(2.2945, 48.8584, 4326), '2025-06-01 08:15+00'),
        tgeompoint_inst(ST_Point(2.3364, 48.8606, 4326), '2025-06-01 08:30+00'),
        tgeompoint_inst(ST_Point(2.3488, 48.8534, 4326), '2025-06-01 08:45+00')
    ]),
    '2025-06-01'
);

-- Where was vehicle 1 at 08:20?
SELECT valueAtTimestamp(trip, '2025-06-01 08:20+00')
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- What was the average speed?
SELECT twAvg(speed(trip))
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- Total distance traveled
SELECT length(trip)
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- Get the full trajectory as a LineString
SELECT ST_AsGeoJSON(trajectory(trip))
FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';
```

### Example: Spatio-Temporal Intersection Query

Find all trips that passed through a specific area during a given time window:

```sql
-- Define area of interest: a circle around the Eiffel Tower
WITH area AS (
    SELECT ST_Buffer(ST_Point(2.2945, 48.8584, 4326)::geography, 500)::geometry AS geom
)
SELECT t.vehicle_id,
       t.trip_date,
       atGeometry(t.trip, a.geom) AS trip_in_area,
       length(atGeometry(t.trip, a.geom)) AS distance_in_area
FROM trips t, area a
WHERE t.trip && stbox(
    a.geom,
    tstzspan '[2025-06-01 07:00+00, 2025-06-01 10:00+00]'
)
  AND eIntersects(t.trip, a.geom)
ORDER BY t.trip_date;
```

### Aggregate Functions

MobilityDB provides temporal aggregates:

```sql
-- Time-weighted average of a temporal float
SELECT twAvg(temperature) FROM sensor_data WHERE sensor_id = 1;

-- Merge multiple temporal points into one
SELECT tUnion(trip) FROM trips WHERE vehicle_id = 1 AND trip_date = '2025-06-01';

-- Centroid of a set of temporal points at each timestamp
SELECT tCentroid(trip) FROM trips WHERE trip_date = '2025-06-01';
```
