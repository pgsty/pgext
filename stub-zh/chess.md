## 用法

来源：

- [扩展 control 文件](https://github.com/ValerioRocca/chess-postgres-extension/blob/1498da4d82588deb4ba49e7cedb9b0721ece6dae/chess.control)
- [SQL 安装脚本](https://github.com/ValerioRocca/chess-postgres-extension/blob/1498da4d82588deb4ba49e7cedb9b0721ece6dae/chess--1.0.sql)
- [C 实现](https://github.com/ValerioRocca/chess-postgres-extension/blob/1498da4d82588deb4ba49e7cedb9b0721ece6dae/chess.c)

`chess` `1.0` 版定义了用于棋局/走法文本和棋盘局面的 `san` 与 `fen` 类型。它还为 `san` 提供 B-tree 比较运算符，以及提取开局走法、重建指定回合后棋盘、判断棋局是否包含某种开局或局面的辅助函数。

### 示例

```sql
CREATE EXTENSION chess;
SELECT getFirstMoves('1.e4 e5 2.Nf3 Nc6'::san, 2);
SELECT getBoard('1.e4 e5 2.Nf3 Nc6'::san, 4);
```

应把它视为历史演示代码，而不是生产解析器。仓库几乎没有用户文档或兼容性声明；C 语言的存储及输入输出实现在定长值中嵌入指针，并执行不安全的定长复制。畸形乃至仅仅较长的输入都可能破坏内存或让后端崩溃。不要向不可信输入开放；评估时应隔离运行，并优先选择采用安全、自包含存储格式且仍在维护的棋类解析器。
