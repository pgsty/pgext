## 用法

来源：

- [README](https://github.com/planetscale/lead/blob/bd95c7e51b6afce81396790852ee2f2c169570ad/README.md)
- [Control file / 控制文件](https://github.com/planetscale/lead/blob/bd95c7e51b6afce81396790852ee2f2c169570ad/postgres/tin.control)
- [Cargo.toml](https://github.com/planetscale/lead/blob/bd95c7e51b6afce81396790852ee2f2c169570ad/Cargo.toml)
- [postgres/src/lib.rs](https://github.com/planetscale/lead/blob/bd95c7e51b6afce81396790852ee2f2c169570ad/postgres/src/lib.rs)

`tin` 是 Lead 的扩展名，提供兼容 TIN 的文本搜索，适合开发、CI 和预发布测试。1.0.3 版有意采用简单参考实现，牺牲查询性能，不适合生产搜索。

### 核心工作流

```sql
CREATE EXTENSION tin;
CREATE TABLE documents (id integer, body text);
INSERT INTO documents VALUES (1, 'craft beer'), (2, 'wine'), (3, 'beer festival');
CREATE INDEX documents_search ON documents USING tin (body);
SELECT id, tin.full_score(ctid) AS score
FROM documents WHERE body ==> 'beer' ORDER BY score DESC;
```

### 对象与行为

`tin` 访问方法与 `==>` 运算符用于计算 TINQL 搜索条件。`tin.score`、`tin.full_score`、`tin.max_score`、`tin.score_inspect` 提供评分，`tin.highlight` 和 `tin.highlight_ansi` 标记匹配文本。评分调用绑定到对应的搜索表达式和索引。

每次索引扫描都会将全部堆页面作为有损候选返回，由 PostgreSQL 重新检查可见行、表达式和部分索引条件。索引不保存搜索数据，评分也可能再次扫描可见堆行。这样保留了查询语义，但没有提供生产级倒排索引。

### 使用要求

审核源码的目标版本是 PostgreSQL 17 和 18。安装需要超级用户，对象使用固定的 `tin` 模式。库按需加载，无需预加载或重启；测试时仍需控制数据规模。
