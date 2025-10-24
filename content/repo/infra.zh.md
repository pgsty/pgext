---
title: INFRA 仓库
description: 可观测性/PostgreSQL工具软件仓库，Linux发行版大版本无关的软件包
icon: Landmark
weight: 200
---

[`pigsty-infra`](https://github.com/pgsty/infra-pkg) 仓库包含与任何 PostgreSQL 版本和 Linux 主版本无关的通用软件包，
包括 prometheus & grafana 技术栈、PostgreSQL 管理工具，以及许多用 Go 编写的实用工具。

该仓库由 [Pigsty](https://doc.pgsty.com) 维护，您可以在 [https://github.com/pgsty/infra-pkg](https://github.com/pgsty/infra-pkg) 找到所有构建规格。
为 RHEL / Debian / Ubuntu 发行版提供预构建的 RPM / DEB 包，支持 `x86_64` 和 `aarch64` 架构。

| Linux  | 包类型  | x86_64 | aarch64 |
|:------:|:-------:|:------:|:-------:|
|   EL   |  `rpm`  |   ✓    |    ✓    |
| Debian |  `deb`  |   ✓    |    ✓    |

--------

## 快速上手

您可以使用 [`pig`](/cmd/repo/) CLI 工具添加 `pigsty-infra` 仓库，它会自动从 `apt/yum/dnf` 中选择合适的包管理器。

```bash tab="默认"
curl https://repo.pigsty.io/pig | bash  # 下载并安装 pig CLI 工具
pig repo add infra                      # 将 pigsty-infra 仓库文件添加到您的系统
pig repo update                         # 使用 apt / dnf 更新本地仓库缓存
```
```bash tab="镜像"
curl https://repo.pigsty.cc/pig | bash  # 从镜像安装 pig
pig repo add infra                      # 将 pigsty-infra 仓库文件添加到您的系统
pig repo update                         # 使用 apt / dnf 更新本地仓库缓存
```
```bash tab="提示"
# 您可以使用以下命令管理 infra 仓库：
pig repo add infra -u       # 添加仓库文件，并更新缓存
pig repo add infra -ru      # 删除所有现有仓库，添加仓库并创建缓存
pig repo set infra          # = pigsty repo add infra -ru

pig repo add all            # 将 infra、node、pgsql 仓库添加到您的系统
pig repo set all            # 删除现有仓库，添加上述仓库并更新缓存
```



--------

## 手动设置

您也可以不使用 `pig` CLI 工具直接使用此仓库，手动将其添加到您的 Linux 操作系统仓库列表中：

### APT 仓库

在 **Debian / Ubuntu** 兼容的 Linux 发行版上，您可以手动添加 [GPG 密钥](/repo/gpg) 和 APT 仓库文件：

```bash tab="默认"
# 将 Pigsty 的 GPG 公钥添加到您的系统密钥链以验证包签名
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
```bash tab="镜像"
# 将 Pigsty 的 GPG 公钥添加到您的系统密钥链以验证包签名
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

### YUM 仓库

在 **RHEL** 兼容的 Linux 发行版上，您可以手动添加 [GPG 密钥](/repo/gpg) 和 YUM 仓库文件：

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
```bash tab="镜像"
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




--------

## 内容

### Prometheus 技术栈

|                                     名称                                      |   版本    | 许可证 | 备注 |
|:---------------------------------------------------------------------------:|:-------:|:---:|:---|
|    [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics)    | 1.120.0 |     |    |
| [VictoriaLogs](https://github.com/VictoriaMetrics/VictoriaMetrics/releases) | 1.24.0  |     |    |
|           [prometheus](https://github.com/prometheus/prometheus)            |  3.4.2  |     |    |
|          [pushgateway](https://github.com/prometheus/pushgateway)           | 1.11.1  |     |    |
|         [alertmanager](https://github.com/prometheus/alertmanager)          | 0.28.1  |     |    |
|    [blackbox_exporter](https://github.com/prometheus/blackbox_exporter)     | 0.27.0  |     |    |
|             [pg_exporter](https://github.com/Vonng/pg_exporter)             |  1.0.0  |     |    |
|    [pgbackrest_exporter](https://github.com/woblerr/pgbackrest_exporter)    | 0.20.0  |     |    |
|        [node_exporter](https://github.com/prometheus/node_exporter)         |  1.9.1  |     |    |
|     [keepalived_exporter](https://github.com/mehdy/keepalived-exporter)     |  1.7.0  |     |    |
|   [nginx_exporter](https://github.com/nginxinc/nginx-prometheus-exporter)   |  1.4.2  |     |    |
|    [zfs_exporter](https://github.com/waitingsong/zfs_exporter/releases/)    |  3.8.1  |     |    |
|      [mysqld_exporter](https://github.com/prometheus/mysqld_exporter)       | 0.17.2  |     |    |
|        [redis_exporter](https://github.com/oliver006/redis_exporter)        | 1.74.0  |     |    |
|        [kafka_exporter](https://github.com/danielqsj/kafka_exporter)        |  1.9.0  |     |    |
|       [mongodb_exporter](https://github.com/percona/mongodb_exporter)       | 0.44.0  |     |    |
|                  [mtail](https://github.com/google/mtail)                   |  3.0.8  |     |    |

### Grafana 技术栈

|                                                  名称                                                   |   版本   | 许可证 | 备注    |
|:-----------------------------------------------------------------------------------------------------:|:------:|:---:|:------|
|                            [grafana](https://github.com/grafana/grafana/)                             | 12.0.2 |     | 可视化平台 |
|                                [loki](https://github.com/grafana/loki)                                | 3.1.1  |     | 日志平台  |
|                    [promtail](https://github.com/grafana/loki/releases/tag/v3.0.0)                    | 3.0.0  |     | 已废弃   |
|                       [vector](https://github.com/vectordotdev/vector/releases)                       | 0.48.0 |     |       |
|            [grafana-infinity-ds](https://github.com/grafana/grafana-infinity-datasource/)             | 3.3.0  |     |       |
|    [grafana-victorialogs-ds](https://github.com/VictoriaMetrics/victorialogs-datasource/releases/)    | 0.18.1 |     |       |
| [grafana-victoriametrics-ds](https://github.com/VictoriaMetrics/victoriametrics-datasource/releases/) | 0.16.0 |     |       |
|        [grafana-plugins](https://github.com/pgsty/infra-pkg/tree/main/noarch/grafana-plugins)         | 12.0.0 |     |       |

### 数据库组件

PostgreSQL 相关工具、数据库管理系统和其他实用程序

|                            名称                             |       版本       | 许可证 | 备注                 |
|:---------------------------------------------------------:|:--------------:|:---:|:-------------------|
|          [etcd](https://github.com/etcd-io/etcd)          |     3.6.1      |     | 容错分布式协调服务          |
|          [minio](https://github.com/minio/minio)          | 20250613113347 |     | 开源 S3 服务器          |
|            [mcli](https://github.com/minio/mc)            | 20250521015954 |     | 开源 S3 客户端          |
|        [kafka](https://kafka.apache.org/downloads)        |     4.0.0      |     | 消息队列               |
|        [duckdb](https://github.com/duckdb/duckdb)         |     1.3.1      |     | 嵌入式 OLAP           |
|     [ferretdb](https://github.com/FerretDB/FerretDB)      |     2.3.1      |     | 基于 PG 的 MongoDB    |
| [tigerbeetle](https://github.com/tigerbeetle/tigerbeetle) |    0.16.48     |     | 金融 OLTP            |
|     [IvorySQL](https://github.com/IvorySQL/IvorySQL)      |      4.5       |     | Oracle 兼容的 PG 17.5 |

### 数据库工具

Pig 包管理器、PostgreSQL 工具、数据库管理系统和其他实用程序

|                                 名称                                  |   版本   |    许可证     |   备注    |
|:-------------------------------------------------------------------:|:------:|:----------:|:-------:|
|                 [pig](https://github.com/pgsty/pig)                 | 0.6.2  | Apache-2.0 | PG 包管理器 |
|  [vip-manager](https://github.com/cybertec-postgresql/vip-manager)  | 4.0.0  |            |         |
| [pg_timetable](https://github.com/cybertec-postgresql/pg_timetable) | 5.13.0 |            |         |
|  [pev2](https://github.com/pgsty/infra-pkg/tree/main/noarch/pev2)   | 1.16.0 |            |         |
|             [sealos](https://github.com/labring/sealos)             | 5.0.1  |            |         |
|        [rclone](https://github.com/rclone/rclone/releases/)         | 1.71.1 |            |         |
|             [restic](https://github.com/restic/restic)              | 0.18.1 |            |         |
|           [juicefs](https://github.com/juicedata/juicefs)           | 1.3.0  |            |         |
|            [dblab](https://github.com/danvergara/dblab)             | 0.33.0 |            |         |
|            [v2ray](https://github.com/v2fly/v2ray-core)             | 5.28.0 |            |         |
