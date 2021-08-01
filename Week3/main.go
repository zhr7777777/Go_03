package week3

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "hello\n")
	})
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	g, errCtx := errgroup.WithContext(ctx)

	serverApp := &http.Server{Addr: ":8080"}
	serverDebug := &http.Server{Addr: ":8081"}

	sc := make(chan os.Signal)                         // 声明接收系统信号的chan
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM) // 显示注册想要接收的系统信号

	g.Go(func() error {
		fmt.Printf("Server listen at %v\n", serverApp.Addr)
		err := serverApp.ListenAndServe()
		fmt.Printf("Server stopped listening with error: %v\n", err)
		return err
	})
	g.Go(func() error {
		fmt.Printf("Server listen at %v\n", serverDebug.Addr)
		err := serverDebug.ListenAndServe()
		fmt.Printf("Server stopped listening with error: %v\n", err)
		return err
	})
	g.Go(func() error {
		<-errCtx.Done()
		fmt.Println("http server stop")
		serverApp.Shutdown(errCtx)
		serverDebug.Shutdown(errCtx)
		return errCtx.Err()
	})

	g.Go(func() error {
		signal := <-sc
		fmt.Printf("Server stopped by signal: ")
		fmt.Println(signal)
		cancel()
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("Get errors: ", err)
	}

	fmt.Println("Exited")
}
