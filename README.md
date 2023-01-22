### File Server
A simple and temporary file hosting server

Important: Before compiling, please make sure you've changed the default authentic key in `key.go`.

- Build
```sh
git clone https://github.com/kiwimoe/fileserver
# Visit https://go.dev/doc/install and download according your platform, if you haven't
cd fileserver && go build
./fileserver
# Upload files through curl
curl -F'file=@myfile' http://mysite.web/upload
```
- Usage
```sh
# Arguments are optional here
./fileserver -d [url] -p [port]
```
By default server should be listen on `http://localhost:1337`
