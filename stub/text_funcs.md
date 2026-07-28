## Usage

Sources:

- [Official upstream README](https://github.com/dannyi96/sql_custom_extensions/blob/f9d1f52fa200f8ac4b5120adc2e917831f7a84f5/README.md)
- [Official extension control file (text_funcs.control)](https://github.com/dannyi96/sql_custom_extensions/blob/f9d1f52fa200f8ac4b5120adc2e917831f7a84f5/3_text_funcs/text_funcs.control)
- [Official extension SQL (text_funcs--1.0.sql)](https://github.com/dannyi96/sql_custom_extensions/blob/f9d1f52fa200f8ac4b5120adc2e917831f7a84f5/3_text_funcs/sql/text_funcs--1.0.sql)

`text_funcs` — Tutorial 3 text functions. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION text_funcs;

SELECT hello();
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `reverse_text(text)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
