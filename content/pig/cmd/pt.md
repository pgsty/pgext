---
title: "CMD: pt"
description: Run patronictl transparently with Pigsty service and config helpers
icon: SquareTerminal
weight: 660
---

Since v1.6.0, the `pig patroni` command (alias `pig pt`) is a **transparent launcher** for the
installed `patronictl` binary: pig only owns config-file selection and a few local helpers,
while every other command and all of its arguments are **forwarded unchanged** to `patronictl` —
with native flags, prompts, output, and exit codes. New patronictl features work without
waiting for a pig release.

```bash
pig pt transparently launches the installed patronictl binary.

Pig only owns config-file selection and these local helpers:
  pig pt set KEY=VALUE...       update PostgreSQL or scalar Patroni settings
  pig pt start|stop             manage the local patroni service
  pig pt service <action>       start, stop, restart, reload, or show service
  pig pt status                 show local and cluster status
  pig pt log ...                view local Patroni logs

Every other command and all arguments after it are passed unchanged to
patronictl. Use "pig pt -- COMMAND ..." to bypass a local-name collision.
```

The first non-option command token decides the dispatch:

- `set`, `start`/`up`, `stop`/`dn`, `service`/`svc`, `status`/`st`, and `log`/`l`
  select pig's local implementation;
- **every other token** (`list`, `restart`, `reload`, `reinit`, `switchover`, `failover`,
  `pause`, `resume`, `show-config`, `edit-config`, `query`, `history`, `topology`, `dsn`,
  `version`, …) and everything after it is forwarded verbatim to `patronictl`;
- `pig pt -- <command> ...` bypasses local-name collisions (e.g. `pig pt -- set`
  hands `set` to patronictl);
- `pig pt` with no command prints pig's help and does not run patronictl;
  use `pig pt <command> --help` for native subcommand help and `pig pt -- --help`
  for patronictl root help.


## Command Overview

**Forwarded commands** (native `patronictl`):

