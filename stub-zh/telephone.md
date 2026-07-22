## 用法

来源：

- [已复核 commit 的官方 telephone 类型文档](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/doc/telephone.md)
- [已复核 commit 的官方 README](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/README.md)
- [SQL 对象与操作符类](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/sql/telephone.sql)
- [解析器与表示实现](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/src/telephone.c)
- [扩展 control 文件](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/telephone.control)

`telephone` 增加带 B-tree 和哈希索引的规范化电话号码类型。它保留显示元数据，同时派生出忽略标点并把字母映射为电话键盘数字的身份值。它还表示暂停、确认、分机和部分号码指令。0.0.1 版本标为不稳定，其拨号计划数据也已陈旧。

### 核心流程

```sql
CREATE EXTENSION telephone;

CREATE TABLE contact_phone (
  phone telephone PRIMARY KEY,
  label text NOT NULL
);

INSERT INTO contact_phone VALUES ('1 (800) 555-01AZ', 'support');

SELECT phone, telephone_to_format(phone, 'E123')
FROM contact_phone
WHERE phone = '18005550129'::telephone;
```

以 `+` 开头的值进入国家代码模式，并根据扩展内嵌的号码计划数据检查。其他值使用数字模式。需要严格验证国家代码时，应直接转换 `+` 形式；包装器可能捕获错误并退回数字模式，从而掩盖无效输入与不支持输入之间的差异。

### 主要对象

`telephone_mode_get`、`telephone_calling_code_get`、`telephone_service_get`、`telephone_fictitious_get`、`telephone_fictitious_is`、`telephone_geo_is`、`telephone_geo_parts_get`、`telephone_domestic_numbers_get`、`telephone_extension_numbers_get` 和 `telephone_domestic_prefer_get` 用于检查规范化值。`telephone_domestic_assume_set` 与 `telephone_set` 用于构造或重新解释号码。比较操作符、B-tree 与哈希操作符类，以及 `min`/`max` 等聚合都基于电话身份，而非原始格式。

### 数据与安全边界

国际支持并不完整；北美编号计划是最完善的路径，许多国家代码只部分支持或完全不支持。号码计划会随时间变化，因此不能把内嵌数据当作当前运营商、路由、紧急号码或监管验证依据。

SQL 定义了从 `bytea` 到 `telephone` 的二进制兼容赋值转换，且没有验证函数。不可信原始字节可以绕过文本解析器并破坏内部假设；应撤销或避免该路径。在把类型用于持久约束前，应测试等值、格式化、虚构号码范围、地理辅助函数、备份恢复，以及每个接受的拨号计划。
