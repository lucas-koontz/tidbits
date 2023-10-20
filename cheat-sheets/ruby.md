# Ruby

## Fixes

### Can't install ruby on M1 Mac "use of undeclared identifier 'RSA_SSLV23_PADDING'"

The root cause of this is because homebrew now [defaults openssl to openssl@3 instead of openssl@1.1](https://github.com/Homebrew/homebrew-core/commit/92d985af3e3757d045d1ce5f0db52ddf1e557108).

```sh
brew uninstall --ignore-dependencies openssl@3
brew install openssl@1.1
export CFLAGS="-DUSE_FFI_CLOSURE_ALLOC"
export PKG_CONFIG_PATH="$(brew --prefix openssl@1.1)/lib/pkgconfig"
rvm install <version> --with-openssl-dir=/opt/homebrew/opt/openssl@1.1
```
