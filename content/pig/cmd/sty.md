---
title: "CMD: sty"
description: How to manage pigsty with pig sty subcommand
icon: SquareTerminal
weight: 630
---

**pig** can also be used as a CLI tool for Pigsty, a batteries-included free PostgreSQL RDS solution.
It brings HA, PITR, monitoring, infrastructure as code (IaC), and rich extension support to your PostgreSQL clusters.

```bash
pig sty - Init (Download), Bootstrap, Configure, and Deploy Pigsty

  pig sty init    [-mpfvd]        # install pigsty (~/pigsty by default)
  pig sty boot    [-rmpk]         # native controller bootstrap
  pig sty conf    [mode] [flags]  # native Inventory-aware configuration
  pig sty deploy                  # use pigsty to deploy everything (CAUTION!)
  pig sty get                     # download pigsty source tarball
  pig sty list                    # list available pigsty versions
  pig sty grafana <verb>          # manage grafana dashboards (native HTTP)

Examples:
  pig sty init                 # extract and init ~/pigsty
  pig sty boot                 # install ansible & other deps
  pig sty conf                 # generate pigsty.yml config file
  pig sty deploy               # run the deploy.yml playbook
```

| Command | Description | Notes |
|:---|:---|:---|
| `sty init` | Install Pigsty | |
| `sty boot` | Bootstrap the Pigsty controller | Auto-elevates when needed |
| `sty conf` | Generate and validate Inventory | Native Go workflow |
| `sty deploy` | Run deployment playbook | |
| `sty list` | List available Pigsty versions | |
| `sty get` | Download Pigsty source tarball | |
| `sty grafana` | Manage Grafana dashboards (alias `gf`) | New in v1.6.0 |
{.full-width}

