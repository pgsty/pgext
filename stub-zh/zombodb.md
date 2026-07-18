## 用法

来源：

- [版本 3000.2.8 的上游 README](https://github.com/zombodb/zombodb/blob/11d5cdfbd2b3fd238acab0ce785cf374e8f6a909/README.md)
- [扩展 control 文件](https://github.com/zombodb/zombodb/blob/11d5cdfbd2b3fd238acab0ce785cf374e8f6a909/zombodb.control)
- [运维设计说明](https://github.com/zombodb/zombodb/blob/11d5cdfbd2b3fd238acab0ce785cf374e8f6a909/THINGS-TO-KNOW.md)
- [已归档上游仓库](https://github.com/zombodb/zombodb)

`zombodb` 实现由 Elasticsearch 支撑的 PostgreSQL 索引访问方法。它同步管理远端索引，提供 `==>` 搜索运算符，同时保持 PostgreSQL 事务可见性：

```sql
CREATE EXTENSION zombodb;

CREATE TABLE products (
  id bigserial PRIMARY KEY,
  name text NOT NULL,
  description zdb.fulltext
);

CREATE INDEX products_zdb_idx
  ON products
  USING zombodb ((products.*))
  WITH (url = 'http://elasticsearch.example.net:9200/');

SELECT *
FROM products
WHERE products ==> 'name:keyboard';
```

### 归档状态与运维风险

该仓库已归档。最终 control 版本 `3000.2.8` 记录支持 PostgreSQL 12 至 15 以及 Elasticsearch 7.10 或更高版本。每张表只能有一个 ZomboDB 索引，不支持部分索引或 `CREATE INDEX CONCURRENTLY`。写入会同步往返 Elasticsearch，因此网络或 Elasticsearch 故障会中止 PostgreSQL 事务。绝不能从外部修改 ZomboDB 管理的索引，并应考虑远端副本会存储完整行和分析后的词项。应将其视为遗留基础设施，并在升级任一依赖前规划迁移。
