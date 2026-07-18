## Usage

Sources:

- [Extension README](https://github.com/kingluo/pgcat-pgxs/blob/bd8d5975680f6870965240dc9cf1218229d9c3a9/README.md)
- [Extension SQL](https://github.com/kingluo/pgcat-pgxs/blob/bd8d5975680f6870965240dc9cf1218229d9c3a9/pgcat--1.0.sql)
- [Extension control file](https://github.com/kingluo/pgcat-pgxs/blob/bd8d5975680f6870965240dc9cf1218229d9c3a9/pgcat.control)
- [PGXN usage documentation](https://pgxn.org/dist/pgcat/README.html)

`pgcat` is the database-side component of the external pgcat logical-replication daemon. The extension creates configuration/progress tables and privileged helper functions in the fixed `pgcat` schema; the separate daemon reads those tables and applies changes.

The `pgcat` role must exist before extension creation. A minimal one-way setup uses a replication role and publication on the publisher, then a matching role, extension, target table privileges, and subscription row on the subscriber:

```sql
-- Publisher
CREATE ROLE pgcat LOGIN REPLICATION PASSWORD 'replace-me';
CREATE SCHEMA pgcat AUTHORIZATION pgcat;
CREATE EXTENSION pgcat;
CREATE PUBLICATION app_pub FOR TABLE public.orders;
GRANT SELECT ON public.orders TO pgcat;

-- Subscriber
CREATE ROLE pgcat LOGIN PASSWORD 'replace-me';
CREATE SCHEMA pgcat AUTHORIZATION pgcat;
CREATE EXTENSION pgcat;
GRANT SELECT, INSERT, UPDATE, DELETE, TRUNCATE ON public.orders TO pgcat;

INSERT INTO pgcat.pgcat_subscription
  (name, hostname, port, username, password, dbname, publications, copy_data, enabled)
VALUES
  ('orders_from_primary', 'primary.example.net', 5432, 'pgcat', 'replace-me',
   'app', ARRAY['app_pub'], true, true);
```

Start and configure the external `pgcat` process after the database objects are ready.

### Version and security

The extension control exposes version `1.0`, while the PGXN distribution is labeled `1.0.0`. Install the extension on both publisher and subscriber as directed by upstream. Connection passwords are stored as plain `text` in `pgcat.pgcat_subscription`; tightly restrict access to the fixed schema, the `pgcat` role, and daemon configuration. Logical replication also requires suitable `wal_level`, `pg_hba.conf`, publication, replica-identity, and target-table preparation.
