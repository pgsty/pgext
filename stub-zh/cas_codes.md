## 用法

来源：

- [官方控制文件](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/cas_codes/cas_codes.control)
- [官方 CAS Codes README](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/cas_codes/README.md)
- [生成的 PostgreSQL 17 扩展 SQL](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/cas_codes/extension/usr/share/postgresql/17/extension/cas_codes--0.1.0.sql)

`cas_codes` 添加 Chemical Abstracts Service Registry Number 使用的 `CAS` 类型。输入过程会校验格式和校验位，因此可在数据库边界拒绝格式错误的化学物质标识符。

### 核心流程

```sql
CREATE EXTENSION cas_codes;

CREATE TABLE reagents (
    name text NOT NULL,
    cas_number CAS NOT NULL UNIQUE
);

INSERT INTO reagents VALUES ('water', '7732-18-5');
SELECT * FROM reagents WHERE cas_number = '7732-18-5'::CAS;
```

格式或校验和无效的输入会被类型输入函数拒绝。输出值采用常规的连字符格式。

### SQL 对象

- `CAS` 在当前 PGRX 构建中是变长自定义类型；README 指出所用 PGRX 版本当时尚不支持定长自定义类型。
- 提供比较操作符 `=`、`<>`、`<`、`<=`、`>` 和 `>=`。
- `CAS_btree_ops` 支持 B-tree 排序与唯一性，`CAS_hash_ops` 支持哈希。

### 运维说明

0.1.0 版不可重定位。控制文件设置 `superuser = true` 和 `trusted = false`，因此必须由超级用户创建；它未声明前置扩展或预加载要求。生成 SQL 与 PostgreSQL/PGRX 构建相关，恢复使用 `CAS` 的模式前，应成套安装匹配的控制文件、SQL 和共享库。
