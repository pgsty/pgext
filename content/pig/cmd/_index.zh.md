---
title: "命令参考"
description: pig CLI 命令参考概览
icon: SquareTerminal
weight: 600
---

`pig` CLI 提供了全面的工具集，用于管理 PostgreSQL 安装、扩展、软件仓库以及从源码构建扩展。使用 `pig help <command>` 查看命令文档。

- [**pig repo**](/zh/pig/cmd/repo/)：管理软件仓库
- [**pig ext**](/zh/pig/cmd/ext/)：管理 PostgreSQL 扩展
- [**pig build**](/zh/pig/cmd/build/)：从源码构建扩展
- **pig install**：使用原生包管理器安装包，并对 PostgreSQL 别名做翻译
- [**pig sty**](/zh/pig/cmd/sty/)：管理 Pigsty 安装
- [**pig pg**](/zh/pig/cmd/pg/)：管理本地 PostgreSQL 服务器
- [**pig pt**](/zh/pig/cmd/pt/)：管理 Patroni HA 集群
- [**pig pb**](/zh/pig/cmd/pb/)：管理 pgBackRest 备份与恢复
- [**pig pitr**](/zh/pig/cmd/pitr/)：进行完整 PITR 工作流
- **pig context**：输出面向人工和 Agent 的环境上下文快照



## 概览

```bash
pig - PostgreSQL 的 Linux 包管理器

Usage:
  pig [command]

Examples:

  pig repo add -ru            # 覆盖现有仓库并更新缓存
  pig install pg18            # 安装 PostgreSQL 18 PGDG 软件包
  pig install pg_duckdb       # 安装指定的 PostgreSQL 扩展
  pig install pgactive -v 18  # 为特定 PG 版本安装扩展

  访问 https://pigsty.io/ext/ 获取详情

PostgreSQL Extension Manager
  build       构建 Postgres 扩展
  ext         管理 PostgreSQL 扩展 (pgext)
  repo        管理 Linux 软件仓库 (apt/dnf)

Pigsty Management Commands
  do          运行管理任务
  pg          管理本地 PostgreSQL 服务器与数据库
  pt          使用 patronictl 管理 Patroni 集群
  pb          管理 pgBackRest 备份与恢复
  pe          管理 pg_exporter 与监控指标
  pitr        编排式时间点恢复
  sty         管理 Pigsty 安装
  context     显示环境上下文快照

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
      --log-level string   日志级别: debug, info, warn, error, fatal, panic (默认 "info")
      --log-path string    日志文件路径，默认为终端输出
  -o, --output string      输出格式：text, yaml, json, json-pretty（默认 "text"）

使用 "pig [command] --help" 获取命令的详细信息。
```


## pig repo

管理 PostgreSQL 软件包的 APT/YUM 仓库，详情请参考 [`pig repo`](/zh/pig/cmd/repo/)

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

管理 PostgreSQL 扩展和内核包，详情请参考 [`pig ext`](/zh/pig/cmd/ext/)

```bash
pig ext list    duck             # 搜索扩展
pig ext info    pg_duckdb        # 扩展详情
pig ext status                   # 显示已安装的扩展
pig ext add     pg_duckdb -y     # 安装扩展
pig ext rm      old_extension    # 移除扩展
pig ext update pg_duckdb         # 更新指定扩展
pig ext scan                     # 扫描已安装的扩展
pig ext import  pg_duckdb        # 下载以供离线使用
pig ext link    polar            # 链接 PG/PolarDB 安装到 PATH
pig ext reload                   # 刷新扩展目录
```




## pig build

从源码构建 PostgreSQL 扩展，详情请参考 [`pig build`](/zh/pig/cmd/build/)

```bash
# 环境设置
pig build spec                   # 初始化构建规格
pig build repo -m                # 使用镜像/代理设置仓库
pig build tool                   # 安装构建工具
pig build rust -m                # 使用镜像配置安装 Rust
pig build pgrx -b                # 安装 PGRX 并包含 beta PG 自动探测

# 构建扩展
pig build pkg citus              # 完整构建流程 = get + dep + ext
pig build get citus              # 下载源码
pig build dep citus              # 安装依赖
pig build ext citus              # 构建包
```


## pig install

使用系统原生包管理器安装软件包，并对 PostgreSQL 内核、扩展及常用别名做包名翻译。需要直接传递系统包名时，可使用 `-n/--no-translation`。

```bash
pig install pg_duckdb            # 安装扩展并自动翻译包名
pig install pg18                 # 安装 PostgreSQL 18 内核包组
pig install nginx htop vim       # 安装多个系统包
pig install unknown-package -n   # 禁用翻译，按原始包名安装
pig install pg18 --plan          # 预览安装计划
pig install pg_vector -y         # 自动确认安装
```


## pig sty

安装 Pigsty 发行版，详情请参考 [`pig sty`](/zh/pig/cmd/sty/)

```bash
pig sty init -m                  # 从镜像安装 Pigsty 到 ~/pigsty
pig sty boot                     # 安装 Ansible 依赖
pig sty conf                     # 生成配置
pig sty deploy                   # 运行部署 playbook
```


## pig context

输出环境上下文快照，覆盖主机、PostgreSQL、Patroni、pgBackRest 与已安装扩展。该命令适合排障和自动化脚本快速了解当前节点状态。

```bash
pig context                      # 文本输出
pig context -o json              # JSON 输出
pig context -m postgres          # 只输出 PostgreSQL 模块（默认包含 host）
pig context -m postgres,!host    # 排除 host 模块
```


## pig pg

管理本地 PostgreSQL 服务器，详情请参考 [`pig pg`](/zh/pig/cmd/pg/)

```bash
pig pg init                      # 初始化数据目录
pig pg start                     # 启动 PostgreSQL
pig pg stop                      # 停止 PostgreSQL
pig pg status                    # 查看状态
pig pg psql mydb                 # 连接数据库
pig pg ps                        # 查看当前连接
pig pg vacuum mydb               # 清理数据库
pig pg tune -p olap              # 生成调优参数
pig pg fork dev                  # 创建本地一次性物理副本
pig pg fork list                 # 列出本地 fork
pig pg log tail                  # 实时查看日志
```


## pig pt

管理 Patroni HA 集群，详情请参考 [`pig pt`](/zh/pig/cmd/pt/)

```bash
pig pt list                      # 列出集群成员
pig pt config show               # 显示集群配置
pig pt config set ttl=60         # 修改集群配置
pig pt status                    # 查看服务状态
pig pt log -f                    # 实时查看日志
```


## pig pb

管理 pgBackRest 备份与恢复，详情请参考 [`pig pb`](/zh/pig/cmd/pb/)

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

执行编排式时间点恢复（PITR），详情请参考 [`pig pitr`](/zh/pig/cmd/pitr/)

```bash
pig pitr -d                      # 恢复到最新数据
pig pitr -t "2025-01-01 12:00:00+08"  # 恢复到指定时间
pig pitr -I                      # 恢复到备份一致性点
pig pitr -d --plan               # 显示执行计划（不实际执行）
pig pitr -d -y                   # 跳过确认（自动化）
```
