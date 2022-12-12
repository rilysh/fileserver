### File Server
A simple file hosting server

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
