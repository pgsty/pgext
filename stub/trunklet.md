## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/trunklet/trunklet-0.3.3/README.md)
- [Official extension control file (trunklet.control)](https://api.pgxn.org/src/trunklet/trunklet-0.3.3/trunklet.control)
- [Official extension SQL (trunklet--0.2.1--0.3.0.sql)](https://api.pgxn.org/src/trunklet/trunklet-0.3.3/sql/trunklet--0.2.1--0.3.0.sql)

`trunklet` — Be sure that you have pg_config installed and in your path. If you used a package management system such as RPM to install PostgreSQL, be sure that the -devel package is also installed. If necessary tell the build process where to find it:. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION trunklet;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `_trunklet.attnum__get(table_name regclass , field_name name)` is an extension function and returns `pg_attribute`.
- `_trunklet.exec(sql text)` is an extension function and returns `void`.
- `_trunklet.function_name(language_id _trunklet.language.language_id%TYPE , function_type text)` is an extension function and returns `text`.
- `_trunklet.language__get(language_id _trunklet.language.language_id%TYPE)` is an extension function and returns `_trunklet`.
- `_trunklet.language__get(language_name _trunklet.language.language_name%TYPE)` is an extension function and returns `_trunklet`.
- `_trunklet.language__get_id(language_name _trunklet.language.language_name%TYPE)` is an extension function and returns `_trunklet`.
- `_trunklet.language__get_loose(language_id _trunklet.language.language_id%TYPE)` is an extension function and returns `_trunklet`.
- `_trunklet.name_sanity(field_name text , value text)` is an extension function and returns `boolean`.
- `_trunklet.template__get(template_id _trunklet.template.template_id%TYPE , loose boolean DEFAULT false)` is an extension function and returns `_trunklet`.
- `_trunklet.template__get(template_name _trunklet.template.template_name%TYPE , template_version _trunklet.template.template_version%TYPE DEFAULT 1 , loose boolean DEFAULT false)` is an extension function and returns `_trunklet`.
- `_trunklet.verify_type(language_name _trunklet.language.language_name%TYPE , allowed_type regtype , supplied_type regtype , which_type text)` is an extension function and returns `void`.
- `trunklet.execute_into(template_id _trunklet.template.template_id%TYPE , parameters anyelement)` is an extension function and returns `anyelement`.
- `trunklet.execute_into(template_name _trunklet.template.template_name%TYPE , parameters anyelement)` is an extension function and returns `anyelement`.
- `trunklet.execute_into(template_name _trunklet.template.template_name%TYPE , template_version _trunklet.template.template_version%TYPE , parameters anyelement)` is an extension function and returns `anyelement`.

### Requirements and Caveats

- The reviewed control file declares default version `0.3.3`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
