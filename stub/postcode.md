## Usage

Sources:

- [README.md](https://github.com/John-Cameron-Burns/pg-postcode/blob/40f86f6567a25d2ef8beb3bfe80230b923d691a3/README.md)
- [postcode.control](https://github.com/John-Cameron-Burns/pg-postcode/blob/40f86f6567a25d2ef8beb3bfe80230b923d691a3/postcode.control)
- [postcode--2.1.0.sql](https://github.com/John-Cameron-Burns/pg-postcode/blob/40f86f6567a25d2ef8beb3bfe80230b923d691a3/postcode--2.1.0.sql)
- [META.json](https://github.com/John-Cameron-Burns/pg-postcode/blob/40f86f6567a25d2ef8beb3bfe80230b923d691a3/META.json)
- [.github/workflows/test.yml](https://github.com/John-Cameron-Burns/pg-postcode/blob/40f86f6567a25d2ef8beb3bfe80230b923d691a3/.github/workflows/test.yml)
- [CHANGELOG.md](https://github.com/John-Cameron-Burns/pg-postcode/blob/40f86f6567a25d2ef8beb3bfe80230b923d691a3/CHANGELOG.md)

`postcode` 2.1.0 is the maintained PGXN continuation of the original patchsoft project. It retains the compact UK `postcode` and delivery-point-suffix `dps` types and adds the 64-bit international `postal_code` type. The country-qualified text form distinguishes otherwise ambiguous national codes.

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

Version 2.1.0 adds `part`, `parts` and `prefix_of` to both postal types. Cast string literals explicitly because those overloads are ambiguous for untyped text. `postal_code_parts` lists available names. `to_char(postcode, text)` is deprecated but remains in 2.x. This release also fixes a use-after-free when compiling very large custom patterns; built-in patterns are not affected. Run `ALTER EXTENSION postcode UPDATE` after installing the new library; the 2.1 upgrade preserves existing stored values.
