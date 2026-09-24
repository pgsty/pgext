## Usage

Sources:

- [PGXN 0.1.0](https://pgxn.org/dist/macavity/0.1.0/)

`macavity` provides deterministic, session-local fault injection for PostgreSQL 16–18. This testing release is intended only for disposable test clusters. Creating the extension requires a superuser; no shared preload is needed.

### Fault injection

```sql
CREATE EXTENSION macavity;
SELECT * FROM macavity_points();
SELECT macavity_arm('executor_start', 'error', 1);
SELECT 1; -- expected injected error
SELECT * FROM macavity_status();
SELECT macavity_disarm();
```

An armed fault fires at its specified occurrence and is then spent. The counters remain available after an injected error. Each session can arm one fault; a new session starts unarmed.

### Actions and boundaries

The points are `executor_start`, `executor_end`, `before_commit` and `before_abort`. Actions are `error`, `delay` and `crash`; delay is one second. Error injection during abort processing is rejected.

The crash action kills the calling backend. PostgreSQL then disconnects other sessions and performs crash recovery. Never enable this extension on a cluster containing important data. Fault counters are session-local and are lost when that backend exits.
