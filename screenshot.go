package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/google/uuid"
	"os"
	"time"
)

func main() {
	domain := os.Args[1]
	allocatorOptions := chromedp.DefaultExecAllocatorOptions[2:]
	opts := append(
		allocatorOptions,
		chromedp.NoFirstRun,
		chromedp.ExecPath("/usr/bin/google-chrome"),
		chromedp.NoDefaultBrowserCheck,
		chromedp.IgnoreCertErrors,
    chromedp.NoSandbox,
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("single-process", true),
		chromedp.Flag("no-zygote", true),
	)
	allocatorCtx, allocatorCancelFunc := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocatorCtx)
	defer cancel()
	defer allocatorCancelFunc()

	//navigate
	err := chromedp.Run(ctx, chromedp.Navigate(domain))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// wait for 2 seconds
	err = chromedp.Run(ctx, chromedp.Sleep(time.Second*2))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// take screenshot
	var out []interface{}
	var res []byte
	err = chromedp.Run(ctx, chromedp.CaptureScreenshot(&res))
	out = append(out, res)

	//save screenshot to file
	file := fmt.Sprintf("image-%s.png", uuid.New())
	err = os.WriteFile(file, out[0].([]byte), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
