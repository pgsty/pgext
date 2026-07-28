## Usage

Sources:

- [Official upstream README](https://github.com/rom8726/fuzzysearch/blob/b61c645cd0916ca4b7d077e4c5e4e73abcadaa5f/README.md)
- [Official extension control file (fuzzysearch.control)](https://github.com/rom8726/fuzzysearch/blob/b61c645cd0916ca4b7d077e4c5e4e73abcadaa5f/fuzzysearch.control)
- [Official extension SQL (fuzzysearch--1.0.sql)](https://github.com/rom8726/fuzzysearch/blob/b61c645cd0916ca4b7d077e4c5e4e73abcadaa5f/fuzzysearch--1.0.sql)

`fuzzysearch` — PostgreSQL extension for strings fuzzy matching. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION fuzzysearch;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `jaro_winkler_match(text, text)` is an extension function and returns `float`.
- `levenshtein_match(text, text)` is an extension function and returns `integer`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
