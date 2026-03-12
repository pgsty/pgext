


## 用法

> [rum: RUM 索引访问方法](https://github.com/postgrespro/rum)

RUM 是一种索引访问方法，通过在 posting tree 中存储附加信息来扩展 GIN。这使得可以直接访问位置数据，避免在排序、短语搜索和时间戳排序时进行额外的堆扫描。

### 创建索引

```sql
CREATE INDEX idx ON table_name USING rum (column operator_class);
```

带附加运算符（例如，在全文搜索的同时按时间戳排序）：

```sql
CREATE INDEX tsts_idx ON tsts USING rum (t rum_tsvector_addon_ops, d)
    WITH (attach = 'd', to = 't');
```

### 运算符类

| 运算符类 | 描述 |
|---------|------|
| `rum_tsvector_ops` | 存储 tsvector 词素及位置。支持 `<=>` 排序和前缀搜索。 |
| `rum_tsvector_hash_ops` | 存储哈希化的 tsvector 词素及位置。支持 `<=>` 排序，不支持前缀搜索。 |
| `rum_tsvector_addon_ops` | 将 tsvector 与附加字段（时间戳、整数等）组合，用于过滤和排序。 |
| `rum_tsvector_hash_addon_ops` | 支持附加字段的哈希变体，不支持前缀搜索。 |
| `rum_tsquery_ops` | 存储 tsquery 分支，用于快速查询匹配已索引的文档。 |
| `rum_anyarray_ops` | 索引数组类型。支持 `&&`、`@>`、`<@`、`=`、`%` 和 `<=>` 排序。 |
| `rum_anyarray_addon_ops` | 将数组元素与附加字段组合。 |
| `rum_TYPE_ops` | 通用运算符类，适用于 int2、int4、int8、float4、float8、money、oid、time、timetz、date、interval、macaddr、inet、cidr、text、varchar、char、bytea、bit、varbit、numeric、timestamp、timestamptz。 |

### 排序运算符

| 运算符 | 描述 |
|--------|------|
| `<=>` | 距离运算符，适用于 tsvector、timestamp、数值类型、数组 |
| `<=\|` | 左侧距离，适用于 timestamp、int、float、money、oid |
| `\|=>` | 右侧距离，适用于 timestamp、int、float、money、oid |

### 示例

带排序的全文搜索：

```sql
SELECT t, a <=> to_tsquery('english', 'beautiful | place') AS rank
FROM test_rum
WHERE a @@ to_tsquery('english', 'beautiful | place')
ORDER BY a <=> to_tsquery('english', 'beautiful | place');
```

按时间戳排序的全文搜索：

```sql
SELECT id, d, d <=> '2016-05-16 14:21:25' FROM tsts
WHERE t @@ 'wr&qh'
ORDER BY d <=> '2016-05-16 14:21:25'
LIMIT 5;
```

带距离排序的数组匹配：

```sql
SELECT * FROM test_array
WHERE i && '{1}'
ORDER BY i <=> '{1}' ASC;
```
