## Usage

Sources:

- [Project README and accuracy disclaimer](https://github.com/gabriel-russo/postgresql-sgp4/blob/bcf6ea2b45285856f1b7b68c9da78ba129a5249d/README.md)
- [Extension control file](https://github.com/gabriel-russo/postgresql-sgp4/blob/bcf6ea2b45285856f1b7b68c9da78ba129a5249d/sgp4.control)
- [Version 0.1 SQL API](https://github.com/gabriel-russo/postgresql-sgp4/blob/bcf6ea2b45285856f1b7b68c9da78ba129a5249d/sgp4--0.1.sql)
- [PostgreSQL wrapper implementation](https://github.com/gabriel-russo/postgresql-sgp4/blob/bcf6ea2b45285856f1b7b68c9da78ba129a5249d/sgp4.c)

`sgp4` 0.1 embeds an SGP4 orbit propagator in PostgreSQL. `satellite_geographic_position()` accepts the two lines of a TLE and a `timestamp without time zone`, then returns WKT text in the form `POINT Z(longitude latitude altitude)`. Longitude and latitude are degrees on WGS84; altitude is kilometers.

### Propagate a TLE at a UTC instant

```sql
CREATE EXTENSION sgp4;

SELECT satellite_geographic_position(
  '1 44883U 19093E   26042.94689667  .00001181  00000+0  15674-3 0  9994',
  '2 44883  97.7645 114.1422 0001434 192.3199 167.7980 14.81620074332535',
  timestamp '2026-02-12 12:00:00'
);
```

The function performs no time-zone conversion. Convert an authoritative instant to UTC before removing its zone, and document that convention at every call site. Check TLE checksums, epoch, catalog identity, and freshness before propagation; TLE accuracy decays outside its useful epoch window.

### Accuracy and spatial boundaries

Upstream explicitly excludes scientific, astrometric, high-precision geodetic, safety-critical, and regulatory use. Its TEME-to-Earth-fixed conversion omits full IAU precession and nutation, equation-of-equinoxes, polar motion, and real-time UT1 corrections. Longitude differences from higher-precision libraries can reach fractions of a degree depending on orbit and epoch.

The result is text, not a PostGIS geometry. If converted to SRID 4326, remember that the Z value remains kilometers and that the horizontal SRID does not define a vertical reference or unit. Validate output against a trusted SGP4 implementation over representative orbits, malformed TLEs, leap and boundary dates, and expected error tolerances. Do not use database errors or a plausible WKT string as evidence that input was valid.
