## 用法

来源：

- [官方扩展控制文件](https://github.com/Octonica/data_rig/blob/e099c8f3a7a4a471a448c4574b7d5049ef7617c3/data_rig.control)
- [官方上游文档](https://github.com/Octonica/data_rig/blob/e099c8f3a7a4a471a448c4574b7d5049ef7617c3/README.md)

`data_rig` — data_rig：提供用于多维 OLAP 的 fact 数据类型和 GiST 索引支持。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`cube`。

```sql
CREATE EXTENSION "data_rig";
SELECT extversion
FROM pg_extension
WHERE extname = 'data_rig';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
