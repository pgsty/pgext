## Usage

Sources:

- [Official upstream README](https://github.com/supabase-community/supabase-custom-claims/blob/254e656bf8b77e6f09131cf1de6440b26c99f39e/README.md)
- [Official extension control file (supabase_custom_claims.control)](https://github.com/supabase-community/supabase-custom-claims/blob/254e656bf8b77e6f09131cf1de6440b26c99f39e/supabase_custom_claims.control)
- [Official extension SQL (supabase_custom_claims--1.0.sql)](https://github.com/supabase-community/supabase-custom-claims/blob/254e656bf8b77e6f09131cf1de6440b26c99f39e/supabase_custom_claims--1.0.sql)

`supabase_custom_claims` — This is just one way to implement custom claims for a Supabase project. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION supabase_custom_claims;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
