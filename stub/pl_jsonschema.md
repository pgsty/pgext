## Usage

Sources:

- [README.md](https://github.com/postsql/pl-jsonschema/blob/e4f2472671ed5d1d336d8afcdf6b6cde05bb286f/README.md)
- [Makefile](https://github.com/postsql/pl-jsonschema/blob/e4f2472671ed5d1d336d8afcdf6b6cde05bb286f/Makefile)
- [pl_jsonschema.sql.in](https://github.com/postsql/pl-jsonschema/blob/e4f2472671ed5d1d336d8afcdf6b6cde05bb286f/pl_jsonschema.sql.in)
- [pl_jsonschema.control](https://github.com/postsql/pl-jsonschema/blob/e4f2472671ed5d1d336d8afcdf6b6cde05bb286f/pl_jsonschema.control)

`pl_jsonschema` installs the trusted language `pl/jsonschema`, whose function bodies are JSON Schema definitions compiled by Ajv. The hyphenated name `pl-jsonschema` is an alternative installation alias for the same implementation; do not install both.

### Core Workflow

```sql
CREATE EXTENSION pljs;
SET pl_jsonschema.engine = 'pljs';
CREATE EXTENSION pl_jsonschema;
CREATE FUNCTION valid_profile(value json) RETURNS boolean
LANGUAGE "pl/jsonschema" IMMUTABLE STRICT AS $$
{"type":"object","properties":{"name":{"type":"string"}},"required":["name"]}
$$;
SELECT valid_profile('{"name":"Alice"}'::json);
```

### Operational Boundaries

Install either `plv8` or `pljs` with JavaScript language-handler support first. With both present, selection defaults to `plv8`; set `pl_jsonschema.engine` before installation to choose explicitly. The extension control itself is not marked trusted, so creating the extension requires a superuser even though the resulting procedural language is trusted.

A validator must be a scalar function with exactly one IN argument of type json and return boolean. Procedures, set-returning functions and other argument modes are rejected. `$id` and `$ref` link schemas within the session; replacing a validator invalidates its cached schema. Use immutable, strict validators for check constraints when those semantics fit the schema.

Upstream tests PostgreSQL 19 and requires compatible JavaScript handler support; do not assume older PLV8/PLJS packages implement it. No dedicated shared library or preload is required by this SQL/JavaScript extension.
