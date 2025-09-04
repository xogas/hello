FROM python:3.6.12-slim-buster
USER root

COPY ./ /app/
WORKDIR /app/

RUN pip install -r requirements.txt

CMD ["python", "manager runserver 0.0.0.0:8000"]
