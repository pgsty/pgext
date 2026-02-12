---
title: "pig pg"
description: "使用 pig pg 子命令管理本地 PostgreSQL 服务器"
icon: Database
weight: 640
---


`pig pg` 命令（别名 `pig postgres`）用于管理本地 PostgreSQL 服务器和数据库。它封装了 `pg_ctl`、`psql`、`vacuumdb` 等原生工具，提供简化的服务器管理体验。

```bash
pig pg - Manage local PostgreSQL server and databases.

Server Control (via pg_ctl):
  pig pg init     [-v ver] [-D datadir]     initialize data directory
  pig pg start    [-D datadir]              start PostgreSQL server
  pig pg stop     [-D datadir] [-m fast]    stop PostgreSQL server
  pig pg restart  [-D datadir] [-m fast]    restart PostgreSQL server
  pig pg reload   [-D datadir]              reload configuration
  pig pg status   [-D datadir]              show server status
  pig pg promote  [-D datadir]              promote standby to primary
  pig pg role     [-D datadir] [-V]         detect instance role (primary/replica)

Service Management (via systemctl):
  pig pg svc start                          start postgres systemd service
  pig pg svc stop                           stop postgres systemd service
  pig pg svc restart                        restart postgres systemd service
  pig pg svc reload                         reload postgres systemd service
  pig pg svc status                         show postgres service status

Connection & Query:
  pig pg psql     [db] [-c cmd]             connect to database via psql
  pig pg ps       [-a] [-u user]            show current connections
  pig pg kill     [-x] [-u user]            terminate connections (dry-run by default)

Database Maintenance:
  pig pg vacuum   [db] [-a] [-t table]      vacuum tables
  pig pg analyze  [db] [-a] [-t table]      analyze tables
  pig pg freeze   [db] [-a] [-t table]      vacuum freeze tables
  pig pg repack   [db] [-a] [-t table]      repack tables (online rebuild)

Utilities:
  pig pg log <list|tail|cat|less>           view PostgreSQL logs
```

## 命令概览

**服务控制**（pg_ctl 封装）：

| 命令 | 别名 | 描述 | 备注 |
|:----|:----|:----|:----|
| `pg init` | `initdb, i` | 初始化数据目录 | 封装 initdb |
| `pg start` | `boot, up` | 启动 PostgreSQL | 封装 pg_ctl start |
| `pg stop` | `halt, down` | 停止 PostgreSQL | 封装 pg_ctl stop |
| `pg restart` | `reboot` | 重启 PostgreSQL | 封装 pg_ctl restart |
| `pg reload` | `hup` | 重载配置 | 封装 pg_ctl reload |
| `pg status` | `st, stat` | 查看服务状态 | 显示进程与相关服务状态 |
| `pg promote` | `pro` | 提升备库为主库 | 封装 pg_ctl promote |
| `pg role` | `r` | 检测实例角色 | 输出 primary/replica |
{.full-width}

**连接与查询**：

| 命令 | 别名 | 描述 | 备注 |
|:----|:----|:----|:----|
| `pg psql` | `sql, connect` | 连接到数据库 | 封装 psql |
| `pg ps` | `activity, act` | 显示当前连接 | 查询 pg_stat_activity |
| `pg kill` | `k` | 终止连接 | 默认 dry-run 模式 |
{.full-width}

**数据库维护**：

| 命令 | 别名 | 描述 | 备注 |
|:----|:----|:----|:----|
| `pg vacuum` | `vac, vc` | 清理表 | 封装 vacuumdb |
| `pg analyze` | `ana, az` | 分析表 | 封装 vacuumdb --analyze-only |
| `pg freeze` | `frz` | 冻结清理表 | 封装 vacuumdb --freeze |
| `pg repack` | `rp` | 在线重整表 | 需要 pg_repack 扩展 |
{.full-width}

**日志工具**：

