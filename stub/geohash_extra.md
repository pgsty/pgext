## Usage

Sources:

- [Upstream README](https://github.com/geo-mathijs/geohash-extra/blob/f65bdc89f60c041bbf10a7b11badae237e30480f/README.md)
- [Extension control file](https://github.com/geo-mathijs/geohash-extra/blob/f65bdc89f60c041bbf10a7b11badae237e30480f/geohash_extra.control)
- [SQL wrappers](https://github.com/geo-mathijs/geohash-extra/blob/f65bdc89f60c041bbf10a7b11badae237e30480f/geohash_extra--0.0.1.sql)
- [C implementation](https://github.com/geo-mathijs/geohash-extra/blob/f65bdc89f60c041bbf10a7b11badae237e30480f/geohash_extra.c)

`geohash_extra` supplements PostGIS with two set-returning operations: `geohash_neighbours()` emits the eight adjacent hashes, and `geohashfromgeom()` emits a grid of hashes derived from a geometry's bounds. The catalog and control file both report version `0.0.1`.

### Neighbours and geometry coverage

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION geohash_extra;

SELECT * FROM geohash_neighbours('u173zq37x014');

SELECT *
FROM geohashfromgeom(
  ST_MakeEnvelope(4.89, 52.36, 4.91, 52.38, 4326),
  7
);
```

`geohashfromgeom()` transforms non-4326 input to SRID 4326, then constructs a rectangular grid from bounding-box corner hashes. It is not an exact cover of an irregular geometry and can return cells outside that geometry. The control file does not declare PostGIS as a dependency, so install PostGIS first.

The current C implementation allocates the neighbour input/output buffers one byte too short for their terminators and incompletely validates hash characters. Empty or malformed input can therefore corrupt memory or crash a backend. Do not expose `geohash_neighbours()` to untrusted input without fixing and rebuilding the code. The repository provides no modern PostgreSQL compatibility matrix.
