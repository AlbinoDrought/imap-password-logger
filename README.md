# IMAP Password Logger

<a href="https://hub.docker.com/r/albinodrought/imap-password-logger">
  <img alt="albinodrought/imap-password-logger Docker Pulls" src="https://img.shields.io/docker/pulls/albinodrought/imap-password-logger">
</a>
<a href="https://github.com/AlbinoDrought/imap-password-logger/blob/master/LICENSE">
  <img alt="AGPL-3.0 License" src="https://img.shields.io/github/license/AlbinoDrought/imap-password-logger">
</a>

I forgot my email password. The [email client on my phone](https://github.com/k9mail/k-9) wouldn't [export my password](https://github.com/k9mail/k-9/blob/db1dcb46189cd4e62a84ee56c585aec586fa6d6a/app/core/src/main/java/com/fsck/k9/preferences/SettingsExporter.java#L227), but it would let me change the IMAP server. I changed it to this.

See [emersion/go-imap](https://github.com/emersion/go-imap)

## Running

1. Start up the IMAP server (runs on port 1143):

```sh
docker run --rm -it -p 1143:1143 albinodrought/imap-password-logger
```

2. Point your mail client at `<your ip>:1143` with all TLS/SSL features disabled.

3. You should see something like this:

```
2019/12/22 18:00:35 Starting IMAP server at localhost:1143
2019/12/22 18:00:50 albinodrought@example.com hunter2
2019/12/22 18:00:50 albinodrought@example.com hunter2
```

## Building

### Without Docker

```
go get -d -v
go build
```

### With Docker

`docker build -t albinodrought/imap-password-logger .`
