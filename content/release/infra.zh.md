---
title: INFRA 变更日志
description: pigsty-infra 仓库软件变更日志
weight: 400
---


## 2025-10-18

| 名称                         |      旧版本       |      新版本       | 备注 |
|:---------------------------|:--------------:|:--------------:|:---|
| prometheus                 |     3.5.0      |     3.6.0      |    |
| nginx_exporter             |     1.4.2      |     1.5.0      |    |
| mysqld_exporter            |     0.17.2     |     0.18.0     |    |
| redis_exporter             |     1.75.0     |     1.77.0     |    |
| mongodb_exporter           |     0.47.0     |     0.47.1     |    |
| victoria-metrics           |    1.121.0     |    1.126.0     |    |
| vicotira-logs              |     1.25.1     |     1.35.0     |    |
| duckdb                     |     1.3.2      |     1.4.0      |    |
| etcd                       |     3.6.4      |     3.6.5      |    |
| restic                     |     0.18.0     |     0.18.1     |    |
| tigerbeetle                |    0.16.54     |    0.16.60     |    |
| grafana-victorialogs-ds    |     0.19.3     |     0.21.0     |    |
| grafana-victoriametrics-ds |     0.18.3     |     0.19.4     |    |
| grafana-infinity-ds        |     3.3.0      |     3.5.0      |    |
| genai-toolbox              |     0.9.0      |     0.16.0     |    |
| grafana                    |     12.1.0     |     12.2.0     |    |
| vector                     |     0.49.0     |     0.50.0     |    |
| rclone                     |     1.70.3     |     1.71.1     |    |
| minio                      | 20250723155402 | 20250907161309 |    |
| mcli                       | 20250721052808 | 20250813083541 |    |


## 2025-08-15

| 名称                         |   旧版本   |   新版本   | 备注 |
|:---------------------------|:-------:|:-------:|:---|
| grafana                    | 12.0.0  | 12.1.0  |    |
| pg_exporter                |  1.0.1  |  1.0.2  |    |
| pig                        |  0.6.0  |  0.6.1  |    |
| vector                     | 0.48.0  | 0.49.0  |    |
| redis_exporter             | 1.74.0  | 1.75.0  |    |
| mongo_exporter             | 0.46.0  | 0.47.0  |    |
| victoriametrics            | 1.121.0 | 1.123.0 |    |
| victorialogs:              | 1.25.0  | 1.28.0  |    |
| grafana-victoriametrics-ds | 0.17.0  | 0.18.3  |    |
| grafana-victorialogs-ds    | 0.18.3  | 0.19.3  |    |
| grafana-infinity-ds        |  3.3.0  |  3.4.1  |    |
| etcd                       |  3.6.1  |  3.6.4  |    |
| ferretdb                   |  2.3.1  |  2.5.0  |    |
| tigerbeetle                | 0.16.50 | 0.16.54 |    |
| genai-toolbox              |  0.9.0  | 0.12.0  |    |


## 2025-07-24

| 名称       | 旧版本 |         新版本          | 备注                         |
|:---------|:---:|:--------------------:|:---------------------------|
| FerretDB |  -  |        2.4.0         | work with documentdb 1.105 |
| etcd     |  -  |        3.6.3         |                            |
| minio    |  -  |    20250723155402    |                            |
| mcli     |  -  |    20250721052808    |                            |
| ivorysql |  -  | 4.5-0ffca11-20250709 | fix libxcrypt deps issue   |


## 2025-07-16

| 名称                         |   旧版本   |   新版本   | 备注                           |
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

| 名称                         |   旧版本   |   新版本   | 备注 |
|:---------------------------|:-------:|:-------:|:---|
| prometheus                 |  3.4.1  |  3.4.2  | -  |
| grafana                    | 12.0.1  | 12.0.2  | -  |
| vector                     | 0.47.0  | 0.48.0  | -  |
| rclone                     | 1.69.0  | 1.70.2  | -  |
| vip-manager                |  3.0.0  |  4.0.0  | -  |
| blackbox_exporter          | 0.26.0  | 0.27.0  | -  |
| redis_exporter             | 1.72.1  | 1.74.0  | -  |
| duckdb                     |  1.3.0  |  1.3.1  | -  |
| etcd                       |  3.6.0  |  3.6.1  | -  |
| ferretdb                   |  2.2.0  |  2.3.1  | -  |
| dblab                      | 0.32.0  | 0.33.0  | -  |
| tigerbettle                | 0.16.41 | 0.16.48 | -  |
| grafana-victorialogs-ds    | 0.16.3  | 0.18.1  | -  |
| grafana-victoriametrics-ds | 0.15.1  | 0.16.0  | -  |
| grafana-inifinity-ds       |  3.2.1  |  3.3.0  | -  |
| victorialogs               | 1.22.2  | 1.24.0  | -  |
| victoriametrics            | 1.117.1 | 1.120.0 | -  |


