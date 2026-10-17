---
title: "pig pitr"
description: "使用 pig pitr 命令执行编排式时间点恢复（PITR）"
icon: fa-solid fa-clock
weight: 670
---

使用 `pig pitr` 从 pgBackRest 备份恢复本地 PostgreSQL 实例。
托管恢复时，PIG 按需停止服务并恢复数据。
默认在恢复后启动 PostgreSQL，保持 Patroni 停止，便于你验证恢复结果。

恢复会替换目标目录中的数据。先用 `--plan` 检查目录、备份来源、恢复目标和服务操作。

## 恢复前准备

需要已安装的 pgBackRest、可读取的配置文件、可用备份，以及到达恢复目标所需的 WAL 归档。
默认配置文件是 `/etc/pgbackrest/pgbackrest.conf`。使用 `-c` 选择其他配置文件，使用 `-s` 选择 stanza。

在目标数据目录所在主机上执行命令。PIG 依次从 `--dbsu`、`$PIG_DBSU`、`postgres` 选择
运行 PostgreSQL 和 pgBackRest 的操作系统用户。服务操作需要 root 权限，可直接使用 root 或通过 sudo 获取。

| 操作 | `pig pitr` | `pig pb restore` |
|:---|:---|:---|
| 为托管目录恢复停止 Patroni | 正在运行时自动停止 | 操作者停止 |
| 停止 PostgreSQL | 自动检查并停止 | 要求已停止 |
| 恢复备份文件 | pgBackRest | pgBackRest |
| 启动 PostgreSQL | 自动启动，除非指定 `--no-restart` | 操作者启动 |
| 恢复 Patroni 管理 | 操作者验证后恢复 | 操作者处理 |
{.full-width}

