package exercise

type Exercise struct {
	ExerciseID  string `gorm:"primaryKey;default:gen_random_uuid()"`
	UserID      string
	ImageID     string
	GifID       string
	VideoID     string
	Name        string
	Description string
	MuscleGroup string
	Difficulty  string
	Type        string
	SetsCount   int32
	RepsCount   int32
	Duration    int64
}
