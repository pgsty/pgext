## Usage

Sources:

- [Pinned upstream README](https://github.com/geo-mathijs/gbuilder/blob/9e0df5e15885949581b43bf931abd6f4fefff2f1/README.md)
- [Version 1.0 installation SQL](https://github.com/geo-mathijs/gbuilder/blob/9e0df5e15885949581b43bf931abd6f4fefff2f1/gbuilder--1.0.sql)
- [Pinned C entry points](https://github.com/geo-mathijs/gbuilder/blob/9e0df5e15885949581b43bf931abd6f4fefff2f1/gbuilder.c)
- [Pinned hull implementation](https://github.com/geo-mathijs/gbuilder/blob/9e0df5e15885949581b43bf931abd6f4fefff2f1/hull.c)
- [Pinned extension control file](https://github.com/geo-mathijs/gbuilder/blob/9e0df5e15885949581b43bf931abd6f4fefff2f1/gbuilder.control)

gbuilder version 1.0 is a C extension that adds one geometry builder: a Delaunay-based concave hull for point inputs. It depends on PostGIS and the Fast Robust Predicates library. The aggregate uses alpha 1.0; the array form accepts a custom alpha.

### Aggregate and array forms

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION gbuilder;

SELECT ConcaveHull(geom)
FROM demo_points;

SELECT ConcaveHull(
  ARRAY[
    ST_SetSRID(ST_Point(0, 0), 4326),
    ST_SetSRID(ST_Point(1, 0), 4326),
    ST_SetSRID(ST_Point(1, 1), 4326),
    ST_SetSRID(ST_Point(0, 1), 4326)
  ],
  1.0
);
```

The aggregate declaration reuses PostGIS's geometry accumulator and always passes alpha 1.0 in its final function. To tune alpha, aggregate points into a geometry array explicitly and call the two-argument function. The implementation silently skips NULL values and every geometry that is not a POINT; no usable point or no surviving polygon produces NULL.

### Geometry and ABI limits

The C code initializes the output SRID to 4326 rather than deriving it from the inputs. Mixed SRIDs, non-4326 input, Z or M dimensions, duplicate points, collinear points, zero X range, very small inputs, and extreme alpha values therefore require explicit validation. Normalize and validate points before calling the function, then check the output with PostGIS validity and SRID functions.

gbuilder compiles against private liblwgeom structures and recreates an internal PostGIS aggregate state structure that is not a public ABI. It also calls PostgreSQL and PostGIS internal symbols directly. A binary built for one PostgreSQL/PostGIS combination can fail to link, crash, or corrupt memory with another. Build and regression-test against the exact PostgreSQL and PostGIS packages deployed on every node.

The upstream performance and precision comparison with PostGIS has no reproducible benchmark in the repository, and the last reviewed code is from 2019. Compare it with the current PostGIS hull functions on representative geometry, including invalid and adversarial inputs, before choosing it. Treat the extension as legacy native code and isolate first trials on a disposable database.
