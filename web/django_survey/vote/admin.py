from django.contrib import admin
from .models import Survey, Alternative, Log

# Register your models here.

class AlternativeInline(admin.StackedInline):
    model = Alternative

class SurveyAdmin(admin.ModelAdmin):
    inlines = [AlternativeInline]
    list_display = ('title', 'status')

admin.site.register(Survey, SurveyAdmin)


class AlternativeAdmin(admin.ModelAdmin):
    pass

#admin.site.register(Alternative, AlternativeAdmin)


class LogAdmin(admin.ModelAdmin):
    pass

#admin.site.register(Log, LogAdmin)