| Example                                        | Description                                       |
|:-----------------------------------------------|:--------------------------------------------------|
| `pig pt list [CLUSTER]`                        | List cluster members (native output, `--format json`) |
| `pig pt restart CLUSTER [MEMBER]`              | Restart PostgreSQL of a cluster / member (native prompt) |
| `pig pt reload CLUSTER`                        | Reload PostgreSQL configuration                   |
| `pig pt reinit CLUSTER MEMBER`                 | Reinitialize a member (resync from primary)       |
| `pig pt switchover CLUSTER [--candidate X]`    | Planned switchover                                |
| `pig pt failover CLUSTER --candidate MEMBER`   | Manual failover (**positional arg is the cluster**) |
| `pig pt pause / resume CLUSTER`                | Enter / leave maintenance mode                    |
| `pig pt show-config / edit-config`             | Show / edit cluster dynamic configuration         |
| `pig pt query -c 'select 1'`                   | Native query (`-c` here is query's own SQL option) |
{.full-width}

Forwarded positionals follow patronictl's native **cluster-first** semantics; confirmation
prompts, output formats, and exit codes (including Click usage-error exit code `2`) are all
owned by patronictl.

**Local commands** (implemented by pig):

| Command      | Alias | Description                                          |
|:-------------|:------|:-----------------------------------------------------|
| `pt set`     |       | Update PostgreSQL params or scalar Patroni settings  |
| `pt status`  | `st`  | Combined status: systemd + processes + cluster       |
| `pt service` | `svc` | Manage the local patroni systemd service             |
| `pt start`   | `up`  | Hidden shortcut for `pt svc start`                   |
| `pt stop`    | `dn`  | Hidden shortcut for `pt svc stop`                    |
| `pt log`     | `l`   | View local Patroni logs (show / tail / grep)         |
{.full-width}

Note: top-level `pt restart` is **not** a daemon shortcut — it forwards to
`patronictl restart` (restarting PostgreSQL). Use `pt svc restart` to restart
the patroni daemon itself.


## Quick Start

```bash
# Cluster operations (native passthrough, cluster-first)
pig pt list pg-meta                         # list cluster members
pig pt list pg-meta --format json           # native JSON output
pig pt restart pg-test --pending            # apply pending restarts
pig pt restart pg-test pg-test-1            # restart a specific member
pig pt switchover pg-test --candidate pg-2  # planned switchover
pig pt failover pg-test --candidate pg-2    # manual failover
pig pt pause pg-test                        # enter maintenance mode
pig pt show-config pg-test                  # show dynamic configuration

# Configuration sugar (local, one edit-config call)
pig pt set ttl=60                           # scalar Patroni key → --set
pig pt set max_connections=200              # PostgreSQL parameter → --pg
pig pt set ttl=60 max_connections=200 -y    # mixed change, skip confirmation
pig pt set shared_buffers=4GB --plan        # preview the translated native command

# Service and logs (local)
pig pt status                               # combined status
pig pt svc restart                          # restart the patroni daemon
pig pt log -f                               # follow logs
pig pt log grep ERROR                       # search logs
```


## Pig/PT Options

These wrapper-level options **must precede the native command**; once the native command
token appears, all remaining tokens belong to patronictl:

| Option          | Short | Description                                                  |
|:----------------|:------|:-------------------------------------------------------------|
| `--config-file` | `-c`  | Explicit Patroni/patronictl configuration file               |
| `--dbsu`        |       | OS user used to execute patronictl (default `$PIG_DBSU` or `postgres`) |
| `--dcs-url`     | `-d`  | Override the Patroni DCS URL (`--dcs` is an alias)           |
| `--insecure`    | `-k`  | Allow TLS without certificate verification                   |
{.full-width}

```bash
pig pt -c /path/patroni.yml list pg-meta   # wrapper-level -c: config file
pig pt query -c 'select 1'                 # native query's -c: SQL, not stolen by pig
```

> Since v1.6.0, `--dbsu` no longer has the `-U` shorthand (the `pg` / `pb` commands keep it).


## Config-File Resolution

Every config-aware patronictl invocation receives exactly one pig-selected root `-c <path>`,
resolved in this order:

1. explicit `-c/--config-file` given before the command;
2. non-empty `PATRONICTL_CONFIG_FILE` environment variable;
3. `/etc/patroni/patroni.yml` (exists and readable as the DBSU);
4. `/infra/conf/patronictl.yml` (exists and readable as the DBSU);
5. fallback to `/etc/patroni/patroni.yml`, so patronictl's error names the conventional path.

Explicit and environment paths are authoritative: pig does not silently substitute another
candidate when they are missing or unreadable; relative paths are made absolute before
switching OS user. Conventional candidates are probed for readability as the **resolved DBSU**
(not by mode bits). Resolution is lazy: pure systemd operations like `pt svc start` never
probe a config, and native `-h/--help` takes a config-independent fast path that works on
machines without Patroni configured.


## Transparent Execution and Output Modes

Forwarded execution directly inherits stdin / stdout / stderr and the terminal: native
prompts, the `edit-config` editor, `--watch` streaming, and exit codes are all preserved.
Pig performs no output capture, no duplicate error rendering, no retries, and no cluster/DCS
reads before or after execution. The logical invocation is:

```text
patronictl -c <selected-config> [--dcs-url URL] [--insecure] <native argv...>
```

**Structured output**: the forwarding path supports pig's `text` mode only. A pig
`-o json` / `-o yaml` placed before the native command is **rejected** (with guidance to use
patronictl's native output options); anything after the native command is forwarded as-is
and validated by patronictl:

```bash
pig -o json pt list          # rejected: use native output options
pig pt -o json list          # rejected: use native output options
pig pt list --format json    # forwarded: native patronictl JSON
pig pt list -o json          # forwarded unchanged; patronictl validates it
```

Local commands (`set` / `status` / `log` / `service`) keep pig's structured output behavior.


## pig pt set

`pt set` is the only local configuration sugar; it edits the cluster selected by the
resolved configuration:

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

All pairs merge into **one** native `edit-config` call in input order — one diff, one
confirmation, one DCS update:

```bash
pig pt set ttl=60 max_connections=200 synchronous_mode=on
# equivalent: patronictl -c <selected> edit-config --set ttl=60 --pg max_connections=200 --set synchronous_mode=on
```

- `--yes/-y` appends native `--force` to skip confirmation; otherwise patronictl shows the
  diff and owns confirmation;
- `--plan` mutates nothing and renders the selected config plus the translated native
  command (**plans may contain sensitive values**);
- in structured output mode, execution requires an explicit `--yes` (hidden prompts are
  forbidden);
- values parse as YAML: both `KEY=null` and `KEY=` remove a key, passed unchanged;
- after changing restart-required PostgreSQL parameters, pig suggests the follow-up:
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

Top-level `pt start` / `pt stop` (aliases `up` / `dn`) are hidden shortcuts to the same
local implementation. Stopping the Patroni service may also stop PostgreSQL on this node
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
`pt log show` support `-o json` (JSONL snapshot); follow / tail / grep are terminal streams
without structured output.

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
| `--lines`   | `-n`  | 50      | Number of lines shown |
| `--log-dir` |       | auto    | Log directory         |
{.full-width}


## Migrating from v1.5.x

The v1.6.0 passthrough rewrite is a **breaking change** — review automation before upgrading:

| v1.5.x usage                           | v1.6.0 usage                                                |
|:---------------------------------------|:------------------------------------------------------------|
| `pig pt failover <candidate>`          | ⚠ `pig pt failover CLUSTER --candidate MEMBER` (the positional is now the **cluster**) |
| `pig pt restart [member]` (auto scope) | `pig pt restart CLUSTER [MEMBER]` (explicit cluster)        |
| `pig pt list -o json`                  | `pig pt list --format json` (native JSON, different schema) |
| `pig pt config show`                   | `pig pt show-config`                                        |
| `pig pt config edit`                   | `pig pt edit-config`                                        |
| `pig pt config set K=V` / `pg K=V`     | `pig pt set K=V`                                            |
| `pig pt restart -y` (pig gate)         | native patronictl confirmation; `pt set -y` still works     |
| `pig pt list -W` / `-w 5`              | native `pig pt list --watch` (patronictl semantics)         |
| aliases `ls/rs/rl/ri/so/fo/p/r/c`      | removed — use full native command names                     |
| `--dbsu -U`                            | `--dbsu` (`-U` shorthand removed)                           |
{.full-width}

Also: forwarded commands return native patronictl exit codes (usage errors exit `2`), and
the pig-side pause-guard preflight for switchover/failover is gone — maintenance-mode
semantics are fully owned by Patroni.


## Design Notes

**Single authority**: there is one cluster-control path — the installed `patronictl` plus
one selected config, acting through the Patroni REST API / DCS. Pig embeds no DCS client or
REST control engine, keeps no patronictl command inventory, and adds no confirmation,
preflight, retry, or output rewriting to forwarded commands.

**Privilege handling** (unchanged from v1.5.x):

- if the current user is already the DBSU: execute directly;
- if the current user is root: execute via `su - postgres -c "..."`;
- otherwise: execute via `sudo -inu postgres -- ...`.

**Platform**: designed for Linux; service management relies on `systemctl`, and the log
helpers require readable Patroni log files.
