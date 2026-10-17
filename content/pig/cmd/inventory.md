---
title: "CMD: inventory"
description: Inspect, edit, validate, check, and exchange the Pigsty Inventory
icon: fa-solid fa-terminal
weight: 645
---

`pig inventory` (alias `pig inv`) inspects, edits, and validates the Pigsty inventory, `pigsty.yml`. It also checks deployment prerequisites and exchanges configuration with PostgreSQL through the experimental `cmdb` commands.

When you edit a fragment, PIG preserves the YAML content and formatting outside that fragment. Before saving, it parses the complete file again. If parsing fails or another process has changed the original file, PIG keeps the original and reports an error.

| Command        | Alias | Description                                             |
|:---------------|:------|:--------------------------------------------------------|
| `inv status`   |       | Inspect the active inventory source (without executing) |
| `inv list`     | `ls`  | List inventory topology and value kinds (`--depth`)     |
| `inv show`     |       | Show verbatim inventory YAML (**may contain secrets**)  |
| `inv edit`     | `e`   | Edit the inventory or one fragment in `$EDITOR`         |
| `inv validate` | `v`   | Validate one complete static Pigsty inventory           |
| `inv check`    | `ck`  | Check inventory, controller, and target readiness       |
| `inv diff`     |       | Compare declarations of two inventories (value-free)    |
| `inv cmdb`     |       | Exchange with the PostgreSQL CMDB (**experimental**)    |
{.full-width}

Use global `-i, --inventory` to select the inventory or `-H, --home` to select the Pigsty directory. Without these options, PIG resolves the source from its configuration and the Pigsty directory. Commands support structured results with `-o json` or `-o yaml`, except `show`, which always prints the selected YAML verbatim.

## Source selection {#source-selection}

PIG selects the first configured source in this order. If a selected source is missing or invalid,
PIG reports that problem instead of falling through to another inventory.

| Priority | Source | Base for a relative path |
|:---|:---|:---|
| 1 | `-i, --inventory PATH` | Current working directory |
| 2 | `PIGSTY_INVENTORY` | Current working directory |
| 3 | `inventory` in `~/.pig/config.yml` | Resolved Pigsty home |
| 4 | `[defaults] inventory` in the Pigsty home's `ansible.cfg` | Directory containing `ansible.cfg` |
| 5 | `pigsty.yml` | Resolved Pigsty home |

