# URL Shortener Service
The service provides users with the ability to create short links. When navigating through a short link created in the service, the browser redirects to the original address.

## Starting the Application
1. Start the Spring Boot application with the command `./gradlew bootRun`
2. Send a request to shorten a link using the curl command-line utility. The request body is contained in the file [create-req.json](./create-req.json). 
Execute the command: `curl -X POST -d "@create-req.json" -H "Content-Type: application/json" http://localhost:8080/make_shorter`
The response will be a JSON similar to:
    ```
    {
        "shortUrl":"http://localhost:8080/hnh",
        "secretKey": "abcfjiedis"
    }
    ```
    Where `shortUrl` is the address that, when visited, will redirect the request to the original address, and `secretKey` is the key for administrative operations.

3. Open the short URL `http://localhost:8080/hnh` in your browser, and the request will be redirected to the original URL from [create-req.json](./create-req.json).

4. Send a `GET` HTTP request with the previously obtained `secretKey` to the address `http://localhost:8080/admin/abcfjiedis`. The response will include usage statistics for the short URL in JSON format similar to:
    ```
    {
        "creationDate": "2022-06-26T22:06:11.876+00:00",
        "usageCount": 1
    }
    ```
5. Send a `DELETE` HTTP request to `http://localhost:8080/admin/abcfjiedis` to delete the previously created redirection rule.