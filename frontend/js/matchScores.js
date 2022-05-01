let matchId = 1
document.querySelectorAll("#MatchScoresTable tr").forEach( (item) => {
    if (item.children[0].localName === "td") {
        let home = item.children[0]
        let score = item.children[1]
        let visitor = item.children[2]
        let attackCount = item.children[3]
        let playerBasedStatistics = item.children[4]
        let matchResultsSocket = new WebSocket("ws://localhost:8080/match-result");
        matchResultsSocket.onopen = function (event) {
            matchResultsSocket.send(this);
            matchResultsSocket.onmessage =  (event)=> {
                let msg = JSON.parse(event.data);
                home.innerText = msg.Home.Name;
                score.innerText = msg.HomeScore + " - " + msg.VisitorScore;
                visitor.innerText = msg.Visitor.Name;
                attackCount.innerText = msg.Attacks.length;
                playerBasedStatistics.innerText = function (msg){
                    let sb = "-----------------HOME TEAM PLAYERS-----------------\n"
                    msg.Home.Players.forEach(item=>{
                        let twoPointSuccessRate = item.TwoPointAttempts === 0 ? "Not Attempted" : item.TwoPointSuccess*100/item.TwoPointAttempts+"%"
                        let threePointSuccessRate = item.ThreePointAttemts === 0 ? "Not Attempted" : item.ThreePointSuccess*100/item.ThreePointAttemts+"%"

                        let temp = item.Name + " ->"+
                        " Assist : " + item.MatchAssist +
                        ", Score : " + item.MatchScore +
                        ", TwoPointSuccessRate : " + twoPointSuccessRate +
                        ", ThreePointSuccessRate : " + threePointSuccessRate +
                        ", TurnoverCounter : " + item.TurnoverCounter + "\n"
                        sb += temp
                    })
                    sb += "----------------VISITOR TEAM PLAYERS----------------\n"
                    msg.Visitor.Players.forEach(item=>{
                        let twoPointSuccessRate = item.TwoPointAttempts === 0 ? "No Attempted" : item.TwoPointSuccess*100/item.TwoPointAttempts+"%"
                        let threePointSuccessRate = item.ThreePointAttemts === 0 ? "No Attempted" : item.ThreePointSuccess*100/item.ThreePointAttemts+"%"

                        let temp = item.Name + " ->"+
                            " Assist : " + item.MatchAssist +
                            ", Score : " + item.MatchScore +
                            ", TwoPointSuccessRate : " + twoPointSuccessRate +
                            ", ThreePointSuccessRate : " + threePointSuccessRate +
                            ", TurnoverCounter : " + item.TurnoverCounter + "\n"
                        sb += temp
                    })
                    return sb;
                }(msg)
            };
        }.bind(matchId);
        matchId++
    }
});