---
title: "pig pg"
description: "使用 pig pg 子命令管理本地 PostgreSQL 服务器"
icon: fa-solid fa-database
weight: 640
---

`pig postgres`（别名 `pig pg`）管理本地 PostgreSQL 实例和数据库，
用于初始化与控制实例、执行 SQL、维护表和创建数据库副本。
Patroni 集群操作使用 [`pig pt`](/zh/pig/cmd/pt/)，pgBackRest 恢复流程使用 [`pig pitr`](/zh/pig/cmd/pitr/)。

命令以 PostgreSQL 操作系统用户运行，通常为 `postgres`。先安装 PostgreSQL。
初始化新数据目录前，为该用户准备可写的父目录，或创建归该用户所有的空目录。

## 命令概览

**实例控制**：

| 命令           | 别名           | 描述            | 备注                 |
|:-------------|:-------------|:--------------|:-------------------|
| `pg init`    | `initdb, i`  | 初始化数据目录       | 封装 initdb          |
| `pg start`   | `boot, up`   | 启动 PostgreSQL | 封装 pg_ctl start    |
| `pg stop`    | `halt, down` | 停止 PostgreSQL | 封装 pg_ctl stop     |
| `pg restart` | `reboot`     | 重启 PostgreSQL | 封装 pg_ctl restart  |
| `pg reload`  | `hup`        | 重载配置          | 封装 pg_ctl reload   |
| `pg status`  | `st, stat`   | 查看服务状态        | 显示进程与相关服务状态        |
| `pg promote` | `pro`        | 提升备库为主库       | 封装 pg_ctl promote  |
| `pg role`    | `r`          | 检测实例角色        | 输出 primary/replica |
{.full-width}

**连接与查询**：

| 命令         | 别名              | 描述      | 备注                                           |
|:-----------|:----------------|:--------|:---------------------------------------------|
| `pg psql`  | `sql, connect`  | 连接到数据库  | 封装 psql                                      |
| `pg ps`    | `activity, act` | 显示当前连接  | 查询 pg_stat_activity                          |
| `pg kill`  | `k`             | 终止连接    | 默认 dry-run 模式                                |
| `pg clone` |                 | 克隆单个数据库 | `CREATE DATABASE ... TEMPLATE ...` |
{.full-width}

**数据库维护**：

| 命令           | 别名        | 描述    | 备注                         |
|:-------------|:----------|:------|:---------------------------|
| `pg vacuum`  | `vac, vc` | 清理表   | 执行 SQL `VACUUM`                |
| `pg analyze` | `ana, az` | 分析表   | 执行 SQL `ANALYZE` |
| `pg freeze`  |           | 冻结清理表 | 执行 SQL `VACUUM (FREEZE)`       |
| `pg repack`  | `rp`      | 在线重整表 | 需要 pg_repack 扩展            |
{.full-width}

**参数调优**：

| 命令        | 别名       | 描述                 | 备注             |
|:----------|:---------|:-------------------|:---------------|
| `pg tune` | `tuning` | 生成 PostgreSQL 调优参数 | 自动探测硬件，支持结构化输出 |
{.full-width}

**实例 Fork**：

| 命令              | 别名               | 描述                | 备注                        |
|:----------------|:-----------------|:------------------|:--------------------------|
| `pg fork`       |                  | `fork init` 的便捷写法 | 默认创建托管 fork，不自动启动         |
| `pg fork init`  | `create`         | 创建本地一次性物理副本       | 默认 `/pg/data-<name>`      |
| `pg fork list`  |                  | 列出托管 fork         | 扫描 `/pg/data-*`           |
| `pg fork start` |                  | 启动已有 fork         | 支持托管名或 `--dst-data` 非托管目录 |
| `pg fork stop`  |                  | 停止已有 fork         | 支持 shutdown mode          |
| `pg fork rm`    | `remove, delete` | 删除 fork           | 运行中的 fork 需 `--stop`      |
{.full-width}

**日志工具**：

| 命令            | 别名          | 描述        | 备注      |
|:--------------|:------------|:----------|:--------|
| `pg log`      | `l`         | 日志管理      | 父命令     |
| `pg log list` | `ls`        | 列出日志文件    |         |
| `pg log tail` | `t, f`      | 实时查看日志    | tail -f |
| `pg log show` | `cat, c`    | 输出日志内容    |         |
| `pg log less` | `vi, v`     | 用 less 查看 |         |
| `pg log grep` | `g, search` | 搜索日志      |         |
{.full-width}

**服务子命令**（`pg svc`，也可写作 `pg service` 或 `pg s`）：

| 命令               | 别名               | 描述             |
|:-----------------|:-----------------|:---------------|
| `pg svc start`   | `boot, up`       | 启动 postgres 服务 |
| `pg svc stop`    | `halt, dn, down` | 停止 postgres 服务 |
| `pg svc restart` | `reboot, rt`     | 重启 postgres 服务 |
| `pg svc reload`  | `rl, hup`        | 重载 postgres 服务 |
| `pg svc status`  | `st, stat`       | 显示服务状态         |
{.full-width}

