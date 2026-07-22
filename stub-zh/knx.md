## 用法

来源：

- [官方 README](https://github.com/mike-kfed/pgknx/blob/c485a440f7347a0bd840018054466735d8f5f9ce/README.md)
- [类型定义](https://github.com/mike-kfed/pgknx/blob/c485a440f7347a0bd840018054466735d8f5f9ce/knx.sql)
- [地址解析器](https://github.com/mike-kfed/pgknx/blob/c485a440f7347a0bd840018054466735d8f5f9ce/knx.c)
- [操作符与索引类](https://github.com/mike-kfed/pgknx/blob/c485a440f7347a0bd840018054466735d8f5f9ce/knx-operators.sql)

`knx` 1.0 为 KNX 组地址和个体地址增加紧凑的双字节类型。它们可验证常见 KNX 文本格式、保留底层 16 位地址值，并支持数值排序以及 B-tree 和哈希索引。

### 核心流程

根据地址族选择合适的显示形式：

```sql
CREATE EXTENSION knx;

CREATE TABLE knx_endpoint (
  group_address knx_group_address3 PRIMARY KEY,
  device_address knx_individual_address
);

INSERT INTO knx_endpoint VALUES ('17/6/1', '4.3.250');
SELECT group_address, group_address::integer FROM knx_endpoint;
```

双向整数转换会暴露打包后的 16 位数值。三个地址类型都拒绝零值。

### 类型索引

- `knx_group_address3` 显示为 `main/middle/sub`，范围分别为 0–31、0–7 和 0–255。
- `knx_group_address2` 显示为 `main/sub`，范围分别为 0–31 和 0–2047。
- `knx_individual_address` 显示为 `area.line.device`，范围分别为 0–15、0–15 和 0–255。

每种类型还接受十六进制打包地址。比较依据打包后的数值，KNX 地址类型之间也提供跨类型比较操作符。

### 运维说明

两个组地址类型的解析器都会接受二层或三层斜杠输入；声明类型只决定规范输出格式。如果层级差异很重要，应在应用中或 CHECK 约束中验证输入形态。解析器拒绝零值或广播地址，以及超出文档范围的数值。建模 KNX 系统前，应确认广播和保留地址是否需要单独表示，而不是强行放入这些类型。
