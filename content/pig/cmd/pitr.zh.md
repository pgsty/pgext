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
  2. Ensure PostgreSQL is stopped (with retry and fallback)
  3. Execute pgbackrest restore
  4. Start PostgreSQL unless --no-restart is used
  5. Provide post-restore guidance

Recovery Targets (at least one required):
  --default, -d      Recover to end of WAL stream (latest)
  --immediate, -I    Recover to backup consistency point
  --time, -t         Recover to specific timestamp
  --name, -n         Recover to named restore point
  --lsn, -l          Recover to specific LSN
  --xid, -x          Recover to specific transaction ID

Backup and Target Options:
  --set, -b          Select backup set to start recovery from
  --target-action    Action when target is reached: pause, promote, shutdown
  --target-timeline  Recover along timeline: latest, current, N, or 0xN

Time Format:
  - Full: "2025-01-01 12:00:00+08"
  - Date only: "2025-01-01" (defaults to 00:00:00)
  - Time only: "12:00:00" (defaults to today)

Examples:
  pig pitr -d                      # Recover to latest (most common)
  pig pitr -t "2025-01-01 12:00:00+08"  # Recover to specific time
  pig pitr -I                      # Recover to backup consistency point
  pig pitr -d --plan               # Show execution plan without running
  pig pitr -d -y                   # Skip confirmation (for automation)
  pig pitr -d --no-restart         # Leave PostgreSQL stopped after restore
  pig pitr -d -D /tmp/pg-restore -S -N  # Side restore
```

## 命令概览

`pig pitr` 的默认目标是恢复 Pigsty 管理的主数据目录。典型流程如下：

1. 验证恢复目标参数，必须指定 `-d/-I/-t/-n/-l/-x` 之一
2. 解析 pgBackRest 配置与目标数据目录
3. 对默认数据目录恢复时，若 Patroni 正在运行，则停止 Patroni
4. 确保 PostgreSQL 已停止
5. 调用 `pgbackrest restore`
6. 除非指定 `--no-restart`，否则启动 PostgreSQL
7. 输出恢复后验证与 Patroni 恢复指引

**与 `pig pb restore` 的区别：**

| 特性 | `pig pitr` | `pig pb restore` |
|:----|:----------|:-----------------|
| 停止 Patroni | 默认数据目录恢复时自动停止 | 手动处理 |
| 停止 PostgreSQL | 自动检查并停止 | 必须预先停止 |
| 启动 PostgreSQL | 默认自动启动，可用 `--no-restart` 禁用 | 手动处理 |
| Patroni 恢复 | 不自动恢复，验证后人工处理 | 不处理 |
| 适用场景 | 生产恢复编排 | 底层 restore 或脚本集成 |
{.full-width}


## 快速入门

```bash
# 最常用：恢复到最新数据
pig pitr -d

# 恢复到指定时间点
pig pitr -t "2025-01-01 12:00:00+08"

# 恢复到备份一致性点（最快）
pig pitr -I

# 查看执行计划（不实际执行）
pig pitr -d --plan

# 跳过确认（用于自动化）
pig pitr -d -y

# 从特定备份集恢复
pig pitr -d -b 20251225-120000F

# 恢复默认数据目录，但恢复后不启动 PostgreSQL
pig pitr -d --no-restart

# side restore：恢复到自定义目录，不触碰 Patroni 与 /pg/data
pig pitr -d -D /tmp/pg-restore --skip-patroni --no-restart

# 额外 pgBackRest restore 参数写在 -- 之后
pig pitr -d -- --delta
```


## 参数说明

### 恢复目标（必选其一）

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--default` | `-d` | 恢复到 WAL 流末尾（最新数据） |
| `--immediate` | `-I` | 恢复到备份一致性点 |
| `--time` | `-t` | 恢复到指定时间戳 |
| `--name` | `-n` | 恢复到命名还原点 |
| `--lsn` | `-l` | 恢复到指定 LSN |
| `--xid` | `-x` | 恢复到指定事务 ID |
{.full-width}

### 备份与目标选项

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--set` | `-b` | 从特定备份集开始恢复 |
| `--target-action` | | 到达恢复目标后的动作：pause/promote/shutdown |
| `--target-timeline` | `-T` | 恢复时间线：latest/current/N/0xN |
| `--exclusive` | `-X` | 排他模式：在目标前停止 |
| `--promote` | `-P` | 到达手工恢复目标后自动提升 |
{.full-width}

### 流程控制

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--skip-patroni` | `-S` | 跳过 Patroni 停止操作，适用于独立 PostgreSQL 或自定义目录 side restore |
| `--no-restart` | `-N` | restore 后不启动 PostgreSQL |
| `--plan` | | 仅显示执行计划，不执行 |
| `--yes` | `-y` | 跳过破坏性操作确认 |
| `--timeout` | | PostgreSQL 启动/恢复等待超时，默认 120 秒 |
| `--force-stop` | | fast stop 失败时允许 immediate shutdown 与 kill fallback |
{.full-width}

