from django.conf.urls import url
from django.contrib import admin
from web.tushare import deal


urlpatterns = [
    url(r'^admin/', admin.site.urls),
    url(r'^getHistData$', deal.getHistData),
]
