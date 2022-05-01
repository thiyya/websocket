package frontend

func HomePage() string {
	x := "<!DOCTYPE html> <html>	<head>	<style>table {  font-family: arial, sans-serif;  border-collapse: collapse; width: 100%;}td, th {  border: 1px solid #dddddd; text-align: left; padding: 8px;}tr:nth-child(even) {  background-color: #dddddd;}</style>" +
		"</head>	<body>" +
		"<h2>Match Scores</h2>" +
		"<table class = \"table\" id=\"MatchScoresTable\">" +
		"<thead>" +
		"<tr>" +
		"<th>Home</th>" +
		"<th>Score</th>" +
		"<th>Visitor</th>" +
		"<th>Attack Count</th>" +
		"<th>Player Based Statistics</th>" +
		"</tr>" +
		"</thead>" +
		"<tr>" +
		"<td class = \"home\"></td>" +
		"<td class = \"score\"></td>" +
		"<td class = \"visitor\"></td>" +
		"<td class = \"attackCount\"></td>" +
		"<td class = \"playerBasedStatistics\"></td>" +
		"</tr>" +
		"<tr>" +
		"<td class = \"home\"></td>" +
		"<td class = \"score\"></td>" +
		"<td class = \"visitor\"></td>" +
		"<td class = \"attackCount\"></td>" +
		"<td class = \"playerBasedStatistics\"></td>" +
		"</tr>" +
		"<tr>" +
		"<td class = \"home\"></td>" +
		"<td class = \"score\"></td>" +
		"<td class = \"visitor\"></td>" +
		"<td class = \"attackCount\"></td>" +
		"<td class = \"playerBasedStatistics\"></td>" +
		"</tr>" +
		"</table>" +
		"</table>" +
		"<h2>Top Players</h2>" +
		"<table class = \"table\" id=\"TopScorerAndTopAssister\">" +
		"<thead>" +
		"<tr>" +
		"<th>TopScorer</th>" +
		"<th>TopAssister</th>" +
		"</tr>" +
		"</thead>" +
		"<tr>" +
		"<td class =\"team\"></td>" +
		"<td class =\"point\"></td>" +
		"</tr>" +
		"</table>" +
		"<h2>Leauge Table</h2>" +
		"<table class = \"table\" id=\"LeaugeTable\">" +
		"<thead>" +
		"<tr>" +
		"<th>Team</th>" +
		"<th>Point</th>" +
		"</tr>" +
		"</thead>" +
		"<tr>" +
		"<td class =\"team\"></td>" +
		"<td class =\"point\"></td>" +
		"</tr>" +
		"<tr>" +
		"<td class =\"team\"></td>" +
		"<td class =\"point\"></td>" +
		"</tr>" +
		"<tr>" +
		"<td class =\"team\"></td>" +
		"<td class =\"point\"></td>" +
		"</tr>" +
		"<tr>" +
		"<td class =\"team\"></td>" +
		"<td class =\"point\"></td>" +
		"</tr>" +
		"<tr>" +
		"<td class =\"team\"></td>" +
		"<td class =\"point\"></td>" +
		"</tr>" +
		"<tr>" +
		"<td class =\"team\"></td>" +
		"<td class =\"point\"></td>" +
		"</tr>" +
		"<script src = \"../js/matchScores.js\"></script>" +
		"<script src = \"../js/leaugeTable.js\"></script>" +
		"<script src = \"../js/topScorerAndTopAssisterTable.js\"></script>" +
		"</body>" +
		"</html>"
	return x
}
