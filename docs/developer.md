# 開発環境の構築
開発は Docker 上で行う。

## Windows における Docker のセットアップ
以下はWindowsユーザー向けに、`WSL``Ubuntu22.04``Docker`の構築手順を説明する。

### 1. WSL2のセットアップ
#### WSL2のインストール
PowerShellを起動し、次のコマンドを実行します。
```bash
wsl --install
```

インストールが完了したら、WSLを最新バージョンにアップデートします。次のコマンドを実行してください。
```bash
wsl --update
```

#### 仮想マシンの機能を有効にする
```bash
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
```

#### WSLの既定のバージョンを 2 に設定する
次のコマンドを実行する。ただし、既にバージョンが2に設定されているなら不要かもしれない。
```bash
wsl --set-default-version 2
```

### 2. Ubuntu 20.04 のセットアップ
#### Ubuntu 20.04 のインストール
Microsoft Store から Ubuntu 20.04 (Linux ディストリビューション) をインストールする。著者は `Ubuntu 22.04.2 LTS` を用いた。
#### ユーザー名・パスワードの登録
Ubuntu 20.04 を初回起動してユーザー名とパスワードを設定する。ユーザー名とパスワードは `sudo` コマンドの実行時に必要となるので控えておく。
#### Ubuntu の更新
ここからはLinuxターミナル上での操作である。はじめに、Ubuntu 20.04 をアップデートする。
```bash
sudo apt update
```
```bash
sudo apt upgrade -y
```
Docker導入に必要なパッケージをインストールする。
```bash
sudo apt install curl -y
```
```bash
sudo apt install apt-transport-https -y
```

### 3. Docker Engine のセットアップ
#### Docker 公式の GPG 鍵を追加
```bash
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
```
#### 安定版の Docker レポジトリを追加
```bash
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```
#### レポジトリをアップデートし、 Docker Engine をインストール
```bash
sudo apt update
```
```bash
sudo apt install docker-ce docker-ce-cli containerd.io -y
```
#### docker-ompose もインストール
```bash
sudo apt install docker-compose -y
```
#### Docker デーモンを起動
```bash
sudo service docker start
```
起動することが確認出来たら、 `sudo service docker stop` で Docker を停止させて終了する。

## 著者
[connect0459](https://github.com/connect0459)
