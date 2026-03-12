

## 用法

> [plperlu: PL/Perl 不受信过程语言](https://www.postgresql.org/docs/current/plperl.html)

PL/Perl 不受信版本提供完整的 Perl 功能，包括加载外部模块、访问文件系统和网络操作。只有超级用户可以使用此语言创建函数。

```sql
CREATE EXTENSION plperlu;

-- 使用外部 CPAN 模块
CREATE FUNCTION fetch_url(text) RETURNS text
LANGUAGE plperlu AS $$
  use LWP::Simple;
  my ($url) = @_;
  return get($url);
$$;

-- 从服务器文件系统读取文件
CREATE FUNCTION read_server_file(text) RETURNS text
LANGUAGE plperlu AS $$
  my ($path) = @_;
  open my $fh, '<', $path or die "Cannot open $path: $!";
  local $/;
  my $content = <$fh>;
  close $fh;
  return $content;
$$;

-- 使用外部模块处理 JSON
CREATE FUNCTION parse_json_key(text, text) RETURNS text
LANGUAGE plperlu AS $$
  use JSON;
  my ($json_str, $key) = @_;
  my $data = decode_json($json_str);
  return $data->{$key};
$$;

-- 使用 Net::SMTP 发送邮件
CREATE FUNCTION send_email(text, text, text) RETURNS boolean
LANGUAGE plperlu AS $$
  use Net::SMTP;
  my ($to, $subject, $body) = @_;
  my $smtp = Net::SMTP->new('localhost');
  $smtp->mail('postgres@localhost');
  $smtp->to($to);
  $smtp->data();
  $smtp->datasend("Subject: $subject\n\n$body");
  $smtp->dataend();
  $smtp->quit();
  return 1;
$$;
```