## 连接实例 {#快速入门}

对于数据目录为 `/pg/data` 的运行中实例，先查看状态，再连接数据库：

```bash
pig pg status -D /pg/data
pig pg psql -D /pg/data
```

请使用实例的实际数据目录。在 `psql` 中执行 SQL，输入 `\q` 退出。
安装 PostgreSQL 并创建测试实例的步骤见[快速上手](/zh/pig/start/)。

## 全局参数

以下参数适用于所有 `pig pg` 子命令：

| 参数          | 简写   | 默认值        | 说明                          |
|:------------|:-----|:-----------|:----------------------------|
| `--version` | `-v` | 自动检测       | PostgreSQL 主版本号（形如 18，17）   |
| `--data`    | `-D` | 实例操作使用 `/pg/data` | 数据目录路径                      |
| `--dbsu`    | `-U` | `postgres` | PostgreSQL 操作系统用户；`$PIG_DBSU` 可覆盖默认值 |
{.full-width}

使用 `pig pg svc ...` 控制 `postgres` systemd 服务。
使用 `start`、`stop` 等实例命令，通过 `pg_ctl` 操作 `-D` 指定的目录，省略时使用 `/pg/data`。
显式传入空字符串或纯空白的 `-D` / `--data` 会返回参数错误，不会选择默认目标。

显式指定 `-D` 时，`psql`、`ps`、`kill`、`vacuum`、`analyze`、`freeze`、`repack` 连接该目录对应的实例，
`--all` 的数据库枚举也使用同一目标。使用 PID 文件记录的端口和 socket
前，先确认 PID 存活且属于 PostgreSQL 进程；陈旧 PID 或无法确认进程身份时，在执行 SQL 客户端前失败。
元数据缺失、无效、数据目录不匹配或无法确定 socket 时，在启动 SQL 客户端前失败，不回退到默认实例。
省略 `-D` 时，这些 SQL 命令保留原生客户端的默认连接方式，以及客户端可见的 `PGHOST`/`PGPORT`；
PIG 不追加主机或端口参数，也不会将连接绑定到 `/pg/data`。
需要显式指定仅监听 TCP 的实例时，请使用原生客户端的连接参数。`pg clone` 和 `pg fork` 保留独立的端口参数。

`promote`、`vacuum`、`repack` 的计划、确认重试和验证命令保留所选的数据目录、PostgreSQL 版本
和操作系统用户，包括显式的 `-D /pg/data`。省略 `-D` 时仍保留 SQL 客户端的默认连接方式。
`VACUUM FULL` 计划将确认标记为必需，与实际执行一致。
生成的维护命令将数据库名放在 `--` 之后，避免 `--all` 这样的库名在复制执行时变成选项。

`PIG_NON_INTERACTIVE=1` 与 JSON/YAML 输出关闭提示输入，并使用非交互 sudo。
原生透传保留其交互契约；`psql -f -` 等显式数据输入通道仍然可用。

**版本检测逻辑：**

1. 如果指定了 `-v`，使用指定 PG 大版本
2. 否则从数据目录的 `PG_VERSION` 文件读取版本
3. 否则使用 PIG 找到的默认 PostgreSQL 安装。

## 服务控制命令

### pg init

初始化 PostgreSQL 数据目录，封装 `initdb` 命令。

- 使用空目录，或 PostgreSQL 操作系统用户可以创建的路径。
- 默认启用数据校验和，使用 `-K` 关闭。
- 编码为 UTF8。PostgreSQL 17+ 使用内置 `C.UTF-8` locale；旧版本使用系统 `C.UTF-8`，不可用时使用 `C`。
- 已初始化的目录需要 `--force`，确认后会删除原数据。运行中的实例不能覆盖。无法确认运行状态时，
  在删除前失败并报告探测错误，不会误报 PostgreSQL 正在运行。
- 额外 `initdb` 参数放在 `--` 之后。需要其他 locale 或编码时，直接运行 `initdb`。

```bash
pig pg init                       # 使用默认设置初始化
pig pg init -v 18                 # 指定 PostgreSQL 18
pig pg init -D /data/pg18         # 指定数据目录
pig pg init -K                    # 关闭数据校验和
pig pg init -f                    # 强制初始化（删除已有数据）
pig pg init -- --waldir=/wal      # 传递额外参数给 initdb
```

**选项：**

| 参数                   | 简写   | 默认值   | 说明                |
|:---------------------|:-----|:------|:------------------|
| `--no-data-checksums` | `-K` | false | 禁用数据校验和                 |
| `--force`             | `-f` | false | 删除已初始化目录后重新初始化       |
| `--yes`               | `-y` | false | 与 `--force` 配合时跳过覆盖确认提示 |
{.full-width}

### pg start

