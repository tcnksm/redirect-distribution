This is sample web application to enable you install binary which is prepared for multiple platform with **same URL** and **same command** (using curl).

```
$ L=/usr/local/bin/sample-app && curl -sL -A "`uname -sp`" http://localhost:3000/app.zip | zcat >$L && chmod +x $L
```

[http://localhost:3000]() will attempt to detect your OS and CPU architecture based on the User-Agent, then redirect you to the latest release for your platform.

## Environmental variable

This app is supposed to use for binary hosted on Github release page like [https://github.com/tcnksm/sample-multiple-platform-bin/releases/download/v1.0](https://github.com/tcnksm/sample-multiple-platform-bin/releases/download/v1.0)

To run this application you need to set environmental variable belows:

- **BASE_URL** - Base url where binary is hosted e.g., https://github.com/tcnksm/sample-multiple-platform-bin/releases/download
- **DIST_NAME** - Distribution name of binary e.g., hk
- **VERSION** - Version of binary e.g., v1.0

## Reference

- [heroku/hk](https://github.com/heroku/hk)
- [flynn/flynn-cli-dist](https://github.com/flynn/flynn-cli-dist)
