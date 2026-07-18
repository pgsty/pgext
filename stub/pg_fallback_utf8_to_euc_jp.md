## Usage

Sources:

- [Official extension control file](https://github.com/MasaoFujii/pg_fallback_utf8_to_euc_jp/blob/101376656db7619d2f12e51b5b90a67b53966ff1/pg_fallback_utf8_to_euc_jp.control)
- [Official upstream README](https://github.com/MasaoFujii/pg_fallback_utf8_to_euc_jp/blob/101376656db7619d2f12e51b5b90a67b53966ff1/README.md)

`pg_fallback_utf8_to_euc_jp` — Provides a fallback UTF-8 to EUC_JP encoding conversion that substitutes unmappable characters with a question mark.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_fallback_utf8_to_euc_jp";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_fallback_utf8_to_euc_jp';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
