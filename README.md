# gcal-week

Google カレンダーの今週の予定をターミナルに表示する CLI ツール。

```
Week: 2026/03/23 (Mon) – 03/29 (Sun)
────────────────────────────────────────
Mon  03/23  (予定なし)
Tue  03/24  10:00 チームMTG
             14:00 1on1
Wed  03/25  TODO 終日イベント
Thu  03/26  (予定なし)
Fri  03/27  09:30 週次レビュー
Sat  03/28  (予定なし)
Sun  03/29  (予定なし)
────────────────────────────────────────
```

今日の行はハイライト表示されます。カラーテーマは [Everforest Dark](https://github.com/sainnhe/everforest) ベース。

## セットアップ

### 1. Google Cloud で認証情報を取得

1. [Google Cloud Console](https://console.cloud.google.com/) でプロジェクトを作成
2. **APIとサービス > ライブラリ** から「Google Calendar API」を有効化
3. **APIとサービス > 認証情報** で OAuth 2.0 クライアント ID を作成（種類: デスクトップアプリ）
4. `credentials.json` をダウンロードして以下に配置:

```
~/.config/gcal-week/credentials.json
```

### 2. ビルド & 実行

```bash
git clone https://github.com/mako-agawa/gcal-week
cd gcal-week
go build -o gcal-week .
./gcal-week
```

初回実行時にブラウザが開き、Google アカウントの認証を求められます。認証後、トークンは `~/.config/gcal-week/token.json` に保存されます。

### (オプション) PATH に追加

```bash
mv gcal-week /usr/local/bin/
```

## 要件

- Go 1.21+
- Google アカウント + Google Calendar API の有効化
