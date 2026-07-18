## 用法

来源：

- [官方扩展控制文件](https://github.com/pgoracle/pgora-builtin/blob/d7f836d619031a5d180b5c8b48b68994ff8a555c/pgorafunc.control)
- [官方上游文档](https://github.com/pgoracle/pgora-builtin/blob/d7f836d619031a5d180b5c8b48b68994ff8a555c/README.md)

`pgorafunc` — 为 PostgreSQL 提供 Oracle 兼容函数与包的扩展。

已复核目录快照记录的版本为 `9.4`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pgorafunc";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgorafunc';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
