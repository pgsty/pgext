---
title: INFRA Changelog
description: The pigsty-infra repo change log and Observability packages release notes
weight: 400
---

Check [Infra Repo](/repo/infra) for usage instructions.

## 2026-02-08

| Name                       | Old            | New            | Comment |
|:---------------------------|:---------------|:---------------|:--------|
| alertmanager               | 0.30.1         | 0.31.0         |         |
| victoria-metrics           | 1.134.0        | 1.135.0        |         |
| victoria-metrics-cluster   | 1.134.0        | 1.135.0        |         |
| vmutils                    | 1.134.0        | 1.135.0        |         |
| victoria-logs              | 1.43.1         | 1.45.0         |         |
| vlagent                    | 1.43.1         | 1.45.0         |         |
| vlogscli                   | 1.43.1         | 1.45.0         |         |
| grafana-victorialogs-ds    | 0.23.5         | 0.24.1         |         |
| grafana-victoriametrics-ds | 0.20.1         | 0.21.0         |         |
| tigerbeetle                | 0.16.68        | 0.16.70        |         |
| loki                       | 3.1.1          | 3.6.5          |         |
| promtail                   | 3.0.0          | 3.6.5          |         |
| logcli                     | 3.1.1          | 3.6.5          |         |
| redis_exporter             | 1.80.1         | 1.80.2         |         |
| timescaledb-tools          | 0.18.1         | 0.18.2         |         |
| seaweedfs                  | 4.06           | 4.09           |         |
| rustfs                     | 1.0.0-alpha.80 | 1.0.0-alpha.82 |         |
| uv                         | 0.9.26         | 0.10.0         |         |
| garage                     | 2.1.0          | 2.2.0          |         |
| headscale                  | 0.27.1         | 0.28.0         |         |
| hugo                       | 0.154.5        | 0.155.2        |         |
| pev2                       | 1.20.0         | 1.20.1         |         |
| postgrest                  | 14.3           | 14.4           |         |
| npgsqlrest                 | 3.4.7          | 3.7.0          |         |
| opencode                   | 1.1.34         | 1.1.53         |         |
| golang                     | 1.25.6         | 1.25.7         |         |
| nodejs                     | 24.12.0        | 24.13.0        |         |
| claude                     | 2.1.19         | 2.1.37         |         |
| vector                     | 0.52.0         | 0.53.0         |         |
| code                       | 1.108.0        | 1.109.0        |         |
| code-server                | 4.108.0        | 4.108.2        |         |
| rclone                     | 1.72.1         | 1.73.0         |         |
| pg_exporter                | 1.1.2          | 1.2.0          |         |
| grafana                    | 12.3.1         | 12.3.2         |         |
| pig                        | 1.0.0          | 1.1.0          |         |
| cloudflared                | 2026.1.1       | 2026.2.0       |         |


## 2026-01-25

| Name                       | Old     | New      | Comment |
|:---------------------------|:--------|:---------|:--------|
| alertmanager               | 0.30.0  | 0.30.1   |         |
| victoria-metrics           | 1.133.0 | 1.134.0  |         |
| victoria-traces            | 0.5.1   | 0.7.1    |         |
| grafana-victorialogs-ds    | 0.23.3  | 0.23.5   |         |
| grafana-victoriametrics-ds | 0.20.0  | 0.20.1   |         |
| npgsqlrest                 | 3.4.3   | 3.4.7    |         |
| claude                     | 2.1.9   | 2.1.19   |         |
| opencode                   | 1.1.23  | 1.1.34   |         |
| caddy                      | -       | 2.10.2   | new     |
| hugo                       | -       | 0.154.5  | new     |
| cloudflared                | -       | 2026.1.1 | new     |
| headscale                  | -       | 0.27.1   | new     |
| pig                        | 0.9.0   | 1.0.0    |         |
| duckdb                     | 1.4.3   | 1.4.4    |         |

## 2026-01-16

