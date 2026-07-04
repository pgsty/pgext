---
title: "pig pb"
description: "使用 pig pb 子命令管理 pgBackRest 备份与时间点恢复"
icon: Archive
weight: 660
---

`pig pgbackrest` 命令（别名 `pig pb`）用于管理 pgBackRest 备份，并提供低层 restore 原语。
它封装了常用的 `pgbackrest` 操作，提供简化的备份管理体验。所有命令均以数据库超级用户身份（默认 `postgres`）执行。
托管集群的编排式时间点恢复请优先使用 [`pig pitr`](/zh/pig/cmd/pitr/)。

```bash
pig pb - Manage pgBackRest backup and point-in-time recovery.

Information:
  pig pb info                      show backup info
  pig pb list                      list backups
  pig pb list repo                 list configured repositories
  pig pb list stanza               list all stanzas

Backup & Restore:
  pig pb backup                    create backup (auto: full/incr)
  pig pb backup full               create full backup
  pig pb restore                   low-level restore primitive
  pig pb restore -t "..."          restore to specific time
  pig pb expire                    cleanup expired backups

Stanza Management:
  pig pb create                    create stanza (first-time setup)
  pig pb upgrade                   upgrade stanza (after PG upgrade)
  pig pb delete                    delete stanza (DANGEROUS!)

Control:
  pig pb check                     verify backup integrity
  pig pb start                     enable pgBackRest operations
  pig pb stop                      disable pgBackRest operations
  pig pb log                       view pgBackRest logs

Examples:
  pig pb info                      # show all backup info
  pig pb backup                    # auto: full if none, else incr
  pig pb backup full               # full backup
  pig pb restore -d                # restore to latest (end of WAL)
  pig pb restore -t "2025-01-01 12:00:00+08"  # restore to time
  pig pb create                    # initialize stanza
  pig pb expire                    # cleanup per retention policy
```

## 命令概览

**信息查询**：

| 命令        | 缩写   | 描述              | 实现方式              |
|:----------|------|:----------------|:------------------|
| `pb info` | `i`  | 显示备份仓库信息        | `pgbackrest info` |
| `pb list` | `ls` | 列出备份集，仓库，Stanza | `pgbackrest info` |
{.full-width}

**备份与恢复**：

| 命令           | 缩写  | 描述       | 实现方式                 |
|:-------------|-----|:---------|:---------------------|
| `pb backup`  | `b` | 创建备份     | `pgbackrest backup`  |
| `pb restore` | `r` | 低层备份恢复原语 | `pgbackrest restore` |
| `pb expire`  | `e` | 清理过期备份   | `pgbackrest expire`  |
{.full-width}

**Stanza 管理**：

| 命令           | 缩写  | 描述                   | 实现方式                        |
|:-------------|-----|:---------------------|:----------------------------|
| `pb create`  | `c` | 创建 stanza（首次设置）      | `pgbackrest stanza-create`  |
| `pb upgrade` | `u` | 升级 stanza（PG 大版本升级后） | `pgbackrest stanza-upgrade` |
| `pb delete`  | `d` | 删除 stanza（危险操作！）     | `pgbackrest stanza-delete`  |
{.full-width}

**控制命令**：

| 命令         | 别名   | 描述               | 实现方式               |
|:-----------|:-----|:-----------------|:-------------------|
| `pb check` | `ck` | 验证备份仓库完整性        | `pgbackrest check` |
| `pb start` | `up` | 启用 pgBackRest 操作 | `pgbackrest start` |
| `pb stop`  | `dw` | 禁用 pgBackRest 操作 | `pgbackrest stop`  |
| `pb log`   | `l`  | 查看日志             | 最新日志快照 / tail      |
{.full-width}


## 快速入门

