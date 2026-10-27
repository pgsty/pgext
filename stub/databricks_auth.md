## Usage

Sources:

- [Extension versions](https://docs.databricks.com/aws/en/oltp/projects/extensions)

`databricks_auth` is the provider's PostgreSQL authentication extension in managed Databricks Lakebase. The official catalog lists version 1.0 on PostgreSQL 16–18. Its public extension listing does not document a standalone authentication API or a portable source package.

### Inspect and Enable

Use the Lakebase SQL Editor or a connected PostgreSQL client with an authorized database role. Inspect availability and the existing installation first:

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions WHERE name = 'databricks_auth';
```

The provider documents standard extension installation for its listed extensions:

```sql
CREATE EXTENSION IF NOT EXISTS databricks_auth;
```

This does not configure client credentials, grant application access or replace the platform's authentication setup. Follow Databricks account and database access policy; do not remove an existing provider-managed installation.

### Version Boundary

Query the actual compute's available and installed versions before planning an update. Lakebase may need a compute restart to make newly supplied versions available, interrupting connections. Source language, license, SQL functions, schema, dependencies and preload behavior are not established by this catalog; no additional API or local installation steps are asserted here.