### 配置参数

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--stanza` | `-s` | pgBackRest stanza 名称（自动检测） |
| `--config` | `-c` | pgBackRest 配置文件路径 |
| `--repo` | `-r` | 仓库编号（多仓库场景） |
| `--dbsu` | `-U` | 数据库超级用户（默认：`postgres`） |
| `--data` | `-D` | 目标数据目录 |
{.full-width}


## 时间格式

`--time` 参数支持多种时间格式，自动补全时区：

| 格式 | 示例 | 说明 |
|:----|:----|:----|
| 完整格式 | `2025-01-01 12:00:00+08` | 包含时区的完整时间戳 |
| 仅日期 | `2025-01-01` | 自动补全为当天 00:00:00（当前时区） |
| 仅时间 | `12:00:00` | 自动补全为今天（当前时区） |
{.full-width}


## 执行流程

### 第一阶段：预检查

- 验证恢复目标参数（必须且只能指定一个）
- 检查数据目录是否存在且已初始化
- 检测 Patroni 服务状态
- 检测 PostgreSQL 运行状态

### 第二阶段：处理 Patroni

默认数据目录恢复时，如果 Patroni 正在运行，命令会停止 Patroni，让目标 PGDATA 在 restore 期间保持离线。恢复完成后 Patroni 会保持停止。

如果 Patroni 正在运行且恢复默认数据目录，`--skip-patroni` 会被拒绝，因为 Patroni 可能在 restore 期间重新拉起 PostgreSQL。自定义 `-D` side restore 不触碰 `/pg/data`，可以配合 `--skip-patroni --no-restart` 使用。

### 第三阶段：确保 PostgreSQL 停止

命令会先尝试 fast stop。如果 PostgreSQL 无法停止，默认不会直接使用更激进手段；需要显式指定 `--force-stop`，才允许 immediate shutdown 与 kill fallback。

### 第四阶段：执行恢复

调用 pgBackRest 执行 restore，并把恢复目标、备份集、时间线、target action 等参数映射到 `pgbackrest restore`。原生 pgBackRest 参数可以写在 `--` 之后：

```bash
pig pitr -d -- --delta
```

### 第五阶段：启动或保持停止

除非指定 `--no-restart`，命令会在 restore 后启动 PostgreSQL，并等待恢复完成。以下情况必须或通常应使用 `--no-restart`：

- 自定义 `-D` side restore，因为恢复出的配置仍保留原端口，需要手动指定空闲端口启动
- `--target-action=shutdown`，因为 PostgreSQL 到达恢复目标后会退出
- 需要人工检查恢复目录后再决定是否启动

启动 side restore 时请手动指定空闲端口，例如：

```bash
pg_ctl -D /tmp/pg-restore -o "-p 15432" start
```


## 使用示例

### 场景一：误删数据恢复

```bash
# 1. 查看可用的备份
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
# 服务器故障后恢复到最新数据
pig pitr -d
```

### 场景三：快速恢复到备份点

```bash
# 恢复到备份一致性点（不需要回放 WAL）
pig pitr -I
```

### 场景四：自动化脚本

```bash
# 跳过所有确认，适合自动化
pig pitr -d -y
```

### 场景五：自定义目录 side restore

```bash
pig pitr -d -D /tmp/pg-restore --skip-patroni --no-restart

# 使用空闲端口手动启动 side restore
pg_ctl -D /tmp/pg-restore -o "-p 15432" start
```

### 场景六：仅恢复不启动

```bash
# 恢复后手动检查，再决定是否启动
pig pitr -d --no-restart

# 检查恢复的数据目录
ls -la /pg/data/

# 手动启动
pig pg start
```


## 执行计划示例

执行 `pig pitr -d --plan` 会显示类似以下的执行计划：

```
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

[Plan mode] No changes made.
```


## 恢复后操作

成功恢复后，命令会显示详细的后续操作指引：

```
══════════════════════════════════════════════════════════════════
 PITR Complete
══════════════════════════════════════════════════════════════════

[1] Verify recovered data:
   pig pg psql

[2] If satisfied, promote to primary:
   pig pg promote

[3] To resume Patroni cluster management:
   WARNING: Ensure data is correct before starting Patroni!
   systemctl start patroni

   Or if you want this node to be the leader:
   1. Promote PostgreSQL first: pig pg promote
   2. Then start Patroni: systemctl start patroni

[4] Re-create pgBackRest stanza if needed:
   pig pb create

══════════════════════════════════════════════════════════════════
```


## 安全机制

### 确认倒计时

除非使用 `--yes` 参数，命令执行前会显示 5 秒倒计时：

```
WARNING: This will overwrite the current database!
Press Ctrl+C to cancel, or wait for countdown...
Starting PITR in 5 seconds...
```

### 停止策略

默认路径保持保守：先使用 fast shutdown。只有显式指定 `--force-stop` 时，才允许 immediate shutdown 与 kill fallback。

### 恢复验证

如果恢复后启动 PostgreSQL，命令会等待启动/恢复完成；失败时会提示检查日志。


## 设计说明

**与其他命令的关系：**

- `pig pitr` 调用 pgBackRest restore，并在本地处理 Patroni/PostgreSQL 停止与可选启动。
- `pig pitr` 不是集群恢复总控，不负责 Patroni failover、rejoin、VIP 或应用流量切换。
- 需要底层 restore 语义或脚本细粒度控制时，使用 `pig pb restore`。
- 需要手工切换 Patroni 集群时，使用 `pig pt switchover` 或 `pig pt failover`。

**错误处理：**

- 每个阶段都有详细的错误信息
- 失败时提示相关日志位置
- 支持中断后手动继续

**权限执行：**

- 如果当前用户已是 DBSU：直接执行命令
- 如果当前用户是 root：使用 `su - postgres -c "..."` 执行
- 其他用户：使用 `sudo -inu postgres -- ...` 执行

**平台支持：**

此命令专为 Linux 系统设计，依赖 Pigsty 的默认目录结构。
