package db

import "websocket/models"

func GetFixture() (fixture []models.Fixture) {
	fixture = append(fixture, models.Fixture{
		MatchId: "1",
		Home:    &Lakers,
		Visitor: &Celtics,
	})
	fixture = append(fixture, models.Fixture{
		MatchId: "2",
		Home:    &Bulls,
		Visitor: &Jazz,
	})
	fixture = append(fixture, models.Fixture{
		MatchId: "3",
		Home:    &Suns,
		Visitor: &Warriors,
	})

	fixture = append(fixture, models.Fixture{
		MatchId: "4",
		Home:    &Jazz,
		Visitor: &Lakers,
	})
	fixture = append(fixture, models.Fixture{
		MatchId: "5",
		Home:    &Celtics,
		Visitor: &Suns,
	})
	fixture = append(fixture, models.Fixture{
		MatchId: "6",
		Home:    &Warriors,
		Visitor: &Bulls,
	})

	fixture = append(fixture, models.Fixture{
		MatchId: "7",
		Home:    &Jazz,
		Visitor: &Celtics,
	})
	fixture = append(fixture, models.Fixture{
		MatchId: "8",
		Home:    &Suns,
		Visitor: &Lakers,
	})
	fixture = append(fixture, models.Fixture{
		MatchId: "9",
		Home:    &Bulls,
		Visitor: &Warriors,
	})
	return
}
