## Usage

Sources:

- [database.dev package page](https://database.dev/stateless/pg_once)

`pg_once` 0.1.0 is a pure-SQL idempotency-key helper published through database.dev. It reserves a key for a named resource, records completion or failure, optionally caches a JSON response, and expires records using a time-to-live value whose default is 24 hours.

### Install and reserve a key

Install the package through a configured dbdev registry, then create the extension:

```sql
SELECT dbdev.install('pg_once');
CREATE EXTENSION pg_once VERSION '0.1.0';

SELECT pg_once_start('order#42', 'orders');
SELECT pg_once_commit('order#42', '{"ok": true}'::jsonb);
```

The application must inspect the JSON result from `pg_once_start` and perform the protected work only when the key is accepted. Call `pg_once_commit` only after that work succeeds; cache only response data that is safe for every retried caller to receive.

For a failed operation or periodic expiry cleanup:

```sql
SELECT pg_once_fail('order#42');
SELECT pg_once_cleanup();
```

### Transaction and lifecycle boundary

`pg_once_start`, the business operation, and `pg_once_commit` form an application protocol. Keep their transaction boundaries deliberate: committing the reservation before the business write can leave an in-progress key after a crash, while committing a response separately can diverge from the protected write. Design retries for each intermediate state and choose a TTL longer than the longest legitimate operation.

Expiry is not proof that an earlier external side effect never happened. `pg_once_cleanup` deletes expired state, so schedule it only after establishing retention and retry requirements. The registry page is the authoritative public artifact for this package; review the installed SQL in the generated migration before granting untrusted roles access to its functions or tables.
