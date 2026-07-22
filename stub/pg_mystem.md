## Usage

Sources:

- [Official bilingual README](https://github.com/maxoodf/pg_mystem/blob/38554744c42e15c6f10132dcb314118a39bc85f1/README.md)
- [Version 1.0.1 SQL API](https://github.com/maxoodf/pg_mystem/blob/38554744c42e15c6f10132dcb314118a39bc85f1/pg_mystem--1.0.1.sql)
- [Extension control file](https://github.com/maxoodf/pg_mystem/blob/38554744c42e15c6f10132dcb314118a39bc85f1/pg_mystem.control)

`pg_mystem` sends text to long-running Yandex Mystem worker processes and returns Russian morphological normalization from SQL. It is a bridge to the external `mystem` executable, not a built-in PostgreSQL text-search dictionary.

### Core Workflow

Install the architecture-appropriate `mystem` binary in PostgreSQL's shared-data directory. Add `pg_mystem` to `shared_preload_libraries`, reserve enough `max_worker_processes`, restart the server, and then create the extension.

```sql
CREATE EXTENSION pg_mystem;

SELECT mystem_convert('Ехал грека через реку, сунул грека руку в реку');
```

The sole SQL entry point is `mystem_convert(text) RETURNS text`. The extension declares it `STRICT`, so a NULL document returns NULL without invoking a worker.

### Configuration and Safety

`DOC_LEN_MAX` and `MYSTEM_PROCS` are compile-time Makefile settings, not runtime GUCs; changing either requires rebuilding and reinstalling the extension. Upstream recommends at least `MYSTEM_PROCS + 1` worker slots and gives only a hardware-dependent throughput estimate, so benchmark representative text rather than sizing from that estimate alone.

The PostgreSQL server account executes a separately installed third-party binary and background workers. Pin and verify that binary, constrain its filesystem permissions, account for worker failure and queue saturation, and test non-Russian text, encoding, oversized input, timeout, and restart behavior. Although the SQL declaration marks `mystem_convert()` as `IMMUTABLE`, its result depends on the external binary and dictionary version; reindex or refresh stored derived values when that dependency changes.
