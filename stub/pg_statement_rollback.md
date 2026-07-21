## Usage

Sources:

- [pg_statement_rollback v1.6 README](https://github.com/lzlabs/pg_statement_rollback/blob/v1.6/README.md)
- [Changes from v1.5 to v1.6](https://github.com/lzlabs/pg_statement_rollback/compare/v1.5...v1.6)

pg_statement_rollback keeps an explicit transaction usable after a statement error by creating an automatic savepoint before each eligible statement. It emulates the statement-level rollback behavior familiar from some other databases, but the client must still issue ROLLBACK TO SAVEPOINT after an error.

The module is loaded into a backend and does not require CREATE EXTENSION.

### Load the Module

Load it for one session:

    LOAD 'pg_statement_rollback.so';

For a selected role or database, add it to session_preload_libraries and reconnect:

    session_preload_libraries = 'pg_statement_rollback'

Use shared_preload_libraries only if the deployment specifically needs server-wide loading; changing either preload list at server scope requires the corresponding restart or reconnect boundary.

### Recover from a Failed Statement

    BEGIN;
    INSERT INTO accounts(id, balance) VALUES (1, 100);
    INSERT INTO accounts(id, balance) VALUES (1, 200);
    -- duplicate-key error
    ROLLBACK TO SAVEPOINT "PgSLRAutoSvpt";
    UPDATE accounts SET balance = 150 WHERE id = 1;
    COMMIT;

The savepoint name is case-sensitive when quoted. Applications must detect the statement error and send the rollback command before continuing.

### Configuration Index

- pg_statement_rollback.enabled enables automatic savepoints for the current session.
- pg_statement_rollback.savepoint_name changes the automatic savepoint name and is superuser-controlled.
- pg_statement_rollback.enable_writeonly limits savepoint creation to statements that can write.

### Version 1.6 Behavior

Version 1.6 adds PostgreSQL 19 build support and improves detection of read-only transactions. The module no longer creates automatic savepoints in read-only transactions and releases its initial savepoint before SET TRANSACTION ... READ ONLY, which avoids interfering with dump and other read-only sessions.

### Caveats

- This is not transparent retry logic: clients must explicitly roll back to the automatic savepoint.
- Savepoints add overhead to every covered statement. Measure write-heavy workloads before enabling the module broadly.
- The upstream README warns of a crash with assertion-enabled PostgreSQL builds; do not treat development-build behavior as production-safe without testing.
- Transaction-wide errors, connection failures, and errors that invalidate the session cannot be repaired by a savepoint.
