## Usage

Sources:

- [packages/stamps/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/stamps/README.md)
- [packages/stamps/sql/pgpm-stamps--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/stamps/sql/pgpm-stamps--0.47.0.sql)
- [packages/stamps/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/stamps/Makefile)
- [packages/stamps/pgpm-stamps.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/stamps/pgpm-stamps.control)

`pgpm-stamps` supplies `stamps.timestamps()` and `stamps.peoplestamps()` triggers. The first maintains creation/update timestamps; the second uses `jwt_public.current_user_id()` to maintain creation/update actor UUIDs.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-stamps";
CREATE TABLE stamp_demo (
  id integer PRIMARY KEY,
  body text,
  created_at timestamptz,
  updated_at timestamptz
);
CREATE TRIGGER stamp_demo_timestamps
BEFORE INSERT OR UPDATE ON stamp_demo
FOR EACH ROW EXECUTE FUNCTION stamps.timestamps();
```

### Operational Boundaries

Attach the triggers before INSERT or UPDATE. Timestamp tracking needs `created_at` and `updated_at` columns; actor tracking needs `created_by` and `updated_by`. On update, the creation fields are preserved and update fields are replaced. JWT identity is supplied by the application. These triggers record row metadata, not a separate immutable history. Dependencies are `plpgsql`, `pgpm-jwt-claims`, `pgpm-verify`. The installation SQL grants privileges to the pre-existing roles `authenticated`, `anonymous`. This SQL-only module needs no library preload or restart; no PostgreSQL-major support matrix is published for this revision. Its fixed schema layout is not relocatable. Install the matching dependency closure before enabling it.
