## 手順

1. .netrcを生成します
   ```
   echo "machine github.com" > ~/.netrc \
   && echo "login <GitHubのユーザー名>" >> ~/.netrc \
   && echo "password <GitHubのアクセストークン>" >> ~/.netrc
   ```

2. GOPRIVATEを登録します
   ```
   export GOPRIVATE=github.com,direct
   ```

3. モジュールを更新します
   ```
   go mod tidy -e
   ```

4. 下記、コマンドでmain.goを実行します。
   ```
   go run main.go
   ```