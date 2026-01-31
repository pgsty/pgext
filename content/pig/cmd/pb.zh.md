---
title: "pig pb"
description: "使用 pig pb 子命令管理 pgBackRest 备份与时间点恢复"
icon: Archive
weight: 660
---

`pig pgbackrest` 命令（别名 `pig pb`）用于管理 pgBackRest 备份和时间点恢复（PITR）。它封装了常用的 `pgbackrest` 操作，提供简化的备份管理体验。所有命令均以数据库超级用户身份执行。

```bash
pig pb - 管理 pgBackRest 备份和时间点恢复

信息查询：
  pig pb info                      显示备份信息
  pig pb ls                        列出备份
  pig pb ls repo                   列出已配置的仓库
  pig pb ls stanza                 列出所有 stanza

备份与恢复：
  pig pb backup                    创建备份（自动：全量/增量）
  pig pb backup full               创建全量备份
  pig pb restore                   从备份恢复（PITR）
  pig pb restore -t "..."          恢复到指定时间
  pig pb expire                    清理过期备份

Stanza 管理：
  pig pb create                    创建 stanza（首次设置）
  pig pb upgrade                   升级 stanza（PG 升级后）
  pig pb delete                    删除 stanza（危险！）

控制命令：
  pig pb check                     验证备份完整性
  pig pb start                     启用 pgBackRest 操作
  pig pb stop                      禁用 pgBackRest 操作
  pig pb log                       查看 pgBackRest 日志

示例：
  pig pb info                      # 显示所有备份信息
  pig pb backup                    # 自动：无备份则全量，否则增量
  pig pb backup full               # 全量备份
  pig pb restore                   # 恢复到最新（默认）
  pig pb restore -t "2025-01-01 12:00:00+08"  # 恢复到指定时间
  pig pb create                    # 初始化 stanza
  pig pb expire                    # 按保留策略清理
```


## 命令概览

**信息查询**：

| 命令 | 描述 | 实现方式 |
|:----|:----|:--------|
| `pb info` | 显示备份仓库信息 | `pgbackrest info` |
| `pb ls` | 列出备份集 | `pgbackrest info` |
| `pb ls repo` | 列出配置的仓库 | 解析 pgbackrest.conf |
| `pb ls stanza` | 列出所有 stanza | 解析 pgbackrest.conf |

**备份与恢复**：

| 命令 | 描述 | 实现方式 |
|:----|:----|:--------|
| `pb backup` | 创建备份 | `pgbackrest backup` |
| `pb restore` | 从备份恢复（PITR） | `pgbackrest restore` |
| `pb expire` | 清理过期备份 | `pgbackrest expire` |

**Stanza 管理**：

| 命令 | 描述 | 实现方式 |
|:----|:----|:--------|
| `pb create` | 创建 stanza（首次设置） | `pgbackrest stanza-create` |
| `pb upgrade` | 升级 stanza（PG 大版本升级后） | `pgbackrest stanza-upgrade` |
| `pb delete` | 删除 stanza（危险操作！） | `pgbackrest stanza-delete` |

**控制命令**：

| 命令 | 别名 | 描述 | 实现方式 |
|:----|:----|:----|:--------|
| `pb check` | | 验证备份仓库完整性 | `pgbackrest check` |
| `pb start` | | 启用 pgBackRest 操作 | `pgbackrest start` |
| `pb stop` | | 禁用 pgBackRest 操作 | `pgbackrest stop` |
| `pb log` | `l, lg` | 查看日志 | `tail/cat` 日志文件 |


## 快速入门

```bash
# 查看备份信息
pig pb info                          # 显示所有备份信息
pig pb info -o json                  # JSON 格式输出
pig pb ls                            # 列出所有备份
pig pb ls repo                       # 列出配置的仓库
pig pb ls stanza                     # 列出所有 stanza

# 创建备份（必须在主库执行）
pig pb backup                        # 自动模式：无备份则全量，否则增量
pig pb backup full                   # 全量备份
pig pb backup diff                   # 差异备份
pig pb backup incr                   # 增量备份

# 恢复（PITR）
pig pb restore                       # 恢复到最新（WAL 流末尾）
pig pb restore -I                    # 恢复到备份一致性点
pig pb restore -t "2025-01-01 12:00:00+08"  # 恢复到指定时间
pig pb restore -n savepoint          # 恢复到命名还原点

# Stanza 管理
pig pb create                        # 初始化 stanza
pig pb upgrade                       # PG 大版本升级后升级 stanza
pig pb check                         # 验证仓库完整性

# 清理
pig pb expire                        # 按保留策略清理
pig pb expire --dry-run              # 干运行模式
```


