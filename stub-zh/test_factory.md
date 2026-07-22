## 用法

来源：

- [官方 PGXN 文档](https://pgxn.org/dist/test_factory/0.5.0/doc/test_factory.html)
- [官方 README](https://github.com/BlueTreble/test_factory/blob/0eb02e0fd8fe0fbd59dc22d4d3b9f86c27678054/README.md)
- [0.5.0 版扩展 SQL](https://github.com/BlueTreble/test_factory/blob/0eb02e0fd8fe0fbd59dc22d4d3b9f86c27678054/sql/test_factory.sql)
- [扩展控制文件](https://github.com/BlueTreble/test_factory/tree/0eb02e0fd8fe0fbd59dc22d4d3b9f86c27678054)

`test_factory` 0.5.0 为关系测试数据注册具名配方，在首次请求时创建配方行，并持久缓存副本供后续测试使用。它旨在避免重复重建深层测试数据依赖，同时保持引用可读。

### 注册并取得测试集

安装会创建角色和固定模式，因此应以超级用户加载扩展。每个配方的 SQL 必须返回与目标表形状完全一致的行：

```sql
CREATE EXTENSION test_factory;

CREATE TABLE customer (
  customer_id bigint GENERATED ALWAYS AS IDENTITY,
  email text NOT NULL
);

SELECT tf.register(
  table_name => 'customer',
  test_sets => ARRAY[
    ROW(
      'base',
      'INSERT INTO customer(email) VALUES (''base@example.org'') RETURNING *'
    )::tf.test_set
  ]
);

SELECT * FROM tf.get(NULL::customer, 'base');
```

第一次 `tf.get` 会执行注册 SQL 并缓存副本。之后使用相同表类型和集合名调用时，会直接返回副本，不再运行配方。配方可以为其他表调用 `tf.get`，从而表达依赖而无需限制注册顺序。

### 主要对象

- `tf.test_set` 是 `(set_name text, insert_sql text)`。
- `tf.register(table_name text, test_sets tf.test_set[])` 为现有表添加配方，或替换名称相同的配方。
- `tf.get(table_type anyelement, set_name text) RETURNS SETOF anyelement` 通过 `NULL::customer` 这样的类型化 NULL 确定结果行类型。
- 可选扩展 `test_factory_pgtap` 依赖 `pgtap`，并添加 `tf.tap(table_name, set_name)` 作为 `isnt_empty` 便捷断言。

扩展把配方定义存入 `_tf._test_factory`，把缓存行副本存入 `_test_factory_test_data`，并通过 `test_factory__owner` 角色持有。它的配置表会包含在扩展感知的转储中。

### 缓存、权限与清理边界

缓存副本永远不会自动删除。原始应用行的修改不会反映到后续 `tf.get` 调用中；注册新列表也不会删除本次调用省略的旧集合名。应规划显式清理，或在独立测试套件之间重建测试数据库。

配方 SQL 是调用方首次请求集合时执行的动态 SQL。调用方需要相应语句权限，而且配方文本必须是受信任代码，不能来自用户输入。配方必须返回与目标表兼容的行，通常通过 `RETURNING *`；否则模式漂移会导致检索失败，或保留过时的缓存行形状。

此扩展只能用于受控测试数据库或严格隔离的测试数据模式。它持久创建的角色、模式、可执行配方和缓存数据不适合未经审查的生产安装。
