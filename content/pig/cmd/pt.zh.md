---
title: "pig pt"
description: "使用 pig pt 子命令管理 Patroni 服务与集群"
icon: RefreshCw
weight: 650
---

`pig patroni` 命令（别名 `pig pt`）用于管理 Patroni 服务和 PostgreSQL HA 集群。它封装了常用的 `patronictl` 和 `systemctl` 操作，提供简化的集群管理体验。

```bash
pig pt - Manage Patroni cluster using patronictl commands.

Cluster Operations (via patronictl):
  pig pt list [cluster]            list cluster members
  pig pt restart [member]          restart PostgreSQL (rolling restart)
  pig pt reload                    reload PostgreSQL config
  pig pt reinit <member>           reinitialize a member
  pig pt pause                     pause automatic failover
  pig pt resume                    resume automatic failover
  pig pt switchover                perform planned switchover
  pig pt failover [candidate]      perform manual failover
  pig pt config <action>           manage cluster config (edit|show|set|pg)

Service Management (via systemctl):
  pig pt status                    show comprehensive patroni status
  pig pt svc start (pig pt start)  start patroni service
  pig pt svc stop  (pig pt stop)   stop patroni service
  pig pt svc restart               restart patroni service
  pig pt svc reload                reload patroni service
  pig pt svc status                show patroni service status

Logs:
  pig pt log [-f] [-n 50]          view patroni logs
  pig pt log tail [-n 50]          follow patroni logs
  pig pt log show [-n 50]          show patroni log snapshot
  pig pt log grep <pattern>        search patroni logs
```

`pt start` / `pt stop` 是 `pt svc start` / `pt svc stop` 的隐藏快捷入口；`pt svc` 是显式的 Patroni 守护进程管理入口。

短参数契约：如果命令作用域内没有冲突，`--wait` 必须提供 `-w` 缩写。当前适用于 `pt reinit`、`pt pause`、`pt resume`；`pt list -w` 保留为该命令自己的刷新间隔参数。

切换前置检查契约：`pt switchover` 与 `pt failover` 在执行或要求确认前，会通过与 `pig pt list` / `pig pt config show` 相同的结构化接口读取当前拓扑，确认集群名、当前 Leader、候选从库以及 Patroni pause 状态。如果集群处于 pause 模式，pig 会拒绝切换并提示先执行 `pig pt resume`。


## 命令概览

**集群操作**（`patronictl` 封装）：

以下命令用于通过 Patroni 管理 PostgreSQL 集群。

| 命令              | 缩写   | 描述               | 实现方式                                   |
|:----------------|:-----|:-----------------|:---------------------------------------|
| `pt list`       | `ls` | 列出集群成员           | `patronictl list -e -t`                |
| `pt restart`    | `rs` | 重启 PostgreSQL 实例 | `patronictl restart`                   |
| `pt reload`     | `rl` | 重载 PostgreSQL 配置 | `patronictl reload`                    |
| `pt reinit`     | `ri` | 重新初始化成员          | `patronictl reinit`                    |
| `pt switchover` | `so` | 计划内主从切换          | `patronictl switchover`                |
| `pt failover`   | `fo` | 手动故障切换           | `patronictl failover`                  |
| `pt pause`      | `p`  | 暂停自动故障切换         | `patronictl pause`                     |
| `pt resume`     | `r`  | 恢复自动故障切换         | `patronictl resume`                    |
| `pt config`     | `c`  | 查看或修改集群配置        | `patronictl show-config / edit-config` |
{.full-width}

**服务子命令**（`pt service`，封装 `systemctl`）：

以下命令通过 `systemctl` 管理 Patroni 服务本身。

| 命令                   | 缩写          | 描述            |
|:---------------------|:------------|:--------------|
| `pt service start`   | `pt svc up` | 启动 Patroni 服务 |
| `pt service stop`    | `pt svc dn` | 停止 Patroni 服务 |
| `pt service restart` | `pt svc rs` | 重启 Patroni 服务 |
| `pt service reload`  | `pt svc rl` | 重载 Patroni 服务 |
| `pt service status`  | `pt svc st` | 显示服务状态        |
{.full-width}

