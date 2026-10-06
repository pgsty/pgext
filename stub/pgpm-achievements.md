## Usage

Sources:

- [packages/achievements/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/achievements/README.md)
- [packages/achievements/sql/pgpm-achievements--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/achievements/sql/pgpm-achievements--0.47.0.sql)
- [packages/achievements/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/achievements/Makefile)
- [packages/achievements/pgpm-achievements.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/achievements/pgpm-achievements.control)

`pgpm-achievements` records completed steps, achievements and level requirements. Triggers update achievement state in the fixed status schemas.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-achievements" CASCADE;
SELECT * FROM status_public.levels;
SELECT * FROM status_public.level_requirements;
```

### Operational Boundaries

`status_public.user_steps` and `status_public.user_achievements` hold user progress; `status_public.steps_required` and `status_public.user_achieved` inspect it. Review JWT-claim integration, row policies and SECURITY DEFINER helpers before accepting user-driven writes. Provision level requirements before expecting progression.

Version 0.47.0 is a SQL/PLpgSQL extension with no own shared library or preload. Install the matching dependency versions first: `plpgsql`, `pgpm-jwt-claims`, `pgpm-verify`. Its SQL expects the platform roles `anonymous`, `authenticated` to exist; use the upstream role bootstrap and review grants before installation. The control permits non-superuser installation but is not marked trusted; dependency, schema and role-grant privileges still apply. No current PostgreSQL-major matrix is declared.
