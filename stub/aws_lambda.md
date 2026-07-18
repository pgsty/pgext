## Usage

Sources:

- [Official upstream documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL-Lambda.html)
- [Official project or provider page](https://aws.amazon.com/rds/postgresql/)

`aws_lambda` — Amazon RDS for PostgreSQL extension for invoking AWS Lambda functions from SQL.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `aws_commons`.
The curated compatibility set is `12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "aws_lambda";
SELECT extversion
FROM pg_extension
WHERE extname = 'aws_lambda';
```

This is a provider-specific component for `AWS`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
