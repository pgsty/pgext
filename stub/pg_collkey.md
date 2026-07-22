## Usage

Sources:

- [Official README](https://github.com/ccutrer/pg_collkey/blob/42d93fd5182ca94efbc45d7c02b86d95e8e2070a/README)
- [Extension control file](https://github.com/ccutrer/pg_collkey/blob/42d93fd5182ca94efbc45d7c02b86d95e8e2070a/pg_collkey.control)
- [SQL function definitions](https://github.com/ccutrer/pg_collkey/blob/42d93fd5182ca94efbc45d7c02b86d95e8e2070a/collkey_icu.sql)

pg_collkey wraps ICU collation-key generation so a query can choose a locale and comparison strength independently of the database's default collation. It is useful for producing sortable byte keys, including accent-insensitive or numeric-aware ordering, but the old upstream code supports only UTF-8 databases.

### Core Workflow

Generate the key in `ORDER BY` or persist it in an expression index when the locale and options are stable:

```sql
CREATE EXTENSION pg_collkey;

SELECT name
FROM product
ORDER BY collkey(name, 'fr_FR');

CREATE INDEX product_name_root_idx
ON product (collkey(name, 'root'));
```

The two-argument overload uses shifted punctuation handling, the default strength, and numeric sorting. The one-argument overload also selects the ICU root collation.

### Important Objects

- `collkey(text, text, boolean, integer, boolean)` returns an ICU sort key as `bytea`.
- The locale argument selects ICU root or a named locale such as `fr_FR`.
- The first boolean shifts punctuation to the fourth comparison level.
- Strength 1 compares base characters; 2 includes accents; 3 includes case; 4 includes other differences; 5 compares normalized code points.
- The final boolean enables numeric ordering, so digit sequences sort by numeric value.
- `collkey(text, text)` and `collkey(text)` are convenience overloads.

### Operational Notes

Sort keys are derived from the linked ICU library. ICU upgrades or locale-data changes can invalidate ordering assumptions and expression indexes, so reindex and regression-test results after such upgrades. Using many locales within one query adds setup cost. The README documents very old PostgreSQL and ICU baselines, and non-UTF-8 databases are unsupported.
