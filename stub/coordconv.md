## Usage

Sources:

- [Official README](https://github.com/segasai/pg_coordconvert/blob/8acfae816cf373204918c8231eabae26c575dbc3/README)
- [Version 1.0.0 extension SQL](https://github.com/segasai/pg_coordconvert/blob/8acfae816cf373204918c8231eabae26c575dbc3/scripts/coordconv--1.0.0.sql)
- [PostgreSQL wrapper implementation](https://github.com/segasai/pg_coordconvert/blob/8acfae816cf373204918c8231eabae26c575dbc3/pgcoord.c)

`coordconv` converts coordinates between the IAU 1958 Galactic system and the J2000 FK5 equatorial system. All inputs and outputs are decimal degrees; the extension returns each output component through a separate function.

### Core Workflow

```sql
CREATE EXTENSION coordconv;

-- FK5 right ascension/declination to Galactic longitude/latitude.
SELECT fk52gall(266.4051, -28.936175),
       fk52galb(266.4051, -28.936175);

-- Galactic longitude/latitude back to FK5 right ascension/declination.
SELECT gal2fk5ra(0.0, 0.0),
       gal2fk5dec(0.0, 0.0);
```

Pass the same coordinate pair to the two functions for a conversion direction, then combine the two scalar results as the output position.

### Object Index

- `fk52gall(ra, dec)` returns Galactic longitude.
- `fk52galb(ra, dec)` returns Galactic latitude.
- `gal2fk5ra(l, b)` returns J2000 FK5 right ascension.
- `gal2fk5dec(l, b)` returns J2000 FK5 declination.

All four signatures use and return `double precision`; they are `IMMUTABLE`, `STRICT`, and `PARALLEL SAFE`.

### Operational Notes

Version `1.0.0` declares no extension dependency or preload requirement. Inputs are angles in degrees, not radians, sexagesimal strings, or PostGIS geometries. Keep the coordinate frame and epoch with stored values: the functions specifically implement IAU 1958 Galactic and J2000 FK5 conversions, not arbitrary reference-system transformations or proper-motion correction.
