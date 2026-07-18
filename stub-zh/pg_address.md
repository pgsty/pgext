## 用法

来源：

- [扩展 README](https://github.com/lotint/dbtools/blob/3966ce98177cfa3b2e99a1871f08d6bb52a1a4ce/pg_extensions/pg_address/README.md)
- [扩展 control 文件](https://github.com/lotint/dbtools/blob/3966ce98177cfa3b2e99a1871f08d6bb52a1a4ce/pg_extensions/pg_address/pg_address.control)
- [0.0.3 版 SQL 定义](https://github.com/lotint/dbtools/blob/3966ce98177cfa3b2e99a1871f08d6bb52a1a4ce/pg_extensions/pg_address/sql/pg_address.sql)
- [冒烟测试](https://github.com/lotint/dbtools/blob/3966ce98177cfa3b2e99a1871f08d6bb52a1a4ce/pg_extensions/pg_address/tests/smoke.sql)

`pg_address` 0.0.3 定义了名为 `pg_address` 的复合类型，用国家、地区、城市、邮编、街道、门牌号、街区、纬度和经度字段存储地址。

### 存储并读取复合地址

```sql
CREATE EXTENSION pg_address;

CREATE TABLE customer_location (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    address pg_address
);

INSERT INTO customer_location (address)
VALUES (ROW('Germany', 'Berlin', 'Berlin', '10117',
            'Hauptstrasse', '57b', 'Mitte', 52.5200, 13.4050)::pg_address);

SELECT (address).city, (address).zip_code, (address).lat, (address).lon
FROM customer_location;
```

字段顺序依次为 `country`、`region`、`city`、`zip_code`、`street`、`num`、`suburb`、`lat` 和 `lon`。应使用显式类型转换或按字段名赋值，避免位置参数悄然错位。

### 数据模型边界

它只是存储类型，不是地址解析器、地理编码器或校验器。文本字段可接受任意值，坐标没有范围约束，也没有提供规范化功能或索引运算符类。应按应用要求为经纬度、国家代码和必填字段补充约束。

复合类型变更会影响依赖它的表、函数、客户端以及位置式 row 字面量。仓库安装说明面向 PostgreSQL 9.6 时代的开发包，且没有发布当前兼容矩阵；采用前应在实际 PostgreSQL 版本上验证扩展升级和客户端类型解码。
