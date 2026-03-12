

## 用法

> [Address Standardizer：PostGIS 的地址解析与标准化](https://github.com/postgis/postgis)

Address Standardizer 是一个 PostGIS 扩展，使用可配置的词典、语法和规则表将单行地址字符串解析为结构化形式。它是 TIGER 地理编码器中内置 `normalize_address` 函数的更灵活替代方案。

- [Address Standardizer 参考手册](https://postgis.net/docs/Extras.html#Address_Standardizer)

### 安装

```sql
CREATE EXTENSION address_standardizer;
```

--------

## 标准化地址

核心函数接收一个地址字符串和三个表引用（lex、gaz、rules）：

```sql
SELECT *
FROM standardize_address(
    'us_lex',        -- 词典表
    'us_gaz',        -- 地名表
    'us_rules',      -- 规则表
    '1600 Pennsylvania Ave NW, Washington, DC 20500'
);
```

返回结果包含以下结构化字段：

| 字段 | 描述 |
|------|------|
| `building` | 建筑名称或标识 |
| `house_num` | 门牌号 |
| `predir` | 前缀方向（N、S、E、W） |
| `qual` | 限定符 |
| `pretype` | 前缀类型 |
| `name` | 街道名称 |
| `suftype` | 后缀类型（St、Ave、Blvd） |
| `sufdir` | 后缀方向 |
| `ruralroute` | 乡村路线 |
| `extra` | 附加信息 |
| `city` | 城市名 |
| `state` | 州 |
| `country` | 国家 |
| `postcode` | 邮政编码 |
| `box` | 邮政信箱 |
| `unit` | 单元/公寓号 |

--------

## 词典、地名和规则表

标准化器由三个用户可配置的表驱动：

**词典（lex）** -- 将输入词元映射为标准化形式和词元类别：

```sql
CREATE TABLE us_lex (
    id serial PRIMARY KEY,
    seq integer,
    word text,
    stdword text,
    token integer
);
```

**地名表（gaz）** -- 将地名（城市、州）映射为标准形式：

```sql
CREATE TABLE us_gaz (
    id serial PRIMARY KEY,
    seq integer,
    word text,
    stdword text,
    token integer
);
```

**规则（rules）** -- 定义地址解析的语法规则：

```sql
CREATE TABLE us_rules (
    id serial PRIMARY KEY,
    rule text
);
```

对于美国地址，`address_standardizer_data_us` 扩展提供了这些表的预置数据。
