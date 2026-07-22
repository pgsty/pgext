## 用法

来源：

- [Extension control file](https://github.com/df7cb/postgresql-hamradio/blob/4e2181d68b62d224b120ce1bb38d2c23ee63e937/hamradio.control)
- [Extension SQL files](https://github.com/df7cb/postgresql-hamradio/tree/4e2181d68b62d224b120ce1bb38d2c23ee63e937/sql)
- [Official example SQL](https://github.com/df7cb/postgresql-hamradio/blob/4e2181d68b62d224b120ce1bb38d2c23ee63e937/examples.sql)

hamradio 为业余无线电呼号、波段、频率、模式和 Maidenhead 网格定位提供 PostgreSQL 域、枚举、转换函数和参考表。它适合校验通联日志数据和推导地理网格，但只是一个小型领域工具集，并非完整日志应用。

### 核心流程

扩展依赖 `uctext` 和 `postgis`。先保存经过校验的呼号、频率及网格定位，再推导无线电波段或 PostGIS 几何：

```sql
CREATE EXTENSION uctext;
CREATE EXTENSION postgis;
CREATE EXTENSION hamradio;

CREATE TABLE contacts (
  worked_call call NOT NULL,
  frequency qrg NOT NULL,
  grid locator
);

INSERT INTO contacts VALUES ('DL1ABC/P', 14.074, 'JO62QM');

SELECT worked_call,
       band(frequency),
       ST_AsText(ST_LocatorPoint(grid))
FROM contacts;
```

### 重要对象

- `call` 校验由斜杠或连字符分隔的全大写呼号组成部分。
- `band` 枚举从 2190m 到 1mm 的业余波段；`band(numeric)` 把以 MHz 表示的频率映射到波段。
- `qrg` 只接受能映射到已知波段的频率。
- `locator` 校验 2 到 10 个字符的 Maidenhead 网格定位。
- `ST_Locator` 返回 SRID 4326 的定位网格多边形，`ST_LocatorPoint` 返回其点表示。
- `locator2` 和 `locator4` 是带 GiST 索引 PostGIS 网格的预填充参考表。
- `major_mode` 和 `adif_mode` 归一化常见模式标签。

### 运维说明

校验规则与波段界限固化在扩展 SQL 中，可能无法反映后续法规变更或本地波段规划。参考表会在安装时加入预计算的全球定位网格。应针对目标应用检查波段和几何推导语义，并保持各域要求的精确大写格式。
