## 用法

来源：

- [README.md](https://github.com/hhkbble/pg_charabia/blob/e035e7d775f918730c9f3ff51cce8d181b28be31/README.md)
- [pg_charabia.control](https://github.com/hhkbble/pg_charabia/blob/e035e7d775f918730c9f3ff51cce8d181b28be31/pg_charabia.control)
- [Cargo.toml](https://github.com/hhkbble/pg_charabia/blob/e035e7d775f918730c9f3ff51cce8d181b28be31/Cargo.toml)
- [reference.md](https://github.com/hhkbble/pg_charabia/blob/e035e7d775f918730c9f3ff51cce8d181b28be31/docs/reference.md)
- [operations.md](https://github.com/hhkbble/pg_charabia/blob/e035e7d775f918730c9f3ff51cce8d181b28be31/docs/operations.md)
- [pg_charabia--0.4.1--0.4.2.sql](https://github.com/hhkbble/pg_charabia/blob/e035e7d775f918730c9f3ff51cce8d181b28be31/sql/pg_charabia--0.4.1--0.4.2.sql)

`pg_charabia` 0.4.2 为 PostgreSQL 18 原生全文检索提供多语言分析，要求数据库采用 UTF8 编码。资源随扩展内嵌，原生全文检索不依赖外部服务或可选的 BM25 扩展。

### 核心用法

数据库超级用户将不可迁移的扩展安装到固定模式中：

```sql
CREATE EXTENSION pg_charabia;
SELECT charabia.analysis_info();
SELECT * FROM charabia.analyze('PostgreSQL multilingual search');
CREATE TABLE documents (id bigint PRIMARY KEY, body text NOT NULL);
INSERT INTO documents VALUES (1, 'PostgreSQL multilingual search');
CREATE INDEX documents_gin ON documents
USING gin (to_tsvector('public.charabia', body));
SELECT id FROM documents
WHERE to_tsvector('public.charabia', body)
   @@ plainto_tsquery('public.charabia', 'PostgreSQL');
```

文档、查询、索引和高亮应使用同一配置。`ts_headline` 使用原文生成高亮；`charabia.query_any` 提供以 OR 连接词项的普通查询分析。

### 配置与权限

`public.charabia` 和 `public.charabia_linguistic` 提供普通语言分析。需要识别 URL、主机名、邮箱、路径及字母数字混合词元时，应显式选择 `public.charabia_structured`。扩展不会修改 `default_text_search_config`。

固定的 `charabia` 模式不得允许不可信角色创建对象。安装过程保留该模式中的 SQL 对象及三个公共配置。自定义配置和词典应放在应用拥有的模式中。复制配置会共享词典，随后修改词典将影响所有引用它的配置。

### 维护与限制

原生全文检索无需预加载。可选的 BM25 功能要求另外安装兼容的 `pg_textsearch` 及其依赖，并预加载后重启服务器。正式发布验证面向 Linux AMD64；macOS ARM64 运行器用于本地开发，测试覆盖范围较小。

从 0.4.1 升级时，需先停止 PostgreSQL 再替换库文件，重启后在各数据库运行 `ALTER EXTENSION pg_charabia UPDATE TO '0.4.2'`。该升级保留 SQL 对象及分析行为，本身不要求重建索引。更早版本没有直接升级路径。

分析规则或词典变化可能要求从原文重新计算存储的向量并重建索引。恢复时应保留准确版本的库、资源及外部词典文件。Rust 内存分配不受 `work_mem` 限制，分配失败可能中止后端并触发恢复。选择索引策略前应核对上游文档中的语言分析限制。
