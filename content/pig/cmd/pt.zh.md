---
title: "pig pt"
description: "透明运行 patronictl，附带 Pigsty 服务、配置与日志辅助命令"
icon: fa-solid fa-arrows-rotate
weight: 650
---

`pig patroni`（别名 `pig pt`）运行已安装的 `patronictl`，用于查看集群成员、修改集群配置和切换主库。
PIG 负责选择配置文件，并提供本地服务、状态和日志命令。
其他命令使用 patronictl 的参数、交互确认、输出和退出码。

执行集群操作前，先安装 Patroni 并配置集群访问方式。
使用 `pig pt COMMAND --help` 查看原生命令帮助，使用 `pig pt -- --help` 查看 patronictl 根帮助。
单独运行 `pig pt` 显示 PIG 帮助。

## 命令概览

**透传命令**（`patronictl` 原生）：

| 示例                                            | 说明                                  |
|:----------------------------------------------|:------------------------------------|
| `pig pt list [CLUSTER]`                       | 列出集群成员（原生输出，`--format json` 可用）     |
| `pig pt restart CLUSTER [MEMBER]`             | 重启集群 / 成员的 PostgreSQL（原生确认）         |
| `pig pt reload CLUSTER`                       | 重载 PostgreSQL 配置                    |
| `pig pt reinit CLUSTER MEMBER`                | 重新初始化成员（从主库重新同步数据）                  |
| `pig pt switchover CLUSTER [--candidate X]`   | 计划内主从切换                             |
| `pig pt failover CLUSTER --candidate MEMBER`  | 手动故障切换（集群名在前）                |
| `pig pt pause CLUSTER` / `pig pt resume CLUSTER`               | 暂停 / 恢复自动故障切换（维护模式）                 |
| `pig pt show-config CLUSTER` / `pig pt edit-config CLUSTER`            | 查看 / 编辑集群动态配置                       |
| `pig pt query -c 'select 1'`                  | 原生查询（此处 `-c` 是 query 自己的 SQL 参数）    |
{.full-width}