使用 `pg_ctl start` 命令启动 PostgreSQL 服务器。
如果 PostgreSQL 由 Patroni 管理，建议使用 `pig pt start` 通过启动 patroni 的方式来启动。
如果 PostgreSQL 由 Systemd 管理，可以使用 `pig pg svc start` 来启动服务。

```bash
pig pg start                      # 使用默认设置启动
pig pg up                         # 别名
pig pg start -D /data/pg18        # 指定数据目录
pig pg start -l /pg/log/pg.log    # 重定向输出到日志文件
pig pg start -O "-p 5433"         # 传递参数给 postgres
pig pg start -o json              # 结构化输出 JSON
```

**选项：**

| 参数          | 简写   | 说明                      |
|:------------|:-----|:------------------------|
| `--log`     | `-l` | 启动日志；默认 `/tmp/pig-pg-start.log` |
| `--timeout` | `-t` | 等待超时（秒）；默认取 `$PGCTLTIMEOUT`，其次为 60                 |
| `--no-wait` |      | 不等待启动完成                 |
| `--options` | `-O` | 传递给 postgres 的选项        |
{.full-width}

如果 PostgreSQL 已经运行，命令会提示，并打印现有 Postmaster 进程 PID，不会报错。

### pg stop

使用 `pg_ctl stop` 命令停止 PostgreSQL 服务器。

Patroni 管理实例时，直接停止 PostgreSQL 可能触发自动重启或故障切换。
先用 `pig pt svc stop` 停止管理进程，再检查 PostgreSQL 状态。

```bash
pig pg stop                       # 快速停止（默认）
pig pg down                       # 别名
pig pg halt                       # 别名
pig pg stop -m smart              # 等待客户端断开
pig pg stop -m immediate          # 立即关闭
pig pg stop --plan                # 预览停止计划
```

**选项：**

| 参数          | 简写   | 默认值   | 说明                         |
|:------------|:-----|:------|:---------------------------|
| `--mode`    | `-m` | fast  | 关闭模式：smart/fast/immediate  |
| `--timeout` | `-t` | 60    | 等待超时（秒）；默认取 `$PGCTLTIMEOUT`，其次为 60                    |
| `--no-wait` |      | false | 不等待关闭完成                    |
| `--plan`    |      | false | 只预览本地 `pg_ctl stop` 计划，不执行 |
{.full-width}

**关闭模式说明：**

| 模式          | 说明                |
|:------------|:------------------|
| `smart`     | 等待所有客户端断开后关闭      |
| `fast`      | 回滚活动事务，断开客户端，正常关闭 |
| `immediate` | 立即终止所有进程，下次启动需要恢复 |
{.full-width}

### pg restart

使用 `pg_ctl restart` 重启 PostgreSQL 服务器。

```bash
pig pg restart                    # 快速重启
pig pg reboot                     # 别名
pig pg restart -m immediate       # 立即重启
pig pg restart -O "-p 5433"       # 使用新选项重启
pig pg restart --plan             # 预览重启计划
```

**选项：** 与 `pg stop` 相同，另外支持 `--options`（`-O`）传递给 postgres。

### pg reload

使用 `pg_ctl reload` 重载 PostgreSQL 配置。向服务器发送 SIGHUP 信号。

```bash
pig pg reload                     # 重载配置
pig pg hup                        # 别名
pig pg reload -D /data/pg18       # 指定数据目录
```

### pg status

查看所选实例的状态：

```bash
pig pg status
pig pg status -D /data/pg18
pig pg status -D /data/pg18 -o json
```

文本输出包含 `pg_ctl status`、`pg_controldata` 摘要、数据库操作系统用户的进程，以及相关服务状态。
`postgres` 行反映所选实例的进程状态。服务行包括 `patroni`、`pgbouncer`、`vip-manager` 和 `haproxy`。

结构化输出描述所选实例，包括数据目录、版本、PID、端口和可用的控制文件信息。

PID 对应的进程不存在，或独立探测确认 PID 已被其他可执行程序复用时，判定为陈旧 PID。
不能仅凭 `EPERM` 认定 PostgreSQL 已停止；无法确认进程身份时，状态保持未知，不能放行强制初始化或 PITR 恢复。
计划将无法确认的运行状态检查标记为 `unknown`，不提供可用 PID，也不将前置检查标记为通过。
`pg status` 对权限失败返回 `130201`（退出码 `3`）；对 PID 元数据无效、进程身份无法取得等
其他状态探测失败返回 `130808`（退出码 `1`）。这些结果的 `success` 为 `false`，其中的部分数据
不能证明 PostgreSQL 已停止。
陈旧 PID 不再让 PIG 把无关进程误认成 PostgreSQL；`pg_ctl` 与 PostgreSQL 仍执行自身的锁文件检查，
崩溃后可能仍需人工清理。其他目录复制来的 PID 文件不能证明所选实例正在运行。

### pg promote

时间线查询使用所选实例的端口与 socket。无法解析其 PID 元数据时，读取该目录的控制数据，
或将时间线保留为未知。这不改变 promote 的执行过程与主库角色成功判据。