| Name                       | Old      | New      | Comment |
|:---------------------------|:---------|:---------|:--------|
| prometheus                 | 3.8.1    | 3.9.1    |         |
| victoria-metrics           | 1.132.0  | 1.133.0  |         |
| tigerbeetle                | 0.16.65  | 0.16.68  |         |
| kafka                      | 4.0.0    | 4.1.1    |         |
| grafana-victoriametrics-ds | 0.19.7   | 0.20.0   |         |
| grafana-victorialogs-ds    | 0.23.2   | 0.23.3   |         |
| grafana-infinity-ds        | 3.6.0    | 3.7.0    |         |
| uv                         | 0.9.18   | 0.9.26   |         |
| seaweedfs                  | 4.01     | 4.06     |         |
| rustfs                     | alpha.71 | alpha.80 |         |
| v2ray                      | 5.28.0   | 5.44.1   |         |
| sqlcmd                     | 1.8.0    | 1.9.0    |         |
| opencode                   | 1.0.223  | 1.1.23   |         |
| claude                     | 2.1.1    | 2.1.9    |         |
| golang                     | 1.25.5   | 1.25.6   |         |
| asciinema                  | 3.0.1    | 3.1.0    |         |
| code                       | 1.107.0  | 1.108.0  |         |
| code-server                | 4.107.0  | 4.108.0  |         |
| npgsqlrest                 | 3.3.0    | 3.4.3    |         |
| genai-toolbox              | 0.24.0   | 0.25.0   |         |
| pg_exporter                | 1.1.1    | 1.1.2    |         |

## 2026-01-08

| 名称            | 旧版本      | 新版本       | 备注                     |
|:--------------|:---------|:----------|:-----------------------|
| pg_exporter   | `1.1.0`  | `1.1.1`   | new pg_timeline 采集器    |
| npgsqlrest    |          | `3.3.3`   | new                    |
| postgrest     |          | `14.3`    | new                    |
| opencode      |          | `1.0.223` | new                    |
| code-server   |          | `4.107.0` | new                    |
| claude        | `2.0.76` | `2.1.1`   | new                    |
| genai-toolbox | `0.23.0` | `0.24.0`  | remove oracle C driver |
| golang        |          | `1.25.5`  | new                    |
| nodejs        |          | `24.12.0` | new                    |


## 2025-12-25

| Name                | Old    | New    | Comment            |
|:--------------------|:-------|:-------|:-------------------|
| pig                 | 0.8.0  | 0.9.0  | routine update     |
| etcd                | 3.6.6  | 3.6.7  | routine update     |
| uv                  | -      | 0.9.18 | python pkg manager |
| ccm                 | -      | 2.0.76 | claude utils       |
| asciinema           |        | 3.0.1  | new package        |
| IvorySQL            | 5.0    | 5.1    |                    |
| grafana             | 12.3.0 | 12.3.1 |                    |
| vector              | 0.51.1 | 0.52.0 |                    |
| prometheus          | 3.8.0  | 3.8.1  |                    |
| alertmanager        | 0.29.0 | 0.30.0 |                    |
| victorialogs        | 1.41.0 | 1.43.1 |                    |
| pgbackrest_exporter | 0.21.0 | 0.22.0 |                    |


## 2025-12-16

| Name              | Old            | New            | Comment  |
|:------------------|:---------------|:---------------|:---------|
| victoria-metrics  | 1.131.0        | 1.132.0        |          |
| victoria-logs     | 1.40.0         | 1.41.0         |          |
| blackbox_exporter | 0.27.0         | 0.28.0         |          |
| duckdb            | 1.4.2          | 1.4.3          |          |
| rclone            | 1.72.0         | 1.72.1         |          |
| pev2              | 1.17.0         | 1.19.0         |          |
| pg_exporrter      | 1.0.3          | 1.1.0          |          |
| pig               | 0.7.4          | 0.8.0          |          |
| rclone            | 1.72.0         | 1.72.1         |          |
| genai-toolbox     | 0.22.0         | 0.23.0         |          |
| minio             | 20250907161309 | 20251203120000 | by pgsty |

## 2025-12-04

| Name                    | Old     | New       | Comment |
|:------------------------|:--------|:----------|:--------|
| rustfs                  | -       | 1.0.0-a71 | new     |
| seaweedfs               | -       | 4.1.0     | new     |
| garage                  | -       | 2.1.0     | new     |
| rclone                  | 1.71.2  | 1.72.0    |         |
| vector                  | 0.51.0  | 0.51.1    |         |
| prometheus              | 3.7.3   | 3.8.0     |         |
| victoria-metrics        | 0.130.0 | 0.131.0   |         |
| victoria-logs           | 0.38.0  | 0.40.0    |         |
| victoria-traces         | -       | 0.5.1     | new     |
| grafana-victorialogs-ds | 0.22.1  | 0.22.3    |         |
| redis_exporter          | 1.80.0  | 1.80.1    |         |
| mongodb_exporter        | 0.47.1  | 0.47.2    |         |
| genai-toolbox           | 0.21.0  | 0.22.0    |         |

