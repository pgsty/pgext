---
title: "命令参考"
description: pig CLI 命令参考概览
icon: SquareTerminal
weight: 600
---

`pig` CLI 提供了全面的工具集，用于管理 PostgreSQL 安装、扩展、软件仓库以及从源码构建扩展。使用 `pig help <command>` 查看命令文档。

{{< cards cols="5" >}}
{{< card link="/pig/cmd/repo"  title="pig repo"  subtitle="管理软件仓库" icon="library" >}}
{{< card link="/pig/cmd/ext"   title="pig ext"   subtitle="管理 PostgreSQL 扩展"   icon="cube" >}}
{{< card link="/pig/cmd/build" title="pig build" subtitle="从源码构建扩展"  icon="view-grid" >}}
{{< card link="/pig/cmd/sty"   title="pig sty"   subtitle="管理 Pigsty 安装"   icon="cloud-download" >}}
{{< /cards >}}

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

  访问 https://pig.pgsty.com 获取详情

PostgreSQL Extension Manager
  build       构建 Postgres 扩展
  ext         管理 PostgreSQL 扩展 (pgext)
  repo        管理 Linux 软件仓库 (apt/dnf)

Pigsty Management Commands
  do          运行管理任务
  patroni     使用 patronictl 管理 PostgreSQL
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
      --log-level string   日志级别: debug, info, warn, error, fatal, panic (默认 "info")
      --log-path string    日志文件路径，默认为终端输出

使用 "pig [command] --help" 获取命令的详细信息。
```

------

## 快速入门

### 1. 设置软件仓库

在安装 PostgreSQL 或扩展之前，需要配置软件包仓库：

```bash
# 快速设置 - 添加所有必需仓库 (pgdg + pigsty + node)
sudo pig repo add -ru

# 或者：添加特定仓库
sudo pig repo add pgdg pigsty -u
```

### 2. 安装 PostgreSQL

```bash
# 安装 PostgreSQL 17
sudo pig ext install pg17

# 链接到系统 PATH
sudo pig ext link 17
```

### 3. 安装扩展

```bash
# 搜索扩展
pig ext list                    # 列出所有
pig ext list duckdb             # 搜索特定扩展

# 安装扩展
sudo pig ext add pg_duckdb postgis pgvector
```

### 4. 构建扩展（可选）

```bash
# 设置构建环境
pig build spec

# 构建扩展
pig build pkg citus
```

------

## 核心命令

### [`pig repo`](/pig/cmd/repo/) - 仓库管理

管理 PostgreSQL 软件包的 APT/YUM 仓库：

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

**核心特性：**
- 支持 RPM (EL) 和 DEB (Debian/Ubuntu) 系统
- 区域镜像（中国等）
- 离线包支持
- 基于模块的仓库组织

------

### [`pig ext`](/pig/cmd/ext/) - 扩展管理

管理 PostgreSQL 扩展和内核包：

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

**核心特性：**
- 400+ PostgreSQL 扩展
- 多版本 PostgreSQL 支持
- 自动依赖解析
- 基于类别的浏览
- 离线安装支持

------

### [`pig build`](/pig/cmd/build/) - 构建扩展

从源码构建 PostgreSQL 扩展：

```bash
# 环境设置
pig build spec                   # 初始化构建规格
pig build repo                   # 设置仓库
pig build tool                   # 安装构建工具
pig build rust -y                # 安装 Rust（用于 Rust 扩展）
pig build pgrx                   # 安装 PGRX 框架

# 构建扩展
pig build pkg citus              # 完整构建流程
pig build get citus              # 下载源码
pig build dep citus              # 安装依赖
pig build ext citus              # 构建包
```

**核心特性：**
- 支持 100+ 扩展
- RPM 和 DEB 包生成
- Rust/PGRX 扩展支持
- 多 PostgreSQL 版本构建
- 自定义扩展支持

------

### [`pig sty`](/pig/cmd/sty/) - Pigsty 管理

管理 Pigsty PostgreSQL 发行版：

```bash
pig sty init                     # 安装 Pigsty 到 ~/pigsty
pig sty boot                     # 安装 Ansible 依赖
pig sty conf                     # 生成配置
pig sty install                  # 运行安装 playbook
```

**Pigsty 特性：**
- 开箱即用的 PostgreSQL 发行版
- 基于 Patroni 的高可用
- 时间点恢复 (PITR)
- 完整的监控堆栈
- 基础设施即代码 (IaC)

------

## 常用工作流

### 安装 PostgreSQL 及扩展

```bash
# 1. 设置仓库
sudo pig repo add -ru

# 2. 安装 PostgreSQL 17
sudo pig ext install pg17
sudo pig ext link 17

# 3. 安装流行扩展
sudo pig ext add pg_duckdb postgis pgvector timescaledb

# 4. 验证安装
pig ext status
```

### 构建自定义扩展

```bash
# 1. 设置构建环境
pig build spec
pig build tool

# 2. 构建扩展
pig build pkg my_extension

# 3. 安装构建的包
sudo rpm -ivh ~/rpmbuild/RPMS/x86_64/my_extension*.rpm  # EL
sudo dpkg -i ~/my_extension*.deb                         # Debian
```

### 离线安装

```bash
# 在联网机器上：
sudo pig repo add -ru
pig ext import pg17 pg_duckdb postgis
pig repo create
pig repo cache

# 将 /tmp/pkg.tgz 传输到离线机器

# 在离线机器上：
pig repo boot
sudo pig repo add local
sudo pig ext install pg17
sudo pig ext add pg_duckdb postgis
```

------

## 环境检测

`pig` 工具会自动检测您的环境：

```bash
pig status                       # 显示全面状态
```

**检测信息：**
- 操作系统 (EL/Debian/Ubuntu)
- 操作系统版本和架构
- PostgreSQL 安装
- 活动 PostgreSQL 版本
- 已安装的扩展
- 仓库配置

------

## 最佳实践

1. **总是先更新仓库**：在安装前运行 `pig repo add -ru`
2. **使用版本标志**：需要时用 `-v` 指定 PostgreSQL 版本
3. **检查兼容性**：安装前使用 `pig ext info`
4. **变更前备份**：重大更新前创建备份
5. **先在开发环境测试**：生产部署前在开发环境测试扩展
6. **定期更新目录**：定期运行 `pig ext reload`
7. **记录您的软件栈**：保留已安装扩展的记录
8. **使用离线包**：为隔离系统准备离线包
9. **按需构建**：对自定义或新版本使用构建命令
10. **变更后监控**：安装后检查日志

------

## 获取帮助

```bash
# 通用帮助
pig help

# 命令特定帮助
pig repo --help
pig ext --help
pig build --help

# 子命令帮助
pig ext add --help
pig repo info --help
pig build pkg --help
```

问题反馈请访问：https://github.com/pgsty/pig/issues