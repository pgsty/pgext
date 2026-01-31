---
title: "pig pg"
description: "使用 pig pg 子命令管理本地 PostgreSQL 服务器"
icon: Database
weight: 640
---

`pig pg` 命令（别名 `pig postgres`）用于管理本地 PostgreSQL 服务器和数据库。它封装了 `pg_ctl`、`psql`、`vacuumdb` 等原生工具，提供简化的服务器管理体验。

```bash
pig pg - 管理本地 PostgreSQL 服务器（pg_ctl, psql, vacuumdb）

控制命令（通过 pg_ctl 或 systemctl）：
  pig pg init                      初始化 PostgreSQL 数据目录
  pig pg start                     启动 PostgreSQL 服务器
  pig pg stop                      停止 PostgreSQL 服务器
  pig pg restart                   重启 PostgreSQL 服务器
  pig pg reload                    重载 PostgreSQL 服务器配置
  pig pg status                    显示 PostgreSQL 服务器状态
  pig pg promote                   将备库提升为主库
  pig pg role                      检测并打印 PostgreSQL 角色

连接与查询（通过 psql）：
  pig pg psql [db] [-c sql]        连接到 PostgreSQL
  pig pg ps                        显示当前连接
  pig pg kill [-a] [-x] [-u user] [-d db] [-q sql] [-w secs]

维护命令（通过 vacuumdb 和 pg_repack）：
  pig pg vacuum  [db] [-a]         清理数据库
  pig pg analyze [db] [-a]         分析数据库
  pig pg freeze  [db] [-a]         冻结清理表
  pig pg repack  [db] [-a]         在线重整数据库

日志命令：
  pig pg log list                  列出日志文件
  pig pg log tail <logfile>        tail -f 查看日志
  pig pg log cat  <logfile>        cat 查看日志
  pig pg log less <logfile>        less 查看日志
  pig pg log grep <logfile> <pat>  grep 搜索日志

服务管理（通过 systemctl）：
  pig pg svc start                 启动 postgres 服务
  pig pg svc stop                  停止 postgres 服务
  pig pg svc restart               重启 postgres 服务
  pig pg svc reload                重载 postgres 服务
  pig pg svc status                显示 postgres 服务状态
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

**连接与查询**：

| 命令 | 别名 | 描述 | 备注 |
|:----|:----|:----|:----|
| `pg psql` | `sql, connect` | 连接到数据库 | 封装 psql |
| `pg ps` | `activity, act` | 显示当前连接 | 查询 pg_stat_activity |
| `pg kill` | `k` | 终止连接 | 默认 dry-run 模式 |

**数据库维护**：

| 命令 | 别名 | 描述 | 备注 |
|:----|:----|:----|:----|
| `pg vacuum` | `vac, vc` | 清理表 | 封装 vacuumdb |
| `pg analyze` | `ana, az` | 分析表 | 封装 vacuumdb --analyze-only |
| `pg repack` | `rp` | 在线重整表 | 需要 pg_repack 扩展 |

**日志工具**：

| 命令 | 别名 | 描述 |
|:----|:----|:----|
| `pg log list` | `ls` | 列出日志文件 |
| `pg log tail` | `t, f` | 实时查看日志 |
| `pg log cat` | `c` | 输出日志内容 |
| `pg log less` | `vi, v` | 用 less 查看 |
| `pg log grep` | `g, search` | 搜索日志 |

**服务子命令**（`pg svc`）：

| 命令 | 别名 | 描述 |
|:----|:----|:----|
| `pg svc start` | `boot, up` | 启动 postgres 服务 |
| `pg svc stop` | `halt, dn, down` | 停止 postgres 服务 |
| `pg svc restart` | `reboot, rt` | 重启 postgres 服务 |
| `pg svc reload` | `rl, hup` | 重载 postgres 服务 |
| `pg svc status` | `st, stat` | 显示服务状态 |


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
pig pg log grep ERROR             # 搜索错误日志
```


## 全局参数

| 参数 | 简写 | 默认值 | 说明 |
|:----|:----|:------|:----|
| `--version` | `-v` | 自动检测 | PostgreSQL 主版本号 |
| `--data` | `-D` | `/pg/data` | 数据目录路径 |
| `--dbsu` | `-U` | `postgres` | 数据库超级用户 |


## 服务控制命令

### pg init

初始化 PostgreSQL 数据目录。

```bash
pig pg init                       # 使用默认设置初始化
pig pg init -v 17                 # 指定 PostgreSQL 17
pig pg init -D /data/pg17         # 指定数据目录
pig pg init -k                    # 启用数据校验和
pig pg init -f                    # 强制初始化（删除已有数据）
```

### pg start

启动 PostgreSQL 服务器。

