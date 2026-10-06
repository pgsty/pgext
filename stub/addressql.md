## Usage

Sources:

- [extensions/addressql-postgres/README.md](https://github.com/veygrit-sys/AGID-Global-Address-Infrastructure/blob/f2310ffa6750d1f0f2a8e3a903f0a41791da3ff9/extensions/addressql-postgres/README.md)
- [extensions/addressql-postgres/sql/addressql--0.1.0.sql](https://github.com/veygrit-sys/AGID-Global-Address-Infrastructure/blob/f2310ffa6750d1f0f2a8e3a903f0a41791da3ff9/extensions/addressql-postgres/sql/addressql--0.1.0.sql)
- [LICENSE_POLICY.md](https://github.com/veygrit-sys/AGID-Global-Address-Infrastructure/blob/f2310ffa6750d1f0f2a8e3a903f0a41791da3ff9/LICENSE_POLICY.md)
- [extensions/addressql-postgres/addressql.control](https://github.com/veygrit-sys/AGID-Global-Address-Infrastructure/blob/f2310ffa6750d1f0f2a8e3a903f0a41791da3ff9/extensions/addressql-postgres/addressql.control)

`addressql` 0.1.0 is an executable SQL/PLpgSQL scaffold for country metadata, postal formats, address normalization and matching. Its seed data is synthetic; results do not establish real delivery coverage or address identity.

### Core Workflow

```sql
CREATE EXTENSION addressql;
SELECT addressql.address_parse('1 Main Street', 'US');
SELECT addressql.country_address_profile('US');
SELECT addressql.postgis_available();
```

### Operational Boundaries

Functions in the fixed `addressql` schema return JSONB with source versions, confidence, warnings and scope limitations. Tables hold country profiles, postal areas, delivery regions and synthetic examples; load the intended source dataset before expecting useful lookup results.

`addressql.address_parse`, `addressql.address_normalize`, `addressql.country_address_profile` and postal helpers provide the core metadata workflow. Spatial helpers use optional `postgis`; without it they return explicit fallback warnings. The control is trusted and the implementation has no dedicated shared library or preload requirement, despite an unused library path in the control file. No major-version matrix is declared. Do not interpret prototype normalization as verified deliverability or privacy protection.
