## 用法

来源：

- [PGXN 1.0.0 README](https://pgxn.org/dist/gsscode/1.0.0/README.html)
- [gsscode 控制文件](https://api.pgxn.org/src/gsscode/gsscode-1.0.0/gsscode.control)
- [gsscode 1.0.0 SQL 定义](https://api.pgxn.org/src/gsscode/gsscode-1.0.0/gsscode--1.0.0.sql)
- [PostgreSQL 许可证](https://api.pgxn.org/src/gsscode/gsscode-1.0.0/LICENSE)

`gsscode` 将英国国家统计局/政府统计服务的九字符地理编码压缩为 32 位。需要紧凑存储 GSS 编码、btree 排序、精确比较或按国家/类型前缀进行索引匹配时可使用它；它校验编码格式，但不证明某个编码确实存在于 ONS 登记表中。

### 核心流程

```sql
CREATE EXTENSION gsscode;

CREATE TABLE areas (
    code gsscode PRIMARY KEY,
    label text NOT NULL
);

INSERT INTO areas VALUES ('E01000001', 'Example LSOA');

SELECT code, country(code), gss_type(code), area(code)
FROM areas
WHERE code % 'E01';
```

`%` 接受一个国家字母、三字符国家/类型前缀、完整编码或前缀数组。它属于 `gsscode_ops` btree 操作符族，因此 btree 索引可以直接处理受支持的前缀条件。`!%` 是其否定形式。

### 对象与登记数据

- `gsscode` 是压缩类型；其文本形式始终为大写九字符表示。
- `is_valid(text)` 在不抛出输入错误的情况下检查可接受的词法格式。
- `country(gsscode)`、`gss_type(gsscode)` 和 `area(gsscode)` 返回压缩后的各组成部分。
- `description(...)` 与 `type_info(...)` 在 `gsscode_types` 中查询三字符地理类型。
- `isnan(gsscode)` 识别六位区域部分为 `999999` 的 ONS 保留编码。
- `left(gsscode, integer)` 以及 `~`、`~*`、`!~`、`!~*` 可在列类型转换后保留常见文本查询写法。

`gsscode_types` 保存随扩展提供的小型地理类型登记表，而不是数十万个具体区域名称。PostgreSQL 会把该表纳入扩展配置转储，因此本地刷新后的行可在备份恢复后保留。

### 边界

输入必须严格为一个字母后接八位数字。算术压缩有意接受未来的国家/类型组合，所以解析成功不代表 ONS 已分配该编码。通用正则表达式与 `left(...)` 会先把值渲染为文本，不能利用压缩前缀操作符类；字面前缀查询应优先使用 `%`，或为反复使用的 `left(...)` 条件建立表达式索引。

1.0.0 版本没有发布 PostgreSQL 主版本支持矩阵，部署前应在目标服务器上验证构建。刷新 `gsscode_types` 是可选操作，由独立的 `gsscode_ons_refresh` 扩展提供。
