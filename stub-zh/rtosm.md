## 用法

来源：

- [官方扩展控制文件](https://github.com/rtosm/rtosm/blob/fb8fbec1cfb127f6ea07abc7db3f79b9a2b2a7f7/rtosm.control)
- [官方上游文档](https://github.com/rtosm/rtosm/blob/fb8fbec1cfb127f6ea07abc7db3f79b9a2b2a7f7/README.md)

`rtosm` — 面向 OpenStreetMap API 数据库的空间对象实时简化扩展。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`intarray`。

```sql
CREATE EXTENSION "rtosm";
SELECT extversion
FROM pg_extension
WHERE extname = 'rtosm';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
