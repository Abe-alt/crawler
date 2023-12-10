// import { crawlPage } from './crawl.js';
// import { printReport } from './report.js';
//
// let ipnt = document.getElementById("crawl");
// let actionBtn = document.getElementById("actionBtn");
// let result = document.getElementById("result");
//
// actionBtn.addEventListener('click', async () => {
//     if (ipnt.value !== "") {
//         console.log(`Crawler is starting at : ${ipnt.value}`)
//         const pages = await crawlPage(ipnt.value, ipnt.value, {})
//         printReport(pages)
//     } else {
//         console.log('no website provided !');
//     }
// })




const { argv } = require("process")
const { crawlPage } = require('./crawl.js')
const { printReport } = require('./report.js')

async function main(){
    if (process.argv.length < 3) {
        /////console.log('no website provided !')
        process.exit(1)
    }
    if (process.argv.length > 3){
        /////console.log('wrong number of args !')
        process.exit(1)
    }
    /////console.log(`Crawler is starting at : ${process.argv[2]}`)

    const baseUrl = process.argv[2]

    const pages = await crawlPage(baseUrl, baseUrl, {})
    // for (const page of Object.entries(pages)){
    //     console.log(page)
    // }
    console.log(JSON.stringify(pages))
    //printReport(pages)

}

main()