PIG expands `~/` before resolving the path, including a quoted path in `ansible.cfg`.
`pig inv status` reports the selected source without executing it. Static Inventory operations
require a supported YAML source. [`pig do`](/pig/cmd/do/) can forward a directory or executable inventory
to Ansible; [`pig sty deploy`](/pig/cmd/sty/#sty-deploy) requires a regular file. Both pass the resolved
path explicitly to Ansible, so it cannot silently select a different default.

## Quick Start

```bash
pig inventory status                     # which inventory source is active?
pig inventory list                       # browse the inventory topology
pig inventory edit pg-meta               # edit only the pg-meta cluster fragment
pig inventory validate                   # full validation
pig inventory check -p ssh --sudo        # add SSH/sudo readiness probes
pig inventory diff /path/other.yml       # compare declarations with another file
```

## inv status / list / show

```bash
pig inv status                          # show the active path and source type
pig inv list --depth 2                  # limit topology depth
pig inv list pg-meta                    # inspect one cluster's structure
pig inv show pg-meta                    # print the cluster YAML verbatim
```

`status` inspects the source without executing a dynamic inventory. `list` shows groups, hosts, keys, and value types. `show` prints actual values, including any passwords in the selected configuration. Adding `-o json` does not change its YAML output.

## Selectors

`list` / `show` / `edit` accept an optional **selector** that addresses one fragment:

```bash
pig inv e                                # the whole inventory
pig inv e vars                           # all.vars global parameters
pig inv e pg-meta                        # one cluster group
pig inv e pg-meta.vars.pg_databases      # a nested key path
pig inv e 'pg-meta.hosts["10.10.10.10"]' # dotted keys (e.g. IPs) via quoted brackets
```

## inv edit

Open the selected fragment in `$EDITOR`. After you save and exit, PIG restores the indentation and newline style, parses the complete document, checks for concurrent changes, and replaces the original file atomically. If parsing fails, PIG keeps the temporary file so you can correct it.

```bash
pig inventory edit                       # edit the whole inventory
pig inv e pg-meta                        # edit one fragment
pig inv e vars --from vars.yml           # replace the fragment from a file (or stdin: -)
```

| Option   | Description                                                        |
|:---------|:-------------------------------------------------------------------|
| `--from` | Replace the selected fragment from a regular file or stdin (`-`), skipping the editor |
{.full-width}

The saved file belongs to the current user. PIG removes group and other permissions, usually producing mode `0600`, and preserves stricter owner-only permissions. The structured result's `mode_tightened` field reports this change.

## inv validate

Validates one complete static Pigsty inventory: YAML structure, Ansible conventions, and
Pigsty semantics (such as `admin_ip`, the `infra` group, and IPv4 host keys) are checked
layer by layer; diagnostics never echo sensitive values.

```bash
pig inventory validate                   # validate the active inventory
pig inv v ./pigsty.yml                   # validate a specific file
pig inv v --strict                       # treat warnings as failures
pig inv v --ansible                      # cross-check with ansible-inventory
pig inv v --strict -o json               # structured output (CI-friendly)
```

| Option      | Description                                                |
|:------------|:-----------------------------------------------------------|
| `--strict`  | Treat validation warnings as failures                      |
| `--ansible` | Also parse with `ansible-inventory` |
| `--timeout` | Timeout for `--ansible` compatibility validation (default 10s) |
{.full-width}

Validation follows Pigsty's configuration rules, including requirements for `admin_ip`, the `infra` group, and host addresses. The file must contain one YAML document with no duplicate keys. These parsing rules also apply to `show` and `edit`.

`pg_users[].roles` accepts role names and objects containing `name`, optional
`state` (`grant`, `revoke`, `absent`), and nullable boolean `admin`, `set`, and `inherit`.
Editing, serialization, and CMDB YAML export preserve strings that look like YAML 1.1 scalars,
including `"1:30"`, `"010"`, dates, and timestamps supplied as strings. Necessary quotes prevent
downstream YAML consumers from changing their types. Existing unquoted or explicitly tagged timestamps
retain their source bytes and use the opaque `unknown` scalar kind; quote them when a string is intended.
`show` and explicit CMDB exports retain their raw-data contract.

## inv check

Readiness checks on top of static validation: by default only the inventory and local
controller conditions are checked; add probes with `--profile`.

```bash
pig inventory check                      # inventory + controller basics
pig inv ck -p ansible                    # add ansible environment checks
pig inv ck -p network                    # add network reachability probes
pig inv ck -p ssh --user admin           # add SSH login probes
pig inv ck -p ssh --sudo                 # SSH probes plus sudo -n validation
```

| Option      | Short | Description                                    |
|:------------|:------|:-----------------------------------------------|
| `--profile` | `-p`  | Additional probes: `ansible` / `network` / `ssh` |
| `--user`    |       | Explicit SSH user (defaults to the controller user) |
| `--port`    |       | TCP/SSH target port (default 22)               |
| `--sudo`    |       | With the ssh profile, also observe `sudo -n true` |
{.full-width}

`--user` and `--port` use explicit arguments or the defaults above. They do not infer effective SSH variables from the inventory. Combine profiles with commas, for example `--profile ansible,ssh`.

## inv diff

Compare the active inventory with another file. The result lists added, removed, and changed declarations by path and value type, without configuration values:

```bash
pig inv diff other.yml
pig inv diff other.yml -o json
```

## inv cmdb (experimental)

`pig inventory cmdb` exchanges the inventory with Pigsty's existing CMDB schema
(`files/cmdb.sql`, the `pigsty` / `pglog` schemas) over a native PostgreSQL driver.
Connection resolution: `-d/--database` (database name, URI, or libpq conninfo) →
`METADB_URL` environment variable → `service=meta`.

| Subcommand | Description                                                            |
|:-----------|:-----------------------------------------------------------------------|
| `check`    | Read-only: verify CMDB projections, optionally against the static inventory |
| `init`     | Apply the `cmdb.sql` baseline (confirmation required over existing schemas; `--plan` previews) |
| `load`     | **Replace all** CMDB declaration rows from the static inventory (one transaction; confirmation required; `--plan`/`--strict`) |
| `dump`     | Export the CMDB as a static inventory file (`--force` to overwrite a differing target) |
| `enable`   | Switch the `ansible.cfg` inventory to the CMDB (`inventory.sh`) |
| `disable`  | Switch back to the static `pigsty.yml` (both support `--plan`) |
{.full-width}

```bash
pig inventory cmdb check                 # verify CMDB state
pig inventory cmdb init --plan           # preview the init plan
pig inventory cmdb load -y               # load the static inventory into the CMDB
pig inventory cmdb dump -f pigsty.yml    # export the CMDB as a static inventory
pig inventory cmdb enable                # switch the ansible inventory source to the CMDB
pig inventory cmdb disable               # switch back to the static pigsty.yml
```

`init` runs `cmdb.sql` and replaces existing `pigsty` / `pglog` schemas. Back up data you need to retain before running it. `load` replaces all declaration rows in one transaction. Interactive mode asks you to confirm the target; use `--yes` for automation and structured output. Use `--plan` to inspect the operation first.

## Design records

- [Lossless Pigsty Inventory](https://pig.pgsty.com/design/lossless-inventory/)
- [Use the CMDB Pigsty Already Has](https://pig.pgsty.com/design/existing-cmdb-boundary/)
