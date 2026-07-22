## 用法

来源：

- [官方 README](https://github.com/marconeperes/pgBRTypes/blob/5491ecb720e6dab34ba4211509c2b75ca9f05c29/README.md)
- [官方扩展 SQL](https://github.com/marconeperes/pgBRTypes/blob/5491ecb720e6dab34ba4211509c2b75ca9f05c29/sql/pgbrtypes--1.0.sql)
- [官方 GUC 实现](https://github.com/marconeperes/pgBRTypes/blob/5491ecb720e6dab34ba4211509c2b75ca9f05c29/pgBRTypes.c)

`pgbrtypes` 为巴西个人和公司纳税登记号增加了经过校验的 `cpf` 与 `cnpj` 类型。它会规范化带格式或不带格式的输入，在类型值中保留前导零，并可按常用掩码显示。

### 核心流程

```sql
CREATE EXTENSION pgbrtypes;

CREATE TABLE counterparties (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  person_cpf cpf,
  company_cnpj cnpj
);

INSERT INTO counterparties (person_cpf, company_cnpj)
VALUES ('021.516.186-68', '14.320.135/0001-49');

CREATE INDEX counterparties_cpf_idx
ON counterparties (person_cpf);

SELECT person_cpf, company_cnpj
FROM counterparties;
```

非法证件号会被类型输入函数拒绝。前导零有意义时应使用文本输入；虽然提供隐式整数转换，但整数字面量无法表示原始文本宽度。

### 类型与辅助函数

- `cpf` 和 `cnpj` 是带二进制发送/接收函数的固定八字节类型。
- `is_cpf(text)` 与 `is_cnpj(text)` 可在不存储的情况下校验候选值；重载还接受 bigint 和 C-string 输入。
- 比较操作符以及 `cpf_op_class` 与 `cnpj_op_class` B-tree 操作符类支持等值/范围谓词和索引。
- 隐式转换接受 integer/bigint 输入，赋值转换则把类型值转回 bigint。

校验只检查标识符的语法/校验位规则；它不能证明个人或公司真实存在或拥有该号码。

### 输出格式

以下用户级 GUC 只改变显示，不改变存储标识或比较行为：

```sql
SHOW cpf.enable_format;
SHOW cnpj.enable_format;

SET cpf.enable_format = off;
SET cnpj.enable_format = off;
```

由于文本输出随会话设置变化，集成程序应选择稳定设置，或有意转换为 bigint。CPF/CNPJ 属于敏感标识符，应独立于格式化实施访问控制、日志脱敏与保留策略。

### 兼容性

README 要求 PostgreSQL 9.1 或更高版本，但只记录了 9.6 测试；核验的源码最后修改于 2017 年，且没有文档化 Windows 支持。生产使用前应在确切 PostgreSQL/OS 目标上构建并测试索引、二进制备份恢复、升级、非法输入和不受区域设置影响的输出。
