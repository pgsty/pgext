## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/jsonbd/blob/6a8aefcca9999596d75212a0f68c28590b5c3570/jsonbd.control)
- [官方上游文档](https://github.com/postgrespro/jsonbd/blob/6a8aefcca9999596d75212a0f68c28590b5c3570/README.md)

`jsonbd` — jsonbd：为 PostgreSQL 的 JSONB 类型提供压缩方法。

已复核目录快照记录的版本为 `0.1`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "jsonbd";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonbd';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
