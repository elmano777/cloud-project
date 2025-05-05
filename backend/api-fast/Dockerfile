FROM python:3.11-slim

WORKDIR /app
ENV MONGO_URL=mongodb://mongo-dev:27017

COPY requirements.txt /app
RUN pip install --no-cache-dir --upgrade -r requirements.txt

COPY . /app

CMD [ "uvicorn", "main:app", "--host", "0.0.0.0", "--port", "80"]
