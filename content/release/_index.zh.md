---
title: 发布记录
description: PIG CLI 发布历史与仓库软件包的变更记录
weight: 800
---


{{< cards cols=2 >}}
{{< card link="/zh/release/pig/"                         title="PIG 发布记录"    subtitle="PG 包管理器发布历史"          icon="sparkles" />}}
{{< card link="https://pigsty.io/docs/about/release" title="Pigsty 发布记录" subtitle="开箱即用的 PostgreSQL 发行版" icon="play" />}}
{{< /cards >}}


{{< cards >}}
{{< card link="/zh/release/rpm"   title="RPM 发布记录"   subtitle="EL RPM 软件包变更记录" icon="sparkles" />}}
{{< card link="/zh/release/deb"   title="DEB 发布记录"   subtitle="Debian 软件包变更记录" icon="play" />}}
{{< card link="/zh/release/infra" title="Infra 发布记录" subtitle="与OS版本无关的通用软件包变更记录"  icon="save" />}}
{{< /cards >}}


## PIG 版本

最新稳定版本是 [v1.10.0](https://github.com/pgsty/pig/releases/tag/v1.10.0)。

|       版本        |     日期     | 摘要                                              |                                                     GitHub |
|:---------------:|:----------:|-------------------------------------------------|-----------------------------------------------------------:|
| [v1.10.0](/zh/release/pig/#v1100) | 2026-10-17 | 603 条目录记录，PostgreSQL 目标与恢复安全、Inventory 保真、自动化修复 | [v1.10.0](https://github.com/pgsty/pig/releases/tag/v1.10.0) |
| [v1.9.0](/zh/release/pig/#v190) | 2026-10-06 | 600 条扩展记录、Xray 角色、公钥准备、pgrx 0.19.3 | [v1.9.0](https://github.com/pgsty/pig/releases/tag/v1.9.0) |
| [v1.8.1](/zh/release/pig/#v181) | 2026-09-03 | CLI 与仓库安全、默认调试包、pgrx 0.19.2 | [v1.8.1](https://github.com/pgsty/pig/releases/tag/v1.8.1) |
| [v1.8.0](/zh/release/pig/#v180) | 2026-08-14 | 原生 `sty boot` 与 `sty conf`，575 个已打包扩展           | [v1.8.0](https://github.com/pgsty/pig/releases/tag/v1.8.0) |
| [v1.7.0](/zh/release/pig/#v170) | 2026-08-12 | EL 模块策略、中国镜像与 EL7 兼容目录，575 个扩展                  | [v1.7.0](https://github.com/pgsty/pig/releases/tag/v1.7.0) |
| [v1.6.2](/zh/release/pig/#v162) | 2026-08-11 | 572 个扩展，Grafana schema v2，SOW 优先的软件仓库生成         | [v1.6.2](https://github.com/pgsty/pig/releases/tag/v1.6.2) |
| [v1.6.1](/zh/release/pig/#v161) | 2026-07-30 | 扩展目录刷新，内置 Pigsty 版本对齐到 4.5.0                    | [v1.6.1](https://github.com/pgsty/pig/releases/tag/v1.6.1) |
| [v1.6.0](/zh/release/pig/#v160) | 2026-07-28 | 562 个已打包扩展，pt 原生透传，inventory 与 CMDB，Grafana 管理  | [v1.6.0](https://github.com/pgsty/pig/releases/tag/v1.6.0) |
| [v1.5.1](/zh/release/pig/#v151) | 2026-07-08 | PG 内核分支包更新，镜像模式，修复若干问题                          | [v1.5.1](https://github.com/pgsty/pig/releases/tag/v1.5.1) |
| [v1.5.0](/zh/release/pig/#v150) | 2026-07-04 | 531 个扩展，pigsty v4.4，pg/pb/pt/pitr 重做，clone/fork | [v1.5.0](https://github.com/pgsty/pig/releases/tag/v1.5.0) |
| [v1.4.2](/zh/release/pig/#v142) | 2026-06-18 | 524 个扩展，PG19 beta，pgrx 0.18.1，Patroni 修复        | [v1.4.2](https://github.com/pgsty/pig/releases/tag/v1.4.2) |
| [v1.4.1](/zh/release/pig/#v141) | 2026-05-01 | 510 个扩展，支持 Ubuntu 26.04，仓库校准                    | [v1.4.1](https://github.com/pgsty/pig/releases/tag/v1.4.1) |
| [v1.4.0](/zh/release/pig/#v140) | 2026-04-19 | 510 个扩展，pgrx 0.18.0，更多构建规格                      | [v1.4.0](https://github.com/pgsty/pig/releases/tag/v1.4.0) |
| [v1.3.4](/zh/release/pig/#v134) | 2026-04-14 | 504 扩展更新与发布产物校验和刷新                              | [v1.3.4](https://github.com/pgsty/pig/releases/tag/v1.3.4) |
| [v1.3.3](/zh/release/pig/#v133) | 2026-04-10 | 481 扩展与 Go 1.26.2 更新                            | [v1.3.3](https://github.com/pgsty/pig/releases/tag/v1.3.3) |
| [v1.3.2](/zh/release/pig/#v132) | 2026-03-23 | 例行元数据更新，新增 `pg tune` 与构建别名                      | [v1.3.2](https://github.com/pgsty/pig/releases/tag/v1.3.2) |
| [v1.3.1](/zh/release/pig/#v131) | 2026-03-05 | PG13 退役，支持窗口统一为 PG14-18，扩展增至 464                | [v1.3.1](https://github.com/pgsty/pig/releases/tag/v1.3.1) |
| [v1.3.0](/zh/release/pig/#v130) | 2026-02-27 | 构建链路强化，扩展增至 461，新内核支持                           | [v1.3.0](https://github.com/pgsty/pig/releases/tag/v1.3.0) |
| [v1.2.0](/zh/release/pig/#v120) | 2026-02-23 | 统一别名，例行更新，计划模式，仓库修复                             | [v1.2.0](https://github.com/pgsty/pig/releases/tag/v1.2.0) |
| [v1.1.1](https://pig.pgsty.com/zh/release/pig-1.1.1/) | 2026-02-14 | 修复路径、符号链接与构建目录迁移问题 | [v1.1.1](https://github.com/pgsty/pig/releases/tag/v1.1.1) |
| [v1.1.0](/zh/release/pig/#v110) | 2026-02-12 | 451 扩展，Agent-Native CLI 框架                      | [v1.1.0](https://github.com/pgsty/pig/releases/tag/v1.1.0) |
| [v1.0.0](/zh/release/pig/#v100) | 2026-01-26 | 444, 新增 pg/pt/pb/pitr 子命令，可用性矩阵                 | [v1.0.0](https://github.com/pgsty/pig/releases/tag/v1.0.0) |
| [v0.9.0](/zh/release/pig/#v090) | 2025-12-28 | 重命名 sty deploy 并扩展 sty conf 参数 | [v0.9.0](https://github.com/pgsty/pig/releases/tag/v0.9.0) |
| [v0.8.1](https://pig.pgsty.com/zh/release/pig-0.8.1/) | 2025-12-27 | 将 llvmjit 恢复到 pgsql-full 别名 | [v0.8.1](https://github.com/pgsty/pig/releases/tag/v0.8.1) |
| [v0.8.0](/zh/release/pig/#v080) | 2025-12-26 | 440 extensions，移除 sysupdate 仓库                  | [v0.8.0](https://github.com/pgsty/pig/releases/tag/v0.8.0) |
| [v0.7.5](/zh/release/pig/#v075) | 2025-12-12 | 常规扩展更新，使用修复后的阿里云镜像                              | [v0.7.5](https://github.com/pgsty/pig/releases/tag/v0.7.5) |
| [v0.7.4](/zh/release/pig/#v074) | 2025-12-01 | 更新 ivory/pgtde 内核与 pgdg extras 仓库               | [v0.7.4](https://github.com/pgsty/pig/releases/tag/v0.7.4) |
| [v0.7.3](/zh/release/pig/#v073) | 2025-11-24 | 修复 el10 & debian13 仓库配置                         | [v0.7.3](https://github.com/pgsty/pig/releases/tag/v0.7.3) |
| [v0.7.2](/zh/release/pig/#v072) | 2025-11-20 | 437 个扩展，修复 pig build 的一些问题                      | [v0.7.2](https://github.com/pgsty/pig/releases/tag/v0.7.2) |
| [v0.7.1](/zh/release/pig/#v071) | 2025-11-10 | 新网站，改进容器内的使用体验                                  | [v0.7.1](https://github.com/pgsty/pig/releases/tag/v0.7.1) |
| [v0.7.0](/zh/release/pig/#v070) | 2025-11-05 | 强化 build 能力，大批量包更新                              | [v0.7.0](https://github.com/pgsty/pig/releases/tag/v0.7.0) |
| [v0.6.2](/zh/release/pig/#v062) | 2025-10-03 | 正式提供 PG 18 支持                                   | [v0.6.2](https://github.com/pgsty/pig/releases/tag/v0.6.2) |
| [v0.6.1](/zh/release/pig/#v061) | 2025-08-14 | CI/CD, el10 存根，PGDG 中国镜像                        | [v0.6.1](https://github.com/pgsty/pig/releases/tag/v0.6.1) |
| [v0.6.0](/zh/release/pig/#v060) | 2025-07-17 | 423 个扩展，percona pg_tde，mcp 工具箱                  | [v0.6.0](https://github.com/pgsty/pig/releases/tag/v0.6.0) |
| [v0.5.0](/zh/release/pig/#v050) | 2025-06-30 | 422 个扩展，新的扩展目录                                  | [v0.5.0](https://github.com/pgsty/pig/releases/tag/v0.5.0) |
| [v0.4.2](/zh/release/pig/#v042) | 2025-05-27 | 421 个扩展，halo 和 oriole deb                       | [v0.4.2](https://github.com/pgsty/pig/releases/tag/v0.4.2) |
| [v0.4.1](/zh/release/pig/#v041) | 2025-05-07 | 414 个扩展，pg18 别名支持                               | [v0.4.1](https://github.com/pgsty/pig/releases/tag/v0.4.1) |
| [v0.4.0](/zh/release/pig/#v040) | 2025-05-01 | do 和 pt 子命令，halo 和 orioledb                     | [v0.4.0](https://github.com/pgsty/pig/releases/tag/v0.4.0) |
| [v0.3.4](/zh/release/pig/#v034) | 2025-04-05 | 常规更新                                            | [v0.3.4](https://github.com/pgsty/pig/releases/tag/v0.3.4) |
| [v0.3.3](/zh/release/pig/#v033) | 2025-03-25 | 别名、仓库、依赖                                        | [v0.3.3](https://github.com/pgsty/pig/releases/tag/v0.3.3) |
| [v0.3.2](/zh/release/pig/#v032) | 2025-03-21 | 新扩展                                             | [v0.3.2](https://github.com/pgsty/pig/releases/tag/v0.3.2) |
| [v0.3.1](/zh/release/pig/#v031) | 2025-03-19 | 轻微错误修复                                          | [v0.3.1](https://github.com/pgsty/pig/releases/tag/v0.3.1) |
| [v0.3.0](/zh/release/pig/#v030) | 2025-02-24 | 新主页和扩展目录                                        | [v0.3.0](https://github.com/pgsty/pig/releases/tag/v0.3.0) |
| [v0.2.2](/zh/release/pig/#v022) | 2025-02-22 | 404 个扩展                                         | [v0.2.2](https://github.com/pgsty/pig/releases/tag/v0.2.2) |
| [v0.2.0](/zh/release/pig/#v020) | 2025-02-14 | 400 个扩展                                         | [v0.2.0](https://github.com/pgsty/pig/releases/tag/v0.2.0) |
| [v0.1.4](/zh/release/pig/#v014) | 2025-02-12 | 常规错误修复                                          | [v0.1.4](https://github.com/pgsty/pig/releases/tag/v0.1.4) |
| [v0.1.3](/zh/release/pig/#v013) | 2025-01-23 | 390 个扩展                                         | [v0.1.3](https://github.com/pgsty/pig/releases/tag/v0.1.3) |
| [v0.1.2](/zh/release/pig/#v012) | 2025-01-12 | anon 扩展和其他 350 个扩展                              | [v0.1.2](https://github.com/pgsty/pig/releases/tag/v0.1.2) |
| [v0.1.1](https://pig.pgsty.com/zh/release/pig-0.1.1/) | 2025-01-09 | 更新扩展列表                                          | [v0.1.1](https://github.com/pgsty/pig/releases/tag/v0.1.1) |
| [v0.1.0](/zh/release/pig/#v010) | 2024-12-29 | repo、ext、sty 和自更新                               | [v0.1.0](https://github.com/pgsty/pig/releases/tag/v0.1.0) |
| [v0.0.1](/zh/release/pig/#v001) | 2024-12-23 | 创世发布                                            | [v0.0.1](https://github.com/pgsty/pig/releases/tag/v0.0.1) |
