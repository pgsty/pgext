## 用法

来源：

- [官方扩展控制文件](https://github.com/Sednai/xl_global_views/blob/804ab3fbd91ca84441af86d520717ce22583cd98/xl_global_views.control)
- [官方上游文档](https://github.com/Sednai/xl_global_views/blob/804ab3fbd91ca84441af86d520717ce22583cd98/README.md)

`xl_global_views` — xl_global_views：Postgres-XL 全局视图扩展，聚合各协调器和数据节点的系统视图。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "xl_global_views";
SELECT extversion
FROM pg_extension
WHERE extname = 'xl_global_views';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
