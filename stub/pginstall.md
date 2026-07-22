## Usage

Sources:

- [Official extension README](https://github.com/dimitri/pginstall/blob/66e4274fc544c48be03c719f8d3f03955649a994/src/client/README.md)
- [Official project README](https://github.com/dimitri/pginstall/blob/66e4274fc544c48be03c719f8d3f03955649a994/README.md)
- [Official control file](https://github.com/dimitri/pginstall/blob/66e4274fc544c48be03c719f8d3f03955649a994/src/client/pginstall.control)

`pginstall` is an experimental extension repository client that intercepts extension utility commands, downloads platform-specific archives, installs files on the server, and can elevate extension operations to superuser. It predates modern PostgreSQL packaging practice and creates a high-risk software-supply-chain boundary; use only in an isolated legacy evaluation environment.

### Configuration and Workflow

The upstream design loads the hook per session and configures writable archive and installation paths plus a repository URL:

```ini
local_preload_libraries = 'pginstall'
pginstall.archive_dir = '/var/cache/pginstall/fetch'
pginstall.control_dir = '/usr/share/postgresql/extension'
pginstall.extension_dir = '/var/lib/pginstall'
pginstall.repository = 'https://extensions.example.invalid/'
pginstall.custom_path = '/etc/pginstall/custom'
pginstall.sudo = false
```

After loading the module, the database-facing workflow is ordinary extension SQL plus repository views:

```sql
CREATE EXTENSION pginstall;

SELECT * FROM pginstall.pg_available_extensions;
SELECT * FROM pginstall.pg_available_extension_versions;
```

The hook may fetch missing artifacts when `CREATE EXTENSION`, `ALTER EXTENSION ... UPDATE`, or `DROP EXTENSION` runs.

### Security Boundary

Archive download, extraction, control-file rewriting, dynamic-library installation, custom before/after SQL, and optional `pginstall.sudo` all cross trust boundaries. A compromised repository or writable cache can become native code execution as the PostgreSQL service account or superuser. HTTPS alone does not provide artifact provenance; require signed, immutable artifacts, verified digests, strict path ownership, an allowlist, and an offline review process.

### Compatibility and Recovery

The source and documentation target PostgreSQL 9.1-era APIs and mention obsolete settings. Do not assume the ProcessUtility hook, file layout, archive format, privilege escalation, or custom scripts work safely on current PostgreSQL. Test only with a disposable cluster and no production credentials.

Filesystem-installed extension files are outside normal database durability. Backups of database objects do not preserve the repository, cache, shared libraries, control files, or custom scripts. Rebuild and disaster recovery therefore require a separately versioned artifact repository and exact platform inventory.
