
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION nominatim_fdw;
> CREATE SERVER osm FOREIGN DATA WRAPPER nominatim_fdw
>   OPTIONS (url 'https://nominatim.openstreetmap.org');
> ```
>
> Sources: [README](https://github.com/jimjonesbr/nominatim_fdw), [Nominatim API](https://nominatim.org/)

`nominatim_fdw` is a PostgreSQL FDW-style extension for calling Nominatim geocoding services from SQL. The extension is organized around functions rather than foreign tables and maps to Nominatim's `search`, `reverse`, and `lookup` endpoints.

## Server Setup

Create the extension and define a server pointing at a Nominatim endpoint:

```sql
CREATE EXTENSION nominatim_fdw;

CREATE SERVER osm
FOREIGN DATA WRAPPER nominatim_fdw
OPTIONS (url 'https://nominatim.openstreetmap.org');
```

The README documents these server options:

- `url` as the required endpoint URL
- `http_proxy`
- `connect_timeout` with default `300`
- `max_connect_retry` with default `3`
- `max_connect_redirect` where `0` means unlimited redirects

Server options can be changed with `ALTER SERVER`:

```sql
ALTER SERVER osm OPTIONS (ADD max_connect_retry '5');
ALTER SERVER osm OPTIONS (SET url 'https://a.new.url');
ALTER SERVER osm OPTIONS (DROP http_proxy);
```

Proxy credentials belong in a user mapping, not in the server definition:

```sql
CREATE USER MAPPING FOR pguser
SERVER osm
OPTIONS (proxy_user 'myuser', proxy_password 'mysecret');
```

## Geocoding Functions

### Search

`nominatim_search` supports both structured and free-form queries:

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

### Reverse Geocoding

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

### OSM Object Lookup

```sql
SELECT osm_id, display_name
FROM nominatim_lookup(
  server_name => 'osm',
  osm_ids => 'W121736959,R123456'
);
```

The README notes that lookup IDs use OSM type prefixes such as `N` for nodes, `W` for ways, and `R` for relations.

## Notes

The current upstream README lists these requirements:

- PostgreSQL 12 or newer
- `libxml2` 2.5.0 or newer
- `libcurl` 7.74.0 or newer

The extension also exposes `nominatim_fdw_version()` for version checks and supports extension upgrades through `ALTER EXTENSION nominatim_fdw UPDATE`.
