

## 用法

> [debversion: PostgreSQL 的 Debian 版本号类型](https://github.com/ATIX-AG/postgresql-debversion-evr)

`debversion` 扩展提供了 Debian 软件包版本比较功能，实现了与 `dpkg` 相同的排序逻辑。

```sql
CREATE EXTENSION debversion;
```

### 数据类型

`debversion` 类型存储 Debian 软件包版本字符串，并按照 Debian 版本规范（epoch:upstream-revision 格式）进行比较。

```sql
CREATE TABLE packages (
    name    text,
    version debversion
);

INSERT INTO packages VALUES ('foo', '1.0-1'), ('foo', '2.0-1'), ('foo', '1.0-2');

SELECT * FROM packages ORDER BY version;
```

### 版本比较

```sql
SELECT '1.60-26+b1'::debversion < '1.60+git20161116.90da8a0-1'::debversion;
```

### 运算符

支持标准比较运算符：`=`、`<>`、`<`、`>`、`<=`、`>=`。

比较算法与 `dpkg --compare-versions` 一致，确保结果与标准 Debian 包管理工具完全相同。
