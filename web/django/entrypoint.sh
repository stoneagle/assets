#!/bin/bash

# echo "collect static files"
# python manage.py collectstatic --noinput

echo "apply database migrations"
python manage.py migrate

echo "start server"
# pip install -r requirements.txt && python manage.py runserver 0.0.0.0:8000
python manage.py runserver 0.0.0.0:8000
