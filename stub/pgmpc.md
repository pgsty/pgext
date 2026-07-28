## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pgmpc/pgmpc-0.1.0/README)
- [Official extension control file (pgmpc.control)](https://api.pgxn.org/src/pgmpc/pgmpc-0.1.0/pgmpc.control)
- [Official extension SQL (pgmpc--0.1.0.sql)](https://api.pgxn.org/src/pgmpc/pgmpc-0.1.0/pgmpc--0.1.0.sql)

`pgmpc` — pgmpc, mpd client for PostgreSQL ================================. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgmpc;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `mpd_add(path text)` is an extension function and returns `void`.
- `mpd_clear()` is an extension function and returns `void`.
- `mpd_consume()` is an extension function.
- `mpd_load(path text)` is an extension function and returns `void`.
- `mpd_ls(path text, OUT song text)` is an extension function and returns `SETOF`.
- `mpd_lsplaylists(OUT playlist text)` is an extension function and returns `SETOF`.
- `mpd_next()` is an extension function and returns `void`.
- `mpd_pause()` is an extension function.
- `mpd_play()` is an extension function and returns `void`.
- `mpd_playlist(OUT song text)` is an extension function and returns `SETOF`.
- `mpd_playlist(path text, OUT song text)` is an extension function and returns `SETOF`.
- `mpd_prev()` is an extension function and returns `void`.
- `mpd_random()` is an extension function.
- `mpd_repeat()` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
