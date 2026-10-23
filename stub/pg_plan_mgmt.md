## Usage

Sources:

- [Official source](https://github.com/pganalyze/pg_plan_mgmt/blob/cab88e858fe7335d086b46876a19004eb947a52d/README.md)
- [Official source](https://github.com/pganalyze/pg_plan_mgmt/blob/cab88e858fe7335d086b46876a19004eb947a52d/pg_plan_mgmt.control)
- [Official source](https://github.com/pganalyze/pg_plan_mgmt/blob/cab88e858fe7335d086b46876a19004eb947a52d/pg_plan_mgmt--1.0.sql)

`pg_plan_mgmt` captures and selects query-plan baselines using PostgreSQL 19's `pg_plan_advice` module. This first iteration is experimental: SQL interfaces, configuration and storage can change, and upstream does not recommend production use.

### Capture and Inspect

Load both modules, preserving other configured libraries, and restart:

```ini
shared_preload_libraries = 'pg_plan_advice,pg_plan_mgmt'
```

```sql
CREATE EXTENSION pg_plan_mgmt;
SET pg_plan_mgmt.capture_plan_baselines = 'manual';
SELECT count(*) FROM pg_class;
SELECT query_id, plan_no, status FROM plan_mgmt.plans;
SET pg_plan_mgmt.capture_plan_baselines = 'off';
SET pg_plan_mgmt.use_plan_baselines = on;
```

Capture is restricted to superusers, including sessions that subsequently change role. The first captured plan is approved; later plans start unapproved. Enabled preferred plans take precedence over approved plans. If neither applies, the optimizer can fall back to its ordinary minimum-cost plan.

### Management and Boundaries

`plan_mgmt.set_plan_status`, `plan_mgmt.set_plan_enabled`, `plan_mgmt.delete_plan` and `plan_mgmt.validate_plans` manage stored plans; `plan_mgmt.add_plan` imports a SQL statement and plan outline. Validation can update, disable or delete invalid entries, so choose its action deliberately.

The fixed schema is `plan_mgmt`; the extension is not relocatable. SQL interfaces are revoked from public access. Granting management functions affects plans across the database. Captured SQL can contain literals: restrict access and protect persistence files. Shared preload is required for the persistence and retention workers. Session loading permits in-memory operation, but those baselines do not survive server restart. The default retention period is 32 days; configuration limits bound statement and plan counts.
