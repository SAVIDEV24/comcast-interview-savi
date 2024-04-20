Please make notes here to clarify any decisions taken that you wish to communicate.

Rate Limitor:
https://echo.labstack.com/docs/middleware/rate-limiter - Found echo provided this rate limiter functionality. It makes easy to add request per second per server. Made it 5 rps. Would like to learn more on this.


Caching : 
    Thought of adding new fields for most occurance character and it's count. And read it for "/stringinate" post request. Instead of runninng actual code logic.

Validator :
    Though it is binding , it is not working for invalid input query param. had to add sepaarte validator. Looking for better solution/ planning to learn thoroughly about echo bind.

Store:
    Design to keep both persistent and temporary in memory data store:
    -Introduced an interface with method sigantures Save and GET
    - Define structs for each ims and both will implement the same interface
    - This aids to swap the stores. We can use same interface to store inputs in database.
