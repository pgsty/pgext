---
title: Release
description: Changelog of PIG CLI and Repo packages
weight: 800
---


{{< cards cols=2 >}}
{{< card link="/release/pig/"                         title="PIG Release"    subtitle="Changelog of the PG Package Manager"      icon="sparkles" />}}
{{< card link="https://pigsty.io/docs/about/release" title="Pigsty Release" subtitle="Battery-Included PostgreSQL Distribution" icon="play" />}}
{{< /cards >}}


{{< cards >}}
{{< card link="/release/rpm"   title="RPM Changelog"   subtitle="Changelog of EL RPM Packages" icon="sparkles" />}}
{{< card link="/release/deb"   title="DEB Changelog"   subtitle="Changelog of Debian Packages" icon="play" />}}
{{< card link="/release/infra" title="Infra Changelog" subtitle="Changelog of OS-Generic Packages"  icon="save" />}}
{{< /cards >}}


## PIG Releases

The latest stable version is [v1.10.0](https://github.com/pgsty/pig/releases/tag/v1.10.0).

|     Version     |    Date    | Summary                                                            |                           GitHub                           |
|:---------------:|:----------:|--------------------------------------------------------------------|:----------------------------------------------------------:|
| [v1.10.0](/release/pig/#v1100) | 2026-10-17 | 603 catalog entries, PostgreSQL target and recovery safety, faithful Inventory, automation fixes | [v1.10.0](https://github.com/pgsty/pig/releases/tag/v1.10.0) |
| [v1.9.0](/release/pig/#v190) | 2026-10-06 | 600 catalog entries, Xray roles, repository keys, pgrx 0.19.3 | [v1.9.0](https://github.com/pgsty/pig/releases/tag/v1.9.0) |
| [v1.8.1](/release/pig/#v181) | 2026-09-03 | CLI/repository safety, debug packages by default, pgrx 0.19.2 | [v1.8.1](https://github.com/pgsty/pig/releases/tag/v1.8.1) |
| [v1.8.0](/release/pig/#v180) | 2026-08-14 | Native `sty boot` and `sty conf`, 575 packaged extensions          | [v1.8.0](https://github.com/pgsty/pig/releases/tag/v1.8.0) |
| [v1.7.0](/release/pig/#v170) | 2026-08-12 | EL module policy, China mirrors, EL7 compatibility, 575 extensions | [v1.7.0](https://github.com/pgsty/pig/releases/tag/v1.7.0) |
| [v1.6.2](/release/pig/#v162) | 2026-08-11 | 572 extensions, Grafana schema v2, SOW-first repository generation | [v1.6.2](https://github.com/pgsty/pig/releases/tag/v1.6.2) |
| [v1.6.1](/release/pig/#v161) | 2026-07-30 | Catalog refresh and embedded Pigsty 4.5.0                           | [v1.6.1](https://github.com/pgsty/pig/releases/tag/v1.6.1) |
| [v1.6.0](/release/pig/#v160) | 2026-07-28 | 562 packaged extensions, patronictl passthrough, inventory & CMDB, Grafana | [v1.6.0](https://github.com/pgsty/pig/releases/tag/v1.6.0) |
| [v1.5.1](/release/pig/#v151) | 2026-07-08 | PG kernel fork updates, mirror mode, bug fixes                    | [v1.5.1](https://github.com/pgsty/pig/releases/tag/v1.5.1) |
| [v1.5.0](/release/pig/#v150) | 2026-07-04 | 531 extensions, pigsty v4.4, pg/pt/pb/pitr rework, clone & fork    | [v1.5.0](https://github.com/pgsty/pig/releases/tag/v1.5.0) |
| [v1.4.2](/release/pig/#v142) | 2026-06-18 | 524 extensions, PG19 beta, pgrx 0.18.1, Patroni fixes              | [v1.4.2](https://github.com/pgsty/pig/releases/tag/v1.4.2) |
| [v1.4.1](/release/pig/#v141) | 2026-05-01 | 510 extensions, Ubuntu 26.04 support, repo calibration             | [v1.4.1](https://github.com/pgsty/pig/releases/tag/v1.4.1) |
| [v1.4.0](/release/pig/#v140) | 2026-04-19 | 510 extensions, pgrx 0.18.0, more building specs                   | [v1.4.0](https://github.com/pgsty/pig/releases/tag/v1.4.0) |
| [v1.3.4](/release/pig/#v134) | 2026-04-14 | 504 extensions refreshed, release checksums updated                | [v1.3.4](https://github.com/pgsty/pig/releases/tag/v1.3.4) |
| [v1.3.3](/release/pig/#v133) | 2026-04-10 | 481 extensions and Go 1.26.2 update                                | [v1.3.3](https://github.com/pgsty/pig/releases/tag/v1.3.3) |
| [v1.3.2](/release/pig/#v132) | 2026-03-23 | Routine metadata refresh, new `pg tune`, new build aliases         | [v1.3.2](https://github.com/pgsty/pig/releases/tag/v1.3.2) |
| [v1.3.1](/release/pig/#v131) | 2026-03-05 | Retire PG13 defaults, unify PG14-18 support window, 464 extensions | [v1.3.1](https://github.com/pgsty/pig/releases/tag/v1.3.1) |
| [v1.3.0](/release/pig/#v130) | 2026-02-27 | Build pipeline hardening, 461 extensions, new pgedge/ivory support | [v1.3.0](https://github.com/pgsty/pig/releases/tag/v1.3.0) |
| [v1.2.0](/release/pig/#v120) | 2026-02-23 | Unified aliases, routine updates, plan mode, repo fixes            | [v1.2.0](https://github.com/pgsty/pig/releases/tag/v1.2.0) |
| [v1.1.1](https://pig.pgsty.com/release/pig-1.1.1/) | 2026-02-14 | Path, symlink, and build migration fixes | [v1.1.1](https://github.com/pgsty/pig/releases/tag/v1.1.1) |
| [v1.1.0](/release/pig/#v110) | 2026-02-12 | 451 extensions, Agent-Native CLI framework                         | [v1.1.0](https://github.com/pgsty/pig/releases/tag/v1.1.0) |
| [v1.0.0](/release/pig/#v100) | 2026-01-26 | New pg/pt/pb/pitr commands, availability matrix                    | [v1.0.0](https://github.com/pgsty/pig/releases/tag/v1.0.0) |
| [v0.9.0](/release/pig/#v090) | 2025-12-28 | Rename sty deploy and expand sty conf options | [v0.9.0](https://github.com/pgsty/pig/releases/tag/v0.9.0) |
| [v0.8.1](https://pig.pgsty.com/release/pig-0.8.1/) | 2025-12-27 | Restore llvmjit to the pgsql-full alias | [v0.8.1](https://github.com/pgsty/pig/releases/tag/v0.8.1) |
| [v0.8.0](/release/pig/#v080) | 2025-12-26 | 440 extensions, remove sysupdate repo                              | [v0.8.0](https://github.com/pgsty/pig/releases/tag/v0.8.0) |
| [v0.7.5](/release/pig/#v075) | 2025-12-12 | Routine extension update, fixed aliyun mirror                      | [v0.7.5](https://github.com/pgsty/pig/releases/tag/v0.7.5) |
| [v0.7.4](/release/pig/#v074) | 2025-12-01 | Update ivory/pgtde kernel and pgdg extras                          | [v0.7.4](https://github.com/pgsty/pig/releases/tag/v0.7.4) |
| [v0.7.3](/release/pig/#v073) | 2025-11-24 | Fix repo for el10 & debian13                                       | [v0.7.3](https://github.com/pgsty/pig/releases/tag/v0.7.3) |
| [v0.7.2](/release/pig/#v072) | 2025-11-20 | 437 extensions, fix pig build issue                                | [v0.7.2](https://github.com/pgsty/pig/releases/tag/v0.7.2) |
| [v0.7.1](/release/pig/#v071) | 2025-11-10 | New Website, improve in-docker experience                          | [v0.7.1](https://github.com/pgsty/pig/releases/tag/v0.7.1) |
| [v0.7.0](/release/pig/#v070) | 2025-11-05 | Build Enhancement and massive upgrade                              | [v0.7.0](https://github.com/pgsty/pig/releases/tag/v0.7.0) |
| [v0.6.2](/release/pig/#v062) | 2025-10-03 | PG 18 official Repo                                                | [v0.6.2](https://github.com/pgsty/pig/releases/tag/v0.6.2) |
| [v0.6.1](/release/pig/#v061) | 2025-08-14 | CI/CD, el10 stub, PGDG CN Mirror                                   | [v0.6.1](https://github.com/pgsty/pig/releases/tag/v0.6.1) |
| [v0.6.0](/release/pig/#v060) | 2025-07-17 | 423 extension, percona pg_tde, mcp toolbox                         | [v0.6.0](https://github.com/pgsty/pig/releases/tag/v0.6.0) |
| [v0.5.0](/release/pig/#v050) | 2025-06-30 | 422 extension, new extension catalog                               | [v0.5.0](https://github.com/pgsty/pig/releases/tag/v0.5.0) |
| [v0.4.2](/release/pig/#v042) | 2025-05-27 | 421 extension, halo & oriole deb                                   | [v0.4.2](https://github.com/pgsty/pig/releases/tag/v0.4.2) |
| [v0.4.1](/release/pig/#v041) | 2025-05-07 | 414 extension, pg18 alias support                                  | [v0.4.1](https://github.com/pgsty/pig/releases/tag/v0.4.1) |
| [v0.4.0](/release/pig/#v040) | 2025-05-01 | do & pt sub-cmd, halo & orioledb                                   | [v0.4.0](https://github.com/pgsty/pig/releases/tag/v0.4.0) |
| [v0.3.4](/release/pig/#v034) | 2025-04-05 | routine update                                                     | [v0.3.4](https://github.com/pgsty/pig/releases/tag/v0.3.4) |
| [v0.3.3](/release/pig/#v033) | 2025-03-25 | alias, repo, deps                                                  | [v0.3.3](https://github.com/pgsty/pig/releases/tag/v0.3.3) |
| [v0.3.2](/release/pig/#v032) | 2025-03-21 | new extensions                                                     | [v0.3.2](https://github.com/pgsty/pig/releases/tag/v0.3.2) |
| [v0.3.1](/release/pig/#v031) | 2025-03-19 | minor bug fix                                                      | [v0.3.1](https://github.com/pgsty/pig/releases/tag/v0.3.1) |
| [v0.3.0](/release/pig/#v030) | 2025-02-24 | new home page and extension catalog                                | [v0.3.0](https://github.com/pgsty/pig/releases/tag/v0.3.0) |
| [v0.2.2](/release/pig/#v022) | 2025-02-22 | 404 extensions                                                     | [v0.2.2](https://github.com/pgsty/pig/releases/tag/v0.2.2) |
| [v0.2.0](/release/pig/#v020) | 2025-02-14 | 400 extensions                                                     | [v0.2.0](https://github.com/pgsty/pig/releases/tag/v0.2.0) |
| [v0.1.4](/release/pig/#v014) | 2025-02-12 | routine bugfix                                                     | [v0.1.4](https://github.com/pgsty/pig/releases/tag/v0.1.4) |
| [v0.1.3](/release/pig/#v013) | 2025-01-23 | 390 extensions                                                     | [v0.1.3](https://github.com/pgsty/pig/releases/tag/v0.1.3) |
| [v0.1.2](/release/pig/#v012) | 2025-01-12 | the anon extension and 350 other ext                               | [v0.1.2](https://github.com/pgsty/pig/releases/tag/v0.1.2) |
| [v0.1.1](https://pig.pgsty.com/release/pig-0.1.1/) | 2025-01-09 | Update Extension List                                              | [v0.1.1](https://github.com/pgsty/pig/releases/tag/v0.1.1) |
| [v0.1.0](/release/pig/#v010) | 2024-12-29 | repo, ext, sty, and self-update                                    | [v0.1.0](https://github.com/pgsty/pig/releases/tag/v0.1.0) |
| [v0.0.1](/release/pig/#v001) | 2024-12-23 | Genesis Release                                                    | [v0.0.1](https://github.com/pgsty/pig/releases/tag/v0.0.1) |
