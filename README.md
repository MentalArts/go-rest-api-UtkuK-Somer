# Book Library Management API

## 📌 Proje Açıklaması
Bu proje, **Go (Golang)** ve **Gin-Gonic** framework'ü kullanılarak geliştirilmiş bir **RESTful API**'dir.
PostgreSQL ile veritabanı entegrasyonu yapılmış ve Docker kullanılarak container ortamında çalıştırılabilir hale getirilmiştir.

API, bir **kitap yönetim sistemi** sunar ve aşağıdaki varlıkları içerir:
- **Yazarlar (Authors)**
- **Kitaplar (Books)**
- **İncelemeler (Reviews)**

---

## 🚀 Kullanılan Teknolojiler
- **Golang** (1.23+)
- **Gin-Gonic** (API Framework)
- **GORM** (ORM Kütüphanesi)
- **PostgreSQL** (Veritabanı)
- **Docker & Docker Compose**
- **Swagger** (API Dökümantasyonu)

---

## 🔧 Kurulum Adımları
### **1️⃣ Docker ile Çalıştırma (Önerilen)**
Projeyi Docker Compose ile hızlı bir şekilde ayağa kaldırabilirsiniz.

```sh
# Repository'yi klonla
git clone https://github.com/MentalArts/go-rest-api-name-surname.git
cd go-rest-api-name-surname

# Docker Compose ile çalıştır
docker-compose up -d --build
```

📌 **Çalıştırdıktan sonra API, `http://localhost:8000` adresinde aktif olacaktır.**

### **2️⃣ Manuel Kurulum (Geliştirici Modu)**
Eğer Docker kullanmadan çalıştırmak isterseniz:

```sh
# Repository'yi klonla
git clone https://github.com/MentalArts/go-rest-api-name-surname.git
cd go-rest-api-name-surname

# PostgreSQL çalıştır
sudo service postgresql start  # Linux
docker run --name pgdb -e POSTGRES_PASSWORD=1234 -p 5432:5432 -d postgres  # Alternatif Docker

# Go bağımlılıklarını yükle
go mod tidy

# API'yi başlat
go run main.go
```

---

## 📌 API Endpointleri

### **1️⃣ Yazar İşlemleri (Authors)**
| HTTP | Endpoint | Açıklama |
|------|---------|----------|
| `GET` | `/api/v1/authors` | Tüm yazarları getir |
| `GET` | `/api/v1/authors/:id` | Belirtilen yazarı getir |
| `POST` | `/api/v1/authors` | Yeni yazar ekle |
| `PUT` | `/api/v1/authors/:id` | Yazarı güncelle |
| `DELETE` | `/api/v1/authors/:id` | Yazarı sil (Soft Delete) |

### **2️⃣ Kitap İşlemleri (Books)**
| HTTP | Endpoint | Açıklama |
|------|---------|----------|
| `GET` | `/api/v1/books` | Tüm kitapları getir |
| `GET` | `/api/v1/books/:id` | Belirtilen kitabı getir |
| `POST` | `/api/v1/books` | Yeni kitap ekle |
| `PUT` | `/api/v1/books/:id` | Kitabı güncelle |
| `DELETE` | `/api/v1/books/:id` | Kitabı sil (Soft Delete) |

### **3️⃣ İnceleme İşlemleri (Reviews)**
| HTTP | Endpoint | Açıklama |
|------|---------|----------|
| `GET` | `/api/v1/books/:id/reviews` | Belirli bir kitaba ait tüm incelemeleri getir |
| `POST` | `/api/v1/books/:id/reviews` | Yeni inceleme ekle |
| `PUT` | `/api/v1/reviews/:id` | İncelemeyi güncelle |
| `DELETE` | `/api/v1/reviews/:id` | İncelemeyi sil |

---

## 🔍 API Örnek Kullanım
### **1️⃣ Yeni Bir Yazar Ekleme (`POST /api/v1/authors`)
```json
{
  "name": "J.K. Rowling",
  "biography": "British writer, best known for Harry Potter series.",
  "birth_date": "1965-07-31T00:00:00Z"
}
```

### **2️⃣ Yeni Bir Kitap Ekleme (`POST /api/v1/books`)
```json
{
  "title": "Harry Potter and the Philosopher's Stone",
  "author_id": 1,
  "isbn": "9780747532743",
  "publication_year": 1997,
  "description": "First book in the Harry Potter series."
}
```

### **3️⃣ Bir Kitaba İnceleme Ekleyerek (`POST /api/v1/books/:id/reviews`)
```json
{
  "rating": 5,
  "comment": "This book was amazing!",
  "date_posted": "2025-03-09T13:47:49.718232Z"
}
```

---

## 🎯 Katkıda Bulunma
Bu proje açık kaynaklıdır. Katkıda bulunmak için **Pull Request** gönderebilirsiniz.

🚀 **Hazırsanız, hemen kullanmaya başlayın!** 🎉

