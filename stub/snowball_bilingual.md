## Usage

Sources:

- [Official upstream README](https://github.com/thomasweiser/mediatum-view/blob/27310339c24cf204325dc9d3ecdd2efdb539236d/backend/pg-extensions/snowball_bilingual/README.md)
- [Official extension control file (snowball_bilingual.control)](https://github.com/thomasweiser/mediatum-view/blob/27310339c24cf204325dc9d3ecdd2efdb539236d/backend/pg-extensions/snowball_bilingual/snowball_bilingual.control)
- [Official extension SQL (snowball_bilingual--1.0.sql)](https://github.com/thomasweiser/mediatum-view/blob/27310339c24cf204325dc9d3ecdd2efdb539236d/backend/pg-extensions/snowball_bilingual/snowball_bilingual--1.0.sql)

`snowball_bilingual` — The snowball_bilingual extension provides a new dictionary template. It is a copy of the Snowball dictionary template of PostgreSQL. Currently it provides stemming algorithms for the following languages: nepali. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION snowball_bilingual;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `dsnowball_bilingual_init(INTERNAL)` is an extension function and returns `INTERNAL`.
- `dsnowball_bilingual_lexize(INTERNAL, INTERNAL, INTERNAL, INTERNAL)` is an extension function and returns `INTERNAL`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
