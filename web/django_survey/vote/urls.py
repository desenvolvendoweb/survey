from django.conf.urls import url
from django.contrib import admin

from . import views

urlpatterns = [
    url(r'^(?P<id>[0-9]+)/', views.home, name='home'),
    url(r'^$', views.home, name='home'),
]
