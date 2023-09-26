package inform

// 兼职信息
type Information struct {
	ID          int
	UserEmail   string
	Topic       string
	Type        string
	Pay         string
	Times       string
	Address     string
	FullAddress string
	Description string
}

// 储存兼职关系
type Userinfor struct {
	Useremail string
	Toid      string
	Name      string
	Photo     string
	Age       string
	Gender    string
}
