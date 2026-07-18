## Usage

Sources:

- [Google Cloud query plan management guide](https://docs.cloud.google.com/alloydb/docs/query-plan-management)
- [Google Cloud extension configuration guide](https://docs.cloud.google.com/alloydb/docs/instance-configure-extensions)
- [Google Cloud database flag reference](https://docs.cloud.google.com/alloydb/docs/reference/database-flags)

google_plan_management is a provider-only Preview extension for AlloyDB for PostgreSQL. It records query templates, plans, and execution statistics, then lets authorized operators approve or deny plans that AlloyDB may use for later executions.

### Enable and inspect

First set the AlloyDB database flag google_plan_management.enabled to on. Google documents that changing this flag restarts the instance. Then connect to the primary instance and run:

```sql
CREATE EXTENSION google_plan_management;
GRANT google_plan_management_role TO plan_operator;
SELECT * FROM google_plan_management.tracked_plans_view;
SELECT * FROM google_plan_management.managed_plans_view;
```

Replace plan_operator with an existing tightly controlled database role. Review tracked plans and measured behavior before calling approve_plan or deny_plan.

### Caveats

- This is an AlloyDB feature, not a portable open-source PostgreSQL extension. Availability and behavior follow the selected AlloyDB release and Google Cloud service terms.
- The feature is Preview and covered by pre-GA terms. Support can be limited and interfaces can change.
- Tracking is enabled by default after the feature is enabled and adds some workload overhead. The repository contains query text, plans, and execution statistics; restrict the management role accordingly.
- Approving a plan changes future optimizer choices. If multiple plans are approved, AlloyDB selects the one with the lowest estimated cost, which need not be the fastest under every parameter or data distribution.
- Query plan management does not support partitioned tables or grouping sets, works only on the primary instance, stores at most 100,000 unique plans, and has no retention policy.
- Keep a rollback procedure: the google_plan_management.enable_steer_plans flag can temporarily stop enforcement while leaving observations available.
