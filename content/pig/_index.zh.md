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
它屏蔽了不同操作系统，不同芯片架构，以及不同 PG 大版本的管理差异，让您用简单的几行命令，就可以完成 PG 内核与 444+ 扩展的安装与管理。

请注意：对于扩展安装来说，**pig 并非必须组件**，您依然可以使用 apt / dnf 等包管理器直接访问 [**Pigsty PGSQL**](/zh/repo/pgsql) 仓库。

{{< cards >}}
{{< card link="/zh/pig/intro"   title="简介" subtitle="为什么需要专用的PG包管理器？" icon="sparkles" >}}
{{< card link="/zh/pig/start"   title="上手" subtitle="快速上手与样例"  icon="play" >}}
{{< card link="/zh/pig/install" title="安装" subtitle="下载、安装、更新 pig" icon="save" >}}
{{< /cards >}}



## 快速上手

使用以下命令即可在您的系统上 [**安装**](/zh/pig/install) PIG 包管理器：

{{< tabs items="默认,镜像" defaultIndex="1" >}}
{{< tab >}}
```bash
curl -fsSL https://repo.pigsty.io/pig | bash     # 从 Cloudflare 安装
```
{{< /tab >}}
{{< tab >}}
```bash
curl -fsSL https://repo.pigsty.cc/pig | bash     # 从中国 CDN 镜像站安装
```
{{< /tab >}}
{{< /tabs >}}

安装完成后，几行命令即可 [**快速开始**](/zh/pig/start) 。例如，若需安装 PG 18 与相应的 [**`pg_duckdb`**](/e/pg_duckdb/) 扩展：

```bash
$ pig repo set                        # 一次性设置好 Linux, Pigsty + PGDG 仓库（覆盖式！）
$ pig install pg18                    # 安装 PostgreSQL 18 内核（原生 PGDG 包）
$ pig install pg_duckdb -v 18         # 安装 pg_duckdb 扩展（针对当前 pg 18）
$ pig install -y postgis timescaledb  # 针对当前活跃PG版本，安装多个扩展
$ pig install -y vector               # 您可以使用扩展名称（vector）或者扩展包名称（pgvector）来安装扩展！
```



## 命令参考

你可以执行 `pig help <command>` 获取子命令的详细帮助。

{{< cards cols="5" >}}
{{< card link="/zh/pig/cmd/repo"  title="pig repo"  subtitle="管理软件仓库"  icon="library" >}}
{{< card link="/zh/pig/cmd/ext"   title="pig ext"   subtitle="管理PG扩展"   icon="cube" >}}
{{< card link="/zh/pig/cmd/build" title="pig build" subtitle="设置构建环境"  icon="view-grid" >}}
{{< card link="/zh/pig/cmd/sty"   title="pig sty"   subtitle="管理 Pigsty"  icon="cloud-download" >}}
{{< /cards >}}


## 关于

`pig` 命令行工具由 [Vonng](https://vonng.com/en/)（冯若航 rh@vonng.com）开发，并以 [Apache 2.0](https://github.com/pgsty/pig/?tab=Apache-2.0-1-ov-file#readme) 许可证开源。

您还可以参考 [**PIGSTY**](https://pgsty.com) 项目，提供了包括扩展交付在内的完整 PostgreSQL RDS DBaaS 使用体验。

{{< cards cols=4 >}}
{{< card link="https://github.com/github.com/pgsty/pgext"  title="PGEXT"  icon="github" subtitle="本网站，扩展数据与管理工具" >}}
{{< card link="https://github.com/github.com/pgsty/pig"    title="PIG"    icon="github" subtitle="PostgreSQL 包管理器" >}}
{{< card link="https://github.com/github.com/pgsty/pigsty" title="PIGSTY" icon="github" subtitle="开箱即用的 PostgreSQL 发行版" >}}
{{< /cards >}}

