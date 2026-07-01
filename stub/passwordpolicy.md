


## Usage

Sources: [README](https://github.com/fmbiete/passwordpolicy/blob/v2.0.5/README.md), [v2.0.5 release](https://github.com/fmbiete/passwordpolicy/releases/tag/v2.0.5), [control file](https://github.com/fmbiete/passwordpolicy/blob/v2.0.5/passwordpolicy.control)

`passwordpolicy` is a configurable replacement for PostgreSQL's `passwordcheck` module. It checks passwords during `CREATE ROLE` and `ALTER ROLE`, can enforce password history and validity rules, and can simulate soft account locks after repeated failed logins.

### Enable The Hook

Load the module before other password-check modules, then restart PostgreSQL:

```conf
shared_preload_libraries = 'passwordpolicy'
```

Install the SQL extension in the `postgres` database when using account soft-lock or password-history features:

```sql
CREATE EXTENSION passwordpolicy;
```

### Password Complexity

Settings are dynamic, but new values apply to new sessions:

```conf
password_policy.min_password_len = 15
password_policy.min_special_chars = 1
password_policy.min_numbers = 1
password_policy.min_uppercase_letter = 1
password_policy.min_lowercase_letter = 1
password_policy.require_validuntil = off
```

Enable CrackLib dictionary checks only after creating the dictionary file:

```conf
password_policy.cracklib_dictpath = '/var/cache/cracklib/postgresql_dict'
password_policy.enable_dictionary_check = on
```

### Soft Account Lock

Soft-locking tracks failed login attempts and delays/rejects responses after the configured threshold:

```conf
password_policy_lock.number_failures = 5
password_policy_lock.failure_delay = 5
password_policy_lock.auto_unlock = on
password_policy_lock.auto_unlock_after = 0
password_policy_lock.max_number_accounts = 100
```

Inspect and reset lock state:

```sql
SELECT * FROM passwordpolicy.accounts_locked() ORDER BY usename;
SELECT passwordpolicy.account_locked_reset('app_user');
```

If `password_policy_lock.include_all = false`, only roles listed in `passwordpolicy.accounts_lockable` are considered for soft-lock.

### Password History

Password history stores recent password hashes in the `postgres` database and checks new passwords against them:

```conf
password_policy_history.max_password_history = 5
password_policy_history.max_number_accounts = 100
```

### Caveats

- Version 2.0.5 supports PostgreSQL 14-18.
- This module must be preloaded; changing `shared_preload_libraries` requires a restart.
- PostgreSQL cannot truly block authentication before it happens, so soft-lock simulates the lock by delaying and returning an error. It does not mitigate authentication DoS attacks.
- Size `password_policy_lock.max_number_accounts` and `password_policy_history.max_number_accounts` realistically to avoid wasted memory or missed accounts.
