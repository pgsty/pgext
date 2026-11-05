## Usage

Sources:

- [Upstream source 1](https://github.com/vital-ai/vital-graph/blob/b79dc29b4e4652adc49365404d9272fb7ee4f9b3/pg_extensions/vg_planner/test/sql/inert.sql)
- [Upstream source 2](https://github.com/vital-ai/vital-graph/blob/b79dc29b4e4652adc49365404d9272fb7ee4f9b3/pg_extensions/vg_planner/vg_planner.control)
- [Upstream source 3](https://github.com/vital-ai/vital-graph/blob/b79dc29b4e4652adc49365404d9272fb7ee4f9b3/pg_extensions/vg_planner/test/sql/advice.sql)

`vg_planner` 0.2 applies VitalGraph-generated advice to PostgreSQL join cardinalities, methods and order.

### Core Workflow

Append the library and restart before using planner advice. Generated SQL carries structured plan comments; ordinary statements without advice remain unaffected. `vg_planner_last_report` describes which items applied, and `vg_planner.enabled` is the session kill switch.

```ini
shared_preload_libraries = 'vg_planner'
```

```sql
CREATE EXTENSION vg_planner;
SELECT vg_planner_info();
SHOW vg_planner.enabled;
SELECT vg_planner_last_report();
```

### Operational Boundaries

The reviewed source integrates PostgreSQL 18-specific join-planner code and requires superuser installation. Malformed advice is rejected; `vg_planner.on_mismatch` controls the documented mismatch response. Requested methods may be unavailable, so inspect the report and execution plan. This is a development integration for the VitalGraph generator, not a generic compatibility promise for arbitrary hints.