**服务管理**（`systemctl` 封装）：

以下顶层命令用于直接查看或管理 Patroni 服务；其中 `pt start` / `pt stop` 是隐藏快捷入口。

| 命令          | 缩写   | 描述            | 实现方式                                          |
|:------------|:-----|:--------------|:----------------------------------------------|
| `pt start`  | `up` | 启动 Patroni 服务 | `systemctl start patroni`                     |
| `pt stop`   | `dn` | 停止 Patroni 服务 | `systemctl stop patroni`                      |
| `pt status` | `st` | 显示综合状态        | `systemctl status` + `ps` + `patronictl list` |
| `pt log`    | `l`  | 查看 Patroni 日志 | 读取 Patroni 日志目录中的日志文件                         |
{.full-width}



## 快速入门

```bash
# 查看集群成员状态
pig pt list                    # 列出默认集群成员
pig pt list pg-meta            # 列出指定集群成员
pig pt list -W                 # 持续监视模式
pig pt list -w 5               # 每 5 秒刷新一次

# 查看和修改集群配置
pig pt config                  # 显示当前集群配置（默认 show）
pig pt config set ttl=60       # 修改单个配置项（直接生效）
pig pt config set ttl=60 loop_wait=15  # 修改多个配置项
pig pt config pg max_connections=200   # 修改 PostgreSQL 参数

# 集群运维操作
pig pt restart                 # 重启所有成员的 PostgreSQL（需要确认）
pig pt restart pg-test-1       # 重启指定成员
pig pt restart --pending       # 应用待重启成员（直接执行）
pig pt restart -y              # 集群级重启，跳过确认
pig pt switchover              # 计划内主从切换
pig pt pause                   # 暂停自动故障切换
pig pt resume                  # 恢复自动故障切换

# 管理 Patroni 服务
pig pt status                  # 查看服务状态
pig pt start                   # 隐藏快捷入口：等价于 pig pt svc start
pig pt stop                    # 隐藏快捷入口：等价于 pig pt svc stop
pig pt svc start               # 启动服务
pig pt svc stop                # 停止服务
pig pt log -f                  # 实时查看日志
pig pt log grep ERROR          # 搜索日志
```


## 全局参数

以下参数适用于所有 `pig pt` 子命令：

| 参数       | 简写   | 说明                                   |
|:---------|:-----|:-------------------------------------|
| `--dbsu` | `-U` | 数据库超级用户（默认：`$PIG_DBSU` 或 `postgres`） |
{.full-width}


## 集群操作命令

### pt list

列出 Patroni 集群成员状态。该命令封装了 `patronictl list`，并默认添加 `-e`（扩展输出）和 `-t`（显示时间戳）参数。

```bash
pig pt list                    # 列出默认集群成员
pig pt list pg-meta            # 列出指定集群
pig pt list -W                 # 持续监视模式
pig pt list -w 5               # 每 5 秒刷新一次
pig pt list -w 0.5             # 每 0.5 秒刷新一次
pig pt list pg-test -W -w 3    # 监视 pg-test 集群，3 秒刷新
```

**选项：**

| 参数           | 简写   | 说明                    |
|:-------------|:-----|:----------------------|
| `--watch`    | `-W` | 启用持续监视模式              |
| `--interval` | `-w` | 监视刷新间隔（秒，支持 0.5 这类小数） |
{.full-width}

watch 模式使用实时 `patronictl` 透传输出，不能与 `-o json` / `-o yaml` 结构化输出一起使用；结构化输出会返回 `CodePtWatchModeUnsupported`。


### pt restart

通过 Patroni 重启 PostgreSQL 实例。这会触发 PostgreSQL 的滚动重启，而非重启 Patroni 守护进程本身。

```bash
pig pt restart                   # 重启所有成员（交互式）
pig pt restart pg-test-1         # 重启指定成员
pig pt restart -y                # 集群级重启，跳过确认
pig pt restart --role=replica    # 仅重启从库
pig pt restart --pending         # 重启待重启的成员
pig pt restart --plan            # 预览执行计划
```

**选项：**