## 2025-11-20

| Name                       | Old     | New     | Comment        |
|:---------------------------|:--------|:--------|:---------------|
| pgschema                   | -       | 1.4.2   | new            |
| pgflo                      | -       | 0.0.15  | new            |
| vector                     | 0.51.0  | 0.51.1  | bugfix release |
| sealos                     | 5.0.1   | 5.1.1   |                |
| etcd                       | 3.6.5   | 3.6.6   |                |
| duckdb                     | 1.4.1   | 1.4.2   |                |
| pg_exporter                | 1.0.2   | 1.0.3   |                |
| pig                        | 0.7.1   | 0.7.4   |                |
| grafana                    | 12.1.0  | 12.3.0  |                |
| pg_timetable               | 6.1.0   | 6.2.0   |                |
| genai-toolbox              | 0.16.0  | 0.21.0  |                |
| timescaledb-tools          | 0.18.0  | 0.18.1  | move to infra  |
| timescaledb-event-streamer | 0.12.0  | 0.20.0  |                |
| tigerbeetle                | 0.16.60 | 0.16.65 |                |
| victoria-metrics           | 1.129.1 | 1.130.0 |                |
| victorialogs               | 1.37.2  | 1.38.0  |                |
| grafana-victorialogs-ds    | 0.21.4  | 0.22.1  |                |
| grafana-victoriametrics-ds | 0.19.6  | 0.19.7  |                |
| grafana-plugins            | 12.0.0  | 12.3.0  |                |


## 2025-11-08

MinIO no longer provider binary packages, pigsty will take over it since this release.

| Name                       |   Old   |   New   | Comment |
|:---------------------------|:-------:|:-------:|:--------|
| prometheus                 |  3.6.0  |  3.7.3  |         |
| pushgateway                | 1.11.1  | 1.11.2  |         |
| alertmanager               | 0.28.1  | 0.29.0  |         |
| nginx_exporter             |  1.5.0  |  1.5.1  |         |
| node_exporter              |  1.9.1  | 1.10.2  |         |
| pgbackrest_exporter        | 0.20.0  | 0.21.0  |         |
| redis_exporter             | 1.77.0  | 1.80.0  |         |
| duckdb                     |  1.4.0  |  1.4.1  |         |
| dblab                      | 0.33.0  | 0.34.2  |         |
| pg_timetable               | 5.13.0  |  6.1.0  |         |
| vector                     | 0.50.0  | 0.51.0  |         |
| rclone                     | 1.71.1  | 1.71.2  |         |
| victoria-metrics           | 1.126.0 | 1.129.1 |         |
| victoria-logs              | 1.35.0  | 1.37.2  |         |
| grafana-victorialogs-ds    | 0.21.0  | 0.21.4  |         |
| grafana-victoriametrics-ds | 0.19.4  | 0.19.6  |         |
| grafana-infinity-ds        |  3.5.0  |  3.6.0  |         |
| pev2                       | 1.16.0  | 1.17.0  |         |
| pig                        |  0.6.2  |  0.7.1  |         |



## 2025-10-18

| Name                       |      Old       |      New       | Comment |
|:---------------------------|:--------------:|:--------------:|:--------|
| prometheus                 |     3.5.0      |     3.6.0      |         |
| nginx_exporter             |     1.4.2      |     1.5.0      |         |
| mysqld_exporter            |     0.17.2     |     0.18.0     |         |
| redis_exporter             |     1.75.0     |     1.77.0     |         |
| mongodb_exporter           |     0.47.0     |     0.47.1     |         |
| victoria-metrics           |    1.121.0     |    1.126.0     |         |
| vicotira-logs              |     1.25.1     |     1.35.0     |         |
| duckdb                     |     1.3.2      |     1.4.0      |         |
| etcd                       |     3.6.4      |     3.6.5      |         |
| restic                     |     0.18.0     |     0.18.1     |         |
| tigerbeetle                |    0.16.54     |    0.16.60     |         |
| grafana-victorialogs-ds    |     0.19.3     |     0.21.0     |         |
| grafana-victoriametrics-ds |     0.18.3     |     0.19.4     |         |
| grafana-infinity-ds        |     3.3.0      |     3.5.0      |         |
| genai-toolbox              |     0.9.0      |     0.16.0     |         |
| grafana                    |     12.1.0     |     12.2.0     |         |
| vector                     |     0.49.0     |     0.50.0     |         |
| rclone                     |     1.70.3     |     1.71.1     |         |
| minio                      | 20250723155402 | 20250907161309 |         |
| mcli                       | 20250721052808 | 20250813083541 |         |


