package handlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
	"websocket/constants"
	"websocket/db"
	"websocket/models"
)

func MatchResultHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("MatchResult Connection error : ", err)
		return
	}
	log.Println("MatchResult Connection established.")
	defer conn.Close()

	for {
		ticker := time.NewTicker(constants.RefreshTime * time.Second)
		tickerForAttack := time.NewTicker(1 * time.Second)
		done := make(chan bool)
		doneForAttack := make(chan bool)

		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("MatchResult Could not read message from websocket, error : ", err)
			break
		}

		matchResult := startGame(string(p))
		log.Println("The game is started now !")

		// Ilk acıldıgında boş gorunmesin diye initial halini gönderiyor.
		databytes, err := json.Marshal(matchResult)
		if err = conn.WriteMessage(messageType, databytes); err != nil {
			log.Println("MatchResult Could not write message to websocket, error", err)
			return
		}

		// 1 ya da 2 sn de bir atak olusturulur. Yani 240 sn lik bir oyunda 120 ile 240 arasında ataklar oluşacaktır.
		// Bu atak ile maçın gidişatı ile alakalı değişiklikler yapılır.
		go func() {
			for {
				select {
				case <-doneForAttack:
					return
				case <-tickerForAttack.C:
					updateGame(matchResult)
				}
			}
		}()

		// Her refreshTime sn de bir maçın durumunu gönderir.
		go func() {
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					databytes, err = json.Marshal(matchResult)
					if err = conn.WriteMessage(messageType, databytes); err != nil {
						log.Println("MatchResult Could not write message to websocket, error", err)
						return
					}
					log.Printf("%v - %v : %v - %v \n", matchResult.Home.Name, matchResult.Visitor.Name, matchResult.HomeScore, matchResult.VisitorScore)
				}
			}
		}()

		time.Sleep(constants.GameTime * time.Second)
		ticker.Stop()
		tickerForAttack.Stop()
		done <- true
		doneForAttack <- true
		log.Println("MatchResult Ticker stopped")

		// Maç bittikten sonra update edilmiş hali son kez gönderilir ve kapatılır.
		databytes, err = json.Marshal(matchResult)
		if err = conn.WriteMessage(messageType, databytes); err != nil {
			log.Println("MatchResult Could not write message to websocket, error", err)
			return
		}
		log.Printf("Maç Sonu : %v - %v : %v - %v \n", matchResult.Home.Name, matchResult.Visitor.Name, matchResult.HomeScore, matchResult.VisitorScore)

		// Maç bittikten sonra lig tablosu update edilir.
		updateLeaugeTable(matchResult)

		// Maç bittikten sonra maç içerisinde olmus butun ataklar loglanır.
		for i := 0; i < len(matchResult.Attacks); i++ {
			log.Printf("%v. atak : \n %v takımının atağı : \n Assister : %+v\n Scorer : %+v\n Score : %+v\n IsTurnover : %+v\n IsSuccessful : %+v\n", i+1, matchResult.Attacks[i].Team.Name, matchResult.Attacks[i].Assister, matchResult.Attacks[i].Scorer, matchResult.Attacks[i].Score, matchResult.Attacks[i].IsTurnover, matchResult.Attacks[i].IsSuccessful)
		}

		break
	}
}

func startGame(matchId string) *models.MatchResult {
	var matchDetails models.Fixture
	fixture := db.GetFixture()
	for _, match := range fixture {
		if match.MatchId == matchId {
			matchDetails = match
		}
	}
	matchResult := &models.MatchResult{
		MatchId:      matchId,
		Home:         matchDetails.Home,
		HomeScore:    0,
		Visitor:      matchDetails.Visitor,
		VisitorScore: 0,
		Attacks:      []*models.Attack{},
	}
	players := db.GetPlayers()

	for _, player := range players {
		player.ThreePointAttemts = 0
		player.ThreePointSuccess = 0
		player.TwoPointAttempts = 0
		player.TwoPointSuccess = 0
		player.TurnoverCounter = 0
		player.MatchAssist = 0
		player.MatchScore = 0
	}
	return matchResult
}

func updateGame(matchResult *models.MatchResult) *models.MatchResult {
	// Bir oyun en fazla 24 sn surebilir. 24 sn probleme gore 2 sn ye karsılık gelir. O yuzden 1. ya da 2. sn lerde o takımın oyunu bitebilir.
	possibleAttackTime := []int{1, 2}
	attackTime := time.Duration(possibleAttackTime[rand.Intn(len(possibleAttackTime))])
	time.Sleep(attackTime * time.Second)

	// ilk atağı Misafir takıma sonrasında da sırayla vermeyi tercih ettim. İstenirse bunun kararı da random verilebilir.
	// Örn reboundu atak yapan takım kazanırsa tekrar atak yapabilirler gibi
	attackingTeam := matchResult.Home
	if len(matchResult.Attacks)%2 == 0 {
		attackingTeam = matchResult.Visitor
	}

	// Atakların sonucu 0 - 2 - 3 puanlık olabilir.
	// 0 top kaybıdır.
	// 2 ya da 3 puanlık girişimler yapıldıgında sonucu olumlu ya da olumsuz olabılır.
	// Bunu oyuncuların 2-3 puanlık yuzdelerını bulabılmek ıcın yaptım.
	possibleScoreOfAttack := []int{0, 2, 3}
	possibleAttackSuccess := []bool{true, false}
	score := possibleScoreOfAttack[rand.Intn(len(possibleScoreOfAttack))]
	isSuccessful := possibleAttackSuccess[rand.Intn(len(possibleAttackSuccess))]
	assister := attackingTeam.Players[rand.Intn(len(attackingTeam.Players))]
	scorer := attackingTeam.Players[rand.Intn(len(attackingTeam.Players))]
	isTurnover := false
	if score == 0 { // top kaybı
		isTurnover = true
		assister = nil
		scorer.TurnoverCounter++
	} else if isSuccessful { // basarılı atak
		assister.TotalAssist++
		assister.MatchAssist++
		scorer.TotalScore += score
		scorer.MatchScore += score
		if score == 2 {
			scorer.TwoPointAttempts++
			scorer.TwoPointSuccess++
		}
		if score == 3 {
			scorer.ThreePointAttemts++
			scorer.ThreePointSuccess++
		}
	} else { // basarısız atak
		if score == 2 {
			scorer.TwoPointAttempts++
		}
		if score == 3 {
			scorer.ThreePointAttemts++
		}
	}

	attack := models.Attack{
		MatchId:      matchResult.MatchId,
		Team:         attackingTeam,
		Assister:     assister,
		Scorer:       scorer,
		Score:        score,
		IsTurnover:   isTurnover,
		IsSuccessful: isSuccessful,
	}

	if score > 0 && isSuccessful {
		if len(matchResult.Attacks)%2 == 0 {
			matchResult.VisitorScore += score
		} else {
			matchResult.HomeScore += score
		}
	}
	matchResult.Attacks = append(matchResult.Attacks, &attack)
	return matchResult
}

func updateLeaugeTable(result *models.MatchResult) {
	// beraberlikleri ev sahibine yazdım.
	if result.HomeScore >= result.VisitorScore {
		result.Home.Point += 2
		result.Visitor.Point += 1
	} else {
		result.Visitor.Point += 2
		result.Home.Point += 1
	}
}
