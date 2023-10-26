FROM golang:latest
# RUN mkdir/app
# ADD . /app
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 8005
CMD [ "/app/main" ]

# docker image build -t forum .