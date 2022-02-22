FROM python:3.8-alpine

MAINTAINER sntshkmr60@gmail.com

WORKDIR /app

COPY requirements.txt requirements.txt

RUN pip3 install -r requirements.txt

COPY . .

CMD [ "python3", "cotu.py"]
