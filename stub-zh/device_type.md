## 用法

来源：

- [官方扩展控制文件](https://github.com/adjust/pg-device_type/blob/2d19dcc0ab9df0808579aff22bd412e44c64741b/device_type.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/device_type/)

`device_type` — device_type：提供用于移动设备类型的固定字节枚举式数据类型、比较操作符和 btree/hash 操作符类。

已复核目录快照记录的版本为 `0.0.2`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "device_type";
SELECT extversion
FROM pg_extension
WHERE extname = 'device_type';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
