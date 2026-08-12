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
  pig sty boot    [-rmpk]         # install ansible and prepare offline pkg
  pig sty conf    [-cvmrsoxnpg --raw] # configure pigsty and generate config
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
| `sty boot` | Install Ansible prerequisites | Requires sudo or root |
| `sty conf` | Generate configuration | |
| `sty deploy` | Run deployment playbook | |
| `sty list` | List available Pigsty versions | |
| `sty get` | Download Pigsty source tarball | |
| `sty grafana` | Manage Grafana dashboards (alias `gf`) | New in v1.6.0 |
{.full-width}

> Since v1.6.0, the former `pig sty edit` / `validate` / `check` commands moved to the
> root-level [`pig inventory`](https://pig.pgsty.com/inventory/) command group, and the experimental
> `pig sty dashboard` was replaced by `pig sty grafana`.


## Quick Start

Use `pig sty` to bootstrap and deploy Pigsty on the current node.

```bash
pig sty init                     # install Pigsty to ~/pigsty
pig sty boot                     # install Ansible prerequisites
pig sty conf                     # generate configuration
pig sty deploy                   # run deployment playbook
```

See the detailed setup guide: <https://pigsty.io/docs/setup/install/>


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

Install Ansible and its dependencies.

```bash
pig sty boot                     # install Ansible
pig sty boot -r china            # use China region mirror
pig sty boot -m                  # equivalent to --region china
pig sty boot -k                  # keep existing repositories
pig sty boot -p /path/to/pkg     # selected offline package path
```

**Options:**

- `-r|--region`: region, such as default, china, europe
- `-m|--mirror`: equivalent to `--region china`
- `-p|--path`: offline package path
- `-k|--keep`: keep existing repositories

See: <https://pigsty.io/docs/setup/offline/#bootstrap>


## sty conf

Configure Pigsty with `./configure` and generate the configuration file.

```bash
pig sty conf                       # use default meta.yml config
pig sty conf -g                    # generate random passwords (recommended!)
pig sty conf -c rich               # use conf/rich.yml template with more extensions
pig sty conf -c ha/full            # use conf/ha/full.yml 4-node HA template
pig sty conf -c slim               # use conf/slim.yml minimal template
pig sty conf -c supabase           # use conf/supabase.yml self-hosting template
pig sty conf -v 18 -c rich         # use conf/rich.yml template with PostgreSQL 18
pig sty conf -r china -s           # use China mirror and skip IP probing
pig sty conf -m -s                 # use mirror mode and skip IP probing
pig sty conf -x                    # write proxy settings from environment variables
pig sty conf -c full -g -O ha.yml  # full HA template, random passwords, output to ha.yml
pig sty conf --raw                 # use legacy shell configure workflow
```

**Options:**

- `-c|--conf`: configuration template name, such as meta/rich/slim/full/supabase
- `--ip`: primary node IP address
- `-v|--version`: PostgreSQL major version, 18/17/16/15/14; 19 beta can be specified explicitly
- `-r|--region`: upstream repository region, such as default/china/europe
- `-m|--mirror`: equivalent to `--region china`
- `-O|--output-file`: output config file path, default `pigsty.yml`
- `-s|--skip`: use a placeholder IP and skip the admin SSH/sudo preflight
- `-p|--port`: SSH port
- `-x|--proxy`: write proxy settings from environment variables
- `-n|--non-interactive`: non-interactive mode
- `-g|--generate`: generate random default passwords, recommended
- `--raw`: use the legacy shell configure workflow (generated passwords remain visible)

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
