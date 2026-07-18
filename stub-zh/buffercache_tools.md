## 用法

来源：

- [官方扩展控制文件](https://github.com/04ina/buffercache_tools/blob/90f4edbc7ba3fb84dd43853e0cf030beb46f226a/buffercache_tools.control)
- [官方上游 README](https://github.com/04ina/buffercache_tools/blob/90f4edbc7ba3fb84dd43853e0cf030beb46f226a/README.md)

`buffercache_tools` — 检查并操作 PostgreSQL 缓冲区缓存中属于关系、分支、数据库和表空间的缓存项。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "buffercache_tools";
SELECT extversion
FROM pg_extension
WHERE extname = 'buffercache_tools';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