将备库或暂停恢复的实例提升为可写主库。
使用 `--plan` 查看操作。文本模式要求确认，结构化执行必须加 `--yes`。

```bash
pig pg promote                    # 提升备库
pig pg promote -D /data/pg18      # 指定数据目录
pig pg promote --plan             # 预览提升计划
pig pg promote -y                 # 跳过确认提示
```

**选项：**

| 参数          | 简写   | 说明      |
|:------------|:-----|:--------|
| `--timeout` | `-t` | 等待超时（秒）；默认取 `$PGCTLTIMEOUT`，其次为 60 |
| `--no-wait` |      | 不等待提升完成 |
| `--plan`    |      | 仅预览提升计划 |
| `--yes`     | `-y` | 跳过确认提示   |
{.full-width}

### pg role

检测 PostgreSQL 实例的角色（主库或备库）。

```bash
pig pg role                       # 输出：primary、replica 或 unknown
pig pg role -V                    # 详细输出，显示检测过程
pig pg role -D /data/pg18         # 指定数据目录
```

**选项：**

| 参数          | 简写   | 说明       |
|:------------|:-----|:---------|
| `--verbose` | `-V` | 显示详细检测过程 |
{.full-width}

**输出说明：**

- `primary`：当前实例为主库
- `replica`：当前实例为备库
- `unknown`：无法确定实例角色

**检测策略（按优先级）：**

1. **目标进程检测**：读取所选数据目录的 `postmaster.pid`，只检查该 postmaster 的进程树
2. **绑定实例的 SQL 查询**：使用所选实例记录的端口与 socket，并在信任 `pg_is_in_recovery()` 前核对 `data_directory`
3. **数据目录检查**：检查 `standby.signal`、`recovery.signal`、`recovery.conf` 文件

文件只能提供角色配置，不能证明 PostgreSQL 正在运行。已停止的实例仍可报告配置的角色，
但残留 PID 文件不会使它被报告为存活。运行状态探测失败时，命令返回错误，不会将角色报告为已确认。

## 连接与查询命令

### pg psql

通过 psql 连接到 PostgreSQL 数据库。

```bash
pig pg psql                       # 连接到 postgres 数据库
pig pg psql mydb                  # 连接到指定数据库
pig pg psql mydb -c "SELECT 1"    # 执行单条命令
pig pg psql -f script.sql         # 执行 SQL 脚本文件
```

**选项：**

| 参数          | 简写   | 说明          |
|:------------|:-----|:------------|
| `--command` | `-c` | 执行单条 SQL 命令 |
| `--file`    | `-f` | 执行 SQL 脚本文件 |
{.full-width}

指定 `-D` 时，PIG 以所选操作系统用户读取该目录的 `postmaster.pid`。
确认进程存活后，才连接其中记录的端口和 Unix socket。实例停止、身份不明或元数据无效时返回错误。
省略 `-D` 时，psql 保留原生默认连接方式，以及客户端可见的 `PGHOST`/`PGPORT`。

使用 `-o json` 或 `-o yaml` 时，必须指定 `-c` 或 `-f`；交互式 psql 使用文本模式。

### pg ps

显示 PostgreSQL 当前连接。查询 `pg_stat_activity` 视图。

```bash
pig pg ps                         # 显示客户端连接
pig pg ps -a                      # 显示所有连接（包括系统进程）
pig pg ps -u dbuser_monitor       # 按用户筛选
pig pg ps -d meta                 # 按数据库筛选
```

**选项：**

| 参数           | 简写   | 说明             |
|:-------------|:-----|:---------------|
| `--all`      | `-a` | 显示所有连接（包括系统进程） |
| `--user`     | `-u` | 按用户筛选          |
| `--database` | `-d` | 按数据库筛选         |
{.full-width}

### pg kill

终止 PostgreSQL 连接。**默认为 dry-run 模式**，需要 `-x` 参数才会实际执行。

```bash
pig pg kill                       # 显示将被终止的连接（dry-run）
pig pg kill -x                    # 实际终止连接
pig pg kill --pid 12345 -x        # 终止指定 PID
pig pg kill -u admin -x           # 终止指定用户的连接
pig pg kill -d mydb -x            # 终止指定数据库的连接
pig pg kill -s idle -x            # 终止空闲连接
pig pg kill --cancel -x           # 取消查询而非终止连接
pig pg kill --watch 5 -x           # 每 5 秒重复执行
pig pg kill --plan                # 预览终止连接计划
```

