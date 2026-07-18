## 用法

来源：

- [固定提交的上游 README](https://github.com/alesharik/prometheus-fdw/blob/df0989e406128c3d7bab65a86f4004b8314a5956/README.md)
- [固定提交的 Cargo manifest](https://github.com/alesharik/prometheus-fdw/blob/df0989e406128c3d7bab65a86f4004b8314a5956/Cargo.toml)
- [固定提交的扩展 control 文件](https://github.com/alesharik/prometheus-fdw/blob/df0989e406128c3d7bab65a86f4004b8314a5956/prometheusfdw.control)

`prometheusfdw` 是 Rust/pgrx 外部数据包装器，用于从 Prometheus 兼容 HTTP 端点读取即时查询结果。扩展会安装 handler 与 validator 函数；wrapper、server 和 foreign table 需要显式创建。

```sql
CREATE EXTENSION prometheusfdw;

CREATE FOREIGN DATA WRAPPER prometheus_wrapper
  HANDLER prometheus_fdw_handler
  VALIDATOR prometheus_fdw_validator;

CREATE SERVER metrics_server
  FOREIGN DATA WRAPPER prometheus_wrapper
  OPTIONS (address 'http://prometheus:9090/');

CREATE FOREIGN TABLE scrape_duration_seconds (
  __name__ text,
  instance text,
  job text,
  value double precision
)
SERVER metrics_server
OPTIONS (query 'scrape_duration_seconds');

SELECT * FROM scrape_duration_seconds LIMIT 20;
```

foreign-table 选项 `query` 保存 PromQL。README 记录了从 `WHERE` qualifier 替换 `${var}`、以列 `time timestamp` 指定 instant-query 时间，以及面向 VictoriaMetrics 的可选选项 `rate`。range query 仍明确未完成，因此不能从 FDW 接口推断其支持历史时间窗口。

固定提交的 Cargo manifest 与生成的 control 版本线报告 `0.0.2`，feature 覆盖 PostgreSQL 14、15 和 16。其他 release packaging 使用不同版本线，因此必须核实准确 artifact。应限制 server-definition 权限，优先使用带认证的 TLS 端点，设置 query 与 statement timeout，并测试 label-to-column mapping、missing sample、cardinality、remote error 和 planner behavior。
