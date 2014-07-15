This is sample web application to enable you install binary which is prepared for multiple platform with **same URL** and **same command** (using curl).

```
$ L=/usr/local/bin/app && curl -sL -A "`uname -sp`" http://locahost:3000/app.gz | zcat >$L && chmod +x $L
```

[http://localhost:3000]() will attempt to detect your OS and CPU architecture based on the User-Agent, then redirect you to the latest release for your platform.


## Usage

Run application:

```
$ go run web.go
```

## Reference

- [heroku/hk](https://github.com/heroku/hk)
- [flynn/flynn-cli-dist](https://github.com/flynn/flynn-cli-dist)
