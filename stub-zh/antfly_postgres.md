## 用法

来源：

- [上游来源 1](https://github.com/antflydb/antfly/blob/616117cc78ba29dcd6c0212abd78ac076201ee3f/rs/crates/postgres/README.md)
- [上游来源 2](https://github.com/antflydb/antfly/blob/616117cc78ba29dcd6c0212abd78ac076201ee3f/rs/crates/postgres/antfly_postgres.control)

`antfly_postgres` 0.1.0 将 PostgreSQL 搜索索引及 SQL 函数连接到外部 Antfly 服务。

### 核心流程

创建索引时会发送已有行，并同步后续插入。`antfly_search` 返回标识符、得分与 JSONB 数据；`antfly_sync_trigger` 提供独立的行变更同步方式。查询构造函数 `antfly_postgres.search`、`antfly_postgres.semantic` 和 `antfly_postgres.hybrid` 生成结构化搜索请求。

```sql
CREATE EXTENSION antfly_postgres;
SELECT antfly_status('http://localhost:8080/db/v1/');
CREATE TABLE antfly_docs (content text);
CREATE INDEX ON antfly_docs USING antfly (content)
  WITH (url = 'http://localhost:8080/db/v1/', collection = 'docs');
SELECT * FROM antfly_docs WHERE content @@@ 'database search';
```

### 运行边界

需要 PostgreSQL 13–18、超级用户安装权限及可访问的 Antfly API。配置 URL 时应包含 /db/v1/ 前缀，默认的裸地址无法访问该 API。搜索会将数据发送到 PostgreSQL 之外，因此需明确配置网络访问及服务凭据。不能假定 PostgreSQL 回滚会同时撤销远端写入。本扩展是 pgaf 迁入单体仓库并更名后的实现。
