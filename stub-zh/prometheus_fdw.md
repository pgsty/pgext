## 用法

来源：

- [官方文档](https://github.com/Jayko001/prometheus_fdw/blob/d9794a9b89bf00e65bf9d31475b740d7f5fdf9f0/docs/index.md)
- [官方 FDW 实现](https://github.com/Jayko001/prometheus_fdw/blob/d9794a9b89bf00e65bf9d31475b740d7f5fdf9f0/src/lib.rs)
- [官方 PGXN 版本](https://pgxn.org/dist/prometheus_fdw/0.2.1/)

`prometheus_fdw` 是一个只读外部数据包装器，把 Prometheus 范围查询响应转换成 PostgreSQL 行。版本 `0.2.1` 只支持一个 `metrics` 对象，并要求精确的指标名与时间边界谓词才能构造 Prometheus 请求。

### 配置与查询

```sql
CREATE EXTENSION prometheus_fdw;

CREATE FOREIGN DATA WRAPPER prometheus_wrapper
  HANDLER prometheus_fdw_handler
  VALIDATOR prometheus_fdw_validator;

CREATE SERVER my_prometheus_server
  FOREIGN DATA WRAPPER prometheus_wrapper
  OPTIONS (base_url 'https://prometheus.example.com');

CREATE FOREIGN TABLE metrics (
  metric_name text,
  metric_labels jsonb,
  metric_time bigint,
  metric_value float8
)
SERVER my_prometheus_server
OPTIONS (object 'metrics', step '10m');

SELECT metric_name, metric_labels, metric_time, metric_value
FROM metrics
WHERE metric_name = 'container_cpu_usage_seconds_total'
  AND metric_time > 1696046800
  AND metric_time < 1696133000;
```

实现实际要求这三个谓词：`metric_name =`、`metric_time >` 和 `metric_time <`。缺少任意一个都无法产生有效请求。`step` 会传给 Prometheus，默认值为 `10m`。

### 返回列

- `metric_name` 重复请求的指标名。
- `metric_labels` 包含 Prometheus 序列标签对象。
- `metric_time` 是范围查询返回的 Unix 时间戳。
- `metric_value` 是解析为 `float8` 的样本值。

### 运维边界

每次扫描都会发出一个超时 30 秒的同步外部 HTTP 请求，并在后端内存中物化所有返回样本；SQL 排序、限制和多数过滤条件不会下推。必须限制端点、查询窗口、步长、行数、语句超时和网络出口。虽然源码定义了用户名、密码和令牌字段，本次复核的构造器并不会从服务器或用户映射选项填充这些字段，因此在验证打包构建前，不能假定需要认证的 Prometheus 可用。实现还会直接解包 JSON 解析结果，也不会显式拒绝 HTTP 错误状态。生产使用前，应测试损坏/错误响应、TLS 信任、规划器重复扫描、并行查询、时间戳边界语义、大规模序列集与凭据处理。
