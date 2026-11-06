## Usage

Sources:

- [pg_opencc 0.1.0 README](https://github.com/pgsty/pg_opencc/blob/v0.1.0/README.md)
- [Extension control file](https://github.com/pgsty/pg_opencc/blob/v0.1.0/opencc.control)
- [SQL API](https://github.com/pgsty/pg_opencc/blob/v0.1.0/sql/opencc--0.1.0.sql)
- [Release notes](https://github.com/pgsty/pg_opencc/releases/tag/v0.1.0)

`opencc`, distributed as `pg_opencc`, converts Simplified Chinese, Traditional Chinese, Taiwan and Hong Kong conventions, and Japanese Shinjitai. Version 0.1.0 embeds OpenCC 1.4.2 and its dictionaries in the shared library; no system OpenCC or external dictionary files are required at runtime.

### Enable and convert

Supports PostgreSQL 14–18 and requires a `UTF8` database. This trusted extension can be installed by a user with `CREATE` privilege on the database. Objects use the fixed `opencc` schema; the extension is not relocatable. Preloading and a server restart are not required.

```sql
CREATE EXTENSION opencc;

SELECT opencc.s2twp('鼠标里面的硅二极管坏了，导致光标分辨率降低。');
SELECT opencc.convert('繁體中文', 't2s');
SELECT * FROM opencc.configs;
```

`opencc.convert(input, config)` selects a conversion explicitly. Wrappers such as `opencc.s2t()` and `opencc.t2s()` provide one function per configuration:

| Configurations | Conversion |
| --- | --- |
| `s2t`, `t2s` | Simplified and Traditional Chinese |
| `s2tw`, `tw2s`, `s2hk`, `hk2s` | Taiwan or Hong Kong character conventions |
| `s2twp`, `tw2sp` | Taiwan regional phrases |
| `s2hkp`, `hk2sp` | Hong Kong regional phrases; developing |
| `t2tw`, `tw2t`, `t2hk`, `hk2t` | Traditional Chinese variant conversion |
| `t2jp`, `jp2t` | Japanese old/new character forms; experimental |

### Regional conversion and inspection

`opencc.localize(input, target, source)` accepts `zh-Hans`, `zh-Hant`, `zh-CN`, `zh-TW`, `zh-HK`, and `ja`. Its default `source => 'auto'` detects the script, but cannot distinguish Taiwan from Hong Kong vocabulary. Supply the source region when regional phrases matter.

```sql
SELECT opencc.localize('滑鼠裡面的矽二極體壞了', 'zh-CN', source => 'zh-TW');
SELECT opencc.detect('中文'), opencc.script_stats('中文');
SELECT * FROM opencc.routes;
SELECT * FROM opencc.ambiguities('大战文丑的时候，他的头发很干燥', 's2t');
SELECT opencc.candidates('发', 's2t');
SELECT * FROM opencc.explain('数据库连接池', 's2twp');
```

| Object | Purpose |
| --- | --- |
| `opencc.detect()` | Classify the script, including mixed or neutral text |
| `opencc.script_stats()` | Count Han, simplified, traditional, and kana characters |
| `opencc.configs`, `opencc.routes` | Inspect configurations and regional conversion chains |
| `opencc.ambiguities()`, `opencc.candidates()` | Inspect ambiguous spans and dictionary candidates |
| `opencc.explain()` | Inspect each conversion stage |
| `opencc.version()` | Report extension, engine, and dictionary digest |
| `opencc.dependents()` | Find dependent database objects and suggested rebuild statements |

### Operation and upgrades

Conversion functions are `IMMUTABLE STRICT PARALLEL SAFE` within an extension version. They can be used in expression indexes and stored generated columns. PostgreSQL 18 generated columns must explicitly use `STORED` for user-defined conversion functions.

Dictionary changes in a future extension release can change conversion results. Inspect `opencc.version()` and review the statements returned by `opencc.dependents()` when upgrading. Dependency discovery does not find calls hidden inside other user-defined functions.

```sql
SELECT * FROM opencc.version();
SELECT * FROM opencc.dependents();
```

Each backend initializes converters on demand. Optional startup preloading can share their memory across backends; changing `shared_preload_libraries` requires a restart:

```ini
shared_preload_libraries = 'opencc'
opencc.preload = 's2t,t2s,s2twp,tw2sp'
```

Merge `opencc` into any existing preload list. `opencc.preload` also accepts `all`. Dropping the extension leaves its empty schema behind.
