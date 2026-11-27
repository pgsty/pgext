## 用法

来源：

- [Official source](https://github.com/sanchitachaurasia/CS-631/blob/79c0152b5fcd67a2e847c3b128d0dbdef503f642/README.md)
- [Official source](https://github.com/sanchitachaurasia/CS-631/blob/79c0152b5fcd67a2e847c3b128d0dbdef503f642/project/extension/card_feedback.control)
- [Official source](https://github.com/sanchitachaurasia/CS-631/blob/79c0152b5fcd67a2e847c3b128d0dbdef503f642/project/extension/card_feedback--1.0.sql)
- [Official source](https://github.com/sanchitachaurasia/CS-631/blob/79c0152b5fcd67a2e847c3b128d0dbdef503f642/project/extension/card_feedback.c)
- [Official source](https://github.com/sanchitachaurasia/CS-631/blob/79c0152b5fcd67a2e847c3b128d0dbdef503f642/project/extension/Makefile)

`card_feedback` 1.0 为 PostgreSQL 16.13 中研究性基数反馈实现提供检查和重置函数。它依赖该项目修改过的规划器和反馈存储 API，不能通过安装模块就在原版服务器上获得这些能力。

### 检查与重置

在隔离的研究环境中安装配套修改版服务器和模块，再以超级用户创建扩展：

```sql
CREATE EXTENSION card_feedback;
SELECT * FROM card_feedback_entries();
SELECT card_feedback_reset();
```

`card_feedback_entries()` 返回签名、原始估算行数、最近观测行数、指数平滑行数、观测次数和更新时间。`card_feedback_reset()` 清空反馈存储并返回删除条目数，仅在确实需要重置实验时调用。

### 权限与结果解释

两个 SQL 函数均为 volatile，且不允许并行执行。安装脚本没有撤销默认 PUBLIC 执行权限，C 包装函数也没有显式超级用户检查。允许其他用户连接前应限制函数执行权，尤其是重置函数。

SQL 扩展本身不会启用规划器修改或创建底层反馈存储，这些属于修改过的内核。空结果并不能证明工作负载已被观测。应在项目规定的基准配置下比较估算与实际执行统计。该检查包装层未声明共享预加载要求。

### 维护

扩展可重定位，但 ABI 依赖配套内核补丁。变更 PostgreSQL 版本时需一起重建并验证两者。已核对源码提供 1.0 安装脚本，没有 SQL 升级链；实验输入和结果需另行保存。已核对文件中未明确项目自有修改的许可证。
