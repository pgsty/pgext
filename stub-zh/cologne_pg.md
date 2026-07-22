## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/DrSloth/cologne_pg/blob/6871c07bb24c9a05e63b5e2b034bf34eda64426c/README.md)
- [Rust 实现](https://github.com/DrSloth/cologne_pg/blob/6871c07bb24c9a05e63b5e2b034bf34eda64426c/src/lib.rs)
- [Rust 包清单](https://github.com/DrSloth/cologne_pg/blob/6871c07bb24c9a05e63b5e2b034bf34eda64426c/Cargo.toml)

`cologne_pg` 把科隆语音算法暴露为 `cologne(text)`，生成用于近似匹配德语姓名的语音代码。

```sql
CREATE EXTENSION cologne_pg;
SELECT cologne('MORELLO');

SELECT person_id, display_name
FROM people
WHERE cologne(display_name) = cologne('Müller');
```

语音相等只是候选生成启发式，并不等同身份。不同姓名可能共享代码，其他语言拼写可能表现不佳，音译、标点、大小写、Unicode 规范化和版本变化也会影响结果。必须保留原字符串，并加入额外排序或人工复核步骤。

若在表达式索引或生成列中使用该函数，应固定精确实现，并在行为变化后重建相关数据。用于搜索前，应测试代表性德语姓名、变音符号、空串、`NULL`、超长输入及对抗性碰撞率。
