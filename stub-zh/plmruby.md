## 用法

来源：

- [官方扩展控制文件](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/plmruby.control)
- [官方上游文档](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/README.md)

`plmruby` — 在嵌入式 mruby 运行时中执行标量、集合返回、内联与触发器函数

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "plmruby";
SELECT extversion
FROM pg_extension
WHERE extname = 'plmruby';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
