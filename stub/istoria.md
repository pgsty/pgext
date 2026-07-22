## Usage

Sources:

- [Extension control file](https://github.com/KiriakosGeorgiou/istoria/blob/e1a21fc8e6d262c1990d4ad7fbd80fb1a75d981f/istoria.control)
- [Version 1.1 installation SQL](https://github.com/KiriakosGeorgiou/istoria/blob/e1a21fc8e6d262c1990d4ad7fbd80fb1a75d981f/istoria--1.1.sql)
- [Official trigger example](https://github.com/KiriakosGeorgiou/istoria/blob/e1a21fc8e6d262c1990d4ad7fbd80fb1a75d981f/examples/sector.sql)

`istoria` version `1.1` records row versions and provides branching, non-linear undo and redo for application tables. It is a trigger-driven SQL framework, not PostgreSQL transaction recovery.

### Core Workflow

The tracked table must have a primary key. Attach an AFTER row trigger and pass a JSON array naming the columns that define a history session; use an empty array to treat the entire table as one session.

```sql
CREATE EXTENSION istoria;

CREATE TRIGGER sector_history_tr
AFTER INSERT OR UPDATE OR DELETE ON sector
FOR EACH ROW
WHEN (pg_trigger_depth() < 1)
EXECUTE PROCEDURE history_trigger_func('["chart_id"]');

SELECT * FROM history_session_actions(1);
SELECT history_session_undo(1);
SELECT history_session_redo(1);
```

Changing data after moving to an earlier action creates another timeline. `history_session_set_active_action` moves directly to a chosen action; undo and redo replay saved row images through deletes and inserts.

### Objects and Caveats

The extension creates `history_sessions`, `history_timelines`, and `history_actions`, plus global `history_*` functions in the installation schema. Saved rows are JSONB snapshots.

Replay mutates the application table and can activate its constraints and other triggers. Several upstream functions are `SECURITY DEFINER`, use dynamic SQL, and are granted to PUBLIC, so review ownership, grants, and search paths before use. Test branching, concurrent writes, schema changes, and failed replay on a copy of real data; do not treat this old project as an audit log or backup.
