

## Usage

> Source: [`segasai/q3c`](https://github.com/segasai/q3c) | [ADASS Paper](http://adsabs.harvard.edu/abs/2006ASPC..351..735K) | [ASCL](https://ascl.net/1905.008)

Q3C (Quad Tree Cube) is a PostgreSQL extension for fast sky-indexing of astronomical catalogues. It enables efficient spatial queries on spherical coordinates (right ascension and declination), including cone searches, ellipse searches, polygon queries, positional cross-matches, and nearest-neighbor lookups.

All angles (ra, dec, distances) are in **degrees**, proper motions in **mas/year**, and epochs in **years** (e.g. 2000.5, 2010.5). All Q3C function names start with the `q3c_` prefix.

### Table Preparation

To use Q3C, create a spatial index on your table with `ra` and `dec` columns (in degrees):

```sql
CREATE INDEX ON mytable (q3c_ang2ipix(ra, dec));
```

Optionally cluster the table by the index to ensure faster queries on large datasets:

```sql
CLUSTER mytable_q3c_ang2ipix_idx ON mytable;
```

Alternatively, reorder the table before indexing:

```sql
CREATE TABLE mytable1 AS SELECT * FROM mytable ORDER BY q3c_ang2ipix(ra, dec);
```

After indexing, analyze the table:

```sql
ANALYZE mytable;
```


## Functions

- `q3c_ang2ipix(ra, dec)` -- returns the ipix value (64-bit integer pixel identifier) for given ra and dec

- `q3c_dist(ra1, dec1, ra2, dec2)` -- returns the distance in degrees between two points

- `q3c_dist_pm(ra1, dec1, pmra1, pmdec1, cosdec_flag, epoch1, ra2, dec2, epoch2)` -- returns distance in degrees between two points, taking proper motion into account. The `cosdec_flag` (0 or 1) indicates whether the proper motion includes the cos(dec) term (1) or not (0).

- `q3c_join(ra1, dec1, ra2, dec2, radius)` -- returns true if (ra1, dec1) is within `radius` spherical distance of (ra2, dec2). Use when the index on `q3c_ang2ipix(ra2, dec2)` is created.

- `q3c_join_pm(ra1, dec1, pmra1, pmdec1, cosdec_flag, epoch1, ra2, dec2, epoch2, max_delta_epoch, radius)` -- like `q3c_join` but takes proper motion into account. `max_delta_epoch` is the maximum epoch difference possible between two tables.

- `q3c_ellipse_join(ra1, dec1, ra2, dec2, major, ratio, pa)` -- like `q3c_join`, except (ra1, dec1) must be within an ellipse with semi-major axis `major`, axis ratio `ratio`, and position angle `pa` (from north through east)

- `q3c_radial_query(ra, dec, center_ra, center_dec, radius)` -- returns true if (ra, dec) is within `radius` degrees of (center_ra, center_dec). Main function for cone searches. Requires index on `q3c_ang2ipix(ra, dec)`.

- `q3c_ellipse_query(ra, dec, center_ra, center_dec, maj_ax, axis_ratio, PA)` -- returns true if (ra, dec) is within the ellipse from (center_ra, center_dec), specified by semi-major axis, axis ratio, and positional angle.

- `q3c_poly_query(ra, dec, poly)` -- returns true if (ra, dec) is within the spherical polygon specified as an array of RA/DEC values or a PostgreSQL polygon type. Uses the index.

- `q3c_ipix2ang(ipix)` -- returns a two-element array of (ra, dec) corresponding to a given ipix

- `q3c_pixarea(ipix, bits)` -- returns the spherical area corresponding to a given ipix at the pixelisation level given by `bits` (1 is smallest, 30 is the cube face)

- `q3c_ipixcenter(ra, dec, bits)` -- returns the ipix value of the pixel center at certain pixel depth covering the specified (ra, dec)

- `q3c_in_poly(ra, dec, poly)` -- returns true/false if point is inside a polygon. Does **NOT** use the q3c index.

- `q3c_version()` -- returns the installed version of Q3C


## Examples

### Cone Search

Query all objects within 0.1 degrees of (ra, dec) = (11, 12):

```sql
SELECT * FROM mytable WHERE q3c_radial_query(ra, dec, 11, 12, 0.1);
```

The column names of the table must come first, and the search location after, otherwise the index will not be used.

Alternative cone search using `q3c_join` (can be faster for small tables):

```sql
SELECT * FROM mytable WHERE q3c_join(11, 12, ra, dec, 0.1);
```

### Ellipse Search

Search for objects within an ellipse centered at (10, 20) with semi-major axis 1 degree, axis ratio 0.5, and PA of 10 degrees:

```sql
SELECT * FROM mytable WHERE q3c_ellipse_query(ra, dec, 10, 20, 1, 0.5, 10);
```

### Polygon Search

Query objects inside a spherical polygon with vertices (0,0), (2,0), (2,1), (0,1):

```sql
SELECT * FROM mytable WHERE
    q3c_poly_query(ra, dec, ARRAY[0, 0, 2, 0, 2, 1, 0, 1]);
```

Using PostgreSQL polygon type:

```sql
SELECT * FROM mytable WHERE
    q3c_poly_query(ra, dec, '((0, 0), (2, 0), (2, 1), (0, 1))'::polygon);
```

### Positional Cross-Match

Cross-match `table1` and `table2` within 0.001 degrees. The index must exist on `q3c_ang2ipix(ra, dec)` of `table2`:

```sql
SELECT * FROM table1 AS a, table2 AS b WHERE
    q3c_join(a.ra, a.dec, b.ra, b.dec, 0.001);
```

The ra/dec columns from the indexed table must be the 3rd and 4th arguments. This returns **all** pairs within the matching distance, not just nearest neighbors.

With per-object error radius:

```sql
SELECT * FROM table1 AS a, table2 AS b WHERE
    q3c_join(a.ra, a.dec, b.ra, b.dec, a.err);
```

### Ellipse Cross-Match

Cross-match using elliptical error areas (e.g., matching within galaxy elliptical bodies):

```sql
SELECT * FROM table1 AS a, table2 AS b WHERE
    q3c_ellipse_join(a.ra, a.dec, b.ra, b.dec, a.maj_ax, a.axis_ratio, a.PA);
```

### Cross-Match with Proper Motion

Cross-match with proper motion correction. Assumes `table1` has `pmra`, `pmdec` (mas/yr) and `epoch` columns, pmra includes cos(dec) factor, and max epoch difference is 30 years:

```sql
SELECT * FROM table1 AS a, table2 AS b WHERE
    q3c_join_pm(a.ra, a.dec, a.pmra, a.pmdec, 1,
    a.epoch, b.ra, b.dec, b.epoch, 30, 0.001);
```

### Nearest Neighbour (with NULLs for unmatched)

Returns the nearest neighbour for each row, with NULLs if no match exists within 1 arcsecond:

```sql
SELECT t.*, ss.* FROM mytable AS t
LEFT JOIN LATERAL (
    SELECT s.*
    FROM sdssdr9.phototag AS s
    WHERE q3c_join(t.ra, t.dec, s.ra, s.dec, 1./3600)
    ORDER BY q3c_dist(t.ra, t.dec, s.ra, s.dec) ASC
    LIMIT 1
) AS ss ON true;
```

### Nearest Neighbour (matched only)

Returns only objects that have neighbours:

```sql
SELECT t.*, ss.* FROM mytable AS t,
LATERAL (
    SELECT s.*
    FROM sdssdr9.phototag AS s
    WHERE q3c_join(t.ra, t.dec, s.ra, s.dec, 1./3600)
    ORDER BY q3c_dist(t.ra, t.dec, s.ra, s.dec) ASC
    LIMIT 1
) AS ss;
```

### Nearest Neighbour (CTE variant)

Uses a CTE with an object ID column (requires an index on the ID column):

```sql
WITH x AS MATERIALIZED (
    SELECT *, (
        SELECT objid FROM sdssdr9.phototag AS p
        WHERE q3c_join(m.ra, m.dec, p.ra, p.dec, 1./3600)
        ORDER BY q3c_dist(m.ra, m.dec, p.ra, p.dec) ASC
        LIMIT 1
    ) AS match_objid
    FROM mytable AS m
)
SELECT * FROM x, sdssdr9.phototag AS s WHERE x.match_objid = s.objid;
```

### Density Estimation

Estimate object density using pixelation depth of 25:

```sql
SELECT (q3c_ipix2ang(i))[1] AS ra,
       (q3c_ipix2ang(i))[2] AS dec,
       c,
       q3c_pixarea(i, 25) AS area
FROM (
    SELECT q3c_ipixcenter(ra, dec, 25) AS i, count(*) AS c
    FROM mytable
    GROUP BY i
) AS x;
```

Note: Q3C does not have uniform pixel areas (unlike HEALPIX).


## Limitations

- Querying very large polygons with diameter greater than ~25 degrees is not supported
- Polygons with more than 100 vertices are not supported


## Performance Tips

- Ensure correct argument order in Q3C functions (e.g., `q3c_radial_query(ra, dec, 120, 3, 1)` not `q3c_radial_query(120, 3, ra, dec, 1)`)
- Use `EXPLAIN` to verify the query plan uses bitmap scans on the Q3C index
- If the planner chooses a bad plan, try: `SET enable_mergejoin TO off; SET enable_seqscan TO off; SET enable_hashjoin TO off;`
- Cluster the table using the Q3C index for best performance
- When combining `q3c_join()` with additional filter clauses, use CTEs with `MATERIALIZED` to avoid plan issues:

```sql
WITH x AS MATERIALIZED (SELECT * FROM t1 WHERE t1.mag < 1),
     y AS (SELECT *, t2.mag AS t2mag FROM x, t2 WHERE q3c_join(x.ra, x.dec, t2.ra, t2.dec, 1./3600))
SELECT * FROM y WHERE t2mag > 33;
```
