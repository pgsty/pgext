## Usage

Sources:

- [Upstream source 1](https://github.com/veridit/jsonb_stats/blob/97d7b3dd5997a0d1b2ae5d39357f4e3cb7a797d3/README.md)
- [Upstream source 2](https://github.com/veridit/jsonb_stats/blob/97d7b3dd5997a0d1b2ae5d39357f4e3cb7a797d3/jsonb_stats.control)

`jsonb_stats` 3.2.0 builds mergeable statistical summaries for dynamically named JSONB variables.

### Core Workflow

`stat` constructs a typed value. `jsonb_stats_agg` aggregates statistics and `jsonb_stats_merge_agg` combines existing summaries. Numeric state includes count, sum, sum of squares, minimum and maximum; the 3.x representation uses exact arithmetic rather than an order-dependent floating-point accumulator.

```sql
CREATE EXTENSION jsonb_stats;
SELECT stat(42::integer);
```

### Operational Boundaries

This release's build feature targets PostgreSQL 18. Installation is marked non-superuser, but the role still needs database/schema privileges for its objects. When upgrading, install matching files, reconnect to load the new library, then run ALTER EXTENSION UPDATE. Do not mix persisted summaries from incompatible major representations without following upstream migration guidance.
