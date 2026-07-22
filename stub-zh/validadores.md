## 用法

来源：

- [官方上游文档](https://github.com/guedes/validadores/blob/63ea247a692f6995c54a24e80f1a39802ede59b6/README.md)
- [官方扩展 SQL](https://github.com/guedes/validadores/blob/63ea247a692f6995c54a24e80f1a39802ede59b6/sql/validadores.sql)
- [官方回归示例](https://github.com/guedes/validadores/blob/63ea247a692f6995c54a24e80f1a39802ede59b6/test/sql/base.sql)

`validadores` 0.0.2 是纯 SQL 的巴西 CPF 校验位辅助扩展。尽管较宽泛的葡萄牙语项目描述提到 CNPJ，已复核发布只安装 CPF 函数、一元操作符和数值域；其中没有 CNPJ 验证器。

### 验证并约束 CPF 值

创建扩展，调用函数或后缀操作符；只有其准确接受规则符合应用时才使用该域：

```sql
CREATE EXTENSION validadores;

SELECT cpf_valido(59328253241);
SELECT 59328253241 #? AS cpf_is_valid;

CREATE TABLE pessoa (
  nro_cpf cpf
);
```

`cpf_valido(numeric)` 将数值文本左补齐到 11 位，计算两个校验位并返回布尔值。后缀 `#?` 操作符调用同一函数。`cpf` 域是带有 `CHECK (cpf_valido(VALUE))` 约束的 `numeric`。

### 数据模型注意事项

CPF 是标识符而不是数量。数值列会丢弃标点，并可能在应用边界掩盖前导零处理。如果准确输入表示很重要，应存储规范文本，并应用经过明确审查的验证函数或域，而不是依赖隐式数值转换。

源码只实现自己的校验位算术。它没有显式拒绝通常视为无效 CPF 的全相同数字序列，没有验证归属或签发，没有验证格式化文本，也没有规范化 CNPJ。因此，通过该函数不能证明标识符真实，或符合当前业务与监管规则。应添加应用专用的重复数字和规范长度测试，与权威巴西验证规范比对，并在使用域前测试已知有效与无效样例。项目较旧；PostgreSQL 升级前应重新检查行为。
