## Usage

Sources:

- [Official database.dev package page](https://database.dev/garyaustin/custom_roles)
- [Official upstream README](https://github.com/GaryAustin1/custom-properties/blob/41f8434fab75171b0fe4080c0d150d85bcc16b18/README.md)
- [Official custom_roles TLE README](https://github.com/GaryAustin1/custom-properties/blob/41f8434fab75171b0fe4080c0d150d85bcc16b18/tle-custom-roles/README.md)
- [Official custom_roles SQL](https://github.com/GaryAustin1/custom-properties/blob/41f8434fab75171b0fe4080c0d150d85bcc16b18/tle-custom-roles/custom_roles--0.0.3.sql)

`custom_roles` is a Supabase-oriented, table-backed role system distributed as a database.dev Trusted Language Extension. It stores application roles separately from PostgreSQL roles and provides helper functions intended for Row Level Security policies. It depends on Supabase's `auth.users` table and `auth.uid()` function.

### Core Workflow

Generate and review a migration for registry version `0.0.3`; the current CLI uses the registry alias shown below. The generated migration installs the TLE through `pg_tle` and creates the namespaced extension.

```bash
dbdev add -o ./migrations -s public -v 0.0.3 package -n "garyaustin@custom_roles"
```

Add allowed role names first, then assign one or more roles to an authenticated user UUID.

```sql
INSERT INTO public.custom_role_names (role_name)
VALUES ('Teacher'), ('Staff');

INSERT INTO public.custom_user_roles (user_id, role)
VALUES ('00000000-0000-0000-0000-000000000001', 'Teacher');
```

Call role helpers through a scalar subquery in RLS policies, as required by the upstream performance guidance.

```sql
CREATE POLICY teacher_read
ON public.lesson
FOR SELECT
TO authenticated
USING ((SELECT public.user_has_role('Teacher')));
```

The optional JWT synchronization trigger is installed disabled. Enable it only if roles must be copied into `auth.users.raw_app_meta_data`; clients do not see changed claims until their JWT is refreshed.

```sql
ALTER TABLE public.custom_user_roles
ENABLE TRIGGER on_custom_role_change;
```

### Important Objects

- `custom_role_names`: allowed names, seeded with `RoleAdmin`.
- `custom_user_roles`: `(user_id, role)` assignments linked to `auth.users` and the name table.
- `user_has_role(text)`: tests one role for the current `auth.uid()`.
- `user_role_in(text[])`: tests whether the current user has any requested role.
- `user_roles_match(text[])`: tests whether the current user has all requested roles.
- `get_user_roles()`: returns the current user's role array.
- `custom_roles_update_to_app_metadata()`: trigger function that writes a `user_roles` array into Supabase app metadata.

### Security and Caveats

Both tables have RLS enabled. Registry version `0.0.3` lets authenticated users read their own assignments and lets users with `RoleAdmin` manage assignments; PostgreSQL and `service_role` bypass or receive broader access according to the Supabase deployment. Audit grants and every policy before use.

The published `0.0.3` SQL fixes function `search_path` to `public`, so installing into another schema is not safely relocatable despite older prose suggesting otherwise. Its disabled trigger is named `on_custom_role_change`, not `custom_role_change`. The registry SQL's delete branch calculates the affected UUID but then reads `new.user_id`; deleting a role while the trigger is enabled can therefore fail or update metadata incorrectly. Review and patch that function in a controlled migration before enabling JWT synchronization. Large role arrays also enlarge JWTs, and upstream advises measuring `get_user_roles()` behavior for users with more than 1,000 roles.
