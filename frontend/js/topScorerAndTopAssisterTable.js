const scorerSocket = new WebSocket("ws://localhost:8080/top-scorer-assister-table");
scorerSocket.onopen = function (event) {
    scorerSocket.send("top-scorer-assister-table");
    scorerSocket.onmessage =  (event)=> {
        let msg = JSON.parse(event.data);
        document.querySelectorAll("#TopScorerAndTopAssister tr").forEach( function (item)  {
            if (item.children[0].localName === "td") {
                let topScorer = item.children[0]
                let topAssister = item.children[1]
                topScorer.innerText = msg.TopScorer;
                topAssister.innerText = msg.TopAssister;
            }
        });
    };
};

