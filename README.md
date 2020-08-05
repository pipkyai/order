Здесь описана программа, которая парсит информацию с криптовалютной биржи Huobi.pro по трем 
валютным парам eth/btc, btc/usdt и eth/usdt.

Основная идея: было предположение, что из-за резких скачков курсов криптовалют, на бирже может
сложиться ситуация, когда имея на балансе какое-то количество биткоинов, их можно продать за 
доллары, потом на эти доллары купить эфереум и затем снова за этот эфереум купить биткоин. И остаться в плюсе.

Реализация идеи:
1 - Представляем, что у нас есть один условный биткоин.
2 - Парсим информацию с биржи по трем валютным парам.
3 - Продаем Btc за Usdt, нас интересует "bid", по этой цене мы можем продать наш биткоин.
{
  "ch": "market.btcusdt.bbo",
  "ts": 1489474082831, //system update time
  "tick": {
    "symbol": "btcusdt",
    "quoteTime": "1489474082811",
    "bid": "10008.31",    
    "bidSize": "0.01",
    "ask": "10009.54",    
    "askSize": "0.3"
    "seqId": "1276823698734"
  }
}

4 - На полученные от предыдущей сделки USDT мы покупаем ETH, а потом за это ETH покупаем BTC.
5 - Если вложив один условный BTC и прогнав его через этот круг мы получим >1, то это хорошо.

PS Алгоритм, к сожалению, не рабочий, поскольку комиссия биржи делает стратегию не выгодной.
Но сами биржи, скорее всего, этот алгоритм и используют для извлечения прибыли и увеличения 
ликвидности. Иногда коэффициент >1 проскакивает.
