## Usage

Sources:

- [Official README](https://github.com/pgbigm/ludia_funcs/blob/d220429554fbc1f04573f444a58bc5ab8919d4a6/README.md)
- [Official extension SQL](https://github.com/pgbigm/ludia_funcs/blob/d220429554fbc1f04573f444a58bc5ab8919d4a6/ludia_funcs--1.0.sql)
- [Official C implementation](https://github.com/pgbigm/ludia_funcs/blob/d220429554fbc1f04573f444a58bc5ab8919d4a6/ludia_funcs.c)

`ludia_funcs` supplies Ludia-compatible text normalization and snippet helpers for PostgreSQL. It is mainly useful when porting applications that expect Ludia/Senna behavior or combining the helpers with a full-text search extension such as `pg_bigm`.

### Core Workflow

```sql
CREATE EXTENSION ludia_funcs;

SELECT pgs2norm('ｱﾌﾟﾘｹｰｼｮﾝ');
SELECT pgs2snippet1(1, 32, 1, '★', '★', 0,
                    'ｱﾌﾟﾘｹｰｼｮﾝ', '...');
SELECT * FROM pgs2seninfo();
```

The installed SQL surface is compact:

- `pgs2norm(text)` normalizes text using Senna.
- `pgs2snippet1(...)` produces highlighted snippets with caller-supplied width, result count, tags, keyword count, source text, and keyword arguments.
- `pgs2seninfo()` returns the linked Senna version and configure options.
- `pgs2textporter1(text)` invokes the optional TextPorter integration when the extension was compiled with that support.

### Configuration and Boundaries

`ludia_funcs` initializes Senna when its library is first loaded; the control file does not require `shared_preload_libraries`. Its optional GUCs include `ludia_funcs.norm_cache_limit` (negative means use `work_mem`, zero means unlimited) and `ludia_funcs.escape_snippet_keyword`. TextPorter builds can expose additional settings.

The primary functions expect a UTF-8 database. `pgs2textporter1()` operates on server-side application files and temporary files, so it crosses a filesystem and privilege boundary; do not pass untrusted paths, and do not assume TextPorter was compiled into a particular package. The extension declares no SQL dependency on `pg_bigm`; that integration is optional.
