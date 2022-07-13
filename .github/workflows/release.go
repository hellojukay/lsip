name: release

on:
  push:
    tags: '*'

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - uses: actions/setup-go@v2
      with:
        stable: 'true'
        go-version: '^1.16.1'
    - name: build
      run: |
        export PATH=$PATH:$(go env GOPATH)/bin
        go install github.com/mitchellh/gox@latest
        mkdir dist
        perl build.pl
        ls -al dist/
    - name: Upload binaries to release
      uses: svenstaro/upload-release-action@v2
      with:
        # secrets.GITHUB_TOKEN 是内置变量
        repo_token: ${{ secrets.GITHUB_TOKEN }}
        file: dist/*
        asset_name: mything
        tag: ${{ github.ref }}
        overwrite: true
        file_glob: true