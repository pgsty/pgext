## Usage

Sources:

- [Official upstream README](https://github.com/winter-loo/pg-extensions/blob/d5e5ccd2a3ce5300c4b0b5f781be28ddfb2a150a/README.md)
- [Official extension control file (copy_data.control)](https://github.com/winter-loo/pg-extensions/blob/d5e5ccd2a3ce5300c4b0b5f781be28ddfb2a150a/copy_data/copy_data.control)
- [Official extension SQL (copy_data--1.0.sql)](https://github.com/winter-loo/pg-extensions/blob/d5e5ccd2a3ce5300c4b0b5f781be28ddfb2a150a/copy_data/copy_data--1.0.sql)

`copy_data` — some small postgresql extensions for beginners. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION copy_data;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `copy_data(fromtbl text, totbl text, create_table boolean DEFAULT false)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
