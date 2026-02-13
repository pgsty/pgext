---
title: PIG 包管理器
description: PostgreSQL 扩展包管理器
icon: PiggyBank
full: true
cascade:
  type: docs
breadcrumbs: false
comments: false
weight: 100
---


—— **Postgres Install Genius，PostgreSQL 生态中缺失的扩展包管理器**

PIG 包管理器是一个专门用于安装、管理、构建 PostgreSQL 及其扩展的命令行工具，使用 Go 开发，开箱即用，简单易用，小巧玲珑（4MB）。
PIG 包管理器并非重新发明的土鳖轮子，而是 **依托** （PiggyBack）现有 Linux 发行版包管理器 （`apt`/`dnf`）的一个高级抽象层。
它屏蔽了不同操作系统，不同芯片架构，以及不同 PG 大版本的管理差异，让您用简单的几行命令，就可以完成 PG 内核与 451 扩展的安装与管理。

PIG 的命令设计同样适合自动化脚本：提供统一的参数风格、清晰的错误提示，以及如 `--dry-run` 等安全开关。

请注意：对于扩展安装来说，**pig 并非必须组件**，您依然可以使用 apt / dnf 等包管理器直接访问 [**Pigsty PGSQL**](/zh/repo/pgsql/) 仓库。

- [**简介**](/zh/pig/intro)：为什么需要专用的 PG 包管理器？
- [**上手**](/zh/pig/start)：快速上手与样例
- [**安装**](/zh/pig/install)：下载、安装、更新 pig


## 快速上手

使用以下命令即可在您的系统上 [**安装**](/zh/pig/install) PIG 包管理器：

**默认安装**（Cloudflare CDN）：

```bash
curl -fsSL https://repo.pigsty.io/pig | bash
```

**中国镜像**：

```bash
curl -fsSL https://repo.pigsty.cc/pig | bash
```

安装完成后，几行命令即可 [**快速开始**](/zh/pig/start) 。例如，若需安装 PG 18 与相应的 [**`pg_duckdb`**](https://pgext.cloud/e/pg_duckdb) 扩展：

```bash
$ pig repo set                        # 一次性设置好 Linux, Pigsty + PGDG 仓库（覆盖式！）
$ pig install pg18                    # 安装 PostgreSQL 18 内核（原生 PGDG 包）
$ pig install pg_duckdb -v 18         # 安装 pg_duckdb 扩展（针对当前 pg 18）
$ pig install -y postgis timescaledb  # 针对当前活跃PG版本，安装多个扩展
$ pig install -y vector               # 您可以使用扩展名称（vector）或者扩展包名称（pgvector）来安装扩展！
```


## 命令参考

你可以执行 `pig help <command>` 获取子命令的详细帮助。

**扩展管理**：

- [**pig repo**](/zh/pig/cmd/repo)：管理 APT/YUM 软件仓库
- [**pig ext**](/zh/pig/cmd/ext)：管理 PostgreSQL 扩展
- [**pig build**](/zh/pig/cmd/build)：从源码构建扩展

**Pigsty 管理**：

- [**pig sty**](/zh/pig/cmd/sty)：管理 Pigsty 安装
- [**pig postgres**](/zh/pig/cmd/pg)：管理本地 PostgreSQL 服务
- [**pig patroni**](/zh/pig/cmd/pt)：管理 Patroni HA 集群
- [**pig pgbackrest**](/zh/pig/cmd/pb)：管理 pgBackRest 备份
- [**pig pitr**](/zh/pig/cmd/pitr)：编排式时间点恢复


## 关于

`pig` 命令行工具由 [Vonng](https://vonng.com/en/)（冯若航 rh@vonng.com）开发，并以 [Apache 2.0](https://github.com/pgsty/pig/?tab=Apache-2.0-1-ov-file#readme) 许可证开源。

您还可以参考 [**PIGSTY**](https://pgsty.com) 项目，提供了包括扩展交付在内的完整 PostgreSQL RDS DBaaS 使用体验。

- [**PGEXT**](https://github.com/pgsty/pgext)：扩展数据与管理工具
- [**PIG**](https://github.com/pgsty/pig)：PostgreSQL 包管理器
- [**PIGSTY**](https://github.com/pgsty/pigsty)：开箱即用的 PostgreSQL 发行版
