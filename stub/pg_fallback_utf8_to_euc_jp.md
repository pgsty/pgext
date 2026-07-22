## Usage

Sources:

- [Official README](https://github.com/masaofujii/pg_fallback_utf8_to_euc_jp/blob/101376656db7619d2f12e51b5b90a67b53966ff1/README.md)
- [Extension SQL for 1.0](https://github.com/masaofujii/pg_fallback_utf8_to_euc_jp/blob/101376656db7619d2f12e51b5b90a67b53966ff1/pg_fallback_utf8_to_euc_jp--1.0.sql)

`pg_fallback_utf8_to_euc_jp` provides an alternative UTF-8 to EUC_JP encoding conversion. It is intended for databases that need the fallback converter to replace PostgreSQL's built-in default; installing the extension alone does not change which converter PostgreSQL selects.

### Core Workflow

Create the extension in every database where this conversion may occur, then use an administrative session to switch the default conversion flags.

```sql
CREATE EXTENSION pg_fallback_utf8_to_euc_jp;

BEGIN;
UPDATE pg_conversion
SET condefault = false
WHERE conname = 'utf8_to_euc_jp';

UPDATE pg_conversion
SET condefault = true
WHERE conname = 'pg_fallback_utf8_to_euc_jp';
COMMIT;
```

Disconnect and open a new session before relying on the new default.

### Object and Recovery Boundary

- `pg_fallback_utf8_to_euc_jp` is the conversion object created by the extension for UTF-8 to EUC_JP.

Changing `pg_conversion` is a direct system-catalog operation and requires appropriately privileged administration. Preserve exactly one default conversion for the encoding pair, and reverse the two flags before dropping the extension.

Upstream explicitly notes that `pg_dump` and `pg_dumpall` do not preserve these catalog-flag changes. After restoring a dump, create the extension where needed and repeat the default-selection transaction. This extension is relocatable and does not require preloading or a server restart.
