from django.conf.urls import url
from django.contrib import admin
from web.tushare import deal


urlpatterns = [
    url(r'^admin/', admin.site.urls),
    url(r'^deal/getHistData$', deal.getHistData),
    url(r'^deal/getKData$', deal.getKData),
    url(r'^deal/getTodayAll$', deal.getTodayAll),
    url(r'^deal/getTickData$', deal.getTickData),
    url(r'^deal/getRealtimeQuotes$', deal.getRealtimeQuotes),
    url(r'^deal/getTodayTicks$', deal.getTodayTicks),
    url(r'^deal/getIndex$', deal.getIndex),
    url(r'^deal/getSinaDD$', deal.getSinaDD),
]