**选项：**

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--execute` | `-x` | 实际执行（默认为 dry-run） |
| `--pid` | | 终止指定 PID |
| `--user` | `-u` | 按用户筛选 |
| `--database` | `-d` | 按数据库筛选 |
| `--state` | `-s` | 按状态筛选（idle/active/idle in transaction） |
| `--query` | `-q` | 按查询模式筛选 |
| `--all` | `-a` | 包括非客户端后台进程 |
| `--cancel` | `-c` | 取消查询而非终止连接 |
| `--watch` | | 每 N 秒重复执行 |
| `--plan` | | 预览执行计划，不终止连接 |
{.full-width}

`--state` 接受已知连接状态，包括 `idle in transaction`（含空格的值加引号）。
`--query` 不区分大小写匹配查询文本。正整数 `--pid` 直接选择该进程，并优先于其他筛选参数。
不加筛选条件时，选择其他客户端会话；`--all` 还包括非客户端后台进程。

### pg clone

在同一实例内创建数据库副本。复制前会终止源库会话，因此应安排在允许中断访问的时段执行。
命令使用 `CREATE DATABASE ... TEMPLATE`，PostgreSQL 15+ 加上 `STRATEGY FILE_COPY`；
也支持 PostgreSQL 14，使用其默认复制方式。

```bash
pig pg clone meta                       # 克隆 meta 为 meta_1/meta_2/...
pig pg clone meta meta_fork             # 克隆为指定数据库名
pig pg clone meta meta_fork --owner dba # 尝试修改新库 owner
pig pg clone meta meta_fork --port 5433 # 连接指定本地端口
pig pg clone meta meta_fork --plan      # 预览克隆计划
```

**选项：**

| 参数             | 简写   | 说明                                                      |
|:---------------|:-----|:--------------------------------------------------------|
| `--port`       |      | PostgreSQL 端口（默认 `5432` 或 `$PG_PORT`）                   |
| `--conn-db`    |      | 连接库默认为 `postgres`；克隆 `postgres` 时改用 `template1` |
| `--owner`      |      | 克隆后尝试修改新库 owner                                         |
| `--conn-limit` |      | 新库连接数限制（`-1` 无限制，`0` 禁止连接）                              |
| `--plan`       |      | 仅显示执行计划                                                 |
| `--yes`        | `-y` | 跳过确认提示                                                  |
{.full-width}

PostgreSQL 18+ 配置 `file_copy_method=clone` 且文件系统支持时，可使用写时复制（CoW）；
否则使用普通复制。要复制整个实例，使用 `pg fork`。
文件系统预检以所选数据库系统用户执行，支持调用用户无法访问的 `0700` 数据目录。
文本模式要求确认，结构化执行必须加 `--yes`。

## 数据库维护命令

`vacuum`、`analyze`、`freeze`、`repack` 的 `[dbname]` 参数接收数据库名。
包含 `=` 或以 `postgres://`、`postgresql://` 开头的值会被当作连接串，在生成计划、确认和执行前拒绝；
错误信息不回显该值。`vacuum`、`analyze`、`freeze` 使用 `--all` 时，会在任何维护操作前校验完整库名列表。
需要连接串时，请使用原生客户端。库名仍支持空格和中文等非 ASCII 字符。
库名以 `-` 开头时，将全部选项放在前面，再用 `--` 分隔库名；例如
`pig pg vacuum --full --plan -- --all` 只预览名为 `--all` 的单个数据库。

### pg vacuum

通过 psql 执行 SQL `VACUUM`，默认数据库为 `postgres`。

```bash
pig pg vacuum                     # 清理 postgres 数据库
pig pg vac                        # 别名
pig pg vacuum mydb                # 清理指定数据库
pig pg vacuum -a                  # 清理所有数据库
pig pg vacuum mydb -t mytable     # 清理指定表
pig pg vacuum mydb --schema myschema -t mytable  # 清理 schema 中的指定表
pig pg vacuum mydb --full         # VACUUM FULL（需要排他锁）
```

**选项：**

| 参数          | 简写   | 说明                 |
|:------------|:-----|:-------------------|
| `--all`     | `-a` | 处理所有数据库            |
| `--schema`  |      | 指定 schema          |
| `--table`   | `-t` | 指定表名               |
| `--verbose` | `-V` | 详细输出               |
| `--full`    | `-F` | VACUUM FULL（需要排他锁） |
| `--plan` | | 查看操作，不执行 |
| `--yes` | `-y` | 跳过 VACUUM FULL 确认 |
{.full-width}

`vacuum` 和 `freeze` 使用 `--schema` 与 `--table` 选择一个表。
处理整个 schema 时，使用原生 `vacuumdb --schema SCHEMA`。

`--full` 重写表并获取排他锁。文本模式要求确认，结构化执行必须加 `--yes`。

### pg analyze

通过 psql 执行 SQL `ANALYZE`，更新查询规划所用的统计信息。默认数据库为 `postgres`。

```bash
pig pg analyze                    # 分析 postgres 数据库
pig pg ana                        # 别名
pig pg analyze mydb               # 分析指定数据库
pig pg analyze -a                 # 分析所有数据库
pig pg analyze mydb -t mytable    # 分析指定表
```

**选项：**

| 参数          | 简写   | 说明        |
|:------------|:-----|:----------|
| `--all`     | `-a` | 处理所有数据库   |
| `--schema`  |      | 指定 schema |
| `--table`   | `-t` | 指定表名      |
| `--verbose` | `-V` | 详细输出      |
{.full-width}

