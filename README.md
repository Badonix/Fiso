<div style="display:flex; align-items: center">
  <h1 style="position:relative; top: -6px">Fiso</h1>
</div>

---

Golang useful applications collection
#



### Installation

1\. First of all you need to clone this repo from github:

```sh
git clone https://github.com/Badonix/fiso
```

2\. Next step requires you to run `go mod tidy` in order to install all required packages.

```sh
go mod tidy
```

3\. Build the app

```sh
go build -o fiso
```

4\. Move the executable to path

```sh
sudo mv shorten /usr/local/bin
```

### Usage

1\. Currently only tool is shorten, which shortes the url:

```sh
fiso shorten --url https://example.com --name randomname
```
`--url` is required
`--name` is optional

It returns shortened url hopefully
