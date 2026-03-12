


## Usage

> [nominatim_fdw: Nominatim Foreign Data Wrapper for PostgreSQL](https://github.com/jimjonesbr/nominatim_fdw)

This FDW provides access to [Nominatim](https://nominatim.org/) geocoding services directly from PostgreSQL using SQL functions rather than traditional foreign tables.

### Create Server

```sql
CREATE EXTENSION nominatim_fdw;

CREATE SERVER osm FOREIGN DATA WRAPPER nominatim_fdw
  OPTIONS (url 'https://nominatim.openstreetmap.org');
```

**Server Options:** `url` (required, Nominatim endpoint URL), `http_proxy`, `proxy_user`, `proxy_user_password`, `connect_timeout` (default 300 seconds), `max_connect_retry` (default 3), `max_request_redirect` (0 = unlimited).

### Geocoding (Address to Coordinates)

Structured search:

```sql
SELECT osm_id, ref, lon, lat, boundingbox
FROM nominatim_search(
  server_name => 'osm',
  street => 'Neubrueckenstrasse 63',
  city => 'Muenster',
  country => 'Germany'
);
```

Free-form search:

```sql
SELECT osm_id, display_name, lon, lat
FROM nominatim_search(
  server_name => 'osm',
  q => '1600 Pennsylvania Avenue, Washington DC'
);
```

**Parameters:** `q` (free-form query), `street`, `city`, `county`, `state`, `country`, `postalcode`, `amenity`, `limit` (default 10), `addressdetails`, `extratags`, `namedetails`.

### Reverse Geocoding (Coordinates to Address)

```sql
SELECT osm_id, display_name, boundingbox
FROM nominatim_reverse(
  server_name => 'osm',
  lon => -77.0365,
  lat => 38.8977,
  zoom => 18,
  addressdetails => true
);
```

**Parameters:** `lon`, `lat` (required), `zoom` (default 18), `addressdetails`, `extratags`, `namedetails`, `layer`.

### OSM Object Lookup

```sql
SELECT osm_id, display_name
FROM nominatim_lookup(
  server_name => 'osm',
  osm_ids => 'W121736959,R123456'
);
```

**Parameters:** `osm_ids` (comma-separated list of OSM IDs with type prefix: N=node, W=way, R=relation), `addressdetails`, `extratags`, `namedetails`.

### Version Check

```sql
SELECT nominatim_fdw_version();
```
