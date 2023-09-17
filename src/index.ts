import express from 'express'
import bodyParser from 'body-parser'
import router from './routes'

const app: express.Application = express()
const port: number = 3000

app.use(bodyParser.json())

app.use('/receipts', router)

app.listen(port, () => {
  console.log(`Server listening at http://localhost:${port}`)
})
