Distribute application to multiple platform by curl
====

brew? apt? yum? No, curl!

```
$ L=/usr/local/bin/app && curl -sL -A "`uname -sp`" https://locahost:3000/app.gz | zcat >$L && chmod +x $L
```

## Reference

- [heroku/hk](https://github.com/heroku/hk)
- [flynn/flynn-cli-dist](https://github.com/flynn/flynn-cli-dist)
