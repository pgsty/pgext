## 用法

来源：

- [官方扩展控制文件](https://github.com/ttfkam/pg_content_utils/blob/71efbcd005630f6d2c96e221c8a6f9aa859b6676/content_utils.control)
- [官方上游文档](https://github.com/ttfkam/pg_content_utils/blob/71efbcd005630f6d2c96e221c8a6f9aa859b6676/README.md)

`content_utils` — 用于新闻内容管理与全文搜索的 PostgreSQL 工具函数。

已复核目录快照记录的版本为 `1.0.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`file_fdw`。

```sql
CREATE EXTENSION "content_utils";
SELECT extversion
FROM pg_extension
WHERE extname = 'content_utils';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
