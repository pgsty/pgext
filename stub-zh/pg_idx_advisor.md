## 用法

来源：

- [Official README and deprecation notice](https://github.com/cohenjo/pg_idx_advisor/blob/5c7a51797b65b0fd8b8d4b26e2fbc28e6a8ca780/README.md)
- [Extension SQL](https://github.com/cohenjo/pg_idx_advisor/blob/5c7a51797b65b0fd8b8d4b26e2fbc28e6a8ca780/sql/pg_idx_advisor.sql)
- [Planner-hook implementation](https://github.com/cohenjo/pg_idx_advisor/blob/5c7a51797b65b0fd8b8d4b26e2fbc28e6a8ca780/src/idx_adviser.c)

pg_idx_advisor 是一个已弃用的规划器钩子扩展，会把普通 EXPLAIN 计划与加入假想候选索引后的计划进行比较，并给出 CREATE INDEX 建议。上游说明项目不再更新，并引导用户使用 HypoPG。它只适合复现历史行为。

### 核心流程

先创建目录表，再在会话中加载库，然后运行 EXPLAIN：

```sql
CREATE EXTENSION pg_idx_advisor;
LOAD 'pg_idx_advisor';

EXPLAIN
SELECT * FROM orders WHERE customer_id = 42;

SELECT recommendation, benefit, index_size
FROM index_advisory
WHERE backend_pid = pg_backend_pid()
ORDER BY benefit DESC;
```

扩展会打印原始计划和假想计划；除非启用只读模式，还会把建议存入 `index_advisory`。

### 重要对象

- `index_advisory` 存储关系 OID、索引属性、估算收益与大小、表达式或谓词、源查询及建议文本。
- `index_adviser.cols` 列出部分索引考虑的列。
- `index_adviser.schema` 选择 advisory 表所在模式。
- `index_adviser.read_only` 禁止写表，只打印建议。
- `index_adviser.text_pattern_ops` 允许文本模式操作符类。
- `index_adviser.composit_max_cols` 限制复合候选项的列数。

### 运维说明

加载该库会替换当前后端的规划器和 EXPLAIN 钩子，并使用旧版 PostgreSQL 虚拟索引内部接口。建议只是成本模型假设，并不能证明索引在实际数据、写入、缓存状态或并发下有效。执行任何输出 SQL 前，都要审查锁、构建时间、存储、写放大、重复索引和部分谓词。受支持 PostgreSQL 版本应优先采用基于 HypoPG 且仍受维护的工作流。
