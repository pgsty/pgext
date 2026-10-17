---
title: "CMD: sty"
description: How to manage pigsty with pig sty subcommand
icon: fa-solid fa-terminal
weight: 630
---

`pig sty` installs Pigsty, prepares the controller, generates configuration, and runs deployment playbooks. It also manages dashboards through the Grafana HTTP API.

| Command | Purpose |
|:---|:---|
| `sty init` | Download and extract Pigsty into an installation directory |
| `sty boot` | Install Ansible dependencies and prepare the controller |
| `sty conf` | Generate and validate an inventory from a template |
| `sty deploy` | Run the deployment playbook |
| `sty list` / `sty get` | List or download Pigsty releases |
| `sty grafana` | Manage Grafana dashboards, language, and theme; alias `gf` |

## Quick Start

Run these commands on the controller:

```bash
pig sty init                     # download Pigsty
pig sty boot                     # prepare Ansible and the controller
pig sty conf -g                  # generate configuration with random passwords
pig inventory edit              # review target hosts and deployment settings
pig sty deploy                   # run the deployment
```

`deploy` starts the playbook immediately and changes the hosts selected by the inventory. See the [Pigsty installation guide](https://pigsty.io/docs/setup/install/) for the full procedure.

`boot` also attempts to initialize a missing `~/pigsty` directory. Run `init` first to select a version or path. Use [`pig inventory`](/pig/cmd/inventory/) to edit and validate an existing configuration.

## sty init

Download and install the Pigsty distribution into `~/pigsty`.

```bash
pig sty init                   # install latest version to ~/pigsty
pig sty init -f                # install and overwrite existing pigsty directory
pig sty init -m                # prefer the pigsty.cc mirror for downloads
pig sty init -p /tmp/pigsty    # install to selected directory /tmp/pigsty
pig sty init -v {{< param pigsty_version >}}       # install this documented Pigsty version
pig sty init 4                 # install the latest release in the 4.x series
```

**Options:**

- `-p|--path`: target installation directory, default `~/pigsty`
- `-f|--force`: force overwrite of existing pigsty directory
- `-m|--mirror`: prefer the `pigsty.cc` mirror
- `-v|--version`: Pigsty version
- `-d|--dir`: download directory, default `/tmp`
- `--accept-eula`: accept the Pigsty Pro EULA without an interactive prompt; applies to Pro releases

## sty boot

Prepare Ansible and its Python dependencies, configure online or offline repositories, and check localhost SSH and the Pigsty directory.

```bash
pig sty boot                         # online bootstrap with the default region
pig sty boot -r china                # use China-region repositories
pig sty boot -m                      # equivalent to --region china
pig sty boot -k                      # preserve existing repository definitions
pig sty boot -p /path/to/pkg.tgz     # use an explicit offline package
pig sty boot -p https://host/pkg.tgz # download and use an offline package
pig sty boot -o json                 # machine-readable result and warnings
```

You may invoke the command without `sudo`. Pig resolves and downloads an explicit source first,
then re-executes itself through `sudo` once when root access is required. Set `PIG_NO_SUDO=1` to
disable automatic elevation. JSON/YAML output and `PIG_NON_INTERACTIVE=1` make sudo
non-interactive; if elevation requires a password, the command fails without waiting for input.

### Bootstrap stages

1. On Debian 12/13, check and attempt to repair the `en_US.UTF-8` locale.
2. Check that Ansible runs and that its Python environment provides `yaml`, `jmespath`, and either `cryptography` or `OpenSSL`.
3. Prepare the selected offline repository. If Ansible is unusable, install controller dependencies and check it again.
4. Check helper tools, prepare key-based SSH to `127.0.0.1` for the invoking administrator, and attempt to initialize a missing `~/pigsty`.

PIG prepares a selected or discovered offline repository even when Ansible already works. Use this command to stage offline packages on a ready controller.

### Source selection and modes

The result records one of four bootstrap modes:

| Mode | Meaning |
|------|---------|
| `ready` | Ansible was already usable and no offline source needed preparation. |
| `offline` | An explicit, discovered, or existing offline repository was selected. |
| `online` | Regional online repositories were configured to repair the controller. |
| `existing` | An online refresh failed under `--keep`, and existing repository definitions were successfully used as a fallback. |

PIG selects a source as follows:

- `--path` accepts a local archive or an HTTP(S) URL. URL credentials are rejected, and a bad
  explicit source is a hard error rather than an online fallback.
- An automatically discovered `/tmp/pkg.tgz` must be a regular, non-group/world-writable file
  owned by root or the invoking sudo user. Unsafe candidates are ignored with a warning.
- A complete `/www/pigsty` repository takes precedence over a selected package. If both exist,
  the existing repository is reused and the package is left untouched with a warning.
- Archives are downloaded and extracted by Pig itself. If `/www` does not exist, Pig creates
  `/data/nginx` and the expected `/www -> /data/nginx` symbolic link before committing content.

Offline mode enables only the strict `pigsty-local` repository. Online mode configures the
selected region and installs the `node` and `pigsty` controller modules. It reuses or installs
Pigsty's embedded signing key and enables signature checks when key preparation succeeds.
If implicit-key preparation fails, it records a warning and continues with the selected Pigsty
repositories' signature checks disabled. With explicit `gpgkey` or `signed-by` metadata,
failure to prepare the referenced key is fatal; see [`pig repo`](/pig/cmd/repo/#trust-and-configuration-ownership).

### Repository replacement and recovery

By default, PIG backs up existing definitions before writing new ones. If repository setup or package installation fails, it attempts to restore the backup and reports the result. `--keep` preserves existing definitions and allows a fallback to them if an online refresh fails.

Ansible installation or validation failure makes the command fail. Incomplete SSH, locale, or Pigsty-directory setup produces warnings.

**Options:**

- `-r|--region`: region, such as default, china, europe
- `-m|--mirror`: equivalent to `--region china`; mutually exclusive with `--region`
- `-p|--path`: offline package file or HTTP(S) URL; an invalid explicit source is a hard error
- `-k|--keep`: preserve existing repositories instead of replacing them

### Structured output

`-o json` or `-o yaml` returns a `pig.sty.boot/v2` result with Ansible state, mode, sources, repository recovery results, localhost SSH, Pigsty directory status, and warnings.

See [Pigsty offline installation](https://pigsty.io/docs/setup/offline/#bootstrap) for the offline procedure.

## sty conf

Generate an inventory from a template under `<PIGSTY_HOME>/conf`. The command edits the YAML configuration, validates the complete result, and writes the output atomically with mode `0600`.

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
pig sty conf ha/full -g -O ha.yml       # custom owner-only output file
pig sty conf -n --ip 10.0.0.10 -o json
```

The default mode is `meta`; `pig sty c` and `pig sty configure` are command aliases. Note that
`-O` selects the Inventory output file, while the global lowercase `-o` selects text, JSON, or
YAML command output.

### Template and output paths

Template names are relative to `<PIGSTY_HOME>/conf`. The `.yml` suffix is optional; for example, use `ha/full`. Use a relative name without `..`.

Relative output paths resolve under `<PIGSTY_HOME>`. Absolute paths are also accepted. The output must be a different file from the template and cannot be a symlink. A parsing, transformation, or validation error leaves the existing output unchanged.

### Structural mutations

The command applies these changes to the template:

| Input | Configuration change |
|-------|-----------------|
| `--ip A,B,...` | Replaces `10.10.10.10` through `10.10.10.19` with up to ten distinct addresses, in order. Replacements are simultaneous; other VIPs remain unchanged. |
| no `--ip` | Detects local interfaces. Interactive mode asks when selection is ambiguous; `--non-interactive` or closed stdin fails with guidance to pass `--ip`. |
| `--domain NAME` | Replaces only the exact `i.pigsty` token, not names such as `cli.pigsty` or `i.pigsty.cc`; NAME must be a valid DNS domain. |
| small controller | When detected CPU count is below four, rewrites `node_tune: oltp` and `pg_conf: oltp.yml` to their `tiny` profiles. |
| `--region REGION` | Updates `all.vars.region` for non-default regions. `china` also activates template-provided Docker and pip mirrors, using values already present in the template. |
| `--proxy` | Reads `HTTP_PROXY`, `HTTPS_PROXY`, `ALL_PROXY`, and `NO_PROXY`, each falling back to its lowercase spelling. Missing HTTP/HTTPS values fall back to ALL_PROXY. Writes non-empty values to `all.vars.proxy_env` and supplies a default no-proxy list when needed. |
| `--version MAJOR` | Updates generic templates for PostgreSQL 14-18, or explicit 19 beta, and selects the matching locale. Version-pinned `mssql`, `polar`, and `pgNN` modes keep their template version and report a warning. |
| `--generate` | Generates one 24-character value per known credential identifier and consistently replaces its active values and documented placeholders. |

An IP mapping that creates a duplicate key fails. PIG warns if a supplied address has no matching placeholder in the template.

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

PIG checks the operating system, package manager, controller resources, SSH, sudo, and Ansible. `--port` selects the SSH check port. `--skip` bypasses administrator SSH/sudo preflight and preserves placeholder IPs; other checks still run.

The output must pass inventory validation. If `ansible-inventory` is available, PIG also uses it to parse the result. Argument, template, and validation errors prevent writing. Issues that allow configuration to continue are reported as warnings.

### Structured output

`-o json` or `-o yaml` returns `pig.sty.configure/v1`, including the template, output path, region, IPs, domain, SSH port, PostgreSQL versions, and warnings. Generated passwords are written only to the inventory; the result lists their identifiers.

See the [Pigsty configuration guide](https://pigsty.io/docs/setup/install/#configure) for the procedure.

## sty deploy

`sty deploy` uses the Inventory source already selected by configuration. `-i` and
`PIGSTY_INVENTORY` relative paths resolve from the invoking directory; quoted `~/` expands from
the user's home. The selected absolute path is passed with `-i`, including the default source,
while Ansible runs in Pigsty home. A missing source fails without fallback. Deploy still
requires a regular file; dynamic source execution remains Ansible's responsibility.

Deploy Pigsty with the `deploy.yml` playbook.

```bash
pig sty deploy       # run deploy.yml, falling back to install.yml if not found
pig sty d            # short alias
pig sty de           # short alias
```

This command runs the `deploy.yml` playbook from your Pigsty installation directory. For backward compatibility, if `deploy.yml` does not exist but `install.yml` exists, `install.yml` is used instead.

The playbook starts immediately, without a confirmation prompt. Review the selected inventory before running it.

## sty list

List available Pigsty versions.

```bash
pig sty list                     # list recent versions
pig sty list all                 # list all available versions
pig sty list 4                   # list from the latest stable 4.x version onward
```

A version prefix selects the highest matching stable version as the starting point, using the
same completion rule as `sty get`. Matching respects component boundaries (`1.1` does not match
`1.10`). Exact versions, including prereleases, and `latest` keep their existing meaning.

## sty get

Download the Pigsty source tarball.

```bash
pig sty get                      # download latest version
pig sty get {{< param pigsty_version >}}       # download the documented Pigsty version
pig sty get -m                   # prefer the pigsty.cc mirror
```

## sty grafana

`pig sty grafana` (alias `gf`) manages dashboards through the HTTP API. The optional `PATH` selects a Grafana root directory, one direct subdirectory, or one dashboard JSON file. The default is `<PIGSTY_HOME>/files/grafana`.

```bash
pig sty grafana info             # check Grafana health, authentication, and basic info
pig sty grafana list             # list all dashboards in the active organization
pig sty grafana boot             # bootstrap Grafana around the existing pigsty dashboards
pig sty grafana init             # load the complete dashboard corpus, then bootstrap
pig sty grafana load [PATH]      # load dashboards selected by a local path
pig sty grafana dump [PATH]      # export remote dashboards into a local path
pig sty grafana clean [PATH]     # delete dashboards selected by a path (--dry-run/--yes)
pig sty grafana lang zh-Hans     # set the organization and current-user language
pig sty grafana style dark       # set the organization and current-user visual style
```

### Select the operation scope

- `list` shows all dashboards in the current organization.
- `load` creates or updates dashboards by UID from local files and preserves other remote dashboards.
- `boot` renames the current organization to Pigsty, stars the existing `pigsty` dashboard, and makes it the organization and current-user home.
- `init` loads the complete dashboard root, then runs `boot`. Use `load` for a single folder or file.
- `dump -d DIR` exports to a chosen directory. `--dir` and positional `PATH` are mutually exclusive.
- `clean PATH` deletes remote dashboards declared by that local path. `clean --all` deletes every dashboard and folder in the current organization.
- `style` accepts `dark`, `light`, or `auto`.

Preview the deletion scope, then confirm it:

```bash
pig sty grafana clean --dry-run
pig sty grafana clean --yes
```

**Connection and credentials**:

| Option            | Description                                                     |
|:------------------|:----------------------------------------------------------------|
| `--endpoint`      | Grafana origin and optional path prefix (default `http://i.pigsty/ui`) |
| `--username`      | Grafana API username                                            |
| `--password`      | Grafana API password (**insecure**: visible in process arguments and shell history) |
| `--password-file` | Owner-only file containing the password (recommended)           |
{.full-width}

Choose `--password` or `--password-file`; they are mutually exclusive. Without either option, PIG uses `GRAFANA_PASSWORD`, then `all.vars.grafana_admin_password` from the inventory.
The HTTP client enforces timeouts and response-size limits, refuses redirects, and
verifies TLS certificates by default.

### Legacy dashboards and schema v2 resources

`load` and `init` accept both classic Grafana dashboard JSON and the resource form whose identity is exactly:

```json
{"apiVersion":"dashboard.grafana.app/v2","kind":"Dashboard","metadata":{"name":"pgsql-overview","namespace":"default"},"spec":{}}
```

Each format uses its corresponding API:

- Classic JSON takes its UID from top-level `uid` and uses the legacy dashboard API.
- Schema v2 takes its UID from `metadata.name`, defaults a missing namespace to `default`, requires an object-valued `spec`, and uses Grafana's dashboard resource API.
- The JSON filename (without `.json`) must match the resolved UID. Only one local folder layer is allowed; its directory name becomes the Grafana folder UID.
- For schema v2, PIG preserves `spec` and the `grafana.app/message` annotation, sets `grafana.app/folder` from the local folder, and deliberately removes server-managed metadata/status before upsert.
- `dump` preserves schema v2 only when the target file already exists locally in v2 form; it then fetches the native v2 resource using that file's namespace. A new dump target defaults to classic JSON. Local-only files are never deleted by `dump`.

Other `dashboard.grafana.app/*` versions are unsupported. To preserve v2 format during export, keep the existing local v2 file.

## Design records

- [Native sty boot](https://pig.pgsty.com/design/native-sty-boot/)
- [Native sty conf](https://pig.pgsty.com/design/native-sty-conf/)
- [Native Grafana Management](https://pig.pgsty.com/design/native-grafana-management/)
- [Use the CMDB Pigsty Already Has](https://pig.pgsty.com/design/existing-cmdb-boundary/)
