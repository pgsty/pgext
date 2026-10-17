---
title: "pig pb"
description: "使用 pig pb 子命令管理 pgBackRest 备份与时间点恢复"
icon: fa-solid fa-box-archive
weight: 660
---

`pig pgbackrest`（别名 `pig pb`）用于创建备份、恢复数据和管理 pgBackRest stanza。
PIG 以 PostgreSQL 操作系统用户运行 pgBackRest，依次从 `--dbsu`、`$PIG_DBSU` 选择用户，默认 `postgres`。

首次备份前，先安装 pgBackRest，在 `pgbackrest.conf` 中配置 stanza 和备份仓库。
一个 *stanza* 对应一个 PostgreSQL 集群的备份配置。
在 PostgreSQL 中启用 `archive_mode`，并设置 `archive_command` 调用 pgBackRest。

需要同时停止 Patroni、停止 PostgreSQL 并在恢复后启动 PostgreSQL 时，使用 [`pig pitr`](/zh/pig/cmd/pitr/)。
自行管理这些步骤时，使用 `pig pb restore`。

## 命令概览

**信息查询**：

| 命令        | 缩写   | 描述              | 实现方式              |
|:----------|------|:----------------|:------------------|
| `pb info` | `i`  | 显示备份仓库信息        | `pgbackrest info` |
| `pb list` | `ls` | 列出备份集、仓库和 stanza | `pgbackrest info` 与配置文件 |
{.full-width}

**备份与恢复**：

| 命令           | 缩写  | 描述       | 实现方式                 |
|:-------------|-----|:---------|:---------------------|
| `pb backup`  | `b` | 创建备份     | `pgbackrest backup`  |
| `pb restore` | `r` | 恢复备份文件 | `pgbackrest restore` |
| `pb expire`  | `e` | 清理过期备份   | `pgbackrest expire`  |
{.full-width}

**Stanza 管理**：

| 命令           | 缩写  | 描述                   | 实现方式                        |
|:-------------|-----|:---------------------|:----------------------------|
| `pb create`  | `c` | 创建 stanza（首次设置）      | `pgbackrest stanza-create`  |
| `pb upgrade` | `u` | 升级 stanza（PG 大版本升级后） | `pgbackrest stanza-upgrade` |
| `pb delete`  | `d` | 删除 stanza 及其备份     | `pgbackrest stanza-delete`  |
{.full-width}

**控制命令**：

| 命令         | 别名   | 描述               | 实现方式               |
|:-----------|:-----|:-----------------|:-------------------|
| `pb check` | `ck` | 检查 stanza 配置与 WAL 归档        | `pgbackrest check` |
| `pb start` | `up` | 启用 pgBackRest 操作 | `pgbackrest start` |
| `pb stop`  | `dw` | 禁用 pgBackRest 操作 | `pgbackrest stop`  |
| `pb log`   | `l`  | 查看日志             | 最新日志快照 / tail      |
{.full-width}

## 创建首次备份 {#快速入门}

配置好 pgBackRest 与 WAL 归档后，创建 stanza 并检查配置：

```bash
pig pb create
pig pb check
```

检查通过后，创建全量备份并查看结果：

```bash
pig pb backup full
pig pb info
```

## 全局参数

以下参数适用于所有 `pig pb` 子命令：

| 参数         | 简写   | 说明                                   |
|:-----------|:-----|:-------------------------------------|
| `--stanza` | `-s` | pgBackRest stanza 名称（自动检测）           |
| `--config` | `-c` | 配置文件路径                               |
| `--repo`   | `-r` | 仓库编号（多仓库场景）                          |
| `--dbsu`   | `-U` | PostgreSQL 操作系统用户（默认：`$PIG_DBSU` 或 `postgres`） |
{.full-width}

**Stanza 自动检测：**

如果未指定 `-s` 参数，PIG 会从配置文件中自动检测 stanza 名称：

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

创建物理备份。PIG 在执行 pgBackRest 前检查本地实例是否为主库。

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
| (空) | 使用 pgBackRest 配置；默认增量，没有全量备份时改为全量 |
| full | 全量备份：备份所有数据       |
| diff | 差异备份：自上次全量备份以来的变更 |
| incr | 增量备份：自上次任意备份以来的变更 |
{.full-width}

备库、已停止的实例或无法确定的角色都不能通过检查。
如果需要按 pgBackRest 配置处理主备角色，使用 `--force` 跳过 PIG 检查；pgBackRest 仍执行自己的备份检查。

### pb expire