### pg freeze

通过 psql 执行 SQL `VACUUM (FREEZE)`，冻结旧事务 ID。默认数据库为 `postgres`。

```bash
pig pg freeze                     # 冻结清理 postgres 数据库
pig pg freeze mydb                # 冻结清理指定数据库
pig pg freeze -a                  # 冻结清理所有数据库
pig pg freeze mydb -t mytable     # 冻结清理指定表
```

**选项：** 与 `pg analyze` 相同。

### pg repack

使用 `pg_repack` 重建表。先安装可执行文件，并在目标数据库中安装扩展。
默认数据库为 `postgres`。

```bash
pig pg repack mydb                # 重整数据库中所有表
pig pg rp mydb                    # 别名
pig pg repack -a                  # 重整所有数据库
pig pg repack mydb -t mytable     # 重整指定表
pig pg repack mydb --schema myschema  # 重整指定 schema 中的表
pig pg repack mydb -j 4           # 使用 4 个并行任务
pig pg repack mydb --plan         # 预览目标与命令，不执行 pg_repack
```

**选项：**

| 参数          | 简写   | 说明          |
|:------------|:-----|:------------|
| `--all`     | `-a` | 处理所有数据库     |
| `--schema`  |      | 指定 schema   |
| `--table`   | `-t` | 指定表名        |
| `--verbose` | `-V` | 详细输出        |
| `--jobs`    | `-j` | 并行任务数（默认 1） |
| `--plan`    |      | 显示将被重整的表    |
{.full-width}

## 参数调优命令

### pg tune

根据 PostgreSQL 版本、CPU 核数、内存、磁盘容量和所选负载类型，生成内存、并行度和 WAL 参数的初始建议。
存储计算按 SSD 假设。命令只输出建议，不修改配置文件。

```bash
pig pg tune                       # 自动探测硬件，使用 oltp 画像
pig pg tuning                     # 别名
pig pg tune -p olap               # 使用 OLAP 画像
pig pg tune -p tiny               # 小规格实例
pig pg tune -c 8 -m 32768 -d 500  # 覆盖自动探测结果
pig pg tune -C 500                # 覆盖 max_connections
pig pg tune -R 0.30               # 调整 shared_buffers 比例
pig pg tune -o json               # 结构化输出 JSON
pig pg tune -o yaml               # 结构化输出 YAML
```

**选项：**

| 参数              | 简写   | 默认值  | 说明                                      |
|:----------------|:-----|:-----|:----------------------------------------|
| `--profile`     | `-p` | oltp | 调优画像：`oltp` / `olap` / `tiny` / `crit`  |
| `--cpu`         | `-c` | 0    | CPU 核数，0 表示自动探测                         |
| `--mem`         | `-m` | 0    | 内存大小（MB），0 表示自动探测                       |
| `--disk`        | `-d` | 0    | 数据盘容量（GB），0 表示自动探测                      |
| `--max-conn`    | `-C` | 0    | 覆盖 `max_connections`，0 表示使用画像默认值        |
| `--shmem-ratio` | `-R` | 0.25 | `shared_buffers` 占内存比例，取值范围 `0.1 ~ 0.4` |
{.full-width}

**画像说明：**

| 画像     | 适用场景     | 特点                 |
|:-------|:---------|:-------------------|
| `oltp` | 通用在线事务处理 | 平衡连接数、缓存与并行度       |
| `olap` | 分析型负载    | 更激进地使用并行与工作内存      |
| `tiny` | 小规格实例    | 控制内存占用与并行度         |
| `crit` | 延迟敏感场景   | 限制并行 gather，偏向稳态响应 |
{.full-width}

**说明：**

- 生成结果会随当前 PostgreSQL 主版本自动裁剪，例如 `io_workers` 仅会在 PG 18+ 输出。
- 文本输出包含成功消息，只将参数赋值行复制到配置文件，不要将完整输出直接重定向为 PostgreSQL 配置。
- 脚本使用 `-o json` 并读取 `data.parameters`。检查参数后，通过现有配置流程应用，按参数要求重载或重启。

## 实例 Fork

### pg fork

创建一个本地一次性 PostgreSQL 物理副本，适合临时分析、排障、恢复验证和开发测试。托管 fork 默认写入 `/pg/data-<name>`，不会注册到 Pigsty、systemd 或 Patroni；显式指定 `--dst-data` 时会创建非托管 fork，不会被 `fork list` 枚举。

```bash
pig pg fork dev                       # 创建 /pg/data-dev，不启动
pig pg fork init dev --start          # 创建后启动，端口从 15432 开始自动探测
pig pg fork init dev -s --dst-port 15433    # 创建后在指定端口启动
pig pg fork init dev -D /pg/data2 --src-port 15431  # 指定源目录与源端口
pig pg fork init dev --dst-data /tmp/dev    # 创建非托管 fork
pig pg fork list                      # 列出托管 fork
pig pg fork start dev                 # 启动已有托管 fork
pig pg fork stop dev                  # 停止已有托管 fork
pig pg fork rm dev --stop             # 停止并删除运行中的 fork
pig pg fork init dev --plan           # 仅显示执行计划
```

