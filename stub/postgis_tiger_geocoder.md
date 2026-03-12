

## Usage

> [PostGIS TIGER Geocoder: US Census TIGER/Line geocoding for PostGIS](https://github.com/postgis/postgis)

The PostGIS TIGER Geocoder provides geocoding and reverse geocoding capabilities for US addresses using US Census TIGER/Line data. It can parse an address string into a normalized form, find the geographic coordinates, and reverse-geocode coordinates back to an address.

- [TIGER Geocoder Reference](https://postgis.net/docs/Extras.html)

### Setup

```sql
CREATE EXTENSION postgis_tiger_geocoder CASCADE;
```

This creates the `tiger` schema with the geocoder tables and functions.

--------

## Loading TIGER Data

Before geocoding, TIGER/Line data must be loaded for the states you need. The extension provides helper functions to generate the loading scripts:

```sql
-- Generate a script to download and load data for a state
-- (e.g., Massachusetts = 'MA')
SELECT loader_generate_script(ARRAY['MA'], 'sh');
```

This generates a shell script that uses `shp2pgsql` to load TIGER shapefiles. Run the generated script to populate the `tiger_data` schema with address ranges, edges, faces, and other data.

After loading:

```sql
-- Install missing indexes for performance
SELECT install_missing_indexes();

-- Update statistics
ANALYZE tiger.addr;
ANALYZE tiger.edges;
ANALYZE tiger.faces;
```

--------

## Geocoding

Convert a US address string to geographic coordinates:

```sql
-- Basic geocoding
SELECT g.rating, ST_X(g.geomout) AS lon, ST_Y(g.geomout) AS lat,
       pprint_addy(g.addy) AS address
FROM geocode('1600 Pennsylvania Ave NW, Washington, DC 20500') AS g;
```

The `rating` indicates match quality (lower is better, 0 = exact match).

```sql
-- Geocode with a limit on results
SELECT g.rating, ST_AsText(g.geomout), pprint_addy(g.addy)
FROM geocode('100 Main St, Boston, MA', 3) AS g;

-- Batch geocode from a table
SELECT a.id, g.rating, g.geomout, pprint_addy(g.addy)
FROM addresses a, LATERAL geocode(a.address_string, 1) AS g;
```

--------

## Reverse Geocoding

Convert coordinates back to a street address:

```sql
SELECT pprint_addy(r.addy[1]) AS address
FROM reverse_geocode(ST_SetSRID(ST_MakePoint(-77.0365, 38.8977), 4326)) AS r;
```

--------

## Address Normalization

Parse and normalize address strings without geocoding:

```sql
SELECT *
FROM normalize_address('1600 Pennsylvania Avenue NW, Washington, DC 20500');
```

Returns components: `address` (number), `predirAbbrev`, `streetName`, `streetTypeAbbrev`, `postdirAbbrev`, `internal`, `location` (city), `stateAbbrev`, `zip`, `parsed`.

```sql
-- Pretty-print a normalized address
SELECT pprint_addy(normalize_address('100 main street boston ma 02101'));
```

--------

## Configuration

The `tiger.geocode_settings` table controls geocoder behavior:

```sql
-- View current settings
SELECT * FROM tiger.geocode_settings;

-- Adjust settings (e.g., increase debug level)
UPDATE tiger.geocode_settings SET val = 'true' WHERE name = 'debug_geocode_address';
```
