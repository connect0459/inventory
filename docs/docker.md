# Dockerコンテナに関して
## docker-compose.yml
`container_name` は `service` 以下に記述したコンテナに名前を付けるもので、指定しない場合は一意な値がランダムで割り振られる。以下コマンドで確認すると、指定した場合は `NAME` にその文字列が格納され、指定しない場合は `3e7...1ea_web-api_1` のような文字列となっていることがわかる。
```bash
docker ps -a
```
### 指定する場合
docker-compose.ymlでDockerを立ち上げる、つまり`docker-compose` コマンドやVSCodeの `Reopen in Container` を使用したときに以下のようなエラーが生じる。
```bash
ERROR: for web-api  Cannot create container for service web-api: Conflict. The container name "/web-api" is already in use by container "bd6...3d3". You have to remove (or rename) that container to be able to reuse that name.
```
同名のコンテナを立ち上げることはできない。この場合、docker-compose.ymlで `container_name` を別のものにするか、すでに構築している場合は停止しているコンテナを `docker start [コンテナ名]` で起動し、そこにアタッチすることもできる。

### 指定しない場合
`docker-compose` コマンドを毎回使用していると、コンテナが乱立することになる。そのため、使わないコンテナは `docker rm [コンテナ名]` で適宜削除する。

