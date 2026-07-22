## Usage

Sources:

- [AWS Aurora PostgreSQL extension support matrix](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)
- [AWS guide to Aurora PostgreSQL extensions](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Appendix.PostgreSQL.CommonDBATasks.html)
- [AWS extension-version overview](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Extensions.html)

`pgdam` is an AWS-provided extension for Amazon Aurora PostgreSQL. AWS currently lists version `1.7` in supported Aurora engine releases, but does not publish an open-source repository or a public SQL object/function reference for it. Treat it as a provider-managed component whose usable surface must be inspected on the target cluster.

### Core Workflow

First confirm that the exact Aurora engine release offers the extension and that the cluster's `rds.allowed_extensions` policy permits it:

```sql
SELECT name, default_version, installed_version, comment
FROM pg_available_extensions
WHERE name = 'pgdam';

CREATE EXTENSION pgdam;

SELECT extname, extversion, extnamespace::regnamespace
FROM pg_extension
WHERE extname = 'pgdam';
```

Use `\dx+ pgdam` and the PostgreSQL catalogs on that cluster to inspect provider-installed objects and privileges. Do not infer functions from similarly named third-party projects.

### Platform Boundary

Availability and version are tied to the Aurora PostgreSQL engine release. Extension versions do not automatically upgrade during an engine major-version upgrade; AWS recommends checking and upgrading extensions separately. Installation may require `rds_superuser` or delegated extension-management privileges, depending on cluster configuration.

### Caveats

The public AWS material cited here establishes availability of `pgdam` `1.7`, not its application API or behavior. Without a provider object reference, there is not enough authoritative evidence to publish function calls, configuration parameters, or operational promises.

Test the installed extension in a non-production Aurora cluster, record the actual object definitions and grants, and confirm support with AWS before depending on it. Recheck the engine-specific matrix before upgrades because supported versions and allowed-extension policy can change independently of this document.
