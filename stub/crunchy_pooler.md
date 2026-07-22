## Usage

Sources:

- [Crunchy Bridge PgBouncer documentation](https://docs.crunchybridge.com/how-to/pgbouncer)
- [Crunchy Bridge product page](https://www.crunchydata.com/products/crunchy-bridge)

`crunchy_pooler` is the provider-managed helper extension that lets the PgBouncer service on Crunchy Bridge validate database roles against PostgreSQL. It is specific to Crunchy Bridge rather than a portable community package, and it must be enabled separately in every database whose users will connect through the pooler.

### Enable pooled connections

Connect as a superuser to each target database and install the extension:

```sql
CREATE EXTENSION crunchy_pooler;
```

The extension creates the `crunchy_pooler` role and grants it access only to the provider's `user_lookup` authentication helper. Dropping the extension removes this authentication path and prevents PgBouncer from accepting connections to that database.

Use the cluster's normal connection string with port `5431` for PgBouncer; port `5432` connects directly to PostgreSQL. The username, password, database, and TLS settings otherwise remain the same.

### Pooling modes and limits

Crunchy Bridge uses transaction pooling by default. Session-level state must therefore not be assumed to survive across transactions. Prepared statements are supported on the service with PgBouncer 1.22 or later when `max_prepared_statements` is greater than zero; the documented service default is `250`. Session pooling is available when an application needs a stable backend session, while statement pooling cannot run multi-statement transactions.

Superuser and replication roles are intentionally rejected through PgBouncer and must connect directly on port `5432`. If the extension is missing from the selected database, clients receive `FATAL: bouncer config error`. Availability, upgrades, and implementation details remain controlled by Crunchy Data, so check the current service documentation before relying on a particular PgBouncer version or setting.