| 命令 | 别名 | 描述 | 备注 |
|:----|:----|:----|:----|
| `pg log` | `l` | 日志管理 | 父命令 |
| `pg log list` | `ls` | 列出日志文件 | |
| `pg log tail` | `t, f` | 实时查看日志 | tail -f |
| `pg log cat` | `c` | 输出日志内容 | |
| `pg log less` | `vi, v` | 用 less 查看 | |
{.full-width}

> `v1.0.0` 已知问题：`pig pg log grep` 存在参数冲突导致不可用。可使用 `pig pg log cat | grep PATTERN` 作为替代。

**服务子命令**（`pg svc`）：

| 命令 | 别名 | 描述 |
|:----|:----|:----|
| `pg svc start` | `boot, up` | 启动 postgres 服务 |
| `pg svc stop` | `halt, dn, down` | 停止 postgres 服务 |
| `pg svc restart` | `reboot, rt` | 重启 postgres 服务 |
| `pg svc reload` | `rl, hup` | 重载 postgres 服务 |
| `pg svc status` | `st, stat` | 显示服务状态 |
{.full-width}


## 快速入门

```bash
# 服务控制
pig pg init                       # 初始化数据目录
pig pg start                      # 启动 PostgreSQL
pig pg status                     # 查看状态
pig pg stop                       # 停止 PostgreSQL
pig pg restart                    # 重启 PostgreSQL
pig pg reload                     # 重载配置

# 连接与查询
pig pg psql                       # 连接到 postgres 数据库
pig pg psql mydb                  # 连接到指定数据库
pig pg ps                         # 查看当前连接
pig pg kill -x                    # 终止连接（需要 -x 确认执行）

# 数据库维护
pig pg vacuum mydb                # 清理指定数据库
pig pg analyze mydb               # 分析指定数据库
pig pg repack mydb                # 在线重整数据库

# 日志查看
pig pg log tail                   # 实时查看最新日志
pig pg log list --log-dir /var/log/pg  # 使用自定义日志目录
pig pg log cat | grep ERROR       # 在 shell 中过滤日志
```


## 全局参数

以下参数适用于所有 `pig pg` 子命令：

| 参数 | 简写 | 默认值 | 说明 |
|:----|:----|:------|:----|
| `--version` | `-v` | 自动检测 | PostgreSQL 主版本号 |
| `--data` | `-D` | `/pg/data` | 数据目录路径 |
| `--dbsu` | `-U` | `postgres` | 数据库超级用户（或 `$PIG_DBSU` 环境变量） |
{.full-width}

**版本检测逻辑：**

1. 如果指定了 `-v`，使用指定版本
2. 否则从数据目录的 `PG_VERSION` 文件读取版本
3. 如果都无法获取，使用 PATH 中的默认 PostgreSQL


## 服务控制命令

### pg init

初始化 PostgreSQL 数据目录。封装 `initdb` 命令。

```bash
pig pg init                       # 使用默认设置初始化
pig pg init -v 17                 # 指定 PostgreSQL 17
pig pg init -D /data/pg17         # 指定数据目录
pig pg init -k                    # 启用数据校验和
pig pg init -f                    # 强制初始化（删除已有数据）
pig pg init -- --waldir=/wal      # 传递额外参数给 initdb
```

**选项：**

| 参数 | 简写 | 默认值 | 说明 |
|:----|:----|:------|:----|
| `--encoding` | `-E` | UTF8 | 数据库编码 |
| `--locale` | | C | 区域设置 |
| `--data-checksum` | `-k` | false | 启用数据校验和 |
| `--force` | `-f` | false | 强制初始化，删除已有数据（危险！） |
{.full-width}

**安全机制：** 即使使用 `--force`，如果 PostgreSQL 正在运行，命令也会拒绝执行，以防止数据丢失。


### pg start

启动 PostgreSQL 服务器。

