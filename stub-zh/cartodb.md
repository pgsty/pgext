## 用法

来源：

- [官方扩展控制文件](https://github.com/CartoDB/cartodb-postgresql/blob/e0a0597061bef6cdd7d7239ffa07ab44491817f1/cartodb.control.in)

`cartodb` — cartodb：已归档的 PostgreSQL 扩展，用于将数据库转换为 CartoDB 用户数据库。

已复核目录快照记录的版本为 `0.37.1`、类型为 `puresql`、实现语言为 `Python`。
应先安装并验证声明的扩展依赖：`plpython3u`, `postgis`。
整理后的兼容版本集合为 `11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "cartodb";
SELECT extversion
FROM pg_extension
WHERE extname = 'cartodb';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