**创建选项：**

| 参数           | 简写   | 默认值                     | 说明                            |
|:-------------|:-----|:------------------------|:------------------------------|
| `--dst-data` |      | `/pg/data-<name>`       | 非托管目标数据目录                     |
| `--dst-port` |      | 自动探测                    | 目标端口，从 15432 起寻找空闲端口          |
| `--src-data` |      | `/pg/data` 或 `$PG_DATA` | 源数据目录；也可用全局 `pg -D/--data` 设置 |
| `--src-port` |      | `5432` 或 `$PG_PORT`     | 源端口                           |
| `--start`    | `-s` | false                   | 创建后启动 fork                    |
| `--force`    | `-f` | false                   | 覆盖已有且已停止的目标目录，并跳过确认           |
| `--timeout`  | `-t` | 60                      | 启动等待超时（秒）；默认取 `$PGCTLTIMEOUT`，其次为 60                     |
| `--yes`      | `-y` | false                   | 跳过确认提示                        |
| `--plan`     |      | false                   | 只显示执行计划，不执行                   |
{.full-width}

**管理子命令：**

| 命令                                     | 常用参数                                                                                       | 说明                           |
|:---------------------------------------|:-------------------------------------------------------------------------------------------|:-----------------------------|
| `pig pg fork list`                     | `-o json/yaml`                                                                   | 列出托管 fork                    |
| `pig pg fork start <name> or --dst-data <dir>` | `--dst-data`, `--dst-port`, `-t/--timeout`, `--plan`                                | 启动已有 fork                    |
| `pig pg fork stop <name> or --dst-data <dir>`  | `--dst-data`, `-m/--mode`, `-t/--timeout`, `--plan`                                   | 停止已有 fork                    |
| `pig pg fork rm <name> or --dst-data <dir>`    | `--dst-data`, `--stop`, `-m/--mode`, `-t/--timeout`, `-f/--force`, `-y/--yes`, `--plan` | 删除 fork；运行中的 fork 需 `--stop` |
{.full-width}

**行为说明：**

- 源实例运行时，命令会使用 PostgreSQL 低级备份 API 创建一致的物理副本；源实例停止时，可执行冷复制。
- 命令会优先使用 CoW/reflink；如果只能普通复制，会在交互模式中提示空间风险并等待确认。
- 目标目录不能是 `/`、`/pg`、源 PGDATA，或源目录的父目录、子目录。先解析符号链接，再检查源、目标关系。`/pg/data` 与其他数据目录遵循相同的归属、路径关系和运行状态检查。
- 覆盖非空的非托管目标时，必须有有效的非托管 `pg_fork` 元数据，且 `fork.json` 明确记录的 `target.data` 与目标目录匹配。缺少 `PG_VERSION` 的普通目录不能因此被删除。元数据缺失、损坏、不完整或不匹配时拒绝覆盖；空目录仍可配合 `--force` 使用。托管 fork 必须按名称替换。运行中或状态不明的目标不能覆盖。按名称强制删除仍允许清理 `/pg/data-<name>` 下的受管孤儿目录。
- 调用用户无法遍历私有目录时，路径校验以 DBSU 解析；无法解析的路径以及源、目标目录别名仍被拒绝。
- 删除非托管 fork 时，对调用用户无权检查的目标，以所选 DBSU 检查其是否为真实目录；符号链接或检查失败仍会阻止删除。
- 复制完成后会清理 fork 中的运行态与复制状态，并写入 `fork.json`。只有指定 `-s|--start` 时才会启动新实例。
- 托管 fork 必须通过名称管理；非托管 fork 需要通过 `--dst-data` 指定目录来启动、停止或删除。

**列出 fork：**

`pig pg fork list` 扫描 `/pg/data-*`，读取 `fork.json` 并刷新运行状态。
文本列为 `NAME`、`PORT`、`STATE`、`PID`、`AGE`、`SOURCE`、`COPY` 和 `DATA`。
状态为 `running`、`stopped`、`orphan` 或 `unknown`。状态不明时，结构化结果设置
`runtime_unknown: true`，不会声称已启动；此时生命周期操作失败且不重写元数据。

JSON 和 YAML 在结果的 `data` 数组中返回 fork 记录，包括源与目标连接信息、复制与备份元数据、
管理命令和 PIG 构建信息。

**结构化输出：**

```bash
pig pg fork init dev --plan -o yaml
pig pg fork list -o json
```

## 日志命令

日志命令用于查看 PostgreSQL 日志文件。默认日志目录为 `/pg/log/postgres`，可通过 `--log-dir` 参数指定其他目录。默认 `pg log` 动作显示最新 CSV 日志快照；使用 `pg log -f` 或 `pg log tail` 实时跟踪。只有 `pg log` 与 `pg log show` 支持 `-o json` 将 CSV 日志行转换为 JSONL；日志快照不支持 `yaml` 与 `json-pretty`，follow/tail、`less`、`grep` 不支持结构化输出。