```bash
pig pg start                      # 使用默认设置启动
pig pg up                         # 别名
pig pg boot                       # 别名
pig pg start -D /data/pg17        # 指定数据目录
pig pg start -l /pg/log/pg.log    # 重定向输出到日志文件
pig pg start -o "-p 5433"         # 传递参数给 postgres
pig pg start -y                   # 强制启动（跳过运行检查）
```

**选项：**

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--log` | `-l` | 重定向 stdout/stderr 到日志文件 |
| `--timeout` | `-t` | 等待超时（秒） |
| `--no-wait` | `-W` | 不等待启动完成 |
| `--options` | `-o` | 传递给 postgres 的选项 |
| `--yes` | `-y` | 强制启动（即使已运行） |
{.full-width}


### pg stop

停止 PostgreSQL 服务器。

```bash
pig pg stop                       # 快速停止（默认）
pig pg down                       # 别名
pig pg halt                       # 别名
pig pg stop -m smart              # 等待客户端断开
pig pg stop -m immediate          # 立即关闭
```

**选项：**

| 参数 | 简写 | 默认值 | 说明 |
|:----|:----|:------|:----|
| `--mode` | `-m` | fast | 关闭模式：smart/fast/immediate |
| `--timeout` | `-t` | 60 | 等待超时（秒） |
| `--no-wait` | `-W` | false | 不等待关闭完成 |
{.full-width}

**关闭模式说明：**

| 模式 | 说明 |
|:----|:----|
| `smart` | 等待所有客户端断开后关闭 |
| `fast` | 回滚活动事务，断开客户端，正常关闭 |
| `immediate` | 立即终止所有进程，下次启动需要恢复 |
{.full-width}


### pg restart

重启 PostgreSQL 服务器。

```bash
pig pg restart                    # 快速重启
pig pg reboot                     # 别名
pig pg restart -m immediate       # 立即重启
pig pg restart -o "-p 5433"       # 使用新选项重启
```

**选项：** 与 `pg stop` 相同，另外支持 `--options`（`-o`）传递给 postgres。


### pg reload

重载 PostgreSQL 配置。向服务器发送 SIGHUP 信号。

```bash
pig pg reload                     # 重载配置
pig pg hup                        # 别名
pig pg reload -D /data/pg17       # 指定数据目录
```


### pg status

显示 PostgreSQL 服务器状态。此命令不仅显示 `pg_ctl status` 的结果，还会显示 postgres 相关进程和 Pigsty 相关服务的状态。

```bash
pig pg status                     # 查看服务状态
pig pg st                         # 别名
pig pg status -D /data/pg17       # 指定数据目录
```

**输出内容：**

1. `pg_ctl status` 输出（进程是否运行、PID 等）
2. PostgreSQL 进程列表（`ps -u postgres`）
3. 相关服务状态：
   - `postgres`：PostgreSQL systemd 服务
   - `patroni`：Patroni HA 管理服务
   - `pgbouncer`：连接池服务
   - `pgbackrest`：备份服务
   - `vip-manager`：VIP 管理服务
   - `haproxy`：负载均衡服务


### pg promote

将备库提升为主库。

```bash
pig pg promote                    # 提升备库
pig pg pro                        # 别名
pig pg promote -D /data/pg17      # 指定数据目录
```

**选项：**

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--timeout` | `-t` | 等待超时（秒） |
| `--no-wait` | `-W` | 不等待提升完成 |
{.full-width}


### pg role

检测 PostgreSQL 实例的角色（主库或备库）。

```bash
pig pg role                       # 输出：primary、replica 或 unknown
pig pg role -V                    # 详细输出，显示检测过程
pig pg role -D /data/pg17         # 指定数据目录
```

