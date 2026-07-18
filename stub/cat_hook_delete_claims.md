## Usage

Sources:

- [database.dev package page](https://database.dev/joel/cat_hook_delete_claims)
- [Official Supabase custom access token hook documentation](https://supabase.com/docs/guides/auth/auth-hooks/custom-access-token-hook)
- [Official dbdev package installation documentation](https://github.com/supabase/dbdev/blob/master/website/content/docs/install-a-package.mdx)

cat_hook_delete_claims version 0.0.1 is a database.dev Trusted Language Extension package for a Supabase custom access token hook. It installs one PL/pgSQL function that returns the incoming claims after removing the optional user_metadata member. It changes newly issued JWT claims only; it does not erase metadata stored for the Auth user and does not revoke tokens that were already issued.

### Install and test

dbdev requires pg_tle and generates a migration for the selected package version. This package's SQL explicitly grants a public-schema function, so generate it for the public schema and review the migration before applying it:

```bash
dbdev add -o ./migrations -s public -v 0.0.1 package -n joel@cat_hook_delete_claims
```

After applying the generated migration, test the hook contract directly:

```sql
SELECT public.custom_access_token_delete_user_metadata(
  jsonb_build_object(
    'claims',
    jsonb_build_object(
      'sub', '00000000-0000-0000-0000-000000000001',
      'role', 'authenticated',
      'user_metadata', jsonb_build_object('display_name', 'Ada')
    )
  )
);
```

Configure that schema-qualified function as the Supabase custom access token hook only after the result preserves every required claim used by the real Auth flow.

### Claim and privilege boundaries

Supabase runs this hook before issuing or refreshing a token and rejects output that omits required claims. The package deletes only user_metadata, while app_metadata and all other members remain. Test password, OAuth, anonymous, refresh, MFA, and service integrations because each flow can supply a different valid claim set. Malformed input or a missing claims object is outside the package's minimal implementation and should fail closed in an application-owned wrapper.

The package grants execution to supabase_auth_admin and revokes it from authenticated and anon, but it does not revoke the PostgreSQL default EXECUTE privilege from PUBLIC. Therefore the named revokes do not by themselves restrict calls. After installation, revoke PUBLIC, then grant only the Auth hook role. Keep the function in public because its grant statements are hard-coded to that schema, or copy and harden the small function under your own migration rather than relocating the package.

database.dev packages depend on pg_tle. The dbdev documentation warns that logical restore can fail for databases containing TLEs and recommends physical backups. Pin version 0.0.1, retain the generated migration, and rehearse backup, restore, hook disablement, and rollback before production use.
