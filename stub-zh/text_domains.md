## 用法

来源：

- [官方 README](https://github.com/cahutton/text_domains/blob/a4e1aa7463e07ffdf1ce9f2d380f054bf1af1b5c/README.md)
- [官方扩展 SQL](https://github.com/cahutton/text_domains/blob/a4e1aa7463e07ffdf1ce9f2d380f054bf1af1b5c/text_domains--1.0.sql)
- [官方控制文件](https://github.com/cahutton/text_domains/blob/a4e1aa7463e07ffdf1ce9f2d380f054bf1af1b5c/text_domains.control)

`text_domains` 版本 `1.0` 定义了 68 个可复用 PostgreSQL 域，用于非空文本，以及固定形状的字母数字、纯字母、大写、小写和数字字符串。这些域在保留普通 `text` 存储与操作符的同时，集中管理简单验证规则。

### 核心流程

安装扩展，并直接在表定义中使用这些域：

```sql
CREATE EXTENSION text_domains;

CREATE TABLE country_codes (
  code alpha2 NOT NULL PRIMARY KEY,
  display_name nonempty_text NOT NULL,
  numeric_code digit3
);

INSERT INTO country_codes VALUES ('US', 'United States', '840');
```

不满足域检查约束的值会在赋值时被拒绝。

### 域家族

- `nonempty_text` 拒绝空字符串。
- `alnum` 与 `nonempty_alnum` 分别接受零个以上和一个以上字母数字字符；`alnum1` 到 `alnum10`、`alnum16`、`alnum22` 等选定固定宽度要求字符数恰好匹配。
- `alpha`、`upper`、`lower` 与 `digit` 提供对应的 POSIX 字符类家族，包括非空和固定宽度变体。
- 所有带模式的域都在扩展 SQL 中使用 `COLLATE "C"` 与 `SIMILAR TO` 检查。

### 运维注意事项

除非表字段另行声明 `NOT NULL`，否则域检查会接受 NULL；即使 `nonempty_text` 也不会让字段自动变成非空。SQL 文件中的固定宽度是显式且不连续的集合，不能假设每个长度都存在，应直接检查文件。字符类与排序规则行为需要结合部署编码和有代表性的非 ASCII 输入进行测试。这些域刻意只做简单约束，不会统一大小写、修剪空白或提供应用专用格式。修改共享域约束会影响所有依赖字段，可能需要先验证已有数据。
