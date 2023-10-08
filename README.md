# blog_server
This is the backend for a decoupled personal blog project, with the frontend provided by a collaborator. **This project is built upon gin, gorm, MySQL, Redis, and Elasticsearch.**

バックエンドとフロントエンドが分離された個人のブログプロジェクトです。フロントエンドは協力者によって提供されます。**このプロジェクトは、gin、gorm、MySQL、Redis、およびElasticsearchを基に構築されています。**

## Dependency
- go
- gin
- gorm
- mysql
- redis
- elasticsearch
- jwt-go
- aws-sdk-go
- websocket
- crypto
- logrus

## Main features
### English
- creation, read, deletion, update for advertisement, article, news, images, menu, comment
- User bind email, change password, registration/login(jwt + cookie session)
- Online chat/groupchat(websocket)
- Article bookmark, liking, data statistic, full-text search by title and category
- Third-party login (reviewing by the third party now)
- Counting the number of comments, articles, article likes, comment likes, and tracking user logins over the past 7 days.(redis+elasticsearch)

### Japanese
- 広告、記事、ニュース、画像、メニュー、コメントの作成、読み取り、削除、更新
- ユーザーのメールアドレスを紐づけ、パスワードを変更、登録/ログイン（(jwt + cookie session)
- オンラインチャット/グループチャット（WebSocket）
- インログのログ
- 文章のブックマーク、いいね、データ統計、タイトルとカテゴリーによる全文検索
- 第三者ログイン（第三者による審査中）
- コメント数、記事数、記事のいいね数、コメントのいいね数、過去7日間のユーザーログイン数の統計。（Redis+Elasticsearch）

## Project Structure
```
blog
├── api
│   ├── advertise_api
│   ├── article_api
│   ├── chat_api
│   ├── comment_api
│   ├── image_api
│   ├── log_api
│   ├── menu_api
│   ├── message_api
│   ├── news_api
│   ├── setting_api
│   ├── statistic_api
│   ├── tag_api
│   └── user_api
├── config
├── docs
├── flag
├── global
├── initialization
├── middleware
├── models
│   ├── ctype
│   ├── res
├── plugins
│   ├── aws
│   ├── email
│   ├── log_stash
│   └── qq
├── preview
├── routers
├── service
│   ├── common_service
│   ├── enter.go
│   ├── es_service
│   ├── image_service
│   ├── redis_service
│   ├── synchro_service
│   └── user_service
├── test
│   ├── 1.flag
│   ├── 10.randomNameAvatar
│   ├── 11.log
│   ├── 12.getIpLocation
│   ├── 13.get_internal_ips
│   ├── 2.redis
│   ├── 3.email
│   ├── 4.elasticSearch
│   ├── 5.markdownToHtml
│   ├── 7.sychroFullTextIndexToES
│   ├── 8.recursionComment
│   ├── 9.news
│   └── ６.fullTextSearch
├── uploads
│   ├── avatar
│   ├── chat_avatar
│   ├── file
│   └── system
└── utils
    ├── jwts
    ├── mask
    ├── pwd
    └── verification
```

## Frontend(coding)
Vue3

## Deployment
```go
set GOARCH=amd64
set GOOS=linux
go build -o main

mysqldump -uroot -proot blog_db > blog_db.sql
go run main.go -es -dump article_index
go run main.go -es -dump full_text_index


docs
uploads
main
settings.yaml
article_index.json
full_text_index.json
blog_db.sql
```





