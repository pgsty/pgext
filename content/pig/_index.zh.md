---
title: "PIG 1.10.3 文档"
description: PostgreSQL 扩展生态包管理器与管理命令行工具
icon: fa-solid fa-piggy-bank
page_width: full
cascade:
  type: docs
breadcrumb: false
comments: false
weight: 100
---

当前文档版本：**v{{< param pig_version >}}**，查看 [发布说明](/zh/release/pig/#v1103)。

本次版本新增四个扩展并更新扩展版本与包信息，内嵌扩展目录增至 **{{< param pig_catalog_count >}} 个条目**。

PIG 用统一的命令查找、安装和构建 PostgreSQL 扩展包，也提供 PostgreSQL、Patroni、pgBackRest 和 Pigsty 管理命令。软件包安装由系统的 `apt`、`dnf` 或 `yum` 完成。

## 从这里开始

| 你要做什么 | 阅读页面 |
|:---|:---|
| 安装 PostgreSQL 并启用第一个扩展 | [快速上手](/zh/pig/start/) |
| 了解用途、工作方式与支持平台 | [简介](/zh/pig/intro/) |
| 下载、安装、升级或卸载 PIG | [安装](/zh/pig/install/) |
| 查询命令、选项与自动化用法 | [命令参考](/zh/pig/cmd/) |

## 安装 PostgreSQL 与扩展

在有 root 或 `sudo` 权限的 Linux 主机上执行：

```bash
curl -fsSL https://repo.pigsty.cc/pig | bash
sudo pig repo add pgsql --update
sudo pig install -v 18 pg18 vector
```

中国大陆以外可将安装地址换为 `https://repo.pigsty.io/pig`。上述命令安装 PostgreSQL 18 与 pgvector 软件包。接下来按[快速上手](/zh/pig/start/)启动测试实例，并用 `CREATE EXTENSION vector` 在数据库中启用扩展。

## 命令参考

| 任务 | 命令 |
|:---|:---|
| 配置 APT/DNF 软件仓库 | [`pig repo`](/zh/pig/cmd/repo/) |
| 查询目录、安装或更新扩展包 | [`pig ext`](/zh/pig/cmd/ext/)、[`pig install`](/zh/pig/cmd/#pig-install) |
| 从源码构建扩展包 | [`pig build`](/zh/pig/cmd/build/) |
| 管理本地 PostgreSQL 实例与数据库 | [`pig pg`](/zh/pig/cmd/pg/) |
| 通过 patronictl 管理 Patroni 集群 | [`pig pt`](/zh/pig/cmd/pt/) |
| 检查或重载 pg_exporter | [`pig pe`](/zh/pig/cmd/pe/) |
| 使用 pgBackRest 备份与恢复 | [`pig pb`](/zh/pig/cmd/pb/) |
| 执行时间点恢复 | [`pig pitr`](/zh/pig/cmd/pitr/) |
| 安装和配置 Pigsty，管理 Grafana | [`pig sty`](/zh/pig/cmd/sty/) |
| 读取、编辑和校验 Pigsty 配置清单 | [`pig inventory`](/zh/pig/cmd/inventory/) |
| 运行 Pigsty 管理剧本 | [`pig do`](/zh/pig/cmd/do/) |
| 采集主机、数据库与备份状态 | [`pig context`](/zh/pig/cmd/#pig-context) |

执行 `pig help <command>` 查看命令帮助。编写自动化脚本时，用 `pig -o json --help` 查看命令能力，再选择对应的输出与计划选项。

## 查找更多信息

按 {{< kbd "Ctrl" "K" >}}（macOS 为 {{< kbd "⌘" "K" >}}）打开站内搜索。按 {{< kbd "/" >}} 打开页面操作；选择 AI 助手操作后，会将当前页面的公开 URL 与阅读提示传给所选服务。

- [设计归档](https://pig.pgsty.com/zh/design/)：实现决策、取舍与提案。
- [发布注记](/zh/release/pig/)：各版本的变更和下载文件。
- [扩展目录](https://pigsty.cc/ext/list/)：扩展说明与软件包信息。
- [GitHub Issues](https://github.com/pgsty/pig/issues)：报告问题或提出功能建议。

PIG 由 [Vonng](https://vonng.com/) 开发，以 [Apache 2.0](https://github.com/pgsty/pig/blob/main/LICENSE) 许可证开源。[Pigsty](https://pigsty.cc/) 提供 PostgreSQL 部署、监控、备份与高可用管理。
