import { Request, Response } from 'express'
import { v4 } from 'uuid'
import { generatePoints } from '../utils'
import { Receipt } from '../types'

const receipts: Record<string, Receipt> = {}

const processReceipt = (req: Request, res: Response): void => {
  const receipt: Receipt = req.body
  const id: string = v4()

  receipt.points = generatePoints(receipt)
  receipts[id] = receipt

  res.status(200).json({ id })
}

const getPoints = (req: Request, res: Response): void => {
  const id: string = req.params.id
  const receipt: Receipt = receipts[id]

  if (!receipt) {
    res.status(404).json({ error: 'No receipt found for that id' })
  }

  res.status(200).json({ points: receipt.points })
}

export { processReceipt, getPoints }
