---
title: "CMD: inventory"
description: Inspect, edit, validate, check, and exchange the Pigsty Inventory
icon: SquareTerminal
weight: 645
---

The `pig inventory` command group (alias `pig inv`), introduced in v1.6.0, provides
**lossless** inspection, editing, validation, and readiness checks for the Pigsty inventory
(`pigsty.yml`), plus an experimental `cmdb` subtree that exchanges the inventory with the
PostgreSQL CMDB.

The lossless engine preserves YAML comments, formatting, key order, anchors, and line
endings byte-for-byte; `edit` re-parses the whole document before writing and writes
atomically — **invalid YAML can never reach disk**.

```bash
Inspect, edit, validate, check, and exchange Pigsty Inventory

Usage:
  pig inventory [command]

Aliases:
  inventory, inv

Available Commands:
  check       Check Inventory, controller, and opt-in target readiness
  cmdb        Exchange Inventory with Pigsty's existing PostgreSQL CMDB (experimental)
  diff        Compare declarations without emitting Inventory values
  edit        Edit Inventory or one selected YAML fragment
  list        List ordered Inventory topology and value kinds
  show        Show verbatim Inventory YAML (may contain secrets)
  status      Inspect the active Inventory source without executing it
  validate    Validate one complete static Pigsty Inventory
```

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

The inventory path comes from pig's config resolution (the global `-i/--inventory` flag or
the Pigsty home directory), and every command supports `-o json|yaml` structured output.


## Quick Start

```bash
pig inventory status                     # which inventory source is active?
pig inventory list                       # browse the inventory topology
pig inventory edit pg-meta               # edit only the pg-meta cluster fragment
pig inventory validate                   # full validation
pig inventory check -p ssh --sudo        # add SSH/sudo readiness probes
pig inventory diff /path/other.yml       # compare declarations with another file
```


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

Opens the selected fragment in `$EDITOR`. After you save and exit, pig re-applies
indentation and newline style → **re-parses the whole document** (aborting and keeping the
temp file on failure) → checks that the on-disk file was not concurrently modified → writes
atomically.

```bash
pig inventory edit                       # edit the whole inventory
pig inv e pg-meta                        # edit one fragment
pig inv e vars --from vars.yml           # replace the fragment from a file (or stdin: -)
```

| Option   | Description                                                        |
|:---------|:-------------------------------------------------------------------|
| `--from` | Replace the selected fragment from a regular file or stdin (`-`), skipping the editor |
{.full-width}

> **Note**: after a successful edit, pig tightens the inventory file mode to `0600` (the
> file may contain database passwords and other secrets); the result reports this via the
> `mode_tightened` field. If other users or tools read the file directly, adjust permissions
> or ownership accordingly.


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
| `--ansible` | Also parse with a bounded `ansible-inventory` adapter      |
| `--timeout` | Timeout for `--ansible` compatibility validation (default 10s) |
{.full-width}

> `validate` is a **Pigsty-semantics validator**, not a generic Ansible linter; its rules
> stay aligned with Pigsty's own `bin/validate` (intentionally stricter in a few places).
> Also note that parsing outright rejects **duplicate keys** and **multi-document YAML** —
> such files cannot even be used with `show` / `edit`.


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


## inv cmdb (experimental)

> **Experimental**: interfaces and behavior may change.

`pig inventory cmdb` exchanges the inventory with Pigsty's existing CMDB schema
(`files/cmdb.sql`, the `pigsty` / `pglog` schemas) over a native PostgreSQL driver.
Connection resolution: `-d/--database` (database name, URI, or libpq conninfo) →
`METADB_URL` environment variable → `service=meta`.

| Subcommand | Description                                                            |
|:-----------|:-----------------------------------------------------------------------|
| `check`    | Read-only: verify CMDB projections, optionally against the static inventory |
| `init`     | Apply the `cmdb.sql` baseline (`--yes` required over existing schemas; `--plan` preview) |
| `load`     | **Replace all** CMDB declaration rows from the static inventory (single transaction; `--yes`; `--plan`/`--strict`) |
| `dump`     | Export the CMDB as a static inventory file (`--force` to overwrite a differing target) |
| `enable`   | Guarded switch of the `ansible.cfg` inventory to the CMDB (`inventory.sh`) |
| `disable`  | Switch back to the static `pigsty.yml` (both support `--plan`; atomic, rollback-safe writes) |
{.full-width}

```bash
pig inventory cmdb check                 # verify CMDB state
pig inventory cmdb init --plan           # preview the init plan
pig inventory cmdb load -y               # load the static inventory into the CMDB
pig inventory cmdb dump -f pigsty.yml    # export the CMDB as a static inventory
pig inventory cmdb enable                # switch the ansible inventory source to the CMDB
pig inventory cmdb disable               # switch back to the static pigsty.yml
```

> **Note**: `init` applies the `cmdb.sql` baseline directly and does **not** take a backup
> of an existing CMDB — back it up yourself before running against real data; `load`
> replaces all declaration rows. Destructive operations are gated by a target-fingerprint
> confirmation, and structured output mode requires an explicit `--yes`.
