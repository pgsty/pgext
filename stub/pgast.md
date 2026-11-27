## Usage

Sources:

- [Official source](https://github.com/demitri/pgAST/blob/cc49e5a46f15275fc0d75917422d31a11c726d22/README.md)
- [Official source](https://github.com/demitri/pgAST/blob/cc49e5a46f15275fc0d75917422d31a11c726d22/src/pgast.control)
- [Official source](https://github.com/demitri/pgAST/blob/cc49e5a46f15275fc0d75917422d31a11c726d22/src/pgast--1.1.0.sql)
- [Official source](https://github.com/demitri/pgAST/blob/cc49e5a46f15275fc0d75917422d31a11c726d22/src/Makefile)

`pgast` exposes astronomical world-coordinate and region operations from Starlink AST. The reviewed source uses extension version 1.1.0 and requires AST 9.5.0 or later. The current test harness requires PostgreSQL 18 or later, and the Makefile explicitly accounts for PostgreSQL 19 headers. No binary-package availability is implied.

### Enable and Query

Install the matching compiled module and its AST dependencies, then create the extension as a superuser. It is relocatable and needs no shared preload:

```sql
CREATE EXTENSION pgast;
SELECT pgast_point_in_polygon(150.0, 2.0,
  '((150.01,1.99),(149.99,1.99),(149.99,2.01),(150.01,2.01))'::polygon);
SELECT pgast_distance(point(150,2), point(151,2));
```

The data types are PostgreSQL's native point, circle and polygon types. PostGIS is not a dependency. Coordinates represent positions on the celestial sphere, normally ICRS in degrees. Polygon edges follow great circles, and their order selects the interior: reversing vertices can describe the rest of the sky.

### FITS and Region Operations

`pgast_bounding_polygon` converts a FITS header to a footprint. `pgast_bounding_circle` accepts a polygon or header. `pgast_point_in_polygon` has polygon and header overloads. `pgast_icrs_polygon_overlaps_hdu` compares an ICRS polygon to a two-dimensional image header, transforming the frame where needed. `pgast_cutout` takes a header, sky position and side length in arcseconds.

Supply a FITS header as a single string of concatenated 80-character records. Native PostgreSQL geometric operators use planar coordinates; they do not replace these spherical tests across RA zero or at the poles. Bounding circles are enclosing approximations, not necessarily minimum circles.

### Limitations and Maintenance

The upstream library can misclassify the near-exact antipode of a small polygon as inside. Invalid headers or unsuitable WCS values raise errors that abort the statement; the documented SQLSTATEs include `22023` and `XX000`. Validate representative footprints before bulk queries. The reviewed tree has an installation script but no SQL upgrade chain; preserve dependent objects and plan migrations explicitly instead of copying the README's destructive drop/recreate example. No project license file was found in the reviewed revision.
