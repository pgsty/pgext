## Usage

Sources:

- [Official README](https://github.com/snoutdata/snout-utils/blob/331d73c5a37060280d7c29e58b6672f7a493e485/README.md)
- [Package metadata](https://github.com/snoutdata/snout-utils/blob/331d73c5a37060280d7c29e58b6672f7a493e485/Cargo.toml)
- [Distribution recipe](https://github.com/snoutdata/snout-utils/blob/331d73c5a37060280d7c29e58b6672f7a493e485/scripts/build-dist.sh)
- [Configuration implementation](https://github.com/snoutdata/snout-utils/blob/331d73c5a37060280d7c29e58b6672f7a493e485/src/settings.rs)

`snout_utils` 0.1.0 is a preloaded PostgreSQL library that delegates selected superuser operations to a configured database-owner role. It creates no SQL objects and has no per-database installation command. Upstream describes it as built but not yet deployed to its production fleet; the distribution recipe supports PostgreSQL 17 and 18.

### Enable Delegation

An administrator installs the matching library and adds it last in the preload list, so its hook runs before other hooks:

```ini
shared_preload_libraries = 'snout_utils'
snout_utils.superuser = 'platform_admin'
snout_utils.privileged_role = 'app_owner'
snout_utils.privileged_extensions = 'postgis,vector,pg_cron,hstore'
snout_utils.extension_custom_scripts_path = '/usr/local/share/extension-custom-scripts'
snout_utils.privileged_role_allowed_configs = 'session_replication_role,pgaudit.*'
```

The named roles and approved extension libraries must already exist. Restart PostgreSQL to load the library. All five settings use `sighup` context: configuration reloads can change them, while sessions cannot. With no privileged role configured, nothing is delegated.

### Delegated Statements

The configured privileged role and roles granted its privileges can manage extensions on the explicit allowlist. They can also create or alter publications, create eligible foreign-data wrappers and event triggers, and change settings permitted by the configuration allowlist. Delegation applies to the individual statement; the original role is restored afterwards, including on errors.

For a foreign-data wrapper, both handler and validator must belong to the same installed extension. Event-trigger delegation excludes superuser-owned functions, and additional hook checks prevent a superuser's DDL from executing a delegated owner's event-trigger function with elevated rights. Role alterations cannot grant superuser or modify a role already holding privileged-role privileges. Ordinary roles remain subject to PostgreSQL's normal checks.

### Custom Scripts and Security

`snout_utils.extension_custom_scripts_path` selects scripts run as the configured superuser around extension creation. It can contain a global before-create script and per-extension before/after scripts. Missing scripts are skipped. This applies to extension creation by any role, including statements that are not otherwise delegated; protect the directory and its contents as privileged server configuration.

The script substitutions `@extname@`, `@extschema@`, `@extversion@` and `@extcascade@` use validated SQL literals. Unsupported quote or backslash characters are refused. Lists use PostgreSQL-style identifier syntax; a trailing wildcard in a permitted configuration name selects a prefix. Keep both the extension and setting allowlists narrow.

### Operations

The shipped artifact is the shared library only. The repository control file exists for build/test tooling and is not installed by the distribution recipe. Do not try to create this library as a SQL extension.

There is no separate worker or health process. If a required preload library cannot be loaded, PostgreSQL startup fails. Replace the library and restart PostgreSQL for upgrades; no database migration is needed. Dumps do not carry its configuration, so restore and major-upgrade procedures must separately provision the library, roles and server settings.
