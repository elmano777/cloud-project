FROM python:3-slim

WORKDIR /programs/ingesta

COPY . .

RUN pip3 install pymongo
RUN pip3 install boto3
RUN pip3 install pandas

CMD ["python3", "ingesta03.py"]