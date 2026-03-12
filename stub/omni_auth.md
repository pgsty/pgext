


## Usage

> [omni_auth: Basic session management](https://docs.omnigres.org/omni_auth/basics/)

The `omni_auth` extension provides primitives for building authentication systems. It currently supports password-based authentication.

### Authentication Subject

An authentication subject represents the target of authentication -- typically a user, but also covers unrecognized identifiers (such as login or e-mail) to support tracking authentication attempts against non-existent accounts and third-party OAuth authentication-as-signup scenarios.

### Core Functions

**`omni_auth.authenticate(Authenticator, authentication_subject_id)`**

Dispatched over `Authenticator` types, accepts an authenticator and a subject ID. Returns a value of type `Authentication` that implements `omni_auth.successful_authentication`.

**`omni_auth.successful_authentication(Authentication)`**

Returns a boolean indicating whether the authentication was successful.

### Supported Implementations

- Password authentication (dispatched via the password `Authenticator` type)
