# Setup

```shell
go mod init github.com/jllovet/learning-go/07-cli/000-urfave/002-config-options
go get github.com/urfave/cli/v3@latest
go get github.com/urfave/cli-altsrc
go get github.com/urfave/cli-altsrc/json
```

# Building and Testing

```shell
go build -o greet
```

```shell
./greet hello --config config-fr.json
```
