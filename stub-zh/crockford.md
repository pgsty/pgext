## 用法

来源：

- [官方 README](https://github.com/grzm/pgcrockford/blob/774cfee2b77086c096547d1b6767cd1e89076f88/README.markdown)
- [扩展 SQL](https://github.com/grzm/pgcrockford/blob/774cfee2b77086c096547d1b6767cd1e89076f88/crockford.sql)
- [输入输出实现](https://github.com/grzm/pgcrockford/blob/774cfee2b77086c096547d1b6767cd1e89076f88/inout.c)

`crockford` 0.8.34 提供采用 Crockford Base32 文本表示的紧凑无符号整数类型。它适合需要大小写不敏感、能够容忍易混淆字符的人类可读标识符，同时保留数值比较、算术运算和索引能力。

### 核心流程

先选择存储宽度，再插入 Base32 文本；若输入表示数值本身，则使用整数强制转换：

```sql
CREATE EXTENSION crockford;

CREATE TABLE ticket (
  id crockford4 PRIMARY KEY,
  label text NOT NULL
);

INSERT INTO ticket VALUES ('10', 'Base32 text'), (10::crockford4, 'Integer value');
SELECT id, id + 1 FROM ticket ORDER BY id;
```

文本输入按 Base32 解码，而从整数转换会保留整数值。因此，文本字面量 `10` 表示十进制 32，整数 10 则显示为 `A`。

### 类型与操作符索引

- `crockford2` 存储 16 位无符号值，上限为 `1ZZZ`。
- `crockford4` 存储 32 位无符号值，上限为 `3ZZZZZZ`。
- `crockford8` 存储 64 位无符号值，上限为 `FZZZZZZZZZZZZ`。
- 扩展提供比较、算术和位运算操作符，以及 B-tree 与哈希操作符类。

解析器接受大小写输入，并将易混淆的字母 I/L 与 O 映射为 1 和 0；Crockford 字母表排除 U。

### 运维说明

这些类型是无符号类型，因此转换为 PostgreSQL 有符号整数时可能溢出，算术运算也会执行范围检查。若扩展安装模式不在 `search_path` 中，应使用 PostgreSQL 的 `OPERATOR(schema.operator)` 语法限定操作符模式。已复核的上游版本仅声明测试 PostgreSQL 11；在较新大版本上应验证编译、转储恢复、转换与溢出行为。
