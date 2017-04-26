#!/bin/bash

echo "Collectstatic executing ..."
/usr/local/bin/python manage.py collectstatic --noinput --clear
echo "Collectstatic finished !"

echo "Migrate executing ..."
/usr/local/bin/python manage.py migrate
echo "Migrate finished !"

echo "Create Super User ..."
echo "from django.contrib.auth.models import User; User.objects.create_superuser('admin', 'admin@example.com', 'admin')" | ./manage.py shell
echo "Create Super User finished !"

echo "Service gunicorn executing ... "
/usr/local/bin/gunicorn django_survey.wsgi:application -w 2 -b :8000
