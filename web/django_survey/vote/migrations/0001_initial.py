# -*- coding: utf-8 -*-
# Generated by Django 1.11 on 2017-04-19 19:28
from __future__ import unicode_literals

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Alternative',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('title', models.CharField(max_length=255, verbose_name='Título')),
                ('votes', models.IntegerField(default=0, editable=False, verbose_name='Votos')),
                ('date', models.DateTimeField(auto_now_add=True)),
            ],
            options={
                'verbose_name': 'Alternativa',
            },
        ),
        migrations.CreateModel(
            name='Log',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('user_session', models.CharField(max_length=255, verbose_name='Usuário')),
                ('date', models.DateTimeField(auto_now_add=True)),
                ('alternative', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='vote.Alternative', verbose_name='Alternativa')),
            ],
            options={
                'verbose_name': 'Log',
            },
        ),
        migrations.CreateModel(
            name='Survey',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('title', models.CharField(max_length=255, verbose_name='Título')),
                ('date', models.DateTimeField(auto_now_add=True)),
            ],
            options={
                'verbose_name': 'Enquete',
            },
        ),
        migrations.AddField(
            model_name='log',
            name='survey',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='vote.Survey', verbose_name='Enquete'),
        ),
        migrations.AddField(
            model_name='alternative',
            name='survey',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='vote.Survey', verbose_name='Enquete'),
        ),
    ]
