## 用法

来源：

- [官方 README](https://github.com/alexeyklyukin/translate_isbn/blob/dd65576710245916ab14e7ff00d8f94058e63e23/README.md)
- [扩展 SQL API](https://github.com/alexeyklyukin/translate_isbn/blob/dd65576710245916ab14e7ff00d8f94058e63e23/translate_isbn--1.0.sql)
- [转换实现](https://github.com/alexeyklyukin/translate_isbn/blob/dd65576710245916ab14e7ff00d8f94058e63e23/translate_isbn.c)

`translate_isbn` 1.0 提供 `translate_isbn1013(text)`，是与 Evergreen 图书馆系统行为兼容的搜索扩展辅助函数。它识别以空白分隔、类似 ISBN 的 token，去除标点，检查/修正校验位，并在可转换时输出 ISBN-10 与 ISBN-13 备选值。

### 核心流程

```sql
CREATE EXTENSION translate_isbn;

SELECT translate_isbn1013('9789997122575');
-- 9789997122575 9997122577

SELECT translate_isbn1013('ISBN 0-596-52068-9');
```

对于每个识别到的 10 或 13 字符候选值，输出都会包含规范化原值。若校验位错误，还会追加修正版。随后追加有效转换：ISBN-10 会变成带 `978` 前缀的 ISBN-13，只有兼容前缀的 ISBN-13 才能变成 ISBN-10。结果是空格分隔文本，用于扩大搜索词范围。

### 验证边界

函数检查数字形状和校验和，但刻意不验证 ISBN 注册组或出版社代码。因此，校验和有效但未分配的代码也可能被转换。若没有识别任何 token，则原样返回输入；一旦混合输入中找到候选值，非 ISBN token 可能从结果中消失。

不要把该函数作为规范 ISBN 数据类型、目录权威或出版验证器。上游 README 建议使用 PostgreSQL contrib 的 `isn` 扩展进行更完整验证和带类型存储。C 函数声明为 `STRICT` 和 `IMMUTABLE`，只接受文本，并可能通过去标点与连接备选值改变格式。
