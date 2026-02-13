---
title: INFRA Repo
description: Packages that are generic to any PostgreSQL version and Linux major version.
icon: Landmark
weight: 200
---

The [`pigsty-infra`](https://github.com/pgsty/infra-pkg) repo contains packages that are generic to any PostgreSQL version and Linux major version,
including prometheus & grafana stack, admin tools for postgres, and many utils written in go.

This repo is maintained by [Ruohang Feng](https://vonng.com/en/) ([Vonng](https://github.com/Vonng)) @ [Pigsty](https://doc.pgsty.com), 
you can find all the build specs on [https://github.com/pgsty/infra-pkg](https://github.com/pgsty/infra-pkg).
Prebuilt RPM / DEB packages for RHEL / Debian / Ubuntu distros available for `x86_64` and `aarch64` arch.
Hosted on cloudflare CDN for free global access.

| Linux  | Package | x86_64 | aarch64 |
|:------:|:-------:|:------:|:-------:|
|   EL   |  `rpm`  |   ✓    |    ✓    |
| Debian |  `deb`  |   ✓    |    ✓    |

You can check the [**Release - Infra Changelog**](/release/infra) for the latest updates.

## Quick Start

You can add the `pigsty-infra` repo with the [`pig`](/pig) CLI tool, it will automatically choose from `apt/yum/dnf`.

{{< tabs items="Default,Mirror,Hint" >}}
{{< tab >}}
```bash tab="default"
curl https://repo.pigsty.io/pig | bash  # download and install the pig CLI tool
pig repo add infra                      # add pigsty-infra repo file to your system
pig repo update                         # update local repo cache with apt / dnf
```
{{< /tab >}}
{{< tab >}}
```bash tab="mirror"
# use when in mainland china or cloudflare is down
curl https://repo.pigsty.cc/pig | bash  # install pig from china CDN mirror 
pig repo add infra                      # add pigsty-infra repo file to your system
pig repo update                         # update local repo cache with apt / dnf
```
{{< /tab >}}
{{< tab >}}
```bash tab="hint"
# you can manage infra repo with these commands:
pig repo add infra -u       # add repo file, and update cache
pig repo add infra -ru      # remove all existing repo, add repo and make cache
pig repo set infra          # = pigsty repo add infra -ru

pig repo add all            # add infra, node, pgsql repo to your system
pig repo set all            # remove existing repo, add above repos and update cache
```
{{< /tab >}}
{{< /tabs >}}



## Manual Setup

You can also use this repo directly without the `pig` CLI tool, by add them to your linux os repo list manually:

### APT Repo

On **Debian / Ubuntu** compatible Linux distros, you can add the [GPG Key](/repo/gpg) and APT repo file manually with:

{{< tabs items="Default,Mirror,NoKey" >}}
{{< tab >}}
```bash tab="default"
# Add Pigsty's GPG public key to your system keychain to verify package signatures, or just trust
curl -fsSL https://repo.pigsty.io/key | sudo gpg --dearmor -o /etc/apt/keyrings/pigsty.gpg

# Get Debian distribution codename (distro_codename=jammy, focal, bullseye, bookworm)
# and write the corresponding upstream repository address to the APT List file
distro_codename=$(lsb_release -cs)
sudo tee /etc/apt/sources.list.d/pigsty-infra.list > /dev/null <<EOF
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.io/apt/infra generic main
EOF

# Refresh APT repository cache
sudo apt update
```
{{< /tab >}}
{{< tab >}}
```bash tab="mirror"
# use when in mainland china or cloudflare is down
# Add Pigsty's GPG public key to your system keychain to verify package signatures, or just trust
curl -fsSL https://repo.pigsty.cc/key | sudo gpg --dearmor -o /etc/apt/keyrings/pigsty.gpg

# Get Debian distribution codename (distro_codename=jammy, focal, bullseye, bookworm)
# and write the corresponding upstream repository address to the APT List file
distro_codename=$(lsb_release -cs)
sudo tee /etc/apt/sources.list.d/pigsty-infra.list > /dev/null <<EOF
deb [signed-by=/etc/apt/keyrings/pigsty.gpg] https://repo.pigsty.cc/apt/infra generic main
EOF

# Refresh APT repository cache
sudo apt update
```
{{< /tab >}}
{{< tab >}}
```bash tab="nokey"
# If you don't want to trust any GPG key, just trust the repo directly
distro_codename=$(lsb_release -cs)
sudo tee /etc/apt/sources.list.d/pigsty-infra.list > /dev/null <<EOF
deb [trust=yes] https://repo.pigsty.io/apt/infra generic main
EOF

sudo apt update
```
{{< /tab >}}
{{< /tabs >}}




### YUM Repo

On **RHEL** compatible Linux distros, you can add the [GPG Key](/repo/gpg) and Yum repo file manually with:

{{< tabs items="Default,Mirror,NoKey" >}}
{{< tab >}}
```bash tab="default"
# Add Pigsty's GPG public key to your system keychain to verify package signatures
curl -fsSL https://repo.pigsty.io/key | sudo tee /etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty >/dev/null

# Add Pigsty Repo definition files to /etc/yum.repos.d/ directory
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

# Refresh YUM/DNF repository cache
sudo yum makecache;
```
{{< /tab >}}
{{< tab >}}
```bash tab="mirror"
# use when in mainland china or cloudflare is down
# Add Pigsty's GPG public key to your system keychain to verify package signatures
curl -fsSL https://repo.pigsty.cc/key | sudo tee /etc/pki/rpm-gpg/RPM-GPG-KEY-pigsty >/dev/null

# Add Pigsty Repo definition files to /etc/yum.repos.d/ directory
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

# Refresh YUM/DNF repository cache
sudo yum makecache;
```
{{< /tab >}}
{{< tab >}}
```bash tab="nokey"
# If you don't want to trust any GPG key, just trust the repo directly
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



## Content

### Grafana Stack

|                                           Name                                            | Version |  License   | Comment                          |
|:-----------------------------------------------------------------------------------------:|:-------:|:----------:|:---------------------------------|
|                      [grafana](https://github.com/grafana/grafana/)                       | 12.3.1  |   AGPLv3   | Visualization Platform           |
|                          [loki](https://github.com/grafana/loki)                          |  3.1.1  |   AGPLv3   | Logging platform (Deprecated)    |
|              [promtail](https://github.com/grafana/loki/releases/tag/v3.0.0)              |  3.0.0  |   AGPLv3   | Loki log collector (Deprecated)  |
|      [grafana-infinity-ds](https://github.com/grafana/grafana-infinity-datasource/)       |  3.7.0  | Apache-2.0 | versatile datasource             |
|  [grafana-plugins](https://github.com/pgsty/infra-pkg/tree/main/noarch/grafana-plugins)   | 12.3.0  | Apache-2.0 | extra panel & datasource plugins |


### Victoria Stack

|                                                 Name                                                  | Version |  License   | Comment                                        |
|:-----------------------------------------------------------------------------------------------------:|:-------:|:----------:|:-----------------------------------------------|
|                [victoria-metrics](https://github.com/VictoriaMetrics/VictoriaMetrics)                 | 1.133.0 | Apache-2.0 | VictoriaMetrics, Better Prometheus Alternative |
|                   [victoria-logs](https://github.com/VictoriaMetrics/VictoriaMetrics)                 | 1.43.1  | Apache-2.0 | VictoriaLogs, Better Logging platform          |
|                 [victoria-traces](https://github.com/VictoriaMetrics/VictoriaMetrics)                 |  0.5.1  | Apache-2.0 | VictoriaTraces, Better Tracing platform        |
|            [victoria-metrics-cluster](https://github.com/VictoriaMetrics/VictoriaMetrics)             | 1.133.0 | Apache-2.0 | Distributive version of VictoriaMetrics        |
|                     [vmutils](https://github.com/VictoriaMetrics/VictoriaMetrics)                     | 1.133.0 | Apache-2.0 | VictoriaMetrics Utils                          |
|                     [vlogscli](https://github.com/VictoriaMetrics/VictoriaMetrics)                    | 1.43.1  | Apache-2.0 | VictoriaLogs CLI Utils                         |
|                      [vlagent](https://github.com/VictoriaMetrics/VictoriaMetrics)                    | 1.43.1  | Apache-2.0 | VictoriaLogs Logging Agent                     |
|    [grafana-victorialogs-ds](https://github.com/VictoriaMetrics/victorialogs-datasource/releases/)    | 0.23.3  | Apache-2.0 | VictoriaLogs Datasource for Grafana            |
| [grafana-victoriametrics-ds](https://github.com/VictoriaMetrics/victoriametrics-datasource/releases/) | 0.20.0  | Apache-2.0 | VictoriaMetrics Datasource for Grafana         |


### Prometheus Stack

|                                 Name                                 | Version |  License   | Comment                           |
|:--------------------------------------------------------------------:|:-------:|:----------:|:----------------------------------|
|        [prometheus](https://github.com/prometheus/prometheus)        |  3.9.1  | Apache-2.0 | FOSS TSDB and monitoring platform |
|       [pushgateway](https://github.com/prometheus/pushgateway)       | 1.11.2  | Apache-2.0 | push metrics to prometheus        |
|      [alertmanager](https://github.com/prometheus/alertmanager)      | 0.30.0  | Apache-2.0 | manage alerting event             |
| [blackbox_exporter](https://github.com/prometheus/blackbox_exporter) | 0.27.0  | Apache-2.0 | send probe to endpoints           |


### Metric Exporters

|                                  Name                                   | Version |  License   | Comment                            |
|:-----------------------------------------------------------------------:|:-------:|:----------:|:-----------------------------------|
|           [pg_exporter](https://github.com/pgsty/pg_exporter)           |  1.1.2  | Apache-2.0 | Advanced Postgres Metrics Exporter |
|  [pgbackrest_exporter](https://github.com/woblerr/pgbackrest_exporter)  | 0.22.0  |    MIT     | expose pgbackrest metrics          |
|      [node_exporter](https://github.com/prometheus/node_exporter)       | 1.10.2  | Apache-2.0 | expose linux node metrics          |
|   [keepalived_exporter](https://github.com/mehdy/keepalived-exporter)   |  1.7.0  |  GPL-3.0   | expose keepalived/VIP metrics      |
| [nginx_exporter](https://github.com/nginxinc/nginx-prometheus-exporter) |  1.5.1  | Apache-2.0 | expose nginx metrics               |
|  [zfs_exporter](https://github.com/waitingsong/zfs_exporter/releases/)  |  3.8.1  |    MIT     | expose zfs metrics                 |
|    [mysqld_exporter](https://github.com/prometheus/mysqld_exporter)     | 0.18.0  | Apache-2.0 | expose mysql metrics               |
|      [redis_exporter](https://github.com/oliver006/redis_exporter)      | 1.80.1  |    MIT     | expose redis metrics               |
|      [kafka_exporter](https://github.com/danielqsj/kafka_exporter)      |  1.9.0  | Apache-2.0 | expose kafka metrics               |
|     [mongodb_exporter](https://github.com/percona/mongodb_exporter)     | 0.47.2  | Apache-2.0 | expose mongodb metrics             |
|                [mtail](https://github.com/google/mtail)                 |  3.0.8  | Apache-2.0 | tail log and generate metrics      |
|        [vector](https://github.com/vectordotdev/vector/releases)        | 0.52.0  |  MPL-2.0   | the versatile logging collector    |

### Object Storage

|                            Name                            |    Version     |  License   | Comment            |
|:----------------------------------------------------------:|:--------------:|:----------:|:-------------------|
|          [minio](https://github.com/minio/minio)           | 20251203120000 |   AGPLv3   | FOSS S3 Server     |
|            [mcli](https://github.com/minio/mc)             | 20250813083541 |   AGPLv3   | FOSS S3 Client     |
|        [rustfs](https://github.com/rustfs/rustfs)          |   alpha.80     | Apache-2.0 | FOSS MinIO, Alpha  |
| [garage](https://git.deuxfleurs.fr/Deuxfleurs/garage)      |     2.1.0      |  AGPL-3.0  | Lightweight S3     |
|    [seaweedfs](https://github.com/seaweedfs/seaweedfs)     |      4.06      | Apache-2.0 | S3 for small files |
|        [rclone](https://github.com/rclone/rclone/)         |     1.72.1     |    MIT     | S3 CLI             |
|         [restic](https://github.com/restic/restic)         |     0.18.1     |   BSD-2    | Backup tool        |
|       [juicefs](https://github.com/juicedata/juicefs)      |     1.3.1      | Apache-2.0 | FS over S3         |

### Databases

PostgreSQL related tools, DBMS, and other utils

|                           Name                            | Version |  License   | Comment                   |
|:---------------------------------------------------------:|:-------:|:----------:|:--------------------------|
|          [etcd](https://github.com/etcd-io/etcd)          |  3.6.7  | Apache-2.0 | Fault Tolerant DCS        |
|          [kafka](https://github.com/apache/kafka)         |  4.1.1  | Apache-2.0 | Message Queue             |
|        [duckdb](https://github.com/duckdb/duckdb)         |  1.4.3  |    MIT     | Embedded OLAP             |
|     [ferretdb](https://github.com/FerretDB/FerretDB)      |  2.7.0  | Apache-2.0 | MongoDB over PG           |
| [tigerbeetle](https://github.com/tigerbeetle/tigerbeetle) | 0.16.68 | Apache-2.0 | Financial OLTP            |
|     [IvorySQL](https://github.com/IvorySQL/IvorySQL)      |   5.1   | Apache-2.0 | Oracle Compatible PG 18.1 |

### Utils

Pig the package manager, PostgreSQL tools, and other database related utils

|                                         Name                                          | Version |  License   | Comment                                            |
|:-------------------------------------------------------------------------------------:|:-------:|:----------:|:---------------------------------------------------|
|                          [pig](https://github.com/pgsty/pig)                          |  0.9.1  | Apache-2.0 | The pg package manager                             |
|           [vip-manager](https://github.com/cybertec-postgresql/vip-manager)           |  4.0.0  |   BSD-2    | bind L2 vip to pg primary                          |
|                       [pgflo](https://github.com/pgflo/pg_flo)                        | 0.0.15  | Apache-2.0 | Stream, transform, and route PG data in real-time. |
|                 [pgschema](https://github.com/schemagood/pgschema)                    |  1.4.2  | Apache-2.0 | perform pg schema migration                        |
|          [pg_timetable](https://github.com/cybertec-postgresql/pg_timetable)          |  6.2.0  | PostgreSQL | Advanced scheduling for PostgreSQL                 |
|          [timescaledb-tools](https://github.com/timescale/timescaledb-tune)           | 0.18.1  | Apache-2.0 | optimize timescaledb params                        |
| [timescaledb-event-streamer](https://github.com/noctarius/timescaledb-event-streamer) | 0.20.0  | Apache-2.0 | CDC on timescaledb hypertable                      |
|                     [dblab](https://github.com/danvergara/dblab)                      | 0.34.2  |    MIT     | Versatile cli for multiple databases               |
|                   [sqlcmd](https://github.com/microsoft/go-sqlcmd)                    |  1.9.0  |    MIT     | cli for MS SQL Server (and babelfish)              |
|                        [pev2](https://github.com/dalibo/pev2)                         | 1.19.0  | PostgreSQL | PostgreSQL explain visualizer 2                    |
|                      [sealos](https://github.com/labring/sealos)                      |  5.0.1  | Apache-2.0 | Battery-Included Kubernetes distribution           |
|                     [vray](https://github.com/v2fly/v2ray-core)                       | 5.44.1  |    MIT     | Building proxies to bypass network restrictions.   |
|                 [postgrest](https://github.com/PostgREST/postgrest)                   |  14.3   |    MIT     | PostgreSQL RESTful API Server                      |
|              [npgsqlrest](https://github.com/vb-consulting/NpgsqlRest)                |  3.4.3  |    MIT     | .NET PostgreSQL REST API Generator                 |


### AI

AI Agents, MCP toolbox, coding IDEs, Python/Go/Node tools...

|                             Name                              |  Version  |   License   | Comment                               |
|:-------------------------------------------------------------:|:---------:|:-----------:|:--------------------------------------|
|        [claude](https://github.com/anthropics/claude-code)    |   2.1.9   | Proprietary | Claude Code - Anthropic Agentic Coder |
|       [opencode](https://github.com/opencode-ai/opencode)     |  1.1.23   |     MIT     | Terminal AI coding assistant          |
|      [code-server](https://github.com/coder/code-server)      |  4.108.0  |     MIT     | VS Code in the browser                |
|    [genai-toolbox](https://github.com/googleapis/genai-toolbox) |  0.25.0 | Apache-2.0  | Google MCP server for databases       |
|             [uv](https://github.com/astral-sh/uv)             |  0.9.26   |     MIT     | Modern Python package manager         |
|                 [golang](https://golang.org/)                 |  1.25.6   |    BSD-3    | Go Compiler                           |
|                [nodejs](https://nodejs.org/)                  |  24.12.0  |  MIT/Mixed  | Run Javascript on the server          |


> Hint: When using the victoria datasource for grafana, don't forget to set `allow_loading_unsigned_plugins = victoriametrics-logs-datasource,victoriametrics-metrics-datasource` in `/etc/grafana/grafana.ini`
