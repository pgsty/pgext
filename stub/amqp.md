## Usage

Sources:

- [Official extension control file](https://github.com/omniti-labs/pg_amqp/blob/240d477d40c5e7a579b931c98eb29cef4edda164/amqp.control)
- [Official upstream documentation](https://github.com/omniti-labs/pg_amqp/blob/240d477d40c5e7a579b931c98eb29cef4edda164/doc/amqp.md)
- [Official PGXN distribution page](https://pgxn.org/dist/pg_amqp/)

`amqp` — Publish messages to AMQP brokers from PostgreSQL

The reviewed catalog snapshot records version `0.4.2`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "amqp";
SELECT extversion
FROM pg_extension
WHERE extname = 'amqp';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
