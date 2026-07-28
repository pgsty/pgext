## Usage

Sources:

- [Official upstream README](https://github.com/jonatas/pgscraper/blob/f17ddc9de0f00e9308948bb188c0416264c08e25/README.md)
- [Official extension control file (pgscraper.control)](https://github.com/jonatas/pgscraper/blob/f17ddc9de0f00e9308948bb188c0416264c08e25/pgscraper.control)
- [Official implementation source](https://github.com/jonatas/pgscraper/blob/f17ddc9de0f00e9308948bb188c0416264c08e25/src/lib.rs)

`pgscraper` — This is a small extension to allow you to scrape data from web directly from your postgres database. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgscraper;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `html_select` is an extension function.
- `html_select_text` is an extension function.
- `http` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
