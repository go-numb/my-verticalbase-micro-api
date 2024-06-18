package page

import "time"

type Page struct {
	ID          string `firestore:"id,omitempty" csv:"id,omitempty" json:"id,omitempty"`
	Title       string `firestore:"title,omitempty" csv:"title,omitempty" json:"title,omitempty"`
	Slug        string `firestore:"slug,omitempty" csv:"slug,omitempty" json:"slug,omitempty"`
	Description string `firestore:"description,omitempty" csv:"description,omitempty" json:"description,omitempty"`

	// 詳細を map[string]any で持つことで、カスタムな情報を持たせることができる
	// 入れ子にすることも可能、ただし、入れ子にする場合は、構造体を定義して、それを使うことを推奨
	Details map[string]any `firestore:"details,omitempty" csv:"details,omitempty" json:"details,omitempty"`

	Categories []string `firestore:"categories,omitempty" csv:"categories,omitempty" json:"categories,omitempty"`
	Tags       []string `firestore:"tags,omitempty" csv:"tags,omitempty" json:"tags,omitempty"`

	// soft delete 用のフィールド
	Delete bool `firestore:"delete,omitempty" csv:"delete,omitempty" json:"delete,omitempty"`

	UpdatedAt time.Time `firestore:"updated_at,omitempty" csv:"updated_at,omitempty" json:"updated_at,omitempty"`
	CreatedAt time.Time `firestore:"created_at,omitempty" csv:"created_at,omitempty" json:"created_at,omitempty"`
}

type PageOTC struct {
	ID          string `firestore:"id,omitempty" csv:"id,omitempty" json:"id,omitempty"`
	Title       string `firestore:"title,omitempty" csv:"title,omitempty" json:"title,omitempty"`
	Slug        string `firestore:"slug,omitempty" csv:"slug,omitempty" json:"slug,omitempty"`
	Description string `firestore:"description,omitempty" csv:"description,omitempty" json:"description,omitempty"`

	// 詳細を map[string]any で持つことで、カスタムな情報を持たせることができる
	// 入れ子にすることも可能、ただし、入れ子にする場合は、構造体を定義して、それを使うことを推奨
	Details string `firestore:"details,omitempty" csv:"details,omitempty" json:"details,omitempty"`

	Categories string `firestore:"categories,omitempty" csv:"categories,omitempty" json:"categories,omitempty"`
	Tags       string `firestore:"tags,omitempty" csv:"tags,omitempty" json:"tags,omitempty"`

	// soft delete 用のフィールド
	Delete bool `firestore:"delete,omitempty" csv:"delete,omitempty" json:"delete,omitempty"`

	UpdatedAt time.Time `firestore:"updated_at,omitempty" csv:"updated_at,omitempty" json:"updated_at,omitempty"`
	CreatedAt time.Time `firestore:"created_at,omitempty" csv:"created_at,omitempty" json:"created_at,omitempty"`
}