```bash
# 查看备份信息
pig pb info                          # 显示所有备份信息
pig pb info -o json                  # Result 包装的结构化 JSON 输出
pig pb info --raw --raw-output json  # pgBackRest 原生 JSON 输出
pig pb ls                            # 列出所有备份
pig pb ls repo                       # 列出配置的仓库
pig pb ls stanza                     # 列出所有 stanza

# 创建备份（必须在主库执行）
pig pb backup                        # 自动模式：无备份则全量，否则增量
pig pb backup full                   # 全量备份
pig pb backup diff                   # 差异备份
pig pb backup incr                   # 增量备份

# 恢复（低层 restore 原语，至少指定一个恢复目标）
pig pb restore -d                    # 恢复到最新（WAL 流末尾）
pig pb restore -I                    # 恢复到备份一致性点
pig pb restore -t "2025-01-01 12:00:00+08"  # 恢复到指定时间
pig pb restore --name savepoint      # 恢复到命名还原点

# Stanza 管理
pig pb create                        # 初始化 stanza
pig pb upgrade                       # PG 大版本升级后升级 stanza
pig pb check                         # 验证仓库完整性

# 清理
pig pb expire                        # 按保留策略清理
pig pb expire --plan                 # 仅预览清理计划
```


## 全局参数

以下参数适用于所有 `pig pb` 子命令：

| 参数         | 简写   | 说明                                   |
|:-----------|:-----|:-------------------------------------|
| `--stanza` | `-s` | pgBackRest stanza 名称（自动检测）           |
| `--config` | `-c` | 配置文件路径                               |
| `--repo`   | `-r` | 仓库编号（多仓库场景）                          |
| `--dbsu`   | `-U` | 数据库超级用户（默认：`$PIG_DBSU` 或 `postgres`） |
{.full-width}

**Stanza 自动检测：**

如果未指定 `-s` 参数，pig 会从配置文件中自动检测 stanza 名称：

1. 读取配置文件（默认 `/etc/pgbackrest/pgbackrest.conf`）
2. 查找非 `[global*]` 开头的 section
3. 使用找到的第一个 stanza

如果配置文件中有多个 stanza，会发出警告并使用第一个。此时应显式指定 `--stanza` 参数。

**多仓库支持：**

pgBackRest 支持配置多个仓库（repo1、repo2 等）。使用 `-r` 参数指定操作的目标仓库：

```bash
pig pb backup -r 1                   # 备份到 repo1
pig pb backup -r 2                   # 备份到 repo2
pig pb info -r 2                     # 查看 repo2 的备份信息
```


## 信息查询命令

### pb info

显示备份仓库详细信息，包括所有备份集和 WAL 归档状态。

```bash
pig pb info                          # 显示所有备份信息
pig pb info --raw --raw-output json  # 原始 JSON 输出
pig pb info --set 20250101-120000F   # 显示特定备份集详情
```

**选项：**

| 参数             | 简写   | 说明                             |
|:---------------|:-----|:-------------------------------|
| `--raw`        | `-R` | 原始输出模式（透传 pgBackRest 输出）       |
| `--raw-output` |      | 原始输出格式：text、json（仅 `--raw` 模式） |
| `--set`        |      | 显示特定备份集详情                      |
{.full-width}


### pb ls

列出备份仓库中的资源。

```bash
pig pb ls                            # 列出所有备份（默认）
pig pb ls backup                     # 列出所有备份（显式）
pig pb ls repo                       # 列出配置的仓库
pig pb ls stanza                     # 列出所有 stanza
```

**类型说明：**

| 类型     | 描述          | 数据来源               |
|:-------|:------------|:-------------------|
| backup | 列出所有备份集（默认） | pgbackrest info    |
| repo   | 列出配置的仓库     | 解析 pgbackrest.conf |
| stanza | 列出所有 stanza | 解析 pgbackrest.conf |
{.full-width}


## 备份命令

### pb backup

创建物理备份。备份只能在主库实例上执行。

```bash
pig pb backup                        # 自动模式
pig pb backup full                   # 全量备份
pig pb backup diff                   # 差异备份
pig pb backup incr                   # 增量备份
pig pb backup --force                # 跳过主库角色检查
```

**选项：**

| 参数        | 简写   | 说明       |
|:----------|:-----|:---------|
| `--force` | `-f` | 跳过主库角色检查 |
{.full-width}

**备份类型：**

