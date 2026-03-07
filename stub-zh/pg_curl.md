

## 用法


```sql
CREATE EXTENSION pg_curl;
```

执行 HTTP Get 请求：

```sql
-- 封装 curl http get 请求
CREATE OR REPLACE FUNCTION get(url TEXT) RETURNS TEXT LANGUAGE SQL AS $BODY$
WITH s AS (SELECT
               curl_easy_reset(),
               curl_easy_setopt_url(url),
               curl_easy_perform(),
               curl_easy_getinfo_data_in()
) SELECT convert_from(curl_easy_getinfo_data_in, 'utf-8') FROM s;
$BODY$;


SELECT get('https://www.postgresql.org/');
```


通过 SMTP 发送邮件：

```bash
CREATE OR REPLACE FUNCTION email(url TEXT, username TEXT, password TEXT, subject TEXT, sender TEXT, recipient TEXT, body TEXT, type TEXT) RETURNS TEXT LANGUAGE SQL AS $BODY$
    WITH s AS (SELECT
        curl_easy_reset(),
        curl_easy_setopt_mail_from(sender),
        curl_easy_setopt_password(password),
        curl_easy_setopt_url(url),
        curl_easy_setopt_username(username),
        curl_header_append('From', sender),
        curl_header_append('Subject', subject),
        curl_header_append('To', recipient),
        curl_mime_data(body, type:=type),
        curl_recipient_append(recipient),
        curl_easy_perform(),
        curl_easy_getinfo_header_in()
    ) SELECT curl_easy_getinfo_header_in FROM s;
$BODY$;
```

通过 FTP 下载文件：

```sql
CREATE OR REPLACE FUNCTION download(url TEXT, username TEXT, password TEXT) RETURNS BYTEA LANGUAGE SQL AS $BODY$
    WITH s AS (SELECT
        curl_easy_reset(),
        curl_easy_setopt_password(password),
        curl_easy_setopt_url(url),
        curl_easy_setopt_username(username),
        curl_easy_perform(),
        curl_easy_getinfo_data_in()
    ) SELECT curl_easy_getinfo_data_in FROM s;
$BODY$;
```
