## Usage

Sources:

- [1.0.2 release](https://github.com/bitsinfra/gnomon/releases/tag/v1.0.2)
- [Official source](https://github.com/bitsinfra/gnomon/blob/v1.0.2/README.md)
- [Official source](https://github.com/bitsinfra/gnomon/blob/v1.0.2/gnomon.control)
- [Official source](https://github.com/bitsinfra/gnomon/blob/v1.0.2/LICENSE)
- [Official source](https://github.com/bitsinfra/gnomon/blob/v1.0.2/Makefile)
- [Official source](https://github.com/bitsinfra/gnomon/blob/v1.0.2/docs/installation.md)
- [Official source](https://github.com/bitsinfra/gnomon/blob/v1.0.2/docs/security.md)
- [Official source](https://github.com/bitsinfra/gnomon/blob/v1.0.2/docs/operations.md)
- [Official source](https://github.com/bitsinfra/gnomon/releases/download/v1.0.2/gnomon-pgtle--1.0.2.sql)

gnomon 1.0.2 is a SQL extension for table history, temporal reads, branches and reviewed three-way merges. It supports PostgreSQL 15–18. The optional gnomon_c companion accelerates and extends routing; the SQL core works independently.

### Start tracking

```sql
CREATE EXTENSION gnomon;
SELECT gnomon.track('public.orders');
SELECT gnomon.coordinate_text();
```

The example assumes public.orders already exists. Tracking takes a brief lock and installs storage/views/triggers, with no historical backfill. Installation needs a superuser and creates cluster-wide roles. Branch selection and routing change the session's effective view of tracked tables; follow the official branch/use workflow rather than writing private storage directly.

### Privileges and durability

Never grant internal gnomon_reader, gnomon_reader_s or gnomon_search roles to application logins. gnomon_merger and gnomon_admin carry BYPASSRLS and are privileged operational roles. Table grants, branch rights and reviewed merge permissions are separate. Merges detect conflicts; undo is another history change and can refuse later edits.

Exact durable replay requires synchronous_commit enabled. Unlogged scratch branches are not replicated and do not survive crashes like logged history. Retention bounds the history that can be read. Back up application tables and Gnomon storage together.

### Time features and upgrades

Run gnomon.heartbeat() approximately every second using a supervised worker, pg_cron or the CLI for time resolution, scheduling and retention. Without it branches/history still work but time features drift. Where event triggers are unavailable, run gnomon.sync() after DDL. For an existing 1.0.0 or 1.0.1 installation, use the documented upgrade to 1.0.2 and retain its data. Server extensions, pg_tle and plain-SQL installations have separate upgrade paths; select the script matching the installed version.

### 1.0.2 correction

The SQL upgrade repairs read sets recorded by 1.0.1 when a branch predates table tracking. Merges now restore the session's selected branch. The C companion rejects system-column reads on branches instead of crashing. Update the SQL core to 1.0.2 first, then the matching C companion if installed; restart processes using a replaced preloaded library.