## 全局参数

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--stanza` | `-s` | 指定 stanza 名称（自动检测） |
| `--config` | `-c` | pgBackRest 配置文件路径 |
| `--repo` | `-r` | 仓库编号（多仓库场景） |
| `--dbsu` | `-U` | 数据库超级用户（默认：`postgres`） |


## 信息查询命令

### pb info

显示备份仓库的详细信息。

```bash
pig pb info                          # 显示所有备份信息
pig pb info -o json                  # JSON 格式输出
pig pb info -s pg-meta               # 指定 stanza
pig pb info --set 20250101-120000F   # 查看特定备份集
```

### pb ls

列出备份集、仓库或 stanza。

```bash
pig pb ls                            # 列出所有备份
pig pb ls repo                       # 列出配置的仓库
pig pb ls stanza                     # 列出所有 stanza
pig pb ls -o json                    # JSON 格式输出
```


## 备份与恢复命令

### pb backup

创建备份。必须在主库上执行。

```bash
pig pb backup                        # 自动模式备份
pig pb backup full                   # 全量备份
pig pb backup diff                   # 差异备份
pig pb backup incr                   # 增量备份
pig pb backup --dry-run              # 干运行模式
```

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--type` | `-t` | 备份类型：full/diff/incr（默认：自动） |
| `--force` | `-f` | 强制备份 |
| `--dry-run` | | 干运行模式 |


### pb restore

从备份恢复数据库（PITR）。

```bash
pig pb restore                       # 恢复到最新（WAL 流末尾）
pig pb restore -I                    # 恢复到备份一致性点
pig pb restore -t "2025-01-01 12:00:00+08"  # 恢复到指定时间
pig pb restore -n savepoint          # 恢复到命名还原点
pig pb restore -l 0/1234567          # 恢复到指定 LSN
pig pb restore -x 12345              # 恢复到指定事务 ID
pig pb restore --set 20250101-120000F  # 从指定备份集恢复
```

| 参数 | 简写 | 说明 |
|:----|:----|:----|
| `--immediate` | `-I` | 恢复到备份一致性点 |
| `--time` | `-t` | 恢复到指定时间戳 |
| `--name` | `-n` | 恢复到命名还原点 |
| `--lsn` | `-l` | 恢复到指定 LSN |
| `--xid` | `-x` | 恢复到指定事务 ID |
| `--set` | `-b` | 从特定备份集恢复 |
| `--target-action` | | 恢复后操作：promote/pause/shutdown |

{{< callout type="warning" >}}
执行恢复前必须停止 PostgreSQL 服务。对于完整的 PITR 工作流，推荐使用 `pig pitr` 命令。
{{< /callout >}}


### pb expire

清理过期备份。

```bash
pig pb expire                        # 按保留策略清理
pig pb expire --dry-run              # 干运行模式
pig pb expire --set 20250101-120000F  # 清理指定备份集
```


## Stanza 管理命令

### pb create

创建 stanza。首次设置备份时使用。

```bash
pig pb create                        # 创建默认 stanza
pig pb create -s pg-meta             # 创建指定 stanza
```

### pb upgrade

升级 stanza。在 PostgreSQL 大版本升级后使用。

```bash
pig pb upgrade                       # 升级默认 stanza
pig pb upgrade -s pg-meta            # 升级指定 stanza
```

### pb delete

删除 stanza。

```bash
pig pb delete                        # 删除默认 stanza
pig pb delete -s pg-meta             # 删除指定 stanza
pig pb delete -f                     # 强制删除
```

{{< callout type="error" >}}
此操作会删除 stanza 及其所有备份数据，无法恢复！
{{< /callout >}}


## 控制命令

### pb check

验证备份仓库完整性。

```bash
pig pb check                         # 验证默认 stanza
pig pb check -s pg-meta              # 验证指定 stanza
```

### pb start

启用 pgBackRest 操作。

```bash
pig pb start                         # 启用所有操作
pig pb start -s pg-meta              # 启用指定 stanza
```

### pb stop

禁用 pgBackRest 操作。

```bash
pig pb stop                          # 禁用所有操作
pig pb stop -s pg-meta               # 禁用指定 stanza
pig pb stop -f                       # 强制停止
```

### pb log

查看 pgBackRest 日志。

```bash
pig pb log                           # 显示最近日志
pig pb log tail                      # 实时跟踪日志
pig pb log -n 100                    # 显示最近 100 行
```


## 设计说明

**与 pgbackrest 的关系：**

`pig pb` 封装了 `pgbackrest` 的常用操作，提供更友好的命令行界面。

**默认配置路径：**

| 配置项 | 默认值 |
|:------|:------|
| 配置文件 | `/etc/pgbackrest/pgbackrest.conf` |
| 日志目录 | `/pg/log/pgbackrest` |

**权限处理：**

- 如果当前用户已是 DBSU：直接执行命令
- 如果当前用户是 root：使用 `su - postgres -c "..."` 执行
- 其他用户：使用 `sudo -inu postgres -- ...` 执行

**平台支持：**

此命令专为 Linux 系统设计。
