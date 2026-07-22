## 用法

来源：

- [官方 README](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/brainfuck/README.md)
- [官方 control 文件](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/brainfuck/brainfuck.control)
- [官方 4.2.0 版 SQL](https://github.com/mansueli/tle/blob/919de6247f4126dd11cd20e36602e0c5baf6fcaf/brainfuck/brainfuck--4.2.0.sql)
- [官方软件包页面](https://database.dev/mansueli/brainfuck)

`brainfuck` 通过 PLV8 在 PostgreSQL 内运行 Brainfuck 程序。它适合演示和实验，不适合承载应用逻辑：任意程序会在数据库后端中执行 JavaScript，并可能消耗大量 CPU。

### 核心流程

先安装 `plv8`，再创建 `brainfuck`：

```sql
CREATE EXTENSION plv8;
CREATE EXTENSION brainfuck;

SELECT brainfuck('++++++++[>++++++++<-]>+.', '');
```

第二个参数提供由 `,` 指令读取的输入；程序不需要输入时传空字符串。

### API

`brainfuck(code text, input text) RETURNS text` 是扩展唯一的 SQL 函数。它声明为 `IMMUTABLE STRICT`，使用 256 个单元的循环纸带，实现八条 Brainfuck 指令，并忽略 `code` 中的其他字符。

### 注意事项

扩展依赖 `plv8`；软件包注册表的安装流程还使用 `pg_tle`/`dbdev`，传统扩展安装则不需要。实际 SQL 签名有两个参数，尽管上游 README 的一个示例只传了一个参数。

解释执行的程序没有专门的运行时间限制或沙箱。括号不匹配会在生成的 JavaScript 编译阶段失败，长时间运行或不终止的程序会持续占用 PostgreSQL 后端。不要把它暴露给不可信 SQL，只应在受控实验中使用。
