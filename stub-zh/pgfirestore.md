## 用法

来源：

- [官方 README](https://github.com/xwkuang5/pgfirestore/blob/bc524d1b5200218b973a157361d74fc55031c907/README.md)
- [官方 control 文件](https://github.com/xwkuang5/pgfirestore/blob/bc524d1b5200218b973a157361d74fc55031c907/pgfirestore.control)

`pgfirestore` 是实验性的 PostgreSQL 查询引擎，用于模拟部分 Firestore 层次化文档与值语义。文档以自定义 `fsvalue` 引用和映射的形式保存在单一 `fs_documents` 表中。上游称它是没有性能预期的探索项目，而且项目已废弃。

### 核心流程

构造带类型的引用与映射，插入文档表，再用 Firestore 风格的比较运算符查询集合组：

```sql
CREATE EXTENSION pgfirestore;

INSERT INTO fs_documents (reference, properties)
VALUES (
  fs_reference('/users/1'),
  fs_map_from_entries(ARRAY['name'], ARRAY[fs_string('Ada')])
);

SELECT reference, properties
FROM fs_collection_group('users')
WHERE properties -> 'name' #= fs_string('Ada');
```

`fs_collection(parent, collection_id)` 返回某个引用下的一层集合；`fs_collection_group(collection_id)` 查找数据库根目录下任意位置具有该 ID 的集合。

### 类型、构造函数与运算符

- `fsvalue`：以 CBOR 存储、以 JSON 表示文本，并实现 Firestore 跨类型排序的值类型。
- `fs_null()`、`fs_nan()`、`fs_boolean(bool)`、`fs_number_from_integer(integer)`、`fs_number_from_double(double precision)`。
- `fs_string(text)`、`fs_reference(text)`、`fs_array(fsvalue[])`、`fs_map_from_entries(text[], fsvalue[])`。
- `#<`、`#>`、`#<=`、`#>=`、`#=`、`#!=`：同类型 Firestore 查询比较；普通 PostgreSQL 比较运算符实现跨类型排序。
- `->`：提取嵌套文档属性。

### 注意事项

`fs_documents` 除主键外没有二级索引，因此集合扫描和属性过滤不面向大规模数据。Firestore 安全规则尚未实现；已发布代码虽然展示了 Date 和 GeoPoint 的预期外部表示，但并未实现这两种值。

它不是 Firestore 服务器、同步层或兼容性保证。应逐项验证应用所需的类型与运算符语义，独立实施授权控制，并且只用于一次性实验数据库。