| 类型   | 说明                |
|:-----|:------------------|
| (空)  | 自动模式：无备份则全量，否则增量  |
| full | 全量备份：备份所有数据       |
| diff | 差异备份：自上次全量备份以来的变更 |
| incr | 增量备份：自上次任意备份以来的变更 |
{.full-width}

**主库检查：**

执行备份前，命令会自动检查当前实例是否为主库。如果是备库，命令会报错退出。使用 `--force` 可跳过此检查。




### pb expire

按保留策略清理过期的备份和 WAL 归档。

```bash
pig pb expire                        # 按策略清理
pig pb expire --set 20250101-*       # 删除特定备份集
pig pb expire --set 20250101-* -y    # 跳过删除备份集确认
pig pb expire --plan                 # 仅预览清理计划，不删除
```

**选项：**

| 参数       | 简写   | 说明                  |
|:---------|:-----|:--------------------|
| `--set`  |      | 删除特定备份集             |
| `--plan` |      | 仅预览清理计划，不删除备份       |
| `--yes`  | `-y` | 与 `--set` 配合时跳过确认提示 |
{.full-width}

**保留策略配置：**

保留策略在 `pgbackrest.conf` 中配置：

```ini
[global]
repo1-retention-full=2               # 保留的全量备份数
repo1-retention-diff=4               # 保留的差异备份数
repo1-retention-archive=2            # WAL 归档保留策略
```


## 恢复命令

### pb restore

从备份恢复，支持时间点恢复目标。
必须显式指定一个恢复目标（`-d/-I/-t/--name/--lsn/--xid`）；不带参数仅显示帮助信息。

```bash
# 恢复目标（互斥选项）
pig pb restore -d                    # 恢复到最新（显式）
pig pb restore -I                    # 恢复到备份一致性点
pig pb restore -t "2025-01-01 12:00:00+08"  # 恢复到指定时间
pig pb restore -t "2025-01-01"       # 恢复到日期（当天 00:00:00）
pig pb restore -t "12:00:00"         # 恢复到时间（今天）
pig pb restore --name my-savepoint   # 恢复到命名还原点
pig pb restore --lsn "0/7C82CB8"     # 恢复到 LSN
pig pb restore --xid 12345           # 恢复到事务 ID

# 备份集选择（必须与恢复目标组合）
pig pb restore -b 20251225-120000F -d  # 从特定备份集恢复到最新

# 其他选项
pig pb restore -t "..." -X           # 排他模式（在目标前停止）
pig pb restore -t "..." --target-action=promote  # 到达目标后提升
pig pb restore -t "..." -T current   # 沿当前时间线恢复
pig pb restore -d --plan             # 仅预览恢复计划
pig pb restore -d -y                 # 跳过 y/yes 交互确认
pig pb restore -d -- --delta         # 在 -- 后透传 pgBackRest restore 参数
```

**恢复目标选项：**

| 参数            | 简写   | 说明                |
|:--------------|:-----|:------------------|
| `--default`   | `-d` | 恢复到 WAL 流末尾（最新数据） |
| `--immediate` | `-I` | 恢复到备份一致性点         |
| `--time`      | `-t` | 恢复到指定时间戳          |
| `--name`      |      | 恢复到命名还原点          |
| `--lsn`       |      | 恢复到指定 LSN         |
| `--xid`       |      | 恢复到指定事务 ID        |
{.full-width}

**备份集和其他选项：**

| 参数                  | 简写   | 说明                                       |
|:--------------------|:-----|:-----------------------------------------|
| `--set`             | `-b` | 从特定备份集恢复（可与目标组合）                         |
| `--data`            | `-D` | 目标数据目录                                   |
| `--exclusive`       | `-X` | 排他模式：在目标前停止                              |
| `--target-action`   |      | 到达恢复目标后的动作：pause/promote/shutdown        |
| `--target-timeline` | `-T` | 恢复时间线：latest/current/N/0xN               |
| `--plan`            |      | 仅预览恢复计划，不执行                              |
| `--yes`             | `-y` | 跳过交互式 y/yes 确认                           |
| `--` 后的原生参数         |      | 透传 pgBackRest restore 参数，例如 `-- --delta` |
{.full-width}

