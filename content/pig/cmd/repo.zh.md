---
title: "命令：repo"
description: 如何使用 pig repo 子命令管理软件仓库？
icon: SquareTerminal
weight: 610
---

`pig repo` 命令是一个综合性的软件包仓库管理工具。它提供了添加、移除、创建和管理软件仓库的功能，支持 RPM 系统（RHEL/CentOS/Rocky/Alma）和 Debian 系统（Debian/Ubuntu）。该命令是安装 PostgreSQL 及其扩展所需仓库设置的关键工具。

------

## 子命令

| 命令 | 描述 |
|---------|------------|
| `list`   | 打印可用仓库列表 |
| `info`   | 获取仓库详细信息 |
| `status` | 显示当前仓库状态 |
| `add`    | 添加新仓库 |
| `set`    | 清空、覆盖并更新仓库 |
| `rm`     | 移除仓库 |
| `update` | 更新仓库缓存 |
| `create` | 创建本地 YUM/APT 仓库 |
| `cache`  | 从本地仓库创建离线包 |
| `boot`   | 从离线包引导仓库 |

------

## 概览

```bash
pig repo - 管理 Linux APT/YUM 仓库

  pig repo list                    # 可用仓库列表                   (info)
  pig repo info   [repo|module...] # 显示仓库信息                   (info)
  pig repo status                  # 显示当前仓库状态               (info)
  pig repo add    [repo|module...] # 添加仓库和模块                 (root)
  pig repo rm     [repo|module...] # 移除仓库和模块                 (root)
  pig repo update                  # 更新仓库包缓存                 (root)
  pig repo create                  # 在当前系统创建仓库             (root)
  pig repo boot                    # 从离线包引导仓库               (root)
  pig repo cache                   # 将仓库缓存为离线包             (root)

Usage:
  pig repo [command]

别名:
  repo, r

可用命令:
  add         添加新仓库
  boot        从离线包引导仓库
  cache       从本地仓库创建离线包
  create      创建本地 YUM/APT 仓库
  info        获取仓库详细信息
  list        打印可用仓库列表
  rm          移除仓库
  set         清空、覆盖并更新仓库
  status      显示当前仓库状态
  update      更新仓库缓存

标志:
  -h, --help   获取帮助信息

全局标志:
      --debug              启用调试模式
  -H, --home string        pigsty 主目录路径
  -i, --inventory string   配置清单路径
      --log-level string   日志级别: debug, info, warn, error, fatal, panic (默认 "info")
      --log-path string    日志文件路径，默认为终端输出
```

------

## 快速入门

为 PostgreSQL 安装设置仓库的最快方式：

```bash
# 方法 1：添加所有必要仓库并更新缓存（推荐）
pig repo add -ru                 # 移除旧仓库，添加所有必要仓库，更新缓存

# 方法 2：温和方式 - 仅添加所需仓库
pig repo add pgdg pigsty -u      # 添加 PGDG 和 Pigsty 仓库并更新缓存

# 方法 3：分步添加仓库
pig repo add pgdg                # 添加 PGDG 官方仓库
pig repo add pigsty              # 添加 Pigsty 扩展仓库
pig repo add node                # 添加操作系统默认仓库
pig repo update                  # 更新包缓存
```

------

## 核心概念

### 模块

在 pig 中，仓库被组织为**模块** - 服务于特定目的的仓库逻辑组。一个模块可以根据以下条件映射到不同的实际仓库：
- 操作系统发行版（RHEL/Rocky/Debian/Ubuntu）
- 操作系统主版本（EL 8/9/10，Debian 12/13，Ubuntu 22/24）
- CPU 架构（x86_64/aarch64）
- 地理区域（default/china）

### 必要模块

PostgreSQL 安装通常需要三个核心模块：

1. **`pgdg`**：PostgreSQL 全球开发组官方仓库
   - 包含 PostgreSQL 内核包
   - 包括 100+ 官方扩展
   - 提供 PostgreSQL 工具和实用程序

2. **`pigsty`**：Pigsty 扩展仓库
   - 包含 200+ PGDG 中没有的额外扩展
   - 包括特殊扩展和实用程序
   - 提供各种扩展的预编译包

3. **`node`**：操作系统默认仓库
   - 包含 Linux 系统库和依赖
   - 构建和运行 PostgreSQL 所需
   - 包括开发工具和编译器

### `all` 别名

`all` 伪模块是一个便捷的别名，包含所有三个必要模块（pgdg + pigsty + node）。这是大多数用户的推荐起点。

------

## 命令参考

### `repo list`

列出当前系统可用的所有仓库模块和单个仓库。

```bash
pig repo list                    # 列出为当前系统过滤的仓库
pig repo list all                # 列出所有仓库（不过滤）
```

