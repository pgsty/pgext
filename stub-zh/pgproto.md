
## 用法

> 语法：
>
> ```sql
> CREATE EXTENSION pgproto;
> INSERT INTO pb_schemas (name, data) VALUES ('MySchema', '\x...');
> CREATE TABLE items (id serial PRIMARY KEY, data protobuf);
> SELECT data #> '{Outer, inner, id}'::text[] FROM items;
> ```
>
> 来源：[README](https://github.com/Apaezmx/pgproto)

`pgproto` 为 PostgreSQL 增加了原生 Protocol Buffers 支持。它提供 `protobuf` 类型、运行时 schema 注册、嵌套字段提取、更新辅助函数，以及面向 protobuf 载荷的索引支持。

## 配置

启用扩展：

```sql
CREATE EXTENSION pgproto;
```

通过加载 `FileDescriptorSet` blob 注册 protobuf schema：

```sql
INSERT INTO pb_schemas (name, data) VALUES ('MySchema', '\x...');
```

使用自定义 `protobuf` 类型创建表：

```sql
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    data protobuf
);
```

## 查询

README 强调可以用 PostgreSQL 风格的操作符进行嵌套字段提取：

```sql
SELECT data #> '{Outer, inner, id}'::text[] FROM items;
SELECT data #> '{Outer, tags, mykey}'::text[] FROM items;
```

它还提到 `->` 和 `#>` 等自定义操作符可用于按 schema 进行导航。

## 修改函数

`pgproto` 提供若干纯函数，会返回一个新的 protobuf 值：

- `pb_set(...)`
- `pb_insert(...)`
- `pb_delete(...)`

由于它们返回的是修改后的值，而不是原地变更，因此通常会在 `UPDATE` 语句中使用：

```sql
UPDATE items SET data = pb_set(data, ARRAY['Outer', 'a'], '42');
UPDATE items SET data = pb_insert(data, ARRAY['Outer', 'scores', '0'], '100');
UPDATE items SET data = pb_delete(data, ARRAY['Outer', 'a']);
```

`||` 操作符用于合并两个同类型的 protobuf 消息。

## 索引

README 记录了对提取字段做 B-tree 表达式索引的方法：

```sql
CREATE INDEX idx_pb_id ON items ((data #> '{Outer, inner, id}'::text[]));
```

项目还宣称支持 GIN 索引以服务查询场景。

## 说明

上游 README 将 `pgproto` 定位为比 JSONB 更节省存储空间的 protobuf 原生载荷方案，并强调 protobuf schema 演进、枚举、`oneof` 以及 map/repeated 字段访问都可支持。
