## 用法

来源：

- [上游 README](https://github.com/neurondb/pg_pageindex/blob/893c95a4846787e65b2276f3016a1852b42e49ff/README.md)
- [1.0 版 SQL API](https://github.com/neurondb/pg_pageindex/blob/893c95a4846787e65b2276f3016a1852b42e49ff/sql/pg_pageindex--1.0.sql)
- [API 参考](https://github.com/neurondb/pg_pageindex/blob/893c95a4846787e65b2276f3016a1852b42e49ff/docs/api.md)
- [上游安全指南](https://github.com/neurondb/pg_pageindex/blob/893c95a4846787e65b2276f3016a1852b42e49ff/docs/security.md)

pg_pageindex 1.0 在 PostgreSQL 后端进程内运行 PageIndex 文档树流水线。它读取服务器端 PDF 或 Markdown 文件、返回 JSON 文档，并可调用语言模型生成摘要和执行树搜索。

### 首先加固安装

安装 SQL 默认向 public 授予函数执行权限。应在同一事务中撤销，再仅向专用可信角色授予所需函数：

```sql
BEGIN;
CREATE EXTENSION pg_pageindex;
REVOKE EXECUTE ON ALL FUNCTIONS IN SCHEMA pageindex FROM PUBLIC;
COMMIT;
SELECT pageindex.version();
```

完成显式授权后，可信角色可以从批准的服务器路径构建：

```sql
SELECT pageindex.build_markdown(
  '/srv/pageindex/manual.md',
  '{"add_node_text":false}'::jsonb
);
```

### 注意事项

- 路径在数据库服务器上解析，并由 PostgreSQL 操作系统账户读取。绝不能接受客户端控制的路径；上游将其信任边界等同于危险的服务器端程序执行。
- PDF 解析以及 Go、MuPDF 与 PageIndex 桥接代码都在调用后端内运行。畸形或超大文档可能消耗内存与 CPU、引发错误或使后端崩溃。
- 摘要、描述与搜索操作可能向配置的 OpenAI 兼容或 Anthropic 端点发起阻塞式外连请求。文档文本、结构、提示词与查询可能离开数据库主机。
- API 密钥应置于 SQL 与数据库表之外，并限制出站网络、设置语句与网络超时、审查供应商保留策略。
- 扩展返回 JSON，而不管理已索引数据行。它不会创建 PostgreSQL 索引访问方法，也不会加速任意表查询。
