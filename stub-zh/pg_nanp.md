## 用法

来源：

- [官方扩展控制文件](https://github.com/rlichtenwalter/pg_nanp/blob/8cfb5213faaeef92a19997978f0b6801f0f61aa0/pg_nanp.control)
- [官方 README](https://github.com/rlichtenwalter/pg_nanp/blob/8cfb5213faaeef92a19997978f0b6801f0f61aa0/README.md)
- [官方扩展 SQL](https://github.com/rlichtenwalter/pg_nanp/blob/8cfb5213faaeef92a19997978f0b6801f0f61aa0/pg_nanp--1.0.sql)

`pg_nanp` 1.0 增加了北美号码规划电话号码的 `nanp` 基础类型。输入函数接受常见分隔符和可选的前置国家代码，校验十位 NANP 结构，并始终把存储值显示为 `AAA-EEE-NNNN`。

### 核心流程

```sql
CREATE EXTENSION pg_nanp;

SELECT '2345678910'::nanp;
SELECT '1 (234) 567-8910'::nanp;

CREATE TABLE contacts (
    contact_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    phone nanp NOT NULL
);

INSERT INTO contacts (phone) VALUES ('234.567.8910');
```

可接受的标点仅限输入正则表达式指定位置上的可选空格、连字符、点号和括号。可选国家前缀必须是 `1`。区号与中心局代码的首位都必须为 2–9，并且会拒绝 211 或 911 之类的 N11 中心局代码。

### 类型行为

`nanp` 以八字节传值整数保存，只定义了文本输入和输出函数。该扩展没有定义比较或相等运算符、与 `text` 或整数类型之间的转换、B-tree 运算符类、其他格式、分机号或地区和分配查询。

由于缺少比较运算符，若不额外定义运算符，就不能直接在 `nanp` 上使用普通 `PRIMARY KEY`、`UNIQUE`、`ORDER BY`、等值连接和 B-tree 索引。该类型主要适合输入校验与显示；需要检索和索引时，应另存规范化文本或数字伴随列。

### 校验边界

正则表达式只验证 1.0 版本编码的 NANP 结构规则。它不会判断区号或交换局是否已经分配、号码是否可达、是否为移动号码或所有权是否变化。上游实现是一个小型历史 C 扩展，采用前应针对精确 PostgreSQL 版本和区域设置进行测试。