需要自行安排服务停止与启动步骤时，使用 [`pig pb restore`](/zh/pig/cmd/pb/#pb-restore)。

## 快速入门

以下示例将已有托管数据目录恢复到指定时间。将时间戳替换为实际恢复目标：

```bash
# 查看可用备份
pig pb info

# 检查恢复计划
pig pitr -t "2025-01-15 09:30:00+08" --plan -- --delta

# 恢复数据并启动 PostgreSQL
pig pitr -t "2025-01-15 09:30:00+08" -- --delta

# 连接恢复后的实例；使用实际数据目录
pig pg psql -D /pg/data
```

`--delta` 让 pgBackRest 按文件校验和恢复已有目录。PIG 会在恢复前要求输入 `y` 或 `yes`。
使用 `--yes` 可跳过交互确认。

指定时间、还原点、LSN、事务 ID 或 immediate 目标时，PostgreSQL 通常在到达目标后暂停恢复。
先检查数据，再提升主库。步骤见[恢复后操作](#post-recovery-actions)。

## 参数说明

### 恢复目标

必须且只能选择一个目标。缺少目标时返回错误，文本模式还会显示帮助。

| 参数 | 简写 | 说明 |
|:---|:---|:---|
| `--default` | `-d` | 恢复到可用 WAL 的末尾 |
| `--immediate` | `-I` | 恢复到备份一致性点 |
| `--time` | `-t` | 恢复到指定时间戳 |
| `--name` | | 恢复到命名还原点 |
| `--lsn` | | 恢复到指定 LSN，例如 `0/7C82CB8` |
| `--xid` | | 恢复到指定正整数事务 ID |
{.full-width}

`--set` 只选择备份，仍需指定恢复目标。

### 备份与目标选项

| 参数 | 简写 | 说明 |
|:---|:---|:---|
| `--set` | `-b` | `pig pb info` 列出的备份标签 |
| `--target-action` | | 到达目标后的动作：`pause`、`promote` 或 `shutdown` |
| `--target-timeline` | `-T` | 时间线：`latest`、`current`、正整数或 `0xHEX` |
| `--exclusive` | `-X` | 在时间、LSN 或事务 ID 目标之前停止 |
{.full-width}

`--target-action=shutdown` 必须配合 `--no-restart`。准备重放 WAL 时再手动启动 PostgreSQL，
它会在到达目标后关闭。

`--target-action` 不能与 `--default` 同用。`--exclusive` 要求同时指定 `--time`、`--lsn` 或 `--xid`。

### 流程控制

| 参数 | 简写 | 默认值 | 说明 |
|:---|:---|:---|:---|
| `--no-restart` | | false | 恢复文件后保持 PostgreSQL 停止 |
| `--plan` | | false | 查看计划，不修改数据或服务 |
| `--yes` | `-y` | false | 跳过确认 |
| `--timeout` | | 120 | 启动和恢复等待超时，单位秒 |
| `--force-stop` | | false | 快速关闭失败时，允许立即关闭及终止进程 |
{.full-width}

### 配置参数

| 参数 | 简写 | 默认值或来源 |
|:---|:---|:---|
| `--stanza` | `-s` | 配置文件中的第一个 stanza；有多个时应显式指定 |
| `--config` | `-c` | `/etc/pgbackrest/pgbackrest.conf` |
| `--repo` | `-r` | 由 pgBackRest 选择仓库 |
| `--dbsu` | `-U` | PostgreSQL 操作系统用户：`$PIG_DBSU`，其次为 `postgres` |
| `--data` | `-D` | 目标目录；配置中的 `pg1-path`、`$PGDATA`、`/pg/data`，按此顺序选择 |
{.full-width}

### 原生恢复参数

额外的 pgBackRest 参数放在 `--` 之后：

```bash
pig pitr -d -- --delta
```

恢复目标、数据目录、配置文件、stanza 和仓库使用 PIG 参数设置。
覆盖这些设置的原生参数会被拒绝，包括 `--repoN-*` 和 `--recovery-option`。
其他恢复参数可以转发，例如 `--delta`、`--tablespace-map`、`--link-map` 和 `--link-all`。

## 时间格式

PIG 按当前日期和本地时区补全缺少的部分：

| 格式 | 示例 | 含义 |
|:---|:---|:---|
| 完整时间戳 | `2025-01-01 12:00:00+08` | 显式日期、时间和时区 |
| 无时区 | `2025-01-01 12:00:00` | 使用本地时区，也接受 `T` 分隔符 |
| 仅日期 | `2025-01-01` | 该日期零点 |
| 仅时间 | `12:00:00` | 今天的该时间 |
{.full-width}

脚本中使用带时区的完整时间戳。计划和建议的后续命令会将简写展开为完整时间戳，
使重复执行时仍使用相同目标。

## 托管目录与旁路恢复

托管目录依次取自所选 stanza 的 `pg1-path`、`$PGDATA`、`/pg/data`，其中必须已有初始化的 PostgreSQL 实例。
显式 `-D` 解析后仍指向此路径时，继续按托管恢复处理，软链接也按实际目标判断。

显式 `-D` 指向其他路径时，执行 **旁路恢复**（side restore）。
PIG 保留 Patroni 和托管实例的当前状态。旁路目录必须已存在、归所选操作系统用户所有，
并允许该用户写入；可以是空目录。

`-D` 只指定数据目录。备份包含表空间或外部链接时，也要将它们映射到独立路径。
在恢复命令的 `--` 后添加 pgBackRest 的 `--tablespace-map` 和 `--link-map` 参数。

旁路恢复必须指定 `--no-restart`。恢复出的配置仍保留原端口，使用空闲端口手动启动：

```bash
# 以 root 或通过 sudo 准备目录
sudo install -d -m 700 -o postgres -g postgres /tmp/pg-restore

# 将备份文件恢复到旁路目录
pig pitr -d -D /tmp/pg-restore --no-restart

# 使用空闲端口启动旁路实例
pig pg start -D /tmp/pg-restore -O "-p 5433"
pig pg psql -D /tmp/pg-restore
```

`pig pg psql -D` 读取恢复目录中的 `postmaster.pid`，使用其中的端口和 socket 连接实例。
无法解析该实例时返回错误。

## 执行流程

1. 检查参数、配置、目标目录、stanza 状态和可用备份。指定 `--set` 时，检查备份标签是否存在。
2. 显示操作并请求确认；指定 `--yes` 时跳过确认。
3. 托管恢复时，停止正在运行的 Patroni 服务。
4. 等待 PostgreSQL 停止，必要时重试快速关闭。`--force-stop` 允许立即关闭及最终的进程终止尝试。
5. 执行 `pgbackrest restore`。
6. 启动 PostgreSQL，除非指定 `--no-restart`。对于 `--default` 或 `--target-action=promote`，
   等待 `pg_is_in_recovery()` 变为 false。其他目标需要自行检查恢复进度。
7. 输出数据验证和后续恢复步骤的命令。

拒绝确认或输入到达 EOF 时，返回 `CodePITRConfirmationRequired`（`161001`，退出码 `7`），
不会停止 Patroni 或 PostgreSQL，也不会执行恢复或启动 PostgreSQL。

PostgreSQL 运行状态无法确定时会报错，不会将其视为已经停止。
PIG 不会覆盖恢复到状态未确认的实例，也不会用未经验证的 PID 强制终止进程。
恢复后状态仍然未知时，不能判定启动成功，也不会向未确认的连接目标执行后续 SQL 检查。

`--plan` 执行预检查并返回计划，需要访问配置、数据目录和备份仓库，但不会停止服务或恢复数据。

## 更多示例

按需要选择一条恢复命令：

```bash
# 恢复到最新可用状态，覆盖已有目录
pig pitr -d -- --delta

# 恢复到备份一致性点
pig pitr -I -- --delta

# 选择具体备份标签
pig pitr -d -b 20251225-120000F -- --delta

# 先恢复文件，检查后再启动 PostgreSQL
pig pitr -d --no-restart -- --delta

# 到达时间目标后自动提升
pig pitr -t "2025-01-15 09:30:00+08" --target-action=promote -- --delta
```

## 恢复后操作 {#post-recovery-actions}

1. 连接恢复后的实例，使用实际数据目录：

   ```bash
   pig pg psql -D /pg/data
   ```

2. 检查恢复状态，再验证需要的表和记录：

   ```sql
   SELECT pg_is_in_recovery() AS recovering,
          CASE WHEN pg_is_in_recovery() THEN pg_is_wal_replay_paused() END AS paused;
   ```

3. 如果恢复已在预期目标暂停，并且需要可写主库，执行提升：

   ```bash
   pig pg promote -D /pg/data
   ```

4. 对于 Patroni 托管集群，先按集群恢复方案确认恢复实例的角色和状态，再用 `pig pt svc start`
   恢复 Patroni 管理。其他成员和应用流量也按该方案处理。

托管恢复在停止 Patroni 后失败时，Patroni 保持停止。目标目录可能只恢复了一部分。
检查 PostgreSQL 和 pgBackRest 日志，修复问题并完成恢复，再启动 Patroni。
启动超时也需要检查原因；它不会撤销已经恢复的文件。

## 结构化输出

使用 `--plan -o json` 查看结构化计划，使用 `--yes -o json` 无交互执行。也可以选择 YAML。

成功结果的 `data` 包含 `requested_data_dir`、`effective_data_dir`、`managed_data_dir` 和 `side_restore`。
后续命令位于结果顶层的 `next_actions` 字段。

## 设计记录

[PIG 运维 CLI 安全契约](https://pig.pgsty.com/zh/design/ops-cli-safety/)解释 `pig pb restore` 与 `pig pitr` 的分工。
