## Usage

Sources:

- [append_only_heap README at the reviewed commit](https://github.com/main--/append_only_heap/blob/9dc0bc189ac378ea146941266928ee44a4a0b793/README.md)
- [append_only_heap.control at the reviewed commit](https://github.com/main--/append_only_heap/blob/9dc0bc189ac378ea146941266928ee44a4a0b793/append_only_heap.control)
- [Version 1.0 install SQL](https://github.com/main--/append_only_heap/blob/9dc0bc189ac378ea146941266928ee44a4a0b793/append_only_heap--1.0.sql)
- [Table-access-method implementation](https://github.com/main--/append_only_heap/blob/9dc0bc189ac378ea146941266928ee44a4a0b793/append_only_heap.c)

`append_only_heap` is a table access method that reuses PostgreSQL heap behavior but adds `HEAP_INSERT_SKIP_FSM` to inserts. New tuples therefore avoid free-space-map holes, which can reduce fragmentation for variable-length workloads that truly append data.

### Create a Trial Table

```sql
CREATE EXTENSION append_only_heap;
CREATE TABLE event_log (
  event_id bigint GENERATED ALWAYS AS IDENTITY,
  payload jsonb NOT NULL
) USING append_only_heap;

INSERT INTO event_log(payload) VALUES ('{"event":"created"}');
SELECT * FROM event_log;
```

### Caveats

- The name does not enforce immutability: `UPDATE`, `DELETE`, and `TRUNCATE` remain inherited heap operations. Enforce an append-only policy separately.
- The version `1.0` implementation calls `mprotect` and overwrites function pointers in PostgreSQL's otherwise constant process-global `TableAmRoutine`. That unsupported internal hook can fail to load, conflict with other modules, or break after a server change.
- Skipping free-space reuse increases relation growth when rows are updated or deleted. Treat this as experimental storage code, benchmark bloat and vacuum behavior, and test crash recovery and upgrades on disposable data.
