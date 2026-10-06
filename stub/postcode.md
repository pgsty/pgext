## Usage

Sources:

- [README.md](https://github.com/ztz88f5f6h-glitch/pg-uk-postcodes/blob/6570ec483670afc164d08fb5a5bcc4ae6f56ca5a/README.md)
- [postcode.control](https://github.com/ztz88f5f6h-glitch/pg-uk-postcodes/blob/6570ec483670afc164d08fb5a5bcc4ae6f56ca5a/postcode.control)
- [postcode--2.0.1.sql](https://github.com/ztz88f5f6h-glitch/pg-uk-postcodes/blob/6570ec483670afc164d08fb5a5bcc4ae6f56ca5a/postcode--2.0.1.sql)
- [META.json](https://github.com/ztz88f5f6h-glitch/pg-uk-postcodes/blob/6570ec483670afc164d08fb5a5bcc4ae6f56ca5a/META.json)
- [.github/workflows/test.yml](https://github.com/ztz88f5f6h-glitch/pg-uk-postcodes/blob/6570ec483670afc164d08fb5a5bcc4ae6f56ca5a/.github/workflows/test.yml)
- [CHANGELOG.md](https://github.com/ztz88f5f6h-glitch/pg-uk-postcodes/blob/6570ec483670afc164d08fb5a5bcc4ae6f56ca5a/CHANGELOG.md)

`postcode` 2.0.1 is the maintained PGXN continuation of the original patchsoft project. It retains the compact UK `postcode` and delivery-point-suffix `dps` types and adds the 64-bit international `postal_code` type. The country-qualified text form distinguishes otherwise ambiguous national codes.

### Core Workflow

```sql
CREATE EXTENSION postcode;
SELECT 'SW1A 1AA'::postcode;
SELECT 'US-90210-1234'::postal_code, 'CA-K1A 0B1'::postal_code;
CREATE TABLE addresses (code postcode);
INSERT INTO addresses VALUES ('SW1A 1AA'), ('LS2 4AA');
SELECT * FROM addresses
WHERE code >= range_lower('SW1A') AND code < range_upper('SW1A');
```

### Operational Boundaries

`postcode` supports formatting, validation and B-tree indexing. Its `%` prefix filter can use planner support for constant fragments; `range_lower` and `range_upper` express explicit indexed bounds. Earlier 1.3.0 incorrectly registered prefix matching as B-tree equality, and the maintained series fixes that as well as a heap-overflow defect. Replace the library and run `ALTER EXTENSION postcode UPDATE`; merely replacing binaries does not repair SQL operator metadata.

For international values, `postal_code` requires a country code, such as the examples above. `country` extracts it; `postal_prefix`, `lower_bound` and `upper_bound` support prefixes. `is_valid_postal_code` and `to_postal_code` help ingest dirty input, and a country typmod can constrain a column. Format validation is not proof of allocation or address existence.

The release tests PostgreSQL 14–18. Installation is superuser-only, the extension is relocatable and no preload is required. Existing UK types and stored values are retained by the additive 2.0 upgrade.
