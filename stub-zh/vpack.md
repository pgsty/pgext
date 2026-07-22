## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/siilike/postgresql-vpack/blob/88148fa952875c6fc3ef49c3add0eec4a48418e0/README.md)
- [0.1 版本 SQL API](https://github.com/siilike/postgresql-vpack/blob/88148fa952875c6fc3ef49c3add0eec4a48418e0/vpack--0.1.sql)
- [扩展控制文件](https://github.com/siilike/postgresql-vpack/blob/88148fa952875c6fc3ef49c3add0eec4a48418e0/vpack.control)

`vpack` 为 PostgreSQL 添加 ArangoDB VelocyPack 文档存储。`vpack` 类型大体类似 `jsonb`，但支持额外的二进制/逻辑类型；`vpackpath` 提供类似 JSONPath 的查询语言。扩展还提供 JSON/JSONB/BSON 转换、文档操作符、B-tree/hash/GIN 操作符类、构造器及路径函数。

```sql
CREATE EXTENSION vpack;
CREATE TABLE events (id bigint PRIMARY KEY, body vpack);
INSERT INTO events VALUES (1, '{"kind":"login","ok":true}'::vpack);
SELECT body ->> 'kind' FROM events WHERE body @> '{"ok":true}'::vpack;
CREATE INDEX ON events USING gin (body);
```

README 说明部分功能和测试尚未完成，`vpackpath` 尚未完整支持标签，类型到标签的映射没有完整规范，还有若干 `jsonb` 风格函数缺失。文档也把部分基准性能优势归因于省略输入校验。必须校验每个不可信载荷，且不要假定它具有与 JSON 相同的规范化或比较语义。

VelocyPack、BSON、JSON 和 PostgreSQL 类型之间不存在无损一一映射。应测试数值范围、重复键、二进制值、标签、路径错误、排序、索引复查及往返转换。应固定上游 VelocyPack 和 Mongo C 驱动 ABI，并在持久使用前验证二进制转储恢复及逻辑复制。
