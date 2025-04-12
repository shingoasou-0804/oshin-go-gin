# Oshin Go Gin

Ginを使用したシンプルなGo Webアプリケーション

## 必要条件

- Go 1.21以上
- Docker

## セットアップ
1. リポジトリのクローン
```bash
git clone https://github.com/shingoasou-0804/oshin-go-gin.git
cd oshin-go-gin
```

2. 依存関係のインストール
```bash
go mod tidy
```

### Dockerコンテナ開発環境

1. イメージのビルド
```bash
docker compose build --no-cache
```

2. コンテナの実行
```bash
docker compose up
```

## 開発

### ライブリロード

Airを使用したライブリロード開発環境を利用できます：

### エンドポイント

- `GET /ping`: サンプルエンドポイント
  - レスポンス: `{"message": "pong"}`

## プロジェクト構造

```
.
├── main.go          # メインアプリケーション
├── .air.toml        # Air設定ファイル
├── go.mod           # Go依存関係
├── go.sum           # Go依存関係のチェックサム
└── README.md        # このファイル
```

## ライセンス

MIT
