## 用法

来源：

- [Official README](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/README.md)
- [Version 1.0 SQL API](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/pgotel--1.0.sql)
- [GUC and SQL-call implementation](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/pgotel.cpp)
- [OTLP gRPC metric exporter implementation](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/metrics_grpc.cpp)
- [Native link dependencies](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/Makefile)

`pgotel` 1.0 是连接 OpenTelemetry C++ SDK 的 C/C++ 概念验证桥接层。当前 SQL 接口通过 OTLP gRPC exporter 发送 double-counter 指标；尽管项目描述更宽泛，它并没有暴露 trace 或 log API。

### 配置并发送计数器

本地库链接 OpenTelemetry metrics SDK 和 OTLP gRPC exporter。应在某个后端首次加载库之前配置 collector endpoint，再创建扩展并发送非负 counter 值：

```conf
pgotel.endpoint = 'localhost:4317'
pgotel.interval = 2000
pgotel.timeout = 500
```

```sql
CREATE EXTENSION pgotel;

SELECT pgotel_counter('requests', 1.0);
SELECT pgotel_counter(
  'rows_processed',
  42.0,
  '{"database":"app","worker":"import-1"}'::jsonb
);
```

两个重载分别是 `pgotel_counter(text, float8)` 和 `pgotel_counter(text, float8, jsonb)`。labels 应为非空、扁平且值为字符串的 JSON 对象。

### 当前实现边界

README 讨论选择 HTTP transport，并展示把 endpoint 作为参数的三参数调用；但已复核代码实际链接并使用 OTLP gRPC metrics exporter，endpoint 也是 GUC 而非函数参数。1.0 版本应以 SQL 和 C++ 实现为准。

`pgotel.enabled` 虽被注册，却没有在指标路径中读取，因此该源码中修改它不会停止发送。`pgotel.endpoint` 没有赋值钩子，应在库首次初始化前设置；修改 `pgotel.interval` 或 `pgotel.timeout` 会重建每后端的指标 provider。

labels 迭代假定所有值都是字符串，无法安全处理任意嵌套、数字、布尔或 null JSON；空 labels 也会被拒绝。counter 名称和 labels 可能产生无界指标基数，而且每次调用都会新建 counter instrument；两者都应限制为规模很小的受控词表。

指标导出属于异步外部 I/O，不具备事务性：即使外围 SQL 事务回滚，指标仍可能已导出；后端在周期导出前退出则可能丢失。扩展还要求兼容的 C++ ABI 及 OpenTelemetry/gRPC 共享库。应固定依赖，测试 collector 故障和后端关闭，并只把它用于实验遥测而非持久计量。
