---
title: "pig pitr"
description: "使用 pig pitr 命令执行编排式时间点恢复（PITR）"
icon: Clock
weight: 670
---

`pig pitr` 命令用于通过 pgBackRest 执行时间点恢复，并以保守方式处理本地 PostgreSQL 与 Patroni 生命周期。与底层的 `pig pb restore` 不同，`pig pitr` 会先做恢复前检查，必要时停止 Patroni 与 PostgreSQL，执行 restore，然后按参数决定是否启动 PostgreSQL。

请注意：对于托管的默认数据目录，`pig pitr` 恢复后会让 Patroni 保持停止。请先验证恢复结果，再由人工恢复 Patroni 管理；该命令不会自动重入 Patroni 集群、执行故障切换，或验证集群成员状态。

```bash
pig pitr - Perform PITR with pgBackRest restore and conservative PostgreSQL stop/start handling.

For the managed default data directory, this command may:
  1. Stop Patroni only to keep the target PGDATA offline during restore
  2. Ensure PostgreSQL is stopped (fast stop with retry; destructive fallback only with --force-stop)
  3. Execute pgbackrest restore
  4. Start PostgreSQL unless --no-restart is used
  5. Leave Patroni stopped; provide post-restore guidance

Recovery Targets (at least one required):
  --default, -d      Recover to end of WAL stream (latest)
  --immediate, -I    Recover to backup consistency point
  --time, -t         Recover to specific timestamp
  --name             Recover to named restore point
  --lsn              Recover to specific LSN
  --xid              Recover to specific transaction ID

Backup and Target Options:
  --set, -b          Select backup set to start recovery from
  --target-action    Action when target is reached: pause, promote, shutdown
  --target-timeline  Recover along timeline: latest, current, N, or 0xN

Use --no-restart with --target-action=shutdown because PostgreSQL exits
after reaching the recovery target.

Custom -D side restores require --no-restart. The custom directory must
already exist, be owned by the configured DBSU (postgres by default), and
be writable by that user. Pig does not create this directory automatically.

Additional pgBackRest arguments:
  Put raw pgBackRest restore arguments after -- so Cobra stops parsing them.
  Example: pig pitr -d -- --delta

Time Format:
  - Full: "2025-01-01 12:00:00+08"
  - Date only: "2025-01-01" (defaults to 00:00:00)
  - Time only: "12:00:00" (defaults to today)

Examples:
  pig pitr -d                                      # Recover to latest
  pig pitr -t "2025-01-01 12:00:00+08"            # Recover to time
  pig pitr -d --plan                               # Preview plan
  pig pitr -d -y                                   # Skip y/yes confirmation
  pig pitr -d --no-restart                         # Leave PostgreSQL stopped
  pig pitr -d -D /tmp/pg-restore --no-restart      # Side restore
```

## 命令概览

`pig pitr` 的默认目标是恢复 Pigsty 管理的主数据目录。典型流程如下：

1. 验证恢复目标参数，必须指定 `-d/-I/-t/--name/--lsn/--xid` 之一
2. 解析 pgBackRest 配置与目标数据目录
3. 对默认数据目录恢复时，若 Patroni 正在运行，则停止 Patroni
4. 确保 PostgreSQL 已停止
5. 调用 `pgbackrest restore`
6. 除非指定 `--no-restart`，否则启动 PostgreSQL
7. 输出恢复后验证与 Patroni 恢复指引

**与 `pig pb restore` 的区别：**

| 特性            | `pig pitr`                  | `pig pb restore` |
|:--------------|:----------------------------|:-----------------|
| 停止 Patroni    | 默认数据目录恢复时自动停止               | 手动处理             |
| 停止 PostgreSQL | 自动检查并停止                     | 必须预先停止           |
| 启动 PostgreSQL | 默认自动启动，可用 `--no-restart` 禁用 | 手动处理             |
| Patroni 恢复    | 不自动恢复，验证后人工处理               | 不处理              |
| 适用场景          | 生产恢复编排                      | 底层 restore 或脚本集成 |
{.full-width}


## 快速入门

