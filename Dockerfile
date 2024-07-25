FROM golang:1.23rc2
# RUN mkdir/app
# ADD . /app
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 8000
CMD [ "/app/main" ]

# docker image build -t forum .