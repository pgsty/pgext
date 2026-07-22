## Usage

Sources:

- [Official AlloyDB index advisor documentation](https://docs.cloud.google.com/alloydb/docs/use-index-advisor)

`google_db_advisor` is the provider-managed AlloyDB index-advisor extension. It analyzes recorded workload statements and recommends conventional or vector indexes; it does not create the proposed indexes automatically.

### Core Workflow

Enable the extension by following the AlloyDB service procedure, run a representative workload, then request an analysis when needed:

```sql
SELECT google_db_advisor_recommend_indexes();

SELECT *
FROM google_db_advisor_recommended_indexes;
```

Review the evidence behind the recommendations before applying any DDL:

```sql
SELECT * FROM google_db_advisor_workload_report;
SELECT * FROM google_db_advisor_workload_statements;
```

After capturing the result or before starting a clean measurement window, reset the advisor state:

```sql
SELECT google_db_advisor_reset();
```

### Important Objects

- `google_db_advisor_recommend_indexes()` starts an on-demand analysis.
- `google_db_advisor_recommended_indexes` exposes proposed index DDL and recommendation details visible to the current role.
- `google_db_advisor_workload_report` summarizes the analyzed workload.
- `google_db_advisor_workload_statements` exposes the statements used by the advisor.
- `google_db_advisor_reset()` clears the advisor's collected workload and recommendations.

### Provider and Safety Boundaries

- This is an AlloyDB service component, not a portable community package. Availability, enablement, versioning, and privileges follow the selected AlloyDB instance and its current documentation.
- Exercise all important query shapes during the collection window. A partial or unrepresentative workload can produce incomplete recommendations.
- Treat generated DDL as a proposal: check duplicate or overlapping indexes, write amplification, storage, build time, locking, and later removal before applying it.
- Results are filtered by the querying role. Use a deliberate review role and do not assume one user's result is a cluster-wide inventory.
- Vector index advice has additional service flags and preview limitations documented by Google; verify current provider requirements before using it in production.
