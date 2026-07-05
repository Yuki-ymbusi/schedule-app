### ブランチ戦略

それぞれのブランチの役割を以下に記す

- main

完成版を置くブランチ。

- develop

普段の開発を行うブランチ。

- feature/〇〇

開発する物ごとに作成するブランチ。

feature/create-schedule
feature/update-schedule
feature/delete-schedule

基本的にfeatureブランチで作成したものをdevelopにマージし完成形をmainにマージする。

### コミットメッセージについて

#### 以下の方法で分類する

- feat: 新機能
- fix: バグ修正
- refactor: リファクタリング
- style: フォーマットのみの変更
- docs: ドキュメント
- chore: 環境構築・設定変更
