To run the image build and run with following commands but docker-compose setup is preferable: 

docker build -t events-service .

docker run -p 8080:8080 events-service