按配置的保留策略删除备份和 WAL 归档。
`--set` 使用完整备份标签，还会删除依赖该备份的其他备份。此次操作仍会执行 WAL 保留策略。
参见 [pgBackRest expire 参考](https://pgbackrest.org/command.html#command-expire)。

```bash
pig pb expire                        # 按策略清理
pig pb expire --set 20250101-120000F       # 删除特定备份集
pig pb expire --set 20250101-120000F -y    # 跳过删除备份集确认
pig pb expire --plan                 # 仅预览清理计划，不删除
```

**选项：**

| 参数       | 简写   | 说明                  |
|:---------|:-----|:--------------------|
| `--set`  |      | 删除特定备份集             |
| `--plan` |      | 仅预览清理计划，不删除备份       |
| `--yes`  | `-y` | 与 `--set` 配合时跳过确认提示 |
{.full-width}

`--plan` 执行 pgBackRest 试运行。使用结构化输出时，PIG 计划会包含可用的试运行结果。
通过该结果检查将受影响的备份。

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

将备份文件恢复到目标数据目录。PostgreSQL 必须已停止。
必须且只能选择一个恢复目标：`-d`、`-I`、`-t`、`--name`、`--lsn` 或 `--xid`。
缺少目标时返回错误，文本模式还会显示帮助。

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
pig pb restore -t "2025-01-01 12:00:00+08" -X           # 排他模式（在目标前停止）
pig pb restore -t "2025-01-01 12:00:00+08" --target-action=promote  # 到达目标后提升
pig pb restore -t "2025-01-01 12:00:00+08" -T current   # 沿当前时间线恢复
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

使用 PIG 参数设置恢复目标、数据目录、stanza、仓库和配置文件。
PIG 会拒绝覆盖这些设置的原生参数，包括 `--repoN-*` 和 `--recovery-option`。
其他恢复参数放在 `--` 之后，例如 `--delta`、`--tablespace-map`、`--link-map` 或 `--link-all`。
恢复到已有数据目录时，使用 `-- --delta` 按校验和复用文件并恢复差异。

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

对于数据目录为 `/pg/data` 的独立实例，先查看计划，停止实例，再执行恢复：

```bash
pig pb restore -D /pg/data -t "2025-01-01 12:00:00+08" --plan
pig pg stop -D /pg/data
pig pb restore -D /pg/data -t "2025-01-01 12:00:00+08" -- --delta
pig pg start -D /pg/data
pig pg psql -D /pg/data
```

使用其他路径时，为每条命令指定相同的 `-D`。提升主库前，先验证恢复的数据。
目标为托管目录且 Patroni 正在运行时，`pb restore` 会拒绝执行；使用 [`pig pitr`](/zh/pig/cmd/pitr/)。

文本模式会在恢复前要求输入 `y` 或 `yes`。脚本使用 `-o json` 或 `-o yaml` 时，
加 `--yes` 执行，或用 `--plan` 查看计划。

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

删除会永久移除所选 stanza 的备份。文本模式下输入 `y` 或 `yes` 确认，或使用 `--yes`。
结构化执行必须加 `--yes`。配置包含多个 stanza 时，用 `--stanza` 指定目标。

## 控制命令

### pb check

检查 stanza 配置、仓库访问和 WAL 归档。

```bash
pig pb check                         # 检查配置与归档
```

此命令检查：

- WAL 归档配置是否正确
- 仓库是否可访问
- stanza 配置是否有效

`check` 不会逐一读取备份文件并校验内容。需要校验备份与归档的校验和时，使用 pgBackRest
[`verify`](https://pgbackrest.org/command.html#command-verify)。

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
| `--follow` | `-f` | false | 使 `pb log` 持续跟踪；`pb log tail` 始终跟踪 |
{.full-width}

**权限处理：**

如果当前用户没有权限读取日志目录，命令会自动使用 `sudo` 重试。

## 路径与权限

| 设置 | 默认值 |
|:---|:---|
| 配置文件 | `/etc/pgbackrest/pgbackrest.conf` |
| 日志目录 | 配置中的 `log-path`，其次为 `/pg/log/pgbackrest` |
| 恢复数据目录 | `--data`、配置中的 `pg1-path`、`$PGDATA`、`/pg/data`，按此顺序选择 |
{.full-width}

当前用户是所选 PostgreSQL 操作系统用户时，PIG 直接运行 pgBackRest。
root 使用 `su`；其他用户使用 `sudo -H -u`。
该用户需要读取配置文件，并访问数据目录和备份仓库。

这些命令采用 Linux 和 Pigsty 默认路径。其他目录布局可显式指定配置文件和 stanza。
PIG 未提供的 pgBackRest 操作，直接使用 `pgbackrest` 命令。

## 设计记录

[PIG 运维 CLI 安全契约](https://pig.pgsty.com/zh/design/ops-cli-safety/)解释`pb restore` 与 `pig pitr` 恢复流程的分工。
