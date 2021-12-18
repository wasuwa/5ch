## 概要
Go で作成した Sample API

## アーキテクチャ
- DDD
- レイアードアーキテクチャ

## 技術構成
- Go
- Echo
- Gorm
- PostgreSQL

## API
### Users
|      | メソッド | URI |
| ---- | ---- | ---- |
| 一覧表示 | GET | /users |
| 個別表示 | GET | /users/:id |
| 作成 | POST | /users |
| 更新 | PATCH | /users/:id |
| 削除 | DELETE | /users/:id |
