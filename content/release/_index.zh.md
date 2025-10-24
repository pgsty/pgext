---
title: 变更
description: PIG CLI 发布历史与仓库软件包的变更记录
weight: 800
---


{{< cards cols=2 >}}
{{< card link="/release/pig"                    title="PIG 发布记录"    subtitle="PG 包管理器发布历史"          icon="sparkles" >}}
{{< card link="https://doc.pgsty.com/release"   title="Pigsty 发布记录" subtitle="开箱即用的 PostgreSQL 发行版" icon="play"     >}}
{{< /cards >}}


{{< cards >}}
{{< card link="/release/rpm"   title="RPM 发布记录"   subtitle="EL RPM 软件包变更记录" icon="sparkles" >}}
{{< card link="/release/deb"   title="DEB 发布记录"   subtitle="Debian 软件包变更记录" icon="play" >}}
{{< card link="/release/infra" title="Infra 发布记录" subtitle="与OS版本无关的通用软件包变更记录"  icon="save" >}}
{{< /cards >}}


## PIG 发布历史

当前 `pig` 最新稳定版本为：[v0.6.2](https://github.com/pgsty/pig/releases/tag/v0.6.2)。

|       版本        |     日期     | 摘要                             |                           GitHub                           |
|:---------------:|:----------:|--------------------------------|:----------------------------------------------------------:|
| [v0.6.2](#v062) | 2025-10-03 | 正式提供 PG 18 支持                  | [v0.6.2](https://github.com/pgsty/pig/releases/tag/v0.6.2) |
| [v0.6.1](#v061) | 2025-08-13 | 添加 CI/CD 管道，使用 PIGSTY PGDG 仓库  | [v0.6.1](https://github.com/pgsty/pig/releases/tag/v0.6.1) |
| [v0.6.0](#v060) | 2025-07-17 | 423 个扩展，percona pg_tde，mcp 工具箱 | [v0.6.0](https://github.com/pgsty/pig/releases/tag/v0.6.0) |
| [v0.5.0](#v050) | 2025-06-30 | 422 个扩展，新的扩展目录                 | [v0.6.0](https://github.com/pgsty/pig/releases/tag/v0.5.0) |
| [v0.4.2](#v042) | 2025-05-27 | 421 个扩展，halo 和 oriole deb      | [v0.4.2](https://github.com/pgsty/pig/releases/tag/v0.4.2) |
| [v0.4.1](#v041) | 2025-05-07 | 414 个扩展，pg18 别名支持              | [v0.4.1](https://github.com/pgsty/pig/releases/tag/v0.4.1) |
| [v0.4.0](#v040) | 2025-05-01 | do 和 pt 子命令，halo 和 orioledb    | [v0.4.0](https://github.com/pgsty/pig/releases/tag/v0.4.0) |
| [v0.3.4](#v034) | 2025-04-05 | 常规更新                           | [v0.3.4](https://github.com/pgsty/pig/releases/tag/v0.3.4) |
| [v0.3.3](#v033) | 2025-03-25 | 别名、仓库、依赖                       | [v0.3.3](https://github.com/pgsty/pig/releases/tag/v0.3.3) |
| [v0.3.2](#v032) | 2025-03-21 | 新扩展                            | [v0.3.2](https://github.com/pgsty/pig/releases/tag/v0.3.2) |
| [v0.3.1](#v031) | 2025-03-19 | 轻微错误修复                         | [v0.3.1](https://github.com/pgsty/pig/releases/tag/v0.3.1) |
| [v0.3.0](#v030) | 2025-02-24 | 新主页和扩展目录                       | [v0.3.0](https://github.com/pgsty/pig/releases/tag/v0.3.0) |
| [v0.2.2](#v022) | 2025-02-22 | 404 个扩展                        | [v0.2.2](https://github.com/pgsty/pig/releases/tag/v0.2.2) |
| [v0.2.0](#v020) | 2025-02-14 | 400 个扩展                        | [v0.2.0](https://github.com/pgsty/pig/releases/tag/v0.2.0) |
| [v0.1.4](#v014) | 2025-02-12 | 常规错误修复                         | [v0.1.4](https://github.com/pgsty/pig/releases/tag/v0.1.4) |
| [v0.1.3](#v013) | 2025-01-23 | 390 个扩展                        | [v0.1.3](https://github.com/pgsty/pig/releases/tag/v0.1.3) |
| [v0.1.2](#v012) | 2025-01-12 | anon 扩展和其他 350 个扩展             | [v0.1.2](https://github.com/pgsty/pig/releases/tag/v0.1.2) |
| [v0.1.1](#v011) | 2025-01-09 | 更新扩展列表                         | [v0.1.1](https://github.com/pgsty/pig/releases/tag/v0.1.1) |
| [v0.1.0](#v010) | 2024-12-29 | repo、ext、sty 和自更新              | [v0.1.0](https://github.com/pgsty/pig/releases/tag/v0.1.0) |
| [v0.0.1](#v001) | 2024-12-23 | 创世发布                           | [v0.0.1](https://github.com/pgsty/pig/releases/tag/v0.0.1) |