**组合规则：** `--target-action` 不能与 `--default` 同时使用，因为 `--default` 已表示恢复到 WAL 末尾。`--exclusive/-X` 必须配合明确的停止目标使用：`--time`、`--lsn` 或 `--xid`。

`--` 后的原生 pgBackRest restore 参数不能覆盖 Pig 已管理的恢复目标、生命周期、数据目录、仓库、配置和选择参数；请使用 Pig 的一等参数设置这些语义。表空间/链接迁移类参数（如 `--tablespace-map`、`--link-map`、`--link-all`）仍可透传。

**时间格式：**

支持多种时间格式输入，自动补全时区（支持非整小时时区如 +05:30）：

| 格式   | 示例                       | 说明                     |
|:-----|:-------------------------|:-----------------------|
| 完整格式 | `2025-01-01 12:00:00+08` | 包含时区的完整时间戳             |
| 仅日期  | `2025-01-01`             | 自动补全为当天 00:00:00（当前时区） |
| 仅时间  | `12:00:00`               | 自动补全为今天（当前时区）          |
{.full-width}

**恢复流程：**

1. 验证参数和环境
2. 检查 PostgreSQL 已停止
3. 显示恢复计划，等待交互式 `y`/`yes` 确认
4. 执行 pgbackrest restore
5. 提供恢复后的操作指引

**重要提示：** 恢复前必须先停止 PostgreSQL；如果该 PGDATA 由 Patroni 管理，应使用 [`pig pitr`](/zh/pig/cmd/pitr/) 编排 Patroni、PostgreSQL 与 pgBackRest：

```bash
pig pg stop                          # 停止 PostgreSQL
pig pb restore -t "..."              # 执行恢复
pig pg start                         # 启动 PostgreSQL
```


## Stanza 管理命令

### pb create

初始化新的 stanza。必须在首次备份前执行。

```bash
pig pb create                        # 创建 stanza
pig pb create --no-online            # PostgreSQL 未运行时创建
pig pb create --force                # 强制创建
```

**选项：**

| 参数            | 简写   | 说明                |
|:--------------|:-----|:------------------|
| `--no-online` |      | PostgreSQL 未运行时创建 |
| `--force`     | `-f` | 强制创建              |
{.full-width}


### pb upgrade

PostgreSQL 大版本升级后更新 stanza。

```bash
pig pb upgrade                       # 升级 stanza
pig pb upgrade --no-online           # PostgreSQL 未运行时升级
```

**选项：**

| 参数            | 说明                |
|:--------------|:------------------|
| `--no-online` | PostgreSQL 未运行时升级 |
{.full-width}

**使用场景：**

当 PostgreSQL 进行大版本升级（如 16 → 17）后，需要执行此命令更新 stanza 元数据。


### pb delete

删除 stanza 及其所有备份。

```bash
pig pb delete                        # 删除 stanza（交互式 y/N 确认）
pig pb delete --yes                  # 跳过 y/yes 交互确认
pig pb delete --plan                 # 预览删除计划
```

**选项：**

| 参数       | 简写   | 说明             |
|:---------|:-----|:---------------|
| `--plan` |      | 仅预览删除计划，不执行    |
| `--yes`  | `-y` | 跳过交互式 y/yes 确认 |
{.full-width}

**警告：** 这是一个**破坏性且不可逆**的操作！所有备份将被永久删除。

命令包含多重安全机制：
1. 文本模式下要求交互式 `y`/`yes` 确认，除非指定 `--yes`
2. 结构化输出模式必须显式指定 `--yes`
3. 配置文件存在多个 stanza 且未显式 `--stanza` 时，拒绝自动选择删除目标


## 控制命令

### pb check

验证备份仓库的完整性和配置。

```bash
pig pb check                         # 验证仓库
```

此命令检查：
- WAL 归档配置是否正确
- 仓库是否可访问
- stanza 配置是否有效


### pb start

启用 pgBackRest 操作。

```bash
pig pb start                         # 启用操作
```

