## 用法

来源：

- [官方扩展控制文件](https://github.com/frost242/monitoring_role/blob/f76f0629de7a8af4e28b8ca92336a5b050ef4e39/monitoring_role.control)
- [官方上游文档](https://github.com/frost242/monitoring_role/blob/f76f0629de7a8af4e28b8ca92336a5b050ef4e39/README.md)

`monitoring_role` — 通过安全定义者函数向非特权监控角色授予受控的统计与文件检查能力。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "monitoring_role";
SELECT extversion
FROM pg_extension
WHERE extname = 'monitoring_role';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
