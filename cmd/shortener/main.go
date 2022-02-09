package main

import (
    "bufio"
    "bytes"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "os"
    "strconv"
    "strings"
)


func UrlHandler(w http.ResponseWriter, r *http.Request) {

  var resp *http.Response

    switch r.Method {
    // если метод POST
    case "POST":
      // извлекаем фрагмент url= из URL запроса ?url=someUrl
      // q := r.URL.Query()..Get("url")
      /*
      if q == "" {
          http.Error(w, "The quering URL is missing", http.StatusBadRequest)
          return
      }*/
      // но можно и в теле
      longUrl, err := io.ReadAll(r.Body)
      if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
      }
      // меняем URL u возвращаем ответ как менять пока не знаю поэтому тут заглушка
      tinyUrl := "http://tinyurl.com/y6F4wwh"
      w.Header().Set("Content-Type", "text/plain; charset=utf-8")
      w.WriteHeader(201)
      fmt.Fprintln(w, tinyUrl)
      return

      // если метод GET
    case "GET":
      tinyUrl, err := io.ReadAll(r.Body)
      if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
      }
      // меняем URL u возвращаем ответ как менять пока не знаю поэтому тут заглушка
      longUrl := "http://longurl.com/sitefolder"
      w.WriteHeader(307)
      //fmt.Fprintln(w, longUrl)
      resp.Header().Set("Location", longUrl)
      return

    default:
      http.Error(w, "400", http.StatusBadRequest)
      return
    }
}


func main() {
  // маршрутизация запросов обработчику
http.HandleFunc("/", UrlHandler)
// запуск сервера с адресом localhost, порт 8080
http.ListenAndServe(":8080", nil)

}