**选项：**

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--verbose` | `-V` | 显示详细检测过程 |
{.full-width}

**输出说明：**

- `primary`：当前实例为主库
- `replica`：当前实例为备库
- `unknown`：无法确定实例角色

**检测策略（按优先级）：**

1. **进程检测**：检查 `walreceiver`、`recovery` 等进程
2. **SQL 查询**：执行 `pg_is_in_recovery()` 查询（需要 PostgreSQL 运行）
3. **数据目录检查**：检查 `standby.signal`、`recovery.signal`、`recovery.conf` 文件


## 连接与查询命令

### pg psql

通过 psql 连接到 PostgreSQL 数据库。

```bash
pig pg psql                       # 连接到 postgres 数据库
pig pg sql                        # 别名
pig pg connect                    # 别名
pig pg psql mydb                  # 连接到指定数据库
pig pg psql mydb -c "SELECT 1"    # 执行单条命令
pig pg psql -f script.sql         # 执行 SQL 脚本文件
```

**选项：**

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--command` | `-c` | 执行单条 SQL 命令 |
| `--file` | `-f` | 执行 SQL 脚本文件 |
{.full-width}


### pg ps

显示 PostgreSQL 当前连接。查询 `pg_stat_activity` 视图。

```bash
pig pg ps                         # 显示客户端连接
pig pg activity                   # 别名
pig pg act                        # 别名
pig pg ps -a                      # 显示所有连接（包括系统进程）
pig pg ps -u admin                # 按用户筛选
pig pg ps -d mydb                 # 按数据库筛选
```

**选项：**

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--all` | `-a` | 显示所有连接（包括系统进程） |
| `--user` | `-u` | 按用户筛选 |
| `--database` | `-d` | 按数据库筛选 |
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
pig pg kill -w 5 -x               # 每 5 秒重复执行
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
| `--all` | `-a` | 包括复制连接 |
| `--cancel` | `-c` | 取消查询而非终止连接 |
| `--watch` | `-w` | 每 N 秒重复执行 |
{.full-width}

**安全说明：** `--state` 和 `--query` 参数会进行标识符验证，只接受简单的字母数字模式，以防止 SQL 注入。


## 数据库维护命令

### pg vacuum

清理数据库表。封装 `vacuumdb` 命令。

```bash
pig pg vacuum                     # 清理当前数据库
pig pg vac                        # 别名
pig pg vacuum mydb                # 清理指定数据库
pig pg vacuum -a                  # 清理所有数据库
pig pg vacuum mydb -t mytable     # 清理指定表
pig pg vacuum mydb -n myschema    # 清理指定 schema 中的表
pig pg vacuum mydb --full         # VACUUM FULL（需要排他锁）
```

**选项：**

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--all` | `-a` | 处理所有数据库 |
| `--schema` | `-n` | 指定 schema |
| `--table` | `-t` | 指定表名 |
| `--verbose` | `-V` | 详细输出 |
| `--full` | `-F` | VACUUM FULL（需要排他锁） |
{.full-width}

**安全说明：** `--schema` 和 `--table` 参数会进行标识符验证，只接受有效的 PostgreSQL 标识符格式。


### pg analyze

分析数据库表以更新统计信息。

```bash
pig pg analyze                    # 分析当前数据库
pig pg ana                        # 别名
pig pg analyze mydb               # 分析指定数据库
pig pg analyze -a                 # 分析所有数据库
pig pg analyze mydb -t mytable    # 分析指定表
```

**选项：** 与 `pg vacuum` 相同（不含 `--full`）。


### pg freeze

对数据库表执行冻结清理（vacuum freeze），防止事务 ID 回卷。

```bash
pig pg freeze                     # 冻结清理当前数据库
pig pg freeze mydb                # 冻结清理指定数据库
pig pg freeze -a                  # 冻结清理所有数据库
pig pg freeze mydb -t mytable     # 冻结清理指定表
```

**选项：** 与 `pg vacuum` 相同（不含 `--full`）。


### pg repack

在线重整数据库表。需要安装 `pg_repack` 扩展。

```bash
pig pg repack mydb                # 重整数据库中所有表
pig pg rp mydb                    # 别名
pig pg repack -a                  # 重整所有数据库
pig pg repack mydb -t mytable     # 重整指定表
pig pg repack mydb -n myschema    # 重整指定 schema 中的表
pig pg repack mydb -j 4           # 使用 4 个并行任务
pig pg repack mydb --dry-run      # 显示将被重整的表
```

