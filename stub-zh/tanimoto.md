## 用法

来源：

- [扩展控制文件](https://github.com/ergo70/tanimoto/blob/df31a86da4098c94c684167a789e34008d051e81/tanimoto.control)
- [扩展 SQL](https://github.com/ergo70/tanimoto/blob/df31a86da4098c94c684167a789e34008d051e81/tanimoto--1.0.sql)
- [C 实现](https://github.com/ergo70/tanimoto/blob/df31a86da4098c94c684167a789e34008d051e81/tanimoto.c)

`tanimoto` 1.0 计算两个二进制指纹的 Tanimoto 或 Jaccard 相似度：交集中置位比特数除以并集中置位比特数。优化后的 `tanimoto_c(bit varying, bit varying)` 用于等长、按字节对齐的化学或特征指纹。

### 设置与基础查询

安装 SQL 还会创建 `cdk.tanimoto(...)`，但既没有声明也没有创建 `cdk` 模式。创建扩展前要确保该模式存在。

```sql
CREATE SCHEMA IF NOT EXISTS cdk;
CREATE EXTENSION tanimoto;

SELECT tanimoto_c(
    B'11001010'::bit varying,
    B'10001110'::bit varying
);
```

当并集中至少有一个置位比特时，结果是 0 到 1 之间的 `real` 值：1 表示启用比特集合完全相同，0 表示没有共同启用比特。

### 函数与前置条件

- `tanimoto_c(bit varying, bit varying)` 使用 C 字节位计数查找表。长度不同或比特长度不能被 8 整除时会抛错。
- `cdk.tanimoto(bit varying, bit varying)` 是使用按位 `&`、`|` 和文本位计数的 SQL 实现。它更慢，并固定依赖 `cdk` 模式。
- 扩展 SQL 把两个函数都声明为严格、不可变、并行安全、无泄漏且使用调用者权限。

两个全零指纹的相似度分母为零，在数学上未定义；应排除这种情况，或定义明确的应用约定。不得悄悄比较由不同算法、长度、位序或特征字典生成的指纹。

### 打包与兼容性

虽然控制文件声明 `relocatable = true`，硬编码的 `cdk.tanimoto` 对象使扩展实际上依赖 `cdk` 模式。SQL 还使用 `PARALLEL SAFE` 和 `LEAKPROOF` 等现代函数属性，后者需要高权限安装。仓库没有 README、发行说明或 PostgreSQL 支持矩阵；使用前必须针对准确服务器编译，并运行有代表性的位模式测试。
