# Go言語の命名規則について

ファイル名 → snake_case（または単語1つ）
ディレクトリ名 → 小文字
パッケージ名 → 小文字
関数名・構造体名 → PascalCase または camelCase

### Javaとの比較

| 対象       | Java                  | Go                        |
| ---------- | --------------------- | ------------------------- |
| クラス名   | `UserService`         | `UserService`             |
| メソッド   | `getUser()`           | `GetUser()` / `getUser()` |
| ファイル   | `UserService.java`    | `user_service.go`         |
| パッケージ | `com.example.service` | `service`                 |

### Goでは先頭が大文字か小文字かで公開・非公開が決まる

HealthCheck → 他のパッケージから使用可能
healthCheck → 同じパッケージ内のみ

### 例

| 対象           | 例                  |
| -------------- | ------------------- |
| ファイル       | `health_handler.go` |
| パッケージ     | `handler`           |
| ディレクトリ   | `handler`           |
| 構造体         | `ScheduleService`   |
| 関数（公開）   | `HealthCheck`       |
| 関数（非公開） | `validateRequest`   |
| 定数           | `DefaultPageSize`   |
