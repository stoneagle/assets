from django.conf.urls import url
from django.contrib import admin
from web.tushare import deal, bank


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
    url(r'^bank/getShiborData$', bank.getShiborData),
    url(r'^bank/getShiborMaData', bank.getShiborMaData),
    url(r'^bank/getShiborQuoteData$', bank.getShiborQuoteData),
    url(r'^bank/getLprData$', bank.getLprData),
    url(r'^bank/getLprMaData$', bank.getLprMaData),
]
