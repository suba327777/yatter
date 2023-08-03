package object

type Relationship struct {
	Follower_id int64 `db:"follower_id"`
	Followed_id int64 `db:"followed_id"`
}
