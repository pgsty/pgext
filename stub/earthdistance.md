

## Usage

> [earthdistance: Great circle distances on the surface of the Earth](https://www.postgresql.org/docs/current/earthdistance.html)

The `earthdistance` module provides two different approaches to calculating great circle distances on the surface of the Earth. The Earth is assumed to be perfectly spherical (for more accuracy, use PostGIS).

```sql
CREATE EXTENSION earthdistance CASCADE; -- requires cube
```

## Cube-Based Earth Distances

Data is stored in cubes using 3 coordinates representing x, y, z distance from Earth's center. A `earth` domain over type `cube` is provided.

### Functions

| Function | Description |
|----------|-------------|
| `earth()` → `float8` | Returns the assumed radius of the Earth |
| `sec_to_gc(float8)` → `float8` | Converts straight line (secant) distance to great circle distance |
| `gc_to_sec(float8)` → `float8` | Converts great circle distance to straight line (secant) distance |
| `ll_to_earth(float8, float8)` → `earth` | Returns location given latitude and longitude in degrees |
| `latitude(earth)` → `float8` | Returns latitude in degrees |
| `longitude(earth)` → `float8` | Returns longitude in degrees |
| `earth_distance(earth, earth)` → `float8` | Returns great circle distance between two points |
| `earth_box(earth, float8)` → `cube` | Returns a box for indexed search using cube `@>` operator |

### Example

```sql
-- Distance between New York and London in meters
SELECT earth_distance(
  ll_to_earth(40.7128, -74.0060),
  ll_to_earth(51.5074, -0.1278)
);

-- Find all points within 1000 meters of a location (indexable)
SELECT *
FROM places
WHERE earth_box(ll_to_earth(40.7128, -74.0060), 1000) @> ll_to_earth(lat, lon);
```


## Point-Based Earth Distances

Represents Earth locations as `point` type values where the first component is longitude and the second is latitude (in degrees).

### Operators

| Operator | Description |
|----------|-------------|
| `point <@> point` → `float8` | Distance in statute miles between two points |

### Example

```sql
-- Distance in statute miles
SELECT point(-74.0060, 40.7128) <@> point(-0.1278, 51.5074);
```

Note: Units are hardwired to statute miles. The point-based approach has edge condition issues near poles and ±180° longitude; the cube-based representation avoids these discontinuities.
