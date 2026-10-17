---
title: 简介
description: 为什么我们还需要一个新的包管理器？尤其是针对 Postgres 扩展？
icon: fa-solid fa-circle-question
weight: 100
breadcrumb: false
---

PIG 是 PostgreSQL 与扩展的命令行管理工具。它将扩展名转换为 RPM 或 DEB 软件包名，再调用 `apt`、`dnf` 或 `yum` 完成安装。

扩展目录收录了 [{{< param pig_catalog_count >}} 个条目](https://pigsty.cc/ext/list/)。PIG 根据 PostgreSQL 大版本、操作系统和 CPU 架构选择对应包名，也提供源码构建、本地 PostgreSQL 管理、备份恢复和 Pigsty 管理命令。

## 可以做什么

| 任务 | 命令 | 参考文档 |
|:---|:---|:---|
| 查找扩展，检查是否有包可用 | `pig ext list`、`pig ext avail` | [扩展管理](/zh/pig/cmd/ext/) |
| 添加 PGDG 与 Pigsty 仓库 | `pig repo add pgsql -u` | [软件仓库](/zh/pig/cmd/repo/) |
| 安装 PostgreSQL 18 与 pgvector | `pig install -v 18 pg18 vector` | [快速上手](/zh/pig/start/) |
| 从源码构建扩展包 | `pig build pkg NAME` | [构建](/zh/pig/cmd/build/) |
| 启动、检查或停止本地 PostgreSQL 实例 | `pig pg` | [PostgreSQL](/zh/pig/cmd/pg/) |
| 管理 Patroni 集群 | `pig pt` | [Patroni](/zh/pig/cmd/pt/) |
| 备份或恢复实例 | `pig pb`、`pig pitr` | [pgBackRest](/zh/pig/cmd/pb/)、[PITR](/zh/pig/cmd/pitr/) |
| 配置 Pigsty 并执行剧本 | `pig sty`、`pig inventory`、`pig do` | [Pigsty](/zh/pig/cmd/sty/) |

从[上手教程](/zh/pig/start/)开始，安装 PostgreSQL 并启用第一个扩展。

## 软件包如何安装

例如，为 PostgreSQL 18 安装 pgvector：

```bash
sudo pig install vector -v 18
```

PIG 将 `vector` 转换为 EL 上的 `pgvector_18`，或 Debian、Ubuntu 上的 `postgresql-18-pgvector`。系统包管理器负责解析依赖并安装文件。

要启用扩展，连接目标数据库后执行：

```sql
CREATE EXTENSION vector;
```

部分扩展还需要设置 `shared_preload_libraries` 并重启 PostgreSQL。用 `pig ext info NAME` 查看这些要求。升级与移除的操作见[软件包层与数据库层](/zh/pig/cmd/ext/#软件包层与数据库层)。

PIG 只在执行命令时运行。PostgreSQL 与已安装的扩展可以独立运行。你也可以直接用 `apt` 或 `dnf` 访问 [Pigsty 软件仓库](https://pigsty.cc/docs/repo/pgsql/)。

## Linux 兼容性

软件仓库覆盖以下 Linux 版本上的 PostgreSQL **14–18**。每个发行版均提供 **x86_64** 和 **aarch64** 两种架构的软件包。

| OS 代码 | 发行版 | 版本 |
|:---|:---|:---|
| `el8` | RHEL、Rocky Linux、AlmaLinux | 8 |
| `el9` | RHEL、Rocky Linux、AlmaLinux | 9 |
| `el10` | RHEL、Rocky Linux、AlmaLinux | 10 |
| `d12` | Debian | 12 |
| `d13` | Debian | 13 |
| `u22` | Ubuntu | 22.04 LTS |
| `u24` | Ubuntu | 24.04 LTS |
| `u26` | Ubuntu | 26.04 LTS |

不同扩展支持的版本与平台组合有所不同。在目标主机上查询：

```bash
pig ext avail vector
```

{{< param pig_catalog_count >}} 个目录条目中包含已打包的 SQL 工具，以及已移除扩展的保留记录。官网与 CLI 使用同一[目录条目计数口径](/zh/pig/cmd/ext/#catalog-count)。选择软件包时，请查看条目状态和可用矩阵。EL 7 与 Debian 11 不在当前软件包矩阵中。

macOS 版支持目录查询、配置清单编辑等操作。PostgreSQL 软件包安装使用 Linux 包管理器。二进制下载见[安装](/zh/pig/install/)。

## 选择工作流

- **在已有主机上安装：** 用 `pig repo add pgsql -u` 添加仓库，再安装所需软件包。
- **统一配置新主机的仓库：** 用 `pig repo set` 备份并替换现有仓库定义。
- **构建软件包：** 准备构建主机，再用 [`pig build`](/zh/pig/cmd/build/) 执行维护好的构建规格。
- **准备离线安装：** 用 `pig ext import` 收集软件包，用 `pig repo create` 生成索引，再用 `pig repo cache` 打包。详见[离线仓库](/zh/pig/cmd/repo/)。
- **编写脚本或 Agent 工作流：** 用 `pig -o json --help` 查看命令能力，再按[命令参考](/zh/pig/cmd/)选择输出格式与计划选项。

[设计归档](https://pig.pgsty.com/zh/design/)解释主要实现决策，[发布注记](/zh/release/pig/)记录各版本的变化。
