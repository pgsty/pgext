## 用法

来源：

- [官方扩展控制文件](https://github.com/fusiongyro/damm/blob/0829d6bc4e51f724a9186a55e26208bcf2152836/damm.control)
- [官方上游文档](https://github.com/fusiongyro/damm/blob/0829d6bc4e51f724a9186a55e26208bcf2152836/README.md)

`damm` — damm：提供 Damm 校验位算法、damm_code 域和相关 SQL 辅助函数。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "damm";
SELECT extversion
FROM pg_extension
WHERE extname = 'damm';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