> Since v1.8.0, `pig sty boot` and `pig sty conf` are native Go workflows. They no longer
> invoke Pigsty's legacy `bootstrap` or `configure` shell scripts. Since v1.6.0, the former
> `pig sty edit` / `validate` / `check` commands moved to the
> root-level [`pig inventory`](https://pig.pgsty.com/inventory/) command group, and the experimental
> `pig sty dashboard` was replaced by `pig sty grafana`.

## Quick Start

Use `pig sty` to bootstrap and deploy Pigsty on the current node.

```bash
pig sty boot                     # prepare the controller; auto-elevates when needed
pig sty conf -g                  # generate and validate pigsty.yml
pig inventory edit              # optional: review and adjust the Inventory
pig sty deploy                   # run deployment playbook
```

See the detailed setup guide: <https://pigsty.io/docs/setup/install/>

`sty boot` initializes a missing default `~/pigsty` tree on a best-effort basis. Use
`pig sty init` first when you need to select an explicit Pigsty version or installation path.

## sty init

Download and install the Pigsty distribution into `~/pigsty`.

```bash
pig sty init                   # install latest version to ~/pigsty
pig sty init -f                # install and overwrite existing pigsty directory
pig sty init -m                # prefer the pigsty.cc mirror for downloads
pig sty init -p /tmp/pigsty    # install to selected directory /tmp/pigsty
pig sty init -v 3.4            # fetch and install selected version v3.4.1
pig sty init 3                 # fetch and install latest v3 major version
```

**Options:**

- `-p|--path`: target installation directory, default `~/pigsty`
- `-f|--force`: force overwrite of existing pigsty directory
- `-m|--mirror`: prefer the `pigsty.cc` mirror
- `-v|--version`: Pigsty version
- `-d|--dir`: download directory, default `/tmp`

## sty boot

Bootstrap the Pigsty controller with the native Go workflow. The command prepares a usable
Ansible environment, supports online and offline repositories, repairs common controller
prerequisites, and reports a structured result. It never delegates to Pigsty's legacy
`bootstrap` script and does not require `curl`, `wget`, `tar`, or `gzip` for package download
and extraction.

```bash
sudo pig sty boot                         # online bootstrap with the default region
sudo pig sty boot -r china                # use China-region repositories
sudo pig sty boot -m                      # equivalent to --region china
sudo pig sty boot -k                      # preserve existing repository definitions
sudo pig sty boot -p /path/to/pkg.tgz     # use an explicit offline package
sudo pig sty boot -p https://host/pkg.tgz # download and use an offline package
sudo pig sty boot -o json                 # machine-readable result and warnings
```

You may invoke the command without `sudo`. Pig resolves and downloads an explicit source first,
then re-executes itself through `sudo` once when root access is required. Set `PIG_NO_SUDO=1` to
disable automatic elevation, or `PIG_NON_INTERACTIVE=1` to make the sudo attempt non-interactive.

### Bootstrap stages

The native workflow performs these stages:

1. On Debian 12/13, check and repair `en_US.UTF-8` when possible so Ansible can start.
2. Execute `ansible-playbook`, discover its Python interpreter, and verify `yaml`, `jmespath`,
   and either `cryptography` or `OpenSSL`. A binary that cannot actually run is not considered ready.
3. Resolve the repository source, prepare offline content when selected, and install the lean
   controller package set only when Ansible is missing or unusable.
4. Verify Ansible again after package installation and retry locale preparation when packages may
   have made locale tools available.
5. Probe controller helpers, repair key-based SSH to `127.0.0.1` for the invoking admin user, and
   initialize a missing default `~/pigsty` tree when possible.

Explicit, automatically discovered, and already committed offline sources are prepared even when
Ansible is already usable. This makes `sty boot` suitable for staging an offline repository on an
otherwise ready controller.

### Source selection and modes

The result records one of four bootstrap modes:

| Mode | Meaning |
|------|---------|
| `ready` | Ansible was already usable and no offline source needed preparation. |
| `offline` | An explicit, trusted automatic, or committed offline repository was selected. |
| `online` | Regional online repositories were configured to repair the controller. |
| `existing` | An online refresh failed under `--keep`, and existing repository definitions were successfully used as a fallback. |

Source precedence and safety rules are deterministic:

- `--path` accepts a local archive or an HTTP(S) URL. URL credentials are rejected, and a bad
  explicit source is a hard error rather than an online fallback.
- An automatically discovered `/tmp/pkg.tgz` must be a regular, non-group/world-writable file
  owned by root or the invoking sudo user. Unsafe candidates are ignored with a warning.
- A completed `/www/pigsty` repository takes precedence over a selected package. If both exist,
  the committed repository is reused and the package is left untouched with a warning.
- Archives are downloaded and extracted by Pig itself. If `/www` does not exist, Pig creates
  `/data/nginx` and the expected `/www -> /data/nginx` symbolic link before committing content.

Offline mode enables only the strict `pigsty-local` repository. Online mode configures the
selected region, installs Pigsty's embedded signing key with repository signature checks enabled,
and installs the `node` and `pigsty` controller modules.

### Repository transaction and failure boundary

By default, repository definitions are backed up before replacement. If repository setup or
package installation fails, Pig attempts to restore the backup and reports whether rollback
succeeded or failed. `--keep` changes the policy to additive operation: existing definitions are
preserved, online refresh may fall back to them, and no replacement rollback is required.

Invalid explicit sources, unsupported package management when installation is required,
repository/package failures, and an unusable post-install Ansible runtime are command failures.
Locale repair, optional helper probes, localhost SSH repair, and Pigsty tree initialization are
advisory finishing steps; their failures are retained as warnings in an otherwise usable result.

**Options:**

- `-r|--region`: region, such as default, china, europe
- `-m|--mirror`: equivalent to `--region china`; mutually exclusive with `--region`
- `-p|--path`: offline package file or HTTP(S) URL; an invalid explicit source is a hard error
- `-k|--keep`: preserve existing repositories instead of replacing them

### Structured output

Use the global `-o json` or `-o yaml` flag for automation. The payload kind is
`pig.sty.boot/v2` and includes the Ansible state, selected mode and package manager, repository
policy and rollback result, source and repository paths, locale, localhost SSH and Pigsty tree
initialization status, whether the machine changed, warnings, and these recommended next steps:

```bash
pig sty conf -g
pig inv edit
pig sty deploy
```

Dynamic progress is suppressed in structured mode, so stdout remains machine-readable.

See: <https://pigsty.io/docs/setup/offline/#bootstrap>

## sty conf

Generate Pigsty Inventory through the native Go workflow. `sty conf` reads one template below
`<PIGSTY_HOME>/conf`, applies bounded structural mutations, validates the complete candidate,
and atomically writes an owner-only Inventory. It does not invoke or fall back to `./configure`.

```bash
pig sty conf                         # use conf/meta.yml and write pigsty.yml
pig sty conf -g                      # generate random passwords (recommended)
pig sty conf rich                    # positional mode selects conf/rich.yml
pig sty conf -c ha/full              # equivalent flag form; do not combine both forms
pig sty conf ha/trio --ip 10.0.0.10,10.0.0.11,10.0.0.12
pig sty conf --domain infra.example.com
pig sty conf rich -v 18              # request PostgreSQL 18 for a generic template
pig sty conf -r china -s             # China region, placeholder IP, no admin preflight
pig sty conf -x                       # materialize proxy environment variables
pig sty conf full -g -O ha.yml       # custom owner-only output file
pig sty conf -n --ip 10.0.0.10 -o json
```

The default mode is `meta`; `pig sty c` and `pig sty configure` are command aliases. Note that
`-O` selects the Inventory output file, while the global lowercase `-o` selects text, JSON, or
YAML command output.

### Template and output safety

- A mode is a slash-separated relative name below `<PIGSTY_HOME>/conf`; `.yml` is optional.
  Absolute paths, traversal, empty segments, and path escape are rejected.
- Relative output paths are resolved below `<PIGSTY_HOME>`; absolute output paths are retained.
- The destination cannot alias the source through the same path, an existing symlink, a
  symlinked parent, or a hard link. An existing output symlink is always refused.
- Pig parses the source and rejects conflicting IP mappings before running external preflight
  checks. A parse, mutation, preflight, or validation failure leaves the destination unchanged.
- A successful result is written atomically with mode `0600`.

### Structural mutations

The command edits parsed YAML structures and bounded scalar tokens rather than applying broad
text substitutions:

| Input | Native behavior |
|-------|-----------------|
| `--ip A,B,...` | Maps up to ten distinct addresses, in order, to `10.10.10.10` through `10.10.10.19`. Replacement is simultaneous, so swaps are safe and unrelated VIPs remain intact. |
| no `--ip` | Detects local interfaces. Interactive mode asks when selection is ambiguous; `--non-interactive` or closed stdin fails with guidance to pass `--ip`. |
| `--domain NAME` | Replaces only the exact `i.pigsty` token, not names such as `cli.pigsty` or `i.pigsty.cc`; NAME must be a valid DNS domain. |
| small controller | When detected CPU count is below four, rewrites `node_tune: oltp` and `pg_conf: oltp.yml` to their `tiny` profiles. |
| `--region REGION` | Updates `all.vars.region` for non-default regions. `china` also activates template-provided Docker and pip mirrors, but never invents values absent from the template. |
| `--proxy` | Writes non-empty `HTTP_PROXY`/`http_proxy`, `HTTPS_PROXY` (falling back to `ALL_PROXY`), `ALL_PROXY`, and `NO_PROXY` values into `all.vars.proxy_env`; a safe default no-proxy list is used when needed. |
| `--version MAJOR` | Updates generic templates for PostgreSQL 14-18, or explicit 19 beta, and selects the matching locale. Version-pinned `mssql`, `polar`, and `pgNN` modes keep their template version and report a warning. |
| `--generate` | Generates one 24-character value per known credential identifier and consistently replaces its active values and documented placeholders. |

An IP mapping that would collide with an unreplaced Inventory key is rejected as an invalid
argument. A supplied address with no matching placeholder slot is retained in structured output
as a discarded-IP warning rather than being silently ignored.

For PostgreSQL 19 beta, Pig also enables the `beta` repository after `pgsql` when the template
contains the expected repository list. Modes below `conf/build/` intentionally bypass IP mapping
and controller-admin preflight so build templates remain portable.

Active password identifiers are `grafana_admin_password`, `pg_admin_password`,
`pg_monitor_password`, `pg_replication_password`, `patroni_password`,
`haproxy_admin_password`, `minio_secret_key`, and `etcd_root_password`. Generation also covers
the documented `DBUser.Meta`, `DBUser.Viewer`, `S3User.Backup`, `S3User.Meta`, `S3User.Data`,
`DBUser.Supa`, and `Vibe.Coding` placeholders. Repeated occurrences of one identifier receive the
same generated value.

**Options:**

- `-c|--conf`: template mode, equivalent to positional `[mode]`; the two forms are exclusive
- `--ip`: up to ten distinct comma-separated IPv4 addresses
- `--domain`: replace the exact `i.pigsty` placeholder with a valid DNS domain
- `-v|--version`: PostgreSQL major version, 18/17/16/15/14; 19 beta can be specified explicitly
- `-r|--region`: upstream repository region, such as default/china/europe
- `-m|--mirror`: equivalent to `--region china`; mutually exclusive with `--region`
- `-O|--output-file`: output config file path, default `pigsty.yml`
- `-s|--skip`: keep the placeholder IP and skip admin SSH/sudo preflight; exclusive with `--ip`
- `-p|--port`: SSH port
- `-x|--proxy`: write non-empty proxy environment variables into `all.vars.proxy_env`
- `-n|--non-interactive`: refuse ambiguous IP selection instead of prompting
- `-g|--generate`: replace known demo credentials with random 24-character values

### Preflight and validation

Unless `--skip` is used, Pig inspects the kernel, architecture, package manager, platform vendor,
controller resources, sudo/admin access, localhost SSH, and Ansible availability. `--port` is used
for the SSH check. These diagnostics are returned as actionable warnings where the Inventory can
still be generated; invalid arguments and unsafe configuration transformations remain errors.

The rendered candidate must pass Pig's native Inventory validation. When `ansible-inventory` is
available, Pig also runs a bounded external parse before committing the file. `--skip` preserves
placeholder IPs and bypasses admin SSH/sudo preflight, but it does not disable template parsing,
safe mutation, Inventory validation, or atomic writing.

### Structured output

With global `-o json` or `-o yaml`, the payload kind is `pig.sty.configure/v1`. It reports the
mode, source and output paths, region, chosen primary address, applied and discarded IPs, domain,
SSH port, requested and effective PostgreSQL versions, native-workflow marker, generated secret
identifiers, and warnings. Secret values are never printed.

See: <https://pigsty.io/docs/setup/install/#configure>

## sty deploy

Deploy Pigsty with the `deploy.yml` playbook.

```bash
pig sty deploy       # run deploy.yml, falling back to install.yml if not found
pig sty d            # short alias
pig sty de           # short alias
```

This command runs the `deploy.yml` playbook from your Pigsty installation directory. For backward compatibility, if `deploy.yml` does not exist but `install.yml` exists, `install.yml` is used instead.

> **Warning**: This operation modifies your system, and **invocation is explicit consent** —
> deploy starts immediately without a `--yes` gate; use Ctrl+C to interrupt a mistaken run.
> (The `pig sty install` / `ins` aliases were removed in v1.6.0.)

## sty list

List available Pigsty versions.

```bash
pig sty list                     # list available versions
```

## sty get

Download the Pigsty source tarball.

```bash
pig sty get                      # download latest version
pig sty get v3.4.0               # download selected version
pig sty get -m                   # prefer the pigsty.cc mirror
```

## sty grafana

Since v1.6.0, `pig sty grafana` (alias `gf`) manages Grafana dashboards through the native
HTTP API, replacing the experimental `pig sty dashboard`. The `PATH` argument may select the
grafana root, one direct folder, or one dashboard JSON file; without `PATH`, pig resolves
`<PIGSTY_HOME>/files/grafana` and never falls back to the current directory.

```bash
pig sty grafana info             # check Grafana health, authentication, and basic info
pig sty grafana list             # list all dashboards in the active organization
pig sty grafana boot             # bootstrap Grafana around the existing pigsty dashboards
pig sty grafana init             # load the complete dashboard corpus, then bootstrap
pig sty grafana load [PATH]      # load dashboards selected by a local path
pig sty grafana dump [PATH]      # export remote dashboards into a local path
pig sty grafana clean [PATH]     # delete dashboards selected by a path (--dry-run/--yes)
pig sty grafana lang zh-Hans     # set the organization and current-user language
pig sty grafana style            # set the organization and current-user visual style
```

**Connection and credentials**:

| Option            | Description                                                     |
|:------------------|:----------------------------------------------------------------|
| `--endpoint`      | Grafana origin and optional path prefix (default `http://i.pigsty/ui`) |
| `--username`      | Grafana API username                                            |
| `--password`      | Grafana API password (**insecure**: visible in process arguments and shell history) |
| `--password-file` | Owner-only file containing the password (recommended)           |
{.full-width}

Password resolution order: `--password` → `--password-file` → the `GRAFANA_PASSWORD`
environment variable → `all.vars.grafana_admin_password` from the inventory.
The HTTP client enforces timeouts and response-size limits, refuses redirects, and
verifies TLS certificates by default.

### Legacy dashboards and schema v2 resources

`load` and `init` accept both classic Grafana dashboard JSON and the resource form whose identity is exactly:

```json
{"apiVersion":"dashboard.grafana.app/v2","kind":"Dashboard","metadata":{"name":"pgsql-overview","namespace":"default"},"spec":{}}
```

The two formats are not flattened into one another during loading:

- Classic JSON takes its UID from top-level `uid` and uses the legacy dashboard API.
- Schema v2 takes its UID from `metadata.name`, defaults a missing namespace to `default`, requires an object-valued `spec`, and uses Grafana's dashboard resource API.
- The JSON filename (without `.json`) must match the resolved UID. Only one local folder layer is allowed; its directory name becomes the Grafana folder UID.
- For schema v2, PIG preserves `spec` and the `grafana.app/message` annotation, sets `grafana.app/folder` from the local folder, and deliberately removes server-managed metadata/status before upsert.
- `dump` preserves schema v2 only when the target file already exists locally in v2 form; it then fetches the native v2 resource using that file's namespace. A new dump target defaults to classic JSON. Local-only files are never deleted by `dump`.

Other `dashboard.grafana.app/*` versions or malformed resource envelopes are rejected instead of being silently treated as classic dashboards. Format preservation therefore depends on keeping the existing local file when round-tripping v2 resources.
