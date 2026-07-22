## Usage

Sources:

- [Aurora PostgreSQL query plan management guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Optimize.html)
- [Query plan management function reference](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Optimize.Functions.html)
- [Aurora PostgreSQL query plan management release notes](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/auroraqpm.updates.html)
- [Aurora PostgreSQL extension version matrix](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`apg_plan_mgmt` is Amazon Aurora PostgreSQL's provider-managed query plan management extension. It captures optimizer plans, records them in `apg_plan_mgmt.dba_plans`, and lets administrators approve, reject, prefer, disable, validate, or delete managed plans. Version 3.0 is currently listed for Aurora PostgreSQL 18.3; other engine releases expose different versions.

### Enable and inspect plan capture

An Aurora administrator must first set the cluster parameter `rds.enable_plan_management` to `1`, attach the required custom parameter groups, reboot the instance, and create the extension as `rds_superuser`. A conservative manual-capture session looks like this:

```sql
CREATE EXTENSION apg_plan_mgmt;

SET apg_plan_mgmt.capture_plan_baselines = manual;
EXPLAIN SELECT * FROM orders WHERE customer_id = 42;
SET apg_plan_mgmt.capture_plan_baselines = off;

SELECT sql_hash, plan_hash, status, enabled
FROM apg_plan_mgmt.dba_plans;
```

Set `apg_plan_mgmt.use_plan_baselines` only after reviewing the captured baseline and testing it against representative parameters. Grant the `apg_plan_mgmt` role only to users who must view or administer plans.

### Operational boundary

Plan states such as `Approved`, `Unapproved`, `Rejected`, and `Preferred` directly influence which execution plans the optimizer may choose. Treat status changes as production configuration changes and keep a rollback path.

`apg_plan_mgmt.evolve_plan_baselines` runs the captured statement through `EXPLAIN ANALYZE` with the caller's privileges. The statement can therefore perform DML or invoke volatile functions; evaluate candidate plans on a non-production clone or restrict the input to statements whose side effects are understood. Availability and the installable extension version are controlled by the Aurora engine release, not by a portable community package.

AWS's version 3.0 release notes report a fix for plans not being captured and support for `sql_hash` changes during `evolve_plan_baselines`. Validate capture and evolution after upgrading rather than assuming existing baselines remain usable unchanged.
