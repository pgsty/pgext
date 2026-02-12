---
title: "命令：repo"
description: 如何使用 pig repo 子命令管理软件仓库？
icon: SquareTerminal
weight: 610
---


`pig repo` 命令是一个综合性的软件包仓库管理工具。它提供了添加、移除、创建和管理软件仓库的功能，支持 RPM 系统（RHEL/CentOS/Rocky/Alma）和 Debian 系统（Debian/Ubuntu）。

```bash
pig repo - Manage Linux APT/YUM Repo

  pig repo list                    # available repo list             (info)
  pig repo info   [repo|module...] # show repo info                  (info)
  pig repo status                  # show current repo status        (info)
  pig repo add    [repo|module...] # add repo and modules            (root)
  pig repo rm     [repo|module...] # remove repo & modules           (root)
  pig repo update                  # update repo pkg cache           (root)
  pig repo create                  # create repo on current system   (root)
  pig repo boot                    # boot repo from offline package  (root)
  pig repo cache                   # cache repo as offline package   (root)

Examples:
  pig repo add -ru                 # add all repo and update cache (brute but effective)
  pig repo add pigsty -u           # gentle version, only add pigsty repo and update cache
  pig repo add node pgdg pigsty    # essential repo to install postgres packages
  pig repo add all                 # all = node + pgdg + pigsty
  pig repo add all extra           # extra module has non-free and some 3rd repo for certain extensions
  pig repo update                  # update repo cache
  pig repo create                  # update local repo /www/pigsty meta
  pig repo boot                    # extract /tmp/pkg.tgz to /www/pigsty
  pig repo cache                   # cache /www/pigsty into /tmp/pkg.tgz
```

| 命令            | 描述              | 备注                |
|:--------------|:----------------|:------------------|
| `repo list`   | 打印可用仓库与模块列表     |                   |
| `repo info`   | 获取仓库详细信息        |                   |
| `repo status` | 显示当前仓库状态        |                   |
| `repo add`    | 添加新仓库           | 需要 sudo 或 root 权限 |
| `repo set`    | 清空、覆盖并更新仓库      | 需要 sudo 或 root 权限 |
| `repo rm`     | 移除仓库            | 需要 sudo 或 root 权限 |
| `repo update` | 更新仓库缓存          | 需要 sudo 或 root 权限 |
| `repo create` | 创建本地 YUM/APT 仓库 | 需要 sudo 或 root 权限 |
| `repo cache`  | 从本地仓库创建离线包      | 需要 sudo 或 root 权限 |
| `repo boot`   | 从离线包引导仓库        | 需要 sudo 或 root 权限 |
| `repo reload` | 刷新仓库目录          |                   |
{.full-width}


## 快速入门

```bash
# 方法 1：清理干净现有仓库，添加所有必要仓库并更新缓存（推荐）
pig repo add all --remove --update    # 移除旧仓库，添加所有必要仓库，更新缓存

# 方法 1 变体：一步到位
pig repo set                          # = pig repo add all --remove --update

# 方法 2：温和方式 - 仅添加所需仓库，保留你目前的仓库配置
pig repo add pgsql                    # 添加 PGDG 和 Pigsty 仓库并更新缓存
pig repo add pigsty --region=china    # 添加 Pigsty 仓库，指定使用中国区域
pig repo add pgdg   --region=default  # 添加 PGDG ，指定使用默认区域
pig repo add infra  --region=europe   # 添加 INFRA 仓库 ，指定使用欧洲区域

# 如果上面没有-u|--update 选项一步到位，请额外执行此命令
pig repo update                       # 更新系统包缓存
```


## 模块

在 pig 中，APT/YUM 仓库被组织为 **模块** —— 服务于特定目的的一组仓库。

| 模块 | 说明 | 仓库列表 |
|:---:|:---|:---|
| `all` | 安装 PG 所需的全部核心模块 | `node` + `infra` + `pgsql` |
| `pgsql` | PGDG + Pigsty PG 扩展 | `pigsty-pgsql` + `pgdg` |
| `pigsty` | Pigsty Infra + PGSQL 仓库 | pigsty-infra, pigsty-pgsql |
| `pgdg` | PGDG 官方仓库 | pgdg-common, pgdg13-18 |
| `node` | Linux 系统仓库 | base, updates, extras, epel, baseos, appstream... |
| `infra` | 基础设施组件仓库 | pigsty-infra, nginx, docker-ce |
| `beta` | PostgreSQL 19 Beta 版本 | pgdg19-beta, pgdg-beta |
| `extra` | PGDG Non-Free 与三方扩展 | pgdg-extras, timescaledb, citus |
| `groonga` | PGroonga 仓库 | groonga |
| `mssql` | WiltonDB 仓库 | wiltondb |
| `percona` | Percona PG + PG_TDE | percona |
| `llvm` | LLVM 工具链仓库 | llvm |
| `kube` | Kubernetes 仓库 | kubernetes |
| `grafana` | Grafana 仓库 | grafana |
| `haproxy` | HAProxy 仓库 | haproxyd, haproxyu |
| `redis` | Redis 仓库 | redis |
| `mongo` | MongoDB 仓库 | mongo |
| `mysql` | MySQL 仓库 | mysql |
| `click` | ClickHouse 仓库 | clickhouse |
| `gitlab` | GitLab 仓库 | gitlab-ce, gitlab-ee |
{.full-width}

