const leaugeTableSocket = new WebSocket("ws://localhost:8080/leauge-table");
leaugeTableSocket.onopen = function (event) {
    leaugeTableSocket.send("leauge-table");
    leaugeTableSocket.onmessage =  (event)=> {
        let msg = JSON.parse(event.data);
        msg.sort((a,b) => b.Point - a.Point);

        let index = 0
        document.querySelectorAll("#LeaugeTable tr").forEach( function (item)  {
            if (item.children[0].localName === "td") {
                let team = item.children[0]
                let point = item.children[1]
                team.innerText = msg[index].Name;
                point.innerText = msg[index].Point;
                index++
            }
        });
    };
};

