**Prerequisites**
Before running this project, ensure you have the following installed:
Docker test

Docker (Installed & Running)
Git (To Clone the Repository)

**Clone the Repository**
git clone https://github.com/niteshjoshi17/food-delivery-api.git
cd food-delivery-api

**Build & Run using Docker**
docker build -t food-delivery-api .
docker run -p 8080:8080 food-delivery-api

**Verify the API is Running**
curl http://localhost:8080
