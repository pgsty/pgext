## Usage

Sources:

- [Upstream source](https://github.com/turbot/steampipe-plugin-aws/releases/tag/v1.36.0)
- [Upstream source](https://github.com/turbot/steampipe-plugin-aws/releases/download/v1.36.0/steampipe_postgres_aws.pg15.darwin_arm64.tar.gz)

- [Published 0.125.0 package page](https://database.dev/toliver38/steampipe_postgres_aws)
- [Official AWS plugin documentation](https://hub.steampipe.io/plugins/turbot/aws#postgres-fdw)
- [Official Postgres FDW configuration guide](https://steampipe.io/docs/steampipe_postgres/configure)

`steampipe_postgres_aws` is the PostgreSQL FDW build of Turbot's AWS Steampipe plugin. It translates SQL scans of imported foreign tables into live AWS API calls, allowing infrastructure inventory and configuration analysis from an existing PostgreSQL database. The package is not the Steampipe CLI and does not copy AWS resources into local tables unless you do so explicitly.

### Core Workflow

After installing binaries that match the operating system, architecture, and PostgreSQL major version, create the extension, configure a foreign server, and import its tables into a schema:

```sql
CREATE EXTENSION steampipe_postgres_aws;

CREATE SERVER steampipe_aws
  FOREIGN DATA WRAPPER steampipe_postgres_aws
  OPTIONS (config 'profile = "default"
regions = ["us-east-1"]');

CREATE SCHEMA aws;
COMMENT ON SCHEMA aws IS 'Steampipe AWS FDW';

IMPORT FOREIGN SCHEMA aws
  FROM SERVER steampipe_aws
  INTO aws;

SELECT name, region
FROM aws.aws_s3_bucket
ORDER BY name;
```

The `config` server option is a newline-sensitive HCL string containing AWS plugin connection arguments. Instead of a profile, it can use the credential mechanisms documented by the AWS plugin. Any environment variables or files under a home directory are resolved in the PostgreSQL server operating-system user's context, not the interactive client user's context.

### Object and Query Model

- Extension/FDW: `steampipe_postgres_aws`.
- Foreign server: user-defined; create one per distinct AWS account/region configuration.
- Imported schema: user-defined; `IMPORT FOREIGN SCHEMA` creates the foreign-table definitions supplied by the plugin.
- Foreign tables: service-oriented `aws_*` tables such as `aws_s3_bucket`, with available columns, filters, permissions, and examples documented individually in the Steampipe Hub.

Queries call AWS APIs and can be slower, rate-limited, paginated, eventually consistent, or denied by IAM. Push supported key columns into `WHERE` clauses when a table's documentation identifies required or optional quals. Materialize a snapshot locally when repeatable analytics, joins, or historical comparison matter.

### Security and Version Boundaries

Prefer an AWS profile, workload role, or other short-lived credential source over static keys in `pg_foreign_server`. Grant only the IAM read permissions needed by the tables you query, and restrict access to server options and imported foreign tables. Foreign scans may incur AWS API activity and related service costs.

The verified official 1.36.0 release includes standalone PostgreSQL AWS artifacts; the inspected PostgreSQL 15 macOS arm64 artifact retains SQL version 1.0. The older database.dev package remains a separate 0.125.0 publication. Match the official artifact to the actual platform and PostgreSQL major; do not infer newer-major support from the plugin tag. SQL extversion alone cannot identify the plugin revision. Review table-definition changes and re-import or recreate foreign tables deliberately, preserving grants and dependent objects. The installer README contains destructive DROP examples, which are unnecessary for a fresh installation and should not be copied over an existing deployment. No shared_preload_libraries setting is required.
