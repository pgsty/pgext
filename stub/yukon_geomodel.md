## Usage

Sources:

- [geomodel/src/yukon_geomodel.control](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/geomodel/src/yukon_geomodel.control)
- [README_EN.md](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/README_EN.md)
- [geomodel/src/yukon_geomodel--1.0.1.sql](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/geomodel/src/yukon_geomodel--1.0.1.sql)
- [doc/yukon_doc/source/installation.rst](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/doc/yukon_doc/source/installation.rst)
- [doc/yukon_doc/source/modular_api/geomodel_api.rst](https://github.com/opengauss-mirror/Yukon/blob/7c1063e45f136dca9d456e5bb56fafd4ea973ef2/doc/yukon_doc/source/modular_api/geomodel_api.rst)

`yukon_geomodel` is a module of Yukon for the openGauss kernel. It supplies GEOMODEL and MODEL_ELEM values. Use the management functions to add/drop model columns and their subsidiary tables; upstream warns against managing these structures through direct DDL.

### Core Workflow

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION yukon_geomodel;
CREATE TABLE model_sample (id integer);
SELECT AddGeoModelColumn('public', 'model_sample', 'model', 4326);
```

### Operational Boundaries

Use a compatible Yukon/openGauss installation, including its adapted PostGIS types and libraries. Enable as an appropriately privileged database administrator; no preload is specified. This catalog entry is not a stock PostgreSQL compatibility claim. Geometry/SRID conventions and model/grid encodings follow the Yukon API; use matching clients and preserve associated metadata during backup and maintenance.
