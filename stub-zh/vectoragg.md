## 用法

来源：

- [官方扩展控制文件](https://github.com/cybertec-postgresql/vectoragg/blob/ffbd58f1c355a1302df9ec1b5a65caad812df5df/vectoragg.control)

`vectoragg` — 浮点数组区间求和、裁剪与降采样函数

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "vectoragg";
SELECT extversion
FROM pg_extension
WHERE extname = 'vectoragg';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
