## 用法

来源：

- [官方 README](https://github.com/alesharik/vmagent-fdw/blob/1ef48ceb1e9fa6d27fc22a668d9385ffb834caea/README.md)
- [官方 FDW 实现](https://github.com/alesharik/vmagent-fdw/blob/1ef48ceb1e9fa6d27fc22a668d9385ffb834caea/src/fdw.rs)
- [官方 HTTP 客户端实现](https://github.com/alesharik/vmagent-fdw/blob/1ef48ceb1e9fa6d27fc22a668d9385ffb834caea/src/client.rs)

`vmagentfdw` 是一个只读 Rust 外部数据包装器，可将 vmagent 的活跃与已移除抓取目标呈现为 PostgreSQL 记录。它会在扫描外部表时同步读取 vmagent `/api/v1/targets` 端点，因此远端延迟和可用性会直接影响执行查询的后端进程。

### 核心流程

安装扩展、注册处理器与验证器，并让外部服务器指向受信任的 vmagent 端点：

```sql
CREATE EXTENSION vmagentfdw;

CREATE FOREIGN DATA WRAPPER vmagent_wrapper
  HANDLER vmagent_fdw_handler
  VALIDATOR vmagent_fdw_validator;

CREATE SERVER local_vmagent
  FOREIGN DATA WRAPPER vmagent_wrapper
  OPTIONS (address 'http://127.0.0.1:8429');
```

只定义查询需要的字段；字段名决定包装器如何映射取值：

```sql
CREATE FOREIGN TABLE vmagent_targets (
  state text,
  health text,
  last_samples_scraped bigint,
  last_scrape_duration double precision,
  last_scrape timestamp,
  last_error text,
  scrape_url text,
  scrape_pool text,
  labels jsonb,
  discovered_labels jsonb,
  labels_job text
)
SERVER local_vmagent;

SELECT state, health, scrape_url, labels_job
FROM vmagent_targets;
```

### 字段映射

- `state` 根据目标所在的响应分区合成为 `active` 或 `dropped`。
- `health`、`last_samples_scraped`、`last_scrape_duration`、`last_scrape`、`last_error`、`scrape_url` 与 `scrape_pool` 映射对应的目标字段。
- `labels` 与 `discovered_labels` 以 JSONB 返回完整标签映射。
- 名为 `labels_<key>` 的字段会投影目标标签映射中的单个条目；无法识别的字段返回 NULL。

### 运维注意事项

实现会在每次扫描时获取全部活跃和已移除目标，不会把 PostgreSQL 的过滤、排序或限制条件下推到 HTTP 请求，也没有实现写入回调。复核的源码声明支持为 PostgreSQL `14`、`15` 与 `16` 构建，但实际兼容性仍取决于已安装的二进制。应限制创建或查询服务器的权限，使用有界且可信的端点，并设置合适的语句与网络超时。文档化的选项只有端点 `address`；认证与 TLS 策略需要由所选 URL 及外围网络控制承担。
