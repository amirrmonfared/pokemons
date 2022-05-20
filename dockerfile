FROM alpine:latest

RUN mkdir /app

COPY pokemonApp /app

CMD [ "/app/pokemonApp" ]