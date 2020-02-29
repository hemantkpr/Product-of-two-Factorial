# Product-of-two-Factorial
Hit this curl command to get Product of factorial of two numbers
curl --location --request POST 'http://localhost:8989/calculate' \
--header 'cache-control: no-cache' \
--header 'Content-Type: application/json' \
--header 'postman-token: 1f1a993d-f3f6-538f-a809-8b2507d28938' \
--data-raw '{
	"a" : 5,
	"b" : 5
}'