```bash
# 最常用：恢复到 WAL 流末尾
pig pitr -d

# 恢复到指定时间点
pig pitr -t "2025-01-01 12:00:00+08"

# 恢复到备份一致性点
pig pitr -I

# 查看执行计划，不实际执行
pig pitr -d --plan

# 跳过确认，适合自动化脚本
pig pitr -d -y

# 从特定备份集恢复
pig pitr -d -b 20251225-120000F

# 恢复默认数据目录，但恢复后不启动 PostgreSQL
pig pitr -d --no-restart

# side restore：恢复到自定义目录，不触碰 Patroni 与 /pg/data
install -d -m 700 -o postgres -g postgres /tmp/pg-restore
pig pitr -d -D /tmp/pg-restore --no-restart

# 额外 pgBackRest restore 参数写在 -- 之后
pig pitr -d -- --delta
```


## 参数说明

### 恢复目标（必选其一）

| 参数            | 简写   | 说明                |
|:--------------|:-----|:------------------|
| `--default`   | `-d` | 恢复到 WAL 流末尾（最新数据） |
| `--immediate` | `-I` | 恢复到备份一致性点         |
| `--time`      | `-t` | 恢复到指定时间戳          |
| `--name`      |      | 恢复到命名还原点          |
| `--lsn`       |      | 恢复到指定 LSN         |
| `--xid`       |      | 恢复到指定事务 ID        |
{.full-width}

### 备份与目标选项

| 参数                  | 简写   | 说明                                |
|:--------------------|:-----|:----------------------------------|
| `--set`             | `-b` | 从特定备份集开始恢复                        |
| `--target-action`   |      | 到达恢复目标后的动作：pause/promote/shutdown |
| `--target-timeline` | `-T` | 恢复时间线：latest/current/N/0xN        |
| `--exclusive`       | `-X` | 排他模式：在目标前停止                       |
{.full-width}

`--target-action=shutdown` 必须配合 `--no-restart`，因为 PostgreSQL 到达目标后会退出。`--target-action` 不能与 `--default` 同时使用，因为 `--default` 已表示恢复到 WAL 末尾。`--exclusive/-X` 必须配合明确的停止目标使用：`--time`、`--lsn` 或 `--xid`。

`--` 后的原生 pgBackRest restore 参数不能覆盖 Pig 已管理的恢复目标、生命周期、数据目录、仓库、配置和选择参数；这些语义应使用 Pig 的一等参数设置。该限制与 `pig pb restore` 的 passthrough blocklist 一致。

### 流程控制

| 参数             | 简写   | 说明                                                 |
|:---------------|:-----|:---------------------------------------------------|
| `--no-restart` |      | restore 后不启动 PostgreSQL                            |
| `--plan`       |      | 仅显示执行计划，不执行                                        |
| `--yes`        | `-y` | 跳过交互式 y/yes 确认                                     |
| `--timeout`    |      | PostgreSQL 启动/恢复等待超时，默认 120 秒                      |
| `--force-stop` |      | fast stop 失败时允许 immediate shutdown 与 kill fallback |
{.full-width}

### 配置参数

| 参数         | 简写   | 说明                     |
|:-----------|:-----|:-----------------------|
| `--stanza` | `-s` | pgBackRest stanza 名称   |
| `--config` | `-c` | pgBackRest 配置文件路径      |
| `--repo`   | `-r` | 仓库编号                   |
| `--dbsu`   | `-U` | 数据库超级用户（默认：`postgres`） |
| `--data`   | `-D` | 目标数据目录                 |
{.full-width}


## 时间格式

`--time` 参数支持多种时间格式，会按当前时区补全缺失部分：

| 格式      | 示例                       | 说明                    |
|:--------|:-------------------------|:----------------------|
| 完整格式    | `2025-01-01 12:00:00+08` | 包含时区的完整时间戳            |
| 无时区日期时间 | `2025-01-01 12:00:00`    | 自动补当前本地时区，`T` 分隔符也可接受 |
| 仅日期     | `2025-01-01`             | 自动补全为当天 00:00:00      |
| 仅时间     | `12:00:00`               | 自动补全为今天的该时间           |
{.full-width}

计划输出与可重放的 next-action 命令会把仅日期、仅时间目标规范化为带时区的确定性时间戳；该规则与 `pig pb restore --plan` 一致。


## 托管目录与 Side Restore

托管 PostgreSQL 数据目录来自有效 pgBackRest 配置中的 `pg1-path` 与命令参数，而不是硬编码 `/pg/data`。例如托管 PGDATA 是 `/var/lib/pgsql/18/data` 时，仍然按托管恢复处理。路径比较会在需要时以数据库超级用户解析软链接，因此指向托管 PGDATA 的软链接不会被误判为 side restore。

