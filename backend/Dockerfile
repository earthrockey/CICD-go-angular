FROM golang

WORKDIR /go/src/github.com/earthrockey/becon-backend
COPY . .
RUN go get github.com/jinzhu/gorm
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/line/line-bot-sdk-go/linebot
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]