| 参数          | 简写   | 说明                        |
|:------------|:-----|:--------------------------|
| `--yes`     | `-y` | 跳过确认                      |
| `--role`    | `-r` | 按角色筛选（leader/replica/any） |
| `--pending` | `-p` | 仅重启待重启的成员                 |
| `--plan`    |      | 仅显示执行计划                   |
{.full-width}

`pt restart` 是条件确认：指定单个成员或使用 `--pending` 时直接执行；未指定成员的集群级滚动重启需要确认，结构化输出中需要显式 `--yes`。
底层 `patronictl restart` 始终由 pig 传入 `--force`，不会再触发 `patronictl` 自己的交互提示。


### pt reload

通过 Patroni 重载 PostgreSQL 配置。这会触发所有成员执行配置重载。

```bash
pig pt reload
```


### pt reinit

重新初始化集群成员。这会从主库重新同步数据。

```bash
pig pt reinit pg-test-1          # 重新初始化指定成员
pig pt reinit pg-test-1 -y       # 跳过确认
pig pt reinit pg-test-1 -w       # 等待完成
pig pt reinit pg-test-1 --plan   # 预览执行计划
```

**选项：**

| 参数       | 简写   | 说明        |
|:---------|:-----|:----------|
| `--yes`  | `-y` | 跳过确认      |
| `--wait` | `-w` | 等待重新初始化完成 |
| `--plan` |      | 仅显示执行计划   |
{.full-width}

**警告：** 此操作会删除目标成员的所有数据并重新同步。文本模式会要求确认；JSON/YAML 执行模式需要显式 `--yes`。



### pt switchover

通过 Patroni 执行计划内的主从切换。（命令别名：`so`）

```bash
pig pt switchover                 # 交互式切换
pig pt switchover -y              # 跳过确认
pig pt switchover -l pg-1 -c pg-2 # 指定当前主库和新主库
pig pt switchover -s "2026-07-01T12:00:00"  # 定时切换
pig pt switchover --plan          # 预览执行计划
pig pt so -c pg-test-1 -y         # 无需确认直接切换至 pg-test-1 实例
```

**选项：**

| 参数            | 简写   | 说明      |
|:--------------|:-----|:--------|
| `--yes`       | `-y` | 跳过确认    |
| `--leader`    | `-l` | 指定当前主库  |
| `--candidate` | `-c` | 指定候选新主库 |
| `--scheduled` | `-s` | 定时切换时间  |
| `--plan`      |      | 仅显示执行计划 |
{.full-width}

相比 `patronictl` 命令行，pig 会从 `/etc/patroni/patroni.yml` 中解析并填充集群 scope，避免用户手动输入集群名称。
执行或确认前，pig 会读取当前拓扑：集群名、当前 Leader、候选从库以及 pause 状态。如果集群已经 pause，命令会拒绝执行并提示先运行 `pig pt resume`。

如果未指定 `--candidate`，pig 不会自行挑选实例，而是将候选选择交给 `patronictl` / Patroni；确认提示会说明“将 leadership 转移给 Patroni 选择的最适格从库”，并列出当前观察到的候选成员。需要指定新主实例时，使用 `--candidate/-c`。



### pt failover

执行手动故障切换。用于主库不可用时强制切换。（命令别名：`fo`）

与 `switchover` 不同，`failover` 不要求当前主库可用，但您 **必须** 指定一个候选新主库。候选可以通过 `--candidate/-c` 指定，也可以作为唯一位置参数传入：`pig pt failover <member>`。

```bash
pig pt failover --candidate pg-2          # 交互式故障切换
pig pt failover pg-2                      # 位置参数形式，等效于 -c pg-2
pig pt failover -c pg-2 -y                # 跳过确认
pig pt failover -c pg-2 --plan            # 预览执行计划
pig pt fo pg-test-2 -y                    # 简写 + 确认
```

**选项：**

| 参数            | 简写   | 说明               |
|:--------------|:-----|:-----------------|
| `--yes`       | `-y` | 跳过确认             |
| `--candidate` | `-c` | 指定候选新主库；也可使用位置参数 |
| `--plan`      |      | 仅显示执行计划          |
{.full-width}

