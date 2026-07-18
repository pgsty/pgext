## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/alexkuz/pg_unidecode/blob/9cb9868332773b4551837eb4eaf48f2903d80fd2/README.md)
- [0.0.6 版本 SQL 对象](https://github.com/alexkuz/pg_unidecode/blob/9cb9868332773b4551837eb4eaf48f2903d80fd2/unidecode--0.0.6.sql)
- [C 实现](https://github.com/alexkuz/pg_unidecode/blob/9cb9868332773b4551837eb4eaf48f2903d80fd2/unidecode.c)

`unidecode` 使用源自 Python Unidecode 项目的表，把 Unicode 文本音译成近似 ASCII 表示。

```sql
CREATE EXTENSION unidecode;
SELECT unidecode('Français, Русский, 漢語');
```

音译有损、不区分语言且不可逆。多个源字符串可能映射为同一 ASCII 输出，文化偏好的拼写可能不同，表变更也会改变结果。绝不能单独用于身份、唯一性、认证、法定姓名或安全规范化；必须保留原文并独立执行标识约束。

上游明确称 C 代码处于早期开发，并要求不要用于生产。0.0.6 没有当前兼容性或内存安全声明，应优先使用持续维护的应用库。遗留数据库若用于搜索提示，应固定精确映射表，测试所有支持文字系统和畸形 UTF-8 边界，限制输入长度，并在任何升级后重建生成列与索引。
