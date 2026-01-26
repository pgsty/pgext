---
title: "pig pt"
description: "使用 pig pt 子命令管理 Patroni 服务与集群"
icon: RefreshCw
weight: 650
---

`pig patroni` 命令（别名 `pig pt`）用于管理 Patroni 服务和 PostgreSQL HA 集群。它封装了常用的 `patronictl` 和 `systemctl` 操作，提供简化的集群管理体验。


## 命令概览

**集群操作**（patronictl 封装）：

| 命令 | 别名 | 描述 | 实现方式 |
|:----|:----|:----|:--------|
| `pt list` | `ls, l` | 列出集群成员 | `patronictl list -e -t` |
| `pt restart` | `reboot, rt` | 重启 PostgreSQL 实例 | `patronictl restart` |
| `pt reload` | `rl, hup` | 重载 PostgreSQL 配置 | `patronictl reload` |
| `pt reinit` | `ri` | 重新初始化成员 | `patronictl reinit` |
| `pt switchover` | `sw` | 计划内主从切换 | `patronictl switchover` |
| `pt failover` | `fo` | 手动故障切换 | `patronictl failover` |
| `pt pause` | `p` | 暂停自动故障切换 | `patronictl pause` |
| `pt resume` | `r` | 恢复自动故障切换 | `patronictl resume` |
| `pt config` | `cfg, c` | 查看或修改集群配置 | `patronictl show-config / edit-config` |

**服务管理**（systemctl 封装）：

| 命令 | 别名 | 描述 | 实现方式 |
|:----|:----|:----|:--------|
| `pt start` | `boot, up` | 启动 Patroni 服务 | `systemctl start patroni` |
| `pt stop` | `halt, dn, down` | 停止 Patroni 服务 | `systemctl stop patroni` |
| `pt status` | `st, stat` | 显示服务状态 | `systemctl status patroni` |
| `pt log` | `l, lg` | 查看 Patroni 日志 | `journalctl -u patroni` |

**服务子命令**（`pt svc`）：

| 命令 | 别名 | 描述 |
|:----|:----|:----|
| `pt svc start` | `boot, up` | 启动 Patroni 服务 |
| `pt svc stop` | `halt, dn, down` | 停止 Patroni 服务 |
| `pt svc restart` | `reboot, rt` | 重启 Patroni 服务 |
| `pt svc reload` | `rl, hup` | 重载 Patroni 服务 |
| `pt svc status` | `st, stat` | 显示服务状态 |


## 快速入门

```bash
# 查看集群成员状态
pig pt list                    # 列出默认集群成员
pig pt list pg-meta            # 列出指定集群成员
pig pt list -W                 # 持续监视模式
pig pt list -w 5               # 每 5 秒刷新一次

# 查看和修改集群配置
pig pt config                  # 显示当前集群配置
pig pt config ttl=60           # 修改单个配置项（直接生效）
pig pt config ttl=60 loop_wait=15  # 修改多个配置项

# 集群运维操作
pig pt restart                 # 重启所有成员的 PostgreSQL
pig pt restart pg-test-1       # 重启指定成员
pig pt switchover              # 计划内主从切换
pig pt pause                   # 暂停自动故障切换
pig pt resume                  # 恢复自动故障切换

# 管理 Patroni 服务
pig pt status                  # 查看服务状态
pig pt start                   # 启动服务
pig pt stop                    # 停止服务
pig pt log -f                  # 实时查看日志
```


## 全局参数

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--dbsu` | `-U` | 数据库超级用户（默认：`$PIG_DBSU` 或 `postgres`） |


## 集群操作命令

### pt list

列出 Patroni 集群成员状态。

```bash
pig pt list                    # 列出默认集群成员
pig pt list pg-meta            # 列出指定集群
pig pt list -W                 # 持续监视模式
pig pt list -w 5               # 每 5 秒刷新一次
```

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--watch` | `-W` | 启用持续监视模式 |
| `--interval` | `-w` | 监视刷新间隔（秒） |


### pt restart

通过 Patroni 重启 PostgreSQL 实例。

```bash
pig pt restart                   # 重启所有成员（交互式）
pig pt restart pg-test-1         # 重启指定成员
pig pt restart -f                # 跳过确认直接重启
pig pt restart --role=replica    # 仅重启从库
pig pt restart --pending         # 重启待重启的成员
```