执行或确认前，pig 会读取当前拓扑并检查 pause 状态。如果集群已经 pause，命令会拒绝执行并提示先运行 `pig pt resume`。确认提示会包含集群名、当前 Leader、指定候选新主库以及当前观察到的候选成员，并保留故障切换可能丢数据的警告。


### pt pause

暂停 Patroni 的自动故障切换，进入维护模式，防止在维护期间触发故障切换。
如果集群已经处于维护模式，则该命令会报错。

```bash
pig pt pause                      # 暂停自动故障切换
pig pt pause -w                   # 等待确认
```

**选项：**

| 参数       | 简写   | 说明     |
|:---------|:-----|:-------|
| `--wait` | `-w` | 等待操作完成 |
{.full-width}

**使用场景：** 在执行维护操作（如大版本升级、存储迁移）时暂停自动故障切换，防止误触发。



### pt resume

恢复 Patroni 的自动故障切换，退出维护模式。
如果集群未处于维护模式，则该命令会报错。

```bash
pig pt resume                     # 恢复自动故障切换
pig pt resume -w                  # 等待确认
```

**选项：**

| 参数       | 简写   | 说明     |
|:---------|:-----|:-------|
| `--wait` | `-w` | 等待操作完成 |
{.full-width}



### pt config

显示或修改集群配置。`show` 显示当前配置，`set` 修改 Patroni 动态配置，`pg` 修改 PostgreSQL 参数。

```bash
pig pt config show                      # 显示当前集群配置
pig pt config edit                      # 交互式编辑配置
pig pt config set ttl=60                # 设置 TTL 为 60 秒
pig pt config set ttl=60 loop_wait=15   # 同时修改多个配置项
pig pt config pg max_connections=200    # 修改 PostgreSQL 参数
pig pt config set ttl=60 --plan         # 预览配置修改
```

**子命令：**

| 子命令             | 说明               |
|:----------------|:-----------------|
| `show`          | 显示当前配置           |
| `edit`          | 交互式编辑配置          |
| `set key=value` | 直接设置配置项          |
| `pg key=value`  | 设置 PostgreSQL 参数 |
{.full-width}

**选项：**

| 参数       | 说明                       |
|:---------|:-------------------------|
| `--plan` | 对 `set` / `pg` 操作仅显示执行计划 |
{.full-width}

当你修改 Patroni 本身的配置参数时，需要使用 `pig pt config set k=v` 命令。`set` 与 `pg` 都要求参数是 `key=value` 形式，非 `key=value` token 会被拒绝，而不是静默忽略。`edit` 是交互式操作，不支持结构化输出。


| 配置项                       | 说明                       | 默认值     |
|:--------------------------|:-------------------------|:--------|
| `ttl`                     | Leader 锁的生存时间（秒）         | 30      |
| `loop_wait`               | 主循环休眠时间（秒）               | 10      |
| `retry_timeout`           | DCS 和 PostgreSQL 操作超时（秒） | 10      |
| `maximum_lag_on_failover` | 故障切换时允许的最大延迟（字节）         | 1048576 |
{.full-width}


当你修改 PostgreSQL 本身的配置参数时，需要使用 `pig pt config pg k=v` 命令。该命令会识别已知的 postmaster-context 参数；若修改需要重启的 PostgreSQL 参数，计划和结构化结果会给出 `pig pt list` 与 `pig pt restart --pending` 后续动作。

**注意：** 此命令修改的是存储在 DCS（如 etcd）中的集群动态配置，而非本地配置文件 `/etc/patroni/patroni.yml`。

例如：

```bash
pig pt config set ttl=60 --plan
pig pt config pg shared_buffers=4GB --plan
```



## 服务管理命令

### pt start

启动 Patroni 服务。

```bash
pig pt start                     # 启动 Patroni 服务
pig pt up                        # 别名
```

等效于执行 `sudo systemctl start patroni`。


### pt stop

停止 Patroni 服务。

```bash
pig pt stop                      # 停止 Patroni 服务
pig pt dn                        # 别名
```

