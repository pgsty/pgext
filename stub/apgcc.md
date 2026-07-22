## Usage

Sources:

- [Aurora PostgreSQL extension versions](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`apgcc` was an AWS-internal extension exposed by early Aurora PostgreSQL 9.6 releases. AWS now states that it is no longer supported. It has no public upstream SQL contract, so do not design new workloads around it or assume that it can be installed outside Aurora.

### Availability Check

Use catalog inspection only; do not issue `CREATE EXTENSION` unless the exact Aurora release and AWS support guidance explicitly permit it.

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'apgcc';

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'apgcc';
```

### Operational Notes

- AWS lists version 1.0 only for old Aurora PostgreSQL 9.6 releases and marks the internal extension unsupported.
- An existing installation is a migration risk. Inventory dependencies, test the application without it, and follow AWS upgrade guidance before moving to a supported Aurora release.
- There is no documented portable API, configuration, or replacement supplied by the cited upstream page. Contact AWS Support when an old cluster still depends on it.
