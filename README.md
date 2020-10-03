# ap2-merihariko-backend
<img src="https://user-images.githubusercontent.com/49857129/94939566-4aa56880-050d-11eb-9bbc-a5866a8fcd36.jpg" width="90%"><img src="https://user-images.githubusercontent.com/49857129/94939566-4aa56880-050d-11eb-9bbc-a5866a8fcd36.jpg" width="90%">

## プロジェクト概要
- めりはりこのWebAPIのリポジトリです

## 開発のルールやタスク管理
### コード規約
- [Effective Go](https://golang.org/doc/effective_go.html)に準拠

### プルリクエストについて
- プルリクエストを出す際は適宜,右にある枠から
- Reviewers (絶対に目を通して欲しい人がいる場合指定する), Labels, Linked issues を設定してください

### ブランチ名について
- master
  - プロダクトとしてリリースするためのブランチ. 基本触らない
- develop
  - 開発ブランチ． コードが安定し,リリース準備ができたら master へマージする. リリース前はこのブランチが最新バージョンとなる.
- feature
  - 機能の追加. develop から分岐し, develop にマージする.
  - feature-#{対応するissue}-{任意で詳細}
- fix
  - 現在のプロダクトのバージョンに対する変更・修正用.
  - fix-#{対応するissue}-{任意で詳細}

### コミットメッセージについて
- add:新機能
- fix:バグ修正
- wip:作業中（WIP：Work In Progress）
- clean:整理（削除も含む）

- 例`[add]ログイン機能を実装`

### タスクの管理について
GitHubProjectを使用します
- To do
    - 作成されたissueをおく場所
- In progress
    - Pull requestを出し， レビューを待っている状態のissueが置いてある場所
- Done
    - 完了し, developブランチにmergeされたissueが置いてある場所