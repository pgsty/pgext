## Usage

Sources:

- [Official README](https://github.com/Geodan/plv8_geo/blob/ee6bd4878ef234aad5dadfeee7d29b19f0525948/README.md)
- [Extension control file](https://github.com/Geodan/plv8_geo/blob/ee6bd4878ef234aad5dadfeee7d29b19f0525948/plv8geo.control)
- [Extension SQL for version 0.0.2](https://github.com/Geodan/plv8_geo/blob/ee6bd4878ef234aad5dadfeee7d29b19f0525948/plv8geo--0.0.2.sql)
- [D3 hexbin wrapper](https://github.com/Geodan/plv8_geo/blob/ee6bd4878ef234aad5dadfeee7d29b19f0525948/src/d3_hexbin.sql)

`plv8geo` bundles selected JavaScript geometry libraries into PostgreSQL and exposes PL/v8 wrappers for TopoJSON, contours, hexbins, triangulation, polygon triangulation, and Voronoi operations. It requires both `plv8` and `postgis`; version `0.0.2` is an old stack tested by upstream with PostgreSQL 9.4 and PL/v8 2.0.3.

### Load Modules and Call Functions

Creating the extension stores bundled JavaScript in `public.plv8_modules` and creates functions under the `plv8` schema. Initialize the loader and load each module required by the function before calling it.

```sql
CREATE EXTENSION plv8;
CREATE EXTENSION postgis;
CREATE EXTENSION plv8geo;

SELECT plv8.plv8_startup();
DO LANGUAGE plv8 'load_module("d3")';
DO LANGUAGE plv8 'load_module("d3-hexbin")';

SELECT plv8.d3_hexbin(
  '[[1,2],[0.5,0.5],[2,2]]'::json,
  '["foo","bar","baz"]'::json,
  1
);
```

The README suggests configuring `plv8.startproc = 'plv8.plv8_startup'` for session startup. PL/v8 startup settings have changed across releases, so verify the exact GUC supported by the installed PL/v8 version; explicit loading as above is the auditable fallback.

### Important Objects

- `public.plv8_modules`: bundled JavaScript source and load flags.
- `plv8.plv8_startup()`: defines the session's `load_module` helper.
- `plv8.d3_totopojson`, `plv8.d3_simplifytopology`, `plv8.d3_mergetopology`, `plv8.d3_topologytofeatures`: TopoJSON conversion and manipulation.
- `plv8.d3_contour`, `plv8.d3_slopecontours`, `plv8.d3_hexbin`: raster contours, slope contours, and point hexbins.
- `plv8.delaunator`, `plv8.earcut`, `plv8.jsts_voronoi`: triangulation and Voronoi wrappers.

### Compatibility and Security

The package embeds dated library versions, including D3 4.x, TopoJSON 3.0.0, GeoTIFF 0.4.1, Delaunator 1.0.2, Earcut 2.1.1, and JSTS 1.4.0. Do not assume current JavaScript APIs, bug fixes, or security properties. Validate compatibility with the exact PostgreSQL, PostGIS, PL/v8, and V8 builds, and compare results and SRIDs against native PostGIS equivalents.

JavaScript is evaluated inside database backends and can consume substantial memory and CPU. Restrict CREATE and EXECUTE privileges, audit the embedded source, and keep untrusted geometries and JSON size-bounded. Several wrappers are declared IMMUTABLE even though the stack uses session-global loaded modules; validate determinism before using them in expression indexes or generated columns. The control file names schema `plv8geo`, but the install SQL explicitly creates and uses `plv8` plus the public module table, so schema-locality assumptions are unsafe.
