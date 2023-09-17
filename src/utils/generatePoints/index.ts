import { Receipt, Item } from '../../types'

const generatePoints = (receipt: Receipt): number => {
  let points: number = 0

  // One point for every alphanumeric character in the retailer name
  points += receipt.retailer.replace(/[^a-zA-Z0-9]/g, '').length

  // 50 points if the total is a round dollar amount with no cents
  if (receipt.total.endsWith('.00')) {
    points += 50
  }

  // 25 points if the total is a multiple of 0.25
  const total = parseFloat(receipt.total.replace('$', ''))
  if (total % 0.25 === 0) {
    points += 25
  }

  // 5 points for every two items on the receipt
  points += Math.floor(receipt.items.length / 2) * 5

  // If the trimmed length of the item description is a multiple of 3,
  // multiply the price by 0.2 and round up to the nearest integer.
  // The result is the number of points earned
  receipt.items.forEach((item: Item) => {
    if (item.shortDescription.trim().length % 3 === 0) {
      const price = parseFloat(item.price.replace('$', ''))
      points += Math.ceil(price * 0.2)
    }
  })

  // 6 points if the day in the purchase data is odd
  const rawDate = receipt.purchaseDate.split('-').map(Number)
  // Date.UTC() expects the month to be 0-indexed
  const day = new Date(
    Date.UTC(rawDate[0], rawDate[1] - 1, rawDate[2])
  ).getUTCDate()
  if (day % 2 === 1) {
    points += 6
  }

  // 10 points if the time of purchase is after 2:00pm and before 4:00pm
  if (receipt.purchaseTime > '14:00' && receipt.purchaseTime < '16:00') {
    points += 10
  }

  return points
}

export default generatePoints
