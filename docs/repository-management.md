# Repository management

`pgext repo` configures Linux RPM/APT package sources. The embedded repository
definitions in `cli/repo/catalogs/{rpm,deb}.yml` are also included verbatim by
`cli/confgen.go` in Pigsty's generated OS configuration. Repository URLs,
modules, release/architecture filters, mirrors and metadata have one source.
These commands do not connect to PostgreSQL. The former `pgext repo` database
inventory has moved to `pgext repo summary`.

## Commands

| Command | Behavior |
| --- | --- |
| `repo` | Show help |
| `repo list [selectors]` | List repositories for the detected/selected system |
| `repo list all` / `repo list --all` | List the complete catalog across platforms |
| `repo info [selectors]` | Show definitions, release/architecture filters, mirrors and metadata |
| `repo show [selectors]` | Print the exact rendered configurations; default selection is `all` |
| `repo status` | List installed source files, including files owned by other tools |
| `repo add [selectors]` | Add or update selected PGEXT files; default selection is `all` |
| `repo set [selectors]` | Replace the PGEXT-managed selection; default selection is `all` |
| `repo rm <selectors>` | Remove selected PGEXT-managed files |
| `repo rm --all` | Remove every PGEXT-managed file, including custom modules |
| `repo update` | Run `dnf makecache` or `apt-get update -o APT::Update::Error-Mode=any` |
| `repo restore <backup-id>` | Undo a successful file change if those files still match its recorded result |
| `repo summary` | Show fetched repository metadata from the database |

Selectors accept spaces or commas. They can name individual repositories
(`pigsty-pgsql`, `pgdg18`, `docker-ce`) or modules (`node`, `pgsql`, `infra`,
`extra`, `docker`, `local`, and other modules shown by `list`).

- `all` means `node + pgsql + infra`, excluding optional and local sources.
- `pgdg` selects stable PGDG repositories within the `pgsql` module.
- `pigsty` selects Pigsty's `pgsql` and `infra` repositories, excluding `local`.
- Selection is filtered by OS and architecture, then deduplicated in catalog
  order. Every explicit selector must match an available repository. A typo or
  unsupported selection fails before any change.

`rm all` removes the default modules; `rm --all` removes all PGEXT-owned source
files. Bare `rm` is an error. The `r` alias is available for `repo`; `ls`/`l`
for `list`, `i` for `info`, and `remove` for `rm`.

## Configure sources

```bash
pgext repo list
pgext repo info pgsql docker
pgext repo show pgsql --mirror
pgext repo add pgdg,pigsty --dry-run
sudo pgext repo add pgdg,pigsty --update
sudo pgext repo add docker --region china
sudo pgext repo rm docker --update
```

`--mirror` is shorthand for `--region china`. Repositories without that region
use their default URL. RPM repositories retain the catalog's ordered list of
mirror URLs. `--admin-ip host[:port]` supplies the endpoint for `local` sources.

`add`, `rm`, and `restore` update package metadata only with `--update` (`-u`).
`set` updates it by default; use `--update=false` to defer the refresh.

PGEXT writes one marked file per repository:

- RPM: `/etc/yum.repos.d/pgext-<name>.repo`.
- APT: `/etc/apt/sources.list.d/pgext-<name>.list`.

A filename prefix alone does not establish ownership: a PGEXT management
header must also be present. `set` removes obsolete PGEXT files and installs
the requested selection. Existing distribution definitions that cover the
same node source are reused, including a configured distribution mirror.

To migrate legacy PGDG or Pigsty source definitions, preview `--replace` with
the repositories to migrate:

```bash
pgext repo set pgdg,pigsty --replace --mirror --dry-run
sudo pgext repo set pgdg,pigsty --replace --mirror
```

`--replace` migrates recognized conflicts in the selection. A mixed file
keeps its unselected RPM sections, APT lines or Deb822 stanzas. Split a
Deb822 stanza if it combines selected and unselected sources. Unrelated
PGEXT-managed files are preserved in this mode.

## Platforms and custom catalogs

Linux detects its OS, full version and architecture from `/etc/os-release`
and the running binary. macOS can list the full catalog and preview or stage
Linux configurations with an explicit target:

```bash
pgext repo show pgdg --platform el9.x86_64 --os-version 9.6 --mirror
pgext repo show pgdg --platform el10.arm64 --os-version 10.1
pgext repo show node --platform u24.arm64
```

