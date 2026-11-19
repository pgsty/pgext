## Usage

Sources:

- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/pg_mdm.control)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/Cargo.toml)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/sql/archive/pg_mdm--0.15.0.sql)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/docs/tutorials/first-entity.md)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/docs/how-to/install.md)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/docs/explanation/security.md)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/docs/reference/compatibility.md)
- [Official source](https://github.com/trickle-labs/pg-mdm/blob/4942af8fb472d1276405b69f032506c7976c906f/sql/configure_helper.sql)

`pg_mdm` 0.15.0 is a synthetic-data demo preview for defining entities, matching records and publishing golden records. Upstream explicitly excludes real workloads and live controls; its release-qualification gates remain blocked. A downloadable package is not evidence of production qualification.

### Prepare the preview

The evidenced setup is PostgreSQL 18 on Linux amd64, with the specifically pinned `pg_trickle` 0.108.4 source build. Both libraries must be preloaded and PostgreSQL restarted before the documented workflow. The extension requires `pg_trickle`, superuser installation and fixed schemas.

```ini
shared_preload_libraries = 'pg_trickle,pg_mdm'
```

```sql
CREATE EXTENSION pg_trickle;
CREATE EXTENSION pg_mdm;
SELECT extname, extversion FROM pg_extension
WHERE extname IN ('pg_trickle', 'pg_mdm');
```

Do not substitute an arbitrary producer build: the integration validates graph/delta capabilities and fails when they are incompatible. Preserve other existing preload entries.

### Define and publish an entity

Follow the first-entity tutorial to create the source tables and register a definition through `mdm.create(mdm.entity(...))`. Definitions contain `mdm.source`, `mdm.field`, `mdm.match` and `mdm.golden_value` values. The example below assumes the tutorial's customer definition and restricted caller are already configured:

```sql
SELECT session_user, current_user;
SELECT mdm_admin.verify_installation();
SELECT mdm.refresh('customer', 'ALLOW');
SELECT * FROM mdm_out.customer;
SELECT * FROM mdm_out.customer_members;
SELECT * FROM mdm_out.customer_review;
```

`mdm.refresh` publishes the value, membership and review tables. Defining an entity does not continuously refresh those public outputs; schedule explicit refreshes. `ALLOW` permits the tutorial's full-refresh path. For `prefer_source`, canonical source-name order determines preference; the `sources` array filters eligibility and its order does not set priority.

### Roles and recovery

Use the supplied `sql/configure_helper.sql` to bind protected objects to a separate helper role with NOLOGIN, NOSUPERUSER, NOBYPASSRLS and no memberships. Keep this owner separate from action roles and application logins. Grant only the required schema/function rights, source SELECT and MAINTAIN privileges, and explicit output access. Both `session_user` and `current_user` are checked; `SET ROLE` from a superuser session alone is insufficient.

Reapply and verify helper binding after upgrade or restore. Policy runtime activation must be re-established deliberately after restart, restore or clone; missing preload prevents activation. Automated policy actions are limited to ASSIGN_QUEUE, ESCALATE and SET_DUE_AT, not automatic merges or golden-value selection. Global publication/digest/snapshot work remains even when some resolution work is incremental.
