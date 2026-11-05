## 用法

来源：

- [上游来源 1](https://github.com/dbgreen/pg_scalpel/blob/3233758db18b2b664360037a43098a114799846f/README.md)
- [上游来源 2](https://github.com/dbgreen/pg_scalpel/blob/3233758db18b2b664360037a43098a114799846f/pg_scalpel.control)

`pg_scalpel` 0.1 在 PostgreSQL 19 的 Linux x86-64 后端中提供 Janet 解释器，用于调试与插桩。

### 核心流程

`janet_eval(text)` 将 Janet 最后一个值以文本返回，虚拟机在整个连接期间保留。scalpel 基础操作可解析符号、检查结构、调用 C 函数、安装钩子和查看机器码。静态符号及布局需要匹配的 DWARF 调试信息。

```sql
CREATE EXTENSION pg_scalpel;
SELECT janet_eval('(+ 40 2)');
```

### 运行边界

无需预加载或重启服务。安装及执行均被强制限制为超级用户操作，撤销 PUBLIC 权限不能替代运行时检查。Janet 与原生基础操作可以执行系统命令、读写任意内存和修改机器码。应在可丢弃的诊断环境中使用，不能将解释器视为沙箱。
