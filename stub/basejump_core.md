## Usage

Sources:

- [basejump_core 2.0.1 on database.dev](https://database.dev/basejump/basejump_core)
- [Basejump product and RPC overview](https://usebasejump.com/)
- [Account API migration at the reviewed commit](https://github.com/usebasejump/basejump/blob/cfaeaa314e2bd336ff296c8b45fa5f0fa7e4856b/supabase/migrations/20240414161947_basejump-accounts.sql)

`basejump_core` is the database package behind Basejump's Supabase account layer. Version 2.0.1 creates personal and team accounts, account membership and roles, invitations, billing records, RLS policies, triggers, and public RPC functions. Install the package into a Supabase migration with dbdev, apply that migration, and then enable the extension:

```bash
dbdev add -o ./migrations -s extensions -v 2.0.1 package -n "basejump@basejump_core"
```

```sql
CREATE EXTENSION IF NOT EXISTS basejump_core WITH SCHEMA extensions;

SELECT create_account('team-slug', 'Your Team Name');
SELECT get_accounts();
```

The public RPC functions use the calling Supabase identity. For example, `create_account` creates an account for the current authenticated user, while `get_accounts` returns accounts visible to that user.

### Caveats

- This is a Supabase-oriented schema package, not a standalone authentication server. It expects Supabase objects and roles such as `auth.users`, `authenticated`, and `service_role`.
- The extension installs security-definer functions, triggers, grants, and RLS policies. Review configuration and privileges before applying its generated migration to an existing project.
- Account and billing data remain in the database. Test upgrades and backups as schema migrations, not as stateless application deployments.
