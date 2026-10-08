# Go Send Transactional Email SMTP

This example demonstrates how to send a basic transactional email using Go's standard `net/smtp` package. It connects to an SMTP server, authenticates, and sends a simple text email, illustrating the core mechanism for programmatic email delivery. This forms the foundation for building more complex transactional email services.

## Language

`go`

## How to Run

1. Set the following environment variables: `SMTP_HOST`, `SMTP_PORT`, `SMTP_USERNAME`, `SMTP_PASSWORD`, `SENDER_EMAIL`, `RECIPIENT_EMAIL`.
2. Run `go run main.go`.

## Original Article

This example accompanies the Turkish article: [Go ile İşlemsel E-posta Servisleri: Hoş Geldin E-postaları, Uyumluluk ve Sekme Takibi](https://fatihsoysal.com/blog/go-ile-islemsel-e-posta-servisleri-hos-geldin-e-postalari-uyumluluk-ve-sekme-takibi/).

## License

MIT — see [LICENSE](LICENSE).
