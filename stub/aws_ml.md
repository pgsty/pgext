## Usage

Sources:

- [Official upstream documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-ml.html)

`aws_ml` — Amazon Aurora PostgreSQL extension for invoking Amazon Comprehend, SageMaker AI, and Bedrock services from SQL.

The reviewed catalog snapshot records version `2.0`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `aws_commons`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "aws_ml";
SELECT extversion
FROM pg_extension
WHERE extname = 'aws_ml';
```

This is a provider-specific component for `AWS`; availability, enablement, privileges, and upgrades follow that service rather than a portable community package.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