显式 `-D/--data` 且解析后不同于托管目录时，才是 side restore。side restore 必须使用 `--no-restart`，不会停止 Patroni，也不会管理默认 PostgreSQL 服务；恢复后请手工使用类似 `pg_ctl -D <dir> -o "-p 5433" start`、`pg_ctl -D <dir> status` 和 `pgbackrest --pg1-path=<dir> stanza-create` 的命令处理。side restore 目录必须已存在且归属 DBSU；与托管 PGDATA 不同，它不要求预先包含 `PG_VERSION` 初始化标记。Pig 不会自动创建该目录，因为命令需要在 destructive restore 前完成路径归类、owner 检查和安全提示。

```bash
install -d -m 700 -o postgres -g postgres /tmp/pg-restore
pig pitr -d -D /tmp/pg-restore --no-restart
```

对于非 `/pg/data` 的托管 PGDATA，恢复后的 runbook 命令会显式带上有效数据目录，例如：

```bash
pig pg start -D /var/lib/pgsql/18/data
pig pg psql -D /var/lib/pgsql/18/data
pig pg promote -D /var/lib/pgsql/18/data
```

`pig pg psql -D <dir>` 会读取该目录的 `postmaster.pid`，使用其中记录的端口和 socket 目录连接恢复后的实例；无法解析 postmaster 信息时不会静默回退到默认连接目标。


## 执行流程

### 第一阶段：预检查

- 验证恢复目标参数，缺失目标时只显示帮助并返回错误
- 解析有效 pgBackRest 配置、stanza、仓库与托管 `pg1-path`
- 检查托管数据目录存在且已初始化
- 对 side restore，检查自定义目录存在且归属 DBSU
- 验证所选 stanza 正常且存在备份；指定 `--set` 时验证对应备份集存在
- 检测 Patroni 服务状态与 PostgreSQL 运行状态

### 第二阶段：处理 Patroni

托管数据目录恢复时，如果 Patroni 正在运行，命令会停止 Patroni，让目标 PGDATA 在 restore 期间保持离线。恢复完成后 Patroni 会保持停止。自定义 `-D` side restore 不触碰托管数据目录，因此不会停止 Patroni。

### 第三阶段：确保 PostgreSQL 停止

命令会先等待 Patroni 停止后 PostgreSQL 自动退出，再重试 `pg_ctl stop -m fast`。如果 PostgreSQL 仍无法停止，默认不会使用更激进手段；只有显式指定 `--force-stop`，才允许 immediate shutdown 与最终的 kill fallback。

### 第四阶段：执行恢复

调用 pgBackRest 执行 restore，并把恢复目标、备份集、时间线、target action 等参数映射到 `pgbackrest restore`。原生 pgBackRest 参数可以写在 `--` 之后：

```bash
pig pitr -d -- --delta
```

### 第五阶段：启动或保持停止

除非指定 `--no-restart`，命令会在 restore 后启动 PostgreSQL，并等待恢复完成。对于 `--default` 与 `--target-action=promote`，命令会等待恢复实例上的 `pg_is_in_recovery()` 变为 false；恢复与后续 SQL 探测会绑定恢复数据目录 `postmaster.pid` 中的端口，并在存在时使用其中的 socket 目录。以下情况必须或通常应使用 `--no-restart`：

- 自定义 `-D` side restore，因为恢复出的配置仍保留原端口，需要手动指定空闲端口启动
- `--target-action=shutdown`，因为 PostgreSQL 到达恢复目标后会退出
- 需要人工检查恢复目录后再决定是否启动


## 使用示例

### 场景一：误删数据恢复

```bash
# 1. 查看可用备份
pig pb info

# 2. 预览恢复计划
pig pitr -t "2025-01-15 09:30:00+08" --plan

# 3. 执行恢复
pig pitr -t "2025-01-15 09:30:00+08"

# 4. 验证数据
pig pg psql
SELECT * FROM important_table;
```

### 场景二：恢复到最新状态

```bash
pig pitr -d
```

### 场景三：恢复到备份一致性点

```bash
pig pitr -I
```

### 场景四：恢复后保持停止

