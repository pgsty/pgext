## Usage

Sources:

- [PGXN 0.1.0](https://pgxn.org/dist/pg_lease/0.1.0/)
- [Official README](https://api.pgxn.org/src/pg_lease/pg_lease-0.1.0/README.md)
- [Extension SQL](https://api.pgxn.org/src/pg_lease/pg_lease-0.1.0/sql/pg_lease--0.1.0.sql)
- [Control file](https://api.pgxn.org/src/pg_lease/pg_lease-0.1.0/pg_lease.control)

`pg_lease` provides expiring ownership of named resources with increasing fencing epochs. It is a pure SQL/PL/pgSQL extension with lazy expiry, no background worker, and no preload requirement. Version 0.1.0 is marked testing on PGXN; upstream requires PostgreSQL 14 or later and documents its compatibility tests on PostgreSQL 14 and 15.

### Core Workflow

The extension is trusted: a role with CREATE privilege on the database can install it. It creates the fixed `lease` schema and is not relocatable.

```sql
CREATE EXTENSION pg_lease;

SELECT acquired, epoch, expires_at
FROM lease.acquire('job-42', 'worker-a', interval '30 seconds');

SELECT status, expires_at
FROM lease.renew('job-42', 'worker-a', 1, interval '30 seconds');

SELECT status FROM lease.release('job-42', 'worker-a', 1);
SELECT held, owner, epoch, expires_at FROM lease.inspect('job-42');
```

The example uses the first epoch, 1. Applications must pass the epoch actually returned by acquisition to later renewal, release, and external resource guards. A successful release advances the epoch; re-acquisition must use the newly returned value.

### Operations

| Operation | Behavior |
|---|---|
| `lease.acquire(key, owner, ttl, wait)` | Acquires a free or expired key. The optional wait defaults to zero; a busy key returns a denial. Positive waits poll until acquisition or timeout. |
| `lease.renew(key, owner, epoch, ttl)` | Extends an active lease without changing its epoch. |
| `lease.release(key, owner, epoch)` | Releases an active lease and advances its epoch. |
| `lease.inspect(key)` | Observes ownership, expiry, and the current fencing counter; an unknown key returns no rows. |

Renewal and release return `ok`, `not_owner`, `lapsed`, or `epoch_mismatch`. Invalid inputs raise SQLSTATE `22023`; an exhausted blocking wait raises `57014`. Re-acquiring an unexpired lease with the same owner is idempotent and does not extend its expiry.

### Transactions and Boundaries

Use READ COMMITTED for blocking acquisition. State changes follow the caller's transaction and are undone by rollback. Expiry follows the database server's clock; session termination does not release a lease immediately. The functions use `SECURITY DEFINER` with a fixed search path; direct public access to `lease.leases` is revoked.

Owner strings are application-supplied identifiers, not authenticated database identities. Applications must control access to the lease API and enforce fencing epochs at the resource being protected. The extension provides no safety guarantee across primary promotion or independent databases. Dropping the extension destroys lease state and resets the epoch history on reinstallation; version 0.1.0 does not register its lease table for extension configuration dumps.
