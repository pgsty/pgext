## Usage

Sources:

- [Extension and client documentation](https://github.com/capraconsulting/skjera/blob/67da21683b0c4d38449c2230178afa64b1f875f0/packages/snapsync/README.md)
- [Extension control](https://github.com/capraconsulting/skjera/blob/67da21683b0c4d38449c2230178afa64b1f875f0/packages/snapsync/extension/snapsync.control)
- [Versioned SQL](https://github.com/capraconsulting/skjera/blob/67da21683b0c4d38449c2230178afa64b1f875f0/packages/snapsync/extension/snapsync--0.1.0.sql)

`snapsync` is a pure SQL extension from the Skjera repository for polling table changes using transaction snapshots. The control version is `0.1.0`; upstream documents PostgreSQL 14 or later. It uses the fixed `snapsync` schema and is not relocatable. It requires no shared library, preload or restart.

### Track a table

After installing the extension files, enable it and track a table as its owner:

```sql
CREATE EXTENSION snapsync;

CREATE TABLE attendees (id uuid PRIMARY KEY, display_name text);
SELECT snapsync.track('attendees');
```

Tracking adds an `xid` column, an index and triggers. Updates stamp the writing transaction ID; deletes create tombstones. One key column is supported, with keys represented as text. Composite keys are unsupported. The installation declares `superuser = false`, but the installing role still needs the database and object privileges used by the SQL script.

### Pull changes consistently

Each page and its row lookup must use the same `REPEATABLE READ` transaction:

```sql
BEGIN ISOLATION LEVEL REPEATABLE READ;
SELECT changed, deleted, next_token, complete, more
FROM snapsync.changes('attendees', NULL, 1000);
COMMIT;
```

For this initial page, pass a null token. Within the transaction, retrieve rows whose keys appear in `changed` using the reader's normal access rules. Apply the returned `deleted` keys to the client copy. Persist `next_token` only after applying the page; request another page when `more` is true.

A true `complete` means the old token was unusable and the client is rebuilding its full copy. Finish all pages before replacing that copy. Concurrent commits may cause repeat delivery; consumers must tolerate duplicate rows.

### Permissions and maintenance

Tracking and maintenance functions have `PUBLIC` execution revoked. Tombstones use row-level security based on the reader's table-level `SELECT` privilege; application authorization and row filtering still apply when loading changed rows. If visibility depends on another table, updates there must also touch the tracked rows.

`snapsync.prune(interval)` removes old tombstones. A client that needs pruned history must start over. Keep retention longer than long-running transactions. After a logical restore, call `snapsync.ensure_current()` before serving pulls so restored transaction IDs cannot silently reuse the old sync history. Truncation also forces a fresh copy. `snapsync.untrack(regclass)` removes tracking but leaves the transaction-ID column.
