## Usage

Sources:

- [geogridcoder/yukon_geogridcoder.control](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/geogridcoder/yukon_geogridcoder.control)
- [README_EN.md](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/README_EN.md)
- [geogridcoder/yukon_geogridcoder--1.0.1.sql](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/geogridcoder/yukon_geogridcoder--1.0.1.sql)
- [doc/yukon_doc/source/installation.rst](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/doc/yukon_doc/source/installation.rst)

`yukon_geogridcoder` is a module of Yukon for the openGauss kernel. It supplies the geosotgrid type, scalar comparison and grid-array operators, and conversions between geometry and GeoSOT grids. Both 2D and altitude-related conversion routines are present.

### Core Workflow

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION yukon_geogridcoder;
SELECT ST_GeoSOTGrid(ST_GeomFromText('POINT(116.4 39.9)', 4326), 15);
```

### Operational Boundaries

Use a compatible Yukon/openGauss installation, including its adapted PostGIS types and libraries. Enable as an appropriately privileged database administrator; no preload is specified. This catalog entry is not a stock PostgreSQL compatibility claim. Geometry/SRID conventions and model/grid encodings follow the Yukon API; use matching clients and preserve associated metadata during backup and maintenance.
