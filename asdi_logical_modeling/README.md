# ASDI: Logical Modeling

## Introduction
這個專案是用來說明我在 Google DSC in NTHU 所推出了課程 <Abstract System Design & Implementation> 中的第一階段實作示範，主要用 Graph base 的方式來說明 ER Model 並且搭配 Dgraph 的圖形資料庫來檢視成果。

## Example Project: Simulate UberEats
針對 UberEats 的模型設計一個 Graph base schema，下圖的列表簡單說明了大致的架構。

* User
    * Consumer
    * Driver
    * Producer
* Product
    * Consumer
* Order
    * Consumer
    * Producer
    * Driver
    * Product

## Practice
* 嘗試擴充這個圖，例如：讓 Driver 和 Consumer 可以聊天
* 嘗試設計一個服務的 Graph

## Author
Dagin Wu / Tachun Wu <daginwu@gmail.com>

