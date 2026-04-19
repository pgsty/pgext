## 用法

来源: [official repo](https://github.com/nmandery/pg_byteamagic), [official doc](https://raw.githubusercontent.com/nmandery/pg_byteamagic/master/doc/byteamagic.md)

`byteamagic` 会对 `bytea` 值运行 `libmagic`，因此 PostgreSQL 可以识别数据库中 blob 的 MIME 类型或人类可读文件类型。包名是 `pg_byteamagic`，但扩展名是 `byteamagic`。

```sql
CREATE EXTENSION byteamagic;

SELECT byteamagic_mime(data) FROM files;
SELECT byteamagic_text(data) FROM files;
```

### 函数

- `byteamagic_mime(bytea) returns text`: 返回 MIME 类型，对应 `file --mime-type`。
- `byteamagic_text(bytea) returns text`: 返回描述性的文件类型文本，对应 `file`。

### 常见用法

```sql
SELECT
  id,
  byteamagic_mime(blob) AS mime,
  byteamagic_text(blob) AS kind
FROM uploads;
```

### 注意事项

- 它只检查 `bytea` 内容；没有运算符、自定义类型或额外的 SQL 管理对象。
- 构建或安装需要 PostgreSQL 开发头文件，以及系统 `libmagic` 开发包。
- 上游文档非常简略；当前面向用户的行为与长期存在的文档页一致。
