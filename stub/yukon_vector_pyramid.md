## Usage

Sources:

- [pyramid/yukon_vector_pyramid.control](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/pyramid/yukon_vector_pyramid.control)
- [README_EN.md](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/README_EN.md)
- [pyramid/yukon_vector_pyramid--1.0.sql](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/pyramid/yukon_vector_pyramid--1.0.sql)
- [doc/yukon_doc/source/installation.rst](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/doc/yukon_doc/source/installation.rst)

`yukon_vector_pyramid` is a module of Yukon for the openGauss kernel. The SQL surface includes ST_BuildPyramid, ST_ListPyramid, ST_HasPyramid, ST_DeletePyramid, ST_BuildTile, ST_UpdatePyramid and ST_AsTile. Pyramid construction creates derived state that must be maintained with source data.

### Core Workflow

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION yukon_vector_pyramid;
SELECT yukon_pyramid_version();
```

### Operational Boundaries

Use a compatible Yukon/openGauss installation, including its adapted PostGIS types and libraries. Enable as an appropriately privileged database administrator; no preload is specified. This catalog entry is not a stock PostgreSQL compatibility claim. Geometry/SRID conventions and model/grid encodings follow the Yukon API; use matching clients and preserve associated metadata during backup and maintenance.
