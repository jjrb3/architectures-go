package persistence

// TennisSet - Model
type TennisSet struct {
	ScoreFirstPlayer  int `bson:"score_first_player" json:"score_first_player"`
	ScoreSecondPlayer int `bson:"score_second_player" json:"score_second_player"`
}

// TennisMatch - Model
type TennisMatch struct {
	ID           uint        `bson:"id" json:"id"`
	FirstPlayer  string      `bson:"first_player" json:"first_player"`
	SecondPlayer string      `bson:"second_player" json:"second_player"`
	Winner       string      `bson:"winner" json:"winner"`
	Status       string      `bson:"status" json:"status"`
	SetList      []TennisSet `bson:"set_list" json:"set_list"`
}
