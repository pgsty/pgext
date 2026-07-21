## Usage

Sources:

- [nominatim_fdw v2.0 README](https://github.com/jimjonesbr/nominatim_fdw/blob/v2.0/README.md)
- [nominatim_fdw v2.0 changelog](https://github.com/jimjonesbr/nominatim_fdw/blob/v2.0/CHANGELOG.md)
- [Extension control file](https://github.com/jimjonesbr/nominatim_fdw/blob/v2.0/nominatim_fdw.control)
- [Official Nominatim API overview](https://nominatim.org/release-docs/develop/api/Overview/)
- [OpenStreetMap Nominatim usage policy](https://operations.osmfoundation.org/policies/nominatim/)

`nominatim_fdw` calls a Nominatim geocoding service from PostgreSQL. Unlike a conventional FDW, it exposes functions for search, reverse geocoding, and OSM-object lookup; it does not create queryable foreign tables.

### Configure a Server

```sql
CREATE EXTENSION nominatim_fdw;

CREATE SERVER osm
  FOREIGN DATA WRAPPER nominatim_fdw
  OPTIONS (
    url 'https://nominatim.openstreetmap.org',
    connect_timeout '10',
    max_connect_retry '2',
    max_connect_redirect '1',
    accept_language 'en'
  );
```

The public OpenStreetMap endpoint has an official usage policy. For sustained or bulk workloads, use an authorized provider or operate your own Nominatim service, identify the application as required, and respect rate limits.

### Core Workflow

Free-form search:

```sql
SELECT osm_id, display_name, lon, lat
FROM nominatim_search(
  server_name => 'osm',
  q => 'Neubrückenstraße 63, Münster, Germany'
);
```

Reverse geocoding and object lookup:

```sql
SELECT osm_id, display_name, addressdetails
FROM nominatim_reverse(
  server_name => 'osm',
  lon => 7.6293,
  lat => 51.9648,
  addressdetails => true
);

SELECT osm_id, display_name
FROM nominatim_lookup(
  server_name => 'osm',
  osm_ids => 'W121736959'
);
```

### Important Objects

- `nominatim_search(...)` implements free-form or structured forward search.
- `nominatim_reverse(...)` resolves longitude and latitude to the nearest suitable OSM address.
- `nominatim_lookup(...)` fetches node, way, or relation identifiers such as `N123`, `W456`, or `R789`.
- `nominatim_fdw_version()` reports the extension and principal library versions.
- `nominatim_fdw_settings` exposes dependency and build versions as rows.
- Server options include `url`, proxy configuration, timeouts, retry and redirect limits, and default `accept_language`.

All endpoint functions are `STRICT`: an explicit SQL `NULL` argument returns no rows without sending a request. In 2.0 they are correctly declared `VOLATILE`, because responses are remote and can change.

### Version 2.0 Changes and Caveats

Version 2.0 validates reverse coordinates, adds `email`, `polygon_threshold`, and `entrances`, exposes dependency settings, and fixes JSON escaping for returned detail fields. It also has user-visible compatibility changes: reverse output uses `display_name`; `addressparts` becomes `addressdetails`; address details default to true for reverse and lookup; and version output is shorter. Review result-column consumers before upgrading from 1.3.

Each call performs network I/O in the database statement. Use finite timeouts, constrain who can create or alter servers, and avoid invoking a public service once per row in a large query. The upstream build requires PostgreSQL 10 or newer, libxml2 2.5 or newer, and libcurl 7.74 or newer.
