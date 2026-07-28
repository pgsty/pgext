## Usage

Sources:

- [Official upstream README](https://github.com/pyramation/inflection/blob/a2d859c40fdc593e4f505f7b69a8383503f20c82/packages/inflection/readme.md)
- [Official extension control file (launchql-inflection.control)](https://github.com/pyramation/inflection/blob/a2d859c40fdc593e4f505f7b69a8383503f20c82/packages/inflection/launchql-inflection.control)
- [Official extension SQL (launchql-inflection--0.0.2.sql)](https://github.com/pyramation/inflection/blob/a2d859c40fdc593e4f505f7b69a8383503f20c82/packages/inflection/sql/launchql-inflection--0.0.2.sql)

`launchql-inflection` — inflection is a port of the functionality from Ruby on Rails' Active Support Inflection classes into PostgreSQL. Use it when SQL needs these specialized functions or aggregates. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "launchql-inflection";

select inflection.plural( 'child' );
-- children

select inflection.singular( 'children' );
-- child

select inflection.camel( 'message_properties' );
-- messageProperties

select inflection.pascal( 'web acl' );
-- WebAcl

select inflection.underscore( 'WebACL' );
-- web_acl
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `inflection.camel(str text)` is an extension function and returns `text`.
- `inflection.dashed(str text)` is an extension function and returns `text`.
- `inflection.lower(str text)` is an extension function and returns `text`.
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

- The reviewed control file declares default version `0.0.2`.
- Install the confirmed extension dependencies first: `plpgsql`, `unaccent`, `uuid-ossp`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
