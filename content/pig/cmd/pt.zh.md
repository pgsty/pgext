---
title: "pig pt"
description: "透明运行 patronictl，附带 Pigsty 服务、配置与日志辅助命令"
icon: fa-solid fa-arrows-rotate
weight: 650
---

`pig patroni` 命令（别名 `pig pt`）自 v1.6.0 起是已安装 `patronictl` 的**透明启动器**：
pig 只负责选择配置文件与少量本地辅助命令，其余一切命令与参数**原样转发**给 `patronictl`，
使用原生的参数、交互确认、输出与退出码——patronictl 的新功能无需等待 pig 发版即可使用。

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

第一个非选项命令词决定分发方式：

- `set`、`start`/`up`、`stop`/`dn`、`service`/`svc`、`status`/`st`、`log`/`l` 走 pig 的本地实现；
- **其余任何命令词**（`list`、`restart`、`reload`、`reinit`、`switchover`、`failover`、
  `pause`、`resume`、`show-config`、`edit-config`、`query`、`history`、`topology`、`dsn`、`version` 等）
  连同其后的全部参数，逐字转发给 `patronictl`；
- `pig pt -- <命令> ...` 可以绕过本地命令名冲突（例如 `pig pt -- set` 会把 `set` 交给 patronictl）；
- `pig pt` 不带命令时打印 pig 帮助，不会运行 patronictl；
  原生子命令帮助用 `pig pt <命令> --help`，patronictl 根帮助用 `pig pt -- --help`。


## 命令概览

**透传命令**（`patronictl` 原生）：

| 示例                                            | 说明                                  |
|:----------------------------------------------|:------------------------------------|
| `pig pt list [CLUSTER]`                       | 列出集群成员（原生输出，`--format json` 可用）     |
| `pig pt restart CLUSTER [MEMBER]`             | 重启集群 / 成员的 PostgreSQL（原生确认）         |
| `pig pt reload CLUSTER`                       | 重载 PostgreSQL 配置                    |
| `pig pt reinit CLUSTER MEMBER`                | 重新初始化成员（从主库重新同步数据）                  |
| `pig pt switchover CLUSTER [--candidate X]`   | 计划内主从切换                             |
| `pig pt failover CLUSTER --candidate MEMBER`  | 手动故障切换（**位置参数是集群名**）                |
| `pig pt pause / resume CLUSTER`               | 暂停 / 恢复自动故障切换（维护模式）                 |
| `pig pt show-config / edit-config`            | 查看 / 编辑集群动态配置                       |
| `pig pt query -c 'select 1'`                  | 原生查询（此处 `-c` 是 query 自己的 SQL 参数）    |
{.full-width}

转发命令的位置参数遵循 patronictl 原生的**集群优先**（CLUSTER-first）语义，
确认提示、输出格式与退出码（含 Click 用法错误退出码 `2`）均由 patronictl 负责。

**本地命令**（pig 实现）：

| 命令           | 别名            | 说明                                        |
|:-------------|:--------------|:------------------------------------------|
| `pt set`     |               | 修改 PostgreSQL 参数或 Patroni 标量配置（语法糖）       |
| `pt status`  | `st`          | 综合状态：systemd 服务 + Patroni 进程 + 集群成员       |
| `pt service` | `svc`         | 管理本地 patroni systemd 服务                   |
| `pt start`   | `up`          | 隐藏快捷入口，等价于 `pt svc start`                 |
| `pt stop`    | `dn`          | 隐藏快捷入口，等价于 `pt svc stop`                  |
| `pt log`     | `l`           | 查看本地 Patroni 日志（show / tail / grep）       |
{.full-width}

注意：顶层 `pt restart` **不是**重启 patroni 守护进程的快捷方式，它会转发给
`patronictl restart` 用于重启 PostgreSQL；重启守护进程请使用 `pt svc restart`。


## 快速入门

```bash
# 集群操作（原生透传，集群优先语义）
pig pt list pg-meta                         # 列出集群成员
pig pt list pg-meta --format json           # 原生 JSON 输出
pig pt restart pg-test --pending            # 应用待重启成员
pig pt restart pg-test pg-test-1            # 重启指定成员
pig pt switchover pg-test --candidate pg-2  # 计划内切换
pig pt failover pg-test --candidate pg-2    # 手动故障切换
pig pt pause pg-test                        # 进入维护模式
pig pt show-config pg-test                  # 查看动态配置

# 配置修改（本地语法糖，一次 edit-config 调用）
pig pt set ttl=60                           # Patroni 标量配置 → --set
pig pt set max_connections=200              # PostgreSQL 参数 → --pg
pig pt set ttl=60 max_connections=200 -y    # 混合修改，跳过确认
pig pt set shared_buffers=4GB --plan        # 仅预览翻译后的原生命令

# 服务与日志（本地实现）
pig pt status                               # 综合状态
pig pt svc restart                          # 重启 patroni 守护进程
pig pt log -f                               # 实时跟踪日志
pig pt log grep ERROR                       # 搜索日志
```


