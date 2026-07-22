## 用法

来源：

- [官方代码仓库](https://github.com/BlueTreble/pg_thud)
- [官方 control 文件](https://github.com/BlueTreble/pg_thud/blob/4d2aba37d20c942f9c05412a6f045f328106cfb3/pg_thud.control)
- [官方 PGXN 发布页](https://pgxn.org/dist/pg_thud/)

`pg_thud` 是 PostgreSQL 测试数据工厂。它为某个表的行类型登记命名 SQL 查询，并在首次使用时物化结果，让测试可以重复请求固定夹具集。项目已经停止维护，文字文档也只有脚手架，因此只应在隔离测试数据库中使用 SQL 所定义的 API。

### 核心流程

安装声明的 `pgtap` 依赖，登记一个或多个命名数据集，再用带类型的 `NULL` 指明目标行类型并取回数据集：

```sql
CREATE EXTENSION pgtap;
CREATE EXTENSION pg_thud;

CREATE TABLE widget (id integer, name text);

SELECT tf.register(
  'widget',
  ARRAY[
    ROW(
      'small',
      'SELECT * FROM (VALUES (1, ''one''), (2, ''two'')) AS v(id, name)'
    )::tf.test_set
  ]
);

SELECT * FROM tf.get(NULL::widget, 'small');
```

首次调用 `tf.get` 时，会在扩展的 `_test_data` 模式中创建缓存数据；后续调用会复用它。

### API 与对象

- `tf.test_set`：包含 `set_name` 与 `insert_sql` 的复合值。
- `tf.register(table_name text, test_sets tf.test_set[])`：为一张表登记数据工厂。
- `tf.get(anyelement, set_name text) RETURNS SETOF anyelement`：以调用者指定的复合类型返回命名夹具集。
- `_tf.test_factory`：内部注册表；`_test_data` 保存物化夹具表。

### 注意事项

两个公开函数都是 `SECURITY DEFINER`，登记过程会保存 SQL 文本，随后执行它来构建夹具数据。应限制 `EXECUTE` 权限，绝不能接收不可信的工厂 SQL。当表定义或工厂查询改变时，缓存数据可能过期；公开 API 没有记录刷新命令，因此定义改变后应重建一次性测试状态。

`pg_thud` 负责生成测试数据，本身不执行结果断言。它的依赖和源代码面向旧版 PostgreSQL，且上游项目已经废弃。必须针对确切服务器版本验证安装，并避免用于生产数据库。
