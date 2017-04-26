from django.db import models
from django.db.models import Q

# Create your models here.

class Survey(models.Model):
    title = models.CharField(max_length=255, verbose_name='Título')
    status = models.BooleanField(default=True, verbose_name='Habilitado')
    date = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return self.title

    def alternatives(self):
        return Alternative.objects.filter(Q(survey=self) & Q(status=True))

    class Meta:
        verbose_name = 'Enquete'


class Alternative(models.Model):
    survey = models.ForeignKey(Survey, verbose_name='Enquete')
    title = models.CharField(max_length=255, verbose_name='Título')
    votes = models.IntegerField(verbose_name='Votos', default=0, editable=False)
    status = models.BooleanField(default=True, verbose_name='Habilitado')
    date = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return self.title

    class Meta:
        verbose_name = 'Alternativa'


class Log(models.Model):
    user_session = models.CharField(max_length=255, verbose_name='Usuário')
    survey = models.ForeignKey(Survey, verbose_name='Enquete')
    alternative = models.ForeignKey(Alternative, verbose_name='Alternativa')
    date = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return self.survey.title

    class Meta:
        verbose_name = 'Log'
