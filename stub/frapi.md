## Usage

Sources:

- [Official README](https://github.com/adauhr/pg_frapi/blob/f6ac6decaf54558ed34beda4f35b9f027cb55835/README.md)
- [Official extension SQL](https://github.com/adauhr/pg_frapi/blob/f6ac6decaf54558ed34beda4f35b9f027cb55835/frapi--0.1.2.sql)
- [Official extension control file](https://github.com/adauhr/pg_frapi/blob/f6ac6decaf54558ed34beda4f35b9f027cb55835/frapi.control)

`frapi` is an alpha-stage SQL wrapper around the French national address API at `api-adresse.data.gouv.fr`. Version `0.1.2` converts search and reverse-geocoding responses into rows with PostGIS points, but it executes network requests through untrusted `plsh` code and is appropriate only for isolated, trusted experiments.

### Core Workflow

The extension is fixed to the `frapi` schema and requires both `plsh` and `postgis`:

```sql
CREATE EXTENSION plsh;
CREATE EXTENSION postgis;
CREATE EXTENSION frapi;

SELECT *
FROM frapi.adresse_search('8 boulevard du Port', NULL, NULL, NULL, 5);

SELECT *
FROM frapi.adresse_reverse(2.37, 48.85, 5);
```

`adresse_search_json(...)` and `adresse_reverse_json(...)` return raw JSONB responses. `adresse_search_format(jsonb)` maps a response into the `adresse_search` composite type, including a `geometry(Point,4326)` value. The convenience functions `adresse_search(...)` and `adresse_reverse(...)` combine those steps.

### Security and Service Boundary

`get_url(text, numeric, numeric, integer)` builds a shell command around `sleep` and `wget`. The SQL also concatenates query parameters without complete URL encoding. Do not pass untrusted strings: shell metacharacters or malformed parameters can change command behavior, and the upstream README explicitly says the PL/sh path has not received a security audit.

Results depend on outbound DNS, TLS, API availability, rate limits, and the provider's response schema. The upstream documentation reports testing on Ubuntu 14.04 and PostgreSQL 9.5 only. Restrict execution privileges, apply network controls, and validate current platform compatibility and provider terms before any non-laboratory use.