## 2025-08-15

| Name                       |   Old   |   New   | Comment |
|:---------------------------|:-------:|:-------:|:--------|
| grafana                    | 12.0.0  | 12.1.0  |         |
| pg_exporter                |  1.0.1  |  1.0.2  |         |
| pig                        |  0.6.0  |  0.6.1  |         |
| vector                     | 0.48.0  | 0.49.0  |         |
| redis_exporter             | 1.74.0  | 1.75.0  |         |
| mongo_exporter             | 0.46.0  | 0.47.0  |         |
| victoriametrics            | 1.121.0 | 1.123.0 |         |
| victorialogs:              | 1.25.0  | 1.28.0  |         |
| grafana-victoriametrics-ds | 0.17.0  | 0.18.3  |         |
| grafana-victorialogs-ds    | 0.18.3  | 0.19.3  |         |
| grafana-infinity-ds        |  3.3.0  |  3.4.1  |         |
| etcd                       |  3.6.1  |  3.6.4  |         |
| ferretdb                   |  2.3.1  |  2.5.0  |         |
| tigerbeetle                | 0.16.50 | 0.16.54 |         |
| genai-toolbox              |  0.9.0  | 0.12.0  |         |


## 2025-07-24

| Name     | Old |         New          | Comment                    |
|:---------|:---:|:--------------------:|:---------------------------|
| FerretDB |  -  |        2.4.0         | work with documentdb 1.105 |
| etcd     |  -  |        3.6.3         |                            |
| minio    |  -  |    20250723155402    |                            |
| mcli     |  -  |    20250721052808    |                            |
| ivorysql |  -  | 4.5-0ffca11-20250709 | fix libxcrypt deps issue   |


## 2025-07-16

| Name                       |   Old   |   New   | Comment                      |
|:---------------------------|:-------:|:-------:|:-----------------------------|
| genai-toolbox              |  0.8.0  |  0.9.0  | MCP toolbox for various DBMS |
| victoriametrics            | 1.120.0 | 1.121.0 | split into various packages  |
| victorialogs               | 1.24.0  | 1.25.0  | split into various packages  |
| prometheus                 |  3.4.2  |  3.5.0  |                              |
| duckdb                     |  1.3.1  |  1.3.2  |                              |
| etcd                       |  3.6.1  |  3.6.2  |                              |
| tigerbeetle                | 0.16.48 | 0.16.50 |                              |
| grafana-victoriametrics-ds | 0.16.0  | 0.17.0  |                              |
| rclone                     | 1.69.3  | 1.70.3  |                              |
| pig                        |  0.5.0  |  0.6.0  |                              |
| pev2                       | 1.15.0  | 1.16.0  |                              |
| pg_exporter                |  1.0.0  |  1.0.1  |                              |


## 2025-07-04

| Name                       |   Old   |   New   | Comment |
|:---------------------------|:-------:|:-------:|:--------|
| prometheus                 |  3.4.1  |  3.4.2  | -       |
| grafana                    | 12.0.1  | 12.0.2  | -       |
| vector                     | 0.47.0  | 0.48.0  | -       |
| rclone                     | 1.69.0  | 1.70.2  | -       |
| vip-manager                |  3.0.0  |  4.0.0  | -       |
| blackbox_exporter          | 0.26.0  | 0.27.0  | -       |
| redis_exporter             | 1.72.1  | 1.74.0  | -       |
| duckdb                     |  1.3.0  |  1.3.1  | -       |
| etcd                       |  3.6.0  |  3.6.1  | -       |
| ferretdb                   |  2.2.0  |  2.3.1  | -       |
| dblab                      | 0.32.0  | 0.33.0  | -       |
| tigerbettle                | 0.16.41 | 0.16.48 | -       |
| grafana-victorialogs-ds    | 0.16.3  | 0.18.1  | -       |
| grafana-victoriametrics-ds | 0.15.1  | 0.16.0  | -       |
| grafana-inifinity-ds       |  3.2.1  |  3.3.0  | -       |
| victorialogs               | 1.22.2  | 1.24.0  | -       |
| victoriametrics            | 1.117.1 | 1.120.0 | -       |


