## 用法

来源：

- [官方扩展 SQL](https://github.com/cahutton/numeric_domains/blob/50496725313939a340f0c8129d9e9f2491009e51/numeric_domains--1.0.sql)
- [官方 README](https://github.com/cahutton/numeric_domains/blob/50496725313939a340f0c8129d9e9f2491009e51/README.md)
- [官方扩展控制文件](https://github.com/cahutton/numeric_domains/blob/50496725313939a340f0c8129d9e9f2491009e51/numeric_domains.control)

`numeric_domains` 在 `smallint`、`integer`、`bigint` 和 `numeric` 上安装可复用的符号约束域，并提供有界的 `latitude` 与 `longitude` 域。它可以把简单数值不变量紧凑地放进 PostgreSQL 类型系统。

### 核心流程

创建扩展后，可在列或函数签名中使用这些域：

```sql
CREATE EXTENSION numeric_domains;

CREATE TABLE measurements (
  retry_count nonnegative_integer,
  adjustment nonzero_numeric,
  lat latitude,
  lon longitude
);

INSERT INTO measurements VALUES (0, -1.25, 39.9042, 116.4074);
```

超出范围的非空值会违反域检查：

```sql
INSERT INTO measurements(lat, lon) VALUES (91, 0);
```

### 已安装的域

`smallint`、`integer`、`bigint` 和 `numeric` 各自都有 `negative_*`、`nonpositive_*`、`nonzero_*`、`nonnegative_*` 和 `positive_*` 变体。`latitude` 接受 `-90` 到 `90` 的值，`longitude` 接受 `-180` 到 `180` 的值。

### 语义注意事项

- 所有域都没有声明 `NOT NULL`。域的 `CHECK` 对 `NULL` 求值为未知并会通过；如果业务不允许缺失值，应再添加列级 `NOT NULL` 约束。
- 这些是域，而不是独立算术类型。表达式通常使用底层数值类型，结果重新赋给域时会再次检查约束。
- `latitude` 与 `longitude` 的边界是包含式的；扩展没有定义坐标参考系、轴顺序、规范化或地理空间运算符。
- 安装 SQL 中多个 `COMMENT ON DOMAIN` 目标存在复制粘贴错误，会覆盖若干 `negative_*` 域的注释。检查仍附着在正确对象上，但系统目录描述可能误导。
- 采用这些无模式限定的全局名称前，应检查是否与已有域冲突，并决定它们应由哪个模式承载；该扩展在安装时可重定位。
