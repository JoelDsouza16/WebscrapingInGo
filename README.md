# Webscraping In Go
A quick demo into web scraping in GoLang


Build in Go Programming Langauge
Used "Colly" to get the scraping work done on "internshala.com" by scraping all the internships available on the first page.

1. Clone the repository
2. CD into the directory
3. Run the below command. It would run web server. The application would be running on :8080
  
    > go run server.go

4. You can request information either through "Postman" or local "web browser"

      4.1 "/" 

      4.2 "/internshala"
5. Testing
    > go test 

6. Benchmarking
    > go test -bench=. 
