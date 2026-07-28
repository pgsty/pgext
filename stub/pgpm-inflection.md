## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/inflection/README.md)
- [Official extension control file (pgpm-inflection.control)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/inflection/pgpm-inflection.control)
- [Official extension SQL (pgpm-inflection--0.30.0.sql)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/inflection/sql/pgpm-inflection--0.30.0.sql)

`pgpm-inflection` — String inflection utilities for PostgreSQL naming conventions. Use it when SQL needs these specialized functions or aggregates. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-inflection";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `inflection.camel(str text)` is an extension function and returns `text`.
- `inflection.dashed(str text)` is an extension function and returns `text`.
- `inflection.dns_1123(value text)` is an extension function and returns `text`.
- `inflection.no_consecutive_caps(str text)` is an extension function and returns `text`.
- `inflection.no_consecutive_caps_till_end(str text)` is an extension function and returns `text`.
- `inflection.no_consecutive_caps_till_lower(str text)` is an extension function and returns `text`.
- `inflection.no_single_underscores(str text)` is an extension function and returns `text`.
- `inflection.no_single_underscores_at_end(str text)` is an extension function and returns `text`.
- `inflection.no_single_underscores_in_beginning(str text)` is an extension function and returns `text`.
- `inflection.no_single_underscores_in_middle(str text)` is an extension function and returns `text`.
- `inflection.pascal(str text)` is an extension function and returns `text`.
- `inflection.pg_slugify(text)` is an extension function and returns `text`.
- `inflection.pg_slugify(value text, allow_unicode boolean)` is an extension function and returns `text`.
- `inflection.plural(str text)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `0.30.0`.
- Install the confirmed extension dependencies first: `plpgsql`, `unaccent`, `pgpm-verify`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