## 2025-06-01

| Name                    | Old |   New   | Comment |
|:------------------------|:---:|:-------:|:--------|
| grafana                 |  -  | 12.0.1  | -       |
| prometheus              |  -  |  3.4.1  | -       |
| keepalived_exporter     |  -  |  1.7.0  | -       |
| redis_exporter          |  -  | 1.73.0  | -       |
| victoriametrics         |  -  | 1.118.0 | -       |
| victorialogs            |  -  | 1.23.1  | -       |
| tigerbeetle             |  -  | 0.16.42 | -       |
| grafana-victorialogs-ds |  -  | 0.17.0  | -       |
| grafana-infinity-ds     |  -  |  3.2.2  | -       |


## 2025-05-22

| Name                       | Old |      New       | Comment                         |
|:---------------------------|:---:|:--------------:|:--------------------------------|
| dblab                      |  -  |     0.32.0     | -                               |
| prometheus                 |  -  |     3.4.0      | -                               |
| duckdb                     |  -  |     1.3.0      | -                               |
| etcd                       |  -  |     3.6.0      | -                               |
| pg_exporter                |  -  |     1.0.0      | -                               |
| ferretdb                   |  -  |     2.2.0      | -                               |
| rclone                     |  -  |     1.69.3     | -                               |
| minio                      |  -  | 20250422221226 | The last version with admin GUI |
| mcli                       |  -  | 20250416181326 | -                               |
| nginx_exporter             |  -  |     1.4.2      | -                               |
| keepalived_exporter        |  -  |     1.6.2      | -                               |
| pgbackrest_exporter        |  -  |     0.20.0     | -                               |
| redis_exporter             |  -  |     1.27.1     | -                               |
| victoriametrics            |  -  |    1.117.1     | -                               |
| victorialogs               |  -  |     1.22.2     | -                               |
| pg_timetable               |  -  |     5.13.0     | -                               |
| tigerbeetle                |  -  |    0.16.41     | -                               |
| pev2                       |  -  |     1.15.0     | -                               |
| grafana                    |  -  |     12.0.0     | -                               |
| grafana-victorialogs-ds    |  -  |     0.16.3     | -                               |
| grafana-victoriametrics-ds |  -  |     0.15.1     | -                               |
| grafana-infinity-ds        |  -  |     3.2.1      | -                               |
| grafana_plugins            |  -  |     12.0.0     | -                               |


## 2025-04-23

| Name                | Old |      New       | Comment |
|:--------------------|:---:|:--------------:|:--------|
| mtail               |  -  |     3.0.8      | new     |
| pig                 |  -  |     0.4.0      | -       |
| pg_exporter         |  -  |     0.9.0      | -       |
| prometheus          |  -  |     3.3.0      | -       |
| pushgateway         |  -  |     1.11.1     | -       |
| keepalived_exporter |  -  |     1.6.0      | -       |
| redis_exporter      |  -  |     1.70.0     | -       |
| victoriametrics     |  -  |    1.115.0     | -       |
| victoria_logs       |  -  |     1.20.0     | -       |
| duckdb              |  -  |     1.2.2      | -       |
| pg_timetable        |  -  |     5.12.0     | -       |
| vector              |  -  |     0.46.1     | -       |
| minio               |  -  | 20250422221226 | -       |
| mcli                |  -  | 20250416181326 | -       |


## 2025-04-05

| Name             | Old |      New       | Comment |
|:-----------------|:---:|:--------------:|:--------|
| pig              |  -  |     0.3.4      | -       |
| etcd             |  -  |     3.5.21     | -       |
| restic           |  -  |     0.18.0     | -       |
| ferretdb         |  -  |     2.1.0      | -       |
| tigerbeetle      |  -  |    0.16.34     | -       |
| pg_exporter      |  -  |     0.8.1      | -       |
| node_exporter    |  -  |     1.9.1      | -       |
| grafana          |  -  |     11.6.0     | -       |
| zfs_exporter     |  -  |     3.8.1      | -       |
| mongodb_exporter |  -  |     0.44.0     | -       |
| victoriametrics  |  -  |    1.114.0     | -       |
| minio            |  -  | 20250403145628 | -       |
| mcli             |  -  | 20250403170756 | -       |


