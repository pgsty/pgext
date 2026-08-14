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
| `sty boot` | Bootstrap the Pigsty controller | Requires root |
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
sudo pig sty boot                # prepare the controller and ~/pigsty
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

Bootstrap the Pigsty controller with the native Go workflow. The command requires root,
prepares a usable Ansible environment, supports online and offline repositories, repairs
common controller prerequisites, and reports a structured result. It never delegates to
Pigsty's legacy `bootstrap` script.

```bash
sudo pig sty boot                         # online bootstrap with the default region
sudo pig sty boot -r china                # use China-region repositories
sudo pig sty boot -m                      # equivalent to --region china
sudo pig sty boot -k                      # preserve existing repository definitions
sudo pig sty boot -p /path/to/pkg.tgz     # use an explicit offline package
sudo pig sty boot -p https://host/pkg.tgz # download and use an offline package
sudo pig sty boot -o json                 # machine-readable result and warnings
```

The native bootstrap performs these stages:

1. On Debian 12/13, check and repair `en_US.UTF-8` when possible so Ansible can start.
2. Verify `ansible-playbook` and its Python dependencies instead of checking only the binary.
3. Reuse a committed `/www/pigsty` repository, consume an explicit package or URL, accept a
   trusted automatic `/tmp/pkg.tgz`, or configure the selected online repositories.
4. Install the controller package set when Ansible is missing or unusable. Explicit offline
   input is still prepared when Ansible is already installed.
5. Unless `--keep` is set, back up repository definitions before replacement and restore them
   automatically when local or online package setup fails.
6. Best-effort repair key-based localhost SSH and initialize a missing default `~/pigsty` tree.
   These final convenience checks produce warnings rather than turning a usable controller into
   a false bootstrap failure.

**Options:**

- `-r|--region`: region, such as default, china, europe
- `-m|--mirror`: equivalent to `--region china`
- `-p|--path`: offline package file or HTTP(S) URL; an invalid explicit source is a hard error
- `-k|--keep`: preserve existing repositories instead of replacing them

An automatically discovered `/tmp/pkg.tgz` must be a regular, non-group/world-writable file
owned by root or the invoking sudo user. Unsafe automatic candidates are ignored with a warning.
Use `-o json` or `-o yaml` to consume bootstrap mode, repository policy, rollback state, locale,
SSH, Pigsty initialization status, warnings, and recommended next commands without scraping text.

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

**Options:**

- `-c|--conf`: template mode, equivalent to positional `[mode]`; the two forms are exclusive
- `--ip`: up to ten distinct comma-separated IPv4 addresses
- `--domain`: replace the exact `i.pigsty` placeholder domain
- `-v|--version`: PostgreSQL major version, 18/17/16/15/14; 19 beta can be specified explicitly
- `-r|--region`: upstream repository region, such as default/china/europe
- `-m|--mirror`: equivalent to `--region china`
- `-O|--output-file`: output config file path, default `pigsty.yml`
- `-s|--skip`: keep the placeholder IP and skip admin SSH/sudo preflight; exclusive with `--ip`
- `-p|--port`: SSH port
- `-x|--proxy`: write non-empty proxy environment variables into `all.vars.proxy_env`
- `-n|--non-interactive`: refuse ambiguous IP selection instead of prompting
- `-g|--generate`: replace known demo credentials with random 24-character values

`--ip` maps addresses in order to the template slots `10.10.10.10` through
`10.10.10.19`; replacement is simultaneous, and unrelated addresses such as VIPs remain intact.
Without `--ip`, interactive mode lists detected interfaces for selection; non-interactive or
closed-input execution fails with guidance to specify an address. `--domain` only replaces the
exact `i.pigsty` token, not names such as `cli.pigsty` or `i.pigsty.cc`.

Template modes are safe relative paths below `conf`; absolute paths, traversal, and path escape
are rejected. Output cannot alias the source template through a direct path, symlink, symlinked
parent, or hard link. The rendered Inventory is validated before the atomic `0600` write, so a
parse, mutation, preflight, or validation failure does not replace the destination. Structured
output reports applied IP mappings, effective PostgreSQL version, generated secret identifiers,
and warnings, but never generated secret values.

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
