
## Usage

Sources: [README](https://github.com/jimjonesbr/nominatim_fdw/blob/master/README.md), [Nominatim API](https://nominatim.org/)

`nominatim_fdw` is a PostgreSQL foreign data wrapper for Nominatim geocoding services. Upstream exposes it through SQL functions mapped to the Nominatim `search`, `reverse`, and `lookup` endpoints rather than through foreign tables.

### Create a server

```sql
CREATE EXTENSION nominatim_fdw;

CREATE SERVER osm
FOREIGN DATA WRAPPER nominatim_fdw
OPTIONS (url 'https://nominatim.openstreetmap.org');
```

Documented server options:

- `url` (required)
- `http_proxy`
- `connect_timeout`
- `max_connect_retry`
- `max_connect_redirect`

Proxy credentials belong in a user mapping:

```sql
CREATE USER MAPPING FOR pguser
SERVER osm
OPTIONS (proxy_user 'myuser', proxy_password 'mysecret');
```

### Geocoding functions

Structured or free-form search:

```sql
SELECT osm_id, ref, lon, lat, boundingbox
FROM nominatim_search(
  server_name => 'osm',
  street => 'Neubrueckenstrasse 63',
  city => 'Muenster',
  country => 'Germany'
);

SELECT osm_id, display_name, lon, lat
FROM nominatim_search(
  server_name => 'osm',
  q => '1600 Pennsylvania Avenue, Washington DC'
);
```

Reverse lookup:

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

OSM object lookup:

```sql
SELECT osm_id, display_name
FROM nominatim_lookup(
  server_name => 'osm',
  osm_ids => 'W121736959,R123456'
);
```

The README notes OSM ID prefixes such as `N` for nodes, `W` for ways, and `R` for relations.

### Notes

- Upstream documents PostgreSQL 12+, `libxml2` 2.5.0+, and `libcurl` 7.74.0+.
- The extension exposes `nominatim_fdw_version()`.
- The current README documents `CREATE EXTENSION ... WITH VERSION '1.3'` and `ALTER EXTENSION ... UPDATE TO '1.3'`, so upstream has moved past the requested `1.2` refresh target.
