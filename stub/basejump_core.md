## Usage

Sources:

- [basejump_core 2.0.1 package](https://database.dev/basejump/basejump_core)
- [Official Basejump README](https://github.com/usebasejump/basejump/blob/cfaeaa314e2bd336ff296c8b45fa5f0fa7e4856b/README.md)
- [Account schema and RPC implementation](https://github.com/usebasejump/basejump/blob/cfaeaa314e2bd336ff296c8b45fa5f0fa7e4856b/supabase/migrations/20240414161947_basejump-accounts.sql)
- [Invitation implementation](https://github.com/usebasejump/basejump/blob/cfaeaa314e2bd336ff296c8b45fa5f0fa7e4856b/supabase/migrations/20240414162100_basejump-invitations.sql)
- [Billing implementation](https://github.com/usebasejump/basejump/blob/cfaeaa314e2bd336ff296c8b45fa5f0fa7e4856b/supabase/migrations/20240414162131_basejump-billing.sql)

`basejump_core` 2.0.1 is the database package behind Basejump's Supabase account layer. It adds personal and team accounts, memberships and roles, invitations, billing state, row-level security policies, triggers, and public RPC functions. It is designed for a Supabase database and is not a standalone authentication or billing service.

### Install and Enable

Generate a migration with the official database.dev package name, apply it through the normal Supabase migration workflow, and then enable the generated extension:

```bash
dbdev add -o ./migrations -s extensions -v 2.0.1 package -n "basejump@basejump_core"
```

```sql
CREATE EXTENSION IF NOT EXISTS basejump_core WITH SCHEMA extensions;
```

The package expects Supabase identities and roles such as `auth.users`, `authenticated`, and `service_role`. Its SQL also references functions including `auth.uid()`, `extensions.uuid_generate_v4()`, and `gen_random_bytes()`; verify that the Supabase project supplies those prerequisites before applying the migration.

### Account Workflow

Public functions are intended to be called through Supabase RPC with the caller's authenticated identity:

```sql
SELECT create_account('engineering', 'Engineering');
SELECT get_accounts();
SELECT get_account_by_slug('engineering');
SELECT get_account_members(get_account_id('engineering'));
```

Owners can manage membership and invitations through functions such as `update_account_user_role`, `remove_account_member`, `create_invitation`, `get_account_invitations`, `lookup_invitation`, and `accept_invitation`. Billing consumers use `get_account_billing_status`; synchronization from a trusted billing webhook uses `service_role_upsert_customer_subscription`.

### Stored Objects

- `basejump.accounts` and `basejump.account_user`: accounts, members, and account roles.
- `basejump.invitations`: one-time or 24-hour invitation tokens.
- `basejump.billing_customers` and `basejump.billing_subscriptions`: provider customer and subscription state.
- `basejump.config`: feature flags for team accounts, billing, and the billing provider.
- Public RPC functions: the supported application-facing interface; direct table access is constrained by grants and RLS.

### Security and Maintenance

The package installs RLS policies, triggers, grants, and multiple `SECURITY DEFINER` functions. Review the generated versioned migration before applying it, preserve the documented `search_path` settings, and test calls as `anon`, `authenticated`, and `service_role` rather than only as an administrator.

Account, invitation, and billing rows are durable application state. Back up the schema and data together, test version upgrades as database migrations, and keep webhook credentials and service-role access outside untrusted clients. Invitation tokens and billing payloads are sensitive even though RLS limits ordinary reads.
