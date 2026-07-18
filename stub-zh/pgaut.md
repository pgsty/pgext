## 用法

来源：

- [已复核 commit 的归档上游 README](https://github.com/redraiment/pgaut/blob/7db20e740ec8167c781b4456d77a2090d7ef2782/README.md)
- [1.0.0 版本 SQL 实现](https://github.com/redraiment/pgaut/blob/7db20e740ec8167c781b4456d77a2090d7ef2782/pgaut--1.0.0.sql)
- [扩展 control 文件](https://github.com/redraiment/pgaut/blob/7db20e740ec8167c781b4456d77a2090d7ef2782/pgaut.control)

`pgaut` 通过 `auto_update_timestamp` domain 与 PostgreSQL event trigger 提供 MySQL 风格的自动更新时间戳。创建表时若有该 domain 类型的列，扩展会添加行触发器，在其他列变化时写入 `clock_timestamp()`。

```sql
CREATE EXTENSION pgaut;

CREATE TABLE app_user (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  updated_at auto_update_timestamp DEFAULT current_timestamp
);

INSERT INTO app_user (name) VALUES ('alice');
UPDATE app_user SET name = 'Alice' WHERE name = 'alice';
```

### 注意事项

安装会为 `CREATE TABLE`、`ALTER TABLE` 与 `DROP TABLE` 创建数据库级 event trigger，并动态创建每张表的 PL/pgSQL 触发器函数。这需要较高权限，在共享数据库中使用前必须审查。仓库已归档，最后一次源码变更在 2018 年；上游关于支持 PostgreSQL 9.1 及更高版本的声明没有针对当前大版本进行验证。
