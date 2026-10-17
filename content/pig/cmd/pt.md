---
title: "CMD: pt"
description: Run patronictl transparently with Pigsty service and config helpers
icon: fa-solid fa-terminal
weight: 660
---

`pig patroni` (alias `pig pt`) runs the installed `patronictl` command.
Use it to inspect cluster members, change cluster configuration, and switch the primary.
PIG selects the configuration file and provides local service, status, and log commands.
Other commands use patronictl's arguments, prompts, output, and exit codes.

Install Patroni and configure access to its cluster before running cluster operations.
Use `pig pt COMMAND --help` for native command help, or `pig pt -- --help` for patronictl root help.
Running `pig pt` alone shows PIG's help.

## Command Overview

**Forwarded commands** (native `patronictl`):

| Example                                        | Description                                       |
|:-----------------------------------------------|:--------------------------------------------------|
| `pig pt list [CLUSTER]`                        | List cluster members (native output, `--format json`) |
| `pig pt restart CLUSTER [MEMBER]`              | Restart PostgreSQL of a cluster / member (native prompt) |
| `pig pt reload CLUSTER`                        | Reload PostgreSQL configuration                   |
| `pig pt reinit CLUSTER MEMBER`                 | Reinitialize a member (resync from primary)       |
| `pig pt switchover CLUSTER [--candidate X]`    | Planned switchover                                |
| `pig pt failover CLUSTER --candidate MEMBER`   | Manual failover (cluster name comes first) |
| `pig pt pause CLUSTER` / `pig pt resume CLUSTER`                | Enter / leave maintenance mode                    |
| `pig pt show-config CLUSTER` / `pig pt edit-config CLUSTER`             | Show / edit cluster dynamic configuration         |
| `pig pt query -c 'select 1'`                   | Native query (`-c` here is query's own SQL option) |
{.full-width}

Put the cluster name before member names. These commands use the installed
[patronictl interface](https://patroni.readthedocs.io/en/latest/patronictl.html).

**Local commands** (implemented by PIG):

| Command      | Alias | Description                                          |
|:-------------|:------|:-----------------------------------------------------|
| `pt set`     |       | Update PostgreSQL params or scalar Patroni settings  |
| `pt status`  | `st`  | Combined status: systemd + processes + cluster       |
| `pt service` | `svc` | Manage the local patroni systemd service             |
| `pt start`   | `up`  | Shortcut for `pt svc start`                   |
| `pt stop`    | `dn`  | Shortcut for `pt svc stop`                    |
| `pt log`     | `l`   | View local Patroni logs (show / tail / grep)         |
{.full-width}

Use `pig pt restart CLUSTER` to restart PostgreSQL through Patroni.
Use `pig pt svc restart` to restart the local Patroni daemon.

## Inspect a cluster {#quick-start}

On a configured Patroni node, inspect the cluster members and dynamic configuration.
Replace `pg-meta` with your cluster name:

```bash
pig pt list pg-meta
pig pt show-config pg-meta
```

Use `pig pt status` to inspect the local service and `pig pt log -f` to follow its logs.

## PIG options

These PIG options **must precede the native command**; once the native command
token appears, all remaining tokens belong to patronictl:

| Option          | Short | Description                                                  |
|:----------------|:------|:-------------------------------------------------------------|
| `--config-file` | `-c`  | Explicit Patroni/patronictl configuration file               |
| `--dbsu`        |       | OS user used to execute patronictl (default `$PIG_DBSU` or `postgres`) |
| `--dcs-url`     | `-d`  | Override the Patroni DCS URL (`--dcs` is an alias)           |
| `--insecure`    | `-k`  | Allow TLS without certificate verification                   |
{.full-width}

```bash
pig pt -c /path/patroni.yml list pg-meta   # PIG -c: config file
pig pt query -c 'select 1'                 # native query -c: SQL
```

Use the full `--dbsu` option; `pig pt` has no `-U` shorthand.

## Config-File Resolution

PIG selects the patronictl configuration file in this order:

1. explicit `-c/--config-file` given before the command;
2. non-empty `PATRONICTL_CONFIG_FILE` environment variable;
3. `/etc/patroni/patroni.yml` (exists and readable as the DBSU);
4. `/infra/conf/patronictl.yml` (exists and readable as the DBSU);
5. otherwise `/etc/patroni/patroni.yml`; patronictl reports a missing file.

PIG uses an explicit path or environment path even if the file is missing or unreadable.
It resolves relative paths before switching operating-system user. It checks the two conventional
paths as the selected user. Service commands and native help do not need a configuration file.

## Command dispatch and output

PIG handles the local commands listed above. It forwards any other command and every argument
after that command to patronictl. Put PIG options before the native command:

```bash
pig pt -c /etc/patroni/patroni.yml list pg-meta
pig pt list pg-meta --format json
```

Native commands keep interactive prompts, editor sessions, streamed output, and exit codes.
Use the native `--format json` option for JSON. PIG's `-o json` and `-o yaml` cannot format
these commands and are rejected before the native command token.

Use `pig pt -- COMMAND ...` to forward a command whose name matches a local command.
The local `set`, `status`, and `service` commands support PIG structured output.
Log output has the formats described below.

## pt set

`pt set` updates the cluster selected by the configuration file:

```bash
pig pt set KEY=VALUE [KEY=VALUE ...] [--yes] [--plan]
```

**Key classification**:

- These scalar Patroni dynamic keys translate to native `--set`:
  `loop_wait`, `ttl`, `retry_timeout`, `primary_race_backoff`,
  `maximum_lag_on_failover`, `maximum_lag_on_syncnode`, `max_timelines_history`,
  `primary_start_timeout`, `primary_stop_timeout`, `synchronous_mode`,
  `synchronous_mode_strict`, `synchronous_node_count`, `failsafe_mode`,
  `check_timeline`, `member_slots_ttl` (`pause` is deliberately excluded —
  use native `pause`/`resume`);
- every other key is treated as a PostgreSQL parameter and translates to native `--pg`
  (including dotted custom GUCs like `timescaledb.telemetry_level`);
- structural keys starting with `postgresql.`, `standby_cluster.`, `slots.`, or
  `ignore_slots.` are **rejected**, pointing to native `pig pt edit-config --set`.

All pairs merge into **one** native `edit-config` call in input order — one configuration diff and one confirmation:

```bash
pig pt set ttl=60 max_connections=200 synchronous_mode=on
# equivalent: patronictl -c <selected> edit-config --set ttl=60 --pg max_connections=200 --set synchronous_mode=on
```

- `--yes/-y` appends native `--force` to skip confirmation; otherwise patronictl shows the
  diff and owns confirmation;
- `--plan` mutates nothing and renders the selected config plus the translated native
  command (including the values you supplied);
- in structured output mode, execution requires an explicit `--yes`;
- values parse as YAML: both `KEY=null` and `KEY=` remove a key, passed unchanged;
- after changing restart-required PostgreSQL parameters, PIG suggests the follow-up:
  first `pig pt list`, then `pig pt restart CLUSTER --pending` (explicit cluster required).

## Service Management

`pt service` (alias `pt svc`) manages the local `patroni` systemd unit:

| Command              | Alias       | Description                 |
|:---------------------|:------------|:----------------------------|
| `pt service start`   | `pt svc up` | Start the Patroni service   |
| `pt service stop`    | `pt svc dn` | Stop the Patroni service    |
| `pt service restart` | `pt svc rs` | Restart the Patroni service |
| `pt service reload`  | `pt svc rl` | Reload the Patroni service  |
| `pt service status`  | `pt svc st` | Show service status         |
{.full-width}

`pt start` / `pt stop` (aliases `up` / `dn`) call the same service actions. Stopping the Patroni service may also stop PostgreSQL on this node
(depending on Patroni configuration).

### pt status

Shows combined status: systemd service state, Patroni process info, and cluster members
from patronictl.

```bash
pig pt status
pig pt st -o json                # structured output
```

### pt log

View local Patroni logs. The log directory comes from `log.dir` in the selected config,
falling back to `/pg/log/patroni`, or set it explicitly with `--log-dir`. Only `pt log` and
`pt log show` support `-o json` (JSONL snapshot); `-f`, `tail`, and `grep` use text output. Log snapshots do not support YAML or `json-pretty`.

```bash
pig pt log                     # show last 50 lines
pig pt log -f                  # follow log output
pig pt log -n 100              # show last 100 lines
pig pt log show -o json        # JSONL snapshot
pig pt log tail -n 100         # follow logs
pig pt log grep ERROR          # search logs
```

| Subcommand | Aliases        | Description               |
|:-----------|:---------------|:--------------------------|
| `show`     | `cat, c, s`    | Print recent Patroni logs |
| `tail`     | `t, f, follow` | Follow Patroni logs       |
| `grep`     | `g, search`    | Search Patroni logs       |
{.full-width}

| Option      | Short | Default | Description           |
|:------------|:------|:--------|:----------------------|
| `--follow`  | `-f`  | false   | Follow log output     |
| `--lines`   | `-n`  | 50      | Lines for show/tail; grep searches all lines unless `-n` is set |
| `--log-dir` |       | auto    | Log directory         |
{.full-width}

## Migrating from v1.5.x

When upgrading from v1.5.x, update older scripts to use patronictl syntax:

| v1.5.x usage                           | v1.6.0 usage                                                |
|:---------------------------------------|:------------------------------------------------------------|
| `pig pt failover <candidate>`          | `pig pt failover CLUSTER --candidate MEMBER` (the positional is now the **cluster**) |
| `pig pt restart [member]` (auto scope) | `pig pt restart CLUSTER [MEMBER]` (explicit cluster)        |
| `pig pt list -o json`                  | `pig pt list --format json` (native JSON, different schema) |
| `pig pt config show`                   | `pig pt show-config`                                        |
| `pig pt config edit`                   | `pig pt edit-config`                                        |
| `pig pt config set K=V` / `pg K=V`     | `pig pt set K=V`                                            |
| `pig pt restart -y` (PIG confirmation)         | native patronictl confirmation; `pt set -y` still works     |
| `pig pt list -W` / `-w 5`              | native `pig pt list --watch` (patronictl semantics)         |
| aliases `ls/rs/rl/ri/so/fo/p/r/c`      | removed — use full native command names                     |
| `--dbsu -U`                            | `--dbsu` (`-U` shorthand removed)                           |
{.full-width}

## Requirements and privileges

Service commands require Linux and systemd. Log commands need readable Patroni log files.
PIG runs patronictl as the operating-system user selected by `--dbsu`, then `$PIG_DBSU`,
then `postgres`. It executes directly as that user, uses `su` from root, or uses `sudo -H -u`
from another account. Service management uses root privileges.

## Design record

[Let Patronictl Speak for Itself](https://pig.pgsty.com/design/patronictl-passthrough/) records why PIG stopped
mirroring the upstream command tree and which local responsibilities remain.