**日志命令全局参数：**

| 参数                | 说明                            |
|:------------------|:------------------------------|
| `--log-dir`       | 日志目录路径（默认：`/pg/log/postgres`） |
| `--lines` / `-n`  | 显示行数（默认 50）                   |
| `--follow` / `-f` | 跟踪最新日志（仅 `pg log` 父命令）        |
{.full-width}

**权限处理：** 如果当前用户没有权限读取日志目录，命令会自动使用 `sudo` 重试。

### pg log

显示最新日志快照；配合 `-f` 时跟踪最新日志。

```bash
pig pg log                        # 显示最新 50 行
pig pg log -n 100                 # 显示最新 100 行
pig pg log -f                     # 跟踪最新日志
```

### pg log list

列出日志目录中的日志文件。

```bash
pig pg log list                              # 列出默认目录中的日志
pig pg log ls                                # 别名
pig pg log list --log-dir /var/log/postgres  # 列出指定目录中的日志
```

### pg log tail

实时查看日志文件（类似 `tail -f`）。默认查看最新的 CSV 日志文件。

```bash
pig pg log tail                   # 查看最新日志
pig pg log t                      # 别名
pig pg log f                      # 别名
pig pg log tail postgresql.csv    # 查看指定日志文件
pig pg log tail -n 100            # 显示最后 100 行后开始跟踪
pig pg log tail --log-dir /var/log/postgres  # 使用自定义目录
```

**选项：**

| 参数         | 简写   | 默认值   | 说明                      |
|:-----------|:-----|:------|:------------------------|
| `--lines`  | `-n` | 50    | 显示的行数                   |
| `--follow` | `-f` | false | 兼容旧用法；`tail` 始终跟踪 |
{.full-width}

### pg log show

输出日志文件内容。

```bash
pig pg log show                   # 输出最新日志
pig pg log cat                    # show 的别名
pig pg log c                      # show 的别名
pig pg log show -n 100            # 输出最后 100 行
pig pg log show postgresql.csv    # 输出指定日志文件
```

**选项：**

| 参数        | 简写   | 默认值 | 说明    |
|:----------|:-----|:----|:------|
| `--lines` | `-n` | 50  | 显示的行数 |
{.full-width}

### pg log less

用 less 打开日志文件。默认定位到文件末尾（`+G`）。

```bash
pig pg log less                   # 用 less 打开最新日志
pig pg log vi                     # 别名
pig pg log v                      # 别名
pig pg log less postgresql.csv    # 打开指定日志文件
```

### pg log grep

搜索日志文件内容。

```bash
pig pg log grep ERROR             # 搜索 ERROR
pig pg log grep --ignore-case error  # 忽略大小写
pig pg log grep -C 3 ERROR        # 显示上下文
pig pg log grep ERROR pg.csv      # 搜索指定日志文件
```

**选项：**

| 参数              | 简写   | 说明      |
|:----------------|:-----|:--------|
| `--ignore-case` |      | 忽略大小写   |
| `--context`     | `-C` | 显示上下文行数 |
{.full-width}

## pg svc 子命令

`pg svc`（也可写作 `pg service` 或 `pg s`）控制 Pigsty 使用的 `postgres` systemd 服务，要求该单元已存在。
对于发行版提供的 `postgresql-18` 或 `postgresql@18-main` 等单元，直接使用实际单元名调用 `systemctl`。

```bash
pig pg svc start                 # 启动 postgres 服务
pig pg svc stop                  # 停止 postgres 服务
pig pg svc restart               # 重启 postgres 服务
pig pg svc reload                # 重载 postgres 服务
pig pg svc status                # 显示服务状态
```

**别名对照：**

| 命令               | 别名               |
|:-----------------|:-----------------|
| `pg svc start`   | `boot, up`       |
| `pg svc stop`    | `halt, dn, down` |
| `pg svc restart` | `reboot, rt`     |
| `pg svc reload`  | `rl, hup`        |
| `pg svc status`  | `st, stat`       |
{.full-width}

## 工具与权限

PIG 使用 `initdb` 初始化数据，用 `pg_ctl` 控制实例，通过 psql 执行 SQL 和表维护，
用 `pg_repack` 重建表，通过系统工具读取日志。需要本命令未提供的选项时，直接使用相应工具。

当前用户是所选 PostgreSQL 操作系统用户时直接执行。
root 使用 `su`，其他用户使用 `sudo -H -u`。服务管理使用 root 权限。
这些命令面向 Linux，服务管理需要 systemd。

## 设计记录

- [七成正确的调优器：界定 pig pg tune 的边界](https://pig.pgsty.com/zh/design/pg-tune/)
- [PIG 运维 CLI 安全契约](https://pig.pgsty.com/zh/design/ops-cli-safety/)