```bash
pig pg start                      # 使用默认设置启动
pig pg up                         # 别名
pig pg start -D /data/pg17        # 指定数据目录
pig pg start -l /pg/log/pg.log    # 重定向输出到日志文件
```

### pg stop

停止 PostgreSQL 服务器。

```bash
pig pg stop                       # 快速停止（默认）
pig pg down                       # 别名
pig pg stop -m smart              # 等待客户端断开
pig pg stop -m immediate          # 立即关闭
```

### pg restart

重启 PostgreSQL 服务器。

```bash
pig pg restart                    # 快速重启
pig pg reboot                     # 别名
pig pg restart -m immediate       # 立即重启
```

### pg reload

重载 PostgreSQL 配置。

```bash
pig pg reload                     # 重载配置
pig pg hup                        # 别名
```

### pg status

显示 PostgreSQL 服务器状态，包括进程信息和相关服务状态。

```bash
pig pg status                     # 查看服务状态
pig pg st                         # 别名
```

### pg promote

将备库提升为主库。

```bash
pig pg promote                    # 提升备库
pig pg pro                        # 别名
```

### pg role

检测 PostgreSQL 实例的角色。

```bash
pig pg role                       # 输出：primary、replica 或 unknown
pig pg role -V                    # 详细输出
```


## 连接与查询命令

### pg psql

通过 psql 连接到 PostgreSQL 数据库。

```bash
pig pg psql                       # 连接到 postgres 数据库
pig pg psql mydb                  # 连接到指定数据库
pig pg psql mydb -c "SELECT 1"    # 执行单条命令
pig pg psql -f script.sql         # 执行 SQL 脚本文件
```

### pg ps

显示 PostgreSQL 当前连接。

```bash
pig pg ps                         # 显示客户端连接
pig pg ps -a                      # 显示所有连接
pig pg ps -u admin                # 按用户筛选
pig pg ps -d mydb                 # 按数据库筛选
```

### pg kill

终止 PostgreSQL 连接。**默认为 dry-run 模式**。

```bash
pig pg kill                       # 显示将被终止的连接（dry-run）
pig pg kill -x                    # 实际终止连接
pig pg kill --pid 12345 -x        # 终止指定 PID
pig pg kill -u admin -x           # 终止指定用户的连接
pig pg kill -s idle -x            # 终止空闲连接
```


## 数据库维护命令

### pg vacuum

清理数据库表。

```bash
pig pg vacuum                     # 清理当前数据库
pig pg vacuum mydb                # 清理指定数据库
pig pg vacuum -a                  # 清理所有数据库
pig pg vacuum mydb -t mytable     # 清理指定表
pig pg vacuum mydb --full         # VACUUM FULL
```

### pg analyze

分析数据库表以更新统计信息。

```bash
pig pg analyze                    # 分析当前数据库
pig pg analyze mydb               # 分析指定数据库
pig pg analyze -a                 # 分析所有数据库
```

### pg repack

在线重整数据库表。需要 `pg_repack` 扩展。

```bash
pig pg repack mydb                # 重整数据库中所有表
pig pg repack -a                  # 重整所有数据库
pig pg repack mydb -t mytable     # 重整指定表
pig pg repack mydb -j 4           # 使用 4 个并行任务
```


## 日志命令

### pg log list

列出日志目录中的日志文件。

```bash
pig pg log list                   # 列出默认目录中的日志
pig pg log ls                     # 别名
```

### pg log tail

实时查看日志文件。

```bash
pig pg log tail                   # 查看最新日志
pig pg log tail -n 100            # 显示最后 100 行后开始跟踪
```

### pg log cat

输出日志文件内容。

```bash
pig pg log cat                    # 输出最新日志
pig pg log cat -n 100             # 输出最后 100 行
```

### pg log less

用 less 打开日志文件。

```bash
pig pg log less                   # 用 less 打开最新日志
pig pg log vi                     # 别名
```

### pg log grep

搜索日志文件。

```bash
pig pg log grep ERROR             # 搜索包含 ERROR 的行
pig pg log grep -i error          # 忽略大小写
pig pg log grep -C 3 ERROR        # 显示前后 3 行上下文
```


## pg svc 子命令

通过 systemctl 管理 PostgreSQL 服务：

```bash
pig pg svc start                 # 启动 postgres 服务
pig pg svc stop                  # 停止 postgres 服务
pig pg svc restart               # 重启 postgres 服务
pig pg svc reload                # 重载 postgres 服务
pig pg svc status                # 显示服务状态
```


## 设计说明

**权限处理：**

- 如果当前用户已是 DBSU：直接执行命令
- 如果当前用户是 root：使用 `su - postgres -c "..."` 执行
- 其他用户：使用 `sudo -inu postgres -- ...` 执行

**平台支持：**

此命令专为 Linux 系统设计，部分功能依赖 `systemctl`。
