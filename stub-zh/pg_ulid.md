## 用法

来源：

- [官方 README](https://github.com/edoceo/pg-ulid/blob/eda5ace98b0f1668297e95476cde5136ff06c97b/README.md)
- [0.0.4 版 SQL API](https://github.com/edoceo/pg-ulid/blob/eda5ace98b0f1668297e95476cde5136ff06c97b/pg_ulid--0.0.4.sql)
- [生成器实现](https://github.com/edoceo/pg-ulid/blob/eda5ace98b0f1668297e95476cde5136ff06c97b/pg_ulid.c)

`pg_ulid` 0.0.4 提供一个生成 26 字符 ULID 字符串的 C 函数。尽管 README 有更广泛的声明与示例，实际可安装 SQL 并没有定义 PostgreSQL `ulid` 类型、比较操作符、验证辅助函数、格式化辅助函数或接受时间戳的重载。

### 实际 SQL API

```sql
CREATE EXTENSION pg_ulid;

CREATE TABLE events (
    id text PRIMARY KEY DEFAULT ulid_create(),
    payload jsonb NOT NULL
);

INSERT INTO events (payload) VALUES ('{"kind":"created"}');
SELECT id, payload FROM events ORDER BY id;
```

唯一可调用对象是 `ulid_create() RETURNS text`，被声明为 volatile 与 strict。README 展示的 `ulid_create(now())` 调用并未由 `0.0.4` 版定义，值使用普通文本比较与索引语义，而不是自定义二进制类型。

### 标识符边界

实现会在每次调用时初始化内嵌的 `ulid-c` 生成器并返回其文本输出。它没有提供时间戳输入、熵策略选择、规范编码验证、内嵌时间解码或同一毫秒内单调生成的接口。如果需要这些属性，应在应用中验证格式与唯一性，也不要假定它具有集群级单调顺序。

只有所有生成方都遵循相同规范格式时，ULID 才会暴露近似创建时间与词法顺序。该项目没有发行制品或升级链，只停留在一个已废弃的小型原型。持久标识符应优先采用对类型与随机性保证有明确说明且持续维护的 ULID 扩展，或在受检查的契约下存储应用生成值。