**示例输出：**
```yaml
repo_modules:   # 可用模块: 19
  - all       : pigsty-infra, pigsty-pgsql, pgdg, baseos, appstream, extras, powertools, crb, epel
  - pigsty    : pigsty-infra, pigsty-pgsql
  - pgdg      : pgdg
  - node      : baseos, appstream, extras, powertools, crb, epel
  - infra     : pigsty-infra, nginx
  - pgsql     : pigsty-pgsql, pgdg-common, pgdg13, pgdg14, pgdg15, pgdg16, pgdg17
  - extra     : pgdg-extras, timescaledb, citus
  - docker    : docker-ce
  - kube      : kubernetes
  # ... 更多模块
```

仓库定义存储在 `~/.pig/repo.yml` 或使用内置默认值。

------

### `repo info`

显示特定仓库或模块的详细信息，包括 URL、元数据和区域镜像。

```bash
pig repo info pgdg               # 显示 pgdg 模块的信息
pig repo info pigsty pgdg        # 显示多个模块的信息
pig repo info all                # 显示所有模块的信息
```

**示例输出：**
```bash
#-------------------------------------------------
名称       : pgdg
摘要       : PostgreSQL 全球开发组官方仓库
可用       : 是 (debian d12 amd64)
模块       : pgsql
系统架构   : [x86_64, aarch64]
系统发行版 : deb [11,12,20,22,24], el [8,9]
元数据     : trusted=yes
基础 URL   : http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg main
     中国  : https://mirrors.tuna.tsinghua.edu.cn/postgresql/repos/apt/ ${distro_codename}-pgdg main

# 当前系统的仓库内容:
deb [trusted=yes] http://apt.postgresql.org/pub/repos/apt/ bookworm-pgdg main

# 中国镜像 (使用 --region china 时):
deb [trusted=yes] https://mirrors.tuna.tsinghua.edu.cn/postgresql/repos/apt/ bookworm-pgdg main
```

------

### `repo status`

显示系统上的当前仓库配置。

```bash
pig repo status
```

**示例输出：**
```bash
# 对于 EL 系统：
仓库目录: /etc/yum.repos.d
活动仓库:
  - pgdg.repo
  - pigsty.repo
  - rocky.repo

总计: 15 个仓库已启用

# 对于 Debian/Ubuntu 系统：
仓库目录: /etc/apt/sources.list.d
活动仓库:
  - pgdg.list
  - pigsty.list
  - debian.list

总计: 8 个仓库已启用
```

------

### `repo add`

添加仓库配置文件到系统。需要 root/sudo 权限。

```bash
# 基本用法
pig repo add pgdg                # 添加 PGDG 仓库
pig repo add pgdg pigsty         # 添加多个仓库
pig repo add all                 # 添加所有必要仓库 (pgdg + pigsty + node)

# 带选项
pig repo add pigsty -u           # 添加并更新缓存
pig repo add all -r              # 添加前移除现有仓库
pig repo add all -ru             # 移除、添加并更新（完全重置）

# 区域镜像
pig repo add pgdg --region china  # 使用中国镜像
```

**选项：**
- `-u, --update`：添加仓库后运行包缓存更新
- `-r, --remove`：添加新仓库前移除现有仓库
- `--region <region>`：使用区域镜像（默认/中国）

**文件位置：**
- EL 系统：`/etc/yum.repos.d/<module>.repo`
- Debian 系统：`/etc/apt/sources.list.d/<module>.list`
- 备份：仓库目录中的 `backup/` 子目录

**常用工作流：**

```bash
# 全新安装 - 推荐方法
sudo pig repo add -ru            # 清洁方式

# 添加特定扩展仓库
sudo pig repo add pgdg -u        # 仅 PGDG
sudo pig repo add extra -u       # 额外的第三方仓库

# 为离线环境设置
sudo pig repo add local          # 使用本地仓库
```

------

### `repo set`

等同于 `repo add --remove`。清空现有仓库并设置新仓库。

```bash
pig repo set                     # 替换为默认仓库
pig repo set pgdg pigsty -u      # 替换为特定仓库并更新
pig repo set all --region china  # 使用中国镜像
```

当您想要确保干净的仓库配置，没有旧的或冲突的仓库时，这很有用。

------

### `repo rm`

移除仓库配置文件并备份它们。

```bash
pig repo rm                      # 移除所有仓库
pig repo rm pgdg                 # 移除特定仓库
pig repo rm pgdg pigsty -u       # 移除并更新缓存
```

**备份位置：**
- EL：`/etc/yum.repos.d/backup/`
- Debian：`/etc/apt/sources.list.d/backup/`

------

### `repo update`

更新包管理器缓存以反映仓库更改。

```bash
pig repo update                  # 更新包缓存
```

**等效命令：**
- EL 系统：`yum makecache` 或 `dnf makecache`
- Debian 系统：`apt update`

------

### `repo create`

为离线安装或缓存创建本地包仓库。

```bash
pig repo create                  # 在默认位置创建 (/www/pigsty)
pig repo create /srv/repo        # 在自定义位置创建
pig repo create /www/pigsty /www/backup  # 多个位置
```

**要求：**
- EL 系统：`createrepo_c` 包
- Debian 系统：`dpkg-dev` 包

