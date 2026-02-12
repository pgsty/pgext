---
title: "命令参考"
description: pig CLI 命令参考概览
icon: SquareTerminal
weight: 600
---


`pig` CLI 提供了全面的工具集，用于管理 PostgreSQL 安装、扩展、软件仓库以及从源码构建扩展。使用 `pig help <command>` 查看命令文档。

- [**pig repo**](/zh/pig/cmd/repo)：管理软件仓库
- [**pig ext**](/zh/pig/cmd/ext)：管理 PostgreSQL 扩展
- [**pig build**](/zh/pig/cmd/build)：从源码构建扩展
- [**pig sty**](/zh/pig/cmd/sty)：管理 Pigsty 安装
- [**pig pg**](/zh/pig/cmd/pg)：管理本地 PostgreSQL 服务器
- [**pig pt**](/zh/pig/cmd/pt)：管理 Patroni HA 集群
- [**pig pb**](/zh/pig/cmd/pb)：管理 pgBackRest 备份与恢复

## 概览

```bash
pig - PostgreSQL 的 Linux 包管理器

Usage:
  pig [command]

Examples:

  pig repo add -ru            # 覆盖现有仓库并更新缓存
  pig install pg17            # 安装 PostgreSQL 17 PGDG 软件包
  pig install pg_duckdb       # 安装指定的 PostgreSQL 扩展
  pig install pgactive -v 18  # 为特定 PG 版本安装扩展

  访问 https://pgext.cloud 获取详情

PostgreSQL Extension Manager
  build       构建 Postgres 扩展
  ext         管理 PostgreSQL 扩展 (pgext)
  repo        管理 Linux 软件仓库 (apt/dnf)

Pigsty Management Commands
  do          运行管理任务
  patroni     使用 patronictl 管理 Patroni 集群
  pg_exporter 管理 pg_exporter 与监控指标
  pgbackrest  管理 pgBackRest 备份与恢复
  pitr        编排式时间点恢复
  postgres    管理本地 PostgreSQL 服务器与数据库
  sty         管理 Pigsty 安装

Additional Commands:
  completion  生成指定 shell 的自动补全脚本
  help        获取任何命令的帮助信息
  install     使用原生包管理器安装软件包
  status      显示环境状态
  update      升级 pig 自身
  version     显示 pig 版本信息

Flags:
      --debug              启用调试模式
  -h, --help               获取帮助信息
  -H, --home string        Pigsty 主目录路径
  -i, --inventory string   配置清单路径
  -t, --toggle             帮助信息中的占位参数
      --log-level string   日志级别: debug, info, warn, error, fatal, panic (默认 "info")
      --log-path string    日志文件路径，默认为终端输出

使用 "pig [command] --help" 获取命令的详细信息。
```


## pig repo

管理 PostgreSQL 软件包的 APT/YUM 仓库，详情请参考 [`pig repo`](/zh/pig/cmd/repo)

```bash
pig repo list                    # 列出可用仓库
pig repo info   pgdg             # 显示仓库详情
pig repo status                  # 检查当前仓库状态
pig repo add    pgdg pigsty -u   # 添加仓库
pig repo rm     old-repo         # 移除仓库
pig repo update                  # 更新软件包缓存
pig repo create /www/pigsty      # 创建本地仓库
pig repo cache                   # 创建离线包
pig repo boot                    # 从离线包引导
```




## pig ext

管理 PostgreSQL 扩展和内核包，详情请参考 [`pig ext`](/zh/pig/cmd/ext)

```bash
pig ext list    duck             # 搜索扩展
pig ext info    pg_duckdb        # 扩展详情
pig ext status                   # 显示已安装的扩展
pig ext add     pg_duckdb -y     # 安装扩展
pig ext rm      old_extension    # 移除扩展
pig ext update                   # 更新扩展
pig ext scan                     # 扫描已安装的扩展
pig ext import  pg_duckdb        # 下载以供离线使用
pig ext link    17               # 链接 PG 版本到 PATH
pig ext reload                   # 刷新扩展目录
```




## pig build

从源码构建 PostgreSQL 扩展，详情请参考 [`pig build`](/zh/pig/cmd/build)

```bash
# 环境设置
pig build spec                   # 初始化构建规格
pig build repo                   # 设置仓库
pig build tool                   # 安装构建工具
pig build rust -y                # 强制重装 Rust（默认不重装）
pig build pgrx                   # 安装 PGRX 框架

# 构建扩展
pig build pkg citus              # 完整构建流程 = get + dep + ext
pig build get citus              # 下载源码
pig build dep citus              # 安装依赖
pig build ext citus              # 构建包
```


## pig sty

安装 Pigsty 发行版，详情请参考 [`pig sty`](/zh/pig/cmd/sty)

```bash
pig sty init                     # 安装 Pigsty 到 ~/pigsty
pig sty boot                     # 安装 Ansible 依赖
pig sty conf                     # 生成配置
pig sty deploy                   # 运行部署 playbook
```


## pig pg

管理本地 PostgreSQL 服务器，详情请参考 [`pig pg`](/zh/pig/cmd/pg)

```bash
pig pg init                      # 初始化数据目录
pig pg start                     # 启动 PostgreSQL
pig pg stop                      # 停止 PostgreSQL
pig pg status                    # 查看状态
pig pg psql mydb                 # 连接数据库
pig pg ps                        # 查看当前连接
pig pg vacuum mydb               # 清理数据库
pig pg log tail                  # 实时查看日志
```


## pig pt

管理 Patroni HA 集群，详情请参考 [`pig pt`](/zh/pig/cmd/pt)

```bash
pig pt list                      # 列出集群成员
pig pt config                    # 显示集群配置
pig pt config ttl=60             # 修改集群配置
pig pt status                    # 查看服务状态
pig pt log -f                    # 实时查看日志
```


## pig pb

管理 pgBackRest 备份与恢复，详情请参考 [`pig pb`](/zh/pig/cmd/pb)

```bash
pig pb info                      # 显示备份信息
pig pb ls                        # 列出所有备份
pig pb backup                    # 创建备份
pig pb backup full               # 全量备份
pig pb restore -d                # 恢复到最新
pig pb restore -t "2025-01-01"   # 恢复到指定时间
pig pb log tail                  # 实时查看日志
```


## pig pitr

执行编排式时间点恢复（PITR），详情请参考 [`pig pitr`](/zh/pig/cmd/pitr)

```bash
pig pitr -d                      # 恢复到最新数据
pig pitr -t "2025-01-01 12:00"   # 恢复到指定时间
pig pitr -I                      # 恢复到备份一致性点
pig pitr -d --dry-run            # 显示执行计划（不实际执行）
pig pitr -d -y                   # 跳过确认（自动化）
```
