# PC3R
1. Run go build to get the .exe file
2. Execute it
3. Since its on port 5000, run http://localhost:5000/players/Pepper to get a result
4. To send a post run `curl -X POST http://localhost:5000/players/Pepper` or `curl -X POST http://localhost:5000/players/User`
5. To run test files you can do `go test`

---
How to use go mod: `go mod init PC3R_Project/PC3R/Part3` <br>
`go mod init %pathtofile%`

---
Calls:
Player : `http://localhost:5000/players/Pepper`
League : `http://localhost:5000/league`

