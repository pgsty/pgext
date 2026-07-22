## 用法

来源：

- [官方 README](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/README.md)
- [扩展控制文件](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/postgresml.control)
- [实际安装的扩展 SQL](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/postgresml--1.0.sql)
- [独立的朴素贝叶斯演示](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/naive_bayes.sql)
- [构建清单](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/Makefile)

`postgresml` 1.0 不是通用机器学习平台。安装扩展只会暴露一个辅助函数 `ts_lexemes(tsvector)`；仓库中的二元朴素贝叶斯垃圾短信分类器是单独的演示脚本，不会被 `CREATE EXTENSION` 安装。

### 核心流程

```sql
CREATE EXTENSION postgresml;

SELECT lexeme, n
FROM ts_lexemes(to_tsvector('english', 'red fox red dog'))
ORDER BY lexeme;
```

`ts_lexemes` 为输入 `tsvector` 中的每个词位返回一行。字段 `n` 是该词位保存的位置条目数。它适合把全文检索向量展开成关系行；但向量经过剥离或以其他方式构造时，位置计数可能缺失或发生变化，因此不能假设 `n` 总是原始词元频数。

### 演示脚本边界

README 中的 `test_naive_bayes(0.3)` 流程属于 `naive_bayes.sql`，不在 Makefile 列出的扩展安装脚本中。该演示会创建通用名称的表和函数，依赖 `tablefunc` 的 `crosstab` 功能，使用 psql 文件加载命令，并要求固定仓库中并不存在的外部 SMS Spam Collection 数据文件。它还随机划分训练集，因此结果不确定。

应显式审查并改写演示，不能假设安装扩展后就会出现。实验时应使用受控模式、版本化数据、确定性划分、留出评估以及适合应用的隐私和保留策略。C 辅助函数访问 PostgreSQL 全文检索内部结构，固定源码来自 2013 年，尚未建立当前大版本兼容性；必须针对确切服务器大版本编译、测试并验证转储恢复。不需要预加载。
