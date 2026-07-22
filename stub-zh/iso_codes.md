## 用法

来源：

- [官方控制文件](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/iso_codes/iso_codes.control)
- [官方 ISO Codes README](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/iso_codes/README.md)
- [生成的 PostgreSQL 17 扩展 SQL](https://github.com/earth-metabolome-initiative/emi-monorepo/blob/9167c187f6a6f91f629187e3c3768c54f4eb0703/web/web_common/iso_codes/extension/usr/share/postgresql/17/extension/iso_codes--0.1.0.sql)

`iso_codes` 添加 `CountryCode` PostgreSQL 枚举，包含该扩展构建固化的两字母国家代码集合。需要把列限制在固定词表而不是任意文本时，可以使用它。

### 核心流程

```sql
CREATE EXTENSION iso_codes;

CREATE TABLE offices (
    name text NOT NULL,
    country CountryCode NOT NULL
);

INSERT INTO offices VALUES ('New York', 'US'), ('Berlin', 'DE');
SELECT * FROM offices WHERE country = 'US'::CountryCode;
```

枚举之外的值会被拒绝。生成的 0.1.0 SQL 包含该版本编译时收录的国家代码标签，其中包括 `XK`。

### SQL 对象

- `CountryCode` 是所核验生成 SQL 中唯一面向用户的 SQL 对象。
- 枚举顺序遵循生成产物的声明顺序；不要假设该顺序代表地理或业务优先级。

### 运维说明

0.1.0 版不可重定位。控制文件设置 `superuser = true` 和 `trusted = false`，因此创建需要超级用户；它未声明前置扩展或预加载要求。PostgreSQL 枚举成员固定在已安装 SQL 中，依赖新分配或变更的代码前应审查版本差异。