**示例工作流：**
```bash
# 创建本地仓库结构
sudo pig repo create /www/pigsty

# 复制包到仓库
cp *.rpm /www/pigsty/     # 对于 EL
cp *.deb /www/pigsty/     # 对于 Debian

# 重新创建仓库元数据
sudo pig repo create /www/pigsty

# 添加本地仓库到系统
sudo pig repo add local
```

------

### `repo cache`

创建仓库内容的压缩 tarball 用于离线分发。

```bash
pig repo cache                   # 默认：/www 到 /tmp/pkg.tgz
pig repo cache -f                # 强制覆盖现有文件
pig repo cache -d /srv           # 自定义源目录
pig repo cache pigsty mssql      # 仅特定仓库
```

**选项：**
- `-d, --dir`：源目录（默认：`/www/`）
- `-p, --path`：输出路径（默认：`/tmp/pkg.tgz`）
- `-f`：强制覆盖现有包

**示例工作流：**
```bash
# 在线系统上：创建离线包
pig repo add -ru                 # 设置仓库
pig install pg17 pg_duckdb       # 安装包
pig repo create                  # 创建本地仓库
pig repo cache                   # 创建离线包

# 将 pkg.tgz 传输到离线系统

# 离线系统上：从包引导
pig repo boot                    # 解压并设置
pig repo add local              # 使用本地仓库
pig install pg17 pg_duckdb      # 从本地安装
```

------

### `repo boot`

从离线包解压并设置本地仓库。

```bash
pig repo boot                    # 默认：/tmp/pkg.tgz 到 /www
pig repo boot -p /mnt/pkg.tgz   # 自定义包路径
pig repo boot -d /srv           # 自定义目标目录
```

**选项：**
- `-p, --path`：包路径（默认：`/tmp/pkg.tgz`）
- `-d, --dir`：目标目录（默认：`/www/`）

------

## 常见场景

### 场景 1：全新 PostgreSQL 安装

```bash
# 设置仓库
sudo pig repo add -ru

# 安装 PostgreSQL 17
sudo pig ext install pg17

# 安装流行扩展
sudo pig ext add pg_duckdb postgis timescaledb
```

### 场景 2：离线环境

```bash
# 在联网机器上：
sudo pig repo add -ru
sudo pig ext install pg17
sudo pig ext add pg_duckdb postgis
sudo pig repo create
sudo pig repo cache

# 将 /tmp/pkg.tgz 传输到离线机器

# 在离线机器上：
sudo pig repo boot
sudo pig repo add local
sudo pig ext install pg17
sudo pig ext add pg_duckdb postgis
```

### 场景 3：使用区域镜像

```bash
# 对于中国用户
sudo pig repo add all --region china -u

# 检查镜像 URL
pig repo info pgdg
```

### 场景 4：故障排除仓库问题

```bash
# 检查当前状态
pig repo status

# 备份并清理现有仓库
sudo pig repo rm

# 添加新仓库
sudo pig repo add all -u

# 验证
pig repo status
```

------

## 高级模块参考

### 基础设施模块

- **`infra`**：监控和可观测性堆栈（Grafana、Prometheus）
- **`docker`**：Docker CE 仓库
- **`kube`**：Kubernetes 仓库
- **`grafana`**：Grafana 官方仓库

### 数据库模块

- **`mssql`**：WiltonDB（SQL Server 兼容性）
- **`mysql`**：MySQL/MariaDB 仓库
- **`mongo`**：MongoDB 仓库
- **`redis`**：Redis 仓库

### PostgreSQL 扩展模块

- **`extra`**：额外的第三方扩展
- **`pgml`**：PostgresML 机器学习
- **`groonga`**：PGroonga 全文搜索
- **`ivory`**：IvorySQL Oracle 兼容性

### 特殊模块

- **`local`**：位于 `/www/pigsty` 的本地仓库
- **`haproxy`**：HAProxy 负载均衡器

------

## 最佳实践

1. **添加仓库后始终更新**：使用 `-u` 标志或运行 `pig repo update` 以确保包列表是最新的。

2. **使用模块组**：使用模块组如 `all` 而不是添加单个仓库，以简化管理。

3. **移除前备份**：`rm` 命令会自动备份仓库，但对于关键系统，考虑手动备份。

4. **检查兼容性**：添加前使用 `pig repo info` 验证仓库对您的操作系统的可用性。

5. **区域镜像**：如果有可用的区域镜像，使用 `--region` 标志以获得更快的下载速度。

6. **离线准备**：对于离线环境，先在联网系统上准备包。

------

## 故障排除

### 权限被拒绝

大多数仓库操作需要 root/sudo 权限：
```bash
sudo pig repo add all -u
```

### 仓库不可用

检查操作系统兼容性：
```bash
pig repo info <repo_name>
pig status  # 检查检测到的操作系统
```

### 添加仓库后找不到包

更新包缓存：
```bash
sudo pig repo update
```

### 冲突的仓库

清理并重置：
```bash
sudo pig repo set all -u
```

### 下载缓慢

使用区域镜像：
```bash
sudo pig repo add all --region china -u
```