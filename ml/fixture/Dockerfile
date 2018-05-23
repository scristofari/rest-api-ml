FROM python:3

RUN apt-get update -y && apt-get upgrade -y

COPY . /src/app
WORKDIR /src/app

RUN pip install -r requirements.txt

CMD python3 algo.py -V /src/app