## Usage

Sources:

- [docs/history.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/history.rst)
- [history/history--0.1.sql](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/history/history--0.1.sql)
- [docs/install.rst](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/docs/install.rst)
- [history/history.control](https://github.com/waveform-computing/oliphant/blob/694df2ceb61334831ad8b62d76cf532ab4a9feef/history/history.control)

`history` is Oliphant’s temporal-table extension. It creates a history table and separate triggers that retain versions of a base table at a selected time resolution.

### Core Workflow

```sql
CREATE EXTENSION auth;
CREATE EXTENSION "merge";
CREATE EXTENSION history;
CREATE TABLE history_demo (id integer PRIMARY KEY, label text);
SELECT create_history_table('history_demo', 'day');
SELECT create_history_triggers('history_demo', 'day');
```

### Operational Boundaries

These are legacy pure-SQL/PLpgSQL utilities originally documented for PostgreSQL 9.x. Current-major compatibility is unverified. Installation uses the control defaults and needs administrator privileges; no shared library or preload is required. Dependencies `auth` and `merge` must be installed. `create_history_table` copies structure and current rows; `create_history_triggers` then links subsequent writes. Use matching resolutions in both calls. A day-level history represents daily states, not every intermediate edit. The base table needs a suitable key. Historical rows use `effective` and `expiry`, and foreign keys are not copied. Generated history structures must be maintained when the base schema changes; dropping the extension is not a substitute for a history-retention plan.