除此之外，pig 还自带了一些其他数据库的 APT/DNF 仓库：`redis`, `kubernetes`, `grafana`, `clickhouse`, `gitlab`, `haproxy`, `mongodb`, `mysql`，在此不再展开。

通常来说，为了安装 PostgreSQL `node` （Linux 系统仓库） 和 `pgsql`（PGDG + Pigsty）是必选项，`infra` 仓库是可选项（包含了一些工具，IvorySQL Kernel 等）。
您可以使用特殊的 `all` 模块，一次性添加所有需要的仓库到系统中，对绝大多数用户来说，这是合适的起点。

```bash
pig repo add all      # 添加 node,pgsql,infra 三个仓库到系统中
pig repo add          # 不添加任何参数时，默认使用 all 模块
pig repo set          # 使用 set 替代 add 时，将清理备份现有仓库定义并覆盖式更新
```


## 仓库定义

Pigsty 中可用仓库的完整定义位于 [`cli/repo/assets/repo.yml`](https://github.com/pgsty/pig/blob/main/cli/repo/assets/repo.yml)。

您可以创建 `~/.pig/repo.yml` 文件，显式修改并覆盖 pig 的仓库定义。在编辑仓库定义文件时，您可以在 `baseurl` 处添加额外的区域镜像，例如指定中国，欧洲地区的镜像仓库 URL。当 pig 使用 `--region` 参数指定特定的区域时，pig 会优先查找对应区域的仓库 URL，如果不存在，则会 Fallback 到 `default` 的仓库 URL。


## repo list

`pig repo list` 将列出当前系统可用的所有仓库模块。

```bash
pig repo list                # 列出当前系统可用仓库
pig repo list all            # 列出所有仓库（不过滤）
```


## repo info

显示特定仓库或模块的详细信息，包括 URL、元数据和区域镜像，以及 `.repo` / `.list` 仓库文件内容。

```bash
pig repo info pgdg               # 显示 pgdg 模块的信息
pig repo info pigsty pgdg        # 显示多个模块的信息
pig repo info all                # 显示所有模块的信息
```


## repo status

显示系统上的当前仓库配置。

```bash
pig repo status
```


## repo add

添加仓库配置文件到系统。需要 root/sudo 权限。

```bash
pig repo add pgdg                # 添加 PGDG 仓库
pig repo add pgdg pigsty         # 添加多个仓库
pig repo add all                 # 添加所有必要仓库 (pgdg + pigsty + node)
pig repo add pigsty -u           # 添加并更新缓存
pig repo add all -r              # 添加前移除现有仓库
pig repo add all -ru             # 移除、添加并更新（完全重置）
pig repo add pgdg --region=china # 使用中国镜像
```

**选项：**

- `-r|--remove`：添加新仓库前移除现有仓库
- `-u|--update`：添加仓库后运行包缓存更新
- `--region <region>`：使用区域镜像仓库（`default` / `china` / `europe`）

| 平台 | 模块位置 |
|:---:|:---|
| EL | `/etc/yum.repos.d/<module>.repo` |
| Debian | `/etc/apt/sources.list.d/<module>.list` |
{.full-width}


## repo set

等同于 `repo add --remove --update`。清空现有仓库并设置新仓库，然后更新缓存。

```bash
pig repo set                     # 替换为默认仓库
pig repo set pgdg pigsty         # 替换为特定仓库并更新
pig repo set all --region=china  # 使用中国镜像
```


## repo rm

移除仓库配置文件并备份它们。

```bash
pig repo rm                      # 移除所有仓库
pig repo rm pgdg                 # 移除特定仓库
pig repo rm pgdg pigsty -u       # 移除并更新缓存
```

| 平台 | 备份位置 |
|:---:|:---|
| EL | `/etc/yum.repos.d/backup/` |
| Debian | `/etc/apt/sources.list.d/backup/` |
{.full-width}


## repo update

更新包管理器缓存以反映仓库更改。

```bash
pig repo update                  # 更新包缓存
```

| 平台 | 等效命令 |
|:---:|:---|
| EL | `dnf makecache` |
| Debian | `apt update` |
{.full-width}


## repo create

为离线安装创建本地包仓库。

```bash
pig repo create                  # 在默认位置创建 (/www/pigsty)
pig repo create /srv/repo        # 在自定义位置创建
```

| 平台 | 依赖软件 |
|:---:|:---|
| EL | `createrepo_c` |
| Debian | `dpkg-dev` |
{.full-width}


## repo cache

创建仓库内容的压缩 tarball 用于离线分发。

```bash
pig repo cache                   # 默认：/www 到 /tmp/pkg.tgz
pig repo cache -d /srv           # 自定义源目录
```

**选项：**
- `-d, --dir`：源目录（默认：`/www/`）
- `-p, --path`：输出路径（默认：`/tmp/pkg.tgz`）


## repo boot

从离线包解压并设置本地仓库。

```bash
pig repo boot                    # 默认：/tmp/pkg.tgz 到 /www
pig repo boot -p /mnt/pkg.tgz   # 自定义包路径
pig repo boot -d /srv           # 自定义目标目录
```

**选项：**
- `-p, --path`：包路径（默认：`/tmp/pkg.tgz`）
- `-d, --dir`：目标目录（默认：`/www/`）


## repo reload

从 GitHub 刷新仓库元数据到最新版本。

```bash
pig repo reload                  # 刷新仓库目录
```

更新后的文件会放置于 `~/.pig/repo.yml` 中。
