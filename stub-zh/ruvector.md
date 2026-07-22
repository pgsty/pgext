## 用法

来源：

- [官方扩展 README](https://github.com/ruvnet/RuVector/blob/f3de1724fa5d8ff871ac24a528575f115b2a9df7/crates/ruvector-postgres/README.md)
- [官方扩展控制文件](https://github.com/ruvnet/RuVector/blob/f3de1724fa5d8ff871ac24a528575f115b2a9df7/crates/ruvector-postgres/ruvector.control)
- [官方扩展构建清单](https://github.com/ruvnet/RuVector/blob/f3de1724fa5d8ff871ac24a528575f115b2a9df7/crates/ruvector-postgres/Cargo.toml)

`ruvector` 为 PostgreSQL 增加定长向量类型、距离运算符与 RuVector HNSW 访问方法。它适合近邻检索，但应先确认实际安装的构建确实包含所需函数：本次核验的扩展控制版本是 `0.3.0`，而当前 Rust crate 清单标为 `2.0.1`，且许多高级模块受编译期特性开关控制。

### 核心流程

创建扩展，以明确维度存储向量，并按选定的距离运算符排序：

```sql
CREATE EXTENSION ruvector;

CREATE TABLE documents (
    id bigserial PRIMARY KEY,
    content text NOT NULL,
    embedding ruvector(3) NOT NULL
);

INSERT INTO documents (content, embedding) VALUES
    ('alpha', '[0.1,0.2,0.3]'),
    ('beta',  '[0.8,0.1,0.2]');

CREATE INDEX documents_embedding_idx
    ON documents USING ruhnsw (embedding ruvector_l2_ops);

SELECT id, content,
       embedding <-> '[0.2,0.2,0.3]'::ruvector AS distance
FROM documents
ORDER BY embedding <-> '[0.2,0.2,0.3]'::ruvector
LIMIT 5;
```

查询必须包含用于行排序的同一个 `ORDER BY` 距离表达式。请用有代表性的数据验证召回率、延迟与索引行为，不要把近似近邻结果当作精确结果。

### 重要对象

- `ruvector(n)` 存储维度由类型修饰符固定的向量。
- `<->` 计算用于排序的 L2 距离；上游扩展还为 L2、余弦、内积与曼哈顿距离定义了函数形式。
- `ruhnsw` 是 HNSW 索引访问方法，`ruvector_l2_ops` 选择其 L2 运算符类。
- 高级图、注意力、求解器、本地嵌入及数学函数取决于构建特性。在依赖新版上游文档展示的函数前，应检查已安装构建的 `pg_proc`。

### 运维注意事项

本次核验的上游 README 声明支持 PostgreSQL 14 至 17；请在实际服务器大版本上测试软件包。控制文件将扩展标为可信且不可迁移，但对象权限仍需按常规审查。

HNSW 索引用构建时间、内存、磁盘空间与精确性换取检索速度。大版本升级或批量数据变更后应重建并复测召回率，备份恢复演练也要涵盖表和索引。若构建包含本地嵌入或 AI 函数，模型加载与推理会在数据库后端内运行，因此在用于生产流量前须测量后端内存、CPU、语句超时与并发影响。