## 2025-06-01

| 名称                      | 旧版本 |   新版本   | 备注 |
|:------------------------|:---:|:-------:|:---|
| grafana                 |  -  | 12.0.1  | -  |
| prometheus              |  -  |  3.4.1  | -  |
| keepalived_exporter     |  -  |  1.7.0  | -  |
| redis_exporter          |  -  | 1.73.0  | -  |
| victoriametrics         |  -  | 1.118.0 | -  |
| victorialogs            |  -  | 1.23.1  | -  |
| tigerbeetle             |  -  | 0.16.42 | -  |
| grafana-victorialogs-ds |  -  | 0.17.0  | -  |
| grafana-infinity-ds     |  -  |  3.2.2  | -  |


## 2025-05-22

| 名称                         | 旧版本 |      新版本       | 备注            |
|:---------------------------|:---:|:--------------:|:--------------|
| dblab                      |  -  |     0.32.0     | -             |
| prometheus                 |  -  |     3.4.0      | -             |
| duckdb                     |  -  |     1.3.0      | -             |
| etcd                       |  -  |     3.6.0      | -             |
| pg_exporter                |  -  |     1.0.0      | -             |
| ferretdb                   |  -  |     2.2.0      | -             |
| rclone                     |  -  |     1.69.3     | -             |
| minio                      |  -  | 20250422221226 | 最后一个带 GUI 的版本 |
| mcli                       |  -  | 20250416181326 | -             |
| nginx_exporter             |  -  |     1.4.2      | -             |
| keepalived_exporter        |  -  |     1.6.2      | -             |
| pgbackrest_exporter        |  -  |     0.20.0     | -             |
| redis_exporter             |  -  |     1.27.1     | -             |
| victoriametrics            |  -  |    1.117.1     | -             |
| victorialogs               |  -  |     1.22.2     | -             |
| pg_timetable               |  -  |     5.13.0     | -             |
| tigerbeetle                |  -  |    0.16.41     | -             |
| pev2                       |  -  |     1.15.0     | -             |
| grafana                    |  -  |     12.0.0     | -             |
| grafana-victorialogs-ds    |  -  |     0.16.3     | -             |
| grafana-victoriametrics-ds |  -  |     0.15.1     | -             |
| grafana-infinity-ds        |  -  |     3.2.1      | -             |
| grafana_plugins            |  -  |     12.0.0     | -             |


## 2025-04-23

| 名称                  | 旧版本 |      新版本       | 备注 |
|:--------------------|:---:|:--------------:|:---|
| mtail               |  -  |     3.0.8      | 新增 |
| pig                 |  -  |     0.4.0      | -  |
| pg_exporter         |  -  |     0.9.0      | -  |
| prometheus          |  -  |     3.3.0      | -  |
| pushgateway         |  -  |     1.11.1     | -  |
| keepalived_exporter |  -  |     1.6.0      | -  |
| redis_exporter      |  -  |     1.70.0     | -  |
| victoriametrics     |  -  |    1.115.0     | -  |
| victoria_logs       |  -  |     1.20.0     | -  |
| duckdb              |  -  |     1.2.2      | -  |
| pg_timetable        |  -  |     5.12.0     | -  |
| vector              |  -  |     0.46.1     | -  |
| minio               |  -  | 20250422221226 | -  |
| mcli                |  -  | 20250416181326 | -  |


## 2025-04-05

| 名称               | 旧版本 |      新版本       | 备注 |
|:-----------------|:---:|:--------------:|:---|
| pig              |  -  |     0.3.4      | -  |
| etcd             |  -  |     3.5.21     | -  |
| restic           |  -  |     0.18.0     | -  |
| ferretdb         |  -  |     2.1.0      | -  |
| tigerbeetle      |  -  |    0.16.34     | -  |
| pg_exporter      |  -  |     0.8.1      | -  |
| node_exporter    |  -  |     1.9.1      | -  |
| grafana          |  -  |     11.6.0     | -  |
| zfs_exporter     |  -  |     3.8.1      | -  |
| mongodb_exporter |  -  |     0.44.0     | -  |
| victoriametrics  |  -  |    1.114.0     | -  |
| minio            |  -  | 20250403145628 | -  |
| mcli             |  -  | 20250403170756 | -  |


