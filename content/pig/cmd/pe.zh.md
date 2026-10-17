---
title: "pig pe"
description: "读取 PostgreSQL 指标并重载 pg_exporter 配置"
weight: 675
icon: fas fa-chart-line
---

使用 `pig pe` 读取运行中的 [pg_exporter](https://github.com/Vonng/pg_exporter) 指标，或重载其配置。
完整命令名为 `pig pg_exporter`，别名为 `pe`、`pgexp` 和 `pgexporter`。

```bash
pig pe get                     # 查看 PostgreSQL 指标
pig pe list                    # 列出 PostgreSQL 指标族
pig pe stat                    # 查看 exporter 统计信息
pig pe reload                  # 重载 exporter 配置
```

## 选择 exporter

默认端点为 `http://127.0.0.1:9630`。使用以下参数选择其他实例：

| 参数 | 默认值 | 说明 |
|:---|:---|:---|
| `--host` | `127.0.0.1` | 主机名或 IPv4 地址 |
| `-p, --port` | `9630` | HTTP 端口 |
{.full-width}

```bash
pig pe --host pg-meta -p 9630 stat
pig pe --host 10.10.10.10 get
```

PIG 使用 HTTP 和标准的 `/metrics`、`/stat`、`/reload` 路径。
需要 HTTPS、身份认证、客户端证书或自定义路径时，使用 `curl` 等 HTTP 客户端。

## pe get

读取 `/metrics`，输出名称以 `pg_` 开头的指标样本，以及对应的 Prometheus `HELP` 和 `TYPE` 行。

```bash
pig pe get
pig pe --host pg-meta get
```

要同时读取非 `pg_` 开头的进程与运行时指标，直接访问完整端点：

```bash
curl -fsS http://127.0.0.1:9630/metrics
```

## pe list

读取 `/metrics`，为响应中的每个 `pg_` 指标族输出一条 `HELP` 行。
列表反映 exporter 的当前配置。

```bash
pig pe list
```

## pe stat

读取 `/stat` 并输出响应。响应格式取决于已安装的 pg_exporter 版本。

```bash
pig pe stat
```

## pe reload

修改 exporter 配置后，通过 `/reload` 请求重载：

```bash
pig pe reload
```

## 输出与错误

使用 `-o json` 或 `-o yaml` 获取包含命令文本输出的 PIG 结果。
Prometheus 指标在结果中仍以文本保存。

网络错误、响应读取失败及 HTTP 4xx/5xx 都返回非零退出码。
HTTP 错误详情只包含状态码与标准说明，不包含响应体；正常响应保留上述筛选或原文输出行为。

使用 `pig ext add pg-exporter` 安装 pg_exporter。配置并启动 exporter 后，再使用 `pig pe` 访问其 HTTP API。