### pt reload

通过 Patroni 重载 PostgreSQL 配置。

```bash
pig pt reload
```


### pt reinit

重新初始化集群成员。此操作会从主库重新同步数据。

```bash
pig pt reinit pg-test-1          # 重新初始化指定成员
pig pt reinit pg-test-1 -f       # 跳过确认
pig pt reinit pg-test-1 --wait   # 等待完成
```

{{< callout type="warning" >}}
此操作会删除目标成员的所有数据并重新同步。
{{< /callout >}}


### pt switchover

执行计划内的主从切换。

```bash
pig pt switchover                 # 交互式切换
pig pt switchover -f              # 跳过确认
pig pt switchover --leader pg-1   # 指定当前主库
pig pt switchover --candidate pg-2  # 指定新主库
```


### pt failover

执行手动故障切换。用于主库不可用时强制切换。

```bash
pig pt failover                   # 交互式故障切换
pig pt failover -f                # 跳过确认
pig pt failover --candidate pg-2  # 指定新主库
```


### pt pause

暂停 Patroni 的自动故障切换。

```bash
pig pt pause                      # 暂停自动故障切换
pig pt pause --wait               # 等待确认
```

**使用场景：** 在执行维护操作（如大版本升级、存储迁移）时暂停自动故障切换，防止误触发。


### pt resume

恢复 Patroni 的自动故障切换。

```bash
pig pt resume                     # 恢复自动故障切换
pig pt resume --wait              # 等待确认
```


### pt config

显示或修改集群配置。

```bash
pig pt config                           # 显示当前集群配置
pig pt config show                      # 显示配置（显式）
pig pt config edit                      # 交互式编辑配置
pig pt config set ttl=60                # 设置 TTL 为 60 秒
pig pt config set ttl=60 loop_wait=15   # 同时修改多个配置项
pig pt config pg max_connections=200    # 修改 PostgreSQL 参数
```

**常用配置项：**

| 配置项 | 说明 | 默认值 |
|:------|:----|:------|
| `ttl` | Leader 锁的生存时间（秒） | 30 |
| `loop_wait` | 主循环休眠时间（秒） | 10 |
| `retry_timeout` | DCS 和 PostgreSQL 操作超时（秒） | 10 |
| `maximum_lag_on_failover` | 故障切换时允许的最大延迟（字节） | 1048576 |

{{< callout type="info" >}}
此命令修改的是存储在 DCS（如 etcd）中的集群动态配置，而非本地配置文件。
{{< /callout >}}


## 服务管理命令

### pt start

启动 Patroni 服务。

```bash
pig pt start                     # 启动 Patroni 服务
pig pt up                        # 别名
pig pt boot                      # 别名
```

### pt stop

停止 Patroni 服务。

```bash
pig pt stop                      # 停止 Patroni 服务
pig pt down                      # 别名
pig pt halt                      # 别名
```

{{< callout type="warning" >}}
停止 Patroni 服务会导致该节点上的 PostgreSQL 实例也被停止。
{{< /callout >}}

### pt status

显示 Patroni 服务的综合状态。

```bash
pig pt status
```

### pt log

查看 Patroni 服务日志。

```bash
pig pt log                     # 显示最近 50 行日志
pig pt log -f                  # 实时跟踪日志输出
pig pt log -n 100              # 显示最近 100 行日志
```


## pt svc 子命令

`pt svc` 提供与顶层服务命令相同的功能：

```bash
pig pt svc start                 # 启动 Patroni 服务
pig pt svc stop                  # 停止 Patroni 服务
pig pt svc restart               # 重启 Patroni 服务
pig pt svc reload                # 重载 Patroni 服务
pig pt svc status                # 显示服务状态
```


## 设计说明

**与 patronictl 的关系：**

`pig pt` 封装了 `patronictl` 的常用操作：
- 集群查询：`list`、`config show`
- 集群管理：`restart`、`reload`、`reinit`、`switchover`、`failover`、`pause`、`resume`
- 配置修改：`config set`、`config edit`
- 服务管理命令调用 `systemctl`
- `log` 命令调用 `journalctl`

**平台支持：**

此命令专为 Linux 系统设计，依赖 `systemctl` 和 `journalctl`。
