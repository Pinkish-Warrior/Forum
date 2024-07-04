FROM golang:1.22.5
# RUN mkdir/app
# ADD . /app
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 8000
CMD [ "/app/main" ]

# docker image build -t forum .