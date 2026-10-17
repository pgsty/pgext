---
title: "pig do"
description: "使用 pig do 执行 Pigsty 管理任务"
weight: 620
icon: fas fa-list-check
---

`pig do` 在选定的 Pigsty 目录中运行管理剧本，按 Inventory 配置添加集群、创建数据库和用户、安装软件包或移除服务。

先在 Inventory 中写好目标配置，再执行对应命令。命令会直接调用 Ansible；该命令组不提供通用 `--plan` 模式。

```bash
pig do pgsql-add  pg-test 10.10.10.12     # 添加一个 PostgreSQL 实例
pig do pgsql-rm   pg-test 10.10.10.12     # 移除一个 PostgreSQL 实例
pig do pgsql-db   pg-meta app             # 创建或更新数据库
pig do pgsql-user pg-meta dbuser_app      # 创建或更新用户
pig do node-pkg   pg-meta restic          # 在选中节点安装软件包
pig do redis-add  redis-meta              # 初始化 Redis 集群
```

Inventory 来源与 [`sty deploy`](/zh/pig/cmd/sty/#sty-deploy) 使用同一配置解析结果，不再按 Pigsty home
重新解释路径。显式、环境或默认来源不存在时，在执行剧本前失败。既有的目录与可执行 Inventory
来源仍交给 Ansible；静态 Inventory 命令不会执行它们。

## 命令总览

| 命令 | 用途 | 必需参数 |
|:---|:---|:---|
| `pgsql-add` | 添加 PostgreSQL 集群或实例 | 集群，可选 IP |
| `pgsql-rm` | 移除 PostgreSQL 集群或实例 | 集群，可选 IP |
| `pgsql-db` | 应用数据库声明 | 集群、数据库 |
| `pgsql-user` | 应用用户声明 | 集群、用户 |
| `pgsql-ext` | 通过 Pigsty 安装扩展 | 集群，可选扩展 |
| `pgsql-hba` | 刷新 PostgreSQL HBA 规则 | 集群 |
| `pgsql-svc` | 刷新 PostgreSQL 服务定义 | 集群 |
| `pgmon-add` / `pgmon-rm` | 添加或移除远程监控目标 | 集群 |
| `node-add` / `node-rm` | 添加或移除选中节点 | 一个或多个选择器 |
| `node-repo` | 在选中节点配置仓库模块 | 可选选择器与模块 |
| `node-pkg` | 在选中节点安装或更新软件包 | 选择器，可选软件包 |
| `repo-build` | 重建 Pigsty 基础设施仓库 | 无 |
| `redis-add` / `redis-rm` | 添加或移除 Redis 集群、节点或实例 | 选择器，可选端口 |
{.full-width}

使用 `pig do COMMAND --help` 查看单项操作的别名与精确示例。

## PostgreSQL 操作

```bash
pig do pgsql-add pg-meta                  # 初始化声明的集群
pig do pgsql-add pg-test 10.10.10.12      # 添加一个声明的副本
pig do pgsql-rm pg-test 10.10.10.13       # 移除一个实例
pig do pgsql-rm pg-test --uninstall       # 移除时同时卸载软件包
pig do pgsql-db pg-meta meta              # 创建或更新 meta 数据库
pig do pgsql-user pg-meta dbuser_view     # 创建或更新用户
pig do pgsql-ext pg-meta postgis vector   # 安装扩展软件包
pig do pgsql-hba pg-meta                  # 刷新 pg_hba 规则
pig do pgsql-svc pg-meta                  # 刷新 PostgreSQL 服务
```

`pgsql-db` 与 `pgsql-user` 使用 Inventory 中已有的 `pg_databases` 和 `pg_users` 声明，不修改清单。`pgsql-rm --uninstall` 在移除服务时还会卸载软件包；同一主机仍需这些包时，省略此选项。

### 名称规则

| 对象 | 允许的名称 |
|:---|:---|
| PostgreSQL 集群 | 字母、数字与连字符；首字符为字母或数字；保留 `root`、`all`、`ungrouped` |
| PostgreSQL 用户 | 最多 63 个字符，首字符为小写字母或 `_`，后续为小写字母、数字、`_`、`@`、`.`、`-`；保留 `postgres` |
| 数据库 | 最多 63 个字符，首字符为字母或 `_`，后续为字母、数字、`_`、`@`、`.`、`-` |

以上规则适用于明确要求集群名、用户名或数据库名的参数。选择器参数支持下节所述的目标选择方式。

## 节点与仓库操作

选择器可以是集群名、主机名、IP 地址或 Ansible 主机模式。实际目标由当前 Inventory 决定。

```bash
pig do node-add pg-test                    # 按集群选择并添加节点
pig do node-rm 10.10.10.13                 # 移除一个选中节点
pig do node-repo pg-meta node,infra        # 配置选中的仓库模块
pig do node-pkg pg-meta openssh restic     # 安装或更新软件包
pig do repo-build                          # 重建 infra 仓库
```

`node-repo` 省略选择器时作用于 `node.yml` 选中的全部主机；省略模块时使用 Inventory 中的设置。仓库模块来自已安装的 Pigsty 版本，例如 `local`、`infra`、`pgsql`、`node` 与 `extra`。
安装软件包前，应确认目标节点能够访问所需仓库。

## Redis 操作

```bash
pig do redis-add redis-meta                # 初始化声明的 Redis 集群
pig do redis-add 10.10.10.11 6379 6380     # 添加选中实例
pig do redis-rm 10.10.10.11 6379           # 移除一个选中实例
pig do redis-rm redis-test --uninstall     # 移除并卸载软件包
```

`redis-rm --uninstall` 在移除节点或集群时还会卸载软件包。

## 配置与输出

`pig do` 使用 PIG 全局选项：

| 选项 | 用途 |
|:---|:---|
| `-H, --home` | 选择 Pigsty 主目录 |
| `-i, --inventory` | 选择 Pigsty Inventory |
| `-o, --output` | 选择文本、JSON、YAML 或格式化 JSON 包装 |
| `--log-level`、`--log-path` | 配置 PIG 诊断日志 |
{.full-width}

结构化模式报告请求的操作与捕获到的执行结果。底层 Ansible 与 Playbook 行为仍由已安装的 Pigsty 版本决定。

完整剧本变量与操作步骤见 [Pigsty 文档](https://pigsty.cc/docs/)。