```bash
pig pitr -d --no-restart

# PostgreSQL 与 Patroni 都保持停止；手动检查目录或日志后再启动
pig pg start
```

### 场景五：自定义目录 side restore

```bash
install -d -m 700 -o postgres -g postgres /tmp/pg-restore
pig pitr -d -D /tmp/pg-restore --no-restart

# 使用空闲端口手动启动 side restore
pg_ctl -D /tmp/pg-restore -o "-p 5433" start
```


## 执行计划示例

执行 `pig pitr -d --plan` 会显示类似以下的计划：

```text
══════════════════════════════════════════════════════════════════
 PITR Execution Plan
══════════════════════════════════════════════════════════════════

Current State:
  Data Directory:  /pg/data
  Database User:   postgres
  Patroni Service: active
  PostgreSQL:      running (PID: 12345)

Recovery Target:
  Latest (end of WAL stream)

Execution Steps:
  [1] Stop Patroni service
  [2] Ensure PostgreSQL is stopped
  [3] Execute pgBackRest restore
  [4] Start PostgreSQL
  [5] Print post-restore guidance

══════════════════════════════════════════════════════════════════
```


## 恢复后操作

成功恢复后，请先验证数据，再决定如何恢复服务编排：

```bash
# 验证数据
pig pg psql

# 如果恢复到了手工目标，且需要提升为主库
pig pg promote

# 验证完成后再恢复 Patroni 管理
systemctl start patroni

# 如有需要，重新创建 pgBackRest stanza
pig pb create
```

托管数据目录恢复后，Patroni 保持停止是刻意设计：避免恢复出的旧状态在未经确认前重新进入 HA 编排。


## 安全机制

**恢复目标必填：** 不指定 `-d/-I/-t/--name/--lsn/--xid` 时，命令只显示帮助，不执行 restore。

**确认机制：** 文本模式下，破坏性恢复会在执行前要求交互式 `y`/`yes` 确认；自动化脚本可使用 `-y|--yes`。结构化输出模式不会交互提示，必须使用 `--yes` 执行或 `--plan` 预览。

**Patroni 边界：** 托管数据目录恢复时，命令会在需要时停止 Patroni，防止 Patroni 在 restore 期间重新拉起 PostgreSQL。恢复后不会自动重入 Patroni。

**Side restore 边界：** 自定义 `-D` side restore 必须使用 `--no-restart`，因为恢复出来的 PostgreSQL 配置仍使用原端口；side restore 不管理 Patroni 或默认 PostgreSQL 服务。

**失败边界：** 如果 restore 在 Patroni 已停止后失败，Patroni 会保持停止，目标数据目录可能已经部分恢复。修复底层问题后应重新执行 PITR，或先验证恢复状态；不要在未确认前启动 Patroni。如果 restore 已执行但 PostgreSQL 启动失败，同样应先检查 PostgreSQL/pgBackRest 日志并验证数据目录，再决定是否恢复 Patroni 管理。

**结构化输出：** 结构化执行需要 `--yes`；`--plan` 是预览路径。成功执行后的结构化 PITR 结果会把恢复后操作放在 Result envelope 的 `next_actions`，而不是 `data` 内部。`data` 包含 `requested_data_dir`、`effective_data_dir`、`managed_data_dir` 与 `side_restore`，便于自动化区分用户输入和实际恢复目标。


## 设计说明

- `pig pitr` 调用 pgBackRest restore，并在本地处理 Patroni/PostgreSQL 停止与可选启动。
- `pig pitr` 不是集群恢复总控，不负责 Patroni failover、rejoin、VIP 或应用流量切换。
- 需要底层 restore 语义或脚本细粒度控制时，使用 [`pig pb restore`](/zh/pig/cmd/pb/#pb-restore/)。
- 需要手工切换 Patroni 集群时，使用 [`pig pt switchover`](/zh/pig/cmd/pt/#pt-switchover/) 或 [`pig pt failover`](/zh/pig/cmd/pt/#pt-failover/)。

**权限执行：**

- 如果当前用户已是 DBSU：直接执行命令
- 如果当前用户是 root：使用 `su - postgres -c` 执行
- 其他用户：使用 `sudo -inu postgres --` 执行

**平台支持：**

此命令专为 Linux 系统设计，依赖 pgBackRest、systemd（托管服务场景）以及 DBSU 可访问的数据目录与日志路径。
