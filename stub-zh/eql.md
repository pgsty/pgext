## 用法

来源：

- [官方上游文档](https://github.com/cipherstash/encrypt-query-language/blob/4119c4246427c54ffb29da0f70a6935a1c7b5215/README.md)
- [官方项目或服务商页面](https://cipherstash.com/docs)
- [官方主文档](https://database.dev/cipherstash/eql)

`eql` — Encrypt Query Language 为 PostgreSQL 提供可搜索加密数据所需的类型、函数、操作符和索引支持。

已复核目录快照记录的版本为 `2.2.1`、类型为 `puresql`、实现语言为 `SQL`。
整理后的兼容版本集合为 `14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "eql";
SELECT extversion
FROM pg_extension
WHERE extname = 'eql';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
