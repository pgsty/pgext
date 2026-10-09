## Usage

Sources:

- [packages/infra-utils/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/infra-utils/README.md)
- [packages/infra-utils/sql/infra-utils--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/infra-utils/sql/infra-utils--0.47.0.sql)
- [packages/infra-utils/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/infra-utils/Makefile)
- [packages/infra-utils/infra-utils.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/infra-utils/infra-utils.control)

`infra-utils` compiles infrastructure resource parameters inside PostgreSQL. `infra_utils.compile_resource_spec` accepts a base specification, parameter declarations, supplied values and a member name. `infra_utils.quantity_to_numeric` converts supported decimal and binary quantity suffixes.

### Core Workflow

```sql
CREATE EXTENSION "infra-utils";
SELECT infra_utils.quantity_to_numeric('4Gi');
SELECT infra_utils.compile_resource_spec(
  '{}'::jsonb,
  '[{"key":"replicas","type":"int","min":1,"max":10,"bindings":[{"path":"replicas"}]}]'::jsonb,
  '{"replicas":2}'::jsonb,
  'web'
);
```

### Operational Boundaries

Declarations constrain parameter types, bounds and explicit output paths. `infra_utils.validate_params_schema` and `infra_utils.validate_bundle_params` check definitions and bundle inputs. Unknown declared-member parameters raise errors. If no member declares a parameter interface, the compilation path instead merges supplied JSONB; do not treat that mode as an allowlist. Application admission helpers additionally query provisioned infrastructure schemas. Invalid or unsupported quantity strings return NULL. Dependencies are `plpgsql`, `pgpm-verify`, `errors`, `db-utils`. This SQL-only module needs no library preload or restart; no PostgreSQL-major support matrix is published for this revision. Its fixed schema layout is not relocatable. Install the matching dependency closure before enabling it.
