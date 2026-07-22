## 用法

来源：

- [官方 README](https://github.com/jjw12345/postgresql-extension/blob/98b82c8a4f621d5f7f6050c57a8313502581b3fb/README.md)
- [官方 1.0 版 SQL](https://github.com/jjw12345/postgresql-extension/blob/98b82c8a4f621d5f7f6050c57a8313502581b3fb/recommend--1.0.sql)
- [官方 C 实现](https://github.com/jjw12345/postgresql-extension/blob/98b82c8a4f621d5f7f6050c57a8313502581b3fb/recommend.c)
- [官方 control 文件](https://github.com/jjw12345/postgresql-extension/blob/98b82c8a4f621d5f7f6050c57a8313502581b3fb/recommend.control)

`recommend` 是 PostgreSQL 9.5 时代的实验扩展，提供数值距离、文本字符向量相似度，以及源自 `smlar` 的一维数组相似度函数。它不是完整推荐系统：候选选择、排序、持久化与评估都需要调用方围绕这些底层分数自行构建。

### 核心流程

安装经过下文所述审查与修复的构建后，预期 API 如下：

```sql
CREATE EXTENSION recommend;

SELECT twoint(10, 4);
SELECT two_float(1.5, 4.0);
SELECT arraysml(ARRAY[1, 2, 3], ARRAY[2, 3, 4]);
```

`twoint` 与 `two_float` 返回绝对差；`arraysml` 比较元素类型相同的数组，捆绑代码默认使用余弦风格的集合相似度公式。

### API

- `twoint(integer, integer) RETURNS integer`：整数绝对差。
- `two_float(double precision, double precision) RETURNS double precision`：浮点绝对差。
- `similarity(text, text) RETURNS real`：自定义字符计数向量的余弦相似度。
- `arraysml(anyarray, anyarray) RETURNS real`：可比较一维数组的相似度，包含捆绑的 `smlar` 公式/配置代码。

### 注意事项

固定的官方 C 源码按发布内容无法复现构建：`generate_vect` 函数缺少开始花括号。其 `similarity` 实现还直接用每个输入字节减去小写 `a` 作为数组下标，且没有边界检查；大写、标点、多字节文本及其他字节都可能访问向量范围之外的内存。没有审查并修补源码前，不能开放该函数或宣称构建已经验证。

代码面向旧版 PostgreSQL 内部 API，除了 README 声称的 PostgreSQL 9.5 外没有维护中的兼容说明。`twoint` 在极端整数上可能溢出，空文本的余弦计算可能除零，数组评分还取决于元素比较与规范化。所有结果都应视为实验分数，并用对抗性输入验证修补后的构建。
