## Usage

Sources:

- [README v0.1.0](https://github.com/PG-Circuit/pg-circuit/blob/v0.1.0/README.md)

`pg_circuit` inspects potentially dangerous DML and DDL on PostgreSQL 16–18. The Community edition can observe, warn or block statements. Add it to `shared_preload_libraries`, restart PostgreSQL, then create the extension as a superuser.

### Basic protection

```conf
shared_preload_libraries = 'pg_circuit'
```

```sql
CREATE EXTENSION pg_circuit;
CREATE TABLE circuit_demo (id integer);
INSERT INTO circuit_demo VALUES (1);
SET pg_circuit.mode = 'enforce';
DELETE FROM circuit_demo; -- blocked
DELETE FROM circuit_demo WHERE id = 1;
SELECT * FROM pg_circuit_status();
```

### Configuration and diagnostics

`pg_circuit.mode` defaults to warn; observe collects risk without warnings, and enforce blocks scores at or above the configured threshold. `pg_circuit_runtime_state()` reports pressure signals and `pg_circuit_events()` exposes recent events.

Community always reports effective runtime mode NORMAL. Pressure readings do not automatically escalate enforcement. This is a policy aid; application transactions, authorization and backups still determine data protection. Review rules and thresholds on representative queries before enabling enforcement.
