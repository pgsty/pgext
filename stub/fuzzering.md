## Usage

Sources:

- [Official upstream README](https://gitlab.com/shanenoi/fuzzering/-/blob/master/README.md)
- [Official extension control file](https://gitlab.com/shanenoi/fuzzering/-/blob/master/fuzzering.control)
- [Official extension SQL](https://gitlab.com/shanenoi/fuzzering/-/blob/master/fuzzering--0.0.1.sql)

`fuzzering` — > You need to install some package for build the extension: Clang. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION fuzzering;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `fuzzering(text, text)` is an extension function and returns `float8`.
- `levenshtein(text, text)` is an extension function and returns `int32`.
- `simular(int32, int32, int32)` is an extension function and returns `float8`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
