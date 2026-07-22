## Usage

Sources:

- [Official README](https://github.com/tvondra/shared_ispell/blob/8e3f3e1022b2f45b07b0ac84481dfb68460cb9e6/README.md)
- [Extension SQL for 1.0.0](https://github.com/tvondra/shared_ispell/blob/8e3f3e1022b2f45b07b0ac84481dfb68460cb9e6/sql/shared_ispell--1.0.0.sql)
- [Extension control file](https://github.com/tvondra/shared_ispell/blob/8e3f3e1022b2f45b07b0ac84481dfb68460cb9e6/shared_ispell.control)

`shared_ispell` stores Ispell dictionaries in a fixed shared-memory segment so PostgreSQL sessions can reuse one loaded copy instead of parsing and storing a private copy per backend. It is useful for large Ispell dictionaries; Snowball dictionaries do not benefit from this extension.

### Core Workflow

Choose a shared-memory budget, preload the module, and restart PostgreSQL before creating the extension.

```conf
shared_preload_libraries = 'shared_ispell'
shared_ispell.max_size = '32MB'
```

```sql
CREATE EXTENSION shared_ispell;

CREATE TEXT SEARCH DICTIONARY czech_shared (
  TEMPLATE = shared_ispell,
  DictFile = czech,
  AffFile = czech,
  StopWords = czech
);

SELECT ts_lexize('czech_shared', 'automobile');
```

Dictionary, affix, and stop-word files follow PostgreSQL text-search file placement rules. The first use loads a dictionary into the shared segment.

### Management Objects

- `shared_ispell_mem_used()` and `shared_ispell_mem_available()` report shared-memory consumption.
- `shared_ispell_dicts()` and `shared_ispell_stoplists()` list loaded dictionaries and stop lists.
- `shared_ispell_reset()` invalidates loaded data so sessions reload changed files.
- The `shared_ispell` text-search template is the basis for custom dictionaries.

The shared-memory allocation cannot grow at runtime. Size it for every dictionary plus headroom; changing `shared_ispell.max_size` requires a restart. Upstream warns that affixes requiring full regular expressions are unsupported and fail during dictionary initialization. The project is unmaintained, so validate compatibility with the exact PostgreSQL build before production use.
