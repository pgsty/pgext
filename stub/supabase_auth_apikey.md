## Usage

Sources:

- [Official Database.dev package page for version 0.0.3](https://database.dev/martindonadieu/supabase_auth_apikey)
- [Official Database.dev package collection](https://database.dev/martindonadieu/pg-extensions)

`supabase_auth_apikey` is a Database.dev SQL package that adds API-key authentication helpers around Supabase Auth users and permission arrays. It targets non-interactive clients such as IoT devices and command-line tools when email/password authentication is unsuitable.

### Installation and Core Workflow

The published package version is `0.0.3`. Install it with the SQL package identifier shown on the official page, then create the canonical PostgreSQL extension, preferably in a dedicated schema:

```sql
SELECT * FROM dbdev('riderx-supabase_auth_apikey');
CREATE EXTENSION supabase_auth_apikey WITH SCHEMA extensions;

SELECT extensions.create_apikey(ARRAY['read', 'write']);

SELECT extensions.is_allowed_apikey(
    (current_setting('request.headers', true)::json ->> 'apikey'),
    ARRAY['read', 'write']
);
```

The official package page documents these user-facing functions:

- `create_apikey(user_id, permissions)` for a backend acting for a user.
- `create_apikey(permissions)` for the current authenticated user.
- `delete_apikey(user_id, apikey)` and `delete_apikey(apikey)` for the corresponding deletion paths.
- `get_user_id(apikey)` to resolve a key's user.
- `is_allowed_apikey(apikey, permissions)` for permission checks, including row-level security policies.

### Security Boundaries

The API key comes from the Supabase/PostgREST `request.headers` setting in the documented RLS pattern. Missing headers, malformed JSON, revoked keys, and insufficient permissions must deny access. Keep backend-only overloads unreachable from client roles and grant only the schema/function privileges needed by each caller.

The package's published source link currently points to a repository path that is not publicly available, so the public package page does not provide enough evidence to assert how keys are generated, hashed, compared, or stored. Audit the exact generated migration before installation, test rotation and revocation, and never log or persist returned plaintext keys unnecessarily. This package is separate from Supabase's platform-supported Auth API surface.
