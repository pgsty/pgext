## 用法

来源：

- [官方扩展控制文件](https://github.com/ttfkam/pg_newsfeeds/blob/02cb001324498999d7d055453b62dac93a4c4e9e/newsfeeds.control)
- [官方上游文档](https://github.com/ttfkam/pg_newsfeeds/blob/02cb001324498999d7d055453b62dac93a4c4e9e/README.md)

`newsfeeds` — newsfeeds：用于采集新闻的 PostgreSQL 扩展。

已复核目录快照记录的版本为 `1.0.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "newsfeeds";
SELECT extversion
FROM pg_extension
WHERE extname = 'newsfeeds';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
