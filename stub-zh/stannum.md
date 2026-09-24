## 用法

来源：

- [README](https://github.com/TeamSpringbird/stannum/blob/e163585cb9d6b6f78b8de067a9c4aa33ea238063/README.md)
- [Control file / 控制文件](https://github.com/TeamSpringbird/stannum/blob/e163585cb9d6b6f78b8de067a9c4aa33ea238063/postgres/stannum.control)
- [Cargo.toml](https://github.com/TeamSpringbird/stannum/blob/e163585cb9d6b6f78b8de067a9c4aa33ea238063/Cargo.toml)
- [postgres/sql/stannum--0.1.0.sql](https://github.com/TeamSpringbird/stannum/blob/e163585cb9d6b6f78b8de067a9c4aa33ea238063/postgres/sql/stannum--0.1.0.sql)

`stannum` 是基于 Lead 开发的实验性文本搜索扩展，提供持久索引、布尔和短语查询以及 BM25 评分。审核的 0.1.0 源码使用独立的扩展名称和存储格式。

### 核心工作流

```sql
CREATE EXTENSION stannum;
CREATE TABLE documents (id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY, body text);
INSERT INTO documents (body) VALUES
  ('PostgreSQL supports full text search'), ('Search with exact phrase matching');
CREATE INDEX documents_search ON documents USING stannum (body);
ANALYZE documents;
SELECT id, stannum.full_score(ctid) AS score
FROM documents WHERE body ==> 'search' ORDER BY score DESC LIMIT 10;
SELECT * FROM stannum.segment_info('documents_search');
SELECT * FROM stannum.verify_index('documents_search', heap_check => true);
```

### 运行与兼容性

`==>` 运算符接受 TINQL。`stannum.highlight` 标记匹配文本；`stannum.segment_info` 检查索引段，`stannum.verify_index` 验证一致性，并可同时检查堆数据。索引随 PostgreSQL 写入和清理操作维护。

构建目标为 PostgreSQL 17 和 18，目前大部分生命周期验证集中在 18。超级用户将扩展安装到固定的 `stannum` 模式。主库使用时按需加载；如需在备库读取索引，须在主备两侧的 `shared_preload_libraries` 中加入该库并重启服务器。

### 迁移边界

不要在同一数据库中同时安装 `tin` 和 `stannum`，两者的全局 `==>` 运算符冲突。上游不支持从旧扩展名称或旧存储布局原地迁移。应按实验性软件使用，并针对目标负载验证恢复和维护行为。
