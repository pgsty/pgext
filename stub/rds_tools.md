## Usage

Sources:

- [AWS guide to RDS PostgreSQL version numbers](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.PostgreSQL.rds.version.html)
- [AWS guide to PostgreSQL password encryption](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL_Password_Encryption_configuration.html)
- [AWS autovacuum diagnostic function reference](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.Autovacuum_Monitoring.Functions.html)
- [AWS extension version matrix](https://docs.aws.amazon.com/AmazonRDS/latest/PostgreSQLReleaseNotes/postgresql-extensions.html)

`rds_tools` is an AWS-maintained, provider-only extension containing administrative functions for Amazon RDS for PostgreSQL. It is not a portable open-source extension and its available version and function set depend on the exact RDS engine release. Version `1.9` is listed in the current AWS extension matrix for newer engine releases.

### Create and Inspect

Connect to the target RDS database with an appropriately privileged role, verify availability, and create the extension in every database where its database-local objects are needed.

```sql
SHOW rds.extensions;

SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'rds_tools';

CREATE EXTENSION rds_tools;
```

Do not infer support solely from a different RDS instance or major version. AWS notes that extensions do not automatically upgrade during an engine upgrade; inspect `pg_extension.extversion` and the current AWS matrix after maintenance.

### Confirmed Workflows

On RDS for PostgreSQL 15.2-R2 and later applicable releases, return the AWS RDS patch version rather than only PostgreSQL's community version.

```sql
SELECT rds_tools.rds_version();
```

Audit how role passwords are encrypted before changing the instance to require SCRAM. AWS directs administrators to repeat this check in every database/instance in scope.

```sql
SELECT *
FROM rds_tools.role_password_encryption_type();
```

On the RDS minor versions listed by AWS for the feature, diagnose aggressive-autovacuum blockers. Run the query in the database with the oldest transaction ID for accurate results.

```sql
SELECT blocker, database, blocker_identifier, wait_event,
       autovacuum_lagging_by, suggestion, suggested_action
FROM rds_tools.postgres_get_av_diag()
ORDER BY autovacuum_lagging_by DESC;
```

### Important Objects

- `rds_tools.rds_version()`: returns the RDS engine patch identifier on supported releases.
- `rds_tools.role_password_encryption_type()`: lists roles and their password encryption type for SCRAM migration audits.
- `rds_tools.postgres_get_av_diag()`: reports identifiable aggressive-vacuum blockers and suggested actions on supported minor releases.

### Provider and Safety Boundaries

Function availability is more granular than extension-version availability. For example, AWS documents `postgres_get_av_diag()` only from specific minor releases within PostgreSQL 13–17; always check the function catalog and current documentation on the target instance.

Diagnostic output can contain SQL text, role information, process identifiers, and suggested terminating or destructive actions. Restrict access and review every `suggested_action`; never execute it automatically. `role_password_encryption_type()` exposes password *methods*, not plaintext passwords, but the result is still security-sensitive. Because AWS does not publish this extension's implementation SQL as an independent upstream project, do not invent or rely on undocumented functions, signatures, or behavior.
