## Usage

Sources:

- [Official extension SQL](https://github.com/ttfkam/pg_content_utils/blob/71efbcd005630f6d2c96e221c8a6f9aa859b6676/content_utils--1.0.0.sql)
- [Official extension control file](https://github.com/ttfkam/pg_content_utils/blob/71efbcd005630f6d2c96e221c8a6f9aa859b6676/content_utils.control)
- [Official README](https://github.com/ttfkam/pg_content_utils/blob/71efbcd005630f6d2c96e221c8a6f9aa859b6676/README.md)

`content_utils` provides small SQL helpers for unusual-word scoring, age-based ranking, and protocol-relative URLs. Version `1.0.0` also imports the server's system dictionary through `file_fdw`, so installation depends on a specific server-side file layout.

### Core Workflow

Install `file_fdw` first and verify that `/usr/share/dict/words` exists and is readable by PostgreSQL before creating the extension:

```sql
CREATE EXTENSION file_fdw;
CREATE EXTENSION content_utils;

SELECT count_unusual('ordinary quuxword');
SELECT match_unusual('ordinary quuxword');
SELECT reify_url(true, '//example.com/path');
```

The extension creates the `dictionary_word_files` foreign server, the `words_fdt` foreign table, and the materialized view `words`. Refresh the snapshot explicitly after the dictionary file changes:

```sql
REFRESH MATERIALIZED VIEW words;
```

### Functions and Caveats

`count_unusual(text)` counts terms absent from the dictionary, while `match_unusual(text)` reports whether such a term exists. `rank_modifier(interval)` and its timestamp overload compute a week-based score. `reify_url(boolean, varchar)` prefixes a protocol-relative URL with either HTTPS or HTTP.

The control file does not declare `file_fdw` as an automatic dependency, and installation fails when the hard-coded dictionary file is absent, unreadable, or incompatible with the expected single-column format. The materialized view is not updated automatically. Test zero-length and future intervals before relying on `rank_modifier(interval)`, because its calculation divides by a rounded number of elapsed weeks.
