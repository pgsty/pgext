## 用法

来源：

- [上游 README](https://github.com/palestamp/hamming_distance/blob/22c4c95bc1e8b92bfc32bc0c6024a1bacd5ce4aa/README.md)
- [扩展 control 文件](https://github.com/palestamp/hamming_distance/blob/22c4c95bc1e8b92bfc32bc0c6024a1bacd5ce4aa/hamming_distance.control)
- [SQL 安装脚本](https://github.com/palestamp/hamming_distance/blob/22c4c95bc1e8b92bfc32bc0c6024a1bacd5ce4aa/hamming_distance--0.0.1.sql)
- [C 实现](https://github.com/palestamp/hamming_distance/blob/22c4c95bc1e8b92bfc32bc0c6024a1bacd5ce4aa/hamming_distance.c)

`hamming_distance` `0.0.1` 版按位比较等长十六进制字符串，提供原始及归一化的距离与相似度函数。

### 示例

```sql
CREATE EXTENSION hamming_distance;
SELECT hamming_distance('1d3f', '1110');
SELECT hamming_similarity('1d3f', '1110');
SELECT hamming_distance_normalized('1d3f', '1110');
```

输入应预先验证为非空、等长、长度为偶数且匹配 `^[0-9A-Fa-f]+$` 的字符串。旧 C 解析器无法可靠拒绝所有非十六进制字符，也不会把奇数长度输入末尾的半字节计入结果。仓库没有当前 PostgreSQL 兼容性声明，因此依赖结果前应在隔离环境中测试具体构建和输入。
