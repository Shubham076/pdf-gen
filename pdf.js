//using puppeteer

async function createPdf(htmlContent){
    const browser = await puppeteer.launch({
        args: [
            '--headless',
            '--disable-gpu',
            '--unlimited-storage',
            '--no-sandbox',
            '--disable-setuid-sandbox',
            '--disable-dev-shm-usage',
            '--ignore-certificate-errors',
            '--disable-accelerated-2d-canvas'
        ],
    });
    const page = await browser.newPage();
    await page.setContent(htmlContent, {
        waitUntil: "networkidle0",
        timeout: 600000,
    });
    page.on('error', err => {
        logger.error('Puppeteer error.', err);
    });

    try {
        const pdf = await page.pdf({
            timeout: 600000,
            format: "A4",
            margin: {
                top: '15mm',
                bottom: '15mm',
                left: '15mm',
                right: '15mm',
            },
            displayHeaderFooter: true,
            headerTemplate: '<div></div>',
            footerTemplate: '<div></div>',
            printBackground: true,
        });
        await page.close();
        await browser.close();
        return pdf;
    } catch (error) {
        try {
            await page.close();
            await browser.close();
        } catch (err) {
            logger.error("error in closing browser for pdf report generation : ", err);
        }
        throw error;
    }
};