package building

//Building интерфейс зданий
type Building interface {
	GetWall() string
	GetDoor() string
}
