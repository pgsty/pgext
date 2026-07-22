## Usage

Sources:

- [Official README](https://github.com/matlads/mapcode-postgres/blob/c74956e43492bcf9a4761af641753ede07ce122b/README.md)
- [SQL function definitions](https://github.com/matlads/mapcode-postgres/blob/c74956e43492bcf9a4761af641753ede07ce122b/sql/mapcode.sql)
- [Function implementation](https://github.com/matlads/mapcode-postgres/blob/c74956e43492bcf9a4761af641753ede07ce122b/src/funcs.c)
- [Official PGXN distribution page](https://pgxn.org/dist/mapcode/)

`mapcode` wraps the Mapcode C++ library to encode a latitude/longitude pair into short territory-relative mapcodes and decode a mapcode back to coordinates. The package is named `mapcode-postgres`, but the PostgreSQL extension name is `mapcode`.

### Encode and Decode

```sql
CREATE EXTENSION mapcode;

SELECT mapcode_version();
SELECT mapcode_encode(0.341637, 32.593781, 'UGA');

SELECT mapcode_decode('N7.FR', 'UGA');
```

`mapcode_encode(float8, float8, cstring)` takes latitude, longitude, and a territory code. It returns text containing one or more comma-separated mapcodes, or `NULL` when no result is available. `mapcode_decode(cstring, cstring)` returns a `latitude,longitude` text value rather than a geometric or structured SQL type. `mapcode_version()` reports the bundled mapcode library version.

### Data Handling

Validate coordinates before calling the extension: the implementation normalizes out-of-range values instead of rejecting them. Do not assume the first encoded value is the only valid code, and parse comma-separated results and decoded coordinates explicitly. Store the territory alongside every mapcode because decoding depends on it.

This old wrapper vendors a particular Mapcode library snapshot and does not integrate with PostGIS. Pin and record `mapcode_version()`, test representative territories and boundary coordinates against the application that consumes the codes, and define how updated mapcode library data will be migrated.
