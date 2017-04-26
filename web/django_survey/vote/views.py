from django.shortcuts import render, HttpResponse
from django.template import RequestContext
from django.http import Http404

from .models import Survey, Log, Alternative


# Create your views here.

def home(request, id=None):

    p = {}

    p['list_survey_menu'] = Survey.objects.filter(status=True).order_by('-date')

    if request.POST:
        alternative = request.POST.get('alternative', None)
        survey      = request.POST.get('survey', None)
        if not alternative or not survey:
            p['error'] = 'Selecione uma das alternativas acima.'
        else:
            try:
                alt = Alternative.objects.get(id=int(alternative))
                alt.votes += 1
                alt.save()
                log = Log()
                log.survey = alt.survey
                log.alternative = alt
                if hasattr(request.session, 'session_key') and request.session.session_key:
                    log.user_session = request.session.session_key
                else:
                    log.user_session = 'anonymus'
                log.save()
                if not request.session.get('survey_itens', None):
                    request.session['survey_itens'] = [alt.survey.id]
                else:
                    request.session['survey_itens'] += [alt.survey.id]
            except Alternative.DoesNotExist:
                p['error'] = 'Houve um erro na hora de votar.'

    if id and int(id) > 0:
        try:
            p['survey'] = p['list_survey_menu'].get(id=id)
        except Survey.DoesNotExist:
            raise Http404
    elif p['list_survey_menu'].exists():
        p['survey'] = p['list_survey_menu'][0]
    else:
        raise Http404

    if request.session.get('survey_itens', None) and request.session.get('survey_itens', None).count(p['survey'].id) > 0:
        return render(request, 'vote/chart.html', p, content_type=RequestContext(request))
    return render(request, 'vote/home.html', p, content_type=RequestContext(request))
