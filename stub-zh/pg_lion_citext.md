## 用法

来源：

- [上游来源 1](https://github.com/cgkanchi/pg_lion/blob/e8c3fea6d9c9fa1da8d501966f61703ba53d17fa/README.md)
- [上游来源 2](https://github.com/cgkanchi/pg_lion/blob/e8c3fea6d9c9fa1da8d501966f61703ba53d17fa/pg_lion_citext.control)

`pg_lion_citext` 0.3.0 为 lion 索引中的 citext 值提供大小写不敏感的相等比较支持。

### 核心流程

先安装 `citext` 与 `pg_lion`。此伴随扩展注册运算符类，应用仍对列使用普通相等条件。

```sql
CREATE EXTENSION citext;
CREATE EXTENSION pg_lion;
CREATE EXTENSION pg_lion_citext;
CREATE TABLE lion_contacts (email citext);
INSERT INTO lion_contacts VALUES ('Ada@example.com');
CREATE INDEX ON lion_contacts USING lion (email);
SELECT * FROM lion_contacts WHERE email = 'ada@example.com';
```

### 运行边界

需要 PostgreSQL 16 及以上版本，并由超级用户安装。应使用匹配的 pg_lion 分发版本。大小写折叠遵循 citext 和排序规则语义，并不等同于任意 Unicode 等价关系。父扩展索引的可选预加载、自定义 WAL 约束及 VACUUM 行为同样适用。