## 2025-03-23

| 名称                  | 旧版本 |  新版本   | 备注      |
|:--------------------|:---:|:------:|:--------|
| etcd                |  -  | 3.5.20 | -       |
| pgbackrest_exporter |  -  | 0.19.0 | rebuild |
| victorialogs        |  -  | 1.17.0 | -       |
| vslogcli            |  -  | 1.17.0 | -       |


## 2025-03-17

| 名称                         | 旧版本 |   新版本   | 备注  |
|:---------------------------|:---:|:-------:|:----|
| kafka                      |  -  |  4.0.0  | -   |
| Prometheus                 |  -  |  3.2.1  | -   |
| AlertManager               |  -  | 0.28.1  | -   |
| blackbox_exporter          |  -  | 0.26.0  | -   |
| node_exporter              |  -  |  1.9.0  | -   |
| mysqld_exporter            |  -  | 0.17.2  | -   |
| kafka_exporter             |  -  |  1.9.0  | -   |
| redis_exporter             |  -  | 1.69.0  | -   |
| DuckDB                     |  -  |  1.2.1  | -   |
| etcd                       |  -  | 3.5.19  | -   |
| FerretDB                   |  -  |  2.0.0  | -   |
| tigerbeetle                |  -  | 0.16.31 | -   |
| vector                     |  -  | 0.45.0  | -   |
| VictoriaMetrics            |  -  | 1.114.0 | -   |
| VictoriaLogs               |  -  | 1.16.0  | -   |
| rclone                     |  -  | 1.69.1  | -   |
| pev2                       |  -  | 1.14.0  | -   |
| grafana-victorialogs-ds    |  -  | 0.16.0  | -   |
| grafana-victoriametrics-ds |  -  | 0.14.0  | -   |
| grafana-infinity-ds        |  -  |  3.0.0  | -   |
| timescaledb-event-streamer |  -  | 0.12.0  | new |
| restic                     |  -  | 0.17.3  | new |
| juicefs                    |  -  |  1.2.3  | new |

## 2025-02-12

| 名称                  |      旧版本       |      新版本       | 备注 |
|:--------------------|:--------------:|:--------------:|:---|
| pushgateway         |     1.10.0     |     1.11.0     | -  |
| alertmanager        |     0.27.0     |     0.28.0     | -  |
| nginx_exporter      |     1.4.0      |     1.4.1      | -  |
| pgbackrest_exporter |     0.18.0     |     0.19.0     | -  |
| redis_exporter      |     1.66.0     |     1.67.0     | -  |
| mongodb_exporter    |     0.43.0     |     0.43.1     | -  |
| VictoriaMetrics     |    1.107.0     |    1.111.0     | -  |
| VictoriaLogs        |     v1.3.2     |     1.9.1      | -  |
| DuckDB              |     1.1.3      |     1.2.0      | -  |
| Etcd                |     3.5.17     |     3.5.18     | -  |
| pg_timetable        |     5.10.0     |     5.11.0     | -  |
| FerretDB            |     1.24.0     |     2.0.0      | -  |
| tigerbeetle         |    0.16.13     |    0.16.27     | -  |
| grafana             |     11.4.0     |     11.5.1     | -  |
| vector              |     0.43.1     |     0.44.0     | -  |
| minio               | 20241218131544 | 20250207232109 | -  |
| mcli                | 20241121172154 | 20250208191421 | -  |
| rclone              |     1.68.2     |     1.69.0     | -  |


## 2024-11-19

| 名称                  |   旧版本   |   新版本   | 备注 |
|:--------------------|:-------:|:-------:|:---|
| Prometheus          | 2.54.0  |  3.0.0  | -  |
| VictoriaMetrics     | 1.102.1 | 1.106.1 | -  |
| VictoriaLogs        | v0.28.0 |  1.0.0  | -  |
| MySQL Exporter      | 0.15.1  | 0.16.0  | -  |
| Redis Exporter      | 1.62.0  | 1.66.0  | -  |
| MongoDB Exporter    | 0.41.2  | 0.42.0  | -  |
| Keepalived Exporter |  1.3.3  |  1.4.0  | -  |
| DuckDB              |  1.1.2  |  1.1.3  | -  |
| etcd                | 3.5.16  | 3.5.17  | -  |
| tigerbeetle         |  16.8   | 0.16.13 | -  |
| grafana             |    -    | 11.3.0  | -  |
| vector              |    -    | 0.42.0  | -  |
