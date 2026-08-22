---
title: 关于
linkTitle: 关于
description: PGEXT.CLOUD 是什么，由谁维护，源码在哪里
weight: 950
icon: fa-solid fa-circle-info
sidebar_enabled: false
breadcrumb: false
---

**PGEXT.CLOUD** 是一份 PostgreSQL 扩展生态的目录，也是把其中相当一部分构建成原生 Linux 软件包的那个仓库的门面。

它之所以存在，是因为这个生态最大的优点 —— 任何人都能扩展数据库 —— 同时也是它最难用的地方。扩展散落在 GitHub、PGXN、厂商站点、高校的 FTP 目录，以及十年前某封邮件列表帖子的附件里。搞清楚「有没有这个扩展」「它还能不能编译」「运行时需要什么」「有没有人为我的发行版打过包」，过去平均每个扩展要花掉一个下午。

## 它提供什么

{{< cards cols=3 >}}
{{< card link="/list" title="扩展目录" icon="clipboard-list" subtitle="收录 2241 个扩展，按分类、语言与许可证建立索引" />}}
{{< card link="/repo" title="软件仓库" icon="warehouse" subtitle="其中 576 个已构建为签名 RPM / DEB，覆盖 16 个 Linux 目标" />}}
{{< card link="/pig" title="包管理器" icon="cash" subtitle="一个命令行工具，用扩展本名安装内核与扩展" />}}
{{< /cards >}}

这三层可以各自独立使用。即便你从不从这个仓库安装任何东西，目录本身也值得一读；而如果你不想再多装一个命令行工具，用原生的 `apt` 或 `dnf` 同样可以使用这个仓库。

## 设计承诺

- **兼容 PGDG**：PostgreSQL 内核直接来自官方
  [PGDG](https://www.postgresql.org/download/linux/) 仓库，未经改动。
  这里的软件包遵循 PGDG 的命名、路径与约定，二者是互补而非竞争关系。
- **Linux 原生**：真正的 RPM 与 DEB 软件包，带真实依赖关系 ——
  不是私有包格式，不是容器镜像，也不是一段构建脚本。
- **可复现**：每个软件包都由公开的构建规范
  （[RPM](https://github.com/pgsty/rpm)、[DEB](https://github.com/pgsty/deb)）
  针对锁定的源码修订版本构建而成。
- **双架构对等**：`x86_64` 与 `aarch64` 按同等标准构建，
  [可用性矩阵](/os/matrix) 会精确标出二者的差异。
- **免费**：目录、构建规范、命令行工具与软件仓库全部开源免费，
  无需注册，没有配额。

## 维护者

PGEXT.CLOUD 由 [**PGSTY**](https://github.com/pgsty) / [**Vonng**](https://vonng.com/) (rh@vonng.com) 开发维护，以 [**Apache 2.0**](https://github.com/pgsty/pig/?tab=Apache-2.0-1-ov-file#readme) 许可证开源。

| GitHub 仓库 | 内容 |
|:------------|:-----|
| [github.com/pgsty](https://github.com/pgsty) | **PGSTY** 组织主页 |
| [github.com/pgsty/pgext](https://github.com/pgsty/pgext) | 本网站、扩展元数据，以及生成两者的工具 |
| [github.com/pgsty/pigsty](https://github.com/pgsty/pigsty) | 构建于其上的 PostgreSQL 发行版 |
| [github.com/pgsty/pig](https://github.com/pgsty/pig) | `pig` 包管理器 |
| [github.com/pgsty/rpm](https://github.com/pgsty/rpm) | RPM 构建规范 |
| [github.com/pgsty/deb](https://github.com/pgsty/deb) | DEB 构建规范 |
| [github.com/pgsty/infra-pkg](https://github.com/pgsty/infra-pkg) | 基础设施软件包构建规范 |

## 参与贡献

最有用的贡献是一处订正。如果发现扩展缺失、版本过期、许可证标错，或者某个软件包装不上，请到 [pgsty/pgext](https://github.com/pgsty/pgext/issues) 提一个 issue —— 那 576 个已打包扩展里，有相当一部分正是这么来的。

延伸阅读：[***PostgreSQL 正在吞噬数据库世界***](https://pigsty.cc/blog/db/pg-is-eating-db-world/)。
