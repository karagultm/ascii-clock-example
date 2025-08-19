# ⏰ ASCII Clock (Go)

Bu proje, **Go** dili ile yazılmış bir **terminal tabanlı dijital saat**
uygulamasıdır.\
ASCII karakterleri kullanarak saat, dakika ve saniyeyi gösterir.\
Ekrandaki ":" işareti her saniye yanıp sönerek gerçek bir dijital saat
efekti verir.

------------------------------------------------------------------------

## 🚀 Nasıl Çalıştırılır?

### Gereksinimler

-   [Go](https://go.dev/dl/) kurulu olmalı.
-   [inancgumus/screen](https://github.com/inancgumus/screen) paketi
    kullanılıyor.

### Adımlar

``` bash
# Depoyu klonla
git clone https://github.com/karagultm/ascii-clock-example.git
cd ascii-clock-example

# Gerekli bağımlılıkları indir
go mod tidy

# Programı çalıştır
go run main.go
```

------------------------------------------------------------------------

## 🎮 Özellikler

-   Terminal üzerinde **büyük ASCII rakamları** ile dijital saat
    görünümü.\
-   Saat, dakika ve saniye gösterimi.\
-   Yanıp sönen kolon (`:`) animasyonu.\
-   Sürekli güncellenen canlı saat.

------------------------------------------------------------------------

## 📷 Örnek Çalışma

    ███  █ █     ███  ███     ███  █  █     ███  ███
    █ █  █ █  ░  █    █ █  ░  █ █  █  █  ░  █    █ █
    █ █  █ █     ███  ███     ███  ████     ███  ███
    █ █   █   ░    █    █  ░    █     █  ░    █    █
    ███   █       ███  ███     ███     █     ███  ███

------------------------------------------------------------------------

## 🛠 Kullanılan Teknolojiler

-   Go programlama dili\
-   `inancgumus/screen` paketi (terminal ekranını temizlemek ve yeniden
    çizmek için)

------------------------------------------------------------------------

## 📌 Not

Bu uygulama tamamen **terminal tabanlıdır**. Herhangi bir GUI içermez.
