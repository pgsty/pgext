

## 用法

> [plperl: PL/Perl 受信过程语言](https://www.postgresql.org/docs/current/plperl.html)

PL/Perl 允许使用 Perl 编写 PostgreSQL 函数。作为受信语言，它在受限环境中运行，无法访问文件系统或外部模块。

```sql
CREATE EXTENSION plperl;

-- 简单标量函数
CREATE FUNCTION perl_hello(text) RETURNS text
LANGUAGE plperl AS $$
  my ($name) = @_;
  return "Hello, $name!";
$$;

SELECT perl_hello('world');

-- 使用 Perl 正则表达式处理文本
CREATE FUNCTION clean_whitespace(text) RETURNS text
LANGUAGE plperl AS $$
  my ($str) = @_;
  $str =~ s/^\s+|\s+$//g;   # 去除首尾空白
  $str =~ s/\s+/ /g;         # 合并内部空白
  return $str;
$$;

-- 返回复合类型
CREATE TYPE name_parts AS (first_name text, last_name text);

CREATE FUNCTION split_name(text) RETURNS name_parts
LANGUAGE plperl AS $$
  my ($full) = @_;
  my ($first, $last) = split(/\s+/, $full, 2);
  return { first_name => $first, last_name => $last };
$$;

-- 集合返回函数
CREATE FUNCTION perl_generate_series(integer, integer) RETURNS SETOF integer
LANGUAGE plperl AS $$
  my ($start, $stop) = @_;
  for my $i ($start .. $stop) {
    return_next($i);
  }
  return undef;
$$;

-- 触发器函数
CREATE FUNCTION perl_audit_trigger() RETURNS trigger
LANGUAGE plperl AS $$
  $_TD->{new}{modified_at} = localtime();
  return "MODIFY";
$$;
```

数据库访问使用 `spi_exec_query`：

```sql
CREATE FUNCTION perl_row_count(text) RETURNS integer
LANGUAGE plperl AS $$
  my ($table) = @_;
  my $rv = spi_exec_query("SELECT count(*) AS cnt FROM $table");
  return $rv->{rows}[0]{cnt};
$$;
```
