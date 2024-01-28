package main

import (
	"context"
	"os"
	"sync"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func CreatePdf(content, fileName string) error {
	var err error

	//chromedp.WithDebugf(log.Printf)
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	err = chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			lctx, cancel := context.WithCancel(ctx)
			chromedp.ListenTarget(lctx, func(ev interface{}) {
				if ev, ok := ev.(*page.EventLifecycleEvent); ok {
					if ev.Name == "networkIdle" {
						wg.Done()
						cancel()
					}
				}
			})
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, content).Do(ctx)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			wg.Wait()
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPaperWidth(8.28).  //A4
				WithPaperHeight(11.7). //A4
				WithMarginRight(0.55).WithMarginLeft(0.55).
				WithMarginBottom(0.55).WithMarginTop(0.55).
				WithHeaderTemplate("<div></div>").
				WithFooterTemplate("<div></div>").
				WithDisplayHeaderFooter(true).
				WithPrintBackground(true).Do(ctx)
			if err != nil {
				return err
			}
			return os.WriteFile(fileName, buf, 0o644)
		}),
	)
	return err
}
