# nim-itb-scraper
## What is this?
This is a web scraper. Scraping ITB student's data.
It use Golang with concurrency implementation so it's **so fast**.
It does 100 HTTP request in concurrency.
It's my first time using golang concurrency so i think this code didn't optimze enough.
My point to do this project because i want to try implement concurrency.
And it's so satisfying because it so fast :)
Feel free to use it.
## HOW TO USE
1. Install Golang
2. Clone this repo
```bash
git clone https://github.com/aeramu/nim-itb-scraper
```
3. Open https://ditsti.itb.ac.id/nic/manajemen_akun/pengecekan_user and login with your ITB SSO Login
4. Open storage inspector on your browser, and see cookies section. Copy *ci_session* value
5. Open cmd/main.go in your text editor, and replace variable ci_session with value in step 4
```go
//paste your cookie here, (19-09-2020) ci_session and ITBnic)
var session = "REPLACE_THIS_WITH_YOUR_COPIED_VALUE"
```
6. Open terminal/cmd run command in project directory
```bash
go run ./cmd
```
