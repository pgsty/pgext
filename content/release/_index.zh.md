---
title: 发布记录
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

当前 `pig` 最新稳定版本为：[v0.7.4](https://github.com/pgsty/pig/releases/tag/v0.7.4)。

| 版本                             |     日期     | 摘要                                |                           GitHub                           |
|:-------------------------------|:----------:|-----------------------------------|:----------------------------------------------------------:|
| [v0.7.4](/zh/release/pig#v074) | 2025-12-01 | 更新 ivory/pgtde 内核与 pgdg extras 仓库 | [v0.7.4](https://github.com/pgsty/pig/releases/tag/v0.7.4) |
| [v0.7.3](/zh/release/pig#v073) | 2025-11-24 | 修复 el10 & debian13 仓库配置           | [v0.7.3](https://github.com/pgsty/pig/releases/tag/v0.7.3) |
| [v0.7.2](/zh/release/pig#v072) | 2025-11-20 | 437 个扩展，修复 pig build 的一些问题        | [v0.7.2](https://github.com/pgsty/pig/releases/tag/v0.7.2) |
| [v0.7.1](/zh/release/pig#v071) | 2025-11-10 | 新网站，改进容器内的使用体验                    | [v0.7.1](https://github.com/pgsty/pig/releases/tag/v0.7.1) |
| [v0.7.0](/zh/release/pig#v070) | 2025-11-05 | 强化 build 能力，大批量包更新                | [v0.7.0](https://github.com/pgsty/pig/releases/tag/v0.7.0) |
| [v0.6.2](/zh/release/pig#v062) | 2025-10-03 | 正式提供 PG 18 支持                     | [v0.6.2](https://github.com/pgsty/pig/releases/tag/v0.6.2) |
| [v0.6.1](/zh/release/pig#v061) | 2025-08-13 | 添加 CI/CD 管道，使用 PIGSTY PGDG 仓库     | [v0.6.1](https://github.com/pgsty/pig/releases/tag/v0.6.1) |
| [v0.6.0](/zh/release/pig#v060) | 2025-07-17 | 423 个扩展，percona pg_tde，mcp 工具箱    | [v0.6.0](https://github.com/pgsty/pig/releases/tag/v0.6.0) |
| [v0.5.0](/zh/release/pig#v050) | 2025-06-30 | 422 个扩展，新的扩展目录                    | [v0.6.0](https://github.com/pgsty/pig/releases/tag/v0.5.0) |
| [v0.4.2](/zh/release/pig#v042) | 2025-05-27 | 421 个扩展，halo 和 oriole deb         | [v0.4.2](https://github.com/pgsty/pig/releases/tag/v0.4.2) |
| [v0.4.1](/zh/release/pig#v041) | 2025-05-07 | 414 个扩展，pg18 别名支持                 | [v0.4.1](https://github.com/pgsty/pig/releases/tag/v0.4.1) |
| [v0.4.0](/zh/release/pig#v040) | 2025-05-01 | do 和 pt 子命令，halo 和 orioledb       | [v0.4.0](https://github.com/pgsty/pig/releases/tag/v0.4.0) |
| [v0.3.4](/zh/release/pig#v034) | 2025-04-05 | 常规更新                              | [v0.3.4](https://github.com/pgsty/pig/releases/tag/v0.3.4) |
| [v0.3.3](/zh/release/pig#v033) | 2025-03-25 | 别名、仓库、依赖                          | [v0.3.3](https://github.com/pgsty/pig/releases/tag/v0.3.3) |
| [v0.3.2](/zh/release/pig#v032) | 2025-03-21 | 新扩展                               | [v0.3.2](https://github.com/pgsty/pig/releases/tag/v0.3.2) |
| [v0.3.1](/zh/release/pig#v031) | 2025-03-19 | 轻微错误修复                            | [v0.3.1](https://github.com/pgsty/pig/releases/tag/v0.3.1) |
| [v0.3.0](/zh/release/pig#v030) | 2025-02-24 | 新主页和扩展目录                          | [v0.3.0](https://github.com/pgsty/pig/releases/tag/v0.3.0) |
| [v0.2.2](/zh/release/pig#v022) | 2025-02-22 | 404 个扩展                           | [v0.2.2](https://github.com/pgsty/pig/releases/tag/v0.2.2) |
| [v0.2.0](/zh/release/pig#v020) | 2025-02-14 | 400 个扩展                           | [v0.2.0](https://github.com/pgsty/pig/releases/tag/v0.2.0) |
| [v0.1.4](/zh/release/pig#v014) | 2025-02-12 | 常规错误修复                            | [v0.1.4](https://github.com/pgsty/pig/releases/tag/v0.1.4) |
| [v0.1.3](/zh/release/pig#v013) | 2025-01-23 | 390 个扩展                           | [v0.1.3](https://github.com/pgsty/pig/releases/tag/v0.1.3) |
| [v0.1.2](/zh/release/pig#v012) | 2025-01-12 | anon 扩展和其他 350 个扩展                | [v0.1.2](https://github.com/pgsty/pig/releases/tag/v0.1.2) |
| [v0.1.1](/zh/release/pig#v011) | 2025-01-09 | 更新扩展列表                            | [v0.1.1](https://github.com/pgsty/pig/releases/tag/v0.1.1) |
| [v0.1.0](/zh/release/pig#v010) | 2024-12-29 | repo、ext、sty 和自更新                 | [v0.1.0](https://github.com/pgsty/pig/releases/tag/v0.1.0) |
| [v0.0.1](/zh/release/pig#v001) | 2024-12-23 | 创世发布                              | [v0.0.1](https://github.com/pgsty/pig/releases/tag/v0.0.1) |


