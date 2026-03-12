


## 用法

> [omni_auth: 基础会话管理](https://docs.omnigres.org/omni_auth/basics/)

`omni_auth` 扩展提供构建认证系统的基本原语，目前支持基于密码的认证。

### 认证主体

认证主体代表认证的目标——通常是用户，但也涵盖未识别的标识符（如登录名或电子邮件），以支持对不存在账户的认证尝试跟踪以及第三方 OAuth 认证即注册场景。

### 核心函数

**`omni_auth.authenticate(Authenticator, authentication_subject_id)`**

基于 `Authenticator` 类型分发，接受认证器和主体 ID。返回实现了 `omni_auth.successful_authentication` 接口的 `Authentication` 类型值。

**`omni_auth.successful_authentication(Authentication)`**

返回布尔值，表示认证是否成功。

### 支持的实现

- 密码认证（通过密码 `Authenticator` 类型分发）
