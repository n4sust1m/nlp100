FROM perl:5.32.1-slim

RUN mkdir /app
WORKDIR /app

COPY ./cpanfile /app/
COPY ./cpanfile.snapshot /app/

RUN cpanm Carton
RUN carton install