## 2025-03-23

| Name                | Old |  New   | Comment |
|:--------------------|:---:|:------:|:--------|
| etcd                |  -  | 3.5.20 | -       |
| pgbackrest_exporter |  -  | 0.19.0 | rebuild |
| victorialogs        |  -  | 1.17.0 | -       |
| vslogcli            |  -  | 1.17.0 | -       |


## 2025-03-17

| Name                       | Old |   New   | Comment |
|:---------------------------|:---:|:-------:|:--------|
| kafka                      |  -  |  4.0.0  | -       |
| Prometheus                 |  -  |  3.2.1  | -       |
| AlertManager               |  -  | 0.28.1  | -       |
| blackbox_exporter          |  -  | 0.26.0  | -       |
| node_exporter              |  -  |  1.9.0  | -       |
| mysqld_exporter            |  -  | 0.17.2  | -       |
| kafka_exporter             |  -  |  1.9.0  | -       |
| redis_exporter             |  -  | 1.69.0  | -       |
| DuckDB                     |  -  |  1.2.1  | -       |
| etcd                       |  -  | 3.5.19  | -       |
| FerretDB                   |  -  |  2.0.0  | -       |
| tigerbeetle                |  -  | 0.16.31 | -       |
| vector                     |  -  | 0.45.0  | -       |
| VictoriaMetrics            |  -  | 1.114.0 | -       |
| VictoriaLogs               |  -  | 1.16.0  | -       |
| rclone                     |  -  | 1.69.1  | -       |
| pev2                       |  -  | 1.14.0  | -       |
| grafana-victorialogs-ds    |  -  | 0.16.0  | -       |
| grafana-victoriametrics-ds |  -  | 0.14.0  | -       |
| grafana-infinity-ds        |  -  |  3.0.0  | -       |
| timescaledb-event-streamer |  -  | 0.12.0  | new     |
| restic                     |  -  | 0.17.3  | new     |
| juicefs                    |  -  |  1.2.3  | new     |

## 2025-02-12

| Name                |      Old       |      New       | Comment |
|:--------------------|:--------------:|:--------------:|:--------|
| pushgateway         |     1.10.0     |     1.11.0     | -       |
| alertmanager        |     0.27.0     |     0.28.0     | -       |
| nginx_exporter      |     1.4.0      |     1.4.1      | -       |
| pgbackrest_exporter |     0.18.0     |     0.19.0     | -       |
| redis_exporter      |     1.66.0     |     1.67.0     | -       |
| mongodb_exporter    |     0.43.0     |     0.43.1     | -       |
| VictoriaMetrics     |    1.107.0     |    1.111.0     | -       |
| VictoriaLogs        |     v1.3.2     |     1.9.1      | -       |
| DuckDB              |     1.1.3      |     1.2.0      | -       |
| Etcd                |     3.5.17     |     3.5.18     | -       |
| pg_timetable        |     5.10.0     |     5.11.0     | -       |
| FerretDB            |     1.24.0     |     2.0.0      | -       |
| tigerbeetle         |    0.16.13     |    0.16.27     | -       |
| grafana             |     11.4.0     |     11.5.1     | -       |
| vector              |     0.43.1     |     0.44.0     | -       |
| minio               | 20241218131544 | 20250207232109 | -       |
| mcli                | 20241121172154 | 20250208191421 | -       |
| rclone              |     1.68.2     |     1.69.0     | -       |


## 2024-11-19

| Name                |   Old   |   New   | Comment |
|:--------------------|:-------:|:-------:|:--------|
| Prometheus          | 2.54.0  |  3.0.0  | -       |
| VictoriaMetrics     | 1.102.1 | 1.106.1 | -       |
| VictoriaLogs        | v0.28.0 |  1.0.0  | -       |
| MySQL Exporter      | 0.15.1  | 0.16.0  | -       |
| Redis Exporter      | 1.62.0  | 1.66.0  | -       |
| MongoDB Exporter    | 0.41.2  | 0.42.0  | -       |
| Keepalived Exporter |  1.3.3  |  1.4.0  | -       |
| DuckDB              |  1.1.2  |  1.1.3  | -       |
| etcd                | 3.5.16  | 3.5.17  | -       |
| tigerbeetle         |  16.8   | 0.16.13 | -       |
| grafana             |    -    | 11.3.0  | -       |
| vector              |    -    | 0.42.0  | -       |
