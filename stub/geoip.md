

## Usage

> [geoip: IP-based geolocation for PostgreSQL](https://github.com/tvondra/geoip)

This extension provides IP-based geolocation — you provide an IPv4 or IPv6 address and the extension looks up country, city, GPS coordinates, ASN and more. It requires the `ip4r` extension and GeoLite2 data from [MaxMind](https://www.maxmind.com).

```sql
CREATE EXTENSION ip4r;
CREATE EXTENSION geoip;
```

### Functions

| Function | Description |
|----------|-------------|
| `geoip_country_code(ip4\|ip6)` | Returns country code (2 chars) |
| `geoip_country(ip4\|ip6)` | Returns all country info (code, name, network) |
| `geoip_city_location(ip4\|ip6)` | Returns just the location ID (INT) |
| `geoip_city(ip4\|ip6)` | Returns all city info (GPS, ZIP code, etc.) |
| `geoip_asn(ip4\|ip6)` | Returns ASN name and IP range |

### Examples

```sql
SELECT geoip_country_code('78.45.133.255'::ip4);
-- CZ

SELECT * FROM geoip.geoip_city('78.45.133.255'::ip4);
--  geoname_id | country_iso_code | city_name | postal_code | ...
-- ------------+------------------+-----------+-------------+----
--     3066399 | CZ               | Sardice   | 696 13      | ...

SELECT * FROM geoip.geoip_country('78.45.133.255'::ip4);
--     network     | country_iso_code | country_name
-- ----------------+------------------+--------------
--  78.45.128.0/17 | CZ               | Czechia

SELECT * FROM geoip.geoip_asn('78.45.133.255'::ip4);
--    network    | asn_number |      asn_name
-- --------------+------------+---------------------
--  78.44.0.0/15 |       6830 | Liberty Global B.V.
```


## Loading Data

The extension requires GeoLite2 CSV data from MaxMind. Download the City, Country, and ASN datasets in CSV format from [MaxMind GeoLite2](https://dev.maxmind.com/geoip/geoip2/geolite2/), then load:

```bash
cat GeoLite2-Country-Locations-en.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_country_locations FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-Country-Blocks-IPv4.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_country_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-Country-Blocks-IPv6.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_country_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-City-Locations-en.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_city_locations FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-City-Blocks-IPv4.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_city_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-City-Blocks-IPv6.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_city_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-ASN-Blocks-IPv4.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_asn_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-ASN-Blocks-IPv6.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_asn_blocks FROM stdin WITH (FORMAT CSV, HEADER)'
```

The "locations" files have multiple language variants — pick the one that works for you.
