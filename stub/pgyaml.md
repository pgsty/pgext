## Usage

Sources:

- [Official upstream README](https://github.com/changtonghf/pgyaml/blob/6472c3a2b0fd7d2c6affc735f358fb6ef970ee0a/README.md)
- [Official extension control file (pgyaml.control)](https://github.com/changtonghf/pgyaml/blob/6472c3a2b0fd7d2c6affc735f358fb6ef970ee0a/pgyaml.control)
- [Official extension SQL (pgyaml--1.0.sql)](https://github.com/changtonghf/pgyaml/blob/6472c3a2b0fd7d2c6affc735f358fb6ef970ee0a/pgyaml--1.0.sql)

`pgyaml` — transform between yaml and jsonb. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgyaml;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `yaml_to_jsonb(text)` is an extension function and returns `jsonb`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
