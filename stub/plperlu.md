

## Usage

> [plperlu: PL/Perl untrusted procedural language](https://www.postgresql.org/docs/current/plperl.html)

PL/Perl Untrusted provides full Perl capabilities including external module loading, filesystem access, and network operations. Only superusers can create functions in this language.

```sql
CREATE EXTENSION plperlu;

-- Use external CPAN modules
CREATE FUNCTION fetch_url(text) RETURNS text
LANGUAGE plperlu AS $$
  use LWP::Simple;
  my ($url) = @_;
  return get($url);
$$;

-- Read a file from the server filesystem
CREATE FUNCTION read_server_file(text) RETURNS text
LANGUAGE plperlu AS $$
  my ($path) = @_;
  open my $fh, '<', $path or die "Cannot open $path: $!";
  local $/;
  my $content = <$fh>;
  close $fh;
  return $content;
$$;

-- JSON processing with external module
CREATE FUNCTION parse_json_key(text, text) RETURNS text
LANGUAGE plperlu AS $$
  use JSON;
  my ($json_str, $key) = @_;
  my $data = decode_json($json_str);
  return $data->{$key};
$$;

-- Send email using Net::SMTP
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
