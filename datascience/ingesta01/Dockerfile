FROM python:3-slim

WORKDIR /programs/ingesta

COPY . .

RUN pip3 install mysql-connector-python
RUN pip3 install boto3

CMD ["python3", "ingesta01.py"]