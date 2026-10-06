## Usage

Sources:

- [packages/errors/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/errors/README.md)
- [packages/errors/sql/errors--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/errors/sql/errors--0.47.0.sql)
- [packages/errors/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/errors/Makefile)
- [packages/errors/errors.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/errors/errors.control)

`errors` supplies the single `errors.raise_error` helper. It raises SQLSTATE P0001 with a code in MESSAGE and a JSON object containing code, context and class in DETAIL.

### Core Workflow

```sql
CREATE EXTENSION "errors" CASCADE;
SELECT errors.raise_error('DEMO_ERROR', '{"field":"name"}'::jsonb, 'public');
```

### Operational Boundaries

The example intentionally raises an exception. The class is application metadata, not access control: callers and logs still receive the supplied context. Keep secrets out of error payloads.

Version 0.47.0 is a SQL/PLpgSQL extension with no own shared library or preload. Install the matching dependency versions first: `plpgsql`. The control permits non-superuser installation but is not marked trusted; dependency, schema and role-grant privileges still apply. No current PostgreSQL-major matrix is declared.
