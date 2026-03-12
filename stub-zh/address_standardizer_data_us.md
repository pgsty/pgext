

## 用法

> [Address Standardizer Data US：address_standardizer 扩展的美国地址数据](https://github.com/postgis/postgis)

此扩展为 `address_standardizer` 扩展提供预置的美国词典、地名和规则数据。它包含常见美国街道类型、方向缩写、州名以及解析美国地址所需的语法规则表。

- [Address Standardizer 参考手册](https://postgis.net/docs/Extras.html#Address_Standardizer)

### 安装

```sql
CREATE EXTENSION address_standardizer_data_us;
```

这将在 public 模式中创建预填充了美国地址数据的 `us_lex`、`us_gaz` 和 `us_rules` 表。

--------

## 与 address_standardizer 配合使用

安装后即可立即标准化美国地址：

```sql
SELECT *
FROM standardize_address(
    'us_lex', 'us_gaz', 'us_rules',
    '123 Main Street NW, Apt 4B, Springfield, IL 62704'
);
```

提供的数据涵盖常见的美国地址模式，包括：

- 街道类型（Street、Avenue、Boulevard、Drive、Lane、Court 等）
- 方向前缀和后缀（North、South、N、S、NW、SE 等）
- 州名及其缩写
- 单元标识符（Apt、Suite、Unit 等）
- 公路标识符（US、State、County、Interstate 等）
