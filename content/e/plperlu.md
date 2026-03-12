---
title: "plperlu"
linkTitle: "plperlu"
description: "PL/PerlU untrusted procedural language"
weight: 3270
categories: ["LANG"]
width: full
---

[**plperlu**](https://www.postgresql.org/docs/current/plperl.html) : PL/PerlU untrusted procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3270** | {{< badge content="plperlu" link="https://www.postgresql.org/docs/current/plperl.html" >}} | {{< ext "plperlu" >}} | `1.0` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plperlu" >}} |
|    **Need By**    | {{< ext "bool_plperlu" >}} {{< ext "hstore_plperlu" >}} {{< ext "jsonb_plperlu" >}} {{< ext "plperlu" >}} {{< ext "pg_utl_smtp" >}} {{< ext "sparql" >}} |
|   **See Also**    | {{< ext "plperl" >}} {{< ext "plluau" >}} {{< ext "pltclu" >}} {{< ext "bool_plperl" >}} {{< ext "hstore_plperl" >}} {{< ext "jsonb_plperl" >}} {{< ext "plpgsql" >}} {{< ext "plv8" >}} |
|    **Siblings**   | {{< ext "bool_plperlu" >}} {{< ext "jsonb_plperlu" >}} {{< ext "hstore_plperlu" >}} |


## Packages

| **PG18** | **PG17** | **PG16** | **PG15** | **PG14** |
|:--------:|:--------:|:--------:|:--------:|:--------:|
| {{< bg "1.0" "PostgreSQL 18: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 17: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 16: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 15: version 1.0" "green" >}} | {{< bg "1.0" "PostgreSQL 14: version 1.0" "green" >}} |

> [!Tip] This is a built-in contrib extension ship with the PostgreSQL kernel


## Install


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plperlu CASCADE; -- requires plperlu
```



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
