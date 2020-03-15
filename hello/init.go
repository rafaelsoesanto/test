package hello

import (
	"bytes"
	"context"
	"expvar"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/tealeg/xlsx"
	"github.com/tokopedia/partnerapp/pkg/common"
	"github.com/yeka/zip"

	"github.com/opentracing/opentracing-go"
	"gopkg.in/tokopedia/logging.v1"
)

type ServerConfig struct {
	Name string
}

type Config struct {
	Server ServerConfig
}

type HelloWorldModule struct {
	cfg       *Config
	something string
	stats     *expvar.Int
}

func NewHelloWorldModule() *HelloWorldModule {

	var cfg Config

	ok := logging.ReadModuleConfig(&cfg, "config", "hello") || logging.ReadModuleConfig(&cfg, "files/etc/gosample", "hello")
	if !ok {
		// when the app is run with -e switch, this message will automatically be redirected to the log file specified
		log.Fatalln("failed to read config")
	}

	// this message only shows up if app is run with -debug option, so its great for debugging
	logging.Debug.Println("hello init called", cfg.Server.Name)

	return &HelloWorldModule{
		cfg:       &cfg,
		something: "John Doe",
		stats:     expvar.NewInt("rpsStats"),
	}

}

func (hlm *HelloWorldModule) SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), r.URL.Path)
	defer span.Finish()

	hlm.stats.Add(1)
	hlm.someSlowFuncWeWantToTrace(ctx, w)
}

func (hlm *HelloWorldModule) someSlowFuncWeWantToTrace(ctx context.Context, w http.ResponseWriter) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "someSlowFuncWeWantToTrace")
	defer span.Finish()

	w.Write([]byte("Hello " + hlm.something))
}

func (hlm *HelloWorldModule) TestFunc(w http.ResponseWriter, r *http.Request) {
	testZip()
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Master Supplier")
	if err != nil {
		log.Println(err)
		return
	}

	row := sheet.AddRow()
	row.AddCell().SetValue("test")

	row = sheet.AddRow()
	row.AddCell().SetValue("asd")

	buf := new(bytes.Buffer)
	file.Write(buf)

	filename := "uhuy"
	buf = EncryptFileToZip(buf, filename, 1)

	if err = file.Save("/home/nakama/Rafael/" + filename + ".zip"); err != nil {
		log.Println(err)
		return
	}

	// // write a zip password
	// raw, _ := io.Create(`./test.zip`)
	// os.Open
	// io.write()
	// zipw := zip.NewWriter(raw)
	// waw, err := zipw.Encrypt("hello-adele.txt", "golang")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//_, err = io.Copy(waw, bytes.NewReader(contents))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// zipw.Close()

	// // write a password zip
	// fzip, err := os.Create(`./test.zip`)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// zipw := zip.NewWriter(fzip)
	// waw, err := zipw.Encrypt("hello.txt", "golang")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = io.Copy(waw, bytes.NewReader(contents))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// zipw.Close()

	// read the password zip
	// zipr, err := zip.NewReader(bytes.NewReader(raw.Bytes()), int64(raw.Len()))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, z := range zipr.File {
	// 	z.SetPassword("golang")
	// 	rr, err := z.Open()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	_, err = io.Copy(os.Stdout, rr)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	rr.Close()
	// }

}

// EncryptFileToZip is Encrypt a single file to a protected zip file
func EncryptFileToZip(buffer *bytes.Buffer, fileName string, encryptionMethod int) *bytes.Buffer {
	raw := new(bytes.Buffer)

	zipw := zip.NewWriter(raw)
	defer zipw.Close()

	//password := GenerateRandomPassword([]common.CharacterSet{common.LowerCaseCharacterSet, common.UpperCaseCharacterSet, common.NumberCharacterSet}, common.CharacterSetMap[common.SimilarCharacterSet], 7)

	w, err := zipw.Encrypt(fileName+common.XlsxExtension, "golang", zip.StandardEncryption)
	if err != nil {
		return nil
	}

	_, err = w.Write(buffer.Bytes())
	if err != nil {
		return nil
	}
	return raw
}

func testZip() {

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Master Supplier")
	if err != nil {
		log.Println(err)
		return
	}

	row := sheet.AddRow()
	row.AddCell().SetValue("test")

	row = sheet.AddRow()
	row.AddCell().SetValue("asd")

	buf := new(bytes.Buffer)
	file.Write(buf)

	fzip, err := os.Create(`./test.zip`)
	if err != nil {
		log.Fatalln(err)
	}
	zipw := zip.NewWriter(fzip)

	defer zipw.Close()
	w, err := zipw.Encrypt("aw"+common.XlsxExtension, `golang`, zip.StandardEncryption)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(w, bytes.NewReader(buf.Bytes()))
	if err != nil {
		log.Fatal(err)
	}
	zipw.Flush()
}
