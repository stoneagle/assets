from django.conf.urls import url
from django.contrib import admin
from web.tushare import deal, bank, billboard, classify


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

    url(r'^billboard/getBrokerTops$', billboard.getBrokerTops),
    url(r'^billboard/getCapTop', billboard.getCapTop),
    url(r'^billboard/getInstDetail$', billboard.getInstDetail),
    url(r'^billboard/getInstTop$', billboard.getInstTop),
    url(r'^billboard/getTopList$', billboard.getTopList),

    url(r'^classify/getIndustryClassified$', classify.getIndustryClassified),
    url(r'^classify/getConceptClassified$', classify.getConceptClassified),
    url(r'^classify/getAreaClassified$', classify.getAreaClassified),
    url(r'^classify/getSMEClassified$', classify.getSMEClassified),
    url(r'^classify/getGEMClassified$', classify.getGEMClassified),
    url(r'^classify/getSTClassified$', classify.getSTClassified),
    url(r'^classify/getHS300s$', classify.getHS300s),
    url(r'^classify/getSZ50s$', classify.getSZ50s),
    url(r'^classify/getZZ500s$', classify.getZZ500s),
    url(r'^classify/getTerminated$', classify.getTerminated),
    url(r'^classify/getSuspended$', classify.getSuspended),
]
