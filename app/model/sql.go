package model

import (
    "context"
    "gorm.io/gorm"
)

// GormAdapterは、データベース操作を抽象化するためのアダプターです。
// GORMライブラリを使用して、具体的なデータベース操作を実装します。
type GormAdapter struct {
    DB *gorm.DB  // DBは、GORMライブラリによるデータベース接続を持ちます。
}

func (ga *GormAdapter) Store(ctx context.Context, table string, entity interface{}) error {
    return ga.DB.Table(table).Create(entity).Error
}

// todo 共通化却下 sql.goがいらない
func (ga *GormAdapter) FindOne(ctx context.Context, table string, id string, result interface{}) error {
    return ga.DB.Table(table).Where("id = ?", id).
    Preload("FishCategory").
    Preload("FishingMethods").
	First(result).Error
}

func (ga *GormAdapter) FindAll(ctx context.Context, table string, query interface{}, result interface{}) error {
    return ga.DB.Table(table).Where(query).Find(result).Error
}

// SQLActionは、データベースアクションのためのインターフェースです。
// これにより、データベースへのアクセスを抽象化し、異なるデータベース技術への依存を減らします。
// DBの依存していいだろ リリースしてからmysqlからpostgresに変えるとか聞いたことねーよ
type DBMethods interface {
    Store(context.Context, string, interface{}) error
	FindOne(context.Context, string, string, interface{}) error
    FindAll(context.Context, string, interface{}, interface{}) error
}

// Sessionは、トランザクション管理のためのインターフェースです。
// これにより、データベース操作をトランザクションスコープ内で安全に実行できます。
type Session interface {
    // WithTransactionは、指定された関数をトランザクション内で実行します。
    // トランザクションは、関数の実行が成功した場合はコミットされ、エラーが発生した場合はロールバックされます。
    WithTransaction(context.Context, func(context.Context) error) error
    
    // EndSessionは、セッションを終了し、使用していたリソースを解放します。
    EndSession(context.Context)
}
