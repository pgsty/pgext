## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/moritetu/walreader/blob/f54a68f9f925da149a2b33030405cfeaea8d6197/README.md)
- [Version 1.0 SQL objects](https://github.com/moritetu/walreader/blob/f54a68f9f925da149a2b33030405cfeaea8d6197/walreader--1.0.sql)
- [WAL reader implementation](https://github.com/moritetu/walreader/blob/f54a68f9f925da149a2b33030405cfeaea8d6197/walreader.c)

`walreader` exposes PostgreSQL WAL record headers and resource-manager descriptions through SQL. It is an educational implementation modeled on `pg_waldump`; upstream recommends the standard tool as lighter and more capable.

```sql
CREATE EXTENSION walreader;
SET walreader.read_limit = 100;
SELECT timeline, walseg, lsn, rmgr, identify, rmgr_desc
FROM read_wal_segment(pg_walfile_name(pg_current_wal_lsn()))
LIMIT 20;
```

`walreader.default_wal_directory` can point a session at another server directory, and functions can read by segment name or LSN. WAL contains sensitive operational and data-change information. Revoke all access from `PUBLIC`, restrict use to a dedicated superuser diagnostic role, and do not accept caller-controlled paths.

WAL formats and internal decoders are major-version-specific. Never parse segments with a mismatched binary. Reading an active final record can legitimately report an incomplete record. Set strict record and statement limits, copy archived segments for analysis where possible, and test corrupted/truncated input without risking the primary. Prefer `pg_waldump` outside the server process to isolate crashes and resource use.
