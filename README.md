# puncorp

A software platform to serve up delicious puns to the world. (Actually: a quick example/demo/hack to learn how Bazel works)

![Screenshot](screenshot.png)

Naturally, it's made up of many services in a monorepo:

* [pung](pung) - A binary that produces random puns
* [puntastic](puntastic) - A web interface that `pung` serves up puns in delicious Comic Neue

## Running

```bash
bazel run //puntastic
```

## Testing

```bash
bazel test //...
```

## Local development

```bash
cd puntastic
bazel build //pung
cp ../bazel-bin/pung/darwin_amd64_stripped/pung bin
./puntastic.sh
# Open localhost:8080
```

## TODO

- [x] Get `bazel run //pung` to work
- [ ] Get `bazel run //puntastic` to work