**选项：**

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--all` | `-a` | 处理所有数据库 |
| `--schema` | `-n` | 指定 schema |
| `--table` | `-t` | 指定表名 |
| `--verbose` | `-V` | 详细输出 |
| `--jobs` | `-j` | 并行任务数（默认 1） |
| `--dry-run` | `-N` | 显示将被重整的表 |
{.full-width}


## 日志命令

日志命令用于查看 PostgreSQL 日志文件。默认日志目录为 `/pg/log/postgres`，可通过 `--log-dir` 参数指定其他目录。

**日志命令全局参数：**

| 参数 | 说明 |
|:----|:----|
| `--log-dir` | 日志目录路径（默认：`/pg/log/postgres`） |
{.full-width}

**权限处理：** 如果当前用户没有权限读取日志目录，命令会自动使用 `sudo` 重试。


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

| 参数 | 简写 | 默认值 | 说明 |
|:----|:----|:------|:----|
| `--lines` | `-n` | 50 | 显示的行数 |
{.full-width}


### pg log cat

输出日志文件内容。

```bash
pig pg log cat                    # 输出最新日志
pig pg log c                      # 别名
pig pg log cat -n 100             # 输出最后 100 行
pig pg log cat postgresql.csv     # 输出指定日志文件
```

**选项：**

| 参数 | 简写 | 默认值 | 说明 |
|:----|:----|:------|:----|
| `--lines` | `-n` | 100 | 显示的行数 |
{.full-width}


### pg log less

用 less 打开日志文件。默认定位到文件末尾（`+G`）。

```bash
pig pg log less                   # 用 less 打开最新日志
pig pg log vi                     # 别名
pig pg log v                      # 别名
pig pg log less postgresql.csv    # 打开指定日志文件
```


## pg svc 子命令

`pg svc` 提供通过 systemctl 管理 PostgreSQL 服务的功能：

```bash
pig pg svc start                 # 启动 postgres 服务
pig pg svc stop                  # 停止 postgres 服务
pig pg svc restart               # 重启 postgres 服务
pig pg svc reload                # 重载 postgres 服务
pig pg svc status                # 显示服务状态
```

**别名对照：**

| 命令 | 别名 |
|:----|:----|
| `pg svc start` | `boot, up` |
| `pg svc stop` | `halt, dn, down` |
| `pg svc restart` | `reboot, rt` |
| `pg svc reload` | `rl, hup` |
| `pg svc status` | `st, stat` |
{.full-width}


## 设计说明

**与原生工具的关系：**

`pig pg` 并非对 PostgreSQL 原生工具的简单封装，而是针对常用操作的上层抽象：

- 服务控制命令（init/start/stop/restart/reload/promote）调用 `pg_ctl`
- `status` 命令除了 `pg_ctl status` 外，还显示进程和相关服务状态
- 连接管理命令（psql/ps/kill）调用 `psql`
- 维护命令（vacuum/analyze）调用 `vacuumdb`
- repack 命令调用 `pg_repack`
- 日志命令调用 `tail`、`less`、`grep` 等系统工具
- `pg svc` 命令调用 `systemctl`

如需使用原生工具的完整功能，可直接调用相应命令。

**权限处理：**

- 如果当前用户已是 DBSU：直接执行命令
- 如果当前用户是 root：使用 `su - postgres -c "..."` 执行
- 其他用户：使用 `sudo -inu postgres -- ...` 执行

**安全性考虑：**

- `--state`、`--query`、`--schema`、`--table` 等参数都经过标识符验证，防止 SQL 注入
- `pg kill` 默认为 dry-run 模式，避免误操作
- 日志命令在权限不足时自动使用 sudo

**平台支持：**

此命令专为 Linux 系统设计，部分功能依赖 `systemctl`。
