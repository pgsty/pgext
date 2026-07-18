## Usage

Sources:

- [Extension control file](https://github.com/benchub/flux/blob/836e16f6631907c1b176586e960c460e1d225140/flux.control)
- [Usage guide](https://github.com/benchub/flux/blob/836e16f6631907c1b176586e960c460e1d225140/docs/howto.md)
- [Documented limitations](https://github.com/benchub/flux/blob/836e16f6631907c1b176586e960c460e1d225140/docs/limitations.md)
- [SQL implementation](https://github.com/benchub/flux/blob/836e16f6631907c1b176586e960c460e1d225140/flux--0.3.sql)

`flux` is a pure-SQL, trigger-based row-history extension. For each enabled table it records reverse changes for `INSERT`, `UPDATE`, and `DELETE`, allowing a row to be reconstructed at an earlier timestamp or its complete history to be returned. It requires `hstore`, a primary key, and the fixed `_flux` schema.

### Enable and query history

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION flux;

CREATE TABLE public.accounts (
  id bigint PRIMARY KEY,
  status text NOT NULL
);

SELECT _flux.enable_change_logging(
  'public', 'accounts', '_accounts_history'
);

INSERT INTO public.accounts VALUES (1, 'new');
UPDATE public.accounts SET status = 'active' WHERE id = 1;

SELECT *
FROM _flux.get_row_history('public', 'accounts', '"id"=>"1"'::hstore);
```

`_flux.enable_change_logging()` also supports `include` or `exclude` column lists and an optional retention interval. `_flux.get_row_from_history()` reconstructs one row in the table's current composite type. `_flux.disable_change_logging()` removes the triggers, while `_flux.cleanup()` later drops disabled log tables and applies retention.

Flux stores reverse diffs rather than an initial snapshot. Consequently, `TRUNCATE` loses the current state needed to reconstruct live rows. Primary-key value changes are rejected; renaming a tracked table breaks its metadata; changing primary keys or column types requires a carefully planned disable/cleanup/re-enable cycle; and dropped columns disappear from typed historical reconstruction. Disabling needs an exclusive table lock, while cleanup is destructive to the selected history tables. Version `0.3` is old and has no current compatibility matrix, so validate locking, storage growth, schema migration, and restore behavior before operational use.
