## 用法

来源：

- [官方 README](https://github.com/tatut/pgprolog/blob/e3629cc2f5a56da98c22939436d9d803e2f9a6d5/README.md)
- [Rust 语言处理器实现](https://github.com/tatut/pgprolog/blob/e3629cc2f5a56da98c22939436d9d803e2f9a6d5/src/lib.rs)
- [扩展控制文件](https://github.com/tatut/pgprolog/blob/e3629cc2f5a56da98c22939436d9d803e2f9a6d5/pgprolog.control)

`pgprolog` 是一个把 Scryer Prolog 嵌入 PostgreSQL 的概念验证型过程语言处理器。函数体可以包含 Prolog 事实与 `handle/2` 谓词，但上游明确说明它尚不能用于实际场景。

### 最小流程

安装扩展仅限超级用户。扩展不会自动创建语言，因此定义函数前要用已安装处理器创建 `plprolog`。

```sql
CREATE EXTENSION pgprolog;
CREATE LANGUAGE plprolog HANDLER plprolog_call_handler;

CREATE FUNCTION parent_of(text)
RETURNS text
AS $$
parent(alice,bob).
handle(In,Out) :- parent(In,Out).
$$ LANGUAGE plprolog;

SELECT parent_of('alice');
```

发布的示例只演示文本输入与文本输出。项目没有记录生产级类型映射矩阵、触发器支持、集合返回函数、事务/SPI 访问、模块、持久 Prolog 状态或资源控制。

### 信任与原型边界

不带 `TRUSTED` 的 `CREATE LANGUAGE` 会创建不可信语言，因此只有超级用户能用它创建函数。应保持其不可信属性：任意 Prolog 程序会在 PostgreSQL 后端中执行一个年轻的嵌入式运行时，解析、分配、递归或 panic 缺陷都可能终止或耗尽该后端。

`0.0.0` 版固定使用 pgrx `0.11.3` 与 Scryer Prolog `0.9.4`；声明的 Cargo 特性只到 PostgreSQL 16。只能在带语句和进程限制的隔离实验室中使用。若要进一步采用，处理器需要明确的类型/错误语义、取消测试、内存和递归边界、权限审查、模糊测试、崩溃隔离以及持续维护的兼容策略。
