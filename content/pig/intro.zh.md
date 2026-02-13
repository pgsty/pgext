---
title: 简介
description: 为什么我们还需要一个新的包管理器？尤其是针对 Postgres 扩展？
icon: CircleHelp
weight: 100
breadcrumbs: false
---


你是否曾因安装或升级 PostgreSQL 扩展而头疼？翻查过时的文档、晦涩难懂的配置脚本，或是在 GitHub 上苦寻分支与补丁？
Postgres 丰富的扩展生态同时意味着复杂的部署流程 —— 在多发行版、多架构环境下尤为棘手。而 PIG 可以为您解决这些烦恼。

这正是 **Pig** 诞生的初衷。Pig 由 Go 语言开发，致力于一站式管理 Postgres 及其 [451](https://pgext.cloud/zh/list) 扩展。
无论是 TimescaleDB、Citus、PGVector，还是 30+ Rust 扩展，亦或 自建 Supabase 所需的全部组件 —— Pig 统一的 CLI 让一切触手可及。
它彻底告别源码编译与杂乱仓库，直接提供版本对齐的 RPM/DEB 包，完美兼容 Debian、Ubuntu、RedHat 等主流发行版，支持 x86 与 Arm 架构，无需猜测，无需折腾。

Pig 并非重复造轮子，而是充分利用系统原生包管理器（APT、YUM、DNF），严格遵循 [PGDG 官方](/zh/repo/pgdg/) 打包规范，确保无缝集成。
你无需在"标准做法"与"快捷方式"之间权衡；Pig 尊重现有仓库，遵循操作系统最佳实践，与现有仓库和软件包和谐共存。
如果你的 Linux 系统和 PostgreSQL 大版本不在 [支持的列表](#linux-兼容性) 中，你还可以使用 [`pig build`](/zh/pig/cmd/build) 直接针对特定组合编译扩展。

想让你的 Postgres 如虎添翼、远离繁琐？欢迎访问 [PIG 官方文档](https://pigsty.io/docs/pig) 获取文档、指南，并查阅庞大的 [扩展列表](https://pgext.cloud/zh/list)，
让你的本地 Postgres 数据库一键进化为全能的多模态数据中台。
如果说 [Postgres 的未来是无可匹敌的可扩展性](https://medium.com/@fengruohang/postgres-is-eating-the-database-world-157c204dcfc4)，那么 Pig 就是帮你解锁它的神灯。毕竟，从没有人抱怨 "扩展太多"。

## 自动化友好

PIG 的命令体系可直接用于自动化脚本：参数风格统一、输出稳定，并在高风险操作中提供 `--dry-run` 或确认步骤，减少误操作风险。


> 《[ANNOUNCE pig: The Postgres Extension Wizard](https://www.postgresql.org/about/news/announce-pig-the-postgres-extension-wizard-2988/)》


--------

## Linux 兼容性

PIG 与 Pigsty 扩展仓库支持以下 Linux 发行版和 PostgreSQL 版本组合：

| OS 代码 | 厂商 | 大版本 | 小版本 | 全名 | PG 版本 | 备注 |
|:---|:---|:---:|:---:|:---|:---|:---:|
| `el7.x86_64` | EL | 7 | 7.9 | CentOS 7 x86 | 13-15 | EOL |
| `el8.x86_64` | EL | 8 | 8.10 | RockyLinux 8 x86 | 13-18 | 即将EOL |
| `el8.aarch64` | EL | 8 | 8.10 | RockyLinux 8 ARM | 13-18 | 即将EOL |
| `el9.x86_64` | EL | 9 | 9.6 | RockyLinux 9 x86 | 13-18 | ✅ |
| `el9.aarch64` | EL | 9 | 9.6 | RockyLinux 9 ARM | 13-18 | ✅ |
| `el10.x86_64` | EL | 10 | 10.0 | RockyLinux 10 x86 | 13-18 | ✅ |
| `el10.aarch64` | EL | 10 | 10.0 | RockyLinux 10 ARM | 13-18 | ✅ |
| `d11.x86_64` | Debian | 11 | 11.11 | Debian 11 x86 | 13-18 | EOL |
| `d11.aarch64` | Debian | 11 | 11.11 | Debian 11 ARM | 13-18 | EOL |
| `d12.x86_64` | Debian | 12 | 12.12 | Debian 12 x86 | 13-18 | ✅ |
| `d12.aarch64` | Debian | 12 | 12.12 | Debian 12 ARM | 13-18 | ✅ |
| `d13.x86_64` | Debian | 13 | 13.1 | Debian 13 x86 | 13-18 | ✅ |
| `d13.aarch64` | Debian | 13 | 13.1 | Debian 13 ARM | 13-18 | ✅ |
| `u20.x86_64` | Ubuntu | 20 | 20.04.6 | Ubuntu 20.04 x86 | 13-18 | EOL |
| `u20.aarch64` | Ubuntu | 20 | 20.04.6 | Ubuntu 20.04 ARM | 13-18 | EOL |
| `u22.x86_64` | Ubuntu | 22 | 22.04.5 | Ubuntu 22.04 x86 | 13-18 | ✅ |
| `u22.aarch64` | Ubuntu | 22 | 22.04.5 | Ubuntu 22.04 ARM | 13-18 | ✅ |
| `u24.x86_64` | Ubuntu | 24 | 24.04.3 | Ubuntu 24.04 x86 | 13-18 | ✅ |
| `u24.aarch64` | Ubuntu | 24 | 24.04.3 | Ubuntu 24.04 ARM | 13-18 | ✅ |
{.full-width}

**说明：**

- **EL** 指 RHEL 兼容发行版，包括 RHEL、CentOS、RockyLinux、AlmaLinux、OracleLinux 等
- **EOL** 表示该操作系统已经或即将停止支持，建议升级到更新版本
- **✅** 表示完整支持，推荐使用
- PG 版本 13-18 表示支持 PostgreSQL 13、14、15、16、17、18 六个大版本
