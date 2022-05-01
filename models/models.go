package models

type Player struct {
	Name              string
	TotalAssist       int
	TotalScore        int
	MatchAssist       int
	MatchScore        int
	TwoPointAttempts  int
	ThreePointAttemts int
	TwoPointSuccess   int
	ThreePointSuccess int
	TurnoverCounter   int
}

type Attack struct {
	MatchId      string
	Team         *Team
	Assister     *Player
	Scorer       *Player
	Score        int
	IsTurnover   bool
	IsSuccessful bool
}

type Team struct {
	Name    string
	Point   int
	Players []*Player
}

type MatchResult struct {
	MatchId      string
	Home         *Team
	HomeScore    int
	Visitor      *Team
	VisitorScore int
	Attacks      []*Attack
}

type Fixture struct {
	MatchId string
	Home    *Team
	Visitor *Team
}

type LeaugeTable struct {
	Team  *Team
	Point int
}
type TopPlayer struct {
	TopScorer   string
	TopAssister string
}
