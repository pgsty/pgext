## Usage

Sources:

- [Official README for version 0.0.3](https://github.com/adjust/pg-geoip2lookup/blob/5c8144578f1c1461cc6ec0a12e18e9708de77220/Readme.md)
- [Extension SQL for version 0.0.3](https://github.com/adjust/pg-geoip2lookup/blob/5c8144578f1c1461cc6ec0a12e18e9708de77220/extension/geoip2lookup--0.0.3.sql)
- [PGXN distribution page](https://pgxn.org/dist/geoip2lookup/)

`geoip2lookup` is an untrusted PL/Perl extension that reads MaxMind MMDB files on the PostgreSQL server and returns either raw JSONB or typed geolocation records. It is useful for occasional address enrichment, but each call opens a database file and the upstream implementation is not optimized for bulk lookup.

### Configuration and Lookup

Version 0.0.3 installs into the fixed `geoip2lookup` schema and requires `plperlu` plus the Perl `MaxMind::DB::Reader` and `JSON` modules. Configure the server-side directory and default language, then call schema-qualified wrappers:

```sql
CREATE EXTENSION geoip2lookup;

SET geoip2lookup.path = '/var/lib/GeoIP';
SET geoip2lookup.language = 'en';

SELECT *
FROM geoip2lookup.city('203.0.113.10'::inet);

SELECT geoip2lookup.raw_geoip2_json(
  '203.0.113.10'::inet,
  'City'
);
```

The configured path must contain files named as expected by the code, such as `GeoIP2-City.mmdb`. Low-level overloads accept an explicit path, and city/country overloads also accept a language. Language keys are case-sensitive.

### Objects and Exact 0.0.3 Surface

- `raw_geoip2_json` returns the complete matching MMDB record as JSONB.
- `city` and `country` return localized names, ISO/continent identifiers, and GeoName identifiers.
- `isp` returns ISP, organization, and autonomous-system fields.
- `connection_type` and `anonymous_ip` return the corresponding MaxMind database fields.
- Composite types with the same names define the row layouts.

There are defects in the published 0.0.3 SQL that callers must not paper over in documentation: the one-argument anonymous-IP wrapper is created as `anonymous_id`, and the ISP row type spells one field `autononous_system_number`. Inspect `\df geoip2lookup.*` and `\dT+ geoip2lookup.*` on the installed build and use those exact identifiers.

### Security and Operations

The control file marks the extension superuser-only, non-relocatable, and dependent on untrusted `plperlu`; there is no preload requirement. Functions read paths with the PostgreSQL server process's filesystem permissions. Restrict execute rights and do not expose overloads that accept arbitrary paths to untrusted users.

MMDB data is a separately licensed, periodically updated input. Pin and audit database editions, refresh them atomically, and test missing-address behavior. For high-volume enrichment, load results into a maintained table or use a service designed for batching rather than repeatedly opening files inside queries. The project is old and lacks automated regression fixtures, so validate it on the exact PostgreSQL, Perl, module, and MMDB versions deployed.