等效于执行 `sudo systemctl stop patroni`。

**注意：** 停止 Patroni 服务会导致该节点上的 PostgreSQL 实例也被停止（取决于 Patroni 配置）。


### pt status

显示 Patroni 服务的综合状态，包括：
- systemd 服务状态
- Patroni 进程信息
- 集群成员状态

```bash
pig pt status
```


### pt log

查看 Patroni 服务日志。日志目录默认从 `/etc/patroni/patroni.yml` 的 `log.dir` 读取，未配置时回退到 `/pg/log/patroni`；也可用 `--log-dir` 显式指定。只有 `pt log` 与 `pt log show` 支持 `-o json` 输出 JSONL；日志快照不支持 `yaml` 与 `json-pretty`，follow/tail/grep 不支持结构化输出。

```bash
pig pt log                     # 显示最近 50 行日志
pig pt log -f                  # 实时跟踪日志输出
pig pt log show                # 显示最近日志
pig pt log tail                # 跟踪日志
pig pt log grep ERROR          # 搜索日志
pig pt log -n 100              # 显示最近 100 行日志
pig pt log -f -n 200           # 显示最近 200 行并持续跟踪
```

**子命令：**

| 子命令    | 别名             | 说明              |
|:-------|:---------------|:----------------|
| `show` | `cat, c, s`    | 输出最近 Patroni 日志 |
| `tail` | `t, f, follow` | 持续跟踪 Patroni 日志 |
| `grep` | `g, search`    | 搜索 Patroni 日志   |
{.full-width}

**选项：**

| 参数          | 简写   | 默认值   | 说明       |
|:------------|:-----|:------|:---------|
| `--follow`  | `-f` | false | 实时跟踪日志输出 |
| `--lines`   | `-n` | 50    | 显示的日志行数  |
| `--log-dir` |      | 自动解析  | 日志目录     |
{.full-width}




## pt svc 子命令

`pt svc`（也可写作 `pt service`）提供与顶层服务命令相同的功能，用于明确操作的是 Patroni 守护进程：

```bash
pig pt svc start                 # 启动 Patroni 服务
pig pt svc stop                  # 停止 Patroni 服务
pig pt svc restart               # 重启 Patroni 服务
pig pt svc reload                # 重载 Patroni 服务
pig pt svc status                # 显示服务状态
```

**别名对照：**

| 命令               | 别名   |
|:-----------------|:-----|
| `pt svc start`   | `up` |
| `pt svc stop`    | `dn` |
| `pt svc restart` | `rs` |
| `pt svc reload`  | `rl` |
| `pt svc status`  | `st` |
{.full-width}

其中 `start` 和 `stop` 有专门的 `pt start` 和 `pt stop` 的快捷方式，方便用户直接管理 Patroni 服务。
但请注意，`pt restart` 并非是 `pt svc restart` 的快捷方式，而是用于 Patroni 重启 PostgreSQL 集群的命令，二者功能不同。



## 设计说明

**与 patronictl 的关系：**

`pig pt` 封装了 `patronictl` 的常用操作：
- 集群查询：`list`、`config show`
- 集群管理：`restart`、`reload`、`reinit`、`switchover`、`failover`、`pause`、`resume`
- 配置修改：`config set`、`config pg`、`config edit`
- 服务管理命令（start/stop/restart/reload/status）调用 `systemctl`
- `log` 命令读取 Patroni 日志目录中的日志文件

**默认配置路径：**

| 配置项          | 默认值                                 |
|:-------------|:------------------------------------|
| Patroni 配置文件 | `/etc/patroni/patroni.yml`          |
| 日志目录         | 配置文件 `log.dir`，回退 `/pg/log/patroni` |
| 服务名称         | `patroni`                           |
{.full-width}

**权限处理：**

- 如果当前用户已是 DBSU：直接执行命令
- 如果当前用户是 root：使用 `su - postgres -c "..."` 执行
- 其他用户：使用 `sudo -inu postgres -- ...` 执行

**平台支持：**

此命令专为 Linux 系统设计，服务管理依赖 `systemctl`，日志功能依赖可读取的 Patroni 日志文件。
