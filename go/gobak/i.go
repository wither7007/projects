import "github.com/cavaliercoder/grab"

resp, err := grab.Get(".", "https://example.com/image.jpg")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Downloaded", resp.Filename)