## Pig/PT 选项

以下包装层选项**必须出现在原生命令之前**；一旦出现原生命令词，其后的所有参数都属于 patronictl：

| 参数              | 简写   | 说明                                     |
|:----------------|:-----|:---------------------------------------|
| `--config-file` | `-c` | 显式指定 Patroni/patronictl 配置文件           |
| `--dbsu`        |      | 执行 patronictl 的 OS 用户（默认 `$PIG_DBSU` 或 `postgres`） |
| `--dcs-url`     | `-d` | 覆盖 Patroni DCS URL（`--dcs` 为其别名）       |
| `--insecure`    | `-k` | 允许 TLS 连接不校验证书                         |
{.full-width}

```bash
pig pt -c /path/patroni.yml list pg-meta   # 包装层 -c：配置文件
pig pt query -c 'select 1'                 # 原生 query 的 -c：SQL 语句，不会被 pig 截获
```

> v1.6.0 起 `--dbsu` 不再提供 `-U` 简写（`pg` / `pb` 命令不受影响）。


## 配置文件解析

每次配置相关的 patronictl 调用都会由 pig 注入唯一的根 `-c <路径>`，解析优先级：

1. 命令前显式给出的 `-c/--config-file`；
2. 非空的 `PATRONICTL_CONFIG_FILE` 环境变量；
3. `/etc/patroni/patroni.yml`（存在且以 DBSU 身份可读）；
4. `/infra/conf/patronictl.yml`（存在且以 DBSU 身份可读）；
5. 兜底回退 `/etc/patroni/patroni.yml`，让 patronictl 的报错指向常规位置。

显式路径与环境变量路径具有权威性：文件缺失或不可读时 pig 不会静默换用其他候选；
相对路径会在切换 OS 用户前转换为绝对路径。常规候选文件的可读性按**实际执行用户（DBSU）**探测，
而非简单检查权限位。解析是惰性的：`pt svc start` 这类纯 systemd 操作不会触发配置解析；
原生 `-h/--help` 走免配置快速路径，在未配置 Patroni 的机器上也能查看帮助。


## 透传执行与输出模式

转发执行直接继承 stdin / stdout / stderr 与终端：原生交互确认、`edit-config` 的编辑器、
`--watch` 流式刷新、退出码全部保真；pig 不做输出捕获、不重复渲染错误、不重试、
也不在执行前后读取集群或 DCS 状态。逻辑调用形态为：

```text
patronictl -c <选中配置> [--dcs-url URL] [--insecure] <原生命令与参数...>
```

**结构化输出**：转发路径只支持 pig 的 `text` 模式。在原生命令之前出现的
`-o json` / `-o yaml` 会被**明确拒绝**（提示改用原生输出选项）；
出现在原生命令之后的参数原样转发，由 patronictl 自行校验：

```bash
pig -o json pt list          # 拒绝：请使用原生输出选项
pig pt -o json list          # 拒绝：请使用原生输出选项
pig pt list --format json    # 原样转发，patronictl 原生 JSON
pig pt list -o json          # 原样转发，由 patronictl 校验
```

本地命令（`set` / `status` / `log` / `service`）保留 pig 的结构化输出行为。


## pt set

`pt set` 是唯一的本地配置语法糖，作用于选中配置对应的集群：

```bash
pig pt set KEY=VALUE [KEY=VALUE ...] [--yes] [--plan]
```

**键分类规则**：

- 以下 Patroni 顶层标量键翻译为原生 `--set`：
  `loop_wait`、`ttl`、`retry_timeout`、`primary_race_backoff`、
  `maximum_lag_on_failover`、`maximum_lag_on_syncnode`、`max_timelines_history`、
  `primary_start_timeout`、`primary_stop_timeout`、`synchronous_mode`、
  `synchronous_mode_strict`、`synchronous_node_count`、`failsafe_mode`、
  `check_timeline`、`member_slots_ttl`（`pause` 不在其列——请使用原生 `pause`/`resume`）；
- 其余键一律视为 PostgreSQL 参数，翻译为原生 `--pg`（含 `timescaledb.telemetry_level` 这类带点自定义 GUC）；
- 以 `postgresql.`、`standby_cluster.`、`slots.`、`ignore_slots.` 开头的结构性键会被**拒绝**，
  并提示改用原生 `pig pt edit-config --set`。

所有键值对按输入顺序合并为**一次**原生 `edit-config` 调用，产生一次 diff、一次确认、一次 DCS 更新：

```bash
pig pt set ttl=60 max_connections=200 synchronous_mode=on
# 等价于：patronictl -c <选中配置> edit-config --set ttl=60 --pg max_connections=200 --set synchronous_mode=on
```