将集群名放在成员名前。这些命令使用已安装版本的
[patronictl 接口](https://patroni.readthedocs.io/en/latest/patronictl.html)。

**本地命令**（PIG 实现）：

| 命令           | 别名            | 说明                                        |
|:-------------|:--------------|:------------------------------------------|
| `pt set`     |               | 修改 PostgreSQL 参数或 Patroni 标量配置       |
| `pt status`  | `st`          | 综合状态：systemd 服务 + Patroni 进程 + 集群成员       |
| `pt service` | `svc`         | 管理本地 patroni systemd 服务                   |
| `pt start`   | `up`          | 等价于 `pt svc start`                 |
| `pt stop`    | `dn`          | 等价于 `pt svc stop`                  |
| `pt log`     | `l`           | 查看本地 Patroni 日志（show / tail / grep）       |
{.full-width}

使用 `pig pt restart CLUSTER` 通过 Patroni 重启 PostgreSQL。
使用 `pig pt svc restart` 重启本地 Patroni 守护进程。

## 查看集群 {#快速入门}

在已配置好 Patroni 的节点上，查看集群成员和动态配置。
将 `pg-meta` 替换为你的集群名：

```bash
pig pt list pg-meta
pig pt show-config pg-meta
```

使用 `pig pt status` 查看本地服务状态，使用 `pig pt log -f` 持续查看日志。

## PIG 选项

以下 PIG 选项 **必须出现在原生命令之前**；一旦出现原生命令词，其后的所有参数都属于 patronictl：

| 参数              | 简写   | 说明                                     |
|:----------------|:-----|:---------------------------------------|
| `--config-file` | `-c` | 显式指定 Patroni/patronictl 配置文件           |
| `--dbsu`        |      | 执行 patronictl 的 OS 用户（默认 `$PIG_DBSU` 或 `postgres`） |
| `--dcs-url`     | `-d` | 覆盖 Patroni DCS URL（`--dcs` 为其别名）       |
| `--insecure`    | `-k` | 允许 TLS 连接不校验证书                         |
{.full-width}

```bash
pig pt -c /path/patroni.yml list pg-meta   # PIG -c：配置文件
pig pt query -c 'select 1'                 # 原生 query 的 -c：SQL 语句，不会被 PIG 截获
```

使用完整的 `--dbsu` 参数；`pig pt` 没有 `-U` 简写。

## 配置文件解析

PIG 按以下顺序选择 patronictl 配置文件：

1. 命令前显式给出的 `-c/--config-file`；
2. 非空的 `PATRONICTL_CONFIG_FILE` 环境变量；
3. `/etc/patroni/patroni.yml`（存在且以 DBSU 身份可读）；
4. `/infra/conf/patronictl.yml`（存在且以 DBSU 身份可读）；
5. 最后使用 `/etc/patroni/patroni.yml`，文件缺失由 patronictl 报错。

PIG 使用显式路径或环境变量路径，即使文件不存在或不可读。相对路径在切换操作系统用户前解析。
两个常规路径按所选用户的读取权限检查。服务命令和原生帮助不需要配置文件。

## 命令分发与输出

PIG 处理上表中的本地命令，其他命令及其后的全部参数原样转发给 patronictl。
将 PIG 参数放在原生命令之前：

```bash
pig pt -c /etc/patroni/patroni.yml list pg-meta
pig pt list pg-meta --format json
```

原生命令保留交互确认、编辑器、流式输出和退出码。
需要 JSON 时，使用原生 `--format json`。PIG 的 `-o json` 和 `-o yaml` 无法格式化这些命令，
放在原生命令之前会报错。

如果原生命令与本地命令同名，使用 `pig pt -- COMMAND ...` 转发。
本地 `set`、`status` 和 `service` 支持 PIG 结构化输出。日志命令的格式见下文。

## pt set

`pt set` 修改配置文件所选集群的参数：

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
- 以 `postgresql.`、`standby_cluster.`、`slots.`、`ignore_slots.` 开头的结构性键会被 **拒绝**，
  并提示改用原生 `pig pt edit-config --set`。

所有键值对按输入顺序合并为 **一次** 原生 `edit-config` 调用，显示一次配置差异，并进行一次确认：

```bash
pig pt set ttl=60 max_connections=200 synchronous_mode=on
# 等价于：patronictl -c <选中配置> edit-config --set ttl=60 --pg max_connections=200 --set synchronous_mode=on
```

- `--yes/-y` 追加原生 `--force` 跳过确认；默认由 patronictl 显示 diff 并负责确认；
- `--plan` 不执行任何修改，仅渲染选中配置与翻译后的原生命令（包含输入的参数值）；
- 结构化输出模式下执行必须携带 `--yes`；
- 值按 YAML 解析：`KEY=null` 与 `KEY=` 均表示删除该键，PIG 原样传递；
- 修改需要重启的 PostgreSQL 参数后，PIG 会提示后续操作：
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

`pt start` / `pt stop`（别名 `up` / `dn`）调用相同的服务操作。
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
`-f`、`tail`、`grep` 使用文本输出。日志快照不支持 YAML 或 `json-pretty`。

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
| `--lines`   | `-n` | 50    | show/tail 的行数；grep 默认搜索全部，指定 `-n` 后只搜索末尾行 |
| `--log-dir` |      | 自动解析  | 日志目录     |
{.full-width}

## 从 v1.5.x 迁移

从 v1.5.x 升级后，按下表将旧脚本改为 patronictl 原生语法：

| v1.5.x 用法                          | v1.6.0 用法                                        |
|:-----------------------------------|:-------------------------------------------------|
| `pig pt failover <候选成员>`           | `pig pt failover CLUSTER --candidate MEMBER`（位置参数为集群名） |
| `pig pt restart [成员]`（自动定位集群）      | `pig pt restart CLUSTER [MEMBER]`（需显式集群名）        |
| `pig pt list -o json`              | `pig pt list --format json`（原生 JSON，schema 不同）   |
| `pig pt config show`               | `pig pt show-config`                             |
| `pig pt config edit`               | `pig pt edit-config`                             |
| `pig pt config set K=V` / `pg K=V` | `pig pt set K=V`                                 |
| `pig pt restart -y`（PIG 确认）       | 原生确认由 patronictl 负责；`pt set -y` 仍有效              |
| `pig pt list -W` / `-w 5`          | 原生 `pig pt list --watch`（patronictl 语义）          |
| 别名 `ls/rs/rl/ri/so/fo/p/r/c`       | 已移除，使用完整原生命令名                                    |
| `--dbsu -U`                        | `--dbsu`（`-U` 简写已移除）                             |
{.full-width}

## 运行条件与权限

服务命令需要 Linux 和 systemd。日志命令需要可读取的 Patroni 日志文件。
PIG 按 `--dbsu`、`$PIG_DBSU`、`postgres` 的顺序选择运行 patronictl 的操作系统用户。
当前用户相同时直接执行；root 使用 `su`；其他用户使用 `sudo -H -u`。
服务管理使用 root 权限。

## 设计记录

[让 Patronictl 自己说话](https://pig.pgsty.com/zh/design/patronictl-passthrough/)记录了 PIG 为什么停止镜像上游命令树，
以及仍然保留了哪些本地职责。
