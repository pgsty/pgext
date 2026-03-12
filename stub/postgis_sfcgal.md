

## Usage

> [PostGIS SFCGAL: 3D geometry and advanced operations powered by SFCGAL](https://github.com/postgis/postgis)

PostGIS SFCGAL provides advanced 2D and 3D spatial operations by wrapping the [SFCGAL](https://sfcgal.gitlab.io/SFCGAL/) library. It adds support for 3D geometry operations, volume calculations, extrusion, tesselation, and other functions not available in the core PostGIS GEOS backend.

- [SFCGAL Function Reference](https://postgis.net/docs/reference_sfcgal.html)

### Setup

```sql
CREATE EXTENSION postgis_sfcgal;
```

--------

## 3D Operations

### 3D Intersection and Difference

```sql
-- 3D intersection of two solids
SELECT ST_3DIntersection(
    ST_GeomFromText('POLYHEDRALSURFACE Z(((0 0 0,1 0 0,1 1 0,0 1 0,0 0 0)),((0 0 1,1 0 1,1 1 1,0 1 1,0 0 1)),((0 0 0,0 1 0,0 1 1,0 0 1,0 0 0)),((1 0 0,1 1 0,1 1 1,1 0 1,1 0 0)),((0 0 0,1 0 0,1 0 1,0 0 1,0 0 0)),((0 1 0,1 1 0,1 1 1,0 1 1,0 1 0)))'),
    ST_GeomFromText('POLYHEDRALSURFACE Z(((0.5 0.5 0.5,1.5 0.5 0.5,1.5 1.5 0.5,0.5 1.5 0.5,0.5 0.5 0.5)),((0.5 0.5 1.5,1.5 0.5 1.5,1.5 1.5 1.5,0.5 1.5 1.5,0.5 0.5 1.5)),((0.5 0.5 0.5,0.5 1.5 0.5,0.5 1.5 1.5,0.5 0.5 1.5,0.5 0.5 0.5)),((1.5 0.5 0.5,1.5 1.5 0.5,1.5 1.5 1.5,1.5 0.5 1.5,1.5 0.5 0.5)),((0.5 0.5 0.5,1.5 0.5 0.5,1.5 0.5 1.5,0.5 0.5 1.5,0.5 0.5 0.5)),((0.5 1.5 0.5,1.5 1.5 0.5,1.5 1.5 1.5,0.5 1.5 1.5,0.5 1.5 0.5)))')
);

-- 3D difference
SELECT ST_3DDifference(solid_a, solid_b) FROM solids;

-- 3D union
SELECT ST_3DUnion(solid_a, solid_b) FROM solids;
```

### 3D Measurements

```sql
-- 3D area of a surface
SELECT ST_3DArea(geom) FROM surfaces;

-- Volume of a solid
SELECT ST_Volume(geom) FROM solids;
```

### Extrusion

```sql
-- Extrude a 2D polygon into a 3D solid
SELECT ST_Extrude(
    ST_GeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))'),
    0, 0, 10  -- dx, dy, dz
);
```

### Tesselation and Triangulation

```sql
-- Tesselate a polygon into triangles
SELECT ST_Tesselate(
    ST_GeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))')
);

-- Constrained Delaunay triangulation
SELECT ST_ConstrainedDelaunayTriangles(
    ST_GeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))')
);
```

### Other Functions

```sql
-- Straight skeleton of a polygon
SELECT ST_StraightSkeleton(
    ST_GeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))')
);

-- Approximate medial axis
SELECT ST_ApproximateMedialAxis(
    ST_GeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))')
);

-- Minkowski sum
SELECT ST_MinkowskiSum(
    ST_GeomFromText('LINESTRING(0 0, 4 0)'),
    ST_GeomFromText('POLYGON((0 0, 1 0, 1 1, 0 1, 0 0))')
);

-- Check planarity of a surface
SELECT ST_IsPlanar(geom) FROM surfaces;
```
