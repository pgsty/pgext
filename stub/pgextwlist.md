## Usage

Sources:

- [pgextwlist v1.20 README](https://github.com/dimitri/pgextwlist/blob/v1.20/README.md)
- [Changes from v1.19 to v1.20](https://github.com/dimitri/pgextwlist/compare/v1.19...v1.20)

pgextwlist lets selected non-superusers run extension lifecycle commands for an explicit allowlist. It temporarily executes those commands with bootstrap-superuser authority, so the allowlist and any custom scripts are part of the database's security boundary.

The module is preloaded and does not itself need CREATE EXTENSION.

### Configure the Allowlist

Load the module for each backend:

    local_preload_libraries = 'pgextwlist'
    extwlist.extensions = 'hstore,cube,pg_stat_statements'

The list may also be assigned per role:

    ALTER ROLE extension_admin
      SET extwlist.extensions = 'hstore,pg_stat_statements';

Reconnect after changing local_preload_libraries. A whitelisted user can then run:

    CREATE EXTENSION hstore;
    ALTER EXTENSION hstore UPDATE;
    COMMENT ON EXTENSION hstore IS 'approved utility';
    DROP EXTENSION hstore;

An extension not named in extwlist.extensions is rejected.

### Restrict Database Ownership

Version 1.20 adds:

    extwlist.restrict_to_database_owner = on

When enabled, the caller must also own the current database. This is off by default for compatibility. Enable it when extension administration should not cross database-ownership boundaries.

### Custom Lifecycle Scripts

Set extwlist.custom_path to an existing readable directory. Version 1.20 raises an error for a missing or unreadable path rather than silently skipping it.

For extension extname, scripts under extname/ can include:

- before--1.0.sql and after--1.0.sql around creation of a specific version.
- before-create.sql and after-create.sql as creation fallbacks.
- before-update.sql and after-update.sql around an update.
- before-drop.sql and after-drop.sql around removal.

Templates can use @extschema@, @current_user@, and @database_owner@. Only trusted administrators should be able to write this directory because scripts execute with elevated authority.

### Configuration Index

- local_preload_libraries: loads pgextwlist into new sessions.
- extwlist.extensions: comma-separated allowlist.
- extwlist.custom_path: base directory for lifecycle scripts.
- extwlist.restrict_to_database_owner: additionally require database ownership.

### Security and Compatibility Notes

- Version 1.20 rejects substitution names containing quote, dollar, apostrophe, or backslash characters, addressing the command-injection class tracked as CVE-2023-39417.
- CREATE EXTENSION, DROP EXTENSION, ALTER EXTENSION UPDATE, and COMMENT ON EXTENSION are supported. ALTER EXTENSION ADD and DROP are not supported.
- Objects created through the elevated path are owned according to the extension/bootstrap-superuser behavior, not necessarily by the requesting role.
- Review extension SQL and custom scripts before adding a name. Whitelisting an extension grants the caller the power embodied by its install and update scripts.
