## Usage

Sources:

- [Official upstream documentation](https://database.dev/basejump/supabase_test_helpers)
- [Official project or provider page](https://usebasejump.com/blog/testing-on-supabase-with-pgtap)

`basejump-supabase_test_helpers` — A collection of functions designed to make testing Supabase projects easier.

The reviewed catalog snapshot records version `0.0.6`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `pgtap`.

```sql
CREATE EXTENSION "basejump-supabase_test_helpers";
SELECT extversion
FROM pg_extension
WHERE extname = 'basejump-supabase_test_helpers';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
