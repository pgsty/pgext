## Usage

Sources:

- [Official README](https://github.com/VitoVan/pg_opencc/blob/main/README.md)
- [Extension SQL](https://github.com/VitoVan/pg_opencc/blob/main/pg_opencc--1.0.sql)
- [C implementation](https://github.com/VitoVan/pg_opencc/blob/main/pg_opencc.c)
- [OpenCC project](https://github.com/BYVoid/OpenCC)

`pg_opencc` 1.0 wraps OpenCC preset configurations as PostgreSQL text functions for Simplified Chinese, Traditional Chinese regional conventions, and Japanese Shinjitai. Use it for vocabulary-aware script normalization when the server's OpenCC dictionaries and presets are controlled.

### Core Workflow

Install the matching OpenCC library and preset data on the database host, create the extension, and call the desired conversion:

```sql
CREATE EXTENSION pg_opencc;

SELECT opencc_s2t('汉字');
SELECT opencc_s2twp('鼠标和软件');
SELECT opencc_tw2sp('滑鼠和軟體');
```

Conversion can be phrase-aware and is not guaranteed to round-trip; compare output against product terminology and regional editorial rules.

### Function Index

- Basic pairs: `opencc_s2t`, `opencc_t2s`.
- Taiwan and Hong Kong: `opencc_s2tw`, `opencc_tw2s`, `opencc_s2hk`, `opencc_hk2s`, `opencc_t2tw`, `opencc_tw2t`, `opencc_t2hk`, `opencc_hk2t`.
- Phrase-aware pairs: `opencc_s2twp`, `opencc_tw2sp`.
- Japanese variants: `opencc_t2jp`, `opencc_jp2t`.

Each function accepts and returns text and opens the corresponding OpenCC JSON preset for the call.

### Operational Notes

The SQL marks these functions immutable, but results depend on host-installed OpenCC libraries, configuration files, and dictionaries. An OpenCC upgrade can therefore change expression-index or stored generated-column semantics; rebuild derived values after a controlled dictionary change. The C implementation opens a converter for every call and does not free the buffer returned by `opencc_convert_utf8`, causing per-call process-heap leakage. Avoid high-volume use until that defect is fixed and verified. Upstream only reports tests with OpenCC 1.1.1/1.1.2 and PostgreSQL 13.
