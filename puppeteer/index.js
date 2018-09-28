const fs = require('fs');
const assert = require('assert');
const puppeteer = require('puppeteer');

(async () => {
  const browser = await puppeteer.launch();
  const page = await browser.newPage();

  // navigate to target page
  await page.goto('https://google.com');

  // save PDF
  await page.pdf({
    path: 'puppeter_sample.pdf',
    printBackground: true,
  });

  browser.close();
})();