Both `amd64`/`x86_64` and `arm64`/`aarch64` are accepted. PGDG uses EL9's minor
version from 9.6 onward and the full EL10 version. EL10 PGDG previews require
`--os-version`; actual Linux detection supplies it automatically. EPEL uses
`10z` on EL10. Other RPM sources use the OS major unless `minor: true` says
otherwise. Debian/Ubuntu sources expand their vendor and codename; Ubuntu ARM
uses the separate ports repositories.

Use a generated Pigsty OS YAML or a small custom file:

```bash
pgext repo list --catalog ./u24.aarch64.yml --platform u24.arm64
pgext repo show custom --catalog ./repos.yml --platform u24.arm64
```

```yaml
repo_upstream:
  - name: custom
    description: Internal packages
    module: custom
    releases: [24]
    arch: [x86_64, aarch64]
    baseurl:
      default: 'https://packages.example.org/apt ${distro_codename} main'
    meta:
      signed-by: /etc/apt/keyrings/custom.asc
```

An explicit `repo_upstream`, including an empty list, takes precedence over
`repo_upstream_default`. Other Pigsty settings in the file are ignored.

RPM defaults to `gpgcheck=1`. APT uses the system trust policy for sources
without a dedicated key. PGEXT installs its bundled Pigsty and PGDG public
keys for the selected definitions and points the renderer at those files.
The explicitly selected `pigsty-local` intranet source supports unsigned
metadata. Custom catalogs can supply explicit `gpgcheck`/`gpgkey` or
`signed-by`/`trusted` settings. Provision any custom key files before
refreshing metadata.

`pgext env repo` calls the same implementation as `pgext repo set`. PGDG
keys are sourced from the PostgreSQL project. [APT key](https://www.postgresql.org/download/linux/debian/)
· [RPM keys](https://download.postgresql.org/pub/repos/yum/keys/)

## Preview, backups and restore

`--dry-run` and its alias `--plan` perform selection, rendering, conflict and
restore checks without creating directories, files, locks, backups or caches.
`--json` returns structured catalog information or a change result with before
and after content, commands, success/error, and backup ID. Native package-manager
output goes to stderr.

An existing directory passed with `--root` contains an isolated filesystem tree:

```bash
mkdir -p /tmp/pgext-repo-demo
pgext repo set all --root /tmp/pgext-repo-demo --platform u24.arm64 --update=false
pgext repo status --root /tmp/pgext-repo-demo --platform u24.arm64
pgext repo rm pigsty --root /tmp/pgext-repo-demo --platform u24.arm64
pgext repo restore <backup-id> --root /tmp/pgext-repo-demo --platform u24.arm64
```

Cache updates are rejected for staging roots. System mutations require root
on the detected Linux platform; a platform override cannot redirect a host
operation to another OS. Symlinks beneath the selected root are rejected.

Every actual file change records its before/after content and permissions in
a private JSON backup under `/var/lib/pgext/repo/backups/` (beneath `--root`
when staging). The backup includes signing-key changes. Restoring removes
files first created by that operation. Identical configurations are a no-op
and create no backup.

Operations serialize through a process lock and check for changes between
planning and writing. Individual writes use atomic rename. A write or cache
refresh failure restores applied file changes, reports whether rollback
succeeded, and retains the backup. Package-manager cache contents themselves
are not transactional. If a concurrent edit prevents rollback, the error
identifies it instead of overwriting that edit.

Use the reported backup ID to undo a successful operation. Restore verifies
the recorded OS/version/architecture and the current content and permissions
of each affected file. A subsequent edit blocks restore. Restoring also creates
a backup, allowing the restore itself to be undone. Keep the same `--root`
and `--platform`/`--os-version` settings when restoring a staged configuration.

PostgreSQL 19 is an active catalog and repository target, while generated
configurations keep PG18 as the default. PG19 beta definitions and PG14 RPM
repository definitions (`pgdg14`, `pgdg14-nonfree`) have been retired.

The current active targets are PG15–19 on EL9/10, Debian 12/13 and Ubuntu
22.04/24.04/26.04, plus PG15–18 on EL8, across both `x86_64` and `aarch64`.
`all`, `pgsql` and `pgdg` select stable PG15–19 repositories on EL9/10 (and
PG15–18 on EL8), alongside the shared APT `main` component on Debian/Ubuntu.
`pgext env repo`, `env tool`, `env pgrx` and build commands follow the active
platform/major policy. The catalog fetch/scan pipeline skips retired PG/OS
repositories.