在执行 `pb stop` 后使用此命令恢复正常操作。


### pb stop

禁用 pgBackRest 操作（用于维护）。

```bash
pig pb stop                          # 禁用操作
pig pb stop --force                  # 终止正在运行的操作
```

**选项：**

| 参数        | 简写   | 说明        |
|:----------|:-----|:----------|
| `--force` | `-f` | 终止正在运行的操作 |
{.full-width}

**使用场景：**

在进行系统维护时，使用此命令阻止新的备份操作启动。


## 日志命令

### pb log

查看 pgBackRest 日志文件。日志目录优先读取 pgBackRest 配置中的 `log-path`，未配置时使用 `/pg/log/pgbackrest/`。父命令默认显示最新日志快照，实时跟踪请使用 `tail` 或 `-f`。只有 `pb log` 与 `pb log show` 支持 `-o json` 输出 JSONL；日志快照不支持 `yaml` 与 `json-pretty`，follow/tail 不支持结构化输出。

```bash
pig pb log                           # 显示最新日志行
pig pb log list                      # 列出日志文件
pig pb log tail                      # 实时查看最新日志
pig pb log tail -n 100               # 显示最后 100 行并跟踪
pig pb log show                      # 显示最新日志内容
pig pb log show -n 50                # 显示最后 50 行
pig pb log cat                       # show 的别名
```

**子命令：**

| 子命令    | 别名             | 说明           |
|:-------|:---------------|:-------------|
| list   | ls             | 列出日志文件       |
| show   | cat, c         | 显示最新日志内容     |
| tail   | t, f, follow   | 实时跟踪最新日志     |
{.full-width}

**选项：**

| 参数         | 简写   | 默认值   | 说明                                                         |
|:-----------|:-----|:------|:-----------------------------------------------------------|
| `--lines`  | `-n` | 50    | 显示的行数                                                      |
| `--follow` | `-f` | false | 父命令 `pb log` 使用该参数进入跟踪；`pb log tail` 中为 no-op，因为 tail 总是跟踪 |
{.full-width}

**权限处理：**

如果当前用户没有权限读取日志目录，命令会自动使用 `sudo` 重试。


## 设计说明

**命令执行方式：**

所有 `pig pb` 命令都以数据库超级用户（DBSU）身份执行。这是因为 pgBackRest 需要访问 PostgreSQL 数据文件和 WAL 归档。

执行逻辑：
- 如果当前用户已是 DBSU：直接执行命令
- 如果当前用户是 root：使用 `su - postgres -c "..."` 执行
- 其他用户：使用 `sudo -inu postgres -- ...` 执行

**与 pgbackrest 的关系：**

`pig pb` 并非 `pgbackrest` 的完整封装，而是针对常用操作的上层抽象：

- 自动检测 stanza 名称，无需每次指定
- 备份前自动检查主库角色
- 恢复时显示计划并要求交互式 `y`/`yes` 确认
- 提供人性化的时间格式输入
- 恢复后提供操作指引

如需使用 `pgbackrest` 的完整功能，请直接使用 `pgbackrest` 命令。

**默认配置路径：**

| 配置项  | 默认值                                             |
|:-----|:------------------------------------------------|
| 配置文件 | `/etc/pgbackrest/pgbackrest.conf`               |
| 日志目录 | `/pg/log/pgbackrest`                            |
| 数据目录 | 配置文件中的 `pg1-path`，或 `$PGDATA` 环境变量，或 `/pg/data` |
{.full-width}

**安全考虑：**

- `pb delete` 在没有 `--yes` 时要求交互式 `y`/`yes` 确认，多 stanza 配置下必须显式指定 `--stanza`
- `pb restore` 需要显式恢复目标，校验 `--time`，并在没有 `--yes` 时要求交互式 `y`/`yes` 确认
- `pb backup` 默认检查主库角色，防止在备库执行
- `pb log tail` 不支持结构化输出；需要 JSONL 快照时使用 `pb log show -n N -o json`

**平台支持：**

此命令专为 Linux 系统设计，依赖 Pigsty 的默认目录结构。
