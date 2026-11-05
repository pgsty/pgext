## Usage

Sources:

- [Upstream source 1](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_graphql/alohadb_graphql--1.0.sql)
- [Upstream source 2](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_graphql/alohadb_graphql.control)

`alohadb_graphql` SQL version 1.0 is the GraphQL helper extension included in AlohaDB v20.0.0.

### Core Workflow

`graphql_schema` generates SDL from public-schema tables and `graphql_schema_json` returns a JSONB representation. `graphql_execute(query, variables, operation_name)` parses a request, translates it to SQL and returns a JSONB data wrapper.

```sql
CREATE EXTENSION alohadb_graphql;
SELECT graphql_schema();
SELECT graphql_schema_json();
```

### Operational Boundaries

Use the matching AlohaDB runtime; stock PostgreSQL compatibility is not established by the bundled control file. Installation requires a privileged role. These functions are not an HTTP authentication gateway. Review exposed tables, caller privileges and resolver behavior before allowing untrusted requests, and test supported query/mutation semantics rather than assuming full GraphQL conformance.
