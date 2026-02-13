---
title: INFRA 仓库
description: 可观测性/PostgreSQL工具软件仓库，Linux发行版大版本无关的软件包
icon: Landmark
weight: 200
---

[`pigsty-infra`](https://github.com/pgsty/infra-pkg) 仓库包含与任何 PostgreSQL 版本和 Linux 主版本无关的通用软件包，
包括 prometheus & grafana 技术栈、PostgreSQL 管理工具，以及许多用 Go 编写的实用工具。

该仓库由 [冯若航](https://vonng.com/) ([Vonng](https://github.com/Vonng)) @ [Pigsty](https://doc.pgsty.com) 维护，
您可以在 [https://github.com/pgsty/infra-pkg](https://github.com/pgsty/infra-pkg) 找到所有构建源代码与命令。
为 RHEL / Debian / Ubuntu 发行版提供预构建的 RPM / DEB 包，支持 `x86_64` 和 `aarch64` 架构。
托管于 Cloudflare CDN，提供免费的全球访问。

| Linux  |  包类型  | x86_64 | aarch64 |
|:------:|:-----:|:------:|:-------:|
|   EL   | `rpm` |   ✓    |    ✓    |
| Debian | `deb` |   ✓    |    ✓    |

Infra 仓库的更新记录可以参考 [**Release - Infra 变更日志**](/zh/release/infra)。

## 快速上手

您可以使用 [`pig`](/zh/pig) CLI 工具添加 `pigsty-infra` 仓库，它会自动从 `apt/yum/dnf` 中选择合适的包管理器。

{{< tabs items="默认,镜像,提示" >}}
{{< tab >}}
```bash tab="默认"
curl https://repo.pigsty.io/pig | bash  # 下载并安装 pig CLI 工具
pig repo add infra                      # 将 pigsty-infra 仓库文件添加到您的系统
pig repo update                         # 使用 apt / dnf 更新本地仓库缓存
```
{{< /tab >}}
{{< tab >}}
```bash tab="镜像"
# 在中国大陆或 Cloudflare 不可用时使用
curl https://repo.pigsty.cc/pig | bash  # 从中国 CDN 镜像安装 pig
pig repo add infra                      # 将 pigsty-infra 仓库文件添加到您的系统
pig repo update                         # 使用 apt / dnf 更新本地仓库缓存
```
{{< /tab >}}
{{< tab >}}
```bash tab="提示"
# 您可以使用以下命令管理 infra 仓库：
pig repo add infra -u       # 添加仓库文件，并更新缓存
pig repo add infra -ru      # 删除所有现有仓库，添加仓库并创建缓存
pig repo set infra          # = pigsty repo add infra -ru

pig repo add all            # 将 infra、node、pgsql 仓库添加到您的系统
pig repo set all            # 删除现有仓库，添加上述仓库并更新缓存
```
{{< /tab >}}
{{< /tabs >}}



## 手动设置

您也可以不使用 `pig` CLI 工具直接使用此仓库，手动将其添加到您的 Linux 操作系统仓库列表中：

### APT 仓库

在 **Debian / Ubuntu** 兼容的 Linux 发行版上，您可以手动添加 [GPG 密钥](/zh/repo/gpg) 和 APT 仓库文件：

{{< tabs items="默认,镜像,免密钥" >}}
{{< tab >}}
```bash tab="默认"
# 将 Pigsty 的 GPG 公钥添加到您的系统密钥链以验证包签名，或者直接信任
curl -fsSL https://repo.pigsty.io/key | sudo gpg --dearmor -o /etc/apt/keyrings/pigsty.gpg

# 获取 Debian 发行版代号（distro_codename=jammy, focal, bullseye, bookworm）
# 并将相应的上游仓库地址写入 APT List 文件
distro_codename=$(lsb_release -cs)
sudo tee /etc/apt/sources.list.d/pigsty-infra.list > /dev/null <<EOF
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.io/apt/infra generic main
EOF

# 刷新 APT 仓库缓存
sudo apt update
```
{{< /tab >}}
{{< tab >}}
```bash tab="镜像"
# 在中国大陆或 Cloudflare 不可用时使用
# 将 Pigsty 的 GPG 公钥添加到您的系统密钥链以验证包签名，或者直接信任
curl -fsSL https://repo.pigsty.cc/key | sudo gpg --dearmor -o /etc/apt/keyrings/pigsty.gpg

# 获取 Debian 发行版代号（distro_codename=jammy, focal, bullseye, bookworm）
# 并将相应的上游仓库地址写入 APT List 文件
distro_codename=$(lsb_release -cs)
sudo tee /etc/apt/sources.list.d/pigsty-infra.list > /dev/null <<EOF
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.cc/apt/infra generic main
EOF

# 刷新 APT 仓库缓存
sudo apt update
```
{{< /tab >}}
{{< tab >}}
```bash tab="免密钥"
# 如果您不想信任任何 GPG 密钥，直接信任仓库
distro_codename=$(lsb_release -cs)
sudo tee /etc/apt/sources.list.d/pigsty-infra.list > /dev/null <<EOF
deb [trust=yes] https://repo.pigsty.io/apt/infra generic main
EOF

sudo apt update
```
{{< /tab >}}
{{< /tabs >}}




### YUM 仓库

在 **RHEL** 兼容的 Linux 发行版上，您可以手动添加 [GPG 密钥](/zh/repo/gpg) 和 YUM 仓库文件：

{{< tabs items="默认,镜像,免密钥" >}}
{{< tab >}}
```bash tab="默认"
# 将 Pigsty 的 GPG 公钥添加到您的系统密钥链以验证包签名
curl -fsSL https://repo.pigsty.io/key | sudo tee /etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty >/dev/null

# 将 Pigsty 仓库定义文件添加到 /etc/yum.repos.d/ 目录
sudo tee /etc/yum.repos.d/pigsty-infra.repo > /dev/null <<-'EOF'
[pigsty-infra]
name=Pigsty Infra for $basearch
baseurl=https://repo.pigsty.io/yum/infra/$basearch
skip_if_unavailable = 1
enabled = 1
priority = 1
gpgcheck = 1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty
module_hotfixes=1
EOF

# 刷新 YUM/DNF 仓库缓存
sudo yum makecache;
```
{{< /tab >}}
{{< tab >}}
```bash tab="镜像"
# 在中国大陆或 Cloudflare 不可用时使用
# 将 Pigsty 的 GPG 公钥添加到您的系统密钥链以验证包签名
curl -fsSL https://repo.pigsty.cc/key | sudo tee /etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty >/dev/null

# 将 Pigsty 仓库定义文件添加到 /etc/yum.repos.d/ 目录
sudo tee /etc/yum.repos.d/pigsty-infra.repo > /dev/null <<-'EOF'
[pigsty-infra]
name=Pigsty Infra for $basearch
baseurl=https://repo.pigsty.cc/yum/infra/$basearch
skip_if_unavailable = 1
enabled = 1
priority = 1
gpgcheck = 1
gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty
module_hotfixes=1
EOF

# 刷新 YUM/DNF 仓库缓存
sudo yum makecache;
```
{{< /tab >}}
{{< tab >}}
```bash tab="免密钥"
# 如果您不想信任任何 GPG 密钥，直接信任仓库
sudo tee /etc/yum.repos.d/pigsty-infra.repo > /dev/null <<-'EOF'
[pigsty-infra]
name=Pigsty Infra for $basearch
baseurl=https://repo.pigsty.io/yum/infra/$basearch
skip_if_unavailable = 1
enabled = 1
priority = 1
gpgcheck = 0
module_hotfixes=1
EOF

sudo yum makecache;
```
{{< /tab >}}
{{< /tabs >}}



## 内容

### Grafana 技术栈

|                                           名称                                           |   版本   |    许可证     | 备注          |
|:--------------------------------------------------------------------------------------:|:------:|:----------:|:------------|
|                     [grafana](https://github.com/grafana/grafana/)                     | 12.3.1 |   AGPLv3   | 可视化平台       |
|                        [loki](https://github.com/grafana/loki)                         | 3.1.1  |   AGPLv3   | 日志平台（已弃用）   |
|            [promtail](https://github.com/grafana/loki/releases/tag/v3.0.0)             | 3.0.0  |   AGPLv3   | Loki 日志采集代理（已弃用） |
|     [grafana-infinity-ds](https://github.com/grafana/grafana-infinity-datasource/)     | 3.7.0  | Apache-2.0 | 多功能数据源      |
| [grafana-plugins](https://github.com/pgsty/infra-pkg/tree/main/noarch/grafana-plugins) | 12.3.0 | Apache-2.0 | 额外的面板与数据源插件 |


### Victoria 技术栈

|                                                  名称                                                   |   版本    |    许可证     | 备注                                |
|:-----------------------------------------------------------------------------------------------------:|:-------:|:----------:|:----------------------------------|
|                [victoria-metrics](https://github.com/VictoriaMetrics/VictoriaMetrics)                 | 1.133.0 | Apache-2.0 | VictoriaMetrics，更好的 Prometheus 替代 |
|                   [victoria-logs](https://github.com/VictoriaMetrics/VictoriaMetrics)                 | 1.43.1  | Apache-2.0 | VictoriaLogs，更好的日志平台              |
|                 [victoria-traces](https://github.com/VictoriaMetrics/VictoriaMetrics)                 |  0.5.1  | Apache-2.0 | VictoriaTraces，更好的链路追踪平台          |
|            [victoria-metrics-cluster](https://github.com/VictoriaMetrics/VictoriaMetrics)             | 1.133.0 | Apache-2.0 | VictoriaMetrics 分布式版本             |
|                     [vmutils](https://github.com/VictoriaMetrics/VictoriaMetrics)                     | 1.133.0 | Apache-2.0 | VictoriaMetrics 工具                |
|                     [vlogscli](https://github.com/VictoriaMetrics/VictoriaMetrics)                    | 1.43.1  | Apache-2.0 | VictoriaLogs 命令行工具                |
|                      [vlagent](https://github.com/VictoriaMetrics/VictoriaMetrics)                    | 1.43.1  | Apache-2.0 | VictoriaLogs 日志 Agent             |
|    [grafana-victorialogs-ds](https://github.com/VictoriaMetrics/victorialogs-datasource/releases/)    | 0.23.3  | Apache-2.0 | VictoriaLogs 的 Grafana 数据源        |
| [grafana-victoriametrics-ds](https://github.com/VictoriaMetrics/victoriametrics-datasource/releases/) | 0.20.0  | Apache-2.0 | VictoriaMetrics 的 Grafana 数据源     |


### Prometheus 技术栈

|                                  名称                                  |   版本   |    许可证     | 备注                |
|:--------------------------------------------------------------------:|:------:|:----------:|:------------------|
|        [prometheus](https://github.com/prometheus/prometheus)        | 3.9.1  | Apache-2.0 | 开源时序数据库与监控平台      |
|       [pushgateway](https://github.com/prometheus/pushgateway)       | 1.11.2 | Apache-2.0 | 向 Prometheus 推送指标 |
|      [alertmanager](https://github.com/prometheus/alertmanager)      | 0.30.0 | Apache-2.0 | 管理告警事件            |
| [blackbox_exporter](https://github.com/prometheus/blackbox_exporter) | 0.27.0 | Apache-2.0 | 向端点发送探针           |


### 指标导出器

|                                   名称                                    |   版本   |    许可证     | 备注                   |
|:-----------------------------------------------------------------------:|:------:|:----------:|:---------------------|
|           [pg_exporter](https://github.com/pgsty/pg_exporter)           | 1.1.2  | Apache-2.0 | 高级 Postgres 指标导出器    |
|  [pgbackrest_exporter](https://github.com/woblerr/pgbackrest_exporter)  | 0.22.0 |    MIT     | 导出 pgbackrest 指标     |
|      [node_exporter](https://github.com/prometheus/node_exporter)       | 1.10.2 | Apache-2.0 | 导出 Linux 节点指标        |
|   [keepalived_exporter](https://github.com/mehdy/keepalived-exporter)   | 1.7.0  |  GPL-3.0   | 导出 keepalived/VIP 指标 |
| [nginx_exporter](https://github.com/nginxinc/nginx-prometheus-exporter) | 1.5.1  | Apache-2.0 | 导出 nginx 指标          |
|  [zfs_exporter](https://github.com/waitingsong/zfs_exporter/releases/)  | 3.8.1  |    MIT     | 导出 zfs 指标            |
|    [mysqld_exporter](https://github.com/prometheus/mysqld_exporter)     | 0.18.0 | Apache-2.0 | 导出 mysql 指标          |
|      [redis_exporter](https://github.com/oliver006/redis_exporter)      | 1.80.1 |    MIT     | 导出 redis 指标          |
|      [kafka_exporter](https://github.com/danielqsj/kafka_exporter)      | 1.9.0  | Apache-2.0 | 导出 kafka 指标          |
|     [mongodb_exporter](https://github.com/percona/mongodb_exporter)     | 0.47.2 | Apache-2.0 | 导出 mongodb 指标        |
|                [mtail](https://github.com/google/mtail)                 | 3.0.8  | Apache-2.0 | 分析日志并生成指标            |
|        [vector](https://github.com/vectordotdev/vector/releases)        | 0.52.0 |  MPL-2.0   | 多功能日志收集器             |

### 对象存储

|                         名称                          |       版本       |    许可证     | 备注             |
|:---------------------------------------------------:|:--------------:|:----------:|:---------------|
|       [minio](https://github.com/minio/minio)       | 20251203120000 |   AGPLv3   | 开源 S3 服务器      |
|         [mcli](https://github.com/minio/mc)         | 20250813083541 |   AGPLv3   | 开源 S3 客户端      |
|      [rustfs](https://github.com/rustfs/rustfs)     |   alpha.80     | Apache-2.0 | 开源 MinIO，Alpha |
|  [garage](https://git.deuxfleurs.fr/Deuxfleurs/garage) |     2.1.0      |  AGPL-3.0  | 轻量级 S3         |
| [seaweedfs](https://github.com/seaweedfs/seaweedfs) |      4.06      | Apache-2.0 | 小文件 S3         |
|     [rclone](https://github.com/rclone/rclone/)     |     1.72.1     |    MIT     | S3 命令行工具       |
|     [restic](https://github.com/restic/restic)      |     0.18.1     |   BSD-2    | 备份工具           |
|    [juicefs](https://github.com/juicedata/juicefs)  |     1.3.1      | Apache-2.0 | S3 上的文件系统      |

### 数据库

PostgreSQL 相关工具、数据库管理系统和其他实用程序

|                            名称                             |   版本    |    许可证     | 备注                 |
|:---------------------------------------------------------:|:-------:|:----------:|:-------------------|
|          [etcd](https://github.com/etcd-io/etcd)          |  3.6.7  | Apache-2.0 | 容错分布式协调服务          |
|        [kafka](https://github.com/apache/kafka)           |  4.1.1  | Apache-2.0 | 消息队列               |
|        [duckdb](https://github.com/duckdb/duckdb)         |  1.4.3  |    MIT     | 嵌入式 OLAP           |
|     [ferretdb](https://github.com/FerretDB/FerretDB)      |  2.7.0  | Apache-2.0 | 基于 PG 的 MongoDB    |
| [tigerbeetle](https://github.com/tigerbeetle/tigerbeetle) | 0.16.68 | Apache-2.0 | 金融 OLTP            |
|     [IvorySQL](https://github.com/IvorySQL/IvorySQL)      |   5.1   | Apache-2.0 | Oracle 兼容的 PG 18.1 |

### 工具

Pig 包管理器、PostgreSQL 工具和其他数据库相关实用程序

|                                          名称                                           |   版本   |    许可证     | 备注                           |
|:-------------------------------------------------------------------------------------:|:------:|:----------:|:-----------------------------|
|                          [pig](https://github.com/pgsty/pig)                          | 0.9.1  | Apache-2.0 | PG 包管理器                      |
|           [vip-manager](https://github.com/cybertec-postgresql/vip-manager)           | 4.0.0  |   BSD-2    | 将 L2 VIP 绑定到 PG 主节点          |
|                       [pgflo](https://github.com/pgflo/pg_flo)                        | 0.0.15 | Apache-2.0 | 实时流式传输、转换和路由 PG 数据           |
|                   [pgschema](https://github.com/schemagood/pgschema)                  | 1.4.2  | Apache-2.0 | 执行 PG 模式迁移                   |
|          [pg_timetable](https://github.com/cybertec-postgresql/pg_timetable)          | 6.2.0  | PostgreSQL | PostgreSQL 高级调度              |
|          [timescaledb-tools](https://github.com/timescale/timescaledb-tune)           | 0.18.1 | Apache-2.0 | 优化 timescaledb 参数            |
| [timescaledb-event-streamer](https://github.com/noctarius/timescaledb-event-streamer) | 0.20.0 | Apache-2.0 | timescaledb 超表 CDC           |
|                     [dblab](https://github.com/danvergara/dblab)                      | 0.34.2 |    MIT     | 多数据库命令行工具                    |
|                   [sqlcmd](https://github.com/microsoft/go-sqlcmd)                    | 1.9.0  |    MIT     | MS SQL Server 数据库客户端         |
|                        [pev2](https://github.com/dalibo/pev2)                         | 1.19.0 | PostgreSQL | PostgreSQL 执行计划可视化工具 2       |
|                      [sealos](https://github.com/labring/sealos)                      | 5.0.1  | Apache-2.0 | 开箱即用的 Kubernetes 发行版         |
|                     [vray](https://github.com/v2fly/v2ray-core)                       | 5.44.1 |    MIT     | 构建代理以绕过网络限制                  |
|                   [postgrest](https://github.com/PostgREST/postgrest)                 |  14.3  |    MIT     | PostgreSQL RESTful API 服务器   |
|               [npgsqlrest](https://github.com/vb-consulting/NpgsqlRest)               | 3.4.3  |    MIT     | .NET PostgreSQL REST API 生成器 |


### AI

AI Agent，MCP 工具箱，编码 IDE，Python/Go/Node 工具……

|                              名称                               |    版本    |    许可证     | 备注                             |
|:-------------------------------------------------------------:|:--------:|:----------:|:-------------------------------|
|         [claude](https://github.com/anthropics/claude-code)   |  2.1.9   | Proprietary | Claude Code - Anthropic 代理编程工具 |
|        [opencode](https://github.com/opencode-ai/opencode)    | 1.1.23   |    MIT     | 终端 AI 编程助手                     |
|       [code-server](https://github.com/coder/code-server)     | 4.108.0  |    MIT     | 浏览器中的 VS Code                  |
|      [genai-toolbox](https://github.com/googleapis/genai-toolbox) | 0.25.0 | Apache-2.0 | Google 数据库 MCP 服务器             |
|              [uv](https://github.com/astral-sh/uv)            |  0.9.26  |    MIT     | 新一代 Python 包管理器                |
|                  [golang](https://golang.org/)                | 1.25.6   |   BSD-3    | Go 编译器                         |
|                 [nodejs](https://nodejs.org/)                 | 24.12.0  | MIT/Mixed  | 在服务端运行 Javascript              |


> 提示：使用 Victoria 数据源时，别忘了在 `/etc/grafana/grafana.ini` 中设置 `allow_loading_unsigned_plugins = victoriametrics-logs-datasource,victoriametrics-metrics-datasource`
