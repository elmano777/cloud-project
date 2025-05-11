FROM python:3-slim

WORKDIR /programs/ingesta

COPY . .

RUN pip3 install psycopg2
RUN pip3 install boto3

CMD ["python3", "ingesta02.py"]