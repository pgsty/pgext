## 用法

来源：

- [官方扩展控制文件](https://github.com/ibarwick/conninfo/blob/6eefa30a916761963dc2727a08bf6c958df5d8fe/conninfo.control)
- [官方上游文档](https://github.com/ibarwick/conninfo/blob/6eefa30a916761963dc2727a08bf6c958df5d8fe/README.md)

`conninfo` — 解析、校验 PostgreSQL libpq conninfo 连接字符串，并查询默认连接参数

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "conninfo";
SELECT extversion
FROM pg_extension
WHERE extname = 'conninfo';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