- `--yes/-y` 追加原生 `--force` 跳过确认；默认由 patronictl 显示 diff 并负责确认；
- `--plan` 不执行任何修改，仅渲染选中配置与翻译后的原生命令（**计划可能包含敏感值**）；
- 结构化输出模式下执行必须携带 `--yes`（禁止隐藏交互提示）；
- 值按 YAML 解析：`KEY=null` 与 `KEY=` 均表示删除该键，pig 原样传递；
- 修改需要重启的 PostgreSQL 参数后，pig 会提示后续操作：
  先 `pig pt list`，再 `pig pt restart CLUSTER --pending`（需显式集群名）。


## 服务管理

`pt service`（别名 `pt svc`）管理本地 `patroni` systemd 服务：

| 命令                   | 别名          | 说明            |
|:---------------------|:------------|:--------------|
| `pt service start`   | `pt svc up` | 启动 Patroni 服务 |
| `pt service stop`    | `pt svc dn` | 停止 Patroni 服务 |
| `pt service restart` | `pt svc rs` | 重启 Patroni 服务 |
| `pt service reload`  | `pt svc rl` | 重载 Patroni 服务 |
| `pt service status`  | `pt svc st` | 显示服务状态        |
{.full-width}

顶层 `pt start` / `pt stop`（别名 `up` / `dn`）是隐藏快捷入口，调用同一本地实现。
停止 Patroni 服务可能导致该节点上的 PostgreSQL 一并停止（取决于 Patroni 配置）。

### pt status

显示综合状态：systemd 服务状态、Patroni 进程信息、以及来自 patronictl 的集群成员状态。

```bash
pig pt status
pig pt st -o json                # 结构化输出
```

### pt log

查看本地 Patroni 日志。日志目录取自选中配置的 `log.dir`，未配置时回退到 `/pg/log/patroni`，
也可用 `--log-dir` 显式指定。仅 `pt log` 与 `pt log show` 支持 `-o json`（JSONL 快照）；
follow / tail / grep 是终端流式操作，不支持结构化输出。

```bash
pig pt log                     # 显示最近 50 行日志
pig pt log -f                  # 实时跟踪日志输出
pig pt log -n 100              # 显示最近 100 行
pig pt log show -o json        # JSONL 快照
pig pt log tail -n 100         # 跟踪日志
pig pt log grep ERROR          # 搜索日志
```

| 子命令    | 别名             | 说明              |
|:-------|:---------------|:----------------|
| `show` | `cat, c, s`    | 输出最近 Patroni 日志 |
| `tail` | `t, f, follow` | 持续跟踪 Patroni 日志 |
| `grep` | `g, search`    | 搜索 Patroni 日志   |
{.full-width}

| 参数          | 简写   | 默认值   | 说明       |
|:------------|:-----|:------|:---------|
| `--follow`  | `-f` | false | 实时跟踪日志输出 |
| `--lines`   | `-n` | 50    | 显示的日志行数  |
| `--log-dir` |      | 自动解析  | 日志目录     |
{.full-width}


## 从 v1.5.x 迁移

v1.6.0 的透传重写是**破坏性变更**，升级前请检查自动化脚本：

| v1.5.x 用法                          | v1.6.0 用法                                        |
|:-----------------------------------|:-------------------------------------------------|
| `pig pt failover <候选成员>`           | ⚠ `pig pt failover CLUSTER --candidate MEMBER`（位置参数语义反转：现在是**集群名**） |
| `pig pt restart [成员]`（自动定位集群）      | `pig pt restart CLUSTER [MEMBER]`（需显式集群名）        |
| `pig pt list -o json`              | `pig pt list --format json`（原生 JSON，schema 不同）   |
| `pig pt config show`               | `pig pt show-config`                             |
| `pig pt config edit`               | `pig pt edit-config`                             |
| `pig pt config set K=V` / `pg K=V` | `pig pt set K=V`                                 |
| `pig pt restart -y`（pig 确认门）       | 原生确认由 patronictl 负责；`pt set -y` 仍有效              |
| `pig pt list -W` / `-w 5`          | 原生 `pig pt list --watch`（patronictl 语义）          |
| 别名 `ls/rs/rl/ri/so/fo/p/r/c`       | 已移除，使用完整原生命令名                                    |
| `--dbsu -U`                        | `--dbsu`（`-U` 简写已移除）                             |
{.full-width}

此外：转发命令的退出码为 patronictl 原生值（用法错误为 `2`）；
switchover / failover 不再有 pig 侧的 pause 预检——维护模式语义完全由 Patroni 负责。


## 设计说明

**单一权威**：集群控制路径只有一条——已安装的 `patronictl` + 一份选中配置，
经由 Patroni REST API / DCS 生效。pig 不内嵌 DCS 客户端或 REST 控制引擎，
不维护 patronictl 命令清单，不对转发命令做确认、预检、重试或结果改写。

**权限处理**（与 v1.5.x 相同）：

- 当前用户已是 DBSU：直接执行；
- 当前用户是 root：`su - postgres -c "..."` 执行；
- 其他用户：`sudo -inu postgres -- ...` 执行。

**平台支持**：此命令专为 Linux 设计，服务管理依赖 `systemctl`，日志功能依赖可读取的 Patroni 日志文件。
