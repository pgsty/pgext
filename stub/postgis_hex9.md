## Usage

Sources:

- [extension/postgis_hex9/README.md](https://github.com/MrBenGriffin/libhex9/blob/1eb0be7d43255b30ecb0c46611fd56b10bf7d773/extension/postgis_hex9/README.md)
- [extension/postgis_hex9/postgis_hex9--2.3.0.sql](https://github.com/MrBenGriffin/libhex9/blob/1eb0be7d43255b30ecb0c46611fd56b10bf7d773/extension/postgis_hex9/postgis_hex9--2.3.0.sql)
- [README.md](https://github.com/MrBenGriffin/libhex9/blob/1eb0be7d43255b30ecb0c46611fd56b10bf7d773/README.md)
- [extension/postgis_hex9/postgis_hex9.control](https://github.com/MrBenGriffin/libhex9/blob/1eb0be7d43255b30ecb0c46611fd56b10bf7d773/extension/postgis_hex9/postgis_hex9.control)

`postgis_hex9` adds Hex9 geographic grids to PostGIS, with UUID cell identifiers and hierarchical spatial aggregation. It requires `postgis` and native libhex9. The default layout reaches layer 30; the legacy layout reaches 29.

### Core Workflow

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_hex9;
SELECT h9_encode('SRID=4326;POINT(-3.19 55.95)'::geometry);
SELECT h9_bin(h9_encode('SRID=4326;POINT(-3.19 55.95)'::geometry), 8);
```

### Operational Boundaries

`h9_encode()` and `h9_decode()` convert points and full addresses; `h9_bin()` returns aggregation keys. `h9_cell()` and `h9_grid()` produce geometry, and neighborhood functions operate on full addresses. No preload or explicit library load is needed. A bin key is not a full point address. WGS84 and spherical inputs produce different addresses, and the UUID records neither datum nor projection regime. Keep that metadata with the dataset. Version 1.x addresses must be re-derived from original geometry for 2.x; do not decode and re-encode old addresses or mix regimes. The control requires superuser installation by default; a PostgreSQL-major support matrix is not declared.
