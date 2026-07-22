## Usage

Sources:

- [Extension control file](https://github.com/df7cb/postgresql-hamradio/blob/4e2181d68b62d224b120ce1bb38d2c23ee63e937/hamradio.control)
- [Extension SQL files](https://github.com/df7cb/postgresql-hamradio/tree/4e2181d68b62d224b120ce1bb38d2c23ee63e937/sql)
- [Official example SQL](https://github.com/df7cb/postgresql-hamradio/blob/4e2181d68b62d224b120ce1bb38d2c23ee63e937/examples.sql)

hamradio provides PostgreSQL domains, enums, conversion functions, and reference tables for amateur-radio callsigns, bands, frequencies, modes, and Maidenhead locators. It is useful for validating logbook data and deriving geographic cells, but it is a small domain toolkit rather than a complete logging application.

### Core Workflow

The extension depends on `uctext` and `postgis`. Store validated callsigns, frequencies, and locators, then derive a radio band or PostGIS geometry:

```sql
CREATE EXTENSION uctext;
CREATE EXTENSION postgis;
CREATE EXTENSION hamradio;

CREATE TABLE contacts (
  worked_call call NOT NULL,
  frequency qrg NOT NULL,
  grid locator
);

INSERT INTO contacts VALUES ('DL1ABC/P', 14.074, 'JO62QM');

SELECT worked_call,
       band(frequency),
       ST_AsText(ST_LocatorPoint(grid))
FROM contacts;
```

### Important Objects

- `call` validates upper-case callsign components separated by slash or hyphen.
- `band` enumerates amateur bands from 2190m through 1mm; `band(numeric)` maps a frequency in MHz to a band.
- `qrg` accepts only frequencies that map to a known band.
- `locator` validates 2- to 10-character Maidenhead locators.
- `ST_Locator` returns the locator-cell polygon in SRID 4326, while `ST_LocatorPoint` returns its point representation.
- `locator2` and `locator4` are prepopulated reference tables with GiST-indexed PostGIS cells.
- `major_mode` and `adif_mode` normalize common mode labels.

### Operational Notes

The validation rules and band limits are frozen in the extension SQL and may not reflect later regulatory changes or local band plans. The reference tables add precomputed global locator cells at install time. Check the derived band and geometry semantics for the target application, and preserve the exact upper-case formats expected by the domains.
