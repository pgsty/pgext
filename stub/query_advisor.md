## Usage

Sources:

- [Official EDB Query Advisor documentation](https://www.enterprisedb.com/docs/pg_extensions/query_advisor/)
- [Official configuration guide](https://www.enterprisedb.com/docs/pg_extensions/query_advisor/configuring/)
- [Official usage guide](https://www.enterprisedb.com/docs/pg_extensions/query_advisor/using/)

`query_advisor` samples workload predicates and planning errors to recommend useful indexes and multicolumn extended statistics without creating every candidate object first.

### Enablement

The product package is named EDB Query Advisor, but the canonical extension and preload identity is `query_advisor`. Install version 1.2.2, preload it, restart, and create it:

```ini
shared_preload_libraries = 'query_advisor'
```

```sql
CREATE EXTENSION query_advisor;
```

Use `query_advisor.sample_rate` and `query_advisor.exclude_schema_list` to scope collection. Capacity parameters such as `query_advisor.max_qual_entries` require a restart.

### Generate Recommendations

```sql
SELECT *
FROM query_advisor_index_recommendations();

SELECT *
FROM query_advisor_statistics_recommendations();

SELECT *
FROM query_advisor_qualstats_pretty;
```

Index recommendations are costed against captured workload queries with hypothetical indexes. Statistics recommendations focus on two-column combinations and include weights and benefited query IDs.

### Review Boundary

Collected data is held in memory and is lost at server restart. Recommendations are candidates, not automatic DDL: validate write amplification, storage, maintenance, redundant indexes, and production plan changes before executing the generated `CREATE INDEX` or `CREATE STATISTICS` command.

