FROM python:3.8-alpine

WORKDIR /app

# COPY requirements.txt requirements.txt
# RUN pip3 install click

COPY . .

CMD [ "python3", "cotu.py"]
