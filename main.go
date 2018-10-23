package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Sayfa adında bir şema oluşturup kullanacağımız değişkenleri burada tanımlıyoruz.
type Page struct {
	Title string
	Static string
}

func main() {

	// Oluşturduğumuz şemadan bir obje oluşturalım. OOP mantığı.
	page := Page{"Can", "static"}

	// ParseFiles ile templatimizi ilgili html sayfası olarak tanımlıyoruz.
	// Must burada hata yakalamak için yardımcı konumda.
	templates := template.Must(template.ParseFiles("templates/main.html"))

	// Statik doyalarımız için yaygın olarak kullanılan statik dosya dizinini kullanıyoruz.
	// http.Handle ile sayfa içinde çağırılan /static/ isteklerini statik dizinine yönlendiriyoruz.
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	// '/' yani Anasayfa pathi çağırıldığıda işlenecek adımları belirtir.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Render işlemi. Sayfamız ile sayfa içinde kullandığımız değişkenleri birleştiriyoruz.
		err := templates.ExecuteTemplate(w, "main.html", page);

		// Hata kontrolü.
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("Dinlenen port 8080");
	fmt.Println(http.ListenAndServe(":8080", nil));
}
