## Usage

Sources:

- [Official upstream README](https://github.com/wuriyanto48/pgwbase64/blob/963c280ca190b23f555e199aacd8feb5a222dde8/README.md)
- [Official extension control file (wbase64.control)](https://github.com/wuriyanto48/pgwbase64/blob/963c280ca190b23f555e199aacd8feb5a222dde8/wbase64.control)
- [Official extension SQL (wbase64--0.0.1.sql)](https://github.com/wuriyanto48/pgwbase64/blob/963c280ca190b23f555e199aacd8feb5a222dde8/wbase64--0.0.1.sql)

`wbase64` — Base64 PostgreSQL Extension. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION wbase64;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgw_b64_decode(text)` is an extension function and returns `text`.
- `pgw_b64_encode(text)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
