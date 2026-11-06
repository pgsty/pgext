## 用法

来源：

- [README](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/README.md)
- [参考手册](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/docs/REFERENCE.md)
- [扩展控制文件](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion_citext.control)
- [扩展 SQL](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion_citext--0.3.0.sql)

`pg_lion_citext` 0.3.0 为 `lion` 索引注册 `citext_ops` 运算符类，用于大小写不敏感的文本键。它是随 `pg_lion` 分发的 SQL 伴随扩展，没有独立动态库。

### 核心流程

先创建 `citext` 和 `pg_lion`，再启用此伴随扩展。默认运算符类使用 `citext` 的哈希与比较函数，支持相等、不等及范围比较（`=`、`<>`、`<`、`<=`、`>=`、`>`）。

```sql
CREATE EXTENSION citext;
CREATE EXTENSION pg_lion;
CREATE EXTENSION pg_lion_citext;
CREATE TABLE lion_contacts (email citext);
INSERT INTO lion_contacts VALUES ('Ada@example.com'), ('ADA@example.com');
CREATE INDEX lion_email ON lion_contacts USING lion (email);
SELECT count(*) FROM lion_contacts WHERE email = 'ada@example.com';
```

示例计数为两行，因为两种拼写比较相等。应用继续使用普通的 `citext` 条件。

### 运行边界

需要 PostgreSQL 16 及以上版本，并由超级用户安装。大小写折叠与排序遵循 `citext` 语义。应使用同一父扩展版本中的文件。父扩展索引的可见性、`VACUUM` 及可选自定义 WAL 预加载约束仍然适用。此伴随扩展本身不会启用预加载或修改既有索引。
