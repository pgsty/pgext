


## 用法

> [omni_polyfill: Postgres API 填充](https://docs.omnigres.org/omni_polyfill/polyfills/)

`omni_polyfill` 扩展为旧版 PostgreSQL 中可能不可用的函数提供填充实现。

### 配置

要使用填充，需将 `search_path` 设置为将 `omni_polyfill` 置于 `pg_catalog` **之前**：

```sql
SET search_path TO '$user', public, omni_polyfill, pg_catalog;
```

此顺序对于正确的函数解析和填充优先级至关重要。

### 可用的填充

- **`trim_array`** -- 数组裁剪功能
- **UUIDv7** -- UUID 版本 7 生成及相关工具
